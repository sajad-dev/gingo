package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/response"
)

func AuthorizationException(ctx *gin.Context) {
	response.Status401Res(ctx,gin.H{
		"error": "Unauthorized  :( ) ",
	})
}
