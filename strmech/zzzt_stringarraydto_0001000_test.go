package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestStringArrayDto_AddManyStrings_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_AddManyStrings_000100()",
		"")

	expectedArray := []string{
		"Hello World",
		"How Now Brown Cow!",
		"The Cow jumped over the moon.",
		"A stitch in time, saves nine.",
	}

	lenExpectedArray := len(expectedArray)

	strArrayDto := StringArrayDto{}

	strArrayDto.AddManyStrings(
		expectedArray[0],
		expectedArray[1],
		expectedArray[2],
		expectedArray[3])

	actualArray :=
		strArrayDto.GetStringArray()

	lenActualArray := len(actualArray)

	if lenExpectedArray != lenActualArray {
		t.Errorf("%v\n"+
			"Error: lenExpectedArray != lenActualArray\n"+
			"Expected the actual array length = '%v'\n"+
			"Instead, the actual array length = '%v'\n",
			ePrefix.String(),
			lenExpectedArray,
			lenActualArray)
		return
	}

	for i := 0; i < lenExpectedArray; i++ {

		if expectedArray[i] != actualArray[i] {

			t.Errorf("%v\n"+
				"Error: expectedArray[%v] != actualArray[%v]\n"+
				"Expected Array Index[%v] = '%v'\n"+
				"Instead, Array Index[%v] = '%v'\n",
				ePrefix.String(),
				i,
				i,
				i,
				expectedArray[i],
				i,
				actualArray[i])

			return
		}

	}

	return
}
