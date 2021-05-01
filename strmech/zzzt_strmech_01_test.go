package strmech

import (
	"testing"
)

func TestDataFieldProfileDto_ConvertToErrorState_01(t *testing.T) {

	dfProfile := DataFieldProfileDto{}

	dfProfile.DataFieldStr = "Hello World"
	dfProfile.DataFieldTrailingDelimiterType = DfTrailDelimiter.EndOfLine()
	dfProfile.DataFieldTrailingDelimiter = "\n"
	dfProfile.DataFieldIndex = 99
	dfProfile.DataFieldLength = len(dfProfile.DataFieldStr)

	dfProfile.ConvertToErrorState()

	if dfProfile.DataFieldStr != "" {
		t.Errorf("Expected dfProfile.DataFieldStr==Empty String.\n"+
			"Instead dfProfile.DataFieldStr='%v'\n", dfProfile.DataFieldStr)
	}

	if dfProfile.DataFieldIndex != -1 {
		t.Errorf("Expected dfProfile.DataFieldIndex==-1\n"+
			"Instead, dfProfile.DataFieldIndex=='%v'\n", dfProfile.DataFieldIndex)
	}

	if dfProfile.DataFieldLength != 0 {
		t.Errorf("Expected dfProfile.DataFieldLength==0\n"+
			"Instead, dfProfile.DataFieldLength=='%v'\n", dfProfile.DataFieldLength)
	}

	if dfProfile.NextTargetStrIndex != -1 {
		t.Errorf("Expected dfProfile.NextTargetStrIndex==-1\n"+
			"Instead, dfProfile.NextTargetStrIndex=='%v'\n", dfProfile.NextTargetStrIndex)
	}
}

func TestStrOps_BreakTextAtLineLength_01(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_01"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Lorem ipsum dolor sit amet, consectetur%adipiscing elit.%"

	actualTxt, err := StrMech{}.NewPtr().BreakTextAtLineLength(
		tstStr,
		40,
		'\n',
		thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = StrMech{}.Ptr().ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected string='%v'. Instead, string='%v'.",
			expected, actualTxt)
	}

}

func TestStrOps_BreakTextAtLineLength_02(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_02"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Did you know? The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = StrMech{}.Ptr().ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

func TestStrOps_BreakTextAtLineLength_03(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_03()"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Did you know? XX The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? XX The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = StrMech{}.Ptr().ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

func TestStrOps_BreakTextAtLineLength_04(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_04()"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "       Did you know? The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = StrMech{}.Ptr().ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

func TestStrOps_BreakTextAtLineLength_05(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_05()"

	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus eu ex sit amet " +
		"sapien consectetur faucibus eget eu arcu. Lorem ipsum dolor sit amet, consectetur adipiscing " +
		"elit. Curabitur vel aliquet massa. Integer id vehicula mi. Cras elementum, nisi in ultrices " +
		"mollis, dui tellus tristique neque, sed egestas nunc nibh sit amet quam. Suspendisse at maximus " +
		"odio, non ultricies felis. Nam maximus tortor condimentum egestas ultrices. Donec ac vehicula " +
		"nulla, at viverra neque. Etiam lobortis quis tellus ut ornare." +
		"Suspendisse eros metus, mattis at viverra sit amet, hendrerit non eros. Aenean nec sagittis " +
		"ligula. Sed a nisi ultrices, efficitur libero pellentesque, porta nibh. Orci varius natoque " +
		"penatibus et magnis dis parturient montes, nascetur ridiculous mus. Cras dictum, odio nec blandit " +
		"fermentum, arcu justo commodo orci, a varius massa erat ut est. Pellentesque venenatis placerat " +
		"efficitur. Donec dapibus ornare eleifend. Curabitur finibus convallis mauris eget posuere." +
		"Morbi ultricies rutrum nulla ut condimentum. Aliquam vulputate iaculis nisl at lacinia. Donec " +
		"ac ligula consequat, tempor elit ut, congue neque. Donec lobortis massa lorem, vitae mattis " +
		"neque mollis dignissim. Nulla facilities. Donec viverra purus a accumsan pellentesque. Proin " +
		"vestibulum accumsan erat vel commodo. Maecenas sapien mauris, faucibus nec consectetur eu, " +
		"ultricies sit amet elit. Suspendisse. "

	expected := "Lorem ipsum dolor sit amet, consectetur%adipiscing elit. Phasellus eu ex sit%amet " +
		"sapien consectetur faucibus eget eu%arcu. Lorem ipsum dolor sit amet,%consectetur adipiscing elit. " +
		"Curabitur%vel aliquet massa. Integer id vehicula%mi. Cras elementum, nisi in ultrices%mollis, dui " +
		"tellus tristique neque, sed%egestas nunc nibh sit amet quam.%Suspendisse at maximus odio, " +
		"non%ultricies felis. Nam maximus tortor%condimentum egestas ultrices. Donec ac%vehicula nulla, " +
		"at viverra neque. Etiam%lobortis quis tellus ut%ornare.Suspendisse eros metus, mattis at%viverra " +
		"sit amet, hendrerit non eros.%Aenean nec sagittis ligula. Sed a nisi%ultrices, efficitur libero " +
		"pellentesque,%porta nibh. Orci varius natoque%penatibus et magnis dis parturient%montes, nascetur " +
		"ridiculous mus. Cras%dictum, odio nec blandit fermentum, arcu%justo commodo orci, a varius massa " +
		"erat%ut est. Pellentesque venenatis placerat%efficitur. Donec dapibus ornare%eleifend. Curabitur " +
		"finibus convallis%mauris eget posuere.Morbi ultricies%rutrum nulla ut condimentum. Aliquam%vulputate " +
		"iaculis nisl at lacinia. Donec%ac ligula consequat, tempor elit ut,%congue neque. Donec lobortis " +
		"massa%lorem, vitae mattis neque mollis%dignissim. Nulla facilities. Donec%viverra purus a accumsan " +
		"pellentesque.%Proin vestibulum accumsan erat vel%commodo. Maecenas sapien mauris,%faucibus nec " +
		"consectetur eu, ultricies%sit amet elit. Suspendisse.%"

	actualTxt, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			40,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = StrMech{}.Ptr().ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'\n\n. Instead, text='%v'\n",
			expected, actualTxt)
	}
}

func TestStrOps_BreakTextAtLineLength_06(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_06()"

	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus eu ex sit amet " +
		"sapien consectetur faucibus eget eu arcu. Lorem ipsum dolor sit amet, consectetur adipiscing " +
		"elit. Curabitur vel aliquet massa. Integer id vehicula mi. Cras elementum, nisi in ultrices. "

	_, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			0,
			'\n',
			thisFuncName)

	if err == nil {
		t.Error("Error: Expected error return from StrMech{}.BreakTextAtLineLength(...). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_BreakTextAtLineLength_07(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_07()"

	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus eu ex sit amet " +
		"sapien consectetur faucibus eget eu arcu. Lorem ipsum dolor sit amet, consectetur adipiscing " +
		"elit. Curabitur vel aliquet massa. Integer id vehicula mi. Cras elementum, nisi in ultrices. "

	_, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			50,
			0,
			thisFuncName)

	if err == nil {
		t.Error("Error: Expected error return from StrMech{}.BreakTextAtLineLength(...). " +
			"NO ERROR RETURNED!")
	}
}

func TestStrOps_BreakTextAtLineLength_08(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_08()"

	tstStr := "                           "

	returnStr, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			10,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.BreakTextAtLineLength(...).\n"+
			"Error='%v'\n", err.Error())
	}

	if "\n" != returnStr {
		t.Errorf("Error: Expected returnStr= new line character.\n"+
			"Instead, returnStr='%v'\n", returnStr)
	}
}

func TestStrOps_BreakTextAtLineLength_09(t *testing.T) {

	thisFuncName := "TestStrOps_BreakTextAtLineLength_09()"

	tstStr := ""

	_, err := StrMech{}.NewPtr().
		BreakTextAtLineLength(
			tstStr,
			10,
			'\n',
			thisFuncName)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.BreakTextAtLineLength(tstStr, 10, '\\n')" +
			"because tstStr is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrOps_ConvertNonPrintableChars_01(t *testing.T) {

	testStr := "Hello world! How are you doing today?\n"
	testRunes := []rune(testStr)
	expectedStr := "Hello world! How are you doing today?\\n"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_02(t *testing.T) {

	testStr := "Hello world! How are you doing today?\n"
	testRunes := []rune(testStr)
	expectedStr := "Hello[SPACE]world![SPACE]How[SPACE]are[SPACE]you[SPACE]doing[SPACE]today?\\n"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, true)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_03(t *testing.T) {

	testStr := "Hello world!\tHow\rare\ayou\bdoing\ftoday?\v\n"
	testRunes := []rune(testStr)
	expectedStr := "Hello world!\\tHow\\rare\\ayou\\bdoing\\ftoday?\\v\\n"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_04(t *testing.T) {

	testStr := ""
	testRunes := []rune(testStr)
	expectedStr := "[EMPTY]"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_05(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		0}

	expectedStr := "Hello[NULL]"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_06(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		1,
		2,
		3,
		4,
		5,
		6}

	expectedStr := "Hello[SOH][STX][ETX][EOT][ENQ][ACK]"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_07(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		0x5c}

	expectedStr := "Hello\\"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_ConvertNonPrintableChars_08(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		14,
		15}

	expectedStr := "Hello[SO][SI]"

	actualStr := StrMech{}.NewPtr().ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrOps_CopyIn_01(t *testing.T) {

	string1 := "What in the world is Garfield doing!"
	string2 := "Now is the time for all good men to come to the aid of their country."
	string3 := "From this valley they say you are going."

	s2 := StrMech{StrOut: string2, StrIn: string1}
	s2.SetStringData(string3)

	s1 := StrMech{}

	s1.CopyIn(&s2)

	if string1 != s1.StrIn {
		t.Errorf("Error: expected s1.StrIn='%v'. Instead,  "+
			"s1.StrIn='%v'", string1, s1.StrIn)
	}

	if string2 != s1.StrOut {
		t.Errorf("Error: expected s1.StrOut='%v'. Instead,  "+
			"s1.StrOut='%v'", string2, s1.StrOut)
	}

	actualStr := s1.GetStringData()

	if string3 != actualStr {
		t.Errorf("Error: expected s1 StringData='%v'. Instead,  "+
			"s1 StringData='%v'", string3, actualStr)
	}

}

func TestStrOps_CopyOut_01(t *testing.T) {

	string1 := "What in the world is Garfield doing!"
	string2 := "Now is the time for all good men to come to the aid of their country."
	string3 := "From this valley they say you are going."

	s1 := StrMech{StrOut: string2, StrIn: string1}
	s1.SetStringData(string3)

	s2 := s1.CopyOut()

	if string1 != s2.StrIn {
		t.Errorf("Error: expected s2.StrIn='%v'. Instead,  "+
			"s2.StrIn='%v'", string1, s1.StrIn)
	}

	if string2 != s2.StrOut {
		t.Errorf("Error: expected s2.StrOut='%v'. Instead,  "+
			"s2.StrOut='%v'", string2, s2.StrOut)
	}

	actualStr := s2.GetStringData()

	if string3 != actualStr {
		t.Errorf("Error: expected s2 StringData='%v'. Instead,  "+
			"s2 StringData='%v'", string3, actualStr)
	}

}

func TestStrOps_DoesLastCharExist_01(t *testing.T) {

	var lastChar rune

	lastChar = '!'

	testString := "What in the world is Garfield doing!"

	actualReturn := StrMech{}.DoesLastCharExist(testString, lastChar)

	if true != actualReturn {
		t.Errorf("Expected return value='true'. Instead, return value='%v' ",
			actualReturn)
	}

}

func TestStrOps_DoesLastCharExist_02(t *testing.T) {

	var lastChar rune

	lastChar = 'x'

	testString := "What in the world is Garfield doing!"

	actualReturn := StrMech{}.DoesLastCharExist(testString, lastChar)

	if false != actualReturn {
		t.Errorf("Expected return value='false'. Instead, return value='%v' ",
			actualReturn)
	}

}

func TestStrOps_DoesLastCharExist_03(t *testing.T) {

	var lastChar rune

	lastChar = 'x'

	testString := ""

	actualReturn := StrMech{}.DoesLastCharExist(testString, lastChar)

	if false != actualReturn {
		t.Error("Expected return value='false' because 'testString' was an empty string\n" +
			"Instead, the actual return value was 'true'\n")
	}

}
