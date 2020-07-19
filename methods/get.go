package methods

import (
	"io/ioutil"
	"net/http"
)

func Get(dest string) ([]byte, error) {

	resp, err := http.Get(dest)
	if err != nil {
		return nil, err
	}

	body := resp.Body
	defer body.Close()

	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
