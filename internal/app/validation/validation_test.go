package validation

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/gingo/internal/exception"
	"github.com/sajad-dev/gingo-helpers/utils"
	"github.com/stretchr/testify/assert"
)

type Valid struct {
	Name string `validate:"required,max=10"`
}

func handler(ctx *gin.Context) {
	ValidReq = validator.New()
	err := ValidationReq(ctx, &Valid{})
	if ok := exception.BadReqException(ctx, err); ok {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "all good"})
}
func TestValidationReq(t *testing.T) {
	r := utils.CreateServer("/valid", "post", []gin.HandlerFunc{handler}, 8080)
	w, _ := utils.SendRequest(utils.Request{Path: "/valid", Method: http.MethodPost, Inputs: Valid{}, Engin: r})

	ok, err := utils.CheckValidationErr(w.Body.String(), "Name", "required")
	assert.NoError(t, err)
	assert.True(t, ok)
}
