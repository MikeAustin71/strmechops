package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
	"time"
)

// FileInfoPlus
//
// Conforms to the os.FileInfo interface. This structure
// will store os.FileInfo information plus additional
// information related to a file or directory.
type FileInfoPlus struct {

	// isFInfoInitialized - Not part of FileInfo interface.
	// 'true' = structure fields have been properly initialized
	isFInfoInitialized bool

	// isDirPathInitialized - Not part of FileInfo interface.
	//   'true' = structure field 'dirPath' has been successfully initialized
	isDirPathInitialized bool

	// CreateTimeStamp - Not part of FileInfo interface.
	// Date time at which this instance of Type 'FileInfoPlus' was initialized
	CreateTimeStamp time.Time

	dirPath      string      // Not part of FileInfo interface. Directory path associated with file name
	fName        string      // FileInfo.Name() base name of the file
	fSize        int64       // FileInfo.Size() length in bytes for regular files; system-dependent for others
	fMode        os.FileMode // FileInfo.Mode() file mode bits
	fModTime     time.Time   // FileInfo.ModTime() file modification time
	isDir        bool        // FileInfo.IsDir() 'true'= this is a directory not a file
	dataSrc      interface{} // FileInfo.Sys() underlying data source (can return nil)
	origFileInfo os.FileInfo
	lock         *sync.Mutex
}

//////////////////////////////////////////////////////////
// os.FileInfo Interface Methods
//////////////////////////////////////////////////////////

// Name
//
// Returns the Base name of the file.
//
//	Example:
//	            Complete File Name: "newerFileForTest_01.txt"
//	  Base Name returned by Name(): "newerFileForTest_01.txt"
func (fip FileInfoPlus) Name() string {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.fName
}

// Size
//
// Returns the file length in bytes for regular files;
// system-dependent for others.
func (fip FileInfoPlus) Size() int64 {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.fSize
}

// Mode
//
// Returns the file mode bits. See os.FileMode.
//
// A FileMode represents a file's mode and permission
// bits. The bits have the same definition on all
// systems, so that information about files can be moved
// from one system to another as a portable. Not all bits
// apply to all systems. The only required bit is ModeDir
// for directories.
//
// The FileMode is of type uint32.
//
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
// const (
//
//	 // The single letters are the abbreviations
//	 // used by the String method's formatting.
//		ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
//		ModeAppend                                     // a: append-only
//		ModeExclusive                                  // l: exclusive use
//		ModeTemporary                                  // T: temporary file; Plan 9 only
//		ModeSymlink                                    // L: symbolic link
//		ModeDevice                                     // D: device file
//		ModeNamedPipe                                  // p: named pipe (FIFO)
//		ModeSocket                                     // S: Unix domain socket
//		ModeSetuid                                     // u: setuid
//		ModeSetgid                                     // g: setgid
//		ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
//		ModeSticky                                     // t: sticky
//
//		// Mask for the type bits. For regular files, none will be set.
//		ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
//
//		ModePerm FileMode = 0777 // Unix permission bits
//
// )
func (fip FileInfoPlus) Mode() os.FileMode {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.fMode
}

// ModTime
//
// Returns the last file modification time.
func (fip FileInfoPlus) ModTime() time.Time {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.fModTime
}

// IsDir
//
// This method returns a boolean value of 'true'
// if the current instance of FileInfoPlus specifies
// a directory and not a file.
func (fip FileInfoPlus) IsDir() bool {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.isDir
}

// Sys
//
// This method returns the underlying data source for
// the current instance of FileInfoPlus.
//
// For more information on Sys(), reference:
//
//	https://pkg.go.dev/io/fs#FileInfo
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method can return nil.
func (fip FileInfoPlus) Sys() interface{} {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.dataSrc
}

// SysAsString
//
// Returns the underlying data source. If Sys is 'nil',
// this method will return an empty string.
//
// Technically, this method is NOT part of the
// os.FileInfo interface. However, it is often useful in
// interpreting the results of Sys().
//
// Sys() is part of the os.FileInfo interface.
//
// For more information on Sys(), reference:
//
//	https://pkg.go.dev/io/fs#FileInfo
func (fip *FileInfoPlus) SysAsString() string {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	if fip.dataSrc == nil {
		return ""
	}

	str := fmt.Sprintf("%v", fip.dataSrc)

	return str
}

//////////////////////////////////////////////////////////
// END OF os.FileInfo Interface Methods
//////////////////////////////////////////////////////////

// CopyOut
//
// Creates a deep copy of the current FileInfoPlus
// instance and returns it.
//
// This method is NOT part of the FileInfo interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//   - NONE -
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileInfoPlus
//
//		This method will return a new, fully populated
//		instance of FileInfoPlus. This new instance will
//		contain an exact copy of all data values
//	 	contained in the current instance of
//	 	FileInfoPlus.
func (fip *FileInfoPlus) CopyOut() FileInfoPlus {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	newInfo := FileInfoPlus{}

	newInfo.fName = fip.fName
	newInfo.fSize = fip.fSize
	newInfo.fMode = fip.fMode
	newInfo.fModTime = fip.fModTime
	newInfo.isDir = fip.isDir
	newInfo.dataSrc = fip.dataSrc

	newInfo.dirPath = fip.dirPath

	newInfo.isFInfoInitialized = fip.isFInfoInitialized
	newInfo.CreateTimeStamp = fip.CreateTimeStamp
	newInfo.origFileInfo = fip.origFileInfo

	return newInfo
}

// DirPath
//
// Returns the directory path. This field,
// FileInfoPlus.dirPath, is not part of the standard
// FileInfo interface.
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) DirPath() string {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.dirPath
}

// Equal
//
// Compares two FileInfoPlus objects to determine
// if they are equal.
//
// If all data fields in the current FileInfoPlus
// instance are equivalent to the corresponding data
// fields in the incoming FileInfo instance, this method
// returns 'true'.
//
// The Directory Path comparisons are NOT case-sensitive.
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) Equal(fip2 *FileInfoPlus) bool {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	if fip2 == nil {
		return false
	}

	if fip.fName != fip2.fName ||
		fip.fSize != fip2.fSize ||
		fip.fMode != fip2.fMode ||
		fip.isDir != fip2.isDir {

		return false
	}

	if !fip.fModTime.Equal(fip2.fModTime) {
		return false
	}

	if strings.ToLower(fip.dirPath) !=
		strings.ToLower(fip2.dirPath) {

		return false
	}

	if fip.dataSrc == nil && fip2.dataSrc == nil {
		return true
	}

	if fip.dataSrc == nil && fip2.dataSrc != nil {
		return false
	}

	if fip.dataSrc != nil && fip2.dataSrc == nil {
		return false
	}

	strFipSys := fmt.Sprintf("%v", fip.dataSrc)
	strFip2Sys := fmt.Sprintf("%v", fip2.dataSrc)

	if strFipSys != strFip2Sys {

		return false
	}

	return true
}

// Empty
//
// Sets the internal data fields of the current
// FileInfoPlus instance to their zero or nil value.
//
// This method is NOT part of the FileInfo interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values in the current instance of FileInfoPlus
//	to their zero or nil values.
func (fip *FileInfoPlus) Empty() {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.isFInfoInitialized = false

	fip.isDirPathInitialized = false

	fip.CreateTimeStamp = time.Time{}

	fip.dirPath = ""
	fip.fName = ""
	fip.fSize = 0
	fip.fMode = os.FileMode(0000)
	fip.fModTime = time.Time{}
	fip.isDir = false
	fip.dataSrc = nil
	fip.origFileInfo = nil
}

// GetOriginalFileInfo
//
// If the FileInfoPlus instance was initialized with an
// os.FileInfo value, this method will return that
// original os.FileInfo value. This is useful for passing
// parameters to some low level go routines such as
// os.SameFile().
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) GetOriginalFileInfo() os.FileInfo {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.origFileInfo
}

// IsFileInfoInitialized
//
// Returns a boolean value signaling whether this
// instance of FileInfoPlus has been initialized.
//
// A FileInfoPlus instance is properly initialized only
// if one of the following three methods is called:
//
// 1. FileInfoPlus.NewFromFileInfo()
// 2. FileInfoPlus.NewFromPathFileInfo()
// 3. FileInfoPlus.SetIsFInfoInitialized()
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) IsFileInfoInitialized() bool {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.isFInfoInitialized
}

// IsDirectoryPathInitialized
//
// Returns a boolean value signaling whether the
// directory path has been initialized for the current
// instance of FileInfoPlus.
//
// FYI, the fields FileInfoPlus.isDirPathInitialized
// and FileInfoPlus.dirPath do NOT exist in a standard
// os.FileInfo object.
//
// A FileInfoPlus directory path is properly initialized
// only if one of the following two methods is called:
//
// 1. FileInfoPlus.NewFromPathFileInfo()
// 2. FileInfoPlus.SetDirectoryPath
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) IsDirectoryPathInitialized() bool {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.isDirPathInitialized
}

// NewFromDirMgrFileInfo
//
// Creates and returns a new FileInfoPlus object
// populated with a Directory Manager (DirMgr) and
// File Info data (os.FileInfo) received from the input
// parameters 'dMgr' and 'info'.
//
// This method is NOT part of the FileInfo interface.
//
// ----------------------------------------------------------------
//
// # Usage Example
//
//	fip, err := FileInfoPlus{}.NewFromDirMgrFileInfo(dMgr, info)
//	fip is now configured as a newly populated FileInfoPlus instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						DirMgr
//
//		A concrete instance of Directory Manager
//		(DirMgr). Directory information provided by this
//		object will be used to populate the returned
//		instance of FileInfoPlus.
//
//	info os.FileInfo
//
//		An object implementing the FileInfo interface.
//		File information contained in this object will be
//		used to construct the returned instance of
//		FileInfoPlus.
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
//	FileInfoPlus
//
//		If this method completes successfully, a new,
//		fully populated, instance of FileInfoPlus will
//		be returned.
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
func (fip *FileInfoPlus) NewFromDirMgrFileInfo(
	dMgr DirMgr,
	info os.FileInfo,
	errorPrefix interface{}) (
	FileInfoPlus,
	error) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileInfoPlus."+
			"NewFromDirMgrFileInfo()",
		"")

	if err != nil {
		return FileInfoPlus{}, err
	}

	err = dMgr.IsValidInstanceError("")

	if err != nil {
		return FileInfoPlus{},
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'dMgr' is INVALID!\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err.Error())
	}

	if info == nil {
		return FileInfoPlus{},
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'info' is nil !\n",
				ePrefix.String())
	}

	newInfo := FileInfoPlus{}
	newInfo.lock = new(sync.Mutex)

	newInfo.fName = info.Name()
	newInfo.fSize = info.Size()
	newInfo.fMode = info.Mode()
	newInfo.fModTime = info.ModTime()
	newInfo.isDir = info.IsDir()
	newInfo.dataSrc = info.Sys()
	newInfo.isFInfoInitialized = true
	newInfo.origFileInfo = info

	newInfo.dirPath = dMgr.GetAbsolutePath()

	newInfo.isDirPathInitialized = true

	return newInfo, err
}

// NewFromFileInfo
//
// Creates and returns a new FileInfoPlus object
// populated with FileInfo data received from the input
// parameter.
//
// This method is NOT part of the FileInfo interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method does NOT set the Directory path member
//	variable in the returned instance of FileInfoPlus.
//
// ----------------------------------------------------------------
//
// # Usage Example
//
//	fip := FileInfoPlus{}.NewFromFileInfo(info)
//	fip is now a newly populated FileInfoPlus instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	info						os.FileInfo
//
//		An object implementing the FileInfo interface.
//		File information contained in this object will be
//		used to construct the returned instance of
//		FileInfoPlus.
//
//		For more information on the os.FileInfo interface
//		reference:
//			https://pkg.go.dev/io/fs#FileInfo
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileInfoPlus
//
//		This method will return a new, fully populated
//		instance of FileInfoPlus will be returned.
func (fip *FileInfoPlus) NewFromFileInfo(
	info os.FileInfo) FileInfoPlus {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	if info == nil {
		return FileInfoPlus{}
	}

	newInfo := FileInfoPlus{}
	newInfo.lock = new(sync.Mutex)

	newInfo.fName = info.Name()
	newInfo.fSize = info.Size()
	newInfo.fMode = info.Mode()
	newInfo.fModTime = info.ModTime()
	newInfo.isDir = info.IsDir()
	newInfo.dataSrc = info.Sys()
	newInfo.isFInfoInitialized = true
	newInfo.origFileInfo = info

	return newInfo
}

// NewFromPathFileInfo
//
// Creates and returns a new FileInfoPlus object
// populated with directory path and FileInfo data
// received from the input parameters.
//
// This method is NOT part of the FileInfo interface.
//
// ----------------------------------------------------------------
//
// # Usage Example
//
//	fip, err := FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)
//	fip is now a newly populated FileInfoPlus instance.
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
//	FileInfoPlus
//
//		This method will return a new, fully populated
//		instance of FileInfoPlus will be returned. The
//		directory path member variable will be set from
//		input parameter 'dirPath'. Basic file information
//		is taken form input parameter 'info'.
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
func (fip *FileInfoPlus) NewFromPathFileInfo(
	dirPath string,
	info os.FileInfo,
	errorPrefix interface{}) (
	FileInfoPlus,
	error) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileInfoPlus."+
			"NewFromPathFileInfo()",
		"")

	if err != nil {
		return FileInfoPlus{}, err
	}

	errCode := 0

	errCode,
		_,
		dirPath = new(fileHelperElectron).isStringEmptyOrBlank(dirPath)

	if errCode < 0 {
		return FileInfoPlus{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'dirPath' is an EMPTY String!\n",
				ePrefix.String())
	}

	if info == nil {
		return FileInfoPlus{},
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'info' is nil !\n",
				ePrefix.String())
	}

	newInfo := FileInfoPlus{}
	newInfo.lock = new(sync.Mutex)

	newInfo.fName = info.Name()
	newInfo.fSize = info.Size()
	newInfo.fMode = info.Mode()
	newInfo.fModTime = info.ModTime()
	newInfo.isDir = info.IsDir()
	newInfo.dataSrc = info.Sys()
	newInfo.origFileInfo = info

	newInfo.dirPath = dirPath

	newInfo.isDirPathInitialized = true

	return newInfo, nil
}

// SetDirectoryPath
//
// Sets the directory path (dirPath) member data field.
//
// This field is not part of the standard FileInfo data
// structure.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains the directory path which
//		will be stored in the 'dirPath' member variable
//		for the current instances of FileInfoPlus.
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
func (fip *FileInfoPlus) SetDirectoryPath(
	dirPath string,
	errorPrefix interface{}) error {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileInfoPlus."+
			"SetDirectoryPath()",
		"")

	if err != nil {
		return err
	}

	fh := FileHelper{}
	errCode := 0

	errCode,
		_,
		dirPath = fh.IsStringEmptyOrBlank(dirPath)

	if errCode < 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dirPath' is an EMPTY String!\n",
			ePrefix.String())
	}

	dirPath = fh.RemovePathSeparatorFromEndOfPathString(dirPath)

	fip.dirPath = dirPath

	fip.isDirPathInitialized = true

	return err
}

// SetName - Sets the file name field.
func (fip *FileInfoPlus) SetName(name string) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.fName = name
}

// SetSize - Sets the file size field
func (fip *FileInfoPlus) SetSize(fileSize int64) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.fSize = fileSize
}

// SetMode - Sets the file Mode
func (fip *FileInfoPlus) SetMode(fileMode os.FileMode) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.fMode = fileMode
}

// SetModTime - Sets the file modification time
func (fip *FileInfoPlus) SetModTime(fileModTime time.Time) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.fModTime = fileModTime
}

// SetIsDir - Sets is directory field.
func (fip *FileInfoPlus) SetIsDir(isDir bool) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.isDir = isDir
}

// SetSysDataSrc - Sets the dataSrc field
func (fip *FileInfoPlus) SetSysDataSrc(sysDataSrc interface{}) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	fip.dataSrc = sysDataSrc
}

// SetIsFInfoInitialized - Sets the flag for 'Is File Info Initialized'
// If set to 'true' it means that all the File Info fields have
// been initialized.
func (fip *FileInfoPlus) SetIsFInfoInitialized(isInitialized bool) {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	if !isInitialized {
		fip.isFInfoInitialized = false
		fip.CreateTimeStamp = time.Time{}
		return
	}

	fip.isFInfoInitialized = true
	fip.CreateTimeStamp = time.Now().Local()
	return
}
