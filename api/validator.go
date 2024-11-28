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
		var passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
		re := regexp.MustCompile(passwordRegex)
		return re.MatchString(password)
	}
	return false
}
