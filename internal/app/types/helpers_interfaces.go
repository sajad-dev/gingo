package types

import "github.com/gin-gonic/gin"

type CURDController interface {
	GetAll(*gin.Context)
	GetById(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

