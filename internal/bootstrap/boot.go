package bootstrap

import (
	"github.com/gin-gonic/gin"
	boot "github.com/sajad-dev/gingo-helpers/core/bootstrap"
	"github.com/sajad-dev/gingo-helpers/types"
	"github.com/sajad-dev/gingo/internal/app/validation"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/sajad-dev/gingo/internal/middleware"
	"github.com/sajad-dev/gingo/internal/routes"
	"gorm.io/gorm"
)

func UtilsBoot() {
	boot.Boot(types.Bootsterap{
		Config: types.ConfigUtils{
			STORAGE_PATH: config.Config.STORAGE_PATH,
			DATABASE:     table.TablesVerfiy,
			JWT:          config.Config.JWT,
		},
	})
}

func Boot(httpServer *gin.Engine, db *gorm.DB) {
	UtilsBoot()
	validation.ValidationBoot()
	middleware.BootMiddleware(httpServer)
	routes.BootRoute(httpServer, db)
}
