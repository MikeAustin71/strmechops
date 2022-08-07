package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TextFieldTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []TextFieldType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Label",
		"DateTime",
		"Filler",
		"Spacer",
		"BlankLine",
		"SolidLine",
		"LineColumns",
		"TimerStartStop",
		"TextAdHoc",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, TextFieldType(0).None())

	enumValues =
		append(enumValues, TextFieldType(0).Label())

	enumValues =
		append(enumValues, TextFieldType(0).DateTime())

	enumValues =
		append(enumValues, TextFieldType(0).Filler())

	enumValues =
		append(enumValues, TextFieldType(0).Spacer())

	enumValues =
		append(enumValues, TextFieldType(0).BlankLine())

	enumValues =
		append(enumValues, TextFieldType(0).SolidLine())

	enumValues =
		append(enumValues, TextFieldType(0).LineColumns())

	enumValues =
		append(enumValues, TextFieldType(0).TimerStartStop())

	enumValues =
		append(enumValues, TextFieldType(0).TextAdHoc())

	intValues =
		append(intValues, TxtFieldType.None().XValueInt())

	intValues =
		append(intValues, TxtFieldType.Label().XValueInt())

	intValues =
		append(intValues, TxtFieldType.DateTime().XValueInt())

	intValues =
		append(intValues, TxtFieldType.Filler().XValueInt())

	intValues =
		append(intValues, TxtFieldType.Spacer().XValueInt())

	intValues =
		append(intValues, TxtFieldType.BlankLine().XValueInt())

	intValues =
		append(intValues, TxtFieldType.SolidLine().XValueInt())

	intValues =
		append(intValues, TxtFieldType.LineColumns().XValueInt())

	intValues =
		append(intValues, TxtFieldType.TimerStartStop().XValueInt())

	intValues =
		append(intValues, TxtFieldType.TextAdHoc().XValueInt())

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

func TestTextFieldType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		TextFieldTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var textFieldType1, textFieldType2,
		textFieldType3, textFieldType4,
		textFieldType5, textFieldType6 TextFieldType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		textFieldType1 = enumValues[i]

		isValid = textFieldType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: TextFieldType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"TextFieldType1 string value  = '%v'\n"+
					"TextFieldType1 integer value = '%v'\n",
					ePrefix.String(),
					textFieldType1.String(),
					textFieldType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"textFieldType1 string value  = '%v'\n"+
				"textFieldType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				textFieldType1.String(),
				textFieldType1.XValueInt())

			return

		}

		textFieldType2,
			err = textFieldType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  textFieldType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"textFieldType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				textFieldType1.String(),
				err.Error())

			return
		}

		if textFieldType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"textFieldType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"textFieldType2 string value  = '%v'\n"+
				"textFieldType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				textFieldType2.String(),
				textFieldType2.XValueInt())

			return
		}

		textFieldType3 = enumValues[i]

		if textFieldType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: textFieldType3.XValueInt() != intValues[%v]\n"+
				"textFieldType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				textFieldType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		textFieldType4,
			err = textFieldType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by textFieldType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if textFieldType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: textFieldType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"textFieldType4 string value  = '%v'\n"+
				"textFieldType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				textFieldType4.String(),
				textFieldType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		textFieldType5 = textFieldType1.XValue()

		textFieldType6 = textFieldType2.XValue()

		if textFieldType5 != textFieldType6 {
			t.Errorf("%v\n"+
				"Error: textFieldType5 != textFieldType6\n"+
				"textFieldType5 = textFieldType1.XValue()\n"+
				"textFieldType6 = textFieldType2.XValue()\n"+
				"textFieldType5 string value  = '%v'\n"+
				"textFieldType5 integer value = '%v'\n"+
				"textFieldType6 string value  = '%v'\n"+
				"textFieldType6 integer value = '%v'\n",
				ePrefix.String(),
				textFieldType5.String(),
				textFieldType5.XValueInt(),
				textFieldType6.String(),
				textFieldType6.XValueInt())

			return
		}

		_,
			err = textFieldType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textFieldType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textFieldType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textFieldType6.String())

			return
		}

		_,
			err = textFieldType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textFieldType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textFieldType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textFieldType6.String())

			return
		}

		_,
			err = textFieldType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from textFieldType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"textFieldType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				textFieldType6.String())

			return
		}

	}

	return
}

func TestTextFieldType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldType_XReturnNoneIfInvalid_000200()",
		"")

	textFieldType := TextFieldType(-972)

	valueNone := textFieldType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected TextFieldType(-972)\n"+
			"would return name of 'None' from \n"+
			"textFieldType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strTextFieldType := textFieldType.String()

	strTextFieldType = strings.ToLower(strTextFieldType)

	if !strings.Contains(strTextFieldType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected TextFieldType(-972).String()\n"+
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
		TextFieldTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var textFieldType2 TextFieldType

	textFieldType2 = enumValues[1].XReturnNoneIfInvalid()

	if textFieldType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: textFieldType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"textFieldType2 string value  = '%v'\n"+
			"textFieldType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			textFieldType2.String(),
			textFieldType2.XValueInt())
		return
	}

	return
}

func TestTextFieldType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	textFieldType := TextFieldType(expectedIntValue)

	actualIntValue := textFieldType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected textFieldType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected textFieldType integer value = '%v'\n"+
			"Actual textFieldType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := textFieldType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected TextFieldType(-972)\n"+
			"would return name of 'None' from \n"+
			"textFieldType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
