package handler

import (
	"fmt"
	"path/filepath"
)

type Server struct {
	Id             string `xml:"id"`
	Name           string `xml:"name"`
	Description    string `xml:"description"`
	LogPath        string `xml:"logpath"`
	LogMode        string `xml:"logmode"`
	Depends        string `xml:"depends"`
	Executable     string `xml:"executable"`
	StopExecutable string `xml:"stopexecutable"`

	//	杀进程关键字
	KillKeyWorld string `xml:"killkeyworld"`
}

func NewServer(name string) (server *Server) {
	//id := uuid.NewString()
	//id = strings.ReplaceAll(id, "-", "")
	return &Server{
		Id:             name,
		Name:           name,
		Description:    name,
		LogMode:        "roll",
		LogPath:        filepath.Join(BaseDir, "server-logs"),
		Executable:     filepath.Join(BaseDir, fmt.Sprintf("start%s.bat", name)),
		StopExecutable: filepath.Join(BaseDir, fmt.Sprintf("stop%s.bat", name)),
	}
}

type Service struct {
	Id               string `xml:"id"`
	Name             string `xml:"name"`
	Description      string `xml:"description"`
	LogPath          string `xml:"logpath"`
	LogMode          string `xml:"logmode"`
	Depends          string `xml:"depends"`
	Executable       string `xml:"executable"`
	StopExecutable   string `xml:"stopexecutable"`
	WorkingDirectory string `xml:"workingdirectory"`
}
