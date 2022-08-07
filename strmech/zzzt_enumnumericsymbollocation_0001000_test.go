package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumericSymbolLocationTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumericSymbolLocation,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumericSymbolLocationTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Before",
		"Interior",
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
		append(enumValues, NumericSymbolLocation(0).None())

	enumValues =
		append(enumValues, NumericSymbolLocation(0).Before())

	enumValues =
		append(enumValues, NumericSymbolLocation(0).Interior())

	enumValues =
		append(enumValues, NumericSymbolLocation(0).After())

	enumValues =
		append(enumValues, NumericSymbolLocation(0).BeforeAndAfter())

	intValues =
		append(intValues, NumSymLocation.None().XValueInt())

	intValues =
		append(intValues, NumSymLocation.Before().XValueInt())

	intValues =
		append(intValues, NumSymLocation.Interior().XValueInt())

	intValues =
		append(intValues, NumSymLocation.After().XValueInt())

	intValues =
		append(intValues, NumSymLocation.BeforeAndAfter().XValueInt())

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

func TestNumericSymbolLocation_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolLocation_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumericSymbolLocationTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numericSymbolLocation1, numericSymbolLocation2,
		numericSymbolLocation3, numericSymbolLocation4,
		numericSymbolLocation5, numericSymbolLocation6 NumericSymbolLocation

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numericSymbolLocation1 = enumValues[i]

		isValid = numericSymbolLocation1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumericSymbolLocation1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumericSymbolLocation1 string value  = '%v'\n"+
					"NumericSymbolLocation1 integer value = '%v'\n",
					ePrefix.String(),
					numericSymbolLocation1.String(),
					numericSymbolLocation1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numericSymbolLocation1 string value  = '%v'\n"+
				"numericSymbolLocation1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numericSymbolLocation1.String(),
				numericSymbolLocation1.XValueInt())

			return

		}

		numericSymbolLocation2,
			err = numericSymbolLocation1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numericSymbolLocation1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numericSymbolLocation1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSymbolLocation1.String(),
				err.Error())

			return
		}

		if numericSymbolLocation2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numericSymbolLocation2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numericSymbolLocation2 string value  = '%v'\n"+
				"numericSymbolLocation2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSymbolLocation2.String(),
				numericSymbolLocation2.XValueInt())

			return
		}

		numericSymbolLocation3 = enumValues[i]

		if numericSymbolLocation3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSymbolLocation3.XValueInt() != intValues[%v]\n"+
				"numericSymbolLocation3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolLocation3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numericSymbolLocation4,
			err = numericSymbolLocation3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numericSymbolLocation3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numericSymbolLocation4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSymbolLocation4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numericSymbolLocation4 string value  = '%v'\n"+
				"numericSymbolLocation4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numericSymbolLocation4.String(),
				numericSymbolLocation4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numericSymbolLocation5 = numericSymbolLocation1.XValue()

		numericSymbolLocation6 = numericSymbolLocation2.XValue()

		if numericSymbolLocation5 != numericSymbolLocation6 {
			t.Errorf("%v\n"+
				"Error: numericSymbolLocation5 != numericSymbolLocation6\n"+
				"numericSymbolLocation5 = numericSymbolLocation1.XValue()\n"+
				"numericSymbolLocation6 = numericSymbolLocation2.XValue()\n"+
				"numericSymbolLocation5 string value  = '%v'\n"+
				"numericSymbolLocation5 integer value = '%v'\n"+
				"numericSymbolLocation6 string value  = '%v'\n"+
				"numericSymbolLocation6 integer value = '%v'\n",
				ePrefix.String(),
				numericSymbolLocation5.String(),
				numericSymbolLocation5.XValueInt(),
				numericSymbolLocation6.String(),
				numericSymbolLocation6.XValueInt())

			return
		}

		_,
			err = numericSymbolLocation6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolLocation6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolLocation6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolLocation6.String())

			return
		}

		_,
			err = numericSymbolLocation6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolLocation6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolLocation6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolLocation6.String())

			return
		}

		_,
			err = numericSymbolLocation6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolLocation6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolLocation6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolLocation6.String())

			return
		}

	}

	return
}

func TestNumericSymbolLocation_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolLocation_XReturnNoneIfInvalid_000200()",
		"")

	numericSymbolLocation := NumericSymbolLocation(-972)

	valueNone := numericSymbolLocation.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolLocation(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSymbolLocation.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumericSymbolLocation := numericSymbolLocation.String()

	strNumericSymbolLocation = strings.ToLower(strNumericSymbolLocation)

	if !strings.Contains(strNumericSymbolLocation, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolLocation(-972).String()\n"+
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
		NumericSymbolLocationTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numericSymbolLocation2 NumericSymbolLocation

	numericSymbolLocation2 = enumValues[1].XReturnNoneIfInvalid()

	if numericSymbolLocation2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numericSymbolLocation2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numericSymbolLocation2 string value  = '%v'\n"+
			"numericSymbolLocation2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numericSymbolLocation2.String(),
			numericSymbolLocation2.XValueInt())
		return
	}

	return
}

func TestNumericSymbolLocation_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolLocation_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numericSymbolLocation := NumericSymbolLocation(expectedIntValue)

	actualIntValue := numericSymbolLocation.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numericSymbolLocation integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numericSymbolLocation integer value = '%v'\n"+
			"Actual numericSymbolLocation integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numericSymbolLocation.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolLocation(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSymbolLocation.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
