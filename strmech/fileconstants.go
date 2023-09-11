package strmech

import (
	"io"
	"sync"
)

var mSeekerWhenceCodes = map[int]string{
	io.SeekStart:   "io.SeekStart",
	io.SeekCurrent: "io.SeekCurrent",
	io.SeekEnd:     "io.SeekEnd",
}

type FileConstants struct {
	lock *sync.Mutex
}

// GetSeekerWhenceCodes
//
// Returns strings describing 'whence' constant values
// for file 'Seeker' methods.
func (fileConstants *FileConstants) GetSeekerWhenceCodes(
	whenceValue int) (
	okWhenceCode bool,
	whenceCodeStr string) {

	if fileConstants.lock == nil {
		fileConstants.lock = new(sync.Mutex)
	}

	fileConstants.lock.Lock()

	defer fileConstants.lock.Unlock()

	whenceCodeStr,
		okWhenceCode = mSeekerWhenceCodes[whenceValue]

	if !okWhenceCode {
		whenceCodeStr = "Invalid 'whence' code"
	}

	return okWhenceCode, whenceCodeStr
}
