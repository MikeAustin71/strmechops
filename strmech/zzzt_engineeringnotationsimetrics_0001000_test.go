package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestEngineeringNotationSI_Metrics_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestEngineeringNotationSI_Metrics_000100()",
		"")

	lockEngNotationSI.Lock()

	defer lockEngNotationSI.Unlock()

	expectedSymbol := "G"
	index := "10^9"

	var actualSymbol string
	var ok bool

	actualSymbol,
		ok = mEngNotationSISymbols[expectedSymbol]

	if !ok {

		t.Errorf("%v\n"+
			"Error: Lookup of index '%v' failed!\n"+
			"Map = mEngNotationSISymbols[expectedSymbol]\n",
			ePrefix.String(),
			index)

		return
	}

	if actualSymbol != expectedSymbol {

		t.Errorf("%v\n"+
			"Error: Lookup of index '%v' failed!\n"+
			"Expected return of symbol '%v'.\n"+
			"Instead, recevied invalid symbol '%v'\n",
			ePrefix.String(),
			index,
			expectedSymbol,
			actualSymbol)

		return
	}

	return
}

func TestEngineeringNotationSI_Metrics_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestEngineeringNotationSI_Metrics_000200()",
		"")

	lockEngNotationSI.Lock()

	defer lockEngNotationSI.Unlock()

	expectedSymbolName := "G"
	index := "10^−21"

	var actualSymbolName string
	var ok bool

	actualSymbolName,
		ok = mEngNotationSINames[expectedSymbolName]

	if !ok {

		t.Errorf("%v\n"+
			"Error: Lookup of index '%v' failed!\n"+
			"Map = mEngNotationSINames[expectedSymbolName]\n",
			ePrefix.String(),
			index)

		return
	}

	if actualSymbolName != expectedSymbolName {

		t.Errorf("%v\n"+
			"Error: Lookup of index '%v' failed!\n"+
			"Expected return of symbol name '%v'.\n"+
			"Instead, recevied invalid symbol name '%v'\n",
			ePrefix.String(),
			index,
			expectedSymbolName,
			actualSymbolName)

		return
	}

	return
}
