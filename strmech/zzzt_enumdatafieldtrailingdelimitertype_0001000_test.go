package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func DataFieldTrailingDelimiterTypeTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []DataFieldTrailingDelimiterType,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DataFieldTrailingDelimiterTypeTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"EndOfField",
		"Comment",
		"EndOfLine",
		"EndOfString",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, DataFieldTrailingDelimiterType(0).None())

	enumValues =
		append(enumValues, DataFieldTrailingDelimiterType(0).EndOfField())

	enumValues =
		append(enumValues, DataFieldTrailingDelimiterType(0).Comment())

	enumValues =
		append(enumValues, DataFieldTrailingDelimiterType(0).EndOfLine())

	enumValues =
		append(enumValues, DataFieldTrailingDelimiterType(0).EndOfString())

	intValues =
		append(intValues, DfTrailDelimiter.None().XValueInt())

	intValues =
		append(intValues, DfTrailDelimiter.EndOfField().XValueInt())

	intValues =
		append(intValues, DfTrailDelimiter.Comment().XValueInt())

	intValues =
		append(intValues, DfTrailDelimiter.EndOfLine().XValueInt())

	intValues =
		append(intValues, DfTrailDelimiter.EndOfString().XValueInt())

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

func TestDataFieldTrailingDelimiterType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestDataFieldTrailingDelimiterType_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		DataFieldTrailingDelimiterTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var DfTrailingDelimiterType1, DfTrailingDelimiterType2,
		DfTrailingDelimiterType3, DfTrailingDelimiterType4,
		DfTrailingDelimiterType5, DfTrailingDelimiterType6 DataFieldTrailingDelimiterType

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		DfTrailingDelimiterType1 = enumValues[i]

		isValid = DfTrailingDelimiterType1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: DataFieldTrailingDelimiterType1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"DfTrailingDelimiterType1 string value  = '%v'\n"+
					"DfTrailingDelimiterType1 integer value = '%v'\n",
					ePrefix.String(),
					DfTrailingDelimiterType1.String(),
					DfTrailingDelimiterType1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"DfTrailingDelimiterType1 string value  = '%v'\n"+
				"DfTrailingDelimiterType1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				DfTrailingDelimiterType1.String(),
				DfTrailingDelimiterType1.XValueInt())

			return

		}

		DfTrailingDelimiterType2,
			err = DfTrailingDelimiterType1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  DfTrailingDelimiterType1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"DfTrailingDelimiterType1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				DfTrailingDelimiterType1.String(),
				err.Error())

			return
		}

		if DfTrailingDelimiterType2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"DfTrailingDelimiterType2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"DfTrailingDelimiterType2 string value  = '%v'\n"+
				"DfTrailingDelimiterType2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				DfTrailingDelimiterType2.String(),
				DfTrailingDelimiterType2.XValueInt())

			return
		}

		DfTrailingDelimiterType3 = enumValues[i]

		if DfTrailingDelimiterType3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: DfTrailingDelimiterType3.XValueInt() != intValues[%v]\n"+
				"DfTrailingDelimiterType3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				DfTrailingDelimiterType3.XValueInt(),
				i,
				intValues[i])

			return
		}

		DfTrailingDelimiterType4,
			err = DfTrailingDelimiterType3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by DfTrailingDelimiterType3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if DfTrailingDelimiterType4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: DfTrailingDelimiterType4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"DfTrailingDelimiterType4 string value  = '%v'\n"+
				"DfTrailingDelimiterType4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				DfTrailingDelimiterType4.String(),
				DfTrailingDelimiterType4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		DfTrailingDelimiterType5 = DfTrailingDelimiterType1.XValue()

		DfTrailingDelimiterType6 = DfTrailingDelimiterType2.XValue()

		if DfTrailingDelimiterType5 != DfTrailingDelimiterType6 {
			t.Errorf("%v\n"+
				"Error: DfTrailingDelimiterType5 != DfTrailingDelimiterType6\n"+
				"DfTrailingDelimiterType5 = DfTrailingDelimiterType1.XValue()\n"+
				"DfTrailingDelimiterType6 = DfTrailingDelimiterType2.XValue()\n"+
				"DfTrailingDelimiterType5 string value  = '%v'\n"+
				"DfTrailingDelimiterType5 integer value = '%v'\n"+
				"DfTrailingDelimiterType6 string value  = '%v'\n"+
				"DfTrailingDelimiterType6 integer value = '%v'\n",
				ePrefix.String(),
				DfTrailingDelimiterType5.String(),
				DfTrailingDelimiterType5.XValueInt(),
				DfTrailingDelimiterType6.String(),
				DfTrailingDelimiterType6.XValueInt())

			return
		}

		_,
			err = DfTrailingDelimiterType6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from DfTrailingDelimiterType6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"DfTrailingDelimiterType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				DfTrailingDelimiterType6.String())

			return
		}

		_,
			err = DfTrailingDelimiterType6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from DfTrailingDelimiterType6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"DfTrailingDelimiterType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				DfTrailingDelimiterType6.String())

			return
		}

		_,
			err = DfTrailingDelimiterType6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from DfTrailingDelimiterType6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"DfTrailingDelimiterType6 string value = '%v'\n",
				ePrefix.String(),
				i,
				DfTrailingDelimiterType6.String())

			return
		}

	}

	return
}

func TestDataFieldTrailingDelimiterType_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestDataFieldTrailingDelimiterType_XReturnNoneIfInvalid_000200()",
		"")

	DfTrailingDelimiterType := DataFieldTrailingDelimiterType(-972)

	valueNone := DfTrailingDelimiterType.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected DataFieldTrailingDelimiterType(-972)\n"+
			"would return name of 'None' from \n"+
			"DfTrailingDelimiterType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strDataFieldTrailingDelimiterType := DfTrailingDelimiterType.String()

	strDataFieldTrailingDelimiterType = strings.ToLower(strDataFieldTrailingDelimiterType)

	if !strings.Contains(strDataFieldTrailingDelimiterType, "error") {

		t.Errorf("%v\n"+
			"Error: Expected DataFieldTrailingDelimiterType(-972).String()\n"+
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
		DataFieldTrailingDelimiterTypeTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var DfTrailingDelimiterType2 DataFieldTrailingDelimiterType

	DfTrailingDelimiterType2 = enumValues[1].XReturnNoneIfInvalid()

	if DfTrailingDelimiterType2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: DfTrailingDelimiterType2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"DfTrailingDelimiterType2 string value  = '%v'\n"+
			"DfTrailingDelimiterType2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			DfTrailingDelimiterType2.String(),
			DfTrailingDelimiterType2.XValueInt())
		return
	}

	return
}

func TestDataFieldTrailingDelimiterType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestDataFieldTrailingDelimiterType_XValueInt_000300()",
		"")

	expectedIntValue := -972

	DfTrailingDelimiterType := DataFieldTrailingDelimiterType(expectedIntValue)

	actualIntValue := DfTrailingDelimiterType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected DfTrailingDelimiterType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected DfTrailingDelimiterType integer value = '%v'\n"+
			"Actual DfTrailingDelimiterType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := DfTrailingDelimiterType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected DataFieldTrailingDelimiterType(-972)\n"+
			"would return name of 'None' from \n"+
			"DfTrailingDelimiterType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
