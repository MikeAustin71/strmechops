package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecStandardLine_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine01 != stdLine02.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

}

func TestTextLineSpecStandardLine_Equal_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.newLineChars = []rune{'\n', '\n'}

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'newLineChars' are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.numOfStdLines = 2

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'numOfStdLines' are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000400()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.turnLineTerminatorOff = true

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'turnLineTerminatorOff' \n"+
			"values are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000500()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	err = stdLine02.DeleteAtIndex(
		2,
		ePrefix.XCtx(
			"stdLine02"))

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'Text Field Arrays' \n"+
			"are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_EqualITextLine_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainText(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.EqualITextLine(
		plainTextLine01)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'plainTextLine01' is the wrong type.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual = stdLine01.EqualITextLine(
		&stdLine02)

	if areEqual == false {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == true because input parameter\n"+
			"'stdLine02' is identical to 'stdLine01'.\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_EqualITextLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_TESTMETHOD_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	areEqual := stdLine01.EqualITextLine(
		nil)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'iTextLine' is nil.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}
