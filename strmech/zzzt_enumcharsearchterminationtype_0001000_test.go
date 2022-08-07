package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func CharSearchTerminationTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []CharSearchTerminationType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTerminationTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"ProcessError",
		"EndOfTargetString",
		"SearchLengthLimit",
		"TerminationDelimiters",
		"FoundSearchTarget",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, CharSearchTerminationType(0).None())

	enumValues =
		append(enumValues, CharSearchTerminationType(0).ProcessError())

	enumValues =
		append(enumValues, CharSearchTerminationType(0).EndOfTargetString())

	enumValues =
		append(enumValues, CharSearchTerminationType(0).SearchLengthLimit())

	enumValues =
		append(enumValues, CharSearchTerminationType(0).TerminationDelimiters())

	enumValues =
		append(enumValues, CharSearchTerminationType(0).FoundSearchTarget())

	intValues =
		append(intValues, CharSearchTermType.None().XValueInt())

	intValues =
		append(intValues, CharSearchTermType.ProcessError().XValueInt())

	intValues =
		append(intValues, CharSearchTermType.EndOfTargetString().XValueInt())

	intValues =
		append(intValues, CharSearchTermType.SearchLengthLimit().XValueInt())

	intValues =
		append(intValues, CharSearchTermType.TerminationDelimiters().XValueInt())

	intValues =
		append(intValues, CharSearchTermType.FoundSearchTarget().XValueInt())

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

func TestCharSearchTerminationType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		CharSearchTerminationTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var charSearchTerminationType1, charSearchTerminationType2,
		charSearchTerminationType3, charSearchTerminationType4,
		charSearchTerminationType5, charSearchTerminationType6 CharSearchTerminationType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		charSearchTerminationType1 = enumValues[i]

		isValid = charSearchTerminationType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: CharSearchTerminationType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"CharSearchTerminationType1 string value  = '%v'\n"+
					"CharSearchTerminationType1 integer value = '%v'\n",
					ePrefix.String(),
					charSearchTerminationType1.String(),
					charSearchTerminationType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"charSearchTerminationType1 string value  = '%v'\n"+
				"charSearchTerminationType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				charSearchTerminationType1.String(),
				charSearchTerminationType1.XValueInt())

			return

		}

		charSearchTerminationType2,
			err = charSearchTerminationType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  charSearchTerminationType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"charSearchTerminationType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				charSearchTerminationType1.String(),
				err.Error())

			return
		}

		if charSearchTerminationType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"charSearchTerminationType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"charSearchTerminationType2 string value  = '%v'\n"+
				"charSearchTerminationType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				charSearchTerminationType2.String(),
				charSearchTerminationType2.XValueInt())

			return
		}

		charSearchTerminationType3 = enumValues[i]

		if charSearchTerminationType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: charSearchTerminationType3.XValueInt() != intValues[%v]\n"+
				"charSearchTerminationType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				charSearchTerminationType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		charSearchTerminationType4,
			err = charSearchTerminationType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by charSearchTerminationType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if charSearchTerminationType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: charSearchTerminationType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"charSearchTerminationType4 string value  = '%v'\n"+
				"charSearchTerminationType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				charSearchTerminationType4.String(),
				charSearchTerminationType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		charSearchTerminationType5 = charSearchTerminationType1.XValue()

		charSearchTerminationType6 = charSearchTerminationType2.XValue()

		if charSearchTerminationType5 != charSearchTerminationType6 {
			t.Errorf("%v\n"+
				"Error: charSearchTerminationType5 != charSearchTerminationType6\n"+
				"charSearchTerminationType5 = charSearchTerminationType1.XValue()\n"+
				"charSearchTerminationType6 = charSearchTerminationType2.XValue()\n"+
				"charSearchTerminationType5 string value  = '%v'\n"+
				"charSearchTerminationType5 integer value = '%v'\n"+
				"charSearchTerminationType6 string value  = '%v'\n"+
				"charSearchTerminationType6 integer value = '%v'\n",
				ePrefix.String(),
				charSearchTerminationType5.String(),
				charSearchTerminationType5.XValueInt(),
				charSearchTerminationType6.String(),
				charSearchTerminationType6.XValueInt())

			return
		}

		_,
			err = charSearchTerminationType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from charSearchTerminationType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"charSearchTerminationType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				charSearchTerminationType6.String())

			return
		}

		_,
			err = charSearchTerminationType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from charSearchTerminationType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"charSearchTerminationType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				charSearchTerminationType6.String())

			return
		}

		_,
			err = charSearchTerminationType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from charSearchTerminationType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"charSearchTerminationType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				charSearchTerminationType6.String())

			return
		}

	}

	return
}

func TestCharSearchTerminationType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XReturnNoneIfInvalid_000200()",
		"")

	charSearchTerminationType := CharSearchTerminationType(-972)

	valueNone := charSearchTerminationType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharSearchTerminationType(-972)\n"+
			"would return name of 'None' from \n"+
			"charSearchTerminationType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strCharSearchTerminationType := charSearchTerminationType.String()

	strCharSearchTerminationType = strings.ToLower(strCharSearchTerminationType)

	if !strings.Contains(strCharSearchTerminationType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected CharSearchTerminationType(-972).String()\n"+
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
		CharSearchTerminationTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var charSearchTerminationType2 CharSearchTerminationType

	charSearchTerminationType2 = enumValues[1].XReturnNoneIfInvalid()

	if charSearchTerminationType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: charSearchTerminationType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"charSearchTerminationType2 string value  = '%v'\n"+
			"charSearchTerminationType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			charSearchTerminationType2.String(),
			charSearchTerminationType2.XValueInt())
		return
	}

	return
}

func TestCharSearchTerminationType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	charSearchTerminationType := CharSearchTerminationType(expectedIntValue)

	actualIntValue := charSearchTerminationType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected charSearchTerminationType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected charSearchTerminationType integer value = '%v'\n"+
			"Actual charSearchTerminationType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := charSearchTerminationType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharSearchTerminationType(-972)\n"+
			"would return name of 'None' from \n"+
			"charSearchTerminationType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
