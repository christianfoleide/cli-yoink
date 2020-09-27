package yoink

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	acceptJSON      = "application/json"
	contentTypeJSON = "application/json"
	r               io.Reader
	rCloser         io.ReadCloser
)

type Client struct {
	Method      string
	ResourceURI string
	Payload     []byte
	HTTPClient  *http.Client
}

//DefaultClient returns a default client for handling a default GET request
func DefaultClient(resourceURI string) *Client {
	return &Client{
		Method:      "GET",
		ResourceURI: resourceURI,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

//NewClient ...
func NewClient(method, resourceURI string, data []byte) *Client {

	return &Client{
		Method:      method,
		ResourceURI: resourceURI,
		Payload:     data,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) DoRequest() ([]byte, error) {

	req, err := http.NewRequest(c.Method, c.ResourceURI, nil)

	if err != nil {
		return nil, err
	}

	if c.Payload != nil {
		r = bytes.NewReader(c.Payload)
		rCloser = ioutil.NopCloser(r)
		req.Body = rCloser
	}
	req.Header.Add("Accept", acceptJSON)
	req.Header.Add("Content-Type", contentTypeJSON)

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
