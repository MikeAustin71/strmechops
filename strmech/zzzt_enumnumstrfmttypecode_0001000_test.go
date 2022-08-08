package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumStrFormatTypeCodeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumStrFormatTypeCode,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatTypeCodeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"AbsoluteValue",
		"Binary",
		"Currency",
		"Hexadecimal",
		"Octal",
		"SignedNumber",
		"ScientificNotation",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).None())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).AbsoluteValue())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).Binary())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).Currency())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).Hexadecimal())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).Octal())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).SignedNumber())

	enumValues =
		append(enumValues, NumStrFormatTypeCode(0).ScientificNotation())

	intValues =
		append(intValues, NumStrFmtType.None().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.AbsoluteValue().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.Binary().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.Currency().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.Hexadecimal().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.Octal().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.SignedNumber().XValueInt())

	intValues =
		append(intValues, NumStrFmtType.ScientificNotation().XValueInt())

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

func TestNumStrFormatTypeCode_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrFormatTypeCode_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumStrFormatTypeCodeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numStrFormatTypeCode1, numStrFormatTypeCode2,
		numStrFormatTypeCode3, numStrFormatTypeCode4,
		numStrFormatTypeCode5, numStrFormatTypeCode6 NumStrFormatTypeCode

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numStrFormatTypeCode1 = enumValues[i]

		isValid = numStrFormatTypeCode1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumStrFormatTypeCode1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"numStrFormatTypeCode1 string value  = '%v'\n"+
					"numStrFormatTypeCode1 integer value = '%v'\n",
					ePrefix.String(),
					numStrFormatTypeCode1.String(),
					numStrFormatTypeCode1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numStrFormatTypeCode1 string value  = '%v'\n"+
				"numStrFormatTypeCode1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numStrFormatTypeCode1.String(),
				numStrFormatTypeCode1.XValueInt())

			return

		}

		numStrFormatTypeCode2,
			err = numStrFormatTypeCode1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numStrFormatTypeCode1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numStrFormatTypeCode1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numStrFormatTypeCode1.String(),
				err.Error())

			return
		}

		if numStrFormatTypeCode2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numStrFormatTypeCode2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numStrFormatTypeCode2 string value  = '%v'\n"+
				"numStrFormatTypeCode2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numStrFormatTypeCode2.String(),
				numStrFormatTypeCode2.XValueInt())

			return
		}

		numStrFormatTypeCode3 = enumValues[i]

		if numStrFormatTypeCode3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numStrFormatTypeCode3.XValueInt() != intValues[%v]\n"+
				"numStrFormatTypeCode3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numStrFormatTypeCode3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numStrFormatTypeCode4,
			err = numStrFormatTypeCode3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numStrFormatTypeCode3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numStrFormatTypeCode4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numStrFormatTypeCode4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numStrFormatTypeCode4 string value  = '%v'\n"+
				"numStrFormatTypeCode4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numStrFormatTypeCode4.String(),
				numStrFormatTypeCode4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numStrFormatTypeCode5 = numStrFormatTypeCode1.XValue()

		numStrFormatTypeCode6 = numStrFormatTypeCode2.XValue()

		if numStrFormatTypeCode5 != numStrFormatTypeCode6 {
			t.Errorf("%v\n"+
				"Error: numStrFormatTypeCode5 != numStrFormatTypeCode6\n"+
				"numStrFormatTypeCode5 = numStrFormatTypeCode1.XValue()\n"+
				"numStrFormatTypeCode6 = numStrFormatTypeCode2.XValue()\n"+
				"numStrFormatTypeCode5 string value  = '%v'\n"+
				"numStrFormatTypeCode5 integer value = '%v'\n"+
				"numStrFormatTypeCode6 string value  = '%v'\n"+
				"numStrFormatTypeCode6 integer value = '%v'\n",
				ePrefix.String(),
				numStrFormatTypeCode5.String(),
				numStrFormatTypeCode5.XValueInt(),
				numStrFormatTypeCode6.String(),
				numStrFormatTypeCode6.XValueInt())

			return
		}

		_,
			err = numStrFormatTypeCode6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numStrFormatTypeCode6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numStrFormatTypeCode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numStrFormatTypeCode6.String())

			return
		}

		_,
			err = numStrFormatTypeCode6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numStrFormatTypeCode6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numStrFormatTypeCode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numStrFormatTypeCode6.String())

			return
		}

		_,
			err = numStrFormatTypeCode6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numStrFormatTypeCode6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numStrFormatTypeCode6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numStrFormatTypeCode6.String())

			return
		}

	}

	return
}

func TestNumStrFormatTypeCode_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrFormatTypeCode_XReturnNoneIfInvalid_000200()",
		"")

	numStrFormatTypeCode := NumStrFormatTypeCode(-972)

	valueNone := numStrFormatTypeCode.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumStrFormatTypeCode(-972)\n"+
			"would return name of 'None' from \n"+
			"numStrFormatTypeCode.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumStrFormatTypeCode := numStrFormatTypeCode.String()

	strNumStrFormatTypeCode = strings.ToLower(strNumStrFormatTypeCode)

	if !strings.Contains(strNumStrFormatTypeCode, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumStrFormatTypeCode(-972).String()\n"+
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
		NumStrFormatTypeCodeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numStrFormatTypeCode2 NumStrFormatTypeCode

	numStrFormatTypeCode2 = enumValues[1].XReturnNoneIfInvalid()

	if numStrFormatTypeCode2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numStrFormatTypeCode2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numStrFormatTypeCode2 string value  = '%v'\n"+
			"numStrFormatTypeCode2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numStrFormatTypeCode2.String(),
			numStrFormatTypeCode2.XValueInt())
		return
	}

	return
}

func TestNumStrFormatTypeCode_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrFormatTypeCode_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numStrFormatTypeCode := NumStrFormatTypeCode(expectedIntValue)

	actualIntValue := numStrFormatTypeCode.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numStrFormatTypeCode integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numStrFormatTypeCode integer value = '%v'\n"+
			"Actual numStrFormatTypeCode integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numStrFormatTypeCode.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumStrFormatTypeCode(-972)\n"+
			"would return name of 'None' from \n"+
			"numStrFormatTypeCode.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
