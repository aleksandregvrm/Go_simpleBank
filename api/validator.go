package api

import (
	"regexp"

	util "example.com/banking/utils"
	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}

var ValidEmail validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if email, ok := fieldLevel.Field().Interface().(string); ok {

		var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		return re.MatchString(email)
	}
	return false
}

var ValidPassword validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if password, ok := fieldLevel.Field().Interface().(string); ok {
		if len(password) < 5 {
			return false
		}

		lowercaseRegex := regexp.MustCompile(`[a-z]`)
		if !lowercaseRegex.MatchString(password) {
			return false
		}

		digitRegex := regexp.MustCompile(`\d`)
		if !digitRegex.MatchString(password) {
			return false
		}

		// If all checks pass
		return true
	}
	return false
}
