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

// AddFileMgr - Adds a FileMgr object to the collection
func (fMgrs *FileMgrCollection) AddFileMgr(fMgr FileMgr) {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
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

	dMgrHlpr := dirMgrHelper{}
	dMgr := DirMgr{}

	isEmpty, err := dMgrHlpr.setDirMgr(
		&dMgr,
		pathName,
		ePrefix.String(),
		"dMgr",
		"pathName")

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

	dMgrHlpr := dirMgrHelper{}
	dMgr := DirMgr{}

	isEmpty, err := dMgrHlpr.setDirMgr(
		&dMgr,
		pathName,
		ePrefix.String(),
		"dMgr",
		"pathName")

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

// AddFileMgrCollection - Adds another collection of File Manager (FileMgr)
// objects to the current collection.
func (fMgrs *FileMgrCollection) AddFileMgrCollection(fMgrs2 *FileMgrCollection) {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if fMgrs2.fileMgrs == nil {
		fMgrs2.fileMgrs = make([]FileMgr, 0, 50)
	}

	lOmc2 := len(fMgrs2.fileMgrs)

	if lOmc2 == 0 {
		return
	}

	for i := 0; i < lOmc2; i++ {
		fMgrs.AddFileMgr(fMgrs2.fileMgrs[i].CopyOut())
	}

	return
}

// CopyFilesToDir - Copies all the files in the File Manager Collection to
// the specified target directory.
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileMgrCollection."+
			"CopyFilesToDir()",
		"")

	if err != nil {
		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
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
			ePrefix)

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Copy Failure on index='%v'\n"+
				"file='%v'."+
				"Error='%v'",
				ePrefix.String(),
				i, fMgrs.fileMgrs[i].absolutePathFileName,
				err.Error())
		}

	}

	return err
}

// CopyOut - Returns an FileMgrCollection which is an
// exact duplicate of the current FileMgrCollection
func (fMgrs *FileMgrCollection) CopyOut() (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.CopyOut() "

	fMgrs2 := FileMgrCollection{}

	fMgrs2.fileMgrs = make([]FileMgr, 0, 50)

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	lOmc := len(fMgrs.fileMgrs)

	if lOmc == 0 {
		return FileMgrCollection{},
			errors.New(ePrefix +
				"Error: This File Manager Collection ('FileMgrCollection') is EMPTY! ")
	}

	for i := 0; i < lOmc; i++ {
		fMgrs2.AddFileMgr(fMgrs.fileMgrs[i].CopyOut())
	}

	return fMgrs2, nil
}

// DeleteAtIndex - Deletes a member File Manager from the
// collection at the index specified by input parameter 'idx'.
//
// If successful, at the completion of this method, the File
// Manager Collection array will have a length which is one
// less than the starting array length.
func (fMgrs *FileMgrCollection) DeleteAtIndex(idx int) error {

	ePrefix := "FileMgrCollection.DeleteAtIndex() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if idx < 0 {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return errors.New(ePrefix +
			"Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!\n")
	}

	if idx >= arrayLen {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is greater than the "+
			"length of the collection index. Index Out-Of-Range! "+
			"idx='%v' Array Length='%v' ", idx, arrayLen)
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

	return nil
}

// FindFiles - Searches the current FileMgrCollection and returns a new
// FileMgrCollection containing FileMgr objects which match the specified
// search criteria.
func (fMgrs *FileMgrCollection) FindFiles(
	fileSelectionCriteria FileSelectionCriteria) (FileMgrCollection, error) {

	ePrefix := "FileMgrCollection.FindFiles() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	lDirCol := len(fMgrs.fileMgrs)

	if lDirCol == 0 {
		return FileMgrCollection{}.New(), nil
	}

	fh := FileHelper{}

	var isMatchedFile bool
	var err error

	fMgrs2 := FileMgrCollection{}.New()

	for i := 0; i < lDirCol; i++ {
		fMgr := fMgrs.fileMgrs[i]

		if fMgr.actualFileInfo.isFInfoInitialized {

			isMatchedFile,
				err,
				_ = fh.FilterFileName(
				&fMgr.actualFileInfo,
				fileSelectionCriteria,
				ePrefix)

			if err != nil {
				return FileMgrCollection{},
					fmt.Errorf(ePrefix+
						"Error returned by "+
						"fh.FilterFileName(fMgr.actualFileInfo, fileSelectionCriteria) "+
						"fMgr.actualFileInfo.Name()='%v'  Error='%v'",
						fMgr.actualFileInfo.Name(), err.Error())
			}

		} else {

			fip := FileInfoPlus{}

			fip.SetName(fMgr.fileNameExt)

			isMatchedFile,
				err,
				_ = fh.FilterFileName(
				&fip,
				fileSelectionCriteria,
				ePrefix)

			if err != nil {
				return FileMgrCollection{}, fmt.Errorf(ePrefix+
					"Error returned by fh.FilterFileName(fip, fileSelectionCriteria) "+
					"fip.Name()='%v'  Error='%v'", fip.Name(), err.Error())
			}

		}

		if isMatchedFile {
			fMgrs2.AddFileMgr(fMgr)
		}

	}

	return fMgrs2, nil
}

// GetFileMgrArray - Returns the entire Directory Manager Array managed
// by this collection.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	[]FileMgr      - The array of FileMgr instances maintained by this
//	                 collection.
func (fMgrs *FileMgrCollection) GetFileMgrArray() []FileMgr {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 10)
	}

	return fMgrs.fileMgrs
}

// GetFileMgrAtIndex - If successful, this method returns a pointer to
// the FileMgr instance at the array index specified. The 'Peek' and 'Pop'
// methods below return FileMgr objects using a 'deep' copy and therefore
// offer better protection against data corruption.
func (fMgrs *FileMgrCollection) GetFileMgrAtIndex(idx int) (*FileMgr, error) {

	ePrefix := "FileMgrCollection.GetFileMgrAtIndex() "

	emptyFileMgr := FileMgr{}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return &emptyFileMgr,
			fmt.Errorf(ePrefix +
				"Error: This File Manager Collection ('FileMgrCollection') is EMPTY!\n")
	}

	if idx < 0 || idx >= arrayLen {

		return &emptyFileMgr,
			fmt.Errorf(ePrefix+
				"Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
				"The minimum index is '0'. "+
				"The maximum index is '%v'. ", idx, arrayLen-1)

	}

	return &fMgrs.fileMgrs[idx], nil

}

// GetNumOfFileMgrs - returns the array length of the
// File Manager Collection, 'FileMgrCollection'.
// Effectively, the returned integer is a count of the
// number of File Managers (FileMgr's) in the Collection.
func (fMgrs *FileMgrCollection) GetNumOfFileMgrs() int {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	return len(fMgrs.fileMgrs)
}

// GetNumOfFiles
//
// Returns the array length of the File Manager
// Collection, 'FileMgrCollection'.
//
// Effectively, the returned integer is a count of the
// number of files or File Managers (FileMgr's) in the
// Collection.
func (fMgrs *FileMgrCollection) GetNumOfFiles() int {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	return len(fMgrs.fileMgrs)
}

// GetTotalFileBytes - Returns the total number of file bytes
// represented by all files in the collection.
func (fMgrs *FileMgrCollection) GetTotalFileBytes() uint64 {

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
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

// InsertFileMgrAtIndex - Inserts a new File Manager into the collection at
// array 'index'. The new File Manager is passed as input parameter 'fMgr'.
//
// If input parameter 'index' is less than zero, an error will be returned. If
// 'index' exceeds the value of the last index in the collection, 'fMgr' will be
// added to the end of the collection at the next legal index.
func (fMgrs *FileMgrCollection) InsertFileMgrAtIndex(fMgr FileMgr, index int) error {

	ePrefix := "FileMgrCollection.InsertFileMgrAtIndex() "

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	if index < 0 {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'index' is LESS THAN ZERO! "+
			"index='%v' ", index)
	}

	lenfMgrs := len(fMgrs.fileMgrs)

	if index >= lenfMgrs {
		fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
		return nil
	}

	newFileMgrs := make([]FileMgr, 0, 100)

	if index == 0 {
		newFileMgrs = append(newFileMgrs, fMgr.CopyOut())
		fMgrs.fileMgrs = append(newFileMgrs, fMgrs.fileMgrs...)
		return nil
	}

	newFileMgrs = append(newFileMgrs, fMgrs.fileMgrs[index:]...)

	fMgrs.fileMgrs = append(fMgrs.fileMgrs[:index])
	fMgrs.fileMgrs = append(fMgrs.fileMgrs, fMgr.CopyOut())
	fMgrs.fileMgrs = append(fMgrs.fileMgrs, newFileMgrs...)

	return nil
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
