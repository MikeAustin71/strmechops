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
		_,
		_,
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
			"SetEndTimeLabel()\n"+
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

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	totalLabelLength := timerLines05.GetLengthTotalLabel()

	longestLabelLength := timerLines05.GetLengthLongestTextLabel()

	maximumLabelLength :=
		timerLines05.GetMaximumTextLabelLength()

	adjustedTotalTableLength :=
		totalLabelLength - longestLabelLength

	newEndTimeLabelCharCount :=
		maximumLabelLength -
			adjustedTotalTableLength + 1

	newEndTimeLabel = strings.Repeat("X", newEndTimeLabelCharCount)

	err = timerLines05.SetEndTimeLabel(
		newEndTimeLabel,
		ePrefix.XCpy(
			"timerLines05"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines05."+
			"SetEndTimeLabel()\n"+
			"because 'newEndTimeLabel' is invalid.\n"+
			"'newEndTimeLabel' exceeds the maximum Field Length!\n"+
			"Length of 'newEndTimeLabel' = '%v'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			len(newEndTimeLabel))

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetEndTimeLabel_000200(t *testing.T) {

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

	sMechPreon := strMechPreon{}

	var newEndTimeLabelRunesToLong []rune

	newEndTimeLabelRunesToLong,
		err = sMechPreon.getRepeatRuneChar(
		94,
		'X',
		ePrefix.XCpy(
			"newEndTimeLabelRunes"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newEndTimeLabelStr := string(newEndTimeLabelRunesToLong)

	err = timerLines01.SetEndTimeLabel(
		newEndTimeLabelStr,
		ePrefix.XCpy(
			"timerLines01"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetEndTimeLabel()\n"+
			"because 'endTimeLabel' exceeds maximum allowable length.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var timerLines02 *TextLineSpecTimerLines

	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines02.SetEndTimeLabel(
		"",
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines03 *TextLineSpecTimerLines

	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines03.SetEndTimeLabel(
		"Xray",
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines03.SetEndTimeLabel()\n"+
			"Expected an error return because the error prefix\n"+
			"is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

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
			"SetEndTime()\n"+
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

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newZeroTime := time.Time{}

	err = timerLines05.SetEndTime(
		newZeroTime,
		ePrefix.XCpy(
			"timerLines05"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetEndTime()\n"+
			"because 'endTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

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

	var timerLines04 *TextLineSpecTimerLines

	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines04.SetLabelJustification(
		-9,
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines04.SetLabelJustification()\n"+
			"Expected an error return because\n"+
			"input parameter 'labelJustification' is invalid.\n"+
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

	var timerLines04 *TextLineSpecTimerLines

	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines04.SetLabelRightMarginChars(
		"",
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
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

	expectedStr := string(timerLines01.startTimeLabel)

	actualStr := string(timerLines02.startTimeLabel)

	if expectedStr != actualStr {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01.startTimeLabel' \n"+
			"WOULD BE EQUAL to 'timerLines02.startTimeLabel'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

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

	sMechPreon := strMechPreon{}

	var newStartTimeLabelRunesToLong []rune

	newStartTimeLabelRunesToLong,
		err = sMechPreon.getRepeatRuneChar(
		94,
		'X',
		ePrefix.XCpy(
			"newStartTimeLabelRunesToLong"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newStartTimeLabelStr := string(newStartTimeLabelRunesToLong)

	var timerLines04 *TextLineSpecTimerLines

	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines04.SetStartTimeLabel(
		newStartTimeLabelStr,
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error timerLines04.SetStartTimeLabel()\n"+
			"Expected an error return because\n"+
			"input parameter 'startTimeLabel' exceeds maximum string length.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines05.SetStartTimeLabel(
		"",
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLines_SetStartTimeLabel_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetStartTimeLabel_000200()",
		"")

	_,
		timerLines05,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	totalLabelLength := timerLines05.GetLengthTotalLabel()

	longestLabelLength := timerLines05.GetLengthLongestTextLabel()

	maximumLabelLength :=
		timerLines05.GetMaximumTextLabelLength()

	adjustedTotalTableLength :=
		totalLabelLength - longestLabelLength

	newEndTimeLabelCharCount :=
		maximumLabelLength -
			adjustedTotalTableLength + 1

	newStartTimeLabel := strings.Repeat("X", newEndTimeLabelCharCount)

	err = timerLines05.SetStartTimeLabel(
		newStartTimeLabel,
		ePrefix.XCpy(
			"timerLines05"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines05."+
			"SetStartTimeLabel()\n"+
			"because input parameter 'newStartTimeLabel' is invalid.\n"+
			"'newStartTimeLabel' exceeds the maximum Field Length!\n"+
			"Length of 'newStartTimeLabel' = '%v'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			len(newStartTimeLabel))

		return
	}

	var timerLines06 *TextLineSpecTimerLines
	_,
		timerLines06,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newEndTimeLabelCharCount = 12

	newStartTimeLabel = strings.Repeat("X", newEndTimeLabelCharCount)

	timerLines06.textLabelFieldLen = 11

	err = timerLines06.SetStartTimeLabel(
		newStartTimeLabel,
		ePrefix.XCpy(
			"timerLines06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines07 *TextLineSpecTimerLines
	_,
		timerLines07,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newEndTimeLabelCharCount = 5

	newStartTimeLabel = strings.Repeat("X", newEndTimeLabelCharCount)

	err = timerLines07.SetStartTimeLabel(
		newStartTimeLabel,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines07.SetStartTimeLabel()\n"+
			"Expected an error return because the error prefix\n"+
			"is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
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

	newZeroTime := time.Time{}

	var timerLines04 *TextLineSpecTimerLines

	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines04.SetStartTime(
		newZeroTime,
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetStartTime()\n"+
			"because 'startTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
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

	var timerLines05 *TextLineSpecTimerLines
	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines05.SetStartAndEndTime(
		startTime,
		endTime,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: timerLines05.SetStartAndEndTime()\n"+
			"Expected an error return because input\n"+
			"parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetTextLabelFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetTextLabelFieldLength_000100()",
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

	err = timerLines01.SetTextLabelFieldLength(
		-2,
		ePrefix.XCpy(
			"timerLines01 #1"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetFullTimerEvent()\n"+
			"because 'fieldLength' is invalid.\n"+
			"'fieldLength' == -2\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = timerLines01.SetTextLabelFieldLength(
		1000001,
		ePrefix.XCpy(
			"timerLines01 #2"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetFullTimerEvent()\n"+
			"because 'fieldLength' is invalid.\n"+
			"'fieldLength' == 1,000,001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var timerLines02 *TextLineSpecTimerLines
	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	longestTxtLabelLen :=
		timerLines02.GetLengthLongestTextLabel()

	err = timerLines02.SetTextLabelFieldLength(
		-1,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualTextFieldLength :=
		timerLines02.textLabelFieldLen

	if longestTxtLabelLen != actualTextFieldLength {

		t.Errorf("\n%v - ERROR\n"+
			"timerLines02.SetTextLabelFieldLength()\n"+
			"Expected Length of Longest Text Label would"+
			"be equal to actual Text Label Field length.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"Length of Longest Text Label   = '%v'\n"+
			"Actual Text Label Field length = '%v'\n",
			ePrefix.String(),
			longestTxtLabelLen,
			actualTextFieldLength)

		return

	}

	var timerLines03 *TextLineSpecTimerLines
	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	longestTxtLabelLen =
		timerLines03.GetLengthLongestTextLabel()

	newTextLabelFieldLength := longestTxtLabelLen + 10

	err = timerLines03.SetTextLabelFieldLength(
		newTextLabelFieldLength,
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualTextFieldLength =
		timerLines03.textLabelFieldLen

	if newTextLabelFieldLength != actualTextFieldLength {

		t.Errorf("\n%v - ERROR\n"+
			"timerLines03.SetTextLabelFieldLength()\n"+
			"Expected Text Label Field length would be\n"+
			"equal to the NEW Text Label Field length.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"New Text Label Field length    = '%v'\n"+
			"Actual Text Label Field length = '%v\n"+
			"Lenght Of Longest Label        = '%v'\n",
			ePrefix.String(),
			newTextLabelFieldLength,
			actualTextFieldLength,
			longestTxtLabelLen)

		return

	}

	timerLines04 := TextLineSpecTimerLines{}

	newTextLabelFieldLength = 15

	err = timerLines04.SetTextLabelFieldLength(
		newTextLabelFieldLength,
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualTextFieldLength =
		timerLines04.textLabelFieldLen

	if newTextLabelFieldLength != actualTextFieldLength {

		t.Errorf("\n%v - ERROR\n"+
			"timerLines04.SetTextLabelFieldLength()\n"+
			"Expected Text Label Field length would be\n"+
			"equal to the NEW Text Label Field length.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"New Text Label Field length    = '%v'\n"+
			"Actual Text Label Field length = '%v\n"+
			"Lenght Of Longest Label        = '%v'\n",
			ePrefix.String(),
			newTextLabelFieldLength,
			actualTextFieldLength,
			longestTxtLabelLen)

		return

	}

	return
}

func TestTextLineSpecTimerLines_SetTextLabelFieldLength_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetTextLabelFieldLength_000200()",
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

	err = timerLines02.SetTextLabelFieldLength(
		newLabelFieldLen,
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

		t.Errorf("\n%v\n"+
			"Test Inspection #4\n"+
			"Error: Expected timerLines02.textLabelFieldLen\n"+
			"to be EQUAL TO newLabelFieldLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	var timerLines03 *TextLineSpecTimerLines
	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines03.SetTextLabelFieldLength(
		1000001,
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Test Inspection #5\n"+
			"timerLines03.SetTextLabelFieldLength()\n"+
			"Expected an error return because input paramter\n"+
			"'labelFieldLength' is greater than 1,000,000.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var timerLines04 *TextLineSpecTimerLines
	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		timerLines04.SetTextLabelFieldLength(
			-2,
			ePrefix.XCpy(
				"timerLines04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Test Inspection #6\n"+
			"timerLines04.SetTextLabelFieldLength()\n"+
			"Expected an error returen because input parameter\n"+
			"'labelFieldLength' is less than -1.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines05 := TextLineSpecTimerLines{}

	err = timerLines05.SetTextLabelFieldLength(
		-1,
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines06 *TextLineSpecTimerLines
	_,
		timerLines06,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines06.SetTextLabelFieldLength(
		15,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines06.SetTextLabelFieldLength()\n"+
			"Expected an error return because the error prefix\n"+
			"is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
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

	expectedTimeDurationLabel := "Time Duration"

	timerLines01.timeDurationLabel =
		[]rune(expectedTimeDurationLabel)

	err = timerLines02.SetTimeDurationLabel(
		expectedTimeDurationLabel,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualTimeDurationLabel :=
		string(timerLines02.timeDurationLabel)

	if expectedTimeDurationLabel != actualTimeDurationLabel {

		t.Errorf("\n%v\n"+
			"Test Inspection #2\n"+
			"Error: Expected Time Duration Label is\n"+
			"NOT EQUAL TO actual Time Duration Label\n"+
			"Expected Time Duration Label = '%v'\n"+
			"  Actual Time Duration Label = '%v'\n",
			ePrefix.String(),
			expectedTimeDurationLabel,
			actualTimeDurationLabel)

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
		expectedTimeDurationLabel,
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

	var timerLines04 *TextLineSpecTimerLines

	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var newTimeDurationLabelRunesToLong []rune

	newTimeDurationLabelRunesToLong,
		err = sMechPreon.getRepeatRuneArray(
		10,
		[]rune("Xray94"),
		ePrefix.XCpy(
			"newEndTimeLabelRunes"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines04.SetTimeDurationLabel(
		string(newTimeDurationLabelRunesToLong),
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetTimeDurationLabel()\n"+
			"because input parameter 'timeDurationLabel' exceeds\n"+
			"the maximum allowable string length.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines05.SetTimeDurationLabel(
		"",
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLines_SetTimeDurationLabel_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetTimeDurationLabel_000200()",
		"")
	_,
		timerLines05,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	totalLabelLength := timerLines05.GetLengthTotalLabel()

	longestLabelLength := timerLines05.GetLengthLongestTextLabel()

	maximumLabelLength :=
		timerLines05.GetMaximumTextLabelLength()

	adjustedTotalTableLength :=
		totalLabelLength - longestLabelLength

	newStartTimeLabelCharCount :=
		maximumLabelLength -
			adjustedTotalTableLength + 1

	newStartTimeLabel := strings.Repeat("X", newStartTimeLabelCharCount)

	err = timerLines05.SetTimeDurationLabel(
		newStartTimeLabel,
		ePrefix.XCpy(
			"timerLines05"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines05."+
			"SetTimeDurationLabel()\n"+
			"because input parameter 'newStartTimeLabel' is invalid.\n"+
			"'newStartTimeLabel' exceeds the maximum Field Length!\n"+
			"Length of 'newStartTimeLabel' = '%v'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			len(newStartTimeLabel))

		return
	}

	var timerLines07 *TextLineSpecTimerLines
	_,
		timerLines07,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newStartTimeLabelCharCount = 5

	newStartTimeLabel = strings.Repeat("X", newStartTimeLabelCharCount)

	err = timerLines07.SetTimeDurationLabel(
		newStartTimeLabel,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines07.SetTimeDurationLabel()\n"+
			"Expected an error return because the error prefix\n"+
			"is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

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

	timerLines03 := TextLineSpecTimerLines{}

	timerLines03.SetTimeFormat(
		newTimeFormat)

	timerLines04 := TextLineSpecTimerLines{}

	timerLines04.SetTimeFormat(
		"")

	defaultTimeFormat := "2006-01-02 15:04:05.000000000 -0700 MST"

	if timerLines04.timeFormat != defaultTimeFormat {

		t.Errorf("%v\n"+
			"Test Inspection #3\n"+
			"Error: Expected timerLines04.timeFormat\n"+
			"to be EQUAL TO the default time format.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"Expected Time Format: %v\n"+
			"  Actual Time Format: %v\n",
			ePrefix.String(),
			defaultTimeFormat,
			timerLines04.timeFormat)

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

func TestTextLineSpecTimerLines_TextBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_TextBuilder_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	sb0 := strings.Builder{}

	err = timerLines02.TextBuilder(
		&sb0,
		ePrefix.XCpy(
			"timerLines02"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines02.TextBuilder()\n"+
			"because 'timerLines02' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines03 *TextLineSpecTimerLines

	timerLines03,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines03<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sb0.Reset()

	err = timerLines03.TextBuilder(
		&sb0,
		ePrefix.XCpy(
			"timerLines03"))

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb0.String()),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines03.TextBuilder()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	var timerLines04 *TextLineSpecTimerLines

	timerLines04,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines04<-timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sb0.Reset()

	err = timerLines04.TextBuilder(
		&sb0,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines04.TextBuilder()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines05.TextBuilder(
		nil,
		ePrefix.XCpy(
			"timerLines05"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines05.TextBuilder()\n"+
			"because input parameter 'sBuilder' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_TextLineSpecName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_TextLineSpecName_000100()",
		"")

	expectedStr := "TimerLines"

	timerLines01 := TextLineSpecTimerLines{}

	actualStr := timerLines01.TextLineSpecName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.TextLineSpecName()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	_,
		timerLines02,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualStr = timerLines02.TextLineSpecName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.TextLineSpecName()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	return
}

func TestTextLineSpecTimerLines_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_TextTypeName_000100()",
		"")

	expectedStr := "TextLineSpecTimerLines"

	timerLines01 := TextLineSpecTimerLines{}

	actualStr := timerLines01.TextTypeName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.TextTypeName()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	_,
		timerLines02,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualStr = timerLines02.TextTypeName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.TextTypeName()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	return
}
