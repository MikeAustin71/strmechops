package strmech

import "sync"

type FileHelper struct {
	lock *sync.Mutex
}

func (fileHelp *FileHelper) WriteFileString(
	filePathName string,
	stringToWrite interface{},
)
