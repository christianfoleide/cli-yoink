package util

import (
	"net/url"
)

func ValidateUrl(dest string) error {

	_, err := url.ParseRequestURI(dest)
	if err != nil {
		return err
	}
	return nil
}
