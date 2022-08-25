package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func CarryResultTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []CarryResultType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CarryResultTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"MinusOne",
		"Zero",
		"PlusOne",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, CarryResultType(0).MinusOne())

	enumValues =
		append(enumValues, CarryResultType(0).Zero())

	enumValues =
		append(enumValues, CarryResultType(0).PlusOne())

	intValues =
		append(intValues, CarryType.MinusOne().XValueInt())

	intValues =
		append(intValues, CarryType.Zero().XValueInt())

	intValues =
		append(intValues, CarryType.PlusOne().XValueInt())

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

func TestCarryResultType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCarryResultType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		CarryResultTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var carryResultType1, carryResultType2,
		carryResultType3, carryResultType4,
		carryResultType5, carryResultType6 CarryResultType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		carryResultType1 = enumValues[i]

		isValid = carryResultType1.XIsValid()

		if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"carryResultType1 string value  = '%v'\n"+
				"carryResultType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				carryResultType1.String(),
				carryResultType1.XValueInt())

			return

		}

		carryResultType2,
			err = carryResultType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  carryResultType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"carryResultType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				carryResultType1.String(),
				err.Error())

			return
		}

		if carryResultType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"carryResultType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"carryResultType2 string value  = '%v'\n"+
				"carryResultType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				carryResultType2.String(),
				carryResultType2.XValueInt())

			return
		}

		carryResultType3 = enumValues[i]

		if carryResultType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: carryResultType3.XValueInt() != intValues[%v]\n"+
				"carryResultType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				carryResultType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		carryResultType4,
			err = carryResultType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by carryResultType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if carryResultType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: carryResultType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"carryResultType4 string value  = '%v'\n"+
				"carryResultType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				carryResultType4.String(),
				carryResultType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		carryResultType5 = carryResultType1.XValue()

		carryResultType6 = carryResultType2.XValue()

		if carryResultType5 != carryResultType6 {
			t.Errorf("%v\n"+
				"Error: carryResultType5 != carryResultType6\n"+
				"carryResultType5 = carryResultType1.XValue()\n"+
				"carryResultType6 = carryResultType2.XValue()\n"+
				"carryResultType5 string value  = '%v'\n"+
				"carryResultType5 integer value = '%v'\n"+
				"carryResultType6 string value  = '%v'\n"+
				"carryResultType6 integer value = '%v'\n",
				ePrefix.String(),
				carryResultType5.String(),
				carryResultType5.XValueInt(),
				carryResultType6.String(),
				carryResultType6.XValueInt())

			return
		}

		_,
			err = carryResultType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from carryResultType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"carryResultType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				carryResultType6.String())

			return
		}

		_,
			err = carryResultType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from carryResultType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"carryResultType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				carryResultType6.String())

			return
		}

		_,
			err = carryResultType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from carryResultType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"carryResultType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				carryResultType6.String())

			return
		}

	}

	return
}

func TestCarryResultType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCarryResultType_XReturnNoneIfInvalid_000200()",
		"")

	carryResultType := CarryResultType(-972)

	valueNone := carryResultType.XReturnNoneIfInvalid()

	if valueNone.String() != "Zero" {

		t.Errorf("%v\n"+
			"Error: Expected CarryResultType(-972)\n"+
			"would return name of 'Zero' from \n"+
			"carryResultType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strCarryResultType := carryResultType.String()

	strCarryResultType = strings.ToLower(strCarryResultType)

	if !strings.Contains(strCarryResultType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected CarryResultType(-972).String()\n"+
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
		CarryResultTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var carryResultType2 CarryResultType

	carryResultType2 = enumValues[1].XReturnNoneIfInvalid()

	if carryResultType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: carryResultType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"carryResultType2 string value  = '%v'\n"+
			"carryResultType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			carryResultType2.String(),
			carryResultType2.XValueInt())
		return
	}

	return
}

func TestCarryResultType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCarryResultType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	carryResultType := CarryResultType(expectedIntValue)

	actualIntValue := carryResultType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected carryResultType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected carryResultType integer value = '%v'\n"+
			"Actual carryResultType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := carryResultType.XReturnNoneIfInvalid()

	if strName.String() != "Zero" {

		t.Errorf("%v\n"+
			"Error: Expected CarryResultType(-972)\n"+
			"would return name of 'Zero' from \n"+
			"carryResultType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
