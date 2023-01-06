package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"win_server_gen/handler"
)

var (
	t           string
	typeHandler = map[string]func(){}
	server      = &handler.Server{}
	fileOption  = handler.FileOption{}
)

func init() {
	typeHandler = map[string]func(){
		"bat":    genBat,
		"server": genServer,
	}
}

func init() {
	genCmd.Flags().StringVarP(&server.Name, "name", "n", "", "Name of the server")
	_ = genCmd.MarkFlagRequired("name")
	genCmd.Flags().StringVarP(&server.Executable, "executable", "e", "", "server executable")
	_ = genCmd.MarkFlagRequired("executable")
	genCmd.Flags().StringVarP(&server.StopExecutable, "stopexecutable", "s", "", "server stopexecutable")
	genCmd.Flags().StringVarP(&server.KillKeyWorld, "killkeyworld", "k", "", "server stopexecutable")
	_ = genCmd.MarkFlagRequired("killkeyworld")
	genCmd.Flags().BoolVarP(&fileOption.OverWrite, "overwrite", "o", true, "over write file (default: false)")
	genCmd.Flags().StringVarP(&t, "type", "t", "all", "gen server type (all | bat | server ) (default: all)")
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generator",
	Long:  `Generator`,
	Run: func(cmd *cobra.Command, args []string) {
		if t == "all" {
			for _, v := range typeHandler {
				v()
			}
			return
		}

		handleFunc := typeHandler[t]
		if handleFunc == nil {
			fmt.Printf("[ un support type %s ]\n", t)
		} else {
			handleFunc()
		}
	},
}

func genServer() {
	serverHandle := handler.NewServerHandler(server)
	serverHandle.FileOption = fileOption
	err := serverHandle.GenAll()
	if err != nil {
		fmt.Printf("Error generating server: %v\n", err)
	}
	return
}

func genBat() {
	batHandler := handler.NewBatchHandler(server)
	batHandler.FileOption = fileOption
	err := batHandler.GenAll()
	if err != nil {
		fmt.Printf("Error generating script: %v\n", err)
	}
	return
}
