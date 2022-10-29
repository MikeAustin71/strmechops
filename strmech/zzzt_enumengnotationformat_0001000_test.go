package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func EngineeringNotationFormatTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []EngineeringNotationFormat,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"EngineeringNotationFormatTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"Exponential",
		"ENotUprCaseELeadPlus",
		"ENotUprCaseENoLeadPlus",
		"ENotLwrCaseELeadPlus",
		"ENotLwrCaseENoLeadPlus",
		"SIPrefixSymbol",
		"SIPrefixName",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).None())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).Exponential())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).ENotUprCaseELeadPlus())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).ENotUprCaseENoLeadPlus())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).ENotLwrCaseELeadPlus())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).ENotLwrCaseENoLeadPlus())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).SIPrefixSymbol())

	enumValues =
		append(enumValues, EngineeringNotationFormat(0).SIPrefixName())

	intValues =
		append(intValues, EngNotFmt.None().XValueInt())

	intValues =
		append(intValues, EngNotFmt.Exponential().XValueInt())

	intValues =
		append(intValues, EngNotFmt.ENotUprCaseELeadPlus().XValueInt())

	intValues =
		append(intValues, EngNotFmt.ENotUprCaseENoLeadPlus().XValueInt())

	intValues =
		append(intValues, EngNotFmt.ENotLwrCaseELeadPlus().XValueInt())

	intValues =
		append(intValues, EngNotFmt.ENotLwrCaseENoLeadPlus().XValueInt())

	intValues =
		append(intValues, EngNotFmt.ENotLwrCaseENoLeadPlus().XValueInt())

	intValues =
		append(intValues, EngNotFmt.SIPrefixSymbol().XValueInt())

	intValues =
		append(intValues, EngNotFmt.SIPrefixName().XValueInt())

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

func TestEngineeringNotationFormat_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestEngineeringNotationFormat_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		EngineeringNotationFormatTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var engNotationFormat1, engNotationFormat2,
		engNotationFormat3, engNotationFormat4,
		engNotationFormat5, engNotationFormat6 EngineeringNotationFormat

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		engNotationFormat1 = enumValues[i]

		isValid = engNotationFormat1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: engNotationFormat1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"engNotationFormat1 string value  = '%v'\n"+
					"engNotationFormat1 integer value = '%v'\n",
					ePrefix.String(),
					engNotationFormat1.String(),
					engNotationFormat1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"engNotationFormat1 string value  = '%v'\n"+
				"engNotationFormat1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				engNotationFormat1.String(),
				engNotationFormat1.XValueInt())

			return

		}

		engNotationFormat2,
			err = engNotationFormat1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  engNotationFormat1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"engNotationFormat1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				engNotationFormat1.String(),
				err.Error())

			return
		}

		if engNotationFormat2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"engNotationFormat2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"engNotationFormat2 string value  = '%v'\n"+
				"engNotationFormat2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				engNotationFormat2.String(),
				engNotationFormat2.XValueInt())

			return
		}

		engNotationFormat3 = enumValues[i]

		if engNotationFormat3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: engNotationFormat3.XValueInt() != intValues[%v]\n"+
				"engNotationFormat3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				engNotationFormat3.XValueInt(),
				i,
				intValues[i])

			return
		}

		engNotationFormat4,
			err = engNotationFormat3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by engNotationFormat3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if engNotationFormat4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: engNotationFormat4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"engNotationFormat4 string value  = '%v'\n"+
				"engNotationFormat4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				engNotationFormat4.String(),
				engNotationFormat4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		engNotationFormat5 = engNotationFormat1.XValue()

		engNotationFormat6 = engNotationFormat2.XValue()

		if engNotationFormat5 != engNotationFormat6 {
			t.Errorf("%v\n"+
				"Error: engNotationFormat5 != engNotationFormat6\n"+
				"engNotationFormat5 = engNotationFormat1.XValue()\n"+
				"engNotationFormat6 = engNotationFormat2.XValue()\n"+
				"engNotationFormat5 string value  = '%v'\n"+
				"engNotationFormat5 integer value = '%v'\n"+
				"engNotationFormat6 string value  = '%v'\n"+
				"engNotationFormat6 integer value = '%v'\n",
				ePrefix.String(),
				engNotationFormat5.String(),
				engNotationFormat5.XValueInt(),
				engNotationFormat6.String(),
				engNotationFormat6.XValueInt())

			return
		}

		_,
			err = engNotationFormat6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from engNotationFormat6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"engNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				engNotationFormat6.String())

			return
		}

		_,
			err = engNotationFormat6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from engNotationFormat6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"engNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				engNotationFormat6.String())

			return
		}

		_,
			err = engNotationFormat6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from engNotationFormat6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"engNotationFormat6 string value = '%v'\n",
				ePrefix.String(),
				i,
				engNotationFormat6.String())

			return
		}

	}

	return
}

func TestEngineeringNotationFormat_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestEngineeringNotationFormat_XReturnNoneIfInvalid_000200()",
		"")

	engNotationFormat := EngineeringNotationFormat(-972)

	valueNone := engNotationFormat.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected EngineeringNotationFormat(-972)\n"+
			"would return name of 'None' from \n"+
			"engNotationFormat.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strEngineeringNotationFormat := engNotationFormat.String()

	strEngineeringNotationFormat = strings.ToLower(strEngineeringNotationFormat)

	if !strings.Contains(strEngineeringNotationFormat, "error") {

		t.Errorf("%v\n"+
			"Error: Expected EngineeringNotationFormat(-972).String()\n"+
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
		EngineeringNotationFormatTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var engNotationFormat2 EngineeringNotationFormat

	engNotationFormat2 = enumValues[1].XReturnNoneIfInvalid()

	if engNotationFormat2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: engNotationFormat2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"engNotationFormat2 string value  = '%v'\n"+
			"engNotationFormat2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			engNotationFormat2.String(),
			engNotationFormat2.XValueInt())
		return
	}

	return
}

func TestEngineeringNotationFormat_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestEngineeringNotationFormat_XValueInt_000300()",
		"")

	expectedIntValue := -972

	engNotationFormat := EngineeringNotationFormat(expectedIntValue)

	actualIntValue := engNotationFormat.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected engNotationFormat integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected engNotationFormat integer value = '%v'\n"+
			"Actual engNotationFormat integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := engNotationFormat.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected EngineeringNotationFormat(-972)\n"+
			"would return name of 'None' from \n"+
			"engNotationFormat.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
