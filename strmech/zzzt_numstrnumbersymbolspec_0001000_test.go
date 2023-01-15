package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumStrNumberSymbolSpec_NewCurrencyLeadingSymbols_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrNumberSymbolSpec_NewCurrencyLeadingSymbols_000100()",
		"")

	originalLeadingCurrencySymbol := "$ "

	var testName string

	originalCurrencyFieldSymbolPosition :=
		NumFieldSymPos.InsideNumField()

	originalCurrencyNumSignRelPos :=
		CurrNumSignRelPos.OutsideNumSign()

	newNStrNumberSymbolSpec,
		err := new(NumStrNumberSymbolSpec).NewCurrencyLeadingSymbol(
		originalLeadingCurrencySymbol,
		originalCurrencyFieldSymbolPosition,
		originalCurrencyNumSignRelPos,
		ePrefix.XCpy(
			"newNStrNumberSymbolSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLeadingNumSymStr :=
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolStr()

	testName = "Test # 1 - Actual vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != actualLeadingNumSymStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymStr NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"         actualLeadingNumSymStr = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumSymStr,
			originalLeadingCurrencySymbol)

		return
	}

	actualLeadingNumSymRunes :=
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolRunes()

	testName = "Test # 2 - Actual Runes vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != string(actualLeadingNumSymRunes) {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymRunes NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"       actualLeadingNumSymRunes = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			string(actualLeadingNumSymRunes),
			originalLeadingCurrencySymbol)

		return
	}

	var actualLeadingNumSymRunesDto RuneArrayDto

	actualLeadingNumSymRunesDto,
		err =
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolRunesDto(
			ePrefix.XCpy("actualLeadingNumSymRunesDto<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test # 3 - Actual RuneArrayDto vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != actualLeadingNumSymRunesDto.GetCharacterString() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymRunes NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"       actualLeadingNumSymRunesDto = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumSymRunesDto.GetCharacterString(),
			originalLeadingCurrencySymbol)

		return
	}

	var actualLeadingNumFieldSymbolPosition NumberFieldSymbolPosition

	actualLeadingNumFieldSymbolPosition =
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolPosition()

	testName = "Test # 4 - Actual NumberFieldSymbolPosition vs Expected Number Field Symbol Position"

	if originalCurrencyFieldSymbolPosition != actualLeadingNumFieldSymbolPosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumFieldSymbolPosition NOT EQUAL TO actualLeadingNumFieldSymbolPosition\n"+
			"  actualLeadingNumFieldSymbolPosition = '%v'\n"+
			"actualLeadingNumFieldSymbolPosition   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumFieldSymbolPosition.String(),
			originalCurrencyFieldSymbolPosition.String())

		return
	}

	var actualLeadingCurrencyNumSignRelativePosition CurrencyNumSignRelativePosition

	actualLeadingCurrencyNumSignRelativePosition =
		newNStrNumberSymbolSpec.GetCurrencyNumSignRelativePosition()

	testName = "Test # 5 - Actual CurrencyNumSignRelativePosition vs Expected Currency Number Sign Relative Position"

	if originalCurrencyNumSignRelPos != actualLeadingCurrencyNumSignRelativePosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingCurrencyNumSignRelativePosition NOT EQUAL TO originalCurrencyNumSignRelPos\n"+
			"       actualLeadingCurrencyNumSignRelativePosition = '%v'\n"+
			"originalCurrencyNumSignRelPos   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingCurrencyNumSignRelativePosition.String(),
			originalCurrencyNumSignRelPos.String())

		return
	}

	actualTrailingNumSymStr :=
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolStr()

	testName = "Test # 6 - Actual Trailing Number Symbol String Length Should Be Zero"

	if len(actualTrailingNumSymStr) != 0 {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymStr string length is greater than zero!\n"+
			"         actualTrailingNumSymStr = '%v'\n"+
			ePrefix.String(),
			testName,
			actualTrailingNumSymStr)

		return
	}

	return
}

func TestNumStrNumberSymbolSpec_NewCurrencyLeadingTrailingSymbols_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrNumberSymbolSpec_NewCurrencyLeadingTrailingSymbols_000100()",
		"")

	originalLeadingCurrencySymbol := "$ "
	originalTrailingCurrencySymbol := " $"

	var testName string

	originalCurrencyFieldSymbolPosition :=
		NumFieldSymPos.InsideNumField()

	originalCurrencyNumSignRelPos :=
		CurrNumSignRelPos.OutsideNumSign()

	newNStrNumberSymbolSpec,
		err := new(NumStrNumberSymbolSpec).NewCurrencyLeadingTrailingSymbols(
		originalLeadingCurrencySymbol,
		originalTrailingCurrencySymbol,
		originalCurrencyFieldSymbolPosition,
		originalCurrencyNumSignRelPos,
		ePrefix.XCpy(
			"newNStrNumberSymbolSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLeadingNumSymStr :=
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolStr()

	testName = "Test # 1 - Actual vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != actualLeadingNumSymStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymStr NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"         actualLeadingNumSymStr = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumSymStr,
			originalLeadingCurrencySymbol)

		return
	}

	actualLeadingNumSymRunes :=
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolRunes()

	testName = "Test # 2 - Actual Runes vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != string(actualLeadingNumSymRunes) {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymRunes NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"       actualLeadingNumSymRunes = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			string(actualLeadingNumSymRunes),
			originalLeadingCurrencySymbol)

		return
	}

	var actualLeadingNumSymRunesDto RuneArrayDto

	actualLeadingNumSymRunesDto,
		err =
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolRunesDto(
			ePrefix.XCpy("actualLeadingNumSymRunesDto<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test # 3 - Actual RuneArrayDto vs Expected Leading Number Symbol"

	if originalLeadingCurrencySymbol != actualLeadingNumSymRunesDto.GetCharacterString() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymRunes NOT EQUAL TO originalLeadingCurrencySymbol\n"+
			"       actualLeadingNumSymRunesDto = '%v'\n"+
			"originalLeadingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumSymRunesDto.GetCharacterString(),
			originalLeadingCurrencySymbol)

		return
	}

	var actualLeadingNumFieldSymbolPosition NumberFieldSymbolPosition

	actualLeadingNumFieldSymbolPosition =
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolPosition()

	testName = "Test # 4 - Actual NumberFieldSymbolPosition vs Expected Number Field Symbol Position"

	if originalCurrencyFieldSymbolPosition != actualLeadingNumFieldSymbolPosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumFieldSymbolPosition NOT EQUAL TO actualLeadingNumFieldSymbolPosition\n"+
			"  actualLeadingNumFieldSymbolPosition = '%v'\n"+
			"actualLeadingNumFieldSymbolPosition   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingNumFieldSymbolPosition.String(),
			originalCurrencyFieldSymbolPosition.String())

		return
	}

	var actualLeadingCurrencyNumSignRelativePosition CurrencyNumSignRelativePosition

	actualLeadingCurrencyNumSignRelativePosition =
		newNStrNumberSymbolSpec.GetCurrencyNumSignRelativePosition()

	testName = "Test # 5 - Actual CurrencyNumSignRelativePosition vs Expected Currency Number Sign Relative Position"

	if originalCurrencyNumSignRelPos != actualLeadingCurrencyNumSignRelativePosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingCurrencyNumSignRelativePosition NOT EQUAL TO originalCurrencyNumSignRelPos\n"+
			"       actualLeadingCurrencyNumSignRelativePosition = '%v'\n"+
			"originalCurrencyNumSignRelPos   = '%v'\n",
			ePrefix.String(),
			testName,
			actualLeadingCurrencyNumSignRelativePosition.String(),
			originalCurrencyNumSignRelPos.String())

		return
	}

	actualTrailingNumSymStr :=
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolStr()

	testName = "Test # 10 - Actual vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != actualTrailingNumSymStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymStr NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"         actualTrailingNumSymStr = '%v'\n"+
			"originalTrailingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumSymStr,
			originalTrailingCurrencySymbol)

		return
	}

	actualTrailingNumSymRunes :=
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolRunes()

	testName = "Test # 11 - Actual Runes vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != string(actualTrailingNumSymRunes) {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymRunes NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"     actualTrailingNumSymRunes = '%v'\n"+
			"originalTrailingCurrencySymbol = '%v'\n",
			ePrefix.String(),
			testName,
			string(actualTrailingNumSymRunes),
			originalTrailingCurrencySymbol)

		return
	}

	var actualTrailingNumSymRunesDto RuneArrayDto

	actualTrailingNumSymRunesDto,
		err =
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolRunesDto(
			ePrefix.XCpy("actualTrailingNumSymRunesDto<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test # 12 - Actual RuneArrayDto vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != actualTrailingNumSymRunesDto.GetCharacterString() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymRunes NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"       actualTrailingNumSymRunesDto = '%v'\n"+
			"originalTrailingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumSymRunesDto.GetCharacterString(),
			originalTrailingCurrencySymbol)

		return
	}

	var actualTrailingNumFieldSymbolPosition NumberFieldSymbolPosition

	actualTrailingNumFieldSymbolPosition =
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolPosition()

	testName = "Test # 13 - Actual NumberFieldSymbolPosition vs Expected Number Field Symbol Position"

	if originalCurrencyFieldSymbolPosition != actualTrailingNumFieldSymbolPosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumFieldSymbolPosition NOT EQUAL TO actualTrailingNumFieldSymbolPosition\n"+
			"  actualTrailingNumFieldSymbolPosition = '%v'\n"+
			"actualTrailingNumFieldSymbolPosition   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumFieldSymbolPosition.String(),
			originalCurrencyFieldSymbolPosition.String())

		return
	}

	var actualTrailingCurrencyNumSignRelativePosition CurrencyNumSignRelativePosition

	actualTrailingCurrencyNumSignRelativePosition =
		newNStrNumberSymbolSpec.GetCurrencyNumSignRelativePosition()

	testName = "Test # 14 - Actual CurrencyNumSignRelativePosition vs Expected Currency Number Sign Relative Position"

	if originalCurrencyNumSignRelPos != actualTrailingCurrencyNumSignRelativePosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingCurrencyNumSignRelativePosition NOT EQUAL TO originalCurrencyNumSignRelPos\n"+
			"       actualTrailingCurrencyNumSignRelativePosition = '%v'\n"+
			"originalCurrencyNumSignRelPos   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingCurrencyNumSignRelativePosition.String(),
			originalCurrencyNumSignRelPos.String())

		return
	}

	return
}

func TestNumStrNumberSymbolSpec_NewCurrencyTrailingSymbols_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrNumberSymbolSpec_NewCurrencyTrailingSymbols_000100()",
		"")

	originalTrailingCurrencySymbol := "$ "

	var testName string

	originalCurrencyFieldSymbolPosition :=
		NumFieldSymPos.InsideNumField()

	originalCurrencyNumSignRelPos :=
		CurrNumSignRelPos.OutsideNumSign()

	newNStrNumberSymbolSpec,
		err := new(NumStrNumberSymbolSpec).NewCurrencyTrailingSymbol(
		originalTrailingCurrencySymbol,
		originalCurrencyFieldSymbolPosition,
		originalCurrencyNumSignRelPos,
		ePrefix.XCpy(
			"newNStrNumberSymbolSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualTrailingNumSymStr :=
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolStr()

	testName = "Test # 1 - Actual vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != actualTrailingNumSymStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymStr NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"         actualTrailingNumSymStr = '%v'\n"+
			"originalTrailingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumSymStr,
			originalTrailingCurrencySymbol)

		return
	}

	actualTrailingNumSymRunes :=
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolRunes()

	testName = "Test # 2 - Actual Runes vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != string(actualTrailingNumSymRunes) {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymRunes NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"       actualTrailingNumSymRunes = '%v'\n"+
			"originalTrailingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			string(actualTrailingNumSymRunes),
			originalTrailingCurrencySymbol)

		return
	}

	var actualTrailingNumSymRunesDto RuneArrayDto

	actualTrailingNumSymRunesDto,
		err =
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolRunesDto(
			ePrefix.XCpy("actualTrailingNumSymRunesDto<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test # 3 - Actual RuneArrayDto vs Expected Trailing Number Symbol"

	if originalTrailingCurrencySymbol != actualTrailingNumSymRunesDto.GetCharacterString() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumSymRunes NOT EQUAL TO originalTrailingCurrencySymbol\n"+
			"       actualTrailingNumSymRunesDto = '%v'\n"+
			"originalTrailingCurrencySymbol   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumSymRunesDto.GetCharacterString(),
			originalTrailingCurrencySymbol)

		return
	}

	var actualTrailingNumFieldSymbolPosition NumberFieldSymbolPosition

	actualTrailingNumFieldSymbolPosition =
		newNStrNumberSymbolSpec.GetTrailingNumberSymbolPosition()

	testName = "Test # 4 - Actual NumberFieldSymbolPosition vs Expected Number Field Symbol Position"

	if originalCurrencyFieldSymbolPosition != actualTrailingNumFieldSymbolPosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingNumFieldSymbolPosition NOT EQUAL TO actualTrailingNumFieldSymbolPosition\n"+
			"  actualTrailingNumFieldSymbolPosition = '%v'\n"+
			"actualTrailingNumFieldSymbolPosition   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingNumFieldSymbolPosition.String(),
			originalCurrencyFieldSymbolPosition.String())

		return
	}

	var actualTrailingCurrencyNumSignRelativePosition CurrencyNumSignRelativePosition

	actualTrailingCurrencyNumSignRelativePosition =
		newNStrNumberSymbolSpec.GetCurrencyNumSignRelativePosition()

	testName = "Test # 5 - Actual CurrencyNumSignRelativePosition vs Expected Currency Number Sign Relative Position"

	if originalCurrencyNumSignRelPos != actualTrailingCurrencyNumSignRelativePosition {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualTrailingCurrencyNumSignRelativePosition NOT EQUAL TO originalCurrencyNumSignRelPos\n"+
			"       actualTrailingCurrencyNumSignRelativePosition = '%v'\n"+
			"originalCurrencyNumSignRelPos   = '%v'\n",
			ePrefix.String(),
			testName,
			actualTrailingCurrencyNumSignRelativePosition.String(),
			originalCurrencyNumSignRelPos.String())

		return
	}

	actualLeadingNumSymStr :=
		newNStrNumberSymbolSpec.GetLeadingNumberSymbolStr()

	testName = "Test # 6 - Actual Leading Number Symbol String Length Should Be Zero"

	if len(actualLeadingNumSymStr) != 0 {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualLeadingNumSymStr string length is greater than zero!\n"+
			"         actualLeadingNumSymStr = '%v'\n"+
			ePrefix.String(),
			testName,
			actualLeadingNumSymStr)

		return
	}

	return
}
