package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

//Translations 验证器全局翻译中间件新增
func Translations() gin.HandlerFunc {
	fmt.Println("456757")
	fmt.Printf("locale 是什么 %s", "456757")
	return func(ctx *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := ctx.GetHeader("locale")
		fmt.Printf("locale 是什么 %s", locale)
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh2.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en2.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh2.RegisterDefaultTranslations(v, trans)
				break

			}
			fmt.Printf("中间件的trans是什么：%s", trans)
			ctx.Set("trans", trans)
		}
		ctx.Next()
	}
}
