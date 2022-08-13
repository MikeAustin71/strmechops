package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func IntegerGroupingTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []IntegerGroupingType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"IntegerGroupingTypeTestSetup0010()",
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
		append(enumValues, IntegerGroupingType(0).None())

	enumValues =
		append(enumValues, IntegerGroupingType(0).Thousands())

	enumValues =
		append(enumValues, IntegerGroupingType(0).IndiaNumbering())

	enumValues =
		append(enumValues, IntegerGroupingType(0).ChineseNumbering())

	intValues =
		append(intValues, IntGroupingType.None().XValueInt())

	intValues =
		append(intValues, IntGroupingType.Thousands().XValueInt())

	intValues =
		append(intValues, IntGroupingType.IndiaNumbering().XValueInt())

	intValues =
		append(intValues, IntGroupingType.ChineseNumbering().XValueInt())

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

func TestIntegerGroupingType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerGroupingType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		IntegerGroupingTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var integerGroupType1, integerGroupType2,
		integerGroupType3, integerGroupType4,
		integerGroupType5, integerGroupType6 IntegerGroupingType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		integerGroupType1 = enumValues[i]

		isValid = integerGroupType1.XIsValid()

		if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"integerGroupType1 string value  = '%v'\n"+
				"integerGroupType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				integerGroupType1.String(),
				integerGroupType1.XValueInt())

			return

		}

		integerGroupType2,
			err = integerGroupType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  integerGroupType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"integerGroupType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				integerGroupType1.String(),
				err.Error())

			return
		}

		if integerGroupType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"integerGroupType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"integerGroupType2 string value  = '%v'\n"+
				"integerGroupType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				integerGroupType2.String(),
				integerGroupType2.XValueInt())

			return
		}

		integerGroupType3 = enumValues[i]

		if integerGroupType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: integerGroupType3.XValueInt() != intValues[%v]\n"+
				"integerGroupType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				integerGroupType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		integerGroupType4,
			err = integerGroupType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by integerGroupType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if integerGroupType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: integerGroupType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"integerGroupType4 string value  = '%v'\n"+
				"integerGroupType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				integerGroupType4.String(),
				integerGroupType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		integerGroupType5 = integerGroupType1.XValue()

		integerGroupType6 = integerGroupType2.XValue()

		if integerGroupType5 != integerGroupType6 {
			t.Errorf("%v\n"+
				"Error: integerGroupType5 != integerGroupType6\n"+
				"integerGroupType5 = integerGroupType1.XValue()\n"+
				"integerGroupType6 = integerGroupType2.XValue()\n"+
				"integerGroupType5 string value  = '%v'\n"+
				"integerGroupType5 integer value = '%v'\n"+
				"integerGroupType6 string value  = '%v'\n"+
				"integerGroupType6 integer value = '%v'\n",
				ePrefix.String(),
				integerGroupType5.String(),
				integerGroupType5.XValueInt(),
				integerGroupType6.String(),
				integerGroupType6.XValueInt())

			return
		}

		_,
			err = integerGroupType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerGroupType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerGroupType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerGroupType6.String())

			return
		}

		_,
			err = integerGroupType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerGroupType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerGroupType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerGroupType6.String())

			return
		}

		_,
			err = integerGroupType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from integerGroupType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"integerGroupType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				integerGroupType6.String())

			return
		}

	}

	return
}

func TestIntegerGroupingType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerGroupingType_XReturnNoneIfInvalid_000200()",
		"")

	integerGroupType := IntegerGroupingType(-972)

	valueNone := integerGroupType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected IntegerGroupingType(-972)\n"+
			"would return name of 'None' from \n"+
			"integerGroupType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strIntegerGroupingType := integerGroupType.String()

	strIntegerGroupingType = strings.ToLower(strIntegerGroupingType)

	if !strings.Contains(strIntegerGroupingType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected IntegerGroupingType(-972).String()\n"+
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
		IntegerGroupingTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var integerGroupType2 IntegerGroupingType

	integerGroupType2 = enumValues[1].XReturnNoneIfInvalid()

	if integerGroupType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: integerGroupType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"integerGroupType2 string value  = '%v'\n"+
			"integerGroupType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			integerGroupType2.String(),
			integerGroupType2.XValueInt())
		return
	}

	return
}

func TestIntegerGroupingType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestIntegerGroupingType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	integerGroupType := IntegerGroupingType(expectedIntValue)

	actualIntValue := integerGroupType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected integerGroupType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected integerGroupType integer value = '%v'\n"+
			"Actual integerGroupType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := integerGroupType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected IntegerGroupingType(-972)\n"+
			"would return name of 'None' from \n"+
			"integerGroupType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
