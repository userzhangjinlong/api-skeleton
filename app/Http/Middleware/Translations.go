package Middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en2 "github.com/go-playground/validator/v10/translations/en"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"sync"
)

var (
	once        sync.Once
	transGlobal ut.Translator
)

//Translations 验证器全局翻译中间件新增
func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		setTrans(ctx)
		ctx.Set("trans", transGlobal)
		ctx.Next()
	}
}

func setTrans(ctx *gin.Context) ut.Translator {
	//单列模式获取设置的国籍语言单位
	//todo::可以优化处理这里是否存在local为英文是还是是中文的情况
	once.Do(func() {
		uni := ut.New(en.New(), zh.New())
		locale := ctx.GetHeader("locale")
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
			transGlobal = trans

		}
	})

	return transGlobal
}
