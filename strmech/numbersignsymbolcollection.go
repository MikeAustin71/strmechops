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
// currently stored in the the Number Sign Symbol Collection. Upon
// completion, this method will ensure that the internal collection
// of NumberSignSymbolDto objects has a length of zero.
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

// GetCollection - Returns a array of NumberSignSymbolDto objects
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
			ePrefix.XCtx(fmt.Sprintf("Collection Index='%v'",
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

// IsLeadingNumSignAtHostIndex - This method will test a host rune
// array to determine if the leading number sign symbol exists
// at the 'hostStartIndex'. The test will be performed on every
// member NumberSignSymbolDto object in the collection maintained
// by the current NumberSignSymbolCollection instance.
//
// If the leading number sign symbol is located at the
// 'hostStartIndex', tracking information will be recorded.
//
// If multiple leading number sign symbols exist in the host rune
// array, only the last leading number sign symbol encountered
// before the first numeric digit will be tracked and recorded.
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

	lenCol := len(numSignSymCol.numSignSymbols)

	if lenCol == 0 {
		return false
	}

	for i := 0; i < lenCol; i++ {

		foundLeadingNumSign =
			numSignSymCol.numSignSymbols[i].IsLeadingNumSignAtHostIndex(
				hostRunes,
				hostStartIndex)

		if foundLeadingNumSign {
			for j := 0; j < lenCol; j++ {

				if i != j {
					numSignSymCol.numSignSymbols[j].ClearLeadingNumSignTracking()
				}
			}

			return foundLeadingNumSign
		}

	}

	return foundLeadingNumSign
}
