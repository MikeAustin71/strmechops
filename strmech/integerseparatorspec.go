package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	IntegerSeparatorSpec
//
//	Integer Separator Specification.
//
//	Type IntegerSeparatorSpec is designed to manage integer
//	separators, primarily thousands separators, for
//	different countries and cultures. The term
//	'integer separators' is used because this type manages
//	both integer grouping and the characters used to
//	separate integer groups.
//
//	In the United States and many other countries, integer
//	numbers are often separated by commas thereby grouping
//	the number into thousands.
//
//		Example: 1,000,000
//
//	Other countries use characters other than the comma to
//	separate integers into thousands. Some countries and
//	cultures do not use thousands separation and instead
//	rely on multiple integer separation characters and
//	grouping sequences for a single integer number. Notable
//	examples of this are found in the 'India Number System'
//	and 'Chinese Numerals'.
//
//			United States Example
//				One-Million:	1,000,000
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//			French Example
//				One-Million:	1 000 000
//
//			India Number System Example:
//				'6,78,90,00,00,00,00,000'
//
//			Chinese Number System Example:
//				'12,3456,7890,2345'
//
// ----------------------------------------------------------------
//
//	# Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//	https://en.wikipedia.org/wiki/Chinese_numerals
//	https://en.wikipedia.org/wiki/Decimal_separator
//
//	The IntegerSeparatorSpec type provides the flexibility
//	necessary to process these complex number separation
//	formats.
//
// ----------------------------------------------------------------
//
//	# EASIEST WAY TO CREATE TYPE IntegerSeparatorSpec
//
//	The following methods offer the easiest and simplest
//	means of creating a Type IntegerSeparatorSpec. These
//	methods allow the user to create integer separation
//	configurations for the vast majority of the world's
//	countries and cultures using enumeration constants:
//
//		IntegerSeparatorSpec.NewIntGroupEnum()
//		IntegerSeparatorSpec.NewIntGroupEnumRunes()
//
//	To quickly implement some of the most common integer
//	separation conventions, reference the following
//	methods:
//
//		IntegerSeparatorSpec.NewUnitedStatesDefaults()
//		IntegerSeparatorSpec.NewFrenchDefaults()
//		IntegerSeparatorSpec.NewGermanDefaults()
//		IntegerSeparatorSpec.NewIndiaNumberingDefault()
//		IntegerSeparatorSpec.NewChineseNumberingDefault()
//		IntegerSeparatorSpec.NewThousands()
//		IntegerSeparatorSpec.NewThousandsRunes()
//
//	For specialty cases or complex integer separation
//	scenarios requiring granular control over all integer
//	separation parameters, consider the following methods:
//
//		IntegerSeparatorSpec.SetComponents()
//		IntegerSeparatorSpec.SetComponentsRunes()
type IntegerSeparatorSpec struct {
	intSeparatorChars []rune
	//	A series of runes used to separate
	//	integer digits.
	//
	//	A character, or series of characters, used to separate
	//	integer digits in a number string. These characters are
	//	commonly known as the 'thousands separator'. A 'thousands
	//	separator' is used to separate groups of integer digits to
	//	the left of the decimal separator (a.k.a. decimal point).
	//	In the United States, the standard integer digits
	//	separator is the single comma character (',').
	//
	//		United States Example
	//			One-Million:	1,000,000
	//
	//	In many European countries, a single period ('.') is
	//	used as the integer separator character.
	//
	//		German Example
	//	    	One-Million:	1.000.000
	//
	//	Other countries and cultures use spaces, apostrophes
	//	or multiple characters to separate integers.
	//
	//		French Example
	//			One-Million:	1 000 000

	intSeparatorGrouping []uint
	//	Number of integer digits in each group.
	//
	//	This unsigned integer array specifies the number of
	//	integer digits within a series of groups. This value
	//	is used to group integers within a number string.
	//
	//	In most western countries, integer digits to the
	//	left of the radix point, or decimal separator (a.k.a.
	//	decimal point), are separated into groups of three
	//	digits representing a grouping of 'thousands' like
	//	this:
	//
	//			'1,000,000,000'
	//
	//	In this case the 'intSeparatorGrouping' value would
	//	be set to three (uint[]{3}).
	//
	//	In some countries and cultures other integer groupings
	//	are used. In India, for example, a number might be
	//	formatted like this:
	//
	//			'6,78,90,00,00,00,00,000'
	//
	//	In this case the 'intSeparatorGrouping' value would be
	//	set to uint[]{3,2}.
	//
	//	Chinese Numerals have an integer grouping value of four
	//	and are formatted like this:
	//
	//		'12,3456,7890,2345'

	restartIntGroupingSequence bool
	//	If this flag is set to 'true', the grouping sequence
	//	will be restarted at the beginning of the
	//	'intSeparatorGrouping' array after completion of the
	//	last group in the 'intSeparatorGrouping' array.
	//
	//		Example-1:
	//			restartIntGroupingSequence = 'true'
	//			intSeparatorGrouping = uint{3,2}
	//			integer = 1234567890123456
	//			result  = 1,23,456,78,901,23,456
	//
	//	If this flag is set to 'false', the last element or
	//	grouping in the 'intSeparatorGrouping' array will
	//	simply be repeated for all the remaining integer
	//	digits.
	//
	//		Example-2:
	//			restartIntGroupingSequence = 'false'
	//			intSeparatorGrouping = uint{3,2}
	//			integer = 1234567890123456
	//			result  = 1,23,45,67,89,01,23,456
	//
	//	The need to set this value to 'true' is exceedingly
	//	rare. For the vast majority of integer separation
	//	scenarios, set this parameter should be set to
	//	'false'.

	turnOffIntegerSeparation bool
	//	When set to 'true' integer separation will NOT
	//	be applied.
	//
	//	The following examples show the resulting number
	//	string presentations when this parameter is set
	//	to 'true' and 'false'
	//
	//	Example turnOffIntegerSeparation = 'true'
	//		Number String: 1000000000
	//
	//	Example turnOffIntegerSeparation = 'false'
	//		Number String: 1,000,000,000

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming
//	IntegerSeparatorSpec instance to the data fields of
//	the current IntegerSeparatorSpec instance.
//
//	If input parameter 'incomingNStrIntSeparator' is
//	judged to be invalid, this method will return an
//	error.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised, all the data fields in the current
//	IntegerSeparatorSpec instance will be overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingNStrIntSeparator	*IntegerSeparatorSpec
//
//		A pointer to an instance of IntegerSeparatorSpec.
//		The data values in this object will be copied to
//		the current IntegerSeparatorSpec instance.
//
//		If input parameter 'incomingNStrIntSeparator' is
//		judged to be invalid, this method will return an
//		error.
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
//	err							error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) CopyIn(
	incomingNStrIntSeparator *IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMolecule).
		copyIntSepSpec(
			nStrIntSep,
			incomingNStrIntSeparator,
			ePrefix)
}

//	CopyOut
//
//	Creates and returns a deep copy of the current
//	IntegerSeparatorSpec instance.
//
//	If the current IntegerSeparatorSpec instance is
//	judged to be invalid, this method will return an
//	error.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec will be created
//		and returned containing all the data values
//		copied from the current instance of
//		IntegerSeparatorSpec.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyIntSepDto IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec.CopyOut()",
		"")

	if err != nil {
		return deepCopyIntSepDto, err
	}

	err = new(integerSeparatorSpecMolecule).
		copyIntSepSpec(
			&deepCopyIntSepDto,
			nStrIntSep,
			ePrefix.XCpy(
				"deepCopyIntSepDto<-nStrIntSep"))

	return deepCopyIntSepDto, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of IntegerSeparatorSpec to their zero or
//	uninitialized states. This method will leave the
//	current instance of IntegerSeparatorSpec in an invalid
//	state and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in the current instance of IntegerSeparatorSpec.
//	All member variable data values will be reset to their
//	zero or uninitialized states.
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
//	NONE
func (nStrIntSep *IntegerSeparatorSpec) Empty() {
	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	_ = new(integerSeparatorSpecQuark).
		empty(
			nStrIntSep,
			nil)

	nStrIntSep.lock.Unlock()

	nStrIntSep.lock = nil

	return
}

//	Equal
//
//	Receives a pointer to an IntegerSeparatorSpec object
//	and proceeds to determine whether all data elements
//	in this object are equal to all corresponding data
//	elements in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	nStrIntSepDto				*IntegerSeparatorSpec
//
//		This method will compare all data elements in
//		the current IntegerSeparatorSpec object to
//		corresponding data elements in this second
//		IntegerSeparatorSpec object in order to determine
//		equivalency.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	bool
//
//		If all the data elements in the current
//		IntegerSeparatorSpec instance are equal to all
//		the corresponding data elements in 'nStrIntSepDto',
//		this return parameter will be set to 'true'. If all
//		the data elements are NOT equal, this return
//		parameter will be set to 'false'.
//
//
//	error
//
//		If all the data elements in the current
//		IntegerSeparatorSpec instance are equal to all the
//		corresponding data elements	contained in input
//		parameter 'nStrIntSepDto', this return parameter
//		will be set to 'nil'.
//
//		If the corresponding data elements are not equal,
//		a detailed error message identifying the unequal
//		elements will be returned.
//
//		If errors are encountered during processing, the
//		returned error message will incorporate the method
//		chain and text passed by input parameter,
//		'errorPrefix'. The 'errorPrefix' text will be
//		attached to the beginning of the error message.
func (nStrIntSep *IntegerSeparatorSpec) Equal(
	nStrIntSepDto *IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec.Equal())",
		"")

	if err != nil {
		return false, err
	}

	return new(integerSeparatorSpecMolecule).
		equal(
			nStrIntSep,
			nStrIntSepDto,
			ePrefix.XCpy(
				"nStrIntSep vs nStrIntSepDto"))
}

//	GetIntSeparatorChars
//
//	Returns the integer digit separator characters for
//	the current IntegerSeparatorSpec instance as an array
//	of runes.
//
//	If the rune array is zero length, an error will be
//	returned.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
//	[] rune
//
//		If this method completes successfully, an array
//		of runes will be returned representing the integer
//		digit separator characters.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) GetIntSeparatorChars(
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
		"IntegerSeparatorSpec.GetIntSeparatorChars()",
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

//	GetIntSeparatorStr
//
//	Returns the integer digit separator characters
//	for the current IntegerSeparatorSpec instance
//	as a string.
//
//	If the string has a zero length, an error will be
//	returned.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
//		If this method completes successfully, a string
//		will be returned containing the integer digit
//		separator characters.
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
func (nStrIntSep *IntegerSeparatorSpec) GetIntSeparatorStr(
	errorPrefix interface{}) (string, error) {

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
		"IntegerSeparatorSpec."+
			"GetIntSeparatorStr()",
		"")

	if err != nil {
		return "", err
	}

	if nStrIntSep.intSeparatorChars == nil {
		nStrIntSep.intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSepChars := len(nStrIntSep.intSeparatorChars)

	if lenIntSepChars == 0 {
		return "",
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

	return string(newIntSepChars), nil
}

//	GetIntegerGrouping
//
//	Returns an unsigned integer representing the number
//	of integer digits included in an integer group for
//	the current IntegerSeparatorSpec instance.
//
//	This unsigned integer value is used to group integers
//	within a number string.
//
//	In most western countries, integer digits to the left
//	of the radix point or decimal separator (a.k.a.
//	decimal point), are separated into groups of three
//	digits representing a grouping of 'thousands' like
//	this:
//
//			'1,000,000'
//
//	In this case the 'intSeparatorGrouping' value would be
//	set to three (uint[]{3}).
//
//	In some countries and cultures other integer groupings
//	are used. In India, for example, a number might be
//	formatted like this:
//
//		'6,78,90,00,00,00,00,000'
//
//	In this case the 'intSeparatorGrouping' value would
//	be set to uint[]{3,2}.
//
//	Chinese Numerals have an integer grouping value of
//	four and are formatted like this:
//
//			'12,3456,7890,2345'
//
//	In this case the 'intSeparatorGrouping' value would be
//	set to four uint[]{4}.
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
//	[]uint
//
//		An array of unsigned integer detailing the integer
//		grouping sequence for the current instance of
//		IntegerSeparatorSpec.
func (nStrIntSep *IntegerSeparatorSpec) GetIntegerGrouping() []uint {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var newRuneArray []uint = nil

	_ = new(strMechPreon).
		copyUnsignedIntArrays(
			&newRuneArray,
			&nStrIntSep.intSeparatorGrouping,
			true,
			nil)

	return newRuneArray
}

//	GetFmtIntSeparatedNumStr
//
//	This method receives a pure number string
//	consisting of numeric digit characters zero
//	('0') through nine ('9'), inclusive.
//
//	These numeric character digits are then
//	grouped and separated using the Integer
//	Separation parameters provided by the
//	current instance of IntegerSeparatorSpec.
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
//		IntegerSeparatorSpec.
//
//		If this string contains any non-numeric characters,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	string
//
//		If this method completes successfully, a
//		formatted string of integer digits will be
//		returned incorporating the integer grouping
//		and separation characters specified by the
//		current instance of IntegerSeparatorSpec.
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
func (nStrIntSep *IntegerSeparatorSpec) GetFmtIntSeparatedNumStr(
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
		"IntegerSeparatorSpec."+
			"GetFmtIntSeparatedNumStr()",
		"")

	if err != nil {
		return "", err
	}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = new(integerSeparatorSpecMolecule).applyIntSeparators(
		nStrIntSep,
		[]rune(pureNumString),
		ePrefix.XCpy(
			"numStrWithIntSeps<-pureNumString"))

	return string(numStrWithIntSeps), err
}

//	GetFmtIntSeparatedNumRunes
//
//	This method receives an array of runes
//	consisting of numeric digit characters
//	zero ('0') through nine ('9'), inclusive.
//
//	These numeric character digits are then
//	grouped and separated using the Integer
//	Separation parameters provided by the
//	current instance of IntegerSeparatorSpec.
//
//	Example:
//		     Pure Number String: 123456
//		Formatted Number String: 123,456
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberRunes				[]rune
//
//		An array of runes consisting exclusively of numeric
//		text characters falling within the range of zero
//		('0') through nine ('9'), inclusive.
//
//		These numeric character digits are then	grouped and
//		separated using the Integer Separation parameters
//		provided by the	current instance of
//		IntegerSeparatorSpec.
//
//		If this rune array contains any non-numeric characters,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	[]rune
//
//		If this method completes successfully, a
//		formatted rune array of integer digits will
//		be returned incorporating the integer grouping
//		and separation characters specified by the
//		current instance of IntegerSeparatorSpec.
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
func (nStrIntSep *IntegerSeparatorSpec) GetFmtIntSeparatedNumRunes(
	pureNumberRunes []rune,
	errorPrefix interface{}) (
	[]rune,
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
		"IntegerSeparatorSpec."+
			"GetFmtIntSeparatedNumRunes()",
		"")

	if err != nil {
		return []rune{}, err
	}

	var numRunesWithIntSeps []rune

	numRunesWithIntSeps,
		err = new(integerSeparatorSpecMolecule).applyIntSeparators(
		nStrIntSep,
		pureNumberRunes,
		ePrefix.XCpy(
			"numRunesWithIntSeps<-pureNumberRunes"))

	return numRunesWithIntSeps, err
}

//	GetRestartIntGroupingSequence
//
//	Returns the internal member variable
//	'restartIntGroupingSequence' for the current
//	IntegerSeparatorSpec instance.
//
//	If this flag is set to 'true', the grouping
//	sequence will be restarted at the beginning
//	of the 'intSeparatorGrouping' array after
//	completion of the last group in the
//	'intSeparatorGrouping' array.
//
//	Example-1:
//		restartIntGroupingSequence = 'true'
//		intSeparatorGrouping = uint{3,2}
//		integer = 1234567890123456
//		result  = 1,23,456,78,901,23,456
//
//	If this flag is set to 'false', the last element
//	or grouping in the 'intSeparatorGrouping' array
//	will simply be repeated for all the remaining
//	integer digits.
//
//	Example-2:
//		restartIntGroupingSequence = 'false'
//		intSeparatorGrouping = uint{3,2}
//		integer = 1234567890123456
//		result  = 1,23,45,67,89,01,23,456
//
//	The need to set this value to 'true' is exceedingly
//	rare. For the vast majority of integer separation
//	scenarios, set this parameter should be set to
//	'false'.
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
//	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example-1:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example-2:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
func (nStrIntSep *IntegerSeparatorSpec) GetRestartIntGroupingSequence() bool {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return nStrIntSep.restartIntGroupingSequence
}

//	IsValidInstance
//
//	Performs a diagnostic review of the current
//	IntegerSeparatorSpec instance to determine
//	whether the current instance is valid in all
//	respects.
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
//	isValid						bool
//
//		This returned boolean value will signal
//		whether the current IntegerSeparatorSpec
//		instance is valid, or not.
//
//		If the current IntegerSeparatorSpec instance
//		contains valid data, this  method returns
//		'true'.
//
//		If the current IntegerSeparatorSpec instance
//		is invalid, this method will return 'false'.
func (nStrIntSep *IntegerSeparatorSpec) IsValidInstance() (
	isValid bool) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	isValid,
		_ =
		new(integerSeparatorSpecQuark).
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				nil)

	return isValid
}

//	IsValidInstanceError
//	Performs a diagnostic review of the current
//	IntegerSeparatorSpec instance to determine whether
//	the current instance is valid in all respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	err							error
//
//		If the current instance of IntegerSeparatorSpec
//		contains invalid data, a detailed error message
//		will be returned identifying the invalid data item.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errorPrefix'. The 'errorPrefix' text will be attached
//		to the beginning of the error message.
//
//		If the current instance is valid, this error parameter
//	    will be set to nil.
func (nStrIntSep *IntegerSeparatorSpec) IsValidInstanceError(
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
		"IntegerSeparatorSpec.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err =
		new(integerSeparatorSpecQuark).
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				ePrefix)

	return err
}

//	NewChineseNumbering
//
//	Returns a new instance of IntegerSeparatorSpec
//	configured for the Chinese Numbering System
//	integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is a
//	string containing the integer separator
//	character or characters which will separate
//	the integer digit groups.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
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
//	By default, integer separation is turned on and
//	fully engaged. To control the application of
//	integer separators, see methods:
//
//		IntegerSeparatorSpec.TurnIntegerSeparationOff()
//
//		IntegerSeparatorSpec.TurnIntegerSeparationOn()
//
//		IntegerSeparatorSpec.IsIntegerSeparationOn()
//
// ----------------------------------------------------------------
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
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		Be advised that the commonly used integer separator
//		for the Chinese Numbering system is the comma (",")
//		character.
//
//			Chinese Numeral System Example:
//				'6,7890,0000,0000,0000'
//
//		If this input parameter contains a zero length string,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec, configured
//		with Chinese Numbering System integer grouping,
//		will be	created and returned.
//
//		By default, integer separation is turned on and
//		fully engaged. To control the application of
//		integer separators, see methods:
//
//			IntegerSeparatorSpec.TurnIntegerSeparationOff()
//
//			IntegerSeparatorSpec.TurnIntegerSeparationOn()
//
//			IntegerSeparatorSpec.IsIntegerSeparationOn()
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
func (nStrIntSep *IntegerSeparatorSpec) NewChineseNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewChineseNumberingDefault
//
//	Returns a new instance of IntegerSeparatorSpec
//	configured for the Chinese Numbering System
//	integer digit grouping.
//
//	The integer separator character will be
//	defaulted to the comma (',') character.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new instance
//		of IntegerSeparatorSpec, configured with the Chinese
//		Numbering System integer grouping, will be created
//		and returned.
//
//		The integer separator character for the new returned
//		instance will be defaulted to the comma (',') character.
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
func (nStrIntSep *IntegerSeparatorSpec) NewChineseNumberingDefault(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewChineseNumberingRunes
//
//	Returns a new instance of IntegerSeparatorSpec
//	configured for the Chinese Numbering System
//	integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is an
//	array of runes containing the integer separator
//	character or characters used to separate the
//	integer digit groups.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
//
//	Although, users have the option of specifying
//	one or more custom integer separator characters
//	passed through input parameter 'intSeparatorChars',
//	be advised that the commonly used integer separator
//	for the Chinese Numbering system is the comma
//	character ([]rune{','}).
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars		[]rune
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		Be advised that the commonly used integer separator
//		for the Chinese Numbering system is the comma (",")
//		character.
//
//			Chinese Numeral System Example:
//				'6,7890,0000,0000,0000'
//
//		If this input parameter contains a zero length string,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new instance of
//		IntegerSeparatorSpec, configured with Chinese Numbering
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
func (nStrIntSep *IntegerSeparatorSpec) NewChineseNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewComponents
//
//	Creates and returns a new instance of IntegerSeparatorSpec.
//	The new instance is generated based on component elements
//	passed as input parameter strings.
//
//	This method differs from IntegerSeparatorSpec.NewComponentsRunes()
//	in that this method accepts a string for input parameter
//
// 'intSeparatorChars' whereas IntegerSeparatorSpec.NewComponentsRunes()
//
//	is designed to accept rune arrays.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//	intSeparatorGrouping			[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value
//		is used to group integers within a number string.
//
//		In most western countries, integer digits to the left
//		of the radix point or decimal separator (a.k.a.
//		decimal point), are separated into groups of three
//		digits representing a grouping of 'thousands' like
//		this:
//
//				'1,000,000'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to three (uint[]{3}).
//
//		In some countries and cultures other integer groupings
//		are used. In India, for example, a number might be
//		formatted like this:
//
//			'6,78,90,00,00,00,00,000'
//
//		In this case the 'intSeparatorGrouping' value would
//		be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of
//		four and are formatted like this:
//
//				'12,3456,7890,2345'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to four uint[]{4}.
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly
//		rare. For the vast majority of integer separation
//		scenarios, set this parameter should be set to
//		'false'.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	newIntSep					IntegerSeparatorSpec
//
//		If this method completes successfully, the parameter
//		will return a new and fully populated instance of
//		IntegerSeparatorSpec.
//
//	err							error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewComponents(
	intSeparatorChars string,
	intSeparatorGrouping []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec.NewComponents()",
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
		new(integerSeparatorSpecMechanics).
			setWithComponents(
				&newIntSep,
				[]rune(intSeparatorChars),
				intSeparatorGrouping,
				restartIntGroupingSequence,
				ePrefix.XCpy("newIntSep"))

	return newIntSep, err
}

//	NewComponentsRunes
//
//	Creates and returns a new instance of
//	IntegerSeparatorSpec. The new instance is generated
//	based on component elements passed as input
//	parameters.
//
//	This method differs from IntegerSeparatorSpec.NewComponents()
//	in that this method accepts an array of runes for
//	'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this parameter is submitted as a zero length array,
//		an error will be returned.
//
//	intSeparatorGrouping		[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value
//		is used to group integers within a number string.
//
//		In most western countries, integer digits to the left
//		of the radix point or decimal separator (a.k.a.
//		decimal point), are separated into groups of three
//		digits representing a grouping of 'thousands' like
//		this:
//
//				'1,000,000'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to three (uint[]{3}).
//
//		In some countries and cultures other integer groupings
//		are used. In India, for example, a number might be
//		formatted like this:
//
//			'6,78,90,00,00,00,00,000'
//
//		In this case the 'intSeparatorGrouping' value would
//		be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of
//		four and are formatted like this:
//
//				'12,3456,7890,2345'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to four uint[]{4}.
//
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example-1:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example-2:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly
//		rare. For the vast majority of integer separation
//		scenarios, set this parameter should be set to
//		'false'.
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
//	newIntSep					IntegerSeparatorSpec
//
//		If this method completes successfully, the parameter
//		will return a new and fully populated instance of
//		IntegerSeparatorSpec.
//
//
//	err							error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewComponentsRunes(
	intSeparatorChars []rune,
	intSeparatorGrouping []uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec."+
			"NewComponentsRunes()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		new(integerSeparatorSpecMechanics).
			setWithComponents(
				&newIntSep,
				intSeparatorChars,
				intSeparatorGrouping,
				restartIntGroupingSequence,
				ePrefix.XCpy("newIntSep"))

	return newIntSep, err
}

//	NewFrenchDefaults
//
//	Creates and returns a new instance of IntegerSeparatorSpec
//	set to French default integer separators. Integer
//	separator values used in France consist of the one blank
//	or empty space (' '), an integer grouping of three ('3')
//	and unlimited repetitions of this sequence.
//
//		French Integer Separation Example
//		(One-Million): '1 000 000'
//
//	While France is a member of the European Union, various
//	European Union member countries apply different characters
//	for integer separation.
//
//	While some countries follow the French integer separation
//	convention, many other European Union member countries
//	follow the integer separation convention used in Germany.
//	For information on German integer separation usage, see
//	method:
//
//		IntegerSeparatorSpec.NewGermanDefaults()
//
//	If custom integer separator algorithms are required, see
//	method:
//
//		IntegerSeparatorSpec.NewComponents()
//		IntegerSeparatorSpec.NewComponentsRunes()
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		This will return a new and fully populated instance
//		of IntegerSeparatorSpec configured with French
//		default integer separator values.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewFrenchDefaults(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
			"NewFrenchDefaults()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err = new(integerSeparatorSpecMechanics).
		setToFrenchDefaults(
			&newIntSep,
			ePrefix.XCpy(
				"newIntSep"))

	return newIntSep, err
}

//	NewGermanDefaults
//
//	Creates and returns a new instance of IntegerSeparatorSpec
//	set to German default integer separators. Integer
//	separator values used in Germany consist of the period
//	character ('.'), an integer grouping of three ('3')
//	and unlimited repetitions of this sequence.
//
//		German Integer Separation Example (One-Million):
//	      				'1.000.000'
//
//	While Germany is a member of the European Union, various
//	European Union member countries apply different characters
//	for integer separation.
//
//	While some countries follow the German integer separation
//	convention, other European Union member countries
//	follow the integer separation convention used in France.
//	For information on French integer separation usage, see
//	method:
//
//		IntegerSeparatorSpec.NewFrenchDefaults()
//
//	If custom integer separator algorithms are required, see
//	methods:
//
//		IntegerSeparatorSpec.NewComponents()
//		IntegerSeparatorSpec.NewComponentsRunes()
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	IntegerSeparatorSpec
//
//		This will return a new and fully populated instance
//		of IntegerSeparatorSpec configured with German
//		default integer separator values.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewGermanDefaults(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
			"NewGermanDefaults()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err = new(integerSeparatorSpecMechanics).
		setToGermanDefaults(
			&newIntSep,
			ePrefix.XCpy(
				"newIntSep"))

	return newIntSep, err
}

//	NewIndiaNumbering
//
//	Returns a new instance of IntegerSeparatorSpec configured
//	for the India Numbering System integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is a string
//	containing one or more integer separator characters used
//	to separate integer digit groups.
//
//	In conformance with the India Numbering System, the
//	integer digit grouping is defaulted to values of three
//	(3) and two (2).
//
//	This means that the first integer group will consist
//	of three digits and all the remaining groups will consist
//	of two digits.
//
//		Example: '6,78,90,00,00,00,00,000'
//
//	Although, users have the option of specifying one or more
//	custom integer separator characters passed through input
//	parameter 'intSeparatorChars', be advised that the
//	commonly used integer separator for the India Numbering
//	system is the comma (",") character.
//
//	The 'integer separator repetitions' value is defaulted to
//	zero (0) signaling unlimited repetitions.
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this input parameter contains a zero length string,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new instance of
//		IntegerSeparatorSpec, configured with India Numbering
//		System integer grouping, will be created and returned.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewIndiaNumbering(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewIndiaNumberingDefault
//
//	Returns a new instance of IntegerSeparatorSpec
//	configured for the India Numbering System integer
//	digit grouping.
//
//	The integer separator character will be defaulted
//	to a comma (',') character.
//
//	In conformance with the India Numbering System,
//	the integer digit grouping is defaulted to values
//	of three (3) and two (2).
//
//	This means that the first integer group will consist
//	of three digits and all the remaining groups will
//	consist of two digits.
//
//		Example: '6,78,90,00,00,00,00,000'
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec, configured
//		with India Numbering System integer grouping,
//		will be created and returned.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewIndiaNumberingDefault(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewIndiaNumberingRunes
//
//	Returns a new instance of IntegerSeparatorSpec
//	configured for the India Numbering System
//	integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is an
//	array of runes containing one or more integer
//	separator characters used to separate integer
//	digit groups.
//
//	In conformance with the India Numbering System,
//	the integer digit grouping is defaulted to
//	values of three (3) and two (2).
//
//	This means that the first integer group will
//	consist of three digits and all the remaining
//	groups will consist of two digits.
//
//		Example: '6,78,90,00,00,00,00,000'
//
//	Although, users have the option of specifying
//	one or more custom integer separator characters
//	passed through input parameter 'intSeparatorChars',
//	be advised that the commonly used integer separator
//	for the India Numbering system is the comma
//	character (",").
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		One or more characters used to separate groups
//		of integers. This string is used to separate
//		groups of integer digits to the left of the
//		radix point or decimal separator (a.k.a. decimal
//		point). The most common integer separator character
//		used in conjunction with the India Numbering
//		System is the comma character ([]rune{','}).
//
//			Example:  6,78,90,00,00,00,00,000
//
//		If this input parameter contains a zero length
//		array, an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec, configured with
//		India Numbering System integer grouping, will be
//		created and returned.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewIndiaNumberingRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
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

//	NewIntGroupEnum
//
//	Creates a new instance of IntegerSeparatorSpec
//	based on an enumeration value passed as input
//	parameter 'intGroupingSpec'.
//
//	'intGroupingType' is of type IntegerGroupingType
//	and must be set to one of the following valid
//	values:
//
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	If 'intGroupingType' is set to 'IntGroupingType.None()',
//	an empty instance of IntegerSeparatorSpec will be
//	returned. An empty instance of IntegerSeparatorSpec
//	means that no integer separation will be performed
//	in subsequent number string formatting operations.
//
//	The returned instance of IntegerSeparatorSpec will
//	be configured for one of these four Integer Grouping
//	Type values.
//
//	The character or characters separating the integer
//	groups is supplied by input parameter
//	'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType enumeration
//		defines how 'intSep' will be configured for integer
//		grouping. The enumeration value 'intGroupingSpec'
//		must be set to one of these four values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorSpec. An empty
//		instance of IntegerSeparatorSpec means that no
//		integer separation will be performed in
//		subsequent number string formatting operations.
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this input parameter contains a zero length string, an
//		error will be returned.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec will be created
//		and returned. This IntegerSeparatorSpec will be
//		configured to 'Thousands','India Numbering System'
//		or 'Chinese Numbering' depending on the
//		specification provided by input parameter,
//		'intGroupingType'.
//
//		If input parameter 'intGroupingType' is set to
//		'IntGroupingType.None()', an empty instance of
//		IntegerSeparatorSpec will be returned. Be advised
//		that an empty instance of IntegerSeparatorSpec
//		means that no integer separation will be performed
//		in subsequent number string formatting operations.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewIntGroupEnum(
	intGroupingType IntegerGroupingType,
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSepSpec := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
			"NewIntGroupEnum()",
		"")

	if err != nil {
		return newIntSepSpec, err
	}

	err = new(integerSeparatorSpecHelper).
		setFromIntGroupEnum(
			&newIntSepSpec,
			intGroupingType,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"newIntSepSpec<-"))

	return newIntSepSpec, err
}

//	NewIntGroupEnumRunes
//
//	Creates a new instance of IntegerSeparatorSpec
//	based on an enumeration value passed through
//	input parameter 'intGroupingSpec'.
//
//	'intGroupingType' is of type IntegerGroupingType
//	and must be set to one of the following valid
//	values:
//
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	If 'intGroupingType' is set to 'IntGroupingType.None()',
//	an empty instance of IntegerSeparatorSpec will be
//	returned. An empty instance of IntegerSeparatorSpec
//	means that no integer separation will be performed
//	in subsequent number string formatting operations.
//
//	The returned instance of IntegerSeparatorSpec will
//	be configured for one of these four Integer Grouping
//	Type values.
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
//
//		This instance of IntegerGroupingType enumeration defines
//		how 'intSep' will be configured for integer grouping.
//		The enumeration value 'intGroupingSpec' must be set to
//		one of these four values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorSpec. An empty
//		instance of IntegerSeparatorSpec means that no
//		integer separation will be performed in
//		subsequent number string formatting operations.
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new instance
//		of IntegerSeparatorSpec will be created and returned.
//		This IntegerSeparatorSpec will be configured to
//		'Thousands', 'India Numbering System' or 'Chinese
//		Numbering' depending on the specification provided
//		by input parameter,	'intGroupingType'.
//
//		If input parameter 'intGroupingType' is set to
//		'IntGroupingType.None()', an empty instance of
//		IntegerSeparatorSpec will be returned. Be advised
//		that an empty instance of IntegerSeparatorSpec
//		means that no integer separation will be performed
//		in subsequent number string formatting operations.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewIntGroupEnumRunes(
	intGroupingType IntegerGroupingType,
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSepDto := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
			"NewIntGroupEnumRunes()",
		"")

	if err != nil {
		return newIntSepDto, err
	}

	err = new(integerSeparatorSpecHelper).
		setFromIntGroupEnum(
			&newIntSepDto,
			intGroupingType,
			intSeparatorChars,
			ePrefix.XCpy(
				"newIntSepDto<-"))

	return newIntSepDto, err
}

//	NewNoIntegerSeparation
//
//	Returns an instance of IntegerSeparatorSpec configured
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, new instance
//		of IntegerSeparatorSpec will be returned. This new
//		instance will be configured to ensure that integer
//		separation WILL NOT BE APPLIED in the subsequent
//		formatting of Number Strings.
func (nStrIntSep *IntegerSeparatorSpec) NewNoIntegerSeparation() IntegerSeparatorSpec {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	newIntSepDto := IntegerSeparatorSpec{}

	_ = new(integerSeparatorSpecQuark).empty(
		&newIntSepDto,
		nil)

	return newIntSepDto
}

//	NewThousands
//
//	Creates and returns a new instance of
//	IntegerSeparatorSpec. The returned IntegerSeparatorSpec
//	instance represents a basic or simple integer
//	separator object using default values and a minimum
//	number of input parameters.
//
//	The input parameter 'intSeparatorChars' is string
//	containing the integer separator characters. The
//	integer digit grouping is defaulted to a value of
//	three (3). The 'separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
//	This means that integer digits will be separated
//	into 'thousands' with each group containing three
//	digits each (Example: 1,000,000,000). Users have
//	the option of specifying integer separator
//	characters through input parameter
//	'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec will be created
//		and returned. The 'integer digits grouping
//		sequence' will be automatically  set to a
//		default value of 3-digits.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewThousands(
	intSeparatorChars string,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec."+
			"NewThousands()",
		"")

	if err != nil {
		return IntegerSeparatorSpec{}, err
	}

	ePrefix.SetEPref(
		"IntegerSeparatorSpec.NewThousands()")

	newIntSep := IntegerSeparatorSpec{}

	err =
		new(integerSeparatorDtoUtility).
			setThousandsRunes(
				&newIntSep,
				[]rune(intSeparatorChars),
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

//	NewThousandsRunes
//
//	Creates and returns a new instance of
//	IntegerSeparatorSpec. The returned
//	IntegerSeparatorSpec instance represents a basic
//	or simple integer separator object using default
//	values and a minimum number of input parameters.
//
//	This method is an alternative to method
//	IntegerSeparatorSpec.NewThousands() in that this
//	method accepts integer separator characters as an
//	array of runes instead of a string.
//
//	The input parameter 'intSeparatorChars' is an
//	array of runes containing the integer separator
//	characters. The integer digit grouping is defaulted
//	to a value of three (3). The 'separator repetitions'
//	value is defaulted to zero (0) signaling unlimited
//	repetitions. The 'restartIntGroupingSequence' flag
//	will be defaulted to 'false'.
//
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each (Example: 1,000,000,000). Users have the option
//	of specifying integer separator characters through
//	input parameter 'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this parameter is submitted as a zero length array,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a new
//		instance of IntegerSeparatorSpec will be created
//		and returned. The 'integer digits grouping
//		sequence' will be automatically set to a default
//		value of 3-digits.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewThousandsRunes(
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
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
		"IntegerSeparatorSpec."+
			"NewThousandsRunes()",
		"")

	if err != nil {
		return IntegerSeparatorSpec{}, err
	}

	newIntSep := IntegerSeparatorSpec{}

	err =
		new(integerSeparatorDtoUtility).
			setThousandsRunes(
				&newIntSep,
				intSeparatorChars,
				ePrefix.XCpy(
					"newIntSep"))

	return newIntSep, err
}

//	NewUnitedStatesDefaults
//
//	Creates and returns a new instance of IntegerSeparatorSpec
//	set to United States default integer separators. Integer
//	separator values used in the United States consist of the
//	comma character (','), an integer grouping of three ('3')
//	and unlimited repetitions of this sequence.
//
//		United States Integer Separation Example
//		(One-Million): '1,000,000'
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	IntegerSeparatorSpec
//
//		This will return a new and fully populated instance
//		of IntegerSeparatorSpec configured with United States
//		default integer separator values.
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) NewUnitedStatesDefaults(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newIntSep := IntegerSeparatorSpec{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorSpec."+
			"NewUnitedStatesDefaults()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err = new(integerSeparatorSpecMechanics).
		setToUSADefaults(
			&newIntSep,
			ePrefix.XCpy(
				"newIntSep"))

	return newIntSep, err
}

//	SetChineseNumbering
//
//	Overwrites all the member variable data values
//	for current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with Chinese Numbering
//	System integer grouping.
//
//	The input parameter 'intSeparatorChars' is a
//	string containing the integer separator character
//	or characters which will separate the integer
//	digit groups.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
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
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. These characters are used to separate
//		groups of integer digits to the left of the radix
//		point or decimal separator (a.k.a. decimal point).
//
//		The most common integer separator used in conjunction
//		with the Chinese Numbering System is the comma (",").
//
//			Example:  6,7890,0000,0000,0000
//
//		If this input parameter contains a zero length string, an
//		error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetChineseNumbering(
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
		"IntegerSeparatorSpec."+
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

//	SetChineseNumberingDefault
//
//	Overwrites all the member variable data values
//	for current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with Chinese Numbering
//	System integer grouping.
//
//	The integer separator character will be defaulted
//	to a comma (',') character.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will consist
//	of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
//
//	The 'integer separator repetitions' value is defaulted
//	to zero (0) signaling unlimited repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetChineseNumberingDefault(
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
		"IntegerSeparatorSpec."+
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

//	SetChineseNumberingRunes
//
//	Overwrites all the member variable data values
//	for current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with Chinese Numbering
//	System integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is an
//	array of runes containing the integer separator
//	character or characters used to separate the
//	integer digit groups.
//
//	In conformance with the Chinese Numbering or
//	Chinese Numeral System, the integer digit
//	grouping is defaulted to a value of four (4).
//
//	This means that integer digit grouping will
//	consist of groups of four (4) integer digits.
//
//		Example: '6,7890,0000,0000,0000'
//
//	Although, users have the option of specifying one
//	or more custom integer separator characters passed
//	through input parameter 'intSeparatorChars', be
//	advised that the commonly used integer separator
//	for the Chinese Numbering system is the comma
//	character ([]rune{','}).
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Chinese_numerals
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		One or more characters used to separate groups of
//		integers. These characters are used to separate
//		groups of integer digits to the left of the radix
//		point or decimal separator (a.k.a. decimal point).
//
//		The most common integer separator used in conjunction
//		with the Chinese Numbering System is the comma
//		character ([]rune{,}).
//
//			Example:  6,7890,0000,0000,0000
//
//		If this input parameter contains a zero length array,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetChineseNumberingRunes(
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
		"IntegerSeparatorSpec."+
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

//	SetComponents
//
//	This method will overwrite and reset the internal
//	member variable data values for the current
//	IntegerSeparatorSpec instance based on the component
//	elements passed as input parameterS.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data values
//	in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	intSeparatorGrouping			[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value
//		is used to group integers within a number string.
//
//		In most western countries, integer digits to the left
//		of the radix point or decimal separator (a.k.a.
//		decimal point), are separated into groups of three
//		digits representing a grouping of 'thousands' like
//		this:
//
//				'1,000,000'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to three (uint[]{3}).
//
//		In some countries and cultures other integer groupings
//		are used. In India, for example, a number might be
//		formatted like this:
//
//			'6,78,90,00,00,00,00,000'
//
//		In this case the 'intSeparatorGrouping' value would
//		be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of
//		four and are formatted like this:
//
//				'12,3456,7890,2345'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to four uint[]{4}.
//
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly
//		rare. For the vast majority of integer separation
//		scenarios, set this parameter should be set to
//		'false'.
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetComponents(
	intSeparatorChars string,
	intSeparatorGrouping []uint,
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
		"IntegerSeparatorSpec."+
			"SetComponents()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMechanics).
		setWithComponents(
			nStrIntSep,
			[]rune(intSeparatorChars),
			intSeparatorGrouping,
			restartIntGroupingSequence,
			ePrefix.XCpy("nStrIntSep"))
}

//	SetComponentsRunes
//
//	This method will overwrite and reset the internal
//	member variable data values for the current
//	IntegerSeparatorSpec instance based on the component
//	elements passed as input parameters.
//
//	This method differs from IntegerSeparatorSpec.SetComponents()
//	in that this method accepts an array of runes for
//	'intSeparatorChars'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data values
//	in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this parameter is submitted as a zero length array,
//		an error will be returned.
//
//	intSeparatorGrouping		[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value
//		is used to group integers within a number string.
//
//		In most western countries, integer digits to the left
//		of the radix point or decimal separator (a.k.a.
//		decimal point), are separated into groups of three
//		digits representing a grouping of 'thousands' like
//		this:
//
//				'1,000,000'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to three (uint[]{3}).
//
//		In some countries and cultures other integer groupings
//		are used. In India, for example, a number might be
//		formatted like this:
//
//			'6,78,90,00,00,00,00,000'
//
//		In this case the 'intSeparatorGrouping' value would
//		be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of
//		four and are formatted like this:
//
//				'12,3456,7890,2345'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to four uint[]{4}.
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example-1:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example-2:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly
//		rare. For the vast majority of integer separation
//		scenarios, set this parameter should be set to
//		'false'.
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetComponentsRunes(
	intSeparatorChars []rune,
	intSeparatorGrouping []uint,
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
		"IntegerSeparatorSpec."+
			"SetComponentsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMechanics).
		setWithComponents(
			nStrIntSep,
			intSeparatorChars,
			intSeparatorGrouping,
			restartIntGroupingSequence,
			ePrefix.XCpy("nStrIntSep"))
}

//	SetFrenchDefaults
//
//	This method will overwrite and set the all the
//	internal member variable data values for the
//	current instance of IntegerSeparatorSpec to
//	default values used in France. Integer
//	separator values used in France consist of the
//	space character (' '), an integer grouping of three
//	('3') and unlimited repetitions of this sequence.
//
//		French Integer Separation Example
//		(One-Million): '1 000 000'
//
//	While France is a member of the European Union,
//	various European Union member countries apply
//	different characters for integer separation.
//
//	Some countries follow the French integer separation
//	convention but other European Union member countries
//	follow the integer separation convention used in
//	Germany. For information on German integer separation
//	usage, see method:
//
//		IntegerSeparatorSpec.SetGermanDefaults()
//
//	If custom integer separator algorithms are required, see
//	method:
//
//		IntegerSeparatorSpec.SetComponents()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data
//	values in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetFrenchDefaults(
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
		"IntegerSeparatorSpec."+
			"SetFrenchDefaults()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMechanics).
		setToFrenchDefaults(
			nStrIntSep,
			ePrefix)
}

//	SetGermanDefaults
//
//	This method will overwrite and set the all the
//	internal member variable data values for the
//	current instance of IntegerSeparatorSpec to
//	default values used in Germany. Integer
//	separator values used in Germany consist of the
//	period character ('.'), an integer grouping of
//	three ('3') and unlimited repetitions of this
//	sequence.
//
//		German Integer Separation Example
//		(One-Million): '1.000.000'
//
//	While Germany is a member of the European Union,
//	various European Union member countries apply
//	different characters for integer separation.
//
//	Many countries follow the German integer separation
//	convention but other European Union member countries
//	follow the integer separation convention used in
//	France. For information on German integer separation
//	usage, see method:
//
//		IntegerSeparatorSpec.SetFrenchDefaults()
//
//	If custom integer separator algorithms are required, see
//	method:
//
//		IntegerSeparatorSpec.SetComponents()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data
//	values in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetGermanDefaults(
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
		"IntegerSeparatorSpec."+
			"SetGermanDefaults()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMechanics).
		setToGermanDefaults(
			nStrIntSep,
			ePrefix)
}

//	SetIndiaNumbering
//
//	Overwrites all the member variable data values for
//	the current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance of
//	IntegerSeparatorSpec with India Numbering System
//	integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is a string
//	containing one or more integer separator characters
//	used to separate integer digit groups.
//
//	In conformance with the India Numbering System, the
//	integer digit grouping is defaulted to values of
//	three (3) and two (2).
//
//	This means that the first integer group will consist
//	of three digits and all the remaining groups will
//	consist of two digits.
//
//		Example: '6,78,90,00,00,00,00,000'
//
//	Although, users have the option of specifying one or
//	more custom integer separator characters passed through
//	input parameter 'intSeparatorChars', be advised that the
//	commonly used integer separator for the India Numbering
//	system is the comma (",") character.
//
//	The 'integer separator repetitions' value is defaulted
//	to zero (0) signaling unlimited repetitions.
//
// ----------------------------------------------------------------
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
// values in the current instance of IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups
//		of integers. This string is used to separate groups
//		of integer digits to the left of the radix point or
//		decimal separator (a.k.a. decimal point). The most
//		common integer separator character used in conjunction
//		with the India Numbering System is the comma character
//		(",").
//
//			Example:  6,78,90,00,00,00,00,000
//
//		If this input parameter contains a zero length string,
//	    an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIndiaNumbering(
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
		"IntegerSeparatorSpec."+
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

//	SetIndiaNumberingDefault
//
//	Deletes and overwrites all the member variable data
//	values for the current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance of
//	IntegerSeparatorSpec with India Numbering System integer
//	digit grouping.
//
//	The integer separator character will be defaulted to a
//	comma (',') character.
//
//	In conformance with the India Numbering System, the
//	integer digit grouping is defaulted to values of three
//	(3) and two (2).
//
//	This means that the first integer group will consist of
//	three digits and all the remaining groups will consist
//	of two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
//	The 'integer separator repetitions' value is defaulted to
//	zero (0) signaling unlimited repetitions.
//
// ----------------------------------------------------------------
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
// values in the current instance of IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIndiaNumberingDefault(
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
		"IntegerSeparatorSpec."+
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

//	SetIndiaNumberingRunes
//
//	Overwrites all the member variable data values
//	for the current instance of IntegerSeparatorSpec.
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with India Numbering
//	System integer digit grouping.
//
//	The input parameter 'intSeparatorChars' is an
//	array of runes containing one or more integer
//	separator characters used to separate integer
//	digit groups.
//
//	In conformance with the India Numbering System, the
//	integer digit grouping is defaulted to values of
//	three (3) and two (2).
//
//	This means that the first integer group will consist
//	of three digits and all the remaining groups will
//	consist of two digits.
//
//	Example: '6,78,90,00,00,00,00,000'
//
//	Although, users have the option of specifying one or
//	more custom integer separator characters passed
//	through input parameter 'intSeparatorChars', be
//	advised that the commonly used integer separator for
//	the India Numbering system is the comma character
//	([]rune{','}).
//
//	The 'integer separator repetitions' value is
//	defaulted to zero (0) signaling unlimited
//	repetitions.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Indian_numbering_system
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		One or more characters used to separate groups
//		of integers. This string is used to separate groups
//		of integer digits to the left of the radix point or
//		decimal separator (a.k.a. decimal point). The most
//		common integer separator character used in
//		conjunction with the India Numbering System is the
//		comma character ([]rune{','}).
//
//			Example:  6,78,90,00,00,00,00,000
//
//		If this input parameter contains a zero length array,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIndiaNumberingRunes(
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
		"IntegerSeparatorSpec."+
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

//	SetIntGroupEnum
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with new data values
//	based on an enumeration value passed as input
//	parameter 'intGroupingSpec'.
//
//	The integer separator characters are passed as
//	a string ('intSeparatorChars').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values in the current instance
//	of IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType enumeration
//		defines how 'intSep' will be configured for integer
//		grouping. The enumeration value 'intGroupingSpec'
//		must be set to one of these four values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorSpec. An empty
//		instance of IntegerSeparatorSpec means that no
//		integer separation will be performed in
//		subsequent number string formatting operations.
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this input parameter contains a zero length string, an
//		error will be returned.
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
// -----------------------------------------------------------------
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIntGroupEnum(
	intGroupingType IntegerGroupingType,
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
		"IntegerSeparatorSpec."+
			"SetIntGroupEnum()",
		"")

	if err != nil {
		return err
	}

	err = new(integerSeparatorSpecHelper).
		setFromIntGroupEnum(
			nStrIntSep,
			intGroupingType,
			[]rune(intSeparatorChars),
			ePrefix.XCpy(
				"newIntSepSpec<-"))

	return err
}

//	SetIntGroupEnumRunes
//
//	This method will configure the current instance
//	of IntegerSeparatorSpec with new data values
//	based on an enumeration value passed as input
//	parameter 'intGroupingSpec'.
//
//	The integer separator characters are passed as
//	an array of runes ('intSeparatorChars').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values in the current instance
//	of IntegerSeparatorSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType enumeration
//		defines how 'intSep' will be configured for integer
//		grouping. The enumeration value 'intGroupingSpec'
//		must be set to one of these four values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		A value of IntGroupingType.None() will return
//		an empty instance of IntegerSeparatorSpec. An empty
//		instance of IntegerSeparatorSpec means that no
//		integer separation will be performed in
//		subsequent number string formatting operations.
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
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
// -----------------------------------------------------------------
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIntGroupEnumRunes(
	intGroupingType IntegerGroupingType,
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
		"IntegerSeparatorSpec."+
			"SetIntGroupEnumRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(integerSeparatorSpecHelper).
		setFromIntGroupEnum(
			nStrIntSep,
			intGroupingType,
			intSeparatorChars,
			ePrefix.XCpy(
				"newIntSepSpec<-"))

	return err
}

//	SetIntGroupingSequence
//
//	Sets the 'Integer Separator Grouping' specification for
//	the current IntegerSeparatorSpec instance.
//
//	This unsigned integer values specifies the number of
//	integer digits within a group. This value is used to
//	group integers within a number string.
//
//	In most western countries, integer digits to the left
//	of the radix point or decimal separator (a.k.a. decimal
//	point), are separated into groups of three digits
//	representing a grouping of 'thousands' like this:
//
//			'1,000,000,000'
//
//	In this case the 'intSeparatorGrouping' value would be
//	set to three ('3').
//
//	In some countries and cultures other integer groupings
//	are used. In India, for example, a number might be
//	formatted like this:
//
//			'6,78,90,00,00,00,00,000'
//
//	Chinese Numerals have an integer grouping value of four
//	('4').
//
//		Chinese Numerals Example: '12,3456,7890,2345'
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorGrouping			[]uint
//
//		This unsigned integer array specifies the number of
//		integer digits within a series of groups. This value
//		is used to group integers within a number string.
//
//		In most western countries, integer digits to the left
//		of the radix point or decimal separator (a.k.a.
//		decimal point), are separated into groups of three
//		digits representing a grouping of 'thousands' like
//		this:
//
//				'1,000,000,000'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to three (uint[]{3}).
//
//		In some countries and cultures other integer groupings
//		are used. In India, for example, a number might be
//		formatted like this:
//
//				'6,78,90,00,00,00,00,000'
//
//		In this case the 'intSeparatorGrouping' value would
//		be set to uint[]{3,2}.
//
//		Chinese Numerals have an integer grouping value of
//		four and are formatted like this:
//
//				'12,3456,7890,2345'
//
//		In this case the 'intSeparatorGrouping' value would be
//		set to four uint[]{4}.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//	err							error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIntGroupingSequence(
	intSeparatorGrouping []uint,
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
		"IntegerSeparatorSpec."+
			"SetIntGroupingSequence()",
		"")

	if err != nil {
		return err
	}

	lenArray := len(intSeparatorGrouping)

	if lenArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorGrouping'\n"+
			"is a zero length array!\n",
			ePrefix.String())

		return err
	}

	for i := 0; i < lenArray; i++ {

		if intSeparatorGrouping[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intSeparatorGrouping[%v]' is equal to zero!\n",
				ePrefix.String(),
				i)

			return err
		}

		if intSeparatorGrouping[i] > 1000000 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intSeparatorGrouping[%v]' is greater\n"+
				"than 1,000,000!\n"+
				"intSeparatorGrouping[%v]= '%v'",
				ePrefix.String(),
				i,
				i,
				intSeparatorGrouping[i])

			return err
		}

	}

	_ = new(strMechPreon).
		copyUnsignedIntArrays(
			&nStrIntSep.intSeparatorGrouping,
			&intSeparatorGrouping,
			true,
			nil)

	return
}

//	SetIntSeparatorChars
//
//	Sets the 'intSeparatorChars' member variable for
//	the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	           German Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//				French Example: 1 000 000
//
//		If this parameter is submitted as a zero length array,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetIntSeparatorChars(
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
		"IntegerSeparatorSpec."+
			"SetIntSeparatorChars()",
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

//	SetRestartIntGroupingSequence
//
//	Sets the internal member variable
//	'restartIntGroupingSequence' for the current
//	IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	restartIntGroupingSequence	bool
//
//		If this flag is set to 'true', the grouping sequence
//		will be restarted at the beginning of the
//		'intSeparatorGrouping' array after completion of the
//		last group in the 'intSeparatorGrouping' array.
//
//			Example-1:
//				restartIntGroupingSequence = 'true'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,456,78,901,23,456
//
//		If this flag is set to 'false', the last element or
//		grouping in the 'intSeparatorGrouping' array will
//		simply be repeated for all the remaining integer
//		digits.
//
//			Example-2:
//				restartIntGroupingSequence = 'false'
//				intSeparatorGrouping = uint{3,2}
//				integer = 1234567890123456
//				result  = 1,23,45,67,89,01,23,456
//
//		The need to set this value to 'true' is exceedingly
//		rare. For the vast majority of integer separation
//		scenarios, set this parameter should be set to
//		'false'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (nStrIntSep *IntegerSeparatorSpec) SetRestartIntGroupingSequence(
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

//	SetThousands
//
//	Overwrites all the member variable data values for
//	the current IntegerSeparatorSpec. This method is
//	intended to configure a basic or simple integer
//	separator object using default values and a minimum
//	number of input parameters
//
//	Input parameter 'intSeparatorChars' specifies the
//	integer separator character or characters. The integer
//	digit grouping is defaulted to a value of three (3).
//	The 'separator repetitions' value is defaulted to zero
//	(0), signaling unlimited repetitions.
//
//	Finally, the 'restartIntGroupingSequence' flag will be
//	defaulted to 'false'.
//
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each
//	(Example: 1,000,000,000).
//
//	Users have the option of specifying integer separator
//	characters through input parameter 'intSeparatorChars'.
//
//	This method is an alternative to method:
//		IntegerSeparatorSpec.SetThousandsRunes()
//
//	It differs from the alternative in that this method
//	accepts integer separator characters as a string
//	instead of an array of runes.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data values
//	in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intSeparatorChars			string
//
//		A character, or series of characters, used to separate
//		integer digits in a number string. These characters are
//		commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the radix point or decimal
//		separator (a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetThousands(
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
		"IntegerSeparatorSpec."+
			"SetThousands()",
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

//	SetThousandsRunes
//
//	Overwrites all the member variable data values
//	for the current IntegerSeparatorSpec. This method
//	is intended to configure a basic or simple integer
//	separator object using default values and a
//	minimum number of input parameters.
//
//	The input parameter 'integerDigitsSeparators' is
//	an array of runes containing the integer separator
//	characters. The integer digit grouping is defaulted
//	to a value of three (3). The 'separator repetitions'
//	value is defaulted to zero (0) signaling unlimited
//	repetitions.
//
//	This means that integer digits will be separated
//	into 'thousands' with each group containing three
//	digits each (Example: 1,000,000,000). Users have
//	the option of specifying integer separator
//	characters through input parameter
//	'integerDigitsSeparators'.
//
//	This method is an alternative to method:
//		IntegerSeparatorSpec.SetThousands()
//
//	It differs from the alternative in that this method
//	accepts integer separator characters as an array of
//	runes instead of a string.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data
//	values in the current IntegerSeparatorSpec
//	instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSeparatorChars			[]rune
//
//		A series of runes or characters used to separate
//		integer digits in a number string. These characters
//		are commonly known as the 'thousands separator'. A
//		'thousands separator' is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point).
//
//		In the United States, the standard integer digits
//		separator is the single comma character (','). Other
//		countries and cultures use periods, spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//			United States Example
//				One-Million:	1,000,000
//
//		In many European countries, a single period ('.') is
//		used as the integer separator character.
//
//	    	German Example
//	        	One-Million:	1.000.000
//
//		Other countries and cultures use spaces, apostrophes
//		or multiple characters to separate integers.
//
//			French Example
//				One-Million:	1 000 000
//
//		If this parameter is submitted as a zero length array,
//		an error will be returned.
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetThousandsRunes(
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
		"IntegerSeparatorSpec."+
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

//	SetUnitedStatesDefaults
//
//	This method will overwrite and set the all the
//	internal member variable data values for the
//	current instance of IntegerSeparatorSpec to
//	default	values used in the United States.
//
//	Integer separator values used in the United States
//	consist of the comma character (','), an integer
//	grouping of three ('3') and unlimited repetitions
//	of this sequence.
//
//		United States Integer Separation Example
//		(One-Million): '1,000,000'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will overwrite all pre-existing data
//	values in the current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetUnitedStatesDefaults(
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
		"IntegerSeparatorSpec.SetUnitedStatesDefaults()",
		"")

	if err != nil {
		return err
	}

	return new(integerSeparatorSpecMechanics).
		setToUSADefaults(
			nStrIntSep,
			ePrefix)
}

//	SetUnitedStatesDefaultsIfEmpty
//
//	If any of the current IntegerSeparatorSpec instance
//	data values are zero or invalid, this method will
//	reset ALL data elements to United States default
//	integer separation values.
//
//	If the current IntegerSeparatorSpec instance is valid
//	and populated with data, this method will take no
//	action and exit. This method will ONLY implement
//	default United States integer separation values
//	if the current IntegerSeparatorSpec instance is
//	empty or invalid.
//
//	United States default numeric separators are listed
//	as follows:
//
//	Thousands Separator (a.k.a. Integer Digits Separator) = ','
//
//	Integer Digits Grouping Sequence = []uint{3}
//	           United States Example:  1,000,000,000
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// current IntegerSeparatorSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		3. []string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4. [][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5. ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6. *ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7. IBasicErrorPrefix
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
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the beginning
//		of the error message.
func (nStrIntSep *IntegerSeparatorSpec) SetUnitedStatesDefaultsIfEmpty(
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
		"IntegerSeparatorSpec."+
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

// String
//
// Returns a string detailing the components of
// the current IntegerSeparatorSpec instance.
//
// Output includes:
//
//	Integer Separator Characters
//	Integer Separator Grouping
//	Restart Grouping Sequence
func (nStrIntSep *IntegerSeparatorSpec) String() string {

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
		nStrIntSep.intSeparatorGrouping)

	str += fmt.Sprintf("Restart Grouping Sequence  = '%v'\n",
		nStrIntSep.restartIntGroupingSequence)

	var isIntSepOn bool

	if nStrIntSep.turnOffIntegerSeparation == true {

		isIntSepOn = false

	} else {

		isIntSepOn = true
	}

	str += fmt.Sprintf("Is Integer Separation On and Active  = '%v'\n",
		isIntSepOn)

	return str
}

//	TurnIntegerSeparationOff
//
//	Calling this method will turn off integer separation
//	for number string formatting.
//
//	This means that integers will NOT be separated when
//	generating integer digits in number string formatting.
//
//	Example Turn Integer Separation Off:
//		Number String: 1000000000
//
//	Example Turn Integer Separation On:
//		Number String: 1,000,000,000
//
//	By default Integer Separation is on, active and fully
//	engaged upon creating new instances of
//	IntegerSeparatorSpec.
//
//	To verify the status of Integer Separation, see method:
//
//		IntegerSeparatorSpec.IsIntegerSeparationOn()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method TURNS INTEGER SEPARATION OFF!
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrIntSep *IntegerSeparatorSpec) TurnIntegerSeparationOff() {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	nStrIntSep.turnOffIntegerSeparation = true
}

//	TurnIntegerSeparationOn
//
//	Calling this method will turn on integer separation
//	for number string formatting.
//
//	This means that integers WILL BE separated when
//	generating integer digits in number string formatting.
//
//	Example Turn Integer Separation On:
//		Number String: 1,000,000,000
//
//	Example Turn Integer Separation Off:
//		Number String: 1000000000
//
//	By default Integer Separation is on, active and fully
//	engaged upon creating new instances of
//	IntegerSeparatorSpec.
//
//	To verify the status of Integer Separation, see method:
//
//		IntegerSeparatorSpec.IsIntegerSeparationOn()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method TURNS INTEGER SEPARATION ON!
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrIntSep *IntegerSeparatorSpec) TurnIntegerSeparationOn() {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	nStrIntSep.turnOffIntegerSeparation = false
}
