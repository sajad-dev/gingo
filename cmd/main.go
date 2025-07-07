package main

import (
	"github.com/sajad-dev/gingo/cmd/cli"
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/sajad-dev/gingo/internal/db/connection"
)

func main() {
	config.BootConfig(".env")
	connection.Connect()
	
	cli.Cli()
}
