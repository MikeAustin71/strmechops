package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func ScientificNotationFormatTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []ScientificNotationFormat,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"ScientificNotationFormatTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Exponential",
		"ENotation",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, ScientificNotationFormat(0).None())

	enumValues =
		append(enumValues, ScientificNotationFormat(0).Exponential())

	enumValues =
		append(enumValues, ScientificNotationFormat(0).ENotation())

	intValues =
		append(intValues, SciNotFmt.None().XValueInt())

	intValues =
		append(intValues, SciNotFmt.Exponential().XValueInt())

	intValues =
		append(intValues, SciNotFmt.ENotation().XValueInt())

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

func TestScientificNotationFormat_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestScientificNotationFormat_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		ScientificNotationFormatTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var sciNotationFormat1, sciNotationFormat2,
		sciNotationFormat3, sciNotationFormat4,
		sciNotationFormat5, sciNotationFormat6 ScientificNotationFormat

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		sciNotationFormat1 = enumValues[i]

		isValid = sciNotationFormat1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: sciNotationFormat1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"sciNotationFormat1 string value  = '%v'\n"+
					"sciNotationFormat1 integer value = '%v'\n",
					ePrefix.String(),
					sciNotationFormat1.String(),
					sciNotationFormat1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"sciNotationFormat1 string value  = '%v'\n"+
				"sciNotationFormat1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				sciNotationFormat1.String(),
				sciNotationFormat1.XValueInt())

			return

		}

		sciNotationFormat2,
			err = sciNotationFormat1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  sciNotationFormat1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"sciNotationFormat1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				sciNotationFormat1.String(),
				err.Error())

			return
		}

		if sciNotationFormat2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"sciNotationFormat2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"sciNotationFormat2 string value  = '%v'\n"+
				"sciNotationFormat2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				sciNotationFormat2.String(),
				sciNotationFormat2.XValueInt())

			return
		}

		sciNotationFormat3 = enumValues[i]

		if sciNotationFormat3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: sciNotationFormat3.XValueInt() != intValues[%v]\n"+
				"sciNotationFormat3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				sciNotationFormat3.XValueInt(),
				i,
				intValues[i])

			return
		}

		sciNotationFormat4,
			err = sciNotationFormat3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by sciNotationFormat3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if sciNotationFormat4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: sciNotationFormat4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"sciNotationFormat4 string value  = '%v'\n"+
				"sciNotationFormat4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				sciNotationFormat4.String(),
				sciNotationFormat4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		sciNotationFormat5 = sciNotationFormat1.XValue()

		sciNotationFormat6 = sciNotationFormat2.XValue()

		if sciNotationFormat5 != sciNotationFormat6 {
			t.Errorf("%v\n"+
				"Error: sciNotationFormat5 != sciNotationFormat6\n"+
				"sciNotationFormat5 = sciNotationFormat1.XValue()\n"+
				"sciNotationFormat6 = sciNotationFormat2.XValue()\n"+
				"sciNotationFormat5 string value  = '%v'\n"+
				"sciNotationFormat5 integer value = '%v'\n"+
				"sciNotationFormat6 string value  = '%v'\n"+
				"sciNotationFormat6 integer value = '%v'\n",
				ePrefix.String(),
				sciNotationFormat5.String(),
				sciNotationFormat5.XValueInt(),
				sciNotationFormat6.String(),
				sciNotationFormat6.XValueInt())

			return
		}

		_,
			err = sciNotationFormat6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from sciNotationFormat6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"sciNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				sciNotationFormat6.String())

			return
		}

		_,
			err = sciNotationFormat6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from sciNotationFormat6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"sciNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				sciNotationFormat6.String())

			return
		}

		_,
			err = sciNotationFormat6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from sciNotationFormat6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"sciNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				sciNotationFormat6.String())

			return
		}

	}

	return
}

func TestScientificNotationFormat_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestScientificNotationFormat_XReturnNoneIfInvalid_000200()",
		"")

	sciNotationFormat := ScientificNotationFormat(-972)

	valueNone := sciNotationFormat.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected ScientificNotationFormat(-972)\n"+
			"would return name of 'None' from \n"+
			"sciNotationFormat.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strScientificNotationFormat := sciNotationFormat.String()

	strScientificNotationFormat = strings.ToLower(strScientificNotationFormat)

	if !strings.Contains(strScientificNotationFormat, "error") {

		t.Errorf("%v\n"+
			"Error: Expected ScientificNotationFormat(-972).String()\n"+
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
		ScientificNotationFormatTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var sciNotationFormat2 ScientificNotationFormat

	sciNotationFormat2 = enumValues[1].XReturnNoneIfInvalid()

	if sciNotationFormat2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: sciNotationFormat2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"sciNotationFormat2 string value  = '%v'\n"+
			"sciNotationFormat2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			sciNotationFormat2.String(),
			sciNotationFormat2.XValueInt())
		return
	}

	return
}

func TestScientificNotationFormat_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestScientificNotationFormat_XValueInt_000300()",
		"")

	expectedIntValue := -972

	sciNotationFormat := ScientificNotationFormat(expectedIntValue)

	actualIntValue := sciNotationFormat.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected sciNotationFormat integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected sciNotationFormat integer value = '%v'\n"+
			"Actual sciNotationFormat integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := sciNotationFormat.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected ScientificNotationFormat(-972)\n"+
			"would return name of 'None' from \n"+
			"sciNotationFormat.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
