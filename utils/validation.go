package utils

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func CheckTarget(url interface{}) error {
	return validation.Validate(url, is.URL)
}

func CheckInterval(day interface{}) error {
	return validation.Validate(day, is.Int)
}
