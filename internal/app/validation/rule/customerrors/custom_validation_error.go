package customerrors

import (
	"fmt"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type FieldErrorCustom struct {
	ValidationTag    string
	FieldNamespace   string
	FieldValue       interface{}

}

var _ validator.FieldError = &FieldErrorCustom{}

func (e FieldErrorCustom) Tag() string {
	return e.ValidationTag
}

func (e FieldErrorCustom) ActualTag() string {
	return e.ValidationTag
}

func (e FieldErrorCustom) Namespace() string {
	return e.FieldNamespace
}

func (e FieldErrorCustom) StructNamespace() string {
	return e.FieldNamespace
}

func (e FieldErrorCustom) Field() string {
	return e.FieldNamespace
}

func (e FieldErrorCustom) StructField() string {
	return e.FieldNamespace
}

func (e FieldErrorCustom) Value() interface{} {
	return e.FieldValue
}

func (e FieldErrorCustom) Param() string {
	return ""
}

func (e FieldErrorCustom) Kind() reflect.Kind {
	return 0
}

func (e FieldErrorCustom) Type() reflect.Type {
	return nil
}

func (e FieldErrorCustom) Error() string {
	return fmt.Sprintf("Error on field %s: %s in validation tag %s",e.FieldNamespace, e.FieldValue, e.ValidationTag)
}

func (e FieldErrorCustom) Translate(ut ut.Translator) string {
	return ""
}

func ConstructorError(err FieldErrorCustom) validator.ValidationErrors {
	return []validator.FieldError{err}
}

