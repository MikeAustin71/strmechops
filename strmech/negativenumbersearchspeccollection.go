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
// This method will create a new instance of NegativeNumberSearchSpec
// and add it to the end of the collection.
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
//       Symbol used in configuring the NegativeNumberSearchSpec
//       instance will be added to the end of the Negative Number
//       Search Specification Collectionmaintained by the current
//       instance of NegNumSearchSpecCollection.
//
//       If this runa array is empty (has a zero length), an error
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
			"AddLeadingNegNumSearchRunes()",
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

// AddLeadingNegNumSearchStr - Adds a Leading Negative Number search
// profile to the end of the collection of NegativeNumberSearchSpec
// objects maintained by the current instance of
// NegNumSearchSpecCollection.
//
// This method will create a new instance of NegativeNumberSearchSpec
// and add it to the end of the collection.
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
//       configuring the NegativeNumberSearchSpec instance will be
//       added to the end of the Negative Number Search
//       Specification Collection maintained by the current
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
