package validation

import (
	"errors"
	"net/url"
	"path/filepath"
	"strings"
)

var (
	allowedMethods   = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	missingArguments = "missing or no arguments, use yoink -h"
)

//ValidateDefault validates the arguments passed to the root command (no flags)
func ValidateDefault(args []string) error {

	if len(args) == 0 {
		return errors.New(missingArguments)
	}

	if err := validateURL(args[0]); err != nil {
		return err
	}
	return nil
}

//ValidateNonDefault validates the arguments passed when changeMethod flag is present
func ValidateNonDefault(args []string) error {

	if len(args) == 0 {
		return errors.New(missingArguments)
	}

	if err := validateMethod(args[0]); err != nil {
		return err
	}

	if err := validateURL(args[1]); err != nil {
		return err
	}

	if err := validateExtension(args[2]); err != nil {
		return err
	}

	return nil
}

func validateExtension(filename string) error {

	if filepath.Ext(filename) == ".json" {
		return nil
	}
	return errors.New("invalid file extension for payload file")
}

func validateURL(destination string) error {

	_, err := url.ParseRequestURI(destination)
	if err != nil {
		return err
	}
	return nil
}

func validateMethod(requestMethod string) error {

	rm := strings.ToUpper(requestMethod)

	for _, method := range allowedMethods {

		if rm == method {
			return nil
		}

	}
	return errors.New("method not allowed or is invalid")
}
