package response

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetErrorMsg 错误信息
func GetErrorMsg(errors validator.ValidationErrors, formRequest interface{}) string {
	msg := ""
	label := ""
	for _, err := range errors {
		formRequest := reflect.TypeOf(formRequest)
		if formType, ok := formRequest.FieldByName(err.StructField()); ok {
			msg = formType.Tag.Get(err.Tag())
			label = formType.Tag.Get("label")
		}
		if msg == "" {
			if label == "" {
				label = err.StructField()
			}
			msg = label + "不正确"
		}
		break
	}
	return msg
}
