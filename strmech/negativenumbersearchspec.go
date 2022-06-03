package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NegativeNumberSearchSpec - Negative Number Sign Specification.
// This type is used to configure search parameters for identifying
// negative numeric values within text strings when extracting or
// parsing numeric digits.
//
// Parsing functions used in converting strings of numeric
// characters into numeric values assume that those values are
// positive unless a Negative Number Sign symbol or symbols are
// present in the string of numeric digits.
//
// The NegativeNumberSearchSpec type is used to specify the
// criterion used to identify Negative Number Sign Symbols
// found in number strings and properly classify the extracted
// numeric values as positive or negative values.
//
type NegativeNumberSearchSpec struct {
	negNumSignPosition NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	//                                                   Negative Number Signs are classified
	//                                                   by their location relative to the
	//                                                   numeric digits in a number string.
	leadingNegNumSignSymbols  []rune
	trailingNegNumSignSymbols []rune

	// Processing flags

	parentNegNumSignCollection *NegNumSearchSpecCollection // If this pointer is not 'nil',
	//                                                       it signals that this instance is
	//                                                       a member of a collection. This
	//                                                       allows for efficient access to
	//                                                       Target Text Search Characters.

	negNumSignTargetSearchChars []rune // The target search text characters to be
	//                                       searched for a negative number sign
	//                                       symbols

	foundFirstNumericDigitInNumStr bool // Indicates first numeric digit
	//                                       the number string has been found
	foundNegNumSignSymbols bool // Indicates all negative number sign symbols
	//                               in this specification have been found
	foundLeadingNegNumSign       bool
	foundLeadingNegNumSignIndex  int
	foundTrailingNegNumSign      bool
	foundTrailingNegNumSignIndex int
	lock                         *sync.Mutex
}

// Empty - Resets all internal member variables for the current
// instance of NegativeNumberSearchSpec to their initial or zero
// states.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) Empty() {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	negNumSearchSpecAtom{}.ptr().empty(
		negNumSearchSpec)

	negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.lock = nil
}

// EmptyProcessingFlags - Resets all the internal processing flags
// to their initial or zero states.
//
// The NegativeNumberSearchSpec type includes a series of flags which
// are used during the process of identifying Negative Numeric Sign
// symbols within as number string. As part of parsing these number
// strings, the internal flags are used to record the status of a
// search for the Negative Number Sign symbol or symbols defined by
// the current instance of NegativeNumberSearchSpec.
//
// Calling this method will effectively clear all of these flags
// and prepare the current instance of NegativeNumberSearchSpec for
// a new number string parsing operation.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) EmptyProcessingFlags() {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpecElectron{}.ptr().emptyProcessingFlags(
		negNumSearchSpec)

}

// GetFoundFirstNumericDigit - This flag is set during a number
// string parsing operation.
//
// If the first numeric digit in a numeric value has been
// identified in the string parsing operation, the internal member
// variable 'foundFirstNumericDigitInNumStr' is set to the boolean
// value of 'true'. This member variable is typically set by the
// number string parsing routine.
//
// If the first numeric digit has not yet been located in the
// parsing operation, 'foundFirstNumericDigitInNumStr' is set to
// 'false'.
//
// This method returns the status flag
// ('foundFirstNumericDigitInNumStr') indicating whether the first
// numeric digit has been located in the number string parsing
// operation.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundFirstNumericDigit() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundFirstNumericDigitInNumStr
}

// GetFoundNegNumSignSymbols - This processing flag is set during a
// number string parsing operation.
//
// If the all the symbols comprising the Negative Number Sign
// defined by the current instance of NegativeNumberSearchSpec have
// been located within a number string, the internal member
// variable, 'foundNegNumSignSymbols' is set to true.
//
// Otherwise, 'foundNegNumSignSymbols' is set to false signaling
// that a negative number sign matching that defined by the current
// NegativeNumberSearchSpec instance has not yet been identified in
// the target number string.
//
// This internal member variable is typically set by the number
// string parsing routine.
//
// This method returns the status flag ('foundNegNumSignSymbols')
// indicating whether the Negative Number Sign Symbols defined by
// the current NegativeNumberSearchSpec instance have been located by
// the number string parsing routine.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundNegNumSignSymbols() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundNegNumSignSymbols
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current NegativeNumberSearchSpec
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
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
//  err                        error
//     - If any of the internal member data variables contained in
//       the current instance of NegativeNumberSearchSpec are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	if !negNumSearchSpec.negNumSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
			"The internal member variable 'negNumSignPosition' is NOT configured.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.Before() {

		if len(negNumSearchSpec.trailingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, it contains Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSearchSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, no Leading Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err
		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.After() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSearchSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, no Trailing Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.BeforeAndAfter() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign.\n"+
				"However, it contains NO Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		if len(negNumSearchSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains NO Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	return err
}

// NewLeadingNegNumSearchRunes - Returns a fully populated
// specification for a Leading Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.NewLeadingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives a rune array as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec   NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading Negative Number Sign Specification and returned to
//       the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		&newLeadingNegNumSignSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingNegNumSearchStr - Returns a fully populated
// specification for a Leading Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.NewLeadingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the character or characters which
//       comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length) an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading Negative Number Sign Specification and returned to
//       the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingNegNumSearchStr()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		&newLeadingNegNumSignSpec,
		[]rune(leadingNegNumSignSymbols),
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSearchRunes - Returns a fully populated
// specification for a Leading and Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameters
// 'leadingNegNumSignSymbols' and 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewLeadingAndTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives rune arrays as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading and Trailing Negative Number Sign Specification
//       and returned to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingAndTrailingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingAndTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			&leadingAndTrailingNegNumSignSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSearchStr - Returns a fully populated
// specification for a Leading and Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameters
// 'leadingNegNumSignSymbols' and 'trailingNegNumSignSymbols '.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, in the US and Canada
// parentheses "()" are used to indicate negative numeric
// values. Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewLeadingAndTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives strings as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the character or characters which
//       comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading and Trailing Negative Number Sign Specification
//       and returned to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingAndTrailingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingAndTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			&leadingAndTrailingNegNumSignSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewTrailingNegNumSearchRunes - Returns a fully populated specification
// for a Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives a rune array as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Trailing Negative Number Sign Specification and returned
//       to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewTrailingNegNumSearchRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			&trailingNegNumSignSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// NewTrailingNegNumSearchStr - Returns a fully populated specification
// for a Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives a string as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Trailing Negative Number Sign Specification and returned
//       to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewTrailingNegNumSearchStr(
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			&trailingNegNumSignSpec,
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// SearchForNegNumSignSymbols - This method is typically called by
// a number string parsing routine attempting to determine if the
// characters in a search string match the Negative Number Sign
// Symbol defined by this current instance of NegativeNumberSearchSpec.
//
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SearchForNegNumSignSymbols(
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errorPrefix interface{}) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SearchForNegNumSignSymbols()",
		"")

	if err != nil {

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	var err2 error

	err2 = negNumSearchSpec.IsValidInstanceError(
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of NegativeNumberSearchSpec\n"+
			"is invalid. The Number String parsing operation has been aborted.\n"+
			"Validation checks returned the following error for this intance of\n"+
			"NegativeNumberSearchSpec:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	negNumSignAtom := negNumSearchSpecAtom{}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.Before() {

		foundNegNumSignSymbols,
			lastIndex,
			err =
			negNumSignAtom.leadingNegSignSymSearch(
				negNumSearchSpec,
				foundFirstNumericDigitInNumStr,
				startingSearchIndex,
				ePrefix)

	} else if negNumSearchSpec.negNumSignPosition == NSignSymPos.After() {

	} else {
		// Must be: NSignSymPos.BeforeAndAfter()

	}

	return foundNegNumSignSymbols,
		lastIndex,
		err
}

// SetForNumberStringSearch - Call this method once to configure
// internal member variables in preparation for a number string
// parsing operation.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// Call this method once at the BEGINNING of a number string
// processing operation.
//
// DO NOT call this method again until that number string
// parsing operation is completed.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetSearchChars          []rune
//     - An array of runes containing the text characters to be
//       searched as part of a number string parsing operation.
//       This rune array is set once for each instance of
//       NegativeNumberSearchSpec.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetForNumberStringSearch(
	targetSearchChars []rune,
	errorPrefix interface{}) (err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SearchForNegNumSignSymbols()",
		"")

	if err != nil {
		return err
	}

	negNumSearchSpecElectron{}.ptr().emptyProcessingFlags(
		negNumSearchSpec)

	lenTargetSearchChars := len(targetSearchChars)

	if lenTargetSearchChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSearchChars' is empty and invalid!\n",
			ePrefix.String())

	}

	negNumSearchSpec.negNumSignTargetSearchChars =
		make([]rune, lenTargetSearchChars)

	itemsCopied := copy(
		negNumSearchSpec.negNumSignTargetSearchChars,
		targetSearchChars)

	if itemsCopied != lenTargetSearchChars {

		err = fmt.Errorf("%v\n"+
			"Error: 'targetSearchChars' copy operation failed!\n"+
			"Expected %v characters would be copied from\n"+
			"'targetSearchChars' to \n"+
			"'negNumSearchSpec.negNumSignTargetSearchChars'\n"+
			"However, only %v characters out of a total of\n"+
			"%v characters were copied.\n",
			ePrefix.String(),
			lenTargetSearchChars,
			itemsCopied,
			lenTargetSearchChars)

		negNumSearchSpec.negNumSignTargetSearchChars = nil

		return err
	}

	return err
}

// SetFoundNegNumSignSymbols - Sets the processing flag describing
// the results of a number string parsing operation.
//
// If the all the symbols comprising the Negative Number Sign
// defined by the current instance of NegativeNumberSearchSpec have
// been located within a number string, the internal member
// variable, 'foundNegNumSignSymbols' is set to true.
//
// Otherwise, 'foundNegNumSignSymbols' is set to false signaling
// that a negative number sign matching that defined by the
// current NegativeNumberSearchSpec instance has not yet been
// identified in the target number string.
//
// This internal member variable is typically set by the number
// string parsing routine when calling method:
//    NegativeNumberSearchSpec.SearchForNegNumSignSymbols()
//
// This method sets the processing status flag
// ('foundNegNumSignSymbols') indicating whether the Negative
// Number Sign Symbols defined by the current
// NegativeNumberSearchSpec instance have been located by the number
// string parsing routine.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetFoundNegNumSignSymbols(
	foundNegNumSignSymbols bool) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.foundNegNumSignSymbols =
		foundNegNumSignSymbols
}

// SetFoundFirstNumericDigit - Sets the internal member variable,
// 'foundFirstNumericDigitInNumStr'.
//
// This flag is typically set during a number string parsing
// operation. If the first numeric digit in a numeric value has
// been identified in the string parsing operation, this internal
// member variable is set to the boolean value of 'true'. Again,
// this member variable is typically set by the number string
// parsing routine.
//
// If the first numeric digit in a number string has not yet been
// identified, this flag is set to 'false'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetFoundFirstNumericDigit(
	foundFirstNumericDigitInNumStr bool) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr
}

// SetLeadingNegNumSearchRunes - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Leading Negative Number Sign
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.SetLeadingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading Negative Number Sign Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the current NegativeNumberSearchSpec instance as
//       a Leading Negative Number Sign Specification.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		negNumSearchSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingNegNumSearchStr  - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Leading Negative Number Sign
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.SetLeadingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading Negative Number Sign Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A strung identifying the character or characters which
//       comprise the Leading Negative Number Symbol used in
//       configuring the current NegativeNumberSearchSpec instance as
//       a Leading Negative Number Sign Specification.
//
//       If this string is empty (has a zero length) an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(leadingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingAndTrailingNegNumSearchRunes - Reconfigures the current
// instance of NegativeNumberSearchSpec as a Leading and Trailing
// Negative Number Sign Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input
// parameters 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetLeadingAndTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives rune arrays as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading and Trailing Negative Number Sign
// Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the current instance of NegativeNumberSearchSpec.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingAndTrailingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingAndTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			negNumSearchSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return err
}

// SetLeadingAndTrailingNegNumSearchStr - Reconfigures the current
// instance of NegativeNumberSearchSpec as a Leading and Trailing
// Negative Number Sign Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input
// parameters 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetLeadingAndTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives strings as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading and Trailing Negative Number Sign
// Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the character or characters which
//       comprise the Leading Negative Number Symbol used in
//       configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingAndTrailingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingAndTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec"))

	return err
}

// SetTrailingNegNumSearchRunes - Reconfigures the current instance
// of NegativeNumberSearchSpec as a Trailing Negative Number Sign
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Trailing Negative Number Sign Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the current NegativeNumberSearchSpec instance as
//       a Trailing Negative Number Sign Specification.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetTrailingNegNumSearchRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			negNumSearchSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"negNumSearchSpec<-trailingNegNumSignSymbols"))

	return err
}

// SetTrailingNegNumSearchStr - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Trailing Negative Number Sign
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Trailing Negative Number Sign Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the current NegativeNumberSearchSpec instance as
//       a Trailing Negative Number Sign Specification.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Trailing Negative Number Sign Specification and returned
//       to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetTrailingNegNumSearchStr(
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec<-trailingNegNumSignSymbols"))

	return err
}
