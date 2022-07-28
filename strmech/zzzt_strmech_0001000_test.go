package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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

func TestStrMech_BreakTextAtLineLength_01(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_01"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Lorem ipsum dolor sit amet, consectetur%adipiscing elit.%"

	sMech := StrMech{}

	actualTxt, err := sMech.BreakTextAtLineLength(
		tstStr,
		40,
		'\n',
		thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = new(StrMech).ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected string='%v'. Instead, string='%v'.",
			expected, actualTxt)
	}

	actualTxt, err = sMech.BreakTextAtLineLength(
		tstStr,
		40,
		'\n',
		strMechAtom{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"BreakTextAtLineLength()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			thisFuncName)

		return

	}

}

func TestStrMech_BreakTextAtLineLength_02(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_02"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Did you know? The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := new(StrMech).
		BreakTextAtLineLength(tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = new(StrMech).ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

func TestStrMech_BreakTextAtLineLength_03(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_03()"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "Did you know? XX The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? XX The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := new(StrMech).
		BreakTextAtLineLength(
			tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = new(StrMech).ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

func TestStrMech_BreakTextAtLineLength_04(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_04()"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	tstStr := "       Did you know? The Cow Jumped Over The Moon!"

	//         0         1         2         3         4         5
	//         012345678901234567890123456789012345678901234567890
	expected := "Did you know? The%" +
		"Cow Jumped Over The%Moon!%"

	actualTxt, err := new(StrMech).
		BreakTextAtLineLength(
			tstStr,
			20,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = new(StrMech).ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'. Instead, text='%v' ",
			expected, actualTxt)
	}
}

//goland:noinspection ALL
func TestStrMech_BreakTextAtLineLength_05(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_05()"

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

	actualTxt, err := new(StrMech).
		BreakTextAtLineLength(
			tstStr,
			40,
			'\n',
			thisFuncName)

	if err != nil {
		t.Errorf("Error returned from StrMech{}.BreakTextAtLineLength("+
			"tstStr, 40, '\n' ). Error='%v' ", err.Error())
	}

	actualTxt = new(StrMech).ReplaceNewLines(actualTxt, "%")

	if expected != actualTxt {
		t.Errorf("Error: Expected text='%v'\n\n. Instead, text='%v'\n",
			expected, actualTxt)
	}
}

func TestStrMech_BreakTextAtLineLength_06(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_06()"

	//goland:noinspection ALL
	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus eu ex sit amet " +
		"sapien consectetur faucibus eget eu arcu. Lorem ipsum dolor sit amet, consectetur adipiscing " +
		"elit. Curabitur vel aliquet massa. Integer id vehicula mi. Cras elementum, nisi in ultrices. "

	_, err := new(StrMech).
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

func TestStrMech_BreakTextAtLineLength_07(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_07()"

	//goland:noinspection ALL
	tstStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus eu ex sit amet " +
		"sapien consectetur faucibus eget eu arcu. Lorem ipsum dolor sit amet, consectetur adipiscing " +
		"elit. Curabitur vel aliquet massa. Integer id vehicula mi. Cras elementum, nisi in ultrices. "

	_, err := new(StrMech).
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

func TestStrMech_BreakTextAtLineLength_08(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_08()"

	tstStr := "                           "

	returnStr, err := new(StrMech).
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

func TestStrMech_BreakTextAtLineLength_09(t *testing.T) {

	thisFuncName := "TestStrMech_BreakTextAtLineLength_09()"

	tstStr := ""

	_, err := new(StrMech).
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

func TestStrMech_ConvertNonPrintableChars_01(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ConvertNonPrintableChars_01()",
		"")

	testStr := "Hello world! How are you doing today?\n"
	testRunes := []rune(testStr)
	expectedStr := "Hello world! How are you doing today?\\n"

	sMech := StrMech{}

	actualStr := sMech.ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)

		return
	}

	actualStr = sMech.ConvertNonPrintableChars(
		nil,
		false)

	expectedStr = "[EMPTY]"

	if expectedStr != actualStr {
		t.Errorf("%v\n"+
			"ERROR: sMech.ConvertNonPrintableChars()\n"+
			"Expected actualStr = '%v'\n"+
			"Instead, actualStr = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return

	}

	return
}

func TestStrMech_ConvertNonPrintableChars_02(t *testing.T) {

	testStr := "Hello world! How are you doing today?\n"
	testRunes := []rune(testStr)
	expectedStr := "Hello[SPACE]world![SPACE]How[SPACE]are[SPACE]you[SPACE]doing[SPACE]today?\\n"

	actualStr := new(StrMech).
		ConvertNonPrintableChars(testRunes, true)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_03(t *testing.T) {

	testStr := "Hello world!\tHow\rare\ayou\bdoing\ftoday?\v\n"
	testRunes := []rune(testStr)
	//goland:noinspection ALL
	expectedStr := "Hello world!\\tHow\\rare\\ayou\\bdoing\\ftoday?\\v\\n"

	actualStr := new(StrMech).
		ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_04(t *testing.T) {

	testStr := ""
	testRunes := []rune(testStr)
	expectedStr := "[EMPTY]"

	actualStr := new(StrMech).ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_05(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		0}

	expectedStr := "Hello[NULL]"

	actualStr := new(StrMech).ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_06(t *testing.T) {

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

	actualStr := new(StrMech).ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_07(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		0x5c}

	expectedStr := "Hello\\"

	actualStr := new(StrMech).ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_08(t *testing.T) {

	testRunes := []rune{
		'H',
		'e',
		'l',
		'l',
		'o',
		14,
		15}

	expectedStr := "Hello[SO][SI]"

	actualStr := new(StrMech).ConvertNonPrintableChars(testRunes, false)

	if expectedStr != actualStr {
		t.Errorf("ERROR: Expected result string='%v'\n"+
			"Instead, actual result string='%v'\n",
			expectedStr, actualStr)
	}

}

func TestStrMech_ConvertNonPrintableChars_09(t *testing.T) {

	tRunes := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	expectedStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	sMech := StrMech{}

	printableChars :=
		sMech.ConvertNonPrintableChars(
			tRunes,
			true)

	if printableChars != expectedStr {
		t.Errorf("ERROR:\n"+
			"Expected printableChars == expectedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"printableChars='%v'\n"+
			"expectedStr='%v'\n",
			printableChars,
			expectedStr)
	}

}

func TestStrMech_ConvertNonPrintableString_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ConvertNonPrintableString_000100()",
		"")

	tRunes := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	expectedStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	sMech := StrMech{}

	printableString :=
		sMech.ConvertNonPrintableString(
			string(tRunes),
			true)

	if printableString != expectedStr {
		t.Errorf("%v\n"+
			"ERROR:\n"+
			"sMech.ConvertNonPrintableString()\n"+
			"Expected printableString == expectedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"printableString='%v'\n"+
			"expectedStr='%v'\n",
			ePrefix.String(),
			printableString,
			expectedStr)

		return
	}

	printableString =
		sMech.ConvertNonPrintableString(
			"",
			true)

	expectedStr = "[EMPTY]"

	if printableString != expectedStr {
		t.Errorf("\n%v\n"+
			"ERROR:\n"+
			"sMech.ConvertNonPrintableString()\n"+
			"'nonPrintableString' parameter is empty.\n"+
			"Expected printableString == expectedStr\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"printableString='%v'\n"+
			"expectedStr='%v'\n",
			ePrefix.String(),
			printableString,
			expectedStr)

		return
	}

	return
}

func TestStrMech_ConvertPrintableChars_01(t *testing.T) {

	funcName := "TestStrMech_ConvertPrintableChars_01"

	nonPrintableRuneArray := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	printableCharsStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	sMech := StrMech{}

	runeArray,
		err :=
		sMech.ConvertPrintableChars(
			printableCharsStr,
			funcName)

	if err != nil {
		t.Errorf("Error:\n"+
			"Error returned from StrMech{}.ConvertPrintableChars()\n"+
			"Error = '%v'\n",
			err.Error())
		return
	}

	lenExpectedRuneArray := len(nonPrintableRuneArray)

	if lenExpectedRuneArray != len(runeArray) {
		t.Errorf("Error:\n"+
			"Expected lenExpectedRuneArray == len(runeArray).\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"lenExpectedRuneArray='%v'\n"+
			"      len(runeArray)='%v'\n",
			lenExpectedRuneArray,
			len(runeArray))
		return
	}

	for i := 0; i < len(nonPrintableRuneArray); i++ {
		if nonPrintableRuneArray[i] != runeArray[i] {
			t.Errorf("ERROR:\n"+
				"nonPrintableRuneArray[%v] != runeArray[%v]\n"+
				"nonPrintableRuneArray[%v]='%v'\n"+
				"runeArray[%v]='%v'\n",
				i,
				i,
				i,
				nonPrintableRuneArray[i],
				i,
				runeArray[i])
		}
	}

	_,
		err =
		sMech.ConvertPrintableChars(
			printableCharsStr,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ConvertPrintableChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			funcName)

		return
	}

	_,
		err =
		sMech.ConvertPrintableChars(
			"",
			funcName)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ConvertPrintableChars()\n"+
			"because 'printableChars' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			funcName)

		return
	}

	return
}

func TestStrMech_ConvertPrintableString_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ConvertPrintableString_000100()",
		"")

	nonPrintableRuneArray := []rune{
		0,    // [NULL]
		1,    // [SOH]
		2,    // [STX]
		3,    // [ETX]
		4,    // "[EOT]"
		5,    // [ENQ]
		6,    // [ACK]
		7,    // "\\a"
		8,    // "\\b"
		9,    // "\\t"
		0x0a, // "\\n"
		0x0b, // "\\v"
		0x0c, // "\\f"
		0x0d, // "\\r"
		0x0e, // "[SO]"
		0x0f, // "[SI]"
		0x5c, // "\\"
		0x20, // "[SPACE]"
	}

	printableCharsStr :=
		"[NULL]" +
			"[SOH]" +
			"[STX]" +
			"[ETX]" +
			"[EOT]" +
			"[ENQ]" +
			"[ACK]" +
			"\\a" +
			"\\b" +
			"\\t" +
			"\\n" +
			"\\v" +
			"\\f" +
			"\\r" +
			"[SO]" +
			"[SI]" +
			"\\" +
			"[SPACE]"

	sMech := StrMech{}

	var nonPrintableString string
	var err error

	nonPrintableString,
		err =
		sMech.ConvertPrintableString(
			printableCharsStr,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedRuneArray := len(nonPrintableRuneArray)

	runeArray := []rune(nonPrintableString)

	if lenExpectedRuneArray != len(runeArray) {
		t.Errorf("Error:\n"+
			"Expected lenExpectedRuneArray == len(runeArray).\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"lenExpectedRuneArray='%v'\n"+
			"      len(runeArray)='%v'\n",
			lenExpectedRuneArray,
			len(runeArray))
		return
	}

	for i := 0; i < len(nonPrintableRuneArray); i++ {
		if nonPrintableRuneArray[i] != runeArray[i] {
			t.Errorf("ERROR:\n"+
				"nonPrintableRuneArray[%v] != runeArray[%v]\n"+
				"nonPrintableRuneArray[%v]='%v'\n"+
				"runeArray[%v]='%v'\n",
				i,
				i,
				i,
				nonPrintableRuneArray[i],
				i,
				runeArray[i])
		}
	}

	_,
		err =
		sMech.ConvertPrintableString(
			printableCharsStr,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ConvertPrintableString()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		sMech.ConvertPrintableString(
			"",
			ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ConvertPrintableString()\n"+
			"because 'printableString' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStrMech_CopyIn_01(t *testing.T) {

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

	s1.CopyIn(nil)
}

func TestStrMech_CopyRuneArrays000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000100()",
		"")

	var targetRuneArray []rune

	sourceRuneArray := []rune("Hello World!")

	err := new(StrMech).CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if len(targetRuneArray) == 0 {
		t.Errorf("%v - Error\n"+
			"After copy operation the length of\n"+
			"the target rune array is Zero!\n",
			ePrefix.String())

		return
	}

	lenSrcArray := len(sourceRuneArray)

	if len(targetRuneArray) != lenSrcArray {
		t.Errorf("%v - Error\n"+
			"len(targetRuneArray) != len(sourceRuneArray)\n"+
			"len(targetRuneArray)= '%v'\n"+
			"len(sourceRuneArray)= '%v'\n",
			ePrefix.String(),
			len(targetRuneArray),
			len(sourceRuneArray))

		return
	}

	for i := 0; i < lenSrcArray; i++ {
		if targetRuneArray[i] != sourceRuneArray[i] {
			t.Errorf("%v - Error\n"+
				"targetRuneArray[i] != sourceRuneArray[i]\n"+
				"i='%v'\n"+
				"sourceRuneArray[i]='%v'\n"+
				"sourceRuneArray='%v'\n",
				ePrefix.String(),
				i,
				string(sourceRuneArray[i]),
				string(sourceRuneArray))

			return
		}
	}

	err = new(StrMech).CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		false,
		strMechAtom{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"CopyOut()\n"+
			"because 'errorPrefix' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStrMech_CopyRuneArrays000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000200()",
		"")

	targetRuneArray := []rune("Hello World")

	sourceRuneArray := []rune("Good Morning Universe!")

	sMechPreon := strMechPreon{}

	err := sMechPreon.copyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		false,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if len(targetRuneArray) == 0 {
		t.Errorf("%v - Error\n"+
			"After copy operation the length of\n"+
			"the target rune array is Zero!\n",
			ePrefix.String())

		return
	}

	lenSrcArray := len(sourceRuneArray)

	if len(targetRuneArray) != lenSrcArray {
		t.Errorf("%v - Error\n"+
			"len(targetRuneArray) != len(sourceRuneArray)\n"+
			"len(targetRuneArray)= '%v'\n"+
			"len(sourceRuneArray)= '%v'\n",
			ePrefix.String(),
			len(targetRuneArray),
			len(sourceRuneArray))

		return
	}

	areEqual := new(strMechQuark).equalRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyRuneArrays000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000300()",
		"")

	targetRuneArray := []rune("Hello World")

	var sourceRuneArray = make([]rune, 0)

	sMechPreon := strMechPreon{}

	err := sMechPreon.copyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		true,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if targetRuneArray != nil {
		t.Errorf("%v - Error\n"+
			"After copy operation, expected that\n"+
			"'targetRuneArray' == nil. Instead,\n"+
			"len(targetRuneArray) = '%v'\n",
			ePrefix.String(),
			len(targetRuneArray))

		return
	}

	areEqual := new(strMechQuark).equalRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if areEqual {
		t.Errorf("%v - Error\n"+
			"'targetRuneArray' and 'sourceRuneArray' are EQUAL!"+
			"This is wrong. 'targetRuneArray' should be 'nil'"+
			"and sourceRuneArray should be a zero length array!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyRuneArrays000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000400()",
		"")

	var targetRuneArray []rune = nil

	expectedStr := "How now brown cow!\n"

	var sourceRuneArray = []rune(expectedStr)

	sMech := StrMech{}

	err := sMech.CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		true,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	sMech2 := StrMech{}

	areEqual := sMech2.EqualRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"sMech2.EqualRuneArraysNil()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
		return
	}

	actualStr := string(targetRuneArray)

	if expectedStr != actualStr {
		t.Errorf("%v - Error\n"+
			"Expected that targetRuneArray string == sourceRuneArray string\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"Expected string =        '%v'\n"+
			"targetRuneArray string = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	actualStr = string(sourceRuneArray)

	if expectedStr != actualStr {
		t.Errorf("%v - Error\n"+
			"Expected that sourceRuneArray string == original rune array string\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"Original Rune Array string = '%v'\n"+
			"sourceRuneArray string     = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	areEqual = new(StrMech).EqualRuneArrays(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"new(StrMech).EqualRuneArrays()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyRuneArrays000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000500()",
		"")

	var targetRuneArray = []rune("How now brown cow!")

	var sourceRuneArray = make([]rune, 0)

	sMech := StrMech{}

	err := sMech.CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		false,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if targetRuneArray == nil {
		t.Errorf("%v - Error\n"+
			"targetRuneArray == nil!\n"+
			"This is WRONG!\n"+
			"targetRuneArray should be equal to make( []rune, 0)!\n",
			ePrefix.String())

		return
	}

	if len(targetRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(targetRuneArray) > 0! This is WRONG!\n"+
			"len(targetRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(targetRuneArray) = '%v'\n"+
			"targetRuneArray string = '%v'\n",
			ePrefix.String(),
			len(targetRuneArray),
			string(targetRuneArray))

		return
	}

	if sourceRuneArray == nil {
		t.Errorf("%v - Error\n"+
			"sourceRuneArray == nil!\n"+
			"This is WRONG!\n"+
			"sourceRuneArray should be equal to make( []rune, 0)!\n",
			ePrefix.String())

		return
	}

	if len(sourceRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(sourceRuneArray) > 0! This is WRONG!\n"+
			"len(sourceRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(sourceRuneArray) = '%v'\n"+
			"sourceRuneArray string = '%v'\n",
			ePrefix.String(),
			len(sourceRuneArray),
			string(sourceRuneArray))

		return
	}

	sMech2 := StrMech{}

	areEqual := sMech2.EqualRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"sMech2.EqualRuneArraysNil()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
		return
	}

	areEqual = new(StrMech).EqualRuneArrays(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"new(StrMech).EqualRuneArrays()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyRuneArrays000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000600()",
		"")

	var targetRuneArray = []rune("How now brown cow!")

	var sourceRuneArray = make([]rune, 0)

	sMech := StrMech{}

	err := sMech.CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		true,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if targetRuneArray != nil {
		t.Errorf("%v - Error\n"+
			"targetRuneArray != nil!\n"+
			"This is WRONG!\n"+
			"targetRuneArray should be equal to 'nil'"+
			"because 'setZeroLenArrayToNil == 'true'!\n",
			ePrefix.String())

		return
	}

	if len(targetRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(targetRuneArray) > 0! This is WRONG!\n"+
			"len(targetRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(targetRuneArray) = '%v'\n"+
			"targetRuneArray string = '%v'\n",
			ePrefix.String(),
			len(targetRuneArray),
			string(targetRuneArray))

		return
	}

	if sourceRuneArray == nil {
		t.Errorf("%v - Error\n"+
			"sourceRuneArray == nil!\n"+
			"This is WRONG!\n"+
			"sourceRuneArray should be equal to make([]rune, 0)!\n",
			ePrefix.String())

		return
	}

	if len(sourceRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(sourceRuneArray) > 0! This is WRONG!\n"+
			"len(sourceRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(sourceRuneArray) = '%v'\n"+
			"sourceRuneArray string = '%v'\n",
			ePrefix.String(),
			len(sourceRuneArray),
			string(sourceRuneArray))

		return
	}

	sMech2 := StrMech{}

	areEqual := sMech2.EqualRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if areEqual {
		t.Errorf("%v - Error\n"+
			"sMech2.EqualRuneArraysNil()\n"+
			"'targetRuneArray' and 'sourceRuneArray' ARE EQUAL!"+
			"This is wrong. 'targetRuneArray' is 'nil' and\n"+
			"sourceRuneArray = make([]rune, 0)!!\n",
			ePrefix.String())
		return
	}

	areEqual = new(StrMech).EqualRuneArrays(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"new(StrMech).EqualRuneArrays()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyRuneArrays000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_CopyRuneArrays000600()",
		"")

	var targetRuneArray []rune = nil

	var sourceRuneArray = make([]rune, 0)

	sMech := StrMech{}

	err := sMech.CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		false,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if targetRuneArray == nil {
		t.Errorf("%v - Error\n"+
			"targetRuneArray == nil!\n"+
			"This is WRONG!\n"+
			"targetRuneArray should be equal to 'make([]rune, 0)'"+
			"because 'setZeroLenArrayToNil == 'false'!\n",
			ePrefix.String())

		return
	}

	if len(targetRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(targetRuneArray) > 0! This is WRONG!\n"+
			"len(targetRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(targetRuneArray) = '%v'\n"+
			"targetRuneArray string = '%v'\n",
			ePrefix.String(),
			len(targetRuneArray),
			string(targetRuneArray))

		return
	}

	if sourceRuneArray == nil {
		t.Errorf("%v - Error\n"+
			"sourceRuneArray == nil!\n"+
			"This is WRONG!\n"+
			"sourceRuneArray should be equal to make([]rune, 0)!\n",
			ePrefix.String())

		return
	}

	if len(sourceRuneArray) != 0 {
		t.Errorf("%v - Error\n"+
			"len(sourceRuneArray) > 0! This is WRONG!\n"+
			"len(sourceRuneArray) SHOULD EQUAL '0'!\n"+
			"Actual len(sourceRuneArray) = '%v'\n"+
			"sourceRuneArray string = '%v'\n",
			ePrefix.String(),
			len(sourceRuneArray),
			string(sourceRuneArray))

		return
	}

	sMech2 := StrMech{}

	areEqual := sMech2.EqualRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"sMech2.EqualRuneArraysNil()\n"+
			"'targetRuneArray' and 'sourceRuneArray' ARE NOT EQUAL!"+
			"This is wrong. Both 'targetRuneArray' and\n"+
			"sourceRuneArray should be equal to make([]rune, 0)!!\n",
			ePrefix.String())
		return
	}

	areEqual = new(StrMech).EqualRuneArrays(
		targetRuneArray,
		sourceRuneArray)

	if !areEqual {
		t.Errorf("%v - Error\n"+
			"new(StrMech).EqualRuneArrays()\n"+
			"'targetRuneArray' and 'sourceRuneArray' are NOT EQUAL!"+
			"This is wrong. 'targetRuneArray' should be equal\n"+
			"to sourceRuneArray!\n",
			ePrefix.String())
	}

}

func TestStrMech_CopyOut_01(t *testing.T) {

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

	sMech := StrMech{}

	_ = sMech.CopyOut()
}

func TestStrMech_CopyOut_02(t *testing.T) {

	string1 := "What in the world is Garfield doing!"
	string2 := "Now is the time for all good men to come to the aid of their country."
	string3 := "From this valley they say you are going."

	s1 := StrMech{}

	s1.StrOut = string2
	s1.StrIn = string1

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

func TestStrMech_DoesLastCharExist_01(t *testing.T) {

	var lastChar rune

	lastChar = '!'

	testString := "What in the world is Garfield doing!"

	actualReturn := new(StrMech).DoesLastCharExist(testString, lastChar)

	if true != actualReturn {
		t.Errorf("Expected return value='true'. Instead, return value='%v' ",
			actualReturn)
	}

}

func TestStrMech_DoesLastCharExist_02(t *testing.T) {

	var lastChar rune

	lastChar = 'x'

	testString := "What in the world is Garfield doing!"

	actualReturn := new(StrMech).DoesLastCharExist(testString, lastChar)

	if false != actualReturn {
		t.Errorf("Expected return value='false'. Instead, return value='%v' ",
			actualReturn)
	}

}

func TestStrMech_DoesLastCharExist_03(t *testing.T) {

	var lastChar rune

	lastChar = 'x'

	testString := ""

	actualReturn := new(StrMech).DoesLastCharExist(testString, lastChar)

	if false != actualReturn {
		t.Error("Expected return value='false' because 'testString' was an empty string\n" +
			"Instead, the actual return value was 'true'\n")
	}

}
