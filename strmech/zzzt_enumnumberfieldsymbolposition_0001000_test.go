package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumberFieldSymbolPositionTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumberFieldSymbolPosition,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberFieldSymbolPositionTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"InsideNumField",
		"OutsideNumField",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumberFieldSymbolPosition(0).None())

	enumValues =
		append(enumValues, NumberFieldSymbolPosition(0).InsideNumField())

	enumValues =
		append(enumValues, NumberFieldSymbolPosition(0).OutsideNumField())

	intValues =
		append(intValues, NumFieldSymPos.None().XValueInt())

	intValues =
		append(intValues, NumFieldSymPos.InsideNumField().XValueInt())

	intValues =
		append(intValues, NumFieldSymPos.OutsideNumField().XValueInt())

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

func TestNumberFieldSymbolPosition_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberFieldSymbolPosition_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumberFieldSymbolPositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numFieldSymbolPos1, numFieldSymbolPos2,
		numFieldSymbolPos3, numFieldSymbolPos4,
		numFieldSymbolPos5, numFieldSymbolPos6 NumberFieldSymbolPosition

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numFieldSymbolPos1 = enumValues[i]

		isValid = numFieldSymbolPos1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumberFieldSymbolPosition1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"numFieldSymbolPos1 string value  = '%v'\n"+
					"numFieldSymbolPos1 integer value = '%v'\n",
					ePrefix.String(),
					numFieldSymbolPos1.String(),
					numFieldSymbolPos1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numFieldSymbolPos1 string value  = '%v'\n"+
				"numFieldSymbolPos1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numFieldSymbolPos1.String(),
				numFieldSymbolPos1.XValueInt())

			return

		}

		numFieldSymbolPos2,
			err = numFieldSymbolPos1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numFieldSymbolPos1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numFieldSymbolPos1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numFieldSymbolPos1.String(),
				err.Error())

			return
		}

		if numFieldSymbolPos2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numFieldSymbolPos2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numFieldSymbolPos2 string value  = '%v'\n"+
				"numFieldSymbolPos2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numFieldSymbolPos2.String(),
				numFieldSymbolPos2.XValueInt())

			return
		}

		numFieldSymbolPos3 = enumValues[i]

		if numFieldSymbolPos3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numFieldSymbolPos3.XValueInt() != intValues[%v]\n"+
				"numFieldSymbolPos3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numFieldSymbolPos3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numFieldSymbolPos4,
			err = numFieldSymbolPos3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numFieldSymbolPos3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numFieldSymbolPos4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numFieldSymbolPos4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numFieldSymbolPos4 string value  = '%v'\n"+
				"numFieldSymbolPos4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numFieldSymbolPos4.String(),
				numFieldSymbolPos4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numFieldSymbolPos5 = numFieldSymbolPos1.XValue()

		numFieldSymbolPos6 = numFieldSymbolPos2.XValue()

		if numFieldSymbolPos5 != numFieldSymbolPos6 {
			t.Errorf("%v\n"+
				"Error: numFieldSymbolPos5 != numFieldSymbolPos6\n"+
				"numFieldSymbolPos5 = numFieldSymbolPos1.XValue()\n"+
				"numFieldSymbolPos6 = numFieldSymbolPos2.XValue()\n"+
				"numFieldSymbolPos5 string value  = '%v'\n"+
				"numFieldSymbolPos5 integer value = '%v'\n"+
				"numFieldSymbolPos6 string value  = '%v'\n"+
				"numFieldSymbolPos6 integer value = '%v'\n",
				ePrefix.String(),
				numFieldSymbolPos5.String(),
				numFieldSymbolPos5.XValueInt(),
				numFieldSymbolPos6.String(),
				numFieldSymbolPos6.XValueInt())

			return
		}

		_,
			err = numFieldSymbolPos6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numFieldSymbolPos6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numFieldSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numFieldSymbolPos6.String())

			return
		}

		_,
			err = numFieldSymbolPos6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numFieldSymbolPos6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numFieldSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numFieldSymbolPos6.String())

			return
		}

		_,
			err = numFieldSymbolPos6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numFieldSymbolPos6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numFieldSymbolPos6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numFieldSymbolPos6.String())

			return
		}

	}

	return
}

func TestNumberFieldSymbolPosition_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberFieldSymbolPosition_XReturnNoneIfInvalid_000200()",
		"")

	numFieldSymbolPos := NumberFieldSymbolPosition(-972)

	valueNone := numFieldSymbolPos.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumberFieldSymbolPosition(-972)\n"+
			"would return name of 'None' from \n"+
			"numFieldSymbolPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumberFieldSymbolPosition := numFieldSymbolPos.String()

	strNumberFieldSymbolPosition = strings.ToLower(strNumberFieldSymbolPosition)

	if !strings.Contains(strNumberFieldSymbolPosition, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumberFieldSymbolPosition(-972).String()\n"+
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
		NumberFieldSymbolPositionTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numFieldSymbolPos2 NumberFieldSymbolPosition

	numFieldSymbolPos2 = enumValues[1].XReturnNoneIfInvalid()

	if numFieldSymbolPos2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numFieldSymbolPos2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numFieldSymbolPos2 string value  = '%v'\n"+
			"numFieldSymbolPos2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numFieldSymbolPos2.String(),
			numFieldSymbolPos2.XValueInt())
		return
	}

	return
}

func TestNumberFieldSymbolPosition_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberFieldSymbolPosition_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numFieldSymbolPos := NumberFieldSymbolPosition(expectedIntValue)

	actualIntValue := numFieldSymbolPos.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numFieldSymbolPos integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numFieldSymbolPos integer value = '%v'\n"+
			"Actual numFieldSymbolPos integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numFieldSymbolPos.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumberFieldSymbolPosition(-972)\n"+
			"would return name of 'None' from \n"+
			"numFieldSymbolPos.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
