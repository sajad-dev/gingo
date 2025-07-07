package bootsterap

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/middleware"
	"github.com/sajad-dev/gingo/internal/routes"
	"gorm.io/gorm"
)

func Boot(httpServer *gin.Engine,db *gorm.DB) {
	validation.ValidationBoot()
	middleware.BootMiddleware(httpServer)
	routes.BootRoute(httpServer, db)
}
