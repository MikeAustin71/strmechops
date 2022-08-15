package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrPositiveNumberSignSpec - Contains all the necessary
// parameters to insert and display a positive number sign
// or symbol for a positive numeric value formatted as a
// number string.
//
// Typically, positive number sigh symbols like the plus
// sign ('+') are not displayed in number strings. Rather,
// the absence of a positive number sign symbol and the
// absence of a negative number sign symbol implies that
// the numeric value is positive.
//
// However, the user has the option to customize this
// behavior by configuring leading or trailing number
// sign symbols.
//
// Different countries and cultures have different ways
// of defining positive number sign symbols. Typically,
// when a positive number sign symbol is specified, it
// is positioned either in front of the numeric value
// or after the numeric value. Again, users have the
// flexibility to specify either leading, trailing or
// both leading and trailing positive number sign
// symbols.
type NumStrPositiveNumberSignSpec struct {
	leadingPosNumSign RuneArrayDto
	// Contains the character or character which
	// will be formatted and displayed in front of
	// a positive numeric value displayed in a
	// number string.

	trailingPosNumSign RuneArrayDto
	// Contains the character or character which
	// will be formatted and displayed after a
	// positive numeric value displayed in a
	// number string.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrPositiveNumberSignSpec ('incomingNStrPosNumSignSpec')
// to the data fields of the current NumStrPositiveNumberSignSpec
// instance ('nStrPosNumSignSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrPositiveNumberSignSpec
// instance ('nStrPosNumSignSpec') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 incomingNStrPosNumSignSpec *NumStrPositiveNumberSignSpec
//	    - A pointer to an instance of NumStrPositiveNumberSignSpec.
//	      This method will NOT change the values of internal member
//	      variables contained in this instance.
//
//	      All data values in this NumStrPositiveNumberSignSpec instance
//	      will be copied to current NumStrPositiveNumberSignSpec
//	      instance ('nStrPosNumSignSpec').
//
//	      If parameter 'incomingNStrPosNumSignSpec' is determined to
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
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (nStrPosNumSignSpec *NumStrPositiveNumberSignSpec) CopyIn(
	incomingNStrPosNumSignSpec *NumStrPositiveNumberSignSpec,
	errorPrefix interface{}) error {

	if nStrPosNumSignSpec.lock == nil {
		nStrPosNumSignSpec.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpec.lock.Lock()

	defer nStrPosNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrPositiveNumberSignSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrPosNumSignSpecNanobot).copyNStrPosNumSignSpec(
		nStrPosNumSignSpec,
		incomingNStrPosNumSignSpec,
		ePrefix.XCpy(
			"nStrPosNumSignSpec<-"+
				"incomingNStrPosNumSignSpec"))
}

// CopyOut - Returns a deep copy of the current
// NumStrPositiveNumberSignSpec instance.
//
// If the current NumStrPositiveNumberSignSpec instance contains
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
//	deepCopyNStrPosNumSignSpec NumStrPositiveNumberSignSpec
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current NumStrPositiveNumberSignSpec instance.
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
func (nStrPosNumSignSpec *NumStrPositiveNumberSignSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNStrPosNumSignSpec *NumStrPositiveNumberSignSpec,
	err error) {

	if nStrPosNumSignSpec.lock == nil {
		nStrPosNumSignSpec.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpec.lock.Lock()

	defer nStrPosNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrPositiveNumberSignSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNStrPosNumSignSpec, err
	}

	err = new(numStrPosNumSignSpecNanobot).copyNStrPosNumSignSpec(
		deepCopyNStrPosNumSignSpec,
		nStrPosNumSignSpec,
		ePrefix.XCpy(
			"deepCopyNStrPosNumSignSpec<-"+
				"nStrPosNumSignSpec"))

	return deepCopyNStrPosNumSignSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrPositiveNumberSignSpec to their zero or
// uninitialized states. This method will leave the current
// instance of NumStrPositiveNumberSignSpec in an invalid state
// and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of NumStrPositiveNumberSignSpec. All member
// variable data values will be reset to their zero or
// uninitialized states.
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
func (nStrPosNumSignSpec *NumStrPositiveNumberSignSpec) Empty() {

	if nStrPosNumSignSpec.lock == nil {
		nStrPosNumSignSpec.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpec.lock.Lock()

	new(numStrPosNumSignSpecAtom).empty(
		nStrPosNumSignSpec)

	nStrPosNumSignSpec.lock.Unlock()

	nStrPosNumSignSpec.lock = nil
}

// EmptyLeadingPosNumSign - Resets the member variable data for
// the leading positive number sign contained in the current
// instance of NumStrPositiveNumberSignSpec to an initial or
// zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading positive
// number sign contained in the current instance of
// NumStrPositiveNumberSignSpec will be deleted and reset
// to an empty or zero value.
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
func (nStrPosNumSignSpec *NumStrPositiveNumberSignSpec) EmptyLeadingPosNumSign() {

	if nStrPosNumSignSpec.lock == nil {
		nStrPosNumSignSpec.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpec.lock.Lock()

	defer nStrPosNumSignSpec.lock.Unlock()

	new(numStrPosNumSignSpecAtom).emptyLeadingPosNumSign(
		nStrPosNumSignSpec)
}

// EmptyTrailingPosNumSign - Resets the member variable data for
// the trailing positive number sign contained in the current
// instance of NumStrPositiveNumberSignSpec to an initial or
// zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing positive
// number sign contained in the current instance of
// NumStrPositiveNumberSignSpec will be deleted and reset
// to an empty or zero value.
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
func (nStrPosNumSignSpec *NumStrPositiveNumberSignSpec) EmptyTrailingPosNumSign() {

	if nStrPosNumSignSpec.lock == nil {
		nStrPosNumSignSpec.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpec.lock.Lock()

	defer nStrPosNumSignSpec.lock.Unlock()

	new(numStrPosNumSignSpecAtom).emptyTrailingPosNumSign(
		nStrPosNumSignSpec)
}

// numStrPosNumSignSpecNanobot - This type provides
// helper methods for NumStrPositiveNumberSignSpec
type numStrPosNumSignSpecNanobot struct {
	lock *sync.Mutex
}

// copyNStrPosNumSignSpec - Copies all data from input parameter
// 'sourcePosNumSignSpec' to input parameter
// 'destinationPosNumSignSpec'. Both instances are of type
// NumStrPositiveNumberSignSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationPosNumSignSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourcePosNumSignSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationPosNumSignSpec  *NumStrPositiveNumberSignSpec
//	   - A pointer to a NumStrPositiveNumberSignSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourcePosNumSignSpec'.
//
//	     'destinationPosNumSignSpec' is the destination for this
//	     copy operation.
//
//
//	sourcePosNumSignSpec       *NumStrPositiveNumberSignSpec
//	   - A pointer to another NumStrPositiveNumberSignSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationPosNumSignSpec'.
//
//	     'sourcePosNumSignSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourcePosNumSignSpec'.
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
func (nStrPosNumSignSpecNanobot *numStrPosNumSignSpecNanobot) copyNStrPosNumSignSpec(
	destinationPosNumSignSpec *NumStrPositiveNumberSignSpec,
	sourcePosNumSignSpec *NumStrPositiveNumberSignSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrPosNumSignSpecNanobot.lock == nil {
		nStrPosNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpecNanobot.lock.Lock()

	defer nStrPosNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrPosNumSignSpecNanobot."+
			"copyNStrPosNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationPosNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationPosNumSignSpec' is invalid!\n"+
			"'destinationPosNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourcePosNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourcePosNumSignSpec' is invalid!\n"+
			"'sourcePosNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrPosNumSignSpecAtom).empty(
		destinationPosNumSignSpec)

	err = destinationPosNumSignSpec.leadingPosNumSign.
		CopyIn(
			&sourcePosNumSignSpec.leadingPosNumSign,
			ePrefix.XCpy(
				"destinationPosNumSignSpec.leadingPosNumSign<-"+
					"sourcePosNumSignSpec"))

	if err != nil {
		return err
	}

	err = destinationPosNumSignSpec.trailingPosNumSign.
		CopyIn(
			&sourcePosNumSignSpec.trailingPosNumSign,
			ePrefix.XCpy(
				"destinationPosNumSignSpec.trailingPosNumSign<-"+
					"sourcePosNumSignSpec"))

	return err
}

// setLeadingNStrPosNumSignSpec - Deletes and resets the data
// value of the Leading Positive Number Sign contained in an
// instance of NumStrPositiveNumberSignSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		posNumSignSpec             *NumStrPositiveNumberSignSpec
//		   - A pointer to a NumStrPositiveNumberSignSpec instance.
//		     The Leading Positive Number Sign contained in this
//	      instance will be deleted and reset to the value
//	      specified by input parameter, ''.
//
//
//		leadingPosNumSign          []rune
//		   - An array of runes specifying the character or
//		     characters which will be copied to the Leading
//		     Positive Number Sign contained in input parameter,
//	      'posNumSignSpec'.
//
//
//		errPrefDto                 *ePref.ErrPrefixDto
//		   - This object encapsulates an error prefix string which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods listed
//		     as a function chain.
//
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
//
//		     Type ErrPrefixDto is included in the 'errpref' software
//		     package, "github.com/MikeAustin71/errpref".
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
func (nStrPosNumSignSpecNanobot *numStrPosNumSignSpecNanobot) setLeadingNStrPosNumSignSpec(
	posNumSignSpec *NumStrPositiveNumberSignSpec,
	leadingPosNumSign []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrPosNumSignSpecNanobot.lock == nil {
		nStrPosNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpecNanobot.lock.Lock()

	defer nStrPosNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrPosNumSignSpecNanobot."+
			"setLeadingNStrPosNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if posNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'posNumSignSpec' is invalid!\n"+
			"'posNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrPosNumSignSpecAtom).emptyLeadingPosNumSign(
		posNumSignSpec)

	if len(leadingPosNumSign) == 0 {
		return err
	}

	err = posNumSignSpec.leadingPosNumSign.SetRuneArray(
		leadingPosNumSign,
		ePrefix.XCpy(
			"posNumSignSpec.leadingPosNumSign"+
				"<-leadingPosNumSign"))

	return err
}

// setTrailingNStrPosNumSignSpec - Deletes and resets the data
// value of the Trailing Positive Number Sign contained in an
// instance of NumStrPositiveNumberSignSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		posNumSignSpec             *NumStrPositiveNumberSignSpec
//		   - A pointer to a NumStrPositiveNumberSignSpec instance.
//		     The Trailing Positive Number Sign contained in this
//	      instance will be deleted and reset to the value
//	      specified by input parameter, ''.
//
//
//		trailingPosNumSign          []rune
//		   - An array of runes specifying the character or
//		     characters which will be copied to the Trailing
//		     Positive Number Sign contained in input parameter,
//	      'posNumSignSpec'.
//
//
//		errPrefDto                 *ePref.ErrPrefixDto
//		   - This object encapsulates an error prefix string which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods listed
//		     as a function chain.
//
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
//
//		     Type ErrPrefixDto is included in the 'errpref' software
//		     package, "github.com/MikeAustin71/errpref".
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
func (nStrPosNumSignSpecNanobot *numStrPosNumSignSpecNanobot) setTrailingNStrPosNumSignSpec(
	posNumSignSpec *NumStrPositiveNumberSignSpec,
	trailingPosNumSign []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrPosNumSignSpecNanobot.lock == nil {
		nStrPosNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrPosNumSignSpecNanobot.lock.Lock()

	defer nStrPosNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrPosNumSignSpecNanobot."+
			"setTrailingNStrPosNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if posNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'posNumSignSpec' is invalid!\n"+
			"'posNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrPosNumSignSpecAtom).emptyTrailingPosNumSign(
		posNumSignSpec)

	if len(trailingPosNumSign) == 0 {
		return err
	}

	err = posNumSignSpec.trailingPosNumSign.SetRuneArray(
		trailingPosNumSign,
		ePrefix.XCpy(
			"posNumSignSpec.trailingPosNumSign"+
				"<-trailingPosNumSign"))

	return err
}

// numStrPosNumSignSpecAtom - This type provides
// helper methods for NumStrPositiveNumberSignSpec
type numStrPosNumSignSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrPositiveNumberSignSpec and proceeds to reset the
// data values for all member variables to their initial or
// zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrPosNumSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	nStrPosNumSpec           *NumStrPositiveNumberSignSpec
//	   - A pointer to an instance of NumStrPositiveNumberSignSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrPosNumSpecAtom *numStrPosNumSignSpecAtom) empty(
	nStrPosNumSpec *NumStrPositiveNumberSignSpec) {

	if nStrPosNumSpecAtom.lock == nil {
		nStrPosNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrPosNumSpecAtom.lock.Lock()

	defer nStrPosNumSpecAtom.lock.Unlock()

	if nStrPosNumSpec == nil {
		return
	}

	nStrPosNumSpec.leadingPosNumSign.Empty()

	nStrPosNumSpec.trailingPosNumSign.Empty()
}

// emptyLeadingPosNumSign - Receives a pointer to an instance
// of NumStrPositiveNumberSignSpec and proceeds to reset the
// member variable data for the leading positive number sign
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading positive
// number sign contained in input parameter 'nStrPosNumSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrPosNumSpec           *NumStrPositiveNumberSignSpec
//	   - A pointer to an instance of NumStrPositiveNumberSignSpec.
//	     The Leading Positive Number Sign contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrPosNumSpecAtom *numStrPosNumSignSpecAtom) emptyLeadingPosNumSign(
	nStrPosNumSpec *NumStrPositiveNumberSignSpec) {

	if nStrPosNumSpecAtom.lock == nil {
		nStrPosNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrPosNumSpecAtom.lock.Lock()

	defer nStrPosNumSpecAtom.lock.Unlock()

	if nStrPosNumSpec == nil {
		return
	}

	nStrPosNumSpec.leadingPosNumSign.Empty()

}

// emptyTrailingPosNumSign - Receives a pointer to an instance
// of NumStrPositiveNumberSignSpec and proceeds to reset the
// member variable data for the trailing positive number sign
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing positive
// number sign contained in input parameter 'nStrPosNumSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrPosNumSpec           *NumStrPositiveNumberSignSpec
//	   - A pointer to an instance of NumStrPositiveNumberSignSpec.
//	     The Trailing Positive Number Sign contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrPosNumSpecAtom *numStrPosNumSignSpecAtom) emptyTrailingPosNumSign(
	nStrPosNumSpec *NumStrPositiveNumberSignSpec) {

	if nStrPosNumSpecAtom.lock == nil {
		nStrPosNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrPosNumSpecAtom.lock.Lock()

	defer nStrPosNumSpecAtom.lock.Unlock()

	if nStrPosNumSpec == nil {
		return
	}

	nStrPosNumSpec.trailingPosNumSign.Empty()
}

// equal - Receives a pointer to two instances of
// NumStrPositiveNumberSignSpec and proceeds to compare their
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
//	nStrPosNumSpec1    *NumStrPositiveNumberSignSpec
//	   - An instance of NumStrPositiveNumberSignSpec.
//	     Internal member variables from 'nStrPosNumSpec1'
//	     will be compared to those of 'nStrPosNumSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrPosNumSpec2    *NumStrPositiveNumberSignSpec
//	   - An instance of NumStrPositiveNumberSignSpec.
//	     Internal member variables from 'nStrPosNumSpec2'
//	     will be compared to those of 'nStrPosNumSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrPosNumSpec1' and
//	     'nStrPosNumSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrPosNumSpecAtom *numStrPosNumSignSpecAtom) equal(
	nStrPosNumSpec1 *NumStrPositiveNumberSignSpec,
	nStrPosNumSpec2 *NumStrPositiveNumberSignSpec) bool {

	if nStrPosNumSpecAtom.lock == nil {
		nStrPosNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrPosNumSpecAtom.lock.Lock()

	defer nStrPosNumSpecAtom.lock.Unlock()

	if nStrPosNumSpec1 == nil ||
		nStrPosNumSpec2 == nil {
		return false
	}

	if !nStrPosNumSpec1.leadingPosNumSign.Equal(
		&nStrPosNumSpec2.leadingPosNumSign) {

		return false
	}

	if !nStrPosNumSpec1.trailingPosNumSign.Equal(
		&nStrPosNumSpec2.trailingPosNumSign) {

		return false
	}

	return true
}
