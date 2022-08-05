package strmech

import (
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

// SetStringArray - Resets the value of the internal string array
// maintained by the current instance of StringArrayDto to that
// of input parameter 'strArray'.
//
// A copy of 'strArray' is used to populated internal member
// variable 'StringArrayDto.StrArray'.
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
