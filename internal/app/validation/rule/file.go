package rule

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	cu "github.com/sajad-dev/gingo/internal/app/validation/rule/customerrors"
)

type FileValidation struct {
}

func (f FileValidation) Struct(st any) error {

	tType := reflect.TypeOf(st).Elem()
	tValue := reflect.ValueOf(st).Elem()

	for i := 0; i < tType.NumField(); i++ {

		field := tType.Field(i)
		val := tValue.Field(i)

		fileValidationTag := field.Tag.Get("validate_file")
		if len(fileValidationTag) != 0 {
			fileValidationTagSlice := strings.Split(fileValidationTag, ",")
			for _, value := range fileValidationTagSlice {
				validationCommand := strings.Split(value, "=")
				switch validationCommand[0] {
				case "size":
					n, _ := strconv.ParseInt(validationCommand[1], 10, 64)
					return f.FileSizeValidate(val.Interface(), field.Tag.Get("json"), n)
				}
			}
		}
	}
	return nil
}

func New() FileValidation {
	return FileValidation{}
}

func (f FileValidation) FileSizeValidate(fileInterface any, field string, size int64) error {
	fileSl, ok := fileInterface.([]*multipart.FileHeader)
	if !ok {
		return fmt.Errorf("Not avalible file  :(")
	}

	errorsValidator := []validator.FieldError{}
	for _, file := range fileSl {
		if file.Size > size*1024*1024 {
			errorsValidator = append(errorsValidator, cu.FieldErrorCustom{
				ValidationTag:  "file_size",
				FieldNamespace: field,
				FieldValue:     file.Filename,
			})
		}
	}

	if len(errorsValidator) == 0{
		return nil
	}
	var err error = validator.ValidationErrors(errorsValidator)
	return err

}
