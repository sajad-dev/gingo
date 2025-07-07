package exception

import (
	// "errors"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/response"
	// "gorm.io/gorm"
)

func ServerException(ctx *gin.Context, err error) bool {
	if err != nil {
		if NotFoundException(ctx, err) {
			return true
		}
		if DEBUG == "false" {
			response.Status500Res(ctx, gin.H{
				"error": "Server has error",
			})
		} else {
			response.Status500Res(ctx, gin.H{
				"error": Exception(err),
			})
		}
		return true
	}
	return false
}
