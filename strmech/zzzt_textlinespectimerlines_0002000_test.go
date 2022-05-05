package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
	"time"
)

func TestTextLineSpecTimerLines_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, lenExpectedStr+1)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"timerLines01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: timerLines01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000200()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 1)

	var actualStr string

	var n, readBytesCnt int

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: timerLines01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"timerLines01.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.String())
		return
	}

	p = make([]byte, 200)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000300()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	txtSpecAtom := textSpecificationAtom{}

	var n int
	p := make([]byte, 100)

	n,
		err = txtSpecAtom.readBytes(
		nil,
		p,
		ePrefix.XCpy("textReader == 'nil'"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'textReader' == 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var formattedTxtStr string
	timerLinesMolecule := textLineSpecTimerLinesMolecule{}

	formattedTxtStr,
		err =
		timerLinesMolecule.getFormattedText(
			timerLines01,
			ePrefix.XCpy("timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	timerLines01.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		timerLines01.textLineReader,
		p,
		ePrefix.XCpy("p == zero length"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'p' is a zero length byte array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 100)

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtSpecAtom.readBytes(
			timerLines01.textLineReader,
			p,
			ePrefix.XCpy("timerLines01 is valid"))

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSpecAtom.readBytes(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtSpecAtom.readBytes(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000400()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	p := make([]byte, 15)

	var n, readBytesCnt int
	sb := strings.Builder{}
	sb.Grow(512)

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		sb.Write(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
	}

	return
}

func TestTextLineSpecTimerLines_Read_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000500()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	badEndTime := timerLines01.startTime
	timerLines01.startTime = timerLines01.endTime
	timerLines01.endTime = badEndTime
	timerLines01.textLabelFieldLen = -99
	timerLines01.textLabelJustification = -99

	p := make([]byte, 15)

	_,
		err = timerLines01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01.Read(p)\n"+
			"because 'timerLines01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000600()",
		"")

	p := make([]byte, 5)

	timerLines01 := TextLineSpecTimerLines{}

	_,
		err := timerLines01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01.Read(p)\n"+
			"because 'timerLines01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecTimerLines_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_ReaderInitialize_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 5)

	var n int

	n,
		err = timerLines01.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by timerLines01.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.String(),
			n)

		return
	}

	p = make([]byte, 200)

	timerLines01.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but timerLines01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 200)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but timerLines01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	timerLines02.ReaderInitialize()

	return
}

func TestTextLineSpecTimerLines_SetDefaultFullTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetDefaultFullTimerEvent_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.SetDefaultFullTimerEvent(
		timerLines01.startTime,
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetDefaultFullTimerEvent(
		timerLines01.startTime,
		timerLines01.endTime,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetDefaultFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetDefaultShellTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetDefaultShellTimerEvent_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	timerLines02.SetDefaultShellTimerEvent()

	err = timerLines02.SetStartAndEndTime(
		timerLines01.startTime,
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetEndTimeLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetEndTimeLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.CopyIn(timerLines01,
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newEndTimeLabel := "Xray-97"

	timerLines01.endTimeLabel = []rune(newEndTimeLabel)

	err = timerLines02.SetEndTimeLabel(
		newEndTimeLabel,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.CopyIn(timerLines01,
		ePrefix.XCpy(
			"timerLines03<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines03.SetEndTimeLabel(
		newEndTimeLabel,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetDefaultFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines04 := TextLineSpecTimerLines{}

	err = timerLines04.SetEndTimeLabel(
		newEndTimeLabel,
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLines_SetEndTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetEndTime_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	timerLines02.SetDefaultShellTimerEvent()

	err = timerLines02.SetStartTime(
		timerLines01.startTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines02.SetEndTime(
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.CopyIn(timerLines01,
		ePrefix.XCpy(
			"timerLines03<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines03.SetEndTime(
		timerLines01.endTime,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetDefaultFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines04 := TextLineSpecTimerLines{}

	err = timerLines04.SetEndTime(
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLines_SetFullTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetFullTimerEvent_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.SetFullTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		timerLines01.startTime,
		string(timerLines01.endTimeLabel),
		timerLines01.endTime,
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetFullTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		timerLines01.startTime,
		string(timerLines01.endTimeLabel),
		timerLines01.endTime,
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetFullTimerEventRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetFullTimerEventRunes_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.SetFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetFullTimerEventRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetShellTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetShellTimerEvent_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.SetShellTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		string(timerLines01.endTimeLabel),
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines02.SetStartAndEndTime(
		timerLines01.startTime,
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' would be equal to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetShellTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		string(timerLines01.endTimeLabel),
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.SetShellTimerEvent()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetStartTimeLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetStartTimeLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		[]rune("Beginning Time"),
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if timerLines01.Equal(timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = timerLines02.SetStartTimeLabel(
		string(timerLines01.startTimeLabel),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' WOULD BE EQUAL to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetStartTimeLabel(
		string(timerLines01.startTimeLabel),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.SetStartTimeLabel()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetStartTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetStartTime_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"\n%v - ERROR\n"+
				"time.LoadLocation(\"America/Chicago\")"+
				"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	newStartTime := time.Date(
		2022,
		4,
		5,
		9,
		0,
		0,
		0,
		loc)

	timerLines01.startTime = newStartTime

	err = timerLines02.SetStartTime(
		newStartTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if !timerLines02.startTime.Equal(
		timerLines01.startTime) {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines02.startTime == timerLines01.startTime\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if !timerLines02.startTime.Equal(newStartTime) {

		t.Errorf("%v\n"+
			"Test Inspection #4\n"+
			"Error: Expected timerLines02.startTime == newStartTime\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetStartTime(
		newStartTime,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines03.SetStartTime()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetStartAndEndTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetStartAndEndTime_000100()",
		"")

	var loc *time.Location
	var err error

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	timerLines01 :=
		TextLineSpecTimerLines{}.NewDefaultShellTimerEvent()

	err = timerLines01.SetStartAndEndTime(
		endTime,
		startTime,
		ePrefix.XCpy(
			"timerLines01 start and end times reversed"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines01.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"start and end times are invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines02 TextLineSpecTimerLines

	err = timerLines02.SetStartAndEndTime(
		startTime,
		endTime,
		ePrefix.XCpy(
			"timerLines02 is empty"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetStartAndEndTime(
		time.Time{},
		endTime,
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: timerLines03.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"'startTime' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	timerLines04 := TextLineSpecTimerLines{}

	err = timerLines04.SetStartAndEndTime(
		startTime,
		time.Time{},
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: timerLines04.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"'endTime' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetLabelFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetLabelFieldLength_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	newLabelFieldLen := 75

	timerLines01.textLabelFieldLen = newLabelFieldLen

	timerLines02.SetLabelFieldLength(newLabelFieldLen)

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if timerLines02.textLabelFieldLen !=
		timerLines01.textLabelFieldLen {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines02.textLabelFieldLen\n"+
			"to be EQUAL TO timerLines01.textLabelFieldLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if timerLines02.textLabelFieldLen != newLabelFieldLen {

		t.Errorf("%v\n"+
			"Test Inspection #4\n"+
			"Error: Expected timerLines02.textLabelFieldLen\n"+
			"to be EQUAL TO newLabelFieldLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetLabelJustification_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetLabelJustification_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines01.textLabelJustification = TxtJustify.Left()

	err = timerLines02.SetLabelJustification(
		TxtJustify.Left(),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if timerLines02.textLabelJustification !=
		timerLines01.textLabelJustification {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines02.textLabelJustification\n"+
			"to be EQUAL TO timerLines01.textLabelJustification\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetLabelJustification(
		TxtJustify.Left(),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines03.SetLabelJustification()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetLabelRightMarginChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetLabelRightMarginChars_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines01.labelRightMarginChars = []rune("- ")

	err = timerLines02.SetLabelRightMarginChars(
		"- ",
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		timerLines01.labelRightMarginChars,
		timerLines02.labelRightMarginChars)

	if !areEqual {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines02.labelRightMarginChars\n"+
			"to be EQUAL TO timerLines01.labelRightMarginChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetLabelRightMarginChars(
		"- ",
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines03.SetLabelRightMarginChars()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetTimeDurationLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetTimeDurationLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	newTimeDurationLabel := "Time Duration"

	timerLines01.timeDurationLabel = []rune(newTimeDurationLabel)

	err = timerLines02.SetTimeDurationLabel(
		newTimeDurationLabel,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		timerLines01.timeDurationLabel,
		timerLines02.timeDurationLabel)

	if !areEqual {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines02.timeDurationLabel\n"+
			"to be EQUAL TO timerLines01.timeDurationLabel\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetTimeDurationLabel(
		newTimeDurationLabel,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines03.SetTimeDurationLabel()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetTimeFormat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetTimeFormat_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Setup #1\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	newTimeFormat := "Monday 2006-01-02 15:04:05.000000000 -0700 MST"

	timerLines01.timeFormat = newTimeFormat

	timerLines02.SetTimeFormat(
		newTimeFormat)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if timerLines01.timeFormat != timerLines02.timeFormat {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines01.timeFormat\n"+
			"to be EQUAL TO timerLines02.timeFormat\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_String_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr = timerLines01.String()

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.String()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_String_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_String_000200()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr := timerLines01.String()

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.String()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_String_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_String_000300()",
		"")

	timerLines01 := TextLineSpecTimerLines{}

	actualStr := timerLines01.String()

	containsError := strings.Contains(
		actualStr, "Error")

	if !containsError {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01.String()\n"+
			"because 'timerLines01' is empty.\n"+
			"HOWEVER, THE RETURNED STRING DOES NOT CONTAIN THE WORD 'Error'!\n",
			ePrefix.String())

		return
	}

	return
}
