package validator

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// validTime 有效时间
func validTime(fl validator.FieldLevel) bool {
	if fl.Field().Int() > 0 && fl.Field().Int() < time.Now().Unix() {
		return false
	}
	return true
}
