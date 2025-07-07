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

	adminRouter := http.Group("/admin", middleware.AuthenticationMiddleware)

	portfolio := adminRouter.Group("/portfolio")
	portfolio.GET("", app.Portfolio.Controller.GetAll)
	portfolio.PUT("", middleware.ValidationMiddleware(validation.Portfolio{}), app.Portfolio.Controller.Create)
	portfolio.GET("/:id", app.Portfolio.Controller.GetById)
	portfolio.PATCH("/:id", middleware.ValidationMiddleware(validation.PortfolioUpdate{}), app.Portfolio.Controller.Update)
	portfolio.DELETE("/:id", app.Portfolio.Controller.Delete)

	projects := adminRouter.Group("/projects")
	projects.GET("", app.Projects.Controller.GetAll)
	projects.PUT("", middleware.ValidationMiddleware(validation.Project{}), app.Projects.Controller.Create)
	projects.GET("/:id", app.Projects.Controller.GetById)
	projects.PATCH("/:id", middleware.ValidationMiddleware(validation.ProjectsUpdate{}), app.Projects.Controller.Update)
	projects.DELETE("/:id", app.Projects.Controller.Delete)

	authRouter := http.Group("/auth")
	authRouter.POST("/login", middleware.ValidationMiddleware(validation.Login{}), app.Auth.Controller.Login)
	authRouter.POST("/register", middleware.ValidationMiddleware(validation.Register{}), app.Auth.Controller.Register)

	http.GET("/portfolio", app.Portfolio.Controller.GetPortfolio)
	http.GET("/projects", app.Projects.Controller.GetAll)
}
