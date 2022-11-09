package handler

import (
	"fmt"
	"io"
	"os"
	"win_server_gen/pkg/util/util"
)

func HandleCommandLine(msg string) bool {
	fmt.Printf("%s (yes/no (default no))", msg)
	command := ""
	_, _ = fmt.Scanln(&command)
	fmt.Println(command)
	for _, line := range []string{"1", "yes", "True", "true"} {
		if command == line {
			return true
		}
	}
	return false
}

func Write(file string, data string) (err error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()
	_, err = io.WriteString(f, data)
	if err != nil {
		err = fmt.Errorf("error writing file: %v", err)
	}
	return nil
}

func WriteFile(data string, dest string, option FileOption, fileExistHandler FileExistHandler) (err error) {
	if fileExistHandler == nil {
		fileExistHandler = FileExistDefaultHandler
	}
	writeFlag, err := fileExistHandler(FileExistOption{
		DistFile: dest,
		Option:   option,
	})
	if writeFlag {
		err = Write(dest, data)
	}
	return
}

type FileOption struct {
	OverWrite bool `json:"overwrite"`
}

func CopyFile(src string, dest string, option FileOption, fileExistHandler FileExistHandler) (err error) {
	if !util.FileExist(src) {
		return fmt.Errorf("file does not exist: %v", src)
	}
	srcFile, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer srcFile.Close()

	if fileExistHandler == nil {
		fileExistHandler = FileExistDefaultHandler
	}
	writeFlag, err := fileExistHandler(FileExistOption{
		DistFile: dest,
		SrcFile:  src,
		Option:   option,
	})
	if writeFlag {
		destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			err = fmt.Errorf("error creating %s: %v", dest, err)
		}
		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			err = fmt.Errorf("error copying: %v", err)
		}
	}

	return nil
}

type FileExistHandler func(option FileExistOption) (writeFlag bool, err error)

type FileExistOption struct {
	DistFile string
	SrcFile  string
	Option   FileOption
}

func FileExistQuestionHandler(option FileExistOption) (writeFlag bool, err error) {
	if !util.FileExist(option.DistFile) {
		return true, nil
	}
	if option.Option.OverWrite {
		_ = os.Remove(option.DistFile)
		return true, nil
	}
	if HandleCommandLine("file already exists. rewriting?") {
		_ = os.Remove(option.DistFile)
		return true, nil
	}
	return
}

func FileExistDefaultHandler(option FileExistOption) (writeFlag bool, err error) {
	if !util.FileExist(option.DistFile) {
		return true, nil
	}
	if option.Option.OverWrite {
		_ = os.Remove(option.DistFile)
		return true, nil
	}
	if util.FileExist(option.DistFile) {
		return false, fmt.Errorf("file already exists: %v", option.DistFile)
	}
	return
}

func FileExistIgnoreHandler(option FileExistOption) (writeFlag bool, err error) {
	if !util.FileExist(option.DistFile) {
		return true, nil
	}
	if option.Option.OverWrite {
		_ = os.Remove(option.DistFile)
		return true, nil
	}

	if option.SrcFile != "" {
		distFileInfo, _ := os.Lstat(option.DistFile)
		srcFileInfo, _ := os.Lstat(option.SrcFile)
		if distFileInfo.Size() != srcFileInfo.Size() {
			_ = os.Remove(option.DistFile)
			return true, nil
		} else {
			VerboseLog(fmt.Sprintf("src file & dest file is same ignore: \n%v\n%v\n", option.SrcFile, option.DistFile))
		}
	} else {
		VerboseLog(fmt.Sprintf("file exist ignore: %v\n", option.DistFile))
	}
	return false, err
}
