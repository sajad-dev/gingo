package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sajad-dev/gingo/internal/bootstrap"
	"github.com/sajad-dev/gingo/internal/exception"
	"gorm.io/gorm"
)

func Http(port int, db *gorm.DB) (*gin.Engine, error) {
	httpServer := gin.Default()
	bootstrap.Boot(httpServer, db)
	go func() {
		err := httpServer.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(exception.Exception(err))
		}
	}()

	return httpServer, nil
}
