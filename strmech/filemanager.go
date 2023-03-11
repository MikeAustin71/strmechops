package strmech

import (
	"bufio"
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// FileMgr
//
// This type and its associated methods are used to
// manage organize and control disk files and file
// permissions.
//
// To create an instance of type 'FileMgr' use one of the
// 'FileMgr.New' methods.
type FileMgr struct {
	isInitialized                   bool
	originalPathFileName            string
	dMgr                            DirMgr
	absolutePathFileName            string
	isAbsolutePathFileNamePopulated bool
	doesAbsolutePathFileNameExist   bool
	fileName                        string
	isFileNamePopulated             bool
	fileExt                         string
	isFileExtPopulated              bool
	fileNameExt                     string
	isFileNameExtPopulated          bool
	filePtr                         *os.File
	fileBufRdr                      *bufio.Reader
	fileRdrBufSize                  int
	fileBufWriter                   *bufio.Writer
	fileWriterBufSize               int
	isFilePtrOpen                   bool
	fileAccessStatus                FileAccessControl
	actualFileInfo                  FileInfoPlus
	fileBytesWritten                uint64
	buffBytesWritten                uint64
	lock                            *sync.Mutex
	// Used internally to ensure thread safe operations
}

// ChangePermissionMode
//
// This method is a wrapper for os.Chmod().
//
// ChangePermissionMode changes the permissions mode of
// the current FileMgr file to input paramter 'mode'. If
// the file is a symbolic link, it changes the mode of
// the link's target. If there is an error, it will be of
// type *PathError.
//
// A different subset of the mode bits are used, depending
// on the operating system.
//
// On Unix/Linux, the mode's permission bits, ModeSetuid,
// ModeSetgid, and ModeSticky are used.
//
// On Windows, the mode must be non-zero but otherwise only
// the 0200 bit (owner writable) of mode is used; it controls
// whether the file's read-only attribute is set or cleared.
// The other bits are currently unused. Use mode 0400 for a
// read-only file and 0600 for a readable+writable file.
//
// For Windows - Eligible permission codes are:
//
//	'modeStr'      Octal
//	 Symbolic      Mode Value     File Access
//
//	 -r--r--r--     0444          File - read only
//	 -rw-rw-rw-     0666          File - read & write
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	mode						FilePermissionConfig
//
//		An instance of FilePermissionConfig containing
//		the file permission specifications.
//
//		See the source code documentation for method
//		'FilePermissionConfig.New()' for an explanation
//		of permission codes.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) ChangePermissionMode(
	mode FilePermissionConfig,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"ChangePermissionMode()",
		"")

	if err != nil {
		return err
	}

	fMgrHelprAtom := fileMgrHelperAtom{}

	filePathDoesExist,
		err := fMgrHelprAtom.
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	if !filePathDoesExist {
		return fmt.Errorf("%v\n"+
			"Error: This file does NOT exist!\n"+
			"(FileMgr) File Name:'%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	err = mode.IsValidInstanceError(ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'mode' is invalid!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err.Error())
	}

	var permissionModeText string

	permissionModeText,
		err = mode.
		GetPermissionFileModeValueText(
			ePrefix.XCpy(
				"permissionModeText<-mode"))

	if err != nil {

		return err
	}

	var fileMode os.FileMode

	fileMode,
		err = mode.GetCompositePermissionMode(
		ePrefix.XCpy(
			"fileMode<-mode"))

	if err != nil {
		return err
	}

	err = os.Chmod(fMgr.absolutePathFileName, fileMode)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned by os.Chmod(fMgr.absolutePathFileName, fileMode).\n"+
			"fileMode='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			permissionModeText,
			err.Error())
	}

	err = nil
	_,
		err = fMgrHelprAtom.doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"Verify fMgr.absolutePathFileName")

	return err
}

// CloseThisFile
//
// This method will call the Close() method on the
// current file pointer, FileMgr.filePtr.
//
// In addition, if the file has been opened for Write or
// Read-Write, this method will automatically flush the
// file buffers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CloseThisFile(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CloseThisFile()",
		"")

	if err != nil {
		return err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.closeFile(fMgr, ePrefix)

	return err
}

// CopyFileMgrByIo
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter 'fMgrDest', an instance of
// type FileMgr.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// One attempt will be made to copy the source file to
// the specified destination file using a technique known
// as 'io.Copy'. This technique will create a new
// destination file and copy the source file contents to
// that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
// https://golang.org/pkg/io/#Copy
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest					*FileMgr
//
//		A pointer to a destination FileMgr instance.
//
//		This File Manager type specifies the path and
//		file name of the destination file to which the
//		source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with 'fMgrDest'
//		this method will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//	    Note:
//	    An error will be returned if the file identified
//	    by the current source File Manager instance does
//	    NOT exist.
func (fMgr *FileMgr) CopyFileMgrByIo(
	fMgrDest *FileMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileMgrByIo()",
		"")

	if err != nil {
		return err
	}

	fMgrHlpr := fileMgrHelper{}

	sourceFMgrLabel := "fMgrSource"
	destFMgrLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true,  // create target/destination directory if it does not exist
		false, // delete existing target/destination file
		ePrefix,
		sourceFMgrLabel,
		destFMgrLabel)

	if err == nil {
		err = fMgrHlpr.lowLevelCopyByIO(
			fMgr,
			fMgrDest,
			0,
			ePrefix,
			sourceFMgrLabel,
			destFMgrLabel)
	}

	return err
}

// CopyFileMgrByIoByLink
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter 'fMgrDest', an instance of
// type FileMgr.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the
// source file to the destination by creating a new
// destination file and copying the source file contents
// to that new destination file using a technique known
// as 'io.Copy'.
//
// If this first file copy operation fails, a second
// attempt will be made using a technique known as a
// 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest					*FileMgr
//
//		A pointer to a destination FileMgr instance.
//
//		This File Manager type specifies the path and
//		file name of the destination file to which the
//		source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with 'fMgrDest'
//		this method will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileMgrByIoByLink(
	fMgrDest *FileMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileMgrByIoByLink()",
		"")

	if err != nil {
		return err
	}

	fMgrHlpr := new(fileMgrHelper)

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true, //create destination directory
		true,
		ePrefix,
		"fMgrSource",
		"fMgrDest")

	if err != nil {
		return err
	}

	// See Reference:
	// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		fMgrDest,
		0, // local buffer size - take default
		ePrefix,
		"fMgrSource",
		"fMgrDest")

	if err != nil {
		err = fMgrHlpr.lowLevelCopyByLink(
			fMgr,
			fMgrDest,
			ePrefix,
			"fMgrSource",
			"fMgrDest")
	}

	return err
}

// CopyFileMgrByIoWithBuffer
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter 'fMgrDest', another
// instance of type FileMgr.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// One attempt will be made to copy the source file to
// the specified destination file using a technique known
// as 'io.Copy'. This technique will create a new
// destination file and copies the source file contents to
// that new destination file.
//
// If this attempted 'io.Copy' operation fails, and error
// will be returned.
//
// The size of the internal buffer used to copy the
// source file to the destination file may be specified
// by the calling function. The ability to control the
// internal buffer size is useful when copying large
// files.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/io/#CopyBuffer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest					FileMgr
//
//		This File Manager type specifies the path and
//		file name of the destination file to which the
//		source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with 'fMgrDest'
//		does not exist, this method will attempt to
//		create it.
//
//	bufferSize					int
//
//		The size in bytes of a local buffer which will be
//		used to copy the source File Manager (fMgr) to
//		the destination File Manager (fMgrDest). This is
//		useful when copying large files.
//
//		If this value is set equal to or less than zero,
//		the default internal buffer will be used.
//
//		Reference: https://golang.org/pkg/io/#CopyBuffer
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileMgrByIoWithBuffer(
	fMgrDest *FileMgr,
	bufferSize int,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileMgrByIoWithBuffer()",
		"")

	if err != nil {
		return err
	}

	sourceFMgrLabel := "fMgrSource"

	destFMgrLabel := "fMgrDest"

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true,  // create target/destination directory if it does not exist
		false, // delete target/destination file if it exists
		ePrefix.XCpy(
			"fMgr-fMgrDest"),
		sourceFMgrLabel,
		destFMgrLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		fMgrDest,
		bufferSize,
		ePrefix,
		sourceFMgrLabel,
		destFMgrLabel)

	return err
}

// CopyFileMgrByLink
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter 'fMgrDest', an
// instance of type FileMgr.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// This method will make one attempt to copy the
// source file to the specified destination using a
// technique known as a 'Hard Link'. This technique will
// utilize a hard symbolic link to the existing source
// file in order to create the destination file.
//
// By creating a 'linked' file, changing the contents of
// one file will be reflected in the second. The two
// linked files are 'mirrors' of each other.
//
// Consider using FileMgr.CopyFileMgrByIo() if the
// 'mirror' feature causes problems.
//
// If the 'Hard Link' copy operation fails, and error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/os/#Link
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest 					*FileMgr
//
//		A pointer to an instance of type FileMgr.
//
//		This File Manager type specifies the path and
//		file name of the destination file to which the
//		source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with 'fMgrDest'
//		does not exist, this method will attempt to
//		create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileMgrByLink(
	fMgrDest *FileMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileMgrByLink()",
		"")

	if err != nil {
		return err
	}

	fMgrHlpr := fileMgrHelper{}

	sourceFMgrLabel := "fMgrSource"

	destFMgrLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true, // create target/destination directory
		true, // delete existing target/destination file
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		fMgrDest,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	return err
}

// CopyFileMgrByLinkByIo
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter 'fMgrDest', an
// instance of type FileMgr.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the
// source file to the destination using a technique known
// as a 'Hard Link'.  This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If the first copy attempt fails, this method will try
// to copy the file to the destination by creating a new
// file and copying the source file contents to that new
// destination file. This technique is known as
// 'io.Copy'.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest 					*FileMgr
//
//		This File Manager type specifies the path and
//		file name of the destination file to which the
//		source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with 'fMgrDest'
//		this method will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileMgrByLinkByIo(
	fMgrDest *FileMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"CopyFileMgrByLinkByIo()",
		"")

	fMgrHlpr := fileMgrHelper{}

	sourceFMgrLabel := "fMgrSource"
	destFMgrLabel := "fMgrDestDir"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true,
		true,
		ePrefix.XCpy("fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		fMgrDest,
		ePrefix.XCpy("fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	if err != nil {

		err = fMgrHlpr.lowLevelCopyByIO(
			fMgr,
			fMgrDest,
			0,
			ePrefix.XCpy("fMgrDest<-fMgr"),
			sourceFMgrLabel,
			destFMgrLabel)
	}

	return err
}

// CopyFileMgrWithBuffer
//
// Performs a Copy By IO operation which copies the file
// identified by the current File Manager to the file
// specified by the input parameter 'fMgrDest'.
//
// 'fMgrDest' is another instance of 'FileMgr'.
//
// This method differs from a similar method,
// 'FileMgr.CopyFileMgrByIo()' in that the calling
// function must supply a local copy buffer of type
// []byte. The copy operation stages through the provided
// buffer (if one is required) rather than allocating a
// temporary internal buffer. This approach can be useful
// when copying a series of large files.
//
// Internally, this method calls io.CopyBuffer().
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://golang.org/pkg/io/#CopyBuffer
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrDest					*FileMgr
//
//		A pointer to an instance of type FileMgr.
//
//		This instance of FileMgr identifies the
//		destination file to which the source file
//		identified by the current FileMgr will be copied.
//
//	localCopyBuffer				[]byte
//
//		A byte array which will be used as an internal
//		buffer through which to stage the copy operation.
//		If this buffer is 'nil' or if this is a zero
//		length byte array, an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Note:
//		An error will be returned if the source file
//		identified by the current FileMgr instance does
//		NOT exist. Likewise, an error will also be
//		triggered if input parameter 'localCopyBuffer' is
//		'nil' or if it is a zero length byte array.
func (fMgr *FileMgr) CopyFileMgrWithBuffer(
	fMgrDest *FileMgr,
	localCopyBuffer []byte,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileMgrWithBuffer()",
		"")

	if err != nil {
		return err
	}

	fMgrHlpr := fileMgrHelper{}

	sourceFMgrLabel := "fMgrSource"
	destFMgrLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		fMgrDest,
		true,  // create target/destination directory if it does not exist
		false, // delete existing target/destination file
		ePrefix,
		sourceFMgrLabel,
		destFMgrLabel)

	if err == nil {
		err = fMgrHlpr.lowLevelCopyByIOBuffer(
			fMgr,
			fMgrDest,
			localCopyBuffer,
			ePrefix,
			sourceFMgrLabel,
			destFMgrLabel)
	}

	return err
}

// CopyFileStrByIo
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter. The destination input
// parameter, 'dstPathFileNameExt' is a string containing
// the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist, this method will
// attempt to create it.
//
// One attempt will be made to copy the source file to
// the specified destination using a technique known as
// 'Copy By IO'. This technique create a new destination
// file and copies the source file contents to that new
// destination file.
//
// If this attempted 'io.Copy' operation fails, and error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/io/#Copy
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dstPathFileNameExt			string
//
//		'dstPathFileNameExt' is a string containing
//		the path, file name and file extension of the
//		destination file.
//
//		If the destination directory path contained in
//		this string does not exist, this method will
//		attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileStrByIo(
	dstPathFileNameExt string,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileMgr.CopyFileStrByIo()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelper).
		setFileMgrPathFileName(
			&fMgrDest,
			dstPathFileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dstPathFileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n"+
			"Error=\n%v\n",
			funcName,
			dstPathFileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n",
			ePrefix.String(),
			dstPathFileNameExt)

	}

	fMgrSourceLabel := "fMgrSource"

	fMgrDestLabel := "fMgrDest"

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true,  // create destination directory
		false, // delete destination file if it exists
		ePrefix,
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		&fMgrDest,
		0, // Local Buffer Size - Use Default
		ePrefix,
		fMgrSourceLabel,
		fMgrDestLabel)

	return err
}

// CopyFileStrByIoByLink
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter. The destination input
// parameter, 'dstPathFileNameExt' is a string
// containing the path, file name and file extension
// of the destination file.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the file
// to the destination by creating a new file and copying
// the source file contents to the new destination file.
// This technique is known as 'io.Copy'.
//
// If that attempted file copy operation fails, a second
// attempt will be made using a technique known as a
// 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dstPathFileNameExt			string
//
//		This string specifies the path, file name and
//		file extension of the destination file to which
//		the source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with
//		'dstPathFileNameExt' does not exist, this method
//		will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileStrByIoByLink(
	dstPathFileNameExt string,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "FileMgr.CopyFileStrByIoByLink()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	fMgrHlpr := fileMgrHelper{}

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = fMgrHlpr.setFileMgrPathFileName(
		&fMgrDest,
		dstPathFileNameExt,
		ePrefix.XCpy("fMgrDest<-dstPathFileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			funcName,
			dstPathFileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n",
			ePrefix.String(),
			dstPathFileNameExt)

	}

	fMgrSourceLabel := "fMgrSource"

	fMgrDestLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // Create Destination Directory
		true, // Delete Destination File if it exists
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		&fMgrDest,
		0,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		err = fMgrHlpr.lowLevelCopyByLink(
			fMgr,
			&fMgrDest,
			ePrefix.XCpy(
				"fMgrDest<-fMgr"),
			fMgrSourceLabel,
			fMgrDestLabel)
	}

	return err
}

// CopyFileStrByLink
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter. The destination input
// parameter, 'dstPathFileNameExt' is a string containing
// the path, file name and file extension of the
// destination file.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// This method will make one attempt to copy the source
// file to the specified destination using a technique
// known as a 'Hard Link'. This technique will utilize a
// hard symbolic link to the existing source file in
// order to create the destination file.
//
// If 'Hard Link' copy operation fails, and error will be
// returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dstPathFileNameExt			string
//
//		This string specifies the path, file name and
//		file extension of the destination file to which
//		the source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with
//		'dstPathFileNameExt' does not exist, this method
//		will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileStrByLink(
	dstPathFileNameExt string,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "FileMgr.CopyFileStrByLink()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	var fMgrDest FileMgr

	fMgrHlpr := fileMgrHelper{}

	var isEmpty bool

	isEmpty,
		err = fMgrHlpr.
		setFileMgrPathFileName(
			&fMgrDest,
			dstPathFileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dstPathFileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt'\n"+
			"Error returned by FileMgr{}.New(dstPathFileNameExt)\n"+
			"dstPathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			funcName,
			dstPathFileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n",
			ePrefix.String(),
			dstPathFileNameExt)

	}

	fMgrSourceLabel := "fMgrSource"

	fMgrDestLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // create destination directory
		true, // delete destination file if it exists
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		&fMgrDest,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	return err
}

// CopyFileStrByLinkByIo
//
// Copies the file represented by the current File
// Manager instance to a location specified by a
// destination input parameter. The destination input
// parameter, 'dstPathFileNameExt' is a string
// containing the path, file name and file extension of
// the destination file.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the
// source file to the destination using a technique known
// as a 'Hard Link'.  This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If the first copy attempt fails, this method will try
// to copy the file to the destination by creating a new
// file and copying the source file contents to the new
// destination file. This technique is known as
// 'io.Copy'.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dstPathFileNameExt			string
//
//		This string specifies the path, file name and
//		file extension of the destination file to which
//		the source file identified by the current File
//		Manager will be copied.
//
//		If the directory path associated with
//		'dstPathFileNameExt' does not exist, this method
//		will attempt to create it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileStrByLinkByIo(
	dstPathFileNameExt string,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "FileMgr.CopyFileStrByLinkByIo()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	var fMgrDest FileMgr

	var isEmpty bool

	fMgrHlpr := fileMgrHelper{}

	isEmpty,
		err = fMgrHlpr.
		setFileMgrPathFileName(
			&fMgrDest,
			dstPathFileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dstPathFileNameExt"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			funcName,
			dstPathFileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Invalid input parameter 'dstPathFileNameExt' is invalid!\n"+
			"dstPathFileNameExt='%v'\n",
			ePrefix.String(),
			dstPathFileNameExt)
	}

	fMgrSourceLabel := "fMgrSource"

	fMgrDestLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // Create Destination Directory
		true, // Delete Destination File if it exists
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		&fMgrDest,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {

		err = fMgrHlpr.lowLevelCopyByIO(
			fMgr,
			&fMgrDest,
			0,
			ePrefix.XCpy(
				"fMgrDest<-fMgr"),
			fMgrSourceLabel,
			fMgrDestLabel)
	}

	return err
}

// CopyFileToDirByIo
//
// Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by
// input parameter 'dir', an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// One attempt will be made to copy the source file to
// the specified destination directory using a technique
// known as 'io.Copy'. This technique creates a new
// destination file and copies the source file contents
// to that new destination file.
//
// If this attempted 'io.Copy' operation fails, an error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dir							DirMgr
//
//		An instance of Directory Manager ('DirMgr').
//
//		This Directory Manager contains the name of the
//		directory to which the file identified by the
//		current FileMgr will be copied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileToDirByIo(
	dir DirMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileToDirByIo()",
		"")

	if err != nil {
		return err
	}

	err = dir.IsDirMgrValid(
		ePrefix.String() + "Input Parameter 'dir' ")

	if err != nil {
		return err
	}

	sourceFMgrLabel := "fMgrSource"
	destFMgrLabel := "fMgrDestDir"

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgrDest,
			&dir,
			fMgr.fileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dir-fMgr.fileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: The FileMgr instance generated by input parameter 'dir' and "+
			"'fMgr.fileNameExt' produced an error!\n"+
			"dMgr='%v'\n"+
			"fMgr.fileNameExt='%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Error: The FileMgr instance generated by input parameter 'dir' and "+
			"'fMgr.fileNameExt' is Empty!\n"+
			"dMgr='%v'\n"+
			"fMgr.fileNameExt='%v'\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt)
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true,  // create target/destination directory if it does not exist
		false, // delete existing target/destination file
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		&fMgrDest,
		0, // Local Buffer Size = default
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		sourceFMgrLabel,
		destFMgrLabel)

	return err
}

// CopyFileToDirByIoByLink
//
// Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by
// input parameter 'dir', an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the file
// to the destination by creating a new file and copying
// the source file contents to the new destination file
// using a technique known as 'io.Copy' or 'Copy by IO'.
//
// If that attempted file copy operation fails, a second
// attempt will be made using a technique known as a
// 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dir							DirMgr
//
//		An instance of Directory Manager ('DirMgr').
//
//		This Directory Manager contains the name of the
//		directory to which the file identified by the
//		current FileMgr will be copied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileToDirByIoByLink(
	dir DirMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"CopyFileToDirByIoByLink()",
		"")

	if err != nil {
		return err
	}

	err = dir.IsDirMgrValid(
		ePrefix.String() +
			"Input Parameter 'dir' ")

	if err != nil {
		return err
	}

	fMgrSourceLabel := "fMgrSource"
	fMgrDestLabel := "fMgrDestDir"

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgrDest,
			&dir,
			fMgr.fileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dir-fMgr.fileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an error!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an empty result!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt)
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // Create Destination Directory
		true, // Delete Destination File if it exists
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByIO(
		fMgr,
		&fMgrDest,
		0,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		err = fMgrHlpr.lowLevelCopyByLink(
			fMgr,
			&fMgrDest,
			ePrefix.XCpy(
				"fMgrDest<-fMgr"),
			fMgrSourceLabel,
			fMgrDestLabel)
	}

	return err
}

// CopyFileToDirByLink
//
// Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by
// input parameter 'dir', an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// This method will make one attempt to copy the source
// file to the specified destination directory using a
// technique known as a 'Hard Link'. This technique will
// utilize a hard symbolic link to the existing source
// file in order to create the destination file.
//
// If the 'Hard Link' copy operation fails, and error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dir							DirMgr
//
//		An instance of Directory Manager ('DirMgr').
//
//		This Directory Manager contains the name of the
//		directory to which the file identified by the
//		current FileMgr will be copied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileToDirByLink(
	dir DirMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CopyFileToDirByLink()",
		"")

	err = dir.IsDirMgrValid(ePrefix.String() +
		"Input Parameter 'dir' ")

	if err != nil {
		return err
	}

	fMgrSourceLabel := "fMgrSource"
	fMgrDestLabel := "fMgrDestDir"

	fMgrHlpr := fileMgrHelper{}

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgrDest,
			&dir,
			fMgr.fileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dir-fMgr.fileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an error!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an empty result!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt)
	}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // Create Destination Directory
		true, // Delete Destination File if it exists
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		&fMgrDest,
		ePrefix.XCpy(
			"fMgrDest<-fMgr"),
		fMgrSourceLabel,
		fMgrDestLabel)

	return err
}

// CopyFileToDirByLinkByIo
//
// Copies the file identified by the current File Manager
// (FileMgr) instance to another directory specified by
// input parameter 'dir', an instance of type 'DirMgr'.
//
// Note that if the destination directory does not exist,
// this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the
// source file to the destination directory using a
// technique known as a 'Hard Link'.  This technique will
// utilize a hard symbolic link to the existing source
// file in order to create the destination file.
//
// If the first copy attempt fails, this method will try
// to copy the file to the destination directory by
// creating a new file and copying the source file
// contents to the new destination file. This technique
// is known as 'io.Copy'.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dir							DirMgr
//
//		An instance of Directory Manager ('DirMgr').
//
//		This Directory Manager contains the name of the
//		directory to which the file identified by the
//		current FileMgr will be copied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CopyFileToDirByLinkByIo(
	dir DirMgr,
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"CopyFileToDirByLinkByIo()",
		"")

	err = dir.IsDirMgrValid(ePrefix.String() +
		"Input Parameter 'dir'")

	if err != nil {
		return err
	}

	fMgrSourceLabel := "fMgrSource"
	fMgrDestLabel := "fMgrDestDir"

	fMgrHlpr := fileMgrHelper{}

	var fMgrDest FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgrDest,
			&dir,
			fMgr.fileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dir-fMgr.fileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an error!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dir' and fMgr.fileNameExt\n"+
			"produced an empty result!\n"+
			"dir='%v'\n"+
			"fMgr.fileNameExt='%v'\n",
			ePrefix.String(),
			dir.absolutePath,
			fMgr.fileNameExt)
	}

	err = fMgrHlpr.copyFileSetup(
		fMgr,
		&fMgrDest,
		true, // Create Destination Directory
		true, // Delete Destination File if it exists
		ePrefix.XCpy(
			"fMgr<-fMgrDest"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		return err
	}

	err = fMgrHlpr.lowLevelCopyByLink(
		fMgr,
		&fMgrDest,
		ePrefix.XCpy(
			"fMgr<-fMgrDest"),
		fMgrSourceLabel,
		fMgrDestLabel)

	if err != nil {
		err = fMgrHlpr.lowLevelCopyByIO(
			fMgr,
			&fMgrDest,
			0,
			ePrefix.XCpy(
				"fMgr<-fMgrDest"),
			fMgrSourceLabel,
			fMgrDestLabel)
	}

	return err
}

// CopyFromStrings
//
// Copies a source file to a destination file. The source
// and destination files are identified by two strings
// specifying path, file name, and file extensions. The
// copy operation employs the 'Copy By IO' technique
// (a.k.a. 'io.Copy').
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
// https://golang.org/pkg/io/#Copy
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the destination path, file name and file extension
//	does NOT exist, this method will create it.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourcePathFileNameExt		string
//
//		A text string which identifies the path, file
//		name and file extension of the source file which
//		will be used in the copy operation.
//
//	destPathFileNameExt			string
//
//		A text string which identifies the path, file
//		name and file extension of the destination file
//		which will be used in the copy operation.
//
//		If this destination path, file name and file
//		extension does NOT exist, this method will create
//		it.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	srcFMgr						FileMgr
//
//		If successful, this method returns a File Manager
//		instance which identifies the source file used in
//		the copy operation.
//
//	destFMgr					FileMgr
//
//		If successful, this method returns a File Manager
//		instance which identifies the destination file
//		used in the copy operation.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Note:
//
//		An error will be returned if the file identified
//		by the input parameter 'sourcePathFileNameExt'
//		does NOT exist.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	srcFileMgr, destFileMgr, err := FileMgr{}.CopyFromStrings(
//	                                "../Dir1/srcFile",
//	                                "../Dir2/destFile")
func (fMgr *FileMgr) CopyFromStrings(
	sourcePathFileNameExt,
	destPathFileNameExt string,
	errorPrefix interface{}) (
	fMgrSrc FileMgr,
	fMgrDest FileMgr,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileMgr.CopyFromStrings()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return fMgrSrc, fMgrDest, err
	}

	fMgrHlpr := fileMgrHelper{}

	fMgrDest = FileMgr{}

	var err2 error

	var isEmpty bool

	isEmpty,
		err2 = fMgrHlpr.
		setFileMgrPathFileName(
			&fMgrSrc,
			sourcePathFileNameExt,
			ePrefix.XCpy(
				"fMgrSrc<-sourcePathFileNameExt"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourcePathFileNameExt' is invalid!\n"+
			"sourcePathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourcePathFileNameExt,
			err2.Error())

		return fMgrSrc, fMgrDest, err
	}

	if isEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourcePathFileNameExt' is invalid!\n"+
			"sourcePathFileNameExt='%v'\n",
			ePrefix.String(),
			sourcePathFileNameExt)

		return fMgrSrc, fMgrDest, err

	}

	isEmpty,
		err2 = fMgrHlpr.
		setFileMgrPathFileName(
			&fMgrDest,
			destPathFileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-destPathFileNameExt"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destPathFileNameExt' is invalid!\n"+
			"destPathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			funcName,
			destPathFileNameExt,
			err2.Error())

		return fMgrSrc, fMgrDest, err
	}

	if isEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destPathFileNameExt' is invalid!\n"+
			"destPathFileNameExt='%v'\n",
			ePrefix.String(),
			destPathFileNameExt)

		return fMgrSrc, fMgrDest, err

	}

	sourceFMgrLabel := "fMgrSource"
	destFMgrLabel := "fMgrDest"

	err = fMgrHlpr.copyFileSetup(
		&fMgrSrc,
		&fMgrDest,
		true,  // create target/destination directory if it does not exist
		false, // delete existing target/destination file
		ePrefix,
		sourceFMgrLabel,
		destFMgrLabel)

	if err == nil {

		err = fMgrHlpr.lowLevelCopyByIO(
			&fMgrSrc,
			&fMgrDest,
			0,
			ePrefix,
			sourceFMgrLabel,
			destFMgrLabel)
	}

	return fMgrSrc, fMgrDest, err
}

// CopyIn - Copies data from an incoming FileMgr object
// into the current FileMgr object.
//
// Note: Internal file pointer will not be copied.
func (fMgr *FileMgr) CopyIn(fmgr2 *FileMgr) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	if fmgr2 == nil {

		fmgr2 = &FileMgr{}
	}

	fMgr.isInitialized = fmgr2.isInitialized
	fMgr.dMgr.CopyIn(&fmgr2.dMgr)
	fMgr.originalPathFileName = fmgr2.originalPathFileName
	fMgr.absolutePathFileName = fmgr2.absolutePathFileName
	fMgr.isAbsolutePathFileNamePopulated = fmgr2.isAbsolutePathFileNamePopulated
	fMgr.doesAbsolutePathFileNameExist = fmgr2.doesAbsolutePathFileNameExist
	fMgr.fileName = fmgr2.fileName
	fMgr.isFileNamePopulated = fmgr2.isFileNamePopulated
	fMgr.fileExt = fmgr2.fileExt
	fMgr.isFileExtPopulated = fmgr2.isFileExtPopulated
	fMgr.fileNameExt = fmgr2.fileNameExt
	fMgr.isFileNameExtPopulated = fmgr2.isFileNameExtPopulated
	fMgr.filePtr = nil
	fMgr.isFilePtrOpen = false
	fMgr.fileAccessStatus = fmgr2.fileAccessStatus.CopyOut()
	fMgr.actualFileInfo = fmgr2.actualFileInfo.CopyOut()
	fMgr.fileBytesWritten = 0
	fMgr.buffBytesWritten = 0
	fMgr.fileRdrBufSize = fmgr2.fileRdrBufSize
	fMgr.fileWriterBufSize = fmgr2.fileWriterBufSize

	return
}

// CopyOut
//
// Duplicates the file information in the current FileMgr
// object and returns it as a new FileMgr object.
//
// Note: Internal File Pointer will not be copied.
func (fMgr *FileMgr) CopyOut() FileMgr {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	fmgr2 := FileMgr{}

	fmgr2.isInitialized = fMgr.isInitialized
	fmgr2.dMgr = fMgr.dMgr.CopyOut()
	fmgr2.originalPathFileName = fMgr.originalPathFileName
	fmgr2.absolutePathFileName = fMgr.absolutePathFileName
	fmgr2.isAbsolutePathFileNamePopulated = fMgr.isAbsolutePathFileNamePopulated
	fmgr2.doesAbsolutePathFileNameExist = fMgr.doesAbsolutePathFileNameExist
	fmgr2.fileName = fMgr.fileName
	fmgr2.isFileNamePopulated = fMgr.isFileNamePopulated
	fmgr2.fileExt = fMgr.fileExt
	fmgr2.isFileExtPopulated = fMgr.isFileExtPopulated
	fmgr2.fileNameExt = fMgr.fileNameExt
	fmgr2.isFileNameExtPopulated = fMgr.isFileNameExtPopulated
	fmgr2.filePtr = nil
	fmgr2.isFilePtrOpen = false
	fmgr2.fileAccessStatus = fMgr.fileAccessStatus.CopyOut()

	fmgr2.actualFileInfo = fMgr.actualFileInfo.CopyOut()

	fmgr2.fileBytesWritten = 0
	fmgr2.buffBytesWritten = 0
	fmgr2.fileRdrBufSize = fMgr.fileRdrBufSize
	fmgr2.fileWriterBufSize = fMgr.fileWriterBufSize

	return fmgr2
}

// CreateDir
//
// Creates the directory previously configured for this
// file manager instance.
//
// If the directory creation process fails, an error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CreateDir(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.CreateDir()",
		"")

	if err != nil {
		return err
	}

	err = new(fileMgrHelper).
		createDirectory(fMgr, ePrefix)

	return err
}

// CreateDirAndFile
//
// Performs two operations:
//
// If the home directory does not currently exist, this
// method will first create the directory tree containing
// the file.
//
// Next, the file will be created. If the file previously
// exists, it will be truncated.
//
// Note that if this method successfully creates the
// file, a File Pointer (*File) will be stored in the
// FileMgr field fMgr.filePtr. Be sure to close the
// File Pointer when finished with it by calling:
//
//	FileMgr.CloseThisFile().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CreateDirAndFile(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"CreateDirAndFile()",
		"")

	if err != nil {
		return err
	}

	return new(fileMgrHelper).createFile(
		fMgr,
		true,
		ePrefix.XCpy(
			"fMgr"))
}

// CreateThisFile
//
// Creates the File identified by FileMgr
// (FileMgr.absolutePathFileName).
//
// If the directory in the path file name designation
// does not exist, this method will return an error.
//
// See Method CreateDirAndFile() which will create both
// the directory and the file.
//
// Note that if the file is actually created, the
// returned file pointer (*File) is stored in the FileMgr
// field, fMgr.filePtr. Be sure to 'close' the File
// Pointer when finished with it by calling:
//
//	FileMgr.CloseThisFile()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) CreateThisFile(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"CreateThisFile()",
		"")

	if err != nil {
		return err
	}

	return new(fileMgrHelper).createFile(
		fMgr,
		false,
		ePrefix.XCpy(
			"fMgr"))
}

// DeleteThisFile
//
// Deletes the file identified by
// FileMgr.absolutePathFileName in the current FileMgr
// instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) DeleteThisFile(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"DeleteThisFile()",
		"")

	if err != nil {
		return err
	}

	return new(fileMgrHelper).
		deleteFile(
			fMgr,
			ePrefix.XCpy("fMgr"))
}

// DoesFileExist
//
// Returns 'true' if the subject FileMgr file does in
// fact exist. If the file does NOT exist, a boolean
// value of 'false' is returned.
//
// This method uses os.Stat() to test for the existence
// of a file path. If a non-path error is returned by
// os.Stat(), it is ignored and the file path is
// classified as 'Does Not Exist'.
//
// This is very similar to 'FileMgr.DoesThisFileExist()'.
// However, unlike this method,
// 'FileMgr.DoesThisFileExist()' will return a non-path
// error. This method only returns a boolean value
// signaling whether the file path exists.
//
// If this method encounters a non-path error or if the
// current FileMgr instance is invalid, a boolean value
// of 'false' will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		Returns 'true' if the subject FileMgr file does
//		in fact exist. If the file does NOT exist, a
//	 	boolean value of 'false' is returned.
func (fMgr *FileMgr) DoesFileExist() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	var pathFileDoesExist bool

	pathFileDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		nil,
		"fMgr.absolutePathFileName")

	if err != nil {
		return false
	}

	return pathFileDoesExist
}

// DoesThisFileExist
//
// Returns a boolean value designating whether the file
// specified by the current FileMgr.absolutePathFileName
// exists.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	fileDoesExist				bool
//
//		Returns a boolean value designating whether the
//		file specified by the current
//		FileMgr.absolutePathFileName exists.
//
//	nonPathError				error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) DoesThisFileExist(
	errorPrefix interface{}) (
	fileDoesExist bool,
	nonPathError error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	fileDoesExist = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.DoesThisFileExist()",
		"")

	if nonPathError != nil {
		return fileDoesExist, nonPathError
	}

	fileDoesExist,
		nonPathError = new(fileMgrHelperAtom).doesFileMgrPathFileExist(fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if nonPathError != nil {
		fileDoesExist = false
	}

	return fileDoesExist, nonPathError
}

// Equal
//
// Compares a second FileMgr data structure to the
// current FileMgr data structure and returns a boolean
// value indicating whether they are equal in all
// respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr2						*FileMgr
//
//		A pointer to an incoming instance of FileMgr.
//		The data fields contained in this instance will
//		be compared to corresponding data fields in the
//		current FileMgr instance to determine if they are
//		equal in all respects.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the data fields contained in input parameter
//		'fMgr2' are equivalent to the corresponding data
//		field values in the current FileMgr instance,
//		this return value is set to 'true'.
//
//		If the data fields in these two FileMgr instances
//		are not equal in all respects, this return value
//		is set to 'false'.
func (fMgr *FileMgr) Equal(
	fMgr2 *FileMgr) bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	if fMgr2 == nil {
		fMgr2 = &FileMgr{}
	}

	result := true

	if fMgr.isInitialized != fMgr2.isInitialized ||
		fMgr.originalPathFileName != fMgr2.originalPathFileName ||
		fMgr.isAbsolutePathFileNamePopulated != fMgr2.isAbsolutePathFileNamePopulated ||
		fMgr.doesAbsolutePathFileNameExist != fMgr2.doesAbsolutePathFileNameExist ||
		fMgr.absolutePathFileName != fMgr2.absolutePathFileName ||
		fMgr.fileName != fMgr2.fileName ||
		fMgr.isFileNamePopulated != fMgr2.isFileNamePopulated ||
		fMgr.fileExt != fMgr2.fileExt ||
		fMgr.isFileExtPopulated != fMgr2.isFileExtPopulated ||
		fMgr.fileNameExt != fMgr2.fileNameExt ||
		fMgr.isFileNameExtPopulated != fMgr2.isFileNameExtPopulated ||
		fMgr.filePtr != fMgr2.filePtr ||
		fMgr.isFilePtrOpen != fMgr2.isFilePtrOpen ||
		fMgr.fileRdrBufSize != fMgr2.fileRdrBufSize ||
		fMgr.fileWriterBufSize != fMgr2.fileWriterBufSize {

		result = false
	}

	if result == true &&
		!fMgr.fileAccessStatus.Equal(&fMgr2.fileAccessStatus) {
		result = false
	}

	if result == true &&
		!fMgr.dMgr.Equal(&fMgr2.dMgr) {
		result = false
	}

	if result == true &&
		!fMgr.actualFileInfo.Equal(&fMgr2.actualFileInfo) {
		result = false
	}

	return result
}

// EqualAbsPaths
//
// Returns 'true' if both the current File Manager and
// the input File Manager ('fMgr2') have the same file
// paths.
//
// In other words, this method answers the question,
// 'Do Both Files have the same directory?'.
//
// The path comparisons are case-insensitive. This means
// that both paths will be converted to lower case before
// making the comparison.
//
// Also, the path comparison will be performed on the
// absolute directory paths associated with the two File
// Managers.
//
// If the directory paths are NOT equal, this method
// returns a boolean value of 'false.
//
// NOTE: This method will NOT test the equality of file
// names and extensions. ONLY the file paths
// (directories) will be compared.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr2						*FileMgr
//
//		A pointer to an incoming instance of FileMgr.
//
//		This method will compare the file path (a.k.a.
//		directory) contained in 'fMgr2' to that contained
//		in the current FileMgr instance to determine if
//		they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the file paths (a.k.a. directories) contained
//		in input parameter 'fMgr2' and the current
//		FileMgr instance are equivalent, this method
//		returns a boolean value of 'true'.
//
//		If they are NOT equivalent, a value of 'false' is
//		returned.
func (fMgr *FileMgr) EqualAbsPaths(
	fMgr2 *FileMgr) bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	if fMgr2 == nil {
		fMgr2 = &FileMgr{}
	}

	fDirMgr := fMgr.GetDirMgr()

	fDirMgr2 := fMgr2.GetDirMgr()

	return fDirMgr.EqualAbsPaths(&fDirMgr2)
}

// EqualFileNameExt
//
// Returns 'true' if both the current File Manager and
// the input File Manager parameter ('fMgr2') have the
// same file name and file extension.
//
// The File Name and File Extension comparisons are
// case-insensitive. This means that both file name and
// file extension will be converted to lower case before
// making the comparison.
//
// Example:
//
//	xray.txt is considered equal to XRAY.TXT
//
// If either the File Name or File Extension are NOT
// equal, this method will return 'false'.
//
// NOTE: This method will NOT test the equality of the
// file paths (a.k.a. directories). ONLY the file name
// and file extension will be tested for equality.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr2						*FileMgr
//
//		A pointer to a incoming instance of FileMgr.
//
//		This method will compare the File Name and File
//		Extension contained in 'fMgr2' to that contained
//		in the current FileMgr instance to determine if
//		they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the File Name and File Extension contained in
//		input parameter 'fMgr2' and the current instance
//		of FileMgr, this method will return a boolean
//		value of 'true'.
//
//		If the File Names and File Extensions are NOT
//		equivalent, this method returns 'false'.
func (fMgr *FileMgr) EqualFileNameExt(
	fMgr2 *FileMgr) bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	if fMgr2 == nil {
		fMgr2 = &FileMgr{}
	}

	f1FileName := ""
	f2FileName := ""
	f1FileExt := ""
	f2FileExt := ""

	f1FileName = fMgr.fileName
	f2FileName = fMgr2.fileName
	f1FileExt = fMgr.fileExt
	f2FileExt = fMgr2.fileExt

	f1FileName = strings.ToLower(f1FileName)

	f2FileName = strings.ToLower(f2FileName)

	if f1FileName != f2FileName {
		return false
	}

	f1FileExt = strings.ToLower(f1FileExt)
	f2FileExt = strings.ToLower(f2FileExt)

	if f1FileExt != f2FileExt {
		return false
	}

	return true
}

// EqualPathFileNameExt
//
// Returns 'true' if both the current File Manager and
// the input File Manager ('fMgr2') have the same
// absolute path, file name and file extension.
//
// The string comparisons are case-insensitive. This
// means that the paths, file names and file extensions
// will all be converted to lower case before making the
// comparison.
//
//	Example:
//		d:\dir1\xray.txt is considered equal to
//			D:\DIR1\XRAY.TXT
//
// If the path, file name or file extensions are NOT
// equal, this method will return 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr2						*FileMgr
//
//		A pointer to an incoming instance of FileMgr.
//
//		The path (a.k.a. directory), file name and
//		file extension contained in 'fMgr2' will be
//		compared to the path, file name and file
//		extension contained in the current instance of
//		FileMgr to determine if they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the path (a.k.a. directory), file name and
//		file extension contained in input parameter
//		'fMgr2' are equivalent to the path, file name and
//		file extension contained in the current instance
//		of FileMgr, this method will return a boolean
//		value of 'true'.
//
//		If path, file name and file extension are NOT
//		equivalent, this method returns 'false'.
func (fMgr *FileMgr) EqualPathFileNameExt(
	fMgr2 *FileMgr) bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	if fMgr2 == nil {
		fMgr2 = &FileMgr{}
	}

	f1AbsolutePathFileName := ""
	f2AbsolutePathFileName := ""

	f1AbsolutePathFileName =
		fMgr.absolutePathFileName

	f2AbsolutePathFileName =
		fMgr2.absolutePathFileName

	f1AbsolutePathFileName =
		strings.ToLower(f1AbsolutePathFileName)

	f2AbsolutePathFileName =
		strings.ToLower(f2AbsolutePathFileName)

	if f1AbsolutePathFileName != f2AbsolutePathFileName {
		return false
	}

	return true
}

// Empty
//
// Resets all data fields in the FileMgr structure to
// their uninitialized or zero states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values in the current instance of FileMgr to
//	their uninitialized or zero states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (fMgr *FileMgr) Empty() {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	_ = new(fileMgrHelperBoson).emptyFileMgr(
		fMgr, nil)

	return
}

// FlushBytesToDisk
//
// After Writing bytes to a file, use this method to
// commit the contents of the current file to stable
// storage.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) FlushBytesToDisk(
	errorPrefix interface{}) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.FlushBytesToDisk()",
		"")

	if err != nil {
		return err
	}

	return new(fileMgrHelperBoson).flushBytesToDisk(
		fMgr, ePrefix.XCpy(
			"fMgr"))
}

// GetAbsolutePath
//
// Returns the absolute path for the current File Manager
// (FileMgr) instance.
//
// Note: The file name and file extension are NOT
// included. Only the absolute path (a.k.a. directory) is
// returned as a 'string'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This string returns the absolute path (a.k.a.
//		directory) for the current FileMgr instance.
func (fMgr *FileMgr) GetAbsolutePath() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	return fMgr.dMgr.GetAbsolutePath()
}

// GetAbsolutePathFileName
//
// Returns the absolute path, file name and file
// extension for the current File Manager (FileMgr)
// instance.
//
// The returned path, file name and file extension string
// may contain both upper and lower case characters.
//
// See the companion method GetAbsolutePathFileNameLc()
// which returns a path/file name string consisting
// entirely of lower case text characters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This string returns the absolute path, file name
//		and file extension for the current File Manager
//		(FileMgr) instance.
func (fMgr *FileMgr) GetAbsolutePathFileName() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var absolutePathFileName string

	if !fMgr.isInitialized {

		absolutePathFileName = ""

	} else {

		absolutePathFileName = fMgr.absolutePathFileName
	}

	return absolutePathFileName
}

// GetAbsolutePathFileNameLc
//
// Returns the absolute path, file name and file
// extension for the current File Manager (FileMgr)
// instance.
//
// The returned path, file name and file extension string
// will consist entirely of lower case characters.
//
// See companion method GetAbsolutePathFileName() which
// returns the same path/file name string with upper and
// lower case characters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This string returns the absolute path, file name
//		and file extension for the current File Manager
//		(FileMgr) instance formatted entirely with lower
//		case text characters.
func (fMgr *FileMgr) GetAbsolutePathFileNameLc() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var absolutePathFileName string

	if !fMgr.isInitialized {
		absolutePathFileName = ""
	} else {

		absolutePathFileName =
			strings.ToLower(fMgr.absolutePathFileName)
	}

	return absolutePathFileName
}

// GetBufioReader
//
// Returns a pointer to the internal bufio.Reader,
// FileMgr.fileBufRdr. This pointer is initialized when
// the file is opened for Read or Read-Write operations.
//
// Be advised that if the file identified by the current
// FileMgr instance has not been opened the returned
// pointer to 'fMgr.fileBufRdr' may be nil.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*bufio.Reader
//
//		A pointer to the internal bufio.Reader,
//		FileMgr.fileBufRdr.
//
//		If the file identified by the current FileMgr
//		instance has not been opened the returned pointer
//		to 'FileMgr.fileBufRdr' may be nil.
func (fMgr *FileMgr) GetBufioReader() *bufio.Reader {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	return fMgr.fileBufRdr
}

// GetBufioWriter
//
// Returns a pointer to the internal bufio.Writer,
// 'FileMgr.fileBufWriter'. This pointer is initialized
// when the file is opened for Write or Read-Write
// operations.
//
// Be advised that if the file identified by the current
// FileMgr instance has not been opened, the returned
// pointer to 'fMgr.fileBufWriter' may be nil.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*bufio.Writer
//
//		A pointer to the internal bufio.Writer,
//		FileMgr.fileBufWriter.
//
//		If the file identified by the current FileMgr
//		instance has not been opened, the returned
//		pointer to 'FileMgr.fileBufWriter' may be nil.
func (fMgr *FileMgr) GetBufioWriter() *bufio.Writer {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	return fMgr.fileBufWriter
}

// GetDirMgr
//
// Returns a deep copy of the Directory Manager (DirMgr)
// for this File Manager (FileMgr) instance. This
// Directory Manager contains information on the file
// path or directory in which the file identified by the
// current File Manager instance resides.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	DirMgr
//
//		Returns a deep copy of the Directory Manager
//		(DirMgr) for the current instance of File Manager
//		(FileMgr).
func (fMgr *FileMgr) GetDirMgr() DirMgr {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	return fMgr.dMgr.CopyOut()
}

// GetFileBytesWritten
//
// Returns the sum of private member variables,
//
//	'FileMgr.buffBytesWritten' +
//			'FileMgr.fileBytesWritten'
//
// These variables record the number of bytes written to
// the FileMgr's target file since it was opened with
// 'Write' or 'Read-Write' permissions.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	uint64
//
//		Returns the total number of bytes written to the
//		FileMgr's target file since it was opened with
//		'Write' or 'Read-Write' permissions.
func (fMgr *FileMgr) GetFileBytesWritten() uint64 {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	return fMgr.buffBytesWritten + fMgr.fileBytesWritten
}

// GetFileExt
//
// Returns a string containing the File Extension for this
// File Manager (FileMgr) instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The returned file extension will contain the preceding
// dot separator.
//
// Examples:
//
//	File Name Plus Extension: "newerFileForTest_01.txt"
//	Returned File Extension: ".txt"
//
//	File Name Plus Extension: "newerFileForTest_01"
//	Returned File Extension: ""
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		Returns a string containing the File Extension
//		for the current File Manager (FileMgr) instance.
func (fMgr *FileMgr) GetFileExt() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var fileExt string

	if fMgr.isInitialized == true {
		fileExt = fMgr.fileExt
	}

	return fileExt
}

// GetFileInfo
//
// Wrapper function for os.Stat(). This method can be
// used to return FileInfo data on the specific file
// identified by the current instance of
// FileMgr.absolutePathFileName.
//
// An error will be triggered if the file path does NOT
// exist!
//
//	 type FileInfo interface {
//		 Name() string       // base name of the file
//		 Size() int64        // length in bytes for regular files; system-dependent for others
//		 Mode() FileMode     // file mode bits
//		 ModTime() time.Time // modification time
//		 IsDir() bool        // abbreviation for Mode().IsDir()
//		 Sys() interface{}   // underlying data source (can return nil)
//	 }
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	fInfo						os.FileInfo
//
//		Returns FileInfo data on the specific file
//		identified by FileMgr.absolutePathFileName.
//
//	 	type FileInfo interface {
//			 Name() string       // base name of the file
//			 Size() int64        // length in bytes for regular files; system-dependent for others
//			 Mode() FileMode     // file mode bits
//			 ModTime() time.Time // modification time
//			 IsDir() bool        // abbreviation for Mode().IsDir()
//			 Sys() interface{}   // underlying data source (can return nil)
//	 	}
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) GetFileInfo(
	errorPrefix interface{}) (
	fInfo os.FileInfo,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	fInfo = nil

	filePathDoesExist := false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.GetFileInfo()",
		"")

	if err != nil {
		return fInfo, err
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(fMgr,
		PreProcPathCode.None(),
		ePrefix.XCpy("filePathDoesExist<-"),
		"fMgr.absolutePathFileName")

	if err != nil {
		return fInfo, err
	}

	if filePathDoesExist {

		fInfo = fMgr.actualFileInfo.GetOriginalFileInfo()

	} else {

		err =
			fmt.Errorf("%v\n"+
				"The current FileMgr file DOES NOT EXIST!\n"+
				"File='%v'\n",
				ePrefix.String(),
				fMgr.absolutePathFileName)
	}

	return fInfo, err
}

// GetFileInfoPlus
//
// Returns a FileInfoPlus instance containing os.FileInfo
// and other data on the current FileManager instance.
//
// An error will be triggered if the file path does NOT
// exist!
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	fileInfoPlus				FileInfoPlus
//
//		This returned FileInfoPlus instance contains
//		os.FileInfo and other data on the current
//		FileManager instance.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) GetFileInfoPlus(
	errorPrefix interface{}) (
	fileInfoPlus FileInfoPlus,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.GetFileInfoPlus()",
		"")

	if err != nil {
		return fileInfoPlus, err
	}

	filePathDoesExist := false

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if err != nil {

		fileInfoPlus = FileInfoPlus{}

	} else if filePathDoesExist {

		fileInfoPlus = fMgr.actualFileInfo.CopyOut()

	} else {

		fileInfoPlus = FileInfoPlus{}
		err = fmt.Errorf("%v\n"+
			"Error: File Manager file DOES NOT EXIST!\n"+
			"File='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	return fileInfoPlus, err
}

// GetFileModTime
//
// Returns the time of the last file modification for the
// current file as a type, 'time.Time'.
//
// If the file does NOT exist, an error is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	time.Time
//
//		Returns the time of the last file modification
//		for the current file.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) GetFileModTime(
	errorPrefix interface{}) (
	time.Time,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	filePathDoesExist := false

	var ePrefix *ePref.ErrPrefixDto

	modTime := time.Time{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.GetFileModTime()",
		"")

	if err != nil {
		return modTime, err
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if err != nil {
		return modTime, err
	}

	if filePathDoesExist {
		modTime = fMgr.actualFileInfo.ModTime()
	}

	if !filePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Current FileMgr file DOES NOT EXIST!\n"+
			"FileMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	return modTime, err
}

// GetFileModTimeStr
//
// Returns as a string the time of the last file
// modification for the file identified by the current
// instance of FileMgr.
//
// If the file does NOT exist, an error is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	timeFormat					string
//
//		A format string used to format the last
//		modification time for the file identified by the
//		current File Manager (FileMgr) instance.
//
//		If the string is empty ("") or if the time format
//		is invalid, the method will automatically format
//		the time using the default format,
//		"2006-01-02 15:04:05".
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	string
//
//		The time at which the current file was last
//		modified formatted as a string. The time format
//		is determined by input parameter 'timeFormat'.
//		If 'timeFormat' is empty or if 'timeFormat' is
//		an invalid format, the default format
//		"2006-01-02 15:04:05" will be substituted.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Note:
//		An error will be returned if the file identified
//		by the current File Manager instance does NOT
//		exist.
func (fMgr *FileMgr) GetFileModTimeStr(
	timeFormat string,
	errorPrefix interface{}) (
	string,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error
	filePathDoesExist := false
	modTime := time.Time{}
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.GetFileModTimeStr()",
		"")

	if err != nil {
		return "", err
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix.XCpy(
				"filePathDoesExist<-fMgr"),
			"fMgr.absolutePathFileName")

	if err != nil {
		return "", err
	}

	if filePathDoesExist {
		modTime = fMgr.actualFileInfo.ModTime()
	}

	defaultFmt := "2006-01-02 15:04:05"

	if !filePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Current FileMgr file DOES NOT EXIST!\n"+
			"FileMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)

		return "", err
	}

	if timeFormat == "" {
		timeFormat = defaultFmt
	}

	tStr := modTime.Format(timeFormat)

	if tStr == timeFormat {
		tStr = modTime.Format(defaultFmt)
	}

	return tStr, err
}

// GetFileName
//
// Returns the file name for the current File Manager
// (FileMgr) instance.
//
// Note that the file extension is NOT included. Only
// the file name is returned.
//
//	Examples:
//
//	File Name Plus Extension: "newerFileForTest_01.txt"
//	Returned File Name: "newerFileForTest_01"
//
//	File Name Plus Extension: "newerFileForTest_01"
//	Returned File Name: "newerFileForTest_01"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		Returns the file name for the current File
//		Manager (FileMgr) instance.
//
//		Note that the file extension is NOT included.
//		Only the file name is returned.
func (fMgr *FileMgr) GetFileName() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var fileName string

	if fMgr.isInitialized == true {
		fileName = fMgr.fileName
	}

	return fileName
}

// GetFileNameExt
//
// Returns a string containing the combination of file
// name and file extension configured for this File
// Manager instance
//
// Example:
//
//	         File Name Plus Extension: "newerFileForTest_01.txt"
//	Returned File Name Plus Extension: "newerFileForTest_01.txt"
//
//	         File Name Plus Extension: "newerFileForTest_01"
//	Returned File Name Plus Extension: "newerFileForTest_01"
func (fMgr *FileMgr) GetFileNameExt() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var fileNameExt string

	if fMgr.isInitialized == true {

		fileNameExt = fMgr.fileNameExt

	}

	return fileNameExt
}

// GetFilePermissionConfig
//
// Returns a FilePermissionConfig instance encapsulating
// all the permission information associated with the
// current FileMgr file.
//
// If the file does NOT exist, this method will return an
// error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	fPerm						FilePermissionConfig
//
//		Returns a FilePermissionConfig instance
//		encapsulating all the permission information
//		associated with the current FileMgr file.
//
//		See the source code documentation for method
//		'FilePermissionConfig.New()' for an explanation
//		of permission codes.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) GetFilePermissionConfig(
	errorPrefix interface{}) (
	fPerm FilePermissionConfig,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileMgr.GetFilePermissionConfig()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return fPerm, err
	}

	var filePathDoesExist bool

	filePathDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix.XCpy("filePathDoesExist<-"),
			"FileMgr.absolutePathFileName")

	if err != nil {
		return fPerm, err
	}

	if !filePathDoesExist {

		err =
			fmt.Errorf("%v\n"+
				"The current FileMgr file DOES NOT EXIST!\n",
				ePrefix.String())

		return fPerm, err
	}

	var err2 error

	fPerm, err2 = new(FilePermissionConfig).
		NewByFileMode(
			fMgr.actualFileInfo.Mode(),
			ePrefix)

	if err2 != nil {

		err =
			fmt.Errorf("%v\n"+
				"\n%v\n",
				funcName,
				err2.Error())
	}

	return fPerm, err
}

// GetFilePermissionTextCodes
//
// If the current FileMgr file exists on disk, this
// method will return the File Permission Codes,
// otherwise known as the unix permission bits, in the
// form of a 10-character string.
//
// If the file does NOT exist, this method will return
// an error.
//
// Examples:
//
//	Permission
//	  Text        Octal            File Access
//	  Codes      Notation          Permission Descriptions
//
//	----------     0000             File - no permissions
//	-rwx------     0700             File - read, write, & execute only for owner
//	-rwxrwx---     0770             File - read, write, & execute for owner and group
//	-rwxrwxrwx     0777             File - read, write, & execute for owner, group and others
//	---x--x--x     0111             File - execute
//	--w--w--w-     0222             File - write only
//	--wx-wx-wx     0333             File - write & execute
//	-r--r--r--     0444             File - read only
//	-r-xr-xr-x     0555             File - read & execute
//	-rw-rw-rw-     0666             File - read & write
//	-rwxr-----     0740             File - Owner can read, write, & execute. Group can only read;
//	                                File - others have no permissions
//	drwxrwxrwx     20000000777      File - Directory - read, write, & execute for owner, group and others
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		Returns the permission bits for the current
//		FileMgr file as string of text.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) GetFilePermissionTextCodes(
	errorPrefix interface{}) (
	string,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	permissionText := ""

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr.GetFilePermissionTextCodes()",
		"")

	if err != nil {
		return permissionText, err
	}

	filePathDoesExist := false

	fPerm := FilePermissionConfig{}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix.XCpy("filePathDoesExist<-"),
		"FileMgr")

	if err != nil {
		return permissionText, err
	}

	if !filePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: The current (FileMgr) path/file name DOES NOT EXIST!\n"+
			"Therefore no file permissions are available.\n"+
			"path/file name='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)

		return permissionText, err
	}

	fPerm,
		err = new(FilePermissionConfig).
		NewByFileMode(
			fMgr.actualFileInfo.Mode(),
			ePrefix)

	if err != nil {
		return permissionText, err
	}

	permissionText,
		err = fPerm.
		GetPermissionTextCode(ePrefix.XCpy(
			"permissionText-fPerm"))

	return permissionText, err
}

// GetFilePtr - will return the internal *os.File pointer
// for this File Manager instance. Depending on circumstances,
// this pointer may be nil.
func (fMgr *FileMgr) GetFilePtr() *os.File {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var filePtr *os.File

	filePtr = fMgr.filePtr

	return filePtr
}

// GetFileSize
//
// Returns os.FileInfo.Size() length in bytes for regular
// files; system-dependent for others.
//
// If the File Manager file does NOT exist, or if there is
// a file-path error, the value returned is -1.
func (fMgr *FileMgr) GetFileSize() int64 {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	filePathDoesExist := false
	var err error
	var fileSize int64

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		nil,
		"FileMgr")

	if err != nil || !filePathDoesExist {

		fileSize = -1

		return fileSize
	}

	fileSize = fMgr.actualFileInfo.Size()

	return fileSize
}

// GetOriginalPathFileName - Returns the path and file name
// used originally to configure this File Manager object.
//
// Note: The original path and file name will be adjusted
// to reflect the path operators used in the operating system
// for the host computer.
func (fMgr *FileMgr) GetOriginalPathFileName() string {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	originalPathName := ""

	originalPathName = fMgr.originalPathFileName

	return originalPathName
}

// GetReaderBufferSize - Returns the size for the internal
// Bufio Reader's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Read' Operation.
func (fMgr *FileMgr) GetReaderBufferSize() int {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	readerBuffSize := 0

	if fMgr.fileBufRdr != nil {
		fMgr.fileRdrBufSize = fMgr.fileBufRdr.Size()
		readerBuffSize = fMgr.fileRdrBufSize
	} else {
		readerBuffSize = fMgr.fileRdrBufSize
	}

	return readerBuffSize
}

// GetWriterBufferSize - Returns the size for the internal
// Bufio Writer's buffer. If the value is less than 1 it means
// that the buffer will be set to the default size at the next
// 'Write' Operation.
func (fMgr *FileMgr) GetWriterBufferSize() int {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	writerBuffSize := 0

	if fMgr.fileBufWriter != nil {
		fMgr.fileWriterBufSize = fMgr.fileBufWriter.Size()
		writerBuffSize = fMgr.fileWriterBufSize
	} else {
		writerBuffSize = fMgr.fileWriterBufSize
	}

	return writerBuffSize
}

// IsAbsolutePathFileNamePopulated - Returns a boolean value
// indicating whether absolute path and file name is
// initialized and populated.
func (fMgr *FileMgr) IsAbsolutePathFileNamePopulated() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	isAbsolutePathFileNamePopulated := false

	_,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		nil,
		"FileMgr")

	if err != nil || len(fMgr.absolutePathFileName) == 0 {

		fMgr.isAbsolutePathFileNamePopulated = false

		isAbsolutePathFileNamePopulated = false

	} else {

		fMgr.isAbsolutePathFileNamePopulated = true

		isAbsolutePathFileNamePopulated = true
	}

	return isAbsolutePathFileNamePopulated
}

// IsFileExtPopulated
//
// Returns a boolean value indicating whether the File
// Extension for this File Manager instance is populated.
func (fMgr *FileMgr) IsFileExtPopulated() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error
	isFileExtPopulated := false

	_,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			nil,
			"FileMgr.absolutePathFileName")

	if err != nil ||
		len(fMgr.fileExt) == 0 {
		fMgr.isFileExtPopulated = false
		isFileExtPopulated = false
	} else {
		fMgr.isFileExtPopulated = true
		isFileExtPopulated = true
	}

	return isFileExtPopulated
}

// IsFileMgrValid
//
// Analyzes the current FileMgr object. If the current
// FileMgr object is INVALID, an error is returned.
//
// If the current FileMgr is VALID, this method returns
// 'nil'
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If any of the internal member data variables
//		contained in the current instance of FileMgr are
//	 	found to be invalid, this method will return an
//	 	error configured with an appropriate message
//	 	identifying the invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any internal member data values are found to be
//		invalid, the returned error Type will encapsulate
//		an appropriate error message. This returned error
//		message will incorporate the method chain and text
//		passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the
//		beginning of the error message.
func (fMgr *FileMgr) IsFileMgrValid(
	errorPrefix interface{}) (err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileMgr." +
		"IsFileMgrValid()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if !fMgr.isInitialized {

		err = fmt.Errorf("%v\n"+
			" Error: The current FileMgr data structure is NOT initialized.\n",
			ePrefix.String())

		return err
	}

	if fMgr.absolutePathFileName == "" {

		fMgr.isAbsolutePathFileNamePopulated = false

		err = fmt.Errorf("%v\n"+
			" Error: FileMgr absolutePathFileName is EMPTY!\n",
			ePrefix.String())

		return err

	}

	var err2 error

	err2 = fMgr.dMgr.IsDirMgrValid(ePrefix.String())

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"FileMgr Directory Manager INVALID\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return err
	}

	_,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	_ = fMgr.dMgr.DoesPathExist()
	_ = fMgr.dMgr.DoesAbsolutePathExist()

	return err
}

// IsFileNameExtPopulated
//
// Returns a boolean value indicating whether both the
// File Name and Extension for this File Manager
// instance have been populated.
//
// If either the File Name or the File Extension is blank
// (empty), this method returns false.
//
// Both the File Name AND the File Extension must be
// populated before this method returns 'true'.
func (fMgr *FileMgr) IsFileNameExtPopulated() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error
	isFileNameExtPopulated := false

	_,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		nil,
		"FileMgr.absolutePathFileName")

	if err != nil {
		return false
	}

	if len(fMgr.fileExt) > 0 &&
		len(fMgr.fileName) > 0 {

		fMgr.isFileNameExtPopulated = true
		isFileNameExtPopulated = true
	} else {

		fMgr.isFileNameExtPopulated = false
		isFileNameExtPopulated = false
	}

	return isFileNameExtPopulated
}

// IsFileNamePopulated
//
// Returns a boolean value indicating whether the file
// name for this File Manager object is populated.
func (fMgr *FileMgr) IsFileNamePopulated() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	_,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			nil,
			"FileMgr.absolutePathFileName")

	if err != nil {

		return false
	}

	var isFileNamePopulated bool

	if len(fMgr.fileName) == 0 {

		fMgr.isFileNamePopulated = false
		isFileNamePopulated = false

	} else {

		fMgr.isFileNamePopulated = true
		isFileNamePopulated = true

	}

	return isFileNamePopulated
}

// IsFilePointerOpen - Returns a boolean value indicating
// whether the File Pointer (*os.File) for this File Manager
// instance is open, or not.
func (fMgr *FileMgr) IsFilePointerOpen() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var isFilePtrOpen bool

	if fMgr.filePtr == nil {
		fMgr.isFilePtrOpen = false
		isFilePtrOpen = false
	} else {
		fMgr.isFilePtrOpen = true
		isFilePtrOpen = true
	}

	return isFilePtrOpen
}

// IsInitialized - Returns a boolean indicating whether the FileMgr
// object is properly initialized.
func (fMgr *FileMgr) IsInitialized() bool {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	isInitialized := false

	isInitialized = fMgr.isInitialized

	return isInitialized
}

// MoveFileToFileMgr - This method will move the file identified
// by the current FileMgr to a new path and file name specified
// by the input parameter 'destinationFMgr' which is of type
// 'FileMgr'.
//
// IMPORTANT:
//
// The file identified by the current FileMgr object will be
// DELETED!
//
// The new file is defined by input parameter 'destinationFMgr'.
//
// If the input parameter 'destinationFMgr' contains a directory path
// which currently does NOT exist, it will be created.
func (fMgr *FileMgr) MoveFileToFileMgr(destinationFMgr FileMgr) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.MoveFileToFileMgr() "
	var err error

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.moveFile(fMgr, &destinationFMgr, ePrefix)

	return err
}

// MoveFileToFileStr - This method will move the file identified
// by the current FileMgr to a new path and file name contained
// in the input parameter 'pathFileNameExt' which is of type
// string.
//
// IMPORTANT:
//
// The file identified by the current FileMgr object will be
// DELETED!
//
// The new file is defined by input parameter 'pathFileNameExt'.
//
// If the input parameter 'pathFileNameExt' contains a directory path
// which currently does NOT exist, it will be created.
func (fMgr *FileMgr) MoveFileToFileStr(
	pathFileNameExt string,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "FileMgr.MoveFileToFileStr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return FileMgr{}, err
	}

	fMgrHlpr := fileMgrHelper{}

	var targetFMgr FileMgr

	var isEmpty bool

	isEmpty,
		err = fMgrHlpr.
		setFileMgrPathFileName(
			&targetFMgr,
			pathFileNameExt,
			ePrefix.XCpy(
				"targetFMgr<-pathFileNameExt"))

	if err != nil {

		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'pathFileNameExt' is invalid!\n"+
				"pathFileName='%v'\n"+
				"Error= \n%v\n",
				funcName,
				pathFileNameExt,
				err.Error())
	}

	if isEmpty {

		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'pathFileNameExt' is invalid!\n"+
				"pathFileName='%v'\n",
				ePrefix.String(),
				pathFileNameExt)

	}

	err = fMgrHlpr.moveFile(
		fMgr,
		&targetFMgr,
		ePrefix.XCpy(
			"targetFMgr<-fMgr"))

	return targetFMgr, err
}

// MoveFileToNewDir - This method will move the current file
// identified by this FileMgr object to a new path designated
// by input parameter string, 'dirPath'.
//
// IMPORTANT:
//
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned in the return
// parameter 'newFMgr'.
//
// If input parameter 'dirPth' contains a directory path which does not
// currently exist, it will be created.
func (fMgr *FileMgr) MoveFileToNewDir(
	dirPath string,
	errorPrefix interface{}) (FileMgr, error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"MoveFileToNewDir()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	dirMgr, err := new(DirMgr).New(dirPath)

	if err != nil {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error returned by DirMgr{}.New(dirPath).\n"+
				"dirPath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirPath,
				err.Error())
	}

	var targetFMgr FileMgr

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&targetFMgr,
			&dirMgr,
			fMgr.fileNameExt,
			ePrefix.XCpy(
				"fMgrDest<-dir-fMgr.fileNameExt"))

	if err != nil {

		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'dirMgr' and fMgr.fileNameExt\n"+
				"produced an error!\n"+
				"dirMgr='%v'\n"+
				"fMgr.fileNameExt='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirMgr.absolutePath,
				fMgr.fileNameExt,
				err.Error())
	}

	if isEmpty {

		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'dirMgr' and fMgr.fileNameExt\n"+
				"produced an empty result!\n"+
				"dirMgr='%v'\n"+
				"fMgr.fileNameExt='%v'\n",
				ePrefix.String(),
				dirMgr.absolutePath,
				fMgr.fileNameExt)
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.moveFile(fMgr, &targetFMgr, ePrefix)

	return targetFMgr, err
}

// MoveFileToNewDirMgr - This method will move the file identified
// by the current FileMgr to a new path contained in the input parameter
// 'dMgr'.
//
// IMPORTANT:
//
// The current file identified by the current FileMgr object will
// be DELETED!
//
// The new file located in the new directory will be returned as a type
// 'FileMgr'.
//
// If the input parameter 'dMgr' does not contain a path which currently
// exists, it will be created.
func (fMgr *FileMgr) MoveFileToNewDirMgr(
	dMgr DirMgr,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileMgr.MoveFileToNewDirMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return FileMgr{}, err
	}

	err2 := dMgr.IsDirMgrValid("")

	if err2 != nil {
		return FileMgr{}, fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is INVALID!\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	targetFile := dMgr.GetAbsolutePathWithSeparator() + fMgr.fileNameExt

	fMgrHlpr := fileMgrHelper{}

	var isEmpty bool

	var targetFMgr FileMgr

	isEmpty,
		err2 = fMgrHlpr.
		setFileMgrPathFileName(
			&targetFMgr,
			targetFile,
			ePrefix.XCpy(
				"targetFMgr<-targetFile"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFile' is invalid!\n"+
			"targetFile='%v'\n"+
			"Error= \n%v\n",
			funcName,
			targetFile,
			err2.Error())

		return FileMgr{}, err
	}

	if isEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFile' is invalid!\n"+
			"targetFile='%v'\n",
			ePrefix.String(),
			targetFile)

		return FileMgr{}, err

	}

	err = fMgrHlpr.moveFile(
		fMgr,
		&targetFMgr,
		ePrefix.XCpy(
			"targetFMgr<-fMgr"))

	return targetFMgr, err
}

// New
//
// Creates a new File Manager instance of File Manager
// ('FileMgr'). This method receives an input parameter
// of type string and parses out the path, file name and
// file extension. This file data is returned in the data
// fields of the new FileMgr object.
//
// This method is identical to method
// 'NewFromPathFileNameExtStr()'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This parameter must consist of a valid path, file
//		name and file extension. The file need not exist.
//
//		Failure to provide a properly formatted path
//		and/or file name will result in an error.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileMgr
//
//		If the method completes successfully, a valid
//		FileMgr instance is returned containing
//		information on the file specified by the input
//		parameter 'pathFileNameExt'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	fmgr := FileMgr{}.New("../mydirectory/fileName.ext")
func (fMgr *FileMgr) New(
	pathFileNameExt string,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fMgr2 := FileMgr{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"New()",
		"")

	if err != nil {
		return fMgr2, err
	}

	fMgrHlpr := fileMgrHelper{}

	var isEmpty bool

	isEmpty,
		err = fMgrHlpr.
		setFileMgrPathFileName(
			&fMgr2,
			pathFileNameExt,
			ePrefix.XCpy(
				"fMgr2<-pathFileNameExt"))

	if err != nil {

		return FileMgr{}, err
	}

	if isEmpty {

		fMgr2 = FileMgr{}

		return FileMgr{}, fmt.Errorf("%v\n"+
			"Error: The FileMgr instance generated by input parameter 'pathFileNameExt' is Empty!\n"+
			"pathFileNameExt='%v'\n",
			ePrefix.String(),
			pathFileNameExt)

	}

	return fMgr2, err
}

// NewFromDirMgrFileNameExt
//
// This method is designed to create a new File Manager
// (FileMgr) object from two input parameters.
//
// Input parameter 'dirMgr' is a valid and correctly
// populated Directory Manager object containing the
// directory path or folder.
//
// Input parameter 'fileNameExt' is a string containing
// the file name and the file extension.
//
// 'dirMgr' and 'fileNameExt' will be combined to create
// a new File Manager (FileMgr) object containing a
// properly configured path, file name and file
// extension.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						DirMgr
//
//		A valid and properly initialized Directory
//		Manager (type: DirMgr). This object contains the
//		file path which will be combined with the second
//		input parameter, 'fileNameExt' to create the new
//		File Manager (type: FileMgr).
//
//	fileNameExt					string
//
//		A string containing the file name and file
//		extension which will be used to create the new
//		File Manager (type: FileMgr).
//
//		If this parameter is submitted as an empty string
//		or if the string consists of all white space, an
//		error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileMgr
//
//		If the method completes successfully, a valid
//		FileMgr instance is returned containing
//		information on the file specified in the two
//		input parameters 'dirMgr' and 'fileNameExt'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	dMgr, err := DirMgr{}.New("D:\\TestDir")
//
//	if err != nil {
//	 fmt.Printf("Error='%v' \n", err.Error())
//	 return
//	}
//
//	fileNameExt := "yourFile.txt"
//
//	fMgr, err := FileMgr{}.NewFromDirMgrFileNameExt(dMgr, fileNameExt)
func (fMgr *FileMgr) NewFromDirMgrFileNameExt(
	dMgr DirMgr,
	fileNameExt string,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fMgr2 := FileMgr{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"NewFromDirMgrFileNameExt()",
		"")

	if err != nil {
		return fMgr2, err
	}

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgr2,
			&dMgr,
			fileNameExt,
			ePrefix)

	if isEmpty {

		fMgr2 = FileMgr{}

		if err == nil {
			return FileMgr{},
				fmt.Errorf("%v\n"+
					"Error: The FileMgr instance generated by input parameters 'dMgr' and "+
					"'fileNameExt' is Empty!\n"+
					"dMgr='%v'\n"+
					"fileNameExt='%v'\n",
					ePrefix.String(),
					dMgr.absolutePath,
					fileNameExt)
		}
	}

	return fMgr2, err
}

// NewFromDirStrFileNameStr
//
// Creates a new file manager object (FileMgr) from a
// directory string and a File Name and Extension string
// passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirStr						string
//
//		A string containing the directory or file path
//		which will be combined with the second input
//	 	parameter, 'fileNameExt' to create the new File
//	 	Manager (type: FileMgr).
//
//	fileNameExtStr				string
//
//		A string containing the file name and file
//		extension which will be used to create the new
//		File Manager (type: FileMgr).
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileMgr
//
//		If the method completes successfully, a valid
//		FileMgr instance is returned containing
//		information on the file specified in the two
//		input parameters 'dirStr' and 'fileNameExtStr'.
//
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	dirStr := "D:\\TestDir"
//
//	fileNameExtStr := "yourFile.txt"
//
//	fMgr, err := FileMgr{}.NewFromDirStrFileNameStr(dirStr, fileNameExtStr)
func (fMgr *FileMgr) NewFromDirStrFileNameStr(
	dirStr string,
	fileNameExtStr string,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fMgr2 := FileMgr{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"NewFromDirStrFileNameStr()",
		"")

	if err != nil {
		return fMgr2, err
	}

	dirMgr, err := new(DirMgr).New(dirStr)

	if err != nil {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error returned by DirMgr{}.NewFromPathFileNameExtStr(dirMgr).\n"+
				"dirStr='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirStr,
				err.Error())
	}

	isEmpty, err := new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgr2,
			&dirMgr,
			fileNameExtStr,
			ePrefix)

	if isEmpty {

		fMgr2 = FileMgr{}

		if err == nil {
			return FileMgr{},
				fmt.Errorf("%v\n"+
					"Error: The FileMgr instance generated by input parameters 'dirStr' and "+
					"'fileNameExtStr' is Empty!\n"+
					"dirStr='%v'\n"+
					"fileNameExtStr='%v'\n",
					ePrefix.String(),
					dirStr,
					fileNameExtStr)
		}
	}

	return fMgr2, err
}

// NewFromFileInfo
//
// Creates and returns a new FileMgr object based on
// input from a directory path string and an os.FileInfo
// or FileInfoPlus object.
//
// ----------------------------------------------------------------
//
// Input Parameters:
//
//	dirPathStr					string
//
//		The directory path. NOTE: This does NOT contain
//		the file name.
//
//	fileInfo					os.FileInfo
//
//		A valid and populated FileInfo structure
//		containing the file name.
//
//		Note:
//
//			An instance of FileInfoPlus may be submitted
//			for this parameter because FileInfoPlus
//			implements the os.FileInfo interface.
//
// ----------------------------------------------------------------
//
// Return Values:
//
//	FileMgr
//
//		If the method completes successfully, a valid
//		FileMgr instance is returned containing
//		information on the file specified by input
//		parameter 'info'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) NewFromFileInfo(
	dirPathStr string,
	fileInfo os.FileInfo,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fMgr2 := FileMgr{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"NewFromFileInfo()",
		"")

	if err != nil {
		return fMgr2, err
	}

	if fileInfo == nil {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fileInfo' is 'nil' and INVALID!\n",
				ePrefix)
	}

	fileName := fileInfo.Name()

	dirMgr, err := new(DirMgr).New(dirPathStr)

	if err != nil {
		return FileMgr{},
			fmt.Errorf("Input Parameter 'dirPathStr' is INVALID! DirMgr{}.New(dirPathStr) "+
				"Error='%v' ", err.Error())
	}

	isEmpty, err := new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgr2,
			&dirMgr,
			fileName,
			ePrefix)

	if isEmpty {

		fMgr2 = FileMgr{}

		if err == nil {
			return FileMgr{},
				fmt.Errorf("%v\n"+
					"Error: The FileMgr instance generated by input parameters 'dirPathStr' and "+
					"'fileInfo' is Empty!\n"+
					"dirPathStr='%v'\n"+
					"fileInfo='%v'\n",
					ePrefix.String(),
					dirPathStr,
					fileName)
		}
	}

	return fMgr2, err
}

// NewFromPathFileNameExtStr
//
// Creates a new File Manager ('FileMgr') instance. This
// method receives an input parameter of type string
// and parses out the path, file name and file extension.
//
// The file data is returned in the data fields of the
// new FileMgr object.
//
// This method is identical to method 'New()'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		Must consist of a valid path, file name and file
//		extension. The file need not exist.
//
//		Failure to provide a properly formatted path
//		and/or file name will result in an error.
//
//		Note: The file extension is optional.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//	fmgr := FileMgr{}.NewFromPathFileNameExtStr("../somedirectory/fileName.ext")
func (fMgr *FileMgr) NewFromPathFileNameExtStr(
	pathFileNameExt string,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"NewFromPathFileNameExtStr()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	fMgr2 := FileMgr{}

	isEmpty, err := new(fileMgrHelper).
		setFileMgrPathFileName(
			&fMgr2,
			pathFileNameExt,
			ePrefix)

	if isEmpty {

		fMgr2 = FileMgr{}

		if err == nil {
			return FileMgr{},
				fmt.Errorf("%v\n"+
					"Error: The FileMgr instance generated by input parameter 'pathFileNameExt' is Empty!\n"+
					"pathFileNameExt='%v'\n",
					ePrefix.String(),
					pathFileNameExt)
		}
	}

	return fMgr2, err
}

// OpenThisFile - Opens the file identified by the current FileMgr object
// using the file open parameters and file permission parameters contained
// in the 'FileAccessControl' instance passed as 'fileAccessCtrl'.
//
// Note: If the FileMgr directory path does not exist, this method will
// create that directory path.
func (fMgr *FileMgr) OpenThisFile(fileAccessCtrl FileAccessControl) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.OpenThisFile() "
	var err error

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.openFile(
		fMgr,
		fileAccessCtrl,
		true,
		true,
		ePrefix)

	return err
}

// OpenThisFileReadOnly - Opens the file identified by the current
// FileMgr object as a 'Read-Only' File. Subsequent operations may
// read from this file but may NOT write to this file.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for reading only.
//
// If FileMgr.absolutePathFileName does not exist, an error will be
// returned.
//
// If successful, the FileMode is set to "-r--r--r--" and the permission
// Mode is set to '0444'.
//
// Note: If the 'FileMgr' directory path or file do not exist, this
// method will return an error.
func (fMgr *FileMgr) OpenThisFileReadOnly() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.OpenThisFileReadOnly() "

	readOnlyAccessCtrl, err :=
		new(FileAccessControl).NewReadOnlyAccess(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.openFile(
		fMgr,
		readOnlyAccessCtrl,
		false,
		false,
		ePrefix)

	return err
}

// OpenThisFileWriteOnly - Opens the current file for 'WriteOnly'
// operations.  If successful, this method will use FileMgr.absolutePathFileName
// to open an *os.File or File Pointer.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for writing only. If FileMgr.absolutePathFileName
// does not exist, it will be created. The FileMode is set to "--w--w--w-" and
// the permission Mode is set to '0222'.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will create them.
func (fMgr *FileMgr) OpenThisFileWriteOnly() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error
	ePrefix := "FileMgr.OpenThisFileWriteOnly() "
	fMgrHlpr := fileMgrHelper{}

	writeOnlyAccessCtrl, err :=
		new(FileAccessControl).NewWriteOnlyAccess(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
	}

	err = fMgrHlpr.openFile(
		fMgr,
		writeOnlyAccessCtrl,
		true,
		true,
		ePrefix)

	return err
}

// OpenThisFileWriteOnlyAppend - Opens the current file for 'Write Only'
// with an 'Append' mode. All bytes written to the file will be written
// at the end of the current file and none of the file's original content
// will be overwritten.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will create both the path and file.
func (fMgr *FileMgr) OpenThisFileWriteOnlyAppend() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.OpenThisFileWriteOnlyAppend() "

	writeOnlyFileAccessCfg, err :=
		new(FileAccessControl).NewWriteOnlyAppendAccess(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.openFile(
		fMgr,
		writeOnlyFileAccessCfg,
		true,
		true,
		ePrefix)

	return err
}

// OpenThisFileWriteOnlyTruncate - Opens the current file for 'Write Only'
// with an 'Truncate' mode. This means that if the file previously exists,
// all the existing file content will be deleted before the 'write'
// operation will begin.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will create both the path and file.
func (fMgr *FileMgr) OpenThisFileWriteOnlyTruncate() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.OpenThisFileWriteOnlyTruncate() "

	writeOnlyTruncateAccessCfg, err :=
		new(FileAccessControl).NewWriteOnlyTruncateAccess(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.openFile(
		fMgr,
		writeOnlyTruncateAccessCfg,
		true,
		true,
		ePrefix)

	return err
}

// OpenThisFileReadWrite - Opens the file identified by the current
// FileMgr object. If successful, this method will use
// FileMgr.absolutePathFileName to open an *os.File or File Pointer.
//
// As the method's name implies, the 'FileMgr.absolutePathFileName'
// will be opened for reading and writing. If FileMgr.absolutePathFileName
// does not exist, it will be created. The FileMode is set to'-rw-rw-rw-' and
// the permission Mode= '0666'.
//
// Note: If the 'FileMgr' directory path and file do not exist, this
// method will create them.
func (fMgr *FileMgr) OpenThisFileReadWrite() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var err error

	ePrefix := "FileMgr.OpenThisFileReadWrite() "

	readWriteAccessCtrl,
		err :=
		new(FileAccessControl).NewReadWriteAccess(
			ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"\n%v\n", err.Error())
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.openFile(
		fMgr,
		readWriteAccessCtrl,
		true,
		true,
		ePrefix)

	return err
}

// ReadAllFile - Reads the file identified by the current FileMgr
// and returns the contents in a byte array.
//
// If no errors are encountered the returned 'error' value is
// nil. This method does not return io.EOF as an error. This is
// because it reads the entire file and therefor no End Of File
// flag is required.
func (fMgr *FileMgr) ReadAllFile() (bytesRead []byte, err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.ReadAllFile() "
	bytesRead = []byte{}
	err = nil
	var err2 error

	readWriteAccessCtrl, err2 :=
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v\n", err2.Error())
		return bytesRead, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.readFileSetup(
		fMgr,
		readWriteAccessCtrl,
		false,
		ePrefix)

	if err != nil {
		return bytesRead, err
	}

	bytesRead, err2 = io.ReadAll(fMgr.filePtr)

	if err2 != nil {
		err =
			fmt.Errorf(ePrefix+
				"Error returned by ioutil.ReadAll(fMgr.filePtr).\n"+
				"fileName='%v'\nErrors='%v'\n",
				fMgr.absolutePathFileName, err2.Error())
	}

	return bytesRead, err
}

// ReadFileBytes - Reads bytes from the file identified by the current FileMgr
// object. Bytes are stored in 'byteBuff', a byte array passed in as an input
// parameter.
//
// If successful, the returned error value is 'nil' or io.EOF.
//
// The returned value 'int' contains the number of bytes read from the current file.
//
// At End of File (EOF), the byte count will be zero and err will be equal to
// 'io.EOF'.
func (fMgr *FileMgr) ReadFileBytes(byteBuff []byte) (bytesRead int, err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.ReadFileBytes() "
	bytesRead = 0
	err = nil
	var err2 error

	readWriteAccessCtrl, err2 :=
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
			"Error='%v'\n", err2.Error())
		return bytesRead, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.readFileSetup(
		fMgr,
		readWriteAccessCtrl,
		false,
		ePrefix)

	if err != nil {
		return bytesRead, err
	}

	bytesRead, err2 = fMgr.fileBufRdr.Read(byteBuff)

	if err2 != nil &&
		err2 == io.EOF {

		err = err2

	} else if err2 != nil {

		err = fmt.Errorf("Error returned by fMgr.fileBufRdr.Read(byteBuff). "+
			"File='%v' Error='%v' ", fMgr.absolutePathFileName, err2.Error())
	}

	return bytesRead, err
}

// ReadFileLine - Effectively, this method reads a file one line
// at a time and returns the line as an array of bytes. The delimiter
// for lines read is specified by input parameter 'delim' of type byte.
//
// This method uses the 'bufio' package.
//
// If End Of File (EOF) is reached, this method returns the 'bytesRead' and
// an error which is equal to 'io.EOF'.
//
// Note that the delimiter is returned as the last character in bytes read.
func (fMgr *FileMgr) ReadFileLine(delim byte) (bytesRead []byte, err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.ReadFileLine() "
	bytesRead = []byte{}
	err = nil
	var err2 error

	readWriteAccessCtrl, err2 :=
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
			"Error='%v'\n", err2.Error())
		return bytesRead, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.readFileSetup(
		fMgr,
		readWriteAccessCtrl,
		false,
		ePrefix)

	if err != nil {
		return bytesRead, err
	}

	bytesRead, err2 = fMgr.fileBufRdr.ReadBytes(delim)

	if err2 != nil &&
		err2 == io.EOF {
		err = err2

	} else if err2 != nil {

		err = fmt.Errorf(ePrefix+
			"\nError returned from fMgr.fileBufRdr.ReadBytes(delim).\n"+
			"Error='%v'\n", err2.Error())
	}

	return bytesRead, err
}

// ReadFileString - Wrapper for bufio.ReadString. ReadFileString reads until the
// first occurrence of 'delim' in the input, returning a string containing the
// data up to and including the delimiter. If ReadFileString encounters an error
// before finding a delimiter, it returns the data read before the error and
// the error itself (often io.EOF). ReadFileString returns err != nil if and
// only if the returned data does not end in 'delim'.
func (fMgr *FileMgr) ReadFileString(delim byte) (stringRead string, err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.ReadFileString() "
	stringRead = ""
	err = nil
	var err2 error

	readWriteAccessCtrl, err2 :=
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by FileAccessControl{}.NewReadWriteAccess().\n"+
			"Error='%v'\n", err2.Error())
		return stringRead, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.readFileSetup(
		fMgr,
		readWriteAccessCtrl,
		false,
		ePrefix)

	if err != nil {
		return stringRead, err
	}

	stringRead, err2 = fMgr.fileBufRdr.ReadString(delim)

	if err2 != nil &&
		err2 == io.EOF {
		err = err2

	} else if err2 != nil {

		err = fmt.Errorf("Error returned from fMgr.fileBufRdr.ReadString(delim). "+
			"Error='%v' ", err2.Error())
	}

	return stringRead, err
}

// ResetFileInfo
//
// Acquires the current os.FileInfo data associated with
// the file identified by the current FileMgr instance.
//
// An error will be triggered if the file path does NOT
// exist!
func (fMgr *FileMgr) ResetFileInfo() error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error
	var filePathDoesExist = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileMgr."+
			"ResetFileInfo()",
		"")

	if err != nil {
		return err
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix.XCpy("filePathDoesExist<-fMgr"),
			"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	if !filePathDoesExist {
		return fmt.Errorf("%v\n"+
			"Current FileMgr file DOES NOT EXIST!\n"+
			"FMgr='%v'",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	return nil
}

// SetReaderBufferSize - Sets the Read Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
func (fMgr *FileMgr) SetReaderBufferSize(readBuffSize int) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	fMgr.fileRdrBufSize = readBuffSize

	return
}

// SetWriterBufferSize - Sets the Write Buffer size in bytes.
// If the value is less than 1, the buffer size will be set
// to the system default size.
func (fMgr *FileMgr) SetWriterBufferSize(writeBuffSize int) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	fMgr.fileWriterBufSize = writeBuffSize

	return
}

// SetFileInfo - Used to initialize the os.FileInfo structure maintained as
// part of the current FileMgr object.
//
// Be careful: Failure of the input parameter to match the current FileMgr file name
// will trigger an error.
func (fMgr *FileMgr) SetFileInfo(info os.FileInfo) error {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.SetFileInfo() "

	if info == nil {
		return errors.New(ePrefix +
			"Error: Input parameter 'info' is 'nil' and INVALID!\n")
	}

	if info.Name() == "" {
		return errors.New(ePrefix +
			"Error: info.Name() is an EMPTY string!\n")
	}

	if info.IsDir() {
		return errors.New(ePrefix +
			"Input parameter info.IsDir()=='true'. This is a Directory NOT A FILE!\n")
	}

	if strings.ToLower(info.Name()) != strings.ToLower(fMgr.fileNameExt) {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'info' does NOT match current FileMgr file name.\n"+
			"FileMgr File Name='%v'\ninfo File Name='%v'\n",
			fMgr.fileNameExt, info.Name())
	}

	fMgr.actualFileInfo =
		new(FileInfoPlus).NewFromFileInfo(info)

	if !fMgr.actualFileInfo.isFInfoInitialized {
		return fmt.Errorf(ePrefix+
			"Error: Failed to initialize fMgr.actualFileInfo object.\n"+
			"info.Name()='%v'", info.Name())
	}

	return nil
}

// SetFileMgrFromDirMgrFileName
//
// Sets the data fields of the current FileMgr object
// based on a DirMgr object and a File Name string which
// are passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						DirMgr
//
//		A valid and properly initialized Directory
//		Manager (type: DirMgr). This object contains the
//		file path which will be combined with the second
//		input parameter, 'fileNameExt' to create the new
//		File Manager (type: FileMgr).
//
//	fileNameExt					string
//
//		A string containing the file name and file
//		extension which will be used to create the new
//		File Manager (type: FileMgr).
//
//		If this parameter is submitted as an empty string
//		or if the string consists of all white space, an
//		error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isEmpty
//
//		If input parameter 'fileNameExt' is submitted as
//		an empty string or if that string consists of all
//		white space, 'isEmpty' will be set to 'true'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgr *FileMgr) SetFileMgrFromDirMgrFileName(
	dMgr DirMgr,
	fileNameExt string,
	errorPrefix interface{}) (
	isEmpty bool,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"SetFileMgrFromDirMgrFileName()",
		"")

	if err != nil {
		return isEmpty, err
	}

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			fMgr,
			&dMgr,
			fileNameExt,
			ePrefix)

	return isEmpty, err
}

// SetFileMgrFromPathFileName - Initializes all the data fields of the
// current FileMgr object based on the path file name string passed to
// this method as an input parameter.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	isEmpty       - This value is set to 'false' if, and only if, all internal
//                  values are set of valid legitimate values.
//
//	error         - If this method completes successfully, the returned error
//	                Type is set equal to 'nil'. If an error condition is encountered,
//	                this method will return an error Type which contains an appropriate
//	                error message.
//

func (fMgr *FileMgr) SetFileMgrFromPathFileName(
	pathFileNameExt string,
	errorPrefix interface{}) (
	isEmpty bool,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	isEmpty = true

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgr."+
			"SetFileMgrFromPathFileName()",
		"")

	if err != nil {
		return isEmpty, err
	}

	fMgrHlpr := fileMgrHelper{}

	isEmpty, err = fMgrHlpr.setFileMgrPathFileName(
		fMgr,
		pathFileNameExt,
		ePrefix)

	return isEmpty, err
}

// WriteBytesToFile a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
//
// If the number of bytes written to the file is less than len(bytes),
// an error will be returned.
//
// This method uses the 'bufio' package.
func (fMgr *FileMgr) WriteBytesToFile(
	bytes []byte) (
	numBytesWritten int,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.WriteBytesToFile() "
	err = nil
	numBytesWritten = 0

	readWriteAccessCtrl, err2 :=
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by FileAccessControl{}."+
			"NewReadWriteAccess()\n."+
			"Error='%v'\n", err2.Error())

		return numBytesWritten, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.writeFileSetup(
		fMgr,
		readWriteAccessCtrl,
		true,
		ePrefix)

	if err != nil {

		return numBytesWritten, err
	}

	numBytesWritten, err2 =
		fMgr.fileBufWriter.Write(bytes)

	if err2 != nil {
		err =
			fmt.Errorf(ePrefix+
				"Error returned from fMgr.fileBufWriter.Write(bytes). Output File='%v'. "+
				"Error='%v'",
				fMgr.absolutePathFileName, err2.Error())

	} else {

		fMgr.buffBytesWritten += uint64(numBytesWritten)

	}

	return numBytesWritten, err
}

// WriteStrToFile - Writes a string to the File identified by
// FileMgr.absolutePathFileName. If the file is not open, this
// method will attempt to open it.
//
// If the number of bytes written to the file is less than len(str),
// an error will be returned.
//
// This method uses the 'bufio' package.
func (fMgr *FileMgr) WriteStrToFile(
	str string) (
	numBytesWritten int,
	err error) {

	if fMgr.lock == nil {
		fMgr.lock = new(sync.Mutex)
	}

	fMgr.lock.Lock()

	defer fMgr.lock.Unlock()

	ePrefix := "FileMgr.WriteStrToFile() "

	numBytesWritten = 0
	err = nil
	var err2 error
	var readWriteAccessCtrl FileAccessControl

	readWriteAccessCtrl, err2 =
		new(FileAccessControl).NewReadWriteAccess(ePrefix)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+
			"Error returned by FileAccessControl{}."+
			"NewReadWriteAccess()\n."+
			"Error='%v'\n", err2.Error())

		return numBytesWritten, err
	}

	fMgrHlpr := fileMgrHelper{}

	err = fMgrHlpr.writeFileSetup(
		fMgr,
		readWriteAccessCtrl,
		true,
		ePrefix)

	if err != nil {
		return numBytesWritten, err
	}

	numBytesWritten, err2 = fMgr.fileBufWriter.WriteString(str)

	if err2 != nil {
		err =
			fmt.Errorf(ePrefix+
				"Error returned from fMgr.filePtr.WriteString(str). Output File='%v'. "+
				"Error='%v'", fMgr.absolutePathFileName, err2.Error())

	} else {
		fMgr.buffBytesWritten += uint64(numBytesWritten)
	}

	return numBytesWritten, err
}
