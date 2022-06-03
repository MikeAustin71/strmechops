package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type NegNumSearchSpecCollection struct {
	negNumSignSpecs             []NegativeNumberSearchSpec
	negNumSignTargetSearchChars []rune
	lock                        *sync.Mutex
}

// AddLeadingNegNumSearchRunes - Adds a Leading Negative Number search
// profile to the end of the collection of NegativeNumberSearchSpec
// objects maintained by the current instance of
// NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Leading Negative Number Search Specification.
//
// This method is identical in function to the method:
//  NegNumSearchSpecCollection.AddLeadingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives an array of runes as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbol used in configuring a NegativeNumberSearchSpec
//       instance. This instance will be added to the end of the
//       Negative Number Search Specification Collection maintained
//       by the current instance of NegNumSearchSpecCollection.
//
//       If this rune array is empty (has a zero length), an error
//       will be returned.
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
func (negNumSignCol *NegNumSearchSpecCollection) AddLeadingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"AddLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingNegNumSearchRunes(
			leadingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingNegNumSign<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSignSpecs =
		append(
			negNumSignCol.negNumSignSpecs,
			newLeadingNegNumSign)

	return err
}

// AddLeadingNegNumSearchStr - Adds a Leading Negative Number search
// profile to the end of the collection of NegativeNumberSearchSpec
// objects maintained by the current instance of
// NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Leading Negative Number Search Specification.
//
// This method is identical in function to the method:
//  NegNumSearchSpecCollection.AddLeadingNegNumSearchRunes()
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
//     - A string identifying the text character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring a NegativeNumberSearchSpec instance, This
//       instance will be added to the end of the Negative Number
//       Search Specification Collection maintained by the current
//       instance of NegNumSearchSpecCollection.
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
func (negNumSignCol *NegNumSearchSpecCollection) AddLeadingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingNegNumSearchStr(
			leadingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingNegNumSign<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSignSpecs =
		append(
			negNumSignCol.negNumSignSpecs,
			newLeadingNegNumSign)

	return err
}

// AddLeadingAndTrailingNegNumSearchRunes - Adds a Leading and
// Trailing Negative Number search profile to the end of the
// collection of NegativeNumberSearchSpec objects maintained by the
// current instance of NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Leading and Trailing Negative Number Search
// Specification.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Signs. As an example, the US and Canada use parentheses "()" to
// indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method receives two input parameters. The first parameter
// specifies the Leading Negative Number Symbols while the second
// specifies the Trailing Negative Number Symbols. Both parameters
// are required in order to construct a Leading and Trailing
// Negative Number Search Specification.
//
// This method is identical in function to the method:
// NegNumSearchSpecCollection.AddLeadingAndTrailingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives arrays of runes as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbols used in configuring a NegativeNumberSearchSpec
//       instance. This instance will be added to the end of the
//       Negative Number Search Specification Collection maintained
//       by the current instance of NegNumSearchSpecCollection.
//
//       If this rune array is empty (has a zero length), an error
//       will be returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
//       Symbols used in configuring a NegativeNumberSearchSpec
//       instance. This instance will be added to the end of the
//       Negative Number Search Specification Collection maintained
//       by the current instance of NegNumSearchSpecCollection.
//
//       If this rune array is empty (has a zero length), an error
//       will be returned.
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
func (negNumSignCol *NegNumSearchSpecCollection) AddLeadingAndTrailingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"AddLeadingAndTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingAndTrailingNegNumSearchRunes(
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingNegNumSign"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSignSpecs =
		append(
			negNumSignCol.negNumSignSpecs,
			newLeadingNegNumSign)

	return err
}

// AddLeadingAndTrailingNegNumSearchStr - Adds a Leading and
// Trailing Negative Number search profile to the end of the
// collection of NegativeNumberSearchSpec objects maintained by the
// current instance of NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Leading and Trailing Negative Number Search
// Specification.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Signs. As an example, the US and Canada use parentheses "()" to
// indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method receives two input parameters. The first parameter
// specifies the Leading Negative Number Symbols while the second
// specifies the Trailing Negative Number Symbols. Both parameters
// are required in order to construct a Leading and Trailing
// Negative Number Search Specification.
//
// This method is identical in function to the method:
// NegNumSearchSpecCollection.AddLeadingAndTrailingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives strings as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the text character or characters
//       which comprise the Leading Negative Number Symbols used in
//       configuring a NegativeNumberSearchSpec instance. This
//       instance will be added to the end of the Negative Number
//       Search Specification Collection maintained by the current
//       instance of NegNumSearchSpecCollection.
//
//       If this string is empty (has a zero length), an error
//       will be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the text character or characters
//       which comprise the Trailing Negative Number Symbols used
//       in configuring a NegativeNumberSearchSpec instance. This
//       instance will be added to the end of the Negative Number
//       Search Specification Collection maintained by the current
//       instance of NegNumSearchSpecCollection.
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
func (negNumSignCol *NegNumSearchSpecCollection) AddLeadingAndTrailingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"AddLeadingAndTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingAndTrailingNegNumSearchStr(
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingNegNumSign"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSignSpecs =
		append(
			negNumSignCol.negNumSignSpecs,
			newLeadingNegNumSign)

	return err
}
