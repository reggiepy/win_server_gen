package cmd

import (
	"github.com/spf13/cobra"
	"win_server_gen/handler"
)

var (
	server     = &handler.Server{}
	fileOption = handler.FileOption{}

	err error
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Generator windows server scripts",
	Long:  `Generator windows server scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
