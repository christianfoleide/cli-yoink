package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error){
	if err := validateExtension(filename); err != nil {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)
	return bytes, nil

}


func validateExtension(basename string) error {

	ext := filepath.Ext(basename)
	if ext != ".json" {
		return errors.New("File must have a .json extension")
	}
	return nil
}
