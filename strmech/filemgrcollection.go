package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sort"
	"strings"
	"sync"
)

// SortFileMgrByAbsPathCaseSensitive
//
// Sorts an array of File Managers (FileMgr) by absolute
// path, filename and file extension. This sorting
// operation is performed as a 'Case Sensitive' sort
// meaning the upper and lower case characters are
// significant.
//
// This method is designed to be used with the 'Go' Sort
// package:
//
//	https://golang.org/pkg/sort/
//
// ----------------------------------------------------------------
//
// # Usage
//
//	sort.Sort(SortFileMgrByAbsPathCaseSensitive(FileMgrArray))
type SortFileMgrByAbsPathCaseSensitive []FileMgr

// Len
//
// Required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Len() int {
	return len(sortAbsPathSens)
}

// Swap
// Required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Swap(i, j int) {
	sortAbsPathSens[i], sortAbsPathSens[j] = sortAbsPathSens[j], sortAbsPathSens[i]
}

// Less
//
// Required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Less(i, j int) bool {

	return sortAbsPathSens[i].absolutePathFileName < sortAbsPathSens[j].absolutePathFileName
}

// SortFileMgrByAbsPathCaseInSensitive
//
// Sort by File Managers by absolute path, filename and
// file extension. This sorting operation is performed as
// a 'Case Insensitive' sort meaning the upper and lower
// case characters are not significant. All sort
// comparisons are therefore made by using lower case
// versions of the absolute path, filename and file
// extension.
//
// This method is designed to be used with the 'Go' Sort
// package:
//
//	https://golang.org/pkg/sort/
//
// ----------------------------------------------------------------
//
// # Usage
//
//	sort.Sort(SortFileMgrByAbsPathCaseInSensitive(FileMgrArray))
type SortFileMgrByAbsPathCaseInSensitive []FileMgr

// Len
//
// Required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Len() int {
	return len(sortAbsPathInSens)
}

// Swap
//
// Required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Swap(i, j int) {
	sortAbsPathInSens[i], sortAbsPathInSens[j] = sortAbsPathInSens[j], sortAbsPathInSens[i]
}

// Less
//
// Required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Less(i, j int) bool {

	return strings.ToLower(sortAbsPathInSens[i].absolutePathFileName) <
		strings.ToLower(sortAbsPathInSens[j].absolutePathFileName)
}

// FileMgrCollection
//
// Manages a collection of FileMgr instances.
//
// Dependencies:
// 'FileMgrCollection' depends on type, 'FileHelper'
// which is located in source code file 'filehelper.go'.
type FileMgrCollection struct {
	fileMgrs []FileMgr

	lock *sync.Mutex
}

// AddFileMgr
//
// Adds a FileMgr object to the File Manager Collection
// maintained by the current instance of
// FileMgrCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No validation is performed on input parameter 'fMgr'
//	before a deep copy of 'fMgr' is added to the File
//	Manager Collection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						FileMgr
//
//		A concrete instance of FileMgr. A deep copy of
//		this FileMgr object will be added to the File
//		Manager Collection maintained by the current
//		instance of FileMgrCollection.
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
func (fMgrs *FileMgrCollection) AddFileMgr(
	fMgr FileMgr,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"AddFileMgr()",
		"")

	if err != nil {
		return err
	}

	// Adds a deep copy of 'fMgr'
	return new(FileMgrCollectionElectron).addFileMgr(
		fMgrs,
		&fMgr,
		ePrefix.XCpy(
			"fMgrs<-fMgr"))
}

// AddFileMgrByDirFileNameExt
//
// Adds a new File Manager using input parameters
// 'directory' and 'pathFileNameExt'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directory					DirMgr
//
//		An instance of Directory Manager ('DirMgr') which
//		contains the directory or folder path to be used
//		in constructing the instance of File Manager
//		(FileMgr) which will be added to the current File
//		Manager Collection (FileMgrCollection).
//
//	fileNameExt					string
//
//		This string holds the file name and file
//		extension which will be combined with the
//		directory path from input parameter 'directory'
//		to create the File Manager (FileMgr) object which
//		will be added to the current File Manager
//		Collection (FileMgrCollection).
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
func (fMgrs *FileMgrCollection) AddFileMgrByDirFileNameExt(
	directory DirMgr,
	fileNameExt string,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"AddFileMgrByDirFileNameExt()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	newFMgr := FileMgr{}

	isEmpty, err :=
		new(fileMgrHelperAtom).setFileMgrDirMgrFileName(
			&newFMgr,
			&directory,
			fileNameExt,
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"Error: The FileMgr instance generated by input parameters 'directory' and "+
			"'fileNameExt' is Empty!\n"+
			"directory='%v'\n"+
			"fileNameExt='%v'\n",
			ePrefix.String(),
			directory.absolutePath,
			fileNameExt)
	}

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, newFMgr)

	return nil
}

// AddFileMgrByPathFileNameExt
//
// Add a new File Manager based on input parameter
// 'pathFileNameExt' which includes the full path name,
// file name and file extension.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string holds the path name, file name and
//		file extension which will be used to construct an
//		instance of File Manager (FileMgr) added to the
//		current File Manager Collection.
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
func (fMgrs *FileMgrCollection) AddFileMgrByPathFileNameExt(
	pathFileNameExt string,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"AddFileMgrByPathFileNameExt()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	fMgr := FileMgr{}

	var isEmpty bool

	isEmpty,
		err = new(fileMgrHelper).
		setFileMgrPathFileName(
			&fMgr,
			pathFileNameExt,
			ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned from fMgrHlpr.setFileMgrPathFileName(pathFileNameExt).\n"+
			"pathFileNameExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			pathFileNameExt,
			err.Error())
	}

	if isEmpty {

		return fmt.Errorf("%v\n"+
			"ERROR: The generated File Manager instance is EMPTY!\n"+
			"pathFileNameExt='%v'\n",
			ePrefix.String(),
			pathFileNameExt)
	}

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

	return err
}

// AddFileMgrByDirStrFileNameStr
//
// Adds a FileMgr object to the File Manager Collection
// based on input parameter strings, 'pathName' and
// 'fileNameExt'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathName					string
//
//		This string holds the directory path which will
//		be combined with the file name and file extension
//		provided by input paramter 'fileNameExt' to
//		construct the File Manager (FileMgr) instance
//		added to the current File Manager Collection.
//
//	fileNameExt					string
//
//		This strings holds the file name and file
//		extension which will be combined with the
//		directory path provided by input parameter
//		'pathName' to construct the File Manager
//		(FileMgr) instance added to the current File
//		Manager Collection.
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
func (fMgrs *FileMgrCollection) AddFileMgrByDirStrFileNameStr(
	pathName string,
	fileNameExt string,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"AddFileMgrByDirStrFileNameStr()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	dMgr := DirMgr{}

	isEmpty, err := new(dirMgrHelperNanobot).
		setDirMgr(
			&dMgr,
			pathName,
			"dMgr",
			"pathName",
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"ERROR: Directory Manager created "+
			"from 'pathName' is EMPTY!\n"+
			"pathName= '%v'\n",
			ePrefix.String(),
			pathName)
	}

	fMgr := FileMgr{}

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgr,
			&dMgr,
			fileNameExt,
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"ERROR: File Manager created "+
			"from 'pathName' and 'fileNameExt' is EMPTY!\n"+
			"pathName='%v'\n"+
			"fileNameExt='%v'\n",
			ePrefix.String(),
			pathName,
			fileNameExt)
	}

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

	return nil
}

// AddFileMgrByFileInfo
//
// Adds a File Manager object to the File Manager
// Collection based on input from a directory path string
// and an os.FileInfo object.
//
//	errorPrefix					interface{}
//
//	pathName					string
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
func (fMgrs *FileMgrCollection) AddFileMgrByFileInfo(
	pathName string,
	fileInfo os.FileInfo,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"AddFileMgrByFileInfo()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	dMgr := DirMgr{}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelperNanobot).
		setDirMgr(
			&dMgr,
			pathName,
			"dMgr",
			"pathName",
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"ERROR: Directory Manager created "+
			"from 'pathName' is EMPTY!\n"+
			"pathName= '%v'\n",
			ePrefix.String(),
			pathName)
	}

	fMgr := FileMgr{}

	isEmpty,
		err = new(fileMgrHelperAtom).
		setFileMgrDirMgrFileName(
			&fMgr,
			&dMgr,
			fileInfo.Name(),
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {
		return fmt.Errorf("%v\n"+
			"ERROR: File Manager created "+
			"from 'pathName' and 'fileInfo' is EMPTY!\n"+
			"pathName='%v'\n"+
			"fileInfo.Name()='%v'\n",
			ePrefix.String(),
			pathName,
			fileInfo.Name())
	}

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

	return nil
}

// AddFileMgrCollection
//
// Adds another collection of File Manager (FileMgr)
// objects to the File Manager Collection maintained
// by the current instance of FileMgrCollection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrs2						*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection.
//
//		Deep copies of all the File Manager objects
//		contained in this File Manager Collection, will
//		be added
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
func (fMgrs *FileMgrCollection) AddFileMgrCollection(
	fMgrs2 *FileMgrCollection,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileMgrCollection." +
		"AddFileMgrCollection()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	if fMgrs2.fileMgrs == nil {
		fMgrs2.fileMgrs = make([]FileMgr, 0, 5)
	}

	lOmc2 := len(fMgrs2.fileMgrs)

	if lOmc2 == 0 {
		return err
	}

	fMgrColElectron := FileMgrCollectionElectron{}

	for i := 0; i < lOmc2; i++ {

		// Adds a deep copy of 'fMgr'
		err = fMgrColElectron.addFileMgr(
			fMgrs,
			&fMgrs2.fileMgrs[i],
			ePrefix.XCpy(
				"fMgrs<-fMgr"))

		if err != nil {

			return fmt.Errorf(
				"%v\n"+
					"Error returned by fMgrColElectron.addFileMgr(fMgrs2.fileMgrs[i])\n"+
					"i Loop Number: '%v'\n"+
					"Error = \n%v\n",
				funcName,
				i,
				err.Error())
		}

	}

	return err
}

// CopyFilesToDir
//
// Copies all the files in the File Manager Collection to
// the specified target directory.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDirectory				DirMgr
//
//		A concrete instance of DirMgr. This instance
//		specifies the target directory where all files
//		in the File Manager Collection, encapsulated by
//		the current FileMgrCollection instance, will be
//		copied.
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
func (fMgrs *FileMgrCollection) CopyFilesToDir(
	targetDirectory DirMgr,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileMgrCollection." +
		"CopyFilesToDir()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	maxLen := len(fMgrs.fileMgrs)

	if maxLen == 0 {
		return fmt.Errorf("%v\n"+
			"ERROR - Collection contains ZERO File Managers!\n",
			ePrefix.String())
	}

	for i := 0; i < maxLen; i++ {

		err = fMgrs.fileMgrs[i].CopyFileToDirByIoByLink(
			targetDirectory,
			ePrefix.XCpy("targetDirectory"))

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Copy Failure on 'fileMgrs' index='%v'\n"+
				"file='%v'."+
				"Error='%v'",
				funcName,
				i,
				fMgrs.fileMgrs[i].absolutePathFileName,
				err.Error())
		}

	}

	return err
}

// CopyOut
//
// Returns a new FileMgrCollection which is an exact
// duplicate of the current instance of
// FileMgrCollection.
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
//	FileMgrCollection
//
//		If this method completes successfully without
//		errors, a deep copy of the current
//		FileMgrCollection instance will be returned
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
func (fMgrs *FileMgrCollection) CopyOut(
	errorPrefix interface{}) (
	FileMgrCollection,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"CopyOut()",
		"")

	if err != nil {
		return FileMgrCollection{}, err
	}

	fMgrs2 := FileMgrCollection{}

	fMgrs2.fileMgrs = make([]FileMgr, 0, 5)

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	lOmc := len(fMgrs.fileMgrs)

	if lOmc == 0 {
		return FileMgrCollection{},
			fmt.Errorf("%v\n"+
				"Error: This File Manager Collection ('FileMgrCollection') is EMPTY!\n",
				ePrefix.String())
	}

	fMgrColElectron := FileMgrCollectionElectron{}

	for i := 0; i < lOmc; i++ {

		err = fMgrColElectron.addFileMgr(
			&fMgrs2,
			&fMgrs.fileMgrs[i],
			ePrefix.XCpy(
				fmt.Sprintf("fMgrs.fileMgrs[%v]",
					i)))

		if err != nil {
			return fMgrs2, err
		}
	}

	return fMgrs2, err
}

// DeleteAtIndex
//
// This method deletes an array element from the File
// Manager Collection maintained by the current instance
// of FileMgrCollection.
//
// The index of the array element to be deleted is
// specified by input parameter 'idx'.
//
// If this method completes successfully, the File
// Manager Collection array will have a length which is
// one less than the starting array length.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be deleted from the File
//		Manager Collection maintained by the current
//		instance of FileMgrCollection.
//
//		If this value is less than zero or greater than
//		the last index in the array, an error will be
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
func (fMgrs *FileMgrCollection) DeleteAtIndex(
	idx int,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"DeleteAtIndex()",
		"")

	if err != nil {
		return err
	}

	return new(FileMgrCollectionElectron).
		deleteAtIndex(
			fMgrs,
			idx,
			ePrefix.XCpy(
				fmt.Sprintf("fMgrs[%v]", idx)))
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of FileMgrCollection to their initial or
//	zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing FileMgr
//	objects contained in the File Manager Collection
//	encapsulated in the current instance of
//	FileMgrCollection. Upon completion of this method,
//	the internal File Manager Collection array for the
//	current FileMgrCollection instance will have a length
//	of zero.
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
func (fMgrs *FileMgrCollection) Empty() {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	colArrayLen := len(fMgrs.fileMgrs)

	for i := 0; i < colArrayLen; i++ {

		fMgrs.fileMgrs[i].Empty()

	}

	fMgrs.fileMgrs = make([]FileMgr, 0)

	fMgrs.lock.Unlock()

	fMgrs.lock = nil

	return
}

// Equal
//
// This method receives a pointer to an incoming instance
// of FileMgrCollection and proceeds to compare the
// encapsulated File Manager Collection with the
// File Manager Collection contained in the current
// instance of FileMgrCollection.
//
// If any of the File Manager (FileMgr) objects in
// the two collections are not equal, this method returns
// a boolean value of 'false'.
//
// A value of 'true' is only returned if all File
// Manager objects in both collections are equal in all
// respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingDMgrCollection		*FileMgrCollection
//
//		A pointer to an external instance of
//		FileMgrCollection. All the File Manager
//		objects in this File Manager Collection will
//		be compared to the File Manager Collection
//		contained in the current instance of
//		FileMgrCollection to determine if all the
//		File Manager objects are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		Two File Manager Collections from input
//		parameter 'incomingDMgrCollection' and the
//		current instance of FileMgrCollection are
//		compared to determine if they are equivalent.
//
//		If any of the corresponding File Manager
//		(FileMgr) objects in the two collections are not
//		equal, this method returns a boolean value of
//		'false'.
//
//		A value of 'true' is only returned if all
//	 	corresponding File Manager objects in both
//	 	collections are equal in all respects.
func (fMgrs *FileMgrCollection) Equal(
	incomingFMgrCollection *FileMgrCollection) bool {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	colArrayLen := len(fMgrs.fileMgrs)

	if colArrayLen != len(incomingFMgrCollection.fileMgrs) {

		return false
	}

	for i := 0; i < colArrayLen; i++ {

		if !fMgrs.fileMgrs[i].Equal(
			&incomingFMgrCollection.fileMgrs[i]) {

			return false

		}

	}

	return true
}

// FindFiles
//
// Searches the current FileMgrCollection and returns a
// new FileMgrCollection containing FileMgr objects which
// match the specified search criteria.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories in the
//		'startPath' will be searched, and all files
//		within those directories WILL BE DELETED.
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
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria are described
//		below:
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
//	FileMgrCollection
//
//		If this method completes successfully without
//		errors, this parameter will return a new instance
//		of FileMgrCollection containing an array of File
//		Managers (FileMgr) identifying the files which
//		match the file selection criteria specified by
//		input parameter 'fileSelectionCriteria'.
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
func (fMgrs *FileMgrCollection) FindFiles(
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	FileMgrCollection,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileMgrCollection.FindFiles()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return FileMgrCollection{}, err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	lDirCol := len(fMgrs.fileMgrs)

	if lDirCol == 0 {
		return FileMgrCollection{}.New(), nil
	}

	fh := FileHelper{}

	var isMatchedFile bool

	fMgrs2 := FileMgrCollection{}.New()

	var fMgr FileMgr

	for i := 0; i < lDirCol; i++ {

		fMgr = fMgrs.fileMgrs[i]

		if fMgr.actualFileInfo.isFInfoInitialized {

			isMatchedFile,
				err,
				_ = fh.FilterFileName(
				&fMgr.actualFileInfo,
				fileSelectCriteria,
				ePrefix.XCpy(
					"fMgr.actualFileInfo"))

			if err != nil {

				return FileMgrCollection{},
					fmt.Errorf("%v\n"+
						"Error returned by "+
						"fh.FilterFileName(fMgr.actualFileInfo, fileSelectCriteria) "+
						"fMgr.actualFileInfo.Name()='%v'\n"+
						"'i' Loop Number: '%v'\n"+
						"Error= \n%v\n",
						funcName,
						fMgr.actualFileInfo.Name(),
						i,
						err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(fMgr.fileNameExt)

			isMatchedFile,
				err,
				_ = fh.FilterFileName(
				&fip,
				fileSelectCriteria,
				ePrefix)

			if err != nil {
				return FileMgrCollection{},
					fmt.Errorf("%v\n"+
						"Error returned by fh.FilterFileName(fip, fileSelectCriteria)\n"+
						"fip.Name()='%v'\n"+
						"'i' Loop Number: '%v'\n"+
						"Error= \n%v\n",
						funcName,
						fip.Name(),
						i,
						err.Error())
			}

		}

		if isMatchedFile {

			// Adds a deep copy of fMgr
			err = new(FileMgrCollectionElectron).addFileMgr(
				&fMgrs2,
				&fMgr,
				ePrefix)

			if err != nil {
				return FileMgrCollection{},
					fmt.Errorf("%v\n"+
						"Error returned by FileMgrCollectionElectron).addFileMgr(fMgrs2<-fMgr)\n"+
						"Error adding matched file to fMgrs2 Collection.\n"+
						"'i' Loop Number: '%v'\n"+
						"Error= \n%v\n",
						funcName,
						i,
						err.Error())
			}

		}

	}

	return fMgrs2, err
}

// GetFileMgrArray
//
// Returns a deep copy of the entire File Manager Array
// generated from the File Manager Collection maintained
// by the current instance of FileMgrCollection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]FileMgr
//		A deep copy of the File Manager array maintained
//		by the current instance of FileMgrCollection.
func (fMgrs *FileMgrCollection) GetFileMgrArray() []FileMgr {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	return fMgrs.fileMgrs
}

// GetNumOfFileMgrs
//
// This method returns the array length of the File
// Manager Collection maintained by the current
// instance of FileMgrCollection.
//
// Effectively, the returned integer is a count of the
// number of File Managers objects (FileMgr's) in the
// File Manager Collection.
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
//	int
//
//		This integer value specifies the number of File
//		Manager (FileMgr) objects in the File Manager
//		Collection maintained by the current instance of
//		FileMgrCollection.
func (fMgrs *FileMgrCollection) GetNumOfFileMgrs() int {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	return len(fMgrs.fileMgrs)
}

// GetTextListing
//
// Receives a pointer to a string builder and adds a text
// listing of the all files in the File Manager
// Collection array encapsulated by the current
// FileMgrCollection instance.
//
// The returned Text Listing will include a title segment
// and a simple listing of all file paths and file names
// in the current File Manager Collection. The file paths
// will be displayed as absolute paths.
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
func (fMgrs *FileMgrCollection) GetTextListing(
	leftMargin string,
	rightMargin string,
	maxLineLength int,
	topTitleDisplay TextLineTitleMarqueeDto,
	bottomTitleDisplay TextLineTitleMarqueeDto,
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"GetTextListing()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

	}

	if maxLineLength < 10 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is less than '10'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	if maxLineLength > 1999 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is greater than '1999'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	return new(FileMgrCollectionMolecule).
		fmtTextListingAllFiles(
			fMgrs,
			leftMargin,
			rightMargin,
			maxLineLength,
			topTitleDisplay,
			bottomTitleDisplay,
			strBuilder,
			ePrefix.XCpy("fMgrs"))
}

// GetTotalFileBytes
//
// Returns the total number of file bytes represented by
// all files in the File Manager Collection maintained by
// the current instance of FileMgrCollection.
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
//	uint64
func (fMgrs *FileMgrCollection) GetTotalFileBytes() uint64 {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
		return 0
	}

	totalFileBytes := uint64(0)

	for i := 0; i < len(fMgrs.fileMgrs); i++ {

		if fMgrs.fileMgrs[i].actualFileInfo.isFInfoInitialized {

			totalFileBytes += uint64(fMgrs.fileMgrs[i].actualFileInfo.Size())

		}
	}

	return totalFileBytes
}

// InsertFileMgrAtIndex
//
// Inserts a new File Manager into the File Manager
// Collection maintained by the current instance of
// FileMgrCollection.
//
// A deep copy of the File Manager object passed as input
// parameter 'fMgr' will be inserted into the File
// Manager Collection at the array index specified by
// input parameter 'index'.
//
// If input parameter 'index' is less than zero, an error
// will be returned.
//
// If 'index' exceeds the value of the last index in the
// collection, 'fMgr' will be added to the end of the
// collection at the next legal index.
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
func (fMgrs *FileMgrCollection) InsertFileMgrAtIndex(
	fMgr FileMgr,
	index int,
	errorPrefix interface{}) error {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection.InsertFileMgrAtIndex()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 1)
	}

	if index < 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'index' is LESS THAN ZERO!\n"+
			"index= '%v'\n",
			ePrefix.String(),
			index)
	}

	lenfMgrs := len(fMgrs.fileMgrs)

	var deepCopyFileMgr FileMgr

	deepCopyFileMgr,
		err = fMgr.CopyOut(ePrefix.XCpy(
		"deepCopyFileMgr<-fMgr"))

	if err != nil {
		return err
	}

	if index >= lenfMgrs {

		fMgrs.fileMgrs = append(
			fMgrs.fileMgrs, deepCopyFileMgr)

		return err
	}

	newFileMgrs := make([]FileMgr, 0, 1)

	if index == 0 {

		newFileMgrs = append(newFileMgrs,
			deepCopyFileMgr)

		fMgrs.fileMgrs = append(
			newFileMgrs,
			fMgrs.fileMgrs...)

		return err
	}

	newFileMgrs = append(
		newFileMgrs, fMgrs.fileMgrs[index:]...)

	fMgrs.fileMgrs = append(fMgrs.fileMgrs[:index])

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, deepCopyFileMgr)

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, newFileMgrs...)

	return err
}

// New
//
// Creates and returns a new, empty and properly
// initialized File Manager Collection
// (FileMgrCollection).
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
//	FileMgrCollection
//
//		This method returns a new, empty and properly
//		initialized File Manager Collection.
func (fMgrs FileMgrCollection) New() FileMgrCollection {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	newFMgrCol := FileMgrCollection{}

	newFMgrCol.fileMgrs = make([]FileMgr, 0, 5)

	newFMgrCol.lock = new(sync.Mutex)

	return newFMgrCol
}

// PopFileMgrAtIndex
//
// Returns a deep copy of the File Manager ('FileMgr')
// object located at index, 'idx', in the File Manager
// Collection ('FileMgrCollection') array for the current
// instance of FileMgrCollection.
//
// As a 'Pop' method, the original File Manager
// ('FileMgr') object residing at the File Collection
// array index identified by input parameter 'idx' WILL
// BE DELETED.
//
// At the completion of this method, the File Manager
// Collection for the current instance of
// FileMgrCollection will have an array length which is
// one less than the starting array length.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the File Manager object
//	in the File Manager Collection residing at the
//	array index specified by input parameter 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be deleted from the File
//		Manager Collection encapsulated by the instance
//		of FileMgrCollection passed by input parameter
//		'fMgrs'.
//
//		If this value is less than zero or greater than
//		the last index in the array, an error will be
//		returned.
//
//		A deep copy of the deleted File Manager object
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at array index 'idx' will be returned
//		through this parameter.
//
//		Remember that the original File Manager object
//		will be deleted from the File Manager Collection
//		maintained by the current instance of
//		FileMgrCollection.
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
func (fMgrs *FileMgrCollection) PopFileMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PopFileMgrAtIndex()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			idx,
			true,
			ePrefix.XCpy(
				fmt.Sprintf(
					"fMgrs[%v]",
					idx)))
}

// PopFirstFileMgr
//
// Returns a deep copy of the first File Manager
// ('FileMgr') object in the File Manager Collection
// array maintained by the current instance of
// FileMgrCollection.
//
// As a 'Pop' method, the original File Manager
// ('FileMgr') object residing at the File Manager
// Collection array index zero('0'), WILL BE DELETED.
//
// Therefore, at the completion of this method, the File
// Manager Collection array will have a length which is
// one less than the starting array length.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the File Manager object
//	in the File Manager Collection residing at array
//	index zero ('0').
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at the first array index in the File
//		Manager Collection will be returned through
//		this parameter.
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
func (fMgrs *FileMgrCollection) PopFirstFileMgr(
	errorPrefix interface{}) (FileMgr, error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PopFirstFileMgr()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			0,
			true,
			ePrefix.XCpy(
				"fMgrs[0]"))
}

// PopLastFileMgr
//
// Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection
// array. As a 'Pop' method, the original File Manager
// ('FileMgr') object is deleted from the File Manager
// Collection ('FileMgrCollection') array.
//
// At the completion of this method, the File Manager
// Collection array will have a length which is one less
// than the starting array length.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the File Manager object
//	residing at the last array index in the File Manager
//	Collection.
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at the last array index in the File
//		Manager Collection will be returned through
//		this parameter.
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
func (fMgrs *FileMgrCollection) PopLastFileMgr(
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PopLastFileMgr()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	lastIndex := len(fMgrs.fileMgrs) - 1

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			lastIndex,
			true,
			ePrefix.XCpy(
				fmt.Sprintf(
					"fMgrs[%v]",
					lastIndex)))
}

// PeekFileMgrAtIndex
//
// Returns a deep copy of the File Manager ('FileMgr')
// object located at array index 'idx' in the File
// Manager Collection maintained by the current instance
// of FileMgrCollection.
//
// This is a 'Peek' method and therefore the original
// File Manager ('FileMgr') object WILL NOT BE DELETED
// from the File Manager Collection ('FileMgrCollection')
// array.
//
// At the completion of this method, the length of the
// File Manager Collection ('FileMgrCollection') array
// will remain unchanged.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method WILL NOT DELETE the File Manager
//	('FileMgr') object located at the array index
//	specified by input parameter 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be deleted from the File
//		Manager Collection encapsulated by the instance
//		of FileMgrCollection passed by input parameter
//		'fMgrs'.
//
//		If this value is less than zero or greater than
//		the last index in the array, an error will be
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at array index 'idx' will be returned
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
func (fMgrs *FileMgrCollection) PeekFileMgrAtIndex(
	idx int,
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PeekFileMgrAtIndex()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			idx,
			false,
			ePrefix.XCpy(
				fmt.Sprintf(
					"fMgrs[%v]",
					idx)))
}

// PeekFirstFileMgr
//
// Returns a deep copy of the first File Manager
// ('FileMgr') object in the File Manager Collection
// maintained by the current instance of
// FileMgrCollection.
//
// This is a 'Peek' method and therefore the original
// File Manager ('FileMgr') object WILL NOT BE DELETED
// from the File Manager Collection ('FileMgrCollection')
// array.
//
// At the completion of this method, the length of the
// File Manager Collection ('FileMgrCollection') array
// will remain unchanged.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method WILL NOT DELETE the File Manager
//	('FileMgr') object located at the first array index
//	in the File Manager Collection.
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at the first array index in the File
//		Manager Collection will be returned through this
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
func (fMgrs *FileMgrCollection) PeekFirstFileMgr(
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PeekFirstFileMgr()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			0,
			false,
			ePrefix.XCpy(
				"fMgrs[0]"))
}

// PeekLastFileMgr
//
// Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection
// ('FileMgrCollection').
//
// This is a 'Peek' method and therefore the original
// File Manager ('FileMgr') object WILL NOT BE DELETED
// from the File Manager Collection ('FileMgrCollection')
// array.
//
// At the completion of this method, the length of the
// File Manager Collection ('FileMgrCollection') array
// will remain unchanged.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method WILL NOT DELETE the File Manager
//	('FileMgr') object located at the last array index in
//	the File Manager Collection.
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at the last array index in the File
//		Manager Collection will be returned through this
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
func (fMgrs *FileMgrCollection) PeekLastFileMgr(
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"PeekFirstFileMgr()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	lastIndex := len(fMgrs.fileMgrs) - 1

	return new(FileMgrCollectionMolecule).
		peekOrPopAtIndex(
			fMgrs,
			lastIndex,
			false,
			ePrefix.XCpy(
				fmt.Sprintf(
					"fMgrs[%v]",
					lastIndex)))
}

// SortByAbsPathFileName
//
// Sorts the collection array of file managers by
// absolute path, file name and file extension.
//
// If the input parameter 'caseInsensitiveSort' is set to
// 'true', it means that upper and lower case characters
// are NOT significant in the sorting operation. The sort
// operation therefore uses lower case versions of
// absolute path, file name and file extension for
// comparison purposes.
//
// If input parameter 'caseInsensitiveSort' is set to
// 'false', it means that upper and lower case characters
// ARE significant to the sort operation.
func (fMgrs *FileMgrCollection) SortByAbsPathFileName(
	caseInsensitiveSort bool) {

	if fMgrs.lock == nil {
		fMgrs.lock = new(sync.Mutex)
	}

	fMgrs.lock.Lock()

	defer fMgrs.lock.Unlock()

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	if len(fMgrs.fileMgrs) == 0 {
		return
	}

	if caseInsensitiveSort {

		sort.Sort(SortFileMgrByAbsPathCaseInSensitive(fMgrs.fileMgrs))

	} else {

		sort.Sort(SortFileMgrByAbsPathCaseSensitive(fMgrs.fileMgrs))

	}
}
