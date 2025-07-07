package test

import (
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/middleware"
	"github.com/sajad-dev/gingo/utils"
	"github.com/stretchr/testify/assert"
)

type InputTest2 struct {
	Name string
	Amount string
	Gg string
}
type InputTest3 struct {
	Name     string
	FileInp3 string `file:"yes"`
}

type InputTest struct {
	Inp1 []InputTest2
	File string `file:"yes"`
}


type InputTest3F struct {
	Name     string
	FileInp3 []*multipart.FileHeader
}
type InputTestF struct {
	Inp1 []InputTest2
	File []*multipart.FileHeader `validate_file:"size=10"`
}
type InputTestG struct {
	NameFFF string
	A       []string
	B       InputTest2
}

func TestMiddelware_validation(t *testing.T) {
	validation.ValidationBoot()

	r := utils.CreateServer("/test", "post", []gin.HandlerFunc{middleware.ValidationMiddleware(InputTestF{})}, 8080)
	req := InputTest{Inp1: []InputTest2{{Name: "ff",Amount: "sd",Gg: "sdfkjkf"}}, File: "../../../storage/files/image.png"}
	w, err := utils.SendRequest(utils.Request{
		Method:  http.MethodPost,
		Path:    "/test",
		Inputs:  req,
		Headers: map[string]string{"Content-Type": "multipart/form-data"},
		Engin:   r,
	})
	assert.NoError(t, err)
	assert.Equal(t, 200, w.Result().StatusCode)
}
