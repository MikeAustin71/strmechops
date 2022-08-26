package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumStrMathQuarkExtendRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathQuarkExtendRunes_000100()",
		"")

	origRuneArray := []rune("123456")

	extensionRuneArray := []rune("0000")

	nStrMathQuark := numStrMathQuark{}

	origRuneArrayDto,
		err := RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<origRuneArrayDto-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var expectedRuneArrayDto RuneArrayDto

	expectedRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<expectedRuneArrayDto-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillChar := '0'

	newCharArrayLen := 2

	var extendRunesRight bool

	extendRunesRight = false

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			&origRuneArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"origRuneArrayDto<-Length=2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !expectedRuneArrayDto.Equal(&origRuneArrayDto) {

		t.Errorf("%v\n"+
			"Error: origRuneArrayDto\n"+
			"Expected expectedRuneArrayDto = origRuneArrayDto\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	newCharArrayLen = 10

	extendRunesRight = true

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			&origRuneArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"origRuneArrayDto2<-Length=10"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedRuneArray :=
		append(origRuneArray, extensionRuneArray...)

	expectedRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		expectedRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<expectedRuneArrayDto2-expectedRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !expectedRuneArrayDto.Equal(&origRuneArrayDto) {

		t.Errorf("%v\n"+
			"Error: origRuneArrayDto\n"+
			"Expected expectedRuneArrayDto2 = origRuneArrayDto2\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	origRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<origRuneArrayDto3-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newCharArrayLen = 10

	extendRunesRight = false

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			&origRuneArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"origRuneArrayDto3<-Length=10"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedRuneArray =
		append(extensionRuneArray, origRuneArray...)

	expectedRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		expectedRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<expectedRuneArrayDto3-expectedRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !expectedRuneArrayDto.Equal(&origRuneArrayDto) {

		t.Errorf("%v\n"+
			"Error: origRuneArrayDto\n"+
			"Expected expectedRuneArrayDto3 = origRuneArrayDto3\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"origRuneArrayDto3    = '%v'\n"+
			"expectedRuneArrayDto3= '%v'\n",
			ePrefix.String(),
			origRuneArrayDto.GetCharacterString(),
			expectedRuneArrayDto.GetCharacterString())

		return

	}

	err =
		nStrMathQuark.extendRunes(
			nil,
			&origRuneArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"origRuneArrayDto3<-Length=10"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from nStrMathQuark.extendRunes()\n"+
			"because input parameter 'incomingCharsToExtend' is a 'nil' pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			nil,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"origRuneArrayDto3<-Length=10"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from nStrMathQuark.extendRunes()\n"+
			"because input parameter 'extendedOutputChars' is a 'nil' pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	origRuneArray = []rune("123456")

	origRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<origRuneArrayDto4-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	extensionRuneArray = []rune("0")

	expectedRuneArray = append(
		extensionRuneArray,
		origRuneArray...)

	expectedRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		expectedRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<expectedRuneArrayDto4-expectedRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	newCharArrayLen = 7

	extendRunesRight = false

	outputArrayDto := RuneArrayDto{}

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			&outputArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"outputArrayDto<-Length=7"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !expectedRuneArrayDto.Equal(&outputArrayDto) {

		t.Errorf("%v\n"+
			"Error: origRuneArrayDto\n"+
			"Expected expectedRuneArrayDto4 = outputArrayDto\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"outputArrayDto       = '%v'\n"+
			"expectedRuneArrayDto4= '%v'\n",
			ePrefix.String(),
			outputArrayDto.GetCharacterString(),
			expectedRuneArrayDto.GetCharacterString())

		return

	}

	origRuneArray = []rune("123456")

	expectedRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<expectedRuneArrayDto5-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	origRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		origRuneArray,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<origRuneArrayDto5-origRuneArray"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	outputArrayDto = RuneArrayDto{}

	newCharArrayLen = -1

	extendRunesRight = false

	err =
		nStrMathQuark.extendRunes(
			&origRuneArrayDto,
			&outputArrayDto,
			fillChar,
			newCharArrayLen,
			extendRunesRight,
			ePrefix.XCpy(
				"outputArrayDto<-Length=-1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !expectedRuneArrayDto.Equal(&outputArrayDto) {

		t.Errorf("%v\n"+
			"Error: origRuneArrayDto\n"+
			"Expected expectedRuneArrayDto5 = outputArrayDto2\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"outputArrayDto2       = '%v'\n"+
			"expectedRuneArrayDto5 = '%v'\n",
			ePrefix.String(),
			outputArrayDto.GetCharacterString(),
			expectedRuneArrayDto.GetCharacterString())

		return

	}

	return
}
