package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberStrStatsDto
//
// A data transport type designed to store and transmit
// information on a numeric value.
type NumberStrStatsDto struct {
	NumOfIntegerDigits uint64
	//	The total number of integer digits to the left of
	//	the radix point or, decimal point, in the subject
	//	numeric value.

	NumOfSignificantIntegerDigits uint64
	//	The number of nonzero integer digits to the left
	//	of the radix point, or decimal point, in the
	//	subject numeric value.

	NumOfFractionalDigits uint64
	//	The total number of fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumOfSignificantFractionalDigits uint64
	//	The number of nonzero fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumberValueType NumericValueType
	//	This enumeration value specifies whether the
	//	subject numeric value is classified either as an
	//	integer or a floating point value.
	//
	//	Possible enumeration values are listed as
	//	follows:
	//  	NumValType.None()
	//  	NumValType.FloatingPoint()
	//  	NumValType.Integer()

	NumberSign NumericSignValueType
	//	An enumeration specifying the number sign
	//	associated with the numeric value. Possible
	//	values are listed as follows:
	//      NumSignVal.None()		= Invalid Value
	//      NumSignVal.Negative()	= -1
	//      NumSignVal.Zero()		=  0
	//      NumSignVal.Positive()	=  1

	IsZeroValue bool
	//	If 'true', the subject numeric value is equal
	//	to zero ('0').
	//
	//	If 'false', the subject numeric value is
	//	greater than or less than zero ('0').

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumberStrStatsDto ('incomingNumStrStatsDto') to the
//	data fields of the current NumberStrStatsDto instance
//	('numberStrStatsDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the data fields in the current NumberStrStatsDto
//	instance ('numberStrStatsDto') will be deleted and
//	overwritten.
//
//	No Data Validation will be performed on
//	'incomingSearchResults'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (numberStrStatsDto *NumberStrStatsDto) CopyIn(
	incomingNumStrStatsDto *NumberStrStatsDto,
	errorPrefix interface{}) error {

	if numberStrStatsDto.lock == nil {
		numberStrStatsDto.lock = new(sync.Mutex)
	}

	numberStrStatsDto.lock.Lock()

	defer numberStrStatsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrStatsDto."+
			"CopyIn()",
		"")

	if err != nil {

		return err

	}

	return new(numberStrStatsDtoNanobot).copyNumStatsDto(
		numberStrStatsDto,
		incomingNumStrStatsDto,
		ePrefix.XCpy(
			"numberStrStatsDto<-"+
				"incomingNumStrStatsDto"))
}

//	CopyOut
//
//	Returns a deep copy of the current NumberStrStatsDto
//	instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No Data Validation will be performed on the current
//	instance of NumberStrStatsDto.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	NumberStrStatsDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		NumberStrStatsDto instance.
//
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (numberStrStatsDto *NumberStrStatsDto) CopyOut(
	errorPrefix interface{}) (
	NumberStrStatsDto,
	error) {

	if numberStrStatsDto.lock == nil {
		numberStrStatsDto.lock = new(sync.Mutex)
	}

	numberStrStatsDto.lock.Lock()

	defer numberStrStatsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newNumStrStatsDto NumberStrStatsDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrStatsDto."+
			"CopyOut()",
		"")

	if err != nil {

		return newNumStrStatsDto, err

	}

	err = new(numberStrStatsDtoNanobot).copyNumStatsDto(
		&newNumStrStatsDto,
		numberStrStatsDto,
		ePrefix.XCpy(
			"newNumStrStatsDto<-"+
				"numberStrStatsDto"))

	return newNumStrStatsDto, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumberStrStatsDto to their zero or
//	uninitialized states. This method will leave the
//	current instance of NumberStrStatsDto in an invalid
//	state and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in this current instance of NumberStrStatsDto.
//	All member variable data values will be reset to
//	their zero or uninitialized states.
//
// ----------------------------------------------------------------
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
func (numberStrStatsDto *NumberStrStatsDto) Empty() {

	if numberStrStatsDto.lock == nil {
		numberStrStatsDto.lock = new(sync.Mutex)
	}

	numberStrStatsDto.lock.Lock()

	new(numberStrStatsDtoMolecule).empty(
		numberStrStatsDto)

	numberStrStatsDto.lock.Unlock()

	numberStrStatsDto.lock = nil

	return
}

//	Equal
//
//	Receives a pointer to another instance of
//	NumberStrStatsDto and proceeds to compare the member
//	variables to those of the current NumberStrStatsDto
//	instance in order to determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables of both instances
//	are equal in all respects, this flag is set to 'true'.
//
//	Otherwise, if the two instances are NOT equal, this
//	method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrStatsDto		*NumberStrStatsDto
//
//		A pointer to an incoming instance of
//		NumberStrStatsDto. This method will compare all
//		member variable data values in this instance
//		against those corresponding data values contained
//		in the current instance of NumberStrStatsDto.
//
//		If the data values in both instances are found to
//		be equal in all respects, this method will return
//		a boolean value of 'true'.
//
//		If the data values in both instances are NOT
//		equivalent, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingNumStrStatsDto' are
//		equal in all respects to those contained in the
//		current instance of NumberStrStatsDto, this
//		method will return a boolean value of 'true'.
//
//		Otherwise, a value of 'false' will be returned.
func (numberStrStatsDto *NumberStrStatsDto) Equal(
	incomingNumStrStatsDto *NumberStrStatsDto) bool {

	if numberStrStatsDto.lock == nil {
		numberStrStatsDto.lock = new(sync.Mutex)
	}

	numberStrStatsDto.lock.Lock()

	defer numberStrStatsDto.lock.Unlock()

	return new(numberStrStatsDtoMolecule).equal(
		numberStrStatsDto,
		incomingNumStrStatsDto)
}

// String
//
// Returns a list of member variables and their values
// for the current instance of NumberStrStatsDto.
func (numberStrStatsDto *NumberStrStatsDto) String() string {

	if numberStrStatsDto.lock == nil {
		numberStrStatsDto.lock = new(sync.Mutex)
	}

	numberStrStatsDto.lock.Lock()

	defer numberStrStatsDto.lock.Unlock()

	var currentValues string

	currentValues += fmt.Sprintf(" NumOfIntegerDigits: %v\n",
		numberStrStatsDto.NumOfIntegerDigits)

	currentValues += fmt.Sprintf(" NumOfSignificantIntegerDigits: %v\n",
		numberStrStatsDto.NumOfSignificantIntegerDigits)

	currentValues += fmt.Sprintf(" NumOfFractionalDigits: %v\n",
		numberStrStatsDto.NumOfFractionalDigits)

	currentValues += fmt.Sprintf(" NumOfSignificantFractionalDigits: %v\n",
		numberStrStatsDto.NumOfSignificantFractionalDigits)

	currentValues += fmt.Sprintf(" NumberValueType: %v\n",
		numberStrStatsDto.NumberValueType.String())

	currentValues += fmt.Sprintf(" NumberSign: %v\n",
		numberStrStatsDto.NumberSign.String())

	currentValues += fmt.Sprintf(" IsZeroValue: %t\n",
		numberStrStatsDto.IsZeroValue)

	return currentValues
}

// numberStrStatsDtoNanobot
//
// Provides helper methods for type NumberStrStatsDto.
type numberStrStatsDtoNanobot struct {
	lock *sync.Mutex
}

// copyNumStatsDto
//
// Copies NumberStrStatsDto information from a source
// object to a destination object.
func (numStrStatsNanobot *numberStrStatsDtoNanobot) copyNumStatsDto(
	destinationNumStatsDto *NumberStrStatsDto,
	sourceNumStatsDto *NumberStrStatsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrStatsNanobot.lock == nil {
		numStrStatsNanobot.lock = new(sync.Mutex)
	}

	numStrStatsNanobot.lock.Lock()

	defer numStrStatsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrStatsDtoNanobot."+
			"copyNumStatsDto()",
		"")

	if err != nil {

		return err

	}

	if destinationNumStatsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationNumStatsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNumStatsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceNumStatsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	destinationNumStatsDto.NumOfIntegerDigits =
		sourceNumStatsDto.NumOfIntegerDigits

	destinationNumStatsDto.NumOfSignificantIntegerDigits =
		sourceNumStatsDto.NumOfSignificantIntegerDigits

	destinationNumStatsDto.NumOfFractionalDigits =
		sourceNumStatsDto.NumOfFractionalDigits

	destinationNumStatsDto.NumOfSignificantFractionalDigits =
		sourceNumStatsDto.NumOfSignificantFractionalDigits

	destinationNumStatsDto.NumberValueType =
		sourceNumStatsDto.NumberValueType

	destinationNumStatsDto.NumberSign =
		sourceNumStatsDto.NumberSign

	destinationNumStatsDto.IsZeroValue =
		sourceNumStatsDto.IsZeroValue

	return err
}

// numberStrStatsDtoMolecule
//
// Provides helper methods for type NumberStrStatsDto.
type numberStrStatsDtoMolecule struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of NumberStrStatsDto
// and proceeds to set all internal member variables to
// their zero or uninitialized states.
//
// This method will therefore delete all data currently
// held by this instance of NumberStrStatsDto.
func (nStrStatsDtoMolecule *numberStrStatsDtoMolecule) empty(
	numStatsDto *NumberStrStatsDto) {

	if nStrStatsDtoMolecule.lock == nil {
		nStrStatsDtoMolecule.lock = new(sync.Mutex)
	}

	nStrStatsDtoMolecule.lock.Lock()

	defer nStrStatsDtoMolecule.lock.Unlock()

	numStatsDto.NumOfIntegerDigits = 0

	numStatsDto.NumOfSignificantIntegerDigits = 0

	numStatsDto.NumOfFractionalDigits = 0

	numStatsDto.NumOfSignificantFractionalDigits = 0

	numStatsDto.NumberValueType = 0

	numStatsDto.NumberSign = 0

	numStatsDto.IsZeroValue = false

}

// equal
//
// Receives pointers to two instances of
// NumberStrStatsDto and proceeds to compare all the
// corresponding member data variables for both
// instances.
//
// If the two instances of NumberStrStatsDto are found
// to be equal in all respects, this method will return
// a boolean value of 'true'.
//
// If the two instances are NOT equal, this method
// returns 'false'.
func (nStrStatsDtoMolecule *numberStrStatsDtoMolecule) equal(
	numStatsDto01 *NumberStrStatsDto,
	numStatsDto02 *NumberStrStatsDto) bool {

	if nStrStatsDtoMolecule.lock == nil {
		nStrStatsDtoMolecule.lock = new(sync.Mutex)
	}

	nStrStatsDtoMolecule.lock.Lock()

	defer nStrStatsDtoMolecule.lock.Unlock()

	if numStatsDto01.NumOfIntegerDigits !=
		numStatsDto02.NumOfIntegerDigits {

		return false
	}

	if numStatsDto01.NumOfSignificantIntegerDigits !=
		numStatsDto02.NumOfSignificantIntegerDigits {

		return false
	}

	if numStatsDto01.NumOfFractionalDigits !=
		numStatsDto02.NumOfFractionalDigits {

		return false
	}

	if numStatsDto01.NumOfSignificantFractionalDigits !=
		numStatsDto02.NumOfSignificantFractionalDigits {

		return false
	}

	if numStatsDto01.NumberValueType !=
		numStatsDto02.NumberValueType {

		return false
	}

	if numStatsDto01.NumberSign !=
		numStatsDto02.NumberSign {

		return false
	}

	if numStatsDto01.IsZeroValue !=
		numStatsDto02.IsZeroValue {

		return false
	}

	return true
}
