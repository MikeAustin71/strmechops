package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestCharSearchTerminationType_XValueInt_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000100()",
		"")

	var isValid bool

	for i := 0; i < 6; i++ {

		charSearchTermType := CharSearchTerminationType(i)

		isValid = charSearchTermType.XIsValid()

		if i == 0 {
			if isValid {

				t.Errorf("%v\n"+
					"Error: CharSearchTerminationType(0).None()\n"+
					"evaluates as 'Valid'. This is actually an\n"+
					"invalid value!\n",
					ePrefix.String())

				return
			}

		} else if isValid == false {

			t.Errorf("%v\n"+
				"Error: Valid value classified as invalid!\n"+
				"CharSearchTerminationType = '%v'\n"+
				"This should be a valid value! It is NOT!\n",
				ePrefix.String(),
				charSearchTermType.String())

			return

		}

	}

}

func TestCharSearchTerminationType_XValueInt_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000200()",
		"")

	names := []string{
		"None",
		"ProcessError",
		"EndOfTargetString",
		"SearchLengthLimit",
		"TerminationDelimiters",
		"FoundSearchTarget",
	}

	var actualName string

	lenNames := len(names)

	for i := 0; i < lenNames; i++ {

		charSearchTermType := CharSearchTerminationType(i)

		actualName = charSearchTermType.String()

		if actualName != names[i] {
			t.Errorf("%v\n"+
				"Error: Invalid name returned from charSearchTermType.String()!\n"+
				"Expected charSearchTermType string  = '%v'\n"+
				"Instead, charSearchTermType string  = '%v'\n"+
				"Actual charSearchTermType int value = '%v'\n\n",
				ePrefix.String(),
				names[i],
				charSearchTermType.String(),
				charSearchTermType.XValueInt())

			return

		}

	}

}

func TestCharSearchTerminationType_XValueInt_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000300()",
		"")

	charSearchTermType := CharSearchTerminationType(-972)

	strName := charSearchTermType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharSearchTerminationType(-972)\n"+
			"would return name of 'None' from \n"+
			"charSearchTermType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}

func TestCharSearchTerminationType_XValueInt_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000400()",
		"")

	expectedIntValue := -972

	charSearchTermType := CharSearchTerminationType(expectedIntValue)

	actualIntValue := charSearchTermType.XValueInt()

	if expectedIntValue != actualIntValue {

		t.Errorf("%v\n"+
			"Error: Expected charSearchTermType integer value\n"+
			" NOT equal to actual integer value\n"+
			"Expected charSearchTermType integer value = '%v'\n"+
			"Actual charSearchTermType integer value   = '%v'\n",
			ePrefix.String(),
			expectedIntValue,
			actualIntValue)

		return

	}

	strName := charSearchTermType.XReturnNoneIfInvalid()

	if strName.String() != "None" {

		t.Errorf("%v\n"+
			"Error: Expected CharSearchTerminationType(-972)\n"+
			"would return name of 'None' from \n"+
			"charSearchTermType.XReturnNoneIfInvalid().\n"+
			"It DID NOT!\n"+
			"strName string value = '%v'\n"+
			"   strName int value = '%v'\n",
			ePrefix.String(),
			strName.String(),
			strName.XValueInt())

		return

	}

}

func TestCharSearchTerminationType_XValueInt_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestCharSearchTerminationType_XValueInt_000500()",
		"")

	ucNames := []string{
		"None",
		"ProcessError",
		"EndOfTargetString",
		"SearchLengthLimit",
		"TerminationDelimiters",
		"FoundSearchTarget",
	}

	lcNames := []string{
		"none",
		"processerror",
		"endoftargetstring",
		"searchlengthlimit",
		"terminationdelimiters",
		"foundsearchtarget",
	}

	lenUcNames := len(ucNames)

	var charSearchTermType CharSearchTerminationType
	var charSearchTermType2 CharSearchTerminationType
	var expectedCharSearchTermType CharSearchTerminationType
	var err error

	for i := 0; i < lenUcNames; i++ {

		charSearchTermType,
			err = CharSearchTerminationType(0).
			XParseString(ucNames[i], true)

		if err != nil {
			t.Errorf("%v\n"+
				"Error: CharSearchTerminationType(0).\n"+
				"XParseString(ucNames[%v], true)\n"+
				"'%v'\n",
				ePrefix.String(),
				i,
				err.Error())

			return

		}

		charSearchTermType2,
			err = CharSearchTerminationType(0).
			XParseString(lcNames[i], false)

		if err != nil {
			t.Errorf("%v\n"+
				"Error: CharSearchTerminationType(0).\n"+
				"XParseString(lcNames[%v], true)\n"+
				"'%v'\n",
				ePrefix.String(),
				i,
				err.Error())

			return

		}

		switch i {
		case 0:
			expectedCharSearchTermType = CharSearchTermType.None()
		case 1:
			expectedCharSearchTermType = CharSearchTermType.ProcessError()
		case 2:
			expectedCharSearchTermType = charSearchTermType.EndOfTargetString()
		case 3:
			expectedCharSearchTermType = charSearchTermType.SearchLengthLimit()
		case 4:
			expectedCharSearchTermType = charSearchTermType.TerminationDelimiters()
		case 5:
			expectedCharSearchTermType = charSearchTermType.FoundSearchTarget()

		}

		if expectedCharSearchTermType != charSearchTermType {
			t.Errorf("%v\n"+
				"Error: expectedCharSearchTermType != charSearchTermType\n"+
				"Exexpected CharSearchTerminationType string value = '%v'\n"+
				"Actual CharSearchTerminationType string value     = '%v'\n"+
				"Actual CharSearchTerminationType integer value    = '%v'\n",
				ePrefix.String(),
				expectedCharSearchTermType.String(),
				charSearchTermType.String(),
				charSearchTermType.XValueInt())

			return
		}

		if expectedCharSearchTermType != charSearchTermType2 {
			t.Errorf("%v\n"+
				"Error: expectedCharSearchTermType != charSearchTermType2\n"+
				"Exexpected CharSearchTerminationType string value = '%v'\n"+
				"Actual CharSearchTerminationType string value     = '%v'\n"+
				"Actual CharSearchTerminationType integer value    = '%v'\n",
				ePrefix.String(),
				expectedCharSearchTermType.String(),
				charSearchTermType2.String(),
				charSearchTermType2.XValueInt())

			return
		}

	}
}
