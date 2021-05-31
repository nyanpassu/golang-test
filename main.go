package main

import (
	"io/fs"
	"log"
	"os"
	"time"
)

// FileInfo .
type FileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    fs.FileMode
	FileModTime time.Time
	Dir         bool
	FileSys     interface{}
}

// Name .
func (f FileInfo) Name() string {
	return f.FileName
}

// Size .
func (f FileInfo) Size() int64 {
	return f.FileSize
}

// Mode .
func (f FileInfo) Mode() fs.FileMode {
	return f.FileMode
}

// ModTime .
func (f FileInfo) ModTime() time.Time {
	return f.FileModTime
}

// IsDir .
func (f FileInfo) IsDir() bool {
	return f.Dir
}

// Sys .
func (f FileInfo) Sys() interface{} {
	return f.FileSys
}

type OS interface {
	Stat(string) (fs.FileInfo, error)
}

type impl struct{}

func (i impl) Stat(path string) (fs.FileInfo, error) {
	return os.Stat(path)
}

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "" {
		log.Fatalln("no given file")
	}
	var (
		a    fs.FileInfo
		err  error
		impl OS = impl{}
	)
	if a, err = impl.Stat(os.Args[1]); err != nil {
		log.Fatalln(err)
	}
	log.Println(a.Name())
}
