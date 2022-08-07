package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumSignSymbolDisplayModeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumSignSymbolDisplayMode,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumSignSymbolDisplayModeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Explicit",
		"Implicit",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumSignSymbolDisplayMode(0).None())

	enumValues =
		append(enumValues, NumSignSymbolDisplayMode(0).Explicit())

	enumValues =
		append(enumValues, NumSignSymbolDisplayMode(0).Implicit())

	intValues =
		append(intValues, NSignSymDisplayMode.None().XValueInt())

	intValues =
		append(intValues, NSignSymDisplayMode.Explicit().XValueInt())

	intValues =
		append(intValues, NSignSymDisplayMode.Implicit().XValueInt())

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

func TestNumSignSymbolDisplayMode_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolDisplayMode_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumSignSymbolDisplayModeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numSignSymbolDisplayMode1, numSignSymbolDisplayMode2,
		numSignSymbolDisplayMode3, numSignSymbolDisplayMode4,
		numSignSymbolDisplayMode5, numSignSymbolDisplayMode6 NumSignSymbolDisplayMode

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numSignSymbolDisplayMode1 = enumValues[i]

		isValid = numSignSymbolDisplayMode1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumSignSymbolDisplayMode1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumSignSymbolDisplayMode1 string value  = '%v'\n"+
					"NumSignSymbolDisplayMode1 integer value = '%v'\n",
					ePrefix.String(),
					numSignSymbolDisplayMode1.String(),
					numSignSymbolDisplayMode1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numSignSymbolDisplayMode1 string value  = '%v'\n"+
				"numSignSymbolDisplayMode1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numSignSymbolDisplayMode1.String(),
				numSignSymbolDisplayMode1.XValueInt())

			return

		}

		numSignSymbolDisplayMode2,
			err = numSignSymbolDisplayMode1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numSignSymbolDisplayMode1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numSignSymbolDisplayMode1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numSignSymbolDisplayMode1.String(),
				err.Error())

			return
		}

		if numSignSymbolDisplayMode2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numSignSymbolDisplayMode2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numSignSymbolDisplayMode2 string value  = '%v'\n"+
				"numSignSymbolDisplayMode2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numSignSymbolDisplayMode2.String(),
				numSignSymbolDisplayMode2.XValueInt())

			return
		}

		numSignSymbolDisplayMode3 = enumValues[i]

		if numSignSymbolDisplayMode3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numSignSymbolDisplayMode3.XValueInt() != intValues[%v]\n"+
				"numSignSymbolDisplayMode3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolDisplayMode3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numSignSymbolDisplayMode4,
			err = numSignSymbolDisplayMode3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numSignSymbolDisplayMode3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numSignSymbolDisplayMode4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numSignSymbolDisplayMode4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numSignSymbolDisplayMode4 string value  = '%v'\n"+
				"numSignSymbolDisplayMode4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numSignSymbolDisplayMode4.String(),
				numSignSymbolDisplayMode4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numSignSymbolDisplayMode5 = numSignSymbolDisplayMode1.XValue()

		numSignSymbolDisplayMode6 = numSignSymbolDisplayMode2.XValue()

		if numSignSymbolDisplayMode5 != numSignSymbolDisplayMode6 {
			t.Errorf("%v\n"+
				"Error: numSignSymbolDisplayMode5 != numSignSymbolDisplayMode6\n"+
				"numSignSymbolDisplayMode5 = numSignSymbolDisplayMode1.XValue()\n"+
				"numSignSymbolDisplayMode6 = numSignSymbolDisplayMode2.XValue()\n"+
				"numSignSymbolDisplayMode5 string value  = '%v'\n"+
				"numSignSymbolDisplayMode5 integer value = '%v'\n"+
				"numSignSymbolDisplayMode6 string value  = '%v'\n"+
				"numSignSymbolDisplayMode6 integer value = '%v'\n",
				ePrefix.String(),
				numSignSymbolDisplayMode5.String(),
				numSignSymbolDisplayMode5.XValueInt(),
				numSignSymbolDisplayMode6.String(),
				numSignSymbolDisplayMode6.XValueInt())

			return
		}

		_,
			err = numSignSymbolDisplayMode6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolDisplayMode6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolDisplayMode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolDisplayMode6.String())

			return
		}

		_,
			err = numSignSymbolDisplayMode6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolDisplayMode6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolDisplayMode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolDisplayMode6.String())

			return
		}

		_,
			err = numSignSymbolDisplayMode6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numSignSymbolDisplayMode6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numSignSymbolDisplayMode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numSignSymbolDisplayMode6.String())

			return
		}

	}

	return
}

func TestNumSignSymbolDisplayMode_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolDisplayMode_XReturnNoneIfInvalid_000200()",
		"")

	numSignSymbolDisplayMode := NumSignSymbolDisplayMode(-972)

	valueNone := numSignSymbolDisplayMode.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolDisplayMode(-972)\n"+
			"would return name of 'None' from \n"+
			"numSignSymbolDisplayMode.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumSignSymbolDisplayMode := numSignSymbolDisplayMode.String()

	strNumSignSymbolDisplayMode = strings.ToLower(strNumSignSymbolDisplayMode)

	if !strings.Contains(strNumSignSymbolDisplayMode, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolDisplayMode(-972).String()\n"+
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
		NumSignSymbolDisplayModeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numSignSymbolDisplayMode2 NumSignSymbolDisplayMode

	numSignSymbolDisplayMode2 = enumValues[1].XReturnNoneIfInvalid()

	if numSignSymbolDisplayMode2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numSignSymbolDisplayMode2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numSignSymbolDisplayMode2 string value  = '%v'\n"+
			"numSignSymbolDisplayMode2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numSignSymbolDisplayMode2.String(),
			numSignSymbolDisplayMode2.XValueInt())
		return
	}

	return
}

func TestNumSignSymbolDisplayMode_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumSignSymbolDisplayMode_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numSignSymbolDisplayMode := NumSignSymbolDisplayMode(expectedIntValue)

	actualIntValue := numSignSymbolDisplayMode.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numSignSymbolDisplayMode integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numSignSymbolDisplayMode integer value = '%v'\n"+
			"Actual numSignSymbolDisplayMode integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numSignSymbolDisplayMode.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumSignSymbolDisplayMode(-972)\n"+
			"would return name of 'None' from \n"+
			"numSignSymbolDisplayMode.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
