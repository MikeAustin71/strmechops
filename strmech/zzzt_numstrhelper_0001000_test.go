package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumStrHelper_DirtyToNativeNumStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrHelper_DirtyToNativeNumStr_000100()",
		"")

	dirtyStr := "($1,256,321.98)"

	expectedStr := "-1256321.98"

	testName := fmt.Sprintf("Test #1 DirtyString to NativeNumStr.\n"+
		"dirtyStr := \"%v\"",
		expectedStr)

	actualNumStr,
		err := new(NumStrHelper).DirtyToNativeNumStr(
		dirtyStr,
		".",
		ePrefix.XCpy(
			"actualStr<-dirtyStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestNumStrHelper_DirtyToNativeNumStr_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrHelper_DirtyToNativeNumStr_000200()",
		"")

	dirtyStr := "1,256,32 1.98-XX"

	expectedStr := "-1256321.98"

	testName := fmt.Sprintf("Test #1 DirtyString to NativeNumStr.\n"+
		"dirtyStr := \"%v\"",
		expectedStr)

	actualNumStr,
		err := new(NumStrHelper).DirtyToNativeNumStr(
		dirtyStr,
		".",
		ePrefix.XCpy(
			"actualStr<-dirtyStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestNumStrHelper_DirtyToNativeNumStr_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrHelper_DirtyToNativeNumStr_000300()",
		"")

	dirtyStr := "1,256,32 1zzz98"

	expectedStr := "1256321.98"

	testName := fmt.Sprintf("Test #1 DirtyString to NativeNumStr.\n"+
		"dirtyStr := \"%v\"",
		expectedStr)

	actualNumStr,
		err := new(NumStrHelper).DirtyToNativeNumStr(
		dirtyStr,
		"zzz",
		ePrefix.XCpy(
			"actualStr<-dirtyStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

type testNativeNumStrProfile struct {
	IntegerDigits    int
	FractionalDigits int
	IsNegativeValue  bool
}

// getTestNativeNumStr
//
// Generates a test Native Number Strings. The length
// of the integer segment and fractional digits segment
// are controlled by input parameters.
func getTestNativeNumStr(
	integerDigits int,
	fractionalDigits int,
	isNegativeValue bool) string {

	var nativeNumStr []rune

	if isNegativeValue {

		nativeNumStr = append(nativeNumStr, '-')

	}

	nextDigit := '1'

	if integerDigits < 1 {
		nativeNumStr = append(nativeNumStr, '0')
	} else {

		for i := 0; i < integerDigits; i++ {

			nativeNumStr = append(nativeNumStr, nextDigit)

			nextDigit += 1

			if nextDigit > '9' {
				nextDigit = '1'
			}

		}

	}

	if fractionalDigits < 1 {

		return string(nativeNumStr)
	}

	nativeNumStr = append(nativeNumStr, '.')

	baseFracDigits := fractionalDigits - 1

	nextDigit = '1'

	for i := 0; i < baseFracDigits; i++ {

		nativeNumStr = append(nativeNumStr, nextDigit)

		nextDigit += 1

		if nextDigit > '9' {
			nextDigit = '1'
		}

	}

	// Last Fractional digit is always 5
	nativeNumStr = append(nativeNumStr, '5')

	return string(nativeNumStr)

}
