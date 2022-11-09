package main

import (
	"fmt"
	"os"
	"win_server_gen/cmd"
	"win_server_gen/handler"
)

func Test() {
	distDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	handler.SetDistDir("dest")
	handler.SetBaseDir(distDir)
	fmt.Println(handler.BaseDir)
	fmt.Println(handler.DistDir)
}

func main() {
	Test()
	_ = cmd.Execute()
}
