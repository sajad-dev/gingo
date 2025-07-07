package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status200Res(ctx *gin.Context, res gin.H) {
	res["status"] = "ok"
	ctx.JSON(http.StatusOK, res)
}

func Status400Res(ctx *gin.Context, res gin.H) {
	res["status"] = "field_not_valid"

	ctx.JSON(http.StatusBadRequest, res)

}

func Status500Res(ctx *gin.Context, res gin.H) {
	res["status"] = "server_error"
	ctx.JSON(http.StatusInternalServerError, res)

}

func Status404Res(ctx *gin.Context, res gin.H) {
	res["status"] = "not_found"
	ctx.JSON(http.StatusNotFound, res)

}

func Status401Res(ctx *gin.Context, res gin.H)  {
	res["status"] = "unauthorized"
	ctx.JSON(http.StatusUnauthorized, res)

}
