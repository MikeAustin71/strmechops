package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestTextLineSpecSolidLine_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, lenExpectedStr+1)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtSolidLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSolidLine01.Read(p)\n"+
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtSolidLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtSolidLine01.Read(p)\n"+
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

func TestTextLineSpecSolidLine_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000200()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"
	var turnAutoLineTerminationOff bool
	turnAutoLineTerminationOff = true

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLineAllParms(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnAutoLineTerminationOff,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 1)

	var actualStr string

	var n, readBytesCnt int

	for {

		n,
			err = txtSolidLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSolidLine01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSolidLine01.Read(p)\n"+
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
			"Byte Length Error: txtSolidLine01.Read(p)\n"+
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"txtSolidLine01.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.String())
		return
	}

	p = make([]byte, 100)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = txtSolidLine01.Read(p)

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
			"Error Returned From txtSolidLine01.Read(p)\n"+
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

func TestTextLineSpecSolidLine_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000300()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

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
	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	formattedTxtStr,
		err =
		txtSolidLineMolecule.getFormattedText(
			&txtSolidLine01,
			ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	txtSolidLine01.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		txtSolidLine01.textLineReader,
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
			txtSolidLine01.textLineReader,
			p,
			ePrefix.XCpy("txtSolidLine01 is valid"))

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

func TestTextLineSpecSolidLine_Read_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000400()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
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
	sb.Grow(128)

	for {

		n,
			err = txtSolidLine01.Read(p)

		if n == 0 {
			break
		}

		sb.Write(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSolidLine01.Read(p)\n"+
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtSolidLine01.textLineReader != 'nil'\n",
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

func TestTextLineSpecSolidLine_Read_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000500()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.solidLineChars = nil
	txtSolidLine01.newLineChars = nil
	txtSolidLine01.leftMarginChars = nil
	txtSolidLine01.rightMarginChars = nil
	txtSolidLine01.solidLineCharsRepeatCount = -1

	p := make([]byte, 15)

	_,
		err = txtSolidLine01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01.Read(p)\n"+
			"because 'txtSolidLine01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	_,
		err = txtSolidLine02.Read(p)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02.Read(p)\n"+
			"because 'txtSolidLine02' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_ReaderInitialize_000100()",
		"")

	txtSolidLine00 := TextLineSpecSolidLine{}

	txtSolidLine00.ReaderInitialize()

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	lenExpectedStr := len(expectedSolidLineStr)

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, 5)

	var n int

	n,
		err = txtSolidLine01.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtSolidLine01.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: txtSolidLine01.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.String(),
			n)

		return
	}

	p = make([]byte, 100)

	txtSolidLine01.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtSolidLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSolidLine01.Read(p)\n"+
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtSolidLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtSolidLine01.Read(p)\n"+
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
			[]rune(expectedSolidLineStr),
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but txtSolidLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 100)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = txtSolidLine01.Read(p)

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
			"Error Returned From txtSolidLine01.Read(p)\n"+
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

	if txtSolidLine01.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but txtSolidLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetFullSolidLineConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetFullSolidLineConfig_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"SetFullSolidLineConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetFullSolidLineRunesConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_Read_000300()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err =
		txtSolidLine02.SetFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"SetFullSolidLineRunesConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetLeftMargin_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetLeftMargin_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	expectedNewLeftMargin := 5

	leftMarginStr := strings.Repeat(" ",
		expectedNewLeftMargin)

	err = txtSolidLine01.SetLeftMargin(
		leftMarginStr,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newLeftMargin :=
		txtSolidLine01.GetLeftMarginLength()

	if expectedNewLeftMargin != newLeftMargin {
		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.SetLeftMargin() set left margin = 5\n"+
			"Expected Left Margin = '%v'\n"+
			"Instead, Left Margin = '%v'\n",
			ePrefix.String(),
			expectedNewLeftMargin,
			newLeftMargin)

		return
	}

	expectedSolidLineStr =
		strings.Repeat(" ", expectedNewLeftMargin) +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText =
		txtSolidLine01.String()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetLeftMargin(
		"   ",
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"SetLeftMargin()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	err = txtSolidLine03.SetLeftMargin(
		strings.Repeat(" ", 1000001),
		ePrefix.XCpy(
			"txtSolidLine03:leftMargin==1000001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"SetLeftMargin()\n"+
			"because leftMargin==1000001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetNewLineChars_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	newLineChars = "\n\n\n\n"

	expectedSolidLineStr =
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			newLineChars

	err = txtSolidLine01.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetNewLineChars(
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02.\n"+
			"SetNewLineChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	newLineChars = "\n\n\n\n"

	expectedSolidLineStr =
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			newLineChars

	err = txtSolidLine01.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine03 TextLineSpecSolidLine

	txtSolidLine03,
		err = TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine03.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newLineChars =
		strings.Repeat("\n", 1000001)

	err = txtSolidLine03.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"SetNewLineChars()\n"+
			"because 'newLineChars' has a string length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	newLineChars = "\n\n"

	var txtSolidLine04 TextLineSpecSolidLine

	txtSolidLine04,
		err = TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine04.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newLineChars = ""

	err = txtSolidLine04.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v - ERROR\n"+
			"txtSolidLine04.SetNewLineChars()\n"+
			"Did not expect an error because 'newLineChars'"+
			"is an empty string and that should not cause an error.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetNewLineRunes_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	newLineChars = "\n\n\n\n"

	err = txtSolidLine01.SetNewLineRunes(
		[]rune(newLineChars),
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedSolidLineStr =
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01-New newLineRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - New newLineRunes\n"+
			"newLineChars = \"\\n\\n\\n\\n\"\n"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	expectedNewLineChars := ""

	err = txtSolidLine01.SetNewLineRunes(
		[]rune(expectedNewLineChars),
		ePrefix.XCpy(
			"txtSolidLine01 expectedNewLineChars = \"\""))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedNewLineChars = "\n"

	actualNewLineChars := txtSolidLine01.GetNewLineChars()

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedNewLineChars),
			true)

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualNewLineChars),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v - ERROR\n"+
			"expectedNewLineChars != actualNewLineChars\n"+
			"Expected actualNewLineChars to equal '\\n'\n"+
			"because submitted newLineChars was an empty rune array.\n"+
			"HOWEVER, actualNewLineChars to equal '\\n'\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	sMechPreon := strMechPreon{}

	var expectedNewLineRunes []rune

	expectedNewLineRunes,
		err =
		sMechPreon.getRepeatRuneChar(
			1000001,
			'\n',
			ePrefix.XCpy(
				"1,000,001 new-lines"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.SetNewLineRunes(
		expectedNewLineRunes,
		ePrefix.XCpy(
			"txtSolidLine01 one-million new-lines"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetNewLineRunes(expectedNewLineRunes)\n"+
			"because 'expectedNewLineRunes' is invalid."+
			"expectedNewLineRunes has an array length of 1,000,001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	expectedNewLineRunes = []rune{'\n', 0, '\n', 0, '\n'}

	err = txtSolidLine01.SetNewLineRunes(
		expectedNewLineRunes,
		ePrefix.XCpy(
			"txtSolidLine01 one-million new-lines"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetNewLineRunes(expectedNewLineRunes)\n"+
			"because 'expectedNewLineRunes' is invalid.\n"+
			"expectedNewLineRunes contains invalid rune characters\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	expectedNewLineRunes = []rune{'\n'}

	err = txtSolidLine01.SetNewLineRunes(
		expectedNewLineRunes,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetNewLineRunes(expectedNewLineRunes)\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetNewLineRunes(
		expectedNewLineRunes,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecSolidLine_SetSolidLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetNewLineRunes_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"
	var turnAutoLineTerminationOff bool
	turnAutoLineTerminationOff = true

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLineAllParms(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnAutoLineTerminationOff,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

		return
	}

	var txtSolidLine01FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err =
		txtSolidLine01.SetSolidLineAllParms(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnAutoLineTerminationOff,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01.\n"+
			"SetSolidLineAllParms()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetSolidLineCharsRepeatCount_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetSolidLineCharsRepeatCount_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	solidLineCharsRepeatCount = 12

	expectedSolidLineStr :=
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			newLineChars

	err = txtSolidLine01.SetSolidLineCharsRepeatCount(
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err = txtSolidLine01.SetSolidLineCharsRepeatCount(
		0,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineCharsRepeatCount()\n"+
			"because 'solidLineCharsRepeatCount' is less than one (1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.SetSolidLineCharsRepeatCount(
		1000001,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineCharsRepeatCount()\n"+
			"because 'solidLineCharsRepeatCount' is greater than one-million.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.SetSolidLineCharsRepeatCount(
		5,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineCharsRepeatCount()\n"+
			"because 'errorPrefix' is greater than one-million.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetSolidLineCharsRepeatCount(
		5,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecSolidLine_SetRightMargin_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetRightMargin_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	expectedNewRightMargin := 5

	err = txtSolidLine01.SetRightMargin(
		strings.Repeat(" ", expectedNewRightMargin),
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newRightMargin :=
		txtSolidLine01.GetRightMarginLength()

	if expectedNewRightMargin != newRightMargin {
		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.SetRightMargin() set right margin = 5\n"+
			"Expected Right Margin = '%v'\n"+
			"Instead, Right Margin = '%v'\n",
			ePrefix.String(),
			expectedNewRightMargin,
			newRightMargin)

		return
	}

	expectedSolidLineStr =
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			strings.Repeat(" ", expectedNewRightMargin) +
			newLineChars

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText =
		txtSolidLine01.String()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err = txtSolidLine02.SetRightMargin(
		strings.Repeat("", 3),
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"SetRightMargin()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	err = txtSolidLine03.SetRightMargin(
		strings.Repeat(" ", 1000001),
		ePrefix.XCpy(
			"txtSolidLine03:rightMargin==1000001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"SetRightMargin()\n"+
			"because rightMargin==1000001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err =
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_SetSolidLineStrChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetSolidLineStrChars_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	solidLineChars = "*"
	solidLineCharsRepeatCount = 20

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars = \"*\""))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	solidLineChars = ""

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars is empty!"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineChars' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineChars = strings.Repeat("=", 1000001)

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars has a length of 1,000,001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineChars' has a string length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineChars = "*"
	solidLineCharsRepeatCount = 0

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineCharsRepeatCount = zero!"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineCharsRepeatCount' has a value of zero (0).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = -1

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineCharsRepeatCount = -1 !"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineCharsRepeatCount' has a value of minus one (-1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = 1000001

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars has a length of 1,000,001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineCharsRepeatCount' has a value of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = 21

	err = txtSolidLine01.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	err = txtSolidLine02.SetSolidLineStrChars(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecSolidLine_SetSolidLineRuneChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_SetSolidLineRuneChars_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	solidLineChars = "*"
	solidLineCharsRepeatCount = 20

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	err = txtSolidLine01.SetSolidLineRuneChars(
		[]rune(solidLineChars),
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars = \"*\""))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	solidLineChars = ""

	err = txtSolidLine01.SetSolidLineRuneChars(
		[]rune(solidLineChars),
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineChars is empty!"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'solidLineChars' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	sMechPreon := strMechPreon{}

	var solidLineRunes []rune

	solidLineRunes,
		err = sMechPreon.getRepeatRuneChar(
		1000001,
		'=',
		ePrefix.XCpy(
			"Repeat '=' 1000001 times!"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineRunes length = 1,000,001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'solidLineRunes' has a string length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineRunes = []rune{'*'}
	solidLineCharsRepeatCount = 0

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineCharsRepeatCount = Zero"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'solidLineCharsRepeatCount' has a value of zero (0).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = -1

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineCharsRepeatCount = -1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'SetSolidLineRuneChars' has a value of minus one (-1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = 1000001

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"solidLineCharsRepeatCount has a value of 1,000,001"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'solidLineCharsRepeatCount' has a value of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = 21

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineRuneChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineRunes = []rune{'-', 0, '*'}

	err = txtSolidLine01.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"SetSolidLineStrChars()\n"+
			"because 'solidLineRunes' is invalid.\n"+
			"'solidLineRunes' contains invalid rune characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineRunes = []rune{'*'}
	solidLineCharsRepeatCount = 40

	var txtSolidLine02 TextLineSpecSolidLine

	err = txtSolidLine02.SetSolidLineRuneChars(
		solidLineRunes,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecSolidLine_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_String_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText string

	txtSolidLine01FmtText = txtSolidLine01.String()

	if strings.Contains(txtSolidLine01FmtText, "Error") {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.String() returned an error\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			txtSolidLine01FmtText)

		return

	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine01.solidLineChars = []rune{'-', 0, '*', 0}

	txtSolidLine01FmtText = txtSolidLine01.String()

	if !strings.Contains(strings.ToLower(txtSolidLine01FmtText), "error") {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01.String()\n"+
			"because txtSolidLine01.solidLineChars contains invalid rune characters."+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return

	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLine01FmtText = txtSolidLine02.String()

	if !strings.Contains(strings.ToLower(txtSolidLine01FmtText), "error") {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02.String()\n"+
			"because txtSolidLine02 is empty."+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecSolidLine_TextBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TextBuilder_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-!"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sb := strings.Builder{}

	sb,
		err = txtSolidLine01.TextBuilder(
		ePrefix.XCpy("txtSolidLine01->sb"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

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

		return
	}

	return
}

func TestTextLineSpecSolidLine_TextBuilder_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TextBuilder_000200()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	sb := strings.Builder{}

	sb,
		err := txtSolidLine01.TextBuilder(
		ePrefix.XCpy("empty txtSolidLine01->sb"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from txtSolidLine01.TextBuilder()\n"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"

	expectedTextStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine02 := TextLineSpecSolidLine{}

	err =
		txtSolidLine02.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sb,
		err = txtSolidLine02.TextBuilder(
		ePrefix.XCpy("valid txtSolidLine02->sb"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

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

		return
	}

	sb,
		err = txtSolidLine02.TextBuilder(
		StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"TextBuilder()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TextTypeName_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	actualTxtTypeName :=
		txtSolidLine01.TextTypeName()

	expectedTxtTypeName := "TextLineSpecSolidLine"

	if expectedTxtTypeName !=
		actualTxtTypeName {
		t.Errorf("%v\n"+
			"Error: txtSolidLine01.TextTypeName()\n"+
			"Expected Text Type Name = '%v'\n"+
			"Instead, Text Type Name = '%v'\n",
			ePrefix.String(),
			expectedTxtTypeName,
			actualTxtTypeName)
	}

	return
}

func TestTextLineSpecSolidLine_TextLineSpecName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TextLineSpecName_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	actualTxtLineSpecName :=
		txtSolidLine01.TextLineSpecName()

	expectedTxtLineSpecName := "SolidLine"

	if expectedTxtLineSpecName !=
		actualTxtLineSpecName {
		t.Errorf("%v\n"+
			"Error: txtSolidLine01.TextLineSpecName()\n"+
			"Expected Text Line Specification Name = '%v'\n"+
			"Instead, Text Line Specification Name = '%v'\n",
			ePrefix.String(),
			expectedTxtLineSpecName,
			actualTxtLineSpecName)
	}

	return
}

func TestTextLineSpecSolidLine_TurnAutoLineTerminationOff_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TurnAutoLineTerminationOff_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"
	var turnAutoLineTerminationOff bool

	turnAutoLineTerminationOff = false

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine01.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	txtSolidLine01.TurnAutoLineTerminationOff()

	if txtSolidLine01.GetTurnLineTerminatorOff() == false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.TurnAutoLineTerminationOff()\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'false'\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLine02.TurnAutoLineTerminationOff()

	if txtSolidLine02.GetTurnLineTerminatorOff() == false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.TurnAutoLineTerminationOff()\n"+
			"txtSolidLine02.GetTurnLineTerminatorOff() == 'false'\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_TurnAutoLineTerminationOn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_TurnAutoLineTerminationOn_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"
	var turnAutoLineTerminationOff bool

	turnAutoLineTerminationOff = true

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine01.GetTurnLineTerminatorOff() == false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'false'\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	txtSolidLine01.TurnAutoLineTerminationOn()

	if txtSolidLine01.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.TurnAutoLineTerminationOn()\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLine02.turnLineTerminatorOff = true

	txtSolidLine02.TurnAutoLineTerminationOn()

	if txtSolidLine02.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.TurnAutoLineTerminationOn()\n"+
			"txtSolidLine02.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	return
}
