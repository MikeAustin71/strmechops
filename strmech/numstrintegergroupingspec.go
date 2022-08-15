package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrIntegerGroupingSpec - Number String Integer Grouping
// Specification. This type encapsulates the parameters required
// to format integer grouping and separation within a number
// string.
type NumStrIntegerGroupingSpec struct {
	integerSeparatorChars RuneArrayDto

	intGroupingType IntegerGroupingType

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
// Input Parameters
//
//	 incomingIntGroupSpec   *NumStrIntegerGroupingSpec
//	    - A pointer to an instance of NumStrIntegerGroupingSpec.
//	      This method will NOT change the values of internal member
//	      variables contained in this instance.
//
//	      All data values in this NumStrIntegerGroupingSpec instance
//	      will be copied to current NumStrIntegerGroupingSpec
//	      instance ('nStrIntGroupSpec').
//
//	      If parameter 'incomingIntGroupSpec' is determined to
//	      be invalid, an error will be returned.
//
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
// Return Values
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
// intGroupingType             IntegerGroupingType
//
//   - This instance of IntegerGroupingType defines the type
//     of IntegerSeparatorDto which will be returned. The
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
//     (a.k.a. decimal point). In the United States, the
//     standard integer digits separator is the comma (",").
//     United States Example:  1,000,000,000
//
//     In many European countries, a single period ('.') is used
//     as the integer separator character.
//     European Example: 1.000.000.000
//
//     Other countries and cultures use spaces, apostrophes or
//     multiple characters to separate integers.
//
//     If this input parameter contains a zero length array, an
//     error will be returned.
//
//     errorPrefix                interface{}
//
//   - This object encapsulates error prefix text which is
//     included in all returned error messages. Usually, it
//     contains the name of the calling method or methods
//     listed as a method or function chain of execution.
//
//     If no error prefix information is needed, set this parameter
//     to 'nil'.
//
//     This empty interface must be convertible to one of the
//     following types:
//
//     1. nil - A nil value is valid and generates an empty
//     collection of error prefix and error context
//     information.
//
//     2. string - A string containing error prefix information.
//
//     3. []string A one-dimensional slice of strings containing
//     error prefix information
//
//     4. [][2]string A two-dimensional slice of strings
//     containing error prefix and error context information.
//
//     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//     from this object will be copied for use in error and
//     informational messages.
//
//     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//     Information from this object will be copied for use in
//     error and informational messages.
//
//     7. IBasicErrorPrefix - An interface to a method generating
//     a two-dimensional slice of strings containing error
//     prefix and error context information.
//
//     If parameter 'errorPrefix' is NOT convertible to one of
//     the valid types listed above, it will be considered
//     invalid and trigger the return of an error.
//
//     Types ErrPrefixDto and IBasicErrorPrefix are included in
//     the 'errpref' software package,
//     "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	newNStrIntGroupSpec        NumStrIntegerGroupingSpec
//	   - If this method completes successfully, a new instance of
//	     NumStrIntegerGroupingSpec will be created and returned.
//	     This integer grouping specification will be configured to
//	     'Thousands', 'India Numbering System' or 'Chinese
//	     Numbering' depending on the specification provided by
//	     input parameter, 'intGroupingType'.
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
//	intGroupingType             IntegerGroupingType
//	 - This instance of IntegerGroupingType defines the type
//	   of IntegerSeparatorDto which will be returned. The
//	   enumeration IntegerGroupingType must be set to one
//	   of the following values:
//	   IntGroupingType.None()
//	   IntGroupingType.Thousands()
//	   IntGroupingType.IndiaNumbering()
//	   IntGroupingType.ChineseNumbering()
//
//	intSeparatorChars          string
//	 - One or more characters used to separate groups of
//	   integers. This separator is also known as the 'thousands'
//	   separator. It is used to separate groups of integer digits
//	   to the left of the decimal separator
//	   (a.k.a. decimal point). In the United States, the standard
//	   integer digits separator is the comma (",").
//	   United States Example:  1,000,000,000
//
//	   In many European countries, a single period ('.') is used
//	   as the integer separator character.
//	   European Example: 1.000.000.000
//
//	   Other countries and cultures use spaces, apostrophes or
//	   multiple characters to separate integers.
//
//	   If this input parameter contains a zero length string, an
//	   error will be returned.
//
//	errorPrefix                interface{}
//	 - This object encapsulates error prefix text which is
//	   included in all returned error messages. Usually, it
//	   contains the name of the calling method or methods
//	   listed as a method or function chain of execution.
//
//	   If no error prefix information is needed, set this parameter
//	   to 'nil'.
//
//	   This empty interface must be convertible to one of the
//	   following types:
//
//	   1. nil - A nil value is valid and generates an empty
//	   collection of error prefix and error context
//	   information.
//
//	   2. string - A string containing error prefix information.
//
//	   3. []string A one-dimensional slice of strings containing
//	   error prefix information
//
//	   4. [][2]string A two-dimensional slice of strings
//	   containing error prefix and error context information.
//
//	   5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	   from this object will be copied for use in error and
//	   informational messages.
//
//	   6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	   Information from this object will be copied for use in
//	   error and informational messages.
//
//	   7. IBasicErrorPrefix - An interface to a method generating
//	   a two-dimensional slice of strings containing error
//	   prefix and error context information.
//
//	   If parameter 'errorPrefix' is NOT convertible to one of
//	   the valid types listed above, it will be considered
//	   invalid and trigger the return of an error.
//
//	   Types ErrPrefixDto and IBasicErrorPrefix are included in
//	   the 'errpref' software package,
//	   "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	newNStrIntGroupSpec        NumStrIntegerGroupingSpec
//	   - If this method completes successfully, a new instance of
//	     NumStrIntegerGroupingSpec will be created and returned.
//	     This integer grouping specification will be configured to
//	     'Thousands', 'India Numbering System' or 'Chinese
//	     Numbering' depending on the specification provided by
//	     input parameter, 'intGroupingType'.
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
func (nStrIntGroupSpec *NumStrIntegerGroupingSpec) NewStr(
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
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

	new(numStrIntGroupingSpecAtom).empty(
		nIntGroupSpec)

	err = nIntGroupSpec.integerSeparatorChars.
		SetRuneArray(
			intSeparatorChars,
			ePrefix.XCpy(
				"nIntGroupSpec<-intSeparatorChars"))

	if err != nil {
		return err
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
