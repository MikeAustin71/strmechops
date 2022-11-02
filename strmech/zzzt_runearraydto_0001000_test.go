package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestRuneArrayDto_DeleteLeadingTrailingChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingCeiling_000100()",
		"")

	var runeArrayDto RuneArrayDto
	var err error

	originalStr := "How now brown cow."
	expectedStr := "now brown cow."
	numOfCharsToDelete := uint64(4)

	runeArrayDto = new(RuneArrayDto).NewStringDefault(
		originalStr)

	err = runeArrayDto.DeleteLeadingTrailingChars(
		numOfCharsToDelete,
		false,
		ePrefix.XCpy(
			"runeArrayDto"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr := runeArrayDto.GetCharacterString()

	if actualStr != expectedStr {

		t.Errorf("\n%v\n"+
			"Test # 1\n"+
			"Error: Returned 'actualStr' is NOT EQUAL\n"+
			"to 'expectedStr'!\n"+
			"expectedStr = '%v'\n"+
			"actualStr   = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	expectedStr = "How now brown"
	numOfCharsToDelete = 5

	var runeArrayDto2 RuneArrayDto

	runeArrayDto2 = new(RuneArrayDto).NewStringDefault(
		originalStr)

	err = runeArrayDto2.DeleteLeadingTrailingChars(
		numOfCharsToDelete,
		true,
		ePrefix.XCpy(
			"runeArrayDto"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = runeArrayDto2.GetCharacterString()

	if actualStr != expectedStr {

		t.Errorf("\n%v\n"+
			"Test # 2\n"+
			"Original String = '%v'\n"+
			"Error: Returned 'actualStr' is NOT EQUAL\n"+
			"to 'expectedStr'!\n"+
			"expectedStr = '%v'\n"+
			"actualStr   = '%v'\n",
			ePrefix.String(),
			originalStr,
			expectedStr,
			actualStr)

		return
	}

	return
}
