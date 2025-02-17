package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberSignSymbolDto - Stores the specification for a single text or
// character number sign symbol. This is usually a single plus
// ('+') or minus ('-) for a numeric sign contained in a number
// string.
//
// The NumberSignSymbolDto type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbolDto type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU'
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Generally, number signs consist of a single text character,
// however there may be cases where multiple characters are used
// to designate positive or negative values.
//
// All of these national or cultural number sign styles are
// supported by the type, NumberSignSymbolDto.
//
// ------------------------------------------------------------------------
//
// IMPORTANT
// NumberSignSymbolDto objects should only be created using one of
// the following two methods:
//
//	NumberSignSymbolDto.New()
//	NumberSignSymbolDto.SetNumberSignSymbol()
type NumberSignSymbolDto struct {
	leadingNumSignChars          []rune
	trailingNumSignChars         []rune
	leadingNumSignFoundInNumber  bool                  // Leading Number Sign Symbol found in target number
	trailingNumSignFoundInNumber bool                  // Trailing Number Sign Symbol found in target number
	leadingNumSignFoundIndex     int                   // Index of Leading Number Sign Symbol in target number.
	trailingNumSignFoundIndex    int                   // Index of Trailing Number Sign Symbol in target number.
	numSignPosition              NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	numSignValueType             NumericSignValueType  // Must be positive or negative
	numSymbolClass               NumericSymbolClass    // Always set to NumericSymbolClass(0).NumberSign()
	numSymbolDisplayMode         NumSignSymbolDisplayMode
	lock                         *sync.Mutex
}

// ClearLeadingNumSignTracking - This method will clear or delete
// all the tracking data collected for the Leading Number Sign
// Symbol encapsulated by the current NumberSignSymbolDto instance.
//
// Specifically, this means that member variables
// 'leadingNumSignFoundInNumber' will be set to 'false' and
// 'leadingNumSignFoundIndex' will be set to zero.
func (nSignSymbol *NumberSignSymbolDto) ClearLeadingNumSignTracking() {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.leadingNumSignFoundInNumber = false

	nSignSymbol.leadingNumSignFoundIndex = 0

	return
}

// ClearTrailingNumSignTracking - This method will clear or delete
// all the tracking data collected for the Trailing Number Sign
// Symbol encapsulated by the current NumberSignSymbolDto instance.
//
// Specifically, this means that member variables
// 'trailingNumSignFoundInNumber' will be set to 'false' and
// 'trailingNumSignFoundIndex' will be set to zero.
func (nSignSymbol *NumberSignSymbolDto) ClearTrailingNumSignTracking() {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.trailingNumSignFoundInNumber = false

	nSignSymbol.trailingNumSignFoundIndex = 0

	return
}

// CopyIn - Copies the data fields from an incoming instance of
// NumberSignSymbolDto ('incomingNumSignSymbol') to the data fields of
// the current NumberSignSymbolDto instance ('nSignSymbol').
//
// IMPORTANT
// All the data fields in current NumberSignSymbolDto instance
// ('nSignSymbol') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumSignSymbol      *NumberSignSymbolDto
//	   - A pointer to an instance of NumberSignSymbolDto. This method
//	     will NOT change the values of internal member variables
//	     contained in this instance.
//
//	     All data values in this NumberSignSymbolDto instance will be
//	     copied to current NumberSignSymbolDto instance ('nSignSymbol').
//
//	     If 'incomingNumSignSymbol' contains invalid data elements,
//	     this method will return an error and the current
//	     NumberSignSymbolDto instance ('nSignSymbol') will remain
//	     unaltered in its original state.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
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
func (nSignSymbol *NumberSignSymbolDto) CopyIn(
	incomingNumSignSymbol *NumberSignSymbolDto,
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
		"NumberSignSymbolDto.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = numberSignSymbolDtoMolecule{}.ptr().
		copyIn(nSignSymbol,
			incomingNumSignSymbol,
			ePrefix)

	return err
}

// CopyOut - Creates and returns a deep copy of the current
// NumberSignSymbolDto instance.
//
// If the current NumberSignSymbolDto instance is judged to be
// invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newNumSignSymbol    NumberSignSymbolDto
//	   - If this method completes successfully, a new instance of
//	     NumberSignSymbolDto will be created and returned
//	     containing all the data values copied from the current
//	     instance of NumberSignSymbolDto.
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
func (nSignSymbol *NumberSignSymbolDto) CopyOut(
	errorPrefix interface{}) (
	newNumSignSymbol NumberSignSymbolDto,
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
		"NumberSignSymbolDto.CopyOut()",
		"")

	if err != nil {
		return newNumSignSymbol, err
	}

	newNumSignSymbol,
		err = numberSignSymbolDtoMolecule{}.ptr().
		copyOut(
			nSignSymbol,
			ePrefix)

	return newNumSignSymbol, err
}

// Empty - Reinitialize all internal member variables for the
// current NumberSignSymbolDto instance to their zero values.
//
// IMPORTANT
// This method will effectively delete all data contained in the
// current NumberSignSymbolDto instance.
func (nSignSymbol *NumberSignSymbolDto) Empty() {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	_ = numberSignSymbolDtoElectron{}.ptr().
		emptyNumSignSymbol(
			nSignSymbol,
			nil)

	nSignSymbol.lock.Unlock()

	nSignSymbol.lock = nil
}

// Equal - Receives a pointer to an incoming instance of
// NumberSignSymbolDto and proceeds to compare the data values with
// those of the current NumberSignSymbolDto instance. If both sets
// of data values are equivalent, this method returns 'true'.
//
// If the incoming NumberSignSymbolDto data values are NOT
// equivalent to those of the current NumberSignSymbolDto instance,
// this method returns 'false'.
func (nSignSymbol *NumberSignSymbolDto) Equal(
	incomingNSignSymbol *NumberSignSymbolDto) bool {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return numberSignSymbolDtoElectron{}.ptr().
		equalNumSignSymbolObjects(
			nSignSymbol,
			incomingNSignSymbol)
}

// EqualNumberSignRunes - Receives a pointer to an incoming
// instance of NumberSignSymbolDto and proceeds to compare the
// leading and trailing number sign runes with those of the current
// NumberSignSymbolDto instance. If both sets of number sign runes
// are equivalent, this method returns 'true'.
//
// If the two number sign rune arrays differ in any respect, this
// method will return false.
//
// This method differs from NumberSignSymbolDto.Equal() in that
// only the number sign rune arrays are compared. No other internal
// member variables are compared.
func (nSignSymbol *NumberSignSymbolDto) EqualNumberSignRunes(
	incomingNSignSymbol *NumberSignSymbolDto) bool {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return numberSignSymbolDtoElectron{}.ptr().
		equalNumSignRuneArrays(
			nSignSymbol,
			incomingNSignSymbol)

}

// GetNumSignSymDisplayMode - Returns the current value of the
// Number Sign Symbol Display Mode for the current
// NumberSignSymbolDto instance.
//
// Number Sign Symbols displayed in a number string can either be
// 'Explicit' or 'Implicit'.
//
// Explicit Number Sign Symbols are created and explicitly
// displayed within the number string using number sign
// characters (plus or minus). Explicit Number Sign Symbols
// are ALWAYS visible and displayed within number strings.
//
//	Example: '+123', '-123'.
//
// Implicit Number Sign Symbols are NOT visible or displayed in a
// number string. Instead, the Number Sign Symbol is assumed or
// implied. Implicit Number Sign Symbols are NEVER visible or
// displayed within number strings.
//
//	Example: '123' is assumed to be a positive value. The
//	         plus sign is implied or 'implicit'.
//
// This method will return an enumeration type which is set to one
// of the following three values:
//
//	NumSignSymbolDisplayMode(0).None()
//	NumSignSymbolDisplayMode(0).Explicit()
//	NumSignSymbolDisplayMode(0).Implicit()
func (nSignSymbol *NumberSignSymbolDto) GetNumSignSymDisplayMode() NumSignSymbolDisplayMode {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSymbolDisplayMode
}

// GetLeadingNumSignChars - Returns a deep copy of the leading
// number sign characters contained in this instance of
// NumberSignSymbolDto.
func (nSignSymbol *NumberSignSymbolDto) GetLeadingNumSignChars() []rune {

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
// current instance of NumberSignSymbolDto.
func (nSignSymbol *NumberSignSymbolDto) GetLeadingNumSignFoundIndex() int {

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
func (nSignSymbol *NumberSignSymbolDto) GetLeadingNumSignFoundInNumber() (
	leadingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.leadingNumSignFoundInNumber
}

// GetNumericSymbolClass - Returns the NumericSymbolClass
// enumeration value associated with instances of
// NumberSignSymbolDto.
//
// All NumberSignSymbolDto objects have NumericSymbolClass values
// set equal to NumericSymbolClass(0).NumberSign().
func (nSignSymbol *NumberSignSymbolDto) GetNumericSymbolClass() NumericSymbolClass {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.numSymbolClass = NumericSymbolClass(0).NumberSign()

	return nSignSymbol.numSymbolClass
}

// GetNumSignArithmeticVal - Returns the arithmetic value of the
// number sign associated with this NumberSignSymbolDto instance.
//
// Valid return values are:
//
//	Negative == -1
//	Zero     ==  0
//	Positive ==  1
//
// If the Number Sign value is invalid, this method will return a
// value of -99
func (nSignSymbol *NumberSignSymbolDto) GetNumSignArithmeticVal() int {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSignValueType.XArithmeticValue()
}

// GetNumSignSymPosition - Returns an enumeration value specifying
// the position of the Number Sign Symbol within a number string.
//
// Number Sign Symbols may be positioned before the numeric value,
// after the numeric value or before and after the numeric value.
//
// This method returns a type of NumSignSymbolPosition. Possible
// return values are listed as follows:
//
//	NumSignSymbolPosition(0).Before()
//	NumSignSymbolPosition(0).After()
//	NumSignSymbolPosition(0).BeforeAndAfter()
func (nSignSymbol *NumberSignSymbolDto) GetNumSignSymPosition() NumSignSymbolPosition {

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
//
//	NumericSignValueType(0).Negative()
//	NumericSignValueType(0).Positive()
func (nSignSymbol *NumberSignSymbolDto) GetNumSignVal() NumericSignValueType {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.numSignValueType
}

// GetTrailingNumSignChars - Returns a deep copy of the trailing
// number sign characters contained in this instance of
// NumberSignSymbolDto.
func (nSignSymbol *NumberSignSymbolDto) GetTrailingNumSignChars() []rune {

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
// current instance of NumberSignSymbolDto.
func (nSignSymbol *NumberSignSymbolDto) GetTrailingNumSignFoundIndex() int {

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
func (nSignSymbol *NumberSignSymbolDto) GetTrailingNumSignFoundInNumber() (
	trailingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	return nSignSymbol.trailingNumSignFoundInNumber
}

// IsEmpty - Returns 'true' if the current NumberSignSymbolDto
// instance is empty or uninitialized.
func (nSignSymbol *NumberSignSymbolDto) IsEmpty() bool {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	if len(nSignSymbol.leadingNumSignChars) == 0 &&
		len(nSignSymbol.trailingNumSignChars) == 0 &&
		nSignSymbol.numSymbolDisplayMode !=
			NumSignSymbolDisplayMode(0).Implicit() {
		return true
	}

	return false
}

// IsLeadingNumSignAtHostIndex - This method will test a host rune
// array to determine if the leading number sign symbol exists
// at the 'hostStartIndex'.
//
// This test will be performed if, and only if, the leading number
// sign symbol has been configured for the current instance of
// NumberSignSymbolDto.
//
// If the leading number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the last leading number sign symbol encountered
// before the first numeric digit will be tracked and recorded.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	hostRunes                  []rune
//	   - An array of runes. This rune array will be searched to
//	     determine if the leading number sign symbol is present in
//	     the array beginning at the 'hostStartIndex'.
//
//	     If 'hostRunes' is a zero length array, this method will
//	     return 'false'.
//
//
//	hostStartIndex             int
//	   - The starting index within the host runes array where
//	     the search operation will commence. If 'hostStartIndex' is
//	     less than zero, it will be automatically set to zero.
//
//	     If the 'hostStartIndex' is greater than or equal to the
//	     length of 'hostRunes', this method will return 'false'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	foundLeadingNumSign        bool
//	   - A boolean flag signaling whether the leading number sign
//	     symbol was located in the host runes array beginning at
//	     the index specified by input parameter 'hostStartIndex'.
//
//	     If the target runes array is found at the staring index in
//	     the host runes array, this method will return 'true'.
func (nSignSymbol *NumberSignSymbolDto) IsLeadingNumSignAtHostIndex(
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
// whether the entire number sign symbol has been located in a
// number string.
//
// A return value of 'true' signals that the entire number sign
// symbol has been located in the target number or number string.
// Depending on the configuration of the NumberSignSymbolDto, this
// could indicate one of the following outcomes:
//  1. The leading number sign was found.
//  2. The trailing number sign was found.
//  3. This symbol has both a leading and trailing component
//     and both components were found.
func (nSignSymbol *NumberSignSymbolDto) IsNumSignSymbolFoundInNumber() bool {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var isNumSignFoundInNumber bool

	isNumSignFoundInNumber,
		_ = numberSignSymbolDtoQuark{}.ptr().
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
// NumberSignSymbolDto.
//
// If the trailing number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded.
//
// If multiple trailing number sign symbols exist in the host rune
// array, only the first trailing number sign symbol encountered
// after the last numeric digit will be tracked and recorded.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	hostRunes                  []rune
//	   - An array of runes. This rune array will be searched to
//	     determine if the trailing number sign symbol is present in
//	     the array beginning at the 'hostStartIndex'.
//
//	     If 'hostRunes' is a zero length array, this method will
//	     return 'false'.
//
//
//	hostStartIndex             int
//	   - The starting index within the host runes array where
//	     the search operation will commence. If 'hostStartIndex' is
//	     less than zero, it will be automatically set to zero.
//
//	     If the 'hostStartIndex' is greater than or equal to the
//	     length of 'hostRunes', this method will return 'false'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	foundTrailingNumSign       bool
//	   - A boolean flag signaling whether the trailing number sign
//	     symbol was located in the host runes array beginning at
//	     the index specified by input parameter 'hostStartIndex'.
//
//	     If the target runes array is found at the staring index in
//	     the host runes array, this method will return 'true'.
func (nSignSymbol *NumberSignSymbolDto) IsTrailingNumSignAtHostIndex(
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

	if nSignSymbol.numSignPosition == NumSignSymPos.BeforeAndAfter() &&
		!nSignSymbol.leadingNumSignFoundInNumber {
		return foundTrailingNumSign
	}

	foundTrailingNumSign = strMechPreon{}.ptr().
		isTargetRunesIndex(
			hostRunes,
			hostStartIndex,
			nSignSymbol.trailingNumSignChars)

	if foundTrailingNumSign {
		nSignSymbol.trailingNumSignFoundInNumber = true
		nSignSymbol.trailingNumSignFoundIndex = hostStartIndex
	}

	return foundTrailingNumSign
}

// IsValidInstance - Performs a diagnostic review of the current
// NumberSignSymbolDto instance to determine whether the current
// instance is valid in all respects.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//	isValid             bool
//	   - This returned boolean value will signal whether the
//	     current NumberSignSymbolDto instance is valid, or not. If the
//	     current NumberSignSymbolDto contains valid data, this method
//	     returns 'true'. If the data is invalid, this method will
//	     return 'false'.
func (nSignSymbol *NumberSignSymbolDto) IsValidInstance() (
	isValid bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	isValid,
		_ = numberSignSymbolDtoElectron{}.ptr().
		testValidityOfNumSignSymbol(
			nSignSymbol,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the
// current NumberSignSymbolDto instance to determine whether the
// current instance is valid in all respects.
//
// If the current NumberSignSymbolDto instance is judged to be
// invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	err                 error
//	   - If the current instance of NumberSignSymbolDto contains
//	     invalid data, a detailed error message will be returned
//	     identifying the invalid data item.
//
//	     If errors are encountered during processing, the returned
//	     error Type (err) will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errorPrefix'. The
//	     'errorPrefix' text will be attached to the beginning of
//	     the error message.
//
//	     If the current instance is valid, this error parameter
//	     will be set to nil.
func (nSignSymbol *NumberSignSymbolDto) IsValidInstanceError(
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
		"NumberSignSymbolDto.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = numberSignSymbolDtoElectron{}.ptr().
		testValidityOfNumSignSymbol(
			nSignSymbol,
			ePrefix)

	return err
}

// New - Creates and returns a new instance of NumberSignSymbolDto.
//
// The NumberSignSymbolDto type stores the specification for a single
// text or character number sign symbol. This is usually a single
// plus ('+') or minus ('-) for a numeric sign contained in a
// number string.
//
// The NumberSignSymbolDto type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbolDto type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU'
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Number Sign Symbols displayed in a number string can either be
// 'Explicit' or 'Implicit'.
//
// Explicit Number Sign Symbols are created and explicitly
// displayed within the number string using number sign
// characters (plus or minus). Explicit Number Sign Symbols
// are ALWAYS visible and displayed within number strings.
//
//	Example: '+123', '-123'.
//
// This method will generate an 'Explicit' Number Sign Symbol.
//
// Implicit Number Sign Symbols are NOT visible or displayed in a
// number string. Instead, the Number Sign Symbol is assumed or
// implied. Implicit Number Sign Symbols are NEVER visible or
// displayed within number strings.
//
//	Example: '123' is assumed to be a positive value. The
//	         plus sign is implied or 'implicit'.
//
// To generate an 'Implicit' Number Sign Symbol, use one of the
// following methods:
//
//	NumberSignSymbolDto.NewDefaultPositive()
//	NewDefaultPositive.SetDefaultPositive()
//
// Generally, number signs consist of a single text character or
// symbol (like '+' or '-'). However, depending on national,
// cultural or language factors, there may be cases where multiple
// characters are used to designate positive or negative numeric
// values.
//
// All national or cultural number sign styles are supported by the
// NumberSignSymbolDto type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	leadingNumberSign   string
//	   - A string comprised of the characters which constitute this
//	     leading number sign.
//
//	     Examples: "-", "+", "("
//
//
//	trailingNumberSign  string
//	   - A string comprised of the characters which constitute this
//	     trailing number sign.
//
//	     Examples: "-", "+", ")"
//
//
//	isNegativeValue     bool
//	   - Number sign symbols will specify either a positive or
//	     negative numeric value. If this parameter is set to
//	     'true', the number sign will be treated as identifying a
//	     negative numeric value. If this parameter is set to
//	     'false', the number sign will be interpreted as
//	     identifying a positive numeric value.
//
//
//	errorPrefix         interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newNumSignSym       NumberSignSymbolDto
//	   - If this method completes successfully, a new, populated
//	     instance of NumberSignSymbolDto will be returned.
//
//
//	err                 error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (nSignSymbol NumberSignSymbolDto) New(
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
	errorPrefix interface{}) (
	newNumSignSym NumberSignSymbolDto,
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
		"NumberSignSymbolDto.New()",
		"")

	if err != nil {
		return newNumSignSym, err
	}

	err = numberSignSymbolDtoMechanics{}.ptr().
		setNumberSignSymbol(
			&newNumSignSym,
			leadingNumberSign,
			trailingNumberSign,
			isNegativeValue,
			NumSignSymbolDisplayMode(0).Explicit(),
			ePrefix)

	return newNumSignSym, err
}

// NewDefaultPositive - Creates and returns a new instance of
// NumberSignSymbolDto configured as an implicit, positive number
// sign symbol.
//
// Number Sign Symbols displayed in a number string can either be
// 'Explicit' or 'Implicit'.
//
// Explicit Number Sign Symbols are created and explicitly
// displayed within the number string using number sign
// characters (plus or minus). Explicit Number Sign Symbols
// are ALWAYS visible and displayed within number strings.
//
//	Example: '+123', '-123'.
//
// To generate an 'Explicit' Number Sign Symbol, use one of the
// following methods:
//
//	NumberSignSymbolDto.New()
//	NumberSignSymbolDto.SetNumberSignSymbol()
//
// Implicit Number Sign Symbols are NOT visible or displayed in a
// number string. Instead, the Number Sign Symbol is assumed or
// implied. Implicit Number Sign Symbols are NEVER visible or
// displayed within number strings.
//
//	Example: '123' is assumed to be a positive value. The
//	         plus sign is implied or 'implicit'.
//
// This method will return a default positive Number Sign Symbol
// with an 'implicit' display mode.
func (nSignSymbol NumberSignSymbolDto) NewDefaultPositive() NumberSignSymbolDto {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	newNumSignSym := NumberSignSymbolDto{}

	_ = numberSignSymbolDtoMechanics{}.ptr().
		setNumberSignSymbol(
			&newNumSignSym,
			"+",
			"",
			false,
			NumSignSymbolDisplayMode(0).Implicit(),
			nil)

	return newNumSignSym
}

// SetDefaultPositive - Resets the internal data values for the
// current instance of NumberSignSymbolDto. The new data values
// will be automatically generated to configure the current
// NumberSignSymbolDto instance as a default, positive Number Sign
// Symbol.
//
// Number Sign Symbols displayed in a number string can either be
// 'Explicit' or 'Implicit'.
//
// Explicit Number Sign Symbols are created and explicitly
// displayed within the number string using number sign
// characters (plus or minus). Explicit Number Sign Symbols
// are ALWAYS visible and displayed within number strings.
//
//	Example: '+123', '-123'.
//
// To generate an 'Explicit' Number Sign Symbol, use one of the
// following methods:
//
//	NumberSignSymbolDto.New()
//	NumberSignSymbolDto.SetNumberSignSymbol()
//
// Implicit Number Sign Symbols are NOT visible or displayed in a
// number string. Instead, the Number Sign Symbol is assumed or
// implied. Implicit Number Sign Symbols are NEVER visible or
// displayed within number strings.
//
//	Example: '123' is assumed to be a positive value. The
//	         plus sign is implied or 'implicit'.
//
// This method will return a default positive Number Sign Symbol
// with an 'Implicit' display mode.
//
// # IMPORTANT
//
// If this method completes successfully, all data values in the
// current NumberSignSymbolDto instance will be deleted and
// overwritten.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	--- NONE ---
func (nSignSymbol *NumberSignSymbolDto) SetDefaultPositive() {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	_ = numberSignSymbolDtoMechanics{}.ptr().
		setNumberSignSymbol(
			nSignSymbol,
			"+",
			"",
			false,
			NumSignSymbolDisplayMode(0).Implicit(),
			nil)

	return
}

// SetNumberSignSymbol - Resets the internal data values for the
// current instance of NumberSignSymbolDto. The new data values will
// be generated from the input parameters listed below.
//
// The NumberSignSymbolDto type stores the specification for a single
// text or character number sign symbol. This is usually a single
// plus ('+') or minus ('-) for a numeric sign contained in a
// number string.
//
// The NumberSignSymbolDto type is designed to support number sign
// symbols used by all nationalities and cultures. As such the
// NumberSignSymbolDto type can process number signs comprised of
// multiple characters. For example, in the USA, negative numeric
// values are often identified by leading and trailing parentheses
// "(55)".
//
// Number Sign Symbols are usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU'
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As shown
// above, the USA uses opening and closing parentheses to designate
// a negative number "(55)".
//
// Number Sign Symbols displayed in a number string can either be
// 'Explicit' or 'Implicit'.
//
// Explicit Number Sign Symbols are created and explicitly
// displayed within the number string using number sign
// characters (plus or minus). Explicit Number Sign Symbols
// are ALWAYS visible and displayed within number strings.
//
//	Example: '+123', '-123'.
//
// This method will generate an 'Explicit' Number Sign Symbol.
//
// Implicit Number Sign Symbols are NOT visible or displayed in a
// number string. Instead, the Number Sign Symbol is assumed or
// implied. Implicit Number Sign Symbols are NEVER visible or
// displayed within number strings.
//
//	Example: '123' is assumed to be a positive value. The
//	         plus sign is implied or 'implicit'.
//
// To generate an 'Implicit' Number Sign Symbol, use one of the
// following methods:
//
//	NumberSignSymbolDto.NewDefaultPositive()
//	NewDefaultPositive.SetDefaultPositive()
//
// Generally, number signs consist of a single text character or
// symbol (like '+' or '-'). However, depending on national,
// cultural or language factors, there may be cases where multiple
// characters are used to designate positive or negative numeric
// values.
//
// All national or cultural number sign styles are supported by the
// NumberSignSymbolDto type.
//
// # IMPORTANT
//
// If this method completes successfully, all data values in the
// current NumberSignSymbolDto instance will be deleted and
// overwritten.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	leadingNumberSign   string
//	   - A string comprised of the characters which constitute this
//	     leading number sign.
//
//	     Examples: "-", "+", "("
//
//
//	trailingNumberSign  string
//	   - A string comprised of the characters which constitute this
//	     trailing number sign.
//
//	     Examples: "-", "+", ")"
//
//
//	isNegativeValue     bool
//	   - Number sign symbols will specify either a positive or
//	     negative numeric value. If this parameter is set to
//	     'true', the number sign will be treated as identifying a
//	     negative numeric value. If this parameter is set to
//	     'false', the number sign will be interpreted as
//	     identifying a positive numeric value.
//
//
//	errorPrefix         interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newNumSignSym       NumberSignSymbolDto
//	   - If this method completes successfully, a new, populated
//	     instance of NumberSignSymbolDto will be returned.
//
//
//	err                 error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (nSignSymbol *NumberSignSymbolDto) SetNumberSignSymbol(
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
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
		"NumberSignSymbolDto.SetNumberSignSymbol()",
		"")

	if err != nil {
		return err
	}

	err = numberSignSymbolDtoMechanics{}.ptr().
		setNumberSignSymbol(
			nSignSymbol,
			leadingNumberSign,
			trailingNumberSign,
			isNegativeValue,
			NumSignSymbolDisplayMode(0).Explicit(),
			ePrefix)

	return err
}

// SetLeadingNumSignFoundIndex - Receives an integer value
// identifying the index in a number rune array or number string
// where the Leading Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbolDto.
//
// The Leading Number Sign Chars rune array must be initialized
// prior to setting the Leading Number Found Index. If the length
// of Leading Number Sign Chars rune array is zero, this method
// will take no action and exit.
func (nSignSymbol *NumberSignSymbolDto) SetLeadingNumSignFoundIndex(
	leadingNumSignFoundIndex int) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	if len(nSignSymbol.leadingNumSignChars) == 0 {
		return
	}

	nSignSymbol.leadingNumSignFoundIndex =
		leadingNumSignFoundIndex

	if leadingNumSignFoundIndex < 0 {
		nSignSymbol.leadingNumSignFoundInNumber = false
	} else {
		nSignSymbol.leadingNumSignFoundInNumber = true
	}
}

// SetLeadingNumSignFoundInNumber - Sets the boolean flag which
// signals whether the Leading Number Sign Symbol has been located
// in a  number or number string.
//
// The Leading Number Sign Chars rune array must be initialized
// prior to setting the Leading Number Sign Found Flag. If the
// length of the Leading Number Sign Chars rune array is zero, this
// method will take no action and exit.
func (nSignSymbol *NumberSignSymbolDto) SetLeadingNumSignFoundInNumber(
	leadingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	if len(nSignSymbol.leadingNumSignChars) == 0 {
		return
	}

	nSignSymbol.leadingNumSignFoundInNumber =
		leadingNumSignFoundInNumber

	if leadingNumSignFoundInNumber == false {
		nSignSymbol.leadingNumSignFoundIndex = 0
	}

}

// SetTrailingNumSignFoundIndex - Receives an integer value
// identifying the index in a number rune array or number string
// where the Trailing Number Sign Symbol is located.
//
// This index is stored as an internal member variable in the
// current instance of NumberSignSymbolDto.
//
// The Trailing Number Sign Chars rune array must be initialized
// prior to setting the Trailing Number Sign Found Index. If the
// length of the Trailing Number Sign Chars rune array is zero,
// this method will take no action and exit.
func (nSignSymbol *NumberSignSymbolDto) SetTrailingNumSignFoundIndex(
	trailingNumSignFoundIndex int) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.trailingNumSignFoundIndex =
		trailingNumSignFoundIndex

	if trailingNumSignFoundIndex < 0 {
		nSignSymbol.trailingNumSignFoundInNumber = false
	} else {
		nSignSymbol.trailingNumSignFoundInNumber = true
	}

}

// SetTrailingNumSignFoundInNumber - Sets the boolean flag which
// signals whether the Trailing Number Sign Symbol has been located
// in a number or number string.
//
// The Trailing Number Sign Chars rune array must be initialized
// prior to setting the Trailing Number Sign Found Flag. If the
// length of the Trailing Number Sign Chars rune array is zero,
// this method will take no action and exit.
func (nSignSymbol *NumberSignSymbolDto) SetTrailingNumSignFoundInNumber(
	trailingNumSignFoundInNumber bool) {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	nSignSymbol.trailingNumSignFoundInNumber =
		trailingNumSignFoundInNumber

	if trailingNumSignFoundInNumber == false {
		nSignSymbol.trailingNumSignFoundIndex = 0
	}

}
