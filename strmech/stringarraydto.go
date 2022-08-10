package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// StringArrayDto - The String Array Data Transfer Object is
// designed to support the creation, management and transfer
// of a string array.
type StringArrayDto struct {
	StrArray []string

	Description1 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of StringArrayDto.

	Description2 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of StringArrayDto.

	lock *sync.Mutex
}

// AddString - Adds a single string to the internal array of
// strings maintained by this instance of StringArrayDto.
//
// Note that no data validation is performed on input parameter
// 'strToAdd'. If 'strToAdd' is an empty string, an empty string
// will be added to the internal string array.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strToAdd                   string
//	   - A string which will be appended to the end of the string
//	     array maintained by the current instance of
//	     StringArrayDto.
//
//	     No data validation is performed on input parameter,
//	     'strToAdd'. If 'strToAdd' is an empty string, an empty
//	     string will be added to the internal string array.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) AddString(
	strToAdd string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.StrArray =
		append(strArrayDto.StrArray, strToAdd)

	return
}

// AddManyStrings - Adds a variable number of strings to the
// string array maintained by the current instance of
// StringArrayDto.
//
// This method is configured as a variadic function with an input
// parameter that accepts a variable number of arguments
// ('stringsToAdd').
//
// Each string in the series of strings passed through parameter
// 'stringsToAdd' is appended to the end of the string array
// maintained by the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringsToAdd               ...string
//	   - This parameter accepts a variable number of string
//	     arguments. Each string argument passed through this
//	     parameter will be appended to the end of the string
//	     array maintained by the current instance of
//	     StringArrayDto.
//
//	     No data validation is performed on this input
//	     parameter. If a string value passed through
//	    'stringsToAdd' is an empty string, an empty string
//	    will be appended to the end of the internal string
//	    array maintained by StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) AddManyStrings(
	stringsToAdd ...string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	for _, val := range stringsToAdd {

		strArrayDto.StrArray =
			append(strArrayDto.StrArray, val)

	}

	return
}

// AddStringArray - Appends a string array to the end of the
// string array contained in the current instance of
// StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strArray                     []string
//	   - This string array will be appended to the end of
//	     the string array contained in the current instance of
//	     StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) AddStringArray(
	strArray []string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.StrArray = append(
		strArrayDto.StrArray, strArray...)

	return
}

// AddStringArrayDto - Receives an instance of StringArrayDto and
// appends that contents of its string array to the end of the
// string array contained in the current instance of
// StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingStrArrayDto          StringArrayDto
//	   - This string array contained in this instance of
//	     StringArrayDto will be appended to the end of the string
//	     array contained in the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) AddStringArrayDto(
	incomingStrArrayDto StringArrayDto) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.StrArray = append(
		strArrayDto.StrArray, incomingStrArrayDto.StrArray...)

	return
}

// CopyIn - Copies the data fields from an incoming instance of
// StringArrayDto ('incomingStrArray') to the data fields of the
// current StringArrayDto instance ('strArrayDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current StringArrayDto instance
// ('strArrayDto') will be deleted and overwritten.
//
// NO DATA VALIDATION is performed on input parameter
// 'incomingStrArray'
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingStrArray           *StringArrayDto
//	   - A pointer to an instance of StringArrayDto. This method
//	     will NOT change the data values of member variables
//	     contained in this instance.
//
//	     All data values in this StringArrayDto instance
//	     ('incomingStrArray') will be copied to the current
//	     StringArrayDto instance ('strArrayDto').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) CopyIn(
	incomingStrArray *StringArrayDto,
	errorPrefix interface{}) error {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(stringArrayDtoNanobot).copyStringArrayDto(
		strArrayDto,
		incomingStrArray,
		ePrefix.XCpy(
			"strArrayDto<-incomingStrArray"))
}

// CopyOut - Returns a deep copy of the current StringArrayDto
// instance.
//
// No data validation is performed on the current StringArrayDto
// instance prior to returning the deep StringArrayDto copy of
// this instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyStrArrayDto        StringArrayDto
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current StringArrayDto instance.
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyStrArrayDto StringArrayDto,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyStrArrayDto, err
	}

	err = new(stringArrayDtoNanobot).copyStringArrayDto(
		&deepCopyStrArrayDto,
		strArrayDto,
		ePrefix.XCpy(
			"deepCopyStrArrayDto<-strArrayDto"))

	return deepCopyStrArrayDto, err
}

// DeleteAtIndex - Deletes an array member element from the
// target string array contained within the current instance
// of StringArrayDto.
//
// After completion of the deletion operation, the target string
// array will have length one less than the length of the original
// string array.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// This method will delete one member element from the string array
// contained within the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	    zeroBasedIndex            int
//	      - The index number of the array element in the target
//			   string array which will be deleted.
//
//		       If the target string array has zero member elements,
//		       this method will return an error.
//
//		       If 'zeroBasedIndex' has a value less than zero, this
//		       method will return an error.
//
//		       If 'zeroBasedIndex' has a value greater than the last
//		       index in the string array, this method will return an
//		       error.
//
//
//	    errorPrefix               interface{}
//	      - This object encapsulates error prefix text which is
//	        included in all returned error messages. Usually, it
//	        contains the name of the calling method or methods
//	        listed as a method or function chain of execution.
//
//	        If no error prefix information is needed, set this
//	        parameter to 'nil'.
//
//	        This empty interface must be convertible to one of the
//	        following types:
//
//
//	        1. nil - A nil value is valid and generates an empty
//	                 collection of error prefix and error context
//	                 information.
//
//	        2. string - A string containing error prefix information.
//
//	        3. []string A one-dimensional slice of strings containing
//	                    error prefix information
//
//	        4. [][2]string A two-dimensional slice of strings
//	           containing error prefix and error context information.
//
//	        5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                          ErrorPrefixInfo from this object will be
//	                          copied to 'errPrefDto'.
//
//	        6. *ErrPrefixDto - A pointer to an instance of
//	                           ErrPrefixDto. ErrorPrefixInfo from this
//	                           object will be copied to 'errPrefDto'.
//
//	        7. IBasicErrorPrefix - An interface to a method generating
//	                               a two-dimensional slice of strings
//	                               containing error prefix and error
//	                               context information.
//
//	        If parameter 'errorPrefix' is NOT convertible to one of
//	        the valid types listed above, it will be considered
//	        invalid and trigger the return of an error.
//
//	        Types ErrPrefixDto and IBasicErrorPrefix are included in
//	        the 'errpref' software package,
//	        "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil' signaling that the designated
//	     Text Line element in the Text Lines Collection has been
//	     deleted. If errors are encountered during processing, the
//	     returned error Type will encapsulate an error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (strArrayDto *StringArrayDto) DeleteAtIndex(
	zeroBasedIndex int,
	errorPrefix interface{}) error {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"DeleteAtIndex()",
		"")

	if err != nil {
		return err
	}

	if len(strArrayDto.StrArray) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of StringArrayDto\n"+
			"has a zero length array. Deletion is invalid\n"+
			"len(strArrayDto.StrArray) == 0\n",
			ePrefix.String())

		return err
	}

	return new(stringArrayDtoElectron).
		deleteStringArrayElement(
			strArrayDto,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf("delete strArrayDto[%v]",
					zeroBasedIndex)))

}

// Empty - Resets all internal member variables for the current
// instance of StringArrayDto to their zero or uninitialized
// states.
//
// This means that the internal string array maintained by this
// instance of StringArrayDto is to a value of 'nil'.
//
//	StringArrayDto.StrArray = nil
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of StringArrayDto. All member variable data
// values will be reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) Empty() {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	new(stringArrayDtoAtom).empty(
		strArrayDto)

	strArrayDto.lock.Unlock()

	strArrayDto.lock = nil

}

// Equal - Receives a pointer to another instance of StringArrayDto
// and proceeds to compare the member variables to those of the
// current StringArrayDto instance in order to determine if they
// are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingStrArray           *StringArrayDto
//	   - A pointer to an incoming instance of StringArrayDto. This
//	     method will compare all member variable data values in
//	     this instance against those contained in the current
//	     instance of StringArrayDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingStrArray' are equal in all respects to
//	     those contained in the current instance of StringArrayDto,
//	     this method will return a boolean value of 'true'.
//	     Otherwise, a value of 'false' will be returned to the
//	     calling function.
func (strArrayDto *StringArrayDto) Equal(
	incomingStrArray *StringArrayDto) bool {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	return new(stringArrayDtoAtom).equal(
		strArrayDto,
		incomingStrArray)
}

// GetDescription1 - Description1 is a name, tag, label or
// narrative text used to describe the current instance of
// StringArrayDto.
//
// 'Description1' is the first of two descriptive text strings.
//
// This method returns the current value of the 'Description1'
// string for the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//		string
//		- This method returns the string value of internal member
//	      variable, 'Description1'. This string is a  name, label
//	      or narrative text used to describe the current instance
//	      of StringArrayDto.
func (strArrayDto *StringArrayDto) GetDescription1() string {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	return strArrayDto.Description1
}

// GetDescription2 - Description2 is a name, label or narrative
// text used to describe the current instance of StringArrayDto.
//
// 'Description2' is the second of two descriptive text strings.
//
// This method returns the current value of the 'Description2'
// string for the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//		string
//		- This method returns the string value of internal member
//	      variable, 'Description2'. This string is a  name, label
//	      or narrative text used to describe the current instance
//	      of StringArrayDto.
func (strArrayDto *StringArrayDto) GetDescription2() string {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	return strArrayDto.Description2
}

// GetStringArray - Returns a deep copy of the string array
// maintained by the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//		[]string
//	    - This method returns a deep copy of the string array
//	      maintained by the current instance of StringArrayDto.
//	      If this internal string array is empty, a value of
//	      'nil' will be returned.
func (strArrayDto *StringArrayDto) GetStringArray() []string {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	lenStrArray := len(strArrayDto.StrArray)

	if lenStrArray == 0 {
		return nil
	}

	strArray := make([]string, lenStrArray)

	for i := 0; i < lenStrArray; i++ {
		strArray[i] =
			strArrayDto.StrArray[i]
	}

	return strArray
}

// GetRuneArrayCollection - Returns an instance of
// RuneArrayCollection generated from the string array contained
// within the current instance of StringArrayDto.
//
// A Rune Array Collection is an array of rune arrays. Each
// member element in the returned Rune Array Collection is created
// from a corresponding string in the string array encapsulated by
// the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) GetRuneArrayCollection() RuneArrayCollection {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	runeArrayCol := RuneArrayCollection{}

	lenStrArray := len(strArrayDto.StrArray)

	if lenStrArray == 0 {
		return runeArrayCol
	}

	runeArrayCol.runeArrayDtoCol =
		make([]RuneArrayDto, lenStrArray)

	for i := 0; i < lenStrArray; i++ {

		runeArrayCol.runeArrayDtoCol[i].CharsArray =
			[]rune(strArrayDto.StrArray[i])

		runeArrayCol.runeArrayDtoCol[i].charSearchType =
			CharSearchType.LinearTargetStartingIndex()

	}

	return runeArrayCol
}

// GetStringArrayLength - Returns the length of the internal string
// maintained by the current instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	int
//	   - This method returns the length of the string array
//	     maintained by the current instance of StringArrayDto.
func (strArrayDto *StringArrayDto) GetStringArrayLength() int {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	return len(strArrayDto.StrArray)
}

// InsertAtIndex - Inserts a new string ('insertStr') into a target
// string array contained within the current instance of
// StringArrayDto.
//
// 'insertStr' is inserted into the target string array at an array
// index specified by input parameter, 'zeroBasedIndex'.
//
// If the target string array has zero member elements, the new
// inserted string will be added as the first and only member of
// the string array.
//
// If 'zeroBasedIndex' has a value less than zero, the new inserted
// string will become the first element in the string array and all
// the old array elements will be appended to that new first element.
//
// If 'zeroBasedIndex' has a value greater than last index in the
// string array, the new inserted string will be appended to the end
// of the current string array.
//
// Otherwise, the new inserted string be inserted at the array element
// specified by 'zeroBasedIndex'. The old member string element which
// formerly occupied index 'zeroBasedIndex' will occur immediately
// after the inserted string in the new string array.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		   insertStr                 string
//	      - This is the string which will be inserted into the
//	        target string array at the index specified by input
//	        parameter, 'zeroBasedIndex'.
//
//
//	    zeroBasedIndex            int
//	      - The index number of the array element in the target
//			   string array where 'insertStr' will be inserted.
//
//		       If the target string array has zero member elements, the
//		       new inserted string will be added as the first and only
//		       member of the string array.
//
//		       If 'zeroBasedIndex' has a value less than zero, the new
//		       inserted string will become the first element in the
//		       string array and all the old array elements will be
//		       appended to that new first element.
//
//		       If 'zeroBasedIndex' has a value greater than the last
//		       index in the string array, the new inserted string will
//		       be appended to the end of the target string array.
//
//		       Otherwise, the new inserted string be inserted at the
//		       array element index specified by 'zeroBasedIndex'. The
//		       old member string element which formerly occupied index
//		       'zeroBasedIndex' will be positioned immediately after
//		       the inserted string in the new string array.
//
//
//	    errorPrefix              interface{}
//	      - This object encapsulates error prefix text which is
//	        included in all returned error messages. Usually, it
//	        contains the name of the calling method or methods
//	        listed as a method or function chain of execution.
//
//	        If no error prefix information is needed, set this
//	        parameter to 'nil'.
//
//	        This empty interface must be convertible to one of the
//	        following types:
//
//
//	        1. nil - A nil value is valid and generates an empty
//	                 collection of error prefix and error context
//	                 information.
//
//	        2. string - A string containing error prefix information.
//
//	        3. []string A one-dimensional slice of strings containing
//	                    error prefix information
//
//	        4. [][2]string A two-dimensional slice of strings
//	           containing error prefix and error context information.
//
//	        5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                          ErrorPrefixInfo from this object will be
//	                          copied to 'errPrefDto'.
//
//	        6. *ErrPrefixDto - A pointer to an instance of
//	                           ErrPrefixDto. ErrorPrefixInfo from this
//	                           object will be copied to 'errPrefDto'.
//
//	        7. IBasicErrorPrefix - An interface to a method generating
//	                               a two-dimensional slice of strings
//	                               containing error prefix and error
//	                               context information.
//
//	        If parameter 'errorPrefix' is NOT convertible to one of
//	        the valid types listed above, it will be considered
//	        invalid and trigger the return of an error.
//
//	        Types ErrPrefixDto and IBasicErrorPrefix are included in
//	        the 'errpref' software package,
//	        "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil' signaling that the designated
//	     Text Line element in the Text Lines Collection has been
//	     deleted. If errors are encountered during processing, the
//	     returned error Type will encapsulate an error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (strArrayDto *StringArrayDto) InsertAtIndex(
	insertStr string,
	zeroBasedIndex int,
	errorPrefix interface{}) error {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(stringArrayDtoElectron).insertAtIndex(
		strArrayDto,
		insertStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto[%v]=insertStr",
				zeroBasedIndex)))
}

// New - Returns a new instance of StringArrayDto containing
// an empty string array.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//		StringArrayDto
//		   - This method returns a new instance of StringArrayDto.
//	      The internal string array maintained by this new
//	      instance is empty and set to 'nil'.
func (strArrayDto StringArrayDto) New() StringArrayDto {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	newStrArray := StringArrayDto{}

	new(stringArrayDtoAtom).empty(
		&newStrArray)

	return newStrArray
}

// NewStringArray - Returns a new instance of StringArrayDto. The
// internal string array for this instance will be populated from
// another string array passed as an input parameter, 'strArray'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		strArray                   []string
//	    - This string array will be used to populate the string
//	      array for the new returned instance of StringArrayDto.
//
//
//	 desc1                      string
//	    - A name, label, tag or narrative text used to describe
//	      the new returned instance of StringArrayDto.
//
//	      This particular text description will be assigned to the
//	      internal member variable, 'StringArrayDto.Description1'.
//
//
//	 desc2                      string
//	    - A name, label, tag or narrative text used to describe
//	      the new returned instance of StringArrayDto.
//
//	      This particular text description will be assigned to the
//	      internal member variable, 'StringArrayDto.Description2'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//		StringArrayDto
//		- This method returns a new instance of StringArrayDto.
//	      The internal string array maintained by this new
//	      instance is empty and set to 'nil'.
func (strArrayDto StringArrayDto) NewStringArray(
	strArray []string,
	desc1 string,
	desc2 string) StringArrayDto {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	newStrArray := StringArrayDto{}

	lenStrArray := len(strArray)

	if lenStrArray > 0 {

		newStrArray.StrArray =
			make([]string, lenStrArray)

		for i := 0; i < lenStrArray; i++ {
			newStrArray.StrArray[i] =
				strArray[i]
		}

	}

	newStrArray.Description1 = desc1

	newStrArray.Description2 = desc2

	return newStrArray
}

// PeekAtFirstStr - Returns a copy of the first array element in
// the string array maintained by the current instance of
// StringArrayDto.
//
// Since this is a 'Peek' operation, the first array element in
// the string array will NOT be deleted.
//
// After completion of this method the string array contained in
// the current instance of StringArrayDto will remain unchanged.
//
// If the string array maintained by the StringArrayDto instance is
// empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	firstArrayStr              string
//	   - If this method completes successfully, a copy of
//	     the first member of the string array maintained by
//	     the StringArrayDto instance will be returned.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PeekAtFirstStr(
	errorPrefix interface{}) (
	firstArrayStr string,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PeekAtFirstStr()",
		"")

	if err != nil {
		return firstArrayStr, err
	}

	firstArrayStr,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		0,
		false,
		ePrefix.XCpy(
			"strArrayDto[0]"))

	return firstArrayStr, err
}

// PeekAtIndex - Returns a copy of the array element in
// the string array, maintained by the current instance of
// StringArrayDto, designated by the index value of parameter
// 'zeroBasedIndex'.
//
// Since this is a 'Peek' operation, the designated array element
// in the string array will NOT be deleted.
//
// After completion of this method the string array contained in
// the current instance of StringArrayDto will remain unchanged.
//
// If the string array maintained by the StringArrayDto instance is
// empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		zeroBasedIndex             int
//		   - The index number of the array element in the
//		     StringArrayDto string array which will be returned
//		     to the calling function.
//
//
//	 errorPrefix                interface{}
//	    - This object encapsulates error prefix text which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods
//	      listed as a method or function chain of execution.
//
//	      If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//	      This empty interface must be convertible to one of the
//	      following types:
//
//
//	      1. nil - A nil value is valid and generates an empty
//	               collection of error prefix and error context
//	               information.
//
//	      2. string - A string containing error prefix information.
//
//	      3. []string A one-dimensional slice of strings containing
//	                  error prefix information
//
//	      4. [][2]string A two-dimensional slice of strings
//	         containing error prefix and error context information.
//
//	      5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                        ErrorPrefixInfo from this object will be
//	                        copied to 'errPrefDto'.
//
//	      6. *ErrPrefixDto - A pointer to an instance of
//	                         ErrPrefixDto. ErrorPrefixInfo from this
//	                         object will be copied to 'errPrefDto'.
//
//	      7. IBasicErrorPrefix - An interface to a method generating
//	                             a two-dimensional slice of strings
//	                             containing error prefix and error
//	                             context information.
//
//	      If parameter 'errorPrefix' is NOT convertible to one of
//	      the valid types listed above, it will be considered
//	      invalid and trigger the return of an error.
//
//	      Types ErrPrefixDto and IBasicErrorPrefix are included in
//	      the 'errpref' software package,
//	      "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	arrayStrAtIndex            string
//	   - If this method completes successfully, a copy of the
//	     string array element, designated by input parameter
//	     'zeroBasedIndex', will be returned.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PeekAtIndex(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	arrayStrAtIndex string,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PeekAtIndex()",
		"")

	if err != nil {
		return arrayStrAtIndex, err
	}

	arrayStrAtIndex,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		zeroBasedIndex,
		false,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]",
				zeroBasedIndex)))

	return arrayStrAtIndex, err
}

// PeekAtLastStr - Returns a copy of the last array element in
// the string array maintained by the current instance of
// StringArrayDto.
//
// Since this is a 'Peek' operation, the last array element in
// the string array will NOT be deleted.
//
// After completion of this method the string array contained in
// the current instance of StringArrayDto will remain unchanged.
//
// If the string array maintained by the current StringArrayDto
// instance is empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	lastArrayStr              string
//	   - If this method completes successfully, a copy of
//	     the last member of the string array maintained by
//	     the StringArrayDto instance will be returned.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PeekAtLastStr(
	errorPrefix interface{}) (
	lastArrayStr string,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PeekAtLastStr()",
		"")

	if err != nil {
		return lastArrayStr, err
	}

	lastIdx := len(strArrayDto.StrArray) - 1

	lastArrayStr,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		lastIdx,
		false,
		ePrefix.XCpy(
			"strArrayDto[0]"))

	return lastArrayStr, err
}

// PopFirstStr - Returns a copy of the first array element in
// the string array maintained by the current instance of
// StringArrayDto.
//
// Since this is a 'Pop' operation, the first array element in
// the string array will be deleted.
//
// After completion of this method the length of the string array
// contained in the current instance of StringArrayDto will be
// reduced by one (1).
//
// If the string array maintained by the StringArrayDto instance is
// empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	firstArrayStr              string
//	   - If this method completes successfully, a copy of
//	     the first member of the string array maintained by
//	     the StringArrayDto instance will be returned.
//
//
//	newArrayLength             int
//	   - If this method completes successfully, the first string
//	     array element in the current the StringArrayDto instance
//	     will be deleted. This parameter returns the new array
//	     length after this deletion.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PopFirstStr(
	errorPrefix interface{}) (
	firstArrayStr string,
	newArrayLength int,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PopFirstStr()",
		"")

	if err != nil {
		return firstArrayStr, newArrayLength, err
	}

	firstArrayStr,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		0,
		true,
		ePrefix.XCpy(
			"strArrayDto[0]"))

	newArrayLength = len(strArrayDto.StrArray)

	return firstArrayStr, newArrayLength, err
}

// PopAtIndex - Returns a copy of the array element in
// the string array, maintained by the current instance of
// StringArrayDto, designated by the index value of parameter
// 'zeroBasedIndex'.
//
// Since this is a 'Pop' operation, the array element designated
// by input parameter 'zeroBasedIndex' will be deleted from the
// string array.
//
// After completion of this method the length of the string array
// contained in the current instance of StringArrayDto will be
// reduced by one (1).
//
// If the string array maintained by the StringArrayDto instance is
// empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//			zeroBasedIndex             int
//			   - The index number of the array element in the
//			     StringArrayDto string array which will be returned
//			     to the calling function.
//
//	          If the value of 'zeroBasedIndex' is less than zero or
//	          greater than the last index in the string array, an
//	          error will be returned.
//
//
//		 errorPrefix                interface{}
//		    - This object encapsulates error prefix text which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods
//		      listed as a method or function chain of execution.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      This empty interface must be convertible to one of the
//		      following types:
//
//
//		      1. nil - A nil value is valid and generates an empty
//		               collection of error prefix and error context
//		               information.
//
//		      2. string - A string containing error prefix information.
//
//		      3. []string A one-dimensional slice of strings containing
//		                  error prefix information
//
//		      4. [][2]string A two-dimensional slice of strings
//		         containing error prefix and error context information.
//
//		      5. ErrPrefixDto - An instance of ErrPrefixDto. The
//		                        ErrorPrefixInfo from this object will be
//		                        copied to 'errPrefDto'.
//
//		      6. *ErrPrefixDto - A pointer to an instance of
//		                         ErrPrefixDto. ErrorPrefixInfo from this
//		                         object will be copied to 'errPrefDto'.
//
//		      7. IBasicErrorPrefix - An interface to a method generating
//		                             a two-dimensional slice of strings
//		                             containing error prefix and error
//		                             context information.
//
//		      If parameter 'errorPrefix' is NOT convertible to one of
//		      the valid types listed above, it will be considered
//		      invalid and trigger the return of an error.
//
//		      Types ErrPrefixDto and IBasicErrorPrefix are included in
//		      the 'errpref' software package,
//		      "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	arrayStrAtIndex            string
//	   - If this method completes successfully, a copy of the
//	     string array element, designated by input parameter
//	     'zeroBasedIndex', will be returned.
//
//
//	newArrayLength             int
//	   - If this method completes successfully, the first string
//	     array element in the current the StringArrayDto instance
//	     will be deleted. This parameter returns the new array
//	     length after this deletion.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PopAtIndex(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	arrayStrAtIndex string,
	newArrayLength int,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PopAtIndex()",
		"")

	if err != nil {
		return arrayStrAtIndex, newArrayLength, err
	}

	arrayStrAtIndex,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		zeroBasedIndex,
		true,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]",
				zeroBasedIndex)))

	newArrayLength = len(strArrayDto.StrArray)

	return arrayStrAtIndex, newArrayLength, err
}

// PopLastStr - Returns a copy of the last array element in
// the string array maintained by the current instance of
// StringArrayDto.
//
// Since this is a 'Pop' operation, the last array element in
// the string array will be deleted.
//
// After completion of this method the length of the string array
// contained in the current instance of StringArrayDto will be
// reduced by one (1) element.
//
// If the string array maintained by the StringArrayDto instance is
// empty, an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	lastArrayStr              string
//	   - If this method completes successfully, a copy of
//	     the last member of the string array maintained by
//	     the StringArrayDto instance will be returned.
//
//
//	newArrayLength             int
//	   - If this method completes successfully, the last string
//	     array element in the current the StringArrayDto instance
//	     will be deleted. This parameter returns the new array
//	     length after this deletion.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (strArrayDto *StringArrayDto) PopLastStr(
	errorPrefix interface{}) (
	lastArrayStr string,
	newArrayLength int,
	err error) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"PopLastStr()",
		"")

	if err != nil {
		return lastArrayStr, newArrayLength, err
	}

	lastArrayStr,
		err = new(stringArrayDtoAtom).peekPopStringArray(
		strArrayDto,
		0,
		true,
		ePrefix.XCpy(
			"strArrayDto[0]"))

	newArrayLength = len(strArrayDto.StrArray)

	return lastArrayStr, newArrayLength, err
}

// ReplaceAtIndex - Replaces a string in a target string array
// contained within the current instance of StringArrayDto.
//
// The string being replaced is specified by a zero based index
// passed as input parameter, 'zeroBasedIndex'.
//
// If the target string array has zero member elements, the new
// string will be added as the first and only member of the
// string array.
//
// If 'zeroBasedIndex' has a value less than zero, the new string
// will become the first element in the string array and all the
// old array elements will be appended to that new first element.
//
// If 'zeroBasedIndex' has a value greater than last index in the
// string array, the new string will be appended to the end of the
// current string array.
//
// Otherwise, the string array element specified by
// 'zeroBasedIndex' will be replaced by input parameter, 'newStr'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	   newStr                     string
//	     - This is the replacement string which will replace the
//	       StringArrayDto string array element at the array index
//	       specified by input parameter, 'zeroBasedIndex'.
//
//
//		  zeroBasedIndex             int
//	     - The index number of the array element in the
//	       StringArrayDto string array which will be replaced by
//	       input parameter, 'zeroBasedIndex'.
//
//	       If the target string array has zero member elements, the
//	       new string will be added as the first and only member of
//	       the string array.
//
//	       If 'zeroBasedIndex' has a value less than zero, the new
//	       string will become the first element in the string array
//	       and all the old array elements will be appended to that
//	       new first element.
//
//	       If 'zeroBasedIndex' has a value greater than last index
//	       in the string array, the new string will be appended to
//	       the end of the current string array.
//
//	       Otherwise, the string array element specified by
//	       'zeroBasedIndex' will be replaced by input parameter,
//	       'newStr'.
//
//
//		  errorPrefix                interface{}
//		    - This object encapsulates error prefix text which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods
//		      listed as a method or function chain of execution.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      This empty interface must be convertible to one of the
//		      following types:
//
//
//		      1. nil - A nil value is valid and generates an empty
//		               collection of error prefix and error context
//		               information.
//
//		      2. string - A string containing error prefix information.
//
//		      3. []string A one-dimensional slice of strings containing
//		                  error prefix information
//
//		      4. [][2]string A two-dimensional slice of strings
//		         containing error prefix and error context information.
//
//		      5. ErrPrefixDto - An instance of ErrPrefixDto. The
//		                        ErrorPrefixInfo from this object will be
//		                        copied to 'errPrefDto'.
//
//		      6. *ErrPrefixDto - A pointer to an instance of
//		                         ErrPrefixDto. ErrorPrefixInfo from this
//		                         object will be copied to 'errPrefDto'.
//
//		      7. IBasicErrorPrefix - An interface to a method generating
//		                             a two-dimensional slice of strings
//		                             containing error prefix and error
//		                             context information.
//
//		      If parameter 'errorPrefix' is NOT convertible to one of
//		      the valid types listed above, it will be considered
//		      invalid and trigger the return of an error.
//
//		      Types ErrPrefixDto and IBasicErrorPrefix are included in
//		      the 'errpref' software package,
//		      "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                          error
//	  - If this method completes successfully, this returned error
//	    Type is set equal to 'nil' signaling that the designated
//	    Text Line element in the Text Lines Collection has been
//	    deleted. If errors are encountered during processing, the
//	    returned error Type will encapsulate an error message.
//
//	    If an error message is returned, the text value for input
//	    parameter 'errPrefDto' (error prefix) will be prefixed or
//	    attached at the beginning of the error message.
func (strArrayDto *StringArrayDto) ReplaceAtIndex(
	newStr string,
	zeroBasedIndex int,
	errorPrefix interface{}) error {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"StringArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(stringArrayDtoElectron).
		replaceStrArrayAtIndex(
			strArrayDto,
			newStr,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"strArrayDto[%v]=newStr",
					zeroBasedIndex)))
}

// SetDescription1 - Receives a string and assigns that string
// value to internal member variable 'StringArrayDto.Description1'.
//
// 'Description1' is a name, tag, label or narrative text used to
// describe the current instance of StringArrayDto.
//
// 'Description1' is the first of two descriptive text strings.
//
// This method sets the value of 'Description1' from the string
// passed as input parameter, 'desc1'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		desc1                      string
//	    - This string will be used to set the value of internal
//	      member variable 'StringArrayDto.Description1'.
//
//	      'Description1' is first of two string used to provide a
//	       name, tag, label or narrative text for the current
//	       instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) SetDescription1(
	desc1 string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.Description1 = desc1
}

// SetDescription2 - Receives a string and assigns that string
// value to internal member variable 'StringArrayDto.Description2'.
//
// 'Description2' is a name, tag, label or narrative text used to
// describe the current instance of StringArrayDto.
//
// 'Description2' is the first of two descriptive text strings.
//
// This method sets the value of 'Description2' from the string
// passed as input parameter, 'desc2'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		desc2                      string
//	    - This string will be used to set the value of internal
//	      member variable 'StringArrayDto.Description2'.
//
//	      'Description2' is first of two string used to provide a
//	       name, tag, label or narrative text for the current
//	       instance of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) SetDescription2(
	desc2 string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.Description2 = desc2
}

// SetManyStrings - Resets the value of the internal string array
// maintained by the current instance of StringArrayDto based on
// one or more strings passed through the variadic input
// parameter, 'newStrs'.
//
// This method is configured as a variadic function with an input
// parameter that accepts a variable number of arguments
// ('newStrs').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The internal string array contained within current
// StringArrayDto instance ('strArrayDto') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		newStrs                    ...string
//		   - This parameter accepts a variable number of string
//		     arguments. Each string argument passed through this
//		     parameter will be used to create a new string array
//		     encapsulated by the current instance of
//		     StringArrayDto.
//
//		     No data validation is performed on input parameter.
//		     If a string value passed through 'newStrs' is an
//		     empty string, an empty string will be added to the
//	         new string array created within the current instance
//	         of StringArrayDto.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) SetManyStrings(
	newStrs ...string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	strArrayDto.StrArray = nil

	for _, val := range newStrs {

		strArrayDto.StrArray =
			append(strArrayDto.StrArray, val)

	}

	return
}

// SetStringArray - Resets the value of the internal string array
// maintained by the current instance of StringArrayDto to that
// of input parameter 'strArray'.
//
// A copy of 'strArray' is used to populated internal member
// variable 'StringArrayDto.StrArray'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The internal string array contained within current
// StringArrayDto instance ('strArrayDto') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//			strArray                   []string
//		    - This string array will be used to populate the string
//		      array for the new returned instance of StringArrayDto.
//
//	       If parameter 'strArray' is submitted as an empty array,
//	       internal member variable 'StringArrayDto.StrArray' will
//	       be set to 'nil'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (strArrayDto *StringArrayDto) SetStringArray(
	strArray []string) {

	if strArrayDto.lock == nil {
		strArrayDto.lock = new(sync.Mutex)
	}

	strArrayDto.lock.Lock()

	defer strArrayDto.lock.Unlock()

	lenStrArray := len(strArray)

	if lenStrArray == 0 {
		strArrayDto.StrArray = nil

		return
	}

	strArrayDto.StrArray =
		make([]string, lenStrArray)

	for i := 0; i < lenStrArray; i++ {
		strArrayDto.StrArray[i] =
			strArray[i]
	}

	return
}
