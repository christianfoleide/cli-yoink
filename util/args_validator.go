package util

import (
	"errors"
)

func ValidateArgs(args []string) error { //used if method flag is present

	if len(args) != 3 {
		return errors.New("Expected 3 arguments, method, url, and a jsonfile")
	}

	return ValidateUrl(args[1])
}
