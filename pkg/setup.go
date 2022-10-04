package getup

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
)

var (
	_homeDirectory string
	_directory     string
	_fileName      string
	_initialized   bool
)

func New(homeDirectory string, directory string, fileName string) {
	if _initialized {
		return
	}
	_homeDirectory = homeDirectory
	_directory = directory
	_fileName = fileName
	os := runtime.GOOS
	switch os {
	case "linux":
		{
			break
		}
	case "windows":
		{
			_fileName += ".exe"
			break
		}
	case "darwin":
		{
			_fileName += ".dmg"
			break
		}
	default:
		{
			panic("go-painless does not support the current platform")
		}
	}
	_initialized = true
}

func Exists(filePath string) (*bool, error) {
	_, err := os.Stat(filePath)
	var output bool
	if err == nil {
		output = true
		return &output, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		output = false
		return &output, nil
	}
	return nil, err
}

func Setup() error {
	if !_initialized {
		return errors.New("not initialized")
	}
	path := fmt.Sprintf("%s/%s/%s", _homeDirectory, _directory, "bin")
	if os.Args[0] == fmt.Sprintf("%s/%s", path, _fileName) {
		return errors.New("invalid setup")
	}
	exists, err := Exists(path)
	if err != nil {
		panic(err)
	}
	if !*exists {
		os.MkdirAll(path, os.ModePerm)
	} else {
		os.Remove(fmt.Sprintf("%s/%s", path, _fileName))
	}
	src, err := os.Open(os.Args[0])
	if err != nil {
		panic(err)
	}
	dest, err := os.Create(fmt.Sprintf("%s/%s", path, _fileName))
	if err != nil {
		panic(err)
	}
	io.Copy(dest, src)
	src.Close()
	dest.Close()
	return nil
}
