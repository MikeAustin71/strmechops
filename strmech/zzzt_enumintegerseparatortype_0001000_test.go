package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func IntegerSeparatorTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []IntegerSeparatorType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerSeparatorTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Thousands",
		"IndiaNumbering",
		"ChineseNumbering",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, IntegerSeparatorType(0).None())

	enumValues =
		append(enumValues, IntegerSeparatorType(0).Thousands())

	enumValues =
		append(enumValues, IntegerSeparatorType(0).IndiaNumbering())

	enumValues =
		append(enumValues, IntegerSeparatorType(0).ChineseNumbering())

	intValues =
		append(intValues, IntSeparatorType.None().XValueInt())

	intValues =
		append(intValues, IntSeparatorType.Thousands().XValueInt())

	intValues =
		append(intValues, IntSeparatorType.IndiaNumbering().XValueInt())

	intValues =
		append(intValues, IntSeparatorType.ChineseNumbering().XValueInt())

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

func TestIntegerSeparatorType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerSeparatorType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		IntegerSeparatorTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var integerSepType1, integerSepType2,
		integerSepType3, integerSepType4,
		integerSepType5, integerSepType6 IntegerSeparatorType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		integerSepType1 = enumValues[i]

		isValid = integerSepType1.XIsValid()

		if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"integerSepType1 string value  = '%v'\n"+
				"integerSepType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				integerSepType1.String(),
				integerSepType1.XValueInt())

			return

		}

		integerSepType2,
			err = integerSepType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  integerSepType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"integerSepType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				integerSepType1.String(),
				err.Error())

			return
		}

		if integerSepType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"integerSepType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"integerSepType2 string value  = '%v'\n"+
				"integerSepType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				integerSepType2.String(),
				integerSepType2.XValueInt())

			return
		}

		integerSepType3 = enumValues[i]

		if integerSepType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: integerSepType3.XValueInt() != intValues[%v]\n"+
				"integerSepType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				integerSepType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		integerSepType4,
			err = integerSepType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by integerSepType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if integerSepType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: integerSepType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"integerSepType4 string value  = '%v'\n"+
				"integerSepType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				integerSepType4.String(),
				integerSepType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		integerSepType5 = integerSepType1.XValue()

		integerSepType6 = integerSepType2.XValue()

		if integerSepType5 != integerSepType6 {
			t.Errorf("%v\n"+
				"Error: integerSepType5 != integerSepType6\n"+
				"integerSepType5 = integerSepType1.XValue()\n"+
				"integerSepType6 = integerSepType2.XValue()\n"+
				"integerSepType5 string value  = '%v'\n"+
				"integerSepType5 integer value = '%v'\n"+
				"integerSepType6 string value  = '%v'\n"+
				"integerSepType6 integer value = '%v'\n",
				ePrefix.String(),
				integerSepType5.String(),
				integerSepType5.XValueInt(),
				integerSepType6.String(),
				integerSepType6.XValueInt())

			return
		}

		_,
			err = integerSepType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerSepType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerSepType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerSepType6.String())

			return
		}

		_,
			err = integerSepType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerSepType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerSepType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerSepType6.String())

			return
		}

		_,
			err = integerSepType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerSepType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerSepType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerSepType6.String())

			return
		}

	}

	return
}

func TestIntegerSeparatorType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerSeparatorType_XReturnNoneIfInvalid_000200()",
		"")

	integerSepType := IntegerSeparatorType(-972)

	valueNone := integerSepType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected IntegerSeparatorType(-972)\n"+
			"would return name of 'None' from \n"+
			"integerSepType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strIntegerSeparatorType := integerSepType.String()

	strIntegerSeparatorType = strings.ToLower(strIntegerSeparatorType)

	if !strings.Contains(strIntegerSeparatorType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected IntegerSeparatorType(-972).String()\n"+
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
		IntegerSeparatorTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var integerSepType2 IntegerSeparatorType

	integerSepType2 = enumValues[1].XReturnNoneIfInvalid()

	if integerSepType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: integerSepType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"integerSepType2 string value  = '%v'\n"+
			"integerSepType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			integerSepType2.String(),
			integerSepType2.XValueInt())
		return
	}

	return
}

func TestIntegerSeparatorType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerSeparatorType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	integerSepType := IntegerSeparatorType(expectedIntValue)

	actualIntValue := integerSepType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected integerSepType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected integerSepType integer value = '%v'\n"+
			"Actual integerSepType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := integerSepType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected IntegerSeparatorType(-972)\n"+
			"would return name of 'None' from \n"+
			"integerSepType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
