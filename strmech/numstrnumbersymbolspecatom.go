package strmech

import "sync"

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
