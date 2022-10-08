package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrIntegerGroupingSpec
//
// Number String Integer Grouping Specification. This
// type encapsulates the parameters required to format
// integer grouping and separation within a number
// string.
type NumStrIntegerGroupingSpec struct {
	integerSeparatorChars RuneArrayDto
	// Contains the integer separator characters
	// which will be used to separate integer groups
	// within a formatted number string.

	intGroupingType IntegerGroupingType
	// Specifies the type of integer grouping
	// algorithm or grouping sequence which will
	// be applied to integer digit grouping in a
	// formatted number string. Possible values
	// are:
	//	IntGroupingType.None()
	//	IntGroupingType.Thousands()
	//	IntGroupingType.IndiaNumbering()
	//	IntGroupingType.ChineseNumbering()

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrIntegerGroupingSpec ('incomingIntGroupSpec') to the data
// fields of the current NumStrIntegerGroupingSpec instance
// ('nStrIntGroupSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrIntegerGroupingSpec
// instance ('nStrIntGroupSpec') will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingIntGroupSpec		*NumStrIntegerGroupingSpec
//		A pointer to an instance of NumStrIntegerGroupingSpec.
//		This method will NOT change the values of internal member
//		variables contained in this instance.
//
//		All data values in this NumStrIntegerGroupingSpec instance
//		will be copied to current NumStrIntegerGroupingSpec
//		instance ('nStrIntGroupSpec').
//
//		If parameter 'incomingIntGroupSpec' is determined to
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
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) CopyIn(
	incomingIntGroupSpec *NumStrIntegerGroupingSpec,
	errorPrefix interface{}) error {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrIntGroupingSpecNanobot).
		copyNStrIntGroupSpec(
			nStrIntGroupSpec,
			incomingIntGroupSpec,
			ePrefix.XCpy(
				"nStrIntGroupSpec<-"+
					"incomingIntGroupSpec"))
}

// CopyOut - Returns a deep copy of the current
// NumStrIntegerGroupingSpec instance.
//
// If the current NumStrIntegerGroupingSpec instance contains
// invalid member variables, this method will return an error.
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
//	deepCopyNStrIntGroupSpec   NumStrIntegerGroupingSpec
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current NumStrIntegerGroupingSpec instance.
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
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNStrIntGroupSpec NumStrIntegerGroupingSpec,
	err error) {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNStrIntGroupSpec, err
	}

	err = new(numStrIntGroupingSpecNanobot).
		copyNStrIntGroupSpec(
			&deepCopyNStrIntGroupSpec,
			nStrIntGroupSpec,
			ePrefix.XCpy(
				"deepCopyNStrIntGroupSpec<-"+
					"nStrIntGroupSpec"))

	return deepCopyNStrIntGroupSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrIntegerGroupingSpec to their initial or
// zero states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NumStrIntegerGroupingSpec.
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
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) Empty() {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	new(numStrIntGroupingSpecAtom).empty(
		nStrIntGroupSpec)

	nStrIntGroupSpec.lock.Unlock()

	nStrIntGroupSpec.lock = nil
}

// Equal - Receives a pointer to another instance of
// NumStrIntegerGroupingSpec and proceeds to compare its
// internal member variables to those of the current
// NumStrIntegerGroupingSpec instance in order to determine
// if they are equivalent.
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
//	incomingNStrIntGroupSpec	*NumStrIntegerGroupingSpec
//		A pointer to an instance of NumStrIntegerGroupingSpec.
//		The internal member variable data values in this instance
//		will be compared to those in the current instance of
//		NumStrIntegerGroupingSpec. The results of this comparison
//		will be returned to the calling function as a boolean
//		value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//		If the internal member variable data values contained in
//		input parameter 'incomingNStrIntGroupSpec' are equivalent
//		in all respects to those contained in the current instance
//		of NumStrIntegerGroupingSpec, this return value will be
//		set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) Equal(
	incomingNStrIntGroupSpec *NumStrIntegerGroupingSpec) bool {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	return new(numStrIntGroupingSpecAtom).equal(
		nStrIntGroupSpec,
		incomingNStrIntGroupSpec)

}

// GetIntegerGroupingType - Returns the Integer Grouping Type
// configured for the current instance of NumStrIntegerGroupingSpec.
//
// Integer Grouping Type is an enumeration specifying the type of
// integer grouping algorithm or grouping sequence which will be
// applied to integer digit grouping in a formatted number string.
// Possible values are listed as follows:
//
//	IntGroupingType.None()
//
//	IntGroupingType.Thousands()
//		Example: 1,000,000,000
//
//	IntGroupingType.IndiaNumbering()
//		Example: 6,78,90,00,00,00,00,000
//
//	IntGroupingType.ChineseNumbering()
//		Example: 12,3456,7890,2345
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
//	IntegerGroupingType
//		This method will the Integer Grouping Type configured for
//		the current instance of NumStrIntegerGroupingSpec.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) GetIntegerGroupingType() IntegerGroupingType {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	return nStrIntGroupSpec.intGroupingType.XReturnNoneIfInvalid()
}

// GetIntegerSeparatorChars - Returns a string containing the
// integer separator character or characters configured for the
// current instance of NumStrIntegerGroupingSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (",").
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period ('.') is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
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
//		This method will return a string containing the Integer
//		Separator character or characters configured for the
//	 	current instance of NumStrIntegerGroupingSpec.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) GetIntegerSeparatorChars() string {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	return nStrIntGroupSpec.integerSeparatorChars.GetCharacterString()
}

// GetIntegerSeparatorDto - Returns an instance of
// IntegerSeparatorSpec based on the configuration parameters
// contained within the current instance of IntegerSeparatorSpec.
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
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
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
// # Return Values
//
//	IntegerSeparatorSpec
//		If this method completes successfully, a new fully
//		populated instance of IntegerSeparatorSpec will be
//		returned.
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
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) GetIntegerSeparatorDto(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"GetIntegerSeparatorSpec()",
		"")

	if err != nil {
		return IntegerSeparatorSpec{}, err
	}

	return new(IntegerSeparatorSpec).NewIntGroupEnumRunes(
		nStrIntGroupSpec.intGroupingType,
		nStrIntGroupSpec.integerSeparatorChars.GetRuneArray(),
		ePrefix.XCpy(
			"<-nStrIntGroupSpec"))
}

// GetIntegerSeparatorRunes - Returns a rune array containing
// the integer separator character or characters configured for
// the current instance of NumStrIntegerGroupingSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (",").
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period ('.') is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
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
//	[]rune
//		This method will return an array of runes containing the
//		Integer Separator character or characters configured for
//	 	the current instance of NumStrIntegerGroupingSpec.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) GetIntegerSeparatorRunes() []rune {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	return nStrIntGroupSpec.integerSeparatorChars.GetRuneArray()
}

// NewRunes - Creates and returns a new and fully populated
// instance of NumStrIntegerGroupingSpec.
//
// Type NumStrIntegerGroupingSpec is used to configure
// integer digit grouping in number strings.
//
// The input parameter 'intSeparatorChars' is an array of runes
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType			IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	intSeparatorChars		[]rune
//
//		One or more characters used to separate groups of
//		integer digits. This separator is also known as the
//		'thousands' separator. It is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the
//		standard integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length array
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	newNStrIntGroupSpec		NumStrIntegerGroupingSpec
//		If this method completes successfully, a new instance of
//		NumStrIntegerGroupingSpec will be created and returned.
//		This integer grouping specification will be configured to
//		'Thousands', 'India Numbering System' or 'Chinese
//		Numbering' depending on the specification provided by
//		input parameter, 'intGroupingSpec'.
//
//	err						error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) NewRunes(
	intGroupingType IntegerGroupingType,
	intSeparatorChars []rune,
	errorPrefix interface{}) (
	newNStrIntGroupSpec NumStrIntegerGroupingSpec,
	err error) {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"NewRunes()",
		"")

	if err != nil {
		return newNStrIntGroupSpec, err
	}

	err = new(numStrIntGroupingSpecNanobot).setNStrIntGroupSpec(
		&newNStrIntGroupSpec,
		intSeparatorChars,
		intGroupingType,
		ePrefix.XCpy(
			"newNStrIntGroupSpec"))

	return newNStrIntGroupSpec, err
}

// NewStr - Creates and returns a new and fully populated
// instance of NumStrIntegerGroupingSpec.
//
// Type NumStrIntegerGroupingSpec is used to configure
// integer digit grouping in number strings.
//
// The input parameter 'intSeparatorChars' is a string
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intGroupingType			IntegerGroupingType
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	intSeparatorChars		string
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
//		If this input parameter contains a zero length string
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	errorPrefix				interface{}
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
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
// Return Values
//
//	newNStrIntGroupSpec		NumStrIntegerGroupingSpec
//		If this method completes successfully, a new instance of
//		NumStrIntegerGroupingSpec will be created and returned.
//		This integer grouping specification will be configured to
//		'Thousands', 'India Numbering System' or 'Chinese
//		Numbering' depending on the specification provided by
//		input parameter, 'intGroupingSpec'.
//
//	err						error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) NewStr(
	intGroupingType IntegerGroupingType,
	intSeparatorChars string,
	errorPrefix interface{}) (
	newNStrIntGroupSpec NumStrIntegerGroupingSpec,
	err error) {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"NewStr()",
		"")

	if err != nil {
		return newNStrIntGroupSpec, err
	}

	err = new(numStrIntGroupingSpecNanobot).setNStrIntGroupSpec(
		&newNStrIntGroupSpec,
		[]rune(intSeparatorChars),
		intGroupingType,
		ePrefix.XCpy(
			"newNStrIntGroupSpec"))

	return newNStrIntGroupSpec, err
}

// SetRunes - Deletes and resets all member variable data values
// in the current instance of NumStrIntegerGroupingSpec.
//
// The input parameter 'intSeparatorChars' is a string
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingType			IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	intSeparatorChars		[]rune
//
//		One or more characters used to separate groups of
//		integer digits. This separator is also known as the
//		'thousands' separator. It is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the
//		standard integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length array
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err						error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) SetRunes(
	intGroupingType IntegerGroupingType,
	intSeparatorChars []rune,
	errorPrefix interface{}) error {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"SetRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrIntGroupingSpecNanobot).setNStrIntGroupSpec(
		nStrIntGroupSpec,
		intSeparatorChars,
		intGroupingType,
		ePrefix.XCpy(
			"newNStrIntGroupSpec"))

}

// SetStr - Deletes and resets all member variable data values
// in the current instance of NumStrIntegerGroupingSpec.
//
// The input parameter 'intSeparatorChars' is a string
// containing one or more integer separator characters used
// to separate integer digit groups.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intGroupingType			IntegerGroupingType
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	intSeparatorChars		string
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
//		If this input parameter contains a zero length string
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	errorPrefix				interface{}
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
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
// Return Values
//
//	err						error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) SetStr(
	intGroupingType IntegerGroupingType,
	intSeparatorChars string,
	errorPrefix interface{}) error {

	if nStrIntGroupSpec.lock == nil {
		nStrIntGroupSpec.lock = new(sync.Mutex)
	}

	nStrIntGroupSpec.lock.Lock()

	defer nStrIntGroupSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrIntegerGroupingSpec."+
			"SetStr()",
		"")

	if err != nil {
		return err
	}

	return new(numStrIntGroupingSpecNanobot).setNStrIntGroupSpec(
		nStrIntGroupSpec,
		[]rune(intSeparatorChars),
		intGroupingType,
		ePrefix.XCpy(
			"newNStrIntGroupSpec"))

}

// numStrIntGroupingSpecNanobot - This type provides
// helper methods for NumStrIntegerGroupingSpec
type numStrIntGroupingSpecNanobot struct {
	lock *sync.Mutex
}

// copyNStrIntGroupSpec - Copies all data from input parameter
// 'sourceIntGroupSpec' to input parameter
// 'destinationIntGroupSpec'. Both instances are of type
// NumStrIntegerGroupingSpec.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in
// 'destinationIntGroupSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceIntGroupSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationIntGroupSpec  *NumStrIntegerGroupingSpec
//	   - A pointer to a NumStrIntegerGroupingSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceIntGroupSpec'.
//
//	     'destinationIntGroupSpec' is the destination for this
//	     copy operation.
//
//
//	sourceIntGroupSpec       *NumStrIntegerGroupingSpec
//	   - A pointer to another NumStrIntegerGroupingSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationIntGroupSpec'.
//
//	     'sourceIntGroupSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceIntGroupSpec'.
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
func (nStrIntGroupSpecNanobot *numStrIntGroupingSpecNanobot) copyNStrIntGroupSpec(
	destinationIntGroupSpec *NumStrIntegerGroupingSpec,
	sourceIntGroupSpec *NumStrIntegerGroupingSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrIntGroupSpecNanobot.lock == nil {
		nStrIntGroupSpecNanobot.lock = new(sync.Mutex)
	}

	nStrIntGroupSpecNanobot.lock.Lock()

	defer nStrIntGroupSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrIntGroupingSpecNanobot."+
			"copyNStrIntGroupSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationIntGroupSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationIntGroupSpec' is invalid!\n"+
			"'destinationIntGroupSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceIntGroupSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceIntGroupSpec' is invalid!\n"+
			"'sourceIntGroupSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrIntGroupingSpecAtom).empty(
		destinationIntGroupSpec)

	err = destinationIntGroupSpec.integerSeparatorChars.
		CopyIn(
			&sourceIntGroupSpec.integerSeparatorChars,
			ePrefix.XCpy(
				"destinationIntGroupSpec.integerSeparatorChars<-"+
					"sourceIntGroupSpec"))

	if err != nil {
		return err
	}

	destinationIntGroupSpec.intGroupingType =
		sourceIntGroupSpec.intGroupingType

	return err
}

// setNStrIntGroupSpec - Deletes and resets all member variable
// data values in the NumStrIntegerGroupingSpec instance passed
// as input parameter 'nIntGroupSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'nIntGroupSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// intGroupingSpec             IntegerGroupingType
//
//   - This instance of IntegerGroupingType defines the type
//     of IntegerSeparatorSpec which will be returned. The
//     enumeration IntegerGroupingType must be set to one
//     of the following values:
//     IntGroupingType.None()
//     IntGroupingType.Thousands()
//     IntGroupingType.IndiaNumbering()
//     IntGroupingType.ChineseNumbering()
//
//     intSeparatorChars       []rune
//
//   - One or more characters used to separate groups of
//     integer digits. This separator is also known as the
//     'thousands' separator. It is used to separate groups of
//     integer digits to the left of the decimal separator
//     (a.k.a. decimal point).
//
//     If this input parameter contains a zero length array
//     and 'intGroupingSpec' is NOT equal to
//     'IntGroupingType.None()', an error will be returned.
//
//     errPrefDto                 *ErrPrefixDto
//
//   - This object encapsulates an error prefix string which is
//     included in all returned error messages. Usually, it
//     contains the names of the calling method or methods.
//
//     Type ErrPrefixDto is included in the 'errpref' software
//     package, "github.com/MikeAustin71/errpref".
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
//	     and text passed by input parameter, 'errPrefDto'. The
//	     'errPrefDto' text will be attached to the beginning of the
//	     error message.
func (nStrIntGroupSpecNanobot *numStrIntGroupingSpecNanobot) setNStrIntGroupSpec(
	nIntGroupSpec *NumStrIntegerGroupingSpec,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrIntGroupSpecNanobot.lock == nil {
		nStrIntGroupSpecNanobot.lock = new(sync.Mutex)
	}

	nStrIntGroupSpecNanobot.lock.Lock()

	defer nStrIntGroupSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrIntGroupingSpecNanobot."+
			"copyNStrIntGroupSpec()",
		"")

	if err != nil {
		return err
	}

	if nIntGroupSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nIntGroupSpec' is invalid!\n"+
			"'nIntGroupSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if !intGroupingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intGroupingType' is invalid!\n"+
			"'nIntGroupSpec' string value  = '%v'.\n"+
			"'nIntGroupSpec' integer value = '%v'.\n",
			ePrefix.String(),
			intGroupingType.String(),
			intGroupingType.XValueInt())

		return err
	}

	if intGroupingType != IntGroupingType.None() &&
		len(intSeparatorChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intGroupingType' is equal to IntGroupingType.None() and\n"+
			"'intSeparatorChars' has a length of zero.\n",
			ePrefix.String())

		return err

	}

	new(numStrIntGroupingSpecAtom).empty(
		nIntGroupSpec)

	if len(intSeparatorChars) > 0 {

		err = nIntGroupSpec.integerSeparatorChars.
			SetRuneArray(
				intSeparatorChars,
				ePrefix.XCpy(
					"nIntGroupSpec<-intSeparatorChars"))

		if err != nil {
			return err
		}

	}

	nIntGroupSpec.intGroupingType = intGroupingType

	return err
}

// numStrIntGroupingSpecAtom - This type provides
// helper methods for NumStrIntegerGroupingSpec
type numStrIntGroupingSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrIntegerGroupingSpec and proceeds to reset the
// data values for all member variables to their initial or
// zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrIntGroupSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntGroupSpec           *NumStrIntegerGroupingSpec
//	   - A pointer to an instance of NumStrIntegerGroupingSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrIntGroupSpecAtom *numStrIntGroupingSpecAtom) empty(
	nStrIntGroupSpec *NumStrIntegerGroupingSpec) {

	if nStrIntGroupSpecAtom.lock == nil {
		nStrIntGroupSpecAtom.lock = new(sync.Mutex)
	}

	nStrIntGroupSpecAtom.lock.Lock()

	defer nStrIntGroupSpecAtom.lock.Unlock()

	if nStrIntGroupSpec == nil {
		return
	}

	nStrIntGroupSpec.integerSeparatorChars.Empty()

	nStrIntGroupSpec.intGroupingType = IntGroupingType.None()
}

// equal - Receives a pointer to two instances of
// NumStrIntegerGroupingSpec and proceeds to compare their
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
//	nStrIntGroupSpec1    *NumStrIntegerGroupingSpec
//	   - An instance of NumStrIntegerGroupingSpec.
//	     Internal member variables from 'nStrIntGroupSpec1'
//	     will be compared to those of 'nStrIntGroupSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrIntGroupSpec2    *NumStrIntegerGroupingSpec
//	   - An instance of NumStrIntegerGroupingSpec.
//	     Internal member variables from 'nStrIntGroupSpec2'
//	     will be compared to those of 'nStrIntGroupSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrIntGroupSpec1' and
//	     'nStrIntGroupSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrIntGroupSpecAtom *numStrIntGroupingSpecAtom) equal(
	nStrIntGroupSpec1 *NumStrIntegerGroupingSpec,
	nStrIntGroupSpec2 *NumStrIntegerGroupingSpec) bool {

	if nStrIntGroupSpecAtom.lock == nil {
		nStrIntGroupSpecAtom.lock = new(sync.Mutex)
	}

	nStrIntGroupSpecAtom.lock.Lock()

	defer nStrIntGroupSpecAtom.lock.Unlock()

	if nStrIntGroupSpec1 == nil ||
		nStrIntGroupSpec2 == nil {

		return false
	}

	if !nStrIntGroupSpec1.integerSeparatorChars.Equal(
		&nStrIntGroupSpec2.integerSeparatorChars) {

		return false
	}

	if nStrIntGroupSpec1.intGroupingType !=
		nStrIntGroupSpec2.intGroupingType {

		return false
	}

	return true
}
