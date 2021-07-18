package strmech

import (
	"strings"
	"testing"
)

func TestTextFieldSpecLabel_CopyIn_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()
	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtLabel,
		err := TextFieldSpecLabel{}.NewConstructor(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel := txtLabel.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}
