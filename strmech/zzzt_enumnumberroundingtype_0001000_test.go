package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumberRoundingTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumberRoundingType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberRoundingTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"HalfUpWithNegNums",
		"HalfDownWithNegNums",
		"HalfAwayFromZero",
		"HalfTowardsZero",
		"HalfToEven",
		"HalfToOdd",
		"Randomly",
		"Floor",
		"Ceiling",
		"Truncate",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumberRoundingType(0).None())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfUpWithNegNums())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfDownWithNegNums())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfAwayFromZero())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfTowardsZero())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfToEven())

	enumValues =
		append(enumValues, NumberRoundingType(0).HalfToOdd())

	enumValues =
		append(enumValues, NumberRoundingType(0).Randomly())

	enumValues =
		append(enumValues, NumberRoundingType(0).Floor())

	enumValues =
		append(enumValues, NumberRoundingType(0).Ceiling())

	enumValues =
		append(enumValues, NumberRoundingType(0).Truncate())

	intValues =
		append(intValues, NumRoundType.None().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfUpWithNegNums().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfDownWithNegNums().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfAwayFromZero().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfTowardsZero().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfToEven().XValueInt())

	intValues =
		append(intValues, NumRoundType.HalfToOdd().XValueInt())

	intValues =
		append(intValues, NumRoundType.Randomly().XValueInt())

	intValues =
		append(intValues, NumRoundType.Floor().XValueInt())

	intValues =
		append(intValues, NumRoundType.Ceiling().XValueInt())

	intValues =
		append(intValues, NumRoundType.Truncate().XValueInt())

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

func TestNumberRoundingType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberRoundingType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumberRoundingTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numberRoundingType1, numberRoundingType2,
		numberRoundingType3, numberRoundingType4,
		numberRoundingType5, numberRoundingType6 NumberRoundingType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numberRoundingType1 = enumValues[i]

		isValid = numberRoundingType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumberRoundingType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumberRoundingType1 string value  = '%v'\n"+
					"NumberRoundingType1 integer value = '%v'\n",
					ePrefix.String(),
					numberRoundingType1.String(),
					numberRoundingType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numberRoundingType1 string value  = '%v'\n"+
				"numberRoundingType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numberRoundingType1.String(),
				numberRoundingType1.XValueInt())

			return

		}

		numberRoundingType2,
			err = numberRoundingType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numberRoundingType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numberRoundingType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numberRoundingType1.String(),
				err.Error())

			return
		}

		if numberRoundingType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numberRoundingType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numberRoundingType2 string value  = '%v'\n"+
				"numberRoundingType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numberRoundingType2.String(),
				numberRoundingType2.XValueInt())

			return
		}

		numberRoundingType3 = enumValues[i]

		if numberRoundingType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numberRoundingType3.XValueInt() != intValues[%v]\n"+
				"numberRoundingType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numberRoundingType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numberRoundingType4,
			err = numberRoundingType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numberRoundingType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numberRoundingType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numberRoundingType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numberRoundingType4 string value  = '%v'\n"+
				"numberRoundingType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numberRoundingType4.String(),
				numberRoundingType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numberRoundingType5 = numberRoundingType1.XValue()

		numberRoundingType6 = numberRoundingType2.XValue()

		if numberRoundingType5 != numberRoundingType6 {
			t.Errorf("%v\n"+
				"Error: numberRoundingType5 != numberRoundingType6\n"+
				"numberRoundingType5 = numberRoundingType1.XValue()\n"+
				"numberRoundingType6 = numberRoundingType2.XValue()\n"+
				"numberRoundingType5 string value  = '%v'\n"+
				"numberRoundingType5 integer value = '%v'\n"+
				"numberRoundingType6 string value  = '%v'\n"+
				"numberRoundingType6 integer value = '%v'\n",
				ePrefix.String(),
				numberRoundingType5.String(),
				numberRoundingType5.XValueInt(),
				numberRoundingType6.String(),
				numberRoundingType6.XValueInt())

			return
		}

		_,
			err = numberRoundingType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numberRoundingType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numberRoundingType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numberRoundingType6.String())

			return
		}

		_,
			err = numberRoundingType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numberRoundingType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numberRoundingType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numberRoundingType6.String())

			return
		}

		_,
			err = numberRoundingType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numberRoundingType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numberRoundingType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numberRoundingType6.String())

			return
		}

	}

	return
}

func TestNumberRoundingType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberRoundingType_XReturnNoneIfInvalid_000200()",
		"")

	numberRoundingType := NumberRoundingType(-972)

	valueNone := numberRoundingType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumberRoundingType(-972)\n"+
			"would return name of 'None' from \n"+
			"numberRoundingType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumberRoundingType := numberRoundingType.String()

	strNumberRoundingType = strings.ToLower(strNumberRoundingType)

	if !strings.Contains(strNumberRoundingType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumberRoundingType(-972).String()\n"+
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
		NumberRoundingTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numberRoundingType2 NumberRoundingType

	numberRoundingType2 = enumValues[1].XReturnNoneIfInvalid()

	if numberRoundingType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numberRoundingType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numberRoundingType2 string value  = '%v'\n"+
			"numberRoundingType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numberRoundingType2.String(),
			numberRoundingType2.XValueInt())
		return
	}

	return
}

func TestNumberRoundingType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberRoundingType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numberRoundingType := NumberRoundingType(expectedIntValue)

	actualIntValue := numberRoundingType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numberRoundingType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numberRoundingType integer value = '%v'\n"+
			"Actual numberRoundingType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numberRoundingType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumberRoundingType(-972)\n"+
			"would return name of 'None' from \n"+
			"numberRoundingType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
