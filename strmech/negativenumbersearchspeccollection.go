package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type NegNumSearchSpecCollection struct {
	negNumSearchSpecsCol []NegativeNumberSearchSpec // The Collection of negative number sign
	//                                                 search specification objects.

	foundNegNumSign bool // When set to 'true' this signals that the current search
	//                      operation has located a valid negative number sign.

	foundNegNumSignColIndex int // The Collection index at which a negative number sign
	//                             symbol was found during the current search operation.

	lock *sync.Mutex
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
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
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

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
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
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
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

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
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

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
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

	var newLeadingTrailingNegNumSign NegativeNumberSearchSpec

	newLeadingTrailingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingAndTrailingNegNumSearchStr(
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingTrailingNegNumSign"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
			newLeadingTrailingNegNumSign)

	return err
}

// AddTrailingNegNumSearchRunes - Adds a Trailing Negative Number
// search profile to the end of the collection of
// NegativeNumberSearchSpec objects maintained by the current
// instance of NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Trailing Negative Number Search Specification.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to the method:
//  NegNumSearchSpecCollection.AddTrailingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives an array of runes as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  TrailingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
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
func (negNumSignCol *NegNumSearchSpecCollection) AddTrailingNegNumSearchRunes(
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
			"AddLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewTrailingNegNumSearchRunes(
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"newLeadingNegNumSign<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
			newLeadingNegNumSign)

	return err
}

// AddTrailingNegNumSearchStr - Adds a Trailing Negative Number search
// profile to the end of the collection of NegativeNumberSearchSpec
// objects maintained by the current instance of
// NegNumSearchSpecCollection.
//
// This method will create a new instance of
// NegativeNumberSearchSpec and add it to the end of the
// NegativeNumberSearchSpec collection. This new instance will be
// configured as a Trailing Negative Number Search Specification.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to the method:
//  NegNumSearchSpecCollection.AddTrailingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  TrailingNegNumSignSymbols   string
//     - A string identifying the text character or characters
//       which comprise the Trailing Negative Number Symbol used in
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
func (negNumSignCol *NegNumSearchSpecCollection) AddTrailingNegNumSearchStr(
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
			"AddTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	var newTrailingNegNumSign NegativeNumberSearchSpec

	newTrailingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewTrailingNegNumSearchStr(
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"newTrailingNegNumSign"))

	if err != nil {
		return err
	}

	negNumSignCol.negNumSearchSpecsCol =
		append(
			negNumSignCol.negNumSearchSpecsCol,
			newTrailingNegNumSign)

	return err
}

// Empty - Resets all internal member variables for the current
// instance of NegNumSearchSpecCollection to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NegNumSearchSpecCollection.
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
func (negNumSignCol *NegNumSearchSpecCollection) Empty() {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	negNumSearchSpecCollectionAtom{}.ptr().empty(
		negNumSignCol)

	negNumSignCol.lock.Unlock()

	negNumSignCol.lock = nil

	return
}

// GetNumberOfNegNumSearchSpecs - Returns the number of elements in
// the Negative Number Search Specification Collection maintained
// by the current instance of NegNumSearchSpecCollection.
//
// Type NegNumSearchSpecCollection manages an internal member
// variable named 'negNumSearchSpecsCol'. This is an array of
// NegativeNumberSearchSpec objects. This method returns the number
// of NegativeNumberSearchSpec objects in that array.
//
// Type NegativeNumberSearchSpec is used in number string parsing
// functions to search for negative number symbols. The presence
// of valid negative number character symbols in a number string
// classify the extracted numeric digits as a negative numeric
// value.
//
// -----------------------------------------------------------------
//
// BE ADVISED
//
// If the number of elements in the 'negNumSearchSpecsCol'
// collection is zero, it means that the current instance of
// NegNumSearchSpecCollection is invalid and cannot be used in
// number string parsing functions to search for negative number
// symbols.
//
func (negNumSignCol *NegNumSearchSpecCollection) GetNumberOfNegNumSearchSpecs() int {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	return len(negNumSignCol.negNumSearchSpecsCol)
}

// GetNegNumSearchSpecCollection - Returns an array of
// NegativeNumberSearchSpec objects.
//
// Type NegNumSearchSpecCollection manages an internal member
// variable named 'negNumSearchSpecsCol'. This is an array of
// NegativeNumberSearchSpec objects. This method returns a deep
// copy of the NegativeNumberSearchSpec objects in that array.
//
// Type NegativeNumberSearchSpec is used in number string parsing
// functions to search for negative number symbols. The presence
// of valid negative number character symbols in a number string
// classify the extracted numeric digits as a negative numeric
// value.
//
// -----------------------------------------------------------------
//
// BE ADVISED
//
// If the length of the 'negNumSearchSpecsCol' collection is zero,
// this method will return an error.
//
// Also, if the number of elements in the 'negNumSearchSpecsCol'
// collection is zero, it means that the current instance of
// NegNumSearchSpecCollection is invalid and cannot be used in
// number string parsing functions to search for negative number
// symbols.
//
func (negNumSignCol *NegNumSearchSpecCollection) GetNegNumSearchSpecCollection(
	errorPrefix interface{}) (
	negNumSearchSpecCol []NegativeNumberSearchSpec,
	err error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	negNumSearchSpecCol = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetTargetSearchString()",
		"")

	if err != nil {

		return negNumSearchSpecCol, err

	}

	lenNegNumSearchSpecsCol :=
		len(negNumSignCol.negNumSearchSpecsCol)

	if lenNegNumSearchSpecsCol == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The length of the Negative Number Search\n"+
			"Spec Collection is zero. The Collection is empty!\n"+
			"len(negNumSignCol.negNumSearchSpecsCol) == 0\n"+
			"The current instance of NegNumSearchSpecCollection\n"+
			"is therefore invalid.\n",
			ePrefix.String())

		return negNumSearchSpecCol, err
	}

	negNumSearchSpecCol =
		make([]NegativeNumberSearchSpec, lenNegNumSearchSpecsCol)

	itemsCopied :=
		copy(negNumSearchSpecCol,
			negNumSignCol.negNumSearchSpecsCol)

	if itemsCopied != lenNegNumSearchSpecsCol {

		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"The number of elements copied from the Negative\n"+
			"Number Search Spec Collection is invalid.\n"+
			"'negNumSearchSpecsCol' elements available = '%v'\n"+
			"   'negNumSearchSpecsCol' elements copied = '%v'\n",
			ePrefix.String(),
			lenNegNumSearchSpecsCol,
			itemsCopied)

		return negNumSearchSpecCol, err
	}

	return negNumSearchSpecCol, err
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current NegNumSearchSpecCollection instance
// to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return a boolean value of 'false'.
//
// If all data elements are determined to be valid, this method
// returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// NegNumSearchSpecCollection.IsValidInstanceError() with the sole
// exceptions being that this method takes no input parameters and
// returns a boolean value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If any of the internal member data variables contained in
//       the current instance of NegNumSearchSpecCollection are found to
//       be invalid, this method will return a boolean value of
//       'false'.
//
//       If all internal member data variables contained in the
//       current instance of NegNumSearchSpecCollection are found to be
//       valid, this method returns a boolean value of 'true'.
//
func (negNumSignCol *NegNumSearchSpecCollection) IsValidInstance() bool {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	negNumSearchColAtom := negNumSearchSpecCollectionAtom{}

	isValid,
		_ := negNumSearchColAtom.testValidityOfNegNumSearchCol(
		negNumSignCol,
		nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current NegNumSearchSpecCollection
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
//       the current instance of NegNumSearchSpecCollection are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (negNumSignCol *NegNumSearchSpecCollection) IsValidInstanceError(
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
		"NegNumSearchSpecCollection."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = negNumSearchSpecCollectionAtom{}.ptr().
		testValidityOfNegNumSearchCol(
			negNumSignCol,
			ePrefix)

	return err
}

// New - Returns an unpopulated or empty instance of
// NegNumSearchSpecCollection.
//
// NegNumSearchSpecCollection or Negative Number Sign Search
// Collection is a collection of NegativeNumberSearchSpec objects.
//
// NegativeNumberSearchSpec or Negative Number Search
// Specifications are used to define Negative Number Sign Symbols
// used in identifying negative numeric values in Number Strings.
//
// A Number String is a string of text characters which contain
// numeric digit characters.
//
// The NegNumSearchSpecCollection is primarily used by Number
// String parsing functions. These functions review a string of
// text characters searching for numeric digits. The numeric digits
// are extracted to form numeric values. Number string parsing
// functions therefore convert numeric text characters to valid
// numeric values. A key feature of this conversion is the
// classification of these resulting numeric values as either
// positive or negative values. If properly configured Negative
// Number Sign Symbols are found in the Number String, the
// resulting numeric value is classified as negative. If Negative
// Number Sign Symbols are NOT found in the Number String, the
// resulting numeric value is classified as positive.
//
// The NegNumSearchSpecCollection allows the user to configure
// multiple Negative Number Sign Symbol objects. If any one of the
// configured Negative Number Sign Symbols is found in a number
// string, the parsing function will classify the numeric value as
// negative.
//
// This method will return a new instance of
// NegNumSearchSpecCollection with an unpopulated or empty
// collection. Thereafter, the user is responsible for calling one
// or more of the 'Add' methods in order to populate the
// collection.
//
// The 'Add' methods are listed as follows:
//  NegNumSearchSpecCollection.AddLeadingNegNumSearchRunes()
//  NegNumSearchSpecCollection.AddLeadingNegNumSearchStr()
//  NegNumSearchSpecCollection.AddLeadingNegNumSearchRunes()
//  NegNumSearchSpecCollection.AddLeadingAndTrailingNegNumSearchRunes()
//  NegNumSearchSpecCollection.AddLeadingAndTrailingNegNumSearchStr()
//  NegNumSearchSpecCollection.AddTrailingNegNumSearchRunes()
//  NegNumSearchSpecCollection.AddTrailingNegNumSearchStr()
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegNumSearchSpecCollection
//     - This method returns an unpopulated or empty instance of
//       NegNumSearchSpecCollection. Thereafter, it is up to the
//       user to add custom NegativeNumberSearchSpec objects.
//
func (negNumSignCol NegNumSearchSpecCollection) New() NegNumSearchSpecCollection {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	newNegNumSearchCol := NegNumSearchSpecCollection{}

	return newNegNumSearchCol
}

// NewUS - Returns an instance of NegNumSearchSpecCollection. This
// Negative Number Search Specification collection is configured
// with negative number signs commonly used in the United States.
//
// This configuration includes two negative number signs:
//   1. Leading Minus Sign   -123.45
//   2. Opposing Parentheses  (123.45)
//
func (negNumSignCol NegNumSearchSpecCollection) NewUS(
	errorPrefix interface{}) (
	newNegNumSearchCol NegNumSearchSpecCollection,
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
		"NegNumSearchSpecCollection."+
			"NewUnitedStates()",
		"")

	if err != nil {
		return newNegNumSearchCol, err
	}

	var newLeadingNegNumSign NegativeNumberSearchSpec

	newLeadingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingNegNumSearchRunes(
			[]rune("-"),
			ePrefix.XCpy(
				"newLeadingNegNumSign<-[]rune(\"-\")"))

	if err != nil {
		return newNegNumSearchCol, err
	}

	newNegNumSearchCol.negNumSearchSpecsCol =
		append(newNegNumSearchCol.negNumSearchSpecsCol,
			newLeadingNegNumSign)

	var newLeadingTrailingNegNumSign NegativeNumberSearchSpec

	newLeadingTrailingNegNumSign,
		err =
		NegativeNumberSearchSpec{}.NewLeadingAndTrailingNegNumSearchStr(
			"(",
			")",
			ePrefix.XCpy(
				"newLeadingTrailingNegNumSign<-'()'"))

	if err != nil {
		return newNegNumSearchCol, err
	}

	newNegNumSearchCol.negNumSearchSpecsCol =
		append(newNegNumSearchCol.negNumSearchSpecsCol,
			newLeadingTrailingNegNumSign)

	return newNegNumSearchCol, err
}

// SearchForNegNumSignSymbols - Performs a Negative Number Sign
// Symbol search on a single index in a Target Search String.
//
// Before beginning a search operation, call method:
// NegNumSearchSpecCollection.SetupForNegativeNumberSignSearch()
//
// This will prepare the current instance of
// NegNumSearchSpecCollection to conduct a Negative Number Sign
// Symbol search operation on an entire Target Search String.
//
func (negNumSignCol *NegNumSearchSpecCollection) SearchForNegNumSignSymbols(
	targetInputParms CharSearchTargetInputParametersDto,
	errorPrefix interface{}) (
	CharSearchNegativeNumberResultsDto,
	error) {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	searchResults := CharSearchNegativeNumberResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegNumSearchSpecCollection."+
			"SearchForNegNumSignSymbols()",
		"")

	if err != nil {

		return searchResults, err

	}

	lenNegNumSpecsCol := len(negNumSignCol.negNumSearchSpecsCol)

	if lenNegNumSpecsCol == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NegNumSearchSpecCollection is invalid!\n"+
			"The Negatvie Number Search Specs collection is empty.\n"+
			"The internal member \n"+
			"for Negative Number Sign symbols.\n",
			ePrefix.String())

		return searchResults, err
	}

	var err2 error

	for i := 0; i < lenNegNumSpecsCol; i++ {

		searchResults,
			err2 = negNumSignCol.negNumSearchSpecsCol[i].
			SearchForNegNumSignSymbols(
				targetInputParms,
				nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned by Negative Number Search Specifiation Collection\n"+
				"negNumSearchSpecsCol[%v].SearchForNegNumSignSymbols().\n"+
				" Error text follows:\n"+
				"%v\n",
				ePrefix,
				i,
				err2.Error())

			return searchResults, err
		}

		if searchResults.FoundNegativeNumberSymbols {

			negNumSignCol.foundNegNumSign = true

			searchResults.CollectionTestObjIndex = i

			negNumSignCol.foundNegNumSignColIndex =
				searchResults.CollectionTestObjIndex

			return searchResults, err
		}
	}

	return searchResults, err
}

// SetupForNegativeNumberSignSearch - Prepares the current instance
// of NegNumSearchSpecCollection for a search operation.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// Call this method each time before calling:
//   NegNumSearchSpecCollection.SearchForNegNumSignSymbols()
//
func (negNumSignCol *NegNumSearchSpecCollection) SetupForNegativeNumberSignSearch() {

	if negNumSignCol.lock == nil {
		negNumSignCol.lock = new(sync.Mutex)
	}

	negNumSignCol.lock.Lock()

	defer negNumSignCol.lock.Unlock()

	negNumSignCol.foundNegNumSign = false

	negNumSignCol.foundNegNumSignColIndex = -1
}
