package yoink

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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

func (f *Filehandler) ConfigToClient() (*Client, error) {

	var config map[string]interface{}

	b, err := ReadFile(f.filename)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	payload, err := ReadFile(config["payload"].(string))
	if err != nil {
		return nil, err
	}
	
	return &Client{
		Method:      config["method"].(string),
		ResourceURI: fmt.Sprintf("http://%s:%s%s", config["hostname"], config["port"], config["path"]),
		Payload:     payload,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (f *Filehandler) ListConfig() error {

	b, err := ReadFile(f.filename)
	if err != nil {
		return err
	}

	pretty := pretty.Pretty(b)
	fmt.Println(string(pretty))
	return nil
}

func (f *Filehandler) WriteChanges(changes map[string]interface{}) error {

	var currentConfig map[string]interface{}

	b, err := ReadFile(f.filename)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &currentConfig); err != nil {
		return err
	}

	for key, newValue := range changes {
		currentConfig[key] = newValue
	}

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

func ReadFile(filename string) ([]byte, error) {

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
