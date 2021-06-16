package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumberSignSymbolDto_New_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000100()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'leadingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(leadingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '1' \n"+
			"len(leadingNumSign) is NOT EQUAL to '1' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	if leadingNumSign[0] != '+' {

		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSign[0] == '+' \n"+
			"leadingNumSign[0] is NOT EQUAL to '+' !\n"+
			"leadingNumSign[0] = '%v'\n",
			ePrefix,
			leadingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.Before() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.Before()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.Before()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Positive() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Positive()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Positive()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_New_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000200()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'leadingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(leadingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '1' \n"+
			"len(leadingNumSign) is NOT EQUAL to '1' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	if leadingNumSign[0] != '-' {

		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSign[0] == '-' \n"+
			"leadingNumSign[0] is NOT EQUAL to '-' !\n"+
			"leadingNumSign[0] = '%v'\n",
			ePrefix,
			leadingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.Before() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.Before()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.Before()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Negative() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Negative()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Negative()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_New_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000300()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"(",
		")",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'leadingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(leadingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '1' \n"+
			"len(leadingNumSign) is NOT EQUAL to '1' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	trailingNumSign := nSignSym.GetTrailingNumSignChars()

	if len(trailingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'trailingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(trailingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(trailingNumSign) == '1' \n"+
			"len(trailingNumSign) is NOT EQUAL to '1' !\n"+
			"len(trailingNumSign) = '%v'\n"+
			"trailingNumSign= '%v'\n",
			ePrefix,
			len(trailingNumSign),
			string(trailingNumSign))

		return
	}

	if leadingNumSign[0] != '(' {

		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSign[0] == '(' \n"+
			"leadingNumSign[0] is NOT EQUAL to '(' !\n"+
			"leadingNumSign[0] = '%v'\n",
			ePrefix,
			leadingNumSign[0])

		return
	}

	if trailingNumSign[0] != ')' {

		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSign[0] == ')' \n"+
			"trailingNumSign[0] is NOT EQUAL to ')' !\n"+
			"trailingNumSign[0] = '%v'\n",
			ePrefix,
			trailingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.BeforeAndAfter() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.BeforeAndAfter()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.BeforeAndAfter()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Negative() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Positive()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Positive()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_New_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000400()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) != 0 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '0' \n"+
			"len(leadingNumSign) is NOT EQUAL to '0' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	trailingNumSign := nSignSym.GetTrailingNumSignChars()

	if len(trailingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'trailingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(trailingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(trailingNumSign) == '1' \n"+
			"len(trailingNumSign) is NOT EQUAL to '1' !\n"+
			"len(trailingNumSign) = '%v'\n"+
			"trailingNumSign= '%v'\n",
			ePrefix,
			len(trailingNumSign),
			string(trailingNumSign))

		return
	}

	if trailingNumSign[0] != '+' {

		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSign[0] == '+' \n"+
			"trailingNumSign[0] is NOT EQUAL to '+' !\n"+
			"trailingNumSign[0] = '%v'\n",
			ePrefix,
			trailingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.After() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.After()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.After()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Positive() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Positive()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Positive()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_New_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000500()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) != 0 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '0' \n"+
			"len(leadingNumSign) is NOT EQUAL to '0' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	trailingNumSign := nSignSym.GetTrailingNumSignChars()

	if len(trailingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'trailingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(trailingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(trailingNumSign) == '1' \n"+
			"len(trailingNumSign) is NOT EQUAL to '1' !\n"+
			"len(trailingNumSign) = '%v'\n"+
			"trailingNumSign= '%v'\n",
			ePrefix,
			len(trailingNumSign),
			string(trailingNumSign))

		return
	}

	if trailingNumSign[0] != '-' {

		t.Errorf("%v - Error:\n"+
			"Expected trailingNumSign[0] == '-' \n"+
			"trailingNumSign[0] is NOT EQUAL to '-' !\n"+
			"trailingNumSign[0] = '%v'\n",
			ePrefix,
			trailingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.After() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.After()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.After()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Negative() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Negative()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Negative()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_New_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_New_000600()",
		"")

	nSignSymMechanics := numberSignSymbolDtoMechanics{}

	err := nSignSymMechanics.setNumberSignSymbol(
		nil,
		"+",
		"",
		false,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from nSignSymMechanics.setNumberSignSymbol()\n"+
			"because 'numSignSymbol' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_New_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberSignSymbolDto_New_000700()",
		"")

	numSignSymbol := NumberSignSymbolDto{}

	nSignSymMechanics := numberSignSymbolDtoMechanics{}

	err := nSignSymMechanics.setNumberSignSymbol(
		&numSignSymbol,
		"",
		"",
		false,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from nSignSymMechanics.setNumberSignSymbol()\n"+
			"because 'leadingNumberSign' and 'trailingNumberSign' are empty strings.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
		return
	}

}

func TestNumberSignSymbolDto_SetNumberSignSymbol_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_New_000100()"

	nSignSym,
		err := NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err = nSignSym.SetNumberSignSymbol(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	leadingNumSign := nSignSym.GetLeadingNumSignChars()

	if len(leadingNumSign) == 0 {

		t.Errorf("%v - Error:\n"+
			"'leadingNumSign' is a zero length array!\n",
			ePrefix)

		return
	}

	if len(leadingNumSign) > 1 {

		t.Errorf("%v - Error:\n"+
			"Expected len(leadingNumSign) == '1' \n"+
			"len(leadingNumSign) is NOT EQUAL to '1' !\n"+
			"len(leadingNumSign) = '%v'\n"+
			"leadingNumSign= '%v'\n",
			ePrefix,
			len(leadingNumSign),
			string(leadingNumSign))

		return
	}

	if leadingNumSign[0] != '+' {

		t.Errorf("%v - Error:\n"+
			"Expected leadingNumSign[0] == '+' \n"+
			"leadingNumSign[0] is NOT EQUAL to '+' !\n"+
			"leadingNumSign[0] = '%v'\n",
			ePrefix,
			leadingNumSign[0])

		return
	}

	err = nSignSym.IsValidInstanceError(ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	symbolPos := nSignSym.GetNumSignSymPosition()

	if symbolPos != NumSymPos.Before() {
		t.Errorf("%v - Error:\n"+
			"Expected Symbol Position == NumSymPos.Before()\n"+
			"Symbol Position is NOT EQUAL to NumSymPos.Before()!\n"+
			"Symbol Position string = '%v'\n"+
			"Symbol Position integer = '%v'\n",
			ePrefix,
			symbolPos.String(),
			symbolPos.XValueInt())

		return
	}

	signValue := nSignSym.GetNumSignVal()

	if signValue != NumSignVal.Positive() {

		t.Errorf("%v - Error:\n"+
			"Expected Sign Value == NumSignVal.Positive()\n"+
			"Sign Value is NOT EQUAL to NumSignVal.Positive()!\n"+
			"Sign Value string = '%v'\n"+
			"Sign Value integer = '%v'\n",
			ePrefix,
			signValue.String(),
			signValue.XValueInt())

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000100()"

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

	nSignSymOne.SetLeadingNumSignFoundIndex(99)

	if nSignSymOne.leadingNumSignFoundIndex != 99 {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '99'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundIndex == '%v' \n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

	if nSignSymOne.leadingNumSignFoundInNumber != true {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'true'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'false' \n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000200()"

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

	nSignSymOne.SetLeadingNumSignFoundIndex(-1)

	if nSignSymOne.leadingNumSignFoundIndex != -1 {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '-1'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundIndex == '%v' \n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

	if nSignSymOne.leadingNumSignFoundInNumber != false {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'true' \n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundIndex_000300()"

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

	nSignSymOne.SetLeadingNumSignFoundIndex(99)

	if nSignSymOne.leadingNumSignFoundIndex != 0 {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '0'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundIndex == '%v' \n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

	if nSignSymOne.leadingNumSignFoundInNumber != false {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'true' \n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000100()"

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

	nSignSymOne.SetLeadingNumSignFoundInNumber(true)

	if nSignSymOne.leadingNumSignFoundInNumber != true {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'true'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'false' \n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000200()"

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

	nSignSymOne.SetLeadingNumSignFoundInNumber(false)

	if nSignSymOne.leadingNumSignFoundInNumber != false {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'true' \n",
			ePrefix)

		return
	}

	if nSignSymOne.leadingNumSignFoundIndex != 0 {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundIndex == '0'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == '%v' \n",
			ePrefix,
			nSignSymOne.leadingNumSignFoundIndex)

		return
	}

}

func TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolDto_SetLeadingNumSignFoundInNumber_000300()"

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

	nSignSymOne.SetLeadingNumSignFoundInNumber(true)

	if nSignSymOne.leadingNumSignFoundInNumber != false {

		t.Errorf("%v - Error:\n"+
			"Expected nSignSymOne.leadingNumSignFoundInNumber == 'false'.\n"+
			"HOWEVER, nSignSymOne.leadingNumSignFoundInNumber == 'true' \n",
			ePrefix)

		return
	}

}
