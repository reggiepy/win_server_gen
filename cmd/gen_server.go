package cmd

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"win_server_gen/handler"
)

var (
	server2     = &handler.Service{}
	fileOption2 = &handler.FileOption{}
)

func init() {
	typeHandler = map[string]func(){
		"bat":    genBat,
		"server": genServer,
	}
}

func init() {
	genServerCmd.Flags().StringVarP(&server2.Id, "id", "i", "", "Id")
	genServerCmd.Flags().StringVarP(&server2.Name, "name", "n", "", "Name")
	_ = genServerCmd.MarkFlagRequired("name")
	genServerCmd.Flags().StringVarP(&server2.Description, "description", "", "", "Description")
	genServerCmd.Flags().StringVarP(&server2.LogPath, "logpath", "", "", "LogPath")
	genServerCmd.Flags().StringVarP(&server2.LogMode, "logmode", "", "roll", "LogMode")
	genServerCmd.Flags().StringVarP(&server2.Depends, "depends", "", "", "Depends")
	genServerCmd.Flags().StringVarP(&server2.WorkingDirectory, "workingdirectory", "", "", "WorkingDirectory")
	genServerCmd.Flags().StringVarP(&server2.Executable, "executable", "e", "", "Executable")
	_ = genServerCmd.MarkFlagRequired("executable")
	genServerCmd.Flags().BoolVarP(&fileOption.OverWrite, "overwrite", "o", true, "over write file (default: false)")
	genServerCmd.Flags().StringVarP(&t, "type", "t", "all", "gen server type (all | bat | server ) (default: all)")
	rootCmd.AddCommand(genServerCmd)
}

var genServerCmd = &cobra.Command{
	Use:   "genServer",
	Short: "Generator server",
	Long:  `Generator server`,
	Run: func(cmd *cobra.Command, args []string) {
		if server2.Id == "" {
			server2.Id = server2.Name
		}
		if server2.Description == "" {
			server2.Description = server2.Name
		}
		dist := fmt.Sprintf("%s_server.xml", server2.Name)
		destFile := filepath.Join(handler.DistDir, dist)
		_, err := os.Lstat(destFile)
		if !os.IsNotExist(err) {
			_ = os.Remove(destFile)
		}
		f, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			return
		}
		defer f.Close()
		encoder := xml.NewEncoder(f)
		encoder.Indent("", "    ")
		err = encoder.Encode(server2)
		if err != nil {
			fmt.Printf("error unmarsh xml: %v", err)
			return
		}
		handler.VerboseLog(fmt.Sprintf("BaseDir: %v", handler.BaseDir))
		handler.VerboseLog(fmt.Sprintf("DistDir: %v", handler.DistDir))
		exePath := filepath.Join(handler.DistDir, fmt.Sprintf("%s_server.exe", server2.Name))
		sourcePath := filepath.Join(filepath.Join(handler.BaseDir, "server2"), "WinSW-x86.exe")
		err = handler.CopyFile(sourcePath, exePath, *fileOption2, handler.FileExistIgnoreHandler)
		if err != nil {
			fmt.Printf("Error copying service %v", err)
			return
		}
	},
}
