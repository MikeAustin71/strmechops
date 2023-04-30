package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
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
// Adds a Directory Manager (DirMgr) using know parent
// path and directory name. This method performs fewer
// string validations then similar methods.
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
func (dMgrs *DirMgrCollection) AddFileInfo(
	parentDirectoryPath string,
	info os.FileInfo,
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
			info.Name(),
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
			info.Name(),
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
			info.Name())
	}

	dMgrs.dirMgrs = append(dMgrs.dirMgrs, newDirMgr)

	return err
}

// AddDirMgrCollection
//
// Adds another collection of File Manager (DirMgr)
// objects to the current collection.
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgrCollection."+
			"AddDirMgrCollection()",
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
			ePrefix.XCpy(fmt.Sprintf(
				"dMgrs2.dirMgrs[%v]",
				i)))

		if err != nil {

			return err
		}

		dMgrs.dirMgrs = append(dMgrs.dirMgrs, dMgrCopy)
	}

	return err
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

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 5)
	}

	lenDMgrs := len(dMgrs.dirMgrs)

	if lenDMgrs == 0 {
		return DirMgrCollection{},
			fmt.Errorf("%v\n"+
				"Error: Empty DirMgrCollection.\n",
				ePrefix.String())
	}

	dMgrs2 := DirMgrCollection{}.New()

	var dMgrCopy DirMgr

	for i := 0; i < lenDMgrs; i++ {

		dMgrCopy,
			err = dMgrs.dirMgrs[i].CopyOut(ePrefix.XCpy(
			fmt.Sprintf("dMgrs.dirMgrs[%v]",
				i)))

		if err != nil {
			return dMgrs2, err
		}

		dMgrs2.dirMgrs = append(dMgrs2.dirMgrs, dMgrCopy)
	}

	return dMgrs2, nil
}

// DeleteAtIndex - Deletes a member Directory Manager from
// the collection at the index specified by input parameter 'idx'.
//
// If successful, at the completion of this method, the Directory
// Manager Collection array will have a length which is one less
// than the starting array length.
func (dMgrs *DirMgrCollection) DeleteAtIndex(idx int) error {

	ePrefix := "DirMgrCollection.DeleteAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if idx < 0 {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return errors.New(ePrefix +
			"Error: The Directory Manager Collection, 'DirMgrCollection', is EMPTY!\n")
	}

	if idx >= arrayLen {
		return fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is greater than the "+
			"length of the collection index. Index Out-Of-Range! "+
			"idx='%v' Array Length='%v' ", idx, arrayLen)
	}

	if arrayLen == 1 {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	} else if idx == 0 {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[1:]
	} else if idx == arrayLen-1 {
		// arrayLen > 1
		dMgrs.dirMgrs = dMgrs.dirMgrs[0 : arrayLen-1]
	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		dMgrs.dirMgrs =
			append(dMgrs.dirMgrs[0:idx], dMgrs.dirMgrs[idx+1:]...)
	}

	return nil
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

// GetDirMgrArray - Returns the entire Directory Manager Array managed
// by this collection.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//	None
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]DirMgr      - The array of DirMgr instances maintained by this
//	                collection.
func (dMgrs *DirMgrCollection) GetDirMgrArray() []DirMgr {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 10)
	}

	return dMgrs.dirMgrs
}

// GetDirMgrAtIndex - If successful, this method returns a pointer to
// the DirMgr instance at the array index specified. The 'Peek' and 'Pop'
// methods below return DirMgr objects using a 'deep' copy and therefore
// offer better protection against data corruption.
func (dMgrs *DirMgrCollection) GetDirMgrAtIndex(idx int) (*DirMgr, error) {

	ePrefix := "DirMgrCollection.GetDirMgrAtIndex() "

	emptyDirMgr := DirMgr{}

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return &emptyDirMgr,
			fmt.Errorf(ePrefix +
				"Error: This Directory Manager Collection ('DirMgrCollection') is EMPTY!\n")
	}

	if idx < 0 || idx >= arrayLen {

		return &emptyDirMgr,
			fmt.Errorf(ePrefix+
				"Error: The input parameter, 'idx', is OUT OF RANGE! idx='%v'.  \n"+
				"The minimum index is '0'. "+
				"The maximum index is '%v'. ", idx, arrayLen-1)

	}

	return &dMgrs.dirMgrs[idx], nil
}

// GetNumOfDirs - returns the number of directories
// contained in this Directory Manager Collection.
func (dMgrs *DirMgrCollection) GetNumOfDirs() int {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

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

// PeekDirMgrAtIndex - Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the Directory
// Manager Collection ('DirMgrCollection'). This is a 'Peek' method
// and therefore the original Directory Manager ('DirMgr') object
// is NOT deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
func (dMgrs *DirMgrCollection) PeekDirMgrAtIndex(idx int) (DirMgr, error) {

	ePrefix := "DirMgrCollection.PeekDirMgrAtIndex() "

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{}, io.EOF

	}

	if idx < 0 {
		return DirMgr{}, fmt.Errorf(ePrefix+
			"Error: Input Parameter 'idx' is less than zero. "+
			"Index Out-Of-Range! idx='%v'", idx)
	}

	if idx >= arrayLen {
		return DirMgr{},
			fmt.Errorf(ePrefix+
				"Error: Input Parameter 'idx' is greater than the "+
				"length of the collection array. Index Out-Of-Range! "+
				"idx='%v' Array Length='%v' ",
				idx, arrayLen)

	}

	return dMgrs.dirMgrs[idx].CopyOut(), nil
}

// PeekFirstDirMgr - Returns a deep copy of the first Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection'). This is a 'Peek' method and therefore
// the original Directory Manager ('DirMgr') object is NOT
// deleted from the Directory Manager Collection
// ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
func (dMgrs *DirMgrCollection) PeekFirstDirMgr() (DirMgr, error) {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	if len(dMgrs.dirMgrs) == 0 {
		return DirMgr{}, io.EOF

	}

	return dMgrs.dirMgrs[0].CopyOut(), nil
}

// PeekLastDirMgr - Returns a deep copy of the last Directory
// Manager ('DirMgr') object in the Directory Manager Collection
// ('DirMgrCollection').
//
// This is a 'Peek' method and therefore the original Directory
// Manager ('DirMgr') object is NOT deleted from the Directory
// Manager Collection ('DirMgrCollection') array.
//
// At the completion of this method, the length of the Directory
// Manager Collection ('DirMgrCollection') array will remain
// unchanged.
func (dMgrs *DirMgrCollection) PeekLastDirMgr() (DirMgr, error) {

	if dMgrs.dirMgrs == nil {
		dMgrs.dirMgrs = make([]DirMgr, 0, 100)
	}

	arrayLen := len(dMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{}, io.EOF
	}

	return dMgrs.dirMgrs[arrayLen-1].CopyOut(), nil
}
