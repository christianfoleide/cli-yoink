package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ACCEPT_JSON       = "application/json"
	CONTENT_TYPE_JSON = "application/json"
	r                 io.Reader
)

type client struct {
	method      string
	resourceURI string
	data        []byte
	HTTPClient  *http.Client
}

func NewClient(method, resourceURI string) *client {

	return &client{
		method:      method,
		resourceURI: resourceURI,
		data:        nil,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *client) WithData(data []byte) *client {
	c.data = data
	return c
}

func (c *client) DoRequest() ([]byte, error) {

	if c.data != nil {
		r = bytes.NewReader(c.data)
	}

	req, err := http.NewRequest(c.method, c.resourceURI, r)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", ACCEPT_JSON)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	resp, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil

}
