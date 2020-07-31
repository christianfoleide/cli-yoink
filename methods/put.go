package methods

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"github.com/christianfoleide/yoink/util"
)

func Put(dest string, filename string) ([]byte, error) {

	file, err := util.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(file)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPut, dest, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()

	respBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil

}
