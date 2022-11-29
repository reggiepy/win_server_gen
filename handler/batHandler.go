package handler

import (
	"fmt"
	"path/filepath"
)

type BatHandler struct {
	Server     *Server
	FileOption FileOption
}

func NewBatchHandler(server *Server) *BatHandler {
	return &BatHandler{
		Server:     server,
		FileOption: FileOption{},
	}
}

func (h *BatHandler) GenStartBatString() string {
	stopCommand := fmt.Sprintf("call stop%s.bat", h.Server.Name)
	startCommand := fmt.Sprintf("call start%s.VBS", h.Server.Name)
	return fmt.Sprintf("%s\n%s", stopCommand, startCommand)
}

func (h *BatHandler) GenStopBatString() string {
	stopCommand := h.Server.StopExecutable
	if stopCommand == "" {
		stopCommand = fmt.Sprintf("taskkill /im %s /t /f", h.Server.KillKeyWorld)
	}
	return stopCommand
}
func (h *BatHandler) GenStartVbsString() string {
	return fmt.Sprintf("CreateObject(\"WScript.Shell\").Run \"%s\",0", h.Server.Executable)
}

func (h *BatHandler) GenStartBatFile() error {
	dest := filepath.Join(DistDir, fmt.Sprintf("start%s.bat", h.Server.Name))
	return WriteFile(h.GenStartBatString(), dest, h.FileOption, FileExistIgnoreHandler)
}

func (h *BatHandler) GenStopBatFile() error {
	dest := filepath.Join(DistDir, fmt.Sprintf("stop%s.bat", h.Server.Name))
	return WriteFile(h.GenStopBatString(), dest, h.FileOption, FileExistIgnoreHandler)
}
func (h *BatHandler) GenStartVbsFile() error {
	dest := filepath.Join(DistDir, fmt.Sprintf("start%s.VBS", h.Server.Name))
	return WriteFile(h.GenStartVbsString(), dest, h.FileOption, FileExistIgnoreHandler)
}

func (h *BatHandler) GenAll() (err error) {
	err = h.GenStartVbsFile()
	if err != nil {
		return err
	}
	err = h.GenStartBatFile()
	if err != nil {
		return err
	}
	err = h.GenStopBatFile()
	if err != nil {
		return err
	}
	return
}
