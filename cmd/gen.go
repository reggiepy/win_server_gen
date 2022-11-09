package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"win_server_gen/handler"
)

func init() {
	genCmd.Flags().StringVarP(&server.Name, "name", "n", "", "Name of the server")
	_ = genCmd.MarkFlagRequired("name")
	genCmd.Flags().StringVar(&server.Executable, "executable", "", "server executable")
	_ = genCmd.MarkFlagRequired("executable")
	genCmd.Flags().StringVar(&server.StopExecutable, "stopexecutable", "", "server stopexecutable")
	genCmd.Flags().BoolVar(&fileOption.OverWrite, "overwrite", false, "over write file (default: false)")
	serverCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generator",
	Long:  `Generator`,
	Run: func(cmd *cobra.Command, args []string) {
		serverHandle := handler.NewServerHandler(server)
		serverHandle.FileOption = fileOption
		err = serverHandle.GenAll()
		if err != nil {
			fmt.Printf("Error generating server: %v\n", err)
		}

		batHandler := handler.NewBatchHandler(server)
		batHandler.FileOption = fileOption
		err = batHandler.GenAll()
		if err != nil {
			fmt.Printf("Error generating script: %v\n", err)
		}
	},
}
