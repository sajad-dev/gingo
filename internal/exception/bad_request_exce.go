package exception

import (

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/gingo/internal/app/response"
)

func BadReqException(ctx *gin.Context, err error) bool {
	if err != nil {
		errValidation, ok := (err).(validator.ValidationErrors)
		if ok {
			errList := [][2]string{}
			for _, val := range errValidation {
				errList = append(errList, [2]string{val.Field(), val.Error()})
			}
			response.Status400Res(ctx, gin.H{
				"errors": errList,
			})
			return true
		}
		ServerException(ctx,err)
		return true
	}
	return false
}
