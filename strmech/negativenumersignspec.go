package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NegativeNumberSignSpec - Negative Number Sign Specification.
// This type is used to configure parameters used in identifying
// negative numeric value when parsing a string of numeric digits.
//
// Parsing functions used in converting strings of numeric
// characters into numeric values assume that those values are
// positive unless a Negative Number Sign symbol or symbols are
// present in the string of numeric digits.
//
// The NegativeNumberSignSpec type is used to identify Negative
// Number Sign symbols found in the parsing of number strings.
//
type NegativeNumberSignSpec struct {
	negNumSignPosition        NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	leadingNegNumSignSymbols  []rune
	trailingNegNumSignSymbols []rune
	// Processing flags
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
// instance of NegativeNumberSignSpec to their initial or zero
// states.
//
func (negNumSignSpec *NegativeNumberSignSpec) Empty() {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	negNumSignSpecAtom{}.ptr().empty(
		negNumSignSpec)

	negNumSignSpec.lock.Unlock()

	negNumSignSpec.lock = nil
}

// EmptyProcessingFlags - Resets all the internal processing flags
// to their initial or zero states.
//
// The NegativeNumberSignSpec type includes a series of flags which
// are used during the process of identifying Negative Numeric Sign
// symbols within as number string. As part of parsing these number
// strings, the internal flags are used to record the status of a
// search for the Negative Number Sign symbol or symbols defined by
// the current instance of NegativeNumberSignSpec.
//
// Calling this method will effectively clear all of these flags
// and prepare the current instance of NegativeNumberSignSpec for
// a new number string parsing operation.
//
func (negNumSignSpec *NegativeNumberSignSpec) EmptyProcessingFlags() {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	negNumSignSpecElectron{}.ptr().emptyProcessingFlags(
		negNumSignSpec)

	negNumSignSpec.lock.Unlock()

	negNumSignSpec.lock = nil
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current NegativeNumberSignSpec
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
//       the current instance of NegativeNumberSignSpec are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (negNumSignSpec *NegativeNumberSignSpec) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	if !negNumSignSpec.negNumSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
			"The internal member variable 'negNumSignPosition' is NOT configured.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	if negNumSignSpec.negNumSignPosition == NSignSymPos.Before() {

		if len(negNumSignSpec.trailingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, it contains Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSignSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, no Leading Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err
		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
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

	if negNumSignSpec.negNumSignPosition == NSignSymPos.After() {

		if len(negNumSignSpec.leadingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSignSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, no Trailing Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
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

	if negNumSignSpec.negNumSignPosition == NSignSymPos.BeforeAndAfter() {

		if len(negNumSignSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign.\n"+
				"However, it contains NO Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		if len(negNumSignSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains NO Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
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

// NewLeadingNegNumSignRunes - Returns a fully populated
// specification for a Leading Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// This method is identical in function to the method:
//  NegativeNumberSignSpec.NewLeadingNegNumSignStr()
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
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingNegNumSignRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewLeadingNegNumSignRunes()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSpecNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSignSpec(
		&newLeadingNegNumSignSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingNegNumSignStr - Returns a fully populated
// specification for a Leading Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// This method is identical in function to the method:
//  NegativeNumberSignSpec.NewLeadingNegNumSignRunes()
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
//     - A strung identifying the character or characters which
//       comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingNegNumSignStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewLeadingNegNumSignStr()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSpecNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSignSpec(
		&newLeadingNegNumSignSpec,
		[]rune(leadingNegNumSignSymbols),
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSignRunes - Returns a fully populated
// specification for a Leading and Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameters
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
//    NegativeNumberSignSpec.NewLeadingAndTrailingNegNumSignStr()
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
//       configuring the NegativeNumberSignSpec instance returned
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
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingAndTrailingNegNumSignRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewLeadingAndTrailingNegNumSignRunes()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setLeadingAndTrailingNegNumSignSpec(
			&leadingAndTrailingNegNumSignSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSignStr - Returns a fully populated
// specification for a Leading and Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameters
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
//    NegativeNumberSignSpec.NewLeadingAndTrailingNegNumSignRunes()
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
//       configuring the NegativeNumberSignSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingAndTrailingNegNumSignStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewLeadingAndTrailingNegNumSignStr()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setLeadingAndTrailingNegNumSignSpec(
			&leadingAndTrailingNegNumSignSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewTrailingNegNumSignRunes - Returns a fully populated specification
// for a Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// This method is identical in function to method:
//    NegativeNumberSignSpec.NewTrailingNegNumSignStr()
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
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewTrailingNegNumSignRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewTrailingNegNumSignRunes()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setTrailingNegNumSignSpec(
			&trailingNegNumSignSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// NewTrailingNegNumSignStr - Returns a fully populated specification
// for a Trailing Negative Number Sign.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// This method is identical in function to method:
//    NegativeNumberSignSpec.NewTrailingNegNumSignStr()
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
//       configuring the NegativeNumberSignSpec instance returned
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
//  NegativeNumberSignSpec     NegativeNumberSignSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSignSpec will be configured as a
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
func (negNumSignSpec NegativeNumberSignSpec) NewTrailingNegNumSignStr(
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewTrailingNegNumSignStr()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setTrailingNegNumSignSpec(
			&trailingNegNumSignSpec,
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// SetLeadingNegNumSignRunes - Reconfigures the current instance of
// NegativeNumberSignSpec as a Leading Negative Number Sign
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSignSpec are reconfigured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// This method is identical in function to the method:
//  NegativeNumberSignSpec.SetLeadingNegNumSignStr()
//
// The only difference between the two methods is that this method
// receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSignSpec
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
//       configuring the current NegativeNumberSignSpec instance as
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
func (negNumSignSpec *NegativeNumberSignSpec) SetLeadingNegNumSignRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec.NewLeadingNegNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	negNumSignNanobot := negNumSignSpecNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSignSpec(
		negNumSignSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"negNumSignSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingNegNumSignStr  - Reconfigures the current instance of
// NegativeNumberSignSpec as a Leading Negative Number Sign
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSignSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// This method is identical in function to the method:
//  NegativeNumberSignSpec.SetLeadingNegNumSignRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSignSpec
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
//       configuring the current NegativeNumberSignSpec instance as
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
func (negNumSignSpec *NegativeNumberSignSpec) SetLeadingNegNumSignStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"SetLeadingNegNumSignStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setLeadingNegNumSignSpec(
			negNumSignSpec,
			[]rune(leadingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSignSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingAndTrailingNegNumSignRunes - Reconfigures the current
// instance of NegativeNumberSignSpec as a Leading and Trailing
// Negative Number Sign Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSignSpec are reconfigured using the input
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
//    NegativeNumberSignSpec.SetLeadingAndTrailingNegNumSignStr()
//
// The only difference between to the two methods is that this
// method receives rune arrays as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSignSpec
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
//       NegativeNumberSignSpec.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the current instance of NegativeNumberSignSpec.
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
func (negNumSignSpec *NegativeNumberSignSpec) SetLeadingAndTrailingNegNumSignRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"SetLeadingAndTrailingNegNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setLeadingAndTrailingNegNumSignSpec(
			negNumSignSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return err
}

// SetLeadingAndTrailingNegNumSignStr - Reconfigures the current
// instance of NegativeNumberSignSpec as a Leading and Trailing
// Negative Number Sign Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSignSpec are reconfigured using the input
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
//    NegativeNumberSignSpec.SetLeadingAndTrailingNegNumSignRunes()
//
// The only difference between to the two methods is that this
// method receives strings as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSignSpec
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
//       NegativeNumberSignSpec.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the current instance of
//       NegativeNumberSignSpec.
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
func (negNumSignSpec *NegativeNumberSignSpec) SetLeadingAndTrailingNegNumSignStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"SetLeadingAndTrailingNegNumSignStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setLeadingAndTrailingNegNumSignSpec(
			negNumSignSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSignSpec"))

	return err
}

// SetTrailingNegNumSignRunes - Reconfigures the current instance of
// NegativeNumberSignSpec as a Trailing Negative Number Sign
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSignSpec are reconfigured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// This method is identical in function to method:
//    NegativeNumberSignSpec.SetTrailingNegNumSignStr()
//
// The only difference between to the two methods is that this
// method receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSignSpec
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
//       configuring the NegativeNumberSignSpec instance returned
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
func (negNumSignSpec *NegativeNumberSignSpec) SetTrailingNegNumSignRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"SetTrailingNegNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSpecNanobot{}.ptr().
		setTrailingNegNumSignSpec(
			negNumSignSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"negNumSignSpec<-trailingNegNumSignSymbols"))

	return err
}
