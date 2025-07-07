package cli

import (
	"github.com/sajad-dev/gingo/internal/config"
	"github.com/spf13/cobra"
	"os"
)

func Cli() {
	rootCmd := &cobra.Command{
		Use:   config.Config.APP_NAME,
		Short: config.Config.DESCRIPTION,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}
	for _, value := range CommandsList {
		for _, val := range value.Flags {
			if flag, ok := val.(FlagString); ok {
				value.Command.Flags().StringVarP(&flag.Value, flag.Flag, flag.ShortFlag, flag.Defualt, flag.Discription)
			}
			if flag, ok := val.(FlagInt); ok {
				value.Command.Flags().IntVarP(&flag.Value, flag.Flag, flag.ShortFlag, flag.Defualt, flag.Discription)
			}
		}
		rootCmd.AddCommand(value.Command)

	}
	Root.Execute()
}
