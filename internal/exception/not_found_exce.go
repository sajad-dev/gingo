package exception

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/response"
	"gorm.io/gorm"
)

func NotFoundException(ctx *gin.Context, err error) bool {
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		response.Status404Res(ctx, gin.H{
			"error": "Not found !!!",
		})
		return true
	}
	return false
}
