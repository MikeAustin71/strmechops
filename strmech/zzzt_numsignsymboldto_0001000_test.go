package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumberSignSymbolDto_ClearLeadingNumSignTracking_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_ClearLeadingNumSignTracking_000100()"

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

	nSignSymOne.leadingNumSignFoundInNumber = true

	nSignSymOne.leadingNumSignFoundIndex = 5

	nSignSymOne.ClearLeadingNumSignTracking()

	if nSignSymOne.leadingNumSignFoundInNumber == true {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'\n"+
			"INSTEAD, nSignSymOne.leadingNumSignFoundInNumber == 'true'\n",
			ePrefix)

		return
	}

	if nSignSymOne.leadingNumSignFoundIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '0'\n"+
			"INSTEAD, nSignSymOne.leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

}

func TestNumberSignSymbolDto_ClearLeadingNumSignTracking_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_ClearLeadingNumSignTracking_000200()"

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

	nSignSymOne.leadingNumSignFoundInNumber = true

	nSignSymOne.leadingNumSignFoundIndex = 5

	nSignSymOne.lock = nil

	nSignSymOne.ClearLeadingNumSignTracking()

	if nSignSymOne.leadingNumSignFoundInNumber == true {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'\n"+
			"INSTEAD, nSignSymOne.leadingNumSignFoundInNumber == 'true'\n",
			ePrefix)

		return
	}

	if nSignSymOne.leadingNumSignFoundIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '0'\n"+
			"INSTEAD, nSignSymOne.leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

}

func TestNumberSignSymbolDto_ClearTrailingNumSignTracking_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_ClearTrailingNumSignTracking_000100()"

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

	nSignSymOne.trailingNumSignFoundInNumber = true

	nSignSymOne.trailingNumSignFoundIndex = 5

	nSignSymOne.ClearTrailingNumSignTracking()

	if nSignSymOne.trailingNumSignFoundInNumber == true {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.trailingNumSignFoundInNumber == 'false'\n"+
			"INSTEAD, nSignSymOne.trailingNumSignFoundInNumber == 'true'\n",
			ePrefix)

		return
	}

	if nSignSymOne.trailingNumSignFoundIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.trailingNumSignFoundIndex == '0'\n"+
			"INSTEAD, nSignSymOne.trailingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignSymOne.trailingNumSignFoundIndex)

		return
	}

}

func TestNumberSignSymbolDto_ClearTrailingNumSignTracking_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_ClearTrailingNumSignTracking_000200()"

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

	nSignSymOne.trailingNumSignFoundInNumber = true

	nSignSymOne.trailingNumSignFoundIndex = 5

	nSignSymOne.lock = nil

	nSignSymOne.ClearTrailingNumSignTracking()

	if nSignSymOne.trailingNumSignFoundInNumber == true {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.trailingNumSignFoundInNumber == 'false'\n"+
			"INSTEAD, nSignSymOne.trailingNumSignFoundInNumber == 'true'\n",
			ePrefix)

		return
	}

	if nSignSymOne.trailingNumSignFoundIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymOne.trailingNumSignFoundIndex == '0'\n"+
			"INSTEAD, nSignSymOne.trailingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignSymOne.trailingNumSignFoundIndex)

		return
	}

}

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

func TestNumberSignSymbolDto_CopyOut_000800(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_CopyOut_000800()"

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

	nSignSymOne.lock = nil

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

	nSignSymTwo.Empty()

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

	ePrefix := "TestNumberSignSymbolDto_Empty_000200()"

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
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nSignSymTwo.IsEmpty() {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo.IsEmpty() == 'false' .\n"+
			"HOWEVER, THE RETURN VALUE WAS 'true'!\n",
			ePrefix)

		return
	}

	nSignSymOne.Empty()

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo != nSignSymOne .\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)

		return
	}

	if !nSignSymOne.IsEmpty() {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.IsEmpty() == 'true' .\n"+
			"HOWEVER, THE RETURN VALUE WAS 'false'!\n",
			ePrefix)

		return
	}

	err = nSignSymOne.IsValidInstanceError(ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymOne.IsValidInstanceError().\n"+
			"because nSignSymOne is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_Empty_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_Empty_000300",
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

func TestNumberSignSymbolDto_Equal_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Equal_000100()"

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

	if !nSignSymOne.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne==nSignSymOne.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_Equal_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Equal_000200()"

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
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo==nSignSymOne.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo.numSignPosition = NumSignSymPos.After()

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.numSignPosition changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.leadingNumSignChars = []rune{'X', 'X'}

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.leadingNumSignChars changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.trailingNumSignChars = []rune{'X', 'X'}

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.trailingNumSignChars changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.numSignValueType = NumSignVal.Negative()

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.numSignValueType changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.numSignPosition = NumSignSymPos.BeforeAndAfter()

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.numSignPosition changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.numSymbolClass = NumSymClass.CurrencySign()

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.numSignPosition changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.leadingNumSignFoundIndex = 99

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.leadingNumSignFoundIndex changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.leadingNumSignFoundInNumber = true

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.leadingNumSignFoundInNumber changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.trailingNumSignFoundIndex = 99

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.trailingNumSignFoundIndex changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo.trailingNumSignFoundInNumber = true

	if nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymTwo!=nSignSymOne\n"+
			"because nSignSymTwo.trailingNumSignFoundInNumber changed.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

	nSignSymTwo,
		err = nSignSymOne.CopyOut(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if !nSignSymTwo.Equal(&nSignSymOne) {
		t.Errorf("%v - Error: Final Run\n"+
			"Expected nSignSymTwo==nSignSymOne\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_Equal_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Equal_000300()"

	var nSignSymOne NumberSignSymbolDto

	var err error

	nSignSymOne,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymElectron := numberSignSymbolDtoElectron{}

	areEqual := nSignSymElectron.equalNumSignSymbolObjects(
		&nSignSymOne,
		nil)

	if areEqual != false {
		t.Errorf("%v - Error:\n"+
			"Expected areEqual=='false' because\n"+
			"'numSignSymbolTwo' is a nil pointer.\n"+
			"HOWEVER, areEqual=='true'!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_Equal_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Equal_000400()"

	var nSignSymTwo NumberSignSymbolDto

	var err error

	nSignSymTwo,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymElectron := numberSignSymbolDtoElectron{}

	areEqual := nSignSymElectron.equalNumSignSymbolObjects(
		nil,
		&nSignSymTwo)

	if areEqual != false {
		t.Errorf("%v - Error:\n"+
			"Expected areEqual=='false' because\n"+
			"'numSignSymbolOne' is a nil pointer.\n"+
			"HOWEVER, areEqual=='true'!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_Equal_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_Equal_000500()"

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

	nSignSymOne.lock = nil

	if !nSignSymOne.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne==nSignSymOne.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000100()"

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

	if !nSignSymOne.Equal(&nSignSymOne) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne==nSignSymOne.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000200()"

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
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("nSignSymTwo - %v", err.Error())
		return
	}

	if !nSignSymOne.EqualNumberSignRunes(&nSignSymTwo) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne==nSignSymTwo.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000250(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000250()"

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
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("nSignSymTwo - %v", err.Error())
		return
	}

	nSignSymOne.lock = nil

	if !nSignSymOne.EqualNumberSignRunes(&nSignSymTwo) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne==nSignSymTwo.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000300()"

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
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("nSignSymTwo - %v", err.Error())
		return
	}

	if nSignSymOne.EqualNumberSignRunes(&nSignSymTwo) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne!=nSignSymTwo.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000400()"

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
		err = NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("nSignSymTwo - %v", err.Error())
		return
	}

	if nSignSymOne.EqualNumberSignRunes(&nSignSymTwo) {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne!=nSignSymTwo.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000500()"

	nSignSymElectron := numberSignSymbolDtoElectron{}

	numSignSymbolTwo := NumberSignSymbolDto{}

	areEqual := nSignSymElectron.equalNumSignRuneArrays(
		nil,
		&numSignSymbolTwo)

	if areEqual {
		t.Errorf("%v - Error:\n"+
			"Expected areEqual=='false' because\n"+
			"'numSignSymbolOne' is a nil pointer.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000600(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000600()"

	nSignSymElectron := numberSignSymbolDtoElectron{}

	numSignSymbolOne := NumberSignSymbolDto{}

	areEqual := nSignSymElectron.equalNumSignRuneArrays(
		&numSignSymbolOne,
		nil)

	if areEqual {
		t.Errorf("%v - Error:\n"+
			"Expected areEqual=='false' because\n"+
			"'numSignSymbolTwo' is a nil pointer.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_EqualNumberSignRunes_000700(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_EqualNumberSignRunes_000600()"

	nSignSymElectron := numberSignSymbolDtoElectron{}

	numSignSymbolOne := NumberSignSymbolDto{}

	numSignSymbolTwo := NumberSignSymbolDto{}

	areEqual := nSignSymElectron.equalNumSignRuneArrays(
		&numSignSymbolOne,
		&numSignSymbolTwo)

	if !areEqual {
		t.Errorf("%v - Error:\n"+
			"Expected areEqual=='true' because\n"+
			"'numSignSymbolOne' and 'numSignSymbolTwo' have empty rune arrays.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_GetLeadingNumSignFoundIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetLeadingNumSignFoundIndex_000100()"

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

	nSignSymOne.leadingNumSignFoundIndex = 99

	leadingNumFoundIdx :=
		nSignSymOne.GetLeadingNumSignFoundIndex()

	if leadingNumFoundIdx != 99 {
		t.Errorf("%v - Error:\n"+
			"Expected leadingNumFoundIdx=='99'.\n"+
			"HOWEVER, leadingNumFoundIdx=='%v'!\n",
			ePrefix,
			leadingNumFoundIdx)
		return
	}

}

func TestNumberSignSymbolDto_GetLeadingNumSignFoundIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetLeadingNumSignFoundIndex_000200()"

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

	nSignSymOne.leadingNumSignFoundIndex = 99

	nSignSymOne.lock = nil

	leadingNumFoundIdx :=
		nSignSymOne.GetLeadingNumSignFoundIndex()

	if leadingNumFoundIdx != 99 {
		t.Errorf("%v - Error:\n"+
			"Expected leadingNumFoundIdx=='99'.\n"+
			"HOWEVER, leadingNumFoundIdx=='%v'!\n",
			ePrefix,
			leadingNumFoundIdx)
		return
	}

}

func TestNumberSignSymbolDto_GetLeadingNumSignFoundInNumber_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetLeadingNumSignFoundInNumber_000100()"

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

	nSignSymOne.leadingNumSignFoundInNumber = true

	leadingNumSignFoundInNum :=
		nSignSymOne.GetLeadingNumSignFoundInNumber()

	if !leadingNumSignFoundInNum {
		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSignFoundInNum=='true'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='false'!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_GetLeadingNumSignFoundInNumber_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetLeadingNumSignFoundInNumber_000200()"

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

	nSignSymOne.leadingNumSignFoundInNumber = true

	nSignSymOne.lock = nil

	leadingNumSignFoundInNum :=
		nSignSymOne.GetLeadingNumSignFoundInNumber()

	if !leadingNumSignFoundInNum {
		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSignFoundInNum=='true'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='false'!\n",
			ePrefix)
		return
	}

}

func TestNumberSignSymbolDto_GetNumericSymbolClass_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumericSymbolClass_000100()"

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

	numSymClass := nSignSymOne.GetNumericSymbolClass()

	if numSymClass != NumSymClass.NumberSign() {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.GetNumericSymbolClass()=='NumberSign'.\n"+
			"HOWEVER, numSymClass string =='%v'\n"+
			"numSymClass integer == '%v'\n",
			ePrefix,
			numSymClass.String(),
			numSymClass.XValueInt())
		return
	}

}

func TestNumberSignSymbolDto_GetNumericSymbolClass_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumericSymbolClass_000200()"

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

	nSignSymOne.lock = nil

	numSymClass := nSignSymOne.GetNumericSymbolClass()

	if numSymClass != NumSymClass.NumberSign() {
		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.GetNumericSymbolClass()=='NumberSign'.\n"+
			"HOWEVER, numSymClass string =='%v'\n"+
			"numSymClass integer == '%v'\n",
			ePrefix,
			numSymClass.String(),
			numSymClass.XValueInt())
		return
	}

}

func TestNumberSignSymbolDto_GetNumSignArithmeticVal_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumSignArithmeticVal_000100()"

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

	arithmeticVal := nSignSymOne.GetNumSignArithmeticVal()

	if arithmeticVal != 1 {
		t.Errorf("%v - Error:\n"+
			"Expected arithmeticVal=='1'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='%v'!\n",
			ePrefix,
			arithmeticVal)
		return
	}

}

func TestNumberSignSymbolDto_GetNumSignArithmeticVal_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumSignArithmeticVal_000200()"

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

	arithmeticVal := nSignSymOne.GetNumSignArithmeticVal()

	if arithmeticVal != -1 {
		t.Errorf("%v - Error:\n"+
			"Expected arithmeticVal=='-1'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='%v'!\n",
			ePrefix,
			arithmeticVal)
		return
	}

}

func TestNumberSignSymbolDto_GetNumSignArithmeticVal_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumSignArithmeticVal_000300()"

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

	nSignSymOne.numSignValueType = NumSignVal.Zero()

	arithmeticVal := nSignSymOne.GetNumSignArithmeticVal()

	if arithmeticVal != 0 {
		t.Errorf("%v - Error:\n"+
			"Expected arithmeticVal=='0'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='%v'!\n",
			ePrefix,
			arithmeticVal)
		return
	}

}

func TestNumberSignSymbolDto_GetNumSignArithmeticVal_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetNumSignArithmeticVal_000400()"

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

	nSignSymOne.lock = nil

	arithmeticVal := nSignSymOne.GetNumSignArithmeticVal()

	if arithmeticVal != 1 {
		t.Errorf("%v - Error:\n"+
			"Expected arithmeticVal=='1'.\n"+
			"HOWEVER, leadingNumSignFoundInNum=='%v'!\n",
			ePrefix,
			arithmeticVal)
		return
	}

}

func TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex_000100()"

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

	nSignSymOne.trailingNumSignFoundIndex = 99

	trailingNumSignFoundIndex :=
		nSignSymOne.GetTrailingNumSignFoundIndex()

	if trailingNumSignFoundIndex != 99 {
		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSignFoundIndex=='99'.\n"+
			"HOWEVER, trailingNumSignFoundIndex=='%v'!\n",
			ePrefix,
			trailingNumSignFoundIndex)
		return
	}

}

func TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex_000200()"

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

	nSignSymOne.trailingNumSignFoundIndex = 99

	nSignSymOne.lock = nil

	trailingNumSignFoundIndex :=
		nSignSymOne.GetTrailingNumSignFoundIndex()

	if trailingNumSignFoundIndex != 99 {
		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSignFoundIndex=='99'.\n"+
			"HOWEVER, trailingNumSignFoundIndex=='%v'!\n",
			ePrefix,
			trailingNumSignFoundIndex)
		return
	}

}

func TestNumberSignSymbolDto_GetTrailingNumSignFoundInNumber_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetTrailingNumSignFoundInNumber_000100()"

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

	nSignSymOne.trailingNumSignFoundInNumber = true

	trailingNumSignFoundInNumber :=
		nSignSymOne.GetTrailingNumSignFoundInNumber()

	if trailingNumSignFoundInNumber != true {
		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSignFoundInNumber=='true'.\n"+
			"HOWEVER, trailingNumSignFoundInNumber=='false'!\n",
			ePrefix)
		return
	}
}

func TestNumberSignSymbolDto_GetTrailingNumSignFoundInNumber_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetTrailingNumSignFoundInNumber_000200()"

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

	nSignSymOne.trailingNumSignFoundInNumber = true

	nSignSymOne.lock = nil

	trailingNumSignFoundInNumber :=
		nSignSymOne.GetTrailingNumSignFoundInNumber()

	if trailingNumSignFoundInNumber != true {
		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSignFoundInNumber=='true'.\n"+
			"HOWEVER, trailingNumSignFoundInNumber=='false'!\n",
			ePrefix)
		return
	}
}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000100()"

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

	hostRunes := []rune(" -123.456")

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(hostRunes, 1)

	if foundLeadingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='true'.\n"+
			"HOWEVER, foundLeadingNumSign=='false'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000200()"

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

	hostRunes := []rune(" -123.456")

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(hostRunes, 0)

	if foundLeadingNumSign != false {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='false'.\n"+
			"HOWEVER, foundLeadingNumSign=='true'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000300()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		" -",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	hostRunes := []rune(" -123.456")

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(
			hostRunes,
			0)

	if foundLeadingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='true'.\n"+
			"HOWEVER, foundLeadingNumSign=='false'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000400()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"- ",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	hostRunes := []rune("- 123.456")

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(
			hostRunes,
			0)

	if foundLeadingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='true'.\n"+
			"HOWEVER, foundLeadingNumSign=='false'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000500()"

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

	hostRunes := []rune("-123.456")

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(
			hostRunes,
			0)

	if foundLeadingNumSign != false {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='false'.\n"+
			"HOWEVER, foundLeadingNumSign=='true'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000600(t *testing.T) {

	sMechPreon := strMechPreon{}

	targetRunes := []rune(" -")

	isTargetRunesIndex :=
		sMechPreon.isTargetRunesIndex(
			nil,
			0,
			targetRunes)

	if isTargetRunesIndex == true {
		t.Error("Error:\n" +
			"Expected isTargetRunesIndex=='false' because\n" +
			"'hostRunes' is nil.\n" +
			"HOWEVER, isTargetRunesIndex=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000700(t *testing.T) {

	sMechPreon := strMechPreon{}

	hostRunes := []rune(" -123489.2")

	isTargetRunesIndex :=
		sMechPreon.isTargetRunesIndex(
			hostRunes,
			0,
			nil)

	if isTargetRunesIndex == true {
		t.Error("Error:\n" +
			"Expected isTargetRunesIndex=='false' because\n" +
			"'targetRunes' is nil.\n" +
			"HOWEVER, isTargetRunesIndex=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000800(t *testing.T) {

	sMechPreon := strMechPreon{}

	hostRunes := []rune(" -123489.2")

	targetRunes := []rune("-")

	isTargetRunesIndex :=
		sMechPreon.isTargetRunesIndex(
			hostRunes,
			99,
			targetRunes)

	if isTargetRunesIndex == true {
		t.Error("Error:\n" +
			"Expected isTargetRunesIndex=='false' because\n" +
			"'hostStartIndex' exceeds length of 'hostRunes'.\n" +
			"HOWEVER, isTargetRunesIndex=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000900(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsLeadingNumSignAtHostIndex_000900()"

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

	hostRunes := []rune(" -123.456")

	nSignSymOne.lock = nil

	foundLeadingNumSign :=
		nSignSymOne.IsLeadingNumSignAtHostIndex(hostRunes, 1)

	if foundLeadingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundLeadingNumSign=='true'.\n"+
			"HOWEVER, foundLeadingNumSign=='false'!\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000100",
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

	nSignSymOne.leadingNumSignFoundInNumber = true

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == false {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='true'.\n" +
			"HOWEVER, isNumSignSymFound=='false'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000200",
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

	nSignSymOne.leadingNumSignFoundInNumber = false

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == true {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='false'.\n" +
			"HOWEVER, isNumSignSymFound=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000300",
		"")

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

	nSignSymOne.trailingNumSignFoundInNumber = true

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == false {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='true'.\n" +
			"HOWEVER, isNumSignSymFound=='false'!\n")

		return
	}
}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000400",
		"")

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

	nSignSymOne.trailingNumSignFoundInNumber = false

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == true {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='false'.\n" +
			"HOWEVER, isNumSignSymFound=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000500",
		"")

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

	nSignSymOne.leadingNumSignFoundInNumber = true
	nSignSymOne.trailingNumSignFoundInNumber = true

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == false {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='true'.\n" +
			"HOWEVER, isNumSignSymFound=='false'!\n")

		return
	}
}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000600",
		"")

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

	nSignSymOne.leadingNumSignFoundInNumber = false
	nSignSymOne.trailingNumSignFoundInNumber = false

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == true {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='false'.\n" +
			"HOWEVER, isNumSignSymFound=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000700",
		"")

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

	nSignSymOne.leadingNumSignFoundInNumber = true
	nSignSymOne.trailingNumSignFoundInNumber = false

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == true {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='false'.\n" +
			"HOWEVER, isNumSignSymFound=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000800",
		"")

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

	nSignSymOne.leadingNumSignFoundInNumber = false
	nSignSymOne.trailingNumSignFoundInNumber = true

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == true {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='false'.\n" +
			"HOWEVER, isNumSignSymFound=='true'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000900(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_000900",
		"")

	nSignSymQuark := numberSignSymbolDtoQuark{}

	_,
		err :=
		nSignSymQuark.isNumberSignSymbolFoundInNumber(
			nil,
			&ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymQuark.isNumberSignSymbolFoundInNumber()\n"+
			"because 'numSignSymbol' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsNumSignSymbolFoundInNumber_001000",
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

	nSignSymOne.leadingNumSignFoundInNumber = true

	nSignSymOne.lock = nil

	isNumSignSymFound :=
		nSignSymOne.IsNumSignSymbolFoundInNumber()

	if isNumSignSymFound == false {
		t.Error("Error:\n" +
			"Expected isNumSignSymFound=='true'.\n" +
			"HOWEVER, isNumSignSymFound=='false'!\n")

		return
	}

}

func TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000100()"

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

	hostRunes := []rune(" 123.456-")

	foundTrailingNumSign :=
		nSignSymOne.IsTrailingNumSignAtHostIndex(
			hostRunes,
			8)

	if foundTrailingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundTrailingNumSign=='true'.\n"+
			"HOWEVER, foundTrailingNumSign=='false'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000200()"

	nSignSymOne,
		err := NumberSignSymbolDto{}.New(
		"",
		"- ",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	hostRunes := []rune(" 123.456- ")

	foundTrailingNumSign :=
		nSignSymOne.IsTrailingNumSignAtHostIndex(
			hostRunes, 8)

	if foundTrailingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundTrailingNumSign=='true'.\n"+
			"HOWEVER, foundTrailingNumSign=='false'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000300()"

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

	hostRunes := []rune(" 123.456- ")

	foundTrailingNumSign :=
		nSignSymOne.IsTrailingNumSignAtHostIndex(
			hostRunes, 9)

	if foundTrailingNumSign != false {

		t.Errorf("%v - Error:\n"+
			"Expected foundTrailingNumSign=='false'.\n"+
			"HOWEVER, foundTrailingNumSign=='true'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000400()"

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

	hostRunes := []rune(" 123.456- ")

	foundTrailingNumSign :=
		nSignSymOne.IsTrailingNumSignAtHostIndex(
			hostRunes, 8)

	if foundTrailingNumSign != false {

		t.Errorf("%v - Error:\n"+
			"Expected foundTrailingNumSign=='false'.\n"+
			"HOWEVER, foundTrailingNumSign=='true'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsTrailingNumSignAtHostIndex_000500()"

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

	hostRunes := []rune(" 123.456-")

	nSignSymOne.lock = nil

	foundTrailingNumSign :=
		nSignSymOne.IsTrailingNumSignAtHostIndex(
			hostRunes,
			8)

	if foundTrailingNumSign != true {

		t.Errorf("%v - Error:\n"+
			"Expected foundTrailingNumSign=='true'.\n"+
			"HOWEVER, foundTrailingNumSign=='false'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsValidInstance_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsValidInstance_000100()"

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

	if !nSignSymOne.IsValidInstance() {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.IsValidInstance()=='true'.\n"+
			"HOWEVER, nSignSymOne.IsValidInstance()=='false'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsValidInstance_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsValidInstance_000200()"

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

	nSignSymOne.leadingNumSignChars = nil
	nSignSymOne.trailingNumSignChars = nil

	if nSignSymOne.IsValidInstance() {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.IsValidInstance()=='false'.\n"+
			"HOWEVER, nSignSymOne.IsValidInstance()=='true'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsValidInstance_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_IsValidInstance_000300",
		"")

	nSignSymElectron := numberSignSymbolDtoElectron{}

	_,
		err := nSignSymElectron.testValidityOfNumSignSymbol(
		nil,
		&ePrefix)

	if err == nil {

		t.Errorf("%v - Error:\n"+
			"Expected error return from nSignSymElectron.testValidityOfNumSignSymbol()\n"+
			"because 'numSignSymbol' is a nil pointer."+
			"HOWEVER, NO ERROR WAS RETURND!\n",
			ePrefix.String())

		return
	}

}

func TestNumberSignSymbolDto_IsValidInstance_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsValidInstance_000400()"

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

	nSignSymOne.lock = nil

	if !nSignSymOne.IsValidInstance() {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.IsValidInstance()=='true'.\n"+
			"HOWEVER, nSignSymOne.IsValidInstance()=='false'!\n",
			ePrefix)

		return
	}
}

func TestNumberSignSymbolDto_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_IsValidInstanceError_000100()"

	nSignSymOne := NumberSignSymbolDto{}

	err := nSignSymOne.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error:\n"+
			"Expected an error return from nSignSymOne.IsValidInstanceError()\n"+
			"because 'nSignSymOne' is an empty object.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
		return
	}

}
