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

	nSignSymTwo.numSignPosition = NumSymPos.After()

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

	nSignSymTwo.numSignPosition = NumSymPos.BeforeAndAfter()

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

func TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_GetTrailingNumSignFoundIndex()"

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
