package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrIntSeparator is designed to be used as an array
// element in an array which will be used to insert integer
// separators, primarily thousands separators, into number
// strings. Some countries/cultures do not use thousands
// separation and instead rely on multiple integer separation
// characters and grouping sequences for a single number
// string. Notable examples of this found in the 'Indian
// Number System' and 'Chinese Numerals'.
//  Reference:
//  https://en.wikipedia.org/wiki/Indian_numbering_system
//  https://en.wikipedia.org/wiki/Chinese_numerals
//  https://en.wikipedia.org/wiki/Decimal_separator
//
// An array of NumStrIntSeparator elements provides the flexibility
// necessary to process these complex number separation formats.
//
type NumStrIntSeparator struct {
	intSeparatorChars       []rune // A series of runes used to separate integer digits.
	intSeparatorGrouping    uint   // Number of integer digits in a group
	intSeparatorRepetitions uint   // Number of times this character/group sequence is repeated
	//                              // A zero value signals unlimited repetitions.
	restartIntGroupingSequence bool // If true, the entire grouping sequence is repeated
	//                              //  beginning at array index zero.
	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming
// NumStrIntSeparator instance to the data fields of the current
// NumStrIntSeparator instance.
//
// If input parameter 'incomingNStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
// Be advised, all of the data fields in the current
// NumStrIntSeparator instance will be overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNStrIntSeparator     *NumStrIntSeparator
//     - A pointer to an instance of NumStrIntSeparator.
//       The data values in this object will be copied to the
//       current NumStrIntSeparator instance.
//
//       If input parameter 'incomingNStrIntSeparator' is judged
//       to be invalid, this method will return an error.
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
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) CopyIn(
	incomingNStrIntSeparator *NumStrIntSeparator,
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
		"NumStrIntSeparator.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorMolecule{}.ptr().
		copyIn(
			nStrIntSep,
			incomingNStrIntSeparator,
			ePrefix)
}

// CopyOut - Creates and returns a deep copy of the current
// NumStrIntSeparator instance.
//
// If the current NumStrIntSeparator instance is judged to be
// invalid, this method will return an error.
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
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned containing
//       all of the data values copied from the current instance of
//       NumStrIntSeparator.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) CopyOut(
	errorPrefix interface{}) (
	NumStrIntSeparator,
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
		"NumStrIntSeparator.CopyOut()",
		"")

	if err != nil {
		return NumStrIntSeparator{}, err
	}

	return numStrIntSeparatorMolecule{}.ptr().
		copyOut(
			nStrIntSep,
			ePrefix)
}

// Empty - Deletes and resets the data values for all member
// variables within the current NumStrIntSeparator instance to
// their 'zero' values.
//
func (nStrIntSep *NumStrIntSeparator) Empty() {
	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	_ = numStrIntSeparatorQuark{}.ptr().empty(
		nStrIntSep,
		nil)

	nStrIntSep.lock.Unlock()

	nStrIntSep.lock = nil

	return
}

// Equal - Receives an NumStrIntSeparator object and proceeds to
// determine whether all data elements in this object are equal to
// all corresponding data elements in the current instance of
// NumStrIntSeparator.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSepTwo       NumStrIntSeparator
//     - This method will compare all data elements in the current
//       NumStrIntSeparator object to corresponding data elements
//       in this second NumStrIntSeparator object in order
//       determine equivalency.
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
//  bool
//     - If all the data elements in the current NumStrIntSeparator
//       instance are equal to all the corresponding data elements
//       in 'nStrIntSepTwo', this return parameter will be set to
//       'true'. If all the data elements are NOT equal, this
//       return parameter will be set to 'false'.
//
//
//  error
//     - If all the data elements in the current NumStrIntSeparator
//       are equal to all the corresponding data elements in
//       'nStrIntSepTwo', this return parameter will be set to
//       'nil'.
//
//       If the corresponding data elements are not equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
//       If errors are encountered during processing, the returned
//       error message will incorporate the method chain  and text
//       passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) Equal(
	nStrIntSepTwo NumStrIntSeparator,
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
		"NumStrIntSeparator.Equal())",
		"")

	if err != nil {
		return false, err
	}

	return numStrIntSeparatorMolecule{}.ptr().
		equal(
			nStrIntSep,
			&nStrIntSepTwo,
			ePrefix.XCtx(
				"nStrIntSep vs nStrIntSepTwo"))
}

// GetIntSeparatorChars - Returns the integer digit separator
// characters for the current NumStrIntSeparator instance as an
// array of runes.
//
// If the rune array is zero length, an error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  [] rune
//     - If this method completes successfully, an array of runes
//       will be returned representing the integer digit separator
//       characters.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) GetIntSeparatorChars(
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
		"NumStrIntSeparator.GetIntSeparatorChars()",
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

// GetIntSeparatorGrouping - Returns an unsigned integer
// representing the number of integer digits included in an integer
// group for the current NumStrIntSeparator instance.
//
// This unsigned integer value is used to group integers within a
// number string.
//
// In most western countries integer digits to the left of
// the decimal separator (a.k.a. decimal point) are separated
// into groups of three digits representing a grouping of
// 'thousands' like this: '1,000,000,000'. In this case the
// intSeparatorGrouping value would be set to three ('3').
//
// In some countries and cultures other integer groupings are
// used. In India, for example, a number might be formatted
// like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
// would be formatted like this: 12,3456,7890,2345
//
func (nStrIntSep *NumStrIntSeparator) GetIntSeparatorGrouping() uint {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return nStrIntSep.intSeparatorGrouping
}

// GetIntSeparatorRepetitions - Returns the Integer Separator
// repetitions specification for the current NumStrIntSeparator
// instance.
//
// The returned unsigned integer value specifies the number of
// times this integer grouping is repeated. A value of zero signals
// that this integer grouping will be repeated indefinitely.
//
func (nStrIntSep *NumStrIntSeparator) GetIntSeparatorRepetitions() uint {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return nStrIntSep.intSeparatorRepetitions
}

// GetRestartIntGroupingSequence - Returns the internal member
// variable 'restartIntGroupingSequence' for the current
// NumStrIntSeparator instance.
//
// The NumStrIntSeparator type is intended to be configured in an
// array of NumStrIntSeparator objects which, taken as a whole,
// provides formatting specifications for complex integer group
// separation operations.
//
// If the current NumStrIntSeparator is the last element in an
// array of NumStrIntSeparator objects, the 'Restart Integer
// Grouping Sequence' flag signals whether the integer separation
// operation will be restarted from the first NumStrIntSeparator
// object in the array.
//
// Again, the NumStrIntSeparator.restartIntGroupingSequence boolean
// flag only has meaning if the current NumStrIntSeparator object
// is last element in an array of NumStrIntSeparator objects.
//
func (nStrIntSep *NumStrIntSeparator) GetRestartIntGroupingSequence() bool {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	return nStrIntSep.restartIntGroupingSequence
}

// IsValidInstance - Performs a diagnostic review of the current
// NumStrIntSeparator instance to determine whether the current
// instance is valid in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - This returned boolean value will signal whether the
//       current NumStrIntSeparator is valid, or not. If the
//       current NumStrIntSeparator contains valid data, this
//       method returns 'true'. If the data is invalid, this method
//       will return 'false'.
//
func (nStrIntSep *NumStrIntSeparator) IsValidInstance() (
	isValid bool) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	isValid,
		_ =
		numStrIntSeparatorQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumStrIntSeparator instance to determine whether the
// current instance is valid in all respects.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the current instance of NumStrIntSeparator contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nStrIntSep *NumStrIntSeparator) IsValidInstanceError(
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
		"NumStrIntSeparator.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err =
		numStrIntSeparatorQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				nStrIntSep,
				ePrefix)

	return err
}

// NewBasic - Creates and returns a new instance of
// NumStrIntSeparator. The returned NumStrIntSeparator instance
// represents a basic or simple integer separator object using
// default values and a minimum number of input parameters.
//
// The input parameter 'integerDigitsSeparators' is string
// containing the integer separator characters. The integer digit
// grouping is defaulted to a value of three (3). The 'separator
// repetitions' value is defaulted to zero (0) signaling unlimited
// repetitions.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators    string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewBasic(
	integerDigitsSeparators string,
	errorPrefix interface{}) (
	NumStrIntSeparator,
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
		"NumStrIntSeparator.NewBasic()",
		"")

	if err != nil {
		return NumStrIntSeparator{}, err
	}

	ePrefix.SetEPref(
		"NumStrIntSeparator.NewBasic()")

	newIntSep := NumStrIntSeparator{}

	err =
		numStrIntSeparatorUtility{}.ptr().
			setBasic(
				&newIntSep,
				integerDigitsSeparators,
				ePrefix.XCtx(
					"newIntSep"))

	return newIntSep, err
}

// NewBasicRunes - Creates and returns a new instance of
// NumStrIntSeparator. The returned NumStrIntSeparator instance
// represents a basic or simple integer separator object using
// default values and a minimum number of input parameters.
//
// This method is an alternative to method
// NumStrIntSeparator.NewBasic() in that this method accepts
// integer separator characters as an array of runes instead
// of a string.
//
// The input parameter 'integerDigitsSeparators' is an array of
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
// 'integerDigitsSeparators'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator (a.k.a. decimal
//       point). In the United States, the standard integer digits
//       separator is the comma (',').
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewBasicRunes(
	integerDigitsSeparators []rune,
	errorPrefix interface{}) (
	NumStrIntSeparator,
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
		"NumStrIntSeparator.NewBasicRunes()",
		"")

	if err != nil {
		return NumStrIntSeparator{}, err
	}

	newIntSep := NumStrIntSeparator{}

	err =
		numStrIntSeparatorUtility{}.ptr().
			setBasicRunes(
				&newIntSep,
				integerDigitsSeparators,
				ePrefix.XCtx(
					"newIntSep"))

	return newIntSep, err
}

// NewDetail - Creates and returns a new instance of
// NumStrIntSeparator. The new instance is generated based on
// component elements passed as input parameters.
//
// This method differs from NumStrIntSeparator.NewDetailRunes() in
// that this method accepts a string for input parameter
// 'intSeparatorChars'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          string
//     - A character, or series of characters, used to separate
//       integer digits in a number string. These characters are
//       commonly known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer value specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       have an integer grouping value of four ('4').
//         Chinese Numerals Example: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - The NumStrIntSeparator type is intended to be configured
//       in an array of NumStrIntSeparator objects which, taken as
//       a whole, provides formatting specifications for complex
//       integer group separation operations.
//
//       If the current NumStrIntSeparator is the last element in
//       an array of NumStrIntSeparator objects, the 'Restart
//       Integer Grouping Sequence' flag signals whether the
//       integer separation operation will be restarted from the
//       first NumStrIntSeparator object in the array.
//
//       In summary, if the NumStrIntSeparator is the last element
//       in an array of NumStrIntSeparator objects, this boolean
//       flag signals whether the entire integer grouping sequence
//       will be restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  newIntSep                  NumStrIntSeparator
//     - If this method completes successfully, the parameter will
//       return a new and fully populated instance of
//       NumStrIntSeparator.
//
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewDetail(
	intSeparatorChars string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep NumStrIntSeparator,
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
		"NumStrIntSeparator.NewDetail()",
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
		numStrIntSeparatorMechanics{}.ptr().
			setWithComponents(
				&newIntSep,
				[]rune(intSeparatorChars),
				intSeparatorGrouping,
				intSeparatorRepetitions,
				restartIntGroupingSequence,
				ePrefix.XCtx("newIntSep"))

	return newIntSep, err
}

// NewDetailRunes - Creates and returns a new instance of
// NumStrIntSeparator. The new instance is generated based on
// component elements passed as input parameters.
//
// This method differs from NumStrIntSeparator.NewDetail() in that
// this method accepts an array of runes for 'intSeparatorChars'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - A character, or series of characters, used to separate
//       integer digits in a number string. These characters are
//       commonly known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer value specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       have an integer grouping value of four ('4').
//         Chinese Numerals Example: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - The NumStrIntSeparator type is intended to be configured
//       in an array of NumStrIntSeparator objects which, taken as
//       a whole, provides formatting specifications for complex
//       integer group separation operations.
//
//       If the current NumStrIntSeparator is the last element in
//       an array of NumStrIntSeparator objects, the 'Restart
//       Integer Grouping Sequence' flag signals whether the
//       integer separation operation will be restarted from the
//       first NumStrIntSeparator object in the array.
//
//       In summary, if the NumStrIntSeparator is the last element
//       in an array of NumStrIntSeparator objects, this boolean
//       flag signals whether the entire integer grouping sequence
//       will be restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  newIntSep                  NumStrIntSeparator
//     - If this method completes successfully, the parameter will
//       return a new and fully populated instance of
//       NumStrIntSeparator.
//
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewDetailRunes(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep NumStrIntSeparator,
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
		"NumStrIntSeparator.NewDetail()",
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
		numStrIntSeparatorMechanics{}.ptr().
			setWithComponents(
				&newIntSep,
				intSeparatorChars,
				intSeparatorGrouping,
				intSeparatorRepetitions,
				restartIntGroupingSequence,
				ePrefix.XCtx("newIntSep"))

	return newIntSep, err
}

// NewUnitedStatesDefaults - Creates and returns a new instance of
// NumStrIntSeparator set to United States default integer
// separators. Integer separator values used in the United States
// consist of the comma character (','), an integer grouping of
// three ('3') and unlimited repetitions of this sequence.
//
//   United States Integer Separation Example:
//         '1,000,000,000,000'
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
// -----------------------------------------------------------------
//
//  NumStrIntSeparator
//     - This will return a new and fully populated instance of
//       NumStrIntSeparator configured with United States default
//       integer separator values.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewUnitedStatesDefaults(
	errorPrefix interface{}) (
	NumStrIntSeparator,
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
		"NumStrIntSeparator."+
			"NewUnitedStatesDefaults()",
		"")

	if err != nil {
		return NumStrIntSeparator{}, err
	}

	newIntSep := NumStrIntSeparator{}

	err = numStrIntSeparatorMechanics{}.ptr().
		setToUSADefaults(
			&newIntSep,
			ePrefix.XCtx(
				"newIntSep"))

	return newIntSep, err
}

// NewWithComponents - Creates and returns a new instance of
// NumStrIntSeparator. The new instance is generated based on
// component elements passed as input parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - A series of runes or characters used to separate integer
//       digits in a number string. These characters are commonly
//       known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer value specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       have an integer grouping value of four ('4').
//         Chinese Numerals Example: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - The NumStrIntSeparator type is intended to be configured
//       in an array of NumStrIntSeparator objects which, taken as
//       a whole, provides formatting specifications for complex
//       integer group separation operations.
//
//       If the current NumStrIntSeparator is the last element in
//       an array of NumStrIntSeparator objects, the 'Restart
//       Integer Grouping Sequence' flag signals whether the
//       integer separation operation will be restarted from the
//       first NumStrIntSeparator object in the array.
//
//       In summary, if the NumStrIntSeparator is the last element
//       in an array of NumStrIntSeparator objects, this boolean
//       flag signals whether the entire integer grouping sequence
//       will be restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  newIntSep                  NumStrIntSeparator
//     - If this method completes successfully, the parameter will
//       return a new and fully populated instance of
//       NumStrIntSeparator.
//
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep NumStrIntSeparator) NewWithComponents(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	errorPrefix interface{}) (
	newIntSep NumStrIntSeparator,
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
		"NumStrIntSeparator.NewWithComponents()",
		"")

	if err != nil {
		return newIntSep, err
	}

	err =
		numStrIntSeparatorMechanics{}.ptr().
			setWithComponents(
				&newIntSep,
				intSeparatorChars,
				intSeparatorGrouping,
				intSeparatorRepetitions,
				restartIntGroupingSequence,
				ePrefix.XCtx("newIntSep"))

	return newIntSep, err
}

// SetBasic - Overwrites all the member variable data values for
// the current NumStrIntSeparator. This method is intended to
// configure a basic or simple integer separator object using
// default values and a minimum number of input parameters
//
// Input parameter 'integerDigitsSeparators' specifies  the integer
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
// 'integerDigitsSeparators'.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators    string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetBasic(
	integerDigitsSeparators string,
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
		"NumStrIntSeparator."+
			"SetBasicRunes()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorUtility{}.ptr().
		setBasic(
			nStrIntSep,
			integerDigitsSeparators,
			ePrefix.XCtx(
				"nStrIntSep"))
}

// SetBasicRunes - Overwrites all the member variable data values for
// the current NumStrIntSeparator. This method is intended to
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
// NumStrIntSeparator.SetBasic() in that this method accepts
// integer separator characters as an array of runes instead
// of a string.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (',').
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NumStrIntSeparator
//     - If this method completes successfully, a new instance of
//       NumStrIntSeparator will be created and returned. The
//       'integer digits grouping sequence' will be automatically
//       set to a default value of 3-digits.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetBasicRunes(
	integerDigitsSeparators []rune,
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
		"NumStrIntSeparator."+
			"SetBasicRunes()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorUtility{}.ptr().
		setBasicRunes(
			nStrIntSep,
			integerDigitsSeparators,
			ePrefix.XCtx(
				"nStrIntSep"))
}

// SetDetail - This method will overwrite and reset the
// internal member variable data values for the current
// NumStrIntSeparator instance based on the component elements
// passed as input parameters.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          string
//     - A character, or series of characters, used to separate
//       integer digits in a number string. These characters are
//       commonly known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer values specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       would be formatted like this: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - If the NumStrIntSeparator is the last element in an array
//       of NumStrIntSeparator objects, this boolean flag signals
//       whether the entire integer grouping sequence will be
//       restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetDetail(
	intSeparatorChars string,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
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
		"NumStrIntSeparator."+
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

	return numStrIntSeparatorMechanics{}.ptr().
		setWithComponents(
			nStrIntSep,
			[]rune(intSeparatorChars),
			intSeparatorGrouping,
			intSeparatorRepetitions,
			restartIntGroupingSequence,
			ePrefix.XCtx("nStrIntSep"))
}

// SetDetailRunes - This method will overwrite and reset the
// internal member variable data values for the current
// NumStrIntSeparator instance based on the component elements
// passed as input parameters.
//
// This method differs from NumStrIntSeparator.SetDetail() in that
// this method accepts an array of runes for 'intSeparatorChars'.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - A character, or series of characters, used to separate
//       integer digits in a number string. These characters are
//       commonly known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer values specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       would be formatted like this: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - If the NumStrIntSeparator is the last element in an array
//       of NumStrIntSeparator objects, this boolean flag signals
//       whether the entire integer grouping sequence will be
//       restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetDetailRunes(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
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
		"NumStrIntSeparator.SetDetailRunes()",
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

	return numStrIntSeparatorMechanics{}.ptr().
		setWithComponents(
			nStrIntSep,
			intSeparatorChars,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			restartIntGroupingSequence,
			ePrefix.XCtx("nStrIntSep"))
}

// SetIntSeparatorChars - Sets the 'intSeparatorChars' member
// variable for the current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - A series of runes or characters used to separate integer
//       digits in a number string. These characters are commonly
//       known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//       If this parameter resolves as a zero length array, an
//       error will be returned.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetIntSeparatorChars(
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
		"NumStrIntSeparator.SetIntSeparatorChars()",
		"")

	if err != nil {
		return err
	}

	lenIntSepChars := len(intSeparatorChars)

	if lenIntSepChars == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a zero length array.\n",
			ePrefix.String())
	}

	nStrIntSep.intSeparatorChars =
		make([]rune, lenIntSepChars, lenIntSepChars+5)

	for i := 0; i < lenIntSepChars; i++ {
		nStrIntSep.intSeparatorChars[i] =
			intSeparatorChars[i]
	}

	return nil
}

// SetIntSeparatorGrouping - Sets the 'Integer Separator Grouping'
// specification for the current NumStrIntSeparator instance.
//
// This unsigned integer values specifies the number of integer
// digits within a group. This value is used to group integers
// within a number string.
//
// In most western countries integer digits to the left of the
// decimal separator (a.k.a. decimal point) are separated into
// groups of three digits representing a grouping of 'thousands'
// like this: '1,000,000,000'. In this case the
// 'intSeparatorGrouping' value would be set to three ('3').
//
// In some countries and cultures other integer groupings are used.
// In India, for example, a number might be formatted like this:
// '6,78,90,00,00,00,00,000'. Chinese Numerals have an integer
// grouping value of four ('4').
//   Chinese Numerals Example: '12,3456,7890,2345'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorGrouping       uint
//     - The 'Integer Separator Grouping' value used to set
//       the integer grouping specification for the current
//       NumStrIntSeparator instance.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  -- NONE --
//
func (nStrIntSep *NumStrIntSeparator) SetIntSeparatorGrouping(
	intSeparatorGrouping uint) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	nStrIntSep.intSeparatorGrouping = intSeparatorGrouping

	return
}

// SetIntSeparatorRepetitions - Sets the 'Integer Separator
// Repetitions' specification for the current NumStrIntSeparator
// instance.
//
// The 'Integer Separator Repetitions' value is an unsigned integer
// which specifies the number of cycles for which the integer
// separation operation defined by separator characters and integer
// digit grouping value in this NumStrIntSeparator instance will be
// repeated.
//
// A value of zero for 'Integer Separator Repetitions' specifies
// that the integer separation operation configured by the current
// NumStrIntSeparator instance will be repeated indefinitely for
// all integer numeric digits within the number string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorRepetitions    uint
//     - The 'Integer Separator Repetitions' value used to
//       determine how many cycles the current integer digit
//       separation operation will be repeated.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  -- NONE --
//
func (nStrIntSep *NumStrIntSeparator) SetIntSeparatorRepetitions(
	intSeparatorRepetitions uint) {

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.lock.Lock()

	defer nStrIntSep.lock.Unlock()

	nStrIntSep.intSeparatorRepetitions =
		intSeparatorRepetitions

	return
}

// SetRestartIntGroupingSequence - Sets the internal member
// variable 'restartIntGroupingSequence' for the current
// NumStrIntSeparator instance.
//
// The NumStrIntSeparator type is intended to be configured in an
// array of NumStrIntSeparator objects which, taken as a whole,
// provides formatting specifications for complex integer group
// separation operations.
//
// If the current NumStrIntSeparator is the last element in an
// array of NumStrIntSeparator objects, the 'Restart Integer
// Grouping Sequence' flag signals whether the integer separation
// operation will be restarted from the first NumStrIntSeparator
// object in the array.
//
// Again, the NumStrIntSeparator.restartIntGroupingSequence boolean
// flag only has meaning if the current NumStrIntSeparator object
// is last element in an array of NumStrIntSeparator objects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  restartIntGroupingSequence bool
//     - The NumStrIntSeparator type is intended to be configured
//       in an array of NumStrIntSeparator objects which, taken as
//       a whole, provides formatting specifications for complex
//       integer group separation operations.
//
//       If the current NumStrIntSeparator is the last element in
//       an array of NumStrIntSeparator objects, the 'Restart
//       Integer Grouping Sequence' flag signals whether the
//       integer separation operation will be restarted from the
//       first NumStrIntSeparator object in the array.
//
//       In summary, if the NumStrIntSeparator is the last element
//       in an array of NumStrIntSeparator objects, this boolean
//       flag signals whether the entire integer grouping sequence
//       will be restarted from array element zero.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  -- NONE --
//
func (nStrIntSep *NumStrIntSeparator) SetRestartIntGroupingSequence(
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

// SetToUSADefaults - This method will overwrite and set the all
// the internal member variable data values to default values used
// in the United States. Integer separator values used in the
// United States consist of the comma character (','), an integer
// grouping of three ('3') and unlimited repetitions of this
// sequence.
//
//   United States Integer Separation Example:
//         '1,000,000,000,000'
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetToUSADefaults(
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
		"NumStrIntSeparator.SetToUSADefaults()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorMechanics{}.ptr().
		setToUSADefaults(
			nStrIntSep,
			ePrefix)
}

// SetToUSADefaultsIfEmpty - If any of the current
// NumStrIntSeparator instance data values are zero or invalid,
// this method will reset ALL data elements to United States
// default values.
//
// If the current NumStrIntSeparator instance is valid and
// populated with data, this method will take no action and exit.
//
// United States default numeric separators are listed as follows:
//
//  Decimal Separator = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits Grouping Sequence = 3
//  Example Floating Point Number String: 1,000,000,000.456
//
// IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetToUSADefaultsIfEmpty(
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
		"NumStrIntSeparator."+
			"SetToUSADefaultsIfEmpty()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorUtility{}.ptr().
		setToUSADefaultsIfEmpty(
			nStrIntSep,
			ePrefix.XCtx(
				"nStrIntSep"))
}

// String - Returns a string detailing the components of this
// NumStrIntSeparator instance. Output includes:
//   Integer Separator Characters
//   Integer Separator Grouping
//   Restart Grouping Sequence
//
func (nStrIntSep *NumStrIntSeparator) String() string {

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

	return str
}

// SetWithComponents - This method will overwrite and reset the
// internal member variable data values for the current
// NumStrIntSeparator instance based on the component elements
// passed as input parameters.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// current NumStrIntSeparator instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSeparatorChars          []rune
//     - A series of runes or characters used to separate integer
//       digits in a number string. These characters are commonly
//       known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer values specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'. Chinese Numerals
//       would be formatted like this: '12,3456,7890,2345'
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - If the NumStrIntSeparator is the last element in an array
//       of NumStrIntSeparator objects, this boolean flag signals
//       whether the entire integer grouping sequence will be
//       restarted from array element zero.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
func (nStrIntSep *NumStrIntSeparator) SetWithComponents(
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
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
		"NumStrIntSeparator.SetWithComponents()",
		"")

	if err != nil {
		return err
	}

	return numStrIntSeparatorMechanics{}.ptr().
		setWithComponents(
			nStrIntSep,
			intSeparatorChars,
			intSeparatorGrouping,
			intSeparatorRepetitions,
			restartIntGroupingSequence,
			ePrefix.XCtx("nStrIntSep"))
}
