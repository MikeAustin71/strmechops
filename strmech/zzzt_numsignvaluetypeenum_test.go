package strmech

import "testing"

func TestNumericSignValueType_XArithmeticValue_000100(t *testing.T) {
	numSignValType := NumericSignValueType(0).None()

	arithmeticVal := numSignValType.XArithmeticValue()

	expectedVal := -2

	if expectedVal != arithmeticVal {
		t.Errorf("Expected that NumericSignValueType(0).None()\n"+
			"would yield an arithmetic value of '%v'.\n"+
			"Instead, arithmetic value=='%v'\n",
			expectedVal,
			arithmeticVal)
	}
}

func TestNumericSignValueType_XArithmeticValue_000200(t *testing.T) {
	numSignValType := NumericSignValueType(-6004)

	arithmeticVal := numSignValType.XArithmeticValue()

	expectedVal := -6004

	if expectedVal != arithmeticVal {
		t.Errorf("Expected that NumericSignValueType(-6004)\n"+
			"would yield an arithmetic value of '%v'.\n"+
			"Instead, arithmetic value=='%v'\n",
			expectedVal,
			arithmeticVal)
	}
}

func TestNumericSignValueType_XArithmeticValue_000300(t *testing.T) {
	numSignValType := NumericSignValueType(0).Negative()

	arithmeticVal := numSignValType.XArithmeticValue()

	expectedVal := -1

	if expectedVal != arithmeticVal {
		t.Errorf("Expected that NumericSignValueType(0).Negative()\n"+
			"would yield an arithmetic value of '%v'.\n"+
			"Instead, arithmetic value=='%v'\n",
			expectedVal,
			arithmeticVal)
	}
}

func TestNumericSignValueType_XArithmeticValue_000400(t *testing.T) {
	numSignValType := NumericSignValueType(0).Zero()

	arithmeticVal := numSignValType.XArithmeticValue()

	expectedVal := 0

	if expectedVal != arithmeticVal {
		t.Errorf("Expected that NumericSignValueType(0).Zero()\n"+
			"would yield an arithmetic value of '%v'.\n"+
			"Instead, arithmetic value=='%v'\n",
			expectedVal,
			arithmeticVal)
	}
}

func TestNumericSignValueType_XArithmeticValue_000500(t *testing.T) {
	numSignValType := NumericSignValueType(0).Positive()

	arithmeticVal := numSignValType.XArithmeticValue()

	expectedVal := 1

	if expectedVal != arithmeticVal {
		t.Errorf("Expected that NumericSignValueType(0).Positive()\n"+
			"would yield an arithmetic value of '%v'.\n"+
			"Instead, arithmetic value=='%v'\n",
			expectedVal,
			arithmeticVal)
	}
}

func TestNumericSignValueType_XIsPositiveOrNegative_000100(t *testing.T) {

	numSignValType := NumericSignValueType(0).Positive()

	isPosNeg := numSignValType.XIsPositiveOrNegative()

	if isPosNeg != true {
		t.Errorf("ERROR\n"+
			"Expected numSignValType.XIsPositiveOrNegative()\n"+
			"to return 'true' because the value is '%v'.\n"+
			"HOWEVER, THE RETURN VALUE IS FALSE!\n",
			numSignValType.String())
	}

}

func TestNumericSignValueType_XIsPositiveOrNegative_000200(t *testing.T) {

	numSignValType := NumericSignValueType(0).Negative()

	isPosNeg := numSignValType.XIsPositiveOrNegative()

	if isPosNeg != true {
		t.Errorf("ERROR\n"+
			"Expected numSignValType.XIsPositiveOrNegative()\n"+
			"to return 'true' because the value is '%v'.\n"+
			"HOWEVER, THE RETURN VALUE IS FALSE!\n",
			numSignValType.String())
	}

}

func TestNumericSignValueType_XIsPositiveOrNegative_000300(t *testing.T) {

	numSignValType := NumericSignValueType(0).Zero()

	isPosNeg := numSignValType.XIsPositiveOrNegative()

	if isPosNeg != false {
		t.Errorf("ERROR\n"+
			"Expected numSignValType.XIsPositiveOrNegative()\n"+
			"to return 'false' because the value is '%v'.\n"+
			"HOWEVER, THE RETURN VALUE IS TRUE!\n",
			numSignValType.String())
	}

}

func TestNumericSignValueType_XIsPositiveOrNegative_000400(t *testing.T) {

	numSignValType := NumericSignValueType(0).None()

	isPosNeg := numSignValType.XIsPositiveOrNegative()

	if isPosNeg != false {
		t.Errorf("ERROR\n"+
			"Expected numSignValType.XIsPositiveOrNegative()\n"+
			"to return 'false' because the value is '%v'.\n"+
			"HOWEVER, THE RETURN VALUE IS TRUE!\n",
			numSignValType.String())
	}

}

func TestNumericSignValueType_XIsPositiveOrNegative_000500(t *testing.T) {

	numSignValType := NumericSignValueType(943)

	isPosNeg := numSignValType.XIsPositiveOrNegative()

	if isPosNeg != false {
		t.Errorf("ERROR\n"+
			"Expected numSignValType.XIsPositiveOrNegative()\n"+
			"to return 'false' because the value is '%v'.\n"+
			"HOWEVER, THE RETURN VALUE IS TRUE!\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XIsValid_000100(t *testing.T) {

	numSignValType := NumericSignValueType(0).None()

	isValid := numSignValType.XIsValid()

	if isValid == true {
		t.Error("Error: Expected NumericSignValueType(0).None() to yield an INVALID enum.\n" +
			"It did NOT! Return value from numSignValType.XIsValid()=='true'\n")
	}
}

func TestNumericSignValueType_XIsValid_000200(t *testing.T) {

	numSignValType := NumericSignValueType(0).Positive()

	isValid := numSignValType.XIsValid()

	if isValid == false {
		t.Error("Error: Expected NumericSignValueType(0).Positive() to yield a VALID enum.\n" +
			"It did NOT! Return value from numSignValType.XIsValid()=='false'\n")
	}
}

func TestNumericSignValueType_XIsValid_000300(t *testing.T) {

	numSignValType := NumericSignValueType(0).Negative()

	isValid := numSignValType.XIsValid()

	if isValid == false {
		t.Error("Error: Expected NumericSignValueType(0).Negative() to yield a VALID enum.\n" +
			"It did NOT! Return value from numSignValType.XIsValid()=='false'\n")
	}
}

func TestNumericSignValueType_XIsValid_000400(t *testing.T) {

	numSignValType := NumericSignValueType(0).Zero()

	isValid := numSignValType.XIsValid()

	if isValid == false {
		t.Error("Error: Expected NumericSignValueType(0).Zero() to yield a VALID enum.\n" +
			"It did NOT! Return value from numSignValType.XIsValid()=='false'\n")
	}
}

func TestNumericSignValueType_XParseString_000100(t *testing.T) {

	testStr := "None"

	numSignValType,
		err := NumSignVal.XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000200(t *testing.T) {

	testStr := "None"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000300(t *testing.T) {

	testStr := "Negative"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Negative() {
		t.Errorf("Error: Expected return of object='numSignValType.Negative()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000400(t *testing.T) {

	testStr := "Zero"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Zero() {
		t.Errorf("Error: Expected return of object='numSignValType.Zero()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000500(t *testing.T) {

	testStr := "Positive"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Positive() {
		t.Errorf("Error: Expected return of object='numSignValType.Positive()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000600(t *testing.T) {

	testStr := "none"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000700(t *testing.T) {

	testStr := "negative"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Negative() {
		t.Errorf("Error: Expected return of object='numSignValType.Negative()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000800(t *testing.T) {

	testStr := "zero"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Zero() {
		t.Errorf("Error: Expected return of object='numSignValType.Zero()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}
}

func TestNumericSignValueType_XParseString_000900(t *testing.T) {

	testStr := "positive"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Positive() {
		t.Errorf("Error: Expected return of object='numSignValType.Positive()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}
}

func TestNumericSignValueType_XValue_000100(t *testing.T) {
	numSignValType := NumericSignValueType(0).None()

	currValue := numSignValType.XValue()

	if currValue != NumSignVal.None() {
		t.Errorf("Error: Expected return of object='NumericSignValueType(0).None()'.\n"+
			"Instead, object integer value = '%v'\n",
			NumSignSymbolPosition(0).XValueInt())
	}
}

func TestNumericSignValueType_XValue_000200(t *testing.T) {
	numSignValType := NumericSignValueType(0).Negative()

	currValue := numSignValType.XValue()

	if currValue != NumericSignValueType(0).Negative() {
		t.Errorf("Error: Expected return of object='NumericSignValueType(0).Negative()'.\n"+
			"Instead, object integer value = '%v'\n",
			NumSignSymbolPosition(0).XValueInt())
	}
}

func TestNumericSignValueType_XValueInt_000100(t *testing.T) {

	numSignValType := NumericSignValueType(0).Negative()

	intValue := numSignValType.XValueInt()

	if intValue != -1 {
		t.Errorf("Error: Expected return of object integer = '-1' (Negative).\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSignValueType_XValueInt_000200(t *testing.T) {

	numSignValType := NumericSignValueType(0).None()

	intValue := numSignValType.XValueInt()

	if intValue != -2 {
		t.Errorf("Error: Expected return of NumericSignValueType(0).None() "+
			"== -2.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSignValueType_XValueInt_000300(t *testing.T) {

	numSignValType := NumericSignValueType(0).Zero()

	intValue := numSignValType.XValueInt()

	if intValue != 0 {
		t.Errorf("Error: Expected return of NumericSignValueType(0).Zero() "+
			"== 0.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSignValueType_XValueInt_000400(t *testing.T) {

	numSignValType := NumericSignValueType(0).Positive()

	intValue := numSignValType.XValueInt()

	if intValue != 1 {
		t.Errorf("Error: Expected return of NumericSignValueType(0).Positive() "+
			"== 1.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSignValueType_String_000100(t *testing.T) {

	testStr := "None"

	numSignValType := NumericSignValueType(0).None()

	actualStr := numSignValType.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"None\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumericSignValueType_String_000200(t *testing.T) {

	testStr := "Negative"

	numSignValType := NumericSignValueType(0).Negative()

	actualStr := numSignValType.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Negative\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumericSignValueType_String_000300(t *testing.T) {

	testStr := "Zero"

	numSignValType := NumericSignValueType(0).Zero()

	actualStr := numSignValType.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Zero\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumericSignValueType_String_000400(t *testing.T) {

	testStr := "Positive"

	numSignValType := NumericSignValueType(0).Positive()

	actualStr := numSignValType.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Positive\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}
