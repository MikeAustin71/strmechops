package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func CharacterSearchTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []CharacterSearchType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharacterSearchTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"LinearTargetStartingIndex",
		"SingleTargetChar",
		"LinearEndOfString",
		"LatinEngAlphaLetter",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, CharacterSearchType(0).None())

	enumValues =
		append(enumValues, CharacterSearchType(0).
			LinearTargetStartingIndex())

	enumValues =
		append(enumValues, CharacterSearchType(0).
			SingleTargetChar())

	enumValues =
		append(enumValues, CharacterSearchType(0).
			LinearEndOfString())

	enumValues =
		append(enumValues, CharacterSearchType(0).
			LatinEngAlphaLetter())

	intValues =
		append(intValues, CharSearchType.None().XValueInt())

	intValues =
		append(intValues, CharSearchType.
			LinearTargetStartingIndex().XValueInt())

	intValues =
		append(intValues, CharSearchType.SingleTargetChar().
			XValueInt())

	intValues =
		append(intValues, CharSearchType.LinearEndOfString().
			XValueInt())

	intValues =
		append(intValues, CharSearchType.LatinEngAlphaLetter().
			XValueInt())

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

func TestCharacterSearchType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharacterSearchType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		CharacterSearchTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var characterSearchType1, characterSearchType2,
		characterSearchType3, characterSearchType4,
		characterSearchType5, characterSearchType6 CharacterSearchType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		characterSearchType1 = enumValues[i]

		isValid = characterSearchType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: CharacterSearchType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"CharacterSearchType1 string value  = '%v'\n"+
					"CharacterSearchType1 integer value = '%v'\n",
					ePrefix.String(),
					characterSearchType1.String(),
					characterSearchType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"characterSearchType1 string value  = '%v'\n"+
				"characterSearchType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				characterSearchType1.String(),
				characterSearchType1.XValueInt())

			return

		}

		characterSearchType2,
			err = characterSearchType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  characterSearchType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"characterSearchType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				characterSearchType1.String(),
				err.Error())

			return
		}

		if characterSearchType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"characterSearchType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"characterSearchType2 string value  = '%v'\n"+
				"characterSearchType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				characterSearchType2.String(),
				characterSearchType2.XValueInt())

			return
		}

		characterSearchType3 = enumValues[i]

		if characterSearchType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: characterSearchType3.XValueInt() != intValues[%v]\n"+
				"characterSearchType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				characterSearchType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		characterSearchType4,
			err = characterSearchType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by characterSearchType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if characterSearchType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: characterSearchType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"characterSearchType4 string value  = '%v'\n"+
				"characterSearchType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				characterSearchType4.String(),
				characterSearchType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		characterSearchType5 = characterSearchType1.XValue()

		characterSearchType6 = characterSearchType2.XValue()

		if characterSearchType5 != characterSearchType6 {
			t.Errorf("%v\n"+
				"Error: characterSearchType5 != characterSearchType6\n"+
				"characterSearchType5 = characterSearchType1.XValue()\n"+
				"characterSearchType6 = characterSearchType2.XValue()\n"+
				"characterSearchType5 string value  = '%v'\n"+
				"characterSearchType5 integer value = '%v'\n"+
				"characterSearchType6 string value  = '%v'\n"+
				"characterSearchType6 integer value = '%v'\n",
				ePrefix.String(),
				characterSearchType5.String(),
				characterSearchType5.XValueInt(),
				characterSearchType6.String(),
				characterSearchType6.XValueInt())

			return
		}

		_,
			err = characterSearchType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from characterSearchType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"characterSearchType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				characterSearchType6.String())

			return
		}

		_,
			err = characterSearchType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from characterSearchType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"characterSearchType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				characterSearchType6.String())

			return
		}

		_,
			err = characterSearchType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from characterSearchType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"characterSearchType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				characterSearchType6.String())

			return
		}

	}

	return
}

func TestCharacterSearchType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharacterSearchType_XReturnNoneIfInvalid_000200()",
		"")

	characterSearchType := CharacterSearchType(-972)

	valueNone := characterSearchType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharacterSearchType(-972)\n"+
			"would return name of 'None' from \n"+
			"characterSearchType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strCharacterSearchType := characterSearchType.String()

	strCharacterSearchType = strings.ToLower(strCharacterSearchType)

	if !strings.Contains(strCharacterSearchType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected CharacterSearchType(-972).String()\n"+
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
		CharacterSearchTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var characterSearchType2 CharacterSearchType

	characterSearchType2 = enumValues[1].XReturnNoneIfInvalid()

	if characterSearchType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: characterSearchType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"characterSearchType2 string value  = '%v'\n"+
			"characterSearchType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			characterSearchType2.String(),
			characterSearchType2.XValueInt())
		return
	}

	return
}

func TestCharacterSearchType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharacterSearchType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	characterSearchType := CharacterSearchType(expectedIntValue)

	actualIntValue := characterSearchType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected characterSearchType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected characterSearchType integer value = '%v'\n"+
			"Actual characterSearchType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := characterSearchType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharacterSearchType(-972)\n"+
			"would return name of 'None' from \n"+
			"characterSearchType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
