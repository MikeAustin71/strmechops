package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fileHelperPreon struct {
	lock *sync.Mutex
}

// makeDirAll
//
// Creates a directory named path, along with any
// necessary parent directories. In other words, all
// directories in the path are created.
//
// The permission bits 'drwxrwxrwx' are used for all
// directories that the method creates.
//
// If path is a directory which already exists, this
// method does nothing and returns and error value of
// 'nil'.
//
// Note:
//
// This method calls MakeDirAllPerm()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains the name of the directory
//		path which will be created by this method.
//
//	dirPathLabel				string
//
//		The name or label associated with input parameter
//		'dirPath' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dirPath" will be
//		automatically applied.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fileHelpPreon *fileHelperPreon) makeDirAll(
	dirPath string,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpPreon.lock == nil {
		fileHelpPreon.lock = new(sync.Mutex)
	}

	fileHelpPreon.lock.Lock()

	defer fileHelpPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperPreon."+
			"makeDirAll()",
		"")

	if err != nil {
		return err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "dirPath"
	}

	var permission FilePermissionConfig

	permission,
		err = new(FilePermissionConfig).
		New("drwxrwxrwx",
			ePrefix)

	if err != nil {
		return fmt.Errorf(
			"Error FilePermissionConfig.New()\n"+
				"Permission Code: \"drwxrwxrwx\"\n"+
				"Error=\n%v\n",
			err.Error())
	}

	err = new(fileHelperPlanck).makeDirAllPerm(
		dirPath,
		permission,
		dirPathLabel,
		ePrefix)

	if err != nil {
		return err
	}

	return nil
}
