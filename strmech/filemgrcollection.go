package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sort"
	"strings"
	"sync"
)

// SortFileMgrByAbsPathCaseSensitive - Sorts an array of File Managers
// (FileMgr) by absolute path, filename and file extension. This sorting
// operation is performed as a 'Case Sensitive' sort meaning the upper
// and lower case characters are significant.
//
// This method is designed to be used with the 'Go' Sort package:
//
//	https://golang.org/pkg/sort/
//
// Example Usage:
//
//	sort.Sort(SortFileMgrByAbsPathCaseSensitive(FileMgrArray))
type SortFileMgrByAbsPathCaseSensitive []FileMgr

// Len - Required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Len() int {
	return len(sortAbsPathSens)
}

// Swap - Required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Swap(i, j int) {
	sortAbsPathSens[i], sortAbsPathSens[j] = sortAbsPathSens[j], sortAbsPathSens[i]
}

// Less - required by the sort.Interface
func (sortAbsPathSens SortFileMgrByAbsPathCaseSensitive) Less(i, j int) bool {

	return sortAbsPathSens[i].absolutePathFileName < sortAbsPathSens[j].absolutePathFileName
}

// SortFileMgrByAbsPathCaseInSensitive - Sort by File Managers by
// absolute path, filename and file extension. This sorting operation
// is performed as a 'Case Insensitive' sort meaning the upper and
// lower case characters are not significant. All sort comparisons
// are therefore made by using lower case versions of the absolute
// path, filename and file extension.
//
// This method is designed to be used with the 'Go' Sort package:
//
//	https://golang.org/pkg/sort/
//
// Example Usage:
//
//	sort.Sort(SortFileMgrByAbsPathCaseInSensitive(FileMgrArray))
type SortFileMgrByAbsPathCaseInSensitive []FileMgr

// Len - Required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Len() int {
	return len(sortAbsPathInSens)
}

// Swap - Required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Swap(i, j int) {
	sortAbsPathInSens[i], sortAbsPathInSens[j] = sortAbsPathInSens[j], sortAbsPathInSens[i]
}

// Less - required by the sort.Interface
func (sortAbsPathInSens SortFileMgrByAbsPathCaseInSensitive) Less(i, j int) bool {

	return strings.ToLower(sortAbsPathInSens[i].absolutePathFileName) <
		strings.ToLower(sortAbsPathInSens[j].absolutePathFileName)
}

// FileMgrCollection - Manages a collection of FileMgr
// instances.
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

	fMgr := FileMgr{}

	isEmpty, err :=
		new(fileMgrHelperAtom).setFileMgrDirMgrFileName(
			&fMgr,
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

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr)

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
//		Manager collection maintained by the current
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

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	if idx < 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'idx' is less than zero.\n"+
			"Index Out-Of-Range!\n"+
			"idx='%v'\n",
			ePrefix.String(),
			idx)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return fmt.Errorf("%v\n"+
			"Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!\n",
			ePrefix.String())
	}

	if idx >= arrayLen {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'idx' is greater than the "+
			"last index in the File Manager Collection.\n"+
			"Index Out-Of-Range!\n"+
			"idx= '%v'\n"+
			"Last Array Index= '%v' ",
			ePrefix.String(),
			idx,
			arrayLen-1)
	}

	if arrayLen == 1 {
		fMgrs.fileMgrs = make([]FileMgr, 0, 100)
	} else if idx == 0 {
		// arrayLen > 1 and requested idx = 0
		fMgrs.fileMgrs = fMgrs.fileMgrs[1:]
	} else if idx == arrayLen-1 {
		// arrayLen > 1 and requested idx = last element index
		fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]
	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		fMgrs.fileMgrs =
			append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)
	}

	return err
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
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
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

	funcName := "FileMgrCollection.InsertFileMgrAtIndex()"

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
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

// New - Creates and returns a new, empty and properly initialized
// File Manager Collection ('FileMgrCollection').
func (fMgrs FileMgrCollection) New() FileMgrCollection {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	newFMgrCol := FileMgrCollection{}

	newFMgrCol.fileMgrs = make([]FileMgr, 0, 100)

	return newFMgrCol
}

// PopFileMgrAtIndex - Returns a deep copy of the File Manager
// ('FileMgr') object located at index, 'idx', in the
// File Manager Collection ('FileMgrCollection') array.
//
// As a 'Pop' method, the original File Manager ('FileMgr')
// object is deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// Therefore, at the completion of this method, the File Manager
// Collection array has a length which is one less than the
// starting array length.
func (fMgrs *FileMgrCollection) PopFileMgrAtIndex(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFileMgrAtIndex() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if idx < 0 {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter is less than zero. Index Out-Of-Range! idx='%v'", idx)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!\n")
	}

	if idx >= arrayLen {
		return FileMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter, 'idx' is greater than the length of the "+
			"collection index. Index Out-Of-Range! "+
			"idx='%v' Array Length='%v' ", idx, arrayLen)
	}

	if idx == 0 {
		return fMgrs.PopFirstFileMgr()
	}

	if idx == arrayLen-1 {
		return fMgrs.PopLastFileMgr()
	}

	fmgr := fMgrs.fileMgrs[idx].CopyOut()

	fMgrs.fileMgrs = append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)

	return fmgr, nil
}

// PopFirstFileMgr - Returns a deep copy of the first File Manager
// ('FileMgr') object in the File Manager Collection array. As a
// 'Pop' method, the original File Manager ('FileMgr') object is
// deleted from the File Manager Collection ('FileMgrCollection')
// array.
//
// Therefore, at the completion of this method, the File Manager
// Collection array has a length which is one less than the starting
// array length.
func (fMgrs *FileMgrCollection) PopFirstFileMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopFirstFileMgr() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if len(fMgrs.fileMgrs) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: The File Manager Collection, 'FileMgrCollection' is EMPTY!\n")
	}

	fMgr := fMgrs.fileMgrs[0].CopyOut()

	fMgrs.fileMgrs = fMgrs.fileMgrs[1:]

	return fMgr, nil
}

// PopLastFileMgr - Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection array. As a
// 'Pop' method, the original File Manager ('FileMgr') object is
// deleted from the File Manager Collection ('FileMgrCollection')
// array.
//
// Therefore, at the completion of this method, the File Manager
// Collection array has a length which is one less than the starting
// array length.
func (fMgrs *FileMgrCollection) PopLastFileMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PopLastFileMgr() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{}, errors.New(ePrefix +
			"Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!\n")
	}

	fmgr := fMgrs.fileMgrs[arrayLen-1].CopyOut()

	fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]

	return fmgr, nil
}

// PeekFileMgrAtIndex - Returns a deep copy of the File Manager
// ('FileMgr') object located at array index 'idx' in the File
// Manager Collection ('FileMgrCollection'). This is a 'Peek'
// method and therefore the original File Manager ('FileMgr')
// object is NOT deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// At the completion of this method, the length of the File
// Manager Collection ('FileMgrCollection') array will remain
// unchanged.
func (fMgrs *FileMgrCollection) PeekFileMgrAtIndex(idx int) (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekFileMgrAtIndex() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: The File Manager Collection, 'FileMgrCollection' is EMPTY!\n")
	}

	if idx < 0 {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= arrayLen {
		return FileMgr{},
			fmt.Errorf(ePrefix+
				"Error: Input Parameter 'idx' is greater than the length "+
				"of the collection array. "+
				"Index Out-Of-Range! idx='%v' Array Length='%v' ",
				idx, arrayLen)
	}

	return fMgrs.fileMgrs[idx].CopyOut(), nil
}

// PeekFirstFileMgr - Returns a deep copy of the first File
// Manager ('FileMgr') object in the File Manager Collection
// ('FileMgrCollection'). This is a 'Peek' method and therefore
// the original File Manager ('FileMgr') object is NOT deleted
// from the File Manager Collection ('FileMgrCollection')
// array.
//
// At the completion of this method, the length of the File
// Manager Collection ('FileMgrCollection') array will remain
// unchanged.
func (fMgrs *FileMgrCollection) PeekFirstFileMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekFirstFileMgr() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if len(fMgrs.fileMgrs) == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: The File Manager Collection ('FileMgrCollection') is EMPTY!\n")
	}

	return fMgrs.fileMgrs[0].CopyOut(), nil
}

// PeekLastFileMgr - Returns a deep copy of the last File Manager
// ('FileMgr') object in the File Manager Collection
// ('FileMgrCollection').
//
// This is a 'Peek' method and therefore the original File Manager
// ('FileMgr') object is NOT deleted from the File Manager Collection
// ('FileMgrCollection') array.
//
// At the completion of this method, the length of the File Manager
// Collection ('FileMgrCollection') array will remain unchanged.
func (fMgrs *FileMgrCollection) PeekLastFileMgr() (FileMgr, error) {

	ePrefix := "FileMgrCollection.PeekLastFileMgr()"

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{},
			errors.New(ePrefix +
				"Error: The File Manager Collection ('FileMgrCollection') is EMPTY!\n")
	}

	return fMgrs.fileMgrs[arrayLen-1].CopyOut(), nil
}

// SortByAbsPathFileName - Sorts the collection array of file managers
// by absolute path, file name and file extension.
//
// If the input parameter 'caseInsensitiveSort' is set to 'true', it means
// that upper and lower case characters are NOT significant in the sorting
// operation. The sort operation therefore uses lower case versions of
// absolute path, file name and file extension for comparison purposes.
//
// On the other hand, if input parameter 'caseInsensitiveSort' is set to 'false',
// it means that upper and lower chase characters ARE significant to the sort
// operation.
func (fMgrs *FileMgrCollection) SortByAbsPathFileName(caseInsensitiveSort bool) {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
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
