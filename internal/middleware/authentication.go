package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/helpers"
	"github.com/sajad-dev/gingo/internal/exception"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		exception.AuthorizationException(ctx)
		ctx.Abort()
		return
	}

	_, valid, err := helpers.ValidJWT(authHeader)

	if err != nil {
		exception.ServerException(ctx, err)
		ctx.Abort()
		return
	}

	if !valid {
		exception.AuthorizationException(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
