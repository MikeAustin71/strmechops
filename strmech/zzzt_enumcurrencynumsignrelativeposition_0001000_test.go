package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func CurrencyNumSignRelativePositionTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []CurrencyNumSignRelativePosition,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CurrencyNumSignRelativePositionTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"OutsideNumSign",
		"InsideNumSign",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, CurrencyNumSignRelativePosition(0).None())

	enumValues =
		append(enumValues, CurrencyNumSignRelativePosition(0).OutsideNumSign())

	enumValues =
		append(enumValues, CurrencyNumSignRelativePosition(0).InsideNumSign())

	intValues =
		append(intValues, CurrNumSignRelPos.None().XValueInt())

	intValues =
		append(intValues, CurrNumSignRelPos.OutsideNumSign().XValueInt())

	intValues =
		append(intValues, CurrNumSignRelPos.InsideNumSign().XValueInt())

	if lenUcNames != len(intValues) {
		err = fmt.Errorf("%v\n"+
			"Error: Length of Upper Case Names ('ucNames')\n"+
			"DOES NOT MATCH the length of 'intVales'\n"+
			"Length Of ucNames   = '%v'\n"+
			"Length of intValues = '%v'\n",
			ePrefix.String(),
			lenUcNames,
			len(intValues))

		return ucNames, lcNames, intValues, enumValues, err
	}

	if len(intValues) != len(enumValues) {
		err = fmt.Errorf("%v\n"+
			"Error: Length of 'intValues' DOES NOT MATCH\n"+
			"the length of 'enumValues'\n"+
			"Length Of intValues  = '%v'\n"+
			"Length of enumValues = '%v'\n",
			ePrefix.String(),
			len(intValues),
			len(enumValues))

		return ucNames, lcNames, intValues, enumValues, err

	}

	for i := 0; i < len(intValues); i++ {

		if intValues[i] != enumValues[i].XValueInt() {
			err = fmt.Errorf("%v\n"+
				"Error: Integer Values DO NOT MATCH!\n"+
				"intValues[%v] != enumValues[%v].XValueInt()\n"+
				"intValues[%v] integer value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				i,
				intValues[i],
				i,
				enumValues[i].XValueInt())

			return ucNames, lcNames, intValues, enumValues, err
		}

	}

	return ucNames, lcNames, intValues, enumValues, err
}

func TestCurrencyNumSignRelativePosition_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCurrencyNumSignRelativePosition_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		CurrencyNumSignRelativePositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var currNumSignRelPos1, currNumSignRelPos2,
		currNumSignRelPos3, currNumSignRelPos4,
		currNumSignRelPos5, currNumSignRelPos6 CurrencyNumSignRelativePosition

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		currNumSignRelPos1 = enumValues[i]

		isValid = currNumSignRelPos1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: CurrencyNumSignRelativePosition1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"currNumSignRelPos1 string value  = '%v'\n"+
					"currNumSignRelPos1 integer value = '%v'\n",
					ePrefix.String(),
					currNumSignRelPos1.String(),
					currNumSignRelPos1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"currNumSignRelPos1 string value  = '%v'\n"+
				"currNumSignRelPos1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				currNumSignRelPos1.String(),
				currNumSignRelPos1.XValueInt())

			return

		}

		currNumSignRelPos2,
			err = currNumSignRelPos1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  currNumSignRelPos1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"currNumSignRelPos1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				currNumSignRelPos1.String(),
				err.Error())

			return
		}

		if currNumSignRelPos2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"currNumSignRelPos2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"currNumSignRelPos2 string value  = '%v'\n"+
				"currNumSignRelPos2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				currNumSignRelPos2.String(),
				currNumSignRelPos2.XValueInt())

			return
		}

		currNumSignRelPos3 = enumValues[i]

		if currNumSignRelPos3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: currNumSignRelPos3.XValueInt() != intValues[%v]\n"+
				"currNumSignRelPos3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				currNumSignRelPos3.XValueInt(),
				i,
				intValues[i])

			return
		}

		currNumSignRelPos4,
			err = currNumSignRelPos3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by currNumSignRelPos3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if currNumSignRelPos4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: currNumSignRelPos4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"currNumSignRelPos4 string value  = '%v'\n"+
				"currNumSignRelPos4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				currNumSignRelPos4.String(),
				currNumSignRelPos4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		currNumSignRelPos5 = currNumSignRelPos1.XValue()

		currNumSignRelPos6 = currNumSignRelPos2.XValue()

		if currNumSignRelPos5 != currNumSignRelPos6 {
			t.Errorf("%v\n"+
				"Error: currNumSignRelPos5 != currNumSignRelPos6\n"+
				"currNumSignRelPos5 = currNumSignRelPos1.XValue()\n"+
				"currNumSignRelPos6 = currNumSignRelPos2.XValue()\n"+
				"currNumSignRelPos5 string value  = '%v'\n"+
				"currNumSignRelPos5 integer value = '%v'\n"+
				"currNumSignRelPos6 string value  = '%v'\n"+
				"currNumSignRelPos6 integer value = '%v'\n",
				ePrefix.String(),
				currNumSignRelPos5.String(),
				currNumSignRelPos5.XValueInt(),
				currNumSignRelPos6.String(),
				currNumSignRelPos6.XValueInt())

			return
		}

		_,
			err = currNumSignRelPos6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from currNumSignRelPos6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"currNumSignRelPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				currNumSignRelPos6.String())

			return
		}

		_,
			err = currNumSignRelPos6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from currNumSignRelPos6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"currNumSignRelPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				currNumSignRelPos6.String())

			return
		}

		_,
			err = currNumSignRelPos6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from currNumSignRelPos6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"currNumSignRelPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				currNumSignRelPos6.String())

			return
		}

	}

	return
}

func TestCurrencyNumSignRelativePosition_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCurrencyNumSignRelativePosition_XReturnNoneIfInvalid_000200()",
		"")

	currNumSignRelPos := CurrencyNumSignRelativePosition(-972)

	valueNone := currNumSignRelPos.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CurrencyNumSignRelativePosition(-972)\n"+
			"would return name of 'None' from \n"+
			"currNumSignRelPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strCurrencyNumSignRelativePosition := currNumSignRelPos.String()

	strCurrencyNumSignRelativePosition = strings.ToLower(strCurrencyNumSignRelativePosition)

	if !strings.Contains(strCurrencyNumSignRelativePosition, "error") {

		t.Errorf("%v\n"+
			"Error: Expected CurrencyNumSignRelativePosition(-972).String()\n"+
			"would return an error because it is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	_,
		_,
		_,
		enumValues,
		err :=
		CurrencyNumSignRelativePositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var currNumSignRelPos2 CurrencyNumSignRelativePosition

	currNumSignRelPos2 = enumValues[1].XReturnNoneIfInvalid()

	if currNumSignRelPos2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: currNumSignRelPos2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"currNumSignRelPos2 string value  = '%v'\n"+
			"currNumSignRelPos2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			currNumSignRelPos2.String(),
			currNumSignRelPos2.XValueInt())
		return
	}

	return
}

func TestCurrencyNumSignRelativePosition_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCurrencyNumSignRelativePosition_XValueInt_000300()",
		"")

	expectedIntValue := -972

	currNumSignRelPos := CurrencyNumSignRelativePosition(expectedIntValue)

	actualIntValue := currNumSignRelPos.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected currNumSignRelPos integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected currNumSignRelPos integer value = '%v'\n"+
			"Actual currNumSignRelPos integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := currNumSignRelPos.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CurrencyNumSignRelativePosition(-972)\n"+
			"would return name of 'None' from \n"+
			"currNumSignRelPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
