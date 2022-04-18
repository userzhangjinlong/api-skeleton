package Request

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func registerValidation(tag string, fn validator.Func) {
	validateObj := validator.New()
	if err := validateObj.RegisterValidation(tag, fn); err != nil {
		logrus.Fatalf("register validator for '%s' error: %v", tag, err)
	}
}

//BindAndValid ShouldBind和验证
func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	registerCustomRules()
	var errs ValidErrors
	err := ctx.ShouldBind(v)
	if err != nil {
		v := ctx.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}

//registerCustomRules 绑定自定义规则
func registerCustomRules() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//绑定自定义规则
		v.RegisterValidation("regPhone", regPhone)
	}
}
