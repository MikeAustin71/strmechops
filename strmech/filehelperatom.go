package strmech

import (
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type fileHelperAtom struct {
	lock *sync.Mutex
}

// adjustPathSlash
//
// This method will standardize path separators according
// to the current operating system.
//
// If input parameter 'path' contains invalid file path
// separators for the current operating system, this
// method will apply standard, compatible file path
// separator characters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	path				string
//
//		This 'path' string will be examined for
//		non-standard path separator characters in the
//		context of the current operating system. If
//		path separator characters are found to be
//		incompatible with the current operating system,
//		they will be replaced and a valid path string
//		will be returned to the calling function.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method will replace invalid path separators
//		found in input parameter 'path' and return a
//		string containing valid path separator characters
//		for the current operating system.
func (fHelperAtom *fileHelperAtom) adjustPathSlash(
	path string) string {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	errCode := 0

	errCode, _, path = new(fileHelperElectron).
		isStringEmptyOrBlank(path)

	if errCode == -1 {
		return ""
	}

	if errCode == -2 {
		return ""
	}

	if os.PathSeparator != '\\' {
		return strings.ReplaceAll(path, "\\", string(os.PathSeparator))
	}

	if os.PathSeparator != '/' {
		return strings.ReplaceAll(path, "/", string(os.PathSeparator))
	}

	return fp.FromSlash(path)
}
