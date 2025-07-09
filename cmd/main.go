package main

import (
	gearConfig "github.com/sajad-dev/gear/config"
	"github.com/sajad-dev/gear/startup"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/sajad-dev/gingo/internal/server"
)

func main() {
	config.BootConfig(".env")
	connection.Connect()

	startup.Boot(gearConfig.ConfigSt{
		Db:          connection.DB,
		Tables:      table.TablesVerfiy,
		Http:        server.Http,
		APP_NAME:    config.Config.APP_NAME,
		DESCRIPTION: config.Config.DESCRIPTION,
	})

}
