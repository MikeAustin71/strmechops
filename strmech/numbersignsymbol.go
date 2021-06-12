package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberSignSymbol - Stores the specification for a single text or
// character number sign symbol. This is usually a single plus
// ('+') or minus ('-) for a numeric sign contained in a number
// string.
//
// The NumberSignSymbol type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbol type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As as shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Generally, number signs consist of a single text character,
// however there may be cases where multiple characters are used
// to designate positive or negative values.
//
// All of these national or cultural number sign styles are
// supported by the type, NumberSignSymbol.
//
type NumberSignSymbol struct {
	leadingNumSignChars          []rune
	trailingNumSignChars         []rune
	leadingNumSignFoundInNumber  bool                  // Leading Number Sign Symbol found in target number
	trailingNumSignFoundInNumber bool                  // Trailing Number Sign Symbol found in target number
	leadingNumSignFoundIndex     int                   // Index of Leading Number Sign Symbol in target number.
	trailingNumSignFoundIndex    int                   // Index of Trailing Number Sign Symbol in target number.
	numSignPosition              NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	numSignType                  NumericSignValueType  // Must be positive or negative
	lock                         *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumberSignSymbol ('incomingNumSignSymbol') to the data fields of
// the current NumberSignSymbol instance ('nSignSymbol').
//
// IMPORTANT
// All of the data fields in current NumberSignSymbol instance
// ('nSignSymbol') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNumSignSymbol      *NumberSignSymbol
//     - A pointer to an instance of NumberSignSymbol. This method
//       will NOT change the values of internal member variables
//       contained in this instance.
//
//       All data values in this NumberSignSymbol instance will be
//       copied to current NumberSignSymbol instance ('nSignSymbol').
//
//       If 'incomingNumSignSymbol' contains invalid data elements,
//       this method will return an error and the current
//       NumberSignSymbol instance ('nSignSymbol') will remain
//       unaltered in its original state.
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
func (nSignSymbol *NumberSignSymbol) CopyIn(
	incomingNumSignSymbol *NumberSignSymbol,
	errorPrefix interface{}) (
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
		"NumberSignSymbol.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = numberSignSymbolMolecule{}.ptr().
		copyIn(nSignSymbol,
			incomingNumSignSymbol,
			ePrefix)

	return err
}

// CopyOut - Creates and returns a deep copy of the current
// NumberSignSymbol instance.
//
// If the current NumberSignSymbol instance is judged to be
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
//  newNumSignSymbol    NumberSignSymbol
//     - If this method completes successfully, a new instance of
//       NumberSignSymbol will be created and returned
//       containing all of the data values copied from the current
//       instance of NumberSignSymbol.
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
func (nSignSymbol *NumberSignSymbol) CopyOut(
	errorPrefix interface{}) (
	newNumSignSymbol NumberSignSymbol,
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
		"NumberSignSymbol.CopyOut()",
		"")

	if err != nil {
		return newNumSignSymbol, err
	}

	newNumSignSymbol,
		err = numberSignSymbolMolecule{}.ptr().
		copyOut(
			nSignSymbol,
			ePrefix)

	return newNumSignSymbol, err
}

// Empty - Reinitializes all internal member variables for the
// current NumberSignSymbol instance to their zero values.
//
// IMPORTANT
// This method will effectively delete all data contained in the
// current NumberSignSymbol instance.
//
func (nSignSymbol *NumberSignSymbol) Empty() {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	_ = numberSignSymbolElectron{}.ptr().
		emptyNumSignSymbol(
			nSignSymbol,
			nil)

	nSignSymbol.lock.Unlock()

	nSignSymbol.lock = nil
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

// GetLeadingNumSignFoundIndex - Returns an integer value
// identifying the index in a number rune array or number string
// where the Leading Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) GetLeadingNumSignFoundIndex() int {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.leadingNumSignFoundIndex
}

// GetLeadingNumSignFoundInNumber - Returns a boolean flag which
// signals whether the Leading Number Sign Symbol has been located
// in a number or number string.
//
func (nSignSymbol *NumberSignSymbol) GetLeadingNumSignFoundInNumber() (
	leadingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.leadingNumSignFoundInNumber
}

// GetNumSignArithmeticVal - Returns the arithmetic value of the
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
func (nSignSymbol *NumberSignSymbol) GetNumSignArithmeticVal() int {

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

// GetNumSignVal - Returns the enumeration type NumericSignValueType
// which identifies this number sign symbol as either positive or
// negative.
//
// Valid return values:
//    NumericSignValueType(0).Negative()
//    NumericSignValueType(0).Positive()
//
func (nSignSymbol *NumberSignSymbol) GetNumSignVal() NumericSignValueType {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSignType
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

// GetTrailingNumSignFoundIndex - Returns an integer value
// identifying the index in a number rune array or number string
// where the Trailing Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) GetTrailingNumSignFoundIndex() int {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.trailingNumSignFoundIndex
}

// GetTrailingNumSignFoundInNumber - Returns a boolean flag which
// signals whether the Trailing Number Sign Symbol has been located
// in a number or number string.
//
func (nSignSymbol *NumberSignSymbol) GetTrailingNumSignFoundInNumber() (
	leadingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.trailingNumSignFoundInNumber
}

// IsLeadingNumSignAtHostIndex - This method will test a host rune
// array to determine if the leading number sign symbol exists
// at the 'hostStartIndex'.
//
// This test will be performed if, and only if, the leading number
// sign symbol has been configured for the current instance of
// NumberSignSymbol.
//
// If the leading number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the last leading number sign symbol encountered
// before the first numeric digit will be tracked and recorded.
//
func (nSignSymbol *NumberSignSymbol) IsLeadingNumSignAtHostIndex(
	hostRunes []rune,
	hostStartIndex int) (
	foundLeadingNumSign bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	if len(nSignSymbol.leadingNumSignChars) == 0 {
		foundLeadingNumSign = false
		return foundLeadingNumSign
	}

	foundLeadingNumSign = strMechPreon{}.ptr().
		isTargetRunesIndex(
			hostRunes,
			hostStartIndex,
			nSignSymbol.leadingNumSignChars)

	if foundLeadingNumSign {
		nSignSymbol.leadingNumSignFoundInNumber = true
		nSignSymbol.leadingNumSignFoundIndex = hostStartIndex
	}

	return foundLeadingNumSign
}

// IsNumSignSymbolFoundInNumber - Returns a boolean flag signaling
// whether the entire number sign symbol, both lead and trailing
// symbols, have been located in a number of number string.
//
// A return value of 'true' signals that the entire number sign
// symbol has been located in the target number or number string.
//
func (nSignSymbol *NumberSignSymbol) IsNumSignSymbolFoundInNumber() bool {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var isNumSignFoundInNumber bool

	isNumSignFoundInNumber,
		_ = numberSignSymbolQuark{}.ptr().
		isNumberSignSymbolFoundInNumber(
			nSignSymbol,
			nil)

	return isNumSignFoundInNumber
}

// IsTrailingNumSignAtHostIndex - This method will test a host rune
// array to determine if the trailing number sign symbol exists
// at the 'hostStartIndex'.
//
// This test will be performed if, and only if, the trailing number
// sign symbol has been configured for the current instance of
// NumberSignSymbol.
//
// If the trailing number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the first trailing number sign symbol encountered
// after the last numeric digit will be tracked and recorded.
//
func (nSignSymbol *NumberSignSymbol) IsTrailingNumSignAtHostIndex(
	hostRunes []rune,
	hostStartIndex int) (
	foundTrailingNumSign bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	foundTrailingNumSign = false

	if len(nSignSymbol.trailingNumSignChars) == 0 {
		return foundTrailingNumSign
	}

	if nSignSymbol.trailingNumSignFoundInNumber {
		return foundTrailingNumSign
	}

	foundTrailingNumSign = strMechPreon{}.ptr().
		isTargetRunesIndex(
			hostRunes,
			hostStartIndex,
			nSignSymbol.leadingNumSignChars)

	if foundTrailingNumSign {
		nSignSymbol.trailingNumSignFoundInNumber = true
		nSignSymbol.trailingNumSignFoundIndex = hostStartIndex
	}

	return foundTrailingNumSign
}

// IsValidInstance - Performs a diagnostic review of the current
// NumberSignSymbol instance to determine whether the current
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
//       current NumberSignSymbol instance is valid, or not. If the
//       current NumberSignSymbol contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
func (nSignSymbol *NumberSignSymbol) IsValidInstance() (
	isValid bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	isValid,
		_ = numberSignSymbolElectron{}.ptr().
		testValidityOfNumSignSymbol(
			nSignSymbol,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumberSignSymbol instance to determine whether the
// current instance is valid in all respects.
//
// If the current NumberSignSymbol instance is judged to be
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
// -----------------------------------------------------------------
//
// Return Values
//
//  err                 error
//     - If the current instance of NumberSignSymbol contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If errors are encountered during processing, the returned
//       error Type (err) will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errorPrefix'. The
//       'errorPrefix' text will be attached to the beginning of
//       the error message.
//
//       If the current instance is valid, this error parameter
//       will be set to nil.
//
func (nSignSymbol *NumberSignSymbol) IsValidInstanceError(
	errorPrefix interface{}) (
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
		"NumberSignSymbol.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = numberSignSymbolElectron{}.ptr().
		testValidityOfNumSignSymbol(
			nSignSymbol,
			ePrefix)

	return err
}

// New - Creates and returns a new instance of NumberSignSymbol.
//
// The NumberSignSymbol type stores the specification for a single
// text or character number sign symbol. This is usually a single
// plus ('+') or minus ('-) for a numeric sign contained in a
// number string.
//
// The NumberSignSymbol type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbol type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As as shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Generally, number signs consist of a single text character (like
// '+' or '-'), however there may be cases where multiple
// characters are used to designate positive or negative values.
//
// All of these national or cultural number sign styles are
// supported by the type, NumberSignSymbol.
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
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
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

	err = numberSignSymbolMechanics{}.ptr().
		setNumberSignSymbol(
			&newNumSignSym,
			leadingNumberSign,
			trailingNumberSign,
			isNegativeValue,
			ePrefix)

	return newNumSignSym, err
}

// SetNumberSignSymbol - Resets the internal data values for the
// current instance of NumberSignSymbol. The new data values will
// be generated from the input parameters listed below.
//
// The NumberSignSymbol type stores the specification for a single
// text or character number sign symbol. This is usually a single
// plus ('+') or minus ('-) for a numeric sign contained in a
// number string.
//
// The NumberSignSymbol type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbol type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As as shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Generally, number signs consist of a single text character (like
// '+' or '-'), however there may be cases where multiple
// characters are used to designate positive or negative values.
//
// All of these national or cultural number sign styles are
// supported by the type, NumberSignSymbol.
//
// IMPORTANT
//
// If this method completes successfully, all data values in the
// current NumberSignSymbol instance will be deleted and
// overwritten.
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
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (nSignSymbol *NumberSignSymbol) SetNumberSignSymbol(
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
	errorPrefix *ePref.ErrPrefixDto) (
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
		"NumberSignSymbol.SetNumberSignSymbol()",
		"")

	if err != nil {
		return err
	}

	err = numberSignSymbolMechanics{}.ptr().
		setNumberSignSymbol(
			nSignSymbol,
			leadingNumberSign,
			trailingNumberSign,
			isNegativeValue,
			ePrefix)

	return err
}

// SetLeadingNumSignFoundIndex - Receives an integer value
// identifying the index in a number rune array or number string
// where the Leading Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) SetLeadingNumSignFoundIndex(
	leadingNumSignFoundIndex int) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.leadingNumSignFoundIndex =
		leadingNumSignFoundIndex
}

// SetLeadingNumSignFoundInNumber - Sets the boolean flag which signals
// whether the Leading Number Sign Symbol has been located in a  number
// or number string.
//
func (nSignSymbol *NumberSignSymbol) SetLeadingNumSignFoundInNumber(
	leadingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.leadingNumSignFoundInNumber =
		leadingNumSignFoundInNumber
}

// SetTrailingNumSignFoundIndex - Receives an integer value
// identifying the index in a number rune array or number string
// where the Trailing Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) SetTrailingNumSignFoundIndex(
	trailingNumSignFoundIndex int) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.trailingNumSignFoundIndex =
		trailingNumSignFoundIndex
}

// SetTrailingNumSignFoundInNumber - Sets the boolean flag which signals
// whether the Trailing Number Sign Symbol has been located in a number
// or number string.
//
func (nSignSymbol *NumberSignSymbol) SetTrailingNumSignFoundInNumber(
	trailingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.leadingNumSignFoundInNumber =
		trailingNumSignFoundInNumber
}
