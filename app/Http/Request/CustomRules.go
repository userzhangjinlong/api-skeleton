package Request

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

//RegPhone 手机号码正则验证
func regPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	regular := "^1[3456789]\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
