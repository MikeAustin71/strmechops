package strmech

import "testing"

func TestDataFieldTrailingDelimiterType_ParseString_01(t *testing.T) {

	testStr := "Comment"

	dTypEnum, err := DfTrailDelimiter.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by DfTrailDelimiter.ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if DfTrailDelimiter.Comment() != dTypEnum {
		t.Errorf("Expected DataFieldTrailingDelimiterType Value =='%v'.\n"+
			"Instead DataFieldTrailingDelimiterType Value = %v. Numeric Value='%v'\n"+
			DfTrailDelimiter.Comment().String(), dTypEnum.String(), int(dTypEnum))
	}

}

func TestDataFieldTrailingDelimiterType_ParseString_02(t *testing.T) {

	testStr := "Co"

	_, err := DfTrailDelimiter.XParseString(testStr, true)

	if err == nil {
		t.Error("Expected and error return from DfTrailDelimiter.ParseString(testStr, true).\n" +
			"because parameter 'testStr' was only 2-characters long.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestDataFieldTrailingDelimiterType_ParseString_03(t *testing.T) {

	testStr := "BadEnumerationStr"

	_, err := DfTrailDelimiter.XParseString(testStr, true)

	if err == nil {
		t.Error("Expected and error return from DfTrailDelimiter.ParseString(testStr, true).\n" +
			"because parameter 'testStr' is an invalid enumeration value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}
func TestDataFieldTrailingDelimiterType_ParseString_04(t *testing.T) {

	testStr := "comment"

	dTypEnum, err := DfTrailDelimiter.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by DfTrailDelimiter.ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if DfTrailDelimiter.Comment() != dTypEnum {
		t.Errorf("Expected DataFieldTrailingDelimiterType Value =='%v'.\n"+
			"Instead DataFieldTrailingDelimiterType Value = %v. Numeric Value='%v'\n"+
			DfTrailDelimiter.Comment().String(), dTypEnum.String(), int(dTypEnum))
	}

}

func TestDataFieldTrailingDelimiterType_ParseString_05(t *testing.T) {

	testStr := "co"

	_, err := DfTrailDelimiter.XParseString(testStr, false)

	if err == nil {
		t.Error("Expected and error return from DfTrailDelimiter.ParseString(testStr, true).\n" +
			"because parameter 'testStr' was only 2-characters long.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestDataFieldTrailingDelimiterType_ParseString_06(t *testing.T) {

	testStr := "badenumerationstr"

	_, err := DfTrailDelimiter.XParseString(testStr, false)

	if err == nil {
		t.Error("Expected and error return from DfTrailDelimiter.ParseString(testStr, true).\n" +
			"because parameter 'testStr' is an invalid enumeration value.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
		return
	}

}

func TestDataFieldTrailingDelimiterType_ParseString_07(t *testing.T) {

	testStr := "EndOfString()"

	dTypEnum, err := DfTrailDelimiter.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by DfTrailDelimiter.ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if DfTrailDelimiter.EndOfString() != dTypEnum {
		t.Errorf("Expected DataFieldTrailingDelimiterType Value =='%v'.\n"+
			"Instead DataFieldTrailingDelimiterType Value = %v. Numeric Value='%v'\n"+
			DfTrailDelimiter.EndOfString().String(), dTypEnum.String(), int(dTypEnum))
	}

}

func TestDataFieldTrailingDelimiterType_StatusIsValid_01(t *testing.T) {
	testDataFieldType := DfTrailDelimiter.EndOfField()

	err := testDataFieldType.StatusIsValid()

	if err != nil {
		t.Error("Error returned on EndOfField Status Is Valid")
	}
}

func TestDataFieldTrailingDelimiterType_StatusIsValid_02(t *testing.T) {
	testDataFieldType := DataFieldTrailingDelimiterType(-99)

	err := testDataFieldType.StatusIsValid()

	if err == nil {
		t.Error("Expected an error return from StatusIsValid() because\n" +
			"DataFieldTrailingDelimiterType(-99) is invalid!\n")
	}
}

func TestDataFieldTrailingDelimiterType_XIsValid_01(t *testing.T) {
	testDataFieldType := DfTrailDelimiter.EndOfField()

	isOk := testDataFieldType.XIsValid()

	if !isOk {
		t.Error("Error returned on EndOfField Status Is Valid")
	}
}

func TestDataFieldTrailingDelimiterType_XIsValid_02(t *testing.T) {
	testDataFieldType := DataFieldTrailingDelimiterType(-99)

	isOk := testDataFieldType.XIsValid()

	if isOk {
		t.Error("Expected an error return from StatusIsValid() because\n" +
			"DataFieldTrailingDelimiterType(-99) is invalid!\n")
	}
}

func TestDataFieldTrailingDelimiterType_String_01(t *testing.T) {

	testDFType := DfTrailDelimiter.EndOfLine()

	if "EndOfLine" != testDFType.String() {
		t.Errorf("Expected DfTrailDelimiter.EndOfLine().String()='EndOfLine'.\n"+
			"Instead, DfTrailDelimiter.EndOfLine().String()='%v'\n",
			testDFType.String())
	}
}

func TestDataFieldTrailingDelimiterType_String_02(t *testing.T) {

	testDFType := DataFieldTrailingDelimiterType(-99)

	if "" != testDFType.String() {
		t.Errorf("Expected testDFType.String() to return an empty string\n"+
			"because, 'testDFType' is an invalid enumeration code.\n"+
			"Instead, testDFType.String()='%v'\n",
			testDFType.String())
	}
}

func TestDataFieldTrailingDelimiterType_Value_01(t *testing.T) {

	testDFType := DataFieldTrailingDelimiterType(0).EndOfString()

	testDFType2 := DataFieldTrailingDelimiterType(testDFType.XValueInt())

	if testDFType2 != testDFType {
		t.Errorf("ERROR: Expected testDFType2 == testDFType. It Does NOT!\n"+
			"testDFType2=='%d'\n"+
			"testDFType=='%d'\n",
			int(testDFType2), int(testDFType))
	}
}
