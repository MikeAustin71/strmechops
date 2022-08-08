package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TextJustifyTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []TextJustify,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextJustifyTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Left",
		"Right",
		"Center",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, TextJustify(0).None())

	enumValues =
		append(enumValues, TextJustify(0).Left())

	enumValues =
		append(enumValues, TextJustify(0).Right())

	enumValues =
		append(enumValues, TextJustify(0).Center())

	intValues =
		append(intValues, TxtJustify.None().XValueInt())

	intValues =
		append(intValues, TxtJustify.Left().XValueInt())

	intValues =
		append(intValues, TxtJustify.Right().XValueInt())

	intValues =
		append(intValues, TxtJustify.Center().XValueInt())

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

func TestTextJustify_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextJustify_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		TextJustifyTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var textJustify1, textJustify2,
		textJustify3, textJustify4,
		textJustify5, textJustify6 TextJustify

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		textJustify1 = enumValues[i]

		isValid = textJustify1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: TextJustify1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"textJustify1 string value  = '%v'\n"+
					"textJustify1 integer value = '%v'\n",
					ePrefix.String(),
					textJustify1.String(),
					textJustify1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"textJustify1 string value  = '%v'\n"+
				"textJustify1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				textJustify1.String(),
				textJustify1.XValueInt())

			return

		}

		textJustify2,
			err = textJustify1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  textJustify1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"textJustify1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				textJustify1.String(),
				err.Error())

			return
		}

		if textJustify2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"textJustify2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"textJustify2 string value  = '%v'\n"+
				"textJustify2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				textJustify2.String(),
				textJustify2.XValueInt())

			return
		}

		textJustify3 = enumValues[i]

		if textJustify3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: textJustify3.XValueInt() != intValues[%v]\n"+
				"textJustify3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				textJustify3.XValueInt(),
				i,
				intValues[i])

			return
		}

		textJustify4,
			err = textJustify3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by textJustify3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if textJustify4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: textJustify4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"textJustify4 string value  = '%v'\n"+
				"textJustify4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				textJustify4.String(),
				textJustify4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		textJustify5 = textJustify1.XValue()

		textJustify6 = textJustify2.XValue()

		if textJustify5 != textJustify6 {
			t.Errorf("%v\n"+
				"Error: textJustify5 != textJustify6\n"+
				"textJustify5 = textJustify1.XValue()\n"+
				"textJustify6 = textJustify2.XValue()\n"+
				"textJustify5 string value  = '%v'\n"+
				"textJustify5 integer value = '%v'\n"+
				"textJustify6 string value  = '%v'\n"+
				"textJustify6 integer value = '%v'\n",
				ePrefix.String(),
				textJustify5.String(),
				textJustify5.XValueInt(),
				textJustify6.String(),
				textJustify6.XValueInt())

			return
		}

		_,
			err = textJustify6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textJustify6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textJustify6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textJustify6.String())

			return
		}

		_,
			err = textJustify6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textJustify6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textJustify6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textJustify6.String())

			return
		}

		_,
			err = textJustify6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textJustify6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textJustify6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textJustify6.String())

			return
		}

	}

	return
}

func TestTextJustify_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextJustify_XReturnNoneIfInvalid_000200()",
		"")

	textJustify := TextJustify(-972)

	valueNone := textJustify.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected TextJustify(-972)\n"+
			"would return name of 'None' from \n"+
			"textJustify.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strTextJustify := textJustify.String()

	strTextJustify = strings.ToLower(strTextJustify)

	if !strings.Contains(strTextJustify, "error") {

		t.Errorf("%v\n"+
			"Error: Expected TextJustify(-972).String()\n"+
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
		TextJustifyTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var textJustify2 TextJustify

	textJustify2 = enumValues[1].XReturnNoneIfInvalid()

	if textJustify2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: textJustify2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"textJustify2 string value  = '%v'\n"+
			"textJustify2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			textJustify2.String(),
			textJustify2.XValueInt())
		return
	}

	return
}

func TestTextJustify_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextJustify_XValueInt_000300()",
		"")

	expectedIntValue := -972

	textJustify := TextJustify(expectedIntValue)

	actualIntValue := textJustify.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected textJustify integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected textJustify integer value = '%v'\n"+
			"Actual textJustify integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := textJustify.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected TextJustify(-972)\n"+
			"would return name of 'None' from \n"+
			"textJustify.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
