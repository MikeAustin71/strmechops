package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberSignSymbol - Defines the text or character value of a numeric
// sign contained in a number string. The most common examples of
// number sign text values are the plus sign ('+') and the minus
// sign ('-').
//
// The number sign text is usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As an
// example, in the USA opening and closing parentheses are used to
// designate a negative number "(55)".
//
// Generally, number signs consist of a single text character,
// however there may be cases where multiple characters are used
// to designate positive or negative values.
//
// Since Go does not directly support enumerations, the 'TextJustify'
// type has been adapted to function in a manner similar to classic
// enumerations. 'TextJustify' is declared as a type 'int'. The
// method names effectively represent an enumeration of text
// justification formats. These methods are listed as follows:
//
//
type NumberSignSymbol struct {
	leadingNumSignChars  []rune
	trailingNumSignChars []rune
	symFoundInNumber     bool                  // Number Sign Symbol found in target number
	numSignPosition      NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	numSignType          NumericSignValueType  // Must be positive or negative
	lock                 *sync.Mutex
}

// GetLeadingNumSignChars - Returns a deep copy of the leading
// number sign characters contained in this instance of
// NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) GetLeadingNumSignChars() []rune {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var leadingNumSignChars []rune

	lenNumSignChars := len(nSignSymbol.leadingNumSignChars)

	if lenNumSignChars == 0 {
		return leadingNumSignChars
	}

	leadingNumSignChars = make([]rune, lenNumSignChars)

	copy(
		leadingNumSignChars,
		nSignSymbol.leadingNumSignChars)

	return leadingNumSignChars
}

// GetNumberSignArithmeticVal - Returns the arithmetic value of the
// number sign associated with this NumberSignSymbol instance.
//
// Valid return values are:
//  Negative == -1
//  Zero     ==  0
//  Positive ==  1
//
// If the Number Sign value is invalid, this method will return a
// value of -99
//
func (nSignSymbol *NumberSignSymbol) GetNumberSignArithmeticVal() int {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSignType.XArithmeticValue()
}

// GetNumSignSymPosition - Returns an enumeration value specifying
// the position of the Number Sign Symbol within a number string.
//
// Number Sign Symbols may be positioned before the numeric value,
// after the numeric value or before and after the numeric value.
//
// This method returns a type of NumSignSymbolPosition. Possible
// return values are listed as follows:
//    NumSignSymbolPosition(0).Before()
//    NumSignSymbolPosition(0).After()
//    NumSignSymbolPosition(0).BeforeAndAfter()
//
func (nSignSymbol *NumberSignSymbol) GetNumSignSymPosition() NumSignSymbolPosition {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSignPosition
}

// GetTrailingNumSignChars - Returns a deep copy of the trailing
// number sign characters contained in this instance of
// NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) GetTrailingNumSignChars() []rune {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var trailingNumSignChars []rune

	lenNumSignChars := len(nSignSymbol.trailingNumSignChars)

	if lenNumSignChars == 0 {
		return trailingNumSignChars
	}

	trailingNumSignChars = make([]rune, lenNumSignChars)

	copy(
		trailingNumSignChars,
		nSignSymbol.trailingNumSignChars)

	return trailingNumSignChars
}

// New - Creates and returns a new instance of NumberSignSymbol.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leadingNumberSign   string
//     - A string comprised of the characters which constitute this
//       leading number sign.
//
//       Examples: "-", "+", "("
//
//
//  trailingNumberSign  string
//     - A string comprised of the characters which constitute this
//       trailing number sign.
//
//       Examples: "-", "+", ")"
//
//
//  isNegativeValue     bool
//     - Number sign symbols will specify either a positive or
//       negative numeric value. If this parameter is set to
//       'true', the number sign will be treated as identifying a
//       negative numeric value. If this parameter is set to
//       'false', the number sign will be interpreted as
//       identifying a positive numeric value.
//
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
//  newNumSignSym       NumberSignSymbol
//     - If this method completes successfully, a new, populated
//       instance of NumberSignSymbol will be returned.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (nSignSymbol NumberSignSymbol) New(
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
	errorPrefix interface{}) (
	newNumSignSym NumberSignSymbol,
	err error) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberSignSymbol.New()",
		"")

	if err != nil {
		return newNumSignSym, err
	}

	lenLeadingNumSign := len(leadingNumberSign)
	lenTrailingNumSign := len(trailingNumberSign)

	if lenLeadingNumSign == 0 &&
		lenTrailingNumSign == 0 {
		err = fmt.Errorf("%v\n" +
			"Error: Input parameters 'leadingNumberSign' and " +
			"'trailingNumberSign' are zero length strings!\n" +
			ePrefix.String())
		return newNumSignSym, err

	} else if lenLeadingNumSign > 0 &&
		lenTrailingNumSign < 0 {

		newNumSignSym.numSignPosition =
			NumSymPos.Before()

		newNumSignSym.leadingNumSignChars =
			[]rune(leadingNumberSign)

	} else if lenLeadingNumSign == 0 &&
		lenTrailingNumSign > 0 {

		newNumSignSym.numSignPosition =
			NumSymPos.After()

		newNumSignSym.trailingNumSignChars =
			[]rune(trailingNumberSign)

	} else {
		// Must be lenLeadingNumSign > 0 &&
		// lenTrailingNumSign > 0
		newNumSignSym.numSignPosition =
			NumSymPos.BeforeAndAfter()

		newNumSignSym.leadingNumSignChars =
			[]rune(leadingNumberSign)

		newNumSignSym.trailingNumSignChars =
			[]rune(trailingNumberSign)

	}

	if isNegativeValue {
		newNumSignSym.numSignType = NumSignVal.Negative()
	} else {
		newNumSignSym.numSignType = NumSignVal.Positive()
	}

	return newNumSignSym, err
}
