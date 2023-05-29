package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// ValidPathStrDto - Used to transfer file/path string attributes and
// associated errors.
type ValidPathStrDto struct {
	isInitialized bool // signals whether the current ValidPathStrDto instance
	//  has been properly initialized.

	originalPathStr string // The original, unformatted path string

	pathStr string // The path string which may or may not be
	//  the absolute path

	pathFInfoPlus FileInfoPlus // Only populated if absValidPath exists on disk.

	pathDoesExist PathExistsStatusCode // -1 = don't know, file/path existence has not been tested
	//  0 - No, tests show the file/path doesn't exist on disk.
	//  1 - Yes, tests show the file/path does exist on disk.

	pathStrLength int // Length of the path string

	absPathStr string // The absolute path version of 'path'

	absPathFInfoPlus FileInfoPlus // Only populated if absValidPath
	// exists on disk.

	absPathDoesExist PathExistsStatusCode // -1 = don't know, has not been tested
	//  0 - No, tests shown path doesn't exist
	//  1 - Yes, tests show path does exist

	absPathStrLength int // Length of the absolute path string

	pathType PathFileTypeCode // The path type. Path File, Path Directory

	pathIsValid PathValidityStatusCode // -1 - don't know
	//  0 - No path is NOT valid
	//  1 - Yes, path is valid

	pathVolumeName string // Volume name associated with current path

	pathVolumeIndex int // Index of the starting character of Volume Name
	// in the path string.

	pathVolumeStrLength int // Length of the Volume name in the path string.

	err error // If no error is encountered
	// this value is nil

	lock *sync.Mutex
}

func (vpDto *ValidPathStrDto) AbsolutePathDoesExist() PathExistsStatusCode {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.absPathDoesExist
}

// GetAbsPath - "getter" method for internal field,
// ValidPathStrDto.absPathStr .
func (vpDto *ValidPathStrDto) GetAbsPath() string {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.absPathStr
}

// GetAbsPathStrLen - "getter" method for internal field,
// ValidPathStrDto.absPathStrLength .
func (vpDto *ValidPathStrDto) GetAbsPathStrLen() int {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.absPathStrLength
}

// GetAbsPathFileInfo - "getter" method for internal field,
// ValidPathStrDto.absPathFInfoPlus .
func (vpDto *ValidPathStrDto) GetAbsPathFileInfo() FileInfoPlus {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.absPathFInfoPlus
}

// GetOriginalPathStr - "getter" method for internal field,
// ValidPathStrDto.originalPathStr .
func (vpDto *ValidPathStrDto) GetOriginalPathStr() string {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.originalPathStr
}

// GetPath - "getter" method for internal field,
// ValidPathStrDto.pathStr .
func (vpDto *ValidPathStrDto) GetPath() string {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathStr
}

// GetPathFileInfo - "getter" method for internal field,
// ValidPathStrDto.pathFInfoPlus .
func (vpDto *ValidPathStrDto) GetPathFileInfo() FileInfoPlus {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathFInfoPlus
}

// GetPathIsValid - "getter" method for internal field,
// ValidPathStrDto.pathIsValid .
//
//	                                      Int
//	   Enumeration                       Value         Definition
//	----------------------------------------------------------------------------
//	PathValidityStatusCode(0).Unknown()   -1  Path/file name validity has NOT been
//	                                          tested and its status as 'Valid' or
//	                                          'Invalid' is Unknown.
//
//	PathValidityStatusCode(0).Invalid()    0  Tests have verified that the Path/file
//	                                          name is 'Invalid'.
//
//	PathValidityStatusCode(0).Valid()     +1  Tests have verified that the Path/file
//	                                          name is 'Valid'.
func (vpDto *ValidPathStrDto) GetPathIsValid() PathValidityStatusCode {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathIsValid
}

// GetPathStrLen - "getter" method for internal field,
// ValidPathStrDto.pathStrLength .
func (vpDto *ValidPathStrDto) GetPathStrLen() int {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathStrLength
}

// GetPathType - "getter" method for internal field,
// ValidPathStrDto.pathType .
func (vpDto *ValidPathStrDto) GetPathType() PathFileTypeCode {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathType
}

// GetPathVolumeName - "getter" method for internal field,
// ValidPathStrDto.pathVolumeName .
func (vpDto *ValidPathStrDto) GetPathVolumeName() string {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathVolumeName
}

// GetPathVolumeIndex - "getter" method for internal field,
// ValidPathStrDto.pathVolumeIndex .
func (vpDto *ValidPathStrDto) GetPathVolumeIndex() int {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathVolumeIndex
}

// GetPathVolumeStrLength - "getter" method for internal field,
// ValidPathStrDto.pathVolumeStrLength .
func (vpDto *ValidPathStrDto) GetPathVolumeStrLength() int {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathVolumeStrLength
}

// GetError - "getter" method for internal field,
// ValidPathStrDto.err . Process errors occurring during
// operations performed on the current ValidPathStrDto
// instance are recorded here.
func (vpDto *ValidPathStrDto) GetError() error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.err
}

// IsDtoValid - Performs an analysis of the internal data structures for
// the current ValidPathStrDto instance to determine its validity.
//
// If the current ValidPathStrDto is invalid, this method will return
// an error. If the instance is valid, this method will return 'nil'.
func (vpDto *ValidPathStrDto) IsDtoValid(
	errorPrefix interface{}) error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ValidPathStrDto.IsDtoValid()",
		"")

	if err != nil {
		return err
	}

	if !vpDto.isInitialized {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"This ValidPathStrDto instance is NOT initialized!\n"+
			"vpDto.isInitialized='false'\n",
			ePrefix.String())
	}

	if vpDto.pathIsValid != PathValidStatus.Valid() {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"The ValidPathStrDto 'Path Is Valid flag' is Invalid!\n"+
			"vpDto.pathIsValid=%v'\n",
			ePrefix.String(),
			vpDto.pathIsValid.String())
	}

	if len(vpDto.pathStr) == 0 {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"The ValidPathStrDto 'pathStr' is EMPTY!\n",
			ePrefix.String())
	}

	if len(vpDto.absPathStr) == 0 {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"The ValidPathStrDto absolute path string is EMPTY!\n",
			ePrefix.String())
	}

	if vpDto.pathDoesExist < PathExistsStatus.Unknown() ||
		vpDto.pathDoesExist > PathExistsStatus.Exists() {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"ValidPathStrDto.pathDoesExist holds an invalid value.\n"+
			"ValidPathStrDto.pathDoesExist='%v'\n",
			ePrefix.String(),
			vpDto.pathDoesExist)
	}

	if vpDto.absPathDoesExist < PathExistsStatus.Unknown() ||
		vpDto.absPathDoesExist > PathExistsStatus.Exists() {
		return fmt.Errorf("%v\n"+
			"ERROR: This ValidPathStrDto is INVALID!\n"+
			"ValidPathStrDto.absPathDoesExist holds an invalid value.\n"+
			"ValidPathStrDto.absPathDoesExist='%v'\n",
			ePrefix.String(),
			vpDto.absPathDoesExist)
	}

	return nil
}

// IsInitialized - Returns a boolean value indicating whether the current instance
// of ValidPathStrDto has been initialized.
func (vpDto *ValidPathStrDto) IsInitialized() bool {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.isInitialized
}

// New
//
// This method returns a new and empty instance of
// ValidPathStrDto.
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
//	ValidPathStrDto
//
//		This method returns a new and empty instance of
//		ValidPathStrDto.
//
//		All member data fields in this new instance are
//		initialized to their zero or empty values.
func (vpDto *ValidPathStrDto) New() ValidPathStrDto {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	newValPathDto := ValidPathStrDto{}
	newValPathDto.pathStr = ""
	newValPathDto.pathStrLength = -1
	newValPathDto.absPathStr = ""
	newValPathDto.absPathStrLength = -1
	newValPathDto.pathDoesExist = PathExistsStatus.Unknown()
	newValPathDto.absPathDoesExist = PathExistsStatus.Unknown()
	newValPathDto.pathIsValid = PathValidStatus.Unknown()
	newValPathDto.isInitialized = false
	newValPathDto.pathVolumeName = ""
	newValPathDto.pathVolumeIndex = -1
	newValPathDto.pathVolumeStrLength = 0
	newValPathDto.err = nil

	return newValPathDto
}

// PathDoesExist - Returns a boolean value indicating whether the path
// defined by this ValidPathStrDto instance actually exists on disk.
//
//	                                      Int
//	   Enumeration                       Value         Definition
//	----------------------------------------------------------------------------
//	PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//	                                            tested and status is 'Unknown'.
//
//	PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//	                                            and path file name does NOT exist on
//	                                            disk.
//
//	PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//	                                            and path file name DOES exist on
//	                                            disk.
func (vpDto *ValidPathStrDto) PathDoesExist() PathExistsStatusCode {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	return vpDto.pathDoesExist
}

// SetIsInitialized - "setter" method for internal data field
// vpDto.isInitialized .
func (vpDto *ValidPathStrDto) SetIsInitialized(isInitialized bool) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.isInitialized = isInitialized

	return
}

// SetPath - "setter" method for internal data fields
// vpDto.pathStr and  vpDto.pathStrLength.
func (vpDto *ValidPathStrDto) SetPath(pathStr string) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.pathStr = pathStr

	vpDto.pathStrLength = len(pathStr)

	return
}

// SetPathFileInfo - "setter" method for internal data field
// vpDto.pathFInfoPlus .
func (vpDto *ValidPathStrDto) SetPathFileInfo(fInfPlus FileInfoPlus) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.pathFInfoPlus = fInfPlus.CopyOut()

	return
}

// SetAbsPath - "setter" method for internal data fields
// vpDto.absPathStr and vpDto.absPathStrLength.
func (vpDto *ValidPathStrDto) SetAbsPath(absPathStr string) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.absPathStr = absPathStr
	vpDto.absPathStrLength = len(absPathStr)

	return
}

// SetAbsPathDoesExistStatus - Input parameter 'absPathExistStatus' is used to set
// the internal field vpDto.absPathDoesExist which is of type PathExistsStatusCode.
//
//	                                      Int
//	   Enumeration                       Value         Definition
//	----------------------------------------------------------------------------
//	PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//	                                            tested and status is 'Unknown'.
//
//	PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//	                                            and path file name does NOT exist on
//	                                            disk.
//
//	PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//	                                            and path file name DOES exist on
//	                                            disk.
func (vpDto *ValidPathStrDto) SetAbsPathDoesExistStatus(
	absPathExistStatus PathExistsStatusCode,
	errorPrefix interface{}) error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ValidPathStrDto."+
			"SetAbsPathDoesExistStatus()",
		"")

	if err != nil {
		return err
	}

	err = absPathExistStatus.StatusIsValid()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Return: absPathExistStatus.StatusIsValid()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	vpDto.absPathDoesExist = absPathExistStatus

	return nil
}

// SetAbsPathFileInfo - "setter" method for internal data field
// vpDto.absPathFInfoPlus .
func (vpDto *ValidPathStrDto) SetAbsPathFileInfo(absFInfo FileInfoPlus) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.absPathFInfoPlus = absFInfo.CopyOut()
}

// SetOriginalPathStr - "setter" method for internal data field
// vpDto.originalPathStr .
func (vpDto *ValidPathStrDto) SetOriginalPathStr(originalPathStr string) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.originalPathStr = originalPathStr
}

// SetPathDoesExist
//
// This is a "setter" method for internal data field
// vpDto.pathDoesExist which is of type PathExistsStatusCode.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathDoesExist				PathExistsStatusCode
//
//		This instance of PathExistsStatusCode is an enumeration
//		which must be set to one of the following values:
//
//		                                      Int
//		   Enumeration                       Value         Definition
//		----------------------------------------------------------------------------
//		PathExistsStatusCode(0).Unknown()     -1    Path file existence has NOT been
//		                                            tested and status is 'Unknown'.
//
//		PathExistsStatusCode(0).DoesNotExist() 0    Path file existence HAS been tested
//		                                            and path file name does NOT exist on
//		                                            disk.
//
//		PathExistsStatusCode(0).Exists()      +1    Path file existence HAD been tested
//		                                            and path file name DOES exist on
//		                                            disk.
//
//		If 'pathDoesExist' is set to an invalid value, an
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
//	error
//
//		If input parameter 'pathDoesExist' is set to a
//		valid value, the returned error Type is set equal
//		to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (vpDto *ValidPathStrDto) SetPathDoesExist(
	pathDoesExist PathExistsStatusCode,
	errorPrefix interface{}) error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ValidPathStrDto."+
			"SetPathDoesExist()",
		"")

	if err != nil {
		return err
	}

	err = pathDoesExist.StatusIsValid()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Return: pathDoesExist.StatusIsValid()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	vpDto.pathDoesExist = pathDoesExist

	return nil
}

// SetPathIsValid
//
// This is a "setter" method for the internal data field
// vpDto.pathIsValid which is of type PathValidityStatusCode.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathIsValid					PathValidityStatusCode
//
//		This instance of 'PathValidityStatusCode' is an
//		enumeration which must be set to one of the following
//		valid values:
//
//		                  Path Validity
//		 Method            Status Code
//		  Name               Constant       Description
//		______________________________________________________________________
//
//		PathValidityStatusCode(0).Unknown()     -1  Path/file name validity has NOT been
//		                                            tested and its status as 'Valid' or
//		                                            'invalid' is 'Unknown'.
//
//		PathValidityStatusCode(0).Invalid()      0  Tests have verified that the Path/file
//		                                            name is 'Invalid'.
//
//		PathValidityStatusCode(0).Valid()       +1  Tests have verified that the Path/file
//		                                            name is 'Valid'.
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
//		If input parameter 'pathIsValid' is set to a
//		valid value, the returned error Type is set equal
//		to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (vpDto *ValidPathStrDto) SetPathIsValid(
	pathIsValid PathValidityStatusCode,
	errorPrefix interface{}) error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ValidPathStrDto.IsDtoValid()",
		"")

	if err != nil {
		return err
	}

	err = pathIsValid.StatusIsValid()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Return: pathIsValid.StatusIsValid()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	vpDto.pathIsValid = pathIsValid

	return err
}

// SetPathType
//
// This "setter" method for internal data field
// vpDto.pathDoesExist which is of type PathFileTypeCode.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathType					PathFileTypeCode
//
//		This instance of 'PathFileTypeCode' is an
//		enumeration which must be set to one of the
//		following valid values:
//
//		                               Path File
//		         Method                Type Code
//		          Name                 Constant            Description
//		_______________________________________________________________________________
//		PathFileTypeCode(0).None()        0           Path/file name type has NOT been
//		                                              tested and its status not known.
//
//		PathFileTypeCode.Path()           1           Tests have established that the
//		                                              Path/file name string is a
//		                                              directory path which does NOT
//		                                              contain a file name.
//
//		PathFileTypeCode.PathFile()       2           Tests have established that the
//		                                              Path/file name string includes
//		                                              both a directory path AND a file
//		                                              name.
//
//		PathFileTypeCode.File()           3           Tests have established that the
//		                                              Path/file name string consists
//		                                              solely of a file name and does
//		                                              NOT include a directory path.
//
//		PathFileTypeCode.Volume()         4           Tests have established that the
//		                                              Path/file name string consists
//		                                              solely of a volume name and does
//		                                              NOT include a directory path or
//		                                              file name.
//
//		PathFileTypeCode.Indeterminate()  5           Tests have been conducted on the
//		                                              Path/file name string, but the
//		                                              string cannot be classified and
//		                                              its status cannot be determined
//		                                              with certainty.
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
//		If input parameter 'pathIsValid' is set to a
//		valid value, the returned error Type is set equal
//		to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (vpDto *ValidPathStrDto) SetPathType(
	pathType PathFileTypeCode,
	errorPrefix interface{}) error {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ValidPathStrDto."+
			"SetPathType()",
		"")

	if err != nil {
		return err
	}

	err = pathType.StatusIsValid()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Return: pathType.StatusIsValid()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	vpDto.pathType = pathType

	return err
}

// SetPathVolumeName - "setter" method for internal data fields
// vpDto.pathVolumeName and vpDto.pathVolumeStrLength .
func (vpDto *ValidPathStrDto) SetPathVolumeName(
	volumeNameStr string) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.pathVolumeName = volumeNameStr
	vpDto.pathVolumeStrLength = len(volumeNameStr)
}

// SetPathVolumeIndex - "setter" method for internal data field
// vpDto.pathVolumeIndex .
func (vpDto *ValidPathStrDto) SetPathVolumeIndex(volIndex int) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.pathVolumeIndex = volIndex
}

// SetError - "setter" method for internal data field
// vpDto.err .
func (vpDto *ValidPathStrDto) SetError(err error) {

	if vpDto.lock == nil {
		vpDto.lock = new(sync.Mutex)
	}

	vpDto.lock.Lock()

	defer vpDto.lock.Unlock()

	vpDto.err = fmt.Errorf("%v", err.Error())
}
