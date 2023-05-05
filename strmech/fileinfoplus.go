package strmech

import (
	"errors"
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
func (fip *FileInfoPlus) Name() string {

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
func (fip *FileInfoPlus) Size() int64 {

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
func (fip *FileInfoPlus) Mode() os.FileMode {

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
func (fip *FileInfoPlus) ModTime() time.Time {

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
func (fip *FileInfoPlus) IsDir() bool {

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
// For more information on Sys, reference:
//
//	https://pkg.go.dev/io/fs#FileInfo
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method can return nil.
func (fip *FileInfoPlus) Sys() interface{} {

	if fip.lock == nil {
		fip.lock = new(sync.Mutex)
	}

	fip.lock.Lock()

	defer fip.lock.Unlock()

	return fip.dataSrc
}

// SysAsString - underlying data source. If Sys is
// 'nil', this method will return an empty string.
//
// Technically, this method is NOT part of the
// os.FileInfo interface. However, it is often
// useful in interpreting the results of Sys().
// Sys() is part of the os.FileInfo interface.
func (fip *FileInfoPlus) SysAsString() string {
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

// GetOriginalFileInfo - If the FileInfoPlus instance was initialized
// with an os.FileInfo value, this method will return that original
// os.FileInfo value. This is useful for passing parameters to some
// low level go routines such as os.SameFile().
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) GetOriginalFileInfo() os.FileInfo {
	return fip.origFileInfo
}

// IsFileInfoInitialized - Returns a boolean value signaling whether
// this instance of FileInfoPlus has been initialized.
//
// A FileInfoPlus instance is properly initialized only if one of the
// following three methods is called:
//
// 1. FileInfoPlus.NewFromFileInfo()
// 2. FileInfoPlus.NewFromPathFileInfo()
// 3. FileInfoPlus.SetIsFInfoInitialized()
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) IsFileInfoInitialized() bool {
	return fip.isFInfoInitialized
}

// IsDirectoryPathInitialized - Returns a boolean value signaling whether
// the directory path has been initialized for this instance of the
// FileInfoPlus instance. FYI, the fields FileInfoPlus.isDirPathInitialized
// and FileInfoPlus.dirPath do NOT exist in a standard os.FileInfo object.
//
// A FileInfoPlus directory path is properly initialized only if one of
// the following two methods is called:
//
// 1. FileInfoPlus.NewFromPathFileInfo()
// 2. FileInfoPlus.SetDirectoryPath
//
// This method is NOT part of the FileInfo interface.
func (fip *FileInfoPlus) IsDirectoryPathInitialized() bool {
	return fip.isDirPathInitialized
}

// NewFromDirMgrFileInfo - Creates and returns a new FileInfoPlus object
// populated with a Directory Manager (DirMgr) and File Info data (os.FileInfo)
// received from the input parameters 'dMgr' and 'info'.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//	fip, err := FileInfoPlus{}.NewFromDirMgrFileInfo(dMgr, info)
//	fip is now configured as a newly populated FileInfoPlus instance.
func (fip *FileInfoPlus) NewFromDirMgrFileInfo(
	dMgr DirMgr,
	info os.FileInfo) (FileInfoPlus, error) {

	ePrefix := "FileInfoPlus.NewFromDirMgrFileInfo() "
	var err error
	err = dMgr.IsDirMgrValid("")

	if err != nil {
		return FileInfoPlus{},
			fmt.Errorf(ePrefix+"ERROR: Input Parameter 'dMgr' is INVALID!\n"+
				"%v", err.Error())
	}

	if info == nil {
		return FileInfoPlus{},
			errors.New(ePrefix + "ERROR: Input Parameter 'info' is nil !\n")
	}

	newInfo := new(FileInfoPlus).NewFromFileInfo(info)

	newInfo.dirPath = dMgr.GetAbsolutePath()

	newInfo.isDirPathInitialized = true

	return newInfo, nil
}

// NewFromFileInfo - Creates and returns a new FileInfoPlus object
// populated with FileInfo data received from the input parameter.
// Notice that this version of the 'NewFromPathFileNameExtStr' method does NOT set the
// Directory path.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//	fip := FileInfoPlus{}.NewFromFileInfo(info)
//	fip is now a newly populated FileInfoPlus instance.
func (fip *FileInfoPlus) NewFromFileInfo(info os.FileInfo) FileInfoPlus {

	if info == nil {
		return FileInfoPlus{}
	}

	newInfo := FileInfoPlus{}

	newInfo.SetName(info.Name())
	newInfo.SetSize(info.Size())
	newInfo.SetMode(info.Mode())
	newInfo.SetModTime(info.ModTime())
	newInfo.SetIsDir(info.IsDir())
	newInfo.SetSysDataSrc(info.Sys())
	newInfo.SetIsFInfoInitialized(true)
	newInfo.origFileInfo = info
	return newInfo
}

// NewFromPathFileInfo - Creates and returns a new FileInfoPlus object
// populated with directory path and FileInfo data received from
// the input parameters.
//
// This method is NOT part of the FileInfo interface.
//
// ------------------------------------------------------------------------
//
// Example Usage:
//
//	fip, err := FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)
//	fip is now a newly populated FileInfoPlus instance.
func (fip *FileInfoPlus) NewFromPathFileInfo(
	dirPath string,
	info os.FileInfo) (FileInfoPlus, error) {

	ePrefix := "FileInfoPlus.NewFromPathFileInfo() "

	errCode := 0

	errCode,
		_,
		dirPath = new(fileHelperElectron).isStringEmptyOrBlank(dirPath)

	if errCode < 0 {
		return FileInfoPlus{},
			fmt.Errorf(ePrefix +
				"\nError: Input parameter 'dirPath' is an EMPTY String!\n")
	}

	if info == nil {
		return FileInfoPlus{},
			errors.New(ePrefix + "ERROR: Input Parameter 'info' is nil !\n")
	}

	newInfo := new(FileInfoPlus).NewFromFileInfo(info)

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
	fip.fName = name
}

// SetSize - Sets the file size field
func (fip *FileInfoPlus) SetSize(fileSize int64) {
	fip.fSize = fileSize
}

// SetMode - Sets the file Mode
func (fip *FileInfoPlus) SetMode(fileMode os.FileMode) {
	fip.fMode = fileMode
}

// SetModTime - Sets the file modification time
func (fip *FileInfoPlus) SetModTime(fileModTime time.Time) {
	fip.fModTime = fileModTime
}

// SetIsDir - Sets is directory field.
func (fip *FileInfoPlus) SetIsDir(isDir bool) {
	fip.isDir = isDir
}

// SetSysDataSrc - Sets the dataSrc field
func (fip *FileInfoPlus) SetSysDataSrc(sysDataSrc interface{}) {
	fip.dataSrc = sysDataSrc
}

// SetIsFInfoInitialized - Sets the flag for 'Is File Info Initialized'
// If set to 'true' it means that all the File Info fields have
// been initialized.
func (fip *FileInfoPlus) SetIsFInfoInitialized(isInitialized bool) {
	if !isInitialized {
		fip.isFInfoInitialized = false
		fip.CreateTimeStamp = time.Time{}
		return
	}

	fip.isFInfoInitialized = true
	fip.CreateTimeStamp = time.Now().Local()
	return
}
