package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
)

//	GetDefaultNumStrFmtSpec
//
//	Returns a deep copy of the default Number String
//	Formatting specifications ('NumStrFormatSpec')
//	configured for the current instance of
//	NumberStrKernel.
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
//	NumStrFormatSpec
//
//		If this method completes successfully, a deep
//		copy of the default Number String Formatting
//		specifications configured for the current
//		instance of NumberStrKernel will be returned.
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
func (numStrKernel *NumberStrKernel) GetDefaultNumStrFmtSpec(
	errorPrefix interface{}) (
	NumStrFormatSpec,
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
			"GetDefaultNumStrFmtSpec()",
		"")

	if err != nil {
		return NumStrFormatSpec{}, err
	}

	if new(numStrFmtSpecNanobot).isNOP(
		&numStrKernel.numStrFormatSpec) {

		// This is a NOP!
		// numStrKernel.numStrFormatSpec is invalid.
		// Set default numStrKernel.numStrFormatSpec
		// to stand US Signed Number Minus Format
		// Specification.
		numStrKernel.numStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumDefaultsUSMinus(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy(
				"numStrKernel.numStrFormatSpec "+
					"Default US NumStrNumberFieldSpec"))

		if err != nil {
			return NumStrFormatSpec{}, err
		}
	}

	return numStrKernel.numStrFormatSpec.CopyOut(
		ePrefix.XCpy("<-numStrKernel.numStrFormatSpec"))
}

//	GetExcessFractionalTrailingZerosCount
//
//	Returns the count of excess fractional trailing
//	zeros for the numeric value contained in the current
//	instance of NumberStrKernel.
//
//	If the fractional digits value is zero, the count
//	will NOT include the first zero to the right of the
//	decimal point.
//
//	This method differs significantly from method:
//
//		NumberStrKernel.GetFractionalTrailingZerosCount()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Numeric Value:  5.1000
//		Count of Excess Fractional Trailing Zeros: 3
//
//	Example-2
//		Numeric Value:  5.000
//		Count of Excess Fractional Trailing Zeros: 2
//
//	Example-3
//		Numeric Value:  50.320
//		Count of Excess Fractional Trailing Zeros: 1
//
//	Example-4
//		Numeric Value:  0.000
//		Count of Excess Fractional Trailing Zeros: 2
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
//		Returns the number of excess trailing zeros for
//		the fractional part of the numeric value
//		contained in the current instance of
//		NumberStrKernel.
//
//		NOTE: If the fractional value is zero, the
//		returned count of excess leading zeros will NOT
//		include the first zero to the right of the
//		decimal point.
//
//			Example
//				Numeric Value:  000
//
//				Count of Excess Fractional Trailing Zeros:
//								2
func (numStrKernel *NumberStrKernel) GetExcessFractionalTrailingZerosCount() uint64 {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var excessFracTrailingZeros uint64 = 0

	lenFracChars :=
		uint64(len(numStrKernel.fractionalDigits.CharsArray))

	if lenFracChars == 0 {

		return excessFracTrailingZeros
	}

	excessFracTrailingZeros =
		numStrKernel.fractionalDigits.GetCountTrailingZeros()

	if excessFracTrailingZeros < lenFracChars {

		return excessFracTrailingZeros
	}

	excessFracTrailingZeros--

	return excessFracTrailingZeros
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

//	GetFractionalTrailingZerosCount
//
//	Returns the count of all fractional digit trailing
//	zeros.
//
//	This method differs significantly from method:
//
//		NumberStrKernel.GetExcessFractionalTrailingZerosCount()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		Numeric Value:  1.000
//		Count of Total Fractional Trailing Zeros: 3
//
//	Example-2
//		Numeric Value:  0.000
//		Count of Total Fractional Trailing Zeros: 3
//
//	Example-3
//		Numeric Value:  0.0
//		Count of Total Fractional Trailing Zeros: 1
//
//	Example-4
//		Numeric Value:  050.3200
//		Count of Total Fractional Trailing Zeros: 2
//
//	Example-5
//		Numeric Value:  50.32
//		Count of Total Fractional Trailing Zeros: 0
//
//	Example-6
//		Numeric Value:  97
//		Count of Total Fractional Trailing Zeros: 0
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
//		Returns the number of total trailing zeros for
//		the fractional part of the numeric value represented
//		by this instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetFractionalTrailingZerosCount() uint64 {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var totalFracTrailingZeros uint64 = 0

	lenFracChars :=
		uint64(len(numStrKernel.fractionalDigits.CharsArray))

	if lenFracChars == 0 {

		return totalFracTrailingZeros
	}

	totalFracTrailingZeros =
		numStrKernel.fractionalDigits.GetCountTrailingZeros()

	return totalFracTrailingZeros
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
// # BE ADVISED
//
//	(1) All leading zeros are deleted from the returned
//		instance of 'consolidatedIntegerKernel'. Any
//		trailing zeros are NOT deleted.
//
//	(2)	The number sign of the returned instance of
//		'consolidatedIntegerKernel' is ALWAYS positive.
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

// GetIntFracDigitsStr
//
// Returns a string containing all the integer and
// fractional numeric digits extracted from the current
// instance of NumberStrKernel. These integer and
// fractional digits are therefore concatenated and
// returned as a single string.
//
// If the combined integer and fractional digits
// represent a negative numeric value, a leading or
// trailing minus sign may be included depending on
// specifications contained in input parameters,
// 'includeNegativeNumSign' and 'leadingMinusSign'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Depending on the rounding specifications
//		contained in input parameters 'roundingType' and
//		'roundToFractionalDigits', the numeric value
//		represented by the current instance of
//		NumberStrKernel will be rounded before being
//		converted into a concatenated string of integer
//		and fractional digits.
//
//	(2)	If the combined integer and fractional digits
//		represent a negative numeric value, a leading or
//		trailing minus sign may be included in the
//		returned string depending on specifications
//		contained in input parameters,
//		'includeNegativeNumSign' and 'leadingMinusSign'.
//
//	(3)	This method will NOT alter, modify or change any
//		data values contained in the current instance of
//		NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	includeNegativeNumSign		bool
//
//		If this parameter is set to 'true' and the
//		numeric value of the current NumberStrKernel
//		instance is negative, a leading or trailing minus
//		sign ('-') will be included in the returned
//		string of integer and fractional digits.
//
//		The positioning of the minus sign as a leading or
//		tailing minus sign is controlled by input
//		parameter 'leadingMinusSign'.
//
//	leadingMinusSign			bool
//
//		If the numeric value of the current
//		NumberStrKernel instance is negative and input
//		parameter 'includeNegativeNumSign' is set to
//		'true', 'leadingMinusSign' will control where
//		the minus sign is included in the returned
//		string of integer and fractional digits.
//
//		When this parameter is set to 'true', the
//		minus sign will be prefixed at the beginning
//		of the returned string.
//
//		If this parameter is set to 'false', the minus
//		sign will be appended to the end of the returned
//		string.
//
//		Again, if parameter 'includeNegativeNumSign' is
//		set to 'false', or if the numeric value is
//		positive, this parameter will be ignored and have
//		no effect.
//
//	roundingType				NumberRoundingType
//
//		The 'roundingType' setting determines whether
//		the numeric value contained in the current
//		instance of NumberStrKernel will be rounded
//		before it is converted into a concatenated
//		string of integer and fractional digits.
//
//		If no rounding operation is required, set this
//		parameter to NumRoundType.NoRounding().
//
//		'roundingType' is used to specify the type of
//		rounding algorithm that will be applied for the
//		rounding of fractional digits contained in the
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//	string
//
//		If this method completes successfully, a string
//		will be returned containing the concatenated
//		integer and fractional digits extracted from the
//		current instance of NumberStrKernel.
//
//		If the current instance of NumberStrKernel
//		represents a negative numeric value, a leading
//		or trailing minus sign may be included in the
//		returned string depending on the settings for
//		input parameters, 'includeNegativeNumSign' and
//		'roundToFractionalDigits'.
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
func (numStrKernel *NumberStrKernel) GetIntFracDigitsStr(
	includeNegativeNumSign bool,
	leadingMinusSign bool,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
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

	var concatenatedIntFracDigits string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntFracDigitsStr()",
		"")

	if err != nil {
		return concatenatedIntFracDigits, err
	}

	var allIntFracDigits RuneArrayDto

	allIntFracDigits,
		err = new(numberStrKernelMolecule).getAllIntFracDigits(
		numStrKernel,
		includeNegativeNumSign,
		leadingMinusSign,
		roundingType,
		roundToFractionalDigits,
		ePrefix.XCpy(
			""))

	return allIntFracDigits.GetCharacterString(), err
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

// GetNumberOfFractionalDigits
//
// Returns the number of fractional digits in the
// current instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfFractionalDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return len(numStrKernel.fractionalDigits.CharsArray)
}

// GetNumberOfIntegerDigits
//
// Returns the number of integer digits in the current
// instance of NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfIntegerDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return len(numStrKernel.integerDigits.CharsArray)
}

//	GetNumberOfNumericDigits
//
//	Returns the number of numeric digits contained in
//	the current instance of NumberStrKernel.
//
//	The number of numeric digits is computed by adding
//	the number of integer digits plus the number of
//	fractional digits contained in the current instance
//	of NumberStrKernel.
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
//	int
//
//		This method returns the number of integer digits
//		plus the number of fractional digits currently
//		residing in the current instance of
//		NumberStrKernel.
func (numStrKernel *NumberStrKernel) GetNumberOfNumericDigits() int {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetRuneArrayLength() +
		numStrKernel.fractionalDigits.GetRuneArrayLength()

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
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
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

// IsValidInstance
//
// Performs a diagnostic review of the data values
// encapsulated in the current NumberStrKernel
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this
// method will return a boolean value of 'false'.
//
// If all data elements are determined to be valid,
// this method returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// NumberStrKernel.IsValidInstanceError() with
// the sole exceptions being that this method takes
// no input parameters and returns a boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal member data values
//		contained in the current instance of
//		NumberStrKernel are found to be invalid, this
//		method will return a boolean value of 'false'.
//
//		If all internal member data variables contained
//		in the current instance of NumberStrKernel are
//		found to be valid, this method returns a boolean
//		value of 'true'.
func (numStrKernel *NumberStrKernel) IsValidInstance() bool {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	isValid,
		_ := new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			nil)

	return isValid
}

// IsValidInstanceError
//
// Performs a diagnostic review of the data values
// encapsulated in the current NumberStrKernel
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this
// method will return an error.
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
//		If any of the internal member data variables
//		contained in the current instance of
//		NumberStrKernel are found to be invalid, this
//		method will return an error configured with
//		an appropriate message identifying the invalid
//		member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any internal member data values are found to be
//		invalid, the returned error Type will encapsulate
//		an appropriate error message. This returned error
//		message will incorporate the method chain and text
//		passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the
//		beginning of the error message.
func (numStrKernel *NumberStrKernel) IsValidInstanceError(
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
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	return err
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
//	Creates a new instance of NumberStrKernel converted
//	from a numeric value passed as a *big.Rat type
//	through input parameter 'bigRatNum'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	big.Rat Rounding Algorithm
//		Before being converted to an instance of
//		NumberStrKernel, the big.Rat numeric value will
//		be rounded by the Golang 'big' package function:
//
//			(*Rat) FloatString
//
//		The rounding algorithm used is described in the
//		big.Rat documentation as:
//
//			The last digit is rounded to nearest, with
//			halves rounded away from zero.
//				https://pkg.go.dev/math/big#Rat
//
//	(2)	Number of fractional digits
//		During the rounding operation, the number of
//		fractional digits to the right of the radix point
//		(decimal point) will be controlled by input
//		paramter 'roundToFractionalDigits'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigRatNum					*big.Rat
//
//		A pointer to an instance of the numeric value
//		type big.Rat. This numeric value will be
//		converted to, and returned as, a Native Number
//		String.
//
//		Before being converted to a Native Number String,
//		this numeric value will be rounded by the Golang
//		'big' package functions as specified by input
//		parameter, 'roundToFractionalDigits'.
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). Effectively this defines
//		the number of fractional digits remaining after
//		completion of the number rounding operation
//		performed by the Golang package functions.
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
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'numericValue'.
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

	var nativeNumStr string

	nativeNumStr,
		err = new(MathBigRatHelper).
		BigRatToNativeNumStr(
			bigRatNum,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {

		return newNumStrKernel, err
	}

	err = new(numberStrKernelQuark).
		setNumStrKernelFromNativeNumStr(
			&newNumStrKernel,
			nativeNumStr,
			ePrefix.XCpy(
				"newNumStrKernel"))

	return newNumStrKernel, err
}

//	NewFromNumericValue
//
//	Creates a new instance of NumberStrKernel based on
//	a numeric value passed as an empty interface.
//
//
//	The numeric value passed to this method by input
//	parameter 'numericValue' MUST BE convertible to one
//	of the types defined below:
//
//		float32, float64, big.Float
//		*float32, *float64, *big.Float
//		*BigFloatDto, BigFloatDto
//		int8, int16, int, int32, int64, big.Int
//		*int8, *int16, *int, *int32, *int64, *big.Int
//		uint8, uint16, uint, uint32, uint64
//		*uint8, *uint16, *uint, *uint32, *uint64
//		*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//		*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//		*NumberStrKernel, NumberStrKernel
//
//	This numeric value is then used to configure and
//	return a new instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	big.Rat Accuracy
//		This method will default big.Rat and *big.Rat
//		values default to 2,000 decimal places of
//		accuracy. Adjust rounding parameters as
//		required or use the following method for more
//		granular control over big.Rat rounding:
//			NewFromNumericValue.NewFromBigRat()
//
//	(2) big.Rat Rounding Algorithm
//		Before being converted to an instance of
//		NumberStrKernel, the big.Rat numeric value will
//		be rounded by the Golang 'big' package function:
//
//			(*Rat) FloatString
//
//		The rounding algorithm used is described in the
//		big.Rat documentation as:
//
//			The last digit is rounded to nearest, with
//			halves rounded away from zero.
//				https://pkg.go.dev/math/big#Rat
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numericValue				interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32, float64, big.Float
//			*float32, *float64, *big.Float
//			*BigFloatDto, BigFloatDto
//			int8, int16, int, int32, int64, big.Int
//			*int8, *int16, *int, *int32, *int64, *big.Int
//			uint8, uint16, uint, uint32, uint64
//			*uint8, *uint16, *uint, *uint32, *uint64
//			*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//			*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//			*NumberStrKernel, NumberStrKernel
//
//		This numeric value will be used to configure and
//		return a new instance of NumberStrKernel.
//
//		If 'numericValue' is NOT convertible to one of
//		the types listed above, an error will be
//		returned.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		returned instance of NumberStrKernel.
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
//		This parameter is ignored if 'roundingType' is
//		set to 'NumRoundType.NoRounding()'.
//
//		This parameter will be used to control the number
//		of fractional digits to the right of the decimal
//		point in the returned instance of
//		NumberStrKernel.
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
//	NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured and
//		populated with the numeric value passed in paramter,
//		'numericValue'.
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
func (numStrKernel *NumberStrKernel) NewFromNumericValue(
	numericValue interface{},
	roundingType NumberRoundingType,
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
			"NewFromNumericValue()",
		"")

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numberStrKernelMolecule).
		convertNumericValueToKernel(
			&newNumStrKernel,
			numericValue,
			ePrefix)

	if err != nil ||
		roundingType == NumRoundType.NoRounding() {

		return newNumStrKernel, err
	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {
		return newNumStrKernel, err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&newNumStrKernel,
		numStrRoundingSpec,
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
// # BE ADVISED
//
//	This method will retain leading integer zeros.
//		Example:
//								Input String: 0012345.678
//			Converted Internal Numeric Value: 0012345.678
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

//	NewFromStringDigits
//
//	Creates and returns a new instance of NumberStrKernel
//	generated from input parameters containing integer
//	digit and fractional digit strings.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will retain leading integer zeros.
//		Example:
//								Input String: 0012345.678
//			Converted Internal Numeric Value: 0012345.678
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				string
//
//		A string of numeric digits used to populate the
//		integer part of the numeric value contained in the
//		returned instance of NumberStrKernel.
//
//		This string should contain only numeric digit
//		characters. Do NOT include number signs such as
//		minus signs.
//
//	fractionalDigits			string
//
//		A string of numeric digits used to populate the
//		fractional part of the numeric value contained in
//		the returned instance of NumberStrKernel
//
//		This string should contain only numeric digit
//		characters. Do NOT include number signs such as
//		minus signs.
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
//		This collection allows the flexibility for
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
		err = new(numStrHelperQuark).extractNumRunes(
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

// NewParseDirtyNumberStr
//
// Receives a Dirty Number String, extracts a valid
// Native Number String, calculates a numeric value,
// and proceeds to configure and return a new instance
// of NumberStrKernel.
//
// A "Dirty Number String" is a malformed number string
// containing numeric digits which will be converted to a
// properly formatted Native Number String and used to
// configure the returned instance of NumberStrKernel.
//
//	Examples Of Dirty Number Strings
//
//		$1,254.65
//		1 000 000,00 €
//		1.000.000,00 €
//		6,78,90,00,00,00,00,000
//		6,7890,0000,0000,0000
//
// All the examples shown above are valid number string
// formats used by different countries and cultures.
// There is nothing wrong with these formats. The term
// "Dirty" simply distinguishes these formats from the
// Native Number String format required by many
// functions and packages in the Go Programming
// Language. The Native Number String format is one
// which is used as a standard format for numeric
// conversions performed in the Go Programming Language
// as well as many other programming languages.
//
// The Dirty Number String passed as input parameter
// 'dirtyNumberStr' is expected to comply with the
// following requirements:
//
//  1. The dirty number string must contain numeric
//     digit characters zero to nine inclusive (0-9).
//
//  2. The dirty number string must contain a radix
//     point or decimal separator to separate
//     integer and fractional digits in a floating
//     point numeric value. This decimal separator
//     is specified by input parameter,
//     'decimalSeparator'.
//
//     If no decimal separator is identified in the
//     dirty number string, the numeric value is
//     assumed to be an integer value.
//
//  3. The dirty number string must designate
//     negative numeric values using one of the
//     following three negative number symbols:
//
//     (a)	A Leading Minus Sign ('-').
//     Example: -123.45
//
//     (b)	A Trailing Minus Sign ('-').
//     Example: 123.45-
//
//     (c) A combination of leading and trailing
//     Parentheses ('()').
//     Example: (123.45)
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirtyNumberStr				string
//
//		This number string contains the numeric digits
//		which will be extracted, converted to a valid
//		numeric value and used to configure the
//		returned instance of NumberStrKernel.
//
//		The 'dirtyNumberStr' is expected to comply with
//		the following requirements:
//
//		1.	The dirty number string must contain numeric
//			digit characters zero to nine inclusive (0-9).
//
//		2.	The dirty number string must contain a radix
//			point or decimal separator to separate
//			integer and fractional digits in a floating
//			point numeric value. This decimal separator
//			is specified by input parameter,
//			'decimalSeparator'.
//
//			If no decimal separator is identified in the
//			dirty number string, the numeric value is
//			assumed to be an integer value.
//
//		3.	The dirty number string must designate
//			negative numeric values using one of the
//			following three negative number symbols:
//
//			(a)	A Leading Minus Sign ('-').
//				Example: -123.45
//
//			(b)	A Trailing Minus Sign ('-').
//				Example: 123.45-
//
//			(c) A combination of leading and trailing
//				Parentheses ('()').
//				Example: (123.45)
//
//		If 'dirtyNumberStr' does not contain any numeric
//		digits, an error will be returned.
//
//	decimalSeparator			string
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		dirty number string passed as input parameter
//		'dirtyNumberStr'.
//
//		The decimal separator will typically consist of
//		one or more non-numeric characters.
//
//		If 'decimalSeparator' consists of an empty
//		or zero length sting, it is assumed that the
//		numeric value contained in input parameter
//		'dirtyNumberStr' is an integer value.
//
//		In the US, Australia, UK, most of Canada and many
//		other countries the period ('.'), or decimal
//		point, separates integer and fractional digits
//		within a floating point numeric value.
//
//		Other countries, including many in the European
//		Union, use the comma (',') to separate integer
//		and fractional digits within a number string.
//
//		If 'decimalSeparator' contains any one of the
//		following invalid characters, an error will be
//		returned.
//
//			Invalid Decimal Separator Characters
//							'-'
//							'('
//							')'
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		new returned instance of NumberStrKernel
//		(newNumStrKernel).
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//	newNumStrKernel				NumberStrKernel
//
//		If this method completes successfully, a new
//		instance of NumberStrKernel will be returned
//		configured with the numeric value parsed from
//		the Dirty Number String input parameter
//		'dirtyNumberStr'.
//
//	numStrStatsDto				NumberStrStatsDto
//
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		numeric value encapsulated in the returned
//		instance of NumberStrKernel, 'newNumStrKernel'.
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
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) NewParseDirtyNumberStr(
	dirtyNumberStr string,
	decimalSeparator string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	newNumStrKernel NumberStrKernel,
	numStrStatsDto NumberStrStatsDto,
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
			"NewParseDirtyNumberStr()",
		"")

	if err != nil {
		return newNumStrKernel,
			numStrStatsDto,
			err
	}

	numStrStatsDto,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedDirtyNumStr(
			&newNumStrKernel,
			dirtyNumberStr,
			decimalSeparator,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"newNumStrKernel"))

	return newNumStrKernel,
		numStrStatsDto,
		err
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
// # Examples of French Number Strings
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
		err = new(numStrHelperQuark).extractNumRunes(
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
// ----------------------------------------------------------------
//
// # German Number String Examples
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
// # BE ADVISED
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
		err = new(numStrHelperQuark).extractNumRunes(
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

// NewParseNativeNumberStr
//
// Receives a Native Number String, calculates the
// numeric value contained therein and proceeds to
// configure and return a new instance of NumberStrKernel
// using that calculated numeric value.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math'
// and 'big' (big.Int, big.Float, big.Rat) routinely
// parse and convert this type of number string to
// numeric values. In addition, Native Number Strings are
// frequently consumed by external library functions such
// as this one (String Mechanics 'strmech') to convert
// strings to numeric values and numeric values to
// strings.
//
// While this format is inconsistent with many national
// and cultural formatting conventions, number strings
// which fail to implement this standardized formatting
// protocol will generate errors in some Golang library
// functions.
//
//	Examples Of Native Number Strings
//		1000000
//		12.5483
//		-1000000
//		-12.5483
//
// A valid Native Number String must conform to the
// standardized formatting criteria defined below:
//
//  1. A Native Number String Consists of numeric
//     character digits zero through nine inclusive
//     (0-9).
//
//  2. A Native Number String will include a period
//     or decimal point ('.') to separate integer and
//     fractional digits within a number string.
//
//     Native Number String Floating Point Value:
//     123.1234
//
//  3. A Native Number String will always format
//     negative numeric values with a leading minus sign
//     ('-').
//
//     Native Number String Negative Value:
//     -123.2
//
//  4. A Native Number String WILL NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Native Number String: 1000000
//
//  5. Native Number Strings will only consist of:
//
//     (a)	Numeric digits zero through nine inclusive (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//  6. A Native Number String will NEVER include
//     currency symbols.
//
//  7. A Native Number String will NEVER include
//     leading integer zeros or trailing fractional
//     zeros.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will delete any leading integer zero values
//	or trailing fractional zeros.
//
//		Example:
//			Input String:  00123.45600
//			Converted Native Number String: 123.456
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		nativeNumStr				string
//
//			A Native Number String containing the numeric
//			character digits which will be converted to, and
//			stored in, the NumberStrKernel object passed as
//			input parameter 'numStrKernel'.
//
//			The term 'Native Number String' means that the
//			number string format is designed to interoperate
//			with the Golang programming language library
//			functions and packages. Types like 'strconv',
//			'strings', 'math' and 'big' (big.Int, big.Float,
//			big.Rat) routinely parse and convert this type of
//			number string to generate numeric values. In
//			addition, Native Number Strings are frequently
//			consumed by external library functions such	as
//			this one (String Mechanics 'strmech') to convert
//			strings to numeric values and numeric values to
//			strings.
//
//			If 'nativeNumStr' fails to meet the formatting
//			criteria for a Native Number String defined
//			below, an error will be returned.
//
//			A valid Native Number String must conform to the
//			standardized formatting criteria defined below:
//
//		 	1.	A Native Number String Consists of numeric
//		 	  	character digits zero through nine inclusive
//		 	  	(0-9).
//
//		 	2.	A Native Number String will include a period
//		 	  	or decimal point ('.') to separate integer and
//		 	  	fractional digits within a number string.
//
//		 	  	Native Number String Floating Point Value:
//		 	   				123.1234
//
//		 	3.	A Native Number String will always format
//		 	  	negative numeric values with a leading minus sign
//		 	  	('-').
//
//		 	  	Native Number String Negative Value:
//		 	  					-123.2
//
//		 	4.	A Native Number String WILL NEVER include integer
//				separators such as commas (',') to separate
//				integer digits by thousands.
//
//		 	   					NOT THIS: 1,000,000
//		 	   		Native Number String: 1000000
//
//		 	5.	Native Number Strings will only consist of:
//
//				(a)	Numeric digits zero through nine inclusive (0-9).
//
//				(b)	A decimal point ('.') for floating point
//					numbers.
//
//				(c)	A leading minus sign ('-') in the case of
//					negative numeric values.
//
//			6.	A Native Number String will NEVER include
//				currency symbols.
//
//	 	7.	A Native Number String will NEVER include
//				leading integer zeros or trailing fractional
//				zeros.
//
//		roundingType				NumberRoundingType
//
//			This enumeration parameter is used to specify the
//			type of rounding algorithm that will be applied for
//			the	rounding of fractional digits contained in the
//			new returned instance of NumberStrKernel
//			(newNumStrKernel).
//
//			If in doubt as to a suitable rounding method,
//			'HalfAwayFromZero' is recommended.
//
//			Possible values are listed as follows:
//				NumRoundType.None()	- Invalid Value
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
//			Definitions:
//
//				NoRounding
//
//					Signals that no rounding operation will be
//					performed on fractional digits. The
//					fractional digits will therefore remain
//					unchanged.
//
//				HalfUpWithNegNums
//
//					Half Round Up Including Negative Numbers.
//					This method is intuitive but may produce
//					unexpected results when applied to negative
//					numbers.
//
//					'HalfUpWithNegNums' rounds .5 up.
//
//						Examples of 'HalfUpWithNegNums'
//						7.6 rounds up to 8
//						7.5 rounds up to 8
//						7.4 rounds down to 7
//						-7.4 rounds up to -7
//						-7.5 rounds up to -7
//						-7.6 rounds down to -8
//
//				HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This
//				method is also considered intuitive but may
//				produce unexpected results when applied to
//				negative numbers.
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
//				HalfAwayFromZero
//
//					The 'HalfAwayFromZero' method rounds .5 further
//					away from zero.	It provides clear and consistent
//					behavior when dealing with negative numbers.
//
//						Examples of HalfAwayFromZero
//
//						7.6 rounds away to 8
//						7.5 rounds away to 8
//						7.4 rounds to 7
//						-7.4 rounds to -7
//						-7.5 rounds away to -8
//						-7.6 rounds away to -8
//
//				HalfTowardsZero
//
//					Round Half Towards Zero. 'HalfTowardsZero' rounds
//					0.5	closer to zero. It provides clear and
//					consistent behavior	when dealing with negative
//					numbers.
//
//						Examples of HalfTowardsZero
//
//						7.6 rounds away to 8
//						7.5 rounds to 7
//						7.4 rounds to 7
//						-7.4 rounds to -7
//						-7.5 rounds to -7
//						-7.6 rounds away to -8
//
//				HalfToEven
//
//					Round Half To Even Numbers. 'HalfToEven' is
//					also called	Banker's Rounding. This method
//					rounds 0.5 to the nearest even digit.
//
//						Examples of HalfToEven
//
//						7.5 rounds up to 8 (because 8 is an even
//						number)	but 6.5 rounds down to 6 (because
//						6 is an even number)
//
//						HalfToEven only applies to 0.5. Other
//						numbers (not ending	in 0.5) round to
//						nearest as usual, so:
//
//						7.6 rounds up to 8
//						7.5 rounds up to 8 (because 8 is an even number)
//						7.4 rounds down to 7
//						6.6 rounds up to 7
//						6.5 rounds down to 6 (because 6 is an even number)
//						6.4 rounds down to 6
//
//				HalfToOdd
//
//					Round Half to Odd Numbers. Similar to 'HalfToEven',
//					but in this case 'HalfToOdd' rounds 0.5 towards odd
//					numbers.
//
//						Examples of HalfToOdd
//
//						HalfToOdd only applies to 0.5. Other numbers
//						(not ending	in 0.5) round to nearest as usual.
//
//						7.5 rounds down to 7 (because 7 is an odd number)
//
//						6.5 rounds up to 7 (because 7 is an odd number)
//
//						7.6 rounds up to 8
//						7.5 rounds down to 7 (because 7 is an odd number)
//						7.4 rounds down to 7
//						6.6 rounds up to 7
//						6.5 rounds up to 7 (because 7 is an odd number)
//						6.4 rounds down to 6
//
//				Randomly
//
//					Round Half Randomly. Uses a Random Number Generator
//					to choose between rounding 0.5 up or down.
//
//					All numbers other than 0.5 round to the nearest as
//					usual.
//
//				Floor
//
//					Yields the nearest integer down. Floor does not apply
//					any	special treatment to 0.5.
//
//					Floor Function: The greatest integer that is less than
//					or equal to x
//
//					Source:
//						https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//					In mathematics and computer science, the floor function
//					is the function that takes as input a real number x,
//					and gives as output the greatest integer less than or
//					equal to x,	denoted floor(x) or ⌊x⌋.
//
//					Source:
//						https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//						Number     Floor
//						 2           2
//						 2.4         2
//						 2.9         2
//						-2.5        -3
//						-2.7        -3
//						-2          -2
//
//				Ceiling
//
//					Yields the nearest integer up. Ceiling does not
//					apply any special treatment to 0.5.
//
//					Ceiling Function: The least integer that is
//					greater than or	equal to x.
//					Source:
//						https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//					The ceiling function maps x to the least integer
//					greater than or equal to x, denoted ceil(x) or
//					⌈x⌉.[1]
//
//					Source:
//						https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//						Examples of Ceiling
//
//							Number    Ceiling
//							 2           2
//							 2.4         3
//							 2.9         3
//							-2.5        -2
//							-2.7        -2
//							-2          -2
//
//				Truncate
//
//					Apply NO Rounding whatsoever. The Round From Digit
//					is dropped or deleted. The Round To Digit is NEVER
//					changed.
//
//					Examples of Truncate
//
//						Example-1
//						Number: 23.14567
//						Objective: Round to two decimal places to
//						the right of the decimal point.
//						Rounding Method: Truncate
//						Round To Digit:   4
//						Round From Digit: 5
//						Rounded Number:   23.14 - The Round From Digit
//						is dropped.
//
//						Example-2
//						Number: -23.14567
//						Objective: Round to two decimal places to
//						the right of the decimal point.
//						Rounding Method: Truncate
//						Round To Digit:   4
//						Round From Digit: 5
//						Rounded Number:  -23.14 - The Round From Digit
//						is dropped.
//
//		roundToFractionalDigits		int
//
//			When set to a positive integer value, this
//			parameter controls the number of digits to the
//			right of the radix point or decimal separator
//			(a.k.a. decimal point). This controls the number
//			of fractional digits remaining after completion
//			of the number rounding operation.
//
//			If input parameter 'roundingType' is set to
//			NumRoundType.NoRounding(),
//			'roundToFractionalDigits' is ignored and no
//			rounding operation is performed.
//
//		errorPrefix					interface{}
//
//			This object encapsulates error prefix text which
//			is included in all returned error messages.
//			Usually, it contains the name of the calling
//			method or methods listed as a method or function
//			chain of execution.
//
//			If no error prefix information is needed, set this
//			parameter to 'nil'.
//
//			This empty interface must be convertible to one of
//			the following types:
//
//			1.	nil
//					A nil value is valid and generates an
//					empty collection of error prefix and
//					error context information.
//
//			2.	string
//					A string containing error prefix
//					information.
//
//			3.	[]string
//					A one-dimensional slice of strings
//					containing error prefix information.
//
//			4.	[][2]string
//					A two-dimensional slice of strings
//			   		containing error prefix and error
//			   		context information.
//
//			5.	ErrPrefixDto
//					An instance of ErrPrefixDto.
//					Information from this object will
//					be copied for use in error and
//					informational messages.
//
//			6.	*ErrPrefixDto
//					A pointer to an instance of
//					ErrPrefixDto. Information from
//					this object will be copied for use
//					in error and informational messages.
//
//			7.	IBasicErrorPrefix
//					An interface to a method
//					generating a two-dimensional slice
//					of strings containing error prefix
//					and error context information.
//
//			If parameter 'errorPrefix' is NOT convertible
//			to one of the valid types listed above, it will
//			be considered invalid and trigger the return of
//			an error.
//
//			Types ErrPrefixDto and IBasicErrorPrefix are
//			included in the 'errpref' software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	newNumStrKernel				NumberStrKernel
//
//		If this method completes successfully, a new
//		instance of NumberStrKernel will be returned
//		configured with the numeric value parsed from
//		the Native Number String input parameter
//		'nativeNumStr'.
//
//	numStrStatsDto				NumberStrStatsDto
//
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		numeric value encapsulated in the returned
//		instance of NumberStrKernel, 'newNumStrKernel'.
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
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) NewParseNativeNumberStr(
	nativeNumStr string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	newNumStrKernel NumberStrKernel,
	numStrStatsDto NumberStrStatsDto,
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
			"NewParseNativeNumberStr()",
		"")

	if err != nil {
		return newNumStrKernel,
			numStrStatsDto,
			err
	}

	numStrStatsDto,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedNativeNumStr(
			&newNumStrKernel,
			nativeNumStr,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"newNumStrKernel"))

	return newNumStrKernel,
		numStrStatsDto,
		err
}

// NewParsePureNumberStr
//
// Receives a Pure Number String and proceeds to return
// the extracted numeric value as a new instance of type
// NumberStrKernel.
//
// A Pure Number String is defined as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters (0-9).
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point or decimal
//     separator. This could be, but is not limited to,
//     a decimal point ('.'). For example, many European
//     countries use the comma (',') as a radix point.
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
//  5. A pure number string will NEVER include currency
//     symbols.
//
// Users also have the option of specifying rounding
// parameters which will be used to round the numeric
// value extracted from 'pureNumberStr' before it is
// converted and returned as a new instance of
// NumberStrKernel.
//
// Note that return parameter 'pureNumStrComponents'
// will provide a statistical profile of the new
// returned instance of NumberStrKernel,
// 'newNumStrKernel'.
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
//		The Decimal Separator is used to identify and
//		separate integer and fractional digits while
//		parsing the Pure Number String.
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
//	newNumStrKernel 			NumberStrKernel
//
//		If this method completes successfully, a new instance
//		of NumberStrKernel will be returned configured with
//		the numeric value parsed and extracted from input
//		paramter, 'pureNumberStr'.
//
//	pureNumStrComponents		PureNumberStrComponents
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		PureNumberStrComponents. This data structure
//		contains an analysis and profile information on
//		the returned instance of NumberStrKernel
//		'newNumStrKernel'.
//
//		type PureNumberStrComponents struct {
//
//			NumStrStats NumberStrStatsDto
//
//				This data transfer object will return key
//				statistics on the numeric value encapsulated
//				by the current instance of NumberStrKernel.
//
//				type NumberStrStatsDto struct {
//
//				NumOfIntegerDigits					uint64
//
//					The total number of integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfSignificantIntegerDigits		uint64
//
//					The number of nonzero integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfFractionalDigits				uint64
//
//					The total number of fractional digits to the
//					right of the radix point or, decimal point,
//					in the subject numeric value.
//
//				NumOfSignificantFractionalDigits	uint64
//
//					The number of nonzero fractional digits to
//					the right of the radix point or, decimal
//					point, in the subject numeric value.
//
//				NumberValueType 					NumericValueType
//
//					This enumeration value specifies whether the
//					subject numeric value is classified either as
//					an integer or a floating point number.
//
//					Possible enumeration values are listed as
//					follows:
//						NumValType.None()
//						NumValType.FloatingPoint()
//						NumValType.Integer()
//
//				NumberSign							NumericSignValueType
//
//					An enumeration specifying the number sign
//					associated with the numeric value. Possible
//					values are listed as follows:
//						NumSignVal.None()		= Invalid Value
//						NumSignVal.Negative()	= -1
//						NumSignVal.Zero()		=  0
//						NumSignVal.Positive()	=  1
//
//				IsZeroValue							bool
//
//					If 'true', the subject numeric value is equal
//					to zero ('0').
//
//					If 'false', the subject numeric value is
//					greater than or less than zero ('0').
//				}
//
//
//
//			AbsoluteValueNumStr string
//
//			The number string expressed as an absolute value.
//			Be advised, this number string may be a floating
//			point number string containing fractional digits.
//
//			AbsoluteValAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			This string DOES NOT contain a leading number
//			sign (a.k.a. minus sign ('-')
//
//			SignedAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			If the numeric value is negative, a leading
//			minus sign will be prefixed at the beginning
//			of the number string.
//
//			NativeNumberStr string
//
//			A Native Number String representing the base
//			numeric value used to generate these profile
//			number string statistics.
//
//			A valid Native Number String must conform to the
//			standardized formatting criteria defined below:
//
//			 	1. A Native Number String Consists of numeric
//			 	   character digits zero through nine inclusive
//			 	   (0-9).
//
//			 	2. A Native Number String will include a period
//			 	   or decimal point ('.') to separate integer and
//			 	   fractional digits within a number string.
//
//			 	   Native Number String Floating Point Value:
//			 	   				123.1234
//
//			 	3. A Native Number String will always format
//			 	   negative numeric values with a leading minus sign
//			 	   ('-').
//
//			 	   Native Number String Negative Value:
//			 	   				-123.2
//
//			 	4. A Native Number String WILL NEVER include integer
//			 	   separators such as commas (',') to separate
//			 	   integer digits by thousands.
//
//			 	   					NOT THIS: 1,000,000
//			 	   		Native Number String: 1000000
//
//			 	5. Native Number Strings will only consist of:
//
//			 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//			 	   (b)	A decimal point ('.') for floating point
//			 	   		numbers.
//
//			 	   (c)	A leading minus sign ('-') in the case of
//			 	   		negative numeric values.
//
//		}
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) NewParsePureNumberStr(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	newNumStrKernel NumberStrKernel,
	pureNumStrComponents PureNumberStrComponents,
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
			"NewParsePureNumberStr()",
		"")

	if err != nil {
		return newNumStrKernel,
			pureNumStrComponents,
			err
	}

	pureNumStrComponents,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedPureNumStr(
			&newNumStrKernel,
			pureNumberStr,
			decSeparatorChars,
			leadingMinusSign,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"newNumStrKernel"))

	return newNumStrKernel,
		pureNumStrComponents,
		err
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
		err = new(numStrHelperQuark).extractNumRunes(
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
		ePrefix.XCpy("numStrKernel"))
}

// SetDefaultNumStrFmtSpec
//
// Sets the default Number String Format Specification
// for the current instance of to NumberStrKernel.
//
// The default Number String Format Specification is
// used and applied in formatting the number string for
// screen display, file output or printing.
//
// The internal member variable controlling default
// number string formatting which will be modified
// by this method is:
//
//	NumberStrKernel.numStrFormatSpec
//
// The default Number String Format Specification is
// used and applied by the following methods:
//
//	NumberStrKernel.FmtSignedNumStrPure()
//	NumberStrKernel.String()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT alter the numeric value
//	configured for the current instance of
//	NumberStrKernel. This method will only reconfigure
//	the internal member variable controlling default
//	number string formatting:
//
//		NumberStrKernel.numStrFormatSpec
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
func (numStrKernel *NumberStrKernel) SetDefaultNumStrFmtSpec(
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
			"SetDefaultNumStrFmtSpec()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(numStrFmtSpecElectron).
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

//	SetDefaultNumStrFmtSpecElements
//
//	Sets the default Number String Format Specification
//	for the current instance of to NumberStrKernel.
//
//	Options include customizing for currency symbols,
//	integer separation, number sign	management, radix
//	point symbols, and floating point number rounding.
//
//	If required, users have the options of
//	implementing the India or Chinese Numbering
//	Systems for integer separation.
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
// # BE ADVISED
//
//	(1)	If Currency Symbol formatting is NOT required,
//		set the Currency Symbol parameter to a 'NOP'.
//		'NOP' is a computer science term for
//		'No Operation'.
//
//			currencySymbol :=
//				new(NumStrNumberSymbolSpec).NewNOP()
//
//	(2)	This method will NOT alter the numeric value
//		configured for the current instance of
//		NumberStrKernel. This method will only
//		reconfigure the internal member variable
//		controlling default number string formatting:
//
//			NumberStrKernel.numStrFormatSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec				DecimalSeparatorSpec
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
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
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
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
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
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	negativeNumberSign				NumStrNumberSymbolSpec
//
//		The Number String Negative Number Sign
//		Specification is used to configure negative
//		number sign symbols for negative numeric
//		values formatted and displayed in number
//		stings.
//
//		If this parameter is submitted as an empty or
//		invalid Negative Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	positiveNumberSign 				NumStrNumberSymbolSpec
//
//		Positive number signs are commonly implied
//		and not specified. However, the user has
//		the option to specify a positive number sign
//		character or characters for positive numeric
//		values using this input parameter.
//
//		If this parameter is submitted as an empty or
//		invalid Positive Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//	zeroNumberSign					NumStrNumberSymbolSpec
//
//		The Number String Zero Number Sign
//		Specification is used to configure number
//		sign symbols for zero numeric values formatted
//		and displayed in number stings. Zero number signs
//		are commonly omitted because zero does not
//		technically qualify as either a positive or
//		negative value. However, the user has the option
//		to configure number sign symbols for zero values
//		if necessary.
//
//		If this parameter is submitted as an empty or
//		invalid Zero Number Sign Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+ "
//			Trailing Symbols: ""
//			Number String:   "+ 0.00"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: "+"
//			Number String:   "0.00+"
//
//	currencySymbol					NumStrNumberSymbolSpec
//
//		A Currency Symbol positioned next to a numeric
//		value designates the number as a monetary amount.
//
//		The Number String Currency Symbol Specification
//		is used to configure currency symbols for
//		positive, negative and zero numeric values
//		formatted and displayed in number stings.
//
//		If Currency Symbol formatting is NOT required,
//		set this parameter to a 'NOP'. 'NOP' is a
//		computer science term for 'No Operation'.
//
//			currencySymbol :=
//				new(NumStrNumberSymbolSpec).NewNOP()
//
//		If this parameter is submitted as an empty or
//		invalid Currency Symbol Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Examples of Currency Symbols include the Dollar
//		sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//		This instance of NumStrNumberSymbolSpec is used
//		to configure leading Currency Symbols, trailing
//		Currency Symbols or both leading and trailing
//		Currency Symbols.
//
//		Example-1: Leading Currency Symbols
//
//			Leading Currency Symbols: "$ "
//			Number String:   "$ 123.456"
//
//		Example-2: Leading Currency Symbols
//
//			Leading Currency Symbols: "$"
//			Number String:   "$123.456"
//
//		Example-3: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: "€"
//			Number String:   "123.456€"
//
//		Example-4: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: " €"
//			Number String:   "123.456 €"
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
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetDefaultNumStrFmtSpecElements(
	decSeparatorSpec DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbol NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetDefaultNumStrFmtSpecElements()",
		"")

	if err != nil {
		return err
	}

	return numStrKernel.numStrFormatSpec.
		SetNumFmtElements(
			decSeparatorSpec,
			intSeparatorSpec,
			negativeNumberSign,
			positiveNumberSign,
			zeroNumberSign,
			currencySymbol,
			numberFieldSpec,
			ePrefix.XCpy(
				"numStrKernel->"))
}

// SetDefaultNumStrFmtSpecPure
//
// Sets the default Number String Format Specification
// for the current instance of to NumberStrKernel.
//
// The default Number String Format Specification is
// used and applied in formatting the number string for
// screen display, file output or printing.
//
// The internal member variable controlling default
// number string formatting which will be modified
// by this method is:
//
//	NumberStrKernel.numStrFormatSpec
//
// The default Number String Format Specification is
// used and applied by the following methods:
//
//	NumberStrKernel.FmtNumStrDefaultRound()
//	NumberStrKernel.String()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT alter the numeric value
//	configured for the current instance of
//	NumberStrKernel. This method will only reconfigure
//	the internal member variable controlling default
//	number string formatting:
//
//		NumberStrKernel.numStrFormatSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned number string.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted floating
//		point number string.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		also known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingNumSymbols			bool
//
//		In Pure Number Strings, positive numeric values
//		are NOT configured with leading or trailing plus
//		signs ('+'). Negative values on the other hand,
//		are always designated by leading or trailing
//		minus sign ('-').
//
//		This parameter, 'leadingNumSymbols', controls
//		the positioning of minus signs for negative
//		numeric values within a	Number String.
//
//		When set to 'true', the returned number string
//		will configure minus signs for negative numbers
//		at the beginning of, or on the left side of, the
//		numeric value. In these cases, the minus sign is
//		said to be configured as a leading minus sign.
//		This is the positioning format used in the US,
//		UK, Australia and most of Canada. In addition,
//		library functions in 'Go' and other programming
//		languages generally expect leading minus signs
//		for negative numbers.
//
//			Example Leading Minus Sign:
//				"-123.456"
//
//		When parameter 'leadingNumSymbols' is set to
//		'false', the returned number string will
//		configure minus signs for negative numbers at the
//		end of, or on the right side of, the numeric
//		value. With this positioning format, the minus
//		sign is said to be configured as a trailing minus
//		sign. This is the positioning format used in
//		France, Germany and many countries in the
//		European Union.
//
//			Example Trailing Minus Sign:
//				"123.456-"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
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
func (numStrKernel *NumberStrKernel) SetDefaultNumStrFmtSpecPure(
	decSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
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
			"SetDefaultNumStrFmtSpecPure()",
		"")

	if err != nil {
		return err
	}

	numStrKernel.numStrFormatSpec.Empty()

	numStrKernel.numStrFormatSpec,
		err = new(NumStrFormatSpec).NewSignedPureNumberStr(
		decSeparatorChars,
		leadingNumSymbols,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"numStrKernel.numStrFormatSpec<-"))

	return err
}

//	SetDefaultSignedNumStrFmtSpecBasic
//
//	Reconfigures the current instance of
//	NumberStrKernel with a new default Signed Number
//	String Format using the basic set of input parameters.
//
//	This Number String Format is designed to format
//	'Signed' integer or floating point numeric values.
//
//	A signed number is an integer or floating point
//	numeric value which does NOT contain currency
//	symbols.
//
//	This method is one step above NewSignedNumSimple()
//	in terms of complexity and customization options.
//
//	This method provides additional options for multiple
//	integer grouping types including India and Chinese
//	Numbering.
//
//	Negative number sign symbols may be configured as
//	leading negative number signs, trailing negative
//	number signs or both. The combination of leading
//	and trailing negative number signs allows for the
//	configuration symbols like parentheses for the
//	formatting of negative numbers.
//
//		Example: (125.34)
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT alter the numeric value
//	configured for the current instance of
//	NumberStrKernel. This method will only reconfigure
//	the internal member variable controlling default
//	number string formatting:
//
//		NumberStrKernel.numStrFormatSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars				string
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
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
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
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//	leadingNegativeNumSign			string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegativeNumSign			string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	fieldLength						int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'fieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of 'fieldLength'
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 15
//				fieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 15
//				fieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = -1
//				fieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 2
//				fieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	fieldJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'fieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'fieldJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('fieldLength')
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
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 15
//				fieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 15
//				fieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = -1
//				fieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				fieldLength = 2
//				fieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
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
func (numStrKernel *NumberStrKernel) SetDefaultSignedNumStrFmtSpecBasic(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	fieldLength int,
	fieldJustification TextJustify,
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
			"SetDefaultSignedNumStrFmtSpecBasic()",
		"")

	if err != nil {
		return err
	}

	numStrKernel.numStrFormatSpec.Empty()

	numStrKernel.numStrFormatSpec,
		err = new(NumStrFormatSpec).NewSignedNumBasic(
		decSeparatorChars,
		intSeparatorChars,
		intGroupingType,
		leadingNegativeNumSign,
		trailingNegativeNumSign,
		fieldLength,
		fieldJustification,
		ePrefix.XCpy(
			"<-numStrKernel.numStrFormatSpec"))

	return err
}

// SetDefaultSignedNumStrFmtSpecSimple
//
// Sets the default Number String Format Specification
// for the current instance of to NumberStrKernel.
//
// The default Number String Format Specification is
// used and applied in formatting the number string for
// screen displays, file output or printing.
//
// The internal member variable controlling default
// number string formatting which will be modified
// by this method is:
//
//	NumberStrKernel.numStrFormatSpec
//
// The default Number String Format Specification is
// used and applied by the following methods:
//
//	NumberStrKernel.FmtNumStrDefaultRound()
//	NumberStrKernel.String()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT alter the numeric value
//	configured for the current instance of
//	NumberStrKernel. This method will only reconfigure
//	the internal member variable controlling default
//	number string formatting:
//
//		NumberStrKernel.numStrFormatSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned number string.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted floating
//		point number string.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		also known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
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
//		string, no error will be returned and integer
//		separation will be turned off. As a result,
//		integer digits will be displayed as a single
//		string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	leadingNumSymbols			bool
//
//		In Simple Number Strings, positive numeric values
//		are NOT configured with leading or trailing plus
//		signs ('+'). Negative values on the other hand,
//		are always designated by leading or trailing
//		minus sign ('-').
//
//		This parameter, 'leadingNumSymbols', controls
//		the positioning of minus signs for negative
//		numeric values within a	Number String.
//
//		When set to 'true', the returned number string
//		will configure minus signs for negative numbers
//		at the beginning of, or on the left side of, the
//		numeric value. In these cases, the minus sign is
//		said to be configured as a leading minus sign.
//		This is the positioning format used in the US,
//		UK, Australia and most of Canada. In addition,
//		library functions in 'Go' and other programming
//		languages generally expect leading minus signs
//		for negative numbers.
//
//			Example Leading Minus Sign:
//				"-123.456"
//
//		When parameter 'leadingNumSymbols' is set to
//		'false', the returned number string will
//		configure minus signs for negative numbers at the
//		end of, or on the right side of, the numeric
//		value. With this positioning format, the minus
//		sign is said to be configured as a trailing minus
//		sign. This is the positioning format used in
//		France, Germany and many countries in the
//		European Union.
//
//			Example Trailing Minus Sign:
//				"123.456-"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
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
func (numStrKernel *NumberStrKernel) SetDefaultSignedNumStrFmtSpecSimple(
	decSeparatorChars string,
	intSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
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
			"SetDefaultSignedNumStrFmtSpecSimple()",
		"")

	if err != nil {
		return err
	}

	numStrKernel.numStrFormatSpec.Empty()

	numStrKernel.numStrFormatSpec,
		err = new(NumStrFormatSpec).NewSignedNumSimple(
		decSeparatorChars,
		intSeparatorChars,
		leadingNumSymbols,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"numStrKernel.numStrFormatSpec<-"))

	return err
}

// SetFromBigRat
//
// Deletes and resets the internal values for the
// current instance of NumberStrKernel using a numeric
// value passed as a pointer to a big.Rat number
// (*big.Rat).
//
// The big.Rat numeric value passed through input
// paramter 'roundToFractionalDigits' will be used
// round that numeric value before it is converted
// and used to configure the current instance of
// NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	big.Rat Rounding Algorithm
//		Before being converted to an instance of
//		NumberStrKernel, the big.Rat numeric value will
//		be rounded by the Golang 'big' package function:
//
//			(*Rat) FloatString
//
//		The rounding algorithm used is described in the
//		big.Rat documentation as:
//
//			The last digit is rounded to nearest, with
//			halves rounded away from zero.
//				https://pkg.go.dev/math/big#Rat
//
//	(2)	Number of fractional digits
//		During the rounding operation, the number of
//		fractional digits to the right of the radix point
//		(decimal point) will be controlled by input
//		paramter 'roundToFractionalDigits'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values contained in the current
//	instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigRatNum					*big.Rat
//		A pointer to an instance of a big.Rat numeric
//		value. This numeric value will be used to
//		reconfigure the current instance of
//		NumberStrKernel.
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). Effectively this defines
//		the number of fractional digits remaining after
//		completion of the number rounding operation
//		performed by the Golang package function:
//
//			func (*Rat) FloatString
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
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//	 	This returned error message will incorporate the method
//	 	chain and text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the beginning
//	 	of the error message.
func (numStrKernel *NumberStrKernel) SetFromBigRat(
	bigRatNum *big.Rat,
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
			"SetFromBigRat()",
		"")

	if err != nil {
		return err
	}

	var nativeNumStr string

	nativeNumStr,
		err = new(MathBigRatHelper).
		BigRatToNativeNumStr(
			bigRatNum,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {

		return err
	}

	err = new(numberStrKernelQuark).
		setNumStrKernelFromNativeNumStr(
			numStrKernel,
			nativeNumStr,
			ePrefix.XCpy(
				"numStrKernel"))

	return err
}

// SetFromDirtyNumberStr
//
// Deletes and resets the internal values for the
// current instance of NumberStrKernel using a numeric
// value extracted and calculated from a Dirty Number
// String passed as input parameter 'dirtyNumberStr'.
//
// A "Dirty Number String" is a malformed number string
// containing numeric digits which will be converted to a
// properly formatted Native Number String and used to
// generate a numeric value used in reconfiguring the
// current instance of NumberStrKernel.
//
//	Examples Of Dirty Number Strings
//
//		$1,254.65
//		1 000 000,00 €
//		1.000.000,00 €
//		6,78,90,00,00,00,00,000
//		6,7890,0000,0000,0000
//
// All the examples shown above are valid number string
// formats used by different countries and cultures.
// There is nothing wrong with these formats. The term
// "Dirty" simply distinguishes these formats from the
// Native Number String format required by many
// functions and packages in the Go Programming
// Language. The Native Number String format is one
// which is used as a standard format for numeric
// conversions performed in the Go Programming Language
// as well as many other programming languages.
//
// The Dirty Number String passed as input parameter
// 'dirtyNumberStr' is expected to comply with the
// following requirements:
//
//  1. The dirty number string must contain numeric
//     digit characters zero to nine inclusive (0-9).
//
//  2. The dirty number string must contain a radix
//     point or decimal separator to separate
//     integer and fractional digits in a floating
//     point numeric value. This decimal separator
//     is specified by input parameter,
//     'decimalSeparator'.
//
//     If no decimal separator is identified in the
//     dirty number string, the numeric value is
//     assumed to be an integer value.
//
//  3. The dirty number string must designate
//     negative numeric values using one of the
//     following three negative number symbols:
//
//     (a)	A Leading Minus Sign ('-').
//     Example: -123.45
//
//     (b)	A Trailing Minus Sign ('-').
//     Example: 123.45-
//
//     (c) A combination of leading and trailing
//     Parentheses ('()').
//     Example: (123.45)
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirtyNumberStr				string
//
//		This number string contains the numeric digits
//		which will be extracted, converted to a valid
//		numeric value and used to configure the
//		returned instance of NumberStrKernel.
//
//		The 'dirtyNumberStr' is expected to comply with
//		the following requirements:
//
//		1.	The dirty number string must contain numeric
//			digit characters zero to nine inclusive (0-9).
//
//		2.	The dirty number string must contain a radix
//			point or decimal separator to separate
//			integer and fractional digits in a floating
//			point numeric value. This decimal separator
//			is specified by input parameter,
//			'decimalSeparator'.
//
//			If no decimal separator is identified in the
//			dirty number string, the numeric value is
//			assumed to be an integer value.
//
//		3.	The dirty number string must designate
//			negative numeric values using one of the
//			following three negative number symbols:
//
//			(a)	A Leading Minus Sign ('-').
//				Example: -123.45
//
//			(b)	A Trailing Minus Sign ('-').
//				Example: 123.45-
//
//			(c) A combination of leading and trailing
//				Parentheses ('()').
//				Example: (123.45)
//
//		If 'dirtyNumberStr' does not contain any numeric
//		digits, an error will be returned.
//
//	decimalSeparator			string
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		dirty number string passed as input parameter
//		'dirtyNumberStr'.
//
//		The decimal separator will typically consist of
//		one or more non-numeric characters.
//
//		If 'decimalSeparator' consists of an empty
//		or zero length sting, it is assumed that the
//		numeric value contained in input parameter
//		'dirtyNumberStr' is an integer value.
//
//		In the US, Australia, UK, most of Canada and many
//		other countries the period ('.'), or decimal
//		point, separates integer and fractional digits
//		within a floating point numeric value.
//
//		Other countries, including many in the European
//		Union, use the comma (',') to separate integer
//		and fractional digits within a number string.
//
//		If 'decimalSeparator' contains any one of the
//		following invalid characters, an error will be
//		returned.
//
//			Invalid Decimal Separator Characters
//							'-'
//							'('
//							')'
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		reconfigured current instance of NumberStrKernel.
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//	numStrStatsDto				NumberStrStatsDto
//
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		reconfigured numeric value encapsulated in the
//		current instance of NumberStrKernel.
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
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) SetFromDirtyNumberStr(
	dirtyNumberStr string,
	decimalSeparator string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	numStrStatsDto NumberStrStatsDto,
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
			"SetFromDirtyNumberStr()",
		"")

	if err != nil {
		return numStrStatsDto, err
	}

	numStrStatsDto,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedDirtyNumStr(
			numStrKernel,
			dirtyNumberStr,
			decimalSeparator,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"numStrKernel"))

	return numStrStatsDto,
		err
}

// SetFromNativeNumberStr
//
// Deletes and resets the internal values for the
// current instance of NumberStrKernel using a numeric
// value extracted and calculated from a Native Number
// String passed as input parameter 'nativeNumStr'.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math'
// and 'big' (big.Int, big.Float, big.Rat) routinely
// parse and convert this type of number string to
// numeric values. In addition, Native Number Strings are
// frequently consumed by external library functions such
// as this one (String Mechanics 'strmech') to convert
// strings to numeric values and numeric values to
// strings.
//
// While this format is inconsistent with many national
// and cultural formatting conventions, number strings
// which fail to implement this standardized formatting
// protocol will generate errors in some Golang library
// functions.
//
//	Examples Of Native Number Strings
//		1000000
//		12.5483
//		-1000000
//		-12.5483
//
// A valid Native Number String must conform to the
// standardized formatting criteria defined below:
//
//  1. A Native Number String Consists of numeric
//     character digits zero through nine inclusive
//     (0-9).
//
//  2. A Native Number String will include a period
//     or decimal point ('.') to separate integer and
//     fractional digits within a number string.
//
//     Native Number String Floating Point Value:
//     123.1234
//
//  3. A Native Number String will always format
//     negative numeric values with a leading minus sign
//     ('-').
//
//     Native Number String Negative Value:
//     -123.2
//
//  4. A Native Number String WILL NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Native Number String: 1000000
//
//  5. Native Number Strings will only consist of:
//
//     (a)	Numeric digits zero through nine inclusive (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//  6. A Native Number String will NEVER include
//     currency symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr				string
//
//		A Native Number String containing the numeric
//		character digits which will be converted to, and
//		stored in, the NumberStrKernel object passed as
//		input parameter 'numStrKernel'.
//
//		The term 'Native Number String' means that the
//		number string format is designed to interoperate
//		with the Golang programming language library
//		functions and packages. Types like 'strconv',
//		'strings', 'math' and 'big' (big.Int, big.Float,
//		big.Rat) routinely parse and convert this type of
//		number string to generate numeric values. In
//		addition, Native Number Strings are frequently
//		consumed by external library functions such	as
//		this one (String Mechanics 'strmech') to convert
//		strings to numeric values and numeric values to
//		strings.
//
//		If 'nativeNumStr' fails to meet the formatting
//		criteria for a Native Number String defined
//		below, an error will be returned.
//
//		A valid Native Number String must conform to the
//		standardized formatting criteria defined below:
//
//	 	1.	A Native Number String Consists of numeric
//	 	  	character digits zero through nine inclusive
//	 	  	(0-9).
//
//	 	2.	A Native Number String will include a period
//	 	  	or decimal point ('.') to separate integer and
//	 	  	fractional digits within a number string.
//
//	 	  	Native Number String Floating Point Value:
//	 	   				123.1234
//
//	 	3.	A Native Number String will always format
//	 	  	negative numeric values with a leading minus sign
//	 	  	('-').
//
//	 	  	Native Number String Negative Value:
//	 	  					-123.2
//
//	 	4.	A Native Number String WILL NEVER include integer
//			separators such as commas (',') to separate
//			integer digits by thousands.
//
//	 	   					NOT THIS: 1,000,000
//	 	   		Native Number String: 1000000
//
//	 	5.	Native Number Strings will only consist of:
//
//			(a)	Numeric digits zero through nine inclusive (0-9).
//
//			(b)	A decimal point ('.') for floating point
//				numbers.
//
//			(c)	A leading minus sign ('-') in the case of
//				negative numeric values.
//
//		6.	A Native Number String will NEVER include
//			currency symbols.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		reconfigured current instance of NumberStrKernel.
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//	numStrStatsDto				NumberStrStatsDto
//
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		updated numeric value encapsulated in the current
//		instance of NumberStrKernel.
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
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) SetFromNativeNumberStr(
	nativeNumStr string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	numStrStatsDto NumberStrStatsDto,
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
			"SetFromNativeNumberStr()",
		"")

	if err != nil {
		return numStrStatsDto, err
	}

	numStrStatsDto,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedNativeNumStr(
			numStrKernel,
			nativeNumStr,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"numStrKernel"))

	return numStrStatsDto,
		err
}

// SetFromNumericValue
//
//	Deletes and resets the internal values for the
//	current instance of NumberStrKernel using a numeric
//	value passed as an empty interface.
//
//	The numeric value passed to this method by input
//	parameter 'numericValue' MUST BE convertible to one
//	of the types defined below:
//
//		float32, float64, big.Float
//		*float32, *float64, *big.Float
//		*BigFloatDto, BigFloatDto
//		int8, int16, int, int32, int64,
//		big.Int, *big.Int
//		*int8, *int16, *int, *int32, *int64,
//		uint8, uint16, uint, uint32, uint64
//		*uint8, *uint16, *uint, *uint32, *uint64
//		*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//		*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//		*NumberStrKernel, NumberStrKernel
//
//	This numeric value is then used to reconfigure the
//	current instance NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	big.Rat Accuracy
//		This method will default big.Rat and *big.Rat
//		values default to 2,000 decimal places of
//		accuracy. Adjust rounding parameters as
//		required or use the following method for more
//		granular control over big.Rat rounding:
//			NewFromNumericValue.SetFromBigRat()
//
//	(2) big.Rat Rounding Algorithm
//		Before being converted to an instance of
//		NumberStrKernel, the big.Rat numeric value will
//		be rounded by the Golang 'big' package function:
//
//			(*Rat) FloatString
//
//		The rounding algorithm used is described in the
//		big.Rat documentation as:
//
//			The last digit is rounded to nearest, with
//			halves rounded away from zero.
//				https://pkg.go.dev/math/big#Rat
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values contained in the current
//	instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numericValue				interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//	     	float32, float64, big.Float
//			*float32, *float64, *big.Float
//			*BigFloatDto, BigFloatDto
//			*big.Rat, big.Rat
//			int8, int16, int, int32, int64, big.Int
//			*int8, *int16, *int, *int32, *int64,
//			*big.Int, big.Int
//			uint8, uint16, uint, uint32, uint64
//			*uint8, *uint16, *uint, *uint32, *uint64
//			*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//			*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//			*NumberStrKernel, NumberStrKernel
//
//		This numeric value will be used to reconfigure
//		the current instance of NumberStrKernel.
//
//		If 'numericValue' is NOT convertible to one of
//		the types listed above, an error will be
//		returned.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		reconfigured instance of NumberStrKernel.
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		This parameter is ignored if 'roundingType' is
//		set to 'NumRoundType.NoRounding()'.
//
//		This parameter will be used to control the number
//		of fractional digits to the right of the decimal
//		point in the reconfigured instance of
//		NumberStrKernel.
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
func (numStrKernel *NumberStrKernel) SetFromNumericValue(
	numericValue interface{},
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
			"SetFromNumericValue()",
		"")

	if err != nil {
		return err
	}

	err = new(numberStrKernelMolecule).
		convertNumericValueToKernel(
			numStrKernel,
			numericValue,
			ePrefix.XCpy(
				"numStrKernel"))

	if err != nil ||
		roundingType == NumRoundType.NoRounding() {

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

// SetFromPureNumberStr
//
// Receives an instance of NumberStrKernel and proceeds
// to reconfigure that instance based on the numeric
// value extracted from a Pure Number String.
//
// A Pure Number String differs from a Native Number
// String in that it offers more options for
// customization. Pure Number Strings are better able to
// match multinational and multicultural number
// formatting conventions. Users have the option to
// specify custom radix points or decimal separator
// characters as well as designating leading or trailing
// minus signs for negative numbers.
//
// A Pure Number String is defined as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters (0-9).
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point or decimal
//     separator. This could be, but is not limited to,
//     a decimal point ('.'). For example, many European
//     countries use the comma (',') as a radix point.
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
//  5. A pure number string will NEVER include currency
//     symbols.
//
// Users also have the option of specifying rounding
// parameters which will be used to round the numeric
// value extracted from 'pureNumberStr' before it is
// converted and stored in the reconfigured current
// instance of 'NumberStrKernel'.
//
// Note that return parameter 'pureNumStrComponents'
// will provide a statistical profile of the new
// returned instance of NumberStrKernel,
// 'newNumStrKernel'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values contained in the current
//	instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberStr				string
//
//		This strings contains the numeric character
//		digits from	which a numeric value will be
//		extracted. This numeric value will be used
//		to reconfigure the instance of NumberStrKernel
//		passed as input parameter 'numStrKernel'.
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
//		The Decimal Separator is used to identify and
//		separate integer and fractional digits while
//		parsing the Pure Number String.
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
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		new returned instance of NumberStrKernel
//		(newNumStrKernel).
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//	pureNumStrComponents		PureNumberStrComponents
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		PureNumberStrComponents. This data structure
//		contains an analysis and profile information on
//		the reconfigured instance of NumberStrKernel
//		passed as input parameter 'numStrKernel'.
//
//		type PureNumberStrComponents struct {
//
//			NumStrStats NumberStrStatsDto
//
//				This data transfer object will return key
//				statistics on the numeric value encapsulated
//				by the current instance of NumberStrKernel.
//
//				type NumberStrStatsDto struct {
//
//				NumOfIntegerDigits					uint64
//
//					The total number of integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfSignificantIntegerDigits		uint64
//
//					The number of nonzero integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfFractionalDigits				uint64
//
//					The total number of fractional digits to the
//					right of the radix point or, decimal point,
//					in the subject numeric value.
//
//				NumOfSignificantFractionalDigits	uint64
//
//					The number of nonzero fractional digits to
//					the right of the radix point or, decimal
//					point, in the subject numeric value.
//
//				NumberValueType 					NumericValueType
//
//					This enumeration value specifies whether the
//					subject numeric value is classified either as
//					an integer or a floating point number.
//
//					Possible enumeration values are listed as
//					follows:
//						NumValType.None()
//						NumValType.FloatingPoint()
//						NumValType.Integer()
//
//				NumberSign							NumericSignValueType
//
//					An enumeration specifying the number sign
//					associated with the numeric value. Possible
//					values are listed as follows:
//						NumSignVal.None()		= Invalid Value
//						NumSignVal.Negative()	= -1
//						NumSignVal.Zero()		=  0
//						NumSignVal.Positive()	=  1
//
//				IsZeroValue							bool
//
//					If 'true', the subject numeric value is equal
//					to zero ('0').
//
//					If 'false', the subject numeric value is
//					greater than or less than zero ('0').
//				}
//
//
//
//			AbsoluteValueNumStr string
//
//			The number string expressed as an absolute value.
//			Be advised, this number string may be a floating
//			point number string containing fractional digits.
//
//			AbsoluteValAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			This string DOES NOT contain a leading number
//			sign (a.k.a. minus sign ('-')
//
//			SignedAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			If the numeric value is negative, a leading
//			minus sign will be prefixed at the beginning
//			of the number string.
//
//			NativeNumberStr string
//
//			A Native Number String representing the base
//			numeric value used to generate these profile
//			number string statistics.
//
//			A valid Native Number String must conform to the
//			standardized formatting criteria defined below:
//
//			 	1. A Native Number String Consists of numeric
//			 	   character digits zero through nine inclusive
//			 	   (0-9).
//
//			 	2. A Native Number String will include a period
//			 	   or decimal point ('.') to separate integer and
//			 	   fractional digits within a number string.
//
//			 	   Native Number String Floating Point Value:
//			 	   				123.1234
//
//			 	3. A Native Number String will always format
//			 	   negative numeric values with a leading minus sign
//			 	   ('-').
//
//			 	   Native Number String Negative Value:
//			 	   				-123.2
//
//			 	4. A Native Number String WILL NEVER include integer
//			 	   separators such as commas (',') to separate
//			 	   integer digits by thousands.
//
//			 	   					NOT THIS: 1,000,000
//			 	   		Native Number String: 1000000
//
//			 	5. Native Number Strings will only consist of:
//
//			 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//			 	   (b)	A decimal point ('.') for floating point
//			 	   		numbers.
//
//			 	   (c)	A leading minus sign ('-') in the case of
//			 	   		negative numeric values.
//
//		}
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernel *NumberStrKernel) SetFromPureNumberStr(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	pureNumStrComponents PureNumberStrComponents,
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
			"SetFromNativeNumberStr()",
		"")

	if err != nil {
		return pureNumStrComponents, err
	}

	pureNumStrComponents,
		err = new(numberStrKernelMechanics).
		setNumStrKernelFromRoundedPureNumStr(
			numStrKernel,
			pureNumberStr,
			decSeparatorChars,
			leadingMinusSign,
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy(
				"numStrKernel"))

	return pureNumStrComponents, err
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
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumberStrKernel.
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
		numStrKernel,
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
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumberStrKernel.
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
		numStrKernel,
		integerDigits,
		fractionalDigits,
		numberSign,
		ePrefix)
}

//	SetStringDigits
//
//	Deletes resets the internal values for the current
//	instance of	NumberStrKernel using input parameters
//	containing integer digit and fractional digit
//	strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumberStrKernel.
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
		numStrKernel,
		[]rune(integerDigits),
		[]rune(fractionalDigits),
		numberSign,
		ePrefix)
}

// String
//
// Returns a formatted number string using the numeric
// value and the Default Number String Format
// Specification configured for the current instance of
// NumberStrKernel.
//
// To explicitly set the Default Number String Format
// Specification, use the following methods:
//
//	NumberStrKernel.SetDefaultNumStrFmtSpec()
//	NumberStrKernel.SetDefaultNumStrFmtSpecPure()
//	NumberStrKernel.SetDefaultSignedNumStrFmtSpecSimple()
//
// If the Default Number String Format is uninitialized
// or invalid, it will be automatically set and defaulted
// to the US (United States) Signed Number String Format
// Specification.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	No rounding will be performed on the numeric
//		value formatted as a returned number string. If
//		rounding is required in conjunction with the
//		Default Format Number String Formatting
//		Specification, the user has two options.
//
//		a.	Call method NumberStrKernel.Round() and
//			then call this method, NumberStrKernel.String().
//
//		b.	Do not call this method. Instead, call
//			method NumberStrKernel.FmtNumStrDefaultRound().
//
//	(2)	This method will NOT delete or modify data values
//		contained in the current instance of
//		NumberStrKernel. However, if the Default
//		NumberStrKernel Number String Format
//		Specification is invalid, internal member
//		variable 'NumberStrKernel.numStrFormatSpec' will
//		be set to the default US (United States) Signed
//		Number String Format Specification.
//
//		All other NumberStrKernel data values will remain
//		unchanged.
//
//	(3)	If errors are encountered, the returned string
//		will be configured with an appropriate error
//		message.
//
//	(4)	This method is very similar to method:
//
//			NumberStrKernel.FmtNumStrDefault()
//
//		The only difference is that in case of error,
//		this method DOES NOT return a type 'error'.
//		Instead, the error message is configured in the
//		returned 'string'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, a
//		formatted Number String will be returned.
//
//		The Number String format used in configuring
//		this Number String is taken from the Default
//		Number String Formatting Specification previously
//		configured in the current instance of
//		NumberStrKernel.
//
//		If errors are encountered, this returned string
//		will be configured with an appropriate error
//		message.
func (numStrKernel *NumberStrKernel) String() string {

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
		nil,
		"NumberStrKernel."+
			"String()",
		"")

	if err != nil {

		numStr = err.Error()

		return numStr
	}

	// If numStrKernel.numStrFormatSpec
	// is a NOP, it will be set to default
	// US Signed Number Minus Format Spec
	_,
		err = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err != nil {

		numStr = err.Error()
		return numStr
	}

	var roundingSpec NumStrRoundingSpec

	roundingSpec,
		err = new(NumStrRoundingSpec).
		NewRoundingSpec(
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy("roundingSpec"))

	if err != nil {

		numStr = err.Error()
		return numStr
	}

	var tempNumStrFormatSpec NumStrFormatSpec

	err = tempNumStrFormatSpec.CopyIn(
		&numStrKernel.numStrFormatSpec,
		ePrefix)

	if err != nil {

		numStr = err.Error()
		return numStr
	}

	numStr,
		err = new(numberStrKernelMolecule).formatNumStr(
		numStrKernel,
		roundingSpec,
		numStrKernel.numStrFormatSpec,
		ePrefix.XCpy("numStrKernel"))

	if err != nil {

		numStr = err.Error()
	}

	return numStr
}
