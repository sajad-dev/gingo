package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/app/bootstrap"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/middleware"
	"gorm.io/gorm"
)

type Route struct {
	Path       string
	Method     string
	Controller []gin.HandlerFunc
}

func BootRoute(http *gin.Engine, db *gorm.DB) {
	app := bootstrap.Counstructor(db)

	authRouter := http.Group("/auth")
	authRouter.POST("/login", middleware.ValidationMiddleware(validation.Login{}), app.Auth.Controller.Login)
	authRouter.POST("/register", middleware.ValidationMiddleware(validation.Register{}), app.Auth.Controller.Register)


}
