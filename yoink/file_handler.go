package yoink

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/pretty"
)

type Filehandler struct {
	filename string
}

func NewFilehandler(filename string) *Filehandler {
	return &Filehandler{
		filename: filename,
	}
}

func (f *Filehandler) ListConfig() error {

	b, err := readFile(f.filename)
	if err != nil {
		return err
	}

	pretty := pretty.Pretty(b)
	fmt.Println(string(pretty))
	return nil
}

func (f *Filehandler) WriteChanges(changes map[string]interface{}) error {

	var currentConfig map[string]interface{}

	b, err := readFile(f.filename)
	if err != nil {
		return err
	}

	json.Unmarshal(b, &currentConfig)

	//keys should be hostname, port, method etc
	for key, newValue := range changes { //range over properties to be changed
		currentConfig[key] = newValue //rewrite value of given changed field
	}

	//unmarshal into bytes for writing
	b, err = json.Marshal(currentConfig)
	if err != nil {
		return err
	}

	pretty := pretty.Pretty(b)

	err = ioutil.WriteFile("config.json", pretty, 0644)
	if err != nil {
		return err
	}

	return nil

}

func readFile(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return b, nil

}
