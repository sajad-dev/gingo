package validation

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/gingo/internal/app/validation/rule"
)

var ValidReq *validator.Validate

func ValidationBoot() {
	ValidReq = validator.New()
	ValidReq.RegisterValidation("uniq", rule.UniqeField)
}

func ValidationReq(http *gin.Context, forms any) error {
	err := http.ShouldBindJSON(forms)
	if err != nil {
		return err
	}
	return ValidReq.Struct(forms)
}

func ValidationReqFile(forms any, field string) error {
	if reflect.TypeOf(forms).Kind() != reflect.Ptr{
		return fmt.Errorf("form not pointer")
	}

	fileValidation := rule.New()
	err := fileValidation.Struct(forms)
	return err

}
