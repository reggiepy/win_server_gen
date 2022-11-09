package handler

import (
	"fmt"
	"path"
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
}

func NewServer(name string) (server *Server) {
	//id := uuid.NewString()
	//id = strings.ReplaceAll(id, "-", "")
	return &Server{
		Id:             name,
		Name:           name,
		Description:    name,
		LogMode:        "roll",
		LogPath:        path.Join(BaseDir, "server-logs"),
		Executable:     path.Join(BaseDir, fmt.Sprintf("start%s.bat", name)),
		StopExecutable: path.Join(BaseDir, fmt.Sprintf("stop%s.bat", name)),
	}
}
