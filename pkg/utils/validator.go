package utils

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var ValidateStruct = validator.New().Struct

func GetValidationStructErrors(err error) string {
	if err == nil {
		return ""
	}

	var errorMessages []string

	// check error type is validator.ValidationErrors
	_, ok := err.(validator.ValidationErrors)
	if !ok {
		// if not, return a generic error message
		return "Validation error: " + err.Error()
	}

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})

	for _, err := range err.(validator.ValidationErrors) {
		msg := err.Error()
		errorMessages = append(errorMessages, msg)
	}

	return strings.Join(errorMessages, " , ")
}
