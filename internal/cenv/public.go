package cenv

import "os"

func FillFieldsByEnvFile(filepath string, strct interface{}) error {
	// get field tag names
	tags := getStructFieldsTags(strct)
	// fill tags by file
	f, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	// set values by tags
	err = fillMapByFile(f, tags)
	if err != nil {
		return err
	}
	// Get tag names field
	return fillElementFieldsByMap(strct, tags)
}
