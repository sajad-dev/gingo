package cli

import (
	"os"

	"github.com/fatih/color"
	"github.com/sajad-dev/gingo/internal/db/connection"
	"github.com/sajad-dev/gingo/internal/db/table"
	"github.com/sajad-dev/gingo/internal/server"
	"github.com/spf13/cobra"
)

var CommandsList []*Commands = []*Commands{
	{
		Command: &cobra.Command{
			Use:   "migration",
			Short: "Run migration list",
			Run: func(cmd *cobra.Command, args []string) {
				if err := table.PublicMigration(connection.DB); err != nil {
					os.Exit(0)
				}
				color.Green("Sucssfuly migrate :) - ;)")
			},
		},
		Flags: []any{},
	},
	{
		Command: &cobra.Command{
			Use:   "db-conn",
			Short: "Check database connection",
			Run: func(cmd *cobra.Command, args []string) {
				err := connection.Connect()
				if err != nil {
					os.Exit(0)
				}
				color.Green("Sucssfuly connection :) - ;)")

			},
		},
		Flags: []any{},
	},
	{
		Command: &cobra.Command{
			Use:   "server",
			Short: "Run http server ",
			Run: func(cmd *cobra.Command, args []string) {
				port, _ := cmd.Flags().GetInt("port")
				_, err := server.Http(port, connection.DB)
				if err != nil {
					os.Exit(0)
				}

			},
		},
		Flags: []any{FlagInt{Flag: "port", ShortFlag: "p", Defualt: 8080, Discription: "Run server in your port"}},
	},
}

type FlagString struct {
	Value       string
	Flag        string
	ShortFlag   string
	Defualt     string
	Discription string
}
type FlagInt struct {
	Value       int
	Flag        string
	ShortFlag   string
	Defualt     int
	Discription string
}
type Commands struct {
	Command *cobra.Command
	Flags   []any
}
