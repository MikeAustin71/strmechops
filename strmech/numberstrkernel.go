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

	numericValueType NumericValueType
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

// CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumberStrKernel ('incomingNumStrKernel') to the	data
//	fields of the current NumberStrKernel instance
//	('numStrKernel').
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// All the data fields in current NumberStrKernel instance
// ('numStrKernel') will be modified and overwritten.
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
//		If parameter 'incomingNumStrKernel' is determined to
//		be invalid, an error will be returned.
//
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
		new(numberStrKernelNanobot).copyIn(
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

	deepCopyNumStrKernel,
		err = new(numberStrKernelNanobot).
		copyOut(
			numStrKernel,
			ePrefix.XCpy(
				"deepCopyNumStrKernel<-"+
					"numStrKernel"))

	return deepCopyNumStrKernel, err
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

	numStrKernel.fractionalDigits.EmptyCharsArray()

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

	numStrKernel.integerDigits.EmptyCharsArray()

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

	return new(numberStrKernelElectron).convertKernelToBigInt(
		numStrKernel,
		roundingType,
		ePrefix)
}

func (numStrKernel *NumberStrKernel) GetIntegerNum(
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
			"GetIntegerNum()",
		"")

	if err != nil {
		return intValue, err
	}

	err = new(numberStrKernelMolecule).convertKernelToIntNum(
		numStrKernel,
		&intValue,
		roundingType,
		ePrefix)

	return intValue, err
}

//	GetIntNum
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as an integer value of type int.
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

	return new(numberStrKernelMolecule).convertKernelToInt(
		numStrKernel,
		roundingType,
		ePrefix)
}

//	GetInt32Num
//
//	Returns the numeric value of the current NumberStrKernel
//	instance as an integer value of type int32.
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
//					the right of the decimal point32.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point32.
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

	intValue := int32(-1)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetInt32Num()",
		"")

	if err != nil {
		return intValue, err
	}

	return new(numberStrKernelMolecule).convertKernelToInt32(
		numStrKernel,
		roundingType,
		ePrefix)
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

// GetIntegerRuneArray - A deep copy of the integer digits rune
// array encapsulated by the current instance of NumberStrKernel.
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
//	   - This method returns a rune array consisting of the integer
//	     digits portion of the numeric value represented by the
//	     current instance of NumberStrKernel. This source rune
//	     array is identified as internal member variable
//	     'NumberStrKernel.integerDigits'.
func (numStrKernel *NumberStrKernel) GetIntegerRuneArray() []rune {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetRuneArray()
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

// GetFractionalString - Converts the fractional rune array
// contained in the current instance of NumberStrKernel to a string
// and returns it to the calling function.
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
//	   - This method returns a string consisting of the fractional
//	     digits portion of the numeric value represented by the
//	     current instance of NumberStrKernel. This source of these
//	     fractional digits is identified as internal member
//	     variable 'NumberStrKernel.fractionalDigits'.
func (numStrKernel *NumberStrKernel) GetFractionalString() string {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.fractionalDigits.GetCharacterString()
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

// GetNumberSignInt - Returns the Number Sign associated with the
// current NumberStrKernel instance as an integer value.
//
// If the current number sign value is invalid, an error will be
// returned.
//
// Valid return values for Number Sign are:
//
//	-1 Minus One
//	 0 Zero
//	 1 Plus 1
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
//		int
//	    - If this method completes successfully, this parameter will
//	      return an integer value specifying the number sign
//	      associated with the numeric value represented by the
//	      current instance of NumberStrKernel. Possible valid return
//	      values are listed as follows:
//	         -1 (Minus One)
//	          0 (Zero)
//	          1 (Plus One)
//
//
//	 error
//	    - If this method completes successfully and no errors are
//	      encountered this return value is set to 'nil'. Otherwise,
//	      if errors are encountered, this return value will contain
//	      an appropriate error message.
//
//	      If an error message is returned, the text value of input
//	      parameter 'errorPrefix' will be inserted or prefixed at
//	      the beginning of the error message.
func (numStrKernel *NumberStrKernel) GetNumberSignInt(
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
			"GetNumberSignInt()",
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
			"GetNumberSignInt()",
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

// IsNonZeroValue
//
// If this method returns 'true', it means that
// the numeric value of the current NumberStrKernel
// instance is non-zero. A non-zero numeric value
// signals that the numeric value is less than or
// greater than zero (0).
//
// If this method returns 'false' it means that the
// numeric value of current NumberStrKernel instance
// is zero ('0').
func (numStrKernel *NumberStrKernel) IsNonZeroValue() bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var isNonZeroValue bool

	isNonZeroValue,
		_ = new(numberStrKernelElectron).isNonZeroValue(
		numStrKernel,
		nil)

	return isNonZeroValue
}

//	NewFromBigFloatValue
//
//	Creates a new instance of NumberStrKernel based on a
//	big float numeric value (*big.Float) passed as an
//	input parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigFloatNum					*big.Float
//
//		This floating point numeric value will be
//		used to create and return a new instance of
//		NumberStrKernel.
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
//		'bigFloatNum'.
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
func (numStrKernel *NumberStrKernel) NewFromBigFloatValue(
	bigFloatNum *big.Float,
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
			"NewFromBigFloatValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	numberSign := NumSignVal.None()

	err = new(numberStrKernelNanobot).setWithNumber(
		&newNumStrKernel,
		bigFloatNum,
		numberSign,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromBigIntValue
//
//	Creates a new instance of NumberStrKernel based on a
//	big integer value (*big.Int) passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigIntNum					*big.Int
//
//		The numeric value which will be used to create and
//		return a new instance of NumberStrKernel.
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
//		'bigIntNum'.
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
func (numStrKernel *NumberStrKernel) NewFromBigIntValue(
	bigIntNum *big.Int,
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

	numberSign := NumSignVal.None()

	err = new(numberStrKernelNanobot).setWithNumber(
		&newNumStrKernel,
		bigIntNum,
		numberSign,
		ePrefix)

	return newNumStrKernel, err
}

//	NewFromFloatValue
//
//	Creates a new instance of NumberStrKernel based on a
//	floating point numeric value passed as an empty
//	interface.
//
//	The floating point numeric value may be any one of
//	the following types:
//
//		float32
//		float64
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	floatingPointValue			interface{}
//
//		Numeric values passed by means of this empty
//		interface must match one of the following
//		types:
//
//			float32
//			float64
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

	numberSign := NumSignVal.None()

	err = new(numberStrKernelNanobot).setWithNumber(
		&newNumStrKernel,
		floatingPointValue,
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
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedIntValue				interface{}
//
//		Numeric values passed by means of this empty
//		interface must match one of the following
//		types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
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

	numberSign := NumSignVal.None()

	err = new(numberStrKernelNanobot).setWithNumber(
		&newNumStrKernel,
		signedIntValue,
		numberSign,
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

	err = new(numberStrKernelNanobot).setWithNumber(
		&newNumStrKernel,
		unsignedIntValue,
		numberSign,
		ePrefix)

	return newNumStrKernel, err
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
//	This method may alter the numeric value of contained
//	withing the current instance of NumberStrKernel.
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

//	SetBigFloatValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the big floating
//	point numeric value passed as an instance of type
//	*big.Float.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of NumberStrKernel, all
//	pre-existing data will be deleted and overwritten
//	with new data generated from the numeric value
//	passed as input parameter, 'bigFloatNum'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigFloatNum					*big.Float
//
//		This floating point numeric value will be used
//		to configure and populate the current instance
//		of NumberStrKernel.
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
func (numStrKernel *NumberStrKernel) SetBigFloatValue(
	bigFloatNum *big.Float,
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
			"SetBigFloatValue()",
		"")

	if err != nil {
		return err
	}

	numberSign := NumSignVal.None()

	return new(numberStrKernelNanobot).setWithNumber(
		numStrKernel,
		bigFloatNum,
		numberSign,
		ePrefix)
}

//	SetBigIntValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the big integer
//	numeric value passed as an instance of type *big.Int.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of NumberStrKernel, all
//	pre-existing data will be deleted and overwritten
//	with new data generated from the numeric value
//	passed as input parameter, 'bigIntNum'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigIntNum					*big.Int
//
//		This big integer numeric value will be used
//		to configure and populate the current instance
//		of NumberStrKernel.
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
func (numStrKernel *NumberStrKernel) SetBigIntValue(
	bigIntNum *big.Int,
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
			"SetBigIntValue()",
		"")

	if err != nil {
		return err
	}

	numberSign := NumSignVal.None()

	return new(numberStrKernelNanobot).setWithNumber(
		numStrKernel,
		bigIntNum,
		numberSign,
		ePrefix)
}

//	SetFloatValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the floating
//	point numeric value passed as an empty interface.
//
//	The floating point numeric value may be any one of
//	the following types:
//
//		float32
//		float64
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
//		interface must match one of the following
//		types:
//
//			float32
//			float64
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

	numberSign := NumSignVal.None()

	return new(numberStrKernelNanobot).setWithNumber(
		numStrKernel,
		floatingPointValue,
		numberSign,
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

//	SetSignedIntValue
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using the signed integer
//	numeric value passed as input parameter,
//	'signedIntValue'.
//
//	'signedIntValue' is an empty interface which means
//	that the signed integer value may be any one of the
//	following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
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
//	signedIntValue				interface{}
//
//		Numeric values passed by means of this empty
//		interface must match one of the following
//		types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
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
func (numStrKernel *NumberStrKernel) SetSignedIntValue(
	signedIntValue interface{},
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
			"SetSignedIntValue()",
		"")

	if err != nil {
		return err
	}

	numberSign := NumSignVal.None()

	err = new(numberStrKernelNanobot).setWithNumber(
		numStrKernel,
		signedIntValue,
		numberSign,
		ePrefix)

	return err
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

	return new(numberStrKernelNanobot).setWithNumber(
		numStrKernel,
		unsignedIntValue,
		numberSign,
		ePrefix)
}

// String - Returns a formatted text string detailing all the
// internal member variable names and their corresponding values
// for the current instance of NumberStrKernel.
//
// If an error is encountered, the error message is included in the
// string returned by this method.
//
// This method implements the Stringer Interface.
func (numStrKernel NumberStrKernel) String() string {

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
		"CharSearchNegativeNumberResultsDto."+
			"String()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"CharSearchNegativeNumberResultsDto.String()",
			err.Error())

		return errOut
	}

	strBuilder := strings.Builder{}

	err = new(numberStrKernelNanobot).
		getParameterTextListing(
			&strBuilder,
			false,
			&numStrKernel,
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
