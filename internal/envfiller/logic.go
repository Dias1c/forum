package envfiller

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

//
func getStructFieldsTags(strct interface{}) map[string]string {
	val := reflect.ValueOf(strct).Elem()
	result := make(map[string]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		tag := string(val.Type().Field(i).Tag)
		if tag == "" {
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
		tag := string(val.Type().Field(i).Tag)
		if tag == "" {
			continue
		} else if _, ok := result[tag]; ok {
			panic(fmt.Sprintf("type %q-s has two equal tag %q", val.Type().Name(), tag))
		}
		switch val.Field(i).Kind() {
		case reflect.Int:
			n, err := strconv.Atoi(m[tag])
			if err != nil {
				return fmt.Errorf("invalid value %q of field %q", m[tag], tag)
			}
			val.Field(i).SetInt(int64(n))
		case reflect.String:
			val.Field(i).SetString(m[tag])
		default:
			panic(fmt.Sprintf("field %q(%q) type of struct %q not supported add it here", val.Type().Field(i).Name, val.Field(i).Kind(), val.Type()))
		}
	}
	return nil
}
