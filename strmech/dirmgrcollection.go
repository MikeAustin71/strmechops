package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
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
//
// Adds a DirMgr object to the collection.
//
// Note that this method does not perform a validity
// check on input parameter, 'dMgr'.
//
// It is recommended that dMgr.IsValidInstanceError() be called
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

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&dMgrCopy,
		&dMgr,
		ePrefix.XCpy(
			"dMgrCopy<-dMgr"))

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
//		If the 'dMgrs2' collection of Directory Manager
//		objects is empty, or has a zero length, this
//		method will exit and no error will be returned.
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

	dMgrs2 := new(dirMgrCollectionHelper).
		newEmptyDMgrCollection()

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

	var errStatus ArrayColErrorStatus

	_,
		errStatus = new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			idx,
			true, // Delete DirMgr object
			ePrefix.XCpy(
				"dMgrs"))

	return errStatus.ProcessingError
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

	dMgrs.lock = nil

	return
}

// Equal
//
// This method receives a pointer to an incoming instance
// of DirMgrCollection and proceeds to compare the
// encapsulated Directory Manager Collection with the
// Directory Manager Collection contained in the current
// instance of DirMgrCollection.
//
// If any of the Directory Manager (DirMgr) objects in
// the two collections are not equal, this method returns
// a boolean value of 'false'.
//
// A value of 'true' is only returned if all Directory
// Manager objects in both collections are equal in all
// respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingDMgrCollection		*DirMgrCollection
//
//		A pointer to an external instance of
//		DirMgrCollection. All the Directory Manager
//		objects in this Directory Manager Collection will
//		be compared to the Directory Manager Collection
//		contained in the current instance of
//		DirMgrCollection to determine if all the
//		Directory Manager objects are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		Two Directory Manager Collections from input
//		parameter 'incomingDMgrCollection' and the
//		current instance of DirMgrCollection are
//		compared to determine if they are equivalent.
//
//		If any of the corresponding Directory Manager
//		(DirMgr) objects in the two collections are not
//		equal, this method returns a boolean value of
//		'false'.
//
//		A value of 'true' is only returned if all
//	 	corresponding Directory Manager objects in both
//	 	collections are	equal in all respects.
func (dMgrs *DirMgrCollection) Equal(
	incomingDMgrCollection *DirMgrCollection) bool {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	return new(dirMgrCollectionHelper).
		equalDMgrCollections(
			dMgrs,
			incomingDMgrCollection)
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
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
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
//			RegularExp			*regexp.Regexp
//
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
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
//		If all the file selection criterion in the FileSelectionCriteria object are
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
		return new(dirMgrCollectionHelper).newEmptyDMgrCollection(), err
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	lDirCol := len(dMgrs.dirMgrs)

	if lDirCol == 0 {
		return new(dirMgrCollectionHelper).newEmptyDMgrCollection(), nil
	}

	fh := FileHelper{}

	var isMatchedFile bool

	dMgrs2 := new(dirMgrCollectionHelper).newEmptyDMgrCollection()

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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) GetDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	*DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	funcName := "DirMgrCollection." +
		"GetDirMgrAtIndex()"

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			funcName,
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return &DirMgr{}, errStatus
	}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {

		errStatus.IsArrayCollectionEmpty = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: This Directory Manager Collection ('DirMgrCollection') is EMPTY!\n",
				ePrefix.String())

		return &DirMgr{}, errStatus
	}

	if idx < 0 {

		errStatus.IsIndexOutOfBounds = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: The input parameter, 'idx', is OUT OF RANGE!\n"+
				"idx='%v'\n"+
				"The minimum index is '0'. \n",
				ePrefix.String(),
				idx)

		return &DirMgr{}, errStatus
	}

	if idx >= arrayLen {

		errStatus.IsIndexOutOfBounds = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: The input parameter, 'idx', is OUT OF RANGE!\n"+
				"idx='%v'\n"+
				"The maximum index is '%v'. ",
				ePrefix.String(),
				idx,
				arrayLen-1)

		return &DirMgr{}, errStatus
	}

	return &dMgrs.dirMgrs[idx], errStatus
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

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	return len(dMgrs.dirMgrs)
}

// GetTotalBytes
//
// Returns the total number of bytes contained in the
// directories which make up the Directory Collection
// encapsulated in the current instance of
// DirMgrCollection.
//
// The byte total represents only the files in the top
// level of each directory in the collection. It does
// not include the files in each directory's tree.
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
//	totalBytesInDirs			uint64
//
//		If this method completes successfully,
//		'totalBytesInDirs' will return the total number
//		of bytes contained within all files residing in
//		the top level of each directory in the Directory
//		Collection encapsulated by the current instance
//		of DirMgrCollection.
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
func (dMgrs *DirMgrCollection) GetTotalBytes(
	errorPrefix interface{}) (
	totalBytesInDirs uint64,
	err error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection.GetTotalBytes()",
		"")

	if err != nil {
		return totalBytesInDirs, err
	}

	lenDirCollection := len(dMgrs.dirMgrs)

	var localTotalBytes uint64

	for i := 0; i < lenDirCollection; i++ {

		localTotalBytes,
			err = dMgrs.dirMgrs[i].GetTotalBytes(
			ePrefix.XCpy(
				fmt.Sprintf("dMgrs.dirMgrs[%v]",
					i)))

		if err != nil {
			return totalBytesInDirs, err
		}

		totalBytesInDirs += localTotalBytes

		localTotalBytes = 0

	}

	return totalBytesInDirs, err
}

// GetTotalBytesCommaSeparated
//
// Returns the total number of bytes contained in the
// directories which make up the Directory Collection
// encapsulated in the current instance of
// DirMgrCollection.
//
// The byte total represents only the files in the top
// level of each directory in the collection. It does
// not include the files in each directory's tree.
//
// The returned numeric value for total bytes will be
// formatted as a number string meaning that thousands
// will be separated by commas.
//
//	Example: 1,645,321
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
//	totalBytesInDirs			string
//
//		If this method completes successfully,
//		'totalBytesInDirs' will return the total number
//		of bytes contained within all files residing in
//		the top level of each directory in the Directory
//		Collection encapsulated by the current instance
//		of DirMgrCollection. The returned numeric value
//		for total bytes will be formatted as a number
//		string meaning that thousands will be separated
//		by commas.
//
//			Example: 1,645,321
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
func (dMgrs *DirMgrCollection) GetTotalBytesCommaSeparated(
	errorPrefix interface{}) (
	totalBytesInDirs string,
	err error) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection.GetTotalBytes()",
		"")

	if err != nil {
		return totalBytesInDirs, err
	}

	lenDirCollection := len(dMgrs.dirMgrs)

	var localTotalBytes uint64
	var bytesSubTotal uint64

	for i := 0; i < lenDirCollection; i++ {

		localTotalBytes,
			err = dMgrs.dirMgrs[i].GetTotalBytes(
			ePrefix.XCpy(
				fmt.Sprintf("dMgrs.dirMgrs[%v]",
					i)))

		if err != nil {
			return totalBytesInDirs, err
		}

		bytesSubTotal += localTotalBytes

		localTotalBytes = 0

	}

	var intSep IntegerSeparatorSpec

	intSep,
		err = new(IntegerSeparatorSpec).
		NewUnitedStatesDefaults(
			ePrefix.XCpy(
				"intSep<-"))

	if err != nil {
		return "", err
	}

	totalBytesInDirs,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				bytesSubTotal),
			ePrefix.XCpy("<-bytesSubTotal"))

	if err != nil {
		return totalBytesInDirs, err
	}

	return totalBytesInDirs, err
}

// GetPathOriginalStrArray
//
// Converts the directories contained in the current
// instance of Directory Manager Collection
// (DirMgrCollection) to an array of strings returned by
// an instance of StringArrayDto.
//
// The Directory Manager Collection encapsulates an array
// of Directory Manager objects (DirMgr). These Directory
// Manager objects are convert to the original path used
// to initialize the Directory Manager instance. This
// returned path may/ be an absolute path or a relative
// path depending on how the DirMgr instance was
// initialized.
//
// If the current Directory Manager Collection instance
// is empty, containing zero directories, the returned
// instance of StringArrayDto is also empty.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
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
//	StringArrayDto
//
//		This returned instance of StringArrayDto
//		encapsulates a string array consisting of the
//		original directory paths stored in the current
//		instance of the Directory Manager Collection.
//
//		The original directory path is the path used to
//		initialize the Directory Manager instance
//		(DirMgr). This original path may be an absolute
//		path or a relative path depending on how the
//		particular Directory Manager instance was
//		initialized.
func (dMgrs *DirMgrCollection) GetPathOriginalStrArray() StringArrayDto {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	newStrArrayDto := StringArrayDto{}.New()

	lenDirMgrs := len(dMgrs.dirMgrs)

	var absPath string

	for i := 0; i < lenDirMgrs; i++ {

		absPath = dMgrs.dirMgrs[i].GetPathOriginal()

		if len(absPath) == 0 {

			absPath = fmt.Sprintf("Error DirMgr[%v]: %v",
				i,
				dMgrs.dirMgrs[i].absolutePath)

		}

		newStrArrayDto.PushStr(absPath)
	}

	return newStrArrayDto
}

// GetPathAbsoluteStrArray
//
// Converts the directories contained in the current
// instance of Directory Manager Collection
// (DirMgrCollection) to an array of strings returned by
// an instance of StringArrayDto.
//
// The Directory Manager Collection encapsulates an array
// of Directory Manager objects (DirMgr). These Directory
// Manager objects are converted to absolute directory
// paths and returned in a string array encapsulated in
// the returned instance of StringArrayDto.
//
// If the current Directory Manager Collection instance
// is empty, containing zero directories, the returned
// instance of StringArrayDto is also empty.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working
//	directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	StringArrayDto
//
//		This returned instance of StringArrayDto
//		encapsulates a string array consisting of the
//		absolute directory paths stored in the current
//		instance of the Directory Manager Collection.
func (dMgrs *DirMgrCollection) GetPathAbsoluteStrArray() StringArrayDto {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	newStrArrayDto := StringArrayDto{}.New()

	lenDirMgrs := len(dMgrs.dirMgrs)

	var absPath string

	for i := 0; i < lenDirMgrs; i++ {

		absPath = dMgrs.dirMgrs[i].GetPathAbsolute()

		if len(absPath) == 0 {

			absPath = fmt.Sprintf("Error DirMgr[%v]: %v",
				i,
				dMgrs.dirMgrs[i].absolutePath)

		}

		newStrArrayDto.PushStr(absPath)
	}

	return newStrArrayDto
}

// GetTextListingAbsPath
//
// Receives a pointer to a string builder and adds a text
// listing of the all directory paths in the Directory
// Manager Collection array encapsulated by the current
// DirMgrCollection instance.
//
// The returned Text Listing will include a title segment
// and a simple listing of all directory paths in the
// current Directory Manager Collection. The directory
// paths will be listed as absolute paths.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leftMargin					string
//
//		This string serves as the left margin for all
//		text lines.
//
//	rightMargin					string
//
//		This string serves as the right margin for all
//		text lines.
//
//
//	maxLineLength				int
//
//		This integer value defines the maximum line
//		length for all text lines. If this value is
//		less than 10, an error will be returned.
//
//	topTitleDisplay				TextLineTitleMarqueeDto
//
//		Contains specifications for the top tile display
//		including title lines and solid line breaks.
//
//		If no title is required, set this parameter to an
//		empty instance of TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no top title text lines are required, and the
//		solid line breaks are still necessary, simply
//		leave the title lines collection empty.
//
//	bottomTitleDisplay			TextLineTitleMarqueeDto
//
//		Contains specifications for the bottom tile
//		display including title lines and solid line
//		breaks.
//
//		If no bottom title is required, set this
//		parameter to an empty instance of
//		TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no bottom title text lines are required, and
//		the solid line breaks are still necessary, simply
//		leave the title lines collection empty.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder.
//		The text listing for directory absolute paths
//		will be added to this instance of
//		strings.Builder.
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
func (dMgrs *DirMgrCollection) GetTextListingAbsPath(
	leftMargin string,
	rightMargin string,
	maxLineLength int,
	topTitleDisplay TextLineTitleMarqueeDto,
	bottomTitleDisplay TextLineTitleMarqueeDto,
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirMgrCollection." +
		"GetTextListingAbsPath()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	return new(dirMgrCollectionHelper).
		fmtTextListingAbsPath(
			dMgrs,
			leftMargin,
			rightMargin,
			maxLineLength,
			topTitleDisplay,
			bottomTitleDisplay,
			strBuilder,
			ePrefix.XCpy("<-dMgrs"))
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

	lenDirMgrs := len(dMgrs.dirMgrs)

	if lenDirMgrs == 0 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 1)
	}

	err = new(dirMgrHelperPlanck).isDirMgrValid(
		&dMgr,
		"dMgr",
		ePrefix.XCpy("dMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: dMgr.IsValidInstanceError()"+
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

	var dMgrCopy DirMgr

	dMgrCopy,
		err = dMgr.CopyOut(ePrefix.XCpy(
		"dMgrCopy<-dMgr"))

	if err != nil {
		return err
	}

	if index >= lenDirMgrs {

		dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)

		return err
	}

	var newDirMgrs []DirMgr

	if index == 0 {

		newDirMgrs = make([]DirMgr, 0, 1)

		newDirMgrs = append(newDirMgrs, dMgrCopy)
		dMgrs.dirMgrs = append(newDirMgrs, dMgrs.dirMgrs...)

		return err
	}

	newDirMgrs = make([]DirMgr, 0, lenDirMgrs-index)

	newDirMgrs = append(
		newDirMgrs, dMgrs.dirMgrs[index:]...)

	dMgrs.dirMgrs = append(dMgrs.dirMgrs[:index])
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)
	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgrs...)

	return err
}

// New
//
// Creates and returns a new, empty and properly
// initialized Directory Manager Collection
// ('DirMgrCollection') instance.
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
//	DirMgrCollection
//
//		This method returns a concrete instance of
//		DirMgrCollection. The returned instance consists
//		of an empty and properly initialized instance of
//		DirMgrCollection.
func (dMgrs *DirMgrCollection) New() DirMgrCollection {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	return new(dirMgrCollectionHelper).newEmptyDMgrCollection()
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
// If an error is encountered, detailed error status
// information will be configured in the
// ArrayColErrorStatus instance returned by this method.
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
//		If this value is less than zero an error will be
//		returned through ArrayColErrorStatus.
//
//		If this value is greater than the last index in
//		the collection array, an error will be returned
//		through ArrayColErrorStatus.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PopDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PopDirMgrAtIndex()",
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			idx,
			true, // Delete 'idx' on exit
			ePrefix.XCpy(
				fmt.Sprintf("dMgrs[%v]", idx)))
}

// PopFirstDirMgr
//
// Returns a deep copy of the first Directory Manager
// ('DirMgr') object in the Directory Manager Collection
// array for the current DirMgrCollection instance.
//
// As a 'Pop' method, the original Directory Manager
// ('DirMgr') object is deleted from the first array
// index of the Directory Manager Collection
// (dMgrs.dirMgrs array index = 0).
//
// If this method is called on an empty Director Manager
// Collection (i.e. length of array dMgrs.dirMgrs = 0),
// detailed error status information will be configured
// in the ArrayColErrorStatus instance returned by this
// method.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PopFirstDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PopFirstDirMgr()",
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			0,
			true, // Delete index 0 on exit
			ePrefix.XCpy(
				"dMgrs[0]"))
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
// If an error is encountered, detailed error status
// information will be configured in the
// ArrayColErrorStatus instance returned by this method.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PopLastDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PopLastDirMgr()",
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	var lastArrayIndex = len(dMgrs.dirMgrs)

	if lastArrayIndex == 0 {

		errStatus.IsArrayCollectionEmpty = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: The Director Managers Collection\n"+
				"for the current DirMgrCollection instance\n"+
				"is EMPTY. The Collection array has an array\n"+
				"length of Zero.\n",
				ePrefix.String())

		return DirMgr{}, errStatus
	}

	lastArrayIndex--

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			lastArrayIndex,
			true, // Delete Last Index in array
			ePrefix.XCpy(
				fmt.Sprintf(
					"dMgrs[%v]",
					lastArrayIndex)))
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
// If an error is encountered, detailed error status
// information will be configured in the
// ArrayColErrorStatus instance returned by this method.
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
//		If this value is less than zero an error will be
//		returned.
//
//		If 'idx' exceeds the last index in the collection,
//		an io.EOF (End-Of-File) error will be returned.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PeekDirMgrAtIndex()",
			"")

	if errStatus.ProcessingError != nil {
		return DirMgr{}, errStatus
	}

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			idx,
			false, // Do NOT Delete 'idx'
			ePrefix.XCpy(
				fmt.Sprintf(
					"dMgrs[%v]",
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
// If an error is encountered, detailed error status
// information will be configured in the
// ArrayColErrorStatus instance returned by this method.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PeekFirstDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PeekFirstDirMgr()",
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			0,
			false, // Do NOT Delete 'dMgrs.dirMgrs[0]'
			ePrefix.XCpy(
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
// If an error is encountered, detailed error status
// information will be configured in the
// ArrayColErrorStatus instance returned by this method.
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
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrs *DirMgrCollection) PeekLastDirMgr(
	errorPrefix interface{}) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrs.lock == nil {
		dMgrs.lock = new(sync.Mutex)
	}

	dMgrs.lock.Lock()

	defer dMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewIEmpty(
			errorPrefix,
			"DirMgrCollection."+
				"PeekLastDirMgr()",
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	lastArrayIndex := len(dMgrs.dirMgrs)

	if lastArrayIndex == 0 {

		errStatus.IsArrayCollectionEmpty = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: The Director Managers Collection\n"+
				"for the current DirMgrCollection instance\n"+
				"is EMPTY. The Collection array has an array\n"+
				"length of Zero.\n",
				ePrefix.String())

		return DirMgr{}, errStatus
	}

	lastArrayIndex--

	return new(dirMgrCollectionHelper).
		peekOrPopAtIndex(
			dMgrs,
			lastArrayIndex,
			false, // Do NOT Delete 'dMgrs.dirMgrs[lastArrayIndex]'
			ePrefix.XCpy(
				fmt.Sprintf("dMgrs.dirMgrs[%v]",
					lastArrayIndex)))
}
