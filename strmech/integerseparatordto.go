package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	IntegerSeparatorDto
//
//	Type IntegerSeparatorDto is designed to manage integer
//	separators, primarily thousands separators, for
//	different countries and cultures. The term
//	'integer separators' is used because this type manages
//	both integer grouping and the characters used to
//	separate integer groups.
//
//	In the USA and many other countries, integer numbers
//	are often separated by commas thereby grouping the
//	number into thousands.
//
//		Example: 1,000,000,000
//
//	Other countries use characters other than the comma to
//	separate integers into thousands. Some countries and
//	cultures do not use thousands separation and instead
//	rely on multiple integer separation characters and
//	grouping sequences for a single integer number. Notable
//	examples of this are found in the 'India Number System'
//	and 'Chinese Numerals'.
//
//	Reference:
//		https://en.wikipedia.org/wiki/Indian_numbering_system
//		https://en.wikipedia.org/wiki/Chinese_numerals
//		https://en.wikipedia.org/wiki/Decimal_separator
//
//	The IntegerSeparatorDto type provides the flexibility
//	necessary to process these complex number separation
//	formats.
//
// ----------------------------------------------------------------
//
// Member Variables
//
//	intSeparatorChars			[]rune
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A 'thousands
//		separator' is used to separate groups of integer digits to
//		the left of the decimal separator (a.k.a. decimal point).
//		In the United States, the standard integer digits
//		separator is the single comma character (',').
//		      United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		      European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//	intGroupingSequence			[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value is
//		used to group integers within a number string.
//
//		In most western countries, integer digits to the left of
//		the decimal separator (a.k.a. decimal point) are separated
//		into groups of three digits representing a grouping of
//		'thousands' like this: '1,000,000,000'. In this case the
//		'intGroupingSequence' value would be set to three
//		(uint[]{3}).
//
//		In some countries and cultures other integer groupings are
//		used. In India, for example, a number might be formatted
//		like this: '6,78,90,00,00,00,00,000'. In this case the
//		'intGroupingSequence' value would be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of four
//		and are formatted like this: '12,3456,7890,2345'. In this
//		case the 'intGroupingSequence' value would be set to four
//		uint[]{4}.
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intGroupingSequence' array after completion of the
//		last group in the 'intGroupingSequence' array.
//		 Example:
//		   restartIntGroupingSequence = 'true'
//		   intGroupingSequence = uint{3,2}
//		   integer = 1234567890123456
//		   result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intGroupingSequence' array will simply be
//		repeated for all the remaining integer digits.
//		 Example:
//		   restartIntGroupingSequence = 'false'
//		   intGroupingSequence = uint{3,2}
//		   integer = 1234567890123456
//		   result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly rare.
//		For the vast majority of integer separation scenarios, set
//		this parameter should be set to 'false'.
type IntegerSeparatorDto struct {
	intSeparatorChars []rune
	// A series of runes used to separate
	// integer digits.

	intGroupingSequence []uint
	// Number of integer digits in each group

	restartIntGroupingSequence bool
	// If 'true', the entire grouping sequence is
	// repeated beginning at array index zero.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// IntegerSeparatorDto instance to the data fields of the current
// IntegerSeparatorDto instance.
//
// If input parameter 'incomingNStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
// Be advised, all the data fields in the current
// IntegerSeparatorDto instance will be overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		incomingNStrIntSeparator     *IntegerSeparatorDto
//		   - A pointer to an instance of IntegerSeparatorDto.
//		     The data values in this object will be copied to the
//		     current IntegerSeparatorDto instance.
//
//		     If input parameter 'incomingNStrIntSeparator' is judged
//		     to be invalid, this method will return an error.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                 error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) CopyIn(
	incomingNStrIntSeparator *IntegerSeparatorDto,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoMolecule).
		copyIntSepDto(
			nStrIntSep,
			incomingNStrIntSeparator,
			ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// IntegerSeparatorDto instance.
//
// If the current IntegerSeparatorDto instance is judged to be
// invalid, this method will return an error.
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
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto will be created and returned containing
//	     all the data values copied from the current instance of
//	     IntegerSeparatorDto.
//
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyIntSepDto IntegerSeparatorDto,
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.CopyOut()",
		"")

	if err != nil {
		return deepCopyIntSepDto, err
	}

	err = new(integerSeparatorDtoMolecule).
		copyIntSepDto(
			&deepCopyIntSepDto,
			nStrIntSep,
			ePrefix.XCpy(
				"deepCopyIntSepDto<-nStrIntSep"))

	return deepCopyIntSepDto, err
}

// Empty - Resets all internal member variables for the current
// instance of IntegerSeparatorDto to their zero or uninitialized
// states. This method will leave the current instance of
// IntegerSeparatorDto in an invalid state and unavailable for
// immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of IntegerSeparatorDto. All member variable
// data values will be reset to their zero or uninitialized
// states.
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
func (nStrIntSep *IntegerSeparatorDto) Empty() {
	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	_ = new(integerSeparatorDtoQuark).
		empty(
			nStrIntSep,
			nil)

	nStrIntSep.lock.Unlock()

	nStrIntSep.lock = nil

	return
}

// Equal - Receives a pointer to an IntegerSeparatorDto object
// and proceeds to determine whether all data elements in this
// object are equal to all corresponding data elements in the
// current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		nStrIntSepDto              *IntegerSeparatorDto
//	    - This method will compare all data elements in the
//	      current IntegerSeparatorDto object to corresponding
//	      data elements in this second IntegerSeparatorDto
//	      object in order to determine equivalency.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If all the data elements in the current IntegerSeparatorDto
//	     instance are equal to all the corresponding data elements
//	     in 'nStrIntSepDto', this return parameter will be set to
//	     'true'. If all the data elements are NOT equal, this
//	     return parameter will be set to 'false'.
//
//
//	error
//	   - If all the data elements in the current IntegerSeparatorDto
//	     are equal to all the corresponding data elements in
//	     'nStrIntSepDto', this return parameter will be set to
//	     'nil'.
//
//	     If the corresponding data elements are not equal, a
//	     detailed error message identifying the unequal elements
//	     will be returned.
//
//	     If errors are encountered during processing, the returned
//	     error message will incorporate the method chain  and text
//	     passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) Equal(
	nStrIntSepDto *IntegerSeparatorDto,
	errorPrefix interface{}) (
	bool,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.Equal())",
		"")

	if err != nil {
		return false, err
	}

	return new(integerSeparatorDtoMolecule).
		equal(
			nStrIntSep,
			nStrIntSepDto,
			ePrefix.XCpy(
				"nStrIntSep vs nStrIntSepDto"))
}

// GetIntSeparatorChars - Returns the integer digit separator
// characters for the current IntegerSeparatorDto instance as an
// array of runes.
//
// If the rune array is zero length, an error will be returned.
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
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// -----------------------------------------------------------------
//
// Return Values
//
//	[] rune
//	   - If this method completes successfully, an array of runes
//	     will be returned representing the integer digit separator
//	     characters.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) GetIntSeparatorChars(
	errorPrefix interface{}) ([]rune, error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.GetIntSeparatorChars()",
		"")

	if err != nil {
		return []rune{}, err
	}

	if nStrIntSep.intSeparatorChars == nil {
		nStrIntSep.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSepChars := len(nStrIntSep.intSeparatorChars)

	if lenIntSepChars == 0 {
		return make([]rune, 0, 5),
			fmt.Errorf("%v\n"+
				"Error: The internal 'intSeparatorChars' array is invalid!\n"+
				"'nStrIntSep.intSeparatorChars' is a zero length array.\n",
				ePrefix.String())
	}

	newIntSepChars :=
		make([]rune, lenIntSepChars, lenIntSepChars+5)

	for i := 0; i < lenIntSepChars; i++ {
		newIntSepChars[i] =
			nStrIntSep.intSeparatorChars[i]
	}

	return newIntSepChars, nil
}

// GetIntegerGroupingSequence - Returns an unsigned integer
// representing the number of integer digits included in an integer
// group for the current IntegerSeparatorDto instance.
//
// This unsigned integer value is used to group integers within a
// number string.
//
// In most western countries, integer digits to the left of
// the decimal separator (a.k.a. decimal point) are separated
// into groups of three digits representing a grouping of
// 'thousands' like this: '1,000,000,000'. In this case the
// intGroupingSequence value would be set to three ('3').
//
// In some countries and cultures other integer groupings are
// used. In India, for example, a number might be formatted
// like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
// would be formatted like this: 12,3456,7890,2345
func (nStrIntSep *IntegerSeparatorDto) GetIntegerGroupingSequence() []uint {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var newRuneArray []uint = nil

	_ = new(strMechPreon).
		copyUnsignedIntArrays(
			&newRuneArray,
			&nStrIntSep.intGroupingSequence,
			true,
			nil)

	return newRuneArray
}

//	GetFormattedIntegerNumStr
//
//	This method receives a pure number string
//	consisting of numeric digit characters zero
//	('0') through nine ('9'), inclusive.
//
//	These numeric character digits are then
//	grouped and separated using the Integer
//	Separation parameters provided by the
//	current instance of IntegerSeparatorDto.
//
//	Example:
//		     Pure Number String: 123456
//		Formatted Number String: 123,456
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumString				string
//
//		A string of numeric digits where all the characters
//		fall within the range of zero ('0') through nine
//		('9'), inclusive.
//
//		These numeric character digits are then	grouped and
//		separated using the Integer Separation parameters
//		provided by the	current instance of
//		IntegerSeparatorDto.
//
//		If this string contains any non-numeric characters,
//		an error will be returned.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorDto
//
//		If this method completes successfully, a new instance
//		of IntegerSeparatorDto, configured with Chinese
//		Numbering System integer grouping, will be created and
//		returned.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errorPrefix'. The 'errorPrefix' text will be attached
//		to the beginning of	the error message.
func (nStrIntSep *IntegerSeparatorDto) GetFormattedIntegerNumStr(
	pureNumString string,
	errorPrefix interface{}) (
	string,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"GetFormattedIntegerNumStr()",
		"")

	if err != nil {
		return "", err
	}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = new(integerSeparatorDtoMolecule).applyIntSeparators(
		nStrIntSep,
		[]rune(pureNumString),
		ePrefix.XCpy(
			"numStrWithIntSeps<-pureNumString"))

	return string(numStrWithIntSeps), err
}

// GetRestartIntGroupingSequence
//
// Returns the internal member variable
// 'restartIntGroupingSequence' for the current
// IntegerSeparatorDto instance.
//
// The IntegerSeparatorDto type is intended to
// be configured in an array of IntegerSeparatorDto
// objects which, taken as a whole, provides
// formatting specifications for complex integer
// group separation operations.
//
// If the current IntegerSeparatorDto is the last
// element in an array of IntegerSeparatorDto objects,
// the 'Restart Integer Grouping Sequence' flag
// signals whether the integer separation operation
// will be restarted from the first IntegerSeparatorDto
// object in the array.
//
// Again, the IntegerSeparatorDto.restartIntGroupingSequence
// boolean flag only has meaning if the current
// IntegerSeparatorDto object is last element in an array of
// IntegerSeparatorDto objects.
func (nStrIntSep *IntegerSeparatorDto) GetRestartIntGroupingSequence() bool {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return nStrIntSep.restartIntGroupingSequence
}

// IsValidInstance - Performs a diagnostic review of the current
// IntegerSeparatorDto instance to determine whether the current
// instance is valid in all respects.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//	isValid             bool
//	   - This returned boolean value will signal whether the
//	     current IntegerSeparatorDto is valid, or not. If the
//	     current IntegerSeparatorDto contains valid data, this
//	     method returns 'true'. If the data is invalid, this method
//	     will return 'false'.
func (nStrIntSep *IntegerSeparatorDto) IsValidInstance() (
	isValid bool) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	isValid,
		_ =
		new(integerSeparatorDtoQuark).
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current IntegerSeparatorDto instance to determine whether the
// current instance is valid in all respects.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
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
//	err                 error
//	   - If the current instance of IntegerSeparatorDto contains
//	     invalid data, a detailed error message will be returned
//	     identifying the invalid data item.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
//
//	     If the current instance is valid, this error parameter
//	     will be set to nil.
func (nStrIntSep *IntegerSeparatorDto) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err =
		new(integerSeparatorDtoQuark).
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				ePrefix)

	return err
}

// NewFromIntGroupEnum - Creates a new instance of
// IntegerSeparatorDto based on an enumeration value passed
// through input parameter 'intGroupingSpec'.
//
// 'intGroupingType' is of type IntegerGroupingType and
// must be set to one of the following valid values:
//
//	IntGroupingType.None()
//	IntGroupingType.Thousands()
//	IntGroupingType.IndiaNumbering()
//	IntGroupingType.ChineseNumbering()
//
// If 'intGroupingType' is set to 'IntGroupingType.None()',
// an empty instance of IntegerSeparatorDto will be returned.
//
// The returned instance of IntegerSeparatorDto will be
// configured for one of these four Integer Grouping
// Type values.
//
// The character or characters separating the integer
// groups is supplied by input parameter
// 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType				IntegerGroupingType
//		This instance of IntegerGroupingType enumeration defines
//		how 'intSep' will be configured for integer grouping.
//		The enumeration value 'intGroupingSpec' must be set to
//		one of these four values:
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorDto.
//
//	intSeparatorChars			string
//		One or more characters used to separate groups of
//		integers. This separator is also known as the 'thousands'
//		separator. It is used to separate groups of integer digits
//		to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the standard
//		integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length string, an
//		error will be returned.
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
//		collection of error prefix and error context
//		information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		from this object will be copied for use in error and
//		informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		Information from this object will be copied for use in
//		error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		a two-dimensional slice of strings containing error
//		prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorDto
//		If this method completes successfully, a new instance of
//		IntegerSeparatorDto will be created and returned. This
//		IntegerSeparatorDto will be configured to 'Thousands',
//		'India Numbering System' or 'Chinese Numbering' depending
//		on the specification provided by input parameter,
//		'intGroupingType'.
//
//		If input parameter 'intGroupingType' is set to
//		'IntGroupingType.None()', an empty instance of
//		IntegerSeparatorDto will be returned.
//
//	error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntSep *IntegerSeparatorDto) NewFromIntGroupEnum(
	intGroupingType IntegerGroupingType,
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSepDto := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewFromIntGroupEnum()",
		"")

	if err != nil {
		return newIntSepDto, err
	}

	err = new(integerSeparatorDtoHelper).
		setFromIntGroupEnum(
			&newIntSepDto,
			intGroupingType,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"newIntSepDto<-"))

	return newIntSepDto, err
}

// NewFromIntGroupEnumRunes - Creates a new instance of
// IntegerSeparatorDto based on an enumeration value passed
// through input parameter 'intGroupingSpec'.
//
// 'intGroupingType' is of type IntegerGroupingType and
// must be set to one of the following valid values:
//
//	IntGroupingType.None()
//	IntGroupingType.Thousands()
//	IntGroupingType.IndiaNumbering()
//	IntGroupingType.ChineseNumbering()
//
// If 'intGroupingType' is set to 'IntGroupingType.None()',
// an empty instance of IntegerSeparatorDto will be returned.
//
// The instance of IntegerSeparatorDto will be configured
// for one of these four Integer Grouping Type values.
//
// The character or characters separating the integer
// groups is supplied by rune array input parameter
// 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType				IntegerGroupingType
//		This instance of IntegerGroupingType enumeration defines
//		how 'intSep' will be configured for integer grouping.
//		The enumeration value 'intGroupingSpec' must be set to
//		one of these four values:
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorDto.
//
//	intSeparatorChars			[]rune
//		One or more characters used to separate groups of
//		integers. It is used to separate groups of integer digits
//		to the left of the decimal separator (a.k.a. decimal
//		point). In the United States, the standard integer digits
//		separator is the comma (',').
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length array, an
//		error will be returned.
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
//		collection of error prefix and error context
//		information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		from this object will be copied for use in error and
//		informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		Information from this object will be copied for use in
//		error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		a two-dimensional slice of strings containing error
//		prefix and error context information.
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
//	IntegerSeparatorDto
//		If this method completes successfully, a new instance of
//		IntegerSeparatorDto will be created and returned. This
//		IntegerSeparatorDto will be configured to 'Thousands',
//		'India Numbering System' or 'Chinese Numbering' depending
//		on the specification provided by input parameter,
//		'intGroupingType'.
//
//		If input parameter 'intGroupingType' is set to
//		'IntGroupingType.None()', an empty instance of
//		IntegerSeparatorDto will be returned.
//
//	error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntSep *IntegerSeparatorDto) NewFromIntGroupEnumRunes(
	intGroupingType IntegerGroupingType,
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSepDto := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewFromIntGroupEnumRunes()",
		"")

	if err != nil {
		return newIntSepDto, err
	}

	err = new(integerSeparatorDtoHelper).
		setFromIntGroupEnum(
			&newIntSepDto,
			intGroupingType,
			intSeparatorChars,
			ePrefix.XCpy(
				"newIntSepDto<-"))

	return newIntSepDto, err
}

// NewDetail - Creates and returns a new instance of
// IntegerSeparatorDto. The new instance is generated based on
// component elements passed as input parameters.
//
// This method differs from IntegerSeparatorDto.NewDetailRunes() in
// that this method accepts a string for input parameter
// 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          string
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	intGroupingSequence        []uint
//	   - This unsigned integer array specifies the number of
//	     integer digits within a series of groups. This value is
//	     used to group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     'intGroupingSequence' value would be set to three
//	     (uint[]{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. In this case the
//	     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//	     Chinese Numerals have an integer grouping value of four
//	     and are formatted like this: '12,3456,7890,2345'. In this
//	     case the 'intGroupingSequence' value would be set to four
//	     uint[]{4}.
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	newIntSep                  IntegerSeparatorDto
//	   - If this method completes successfully, the parameter will
//	     return a new and fully populated instance of
//	     IntegerSeparatorDto.
//
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewDetail(
	intSeparatorChars string,
	intSeparatorGrouping []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep IntegerSeparatorDto,
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.NewDetail()",
		"")

	if err != nil {
		return newIntSep, err
	}

	if len(intSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())

		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				&newIntSep,
				[]rune(intSeparatorChars),
				intSeparatorGrouping,
				restartIntGroupingSequence,
				ePrefix.XCpy("newIntSep"))

	return newIntSep, err
}

// NewDetailRunes - Creates and returns a new instance of
// IntegerSeparatorDto. The new instance is generated based on
// component elements passed as input parameters.
//
// This method differs from IntegerSeparatorDto.NewDetail() in that
// this method accepts an array of runes for 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          []rune
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length array, an
//	     error will be returned.
//
//
//	intGroupingSequence        []uint
//	   - This unsigned integer array specifies the number of
//	     integer digits within a series of groups. This value is
//	     used to group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     'intGroupingSequence' value would be set to three
//	     (uint[]{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. In this case the
//	     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//	     Chinese Numerals have an integer grouping value of four
//	     and are formatted like this: '12,3456,7890,2345'. In this
//	     case the 'intGroupingSequence' value would be set to four
//	     uint[]{4}.
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	newIntSep                  IntegerSeparatorDto
//	   - If this method completes successfully, the parameter will
//	     return a new and fully populated instance of
//	     IntegerSeparatorDto.
//
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewDetailRunes(
	intSeparatorChars []rune,
	intSeparatorGrouping []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep IntegerSeparatorDto,
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.NewDetail()",
		"")

	if err != nil {
		return newIntSep, err
	}

	if len(intSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())

		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				&newIntSep,
				intSeparatorChars,
				intSeparatorGrouping,
				restartIntGroupingSequence,
				ePrefix.XCpy("newIntSep"))

	return newIntSep, err
}

//	NewChineseNumbering
//
//	Returns a new instance of IntegerSeparatorDto
//	configured for the Chinese Numbering System
//	integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is a
//	string containing the integer separator
//	character or characters which will separate
//	the integer digit groups.
//
//	In conformance with the Chinese Numbering
//	System, the integer digit grouping is defaulted
//	to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,000,0000,0000'
//
//	Although, users have the option of specifying one
//	or more custom integer separator characters passed
//	through input parameter 'intSeparatorChars', be
//	advised that the commonly used integer separator
//	for the Chinese Numbering system is the comma (",")
//	character.
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups
//		of integers. These characters are used to separate
//		groups of integer digits to the left of the decimal
//		separator (a.k.a. decimal point). The most common
//		integer separator used in conjunction with the
//		Chinese Numbering System is	the comma (",").
//
//			Example:  6,7890,000,0000,0000
//
//		If this input parameter contains a zero length string, an
//		error will be returned.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorDto
//
//		If this method completes successfully, a new instance of
//		IntegerSeparatorDto, configured with Chinese Numbering
//		System integer grouping, will be created and returned.
//
//	error
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
func (nStrIntSep *IntegerSeparatorDto) NewChineseNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewChineseNumbering()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setChineseNumberingRunes(
				&newIntSep,
				[]rune(intSeparatorChars),
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewChineseNumberingDefault - Returns a new instance
// of IntegerSeparatorDto configured for the Chinese
// Numbering System integer digit grouping.
//
// The integer separator character will be defaulted to
// the comma (',') character.
//
// In conformance with the Chinese Numbering System, the
// integer digit grouping is defaulted to a value of four
// (4).
//
// This means that integer digit grouping will consist of
// groups of four (4) integer digits.
//
//	Example: '6,7890,000,0000,0000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
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
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto, configured with Chinese Numbering
//	     System integer grouping, will be created and returned.
//	     The integer separator character for the new returned
//	     instance will be defaulted to the comma (',') character.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewChineseNumberingDefault(
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewChineseNumberingDefault()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setChineseNumberingRunes(
				&newIntSep,
				[]rune{','},
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewChineseNumberingRunes - Returns a new instance of
// IntegerSeparatorDto configured for the Chinese
// Numbering System integer digit grouping.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing the integer separator character or
// characters used to separate the integer digit groups.
//
// In conformance with the Chinese Numbering System, the
// integer digit grouping is defaulted to a value of four
// (4).
//
// This means that integer digit grouping will consist of
// groups of four (4) integer digits.
//
//	Example: '6,7890,000,0000,0000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the Chinese Numbering system is
// the comma ([]rune{','}) character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          []rune
//		   - One or more characters used to separate groups of
//		     integers. These characters are used to separate groups
//		     of integer digits to the left of the decimal separator
//		     (a.k.a. decimal point). The most common integer separator
//		     used in conjunction with the Chinese Numbering System is
//		     the comma ([]rune{,}).
//		       Example:  6,7890,000,0000,0000
//
//		     If this input parameter contains a zero length array, an
//		     error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto, configured with Chinese Numbering
//	     System integer grouping, will be created and returned.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewChineseNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewChineseNumberingRunes()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setChineseNumberingRunes(
				&newIntSep,
				intSeparatorChars,
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewIndiaNumbering - Returns a new instance of
// IntegerSeparatorDto configured for the India
// Numbering System integer digit grouping.
//
// The input parameter 'intSeparatorChars' is a string
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the India Numbering system is
// the comma (",") character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          string
//		   - One or more characters used to separate groups of
//		     integers. These characters are used to separate groups
//		     of integer digits to the left of the decimal separator
//		     (a.k.a. decimal point). The most common integer separator
//		     used in conjunction with the India Numbering System is
//		     the comma (",").
//		       Example:  6,78,90,00,00,00,00,000
//
//		     If this input parameter contains a zero length string, an
//		     error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto, configured with India Numbering
//	     System integer grouping, will be created and returned.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewIndiaNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewIndiaNumbering()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setIndiaNumberingRunes(
				&newIntSep,
				[]rune(intSeparatorChars),
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewIndiaNumberingDefault - Returns a new instance of
// IntegerSeparatorDto configured for the India
// Numbering System integer digit grouping.
//
// The integer separator character will be defaulted to a
// comma (',') character.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
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
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto, configured with India Numbering
//	     System integer grouping, will be created and returned.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewIndiaNumberingDefault(
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewIndiaNumberingDefault()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setIndiaNumberingRunes(
				&newIntSep,
				[]rune{','},
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewIndiaNumberingRunes - Returns a new instance of
// IntegerSeparatorDto configured for the India
// Numbering System integer digit grouping.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing one or more integer separator characters
// used to separate integer digit groups.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the India Numbering system is
// the comma (",") character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          []rune
//		   - One or more characters used to separate groups of
//		     integers. This string is used to separate groups of
//		     integer digits to the left of the decimal separator
//		     (a.k.a. decimal point). The most common integer separator
//		     character used in conjunction with the India Numbering
//		     System is the comma  ([]rune{','}).
//		       Example:  6,78,90,00,00,00,00,000
//
//		     If this input parameter contains a zero length array,
//		     an error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto, configured with India Numbering
//	     System integer grouping, will be created and returned.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewIndiaNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewIndiaNumberingRunes()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoUtility).
			setIndiaNumberingRunes(
				&newIntSep,
				intSeparatorChars,
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

//	NewNoIntegerSeparation
//
//	Returns an instance of IntegerSeparatorDto configured
//	for NO integer separation.
//
//	Call this method if you want to guarantee that integer
//	separation WILL NOT BE APPLIED in the formatting of
//	Number Strings.
//
//	Example With Integer Separation
//
//		1,234,567,890
//
//	Example With NO Integer Separation
//
//		1234567890
//
//	This method will ensure the integer separation WILL
//	NOT BE APPLIED in the formatting of	Number Strings.
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
//	IntegerSeparatorDto
//
//		If this method completes successfully, new instance
//		of IntegerSeparatorDto will be returned. This new
//		instance will be configured to ensure that integer
//		separation WILL NOT BE APPLIED in the formatting of
//		Number Strings.
func (nStrIntSep *IntegerSeparatorDto) NewNoIntegerSeparation() IntegerSeparatorDto {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	newIntSepDto := IntegerSeparatorDto{}

	_ = new(integerSeparatorDtoQuark).empty(
		&newIntSepDto,
		nil)

	return newIntSepDto
}

// NewThousands - Creates and returns a new instance of
// IntegerSeparatorDto. The returned IntegerSeparatorDto instance
// represents a basic or simple integer separator object using
// default values and a minimum number of input parameters.
//
// The input parameter 'intSeparatorChars' is string
// containing the integer separator characters. The integer digit
// grouping is defaulted to a value of three (3). The 'separator
// repetitions' value is defaulted to zero (0) signaling unlimited
// repetitions.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          string
//	   - One or more characters used to separate groups of
//	     integers. This separator is also known as the 'thousands'
//	     separator. It is used to separate groups of integer digits
//	     to the left of the decimal separator
//	     (a.k.a. decimal point). In the United States, the standard
//	     integer digits separator is the comma (",").
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this input parameter contains a zero length string, an
//	     error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto will be created and returned. The
//	     'integer digits grouping sequence' will be automatically
//	     set to a default value of 3-digits.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewThousands(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.NewThousands()",
		"")

	if err != nil {
		return IntegerSeparatorDto{}, err
	}

	ePrefix.SetEPref(
		"IntegerSeparatorDto.NewThousands()")

	newIntSep := IntegerSeparatorDto{}

	err =
		new(integerSeparatorDtoUtility).
			setThousandsRunes(
				&newIntSep,
				[]rune(intSeparatorChars),
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewThousandsRunes - Creates and returns a new instance of
// IntegerSeparatorDto. The returned IntegerSeparatorDto instance
// represents a basic or simple integer separator object using
// default values and a minimum number of input parameters.
//
// This method is an alternative to method
// IntegerSeparatorDto.NewThousands() in that this method accepts
// integer separator characters as an array of runes instead
// of a string.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing the integer separator characters. The integer
// digit grouping is defaulted to a value of three (3). The
// 'separator repetitions' value is defaulted to zero (0) signaling
// unlimited repetitions. The 'restartIntGroupingSequence' flag
// will be defaulted to 'false'.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          []rune
//		   - One or more characters used to separate groups of
//		     integers. This separator is also known as the 'thousands'
//		     separator. It is used to separate groups of integer digits
//		     to the left of the decimal separator (a.k.a. decimal
//		     point). In the United States, the standard integer digits
//		     separator is the comma (',').
//		           United States Example:  1,000,000,000
//
//		     In many European countries, a single period ('.') is used
//		     as the integer separator character.
//		           European Example: 1.000.000.000
//
//		     Other countries and cultures use spaces, apostrophes or
//		     multiple characters to separate integers.
//
//		     If this input parameter contains a zero length array, an
//		     error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto will be created and returned. The
//	     'integer digits grouping sequence' will be automatically
//	     set to a default value of 3-digits.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewThousandsRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.NewThousandsRunes()",
		"")

	if err != nil {
		return IntegerSeparatorDto{}, err
	}

	newIntSep := IntegerSeparatorDto{}

	err =
		new(integerSeparatorDtoUtility).
			setThousandsRunes(
				&newIntSep,
				intSeparatorChars,
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

// NewUnitedStatesDefaults - Creates and returns a new instance of
// IntegerSeparatorDto set to United States default integer
// separators. Integer separator values used in the United States
// consist of the comma character (','), an integer grouping of
// three ('3') and unlimited repetitions of this sequence.
//
//	United States Integer Separation Example:
//	      '1,000,000,000,000'
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
//	IntegerSeparatorDto
//	   - This will return a new and fully populated instance of
//	     IntegerSeparatorDto configured with United States default
//	     integer separator values.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewUnitedStatesDefaults(
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"NewUnitedStatesDefaults()",
		"")

	if err != nil {
		return IntegerSeparatorDto{}, err
	}

	newIntSep := IntegerSeparatorDto{}

	err = new(integerSeparatorDtoMechanics).
		setToUSADefaults(
			&newIntSep,
			ePrefix.XCpy(
				"newIntSep"))

	return newIntSep, err
}

// NewWithComponents - Creates and returns a new instance of
// IntegerSeparatorDto. The new instance is generated based on
// component elements passed as input parameters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          []rune
//	   - A series of runes or characters used to separate integer
//	     digits in a number string. These characters are commonly
//	     known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (','). Other
//	     countries and cultures use periods, spaces, apostrophes or
//	     multiple characters to separate integers.
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length array, an
//	     error will be returned.
//
//
//	intGroupingSequence        []uint
//	   - This unsigned integer array specifies the number of
//	     integer digits within a series of groups. This value is
//	     used to group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     'intGroupingSequence' value would be set to three
//	     (uint[]{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. In this case the
//	     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//	     Chinese Numerals have an integer grouping value of four
//	     and are formatted like this: '12,3456,7890,2345'. In this
//	     case the 'intGroupingSequence' value would be set to four
//	     uint[]{4}.
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
//
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	newIntSep                  IntegerSeparatorDto
//	   - If this method completes successfully, the parameter will
//	     return a new and fully populated instance of
//	     IntegerSeparatorDto.
//
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) NewWithComponents(
	intSeparatorChars []rune,
	intSeparatorGrouping []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep IntegerSeparatorDto,
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.NewWithComponents()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				&newIntSep,
				intSeparatorChars,
				intSeparatorGrouping,
				restartIntGroupingSequence,
				ePrefix.XCpy("newIntSep"))

	return newIntSep, err
}

// SetDetail - This method will overwrite and reset the
// internal member variable data values for the current
// IntegerSeparatorDto instance based on the component elements
// passed as input parameters.
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          string
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//
//	intGroupingSequence        []uint
//	   - This unsigned integer array specifies the number of
//	     integer digits within a series of groups. This value is
//	     used to group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     'intGroupingSequence' value would be set to three
//	     (uint[]{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. In this case the
//	     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//	     Chinese Numerals have an integer grouping value of four
//	     and are formatted like this: '12,3456,7890,2345'. In this
//	     case the 'intGroupingSequence' value would be set to four
//	     uint[]{4}.
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
//
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetDetail(
	intSeparatorChars string,
	intGroupingSequence []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetDetail()",
		"")

	if err != nil {
		return err
	}

	if len(intSeparatorChars) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())
	}

	return new(integerSeparatorDtoMechanics).
		setWithComponents(
			nStrIntSep,
			[]rune(intSeparatorChars),
			intGroupingSequence,
			restartIntGroupingSequence,
			ePrefix.XCpy("nStrIntSep"))
}

// SetDetailRunes - This method will overwrite and reset the
// internal member variable data values for the current
// IntegerSeparatorDto instance based on the component elements
// passed as input parameters.
//
// This method differs from IntegerSeparatorDto.SetDetail() in that
// this method accepts an array of runes for 'intSeparatorChars'.
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          []rune
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length array, an
//	     error will be returned.
//
//
//	intGroupingSequence        []uint
//	   - This unsigned integer array specifies the number of
//	     integer digits within a series of groups. This value is
//	     used to group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     'intGroupingSequence' value would be set to three
//	     (uint[]{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. In this case the
//	     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//	     Chinese Numerals have an integer grouping value of four
//	     and are formatted like this: '12,3456,7890,2345'. In this
//	     case the 'intGroupingSequence' value would be set to four
//	     uint[]{4}.
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
//
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetDetailRunes(
	intSeparatorChars []rune,
	intGroupingSequence []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.SetDetailRunes()",
		"")

	if err != nil {
		return err
	}

	if len(intSeparatorChars) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())
	}

	return new(integerSeparatorDtoMechanics).
		setWithComponents(
			nStrIntSep,
			intSeparatorChars,
			intGroupingSequence,
			restartIntGroupingSequence,
			ePrefix.XCpy("nStrIntSep"))
}

// SetChineseNumbering - Overwrites all the member variable
// data values for current instance of IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with Chinese Numbering System integer
// grouping.
//
// The input parameter 'intSeparatorChars' is a string
// containing the integer separator character or characters
// which will separate the integer digit groups.
//
// In conformance with the Chinese Numbering System, the integer
// digit grouping is defaulted to a value of four (4).
//
// This means that integer digit grouping will consist of
// groups of four (4) integer digits.
//
//	Example: '6,7890,000,0000,0000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the Chinese Numbering system is
// the comma (",") character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//			intSeparatorChars          string
//			   - One or more characters used to separate groups of
//			     integers. These characters are used to separate groups
//			     of integer digits to the left of the decimal separator
//			     (a.k.a. decimal point). The most common integer separator
//			     used in conjunction with the Chinese Numbering System is
//			     the comma (",").
//			       Example:  6,7890,000,0000,0000
//
//			     If this input parameter contains a zero length string, an
//			     error will be returned.
//
//
//	 errorPrefix                interface{}
//			   - This object encapsulates error prefix text which is
//			     included in all returned error messages. Usually, it
//			     contains the name of the calling method or methods
//			     listed as a method or function chain of execution.
//
//			     If no error prefix information is needed, set this parameter
//			     to 'nil'.
//
//			     This empty interface must be convertible to one of the
//			     following types:
//
//			     1. nil - A nil value is valid and generates an empty
//			        collection of error prefix and error context
//			        information.
//
//			     2. string - A string containing error prefix information.
//
//			     3. []string A one-dimensional slice of strings containing
//			        error prefix information
//
//			     4. [][2]string A two-dimensional slice of strings
//			        containing error prefix and error context information.
//
//			     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//			        from this object will be copied for use in error and
//			        informational messages.
//
//			     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//			        Information from this object will be copied for use in
//			        error and informational messages.
//
//			     7. IBasicErrorPrefix - An interface to a method generating
//			        a two-dimensional slice of strings containing error
//			        prefix and error context information.
//
//			     If parameter 'errorPrefix' is NOT convertible to one of
//			     the valid types listed above, it will be considered
//			     invalid and trigger the return of an error.
//
//			     Types ErrPrefixDto and IBasicErrorPrefix are included in
//			     the 'errpref' software package,
//			     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetChineseNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetChineseNumbering()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setChineseNumberingRunes(
			nStrIntSep,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetChineseNumberingDefault - Overwrites all the member
// variable data values for current instance of
// IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with Chinese Numbering System integer
// grouping.
//
// The integer separator character will be defaulted to a
// comma (',') character.
//
// In conformance with the Chinese Numbering System, the integer
// digit grouping is defaulted to a value of four (4).
//
// This means that integer digit grouping will consist of
// groups of four (4) integer digits.
//
//	Example: '6,7890,000,0000,0000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetChineseNumberingDefault(
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetChineseNumberingDefault()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setChineseNumberingRunes(
			nStrIntSep,
			[]rune{','},
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetChineseNumberingRunes - Overwrites all the member variable
// data values for current instance of IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with Chinese Numbering System integer
// digit grouping.
//
// The input parameter 'intSeparatorChars' is an array of runes
// containing the integer separator character or characters used
// to separate the integer digit groups.
//
// In conformance with the Chinese Numbering System, the integer
// digit grouping is defaulted to a value of four (4).
//
// This means that integer digit grouping will consist of
// groups of four (4) integer digits.
//
//	Example: '6,7890,000,0000,0000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the Chinese Numbering system is
// the comma ([]rune{','}) character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars          []rune
//	   - One or more characters used to separate groups of
//	     integers. These characters are used to separate groups
//	     of integer digits to the left of the decimal separator
//	     (a.k.a. decimal point). The most common integer separator
//	     used in conjunction with the Chinese Numbering System is
//	     the comma ([]rune{,}).
//	       Example:  6,7890,000,0000,0000
//
//	     If this input parameter contains a zero length array, an
//	     error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of the
//	     error message.
func (nStrIntSep *IntegerSeparatorDto) SetChineseNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetChineseNumberingRunes()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setChineseNumberingRunes(
			nStrIntSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetIndiaNumbering - Overwrites all the member variable
// data values for the current instance of IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with India Numbering System integer
// digit grouping.
//
// The input parameter 'intSeparatorChars' is a string
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the India Numbering system is
// the comma (",") character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars          string
//	   - One or more characters used to separate groups of
//	     integers. This string is used to separate groups of
//	     integer digits to the left of the decimal separator
//	     (a.k.a. decimal point). The most common integer separator
//	     character used in conjunction with the India Numbering
//	     System is the comma  (",").
//	       Example:  6,78,90,00,00,00,00,000
//
//	     If this input parameter contains a zero length string,
//	     an error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetIndiaNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetIndiaNumbering()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setIndiaNumberingRunes(
			nStrIntSep,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetIndiaNumberingDefault - Deletes and overwrites all the
// member variable data values for the current instance of
// IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with India Numbering System integer
// digit grouping.
//
// The integer separator character will be defaulted to a
// comma (',') character.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
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
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetIndiaNumberingDefault(
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetIndiaNumberingDefault()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setIndiaNumberingRunes(
			nStrIntSep,
			[]rune{','},
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetIndiaNumberingRunes - Overwrites all the member
// variable data values for the current instance of
// IntegerSeparatorDto.
//
// This method will configure the current instance of
// IntegerSeparatorDto with India Numbering System integer
// digit grouping.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing one or more integer separator characters
// used to separate integer digit groups.
//
// In conformance with the India Numbering System, the integer
// digit grouping is defaulted to values of three (3) and two
// (2).
//
// This means that the first integer group will consist of
// three digits and all the remaining groups will consist of
// two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
// Although, users have the option of specifying one or more
// custom integer separator characters passed through input
// parameter 'intSeparatorChars', be advised that the commonly
// used integer separator for the India Numbering system is
// the comma ([]rune{','}) character.
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of IntegerSeparatorDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars          []rune
//	   - One or more characters used to separate groups of
//	     integers. This string is used to separate groups of
//	     integer digits to the left of the decimal separator
//	     (a.k.a. decimal point). The most common integer separator
//	     character used in conjunction with the India Numbering
//	     System is the comma  ([]rune{','}).
//	       Example:  6,78,90,00,00,00,00,000
//
//	     If this input parameter contains a zero length array,
//	     an error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetIndiaNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetIndiaNumberingRunes()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setIndiaNumberingRunes(
			nStrIntSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"nStrIntSep<-"))
}

// SetIntSeparatorChars - Sets the 'intSeparatorChars' member
// variable for the current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars          []rune
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length array, an
//	     error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
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
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetIntSeparatorChars(
	intSeparatorChars []rune,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.SetIntSeparatorChars()",
		"")

	if err != nil {
		return err
	}

	if len(intSeparatorChars) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a zero length array.\n",
			ePrefix.String())
	}

	err = new(strMechPreon).
		copyRuneArrays(
			&nStrIntSep.intSeparatorChars,
			&intSeparatorChars,
			true,
			ePrefix.XCpy(
				"intSeparatorChars->"+
					"nStrIntSep.intSeparatorChars"))

	return err
}

// SetIntGroupingSequence - Sets the 'Integer Separator Grouping'
// specification for the current IntegerSeparatorDto instance.
//
// This unsigned integer values specifies the number of integer
// digits within a group. This value is used to group integers
// within a number string.
//
// In most western countries, integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000'. In this case the
// 'intGroupingSequence' value would be set to three ('3').
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//
//	Chinese Numerals Example: '12,3456,7890,2345'
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intGroupingSequence        []uint
//		   - This unsigned integer array specifies the number of
//		     integer digits within a series of groups. This value is
//		     used to group integers within a number string.
//
//		     In most western countries, integer digits to the left of
//		     the decimal separator (a.k.a. decimal point) are separated
//		     into groups of three digits representing a grouping of
//		     'thousands' like this: '1,000,000,000'. In this case the
//		     'intGroupingSequence' value would be set to three
//		     (uint[]{3}).
//
//		     In some countries and cultures other integer groupings are
//		     used. In India, for example, a number might be formatted
//		     like this: '6,78,90,00,00,00,00,000'. In this case the
//		     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//		     Chinese Numerals have an integer grouping value of four
//		     and are formatted like this: '12,3456,7890,2345'. In this
//		     case the 'intGroupingSequence' value would be set to four
//		     uint[]{4}.
//
//		     If this input parameter is submitted as a zero length
//		     array or is otherwise invalid, this method will take no
//		     action and exit.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// -----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetIntGroupingSequence(
	intGroupingSequence []uint,
	errorPrefix interface{}) (
	err error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.SetIntGroupingSequence()",
		"")

	if err != nil {
		return err
	}

	lenArray := len(intGroupingSequence)

	if lenArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intGroupingSequence'\n"+
			"is a zero length array!\n",
			ePrefix.String())

		return err
	}

	for i := 0; i < lenArray; i++ {

		if intGroupingSequence[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intGroupingSequence[%v]' is equal to zero!\n",
				ePrefix.String(),
				i)

			return err
		}

		if intGroupingSequence[i] > 1000000 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intGroupingSequence[%v]' is greater\n"+
				"than 1,000,000!\n"+
				"intGroupingSequence[%v]= '%v'",
				ePrefix.String(),
				i,
				i,
				intGroupingSequence[i])

			return err
		}

	}

	_ = new(strMechPreon).
		copyUnsignedIntArrays(
			&nStrIntSep.intGroupingSequence,
			&intGroupingSequence,
			true,
			nil)

	return
}

// SetRestartIntGroupingSequence - Sets the internal member
// variable 'restartIntGroupingSequence' for the current
// IntegerSeparatorDto instance.
//
// The IntegerSeparatorDto type is intended to be configured in an
// array of IntegerSeparatorDto objects which, taken as a whole,
// provides formatting specifications for complex integer group
// separation operations.
//
// If the current IntegerSeparatorDto is the last element in an
// array of IntegerSeparatorDto objects, the 'Restart Integer
// Grouping Sequence' flag signals whether the integer separation
// operation will be restarted from the first IntegerSeparatorDto
// object in the array.
//
// Again, the IntegerSeparatorDto.restartIntGroupingSequence boolean
// flag only has meaning if the current IntegerSeparatorDto object
// is last element in an array of IntegerSeparatorDto objects.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
//
// -----------------------------------------------------------------
//
// Return Values
//
//	-- NONE --
func (nStrIntSep *IntegerSeparatorDto) SetRestartIntGroupingSequence(
	restartIntGroupingSequence bool) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	nStrIntSep.restartIntGroupingSequence =
		restartIntGroupingSequence

	return
}

// SetThousands - Overwrites all the member variable data values for
// the current IntegerSeparatorDto. This method is intended to
// configure a basic or simple integer separator object using
// default values and a minimum number of input parameters
//
// Input parameter 'intSeparatorChars' specifies  the integer
// separator character or characters. The integer digit grouping is
// defaulted to a value of three (3). The 'separator repetitions'
// value is defaulted to zero (0), signaling unlimited repetitions.
// Finally, the 'restartIntGroupingSequence' flag will be defaulted
// to 'false'.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'intSeparatorChars'.
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSeparatorChars    string
//	   - One or more characters used to separate groups of
//	     integers. This separator is also known as the 'thousands'
//	     separator. It is used to separate groups of integer digits
//	     to the left of the decimal separator
//	     (a.k.a. decimal point). In the United States, the standard
//	     integer digits separator is the comma (",").
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this input parameter contains a zero length string, an
//	     error will be returned.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto will be created and returned. The
//	     'integer digits grouping sequence' will be automatically
//	     set to a default value of 3-digits.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetThousands(
	intSeparatorChars string,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetThousandsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setThousandsRunes(
			nStrIntSep,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"nStrIntSep"))
}

// SetThousandsRunes - Overwrites all the member variable data values for
// the current IntegerSeparatorDto. This method is intended to
// configure a basic or simple integer separator object using
// default values and a minimum number of input parameters.
//
// The input parameter 'integerDigitsSeparators' is an array of
// runes containing the integer separator characters. The integer
// digit grouping is defaulted to a value of three (3). The
// 'separator repetitions' value is defaulted to zero (0) signaling
// unlimited repetitions.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'integerDigitsSeparators'.
//
// This method is an alternative to method
// IntegerSeparatorDto.SetThousands() in that this method accepts
// integer separator characters as an array of runes instead
// of a string.
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          []rune
//		   - A character, or series of characters, used to separate
//		     integer digits in a number string. These characters are
//		     commonly known as the 'thousands separator'. A 'thousands
//		     separator' is used to separate groups of integer digits to
//		     the left of the decimal separator (a.k.a. decimal point).
//		     In the United States, the standard integer digits
//		     separator is the single comma character (',').
//		           United States Example:  1,000,000,000
//
//		     In many European countries, a single period ('.') is used
//		     as the integer separator character.
//		           European Example: 1.000.000.000
//
//		     Other countries and cultures use spaces, apostrophes or
//		     multiple characters to separate integers.
//
//		     If this parameter is submitted as a zero length array, an
//		     error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
//	IntegerSeparatorDto
//	   - If this method completes successfully, a new instance of
//	     IntegerSeparatorDto will be created and returned. The
//	     'integer digits grouping sequence' will be automatically
//	     set to a default value of 3-digits.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetThousandsRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetThousandsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setThousandsRunes(
			nStrIntSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"nStrIntSep"))
}

// SetUnitedStatesDefaults - This method will overwrite and set the all
// the internal member variable data values to default values used
// in the United States. Integer separator values used in the
// United States consist of the comma character (','), an integer
// grouping of three ('3') and unlimited repetitions of this
// sequence.
//
//	United States Integer Separation Example:
//	      '1,000,000,000,000'
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetUnitedStatesDefaults(
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.SetUnitedStatesDefaults()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoMechanics).
		setToUSADefaults(
			nStrIntSep,
			ePrefix)
}

// SetUnitedStatesDefaultsIfEmpty - If any of the current
// IntegerSeparatorDto instance data values are zero or invalid,
// this method will reset ALL data elements to United States
// default values.
//
// If the current IntegerSeparatorDto instance is valid and
// populated with data, this method will take no action and exit.
//
// United States default numeric separators are listed as follows:
//
//	Thousands Separator (a.k.a. Integer Digits Separator) = ','
//	Integer Digits Grouping Sequence = []uint{3}
//	           United States Example:  1,000,000,000
//
// # IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
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
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetUnitedStatesDefaultsIfEmpty(
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto."+
			"SetUnitedStatesDefaultsIfEmpty()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoUtility).
		setToUSADefaultsIfEmpty(
			nStrIntSep,
			ePrefix.XCpy(
				"nStrIntSep"))
}

// String - Returns a string detailing the components of this
// IntegerSeparatorDto instance. Output includes:
//
//	Integer Separator Characters
//	Integer Separator Grouping
//	Restart Grouping Sequence
func (nStrIntSep *IntegerSeparatorDto) String() string {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	if nStrIntSep.intSeparatorChars == nil {
		nStrIntSep.intSeparatorChars =
			make([]rune, 0, 5)
	}

	str := fmt.Sprintf("Integer Separator Char(s) = '%v'\n",
		string(nStrIntSep.intSeparatorChars))

	str += fmt.Sprintf("Integer Separator Grouping = '%v'\n",
		nStrIntSep.intGroupingSequence)

	str += fmt.Sprintf("Restart Grouping Sequence  = '%v'\n",
		nStrIntSep.restartIntGroupingSequence)

	return str
}

// SetWithComponents - This method will overwrite and reset the
// internal member variable data values for the current
// IntegerSeparatorDto instance based on the component elements
// passed as input parameters.
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current IntegerSeparatorDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		intSeparatorChars          []rune
//		   - A character, or series of characters, used to separate
//		     integer digits in a number string. These characters are
//		     commonly known as the 'thousands separator'. A 'thousands
//		     separator' is used to separate groups of integer digits to
//		     the left of the decimal separator (a.k.a. decimal point).
//		     In the United States, the standard integer digits
//		     separator is the single comma character (',').
//		           United States Example:  1,000,000,000
//
//		     In many European countries, a single period ('.') is used
//		     as the integer separator character.
//		           European Example: 1.000.000.000
//
//		     Other countries and cultures use spaces, apostrophes or
//		     multiple characters to separate integers.
//
//		     If this parameter is submitted as a zero length array, an
//		     error will be returned.
//
//
//		intGroupingSequence        []uint
//		   - This unsigned integer array specifies the number of
//		     integer digits within a series of groups. This value is
//		     used to group integers within a number string.
//
//		     In most western countries, integer digits to the left of
//		     the decimal separator (a.k.a. decimal point) are separated
//		     into groups of three digits representing a grouping of
//		     'thousands' like this: '1,000,000,000'. In this case the
//		     'intGroupingSequence' value would be set to three
//		     (uint[]{3}).
//
//		     In some countries and cultures other integer groupings are
//		     used. In India, for example, a number might be formatted
//		     like this: '6,78,90,00,00,00,00,000'. In this case the
//		     'intGroupingSequence' value would be set to uint[]{3,2}.
//
//		     Chinese Numerals have an integer grouping value of four
//		     and are formatted like this: '12,3456,7890,2345'. In this
//		     case the 'intGroupingSequence' value would be set to four
//		     uint[]{4}.
//
//
//		restartIntGroupingSequence bool
//		   - If this flag is set to 'true', the grouping sequence
//		     will be restarted at the beginning of the
//		     'intGroupingSequence' array after completion of the last
//		      group in the 'intGroupingSequence' array.
//		      Example:
//		        restartIntGroupingSequence = 'true'
//		        intGroupingSequence = uint{3,2}
//		        integer = 1234567890123456
//		        result  = 1,23,456,78,901,23,456
//
//		     If this flag is set to 'false', the last element or
//		     grouping in the 'intGroupingSequence' array will simply be
//		     repeated for all the remaining integer digits.
//		      Example:
//		        restartIntGroupingSequence = 'false'
//		        intGroupingSequence = uint{3,2}
//		        integer = 1234567890123456
//		        result  = 1,23,45,67,89,01,23,456
//
//		     The need to set this value to 'true' is exceedingly rare.
//		     For the vast majority of integer separation scenarios, set
//		     this parameter should be set to 'false'.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
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
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
func (nStrIntSep *IntegerSeparatorDto) SetWithComponents(
	intSeparatorChars []rune,
	intGroupingSequence []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) error {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorDto.SetWithComponents()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorDtoMechanics).
		setWithComponents(
			nStrIntSep,
			intSeparatorChars,
			intGroupingSequence,
			restartIntGroupingSequence,
			ePrefix.XCpy("nStrIntSep"))
}
