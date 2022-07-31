package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberSignSymbolCollection - This type is used to manage a
// collection of NumberSignSymbolDto objects.
//
type NumberSignSymbolCollection struct {
	numSignSymbols []NumberSignSymbolDto
	lock           *sync.Mutex
}

// AddSymbol - Adds a number sign symbol object
// (NumberSignSymbolDto) to the collection of NumberSignSymbolDto
// objects stored and maintained by the current instance of
// NumberSignSymbolCollection.
//
// A deep copy of the NumberSignSymbolDto object is added to the
// collection.
//
// If the input parameter, 'nSignSymbol' is judged to be invalid,
// this method will return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nSignSymbol         NumberSignSymbolDto
//     - A valid instance of type NumberSignSymbolDto. If this
//       object is judged to be invalid, an error will be returned.
//
//       A deep copy of this NumberSignSymbolDto object will be
//       added to the existing collection of NumberSignSymbolDto
//       objects stored and maintained by the current
//       NumberSignSymbolCollection instance.
//
//       Duplicates are not allowed. If 'nSignSymbol' duplicates
//       the leading and trailing number sign characters of an
//       existing member of the NumberSignSymbolDto collection,
//       this method will take no action and exit without returning
//       an error.
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
func (numSignSymCol *NumberSignSymbolCollection) AddSymbol(
	nSignSymbol NumberSignSymbolDto,
	errorPrefix interface{}) (
	err error) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberSignSymbolCollection.AddSymbol()",
		"")

	if err != nil {
		return err
	}

	var newSymbol NumberSignSymbolDto

	newSymbol,
		err = nSignSymbol.CopyOut(ePrefix)

	if err != nil {
		return err
	}

	// If the new Number Sign Symbol is a duplicate of an
	// existing member of the collection, no action will
	// be taken, this method will exit and no error will
	// be returned.
	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol > 0 {
		for i := 0; i < lenCol; i++ {
			if numSignSymCol.numSignSymbols[i].EqualNumberSignRunes(&newSymbol) {
				return err
			}
		}
	}

	numSignSymCol.numSignSymbols = append(
		numSignSymCol.numSignSymbols, newSymbol)

	return err
}

// AddNewSymbol - Receives the components necessary to create a new
// instance of NumberSignSymbolDto. This method then proceeds to
// create the new NumberSignSymbolDto instance and add it to the
// collection of NumberSignSymbolDto objects stored and maintained
// by the current instance of NumberSignSymbolCollection.
//
// Duplicates are not allowed. If the new instance of
// NumberSignSymbolDto duplicates the leading and trailing number
// sign characters of an existing member of the NumberSignSymbolDto
// collection, this method will take no action and exit without
// returning an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leadingNumberSign   string
//     - A string comprised of the characters which constitute the
//       leading number sign.
//
//       Examples: "-", "+", "("
//
//
//  trailingNumberSign  string
//     - A string comprised of the characters which constitute the
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
func (numSignSymCol *NumberSignSymbolCollection) AddNewSymbol(
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
	errorPrefix interface{}) (
	err error) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberSignSymbolCollection.AddNewSymbol()",
		"")

	if err != nil {
		return err
	}

	var newSymbol NumberSignSymbolDto

	newSymbol,
		err = NumberSignSymbolDto{}.New(
		leadingNumberSign,
		trailingNumberSign,
		isNegativeValue,
		ePrefix)

	if err != nil {
		return err
	}

	// If the new Number Sign Symbol is a duplicate of an
	// existing member of the collection, no action will
	// be taken, this method will exit and no error will
	// be returned.
	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol > 0 {
		for i := 0; i < lenCol; i++ {
			if numSignSymCol.numSignSymbols[i].EqualNumberSignRunes(&newSymbol) {
				return err
			}
		}
	}

	numSignSymCol.numSignSymbols = append(
		numSignSymCol.numSignSymbols, newSymbol)

	return err
}

// EmptyCollection - Deletes all member NumberSignSymbolDto objects
// currently stored in the Number Sign Symbol Collection. Upon
// completion, this method will ensure that the internal collection
// of NumberSignSymbolDto objects has a length of zero.
//
func (numSignSymCol *NumberSignSymbolCollection) EmptyCollection() {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol == 0 {
		return
	}

	for i := 0; i < lenCol; i++ {

		numSignSymCol.numSignSymbols[i].Empty()

	}

	numSignSymCol.numSignSymbols = nil
}

// GetCollection - Returns an array of NumberSignSymbolDto objects
// which comprise the collection maintained by the current
// NumberSignSymbolCollection instance.
//
// The returned array of collection objects represents deep copies
// of the NumberSignSymbolDto objects maintained by the current
// NumberSignSymbolCollection instance.
//
// If the collection has a zero length, meaning there are no
// members in the collection, this method will return 'nil'.
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
func (numSignSymCol *NumberSignSymbolCollection) GetCollection(
	errorPrefix interface{}) (
	[]NumberSignSymbolDto,
	error) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberSignSymbolCollection.GetCollection()",
		"")

	if err != nil {
		return nil, err
	}

	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol == 0 {
		return nil, nil
	}

	col := make([]NumberSignSymbolDto, lenCol)

	for i := 0; i < lenCol; i++ {

		col[i],
			err = numSignSymCol.numSignSymbols[i].CopyOut(
			ePrefix.XCpy(fmt.Sprintf("Collection Index='%v'",
				i)))

		if err != nil {
			return nil, err
		}
	}

	return col, err
}

// GetCollectionLength - Returns the number of NumberSignSymbolDto
// objects currently stored in the Number Sign Symbol Collection
// maintained by the current NumberSignSymbolCollection instance.
func (numSignSymCol *NumberSignSymbolCollection) GetCollectionLength() int {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	return len(numSignSymCol.numSignSymbols)
}

// GetFoundNumberSignSymbol - If one of the number sign symbols in
// this collection was located in a host runes array, this method
// will return a deep copy of that NumberSignSymbolDto object.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  ---- None ----
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  foundNumberSign     bool
//     - If one of the number sign symbols currently residing in
//       this collection of NumberSignSymbolDto objects was located
//       in the host runes array, this return parameter will be set
//       to 'true'.
//
//
//  numSignSymbol       NumberSignSymbolDto
//     - If one of the number sign symbols currently residing in
//       this collection of NumberSignSymbolDto objects was located
//       in the host runes array, this return parameter will be
//       populated with a deep copy of that NumberSignSymbolDto
//       object.
//
//
//  collectionIndex     int
//     - If one of the number sign symbols currently residing in
//       this collection of NumberSignSymbolDto objects was located
//       in the host runes array, this return parameter will be set
//       to the index of the collection array from which the returned
//       number sign symbol (numSignSymbol) was copied.
//
//       If none of the number sign symbols in the NumberSignSymbolDto
//       collection was located in the host rune array (foundNumberSign
//       == false), this parameter is set to -1.
//
func (numSignSymCol *NumberSignSymbolCollection) GetFoundNumberSignSymbol() (
	foundNumberSign bool,
	numSignSymbol NumberSignSymbolDto,
	collectionIndex int) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	collectionIndex = -1

	foundNumberSign,
		collectionIndex =
		numberSignSymbolCollectionAtom{}.ptr().
			getLeadingNSignSymbolFound(
				numSignSymCol.numSignSymbols)

	if foundNumberSign {
		_ =
			numSignSymbol.CopyIn(
				&numSignSymCol.numSignSymbols[collectionIndex],
				nil)

		return foundNumberSign, numSignSymbol, collectionIndex
	}

	foundNumberSign,
		collectionIndex =
		numberSignSymbolCollectionAtom{}.ptr().
			getTrailingNSignSymbolFound(
				numSignSymCol.numSignSymbols)

	if foundNumberSign {
		_ =
			numSignSymbol.CopyIn(
				&numSignSymCol.numSignSymbols[collectionIndex],
				nil)

		return foundNumberSign, numSignSymbol, collectionIndex
	}

	return foundNumberSign, numSignSymbol, collectionIndex
}

// IsCollectionEmpty - Returns 'true' if the collection of
// NumberSignSymbolDto objects is empty and contains zero members.
//
func (numSignSymCol *NumberSignSymbolCollection) IsCollectionEmpty() bool {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	if len(numSignSymCol.numSignSymbols) == 0 {
		return true
	}

	return false
}

// IsLeadingNumSignAtHostIndex - This method will test a host rune
// array to determine if the leading number sign symbol exists
// at the 'hostStartIndex'. The test will be performed on every
// member NumberSignSymbolDto object in the collection maintained
// by the current NumberSignSymbolCollection instance.
//
// If the leading number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded. To be
// clear, the leading number sign symbol must exist in the host
// runes array at the host starting index before a value of 'true'
// is returned.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the last leading number sign symbol encountered
// before the first numeric digit will be tracked and recorded.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hostRunes                  []rune
//     - An array of runes. This rune array will be searched to
//       determine if the leading number sign symbol is present in
//       the array beginning at the 'hostStartIndex'.
//
//       If 'hostRunes' is a zero length array, this method will
//       return 'false'.
//
//
//  hostStartIndex             int
//     - The starting index within the host runes array where
//       the search operation will commence. If 'hostStartIndex' is
//       less than zero, it will be automatically set to zero.
//
//       If the 'hostStartIndex' is greater than or equal to the
//       length of 'hostRunes', this method will return 'false'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  foundLeadingNumSign        bool
//     - A boolean flag signaling whether the leading number sign
//       symbol was located in the host runes array beginning at
//       the index specified by input parameter 'hostStartIndex'.
//
//       If the target runes array is found at the staring index in
//       the host runes array, this method will return 'true'.
//
func (numSignSymCol *NumberSignSymbolCollection) IsLeadingNumSignAtHostIndex(
	hostRunes []rune,
	hostStartIndex int) (
	foundLeadingNumSign bool) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	foundLeadingNumSign = false

	if len(hostRunes) == 0 ||
		hostStartIndex < 0 {

		return foundLeadingNumSign
	}

	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol == 0 {
		return foundLeadingNumSign
	}

	for i := 0; i < lenCol; i++ {

		foundLeadingNumSign =
			numSignSymCol.numSignSymbols[i].
				IsLeadingNumSignAtHostIndex(
					hostRunes,
					hostStartIndex)

		if foundLeadingNumSign {
			return foundLeadingNumSign
		}
	}

	return foundLeadingNumSign
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
// 'hostStartIndex', tracking information will be recorded. To be
// clear, the trailing number sign symbol must exist in the host
// runes array beginning at the host starting index before this
// method will return a value of 'true'.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the first trailing number sign symbol encountered
// after the last numeric digit will be tracked and recorded.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hostRunes                  []rune
//     - An array of runes. This rune array will be searched to
//       determine if the trailing number sign symbol is present in
//       the array beginning at the 'hostStartIndex'.
//
//       If 'hostRunes' is a zero length array, this method will
//       return 'false'.
//
//
//  hostStartIndex             int
//     - The starting index within the host runes array where
//       the search operation will commence. If 'hostStartIndex' is
//       less than zero, it will be automatically set to zero.
//
//       If the 'hostStartIndex' is greater than or equal to the
//       length of 'hostRunes', this method will return 'false'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  foundTrailingNumSign       bool
//     - A boolean flag signaling whether the trailing number sign
//       symbol was located in the host runes array beginning at
//       the index specified by input parameter 'hostStartIndex'.
//
//       If the target runes array is found at the staring index in
//       the host runes array, this method will return 'true'.
//
func (numSignSymCol *NumberSignSymbolCollection) IsTrailingNumSignAtHostIndex(
	hostRunes []rune,
	hostStartIndex int) (
	foundTrailingNumSign bool) {

	if numSignSymCol.lock == nil {
		numSignSymCol.lock = new(sync.Mutex)
	}

	numSignSymCol.lock.Lock()

	defer numSignSymCol.lock.Unlock()

	foundTrailingNumSign = false

	if len(hostRunes) == 0 {

		return foundTrailingNumSign
	}

	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol == 0 {
		return foundTrailingNumSign
	}

	isTrailingNumberSignFound,
		_ :=
		numberSignSymbolCollectionAtom{}.ptr().
			isTrailingNSignSymbolFound(numSignSymCol.numSignSymbols)

	if isTrailingNumberSignFound {
		// Trailing Number Sign Already Found
		// return false
		return foundTrailingNumSign
	}

	for i := 0; i < lenCol; i++ {

		foundTrailingNumSign =
			numSignSymCol.numSignSymbols[i].
				IsTrailingNumSignAtHostIndex(
					hostRunes,
					hostStartIndex)

		if foundTrailingNumSign {
			return foundTrailingNumSign
		}
	}

	return foundTrailingNumSign
}
