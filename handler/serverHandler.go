package handler

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GenServerFile(server *Server, dist string) error {
	destFile := filepath.Join(DistDir, dist)
	_, err := os.Lstat(destFile)
	if !os.IsNotExist(err) {
		os.Remove(destFile)
	}
	f, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()
	encoder := xml.NewEncoder(f)
	encoder.Indent("", "    ")
	err = encoder.Encode(server)
	if err != nil {
		return fmt.Errorf("error unmarsh xml: %v", err)
	}
	return nil
}

func CopyServerDll(DistDir string, option FileOption) error {
	files, err := ioutil.ReadDir(filepath.Join(BaseDir, "server"))
	if err != nil {
		fmt.Printf("error: %v", err)
		return fmt.Errorf("error server directory: %s %v", DistDir, err)
	}
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), "dll") {
			VerboseLog(fmt.Sprintf("start copying %d %s ...\n", i+1, file.Name()))
			srcFile := filepath.Join(filepath.Join(BaseDir, "server"), file.Name())
			destFile := filepath.Join(DistDir, file.Name())
			err := CopyFile(srcFile, destFile, option, FileExistIgnoreHandler)
			if err != nil {
				return fmt.Errorf("error copying %s %v", file, err)
			}
		}
	}
	return nil
}

type ServerHandler struct {
	Server     *Server
	FileOption FileOption
}

func NewServerHandler(server *Server) *ServerHandler {
	return &ServerHandler{
		Server:     server,
		FileOption: FileOption{},
	}
}

func (h *ServerHandler) GenServerXml() error {
	xmlServerName := fmt.Sprintf("%s-server.xml", h.Server.Name)
	err := GenServerFile(h.Server, xmlServerName)
	if err != nil {
		return err
	}
	return nil
}

func (h *ServerHandler) GenServerExe() error {
	VerboseLog(fmt.Sprintf("BaseDir: %v", BaseDir))
	VerboseLog(fmt.Sprintf("DistDir: %v", DistDir))
	exePath := filepath.Join(DistDir, fmt.Sprintf("%s-service.exe", h.Server.Name))
	sourcePath := filepath.Join(filepath.Join(BaseDir, "server"), "service.exe")
	err := CopyFile(sourcePath, exePath, h.FileOption, FileExistIgnoreHandler)
	if err != nil {
		return err
	}
	return nil
}

func (h *ServerHandler) GenServerDll() error {
	err := CopyServerDll(DistDir, h.FileOption)
	if err != nil {
		return err
	}
	return nil
}

func (h *ServerHandler) GenAll() (err error) {
	err = h.GenServerXml()
	if err != nil {
		return err
	}
	err = h.GenServerExe()
	if err != nil {
		return err
	}
	err = h.GenServerDll()
	if err != nil {
		return err
	}
	return
}
