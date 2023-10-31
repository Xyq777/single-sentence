package model

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("correctType", correctTypeValid)
		if err != nil {
			log.Panicln("验证器函数注册失败", err)
		}
	}
}
func correctTypeValid(fl validator.FieldLevel) bool {
	typeMap := map[string]struct{}{
		"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}, "g": {}, "h": {}, "i": {}, "j": {}, "k": {}, "l": {},
	}
	inputType := fl.Field().Interface().(string)
	if _, ok := typeMap[inputType]; !ok {
		return false
	}
	return true
}
