package util

import (
	"encoding/json"
	"io/ioutil"
)

// ParseFile parse json file
func ParseFile(fileName string, out interface{}) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, out); err != nil {
		return err
	}

	return nil
}

// ParseJSONFile is the entrypoint of parse json file
func ParseJSONFile(fileName string, out interface{}) error {
	if err := ParseFile(fileName, out); err != nil {
		return err
	}

	if entrys, ok := out.(IJsonEntrys); ok {
		entrys.InitMap()
	}

	return nil
}

// IJsonEntrys is interface
type IJsonEntrys interface {
	InitMap()
}
