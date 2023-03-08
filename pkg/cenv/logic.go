package cenv

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

const tagKey = "cenv"

//
func getStructFieldsTags(strct interface{}) map[string]string {
	val := reflect.ValueOf(strct).Elem()
	result := make(map[string]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		tag, ok := val.Type().Field(i).Tag.Lookup(tagKey)
		if !ok || tag == "" {
			continue
		} else if _, ok := result[tag]; ok {
			panic(fmt.Sprintf("type %q-s has two equal tag %q", val.Type().Name(), tag))
		}
		result[tag] = ""
	}
	return result
}

//
func fillMapByFile(r io.Reader, m map[string]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		keyValue := strings.SplitN(line, "=", 2)
		if len(keyValue) != 2 || len(keyValue) < 2 {
			return fmt.Errorf("can't split %q, here is no separator '='", line)
		}
		m[keyValue[0]] = keyValue[1]
	}
	return nil
}

func fillElementFieldsByMap(strct interface{}, m map[string]string) error {
	val := reflect.ValueOf(strct).Elem()
	result := make(map[string]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		tag, ok := val.Type().Field(i).Tag.Lookup(tagKey)
		if !ok || tag == "" {
			continue
		} else if _, ok := result[tag]; ok {
			panic(fmt.Sprintf("type %q-s has two equal tag %q", val.Type().Name(), tag))
		}
		ftype := val.Field(i).Kind()
		switch ftype {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.Atoi(m[tag])
			if err != nil {
				return fmt.Errorf("invalid value %q of field %q", m[tag], tag)
			}
			val.Field(i).SetInt(int64(n))
		case reflect.Float32, reflect.Float64:
			var bitsize int = 32
			if ftype == reflect.Float64 {
				bitsize = 64
			}
			n, err := strconv.ParseFloat(m[tag], bitsize)
			if err != nil {
				return fmt.Errorf("%q(n) = strconv.ParseFloat(%q, %v) returns err: %w ", tag, m[tag], bitsize, err)
			}
			val.Field(i).SetFloat(n)
		case reflect.String:
			val.Field(i).SetString(m[tag])
		case reflect.Bool:
			value, err := strconv.ParseBool(m[tag])
			if err != nil {
				return fmt.Errorf("%q(variable) = trconv.ParseBool(%q) returns err: %w ", tag, m[tag], err)
			}
			val.Field(i).SetBool(value)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var bitsize int = 0
			switch ftype {
			case reflect.Uint:
				bitsize = 32
			case reflect.Uint8:
				bitsize = 8
			case reflect.Uint16:
				bitsize = 16
			case reflect.Uint32:
				bitsize = 32
			case reflect.Uint64:
				bitsize = 64
			}
			n, err := strconv.ParseUint(m[tag], 10, bitsize)
			if err != nil {
				return fmt.Errorf("%q(n) = strconv.ParseUint(%q, 10, %v) returns err: %w ", tag, m[tag], bitsize, err)
			}
			val.Field(i).SetUint(n)
		default:
			panic(fmt.Sprintf("field %q(%q) type of struct %q not supported add it here", val.Type().Field(i).Name, val.Field(i).Kind(), val.Type()))
		}
	}
	return nil
}
