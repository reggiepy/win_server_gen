package handler

import (
	"os"
	"path/filepath"
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
	absBaseDir, _ := filepath.Abs(BaseDir)
	BaseDir = filepath.Dir(absBaseDir)
}

func SetBaseDir(dir string) {
	BaseDir = dir
}
func SetDistDir(dir string) {
	DistDir = dir
}
