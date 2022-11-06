package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumericValueTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumericValueType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumericValueTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"FloatingPoint",
		"Integer",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumericValueType(0).None())

	enumValues =
		append(enumValues, NumericValueType(0).FloatingPoint())

	enumValues =
		append(enumValues, NumericValueType(0).Integer())

	intValues =
		append(intValues, NumValType.None().XValueInt())

	intValues =
		append(intValues, NumValType.FloatingPoint().XValueInt())

	intValues =
		append(intValues, NumValType.Integer().XValueInt())

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

func TestNumericValueType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericValueType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumericValueTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numericValueType1, numericValueType2,
		numericValueType3, numericValueType4,
		numericValueType5, numericValueType6 NumericValueType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numericValueType1 = enumValues[i]

		isValid = numericValueType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumericValueType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumericValueType1 string value  = '%v'\n"+
					"NumericValueType1 integer value = '%v'\n",
					ePrefix.String(),
					numericValueType1.String(),
					numericValueType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numericValueType1 string value  = '%v'\n"+
				"numericValueType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numericValueType1.String(),
				numericValueType1.XValueInt())

			return

		}

		numericValueType2,
			err = numericValueType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numericValueType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numericValueType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericValueType1.String(),
				err.Error())

			return
		}

		if numericValueType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numericValueType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numericValueType2 string value  = '%v'\n"+
				"numericValueType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericValueType2.String(),
				numericValueType2.XValueInt())

			return
		}

		numericValueType3 = enumValues[i]

		if numericValueType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numericValueType3.XValueInt() != intValues[%v]\n"+
				"numericValueType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numericValueType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numericValueType4,
			err = numericValueType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numericValueType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numericValueType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numericValueType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numericValueType4 string value  = '%v'\n"+
				"numericValueType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numericValueType4.String(),
				numericValueType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numericValueType5 = numericValueType1.XValue()

		numericValueType6 = numericValueType2.XValue()

		if numericValueType5 != numericValueType6 {
			t.Errorf("%v\n"+
				"Error: numericValueType5 != numericValueType6\n"+
				"numericValueType5 = numericValueType1.XValue()\n"+
				"numericValueType6 = numericValueType2.XValue()\n"+
				"numericValueType5 string value  = '%v'\n"+
				"numericValueType5 integer value = '%v'\n"+
				"numericValueType6 string value  = '%v'\n"+
				"numericValueType6 integer value = '%v'\n",
				ePrefix.String(),
				numericValueType5.String(),
				numericValueType5.XValueInt(),
				numericValueType6.String(),
				numericValueType6.XValueInt())

			return
		}

		_,
			err = numericValueType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericValueType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericValueType6.String())

			return
		}

		_,
			err = numericValueType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericValueType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericValueType6.String())

			return
		}

		_,
			err = numericValueType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericValueType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericValueType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericValueType6.String())

			return
		}

	}

	return
}

func TestNumericValueType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericValueType_XReturnNoneIfInvalid_000200()",
		"")

	numericValueType := NumericValueType(-972)

	valueNone := numericValueType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericValueType(-972)\n"+
			"would return name of 'None' from \n"+
			"numberValueType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumericValueType := numericValueType.String()

	strNumericValueType = strings.ToLower(strNumericValueType)

	if !strings.Contains(strNumericValueType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumericValueType(-972).String()\n"+
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
		NumericValueTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numericValueType2 NumericValueType

	numericValueType2 = enumValues[1].XReturnNoneIfInvalid()

	if numericValueType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numericValueType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numericValueType2 string value  = '%v'\n"+
			"numericValueType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numericValueType2.String(),
			numericValueType2.XValueInt())
		return
	}

	return
}

func TestNumericValueType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericValueType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numericValueType := NumericValueType(expectedIntValue)

	actualIntValue := numericValueType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numberValueType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numberValueType integer value = '%v'\n"+
			"Actual numberValueType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numericValueType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericValueType(-972)\n"+
			"would return name of 'None' from \n"+
			"numberValueType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
