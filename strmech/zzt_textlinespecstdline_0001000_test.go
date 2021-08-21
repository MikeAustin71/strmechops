package strmech

import (
	"testing"
)

func TestTextLineSpecStandardLine_CopyIn_000100(t *testing.T) {

	funcName := "TestTextLineSpecStandardLine_CopyIn_000100()"

	/*

		ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
			funcName,
			"")
	*/

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	numOfStdLines := stdLine.GetNumOfStdLines()

	if numOfStdLines != 1 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '1'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			funcName,
			numOfStdLines)

		return
	}

}
