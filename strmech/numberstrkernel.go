package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
)

// NumberStrKernel
//
//	Contains all the base specifications necessary
//	to construct an integer or floating point number
//	string.
type NumberStrKernel struct {
	integerDigits RuneArrayDto
	// An array of numeric digits constituting the
	// integer portion of the numeric value represented
	// by this instance of NumberStrKernel.

	fractionalDigits RuneArrayDto
	// An array of numeric digits constituting the
	// fractional portion of the numeric value
	// represented by this instance of
	// NumberStrKernel.

	numberValueType NumericValueType
	// This enumeration value specifies the type of
	// numeric value contained in the current instance
	// of NumberStrKernel. The contained numeric value
	// is classified either as an integer or a floating
	// point value.
	//
	// Possible enumeration values are listed as
	// follows:
	//  NumValType.None()
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	numberSign NumericSignValueType
	// An enumeration specifying the number sign associated
	// with the numeric value represented by this instance
	// of NumberStrKernel.  Possible values are listed as
	// follows:
	//      NumSignVal.None() - Invalid Value
	//      NumSignVal.Negative() = -1
	//      NumSignVal.Zero()     =  0
	//      NumSignVal.Positive() =  1

	isNonZeroValue bool
	// If 'true', the Numeric Value is greater than
	// or less than zero ('0'). If 'false', the
	// Numeric Value is equal to zero.

	numStrFormatSpec NumStrFormatSpec
	// The default Number String format. If
	// this format is found to be empty or
	// invalid, the United States Signed Number
	// String Format will be used.
	//
	// To set the Default Number String Format
	// Specification, use this method:
	//
	//	NumberStrKernel.SetDefaultNumberStrFormatSpec()

	lock *sync.Mutex
}

//	AddFractionalDigit
//
//	Appends a single numeric digit to the end of the internal
//	fractional digits rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fractionalDigit				rune
//
//		A rune with a numeric character between '0' (zero)
//		and '9' (nine) inclusive. This numeric digit will
//		be appended to the end of the internal member
//		variable 'NumberStrKernel.fractionalDigits'
//		for the current instance of NumberStrKernel.
//
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) AddFractionalDigit(
	fractionalDigit rune,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"AddFractionalDigit()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelAtom).addFractionalDigit(
		numStrKernel,
		fractionalDigit,
		ePrefix)
}

//	AddIntegerDigit
//
//	Appends a single numeric digit to the end of the
//	internal integer digits rune array contained
//	within the current instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigit				rune
//
//		A rune with a numeric character between '0' (zero)
//		and '9' (nine) inclusive. This numeric digit will
//		be appended to the end of the internal member
//		variable 'NumberStrKernel.integerDigits' for the
//		current	instance of NumberStrKernel.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) AddIntegerDigit(
	integerDigit rune,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"AddIntegerDigit()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelAtom).addIntegerDigit(
		numStrKernel,
		integerDigit,
		ePrefix)
}

//	Compare
//
//	Receives a pointer to an external instance of
//	NumberStrKernel ('incomingNumStrKernel') and
//	proceeds to compare the numeric value with that
//	of the current NumberStrKernel instance.
//
//	The numeric value of the current NumberStrKernel
//	instance is compared to that of
//	'incomingNumStrKernel'. The comparison results are
//	returned as one of three integer values:
//
//		-1	= Current Instance numeric value is less
//				than 'incomingNumStrKernel'
//
//		 0	= Current Instance numeric value equal to
//		 		'incomingNumStrKernel'
//
//		+1	= Current Instance numeric value is greater
//				than 'incomingNumStrKernel'
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	incomingNumStrKernel		*NumberStrKernel
//
//		The numeric value of 'incomingNumStrKernel' will
//		be compared to that of the current
//		NumberStrKernel instance. The comparison results
//		will be returned as an integer value.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	comparisonValue				int
//
//		This parameter will return the results of numeric
//		value comparison for the current NumberStrKernel
//		instance versus the input parameters,
//		'incomingNumStrKernel'. The returned integer
//		comparison result will be set to one of three
//		values:
//
//		-1	= Current Instance numeric value is less
//				than 'incomingNumStrKernel'
//
//		 0	= Current Instance numeric value equal to
//		 		'incomingNumStrKernel'
//
//		+1	= Current Instance numeric value is greater
//				than 'incomingNumStrKernel'
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) Compare(
	incomingNumStrKernel *NumberStrKernel,
	errorPrefix interface{}) (
	comparisonValue int,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"Compare()",
		"")

	if err != nil {
		return comparisonValue, err
	}

	comparisonValue,
		err = new(numberStrKernelMechanics).
		compareNumStrKernels(
			numStrKernel,
			incomingNumStrKernel,
			ePrefix.XCpy(
				"numStrKernel vs "+
					"incomingNumStrKernel"))

	return comparisonValue, err
}

// CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumberStrKernel ('incomingNumStrKernel') to the	data
//	fields of the current NumberStrKernel instance
//	('numStrKernel').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data field values in current NumberStrKernel
//	instance ('numStrKernel') will be deleted and
//	overwritten.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT delete data values contained
//	in input paramter 'incomingNumStrKernel'. However,
//	if the 'incomingNumStrKernel' Number String Format
//	Specification is invalid,
//	'incomingNumStrKernel.numStrFormatSpec' will be set
//	to the default US (United States) Number String
//	Format Specification.
//
//	All other 'incomingNumStrKernel' data values will
//	remain unchanged.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingNumStrKernel		*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will NOT change the values of internal member
//		variables contained in this instance.
//
//		All data values in this NumberStrKernel instance
//		will be copied to the current NumberStrKernel
//		instance ('numStrKernel').
//
//		This method will NOT delete data values contained
//		in input paramter 'incomingNumStrKernel'. However,
//		if the 'incomingNumStrKernel' Number String Format
//		Specification is invalid,
//		'incomingNumStrKernel.numStrFormatSpec' will be set
//		to the default US (United States) Number String
//		Format Specification.
//
//		All other 'incomingNumStrKernel' data values will
//		remain unchanged.
//
//		If parameter 'incomingNumStrKernel' is determined to
//		be invalid, an error will be returned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrKernel *NumberStrKernel) CopyIn(
	incomingNumStrKernel *NumberStrKernel,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	err =
		new(numberStrKernelNanobot).copy(
			numStrKernel,
			incomingNumStrKernel,
			ePrefix.XCpy(
				"numStrKernel<-incomingNumStrKernel"))

	return err
}

// CopyOut - Returns a deep copy of the current NumberStrKernel
// instance.
//
// If the current NumberStrKernel instance contains invalid member
// variables, this method will return an error.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT delete data values contained
//	in the current instance of NumberStrKernel. However,
//	if the Number String Format Specification is invalid,
//	'NumberStrKernel.numStrFormatSpec' will be set
//	to the default US (United States) Number String
//	Format Specification.
//
//	All other NumberStrKernel data values will remain
//	unchanged.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
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
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	deepCopyNumStrKernel       NumberStrKernel
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current NumberStrKernel instance.
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
func (numStrKernel *NumberStrKernel) CopyOut(
	errorPrefix interface{}) (
	deepCopyNumStrKernel NumberStrKernel,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNumStrKernel, err
	}

	err = new(numberStrKernelNanobot).
		copy(
			&deepCopyNumStrKernel,
			numStrKernel,
			ePrefix.XCpy(
				"deepCopyNumStrKernel<-"+
					"numStrKernel"))

	return deepCopyNumStrKernel, err
}

//	DeleteLeadingTrailingFractionalChars
//
//	Deletes Leading or Trailing characters from the
//	internal fractional rune array:
//
//			numStrKernel.fractionalDigits
//
//	Leading numeric digit characters will be deleted
//	from the beginning of the fractional rune array, OR,
//	Trailing numeric digit characters will be deleted
//	from the end of the fractional rune array, depending
//	on the setting for input parameter
//	'deleteTrailingChars'.
//
//	If 'deleteTrailingChars' is set to 'false', Leading
//	numeric digit characters are deleted from the
//	beginning of the fractional rune array.
//
//	If 'deleteTrailingChars' is set to 'true', Trailing
//	numeric digit characters are deleted from the
//	end of the fractional rune array.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Fractional Digits: "123456"
//		numOfCharsToDelete: 2
//		deleteTrailingChars: false
//		New Fractional Digits: "3456"
//
//	Example-2
//		Fractional Digits: "123456"
//		numOfCharsToDelete: 3
//		deleteTrailingChars: true
//		New Fractional Digits: "123"
//
//	Example-3
//		Fractional Digits: "123456"
//		numOfCharsToDelete: 6
//		deleteTrailingChars: true
//		New Fractional Digits: ""
//
//	Example-4
//		Fractional Digits: "123456"
//		numOfCharsToDelete: 7
//		deleteTrailingChars: false
//		New Fractional Digits: ""
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numOfCharsToDelete			uint64
//
//		This uint64 parameter specifies the number of
//		rune numeric digit characters which will be
//		deleted from the fractional rune array contained
//		in the current instance of NumberStrKernel. These
//		rune numeric digit characters will be deleted
//		from the internal rune array:
//
//			'NumberStrKernel.fractionalDigits'
//
//		If 'numOfCharsToDelete' is set to zero, no rune
//		numeric digit characters will be deleted and no
//		error will be returned.
//
//		If this parameter is set to a value greater than
//		or equal to the length of the rune array, the
//		rune array will be set to 'nil', all fractional
//		numeric digit characters will be deleted and no
//		error will be returned.
//
//	deleteTrailingChars			bool
//
//		This parameter determines whether the rune
//		characters deleted from the rune array will be
//		trailing characters or leading characters.
//
//		If this parameter is set to 'true', trailing
//		characters at the end of the fractional rune
//		array will be deleted.
//
//		If this parameter is set to 'false', leading
//		characters at the beginning of the fractional
//		rune array will be deleted.
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
func (numStrKernel *NumberStrKernel) DeleteLeadingTrailingFractionalChars(
	numOfCharsToDelete uint64,
	deleteTrailingChars bool,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"DeleteLeadingTrailingFractionalChars()",
		"")

	if err != nil {
		return err
	}

	err = numStrKernel.fractionalDigits.
		DeleteLeadingTrailingChars(
			numOfCharsToDelete,
			deleteTrailingChars,
			ePrefix.XCpy(
				"numStrKernel.fractionalDigits<-"))

	if err != nil {
		return err
	}
	_,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel after fractional deletion"))

	return err
}

//	DeleteLeadingTrailingIntegerChars
//
//	Deletes Leading or Trailing characters from the
//	internal integer rune array:
//
//			numStrKernel.integerDigits
//
//	Leading numeric digit characters will be deleted
//	from the beginning of the integer rune array, OR,
//	Trailing numeric digit characters will be deleted
//	from the end of the integer rune array, depending
//	on the setting for input parameter
//	'deleteTrailingChars'.
//
//	If 'deleteTrailingChars' is set to 'false', Leading
//	numeric digit characters are deleted from the
//	beginning of the integer rune array.
//
//	If 'deleteTrailingChars' is set to 'true', Trailing
//	numeric digit characters are deleted from the
//	end of the integer rune array.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Integer Digits: "123456"
//		numOfCharsToDelete: 2
//		deleteTrailingChars: false
//		New Integer Digits: "3456"
//
//	Example-2
//		Integer Digits: "123456"
//		numOfCharsToDelete: 3
//		deleteTrailingChars: true
//		New Integer Digits: "123"
//
//	Example-3
//		Integer Digits: "123456"
//		numOfCharsToDelete: 6
//		deleteTrailingChars: true
//		New Integer Digits: ""
//
//	Example-4
//		Integer Digits: "123456"
//		numOfCharsToDelete: 7
//		deleteTrailingChars: false
//		New Integer Digits: ""
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numOfCharsToDelete			uint64
//
//		This uint64 parameter specifies the number of
//		rune numeric digit characters which will be
//		deleted from the integer rune array contained in
//		the current instance of NumberStrKernel. These
//		rune numeric digit characters will be
//		deleted from the internal rune array:
//
//			'NumberStrKernel.integerDigits'
//
//		If 'numOfCharsToDelete' is set to zero, no rune
//		numeric digit characters will be deleted and no
//		error will be returned.
//
//		If this parameter is set to a value greater than
//		or equal to the length of the rune array, the
//		rune array will be set to 'nil', all integer
//		numeric digit characters will be deleted and no
//		error will be returned.
//
//	deleteTrailingChars			bool
//
//		This parameter determines whether the rune
//		characters deleted from the rune array will be
//		trailing characters or leading characters.
//
//		If this parameter is set to 'true', trailing
//		characters at the end of the integer rune array
//		will be deleted.
//
//		If this parameter is set to 'false', leading
//		characters at the beginning of the integer rune
//		array will be deleted.
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
func (numStrKernel *NumberStrKernel) DeleteLeadingTrailingIntegerChars(
	numOfCharsToDelete uint64,
	deleteTrailingChars bool,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"DeleteLeadingTrailingIntegerChars()",
		"")

	if err != nil {
		return err
	}

	err = numStrKernel.integerDigits.
		DeleteLeadingTrailingChars(
			numOfCharsToDelete,
			deleteTrailingChars,
			ePrefix.XCpy(
				"numStrKernel.integerDigits<-"))

	if err != nil {
		return err
	}
	_,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel after integer deletion"))

	return err
}

// Empty - Resets all internal member variables for the current
// instance of NumberStrKernel to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of NumberStrKernel.
//
// ------------------------------------------------------------------------
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
func (numStrKernel *NumberStrKernel) Empty() {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numStrKernel.lock.Unlock()

	numStrKernel.lock = nil

	return
}

// EmptyFractionalDigits - Sets the Fractional Digits rune array
// to 'nil'. This effectively deletes the previous fractional
// digits stored in that array.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the text characters contained in the member variable
// NumberStrKernel.fractionalDigits will be deleted. The
// NumberStrKernel.fractionalDigits rune array will set to a value
// of 'nil'.
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
func (numStrKernel *NumberStrKernel) EmptyFractionalDigits() {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	_ = new(numberStrKernelAtom).emptyFractionalDigits(
		numStrKernel,
		nil)

	return
}

// EmptyIntegerDigits - Sets the Integer Digits rune array
// to 'nil'. This effectively deletes the previous integer
// digits stored in that array.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the text characters contained in the member variable
// NumberStrKernel.integerDigits will be deleted. The
// NumberStrKernel.integerDigits rune array will set to a value of
// 'nil'.
//
// ------------------------------------------------------------------------
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
func (numStrKernel *NumberStrKernel) EmptyIntegerDigits() {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	_ = new(numberStrKernelAtom).emptyIntegerDigits(
		numStrKernel,
		nil)

	return
}

// Equal - Receives a pointer to another instance of
// NumberStrKernel and proceeds to compare its internal member
// variables to those of the current NumberStrKernel instance in
// order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method differs from method
// NumberStrKernel.EqualNumericDigits() in that this method
// includes all internal member variables in the comparison.
//
// In method NumberStrKernel.EqualNumericDigits() only the integer
// and fractional numeric digits are included in the comparison.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrKernel       *NumberStrKernel
//	   - A pointer to an external instance of NumberStrKernel. The
//	     internal member variable data values in this instance will
//	     be compared to those in the current instance of
//	     NumberStrKernel. The results of this comparison will
//	     be returned to the calling function as a boolean value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the internal member variable data values contained in
//	     input parameter 'incomingNumStrKernel' are equivalent in
//	     all respects to those contained in the current instance of
//	     NumberStrKernel, this return value will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (numStrKernel *NumberStrKernel) Equal(
	incomingNumStrKernel *NumberStrKernel) bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	numStrKernelElectron := numberStrKernelElectron{}

	return numStrKernelElectron.equal(
		numStrKernel,
		incomingNumStrKernel)
}

// EqualFractionalDigits - Receives a pointer to another instance
// of NumberStrKernel and proceeds to compare the fractional digits
// contained in both instances. This necessarily involves a
// comparison of the internal member variables,
// 'NumberStrKernel.fractionalDigits'.
//
// If the fractional digits contained in both instances are
// equivalent, this method returns a boolean value of 'true'.
// Otherwise, a value of 'false' is returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrKernel       *NumberStrKernel
//	   - A pointer to an external instance of NumberStrKernel. The
//	     internal member variable 'fractionalDigits' will be
//	     compared to 'fractionalDigits' member variable in the
//	     current instance of NumberStrKernel. If the fractional
//	     digits in both instances are equivalent, a boolean value
//	     of 'true' will be returned to the calling function.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the fractional digits contained in input parameter
//	     'incomingNumStrKernel' are equivalent in all respects to
//	     those contained in the current instance of
//	     NumberStrKernel, this return value will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (numStrKernel *NumberStrKernel) EqualFractionalDigits(
	incomingNumStrKernel *NumberStrKernel) bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	if incomingNumStrKernel == nil {

		return false
	}

	return numStrKernel.fractionalDigits.EqualCharArrays(
		&incomingNumStrKernel.fractionalDigits)
}

// EqualIntegerDigits - Receives a pointer to another instance of
// NumberStrKernel and proceeds to compare the integer digits
// contained in both instances. This necessarily involves a
// comparison of the internal member variables,
// 'NumberStrKernel.integerDigits'.
//
// If the integer digits contained in both instances are equivalent,
// this method returns a boolean value of 'true'. Otherwise, a
// value of 'false' is returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrKernel       *NumberStrKernel
//	   - A pointer to an external instance of NumberStrKernel. The
//	     internal member variable 'integerDigits' will be compared
//	     to the 'integerDigits' member variable in the current
//	     instance of NumberStrKernel. If the integer digits in both
//	     instances are equivalent, a boolean value of 'true' will
//	     be returned to the calling function.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the integer digits contained in input parameter
//	     'incomingNumStrKernel' are equivalent in all respects to
//	     those contained in the current instance of
//	     NumberStrKernel, this return value will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (numStrKernel *NumberStrKernel) EqualIntegerDigits(
	incomingNumStrKernel *NumberStrKernel) bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	if incomingNumStrKernel == nil {

		return false
	}

	return numStrKernel.integerDigits.EqualCharArrays(
		&incomingNumStrKernel.integerDigits)
}

// EqualNumericDigits - Receives a pointer to another instance of
// NumberStrKernel and proceeds to compare the integer and
// fractional digits contained in both instances. This necessarily
// involves a comparison of the internal member variables,
// 'NumberStrKernel.integerDigits' and
// 'NumberStrKernel.fractionalDigits'.
//
// If the integer and fractional digits contained in both instances
// are equivalent, this method returns a boolean value of 'true'.
// Otherwise, a value of 'false' is returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method differs from method NumberStrKernel.Equal() in that
// only numeric digits are included in the comparison for
// NumberStrKernel.EqualNumericDigits().
//
// In method NumberStrKernel.Equal() all internal member variable
// data values are included in the comparison.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrKernel       *NumberStrKernel
//	   - A pointer to an external instance of NumberStrKernel. The
//	     internal member variables 'integerDigits' and
//	     'fractionalDigits' will be compared to their counterparts
//	     in the current instance of NumberStrKernel. If both the
//	     integer digits and the fractional digits are equivalent,
//	     a boolean value of 'true' will be returned to the calling
//	     function.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the integer and fractional digits contained in input
//	     parameter 'incomingNumStrKernel' are equivalent in all
//	     respects to those contained in the current instance of
//	     NumberStrKernel, this return value will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (numStrKernel *NumberStrKernel) EqualNumericDigits(
	incomingNumStrKernel *NumberStrKernel) bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	if incomingNumStrKernel == nil {

		return false
	}

	areEqual := numStrKernel.integerDigits.EqualCharArrays(
		&incomingNumStrKernel.integerDigits)

	if !areEqual {
		return false
	}

	return numStrKernel.fractionalDigits.EqualCharArrays(
		&incomingNumStrKernel.fractionalDigits)
}

//	ExtendFractionalDigitsArray
//
//	Adds one or more characters to the beginning or end
//	of the fractional rune array contained within the
//	current instance of NumberStrKernel.
//
//	The name of the fractional rune array internal member
//	variable modified by this method is:
//
//			NumberStrKernel.fractionalDigits
//
//	A single numeric digit character will be replicated
//	one or more times as specified by input parameter,
//	'numOfCharsToAdd'.
//
//	This numeric digit character or characters will be
//	added either as leading characters at the beginning
//	of the fractional array or as trailing characters at the
//	end of the fractional array depending on the setting for
//	input paramter, 'addTrailingChars'.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Original Fractional Rune Array: "1234"
//		charToAdd: '0'
//		numOfCharsToAdd: 3
//		addTrailingChars: false
//		New Fractional Rune Array: "0001234"
//
//	Example-2
//		Original Fractional Rune Array: "1234"
//		charToAdd: '0'
//		numOfCharsToAdd: 6
//		addTrailingChars: true
//		New Fractional Rune Array: "1234000000"
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numCharToAdd				rune
//
//		This numeric character digit will be replicated
//		one or more times at the beginning or the end of
//		the fractional rune array contained within the
//		current instance of NumberStrKernel.
//
//		This numeric character digit must be greater than
//		or equal to '0' (zero) and less than or equal to
//		'9' (nine). If the numeric character digit is not
//		in the range '0' - '9' inclusive, an error will
//		be returned.
//
//	numOfCharsToAdd				int
//
//		'numCharToAdd' will be added to the fractional
//		rune array 'numOfCharsToAdd' of times.
//
//		'numOfCharsToAdd' is equal to the total number of
//		characters which will be added to the fractional
//		rune array.
//
//		If the value of 'numOfCharsToAdd' is less than
//		one (1), an error will be returned.
//
//	addTrailingChars			bool
//
//		If this parameter is set to 'true', the
//		additional numeric digit characters will be
//		appended to the fractional rune array as trailing
//		characters, at the end of the fractional rune
//		array.
//
//		If this parameter is set to 'false', the
//		additional numeric digit characters will be
//		prepended to the fractional rune array as leading
//		characters, at the beginning of the fractional
//		rune array.
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
func (numStrKernel *NumberStrKernel) ExtendFractionalDigitsArray(
	numCharToAdd rune,
	numOfCharsToAdd int,
	addTrailingChars bool,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"ExtendFractionalDigitsArray()",
		"")

	if err != nil {
		return err
	}

	if numCharToAdd < '0' ||
		numCharToAdd > '9' {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numCharToAdd' is invalid!\n"+
			"'numCharToAdd' is NOT a valid numeric digit character.\n"+
			"'numCharToAdd' = '%v'\n",
			ePrefix.String(),
			string(numCharToAdd))

		return err
	}

	return numStrKernel.fractionalDigits.ExtendRuneArray(
		numCharToAdd,
		numOfCharsToAdd,
		addTrailingChars,
		ePrefix.XCpy(
			"numStrKernel.fractionalDigits"))
}

//	ExtendIntegerDigitsArray
//
//	Adds one or more characters to the beginning or end
//	of the integer rune array contained within the
//	current instance of NumberStrKernel.
//
//	The name of the integer rune array internal member
//	variable modified by this method is:
//
//			NumberStrKernel.integerDigits
//
//	A single numeric digit character will be replicated
//	one or more times as specified by input parameter,
//	'numOfCharsToAdd'.
//
//	This numeric digit character or characters will be
//	added either as leading characters at the beginning
//	of the integer array or as trailing characters at the
//	end of the integer array depending on the setting for
//	input paramter, 'addTrailingChars'.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Original Integer Rune Array: "1234"
//		charToAdd: '0'
//		numOfCharsToAdd: 3
//		addTrailingChars: false
//		New Integer Rune Array: "0001234"
//
//	Example-2
//		Original Integer Rune Array: "1234"
//		charToAdd: '0'
//		numOfCharsToAdd: 6
//		addTrailingChars: true
//		New Integer Rune Array: "1234000000"
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numCharToAdd				rune
//
//		This numeric character digit will be replicated
//		one or more times at the beginning or the end of
//		the integer rune array contained within the
//		current instance of NumberStrKernel.
//
//		This numeric character digit must be greater than
//		or equal to '0' (zero) and less than or equal to
//		'9' (nine). If the numeric character digit is not
//		in the range '0' - '9' inclusive, an error will
//		be returned.
//
//	numOfCharsToAdd				int
//
//		'numCharToAdd' will be added to the integer rune
//		array 'numOfCharsToAdd' of times.
//
//		'numOfCharsToAdd' is equal to the total number of
//		characters which will be added to the integer
//		rune array.
//
//		If the value of 'numOfCharsToAdd' is less than
//		one (1), an error will be returned.
//
//	addTrailingChars			bool
//
//		If this parameter is set to 'true', the
//		additional numeric digit characters will be
//		appended to the integer rune arrays as trailing
//		characters, at the end of the integer rune array.
//
//		If this parameter is set to 'false', the
//		additional numeric digit characters will be
//		prepended to the integer rune array as leading
//		characters, at the beginning of the integer rune
//		array.
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
func (numStrKernel *NumberStrKernel) ExtendIntegerDigitsArray(
	numCharToAdd rune,
	numOfCharsToAdd int,
	addTrailingChars bool,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"ExtendIntegerDigitsArray()",
		"")

	if err != nil {
		return err
	}

	if numCharToAdd < '0' ||
		numCharToAdd > '9' {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numCharToAdd' is invalid!\n"+
			"'numCharToAdd' is NOT a valid numeric digit character.\n"+
			"'numCharToAdd' = '%v'\n",
			ePrefix.String(),
			string(numCharToAdd))

		return err
	}

	return numStrKernel.integerDigits.ExtendRuneArray(
		numCharToAdd,
		numOfCharsToAdd,
		addTrailingChars,
		ePrefix.XCpy(
			"numStrKernel.integerDigits"))
}

//	FmtCountryCurrencyNumStr
//
//	Creates and returns a formatted number string based
//	on the numeric value contained in the current
//	instance of NumberStrKernel and an instance of
//	NumStrFmtCountryCultureSpec passed as an input
//	parameter.
//
//	Additional input parameters for the output number
//	field and numeric value rounding are also provided.
//
//	This method will produce a Number String configured
//	with Currency Numeric Values.
//
//	For signed number formats generated with
//	Country/Culture parameters, see method:
//
//		NumberStrKernel.FmtCountrySignedNumStr()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		currency formatting information for a
//		designated country or culture.
//
//		This method will NOT modify the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameters 'numberFieldSpec'
//		and 'roundingSpec' to construct and return a
//		Number String configured for Currency Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Currency Number
//		String containing the numeric value of the
//		current NumberStrKernel instance.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtCountryCurrencyNumStr(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtCurrencyNumStrFrance()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCountryCurrencyNumFormat(
		countryCultureFormat,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-countryCultureFormat"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtCountrySignedNumStr
//
//	Creates and returns a formatted number string based
//	on the numeric value contained in the current
//	instance of NumberStrKernel and an instance of
//	NumStrFmtCountryCultureSpec passed as an input
//	parameter.
//
//	Additional input parameters for the output number
//	field and numeric value rounding are also provided.
//
//	This method will produce a Number String configured
//	with Signed Numeric Values.
//
//	For currency formats generated with Country/Culture
//	parameters, see method:
//
//		NumberStrKernel.FmtCountryCurrencyNumStr()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureFormat		NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//
//		The Country Culture Specification contains
//		signed number formatting information for a
//		designated country or culture.
//
//		This method will NOT modify the values of
//		internal member variables contained in this
//		instance.
//
//		The data values in 'countryCultureFormat' will be
//		combined with input parameters 'numberFieldSpec'
//		and 'roundingSpec' to construct and return a
//		Number String configured for Signed Numeric
//		Values.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Signed Number
//		String containing the numeric value of the
//		current NumberStrKernel instance.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtCountrySignedNumStr(
	countryCultureFormat NumStrFmtCountryCultureSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSignedNumStrFrance()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCountrySignedNumFormat(
		countryCultureFormat,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-countryCultureFormat"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtCurrencyNumStrFrance
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Currency Number
//	String formatting parameters typically applied
//	in France.
//
//	The returned number string WILL CONTAIN
//	currency symbols.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either France or
//	Germany.
//
//	For information on German Number String
//	formatting conventions, see method:
//
//		NumberStrKernel.FmtCurrencyNumStrGermany()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The currency symbol used in the France is the
//	Euro symbol ('€').
//
//		French Example-3
//		1 000 000,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro symbol ("€").
//
//		French Example-4
//		-1 000 000,00 €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		French Example-5
//		1 000 000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		French Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Currency Number
//		String containing the numeric value of the
//		current NumberStrKernel instance and the
//		currency formatting typically applied in
//		France.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtCurrencyNumStrFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtCurrencyNumStrFrance()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCurrencyNumFmtFrance(
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))

}

//	FmtCurrencyNumStrGermany
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Currency Number
//	String formatting parameters typically applied
//	in Germany.
//
//	The returned number string WILL CONTAIN
//	currency symbols.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either Germany or
//	France.
//
//	For information on French Number String
//	formatting conventions, see method:
//
//		NumberStrKernel.FmtCurrencyNumStrFrance()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	The currency symbol used in the Germany is the
//	Euro symbol ('€').
//
//		German Example-3
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-4
//		1.000.000,00- €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		German Example-5
//		1.000.000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		German Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Currency Number
//		String containing the numeric value of the
//		current NumberStrKernel instance and the
//		currency formatting typically applied in
//		Germany.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtCurrencyNumStrGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtCurrencyNumStrGermany()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCurrencyNumFmtGermany(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))

}

// FmtCurrencyNumStrUK
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Currency Number
//	String formatting parameters typically applied
//	in the United Kingdom.
//
//	The returned number string WILL CONTAIN
//	currency symbols.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United Kingdom Example-2
//				1,000,000
//
//	The currency symbol used in the United Kingdom is the
//	Pound Sterling symbol ('£').
//
//		United Kingdom Example-3
//			£ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United Kingdom Example-4
//			-£ 1,000,000.00
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			£ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-6
//				£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Currency Number
//		String containing the numeric value of the
//		current NumberStrKernel instance and the
//		currency formatting typically applied in
//		the UK (United Kingdom).
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
func (numStrKernel *NumberStrKernel) FmtCurrencyNumStrUK(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtCurrencyNumStrUK()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCurrencyNumFmtUK(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtCurrencyNumStrUS
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Currency Number
//	String formatting parameters typically applied
//	in the US (United States).
//
//	The returned number string WILL CONTAIN
//	currency symbols.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//			1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United States Example-4
//			$ (1,000,000.00)
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//			$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Currency Number
//		String containing the numeric value of the
//		current NumberStrKernel instance and the
//		currency formatting typically applied in
//		the US (United States).
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtCurrencyNumStrUS(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtCurrencyNumStrUS()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewCurrencyNumFmtUS(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtNumStr
//
//	Returns a formatted number string using the
//	numeric value provided by the current instance
//	of NumberStrKernel.
//
//	Options include customizing for currency
//	symbols, integer separation, number sign
//	management, radix point symbol, and
//	floating point number rounding.
//
//	If required, users have the options of
//	implementing the India or Chinese Numbering
//	Systems.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec				NumStrFormatSpec
//
//		This structure includes all parameters
//		necessary for formatting a number string.
//		These customization options provide maximum
//		granularity in controlling the formatting
//		of the returned Number String.
//
//		type NumStrFormatSpec struct {
//
//			decSeparator			DecimalSeparatorSpec
//
//				Contains the radix point or decimal
//				separator character(s) which will
//				separate integer and fractional
//				numeric digits in a floating point
//				number.
//
//			intSeparatorSpec 		IntegerSeparatorSpec
//
//				Integer Separator Specification. This
//				parameter specifies the type of integer
//				specifies the type of integer grouping and
//				integer separator characters which will be
//				applied to the number string formatting
//				operations.
//
//			positiveNumberSign		NumStrNumberSymbolSpec
//
//				Positive number signs are commonly implied
//				and not specified. However, the user as the
//				option to specify a positive number sign
//				character or characters for positive numeric
//				values using a Number String Positive Number
//				Sign Specification.
//
//				This specification can also be used to
//				configure currency symbols.
//
//			negativeNumberSign		NumStrNumberSymbolSpec
//
//				The Number String Negative Number Sign
//				Specification is used to configure negative
//				number sign symbols for negative numeric values
//				formatted and displayed in number stings.
//
//				This specification can also be used to
//				configured currency symbols.
//
//			numberFieldSpec			NumStrNumberFieldSpec
//
//				This Number String Number Field Specification
//				contains the field length and text
//				justification parameter necessary to display
//				a numeric value within a text number field
//				for display as a number string.
//		}
//
//	roundingSpec				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, a formatted
//		Number String will be returned.
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
func (numStrKernel *NumberStrKernel) FmtNumStr(
	numStrFmtSpec NumStrFormatSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtNumStr()",
		"")

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtNumStrComponents
//
//	Creates and returns a fully formatted Number
//	String generated from Number String formatting
//	components passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec			DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//	intSeparatorSpec			IntegerSeparatorSpec
//
//		Integer Separator Specification. This type
//		encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//	numberSymbolsSpec  			NumStrNumberSymbolsSpec
//
//		This instance of NumStrNumberSymbolsSpec contains the
//		Number Symbol Specifications for negative numeric
//		values, positive numeric values and zero numeric
//		values.
//
//		type NumStrNumberSymbolsSpec struct {
//
//			negativeNumberSign NumStrNumberSymbolSpec
//
//				The Number String Negative Number Sign
//				Specification is used to configure negative
//				number sign symbols for negative numeric
//				values formatted and displayed in number
//				stings.
//
//				For currency presentations, the currency
//				symbol is combined with the negative number
//				sign.
//
//				Example-1: Leading Number Symbols
//					Leading Number Symbols for Negative Values
//
//					Leading Symbols: "- "
//					Number String:   "- 123.456"
//
//				Example-2: Leading Number Symbols With Currency
//					Leading Number Symbols for Negative Values
//
//					Leading Symbols: "$-"
//					Number String:   "$-123.456"
//
//
//				Example-3: Trailing Number Symbols
//					Trailing Number Symbols for Negative Values
//
//					Trailing Symbols: " -"
//					Number String:   "123.456 -"
//
//				Example-4: Trailing Number Symbols
//					Trailing Number Symbols for Negative Values
//
//					Trailing Symbols: "-$"
//					Number String:   "123.456-€"
//
//			positiveNumberSign NumStrNumberSymbolSpec
//
//				Positive number signs are commonly implied
//				and not specified. However, the user has
//				the option to specify a positive number sign
//				character or characters for positive numeric
//				values using a Number String Positive Number
//				Sign Specification.
//
//				For currency presentations, the currency
//				symbol is combined with the positive number
//				sign.
//
//				Example-1: Leading Number Symbols
//					Leading Number Symbols for Positive Values
//
//					Leading Symbols: "+ "
//					Number String:   "+ 123.456"
//
//				Example-2: Leading Number Symbols
//					Leading Number Symbols for Positive Values
//
//					Leading Symbols: "$+"
//					Number String:   "$+123.456"
//
//				Example-3: Leading Number Symbols
//					Leading Number Symbols for Positive Values
//
//					Leading Symbols: "$"
//					Number String:   "$123.456"
//
//				Example-4: Trailing Number Symbols
//					Trailing Number Symbols for Positive Values
//
//					Trailing Symbols: " +"
//					Number String:   "123.456 +"
//
//				Example-5: Trailing Number Symbols
//					Trailing Number Symbols for Positive Values
//
//					Trailing Symbols: "+€"
//					Number String:   "123.456+€"
//
//				Example-6: Trailing Number Symbols
//					Trailing Number Symbols for Positive Values
//
//					Trailing Symbols: " €"
//					Number String:   "123.456 €"
//
//			zeroNumberSign NumStrNumberSymbolSpec
//
//				The Number String Zero Number Symbol
//				Specification is used to configure number
//				symbols for zero numeric values formatted
//				and displayed in number stings. Zero number
//				signs are commonly omitted because zero
//				does not technically qualify as either a
//				positive or negative value. However,
//				currency symbols may be required for zero
//				values.
//
//				For currency presentations, the currency
//				symbol is often used as either a leading
//				or trailing symbol for zero numeric
//				values.
//
//				Example-1: Leading Number Symbols
//					Leading Number Symbols for Zero Values
//
//					Leading Symbols: "$"
//					Trailing Symbols: ""
//					Number String:   "$0.00"
//
//				Example-2: Trailing Number Symbols
//					Trailing Number Symbols for Zero Values
//
//					Leading Symbols: ""
//					Trailing Symbols: " €"
//					Number String:   "0.00 €"
//		}
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	numberString				string
//
//		If this method completes successfully, a formatted
//		Number String will be returned.
//
//	err							error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtNumStrComponents(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	numberSymbolsSpec NumStrNumberSymbolsSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	numberString string,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtNumStrComponents()",
		"")

	if err != nil {
		return numberString, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewNumFmtComponents(
		decSeparatorSpec,
		intSeparatorSpec,
		numberSymbolsSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {
		return numberString, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtNumStrCustom
//
//	Creates and returns a fully formatted Number
//	String generated from Number String formatting
//	components passed as input parameters.
//
//	Options include customizing for currency symbols,
//	integer separation, number sign	management, radix
//	point symbols, and floating point number rounding.
//
//	If required, users have the options of
//	implementing the India or Chinese Numbering
//	Systems.
//
//	This method offers the maximum degree of granular
//	control over all aspects of the Number String
//	formatting operation.
//
//	In particular, it offers maximum flexibility in
//	configuring integer separator characters and
//	integer grouping sequences.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		formatted and returned as a Number String.
//
//	decSeparatorSpec			DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//	intSeparatorSpec			IntegerSeparatorSpec
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Example: 1,000,000,000
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the flexibility
//		necessary to process these complex number separation
//		formats.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in the context of
//				a number string, field length and a 'textJustification'
//				object of type TextJustify. This is because number strings
//				with a field length equal to or less than the length of the
//				numeric value string never use text justification. In these
//				cases, text justification is completely ignored.
//
//				If the field length parameter ('fieldLength') is greater
//				than the length of the numeric value string, text
//				justification must be equal to one of these
//				three valid values:
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text justification
//				enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, a formatted
//		Number String will be returned.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtNumStrCustom(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtNumStr()",
		"")

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelAtom).
		formatNumStrComponents(
			numStrKernel,
			decSeparatorSpec,
			intSeparatorSpec,
			roundingSpec,
			negativeNumberSign,
			positiveNumberSign,
			zeroNumberSign,
			numberFieldSpec,
			ePrefix.XCpy(
				"numStrKernel->"))
}

// FmtNumStrDefaultFormat
//
//	Returns a formatted number string using the
//	numeric value provided by the current instance
//	of NumberStrKernel.
//
//	The number string format is taken from the Default
//	Number String Format Specification previously
//	configured for the current instance of
//	NumberStrKernel.
//
//	To explicitly set the Default Number String Format
//	Specification for the current NumberStrKernel
//	instance, use the following method:
//
//		NumberStrKernel.SetDefaultNumberStrFormatSpec()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT delete or modify data values
//	contained in the current instance of NumberStrKernel.
//	However, if the NumberStrKernel Default Number String
//	Format Specification is invalid,
//	'NumberStrKernel.numStrFormatSpec' will be set
//	to the default US (United States) Signed Number
//	String Format Specification.
//
//	All other NumberStrKernel data values will remain
//	unchanged.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingSpec				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, a formatted
//		Number String will be returned.
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
func (numStrKernel *NumberStrKernel) FmtNumStrDefaultFormat(
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtNumStrDefaultFormat()",
		"")

	if err != nil {
		return numStr, err
	}

	var tempNumStrFormatSpec NumStrFormatSpec

	if new(numStrFmtSpecNanobot).isNOP(
		&numStrKernel.numStrFormatSpec) {

		// The current Number String Format
		// Specification is invalid.
		tempNumStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumFmtUS(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy(
				"tempNumStrFormatSpec<-"))

		if err != nil {

			return numStr, err
		}

	} else {

		// The current Number String Format
		// Specification is valid.

		err = tempNumStrFormatSpec.CopyIn(
			&numStrKernel.numStrFormatSpec,
			ePrefix)

		if err != nil {

			return numStr, err
		}

	}

	numStr,
		err = new(numberStrKernelMolecule).formatNumStr(
		numStrKernel,
		tempNumStrFormatSpec,
		roundingSpec,
		ePrefix.XCpy("numStrKernel"))

	return numStr, err
}

//	FmtSignedNumStrFrance
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Signed Number String
//	formatting parameters typically applied in the
//	France.
//
//	The returned number string will NOT contain
//	currency symbols.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on German Signed Number
//	String formatting conventions, see method:
//
//		NumberStrKernel.FmtSignedNumStrGermany()
//
//	If custom decimal separator, integer separators,
//	or negative number sign characters are required,
//	see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		French Example-4
//		1 000 000 000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		French Example-5
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Signed Number String
//		containing the numeric value of	the current
//		NumberStrKernel instance and the formatting
//		typically applied in France.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtSignedNumStrFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSignedNumStrFrance()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewSignedNumFmtFrance(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtSignedNumStrGermany
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Signed Number String
//	formatting parameters typically applied in the
//	Germany.
//
//	The returned number string will NOT contain
//	currency symbols.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either Germany or France.
//
//	For information on French Signed Number
//	String formatting conventions, see method:
//
//		NumberStrKernel.FmtSignedNumStrFrance()
//
//	If custom decimal separator, integer separators,
//	or negative number sign characters are required,
//	see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-3
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		German Example-4
//		1.000.000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		German Example-5
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Signed Number String
//		containing the numeric value of	the current
//		NumberStrKernel instance and the formatting
//		typically applied in Germany.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtSignedNumStrGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSignedNumStrGermany()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewSignedNumFmtGermany(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtSignedNumStrUK
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Signed Number String
//	formatting parameters typically applied in the
//	UK (United Kingdom).
//
//	The returned number string will NOT contain
//	currency symbols.
//
//	If custom decimal separator, integer separators,
//	or negative number sign characters are required,
//	see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://freeformatter.com/united-kingdom-standards-code-snippets.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United Kingdom Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United Kingdom Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Signed Number String
//		containing the numeric value of	the current
//		NumberStrKernel instance and the formatting
//		typically applied in the UK (United Kingdom).
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtSignedNumStrUK(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSignedNumStrUK()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewSignedNumFmtUK(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtSignedNumStrUS
//
//	Returns a formatted number string based on the
//	numeric value contained in the current instance
//	of NumberStrKernel and the Signed Number String
//	formatting parameters typically applied in the
//	US (United States).
//
//	The returned number string will NOT contain
//	currency symbols.
//
//	If custom decimal separator, integer separators,
//	or negative number sign characters are required,
//	see methods:
//
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//	string
//
//		If this method completes successfully, this
//		parameter will return a Signed Number String
//		containing the numeric value of	the current
//		NumberStrKernel instance and the formatting
//		typically applied in the US (United States).
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) FmtSignedNumStrUS(
	numberFieldSpec NumStrNumberFieldSpec,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	string,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	var numStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSignedNumStrUS()",
		"")

	if err != nil {
		return numStr, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewSignedNumFmtUS(
		numberFieldSpec,
		ePrefix.XCpy(
			"newSignedNumFmtSpec<-"))

	if err != nil {
		return numStr, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	FmtSimpleSignedNumber
//
//	Returns a number string configured with Signed
//	Number String Formatting.
//
//	This method provides a simplified means of creating
//	a Signed Number String using default values.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	string formatting, use one of the more advanced
//	configuration methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumberStrKernel.FmtCountryCurrencyNumStr()
//		NumberStrKernel.FmtCountrySignedNumStr()
//		NumberStrKernel.FmtCurrencyNumStrFrance()
//		NumberStrKernel.FmtCurrencyNumStrGermany()
//		NumberStrKernel.FmtCurrencyNumStrUS()
//		NumberStrKernel.FmtNumStr()
//		NumberStrKernel.FmtNumStrCustom()
//		NumberStrKernel.FmtSignedNumStrFrance()
//		NumberStrKernel.FmtSignedNumStrGermany()
//		NumberStrKernel.FmtSignedNumStrUS()
//
// ----------------------------------------------------------------
//
// # Signed Number Defaults
//
//	Integer Grouping
//		Integers are grouped by thousands or groups
//		of three integers.
//
//		Example: 1,000,000,000
//
//	Negative Signed Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456-"
//			US Number String: "-123.456"
//
//	Positive Signed Number Symbol:
//		No Positive Number Sign Symbol. Positive values
//		are assumed.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string, an error will be returned.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrFormatSpec will configure Number Symbols on
//		the left side of the numeric value. Such Number
//		Symbols are therefore configured as leading
//		Number Symbols. This is the positioning format
//		used in the US, UK, Australia and most of Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrFormatSpec will configure Number Symbols on
//		the right side of the numeric value. Such Number
//		Symbols are therefore configured as trailing
//		Number Symbols. This is the positioning format
//		used in France, Germany and many countries in
//		the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	 errorPrefix                interface{}
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
//		7.  IBasicErrorPrefix
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	numberString				string
//
//		If this method completes successfully, this parameter
//		will return a formatted number string based on the
//		numeric value specified in the current instance of
//		NumberStrKernel. The returned number string will
//		be configured with Signed Number String Formatting.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) FmtSimpleSignedNumber(
	decSeparatorChars string,
	intSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) (
	numberString string,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"FmtSimpleSignedNumber()",
		"")

	if err != nil {
		return numberString, err
	}

	var numStrFmtSpec NumStrFormatSpec

	numStrFmtSpec,
		err = new(NumStrFormatSpec).NewSimpleSignedNumber(
		decSeparatorChars,
		intSeparatorChars,
		leadingNumSymbols,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {
		return numberString, err
	}

	return new(numberStrKernelMolecule).
		formatNumStr(
			numStrKernel,
			numStrFmtSpec,
			roundingSpec,
			ePrefix.XCpy("numStrKernel"))
}

//	GetExcessIntegerLeadingZerosCount
//
//	Returns the count of excess integer leading zeros.
//
//	If the integer value is zero, the count will NOT
//	include the first zero to the left of the decimal
//	point.
//
//	This method differs significantly from method:
//
//		NumberStrKernel.GetIntegerLeadingZerosCount()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Numeric Value:  0001
//		Count of Excess Integer Leading Zeros: 3
//
//	Example-2
//		Numeric Value:  000
//		Count of Excess Integer Leading Zeros: 2
//
//	Example-3
//		Numeric Value:  050.0032
//		Count of Excess Integer Leading Zeros: 1
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	uint64
//
//		Returns the number of excess leading zeros for
//		the integer part of the numeric value represented
//		by this instance of NumberStrKernel.
//
//		NOTE: If the integer value is zero, the returned
//		count of excess leading zeros will NOT include
//		the first zero to the left of the decimal point.
//
//			Example
//				Numeric Value:  000
//				Count of Excess Integer Leading Zeros: 2
func (numStrKernel *NumberStrKernel) GetExcessIntegerLeadingZerosCount() uint64 {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var excessIntLeadingZeros uint64 = 0

	lenIntChars :=
		uint64(len(numStrKernel.integerDigits.CharsArray))

	if lenIntChars == 0 {

		return excessIntLeadingZeros
	}

	excessIntLeadingZeros =
		numStrKernel.integerDigits.GetCountLeadingZeros()

	if excessIntLeadingZeros < lenIntChars {

		return excessIntLeadingZeros
	}

	excessIntLeadingZeros--

	return excessIntLeadingZeros
}

//	GetBigFloatNum
//
//	Returns the numeric value of the current NumberStrKernel
//	instance expressed as a floating point number of type
//	*big.Float.
//
//	If numeric value of the current NumberStrKernel instance
//	is a floating point value, the returned integer value
//	will be rounded according to the rounding policy
//	specified by input parameters, 'roundingType' and
//	'roundToFactionalDigits'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		'roundingType' is only applied in cases where the
//		current NumberStrKernel instance consists of a
//		floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to
//		NumRoundType.NoRounding(), 'roundToFractionalDigits' is
//		ignored and has no effect.
//
//		if 'roundToFractionalDigits' is set to a value greater
//		than the number of fractional digits in 'numStrKernel',
//		the number of fractional digits will be extended with
//		zero values and reflected in the numeric value returned
//		through parameter 'numericValue'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bigFloatNum					*big.Float
//
//		If this method completes successfully, the
//		numeric	value represented by the current instance
//		of NumberStrKernel will be returned as a type
//		*big.Float.
//
//	numOfFractionalDigits		int
//
//		The number of fractional digits contained in the
//		returned numeric value 'bigFloatNum'.
//
//		For returned type *big.Float ('bigFloatNum'), the
//		'numOfFractionalDigits' parameter can be used to
//		improve accuracy in conversions from type *big.Float
//		to character strings.
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetBigFloatNum(
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errorPrefix interface{}) (
	bigFloatNum *big.Float,
	numOfFractionalDigits int,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	bigFloatNum = big.NewFloat(0.0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntNum()",
		"")

	if err != nil {
		return bigFloatNum, numOfFractionalDigits, err
	}

	numOfFractionalDigits,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		bigFloatNum,
		roundingType,
		roundToFactionalDigits,
		ePrefix)

	return bigFloatNum, numOfFractionalDigits, err
}

//	GetBigIntNum
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as an integer value of type *big.Int.
//
//	If numeric value of the current NumberStrKernel instance
//	is a floating point value, the returned *big.Int integer
//	value will be rounded according to the rounding policy
//	specified by input parameter, 'roundingType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Int
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a type *big.Int.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) GetBigIntNum(
	roundingType NumberRoundingType,
	errorPrefix interface{}) (
	*big.Int,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	bigIntValue := big.NewInt(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"Round()",
		"")

	if err != nil {
		return bigIntValue, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		bigIntValue,
		roundingType,
		0,
		ePrefix)

	return bigIntValue, err
}

//	GetBigRatNum
//
//	Converts the current instance of NumberStrKernel to
//	a rational number of type *big.Rat.
//
//	Wikipedia defines a rational number as follows:
//
//	"In mathematics, a rational number is a number that
//	can be expressed as the quotient or fraction of two
//	integers, a numerator 'p' and a non-zero denominator
//	'q'."
//		https://en.wikipedia.org/wiki/Rational_number
//
//	A rational number is defined by the Golang Big Math
//	package as:
//
//	"A Rat represents a quotient a/b of arbitrary
//	precision. The zero value for a Rat represents the
//	value 0."
//		https://pkg.go.dev/math/big#Rat
//
//	Users have the option of rounding the numeric value
//	passed as an instance of NumberStrKernel before that
//	value is converted into a type *big.Rat.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The current instance of NumberStrKernel WILL NOT BE
//	MODIFIED by this method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied
//		for the	rounding of fractional digits before the
//		numeric value is converted to a rational number
//		of type *big.Rat.
//
//		If no rounding of the numeric value is required,
//		set this paramter to: NumRoundType.NoRounding().
//
//		'roundingType' is only applied in cases where the
//		current NumberStrKernel instance consists of a
//		floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits 	int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a.
//		decimal point).
//
//		After completion of a number rounding operation, the
//		value of roundToFractionalDigits will be equal to the
//		number of digits to the right of the decimal point.
//
//		If this parameter is set to a value less than zero,
//		an error will be returned.
func (numStrKernel *NumberStrKernel) GetBigRatNum(
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	bigRatNum *big.Rat,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	bigRatNum = big.NewRat(1, 1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetBigRatNum()",
		"")

	if err != nil {
		return bigRatNum, err
	}

	bigRatNum,
		err = new(numberStrKernelAtom).
		convertKernelToBigRat(
			numStrKernel,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"numStrKernel"))

	return bigRatNum, err
}

//		GetFloat32Num
//
//		Returns the numeric value of the current NumberStrKernel
//		instance as a 32-bit floating point value of type
//		float32.
//
//		The returned numeric value of the current NumberStrKernel
//		instance will be rounded according to the rounding policy
//		specified by input parameter, 'roundingType' and the
//	 target number of fractional digits designated by
//		parameter, 'roundToFractionalDigits'
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that the capacity of a type 'float32' is
//	approximately 7 to 8 digits including integer and
//	fractional digits.
//
//	Source:
//	https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
//
//	Type	Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
//							   field       bits     bias    precision	decimal digits
//	         ----  --------	-----------   -----   --------  ---------	--------------
//	Single	 1		  8			23			32	     127	   24		    ~7.2
//
//	Type float64 provides additional capacity. However,
//	Type *big.Float provides the most accurate representation
//	of floating point numeric values with large numbers of
//	integer	and fractional digits.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to
//		NumRoundType.NoRounding(), 'roundToFractionalDigits' is
//		ignored and has no effect.
//
//		if 'roundToFractionalDigits' is set to a value greater
//		than the number of fractional digits in 'numStrKernel',
//		the number of fractional digits will be extended with
//		zero values and reflected in the numeric value returned
//		through parameter 'numericValue'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	float32
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a 32-bit floating
//		point numeric value of type float32.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetFloat32Num(
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errorPrefix interface{}) (
	float32,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	float32Value := float32(-1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetFloat32Num()",
		"")

	if err != nil {
		return float32Value, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		&float32Value,
		roundingType,
		roundToFactionalDigits,
		ePrefix)

	return float32Value, err
}

//	GetFloat64Num
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as a 64-bit floating point value of type
//	float64.
//
//	The returned numeric value of the current NumberStrKernel
//	instance will be rounded according to the rounding policy
//	specified by input parameter, 'roundingType' and the
//	target number of fractional digits designated by
//	parameter, 'roundToFractionalDigits'.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that the capacity of a type 'float64' is
//	approximately 15 to 17 digits including integer and
//	fractional digits.
//
//	Source:
//	https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
//
//	Type	Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
//							   field       bits     bias    precision	decimal digits
//	         ----  --------	-----------   -----   --------  ---------	--------------
//	Double	  1		 11			52			64	    1023	   53		   ~15.9
//
//	Type *big.Float provides the most accurate representation
//	of floating point numeric values with large numbers of
//	integer	and fractional digits.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to
//		NumRoundType.NoRounding(), 'roundToFractionalDigits' is
//		ignored and has no effect.
//
//		if 'roundToFractionalDigits' is set to a value greater
//		than the number of fractional digits in 'numStrKernel',
//		the number of fractional digits will be extended with
//		zero values and reflected in the numeric value returned
//		through parameter 'numericValue'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	float64
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a 64-bit floating
//		point numeric value of type float64.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetFloat64Num(
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errorPrefix interface{}) (
	float64,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	float64Value := float64(-1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetFloat64Num()",
		"")

	if err != nil {
		return float64Value, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		&float64Value,
		roundingType,
		roundToFactionalDigits,
		ePrefix)

	return float64Value, err
}

// GetFractionalDigits - Returns an instance of RuneArrayDto
// containing a deep copy of the fractional digits rune array
// encapsulated by the current instance of NumberStrKernel.
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
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
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
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//		copyOfFractionalDigits     RuneArrayDto
//	    - If this method completes successfully, the parameter will
//	      return an instance of RuneArrayDto populated with a deep
//	      copy of the fractional numeric digits contained in the
//	      current instance of NumberStrKernel.
//
//
//	 err                        error
//	    - If this method completes successfully and no errors are
//	      encountered this return value is set to 'nil'. Otherwise,
//	      if errors are encountered, this return value will contain
//	      an appropriate error message.
//
//	      If an error message is returned, the text value of input
//	      parameter 'errorPrefix' will be inserted or prefixed at
//	      the beginning of the error message.
func (numStrKernel *NumberStrKernel) GetFractionalDigits(
	errorPrefix interface{}) (
	copyOfFractionalDigits RuneArrayDto,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetFractionalDigits()",
		"")

	if err != nil {
		return copyOfFractionalDigits, err
	}

	if len(numStrKernel.fractionalDigits.CharsArray) == 0 {
		return copyOfFractionalDigits, err
	}

	copyOfFractionalDigits,
		err = numStrKernel.fractionalDigits.CopyOut(
		ePrefix.XCpy("copyOfFractionalDigits<-" +
			"numStrKernel.fractionalDigits"))

	return copyOfFractionalDigits, err
}

// GetFractionalRuneArray - A deep copy of the fractional digits
// rune array encapsulated by the current instance of
// NumberStrKernel.
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
//	[]rune
//	   - This method returns a rune array consisting of the
//	     fractional digits portion of the numeric value represented
//	     by the current instance of NumberStrKernel.  This source
//	     rune array is identified as internal member variable
//	     'NumberStrKernel.fractionalDigits'.
func (numStrKernel *NumberStrKernel) GetFractionalRuneArray() []rune {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.fractionalDigits.GetRuneArray()
}

//	GetFractionalString
//
//	Converts the fractional rune array contained in the
//	current instance of NumberStrKernel to a string
//	and returns it to the calling function.
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
//	string
//
//		This method returns a string consisting of the
//		fractional digits portion of the numeric value
//		represented by the current instance of
//		NumberStrKernel. This source of these fractional
//		digits is identified as internal member variable
//		'NumberStrKernel.fractionalDigits'.
func (numStrKernel *NumberStrKernel) GetFractionalString() string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.fractionalDigits.GetCharacterString()
}

//	GetFractionalRuneArrayDto
//
//	Returns a deep copy of the fractional digits
//	RuneArrayDto object maintained by the current
//	instance of NumberStrKernel.
//
//	This RuneArrayDto object encapsulates the fractional
//	digits rune array containing the fractional portion
//	of the numeric value represented by the current
//	NumberStrKernel instance.
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
//	RuneArrayDto
//
//		This method returns an instance of RuneArrayDto
//		consisting of the fractional digits portion of the
//		numeric value represented by the current instance
//		of NumberStrKernel. This RuneArrayDto member
//		variable identified as:
//
//	     	'NumberStrKernel.fractionalDigits'
//
//		This RuneArrayDto object encapsulates the fractional
//		digits rune array for the current instance of
//		NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetFractionalRuneArrayDto(
	errorPrefix interface{}) (
	RuneArrayDto,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var runeArrayDto RuneArrayDto
	var err error

	runeArrayDto,
		err = numStrKernel.fractionalDigits.CopyOut(
		errorPrefix)

	return runeArrayDto, err
}

//	GetIntNum
//
//	Returns the numeric value of the current NumberStrKernel
//	instance expressed as an integer value of type int.
//
//	If numeric value of the current NumberStrKernel instance
//	is a floating point value, the returned integer value
//	will be rounded according to the rounding policy
//	specified by input parameter, 'roundingType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a type int.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetIntNum(
	roundingType NumberRoundingType,
	errorPrefix interface{}) (
	int,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	intValue := -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntNum()",
		"")

	if err != nil {
		return intValue, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		&intValue,
		roundingType,
		0,
		ePrefix)

	return intValue, err
}

//	GetInt32Num
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as a 32-bit integer value of type int32.
//
//	If numeric value of the current NumberStrKernel instance
//	is a floating point value, the returned integer value
//	will be rounded according to the rounding policy
//	specified by input parameter, 'roundingType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int32
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a type int32.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetInt32Num(
	roundingType NumberRoundingType,
	errorPrefix interface{}) (
	int32,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	int32Value := int32(-1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetInt32Num()",
		"")

	if err != nil {
		return int32Value, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		&int32Value,
		roundingType,
		0,
		ePrefix)

	return int32Value, err
}

//	GetInt64Num
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as a 64-bit integer value of type int64.
//
//	If numeric value of the current NumberStrKernel instance
//	is a floating point value, the returned integer value
//	will be rounded according to the rounding policy
//	specified by input parameter, 'roundingType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		Parameter 'roundingType' will only be applied in
//		cases where the current instance of NumberStrKernel
//		is configured as a floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int64
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a type int64.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernel *NumberStrKernel) GetInt64Num(
	roundingType NumberRoundingType,
	errorPrefix interface{}) (
	int64,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	int64Value := int64(-1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetInt64Num()",
		"")

	if err != nil {
		return int64Value, err
	}

	_,
		err = new(numberStrKernelMolecule).convertKernelToNumber(
		numStrKernel,
		&int64Value,
		roundingType,
		0,
		ePrefix)

	return int64Value, err
}

// GetIntegerDigits - Returns an instance of RuneArrayDto
// containing a deep copy of the integer digits rune array
// encapsulated by the current instance of NumberStrKernel.
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
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
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
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//		copyOfIntegerDigits        RuneArrayDto
//	    - If this method completes successfully, this parameter will
//	      return an instance of RuneArrayDto populated with a deep
//	      copy of the integer numeric digits contained in the
//	      current instance of NumberStrKernel.
//
//
//	 err                        error
//	    - If this method completes successfully and no errors are
//	      encountered this return value is set to 'nil'. Otherwise,
//	      if errors are encountered, this return value will contain
//	      an appropriate error message.
//
//	      If an error message is returned, the text value of input
//	      parameter 'errorPrefix' will be inserted or prefixed at
//	      the beginning of the error message.
func (numStrKernel *NumberStrKernel) GetIntegerDigits(
	errorPrefix interface{}) (
	copyOfIntegerDigits RuneArrayDto,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntegerDigits()",
		"")

	if len(numStrKernel.integerDigits.CharsArray) == 0 {
		return copyOfIntegerDigits, err
	}

	copyOfIntegerDigits,
		err = numStrKernel.integerDigits.CopyOut(
		ePrefix.XCpy("copyOfIntegerDigits<-numStrKernel.integerDigits"))

	return copyOfIntegerDigits, err
}

//	GetIntFracDigitsKernel
//
//	Returns a new instance of NumberStrKernel,
//	'consolidatedIntegerKernel', containing integer
//	digits representing a consolidation of the integer
//	and fractional digits in the current NumberStrKernel
//	instance.
//
//	NOTE: All leading integer zeros are deleted from
//	the returned instance of 'consolidatedIntegerKernel'.
//
//	In addition, this method also returns an int64 value,
//	(exponent) specifying the location of the radix
//	point, or decimal point, in the integer digits.
//
//	The actual numeric value is there for equal to:
//		consolidatedIntegerKernel x 10^exponent
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		 Current NumberStrKernel Value:	1234.5678
//		Returned NumberStrKernel Value:	12345678
//			 Returned 'exponent' Value: -4
//
//	Example-2
//		 Current NumberStrKernel Value:	001234.5678
//		Returned NumberStrKernel Value:	12345678
//			 Returned 'exponent' Value: -4
//		NOTE: Leading Integer Zeros are Deleted!
//
//	Example-3
//		 Current NumberStrKernel Value:	1234
//		Returned NumberStrKernel Value:	1234
//			 Returned 'exponent' Value: 0
//
//	Example-4
//		 Current NumberStrKernel Value:	-1234.5678
//		Returned NumberStrKernel Value:	-12345678
//			 Returned 'exponent' Value: -4
//
//	Example-5
//		 Current NumberStrKernel Value:	-1234
//		Returned NumberStrKernel Value:	-1234
//			 Returned 'exponent' Value: 0
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
//	consolidatedIntegerKernel	NumberStrKernel
//
//		This method returns an instance of
//		NumberStrKernel which consolidates the integer
//		and fractional digits of the current
//		NumberStrKernel instance to form a single
//		integer value.
//
//		Example-1:
//		 	Current NumberStrKernel Value:	1234.5678
//			Returned NumberStrKernel Value:	12345678
//
//		Example-2:
//		 	Current NumberStrKernel Value:	0001234.5678
//			Returned NumberStrKernel Value:	12345678
//
//		NOTE: All leading zeros are deleted in the
//		returned instance of 'consolidatedIntegerKernel'.
//
//	exponent					int64
//
//		This returned exponent value identifies the
//		location of the radix point, or decimal point,
//		within the integer digits returned by parameter,
//		'consolidatedIntegerKernel'.
//
//		Example-1
//			 Current NumberStrKernel Value:	1234.5678
//			Returned NumberStrKernel Value:	12345678
//				 Returned 'exponent' Value: -4
//
//		Example-2
//			 Current NumberStrKernel Value:	1234
//			Returned NumberStrKernel Value:	1234
//				 Returned 'exponent' Value: 0
//
//	err							error
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
func (numStrKernel *NumberStrKernel) GetIntFracDigitsKernel(
	errorPrefix interface{}) (
	consolidatedIntegerKernel NumberStrKernel,
	exponent int64,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntFracDigitsKernel()",
		"")

	if err != nil {
		return consolidatedIntegerKernel, exponent, err
	}

	var consolidatedRuneArrayDto RuneArrayDto

	runeArrayDto := RuneArrayDto{}

	consolidatedRuneArrayDto,
		err = runeArrayDto.NewRuneArrayDtos(
		ePrefix.XCpy(
			"numStrKernel integer+fractional"),
		numStrKernel.integerDigits,
		numStrKernel.fractionalDigits)

	if err != nil {
		return consolidatedIntegerKernel, exponent, err
	}

	nopRuneArrayDto := runeArrayDto.NewNOP()

	err = new(numberStrKernelNanobot).setWithRuneArrayDto(
		&consolidatedIntegerKernel,
		&consolidatedRuneArrayDto,
		&nopRuneArrayDto,
		numStrKernel.numberSign,
		ePrefix)

	leadingIntZeros :=
		consolidatedIntegerKernel.GetIntegerLeadingZerosCount()

	if leadingIntZeros > 0 {

		err = consolidatedIntegerKernel.integerDigits.
			DeleteLeadingTrailingChars(
				leadingIntZeros,
				false,
				ePrefix.XCpy(
					"numStrKernel.integerDigits<-"))

	}

	return consolidatedIntegerKernel, exponent, err
}

//	GetIntegerLeadingZerosCount
//
//	Returns the count of all integer leading zeros.
//
//	This method differs significantly from method:
//
//		NumberStrKernel.GetExcessIntegerLeadingZerosCount()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Numeric Value:  0001
//		Count of Total Integer Leading Zeros: 3
//
//	Example-2
//		Numeric Value:  000
//		Count of Total Integer Leading Zeros: 3
//
//	Example-3
//		Numeric Value:  050.0032
//		Count of Total Integer Leading Zeros: 1
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	uint64
//
//		Returns the number of total leading zeros for
//		the integer part of the numeric value represented
//		by this instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetIntegerLeadingZerosCount() uint64 {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var totalIntLeadingZeros uint64 = 0

	lenIntChars :=
		uint64(len(numStrKernel.integerDigits.CharsArray))

	if lenIntChars == 0 {

		return totalIntLeadingZeros
	}

	totalIntLeadingZeros =
		numStrKernel.integerDigits.GetCountLeadingZeros()

	return totalIntLeadingZeros
}

//	GetIntegerRuneArray
//
//	Returns a deep copy of the integer digits rune array
//	encapsulated by the current instance of
//	NumberStrKernel.
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
//	[]rune
//
//		This method returns a rune array consisting of
//		the integer digits portion of the numeric value
//		represented by the current instance of
//		NumberStrKernel. This rune array is encapsulated
//		in the RuneArrayDto member variable identified
//		as:
//
//	     	'NumberStrKernel.integerDigits'.
func (numStrKernel *NumberStrKernel) GetIntegerRuneArray() []rune {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetRuneArray()
}

//	GetIntegerRuneArrayDto
//
//	Returns a deep copy of the integer digits
//	RuneArrayDto object maintained by the current
//	instance of NumberStrKernel.
//
//	This RuneArrayDto object encapsulates the integer
//	digits rune array containing the integer portion of
//	the numeric value represented by the current
//	NumberStrKernel instance.
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
//	RuneArrayDto
//
//		This method returns an instance of RuneArrayDto
//		consisting of the integer digits portion of the
//		numeric value represented by the current instance
//		of NumberStrKernel. This RuneArrayDto member
//		variable identified as:
//
//	     	'NumberStrKernel.integerDigits'
//
//		This RuneArrayDto object encapsulates the integer
//		digits rune array for the current instance of
//		NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetIntegerRuneArrayDto(
	errorPrefix interface{}) (
	RuneArrayDto,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var runeArrayDto RuneArrayDto
	var err error

	runeArrayDto,
		err = numStrKernel.integerDigits.CopyOut(
		errorPrefix)

	return runeArrayDto, err
}

// GetIntegerString - Converts the integer rune array contained in
// the current instance of NumberStrKernel to a string and returns
// it to the calling function.
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
//	string
//	   - This method returns a string consisting of the integer
//	     digits portion of the numeric value represented by the
//	     current instance of NumberStrKernel. This source of these
//	     integer digits is identified as internal member variable
//	     'NumberStrKernel.integerDigits'.
func (numStrKernel *NumberStrKernel) GetIntegerString() string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetCharacterString()
}

// GetInternalParameterDiagnostics
//
// Returns formatted text output detailing the
// NumberStrKernel member variable names and their
// corresponding values contained in the current instance
// of NumberStrKernel.
//
// If an error is encountered, the error message is
// included in the string returned by this method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		Returns a formatted text string detailing all the
//		internal member variable names and their
//		corresponding values for the current instance of
//		NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetInternalParameterDiagnostics() string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"NumberStrKernel."+
			"GetInternalParameterDiagnostics()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"NumberStrKernel."+
				"GetInternalParameterDiagnostics()",
			err.Error())

		return errOut
	}

	strBuilder := strings.Builder{}

	err = new(numberStrKernelNanobot).
		getParameterTextListing(
			&strBuilder,
			false,
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			ePrefix.String(),
			err.Error())

		return errOut
	}

	return strBuilder.String()
}

//	GetIsNonZeroValue
//
//	Returns a boolean value signaling whether the
//	numeric value represented by the current instance
//	of NumberStrKernel is a Non-Zero value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isNonZeroValue				bool
//
//		If this boolean value is set to false it signals
//		that the numeric value of the current
//		NumberStrKernel instance is equal to zero.
//
//		If this boolean value is set to 'true' it signals
//		that the numeric value of the current
//		NumberStrKernel instance is greater than or less
//		than zero.
func (numStrKernel *NumberStrKernel) GetIsNonZeroValue() (
	isNonZeroValue bool) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	isNonZeroValue,
		_ = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		nil)

	return isNonZeroValue
}

// GetNumberSign - Returns the Number Sign associated with the
// current NumberStrKernel instance as an enumeration value of
// type NumericSignValueType.
//
// If the current number sign value is invalid, an error will be
// returned.
//
// Valid return values for Number Sign are:
//
//	NumSignVal.Negative()
//	NumSignVal.Zero()
//	NumSignVal.Positive()
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
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
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
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NumericSignValueType
//	   - If this method completes successfully, this parameter
//	     will return an enumeration value of type
//	     NumericSignValueType specifying the number sign associated
//	     with the numeric value represented by the current instance
//	     of NumberStrKernel. Possible valid return values are
//	     listed as follows:
//	       NumSignVal.Negative()
//	       NumSignVal.Zero()
//	       NumSignVal.Positive()
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (numStrKernel *NumberStrKernel) GetNumberSign(
	errorPrefix interface{}) (
	NumericSignValueType,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetNumberSign()",
		"")

	if err != nil {
		return NumericSignValueType(-2), err
	}

	if !numStrKernel.numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: The current instance of 'NumberStrKernel' is invalid\n"+
			"because member variable Number Sign Value contains an invalid value.\n"+
			"Number Sign should be set to minus one (-1), zero (0) or plus one (1).\n"+
			"The current value of Number Sign is '%v'\n",
			ePrefix.String(),
			numStrKernel.numberSign.XArithmeticValue())

		return NumericSignValueType(-2), err
	}

	return numStrKernel.numberSign, err
}

// GetNumberOfFractionalDigits - Returns the number of fractional
// digits in the current instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfFractionalDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return len(numStrKernel.fractionalDigits.CharsArray)
}

// GetNumberOfIntegerDigits - Returns the number of integer
// digits in the current instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfIntegerDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return len(numStrKernel.integerDigits.CharsArray)
}

// GetNumberOfNumericDigits - Returns the number of numeric digits
// contained in this instance of NumberStrKernel.
//
// The number of numeric digits is computed by adding the number of
// integer digits plus the number of fractional digits.
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
//	int
//	   - This method returns the number of integer digits plus the
//	     number of fractional digits currently residing in this
//	     instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfNumericDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetRuneArrayLength() +
		numStrKernel.fractionalDigits.GetRuneArrayLength()

}

//	GetNumberSignAsInt
//
//	Returns the Number Sign associated with the
//	current NumberStrKernel instance as an integer
//	value.
//
//	If the current number sign value is invalid,
//	an error will be returned.
//
//	Valid return values for Number Sign are:
//
//		-1 Minus One for negative values
//		 0 Zero for value of zero
//		 1 Plus One (+1) for positive values
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int
//
//		If this method completes successfully, this parameter will
//		return an integer value specifying the number sign
//		associated with the numeric value represented by the
//		current instance of NumberStrKernel. Possible valid return
//		values are listed as follows:
//
//
//			-1 Minus One for negative values
//			 0 Zero for value of zero
//			 1 Plus One (+1) for positive values
//
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) GetNumberSignAsInt(
	errorPrefix interface{}) (
	int,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetNumberSignAsInt()",
		"")

	if err != nil {
		return -99, err
	}

	if !numStrKernel.numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: The current instance of 'NumberStrKernel' is invalid\n"+
			"because member variable Number Sign Value contains an invalid value.\n"+
			"Number Sign should be set to minus one (-1), zero (0) or plus one (1).\n"+
			"The current value of Number Sign is '%v'\n",
			ePrefix.String(),
			numStrKernel.numberSign.XArithmeticValue())

		return -99, err
	}

	return numStrKernel.numberSign.XArithmeticValue(), err
}

//	GetNumericValueStats
//
//	Returns a type of NumberStrStatsDto containing
//	statistics on the nature of the numeric value
//	encapsulated by the current instance of
//	NumberStrKernel.
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
//	numStrStatsDto				NumberStrStatsDto
//
//		This data transfer object will return key
//		statistics on the numeric value encapsulated
//		by the current instance of NumberStrKernel.
//
//		type NumberStrStatsDto struct {
//
//		NumOfIntegerDigits					uint64
//
//			The total number of integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfSignificantIntegerDigits		uint64
//
//			The number of nonzero integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfFractionalDigits				uint64
//
//			The total number of fractional digits to the
//			right of the radix point or, decimal point,
//			in the subject numeric value.
//
//		NumOfSignificantFractionalDigits	uint64
//
//			The number of nonzero fractional digits to
//			the right of the radix point or, decimal
//			point, in the subject numeric value.
//
//		NumberValueType 					NumericValueType
//
//			This enumeration value specifies whether the
//			subject numeric value is classified either as
//			an integer or a floating point number.
//
//			Possible enumeration values are listed as
//			follows:
//				NumValType.None()
//				NumValType.FloatingPoint()
//				NumValType.Integer()
//
//		NumberSign							NumericSignValueType
//
//			An enumeration specifying the number sign
//			associated with the numeric value. Possible
//			values are listed as follows:
//				NumSignVal.None()		= Invalid Value
//				NumSignVal.Negative()	= -1
//				NumSignVal.Zero()		=  0
//				NumSignVal.Positive()	=  1
//
//		IsZeroValue							bool
//
//			If 'true', the subject numeric value is equal
//			to zero ('0').
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//		}
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
func (numStrKernel *NumberStrKernel) GetNumericValueStats(
	errorPrefix interface{}) (
	NumberStrStatsDto,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetNumericValueStats()",
		"")

	if err != nil {
		return NumberStrStatsDto{}, err
	}

	return new(numberStrKernelAtom).
		calcNumStrKernelStats(
			numStrKernel,
			ePrefix.XCpy("numStrKernel"))
}

//	GetNumericValueType
//
//	Returns the Numeric Value Type associated with
//	the current instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumericValueType
//		This enumeration value specifies the type of
//		numeric value contained in the current instance
//		of NumberStrKernel. The contained numeric value
//		is classified either as an integer or a floating
//		point value.
//
//		Possible enumeration values are listed as
//		follows:
//			NumValType.None()
//			NumValType.FloatingPoint()
//			NumValType.Integer()
func (numStrKernel *NumberStrKernel) GetNumericValueType() NumericValueType {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var newNumericValueType NumericValueType

	newNumericValueType,
		_ = new(numberStrKernelQuark).getSetNumValueType(
		numStrKernel,
		nil)

	return newNumericValueType
}

// GetParameterTextListing - Returns formatted text output
// detailing the member variable names and their corresponding
// values contained in the current instance of
// NumberStrKernel ('numStrKernel').
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of *strings.Builder. The
//	      formatted text characters produced by this method will be
//	      written to this instance of strings.Builder.
//
//
//		displayFunctionChain       bool
//		   - Set 'displayFunctionChain' to 'true' and a list of the
//		     functions which led to this result will be included in
//		     the formatted text output.
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
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (numStrKernel *NumberStrKernel) GetParameterTextListing(
	strBuilder *strings.Builder,
	displayFunctionChain bool,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetNumberSignAsInt()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelNanobot).
		getParameterTextListing(
			strBuilder,
			displayFunctionChain,
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))
}

//	GetPureNumberStr
//
//	Extracts the numeric value from the current
//	NumberStrKernel instance and returns it as
//	a pure number string.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a floating point number
//		string.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string returned by this method will
//		format negative numeric values with a leading
//		minus sign ('-') at the beginning of the number
//		string.
//
//		Leading minus signs represent the standard means
//		for designating negative numeric values in the
//		US, UK, Australia, most of Canada and many other
//		countries.
//
//		Example Leading Minus Sign:
//			"-123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string returned by this method will
//		format negative numeric values with a trailing
//		minus signs ('-') located at the end of the
//		number string.
//
//		Trailing minus signs represent the standard for
//		France, Germany and many countries in the
//		European Union.
//
//		Example Trailing Number Symbols:
//			"123.456-"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, a pure
//		number string containing the numeric value
//		extracted from the current instance of
//		NumberStrKernel will be returned.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator.
//				Decimal separators separate integer and
//				fractional numeric digits in a pure
//				number string. The decimal separator may
//				consist of one or more text characters.
//
//				In the US, UK, Australia, most of Canada
//				and many other countries, the decimal
//				separator is the period character ('.')
//				known as the decimal point.
//
//				In France, Germany and many countries in
//				the European Union, the Decimal Separator
//				is the comma character (',').
//
//			3.	Optional: A Pure Number String may
//				include a negative number sign symbol
//				consisting of a minus sign ('-'). The
//				minus sign will identify the numeric
//				value contained in the pure number string
//				as a negative number. Only the minus sign
//				('-') classifies a numeric value as a
//				negative number in a pure number string.
//
//				If a leading or trailing minus sign ('-')
//				is NOT present,	the numeric value is
//				positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				included in the pure number string.
func (numStrKernel *NumberStrKernel) GetPureNumberStr(
	decimalSeparatorChars string,
	leadingMinusSign bool) string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	numberStr := ""

	if leadingMinusSign == true &&
		numStrKernel.numberSign == NumSignVal.Negative() {

		numberStr += "-"
	}

	numberStr +=
		numStrKernel.integerDigits.GetCharacterString()

	if len(numStrKernel.fractionalDigits.CharsArray) > 0 {

		numberStr += decimalSeparatorChars

		numberStr +=
			numStrKernel.fractionalDigits.GetCharacterString()
	}

	if leadingMinusSign == false &&
		numStrKernel.numberSign == NumSignVal.Negative() {

		numberStr += "-"
	}

	return numberStr
}

//	GetScientificNotation
//
//	Converts the numeric value contained in the current
//	instance of NumberStrKernel to a scientific notation
//	value and returns that value as an instance of
//	SciNotationKernel.
//
// ----------------------------------------------------------------
//
//	# Definition of Terms
//
//	In scientific notation, nonzero numbers are written
//	in the form	m × 10n or m times ten raised to the
//	power of n, where n is an integer, and the coefficient
//	m is a nonzero real number (usually between 1 and 10
//	in absolute value, and nearly always written as a
//	terminating decimal).
//
//										Wikipedia
//
//	The Scientific Notation value consists of a
//	significand and exponent which are used to express a
//	numeric value in the form:
//
//		significand x 10^exponent = numeric value
//
//	Example
//
//		Numeric Value				=	265,200,000
//		Scientific Notation Value	=	'2.652 x 10^8'
//		significand 				=	'2.652'
//		significand integer digits 	= 	'2'
//	    significand fractional digits =	'652'
//		exponent    				= 	'8'  (10^8)
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	significandRoundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm which will be applied
//		when rounding fractional digits contained in the
//		significand of the returned Scientific Notation
//		value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	significandRoundToFactionalDigits	int
//
//		This parameter in conjunction with
//		'significandRoundingType' to specify the number of digits
//		to the right of the radix point, or decimal point, which
//		will be returned as the Scientific Notation significand
//		value.
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to NumRoundType.NoRounding(),
//		'significandRoundToFactionalDigits' is ignored and has no
//		effect.
//
//		if 'significandRoundToFactionalDigits' is set to a value
//		greater than the number of fractional digits in the
//		'significand', the number of fractional digits will be
//		extended with zero values and reflected in the numeric
//		value returned through parameter 'numericValue'.
//
//		NOTE: Rounding the significand to zero is considered bad
//		form. Common practice always retains at least one digit
//		to the right of the decimal point in Scientific Notation.
//
//
//	errorPrefix							interface{}
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
//	sciNotKernel				SciNotationKernel
//
//		This returned instance of SciNotationKernel will
//		be configured with the numeric value contained in
//		the current instance of NumberStrKernel.
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
func (numStrKernel *NumberStrKernel) GetScientificNotation(
	significandRoundingType NumberRoundingType,
	significandRoundToFactionalDigits int,
	errorPrefix interface{}) (
	SciNotationKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetScientificNotation()",
		"")

	if err != nil {
		return SciNotationKernel{}, err
	}

	return new(numberStrKernelMechanics).
		convertToSciNotation(
			numStrKernel,
			significandRoundingType,
			significandRoundToFactionalDigits,
			ePrefix.XCpy(
				"numStrKernel"))
}

// IsFloatingPointValue
//
// If the current instance of NumberStrKernel
// encapsulates a floating point numeric value, this
// method returns 'true'.
//
// If the numeric value is not a floating point value,
// this method returns 'false'.
func (numStrKernel *NumberStrKernel) IsFloatingPointValue() bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var newNumericValueType NumericValueType

	newNumericValueType,
		_ = new(numberStrKernelQuark).getSetNumValueType(
		numStrKernel,
		nil)

	return newNumericValueType == NumValType.FloatingPoint()
}

// IsZeroValue
//
// If this method returns 'true', it means that the
// numeric value of the current NumberStrKernel instance
// is zero (0).
//
// If this method returns 'false' it means that the
// numeric value of current NumberStrKernel instance is
// nonzero. A nonzero value will be less than or greater
// than zero.
func (numStrKernel *NumberStrKernel) IsZeroValue() bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var isNonZeroValue bool

	isNonZeroValue,
		_ = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		nil)

	return !isNonZeroValue
}

//	NewFromBigRat
//
//	Creates and returns a new instance of NumberStrKernel
//	from a Rational Number (big.Rat) passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	bigRatNum					*big.Rat
//
//		The rational number which will be converted to a
//		new instance of NumberStrKernel.
//
//	roundToFractionalDigits 	int
//
//		Controls the number of fractional digits returned
//		to the new instance of NumberStrKernel.
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
//	NumberStrKernel
//
//		If this method completes successfully, a new
//		instance of NumberStrKernel will be returned
//		configured with the numeric value extracted
//		from the Rational Number passed as input
//		parameter 'bigRatNum'.
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
func (numStrKernel *NumberStrKernel) NewFromBigRat(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromBigRat()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	ratFloatStr := bigRatNum.FloatString(roundToFractionalDigits)

	runeArrayDto := RuneArrayDto{
		CharsArray:     []rune(ratFloatStr),
		Description1:   "",
		Description2:   "",
		charSearchType: CharSearchType.LinearTargetStartingIndex(),
	}

	var decSeparatorSpec DecimalSeparatorSpec

	decSeparatorSpec,
		err = new(DecimalSeparatorSpec).NewStr(
		".",
		ePrefix.XCpy(
			"decSeparatorSpec<-"+
				"decSeparatorChars"))

	newNumStrKernel,
		err = new(numberStrKernelElectron).parsePureNumStr(
		runeArrayDto,
		decSeparatorSpec,
		true,
		ePrefix.XCpy(
			"ratFloatStr"))

	return newNumStrKernel, err
}

//	NewFromFloatValue
//
//	Creates a new instance of NumberStrKernel based on
//	a floating point numeric value passed as an empty
//	interface.
//
//	The floating point numeric value may be any one of
//	the following types:
//
//		float32
//		float64
//		*big.Float
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	floatingPointValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
//
//		If 'floatingPointValue' is NOT convertible to
//		one of the types listed above, an error will be
//		returned.
//
//		This numeric value will be used to populate the
//		returned instance of NumberStrKernel.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'floatingPointValue'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromFloatValue(
	floatingPointValue interface{},
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromFloatValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelMolecule).convertFloatNumberToKernel(
		&newNumStrKernel,
		floatingPointValue,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromRuneDto
//
//	Creates and returns a new instance of NumberStrKernel
//	generated from input parameters containing integer
//	digit and fractional digit Rune Array Data Transfer
//	Objects (RuneArrayDto).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the integer digits array contained within
//		the returned instance of NumberStrKernel.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the fractional digits array contained within
//		the returned instance of NumberStrKernel.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value extracted from input parameters
//		'integerDigits' and 'fractionalDigits'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromRuneDto(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	numberSign NumericSignValueType,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromRuneDto()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelNanobot).setWithRuneArrayDto(
		&newNumStrKernel,
		integerDigits,
		fractionalDigits,
		numberSign,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromRuneDigits
//
//	Creates and returns a new instance of NumberStrKernel
//	generated from input parameters containing integer
//	digit and fractional digit rune arrays.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				[]rune
//
//		A rune array containing numeric digits used to
//		populate the integer part of the numeric value
//		contained in the returned instance of
//		NumberStrKernel.
//
//	fractionalDigits			[]rune
//
//		A rune array containing numeric digits used to
//		populate the fractional part of the numeric
//		value contained in the returned instance of
//		NumberStrKernel.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value extracted from input parameters
//		'integerDigits' and 'fractionalDigits'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromRuneDigits(
	integerDigits []rune,
	fractionalDigits []rune,
	numberSign NumericSignValueType,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromRuneDigits()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelNanobot).setWithRunes(
		&newNumStrKernel,
		integerDigits,
		fractionalDigits,
		numberSign,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromSignedIntValue
//
//	Creates a new instance of NumberStrKernel based on a
//	signed integer value passed as an empty interface.
//
//	This signed integer value may be any one of the
//	following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			*big.Int
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedIntValue				interface{}
//
//		Integer numeric values passed by means of this
//		empty interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			*big.Int
//
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'signedIntValue'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromSignedIntValue(
	signedIntValue interface{},
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromSignedIntValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelMolecule).
		convertIntNumberToKernel(
			&newNumStrKernel,
			signedIntValue,
			ePrefix)

	return newNumStrKernel, err
}

//	NewFromStringDigits
//
//	Creates and returns a new instance of NumberStrKernel
//	generated from input parameters containing integer
//	digit and fractional digit strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				string
//
//		A string of numeric digits used to populate the
//		integer part of the numeric value contained in the
//		returned instance of NumberStrKernel
//
//	fractionalDigits			string
//
//		A string of numeric digits used to populate the
//		fractional part of the numeric value contained in
//		the returned instance of NumberStrKernel
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value extracted from input parameters
//		'integerDigits' and 'fractionalDigits'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromStringDigits(
	integerDigits string,
	fractionalDigits string,
	numberSign NumericSignValueType,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromStringDigits()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelNanobot).setWithRunes(
		&newNumStrKernel,
		[]rune(integerDigits),
		[]rune(fractionalDigits),
		numberSign,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromUnsignedIntValue
//
//	Creates a new instance of NumberStrKernel based on
//	an unsigned integer value passed as an empty
//	interface through input parameter,
//	'unsignedIntValue'.
//
//	Since 'unsignedIntValue' is an empty interface, the
//	unsigned integer value may be any one of the
//	following types:
//
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//	Unsigned integer numeric values always default to
//	a positive number sign (+). With this method, the
//	user has the option to specify the number sign
//	(positive or negative) assigned to the converted
//	numeric value using input parameter, 'numberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	unsignedIntValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface must match one of the following
//		types:
//
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = -2 - Set to positive for
//											unsigned integer
//											values
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
//
//		Unsigned integer values are by default converted as
//		positive numeric values. If this parameter is set
//		to NumSignVal.Negative(), the numeric value returned
//		through parameter 'numStrKernel' will be classified
//		as a negative value.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign for the converted numeric
//		value will be set to 'positive'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'unsignedIntValue'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewFromUnsignedIntValue(
	unsignedIntValue interface{},
	numberSign NumericSignValueType,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromUnsignedIntValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelMolecule).
		convertIntNumberToKernel(
			&newNumStrKernel,
			unsignedIntValue,
			ePrefix)

	if err != nil {
		return newNumStrKernel, err
	}

	if numberSign == NumSignVal.Negative() {

		newNumStrKernel.numberSign =
			numberSign

	}

	return newNumStrKernel, err
}

//	NewParseCustomNumberStr
//
//	Receives a raw or dirty number string and proceeds to
//	return the extracted numeric value as a type
//	NumberStrKernel.
//
//	This method provides a series of custom parsing
//	parameters which allows users to exercise granular
//	control over the number string parsing operation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	rawNumStr					string
//
//		This parameter is the raw number string
//		containing numeric digit characters which will
//		be extracted and processed to identify the
//		numeric value. This numeric value will in
//		turn form the basis of the new NumberStrKernel
//		instance returned to the calling function.
//
//	startSearchIndex			int
//
//		The character index in the Raw Number String
//		('rawNumStr') at which the search for numeric
//		digit characters will begin.
//
//		Set this value to zero (0) in order	begin the
//		search for numeric characters at the first
//		character in the Raw Number String ('rawNumStr').
//
//	characterSearchLength		int
//
//		The actual number of characters within the Raw
//		Number String ('rawNumStr') that are included in
//		the search for numeric character digits.
//
//		If this value is set to -1, the search length will
//		be configured to include the last index in
//		'rawNumStr'. In other words the search will proceed
//		to the end of 'rawNumStr'.
//
//	negativeNumSignSearchSpecs	NegNumSearchSpecCollection
//
//		Type NegNumSearchSpecCollection is a collection of
//		NegativeNumberSearchSpec object. This collection
//		represents all the negative number signs which
//		might be located with the Raw Number String
//		('rawNumStr').
//
//		This collections allows the flexibility for
//		locating such negative number signs as leading
//		minus signs ('-'), trailing minus signs ('-'),
//		parentheses ('()') or any other negative number
//		sign characters which might be associated with
//		a particular national or cultural usage.
//
//	decimalSeparator			DecimalSeparatorSpec
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		Raw Number String ('rawNumStr'). In the US, the
//		period ('.') or decimal point separates integer
//		and fractional digits within a floating point
//		numeric value. In Europe, the comma (',')
//		performs this function.
//
//	numParsingTerminators		[]string
//
//		An array of strings. If any one of these strings
//		is encountered while searching the Raw Number
//		String ('rawNumStr'), the search operation will
//		be immediately terminated.
//
//	requestRemainderString		bool
//
//		If this parameter is set to 'true', the Number
//		String Search Results ('numberStrSearchResults')
//		returned by this method	will include the
//		remaining string characters	which were NOT
//		included in the search operation.
//
//		This feature may be useful when either
//		'characterSearchLength' or 'numParsingTerminators'
//		parameters are used to arbitrarily terminate the
//		search operation.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numberStrSearchResults		CharSearchNumStrParseResultsDto
//
//		This structure contains public member variables
//		containing operational data describing the
//		success or failure of the number string parsing
//		operation and the search for numeric digits in
//		the Raw Number String, 'rawNumStr'.
//
//		Users are advised to check the following member
//		variable values after a number string parsing
//		operation:
//
//			FoundNumericDigits		bool
//
//				Signals a successful Number String Parsing
//				operation. When set to 'true', this means
//				one or more numeric	digit characters ('0'
//				through '9' inclusive) were	located in the
//				Raw Number String, 'rawNumStr'.
//
//			FoundNonZeroValue		bool
//
//				When set to 'true' this signals that the
//				search operation has detected one or more
//				nonzero numeric digits.
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed from the Raw Number String,
//		'rawNumStr'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewParseCustomNumberStr(
	rawNumStr string,
	startSearchIndex int,
	characterSearchLength int,
	negativeNumSignSearchSpecs NegNumSearchSpecCollection,
	decimalSeparator DecimalSeparatorSpec,
	numParsingTerminators []string,
	requestRemainderString bool,
	errorPrefix interface{}) (
	numberStrSearchResults CharSearchNumStrParseResultsDto,
	numberStrKernel NumberStrKernel,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewParseCustomNumberStr()",
		"")

	if err != nil {
		return numberStrSearchResults, numberStrKernel, err
	}

	numParsingTerminatorsCol := RuneArrayCollection{}

	lenStrArray := len(numParsingTerminators)

	for i := 0; i < lenStrArray; i++ {

		err = numParsingTerminatorsCol.AddRuneArrayString(
			numParsingTerminators[i],
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf("numParsingTerminators[%v]",
					i)))

		if err != nil {
			return numberStrSearchResults, numberStrKernel, err
		}
	}

	var runeDto RuneArrayDto

	runeDto,
		err = RuneArrayDto{}.NewString(
		rawNumStr,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy("rawNumStr"))

	if err != nil {
		return numberStrSearchResults, numberStrKernel, err
	}

	numberStrSearchResults,
		numberStrKernel,
		err = new(numberStrKernelQuark).extractNumRunes(
		runeDto,
		"rawNumStr",
		startSearchIndex,
		characterSearchLength,
		negativeNumSignSearchSpecs,
		decimalSeparator,
		numParsingTerminatorsCol,
		requestRemainderString,
		ePrefix)

	return numberStrSearchResults, numberStrKernel, err
}

//	NewParseFrenchNumberStr
//
//	This method parses an incoming number string
//	using standard decimal separator and negative
//	number sign parameters commonly applied in
//	France.
//
//	The radix point or decimal separator used to
//	separate integer and fractional digits in a
//	floating point numeric value is assumed to be
//	the comma (',') character.
//
//	The integer digits separator is assumed to be
//	a space character (' ').
//
//	The negative number sign used to designate
//	numbers with a negative value is assumed to
//	be a leading minus sign ('-').
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the decimal separator and negative number
//	formatting standards implemented by either
//	France or Germany.
//
//	For information on German Number String
//	parsing conventions, see method:
//
//		NumberStrKernel.NewParseGermanNumberStr()
//
//	If custom decimal separator and negative number
//	characters are required, see method:
//
//		NumberStrKernel.NewParseCustomNumberStr()
//
// ----------------------------------------------------------------
//
// # Reference
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//	https://freeformatter.com/france-standards-code-snippets.html
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	rawNumStr					string
//
//		This parameter is the raw number string
//		containing numeric digit characters which will
//		be extracted and processed to identify the
//		numeric value. This numeric value will in
//		turn form the basis of the new NumberStrKernel
//		instance returned to the calling function.
//
//	startSearchIndex			int
//
//		The character index in the Raw Number String
//		('rawNumStr') at which the search for numeric
//		digit characters will begin.
//
//		Set this value to zero (0) in order	begin the
//		search for numeric characters at the first
//		character in the Raw Number String ('rawNumStr').
//
//	breakOnCharSearchLength		int
//
//		The actual number of characters within the Raw
//		Number String ('rawNumStr') that are included in
//		the search for numeric character digits.
//
//		If this parameter is set to a value greater than
//		zero ('0'), the Number String Parsing algorithm
//		will search the specified number of text
//		characters and then automatically terminate the
//		search for numeric digits.
//
//		If this value is set to value less than one (+1),
//		the search length will be configured to include
//		the last character in 'rawNumStr'. In other words
//		the search for numeric characters will proceed to
//		the end of the string, 'rawNumStr'.
//
//	breakOnCharDelimiters		[]string
//
//		An array of strings. If any one of these strings
//		is encountered while searching the Raw Number
//		String ('rawNumStr'), the search for numeric
//		digits will be automatically terminated.
//
//		If this string array is set to nil, this
//		parameter will be ignored by the Number String
//		parsing algorithm.
//
//	requestRemainderString		bool
//
//		If this parameter is set to 'true', the Number
//		String Search Results ('numberStrSearchResults')
//		returned by this method	will include the
//		remaining string characters	which were NOT
//		included in the search operation.
//
//		This feature may be useful when either
//		'characterSearchLength' or 'numParsingTerminators'
//		parameters are used to arbitrarily terminate the
//		search operation.
//
//
//	 errorPrefix                interface{}
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
//	numberStrSearchResults		CharSearchNumStrParseResultsDto
//
//		This structure contains public member variables
//		containing operational data describing the
//		success or failure of the number string parsing
//		operation and the search for numeric digits in
//		the Raw Number String, 'rawNumStr'.
//
//		Users are advised to check the following member
//		variable values after a number string parsing
//		operation:
//
//			FoundNumericDigits		bool
//
//				Signals a successful Number String Parsing
//				operation. When set to 'true', this means
//				one or more numeric	digit characters ('0'
//				through '9' inclusive) were	located in the
//				Raw Number String, 'rawNumStr'.
//
//			FoundNonZeroValue		bool
//
//				When set to 'true' this signals that the
//				search operation has detected one or more
//				nonzero numeric digits.
//
//	nStrKernel					NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed from the Raw Number String,
//		'rawNumStr'.
//
//	err							error
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
func (numStrKernel *NumberStrKernel) NewParseFrenchNumberStr(
	rawNumStr string,
	startSearchIndex int,
	breakOnCharSearchLength int,
	breakOnCharDelimiters []string,
	requestRemainderString bool,
	errorPrefix interface{}) (
	numberStrSearchResults CharSearchNumStrParseResultsDto,
	nStrKernel NumberStrKernel,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewParseFrenchNumberStr()",
		"")

	if err != nil {
		return numberStrSearchResults, nStrKernel, err
	}

	numParsingTerminatorsCol := RuneArrayCollection{}

	lenStrArray := len(breakOnCharDelimiters)

	for i := 0; i < lenStrArray; i++ {

		err = numParsingTerminatorsCol.AddStringDefault(
			breakOnCharDelimiters[i],
			ePrefix.XCpy(
				fmt.Sprintf("breakOnCharDelimiters[%v]",
					i)))

		if err != nil {
			return numberStrSearchResults, nStrKernel, err
		}
	}

	var runeDto RuneArrayDto

	runeDto = RuneArrayDto{}.NewStringDefault(
		rawNumStr)

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewStr(
		",",
		ePrefix.XCpy("decSeparator"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading minus sign '-'"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	numberStrSearchResults,
		nStrKernel,
		err = new(numberStrKernelQuark).extractNumRunes(
		runeDto,
		"rawNumStr",
		startSearchIndex,
		breakOnCharSearchLength,
		negativeNumSearchSpecs,
		decSeparator,
		numParsingTerminatorsCol,
		requestRemainderString,
		ePrefix)

	return numberStrSearchResults, nStrKernel, err
}

//	NewParseGermanNumberStr
//
//	This method parses an incoming number string
//	using standard decimal separator and negative
//	number sign parameters commonly applied in
//	Germany.
//
//	The radix point or decimal separator used to
//	separate integer and fractional digits in a
//	floating point numeric value is assumed to be
//	the comma (',') character.
//
//	The integer digits separator is assumed to be
//	a period character ('.').
//
//	The negative number sign used to designate
//	numbers with a negative value is assumed to
//	be a trailing minus sign ('-').
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the decimal separator and negative number
//	formatting standards implemented by either
//	Germany or France.
//
//	For information on French Number String
//	parsing conventions, see method:
//
//		NumberStrKernel.NewParseFrenchNumberStr()
//
//	If custom decimal separator and negative number
//	characters are required, see method:
//
//		NumberStrKernel.NewParseCustomNumberStr()
//
// ----------------------------------------------------------------
//
// # Reference
//
// https://freeformatter.com/germany-standards-code-snippets.html
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	rawNumStr					string
//
//		This parameter is the raw number string
//		containing numeric digit characters which will
//		be extracted and processed to identify the
//		numeric value. This numeric value will in
//		turn form the basis of the new NumberStrKernel
//		instance returned to the calling function.
//
//	startSearchIndex			int
//
//		The character index in the Raw Number String
//		('rawNumStr') at which the search for numeric
//		digit characters will begin.
//
//		Set this value to zero (0) in order	begin the
//		search for numeric characters at the first
//		character in the Raw Number String ('rawNumStr').
//
//	breakOnCharSearchLength		int
//
//		The actual number of characters within the Raw
//		Number String ('rawNumStr') that are included in
//		the search for numeric character digits.
//
//		If this parameter is set to a value greater than
//		zero ('0'), the Number String Parsing algorithm
//		will search the specified number of text
//		characters and then automatically terminate the
//		search for numeric digits.
//
//		If this value is set to value less than one (+1),
//		the search length will be configured to include
//		the last character in 'rawNumStr'. In other words
//		the search for numeric characters will proceed to
//		the end of the string, 'rawNumStr'.
//
//	breakOnCharDelimiters		[]string
//
//		An array of strings. If any one of these strings
//		is encountered while searching the Raw Number
//		String ('rawNumStr'), the search for numeric
//		digits will be automatically terminated.
//
//		If this string array is set to nil, this
//		parameter will be ignored by the Number String
//		parsing algorithm.
//
//	requestRemainderString		bool
//
//		If this parameter is set to 'true', the Number
//		String Search Results ('numberStrSearchResults')
//		returned by this method	will include the
//		remaining string characters	which were NOT
//		included in the search operation.
//
//		This feature may be useful when either
//		'characterSearchLength' or 'numParsingTerminators'
//		parameters are used to arbitrarily terminate the
//		search operation.
//
//	 errorPrefix                interface{}
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
//	numberStrSearchResults		CharSearchNumStrParseResultsDto
//
//		This structure contains public member variables
//		containing operational data describing the
//		success or failure of the number string parsing
//		operation and the search for numeric digits in
//		the Raw Number String, 'rawNumStr'.
//
//		Users are advised to check the following member
//		variable values after a number string parsing
//		operation:
//
//			FoundNumericDigits		bool
//
//				Signals a successful Number String Parsing
//				operation. When set to 'true', this means
//				one or more numeric	digit characters ('0'
//				through '9' inclusive) were	located in the
//				Raw Number String, 'rawNumStr'.
//
//			FoundNonZeroValue		bool
//
//				When set to 'true' this signals that the
//				search operation has detected one or more
//				nonzero numeric digits.
//
//	nStrKernel					NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed from the Raw Number String,
//		'rawNumStr'.
//
//	err							error
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
func (numStrKernel *NumberStrKernel) NewParseGermanNumberStr(
	rawNumStr string,
	startSearchIndex int,
	breakOnCharSearchLength int,
	breakOnCharDelimiters []string,
	requestRemainderString bool,
	errorPrefix interface{}) (
	numberStrSearchResults CharSearchNumStrParseResultsDto,
	nStrKernel NumberStrKernel,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewParseGermanNumberStr()",
		"")

	if err != nil {
		return numberStrSearchResults, nStrKernel, err
	}

	numParsingTerminatorsCol := RuneArrayCollection{}

	lenStrArray := len(breakOnCharDelimiters)

	for i := 0; i < lenStrArray; i++ {

		err = numParsingTerminatorsCol.AddStringDefault(
			breakOnCharDelimiters[i],
			ePrefix.XCpy(
				fmt.Sprintf("breakOnCharDelimiters[%v]",
					i)))

		if err != nil {
			return numberStrSearchResults, nStrKernel, err
		}
	}

	var runeDto RuneArrayDto

	runeDto = RuneArrayDto{}.NewStringDefault(
		rawNumStr)

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewStr(
		",",
		ePrefix.XCpy("decSeparator"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = negativeNumSearchSpecs.AddTrailingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Trailing minus sign '-'"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	numberStrSearchResults,
		nStrKernel,
		err = new(numberStrKernelQuark).extractNumRunes(
		runeDto,
		"rawNumStr",
		startSearchIndex,
		breakOnCharSearchLength,
		negativeNumSearchSpecs,
		decSeparator,
		numParsingTerminatorsCol,
		requestRemainderString,
		ePrefix)

	return numberStrSearchResults, nStrKernel, err
}

//	NewParsePureNumberStr
//
//	Receives a Pure Number String and proceeds to return
//	the extracted numeric value as a type NumberStrKernel.
//
//	A "Pure Number String" is defined as follows:
//
//		1.	Consists of numeric character digits
//			zero through nine inclusive (0-9).
//
//		2.	Option: A Pure Number String may include
//			a radix point or decimal separator.
//			Decimal separators separate integer and
//			fractional numeric digits in a pure
//			number string. The decimal separator may
//			consist of one or more text characters.
//
//			In the US, UK, Australia, most of Canada
//			and many other countries, the decimal
//			separator is the period character ('.')
//			known as the decimal point.
//
//			In France, Germany and many countries in
//			the European Union, the Decimal Separator
//			is the comma character (',').
//
//		3.	Optional: A Pure Number String may
//			include a negative number sign symbol
//			consisting of a minus sign ('-'). The
//			minus sign will identify the numeric
//			value contained in the pure number string
//			as a negative number. Only the minus sign
//			('-') classifies a numeric value as a
//			negative number in a Pure Number String.
//
//			If a leading or trailing minus sign ('-')
//			is NOT present in the pure number string,
//			the numeric value is assumed to be
//			positive.
//
//		4.	Only numeric characters, the decimal
//			separator and the minus sign will be
//			processed by the pure number string
//			parsing algorithm. All other characters
//			will be	ignored.
//
//		5.	Pure Number Strings consist of a single
//			numeric value. The entire Pure Number String
//			will be parsed, or processed, and only one
//			numeric value per Pure Number String will
//			be returned.
//
//	Pure number strings are routinely produced from
//	methods like fmt.Sprintf() when converting numeric
//	values to strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberStr				string
//
//		This strings contains the numeric character
//		digits from	which a numeric value will be
//		extracted and returned as a NumberStrKernel.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator.
//				Decimal separators separate integer and
//				fractional numeric digits in a pure
//				number string. The decimal separator may
//				consist of one or more text characters.
//
//				In the US, UK, Australia, most of Canada
//				and many other countries, the decimal
//				separator is the period character ('.')
//				known as the decimal point.
//
//				In France, Germany and many countries in
//				the European Union, the Decimal Separator
//				is the comma character (',').
//
//			3.	Optional: A Pure Number String may
//				include a negative number sign symbol
//				consisting of a minus sign ('-'). The
//				minus sign will identify the numeric
//				value contained in the pure number string
//				as a negative number. Only the minus sign
//				('-') classifies a numeric value as a
//				negative number in a Pure Number String.
//
//				If a leading or trailing minus sign ('-')
//				is NOT present in the pure number string,
//				the numeric value is assumed to be
//				positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				processed by the pure number string
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted, floating
//		point Number String.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string parsing algorithm will search
//		for a leading minus sign ('-') at the beginning of
//		the number string. Leading minus signs represent
//		the standard means for designating negative
//		numeric values in the US, UK, Australia, most of
//		Canada and many other parts of world.
//
//		Example Leading Minus Sign:
//			"-123.456" or "- 123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the pure number string input parameter,
//		'pureNumberString', will immediately terminate
//		the search for numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed and extracted from input
//		paramter, 'pureNumberStr'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) NewParsePureNumberStr(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errorPrefix interface{}) (
	NumberStrKernel,
	error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewFromBigIntValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:     []rune(pureNumberStr),
		Description1:   "",
		Description2:   "",
		charSearchType: CharSearchType.LinearTargetStartingIndex(),
	}

	var decSeparatorSpec DecimalSeparatorSpec

	decSeparatorSpec,
		err = new(DecimalSeparatorSpec).NewStr(
		decSeparatorChars,
		ePrefix.XCpy(
			"decSeparatorSpec<-"+
				"decSeparatorChars"))

	newNumStrKernel,
		err = new(numberStrKernelElectron).parsePureNumStr(
		runeArrayDto,
		decSeparatorSpec,
		leadingMinusSign,
		ePrefix.XCpy(
			"pureNumberStr"))

	return newNumStrKernel, err
}

//	NewParseUSNumberStr
//
//	This method parses an incoming number string
//	using standard decimal separator and negative
//	number sign parameters commonly applied in the
//	US (United States).
//
//	The radix point or decimal separator used to
//	separate integer and fractional digits in a
//	floating point numeric value is therefore
//	assumed to be the period ('.') or decimal
//	point.
//
//	The negative number signs used to designate
//	numbers with a negative value are assumed to
//	be either a leading minus sign ('-') or
//	parentheses ('()').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	rawNumStr					string
//
//		This parameter is the raw number string
//		containing numeric digit characters which will
//		be extracted and processed to identify the
//		numeric value. This numeric value will in
//		turn form the basis of the new NumberStrKernel
//		instance returned to the calling function.
//
//	startSearchIndex			int
//
//		The character index in the Raw Number String
//		('rawNumStr') at which the search for numeric
//		digit characters will begin.
//
//		Set this value to zero (0) in order	begin the
//		search for numeric characters at the first
//		character in the Raw Number String ('rawNumStr').
//
//	breakOnCharSearchLength		int
//
//		The actual number of characters within the Raw
//		Number String ('rawNumStr') that are included in
//		the search for numeric character digits.
//
//		If this parameter is set to a value greater than
//		zero ('0'), the Number String Parsing algorithm
//		will search the specified number of text
//		characters and then automatically terminate the
//		search for numeric digits.
//
//		If this value is set to value less than one (+1),
//		the search length will be configured to include
//		the last character in 'rawNumStr'. In other words
//		the search for numeric characters will proceed to
//		the end of the string, 'rawNumStr'.
//
//	breakOnCharDelimiters		[]string
//
//		An array of strings. If any one of these strings
//		is encountered while searching the Raw Number
//		String ('rawNumStr'), the search for numeric
//		digits will be automatically terminated.
//
//		If this string array is set to nil, this
//		parameter will be ignored by the Number String
//		parsing algorithm.
//
//	requestRemainderString		bool
//
//		If this parameter is set to 'true', the Number
//		String Search Results ('numberStrSearchResults')
//		returned by this method	will include the
//		remaining string characters	which were NOT
//		included in the search operation.
//
//		This feature may be useful when either
//		'characterSearchLength' or 'numParsingTerminators'
//		parameters are used to arbitrarily terminate the
//		search operation.
//
//	 errorPrefix                interface{}
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
//	numberStrSearchResults		CharSearchNumStrParseResultsDto
//
//		This structure contains public member variables
//		containing operational data describing the
//		success or failure of the number string parsing
//		operation and the search for numeric digits in
//		the Raw Number String, 'rawNumStr'.
//
//		Users are advised to check the following member
//		variable values after a number string parsing
//		operation:
//
//			FoundNumericDigits		bool
//
//				Signals a successful Number String Parsing
//				operation. When set to 'true', this means
//				one or more numeric	digit characters ('0'
//				through '9' inclusive) were	located in the
//				Raw Number String, 'rawNumStr'.
//
//			FoundNonZeroValue		bool
//
//				When set to 'true' this signals that the
//				search operation has detected one or more
//				nonzero numeric digits.
//
//	nStrKernel					NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed from the Raw Number String,
//		'rawNumStr'.
//
//	err							error
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
func (numStrKernel *NumberStrKernel) NewParseUSNumberStr(
	rawNumStr string,
	startSearchIndex int,
	breakOnCharSearchLength int,
	breakOnCharDelimiters []string,
	requestRemainderString bool,
	errorPrefix interface{}) (
	numberStrSearchResults CharSearchNumStrParseResultsDto,
	nStrKernel NumberStrKernel,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"NewParseUSNumberStr()",
		"")

	if err != nil {
		return numberStrSearchResults, nStrKernel, err
	}

	numParsingTerminatorsCol := RuneArrayCollection{}

	lenStrArray := len(breakOnCharDelimiters)

	for i := 0; i < lenStrArray; i++ {

		err = numParsingTerminatorsCol.AddRuneArrayString(
			breakOnCharDelimiters[i],
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf("breakOnCharDelimiters[%v]",
					i)))

		if err != nil {
			return numberStrSearchResults, nStrKernel, err
		}
	}

	var runeDto RuneArrayDto

	runeDto,
		err = RuneArrayDto{}.NewString(
		rawNumStr,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy("rawNumStr"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewStr(
		".",
		ePrefix.XCpy("decSeparator<-"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading minus sign '-'"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	err = negativeNumSearchSpecs.AddLeadingAndTrailingNegNumSearchStr(
		"(",
		")",
		ePrefix.XCpy("Leading Trailing '()'"))

	if err != nil {

		return numberStrSearchResults, nStrKernel, err
	}

	numberStrSearchResults,
		nStrKernel,
		err = new(numberStrKernelQuark).extractNumRunes(
		runeDto,
		"rawNumStr",
		startSearchIndex,
		breakOnCharSearchLength,
		negativeNumSearchSpecs,
		decSeparator,
		numParsingTerminatorsCol,
		requestRemainderString,
		ePrefix)

	return numberStrSearchResults, nStrKernel, err
}

// RationalizeFractionalIntegerDigits
//
// If fractional digits are present in this instance of
// NumberStrKernel, this method will ensure that integer
// digits are also present.
//
// If fractional digits are present and no integer digits
// are found, this method will configure a zero ('0') in
// the integer digits rune array.
//
// Example:
//
//	.752 will be converted to 0.752
func (numStrKernel *NumberStrKernel) RationalizeFractionalIntegerDigits() {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	_ = new(numberStrKernelElectron).rationalizeFractionalIntegerDigits(
		numStrKernel,
		nil)

	return
}

//	Round
//
//	Rounds the fractional digits contained in the current
//	instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method may alter the numeric value contained
//	within the current instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the radix
//		point or decimal separator (a.k.a. decimal point). This
//		controls the number of fractional digits remaining after
//		completion of the number rounding operation.
//
//	 errorPrefix                interface{}
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
func (numStrKernel *NumberStrKernel) Round(
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"Round()",
		"")

	if err != nil {
		return err
	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {
		return err
	}

	return new(numStrMathRoundingNanobot).roundNumStrKernel(
		numStrKernel,
		numStrRoundingSpec,
		ePrefix)
}

// SetDefaultNumberStrFormatSpec
//
// Sets the default Number String Format to be applied in
// formatting the number string for screen display,
// file output or printing.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	defaultNumStrFmt			NumStrFormatSpec
//
//		An instance of Number String Format Specification
//		('NumStrFormatSpec').
//
//		A deep copy of this instance will be stored in the
//		current instance of NumberStrKernel. This format
//		specification will be used to format the number
//		string for screen displays, file output or
//		printing.
//
//		If 'defaultNumStrFmt' is found to be invalid, an
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
func (numStrKernel *NumberStrKernel) SetDefaultNumberStrFormatSpec(
	defaultNumStrFmt NumStrFormatSpec,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetDefaultNumberStrFormatSpec()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(numStrFmtSpecAtom).
		testValidityNumStrFormatSpec(
			&defaultNumStrFmt,
			ePrefix.XCpy(
				"defaultNumStrFmt invalid!"))

	if err != nil {
		return err
	}

	err = numStrKernel.numStrFormatSpec.CopyIn(
		&defaultNumStrFmt,
		ePrefix.XCpy(
			"numStrKernel.numStrFormatSpec<-defaultNumStrFmt"))

	return err
}

//	SetFloatValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the floating
//	point numeric value passed as an empty interface.
//
//	The floating point numeric value passed to this
//	method by input parameter '' MUST BE convertible
//	to one of the floating point types defined below:
//
//		float32
//		float64
//		*big.Float
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of NumberStrKernel, all
//	pre-existing data will be deleted and overwritten
//	with new data generated from the numeric value
//	passed as input parameter, 'floatingPointValue'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	floatingPointValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
//
//		This floating point numeric value will be used
//		to configure and populate the current instance
//		of NumberStrKernel.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'floatingPointValue'.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetFloatValue(
	floatingPointValue interface{},
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetFloatValue()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelMolecule).
		convertFloatNumberToKernel(
			numStrKernel,
			floatingPointValue,
			ePrefix)
}

// SetNumberSign - Sets the Number Sign for the numeric value
// represented by the current instance of NumberStrKernel.
//
// The Number Sign is specified by means of a
// NumericSignValueType enumeration value.
//
// Possible values are listed as follows:
//
//	NumSignVal.None()     = -2 - Invalid Value
//	NumSignVal.Negative() = -1 - Valid Value
//	NumSignVal.Zero()     =  0 - Valid Value
//	NumSignVal.Positive() =  1 - Valid Value
func (numStrKernel *NumberStrKernel) SetNumberSign(
	numberSign NumericSignValueType,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetNumberSign()",
		"")

	if err != nil {
		return err
	}

	if !numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberSign' is invalid\n"+
			"Number Sign must be set to one of the following"+
			"three valid values:\n"+
			"  NumSignVal.Negative()\n"+
			"  NumSignVal.Zero()\n"+
			"  NumSignVal.Positive()"+
			"Input parameter Number Sign Integer Value = '%v'\n"+
			"Input parameter Number Sign  String Value = '%v'\n",
			ePrefix.String(),
			numberSign.XValueInt(),
			numberSign.String())

		return err
	}

	numStrKernel.numberSign = numberSign

	return err
}

// SetNumberSignInt - Sets the Number Sign of the current
// NumberStrKernel instance based on an integer input value.
//
// Possible valid integer input values are listed as follows:
//
//	-1 - Minus One
//	 0 - Zero
//	 1 - Plus One
//
// Invalid input value will generate an error.
func (numStrKernel *NumberStrKernel) SetNumberSignInt(
	numberSign int,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetNumberSignInt()",
		"")

	if err != nil {
		return err
	}

	if numberSign > 1 ||
		numberSign < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberSign' is invalid!\n"+
			"'numberSign' must be set to one of three values:"+
			"     0 = Numeric Value of Zero ('0').\n"+
			"     1 = + or positive number sign\n"+
			"    -1 = - or negative number sign\n"+
			"Input parameter 'numberSign' = '%v'\n",
			ePrefix.String(),
			numberSign)

		return err
	}

	if numberSign == 0 {
		numStrKernel.isNonZeroValue = false
	} else {
		numStrKernel.isNonZeroValue = true
	}

	numStrKernel.numberSign = NumericSignValueType(numberSign)

	return err
}

//	SetRuneDigits
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using input parameters
//	containing integer digit and fractional digit
//	rune arrays.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that all the data fields within the
//	current instance of NumberStrKernel will be
//	deleted and reset to new values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				[]rune
//
//		This rune array contains the numeric digits used to
//		populate the integer part of the numeric value
//		contained in the current instance of NumberStrKernel.
//
//	fractionalDigits			[]rune
//
//		This rune array contains the numeric digits used to
//		populate the fractional part of the numeric value
//		contained in the current instance of NumberStrKernel.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign for the new numeric value is
//		specified by means of a	NumericSignValueType
//		enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetRuneDigits(
	integerDigits []rune,
	fractionalDigits []rune,
	numberSign NumericSignValueType,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetRuneDigits()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelNanobot).setWithRunes(
		&newNumStrKernel,
		integerDigits,
		fractionalDigits,
		numberSign,
		ePrefix)
}

//	SetRuneDto
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using input parameters
//	containing integer digits and fractional digits
//	formatted as Rune Array Data Transfer Objects.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that all the data fields within the
//	current instance of NumberStrKernel will be
//	deleted and reset to new values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the integer digits array contained within
//		the current instance of NumberStrKernel.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the fractional digits array contained within
//		the current instance of NumberStrKernel.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign for the new numeric value is
//		specified by means of a	NumericSignValueType
//		enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetRuneDto(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	numberSign NumericSignValueType,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetRuneDto()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelNanobot).setWithRuneArrayDto(
		&newNumStrKernel,
		integerDigits,
		fractionalDigits,
		numberSign,
		ePrefix)
}

//	SetSignedIntegerValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the integer
//	numeric value passed as input parameter,
//	'signedIntegerValue'.
//
//	'signedIntegerValue' is an empty interface. This
//	object MUST BE convertible to one of the signed
//	integer value types defined below:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			*big.Int
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of NumberStrKernel, all
//	pre-existing data will be deleted and overwritten
//	with new data generated from the signed numeric
//	value passed as input parameter, 'signedIntValue'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedIntegerValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			*big.Int
//
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetSignedIntegerValue(
	signedIntegerValue interface{},
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetSignedIntegerValue()",
		"")

	if err != nil {
		return err
	}

	err = new(numberStrKernelMolecule).
		convertIntNumberToKernel(
			numStrKernel,
			signedIntegerValue,
			ePrefix)

	return err
}

//	SetStringDigits
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using input parameters
//	containing integer digit and fractional digit
//	strings.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that all the data fields within the
//	current instance of NumberStrKernel will be
//	deleted and reset to new values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				string
//
//		A string of numeric digits used to populate the
//		integer part of the numeric value contained in the
//		current instance of NumberStrKernel
//
//	fractionalDigits			string
//
//		A string of numeric digits used to populate the
//		fractional part of the numeric value contained in
//		the current instance of NumberStrKernel
//
//	numberSign					NumericSignValueType
//
//		The Number Sign for the new numeric value is
//		specified by means of a	NumericSignValueType
//		enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = Ignored. Numeric Value
//									will be set to positive.
//
//			NumSignVal.Negative() = Valid Value
//
//			NumSignVal.Zero()     = Zero numeric values are
//									automatically assigned
//									NumSignVal.Zero().
//
//			NumSignVal.Positive() = Default value for
//									non-zero numeric values.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign will be automatically
//		assigned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetStringDigits(
	integerDigits string,
	fractionalDigits string,
	numberSign NumericSignValueType,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumStrKernel := NumberStrKernel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetStringDigits()",
		"")

	if err != nil {
		return err
	}

	return new(numberStrKernelNanobot).setWithRunes(
		&newNumStrKernel,
		[]rune(integerDigits),
		[]rune(fractionalDigits),
		numberSign,
		ePrefix)
}

//	SetUnsignedIntValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the unsigned
//	integer	numeric value passed as input parameter,
//	'unsignedIntValue'.
//
//	Since 'unsignedIntValue' is an empty interface, the
//	unsigned integer value may be any one of the
//	following types:
//
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//	Unsigned integer numeric values always default to
//	a positive number sign (+). With this method, the
//	user has the option to specify the number sign
//	(positive or negative) assigned to the converted
//	numeric value using input parameter, 'numberSign'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of NumberStrKernel, all
//	pre-existing data will be deleted and overwritten
//	with new data generated from the signed numeric
//	value passed as input parameter, 'unsignedIntValue'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	unsignedIntValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface must match one of the following
//		types:
//
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = -2 - Set to positive for
//											unsigned integer
//											values
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
//
//		Unsigned integer values are by default converted as
//		positive numeric values. If this parameter is set
//		to NumSignVal.Negative(), the converted numeric value
//		will be classified as a negative value.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign for the converted numeric
//		value will be set to 'positive'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetUnsignedIntValue(
	unsignedIntValue interface{},
	numberSign NumericSignValueType,
	errorPrefix interface{}) error {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetUnsignedIntValue()",
		"")

	if err != nil {
		return err
	}

	err = new(numberStrKernelMolecule).
		convertIntNumberToKernel(
			numStrKernel,
			unsignedIntValue,
			ePrefix)

	if err != nil {
		return err
	}

	if numberSign == NumSignVal.Negative() {

		numStrKernel.numberSign = numberSign
	}

	return err
}

// String
//
// Returns a formatted number string using the numeric
// value and the Default Number String Format
// Specification configured for the current instance of
// NumberStrKernel.
//
// To set the Default Number String Format Specification,
// use the following method:
//
//	NumberStrKernel.SetDefaultNumberStrFormatSpec()
//
// If an error is encountered, the error message is
// included in the string returned by this method.
//
// This method implements the Stringer Interface.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	No rounding will be performed on the numeric value
//		formatted as number string. If rounding is
//		required in conjunction with the Default Format
//		Specification, use the following method:
//
//			NumberStrKernel.FmtNumStrDefaultFormat()
//
//	(2)	This method will NOT delete or modify data values
//		contained in the current instance of
//		NumberStrKernel. However, if the Default
//		NumberStrKernel Number String Format
//		Specification is invalid,
//		'NumberStrKernel.numStrFormatSpec' will be set to
//		the default US (United States) Signed Number
//		String Format Specification.
//
//		All other NumberStrKernel data values will remain
//		unchanged.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, a formatted
//		Number String will be returned.
//
//		If errors are encountered, an error message will
//		be formatted in the returned string.
func (numStrKernel NumberStrKernel) String() string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var errOut string

	funcName := "NumberStrKernel.String()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		funcName,
		"")

	if err != nil {
		errOut = fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			funcName,
			err.Error())

		return errOut
	}

	var tempNumStrFormatSpec NumStrFormatSpec

	if new(numStrFmtSpecNanobot).isNOP(
		&numStrKernel.numStrFormatSpec) {

		// The current Number String Format
		// Specification is invalid.
		tempNumStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumFmtUS(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy(
				"tempNumStrFormatSpec<-"))

		if err != nil {

			errOut = fmt.Sprintf("%v\n"+
				"Error Message:\n"+
				"%v",
				funcName,
				err.Error())

			return errOut
		}

	} else {

		// The current Number String Format
		// Specification is valid.

		err = tempNumStrFormatSpec.CopyIn(
			&numStrKernel.numStrFormatSpec,
			ePrefix)

		if err != nil {

			errOut = fmt.Sprintf("%v\n"+
				"Error Message:\n"+
				"%v",
				funcName,
				err.Error())

			return errOut
		}

	}

	var roundingSpec NumStrRoundingSpec

	roundingSpec,
		err = new(NumStrRoundingSpec).
		NewRoundingSpec(
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy("roundingSpec"))

	var numStr string

	numStr,
		err = new(numberStrKernelMolecule).formatNumStr(
		&numStrKernel,
		tempNumStrFormatSpec,
		roundingSpec,
		ePrefix.XCpy("numStrKernel"))

	if err != nil {

		errOut = fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			funcName,
			err.Error())

		return errOut
	}

	return numStr
}
