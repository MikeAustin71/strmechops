package strmech

import (
	"io"
	"os"
	"sync"
)

type FileIoReader struct {
	fileIoReader       *io.Reader
	filePtr            *os.File
	targetReadFileName string

	lock *sync.Mutex
}
