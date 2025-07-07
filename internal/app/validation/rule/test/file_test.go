package test

import (
	"mime/multipart"
	"testing"

	"github.com/sajad-dev/gingo/internal/app/validation/rule"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	File *multipart.FileHeader `validate_file:"size=2"`
}

func TestFileValidation_Struct(t *testing.T) {
	file := &multipart.FileHeader{
		Filename: "testfile.txt",
		Size:     1 * 1024 * 1024, 
	}

	fileValidation := rule.New()

	testStruct := &TestStruct{File: file}

	 fileValidation.Struct(testStruct)

}

func TestFileValidation_FileSizeValidate(t *testing.T) {
	file := []*multipart.FileHeader{{
		Filename: "testfile.txt",
		Size:     3 * 1024 * 1024, 
	}}

	fileValidation := rule.New()

	err := fileValidation.FileSizeValidate(file,"testfile", 2)

	assert.Contains(t,err.Error(), "file_size")

}


