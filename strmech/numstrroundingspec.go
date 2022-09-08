package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrRoundingSpec
//
// This data transfer object contains all the parameters
// required to configure a rounding algorithm for a
// floating point number string.
type NumStrRoundingSpec struct {
	roundingType NumberRoundingType
	// This enumeration parameter is used to specify the type
	// of rounding algorithm that will be applied for the
	// rounding of fractional digits in a number string.
	//
	// Possible values are listed as follows:
	//	NumRoundType.None()
	//  NumRoundType.NoRounding()
	//	NumRoundType.HalfUpWithNegNums()
	//	NumRoundType.HalfDownWithNegNums()
	//	NumRoundType.HalfAwayFromZero()
	//	NumRoundType.HalfTowardsZero()
	//	NumRoundType.HalfToEven()
	//	NumRoundType.HalfToOdd()
	//	NumRoundType.Randomly()
	//	NumRoundType.Floor()
	//	NumRoundType.Ceiling()
	//	NumRoundType.Truncate()
	//
	// NoRounding					1
	//
	//	Signals that no rounding operation will be performed
	//	on fractional digits contained in a number string.
	//	The fractional digits will therefore remain unchanged.
	//
	// HalfUpWithNegNums			2
	//
	//	Half Round Up Including Negative Numbers. This method
	//	is intuitive but may produce unexpected results when
	//	applied to negative numbers.
	//
	//	'HalfUpWithNegNums' rounds .5 up.
	//
	//		Examples of 'HalfUpWithNegNums'
	//		7.6 rounds up to 8
	//		7.5 rounds up to 8
	//		7.4 rounds down to 7
	//		-7.4 rounds up to -7
	//		-7.5 rounds up to -7
	//		-7.6 rounds down to -8
	//
	// HalfDownWithNegNums          3
	//
	//	Half Round Down Including Negative Numbers. This method
	//	is also considered intuitive but may produce unexpected
	//	results when applied to negative numbers.
	//
	//	'HalfDownWithNegNums' rounds .5 down.
	//
	//		Examples of HalfDownWithNegNums
	//
	//		7.6 rounds up to 8
	//		7.5 rounds down to 7
	//		7.4 rounds down to 7
	//		-7.4 rounds up to -7
	//		-7.5 rounds down to -8
	//		-7.6 rounds down to -8
	//
	// HalfAwayFromZero				4
	//
	//	Round Half Away From Zero. This rounding method is treated
	//	as the default and this value is returned by method:
	//	NumberRoundingType(0).XGetDefaultRoundingType()
	//
	//	The 'HalfAwayFromZero' method rounds .5 further away from zero.
	//	It provides clear and consistent behavior when dealing with
	//	negative numbers.
	//
	//		Examples of HalfAwayFromZero
	//
	//		7.6 rounds away to 8
	//		7.5 rounds away to 8
	//		7.4 rounds to 7
	//		-7.4 rounds to -7
	//		-7.5 rounds away to -8
	//		-7.6 rounds away to -8
	//
	// HalfTowardsZero				5
	//
	//	Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
	//	closer to zero. It provides clear and consistent behavior
	//	when dealing with negative numbers.
	//
	//		Examples of HalfTowardsZero
	//
	//		7.6 rounds away to 8
	//		7.5 rounds to 7
	//		7.4 rounds to 7
	//		-7.4 rounds to -7
	//		-7.5 rounds to -7
	//		-7.6 rounds away to -8
	//
	// HalfToEven					6
	//
	//	Round Half To Even Numbers. 'HalfToEven' is also called
	//	Banker's Rounding. This method rounds 0.5 to the nearest
	//	even digit.
	//
	//		Examples of HalfToEven
	//
	//		7.5 rounds up to 8 (because 8 is an even number)
	//		but 6.5 rounds down to 6 (because 6 is an even number)
	//
	//		HalfToEven only applies to 0.5. Other numbers (not ending
	//		in 0.5) round to nearest as usual, so:
	//
	//		7.6 rounds up to 8
	//		7.5 rounds up to 8 (because 8 is an even number)
	//		7.4 rounds down to 7
	//		6.6 rounds up to 7
	//		6.5 rounds down to 6 (because 6 is an even number)
	//		6.4 rounds down to 6
	//
	// HalfToOdd					7
	//
	//	Round Half to Odd Numbers. Similar to 'HalfToEven', but
	//	in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
	//
	//		Examples of HalfToOdd
	//
	//		HalfToOdd only applies to 0.5. Other numbers (not ending
	//		in 0.5) round to nearest as usual.
	//
	//		7.5 rounds down to 7 (because 7 is an odd number)
	//
	//		6.5 rounds up to 7 (because 7 is an odd number)
	//
	//		7.6 rounds up to 8
	//		7.5 rounds down to 7 (because 7 is an odd number)
	//		7.4 rounds down to 7
	//		6.6 rounds up to 7
	//		6.5 rounds up to 7 (because 7 is an odd number)
	//		6.4 rounds down to 6
	//
	// Randomly						8
	//
	//	Round Half Randomly. Uses a Random Number Generator to choose
	//	between rounding 0.5 up or down.
	//
	//	All numbers other than 0.5 round to the nearest as usual.
	//
	// Floor						9
	//
	//	Yields the nearest integer down. Floor does not apply any
	//	special treatment to 0.5.
	//
	//	Floor Function: The greatest integer that is less than or
	//	equal to x
	//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
	//
	//	In mathematics and computer science, the floor function is
	//	the function that takes as input a real number x, and gives
	//	as output the greatest integer less than or equal to x,
	//	denoted floor(x) or ⌊x⌋.
	//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
	//
	//		Examples of Floor
	//
	//		Number     Floor
	//		2           2
	//		2.4         2
	//		2.9         2
	//		-2.5        -3
	//		-2.7        -3
	//		-2          -2
	//
	// Ceiling						10
	//
	//	Yields the nearest integer up. Ceiling does not apply any
	//	special treatment to 0.5.
	//
	//	Ceiling Function: The least integer that is greater than or
	//	equal to x.
	//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
	//
	//	The ceiling function maps x to the least integer greater than
	//	or equal to x, denoted ceil(x) or ⌈x⌉.[1]
	//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
	//
	//		Examples of Ceiling
	//
	//		Number    Ceiling
	//		2           2
	//		2.4         3
	//		2.9         3
	//		-2.5        -2
	//		-2.7        -2
	//		-2          -2
	//
	// Truncate						11
	//
	//	Apply NO Rounding whatsoever. The Round From Digit is dropped
	//	or deleted. The Round To Digit is NEVER changed.
	//
	//		Examples of Truncate
	//
	//		Example-1
	//		Number: 23.14567
	//		Objective: Round to two decimal places to
	//		the right of the decimal point.
	//		Rounding Method: Truncate
	//		Round To Digit:   4
	//		Round From Digit: 5
	//		Rounded Number:   23.14 - The Round From Digit is dropped.
	//
	//		Example-2
	//		Number: -23.14567
	//		Objective: Round to two decimal places to
	//		the right of the decimal point.
	//		Rounding Method: Truncate
	//		Round To Digit:   4
	//		Round From Digit: 5
	//		Rounded Number:  -23.14 - The Round From Digit is dropped.

	roundToFractionalDigits int
	// When set to a positive integer value, this
	// parameter controls the number of digits to
	// the right of the decimal separator (a.k.a.
	// decimal point) which will remain after
	// completion of the number rounding operation.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrRoundingSpec ('incomingNStrRoundingSpec')
// to the data fields of the current NumStrRoundingSpec
// instance ('nStrRoundingSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrRoundingSpec
// instance ('nStrRoundingSpec') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNStrRoundingSpec	*NumStrRoundingSpec
//		A pointer to an instance of NumStrRoundingSpec.
//		This method will NOT change the values of internal member
//		variables contained in this instance.
//
//		All data values in this NumStrRoundingSpec instance
//		will be copied to current NumStrRoundingSpec
//		instance ('nStrRoundingSpec').
//
//		If parameter 'incomingNStrRoundingSpec' is determined to
//		be invalid, an error will be returned.
//
//	errorPrefix					interface{}
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
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
func (nStrRoundingSpec *NumStrRoundingSpec) CopyIn(
	incomingNStrRoundingSpec *NumStrRoundingSpec,
	errorPrefix interface{}) error {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrRoundingSpecNanobot).
		copyNStrRoundingSpec(
			nStrRoundingSpec,
			incomingNStrRoundingSpec,
			ePrefix.XCpy(
				"nStrRoundingSpec<-"+
					"incomingNStrRoundingSpec"))
}

// CopyOut - Returns a deep copy of the current
// NumStrRoundingSpec instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyNStrRoundingSpec   NumStrRoundingSpe
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of
//	     the current NumStrRoundingSpec instance.
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
func (nStrRoundingSpec *NumStrRoundingSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNStrRoundingSpec NumStrRoundingSpec,
	err error) {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNStrRoundingSpec, err
	}

	err = new(numStrRoundingSpecNanobot).
		copyNStrRoundingSpec(
			&deepCopyNStrRoundingSpec,
			nStrRoundingSpec,
			ePrefix.XCpy(
				"deepCopyNStrRoundingSpec<-"+
					"nStrRoundingSpec"))

	return deepCopyNStrRoundingSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrRoundingSpec to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NumStrRoundingSpec.
//
// ------------------------------------------------------------------------
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
func (nStrRoundingSpec *NumStrRoundingSpec) Empty() {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	new(numStrRoundingSpecAtom).empty(
		nStrRoundingSpec)

	nStrRoundingSpec.lock.Unlock()

	nStrRoundingSpec.lock = nil

	return
}

// Equal - Receives a pointer to another instance of
// NumStrRoundingSpec and proceeds to compare its internal
// member variables to those of the current
// NumStrRoundingSpec instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'.
// Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNStrRoundingSpec   *NumStrRoundingSpec
//	   - A pointer to an external instance of
//	     NumStrRoundingSpec. The internal member variable
//	     data values in this instance will be compared to those
//	     in the current instance of NumStrRoundingSpec. The
//	     results of this comparison will be returned to the
//	     calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the internal member variable data values contained in
//	     input parameter 'incomingNStrRoundingSpec' are equivalent
//	     in all respects to those contained in the current
//	     instance of 'incomingNStrRoundingSpec', this return value
//	     will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (nStrRoundingSpec *NumStrRoundingSpec) Equal(
	incomingNStrRoundingSpec *NumStrRoundingSpec) bool {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	return new(numStrRoundingSpecAtom).equal(
		nStrRoundingSpec,
		incomingNStrRoundingSpec)
}

// GetRoundingType - Returns the value of member variable
// 'NumStrRoundingSpec.roundingType' for the current instance
// of NumStrRoundingSpec.
//
// This member variable is an instance of NumberRoundingType.
// NumberRoundingType is an enumeration specifying the
// rounding algorithm to be applied in the fractional digit
// rounding operation. Valid values are listed as follows:
//
//	NumRoundType.None(),
//	NumRoundType.HalfUpWithNegNums(),
//	NumRoundType.HalfDownWithNegNums(),
//	NumRoundType.HalfAwayFromZero(),
//	NumRoundType.HalfTowardsZero(),
//	NumRoundType.HalfToEven(),
//	NumRoundType.HalfToOdd(),
//	NumRoundType.Randomly(),
//	NumRoundType.Floor(),
//	NumRoundType.Ceiling(),
//	NumRoundType.Truncate(),
func (nStrRoundingSpec *NumStrRoundingSpec) GetRoundingType() NumberRoundingType {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	return nStrRoundingSpec.roundingType
}

// GetRoundToFractionalDigits - Returns the value of member
// variable 'NumStrRoundingSpec.roundToFractionalDigits' for
// the current instance of NumStrRoundingSpec.
//
// When set to a positive integer value, this integer value
// controls the number of digits to the right of the decimal
// separator (a.k.a. decimal point) which will remain after
// completion of the number rounding operation.
func (nStrRoundingSpec *NumStrRoundingSpec) GetRoundToFractionalDigits() int {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	return nStrRoundingSpec.roundToFractionalDigits
}

// NewRoundingSpec - Creates and returns a new instance of
// NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	roundingType				NumberRoundingType
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//		 NumRoundType.None(),
//		 NumRoundType.HalfUpWithNegNums(),
//		 NumRoundType.HalfDownWithNegNums(),
//		 NumRoundType.HalfAwayFromZero(),
//		 NumRoundType.HalfTowardsZero(),
//		 NumRoundType.HalfToEven(),
//		 NumRoundType.HalfToOdd(),
//		 NumRoundType.Randomly(),
//		 NumRoundType.Floor(),
//		 NumRoundType.Ceiling(),
//		 NumRoundType.Truncate(),
//
//	roundToFractionalDigits		int
//		This parameter will populate the
//		'NumStrRoundingSpec.roundToFractionalDigits' member
//		variable data value contained in the returned instance
//		of NumStrRoundingSpec.
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	errorPrefix					interface{}
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrRoundingSpec		NumStrRoundingSpec
//		If this method completes successfully, this parameter
//		will return a fully populated instance of
//		NumStrRoundingSpec.
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrRoundingSpec *NumStrRoundingSpec) NewRoundingSpec(
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	newNumStrRoundingSpec NumStrRoundingSpec,
	err error) {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"NewRoundingSpec()",
		"")

	if err != nil {
		return newNumStrRoundingSpec, err
	}

	err = new(numStrRoundingSpecNanobot).
		setNStrNStrRoundingSpec(
			&newNumStrRoundingSpec,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"newNumStrRoundingSpec<-"))

	return newNumStrRoundingSpec, err
}

// SetRoundingSpec - Deletes and overwrites all member variable
// data values in the current instance of NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 roundingType               NumberRoundingType
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundingType' member variable
//	      data value contained in the current instance of
//	      NumStrRoundingSpec.
//
//	      NumberRoundingType is an enumeration specifying the
//	      rounding algorithm to be applied in the fractional digit
//	      rounding operation. Valid values are listed as follows:
//
//	       NumRoundType.None(),
//	       NumRoundType.HalfUpWithNegNums(),
//	       NumRoundType.HalfDownWithNegNums(),
//	       NumRoundType.HalfAwayFromZero(),
//	       NumRoundType.HalfTowardsZero(),
//	       NumRoundType.HalfToEven(),
//	       NumRoundType.HalfToOdd(),
//	       NumRoundType.Randomly(),
//	       NumRoundType.Floor(),
//	       NumRoundType.Ceiling(),
//	       NumRoundType.Truncate(),
//
//	 roundToFractionalDigits    int
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundToFractionalDigits' member
//	      variable data value contained in the current instance
//	      of NumStrRoundingSpec.
//
//	      When set to a positive integer value, this parameter
//	      controls the number of digits to the right of the
//	      decimal separator (a.k.a. decimal point) which will
//	      remain after completion of the number rounding
//	      operation.
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpec *NumStrRoundingSpec) SetRoundingSpec(
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	err error) {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"SetRoundingSpec()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrRoundingSpecNanobot).
		setNStrNStrRoundingSpec(
			nStrRoundingSpec,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"nStrRoundingSpec<-"))

	return err
}

// SetRoundingType - Deletes and resets the
// 'NumStrRoundingSpec.roundingType' member variable
// data value contained in the current instance of
// NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the pre-existing data
// value for the 'NumStrRoundingSpec.roundingType' member
// variable contained in the current instance of
// NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 roundingType               NumberRoundingType
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundingType' member variable
//	      data value contained in the current instance of
//	      NumStrRoundingSpec.
//
//	      NumberRoundingType is an enumeration specifying the
//	      rounding algorithm to be applied in the fractional digit
//	      rounding operation. Valid values are listed as follows:
//
//	       NumRoundType.None(),
//	       NumRoundType.HalfUpWithNegNums(),
//	       NumRoundType.HalfDownWithNegNums(),
//	       NumRoundType.HalfAwayFromZero(),
//	       NumRoundType.HalfTowardsZero(),
//	       NumRoundType.HalfToEven(),
//	       NumRoundType.HalfToOdd(),
//	       NumRoundType.Randomly(),
//	       NumRoundType.Floor(),
//	       NumRoundType.Ceiling(),
//	       NumRoundType.Truncate(),
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpec *NumStrRoundingSpec) SetRoundingType(
	roundingType NumberRoundingType,
	errorPrefix interface{}) (
	err error) {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"SetRoundingType()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrRoundingSpecAtom).
		setRoundingType(
			nStrRoundingSpec,
			roundingType,
			ePrefix.XCpy(
				"nStrRoundingSpec<-"))

	return err
}

// SetRoundToFractionalDigits - Deletes and resets the
// 'NumStrRoundingSpec.roundToFractionalDigits' member
// variable data value contained in the current instance
// of NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the pre-existing data
// value for the 'NumStrRoundingSpec.roundToFractionalDigits'
// member variable contained in the current instance of
// NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 roundToFractionalDigits    int
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundToFractionalDigits' member
//	      variable data value contained in the input parameter,
//	      'nStrRoundingSpec'. When set to a positive integer value,
//	      this parameter controls the number of digits to the right
//	      of the decimal separator (a.k.a. decimal point) which
//	      will remain after completion of the number rounding
//	      operation.
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpec *NumStrRoundingSpec) SetRoundToFractionalDigits(
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	err error) {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrRoundingSpec."+
			"SetRoundToFractionalDigits()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrRoundingSpecAtom).
		setRoundToFractionalDigits(
			nStrRoundingSpec,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"nStrRoundingSpec<-"))

	return err
}

// numStrRoundingSpecNanobot - This type provides
// helper methods for NumStrRoundingSpec
type numStrRoundingSpecNanobot struct {
	lock *sync.Mutex
}

// copyNStrRoundingSpec - Copies all data from input parameter
// 'sourceNStrRoundingSpec' to input parameter
// 'destinationNStrRoundingSpec'. Both instances are of type
// NumStrRoundingSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationNStrRoundingSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceNStrRoundingSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNStrRoundingSpec  *NumStrRoundingSpec
//	   - A pointer to a NumStrRoundingSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNStrRoundingSpec'.
//
//	     'destinationNStrRoundingSpec' is the destination for this
//	     copy operation.
//
//
//	sourceNStrRoundingSpec       *NumStrRoundingSpec
//	   - A pointer to another NumStrRoundingSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationNStrRoundingSpec'.
//
//	     'sourceNStrRoundingSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceNStrRoundingSpec'.
//
//
//	errPrefDto                     *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpecNanobot *numStrRoundingSpecNanobot) copyNStrRoundingSpec(
	destinationNStrRoundingSpec *NumStrRoundingSpec,
	sourceNStrRoundingSpec *NumStrRoundingSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrRoundingSpecNanobot.lock == nil {
		nStrRoundingSpecNanobot.lock = new(sync.Mutex)
	}

	nStrRoundingSpecNanobot.lock.Lock()

	defer nStrRoundingSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrRoundingSpecNanobot."+
			"copyNStrRoundingSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNStrRoundingSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNStrRoundingSpec' is invalid!\n"+
			"'destinationNStrRoundingSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNStrRoundingSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNStrRoundingSpec' is invalid!\n"+
			"'sourceNStrRoundingSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrRoundingSpecAtom).empty(
		destinationNStrRoundingSpec)

	destinationNStrRoundingSpec.roundingType =
		sourceNStrRoundingSpec.roundingType

	destinationNStrRoundingSpec.roundToFractionalDigits =
		sourceNStrRoundingSpec.roundToFractionalDigits

	return err
}

// setNStrNStrRoundingSpec - Deletes and resets all member
// variable data values contained in the instance of
// NumStrRoundingSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		nStrRoundingSpec           *NumStrRoundingSpec
//		    - A pointer to an instance of NumStrRoundingSpec.
//		      All the member variable data values in this instance
//		      will be deleted and reset according to the data
//		      extracted from the following input parameters.
//
//	 roundingType               NumberRoundingType
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundingType' member variable
//	      data value contained in the input parameter,
//	      'nStrRoundingSpec'.
//
//	      NumberRoundingType is an enumeration specifying the
//	      rounding algorithm to be applied in the fractional digit
//	      rounding operation. Valid values are listed as follows:
//
//	       NumRoundType.None(),
//	       NumRoundType.HalfUpWithNegNums(),
//	       NumRoundType.HalfDownWithNegNums(),
//	       NumRoundType.HalfAwayFromZero(),
//	       NumRoundType.HalfTowardsZero(),
//	       NumRoundType.HalfToEven(),
//	       NumRoundType.HalfToOdd(),
//	       NumRoundType.Randomly(),
//	       NumRoundType.Floor(),
//	       NumRoundType.Ceiling(),
//	       NumRoundType.Truncate(),
//
//	 roundToFractionalDigits    int
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundToFractionalDigits' member
//	      variable data value contained in the input parameter,
//	      'nStrRoundingSpec'. When set to a positive integer value,
//	      this parameter controls the number of digits to the right
//	      of the decimal separator (a.k.a. decimal point) which
//	      will remain after completion of the number rounding
//	      operation.
//
//		 errPrefDto                 *ePref.ErrPrefixDto
//		    - This object encapsulates an error prefix string which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods listed
//		      as a function chain.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      Type ErrPrefixDto is included in the 'errpref' software
//		      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpecNanobot *numStrRoundingSpecNanobot) setNStrNStrRoundingSpec(
	nStrRoundingSpec *NumStrRoundingSpec,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrRoundingSpecNanobot.lock == nil {
		nStrRoundingSpecNanobot.lock = new(sync.Mutex)
	}

	nStrRoundingSpecNanobot.lock.Lock()

	defer nStrRoundingSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrRoundingSpecNanobot."+
			"setNStrNStrRoundingSpec()",
		"")

	if err != nil {
		return err
	}

	if nStrRoundingSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrRoundingSpec' is invalid!\n"+
			"'nStrRoundingSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	nStrRoundingSpecAtom := numStrRoundingSpecAtom{}

	nStrRoundingSpecAtom.empty(
		nStrRoundingSpec)

	err = nStrRoundingSpecAtom.setRoundingType(
		nStrRoundingSpec,
		roundingType,
		ePrefix.XCpy(
			"nStrRoundingSpec<-roundingType"))

	if err != nil {
		return err
	}

	err = nStrRoundingSpecAtom.setRoundToFractionalDigits(
		nStrRoundingSpec,
		roundToFractionalDigits,
		ePrefix.XCpy(
			"nStrRoundingSpec<-"+
				"roundToFractionalDigits"))

	if err != nil {
		return err
	}

	return err
}

// numStrRoundingSpecAtom - This type provides
// helper methods for NumStrRoundingSpec
type numStrRoundingSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrRoundingSpec and proceeds to reset the
// data values for all member variables to their
// initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrRoundingSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	nStrRoundingSpec           *NumStrRoundingSpec
//	   - A pointer to an instance of NumStrRoundingSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrRoundingSpecAtom *numStrRoundingSpecAtom) empty(
	nStrRoundingSpec *NumStrRoundingSpec) {

	if nStrRoundingSpecAtom.lock == nil {
		nStrRoundingSpecAtom.lock = new(sync.Mutex)
	}

	nStrRoundingSpecAtom.lock.Lock()

	defer nStrRoundingSpecAtom.lock.Unlock()

	if nStrRoundingSpec == nil {
		return
	}

	nStrRoundingSpec.roundingType = NumRoundType.None()

	nStrRoundingSpec.roundToFractionalDigits = -1
}

// equal - Receives a pointer to two instances of
// NumStrRoundingSpec and proceeds to compare their
// member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'. Otherwise,
// this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrRoundingSpec1    *NumStrRoundingSpec
//	   - An instance of NumStrRoundingSpec.
//	     Internal member variables from 'nStrRoundingSpec1'
//	     will be compared to those of 'nStrRoundingSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrRoundingSpec2    *NumStrRoundingSpec
//	   - An instance of NumStrRoundingSpec.
//	     Internal member variables from 'nStrRoundingSpec2'
//	     will be compared to those of 'nStrRoundingSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrRoundingSpec1' and
//	     'nStrRoundingSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrRoundingSpecAtom *numStrRoundingSpecAtom) equal(
	nStrRoundingSpec1 *NumStrRoundingSpec,
	nStrRoundingSpec2 *NumStrRoundingSpec) bool {

	if nStrRoundingSpecAtom.lock == nil {
		nStrRoundingSpecAtom.lock = new(sync.Mutex)
	}

	nStrRoundingSpecAtom.lock.Lock()

	defer nStrRoundingSpecAtom.lock.Unlock()

	if nStrRoundingSpec1 == nil ||
		nStrRoundingSpec2 == nil {
		return false
	}

	if nStrRoundingSpec1.roundingType !=
		nStrRoundingSpec2.roundingType {

		return false
	}

	if nStrRoundingSpec1.roundToFractionalDigits !=
		nStrRoundingSpec2.roundToFractionalDigits {

		return false
	}

	return true
}

// setRoundingType - Deletes and resets the
// 'NumStrRoundingSpec.roundingType' member variable
// data value contained in the instance of
// NumStrRoundingSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		nStrRoundingSpec           *NumStrRoundingSpec
//		    - A pointer to an instance of NumStrRoundingSpec.
//		      The 'NumStrRoundingSpec.roundingType' member
//		      variable data value in this instance will be deleted
//		      and reset according to the data extracted from input
//		      parameter ,'roundingType' .
//
//	 roundingType               NumberRoundingType
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundingType' member variable
//	      data value contained in the input parameter,
//	      'nStrRoundingSpec'.
//
//	      NumberRoundingType is an enumeration specifying the
//	      rounding algorithm to be applied in the fractional digit
//	      rounding operation. Valid values are listed as follows:
//
//	       NumRoundType.None(),
//	       NumRoundType.HalfUpWithNegNums(),
//	       NumRoundType.HalfDownWithNegNums(),
//	       NumRoundType.HalfAwayFromZero(),
//	       NumRoundType.HalfTowardsZero(),
//	       NumRoundType.HalfToEven(),
//	       NumRoundType.HalfToOdd(),
//	       NumRoundType.Randomly(),
//	       NumRoundType.Floor(),
//	       NumRoundType.Ceiling(),
//	       NumRoundType.Truncate(),
//
//		 errPrefDto                 *ePref.ErrPrefixDto
//		    - This object encapsulates an error prefix string which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods listed
//		      as a function chain.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      Type ErrPrefixDto is included in the 'errpref' software
//		      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpecAtom *numStrRoundingSpecAtom) setRoundingType(
	nStrRoundingSpec *NumStrRoundingSpec,
	roundingType NumberRoundingType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrRoundingSpecAtom.lock == nil {
		nStrRoundingSpecAtom.lock = new(sync.Mutex)
	}

	nStrRoundingSpecAtom.lock.Lock()

	defer nStrRoundingSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrRoundingSpecAtom."+
			"setRoundingType()",
		"")

	if err != nil {
		return err
	}

	if nStrRoundingSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrRoundingSpec' is invalid!\n"+
			"'nStrRoundingSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if !roundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'roundingType' is invalid!\n"+
			"'nStrRoundingSpec' string value  = '%v'.\n"+
			"'nStrRoundingSpec' integer value = '%v'.\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

		return err
	}

	nStrRoundingSpec.roundingType = roundingType

	return err
}

// setRoundToFractionalDigits - Deletes and resets the
// 'NumStrRoundingSpec.roundToFractionalDigits' member
// variable data value contained in the instance of
// NumStrRoundingSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		nStrRoundingSpec           *NumStrRoundingSpec
//		    - A pointer to an instance of NumStrRoundingSpec.
//		      The 'NumStrRoundingSpec.roundToFractionalDigits' member
//		      variable data value in this instance will be deleted
//		      and reset according to the data extracted from input
//		      parameter ,'roundToFractionalDigits' .
//
//	 roundToFractionalDigits    int
//	    - This parameter will replace the current value of
//	      the 'NumStrRoundingSpec.roundToFractionalDigits' member
//	      variable data value contained in the input parameter,
//	      'nStrRoundingSpec'.
//
//	      When set to a positive integer value, this parameter
//	      controls the number of digits to the right of the
//	      decimal separator (a.k.a. decimal point) which will
//	      remain after completion of the number rounding
//	      operation.
//
//		 errPrefDto                 *ePref.ErrPrefixDto
//		    - This object encapsulates an error prefix string which is
//		      included in all returned error messages. Usually, it
//		      contains the name of the calling method or methods listed
//		      as a function chain.
//
//		      If no error prefix information is needed, set this
//		      parameter to 'nil'.
//
//		      Type ErrPrefixDto is included in the 'errpref' software
//		      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrRoundingSpecAtom *numStrRoundingSpecAtom) setRoundToFractionalDigits(
	nStrRoundingSpec *NumStrRoundingSpec,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrRoundingSpecAtom.lock == nil {
		nStrRoundingSpecAtom.lock = new(sync.Mutex)
	}

	nStrRoundingSpecAtom.lock.Lock()

	defer nStrRoundingSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrRoundingSpecAtom."+
			"setRoundToFractionalDigits()",
		"")

	if err != nil {
		return err
	}

	if nStrRoundingSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrRoundingSpec' is invalid!\n"+
			"'nStrRoundingSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if roundToFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'roundToFractionalDigits' is invalid!\n"+
			"'roundToFractionalDigits' has a value less than zero (0).\n"+
			"'roundToFractionalDigits' = '%v'.\n",
			ePrefix.String(),
			roundToFractionalDigits)

		return err
	}

	nStrRoundingSpec.roundToFractionalDigits = roundToFractionalDigits

	return err
}
