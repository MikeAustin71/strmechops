package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

/*
  This source code file contains type 'DirMgrCollection' .

  The Source Repository for this source code file is :
    https://github.com/MikeAustin71/pathfileopsgo.git

  Dependencies:
  -------------

  Type 'DirMgrCollection' depends on types, 'FileHelper',
  'FileMgr' and 'DirMgr' which are contained in source
  code files, 'filehelper.go', 'filemanager.go' and
  'dirmanager.go' located in this directory.

*/

// DirMgrCollection - A collection of Type DirMgr. The collection
// is used to aid in the management of groups of paths and directories.
//
// Dependencies:
//
// Type 'DirMgrCollection' depend on types, 'FileHelper' and
// 'FileMgr' which are contained in source code files: 'filehelper.go'
// and 'filemanager.go' located in this directory.
type DirMgrCollection struct {
	dirMgrs []DirMgr
	lock    *sync.Mutex
}

// AddDirMgr
// Adds a DirMgr object to the collection.
//
// Note that this method does not perform a validity
// check on input parameter, 'dMgr'.
//
// It is recommended that dMgr.IsDirMgrValid() be called
// before adding the directory manager to the collection.
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
//	dMgr						DirMgr
//
//		A concrete instance of DirMgr. A copy of this
//		instance will be added to the internal Directory
//		Manager Collection maintained by the current
//		instance of DirMgrCollection.
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
func (dMgrs *DirMgrCollection) AddDirMgr(
	dMgr DirMgr,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"AddDirMgr()",
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	var dMgrCopy DirMgr

	dMgrCopy,
		err = dMgr.CopyOut(
		ePrefix.XCpy(
			"dMgr->dMgrCopy"))

	if err != nil {
		return err
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)

	return err
}

// AddDirMgrByKnownPathDirName
//
// Adds a Directory Manager (DirMgr) using known parent
// path and directory names.
//
// The two input parameters represent the two components
// of the complete directory path: the parent directory
// and a single subdirectory.
//
// The two components will be combined to create the
// complete and total directory path.
//
// This method performs fewer string validations then
// similar methods. Quoting Davy Crockett, "Be sure your
// right, then go ahead".
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentPathName				string
//
//		This string contains the parent path component
//		of the complete directory path. This parent
//		directory will be combined with the single
//		subdirectory supplied by input parameter
//		'dirName' to create the total directory.
//
//		Example:
//		parentPathName = "c:/papadirectory"
//		dirName = "babydirectory"
//		Final Directory Name:
//			"c:/papadirectory/babydirectory"
//
//	dirName						string
//
//		The actual name of the directory which will be
//		combined with the parent path ('parentPathName')
//		to create the final and complete directory name.
//
//		Example:
//		parentPathName = "c:/papadirectory"
//		dirName = "babydirectory"
//		Final Directory Name:
//			"c:/papadirectory/babydirectory"
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
func (dMgrs *DirMgrCollection) AddDirMgrByKnownPathDirName(
	parentPathName,
	dirName string,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"AddDirMgrByKnownPathDirName()",
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	newDirMgr := DirMgr{}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrFromKnownPathDirName(
			&newDirMgr,
			parentPathName,
			dirName,
			"newDirMgr",
			"parentPathName",
			"dirName",
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"Returned 'DirMgr' is Empty!\n"+
			"dMgrHlpr.setDirMgrFromKnownPathDirName()\n"+
			"parentPathName= '%v'\n"+
			"dirName= '%v'\n",
			ePrefix.String(),
			parentPathName,
			dirName)
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgr)

	return nil
}

// AddDirMgrByPathNameStr
//
// Adds a Directory Manager (DirMgr) to the collections
// based on a string input parameter, 'pathName'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathName					string
//
//		This string contains the complete directory path
//		which will be converted to a Directory Manager
//		type (DirMgr) before being added to the Directory
//		Manager Collection maintained by the current
//		instance of DirMgrCollection.
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
func (dMgrs *DirMgrCollection) AddDirMgrByPathNameStr(
	pathName string,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "DirMgrCollection.AddDirMgrByPathNameStr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	newDirMgr := DirMgr{}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelperNanobot).setDirMgr(
		&newDirMgr,
		pathName,
		"newDirMgr",
		"pathName",
		ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned from dMgrHlpr.setDirMgr(pathName).\n"+
			"pathName='%v'\n"+
			"Error='%v'",
			funcName,
			pathName,
			err.Error())

	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Returned 'DirMgr' is Empty!\n"+
			"dMgrHlpr.setDirMgrFromKnownPathDirName()\n"+
			"parentPathName='%v'\n",
			ePrefix.String(),
			pathName)
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgr)

	return err
}

// AddFileInfo
//
// Adds a Directory Manager object to the collection
// based on input from a parent directory path string and
// an os.FileInfo object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentDirectoryPath			string
//
//		The parent directory path. This path will be
//		combined with any subsidiary directories found
//		in the os.FileInfo object to create the final
//		directory path which will be converted to a
//		Directory Manager (DirMgr) and added to the
//		Directory Manager Collection maintained by the
//		current instance of DirMgrCollection.
//
//	fInfo						os.FileInfo
//
//		An object which implements the os.FileInfo
//		interface. This parameter may transmit an
//		instance of FileInfoPlus which implements
//		the os.FileInfo interface but provides file
//		information over and above that provided by the
//		standard os.FileInfo interface.
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
func (dMgrs *DirMgrCollection) AddFileInfo(
	parentDirectoryPath string,
	fInfo os.FileInfo,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection." +
		"AddFileMgrByFileInfo()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	newDirMgr := DirMgr{}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrFromKnownPathDirName(
			&newDirMgr,
			parentDirectoryPath,
			fInfo.Name(),
			"newDirMgr",
			"parentDirectoryPath",
			"FileInfo.Name()",
			ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned from dMgrHlpr.setDirMgrFromKnownPathDirName("+
			"parentDirectoryPath, FileInfo.Name()).\n"+
			"parentDirectoryPath='%v'\n"+
			"FileInfo.Name()='%v'\n"+
			"Error= \n%v\n",
			funcName,
			parentDirectoryPath,
			fInfo.Name(),
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"Returned 'DirMgr' is Empty!\n"+
			"dMgrHlpr.setDirMgrFromKnownPathDirName()\n"+
			"parentDirectoryPath= '%v'\n"+
			"FileInfo.Name()= '%v'\n",
			ePrefix.String(),
			parentDirectoryPath,
			fInfo.Name())
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgr)

	return err
}

// AddDirMgrCollection
//
// Adds another collection of Directory Manager (DirMgr)
// objects (dMgrs2) to the Directory Manager Collection
// maintained by the current instance of
// DirMgrCollection. The new collection will be appended
// to the end of the Directory Manager Collection
// maintained by the current instance of
// DirMgrCollection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgrs2						*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection. The
//		Directory Manager collection contained in
//		'dMgrs2' will be appended to the end of the
//		Directory Managers Collection maintained by the
//		current instance of DirMgrCollection.
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
func (dMgrs *DirMgrCollection) AddDirMgrCollection(
	dMgrs2 *DirMgrCollection,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection." +
		"AddDirMgrCollection()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if dMgrs2 == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'dMgrs2' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lDmc2 := len(dMgrs2.dirMgrs)

	if lDmc2 == 0 {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'dMgrs2' contains a zero"+
			"length Directory Managers array (dMgrs2.dirMgrs)\n",
			ePrefix.String())

		return err
	}

	var dMgrCopy DirMgr

	for i := 0; i < lDmc2; i++ {

		dMgrCopy,
			err = dMgrs2.dirMgrs[i].CopyOut(
			ePrefix.XCpy("dMgrs2.dirMgrs[i]"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error returned by dMgrs2.dirMgrs[%v].CopyOut()\n"+
				"dMgrs2 index = %v\n"+
				"Directory Manager = %v\n"+
				"Error=\n%v\n",
				funcName,
				i,
				i,
				dMgrs2.dirMgrs[i].absolutePath,
				err.Error())
		}

		dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)
	}

	return err
}

//	CopyIn
//
//	Copies the Directory Manager Collection from an
//	incoming instance of DirMgrCollection
//	('incomingDMgrCollection')	to the Directory Manager
//	Collection maintained by the current DirMgrCollection
//	instance ('dMgrs').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values in the current
//	DirMgrCollection instance ('dMgrs') will be deleted
//	and overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingDMgrCollection 		*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		This method will NOT change the values of
//		internal member variables contained in this
//		instance.
//
//		The Directory Manager Collection contained in
//		this DirMgrCollection instance will be copied to
//		the current DirMgrCollection instance ('dMgrs').
//
//		If 'incomingDMgrCollection' contains an empty
//		Directory Manager Collection, an error will be
//		returned.
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
func (dMgrs *DirMgrCollection) CopyIn(
	incomingDMgrCollection *DirMgrCollection,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(dirMgrCollectionHelper).
		copyCollection(
			dMgrs,
			incomingDMgrCollection,
			ePrefix.XCpy(
				"dMgrs<-incomingDMgrCollection"))
}

// CopyOut
//
// Returns an DirMgrCollection which is an exact
// duplicate of the current DirMgrCollection.
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
//	DirMgrCollection
//
//		If this method completes successfully, an exact
//		copy of the Directory Manager Collection
//		contained in the current instance of
//		DirMgrCollection will be returned in a new
//		instance of DirMgrCollection.
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
func (dMgrs *DirMgrCollection) CopyOut(
	errorPrefix interface{}) (
	DirMgrCollection,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection.CopyOut()",
		"")

	if err != nil {
		return DirMgrCollection{}, err
	}

	dMgrs2 := DirMgrCollection{}.New()

	err = new(dirMgrCollectionHelper).
		copyCollection(
			&dMgrs2,
			dMgrs,
			ePrefix.XCpy(
				"dMgrs2<-dMgrs"))

	return dMgrs2, err
}

// DeleteAtIndex
//
// Deletes a member Directory Manager object from the
// Directory Manager Collection encapsulated by the
// current instance of DirMgrCollection.
//
// The Directory Manager object to be deleted from the
// Directory Manager Collection is specified by the
// array index passed by input parameter 'idx.
//
// If successful, at the completion of this method, the
// Directory Manager Collection array will have a length
// which is one less than the starting array length.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies the Directory
//		Manager array index which will identify the
//		Directory Manager object (DirMgr) to be deleted.
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
func (dMgrs *DirMgrCollection) DeleteAtIndex(
	idx int,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection.DeleteAtIndex()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			idx,
			true, // Delete DirMgr object
			ePrefix.XCpy(
				"dMgrs"))

	return err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of DirMgrCollection to their initial or zero
//	values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing DirMgr
//	objects contained in the Directory Manager Collection
//	encapsulated in the current instance of
//	DirMgrCollection. Upon completion of this method, the
//	internal Directory Manager Collection for the current
//	DirMgrCollection instance will have a length of zero.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (dMgrs *DirMgrCollection) Empty() {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	colArrayLen := len(dMgrs.dirMgrs)

	for i := 0; i < colArrayLen; i++ {

		dMgrs.dirMgrs[i].Empty()

	}

	dMgrs.dirMgrs = make([]DirMgr, 0)

	dMgrs.lock.Unlock()

	return
}

// FindDirectories
//
// Searches through the DirMgrCollection to find DirMgr
// objects matching specified search criteria.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirMgrSelectionCriteria		FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired directory manager selection criteria.
//		Instances of DirMgr in the Directory Manager
//		Collection matching this criteria will be
//		selected and returned as a new and separate
//		Directory Manager Collection.
//
//		type FileSelectionCriteria struct {
//		 FileNamePatterns    []string
//			An array of strings containing File Name Patterns
//
//		 FilesOlderThan      time.Time
//		 	Match files with older modification date times
//
//		 FilesNewerThan      time.Time
//		 	Match files with newer modification date times
//
//		 SelectByFileMode    FilePermissionConfig
//		 	Match file mode (os.FileMode).
//
//		 SelectCriterionModeFileSelectCriterionMode
//		 	Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
//
//			FileNamePatterns		[]string
//
//				An array of strings which may define one or more
//				search patterns. If a file name matches any one
//				of the search pattern strings, it is deemed to be
//				a 'match' for the search pattern criterion.
//
//				Example Patterns:
//					FileNamePatterns = []string{"*.log"}
//					FileNamePatterns = []string{"current*.txt"}
//					FileNamePatterns = []string{"*.txt", "*.log"}
//
//				If this string array has zero length or if
//				all the strings are empty strings, then this
//				file search criterion is considered 'Inactive'
//				or 'Not Set'.
//
//
//			FilesOlderThan		time.Time
//
//				This date time type is compared to file
//				modification date times in order to determine
//				whether the file is older than the
//				'FilesOlderThan' file selection criterion. If
//				the file modification date time is older than
//				the 'FilesOlderThan' date time, that file is
//				considered a 'match' for this file selection
//				criterion.
//
//				If the value of 'FilesOlderThan' is set to
//				time zero, the default value for type
//				time.Time{}, then this file selection
//				criterion is considered to be 'Inactive' or
//				'Not Set'.
//
//			FilesNewerThan      time.Time
//
//				This date time type is compared to the file
//				modification date time in order to determine
//				whether the file is newer than the
//				'FilesNewerThan' file selection criterion. If
//				the file modification date time is newer than
//				the 'FilesNewerThan' date time, that file is
//				considered a 'match' for this file selection
//				criterion.
//
//				If the value of 'FilesNewerThan' is set to
//				time zero, the default value for type
//				time.Time{}, then this file selection
//				criterion is considered to be 'Inactive' or
//				'Not Set'.
//
//			SelectByFileMode  FilePermissionConfig
//
//				Type FilePermissionConfig encapsulates an os.FileMode. The
//				file selection criterion allows for the selection of files
//				by File Mode.
//
//				File modes are compared to the value of 'SelectByFileMode'.
//				If the File Mode for a given file is equal to the value of
//				'SelectByFileMode', that file is considered to be a 'match'
//				for this file selection criterion. Examples for setting
//				SelectByFileMode are shown as follows:
//
//				fsc := FileSelectionCriteria{}
//
//				err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//				err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//			SelectCriterionMode FileSelectCriterionMode
//
//			This parameter selects the manner in which the file selection
//			criteria above are applied in determining a 'match' for file
//			selection purposes. 'SelectCriterionMode' may be set to one of
//			two constant values:
//
//			(1) FileSelectCriterionMode(0).ANDSelect()
//
//				File selected if all active selection criteria
//				are satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will not be judged as 'selected' unless all
//				the active selection criterion are satisfied. In other words, if
//				three active search criterion are provided for 'FileNamePatterns',
//				'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//				selected unless it has satisfied all three criterion in this example.
//
//			(2) FileSelectCriterionMode(0).ORSelect()
//
//				File selected if any active selection criterion is satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will be selected if any one of the active file
//				selection criterion is satisfied. In other words, if three active
//				search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//				and 'FilesNewerThan', then a file will be selected if it satisfies any
//				one of the three criterion in this example.
//
//		------------------------------------------------------------------------
//
//		IMPORTANT:
//
//		If all of the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all the
//		files processed in the directory tree will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected.
//
//		------------------------------------------------------------------------
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
func (dMgrs *DirMgrCollection) FindDirectories(
	dirMgrSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	DirMgrCollection,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection.FindDirectories()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return DirMgrCollection{}.New(), err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	lDirCol := len(dMgrs.dirMgrs)

	if lDirCol == 0 {
		return DirMgrCollection{}.New(), nil
	}

	fh := FileHelper{}

	var isMatchedFile bool

	dMgrs2 := DirMgrCollection{}.New()

	for i := 0; i < lDirCol; i++ {
		dMgr := dMgrs.dirMgrs[i]

		if dMgr.actualDirFileInfo.isFInfoInitialized {

			isMatchedFile,
				err,
				_ =
				fh.FilterFileName(
					&dMgr.actualDirFileInfo,
					dirMgrSelectionCriteria,
					ePrefix)

			if err != nil {

				return DirMgrCollection{},
					fmt.Errorf("%v\n"+
						"Error returned by fh.FilterFileName("+
						"dMgr.actualDirFileInfo, dirMgrSelectionCriteria)\n"+
						"dMgr.actualDirFileInfo.Name()='%v'\n"+
						"Error= \n%v\n",
						funcName,
						dMgr.actualDirFileInfo.Name(),
						err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(dMgr.directoryName)

			isMatchedFile,
				err,
				_ = fh.FilterFileName(
				&fip,
				dirMgrSelectionCriteria,
				ePrefix)

			if err != nil {

				var err2 error

				err2 = fmt.Errorf("%v\n"+
					"Error returned by fh.FilterFileName(fip, dirMgrSelectionCriteria)\n"+
					"fip.Name()= '%v'\n"+
					"Error= \n%v\n",
					funcName,
					fip.Name(),
					err.Error())

				return DirMgrCollection{}, err2
			}

		}

		if isMatchedFile {

			var dMgrCopy DirMgr

			dMgrCopy,
				err = dMgr.CopyOut(
				ePrefix.XCpy("dMgr->dMgrCopy"))

			if err != nil {
				return dMgrs2, err
			}

			dMgrs2.dirMgrs = append(dMgrs2.dirMgrs, dMgrCopy)
		}

	}

	return dMgrs2, err
}

// GetDirMgrArray
//
// Returns the Directory Manager Collection.
//
// This method returns a deep copy of the Directory
// Manager Array maintained by the current instance of
// DirMgrCollection.
//
// If the current DirMgrCollection Directory Manager
// Collection is empty, this method will return an
// empty array and no error will be returned.
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
//	[]DirMgr
//
//		A deep copy of the array of DirMgr objects
//		contained in the Directory Manager Collection
//		encapsulated in the current instance of
//		DirMgrCollection.
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
func (dMgrs *DirMgrCollection) GetDirMgrArray(
	errorPrefix interface{}) (
	[]DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection." +
		"GetDirMgrArray()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return make([]DirMgr, 0), err
	}

	lenDirMgrs := len(dMgrs.dirMgrs)

	if lenDirMgrs == 0 {
		return make([]DirMgr, 0), err
	}

	var newDirMgrCollection []DirMgr

	newDirMgrCollection = make([]DirMgr, lenDirMgrs)

	for i := 0; i < lenDirMgrs; i++ {

		err = newDirMgrCollection[i].CopyIn(
			&dMgrs.dirMgrs[i],
			ePrefix.XCpy(
				"dMgrs.dirMgrs[i]"))

		if err != nil {

			return make([]DirMgr, 0),
				fmt.Errorf("%v\n"+
					"Error: newDirMgrCollection[%v].CopyIn()\n"+
					"dMgrs.dirMgrs Index = %v\n"+
					"dMgrs.dirMgrs[%v] = %v\n"+
					"Error= \n%v\n",
					funcName,
					i,
					i,
					i,
					dMgrs.dirMgrs[i].absolutePath,
					err.Error())
		}
	}

	return newDirMgrCollection, err
}

// GetDirMgrAtIndex
//
// If successful, this method returns a pointer to the
// DirMgr instance at the array index specified. The
// 'Peek' and 'Pop' methods below return DirMgr objects
// using a 'deep' copy and therefore offer better
// protection against data corruption.
//
// If the Directory Manager Collection for the current
// DirMgrCollection is empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	Since this method returns a pointer to Directory
//	Manager object residing in the Directory Manager
//	Collection maintained by the current DirMgrCollection
//	instance, users must be cautious about changing data
//	values.
//
//	Since users are working a pointer to a DirMgr instance,
//	altering data values will change the content of the
//	Directory Manager Collection.
//
//	This feature is provided to promote efficiency. In
//	the interests of data integrity and safety, users
//	are strongly encouraged use one of the 'peek' or
//	'pop' methods which return deep copies of the
//	specified DirMgr object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies the index of the
//		Directory Manager (DirMgr) object in the
//		Directory Manager Collection maintained by the
//		current instance of	DirMgrCollection.
//
//		A pointer to the DirMgr object identified by this
//		Directory Manager Collection array index value
//		will be returned by this method.
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
//	*DirMgr
//
//		If this method completes successfully, a pointer
//		to the DirMgr object identified by the Directory
//		Manager Collection array index value 'idx' will
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
func (dMgrs *DirMgrCollection) GetDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	*DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection." +
		"GetDirMgrAtIndex()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return &DirMgr{}, err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return &DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: This Directory Manager Collection ('DirMgrCollection') is EMPTY!\n",
				ePrefix.String())
	}

	if idx < 0 || idx >= arrayLen {

		return &DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The input parameter, 'idx', is OUT OF RANGE!\n"+
				"idx='%v'\n"+
				"The minimum index is '0'. "+
				"The maximum index is '%v'. ",
				ePrefix.String(),
				idx,
				arrayLen-1)

	}

	return &dMgrs.dirMgrs[idx], err
}

// GetNumOfDirs
//
// Returns the number of Directory Manager objects
// contained in the current Directory Manager Collection.
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
//	int
//
//		The integer value returned by this method
//		specifies the number of Directory Manager objects
//		(DirMgr) contained in the Directory Manager
//		Collection maintained by the current instance of
//		DirMgrCollection
func (dMgrs *DirMgrCollection) GetNumOfDirs() int {

	return len(dMgrs.dirMgrs)
}

// InsertDirMgrAtIndex
//
// Inserts a new Directory Manager into the collection at
// array 'index'. The new Directory Manager is passed as
// input parameter 'dMgr'. A deep copy of 'dMgr' will be
// added to the Directory Manager Collection maintained
// by the current instance of DirMgrCollection.
//
// When inserting the new Directory Manager to the
// collection, the current Directory Manager residing at
// the specified 'index' will NOT be deleted. Instead,
// the entire array will be expanded one element and the
// current resident of array index will reside at
// index+1.
//
// If input parameter 'index' is less than zero, an error
// will be returned.
//
// If 'index' exceeds the value of the last index in the
// collection, 'dMgr' will be added to the end of the
// collection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						DirMgr
//
//		A concrete instance of DirMgr. A deep copy of
//		'dMgr' will be inserted into Directory Manager
//		Collection at the array index specified by input
//		parameter 'index'.
//
//		If 'dMgr' fails validation tests and proves to be
//		invalid, an error will be returned.
//
//	index						int
//
//		An integer value designating the array index
//		at which a deep copy of 'dMgr' will be inserted
//		into the Directory Manager Collection maintained
//		by the current instance of DirMgr.
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
func (dMgrs *DirMgrCollection) InsertDirMgrAtIndex(
	dMgr DirMgr,
	index int,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "DirMgrCollection.InsertDirMgrAtIndex()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	err = dMgr.IsDirMgrValid(ePrefix.XCpy("dMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: dMgr.IsDirMgrValid()"+
			"Input parameter 'dMgr' is INVALID!\n"+
			"Error = \n%v\n",
			funcName,
			err.Error())
	}

	if index < 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'index' is LESS THAN ZERO! "+
			"index='%v' ",
			ePrefix.String(),
			index)
	}

	lenDgrs := len(dMgrs.dirMgrs)

	var dMgrCopy DirMgr

	dMgrCopy,
		err = dMgr.CopyOut(ePrefix.XCpy(
		"dMgr"))

	if err != nil {
		return err
	}

	if index >= lenDgrs {

		dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)

		return err
	}

	newDirMgrs := make([]DirMgr, 0, 100)

	if index == 0 {
		newDirMgrs = append(newDirMgrs, dMgrCopy)
		dMgrs.dirMgrs = append(newDirMgrs, dMgrs.dirMgrs...)
		return nil
	}

	newDirMgrs = append(newDirMgrs, dMgrs.dirMgrs[index:]...)

	dMgrs.dirMgrs = append(dMgrs.dirMgrs[:index])
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgrs...)

	return nil
}

// New - Creates and returns a new and properly initialized
// Directory Manager Collection ('DirMgrCollection').
func (dMgrs DirMgrCollection) New() DirMgrCollection {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 5)
	}

	newDirMgrCol := DirMgrCollection{}
	newDirMgrCol.dirMgrs = make([]DirMgr, 0, 5)

	return newDirMgrCol
}

// PopDirMgrAtIndex
//
//	Returns a deep copy of the Directory Manager
//
// ('DirMgr') object located at index, 'idx', in the
// Directory Manager Collection array maintained by the
// current instance of DirMgrCollection.
//
// As a 'Pop' method, the original Directory Manager
// ('DirMgr') object is deleted from the Directory
// Manager Collection ('DirMgrCollection') array. The
// 'DirMgr' object deleted is located at the index
// specified by input parameter, 'idx'.
//
// Therefore, at the completion of this method, the
// Directory Manager Collection array has a length which
// is one less than the starting array length.
//
// If this method is called on an empty Directory Manager
// Collection (i.e. length of array dMgrs.dirMgrs = 0),
// an error will be returned.
//
// If this method is called with and index ('idx') value
// greater than the last index in the collection, an
// error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		An integer value designating the index value of
//		the array element to be deleted from the
//		Directory Manager Collection. A deep copy of this
//		deleted Directory Manager object will be returned
//		by this method.
//
//		If the value of 'idx' is less than zero or
//		greater than the index value of the last element
//		in the Directory Manager Collection, an error will
//		be returned.
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
//	DirMgr
//
//		If this method completes successfully, a copy of
//		the Directory Manager objected deleted from the
//		Directory Manager Collection will be returned.
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
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PopDirMgrAtIndex()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if idx < 0 {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is less than zero.\n"+
				"Index Out-Of-Range! idx='%v'\n",
				ePrefix.String(),
				idx)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The Directory Manager Collection is EMPTY!\n"+
				"The length of 'dMgrs.dirMgrs' is Zero Elements.\n",
				ePrefix.String())
	}

	if idx >= arrayLen {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				"length of the collection index. Index Out-Of-Range!\n"+
				"idx='%v' Array Length='%v'\n",
				ePrefix.String(),
				idx,
				arrayLen)
	}

	if idx < 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is less than zero.\n"+
				"Index ('idx') is Out-Of-Range!\n"+
				"idx='%v' Array Length='%v'\n",
				ePrefix.String(),
				idx,
				arrayLen)

	}

	var dirMgrCopy DirMgr

	dirMgrCopy,
		err = dMgrs.dirMgrs[idx].CopyOut(
		ePrefix.XCpy(fmt.Sprintf(
			"dMgrs.dirMgrs[%v]",
			idx)))

	if err != nil {

		return dirMgrCopy, err
	}

	if idx == 0 {
		// First Element

		dMgrs.dirMgrs = dMgrs.dirMgrs[1:]

		return dirMgrCopy, err
	}

	if idx == arrayLen-1 {
		// Last Element
		dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]

		return dirMgrCopy, err
	}

	dMgrs.dirMgrs =
		append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)

	return dirMgrCopy, err
}

// PopFirstDirMgr
//
// Returns a deep copy of the first Directory Manager
// ('DirMgr') object in the Directory Manager Collection
// array for the current DirMgrCollection instance. As a
// 'Pop' method, the original Directory Manager
// ('DirMgr') object is deleted from the first array
// index of the Directory Manager Collection
// (dMgrs.dirMgrs array index = 0).
//
// If this method is called on an empty Director Manager
// Collection (i.e. length of array dMgrs.dirMgrs = 0),
// an error will be returned.
//
// After the successful completion of this method, the
// Directory Manager Collection array has a length which
// is one less than the starting array length.
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
//	DirMgr
//
//		If this method completes successfully, a deep
//		copy of the DirMgr object 'popped' or deleted
//		from the first array element position in the
//		Directory Managers Collection (dMgrs.dirMgrs
//		array index = 0) will be returned through this
//		parameter.
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
func (dMgrs *DirMgrCollection) PopFirstDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PopFirstDirMgr()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 5)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The Director Managers Collection\n"+
				"for the current DirMgrCollection instance\n"+
				"is EMPTY. The Collection array has an array\n"+
				"length of Zero.\n",
				ePrefix.String())
	}

	var dMgrCopy DirMgr

	dMgrCopy,
		err = dMgrs.dirMgrs[0].CopyOut(ePrefix.XCpy(
		"dMgrs.dirMgrs[0]"))

	if err != nil {
		return DirMgr{}, err
	}

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)

	} else {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
	}

	return dMgrCopy, err
}

// PopLastDirMgr
//
// Returns a deep copy of the last Directory Manager
// ('DirMgr') object in the Directory Manager Collection
// array. As a 'Pop' method, the original Directory
// Manager ('DirMgr') object is deleted from the last
// array position of the Directory Manager Collection
// ('DirMgrCollection') array (dMgrs.dirMgrs[length-1]).
//
// If this method is called on an empty Director Manager
// Collection (i.e. length of array dMgrs.dirMgrs = 0),
// an error will be returned.
//
// At the successful completion of this method, the
// Directory Manager Collection array will have a length
// which is one less than the starting array length.
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
//	DirMgr
//
//		If this method completes successfully, a deep
//		copy of the DirMgr object 'popped' or deleted
//		from the last array element position in the
//		Directory Managers Collection
//		(dMgrs.dirMgrs[length-1]) will be returned
//		through this parameter.
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
func (dMgrs *DirMgrCollection) PopLastDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PopLastDirMgr()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 10)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The Director Managers Collection\n"+
				"for the current DirMgrCollection instance\n"+
				"is EMPTY. The Collection array has an array\n"+
				"length of Zero.\n",
				ePrefix.String())
	}

	var dirMgrCopy DirMgr

	dirMgrCopy,
		err = dMgrs.dirMgrs[arrayLen-1].CopyOut(
		ePrefix.XCpy(fmt.Sprintf("dMgrs.dirMgrs[%v-1]",
			arrayLen)))

	if err != nil {
		return DirMgr{}, err
	}

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)

	} else {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
	}

	return dirMgrCopy, err
}

// PeekDirMgrAtIndex
//
// Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the
// current Directory Manager Collection
// ('DirMgrCollection').
//
// This is a 'Peek' method and therefore the original
// Directory Manager ('DirMgr') object stored at index
// 'idx' is NOT deleted from the Directory Manager
// Collection ('DirMgrCollection') array.
//
// At the completion of this method, the length of the
// current Directory Manager Collection
// ('DirMgrCollection') array will remain unchanged.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies an index in the
//		Directory Manager Collection array for the
//		current instance of DirMgrCollection. A deep copy
//		of  the Directory Manager object residing at this
//		index will be returned to the calling function.
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
//	DirMgr
//
//		If this method completes successfully, a deep
//		copy of the Directory Manager object residing at
//		the array index 'idx' in the Directory Manager
//		Collection will be returned through this
//		parameter.
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
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PeekDirMgrAtIndex()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The current DirMgrCollection is Empty!\n"+
				"The are zero Directory Manager objects ('DirMgr')\n"+
				"in the collection.\n",
				ePrefix.String())

	}

	if idx < 0 {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is less than zero.\n"+
				"Index Out-Of-Range!\n"+
				"idx='%v'",
				ePrefix.String(),
				idx)
	}

	if idx >= arrayLen {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				"length of the collection array. Index Out-Of-Range!\n"+
				"idx='%v'\n"+
				"Array Length='%v' ",
				ePrefix.String(),
				idx,
				arrayLen)

	}

	return dMgrs.dirMgrs[idx].CopyOut(ePrefix.XCpy(
		fmt.Sprintf("dMgrs.dirMgrs[%v]",
			idx)))
}

// PeekFirstDirMgr
//
// Returns a deep copy of the first Directory Manager
// ('DirMgr') object in the current Directory Manager
// Collection ('DirMgrCollection').
//
// This is a 'Peek' method and therefore the original
// Directory Manager ('DirMgr') object is NOT deleted
// from the current Directory Manager Collection array
// ('DirMgrCollection').
//
// At the completion of this method, the length of the
// Directory Manager Collection array
// ('DirMgrCollection') array will remain unchanged.
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
//	DirMgr
//
//		If this method completes successfully, a deep
//		copy of the first Directory Manager object in the
//		current Directory Manager Collection will be
//		returned through this parameter.
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
func (dMgrs *DirMgrCollection) PeekFirstDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PeekFirstDirMgr()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	if len(dMgrs.dirMgrs) == 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The current DirMgrCollection is Empty!\n"+
				"The are zero Directory Manager objects ('DirMgr')\n"+
				"in the collection.\n",
				ePrefix.String())

	}

	return dMgrs.dirMgrs[0].CopyOut(ePrefix.XCpy(
		"dMgrs.dirMgrs[0]"))
}

// PeekLastDirMgr
//
// Returns a deep copy of the last Directory Manager
// object ('DirMgr') object in the current Directory
// Manager Collection ('DirMgrCollection').
//
// This is a 'Peek' method and therefore, the original
// Directory Manager ('DirMgr') object is NOT deleted
// from the Directory Manager Collection array
// ('DirMgrCollection').
//
// At the completion of this method, the length of the
// current Directory Manager Collection array
// ('DirMgrCollection') will remain unchanged.
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
//	DirMgr
//
//		If this method completes successfully, a deep
//		copy of the last Directory Manager object in the
//		current Directory Manager Collection will be
//		returned through this parameter.
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
func (dMgrs *DirMgrCollection) PeekLastDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"PeekLastDirMgr()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The current DirMgrCollection is Empty!\n"+
				"The are zero Directory Manager objects ('DirMgr')\n"+
				"in the collection.\n",
				ePrefix.String())

	}

	arrayLen--

	return dMgrs.dirMgrs[arrayLen].
		CopyOut(ePrefix.XCpy(
			fmt.Sprintf("dMgrs.dirMgrs[%v]",
				arrayLen)))
}
