package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumSignSymbolPositionTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumSignSymbolPosition,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumSignSymbolPositionTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Before",
		"After",
		"BeforeAndAfter",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumSignSymbolPosition(0).None())

	enumValues =
		append(enumValues, NumSignSymbolPosition(0).Before())

	enumValues =
		append(enumValues, NumSignSymbolPosition(0).After())

	enumValues =
		append(enumValues, NumSignSymbolPosition(0).BeforeAndAfter())

	intValues =
		append(intValues, NumSignSymPos.None().XValueInt())

	intValues =
		append(intValues, NumSignSymPos.Before().XValueInt())

	intValues =
		append(intValues, NumSignSymPos.After().XValueInt())

	intValues =
		append(intValues, NumSignSymPos.BeforeAndAfter().XValueInt())

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
			"Length Of intValues   = '%v'\n"+
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

func TestNumSignSymbolPosition_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolPosition_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumSignSymbolPositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numSignSymbolPos1, numSignSymbolPos2,
		numSignSymbolPos3, numSignSymbolPos4,
		numSignSymbolPos5, numSignSymbolPos6 NumSignSymbolPosition

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numSignSymbolPos1 = enumValues[i]

		isValid = numSignSymbolPos1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: numSignSymbolPos1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"numSignSymbolPos1 string value  = '%v'\n"+
					"numSignSymbolPos1 integer value = '%v'\n",
					ePrefix.String(),
					numSignSymbolPos1.String(),
					numSignSymbolPos1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numSignSymbolPos1 string value  = '%v'\n"+
				"numSignSymbolPos1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numSignSymbolPos1.String(),
				numSignSymbolPos1.XValueInt())

			return

		}

		numSignSymbolPos2,
			err = numSignSymbolPos1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numSignSymbolPos1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numSignSymbolPos1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numSignSymbolPos1.String(),
				err.Error())

			return
		}

		if numSignSymbolPos2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numSignSymbolPos2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numSignSymbolPos2 string value  = '%v'\n"+
				"numSignSymbolPos2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numSignSymbolPos2.String(),
				numSignSymbolPos2.XValueInt())

			return
		}

		numSignSymbolPos3 = enumValues[i]

		if numSignSymbolPos3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numSignSymbolPos3.XValueInt() != intValues[%v]\n"+
				"numSignSymbolPos3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolPos3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numSignSymbolPos4,
			err = numSignSymbolPos3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numSignSymbolPos3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numSignSymbolPos4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numSignSymbolPos4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numSignSymbolPos4 string value  = '%v'\n"+
				"numSignSymbolPos4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numSignSymbolPos4.String(),
				numSignSymbolPos4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numSignSymbolPos5 = numSignSymbolPos1.XValue()

		numSignSymbolPos6 = numSignSymbolPos2.XValue()

		if numSignSymbolPos5 != numSignSymbolPos6 {
			t.Errorf("%v\n"+
				"Error: numSignSymbolPos5 != numSignSymbolPos6\n"+
				"numSignSymbolPos5 = numSignSymbolPos1.XValue()\n"+
				"numSignSymbolPos6 = numSignSymbolPos2.XValue()\n"+
				"numSignSymbolPos5 string value  = '%v'\n"+
				"numSignSymbolPos5 integer value = '%v'\n"+
				"numSignSymbolPos6 string value  = '%v'\n"+
				"numSignSymbolPos6 integer value = '%v'\n",
				ePrefix.String(),
				numSignSymbolPos5.String(),
				numSignSymbolPos5.XValueInt(),
				numSignSymbolPos6.String(),
				numSignSymbolPos6.XValueInt())

			return
		}

		_,
			err = numSignSymbolPos6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolPos6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolPos6.String())

			return
		}

		_,
			err = numSignSymbolPos6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolPos6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolPos6.String())

			return
		}

		_,
			err = numSignSymbolPos6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolPos6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolPos6.String())

			return
		}

	}

	return
}

func TestNumSignSymbolPosition_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolPosition_XReturnNoneIfInvalid_000200()",
		"")

	numSignSymbolPos := NumSignSymbolPosition(-972)

	valueNone := numSignSymbolPos.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolPosition(-972)\n"+
			"would return name of 'None' from \n"+
			"numSignSymbolPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumSignSymbolPosition := numSignSymbolPos.String()

	strNumSignSymbolPosition = strings.ToLower(strNumSignSymbolPosition)

	if !strings.Contains(strNumSignSymbolPosition, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolPosition(-972).String()\n"+
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
		NumSignSymbolPositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numSignSymbolPos2 NumSignSymbolPosition

	numSignSymbolPos2 = enumValues[1].XReturnNoneIfInvalid()

	if numSignSymbolPos2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numSignSymbolPos2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numSignSymbolPos2 string value  = '%v'\n"+
			"numSignSymbolPos2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numSignSymbolPos2.String(),
			numSignSymbolPos2.XValueInt())
		return
	}

	return
}

func TestNumSignSymbolPosition_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolPosition_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numSignSymbolPos := NumSignSymbolPosition(expectedIntValue)

	actualIntValue := numSignSymbolPos.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numSignSymbolPos integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numSignSymbolPos integer value = '%v'\n"+
			"Actual numSignSymbolPos integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numSignSymbolPos.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolPosition(-972)\n"+
			"would return name of 'None' from \n"+
			"numSignSymbolPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
