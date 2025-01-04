package valid

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func Init() {
	zhTranslation := zh.New()
	uni = ut.New(zhTranslation)

	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)

	err := zhtranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("required")
	})
}

// Translator 翻译错误信息
func Translator(err error) string {
	result := ""
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if ok {
		for _, er := range errs {
			result += er.Translate(trans) + ";"
		}
	} else {
		result = "参数错误"
	}
	return result
}
