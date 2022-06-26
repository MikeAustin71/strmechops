package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// CharSearchTargetInputParametersDto - Target Input Parameters are
// more easily understood in the context of text character search
// operations.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usually involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//    Target String        - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
// A comparison of text characters contained in the Target Search
// String and the Test String serves as the basis for determining
// a 'Match' condition or successful outcome from a text character
// search algorithm. The specific criterion for determining a
// 'Match' condition vary between the different Character Search
// Types.
//
// ----------------------------------------------------------------
//
// The Character Search Target Input Parameters Data Transfer
// Object type (CharSearchTargetInputParametersDto) is used to
// transmit Target String input parameters to methods performing
// search operations.
//
type CharSearchTargetInputParametersDto struct {
	TargetInputParametersName string
	// The Name, Label or descriptive Tag associated with this
	// instance of CharSearchTargetInputParametersDto. If empty,
	// this string will be defaulted to "TargetInputParameters"

	TargetString *RuneArrayDto
	// A pointer to the RuneArrayDto containing the Target
	// Search String text characters used in the search
	// algorithm. Target Characters are compared against
	// Test Characters to determine if a 'Match' condition
	// exists.

	TargetStringName string
	// The label or name of the 'TargetString' parameter.
	// Used in error and informational messages.

	TargetStringLength int
	// Actual number of text characters in the entire
	// Target Search String ('TargetString').

	TargetStringLengthName string
	// The label or name of the 'TargetStringLength' parameter.
	// Used in error and informational messages.

	TargetStringStartingSearchIndex int
	// The index in 'TargetString' at which the search
	// operation begins.

	TargetStringStartingSearchIndexName string
	// The label or name of the
	// TargetStringStartingSearchIndex parameter.
	// Used in error and informational messages.

	TargetStringSearchLength int
	// The actual number of characters within the Target
	// Search String that are included in the search
	// operation. This value may be less than the actual
	// length of the Target Search String. If this value
	// is set to -1, the search length will be configured
	// for to include the last index in 'TargetString'. In
	// other words the search will proceed to the end of
	// 'TargetString'.

	TargetStringSearchLengthName string
	// The label or name of the TargetStringSearchLength
	// parameter. Used in error and informational
	// messages.

	TargetStringAdjustedSearchLength int
	// The adjusted or corrected Target String Search
	// Length. This value is guaranteed to be equal to or
	// less than the actual Target String Length.

	TargetStringDescription1 string
	// First of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	TargetStringDescription2 string
	// Second of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	FoundFirstNumericDigitInNumStr bool
	// When set to 'true' this signals that the first
	// numeric digit has been identified in the text
	// characters specified by 'TargetString'

	TextCharSearchType CharacterSearchType
	// Optional. An enumeration value signaling the type
	// of text character search algorithm used to conduct
	// this search operation. When set to a valid value,
	// this specification will override the search
	// specification contained in the Test Input
	// Parameters Data Transfer Object.
	//
	// Valid CharSearch Type values are listed as follows:
	//  CharSearchType.None() - Invalid Value
	//  CharSearchType.LinearTargetStartingIndex() - Default
	//  CharSearchType.SingleTargetChar()
	//  CharSearchType.LinearEndOfString()

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// CharSearchTargetInputParametersDto ('sourceTargetInputParms') to
// the data fields of the current
// CharSearchTargetInputParametersDto instance
// ('searchTargetInputParmsDto').
//
// IMPORTANT
// All the data fields in current
// CharSearchTargetInputParametersDto instance
// ('searchTargetInputParmsDto') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sourceTargetInputParms     *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. This method will NOT
//       change the values of internal member variables contained
//       in this instance.
//
//       All data values in this CharSearchTargetInputParametersDto
//       instance will be copied to current
//       CharSearchTargetInputParametersDto instance
//       ('searchTargetInputParmsDto').
//
//       If parameter 'sourceTargetInputParms' is determined to be
//       invalid, an error will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) CopyIn(
	sourceTargetInputParms *CharSearchTargetInputParametersDto,
	errorPrefix interface{}) (
	err error) {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	targetInputParmsNanobot := charSearchTargetInputParametersDtoNanobot{}

	err =
		targetInputParmsNanobot.copyIn(
			searchTargetInputParmsDto,
			sourceTargetInputParms,
			ePrefix.XCpy(
				"searchTargetInputParmsDto<-"+
					"sourceTargetInputParms"))

	return err
}

// CopyOut - Returns a deep copy of the current
// CharSearchTargetInputParametersDto instance.
//
// If the current CharSearchTargetInputParametersDto instance
// contains invalid member variables, this method will return an
// error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyTargetInputParms   CharSearchTargetInputParametersDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchTargetInputParametersDto instance.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyTargetInputParms CharSearchTargetInputParametersDto,
	err error) {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyTargetInputParms, err
	}

	deepCopyTargetInputParms,
		err = charSearchTargetInputParametersDtoNanobot{}.ptr().
		copyOut(
			searchTargetInputParmsDto,
			ePrefix.XCpy(
				"deepCopyTargetInputParms<-"+
					"searchTargetInputParmsDto"))

	return deepCopyTargetInputParms, err
}

// Empty - Resets all internal member variables for the current
// instance of CharSearchTargetInputParametersDto to their initial
// or zero states.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// CharSearchTargetInputParametersDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) Empty() {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	charSearchTargetInputParametersDtoAtom{}.ptr().
		empty(searchTargetInputParmsDto)

	searchTargetInputParmsDto.lock.Unlock()

	searchTargetInputParmsDto.lock = nil
}

// Equal - Receives a pointer to another instance of
// CharSearchTargetInputParametersDto and proceeds to compare its
// internal member variables to those of the current
// CharSearchTargetInputParametersDto instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTargetInputParms   *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. The internal member
//       variable data values in this instance will be compared to
//       those in the current instance of
//       CharSearchTargetInputParametersDto. The results of this
//       comparison will be returned to the calling functions as a
//       boolean value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingTargetInputParms' are equivalent
//       in all respects to those contained in the current instance
//       of CharSearchTargetInputParametersDto, this return value
//       will be set to 'true'.
//
//       Otherwise, this method will return 'false'.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) Equal(
	incomingTargetInputParms *CharSearchTargetInputParametersDto) bool {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	return charSearchTargetInputParametersDtoAtom{}.ptr().
		equal(searchTargetInputParmsDto,
			incomingTargetInputParms)
}

// EqualTargetStrings - Receives a pointer to another instance of
// CharSearchTargetInputParametersDto and proceeds to compare the
// internal member variable 'TargetString' to the same internal
// member variable 'TargetString' contained in the current
// CharSearchTargetInputParametersDto instance in order to
// determine if the two 'TargetString' member variables are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the 'TargetString' member variables for both
// instances are equal in all respects, this flag is set to 'true'.
// Otherwise, this method returns a boolean value of 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTargetInputParms   *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. The internal member
//       variable 'TargetString' contained in this instance will be
//       compared to the same member variable 'TargetString'
//       contained in the current instance of
//       CharSearchTargetInputParametersDto. The results of this
//       comparison will be returned to the calling functions as a
//       boolean value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable 'TargetString' contained
//       in input parameter 'incomingTargetInputParms' is
//       equivalent in all respects to the same member variable
//       'TargetString' contained in the current instance of
//       CharSearchTargetInputParametersDto, this return value
//       will be set to 'true'.
//
//       Otherwise, this method will return a boolean value of
//       'false'.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) EqualTargetStrings(
	incomingTargetInputParms *CharSearchTargetInputParametersDto) bool {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	return charSearchTargetInputParametersDtoElectron{}.ptr().
		equalTargetStrings(searchTargetInputParmsDto,
			incomingTargetInputParms)
}

// GetFormattedText - Returns a formatted text string detailing all
// member variables and their values for the current instance of
// CharSearchTargetInputParametersDto.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"GetFormattedText()",
		"")

	if err != nil {

		return "", err

	}

	var strBuilder strings.Builder

	strBuilder,
		err = charSearchTargetInputParametersDtoNanobot{}.ptr().
		getFormattedText(
			searchTargetInputParmsDto,
			ePrefix.XCpy(
				"strBuilder<-Formatted Text"))

	if err != nil {
		return "", err
	}

	return strBuilder.String(), err
}

// IsValidInstance - Performs a diagnostic review of the member
// variable data values encapsulated in the current instance of
// CharSearchTargetInputParametersDto to determine if they are
// valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// In addition to performing validation diagnostics on the current
// instance of CharSearchTargetInputParametersDto, this method will
// proceed to set all empty member variable labels or name strings
// to their default values.
//
// Type CharSearchTargetInputParametersDto contains a number of
// string variables which are used to label, name or otherwise
// describe other operational member variables. These labels are
// used for error or informational messages. If any of these label
// strings are empty when this method is called, those empty label
// strings will be set to their default values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If all data elements encapsulated by the current instance
//       of CharSearchTargetInputParametersDto are valid, this
//       returned boolean value is set to 'true'. If any data
//       values are invalid, this return parameter is set to
//       'false'.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) IsValidInstance() (
	isValid bool) {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	isValid,
		_ = charSearchTargetInputParametersDtoAtom{}.ptr().
		testValidityOfTargetInputParms(
			searchTargetInputParmsDto,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current instance of
// CharSearchTargetInputParametersDto to determine if they are all
// valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// This method is functionally equivalent to method:
//  CharSearchTargetInputParametersDto.ValidateTargetParameters()
//
// This method, IsValidInstanceError(), is included for convenience
// and continuity as many other types in this package also use a
// method of this name when checking the validity of internal member
// variables.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix         interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of CharSearchTargetInputParametersDto
//       are found to be invalid, this method will return an error.
//       If the member data variables are determined to be valid,
//       this error return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) IsValidInstanceError(
	errorPrefix interface{}) error {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}
	_,
		err = charSearchTargetInputParametersDtoAtom{}.ptr().
		testValidityOfTargetInputParms(
			searchTargetInputParmsDto,
			ePrefix.XCpy(
				"searchTargetInputParmsDto"))

	return err
}

// New - Returns a new, uninitialized instance of
// CharSearchTargetInputParametersDto. All member variables are
// guaranteed to be set to their zero or uninitialized states.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//
//  CharSearchTargetInputParametersDto
//     - This method will return an empty or uninitialized instance
//       of type CharSearchTargetInputParametersDto. All member
//       variables are guaranteed to be set to their or
//       uninitialized states.
//
func (searchTargetInputParmsDto CharSearchTargetInputParametersDto) New() CharSearchTargetInputParametersDto {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	newEmptyTargetInputParms := CharSearchTargetInputParametersDto{}

	newEmptyTargetInputParms.Empty()

	return newEmptyTargetInputParms
}

// NewTargetString - Returns a new, instance of
// CharSearchTargetInputParametersDto populated from input
// parameters for Target String, Target Input Parameters Name,
// Target String Starting Search Index and Target String Search
// Length.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetString                    *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. 'targetString'
//       contains a rune array internal member variable which
//       specifies the target string to be used a text character
//       search operation.
//
//
//  targetInputParametersName       string
//     - Optional label, name or descriptive text associated
//       with the newly created instance of
//       CharSearchTargetInputParametersDto returned by this
//       method.
//
//
//  targetStringStartingSearchIndex int
//     - An integer containing the index number within
//       'targetString' where the text character search operation
//       will begin.
//
//
//  targetStringSearchLength        int
//     - The number of characters which will be searched in
//       'targetString'.
//
//       Set this parameter to a value of minus one (-1) to ensure
//       that the search operation proceeds from the Starting
//       Search Index to the end of the string in 'targetString'.
//
//
//  errorPrefix                     interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  newTargetInputParms        CharSearchTargetInputParametersDto
//  If this method completes successfully, it will return a new
//  instance of CharSearchTargetInputParametersDto populated with
//  new values for Target String, Target String Starting Search
//  Index and Target String Search Length.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (searchTargetInputParmsDto CharSearchTargetInputParametersDto) NewTargetString(
	targetString *RuneArrayDto,
	targetInputParametersName string,
	targetStringStartingSearchIndex int,
	targetStringSearchLength int,
	errorPrefix interface{}) (
	newTargetInputParms CharSearchTargetInputParametersDto,
	err error) {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"NewTargetString()",
		"")

	if err != nil {
		return newTargetInputParms, err
	}

	if targetString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetString' is a nil pointer!\n",
			ePrefix.String())

		return newTargetInputParms, err
	}

	if targetString.GetRuneArrayLength() == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetString' is empty\n"+
			"and a length of zero!\n",
			ePrefix.String())

		return newTargetInputParms, err
	}

	if targetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetStringStartingSearchIndex' is invalid!\n"+
			"'targetStringStartingSearchIndex' has a value less than zero (0).\n"+
			"targetStringStartingSearchIndex = '%v'\n",
			ePrefix.String(),
			targetStringStartingSearchIndex)

		return newTargetInputParms, err
	}

	if targetStringSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetStringSearchLength' is invalid!\n"+
			"'targetStringSearchLength' has a value less than minus one (-1).\n"+
			"targetStringSearchLength = '%v'\n",
			ePrefix.String(),
			targetStringSearchLength)

		return newTargetInputParms, err
	}

	newEmptyTargetInputParms := CharSearchTargetInputParametersDto{}

	newEmptyTargetInputParms.Empty()

	newEmptyTargetInputParms.TargetString = targetString

	newEmptyTargetInputParms.TargetInputParametersName =
		targetInputParametersName

	newEmptyTargetInputParms.TargetStringStartingSearchIndex =
		targetStringStartingSearchIndex

	newEmptyTargetInputParms.TargetStringSearchLength =
		targetStringSearchLength

	_,
		err = charSearchTargetInputParametersDtoAtom{}.ptr().
		testValidityOfTargetInputParms(
			&newEmptyTargetInputParms,
			ePrefix.XCpy(
				"newEmptyTargetInputParms"))

	return newEmptyTargetInputParms, err
}

// String - Returns a formatted text string detailing all the
// internal member variables in the current instance of
// CharSearchTargetInputParametersDto and their values.
//
// If an error is encountered, the error message is included in the
// string returned by this method.
//
// This method implements the Stringer Interface.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) String() string {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"CharSearchTargetInputParametersDto."+
			"String()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"CharSearchTargetInputParametersDto.String()",
			err.Error())

		return errOut
	}

	var strBuilder strings.Builder

	strBuilder,
		err = charSearchTargetInputParametersDtoNanobot{}.ptr().
		getFormattedText(
			searchTargetInputParmsDto,
			ePrefix.XCpy(
				"strBuilder"))

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

// ValidateTargetParameters - Validates internal member variables
// contained in the current instance of
// CharSearchTargetInputParametersDto.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// In addition to performing validation diagnostics on the current
// instance of CharSearchTargetInputParametersDto, this method will
// proceed to set all empty member variable labels or name strings
// to their default values.
//
// Type CharSearchTargetInputParametersDto contains a number of
// string variables which are used to label, name or otherwise
// describe other operational member variables. These labels are
// used for error or informational messages. If any of these label
// strings are empty when this method is called, those empty label
// strings will be set to their default values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of CharSearchTargetInputParametersDto
//       are found to be invalid, this method will return an error.
//       If the member data variables are determined to be valid,
//       this error return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) ValidateTargetParameters(
	errorPrefix interface{}) error {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"ValidateTargetParameters()",
		"")

	if err != nil {

		return err

	}

	targetInputParmsAtom :=
		charSearchTargetInputParametersDtoAtom{}

	_,
		err = targetInputParmsAtom.
		testValidityOfTargetInputParms(
			searchTargetInputParmsDto,
			ePrefix.XCpy(
				"searchTargetInputParmsDto"))

	return err
}

func (searchTargetInputParmsDto *CharSearchTargetInputParametersDto) ValidateCharSearchType(
	errorPrefix interface{}) error {

	if searchTargetInputParmsDto.lock == nil {
		searchTargetInputParmsDto.lock = new(sync.Mutex)
	}

	searchTargetInputParmsDto.lock.Lock()

	defer searchTargetInputParmsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"ValidateCharSearchType()",
		"")

	if err != nil {

		return err

	}

	if !searchTargetInputParmsDto.TextCharSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The Character Search Type is invalid!\n"+
			"Character Search Type must be set to one of these\n"+
			"enumeration values:\n"+
			"  CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"  CharacterSearchType(0).SingleTargetChar()\n"+
			"  CharacterSearchType(0).LinearEndOfString()\n"+
			"The invalid Input Character Search Type is currently\n"+
			"configured as:\n"+
			" Character Search Type   String Name: %v\n"+
			" Character Search Type Integer Value: %v\n",
			ePrefix.String(),
			searchTargetInputParmsDto.TextCharSearchType.String(),
			searchTargetInputParmsDto.TextCharSearchType.XValueInt())

	}

	return err
}
