package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type stringArrayDtoNanobot struct {
	lock *sync.Mutex
}

// concatenateStrArray
//
// Returns a single string made up of concatenated string
// array elements taken from the string array
// encapsulated from the instance of StringArrayDto
// passed as input parameter 'strArray'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	insertStr					string
//
//		If 'insertStr' has a string length greater than
//		zero, it will be appended to the end of each
//		string extracted for concatenation from the
//		string array encapsulated by the current
//		StringArrayDto instance.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The strings contained in the internal string
//		array encapsulated by StringArrayDto input
//		parameter 'strArray' will be concatenated
//		together and returned as a single string
//		through this parameter.
func (strArrayDtoNanobot *stringArrayDtoNanobot) concatenateStrArray(
	strArray *StringArrayDto,
	insertStr string) string {

	if strArrayDtoNanobot.lock == nil {
		strArrayDtoNanobot.lock = new(sync.Mutex)
	}

	strArrayDtoNanobot.lock.Lock()

	defer strArrayDtoNanobot.lock.Unlock()

	if strArray == nil {
		return ""
	}

	lenStrArray := len(strArray.StrArray)

	if lenStrArray == 0 {
		return ""
	}

	lenInsertStr := len(insertStr)

	conCatStr := ""

	for i := 0; i < lenStrArray; i++ {

		if len(strArray.StrArray[i]) == 0 &&
			lenInsertStr == 0 {
			continue
		}

		conCatStr += strArray.StrArray[i]

		if lenInsertStr > 0 {
			conCatStr += insertStr
		}
	}

	return conCatStr
}

// copyStringArrayDto - Copies all data from input parameter
// 'sourceStrArray' to input parameter 'destinationStrArray'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing data fields for input parameter
// 'destinationStrArray' will be overwritten and deleted.
//
// NO DATA VALIDATION is performed on input parameter
// 'sourceStrArray'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationStrArray        *StringArrayDto
//	   - A pointer to an instance of StringArrayDto. Data
//	     extracted from input parameter 'sourceStrArray' will be
//	     copied to this input parameter, 'destinationStrArray'. If
//	     this method completes successfully, all member data
//	     variables encapsulated in 'destinationStrArray' will be
//	     identical to those contained in input parameter,
//	     'destinationStrArray'.
//
//	     Be advised that the pre-existing data fields in input
//	     parameter 'destinationStrArray' will be overwritten and
//	     deleted.
//
//
//	sourceStrArray             *StringArrayDto
//	   - A pointer to an instance of StringArrayDto.
//
//	     All data values in this StringArrayDto instance will be
//	     copied to input parameter 'destinationStrArray'.
//
//	     The original member variable data values encapsulated in
//	     'sourceStrArray' will remain unchanged and will NOT be
//	     overwritten or deleted.
//
//	     NO DATA VALIDATION is performed on 'sourceStrArray'.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (strArrayDtoNanobot *stringArrayDtoNanobot) copyStringArrayDto(
	destinationStrArray *StringArrayDto,
	sourceStrArray *StringArrayDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if strArrayDtoNanobot.lock == nil {
		strArrayDtoNanobot.lock = new(sync.Mutex)
	}

	strArrayDtoNanobot.lock.Lock()

	defer strArrayDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"stringArrayDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if sourceStrArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceStrArray' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationStrArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingStdLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(stringArrayDtoAtom).empty(
		destinationStrArray)

	lenSourceStrArray := len(sourceStrArray.StrArray)

	if lenSourceStrArray > 0 {

		destinationStrArray.StrArray =
			make([]string, lenSourceStrArray)

		for i := 0; i < lenSourceStrArray; i++ {
			destinationStrArray.StrArray[i] =
				sourceStrArray.StrArray[i]
		}

	}

	destinationStrArray.Description1 =
		sourceStrArray.Description1

	destinationStrArray.Description2 =
		sourceStrArray.Description2

	return err
}
