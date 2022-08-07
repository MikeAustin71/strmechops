package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumericSignValueTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumericSignValueType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumericSignValueTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Negative",
		"Zero",
		"Positive",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumericSignValueType(0).None())

	enumValues =
		append(enumValues, NumericSignValueType(0).Negative())

	enumValues =
		append(enumValues, NumericSignValueType(0).Zero())

	enumValues =
		append(enumValues, NumericSignValueType(0).Positive())

	intValues =
		append(intValues, NumSignVal.None().XValueInt())

	intValues =
		append(intValues, NumSignVal.Negative().XValueInt())

	intValues =
		append(intValues, NumSignVal.Zero().XValueInt())

	intValues =
		append(intValues, NumSignVal.Positive().XValueInt())

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

func TestNumericSignValueType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSignValueType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumericSignValueTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numericSignValueType1, numericSignValueType2,
		numericSignValueType3, numericSignValueType4,
		numericSignValueType5, numericSignValueType6 NumericSignValueType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numericSignValueType1 = enumValues[i]

		isValid = numericSignValueType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumericSignValueType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumericSignValueType1 string value  = '%v'\n"+
					"NumericSignValueType1 integer value = '%v'\n",
					ePrefix.String(),
					numericSignValueType1.String(),
					numericSignValueType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numericSignValueType1 string value  = '%v'\n"+
				"numericSignValueType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numericSignValueType1.String(),
				numericSignValueType1.XValueInt())

			return

		}

		numericSignValueType2,
			err = numericSignValueType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numericSignValueType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numericSignValueType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSignValueType1.String(),
				err.Error())

			return
		}

		if numericSignValueType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numericSignValueType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numericSignValueType2 string value  = '%v'\n"+
				"numericSignValueType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSignValueType2.String(),
				numericSignValueType2.XValueInt())

			return
		}

		numericSignValueType3 = enumValues[i]

		if numericSignValueType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSignValueType3.XValueInt() != intValues[%v]\n"+
				"numericSignValueType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numericSignValueType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numericSignValueType4,
			err = numericSignValueType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numericSignValueType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numericSignValueType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSignValueType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numericSignValueType4 string value  = '%v'\n"+
				"numericSignValueType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numericSignValueType4.String(),
				numericSignValueType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numericSignValueType5 = numericSignValueType1.XValue()

		numericSignValueType6 = numericSignValueType2.XValue()

		if numericSignValueType5 != numericSignValueType6 {
			t.Errorf("%v\n"+
				"Error: numericSignValueType5 != numericSignValueType6\n"+
				"numericSignValueType5 = numericSignValueType1.XValue()\n"+
				"numericSignValueType6 = numericSignValueType2.XValue()\n"+
				"numericSignValueType5 string value  = '%v'\n"+
				"numericSignValueType5 integer value = '%v'\n"+
				"numericSignValueType6 string value  = '%v'\n"+
				"numericSignValueType6 integer value = '%v'\n",
				ePrefix.String(),
				numericSignValueType5.String(),
				numericSignValueType5.XValueInt(),
				numericSignValueType6.String(),
				numericSignValueType6.XValueInt())

			return
		}

		_,
			err = numericSignValueType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSignValueType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSignValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSignValueType6.String())

			return
		}

		_,
			err = numericSignValueType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSignValueType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSignValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSignValueType6.String())

			return
		}

		_,
			err = numericSignValueType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSignValueType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSignValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSignValueType6.String())

			return
		}

	}

	return
}

func TestNumericSignValueType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSignValueType_XReturnNoneIfInvalid_000200()",
		"")

	numericSignValueType := NumericSignValueType(-972)

	valueNone := numericSignValueType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSignValueType(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSignValueType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumericSignValueType := numericSignValueType.String()

	strNumericSignValueType = strings.ToLower(strNumericSignValueType)

	if !strings.Contains(strNumericSignValueType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumericSignValueType(-972).String()\n"+
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
		NumericSignValueTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numericSignValueType2 NumericSignValueType

	numericSignValueType2 = enumValues[1].XReturnNoneIfInvalid()

	if numericSignValueType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numericSignValueType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numericSignValueType2 string value  = '%v'\n"+
			"numericSignValueType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numericSignValueType2.String(),
			numericSignValueType2.XValueInt())
		return
	}

	return
}

func TestNumericSignValueType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSignValueType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numericSignValueType := NumericSignValueType(expectedIntValue)

	actualIntValue := numericSignValueType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numericSignValueType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numericSignValueType integer value = '%v'\n"+
			"Actual numericSignValueType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numericSignValueType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSignValueType(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSignValueType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
