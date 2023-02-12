package strmech

import (
	"sync"
)

type fileHelperMechanics struct {
	lock *sync.Mutex
}

// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
//
//	const (
//		 FILE OPEN TYPE - Select only ONE File Open Type Code
//		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
//		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
//		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
//		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
//
//		FILE OPEN MODES - May Select Multiple File Open Modes
//		// The remaining values may be or'ed in to control behavior.
//		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
//		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
//		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
//		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
//		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
//	)
func (fileHelpMech *fileHelperMechanics) openFile(
	fileOpenType int,
	errorPrefix interface{},
	fileOpenModes ...int) error {

	var err error

	return err
}
