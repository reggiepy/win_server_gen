package handler

import (
	"os"
)

var (
	BaseDir string
	DistDir string
	err     error
)

func init() {
	DistDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	BaseDir, err = os.Executable()
	if err != nil {
		panic(err)
	}
}

func SetBaseDir(dir string) {
	BaseDir = dir
}
func SetDistDir(dir string) {
	DistDir = dir
}
