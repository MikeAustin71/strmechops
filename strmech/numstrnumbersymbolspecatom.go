package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// nStrNumberSymbolSpecAtom - This type provides
// helper methods for NumStrNumberSymbolSpec
type nStrNumberSymbolSpecAtom struct {
	lock *sync.Mutex
}

// emptyLeadingNStrNumSymbol - Receives a pointer to an
// instance of NumStrNumberSymbolSpec and proceeds to
// reset the member variable data for the leading number
// symbol to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading number
// symbol contained in input parameter 'nStrNumSymbolSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec           *NumStrNumberSymbolSpec
//
//		A pointer to an instance of NumStrNumberSymbolSpec.
//		The Leading Number Symbol contained in this
//		instance will be deleted and reset to an empty or zero
//		value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyLeadingNStrNumSymbol(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.leadingNumberSymbols.Empty()

	nStrNumSymbolSpec.leadingNumberFieldSymbolPosition =
		NumFieldSymPos.InsideNumField()

	return
}

// emptyTrailingNStrNumSymbol - Receives a pointer to an instance
// of NumStrNumberSymbolSpec and proceeds to reset the
// member variable data for the trailing number symbol
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing number symbol
// contained in input parameter 'nStrNumSymbolSpec' will be
// deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec           *NumStrNumberSymbolSpec
//	   - A pointer to an instance of NumStrNumberSymbolSpec.
//	     The Trailing Number Symbol contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyTrailingNStrNumSymbol(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.trailingNumberSymbols.Empty()

	nStrNumSymbolSpec.trailingNumberFieldSymbolPosition =
		NumFieldSymPos.InsideNumField()

	return
}
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyCurrNumSignRelPos(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.currencyNumSignRelativePosition =
		CurrNumSignRelPos.None()
}

// equal - Receives a pointer to two instances of
// NumStrNumberSymbolSpec and proceeds to compare their
// member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'. Otherwise,
// this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec1    *NumStrNumberSymbolSpec
//	   - An instance of NumStrNumberSymbolSpec.
//	     Internal member variables from 'nStrNumSymbolSpec1'
//	     will be compared to those of 'nStrNumSymbolSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrNumSymbolSpec2    *NumStrNumberSymbolSpec
//	   - An instance of NumStrNumberSymbolSpec.
//	     Internal member variables from 'nStrNumSymbolSpec2'
//	     will be compared to those of 'nStrNumSymbolSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrNumSymbolSpec1' and
//	     'nStrNumSymbolSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) equal(
	nStrNumSymbolSpec1 *NumStrNumberSymbolSpec,
	nStrNumSymbolSpec2 *NumStrNumberSymbolSpec) bool {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec1 == nil ||
		nStrNumSymbolSpec2 == nil {
		return false
	}

	if !nStrNumSymbolSpec1.leadingNumberSymbols.Equal(
		&nStrNumSymbolSpec2.leadingNumberSymbols) {

		return false
	}

	if nStrNumSymbolSpec1.leadingNumberFieldSymbolPosition !=
		nStrNumSymbolSpec2.leadingNumberFieldSymbolPosition {

		return false

	}

	if !nStrNumSymbolSpec1.trailingNumberSymbols.Equal(
		&nStrNumSymbolSpec2.trailingNumberSymbols) {

		return false
	}

	if nStrNumSymbolSpec1.trailingNumberFieldSymbolPosition !=
		nStrNumSymbolSpec2.trailingNumberFieldSymbolPosition {

		return false
	}

	if nStrNumSymbolSpec1.currencyNumSignRelativePosition !=
		nStrNumSymbolSpec2.currencyNumSignRelativePosition {

		return false
	}

	return true
}

//	testValidityNumStrNumberSymbolSpec
//
//	Performs a diagnostic review of the data values
//	encapsulated in an instance of NumStrNumberSymbolSpec
//	to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolSpec			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of NumStrNumberSymbolSpec.
//		All member variable data values contained in
//		this instance will be reviewed to determine if
//		they are valid.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If any of the internal member data variables
//		contained in 'nStrNumSymbolSpec' are found to be
//		invalid, this method will return a boolean value
//		of 'false'.
//
//		If all internal member data variables contained
//		in 'nStrNumSymbolSpec' are found to be valid,
//		this method returns a boolean value of 'true'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, or
//		if input parameter 'nStrNumSymbolSpec' contains
//		invalid data elements, an error will be returned.
//		The returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) testValidityNumStrNumberSymbolSpec(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrNumberSymbolSpecAtom."+
			"testValidityNumStrNumberSymbolSpec()",
		"")

	if err != nil {
		return isValid, err
	}

	if nStrNumSymbolSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolSpec' is invalid!\n"+
			"'nStrNumSymbolSpec' is a nil pointer.\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrNumSymbolSpec.leadingNumberSymbols.GetRuneArrayLength() == 0 &&
		nStrNumSymbolSpec.trailingNumberSymbols.GetRuneArrayLength() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The NumStrNumberSymbolSpec instance is invalid!\n"+
			"Both 'leadingNumberSymbols' and 'trailingNumberSymbols' are empty.\n"+
			"'nStrNumSymbolSpec.leadingNumberSymbols' has a length of zero.\n"+
			"'nStrNumSymbolSpec.trailingNumberSymbols' has a length of zero.\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrNumSymbolSpec.leadingNumberSymbols.GetRuneArrayLength() > 0 &&
		nStrNumSymbolSpec.leadingNumberFieldSymbolPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: The NumStrNumberSymbolSpec instance is invalid!\n"+
			"The Leading Number Field Symbol Position is invalid!.\n"+
			" 'nStrNumSymbolSpec.leadingNumberFieldSymbolPosition' String Value = %v\n"+
			"'nStrNumSymbolSpec.leadingNumberFieldSymbolPosition' Integer Value = %v\n",
			ePrefix.String(),
			nStrNumSymbolSpec.leadingNumberFieldSymbolPosition.String(),
			nStrNumSymbolSpec.leadingNumberFieldSymbolPosition.XValueInt())

		return isValid, err

	}

	if nStrNumSymbolSpec.trailingNumberSymbols.GetRuneArrayLength() > 0 &&
		nStrNumSymbolSpec.trailingNumberFieldSymbolPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: The NumStrNumberSymbolSpec instance is invalid!\n"+
			"The Trailing Number Field Symbol Position is invalid!.\n"+
			" 'nStrNumSymbolSpec.trailingNumberFieldSymbolPosition' String Value = %v\n"+
			"'nStrNumSymbolSpec.trailingNumberFieldSymbolPosition' Integer Value = %v\n",
			ePrefix.String(),
			nStrNumSymbolSpec.trailingNumberFieldSymbolPosition.String(),
			nStrNumSymbolSpec.trailingNumberFieldSymbolPosition.XValueInt())

		return isValid, err

	}

	isValid = true

	return isValid, err
}
