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
