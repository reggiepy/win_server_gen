package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"win_server_gen/pkg/util/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of win_server_gen",
	Long:  `All software has versions. This is win_server_gen's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("win_server_gen Static Site Generator v%s -- HEAD", version.Full())
	},
}
