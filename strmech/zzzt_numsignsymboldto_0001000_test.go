package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumberSignSymbolDto_CopyIn_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyIn_000100()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	err = nSignSymTwo.CopyIn(
		&nSignSymOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyIn_000200()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	err = nSignSymTwo.CopyIn(
		&nSignSymOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyIn_000300()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	err = nSignSymTwo.CopyIn(
		&nSignSymOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyIn_000400()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	err = nSignSymTwo.CopyIn(
		&nSignSymOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyIn_000500()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"(",
		")",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	err = nSignSymTwo.CopyIn(
		&nSignSymOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_CopyIn_000600()",
		"")

	nSignSymMolecule := numberSignSymbolDtoMolecule{}

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err = nSignSymMolecule.copyIn(
		&nSignSymOne,
		nil,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymMolecule.copyIn()\n"+
			"because 'incomingNumSignSymbol' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_CopyIn_000700()",
		"")

	nSignSymMolecule := numberSignSymbolDtoMolecule{}

	nSignSymTwo,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err = nSignSymMolecule.copyIn(
		nil,
		&nSignSymTwo,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymMolecule.copyIn()\n"+
			"because 'targetNumSignSymbol' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_CopyIn_000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_CopyIn_000800()",
		"")

	nSignSymOne := NumberSignSymbolDto{}

	nSignSymMolecule := numberSignSymbolDtoMolecule{}

	nSignSymTwo,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.leadingNumSignChars = nil
	nSignSymTwo.trailingNumSignChars = nil

	err = nSignSymMolecule.copyIn(
		&nSignSymOne,
		&nSignSymTwo,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymMolecule.copyIn()\n"+
			"because 'incomingNumSignSymbol' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000100()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var nSignSymTwo NumberSignSymbolDto

	nSignSymTwo,
		err = nSignSymOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000200()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var nSignSymTwo NumberSignSymbolDto

	nSignSymTwo,
		err = nSignSymOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000300()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var nSignSymTwo NumberSignSymbolDto

	nSignSymTwo,
		err = nSignSymOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000400()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var nSignSymTwo NumberSignSymbolDto

	nSignSymTwo,
		err = nSignSymOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000500()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"(",
		")",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	var nSignSymTwo NumberSignSymbolDto

	nSignSymTwo,
		err = nSignSymOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_CopyOut_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_CopyOut_000600",
		"")

	nSignSymMolecule := numberSignSymbolDtoMolecule{}

	_,
		err := nSignSymMolecule.copyOut(
		nil,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:"+
			"Expected an error return from nSignSymMolecule.copyOut()\n"+
			"because 'incomingNumSignSymbol' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}
}

func TestNumberSignSymbolDto_CopyOut_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_CopyOut_000700",
		"")

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymOne.leadingNumSignChars = nil
	nSignSymOne.trailingNumSignChars = nil

	nSignSymMolecule := numberSignSymbolDtoMolecule{}

	_,
		err = nSignSymMolecule.copyOut(
		&nSignSymOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:"+
			"Expected an error return from nSignSymMolecule.copyOut()\n"+
			"because 'incomingNumSignSymbol' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_Empty_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Empty_000100()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo := NumberSignSymbolDto{}

	nSignSymOne.Empty()

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo == nSignSymOne .\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_Empty_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_Empty_000200",
		"")

	nSignSymElectron := numberSignSymbolDtoElectron{}

	err :=
		nSignSymElectron.emptyNumSignSymbol(
			nil,
			&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymElectron.emptyNumSignSymbol()\n"+
			"because 'numSignSymbol' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}
