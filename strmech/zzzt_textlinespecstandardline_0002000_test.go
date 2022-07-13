package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestTextLineSpecStandardLine_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		stdLine01.GetFormattedText(
			ePrefix.XCpy(
				"stdLine01"))

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
			err = stdLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From stdLine01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: stdLine01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: stdLine01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if stdLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"stdLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: stdLine01.Read(p)\n"+
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

func TestTextLineSpecStandardLine_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		stdLine01.GetFormattedText(
			ePrefix.XCpy(
				"stdLine01"))

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
			err = stdLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From stdLine01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From stdLine01.Read(p)\n"+
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
			"Byte Length Error: stdLine01.Read(p)\n"+
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

	if stdLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"stdLine01.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.String())
		return
	}

	p = make([]byte, 200)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = stdLine01.Read(p)

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
			"Error Returned From stdLine01.Read(p)\n"+
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

func TestTextLineSpecStandardLine_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		stdLine01.GetFormattedText(
			ePrefix.XCpy(
				"stdLine01"))

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
	stdLineMolecule := textLineSpecStandardLineMolecule{}

	formattedTxtStr,
		_,
		_,
		err =
		stdLineMolecule.getFormattedText(
			&stdLine01,
			ePrefix.XCpy("stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	stdLine01.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		stdLine01.textLineReader,
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
			stdLine01.textLineReader,
			p,
			ePrefix.XCpy("stdLine01 is valid"))

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

func TestTextLineSpecStandardLine_Read_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000400()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		stdLine01.GetFormattedText(
			ePrefix.XCpy(
				"stdLine01"))

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
	sb.Grow(128)

	for {

		n,
			err = stdLine01.Read(p)

		if n == 0 {
			break
		}

		sb.Write(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From stdLine01.Read(p)\n"+
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

	if stdLine01.textLineReader != nil {
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

func TestTextLineSpecStandardLine_Read_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000500()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine01.textFields = nil
	stdLine01.numOfStdLines = -99
	stdLine01.newLineChars = []rune{'-', 0, 0, '-'}

	p := make([]byte, 15)

	_,
		err = stdLine01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01.Read(p)\n"+
			"because 'stdLine01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_Read_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Read_000600()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	p := make([]byte, 5)

	_,
		err :=
		stdLine01.Read(p)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01.Read(p)\n"+
			"because 'stdLine01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_ReaderInitialize_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		stdLine01.GetFormattedText(
			ePrefix.XCpy(
				"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 5)

	var n int

	n,
		err = stdLine01.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine01.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: stdLine01.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.String(),
			n)

		return
	}

	p = make([]byte, 200)

	stdLine01.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = stdLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From stdLine01.Read(p)\n"+
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

	if stdLine01.textLineReader != nil {
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

	if stdLine01.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but stdLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 200)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = stdLine01.Read(p)

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
			"Error Returned From stdLine01.Read(p)\n"+
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

	if stdLine01.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but stdLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	stdLine02 := TextLineSpecPlainText{}

	stdLine02.ReaderInitialize()

	return
}

func TestTextLineSpecStandardLine_ReaderInitialize_000200(t *testing.T) {

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.ReaderInitialize()

	return
}

func TestTextLineSpecStandardLine_ReplaceTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_ReplaceTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedLabelText := "Xray97 where are?"

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.ReplaceTextField(
		&labelTxt,
		1,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	// var labelTxt TextFieldSpecLabel
	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtTextField(
		1,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	labelField, ok := iTxtFieldSpec.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	err = labelField.IsValidInstanceError(
		ePrefix.XCpy(
			"labelField"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelText := labelField.GetTextLabel()

	if expectedLabelText != actualLabelText {

		t.Errorf("%v - Error\n"+
			"Expected expectedLabelText==actualLabelText\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedLabelText,
			actualLabelText)

		return
	}

	expectedNumTextFields := 3

	actualNumTextFields :=
		stdLine01.GetNumOfTextFields()

	if expectedNumTextFields != actualNumTextFields {

		t.Errorf("%v - Error\n"+
			"Expected stdLine01 to contain %v text fields.\n"+
			"Instead, stdLine01 contains %v text fields!\n",
			ePrefix.String(),
			expectedNumTextFields,
			actualNumTextFields)

		return
	}

	return
}

func TestTextLineSpecStandardLine_ReplaceTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_ReplaceTextField_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedLabelText := "Xray97 where are?"

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.ReplaceTextField(
		&labelTxt,
		304,
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.ReplaceTextField()\n"+
			"because input parameter 'replaceAtIndex' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

func TestTextLineSpecStandardLine_ReplaceTextField_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_ReplaceTextField_000300()",
		"")

	labelTxt,
		err := TextFieldSpecLabel{}.NewTextLabel(
		"Xray97 where are?",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.ReplaceTextField(
			&labelTxt,
			-5,
			ePrefix.XCpy(
				"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.ReplaceTextField()\n"+
			"because input parameter 'replaceAtIndex' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_ReplaceTextField_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_ReplaceTextField_000400()",
		"")

	labelTxt,
		err := TextFieldSpecLabel{}.NewTextLabel(
		"Xray97 where are?",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01 := TextLineSpecStandardLine{}

	err =
		stdLine01.ReplaceTextField(
			&labelTxt,
			-5,
			ePrefix.XCpy(
				"stdLine01"))

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.textFields = nil

	err =
		stdLine02.ReplaceTextField(
			&labelTxt,
			2,
			ePrefix.XCpy(
				"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine02.ReplaceTextField()\n"+
			"because input parameter 'stdLine02.textFields == nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	labelTxt.textLabel = nil
	labelTxt.fieldLen = -99

	err =
		stdLine03.ReplaceTextField(
			&labelTxt,
			2,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine03.ReplaceTextField()\n"+
			"because input parameter 'labelTxt' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		"Xray97 where are?",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		stdLine04.ReplaceTextField(
			&labelTxt,
			2,
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.ReplaceTextField()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_SetNumOfStdLines_000100(t *testing.T) {

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.SetNumOfStdLines(-5)

}

func TestTextLineSpecStandardLine_SetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_SetNewLineChars_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	newLineChars := "xyz\n"

	_ = stdLine01.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"stdLine01"))

	var err error
	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars = ""

	err = stdLine02.SetNewLineChars(
		newLineChars,
		ePrefix.XCpy(
			"stdLine02"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine02.SetNewLineChars()\n"+
			"because input parameter 'newLineChars' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars = "\n\n"

	err = stdLine03.SetNewLineChars(
		newLineChars,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine03.SetNewLineChars()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_SetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_SetNewLineRunes_000100()",
		"")

	newLineRunes := []rune{'x', 'y', 'z', '\n'}

	stdLine01 := TextLineSpecStandardLine{}

	_ = stdLine01.SetNewLineRunes(
		newLineRunes,
		ePrefix.XCpy(
			"stdLine01"))

	var err error
	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineRunes = nil

	err = stdLine02.SetNewLineRunes(
		newLineRunes,
		ePrefix.XCpy(
			"stdLine02"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine02.SetNewLineRunes()\n"+
			"because input parameter 'newLineRunes' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineRunes = []rune{'x', 0, 'y', 0, 'z', '\n'}

	err = stdLine03.SetNewLineRunes(
		newLineRunes,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine03.SetNewLineRunes()\n"+
			"because input parameter 'newLineRunes' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	newLineRunes = []rune{'\n'}

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.SetNewLineRunes(
		newLineRunes,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.SetNewLineRunes()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_SetStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_SetStandardLine_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	err = stdLine02.SetStandardLine(
		1,
		textFields,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.SetStandardLine(
		-9,
		textFields,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine03.SetStandardLine()\n"+
			"because input parameter 'numOfStdLines' is '-9' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine04.newLineChars = []rune{'x', 0, 'y', 0, 'z', 0, '\n'}

	err = stdLine04.SetStandardLine(
		1,
		textFields,
		ePrefix.XCpy(
			"stdLine04"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.SetStandardLine()\n"+
			"because input parameter 'stdLine04.newLineChars' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine05 TextLineSpecStandardLine

	stdLine05,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine05.SetStandardLine(
		1,
		textFields,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine05.SetNewLineRunes()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_SetStandardLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_SetStandardLineAllParms_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineRunes := []rune{'\n'}

	stdLine02 := TextLineSpecStandardLine{}

	err = stdLine02.SetStandardLineAllParms(
		1,
		textFields,
		newLineRunes,
		false,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.SetStandardLineAllParms(
		-1,
		textFields,
		newLineRunes,
		false,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine02.SetStandardLineAllParms()\n"+
			"because input parameter 'numOfStdLines' is '-2' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.SetStandardLineAllParms(
		1,
		nil,
		newLineRunes,
		false,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.SetStandardLineAllParms()\n"+
			"because input parameter 'textFields' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine05 TextLineSpecStandardLine

	stdLine05,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine05.SetStandardLineAllParms(
		1,
		textFields,
		[]rune{0, 'x', 0, 'y', 0, 'z', 0, '\n'},
		false,
		ePrefix.XCpy(
			"stdLine05"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine05.SetStandardLineAllParms()\n"+
			"because input parameter 'newLineChars' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine06 TextLineSpecStandardLine

	stdLine06,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine06.SetStandardLineAllParms(
		1,
		textFields,
		nil,
		false,
		ePrefix.XCpy(
			"stdLine06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine07 TextLineSpecStandardLine

	stdLine07,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine07.SetStandardLineAllParms(
		1,
		textFields,
		newLineRunes,
		false,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine07.SetNewLineRunes()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_SetTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_SetTextFields_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	err = stdLine02.SetTextFields(
		textFields,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	err = stdLine03.SetTextFields(
		nil,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine03.SetTextFields()\n"+
			"because input parameter 'textFields' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	textFields[1] = nil

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.SetTextFields(
		textFields,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.SetTextFields()\n"+
			"because input parameter 'textFields' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine05 TextLineSpecStandardLine

	stdLine05,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields,
		err = stdLine05.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine06 TextLineSpecStandardLine

	stdLine06,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine06.SetTextFields(
		textFields,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.SetTextFields()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_String_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_ = stdLine01.String()

	expectedString := "        How Now Brown Cow!       " +
		"\n"

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedString),
			true)

	actualString := stdLine02.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualString),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecStandardLine_TextBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TextBuilder_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		_ = stdLine01.TextBuilder(
		ePrefix.XCpy(
			"stdLine01"))

	expectedString := "        How Now Brown Cow!       " +
		"\n"

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedString),
			true)

	sb := strings.Builder{}

	sb,
		err = stdLine02.TextBuilder(
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine04 TextLineSpecStandardLine

	stdLine04,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sb,
		err = stdLine04.TextBuilder(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine04.TextBuilder()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TextTypeName_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	expectedStr := "TextLineSpecStandardLine"

	actualStr := stdLine01.TextTypeName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected string does NOT match \n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	return
}

func TestTextLineSpecStandardLine_TextLineSpecName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TextLineSpecName_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	expectedStr := "StandardLine"

	actualStr := stdLine01.TextLineSpecName()

	if expectedStr != actualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected string does NOT match \n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	return
}

func TestTextLineSpecStandardLine_TurnAutoLineTerminationOff_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TurnAutoLineTerminationOff_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedString := "        How Now Brown Cow!       "

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedString),
			true)

	stdLine01.TurnAutoLineTerminationOff()

	lineTerminatorStatus := stdLine01.GetTurnLineTerminatorOff()

	if lineTerminatorStatus == false {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected lineTerminatorStatus == 'true'\n"+
			"HOWEVER, 'lineTerminatorStatus' == 'false' \n",
			ePrefix.String())

		return
	}

	actualString := stdLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualString),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	stdLine02.TurnAutoLineTerminationOff()

	return
}

func TestTextLineSpecStandardLine_TurnAutoLineTerminationOn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TurnAutoLineTerminationOn_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedString := "        How Now Brown Cow!       " +
		"\n"

	stdLine01.turnLineTerminatorOff = true

	stdLine01.TurnAutoLineTerminationOn()

	lineTerminatorStatus := stdLine01.GetTurnLineTerminatorOff()

	if lineTerminatorStatus == true {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected lineTerminatorStatus == 'false'\n"+
			"HOWEVER, 'lineTerminatorStatus' == 'true' \n",
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedString),
			true)

	actualString := stdLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualString),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	stdLine02.TurnAutoLineTerminationOn()

	return
}
