package methods

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"yo/util"
)

func Post(dest string, filename string) ([]byte, error) {

	file, err := util.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(file)
	fmt.Println(reqBody.Len())

	resp, err := http.Post(dest, "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
