package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func NumericSymbolClassTestSetup0010(
	errorPrefix interface{}) (
	ucNames []string,
	lcNames []string,

	intValues []int,
	enumValues []NumericSymbolClass,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumericSymbolClassTestSetup0010()",
		"Initial Setup")

	if err != nil {
		return ucNames, lcNames, intValues, enumValues, err
	}

	ucNames = []string{
		"None",
		"NumberSign",
		"CurrencySign",
		"IntegerSeparator",
		"DecimalSeparator",
	}

	lenUcNames := len(ucNames)

	lcNames =
		make([]string, lenUcNames)

	for i := 0; i < lenUcNames; i++ {

		lcNames[i] = strings.ToLower(ucNames[i])

	}

	enumValues =
		append(enumValues, NumericSymbolClass(0).None())

	enumValues =
		append(enumValues, NumericSymbolClass(0).NumberSign())

	enumValues =
		append(enumValues, NumericSymbolClass(0).CurrencySign())

	enumValues =
		append(enumValues, NumericSymbolClass(0).IntegerSeparator())

	enumValues =
		append(enumValues, NumericSymbolClass(0).DecimalSeparator())

	intValues =
		append(intValues, NumSymClass.None().XValueInt())

	intValues =
		append(intValues, NumSymClass.NumberSign().XValueInt())

	intValues =
		append(intValues, NumSymClass.CurrencySign().XValueInt())

	intValues =
		append(intValues, NumSymClass.IntegerSeparator().XValueInt())

	intValues =
		append(intValues, NumSymClass.DecimalSeparator().XValueInt())

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

func TestNumericSymbolClass_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolClass_XValueInt_000100()",
		"")

	ucNames,
		lcNames,
		intValues,
		enumValues,
		err :=
		NumericSymbolClassTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var isValid bool
	var numericSymbolClass1, numericSymbolClass2,
		numericSymbolClass3, numericSymbolClass4,
		numericSymbolClass5, numericSymbolClass6 NumericSymbolClass

	lenUcNames := len(ucNames)

	for i := 0; i < lenUcNames; i++ {

		numericSymbolClass1 = enumValues[i]

		isValid = numericSymbolClass1.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: NumericSymbolClass1.None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n"+
					"NumericSymbolClass1 string value  = '%v'\n"+
					"NumericSymbolClass1 integer value = '%v'\n",
					ePrefix.String(),
					numericSymbolClass1.String(),
					numericSymbolClass1.XValueInt())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"numericSymbolClass1 string value  = '%v'\n"+
				"numericSymbolClass1 integer value = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				numericSymbolClass1.String(),
				numericSymbolClass1.XValueInt())

			return

		}

		numericSymbolClass2,
			err = numericSymbolClass1.XParseString(
			ucNames[i],
			true)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned from  numericSymbolClass1."+
				"XParseString(ucNames[%v]\n"+
				"ucName = %v\n"+
				"numericSymbolClass1 string value = '%v'\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSymbolClass1.String(),
				err.Error())

			return
		}

		if numericSymbolClass2.String() != ucNames[i] {
			t.Errorf("%v\n"+
				"numericSymbolClass2.String() != ucNames[%v]\n"+
				"ucName = '%v'\n"+
				"numericSymbolClass2 string value  = '%v'\n"+
				"numericSymbolClass2 integer value = '%v'\n",
				ePrefix.String(),
				i,
				ucNames[i],
				numericSymbolClass2.String(),
				numericSymbolClass2.XValueInt())

			return
		}

		numericSymbolClass3 = enumValues[i]

		if numericSymbolClass3.XValueInt() != intValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSymbolClass3.XValueInt() != intValues[%v]\n"+
				"numericSymbolClass3.XValueInt() = '%v'\n"+
				"             intValues[%v] = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolClass3.XValueInt(),
				i,
				intValues[i])

			return
		}

		numericSymbolClass4,
			err = numericSymbolClass3.XParseString(
			lcNames[i],
			false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by numericSymbolClass3.XParseString("+
				"lcNames[%v])\n"+
				"Error:\n%v\n",
				ePrefix.String(),
				i,
				err.Error())

			return
		}

		if numericSymbolClass4 != enumValues[i] {
			t.Errorf("%v\n"+
				"Error: numericSymbolClass4 != enumValues[%v]\n"+
				"                 lcNames[%v] = '%v'\n"+
				"numericSymbolClass4 string value  = '%v'\n"+
				"numericSymbolClass4 integer value = '%v'\n"+
				"enumValues[%v] string value  = '%v'\n"+
				"enumValues[%v] integer value = '%v'\n",
				ePrefix.String(),
				i,
				i,
				lcNames[i],
				numericSymbolClass4.String(),
				numericSymbolClass4.XValueInt(),
				i,
				enumValues[i].String(),
				i,
				enumValues[i].XValueInt())

			return
		}

		numericSymbolClass5 = numericSymbolClass1.XValue()

		numericSymbolClass6 = numericSymbolClass2.XValue()

		if numericSymbolClass5 != numericSymbolClass6 {
			t.Errorf("%v\n"+
				"Error: numericSymbolClass5 != numericSymbolClass6\n"+
				"numericSymbolClass5 = numericSymbolClass1.XValue()\n"+
				"numericSymbolClass6 = numericSymbolClass2.XValue()\n"+
				"numericSymbolClass5 string value  = '%v'\n"+
				"numericSymbolClass5 integer value = '%v'\n"+
				"numericSymbolClass6 string value  = '%v'\n"+
				"numericSymbolClass6 integer value = '%v'\n",
				ePrefix.String(),
				numericSymbolClass5.String(),
				numericSymbolClass5.XValueInt(),
				numericSymbolClass6.String(),
				numericSymbolClass6.XValueInt())

			return
		}

		_,
			err = numericSymbolClass6.XParseString(
			"How Now Brown Cow",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolClass6.XParseString()\n"+
				"because value string = 'How Now Brown Cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolClass6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolClass6.String())

			return
		}

		_,
			err = numericSymbolClass6.XParseString(
			"how now brown cow",
			false)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolClass6.XParseString()\n"+
				"because value string = 'now now brown cow'\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolClass6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolClass6.String())

			return
		}

		_,
			err = numericSymbolClass6.XParseString(
			"X",
			true)

		if err == nil {
			t.Errorf("\n%v\n"+
				"Expected an error return from numericSymbolClass6.XParseString()\n"+
				"because value string = 'X' is less than the\n"+
				"minimum required length.\n"+
				"HOWEVER, NO ERROR WAS RETURNED!\n"+
				"i = '%v'\n"+
				"numericSymbolClass6 string value = '%v'\n",
				ePrefix.String(),
				i,
				numericSymbolClass6.String())

			return
		}

	}

	return
}

func TestNumericSymbolClass_XReturnNoneIfInvalid_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolClass_XReturnNoneIfInvalid_000200()",
		"")

	numericSymbolClass := NumericSymbolClass(-972)

	valueNone := numericSymbolClass.XReturnNoneIfInvalid()

	if valueNone.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolClass(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSymbolClass.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"valueNone string value = '%v'\n"+
			"   valueNone int value = '%v'\n",
			ePrefix.String(),
			valueNone.String(),
			valueNone.XValueInt())

		return

	}

	strNumericSymbolClass := numericSymbolClass.String()

	strNumericSymbolClass = strings.ToLower(strNumericSymbolClass)

	if !strings.Contains(strNumericSymbolClass, "error") {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolClass(-972).String()\n"+
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
		NumericSymbolClassTestSetup0010(
			ePrefix)

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	var numericSymbolClass2 NumericSymbolClass

	numericSymbolClass2 = enumValues[1].XReturnNoneIfInvalid()

	if numericSymbolClass2 != enumValues[1] {
		t.Errorf("%v\n"+
			"Error: numericSymbolClass2 != enumValues[1].XReturnNoneIfInvalid()\n"+
			"enumValues[1]  string value  = '%v'\n"+
			"enumValues[1]  integer value = '%v'\n"+
			"numericSymbolClass2 string value  = '%v'\n"+
			"numericSymbolClass2 integer value = '%v'\n",
			ePrefix.String(),
			enumValues[1].String(),
			enumValues[1].XValueInt(),
			numericSymbolClass2.String(),
			numericSymbolClass2.XValueInt())
		return
	}

	return
}

func TestNumericSymbolClass_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumericSymbolClass_XValueInt_000300()",
		"")

	expectedIntValue := -972

	numericSymbolClass := NumericSymbolClass(expectedIntValue)

	actualIntValue := numericSymbolClass.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected numericSymbolClass integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected numericSymbolClass integer value = '%v'\n"+
			"Actual numericSymbolClass integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := numericSymbolClass.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected NumericSymbolClass(-972)\n"+
			"would return name of 'None' from \n"+
			"numericSymbolClass.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}
