package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrNegativeNumberSignSpec - Contains all the necessary
// parameters to insert and display a negative number sign
// or symbol for a negative numeric value formatted as a
// number string.
//
// Different countries and cultures apply different
// definitions for negative number sign symbols. Typically,
// when a negative number sign symbol is specified, it
// is positioned in front of the numeric value, after the
// numeric value or before and after the numeric value.
type NumStrNegativeNumberSignSpec struct {
	leadingNegNumSign RuneArrayDto
	// Contains the character or character which
	// will be formatted and displayed in front of
	// a negative numeric value displayed in a
	// number string.

	trailingNegNumSign RuneArrayDto
	// Contains the character or character which
	// will be formatted and displayed after a
	// negative numeric value displayed in a
	// number string.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrNegativeNumberSignSpec ('incomingNStrNegNumSignSpec')
// to the data fields of the current NumStrNegativeNumberSignSpec
// instance ('nStrNegNumSignSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrNegativeNumberSignSpec
// instance ('nStrNegNumSignSpec') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNStrNegNumSignSpec	*NumStrNegativeNumberSignSpec
//		A pointer to an instance of NumStrNegativeNumberSignSpec.
//		This method will NOT change the values of internal member
//		variables contained in this instance.
//
//		All data values in this NumStrNegativeNumberSignSpec instance
//		will be copied to current NumStrNegativeNumberSignSpec
//		instance ('nStrNegNumSignSpec').
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
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) CopyIn(
	incomingNStrNegNumSignSpec *NumStrNegativeNumberSignSpec,
	errorPrefix interface{}) error {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNegNumSignSpecNanobot).copyNStrNegNumSignSpec(
		nStrNegNumSignSpec,
		incomingNStrNegNumSignSpec,
		ePrefix.XCpy(
			"nStrNegNumSignSpec<-"+
				"incomingNStrNegNumSignSpec"))
}

// CopyOut - Returns a deep copy of the current
// NumStrNegativeNumberSignSpec instance.
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
//	deepCopyNStrNegNumSignSpec NumStrNegativeNumberSignSpec
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current NumStrNegativeNumberSignSpec instance.
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
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNStrNegNumSignSpec *NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNStrNegNumSignSpec, err
	}

	err = new(numStrNegNumSignSpecNanobot).copyNStrNegNumSignSpec(
		deepCopyNStrNegNumSignSpec,
		nStrNegNumSignSpec,
		ePrefix.XCpy(
			"deepCopyNStrNegNumSignSpec<-"+
				"nStrNegNumSignSpec"))

	return deepCopyNStrNegNumSignSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrNegativeNumberSignSpec to their zero or
// uninitialized states. This method will leave the current
// instance of NumStrNegativeNumberSignSpec in an invalid state
// and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of NumStrNegativeNumberSignSpec. All member
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
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) Empty() {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	new(numStrNegNumSignSpecAtom).empty(
		nStrNegNumSignSpec)

	nStrNegNumSignSpec.lock.Unlock()

	nStrNegNumSignSpec.lock = nil
}

// EmptyLeadingNegNumSign - Resets the member variable data for
// the leading negative number sign contained in the current
// instance of NumStrNegativeNumberSignSpec to an initial or
// zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading negative
// number sign contained in the current instance of
// NumStrNegativeNumberSignSpec will be deleted and reset
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
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) EmptyLeadingNegNumSign() {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	new(numStrNegNumSignSpecAtom).emptyLeadingNegNumSign(
		nStrNegNumSignSpec)
}

// EmptyTrailingNegNumSign - Resets the member variable data for
// the trailing negative number sign contained in the current
// instance of NumStrNegativeNumberSignSpec to an initial or
// zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing negative
// number sign contained in the current instance of
// NumStrNegativeNumberSignSpec will be deleted and reset
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
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) EmptyTrailingNegNumSign() {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	new(numStrNegNumSignSpecAtom).emptyTrailingNegNumSign(
		nStrNegNumSignSpec)
}

// Equal - Receives a pointer to an NumStrNegativeNumberSignSpec
// object and proceeds to determine whether all data elements in
// this object are equal to all corresponding data elements in
// the current instance of NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		incomingNStrNegNumSignSpec *NumStrNegativeNumberSignSpec
//	    - This method will compare all data elements in the
//	      current NumStrNegativeNumberSignSpec object to
//	      corresponding data elements in this second
//	      NumStrNegativeNumberSignSpec object in order
//	      to determine equivalency.
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
//	   - If all the data elements in the current
//	     NumStrNegativeNumberSignSpec instance are equal to all the
//	     corresponding data elements in 'incomingNStrNegNumSignSpec',
//	     this return parameter will be set to 'true'. If all the data
//	     elements are NOT equal, this return parameter will be set to
//	     'false'.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) Equal(
	incomingNStrNegNumSignSpec *NumStrNegativeNumberSignSpec) bool {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	return new(numStrNegNumSignSpecAtom).equal(
		nStrNegNumSignSpec,
		incomingNStrNegNumSignSpec)
}

// GetLeadingNegNumSignStr - Returns a string containing the
// leading negative number sign character or characters.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) GetLeadingNegNumSignStr() string {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	return nStrNegNumSignSpec.leadingNegNumSign.GetCharacterString()
}

// GetTrailingNegNumSignStr - Returns a string containing the
// trailing negative number sign character or characters.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) GetTrailingNegNumSignStr() string {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	return nStrNegNumSignSpec.trailingNegNumSign.GetCharacterString()
}

// NewLeadingNegNumberSign - Creates and returns a new instance
// of NumStrNegativeNumberSignSpec configured with a leading
// negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leadingNegativeNumSign     string
//		   - A string containing the leading negative number sign
//		     character or characters used to configure the returned
//		     instance of NumStrNegativeNumberSignSpec.
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
//	newNStrNegNumSign          NumStrNegativeNumberSignSpec
//	   - If this method completes successfully, a new instance of
//		 NumStrNegativeNumberSignSpec, configured with a leading
//		 negative number sign symbol, will be returned through this
//	     parameter.
//
//	err                        error
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) NewLeadingNegNumberSign(
	leadingNegativeNumSign string,
	errorPrefix interface{}) (
	newNStrNegNumSign NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"NewLeadingNegNumberSign()",
		"")

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setLeadingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			[]rune(leadingNegativeNumSign),
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"leadingNegativeNumSign"))

	return newNStrNegNumSign, err
}

// NewLeadingNegNumberSignRunes - Creates and returns a new
// instance of NumStrNegativeNumberSignSpec configured with
// a leading negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leadingNegativeNumSign     []rune
//		   - An array of runes containing the leading negative number
//		     sign character or characters used to configure the
//		     returned instance of NumStrNegativeNumberSignSpec.
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
//	newNStrNegNumSign          NumStrNegativeNumberSignSpec
//	   - If this method completes successfully, a new instance of
//		 NumStrNegativeNumberSignSpec, configured with a leading
//		 negative number sign symbol, will be returned through this
//	     parameter.
//
//	err                        error
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) NewLeadingNegNumberSignRunes(
	leadingNegativeNumSign []rune,
	errorPrefix interface{}) (
	newNStrNegNumSign NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"NewLeadingNegNumberSignRunes()",
		"")

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setLeadingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			leadingNegativeNumSign,
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"leadingNegativeNumSign"))

	return newNStrNegNumSign, err
}

// NewTrailingNegNumberSign - Creates and returns a new
// instance of NumStrNegativeNumberSignSpec configured with a
// trailing negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	trailingNegativeNumSign		string
//		A string containing the trailing negative number
//		sign character or characters used to configure the
//		returned instance of NumStrNegativeNumberSignSpec.
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
//	newNStrNegNumSign		NumStrNegativeNumberSignSpec
//		If this method completes successfully, a new instance of
//		NumStrNegativeNumberSignSpec, configured with a trailing
//		negative number sign symbol, will be returned through this
//		parameter.
//
//	err							error
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) NewTrailingNegNumberSign(
	trailingNegativeNumSign string,
	errorPrefix interface{}) (
	newNStrNegNumSign NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"NewTrailingNegNumberSign()",
		"")

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setTrailingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			[]rune(trailingNegativeNumSign),
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"trailingNegativeNumSign"))

	return newNStrNegNumSign, err
}

// NewTrailingNegNumberSignRunes - Creates and returns a new
// instance of NumStrNegativeNumberSignSpec configured with a
// trailing negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		trailingNegativeNumSign    []rune
//		   - An array of runes containing the trailing negative number
//		     sign character or characters used to configure the
//		     returned instance of NumStrNegativeNumberSignSpec.
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
//	newNStrNegNumSign          NumStrNegativeNumberSignSpec
//	   - If this method completes successfully, a new instance of
//		 NumStrNegativeNumberSignSpec, configured with a trailing
//		 negative number sign symbol, will be returned through this
//	     parameter.
//
//	err                        error
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) NewTrailingNegNumberSignRunes(
	trailingNegativeNumSign []rune,
	errorPrefix interface{}) (
	newNStrNegNumSign NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"NewTrailingNegNumberSignRunes()",
		"")

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setTrailingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			trailingNegativeNumSign,
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"trailingNegativeNumSign"))

	return newNStrNegNumSign, err
}

// NewLeadingTrailingNegNumberSignRunes - Creates and returns a
// new instance of NumStrNegativeNumberSignSpec configured with
// both leading and trailing negative number sign symbols.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leadingNegativeNumSign		[]rune
//		An array of runes containing the leading negative
//		number sign character or characters used to
//		configure the returned instance of
//		NumStrNegativeNumberSignSpec.
//
//	trailingNegativeNumSign		[]rune
//		An array of runes containing the trailing negative
//		number sign character or characters used to
//		configure the returned instance of
//		NumStrNegativeNumberSignSpec.
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
//	newNStrNegNumSignNumStr		NegativeNumberSignSpec
//		If this method completes successfully, a new instance
//		of NumStrNegativeNumberSignSpec, configured with both
//		leading and trailing negative number sign symbols,
//		will be returned through this parameter.
//
//	err							error
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) NewLeadingTrailingNegNumberSignRunes(
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	errorPrefix interface{}) (
	newNStrNegNumSign NumStrNegativeNumberSignSpec,
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"NewTrailingNegNumberSignRunes()",
		"")

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setLeadingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			leadingNegativeNumSign,
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"leadingNegativeNumSign"))

	if err != nil {
		return newNStrNegNumSign, err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setTrailingNStrNegNumSignSpec(
			&newNStrNegNumSign,
			trailingNegativeNumSign,
			ePrefix.XCpy(
				"newNStrNegNumSign<-"+
					"trailingNegativeNumSign"))

	return newNStrNegNumSign, err
}

// SetLeadingNegNumberSign - Resets and configures a leading
// negative number sign character or characters for the current
// instance of NumStrNegativeNumberSignSpec
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the leading negative
// number sign symbol data value in the current instance of
// NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leadingNegativeNumSign     string
//	    - A string containing the leading negative number sign
//			 character or characters used to configure the current
//			 instance of NumStrNegativeNumberSignSpec.
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
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) SetLeadingNegNumberSign(
	leadingNegativeNumSign string,
	errorPrefix interface{}) error {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"SetLeadingNegNumberSign()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setLeadingNStrNegNumSignSpec(
			nStrNegNumSignSpec,
			[]rune(leadingNegativeNumSign),
			ePrefix.XCpy(
				"nStrNegNumSignSpec<-"+
					"leadingNegativeNumSign"))

	return err
}

// SetLeadingNegNumberSignRunes - Resets and configures a leading
// negative number sign character or characters for the current
// instance of NumStrNegativeNumberSignSpec
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the leading negative
// number sign symbol data value in the current instance of
// NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leadingNegativeNumSign     []rune
//	    - An array of runes containing the leading negative number
//	      sign character or characters used to configure the current
//	      instance of NumStrNegativeNumberSignSpec.
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
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) SetLeadingNegNumberSignRunes(
	leadingNegativeNumSign []rune,
	errorPrefix interface{}) error {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"SetLeadingNegNumberSignRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setLeadingNStrNegNumSignSpec(
			nStrNegNumSignSpec,
			leadingNegativeNumSign,
			ePrefix.XCpy(
				"nStrNegNumSignSpec<-"+
					"leadingNegativeNumSign"))

	return err
}

// SetTrailingNegNumberSign - Creates and returns a new instance
// of NumStrNegativeNumberSignSpec configured with a trailing
// negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the trailing negative
// number sign symbol data value in the current instance of
// NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 trailingNegativeNumSign     string
//	    - A string containing the trailing negative number sign
//	      character or characters used to configure the returned
//	      instance of SetTrailingNegNumberSign.
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
//	newNStrNegNumSign          NumStrNegativeNumberSignSpec
//	   - If this method completes successfully, a new instance of
//		 NumStrNegativeNumberSignSpec, configured with a trailing
//		 negative number sign symbol, will be returned through this
//	     parameter.
//
//	err                        error
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) SetTrailingNegNumberSign(
	trailingNegativeNumSign string,
	errorPrefix interface{}) (
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"SetTrailingNegNumberSign()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setTrailingNStrNegNumSignSpec(
			nStrNegNumSignSpec,
			[]rune(trailingNegativeNumSign),
			ePrefix.XCpy(
				"nStrNegNumSignSpec<-"+
					"trailingNegativeNumSign"))

	return err
}

// SetTrailingNegNumberSignRunes - Creates and returns a new
// instance of NumStrNegativeNumberSignSpec configured with a
// trailing negative number sign character or characters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the trailing negative
// number sign symbol data value in the current instance of
// NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		trailingNegativeNumSign     []rune
//	    - An array of runes containing the trailing negative
//	      number sign character or characters used to configure
//	      the returned instance of SetTrailingNegNumberSignRunes.
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
//	newNStrNegNumSign          NumStrNegativeNumberSignSpec
//	   - If this method completes successfully, a new instance of
//		 NumStrNegativeNumberSignSpec, configured with a trailing
//		 negative number sign symbol, will be returned through this
//	     parameter.
//
//	err                        error
//	   - If this method completes successfully, the returned error
//		 Type is set equal to 'nil'.
//
//		 If errors are encountered during processing, the returned
//		 error Type will encapsulate an error message. This
//		 returned error message will incorporate the method chain
//		 and text passed by input parameter, 'errorPrefix'. The
//		 'errorPrefix' text will be attached to the beginning of
//		 the error message.
func (nStrNegNumSignSpec *NumStrNegativeNumberSignSpec) SetTrailingNegNumberSignRunes(
	trailingNegativeNumSign []rune,
	errorPrefix interface{}) (
	err error) {

	if nStrNegNumSignSpec.lock == nil {
		nStrNegNumSignSpec.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpec.lock.Lock()

	defer nStrNegNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNegativeNumberSignSpec."+
			"SetTrailingNegNumberSignRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNegNumSignSpecNanobot).
		setTrailingNStrNegNumSignSpec(
			nStrNegNumSignSpec,
			trailingNegativeNumSign,
			ePrefix.XCpy(
				"nStrNegNumSignSpec<-"+
					"trailingNegativeNumSign"))

	return err
}

// numStrNegNumSignSpecNanobot - This type provides
// helper methods for NumStrNegativeNumberSignSpec
type numStrNegNumSignSpecNanobot struct {
	lock *sync.Mutex
}

// copyNStrNegNumSignSpec - Copies all data from input parameter
// 'sourceNegNumSignSpec' to input parameter
// 'destinationNegNumSignSpec'. Both instances are of type
// NumStrNegativeNumberSignSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationNegNumSignSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceNegNumSignSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNegNumSignSpec  *NumStrNegativeNumberSignSpec
//	   - A pointer to a NumStrNegativeNumberSignSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNegNumSignSpec'.
//
//	     'destinationNegNumSignSpec' is the destination for this
//	     copy operation.
//
//
//	sourceNegNumSignSpec       *NumStrNegativeNumberSignSpec
//	   - A pointer to another NumStrNegativeNumberSignSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationNegNumSignSpec'.
//
//	     'sourceNegNumSignSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceNegNumSignSpec'.
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
func (nStrNegNumSignSpecNanobot *numStrNegNumSignSpecNanobot) copyNStrNegNumSignSpec(
	destinationNegNumSignSpec *NumStrNegativeNumberSignSpec,
	sourceNegNumSignSpec *NumStrNegativeNumberSignSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNegNumSignSpecNanobot.lock == nil {
		nStrNegNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpecNanobot.lock.Lock()

	defer nStrNegNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNegNumSignSpecNanobot."+
			"copyNStrNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNegNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNegNumSignSpec' is invalid!\n"+
			"'destinationNegNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNegNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNegNumSignSpec' is invalid!\n"+
			"'sourceNegNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrNegNumSignSpecAtom).empty(
		destinationNegNumSignSpec)

	err = destinationNegNumSignSpec.leadingNegNumSign.
		CopyIn(
			&sourceNegNumSignSpec.leadingNegNumSign,
			ePrefix.XCpy(
				"destinationNegNumSignSpec.leadingNegNumSign<-"+
					"sourceNegNumSignSpec"))

	if err != nil {
		return err
	}

	err = destinationNegNumSignSpec.trailingNegNumSign.
		CopyIn(
			&sourceNegNumSignSpec.trailingNegNumSign,
			ePrefix.XCpy(
				"destinationNegNumSignSpec.trailingNegNumSign<-"+
					"sourceNegNumSignSpec"))

	return err
}

// setLeadingNStrNegNumSignSpec - Deletes and resets the data
// value of the Leading Negative Number Sign contained in an
// instance of NumStrNegativeNumberSignSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		negNumSignSpec             *NumStrNegativeNumberSignSpec
//		   - A pointer to a NumStrNegativeNumberSignSpec instance.
//		     The Leading Negative Number Sign contained in this
//	      instance will be deleted and reset to the value
//	      specified by input parameter, ''.
//
//
//		leadingNegNumSign          []rune
//		   - An array of runes specifying the character or
//		     characters which will be copied to the Leading
//		     Negative Number Sign contained in input parameter,
//	      'negNumSignSpec'.
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
func (nStrNegNumSignSpecNanobot *numStrNegNumSignSpecNanobot) setLeadingNStrNegNumSignSpec(
	negNumSignSpec *NumStrNegativeNumberSignSpec,
	leadingNegNumSign []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNegNumSignSpecNanobot.lock == nil {
		nStrNegNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpecNanobot.lock.Lock()

	defer nStrNegNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNegNumSignSpecNanobot."+
			"setLeadingNStrNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSignSpec' is invalid!\n"+
			"'negNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSign) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSign' is invalid!\n"+
			"'leadingNegNumSign' is is empty and has a length of zero.\n",
			ePrefix.String())

		return err
	}

	new(numStrNegNumSignSpecAtom).emptyLeadingNegNumSign(
		negNumSignSpec)

	if len(leadingNegNumSign) == 0 {
		return err
	}

	err = negNumSignSpec.leadingNegNumSign.SetRuneArray(
		leadingNegNumSign,
		ePrefix.XCpy(
			"negNumSignSpec.leadingNegNumSign"+
				"<-leadingNegNumSign"))

	return err
}

// setTrailingNStrNegNumSignSpec - Deletes and resets the data
// value of the Trailing Negative Number Sign contained in an
// instance of NumStrNegativeNumberSignSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		negNumSignSpec             *NumStrNegativeNumberSignSpec
//		   - A pointer to a NumStrNegativeNumberSignSpec instance.
//		     The Trailing Negative Number Sign contained in this
//	      instance will be deleted and reset to the value
//	      specified by input parameter, ''.
//
//
//		trailingNegNumSign          []rune
//		   - An array of runes specifying the character or
//		     characters which will be copied to the Trailing
//		     Negative Number Sign contained in input parameter,
//	      'negNumSignSpec'.
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
func (nStrNegNumSignSpecNanobot *numStrNegNumSignSpecNanobot) setTrailingNStrNegNumSignSpec(
	negNumSignSpec *NumStrNegativeNumberSignSpec,
	trailingNegNumSign []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNegNumSignSpecNanobot.lock == nil {
		nStrNegNumSignSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNegNumSignSpecNanobot.lock.Lock()

	defer nStrNegNumSignSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNegNumSignSpecNanobot."+
			"setTrailingNStrNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSignSpec' is invalid!\n"+
			"'negNumSignSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSign) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSign' is invalid!\n"+
			"'trailingNegNumSign' is is empty and has a length of zero.\n",
			ePrefix.String())

		return err
	}

	new(numStrNegNumSignSpecAtom).emptyTrailingNegNumSign(
		negNumSignSpec)

	if len(trailingNegNumSign) == 0 {
		return err
	}

	err = negNumSignSpec.trailingNegNumSign.SetRuneArray(
		trailingNegNumSign,
		ePrefix.XCpy(
			"negNumSignSpec.trailingNegNumSign"+
				"<-trailingNegNumSign"))

	return err
}

// numStrNegNumSignSpecAtom - This type provides
// helper methods for NumStrNegativeNumberSignSpec
type numStrNegNumSignSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrNegativeNumberSignSpec and proceeds to reset the
// data values for all member variables to their initial or
// zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrNegNumSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	nStrNegNumSpec           *NumStrNegativeNumberSignSpec
//	   - A pointer to an instance of NumStrNegativeNumberSignSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNegNumSpecAtom *numStrNegNumSignSpecAtom) empty(
	nStrNegNumSpec *NumStrNegativeNumberSignSpec) {

	if nStrNegNumSpecAtom.lock == nil {
		nStrNegNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrNegNumSpecAtom.lock.Lock()

	defer nStrNegNumSpecAtom.lock.Unlock()

	if nStrNegNumSpec == nil {
		return
	}

	nStrNegNumSpec.leadingNegNumSign.Empty()

	nStrNegNumSpec.trailingNegNumSign.Empty()
}

// emptyLeadingNegNumSign - Receives a pointer to an instance
// of NumStrNegativeNumberSignSpec and proceeds to reset the
// member variable data for the leading negative number sign
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading negative
// number sign contained in input parameter 'nStrNegNumSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNegNumSpec           *NumStrNegativeNumberSignSpec
//	   - A pointer to an instance of NumStrNegativeNumberSignSpec.
//	     The Leading Negative Number Sign contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNegNumSpecAtom *numStrNegNumSignSpecAtom) emptyLeadingNegNumSign(
	nStrNegNumSpec *NumStrNegativeNumberSignSpec) {

	if nStrNegNumSpecAtom.lock == nil {
		nStrNegNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrNegNumSpecAtom.lock.Lock()

	defer nStrNegNumSpecAtom.lock.Unlock()

	if nStrNegNumSpec == nil {
		return
	}

	nStrNegNumSpec.leadingNegNumSign.Empty()

}

// emptyTrailingNegNumSign - Receives a pointer to an instance
// of NumStrNegativeNumberSignSpec and proceeds to reset the
// member variable data for the trailing negative number sign
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing negative
// number sign contained in input parameter 'nStrNegNumSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNegNumSpec           *NumStrNegativeNumberSignSpec
//	   - A pointer to an instance of NumStrNegativeNumberSignSpec.
//	     The Trailing Negative Number Sign contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNegNumSpecAtom *numStrNegNumSignSpecAtom) emptyTrailingNegNumSign(
	nStrNegNumSpec *NumStrNegativeNumberSignSpec) {

	if nStrNegNumSpecAtom.lock == nil {
		nStrNegNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrNegNumSpecAtom.lock.Lock()

	defer nStrNegNumSpecAtom.lock.Unlock()

	if nStrNegNumSpec == nil {
		return
	}

	nStrNegNumSpec.trailingNegNumSign.Empty()
}

// equal - Receives a pointer to two instances of
// NumStrNegativeNumberSignSpec and proceeds to compare their
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
//	nStrNegNumSpec1    *NumStrNegativeNumberSignSpec
//	   - An instance of NumStrNegativeNumberSignSpec.
//	     Internal member variables from 'nStrNegNumSpec1'
//	     will be compared to those of 'nStrNegNumSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrNegNumSpec2    *NumStrNegativeNumberSignSpec
//	   - An instance of NumStrNegativeNumberSignSpec.
//	     Internal member variables from 'nStrNegNumSpec2'
//	     will be compared to those of 'nStrNegNumSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrNegNumSpec1' and
//	     'nStrNegNumSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrNegNumSpecAtom *numStrNegNumSignSpecAtom) equal(
	nStrNegNumSpec1 *NumStrNegativeNumberSignSpec,
	nStrNegNumSpec2 *NumStrNegativeNumberSignSpec) bool {

	if nStrNegNumSpecAtom.lock == nil {
		nStrNegNumSpecAtom.lock = new(sync.Mutex)
	}

	nStrNegNumSpecAtom.lock.Lock()

	defer nStrNegNumSpecAtom.lock.Unlock()

	if nStrNegNumSpec1 == nil ||
		nStrNegNumSpec2 == nil {
		return false
	}

	if !nStrNegNumSpec1.leadingNegNumSign.Equal(
		&nStrNegNumSpec2.leadingNegNumSign) {

		return false
	}

	if !nStrNegNumSpec1.trailingNegNumSign.Equal(
		&nStrNegNumSpec2.trailingNegNumSign) {

		return false
	}

	return true
}
