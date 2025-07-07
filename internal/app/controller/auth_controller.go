package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/response"
	"github.com/sajad-dev/gingo/internal/app/services"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/exception"
	"gorm.io/gorm"
)

type AuthController struct {
	Services services.AuthLogestic
}

type AuthHandller interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

var _ AuthHandller = &AuthController{}

func (a *AuthController) Login(ctx *gin.Context) {
	reqParams, _ := ctx.Get("req")
	input := reqParams.(validation.Login)

	account, token, err := a.Services.Login(ctx, input)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Status400Res(ctx, gin.H{
			"message": "email or password is not valid",
		})
		return
	}

	if ok := exception.ServerException(ctx, err); ok {
		return
	}
	response.Status200Res(ctx, gin.H{
		"account": account,
		"token":   token,
	})
}
func (a *AuthController) Register(ctx *gin.Context) {
	reqParams, _ := ctx.Get("req")
	input := reqParams.(validation.Register)

	token, err := a.Services.Register(ctx, input)
	if ok := exception.ServerException(ctx, err); ok {
		return
	}

	response.Status200Res(ctx, gin.H{
		"token": token,
	})

}
