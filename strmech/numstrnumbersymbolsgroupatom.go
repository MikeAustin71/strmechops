package strmech

import "sync"

// numStrNumberSymbolGroupAtom
//
// Provides helper methods for NumStrNumberSymbolGroup.
type numStrNumberSymbolGroupAtom struct {
	lock *sync.Mutex
}

//	emptyNegativeNumSymbols
//
//	Deletes and resets the Negative Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Negative Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the  Negative Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Negative
//	Number Sign Symbol Specification member variable
//	is identified as:
//
//		nStrNumSymbols.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyNegativeNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.negativeNumberSign.Empty()
}

//	emptyPositiveNumSymbols
//
//	Deletes and resets the Positive Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Positive Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the  Positive Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Positive
//	Number Sign Symbol Specification member variable
//	is identified as:
//
//		nStrNumSymbols.positiveNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyPositiveNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.positiveNumberSign.Empty()
}

//	emptyZeroNumSymbols
//
//	Deletes and resets the Zero Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Zero Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the Zero Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Zero
//	Number Sign Symbol Specification member variable
//	is identified as:
//
//		nStrNumSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyZeroNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.zeroNumberSign.Empty()
}

//	emptyCurrencySymbols
//
//	Deletes and resets the Currency Symbol Specification
//	to its zero or uninitialized state.
//
//	The Currency Symbol Specification object is a member
//	variable in the 'nStrNumSymbols' instance passed as
//	an input parameter.
//
//		nStrNumSymbols.currencySymbol
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the Currency Symbol
//	Specification will be deleted and reset to their
//	zero or uninitialized values. This Currency Symbol
//	Specification member variable is identified as:
//
//		nStrNumSymbols.currencySymbol
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyCurrencySymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.currencySymbol.Empty()
}
