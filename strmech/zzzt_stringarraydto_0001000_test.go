package strmech

import (
	"fmt"
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

func TestStringArrayDto_AddManyStrings_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_AddManyStrings_000200()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
		"7-Str",
		"8-Str",
		"9-Str",
		"10-Str",
	}

	lenBeginningArray := len(beginningArray)

	expectedArray := make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	for i := 0; i < 1000; i++ {

		strArrayDto := StringArrayDto{}

		strArrayDto.AddManyStrings(
			beginningArray[0],
			beginningArray[1],
			beginningArray[2],
			beginningArray[3],
			beginningArray[4],
			beginningArray[5],
			beginningArray[6],
			beginningArray[7],
			beginningArray[8],
			beginningArray[9])

		actualArray := strArrayDto.GetStringArray()

		if len(actualArray) != lenBeginningArray {
			t.Errorf("\n%v\n"+
				"Error: len(actualArray) != lenBeginningArray\n"+
				"Actual Array Length    = '%v'\n"+
				"Beginning Array Length = '%v'\n",
				ePrefix.String(),
				len(actualArray),
				lenBeginningArray)

			return
		}

		for j := 0; j < lenBeginningArray; j++ {
			if actualArray[j] != expectedArray[j] {
				t.Errorf("\n%v\n"+
					"Error: Actual Array is NOT EQUAL to \n"+
					"Expected Array!\n"+
					"  actualArray[%v] = '%v'\n"+
					"expectedArray[%v] = '%v'\n",
					ePrefix.String(),
					j,
					actualArray[j],
					j,
					expectedArray[j])

				return
			}
		}
	}

	return
}

func TestStringArrayDto_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_Equal_000100()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	strArrayDto1 := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	var strArrayDto2 StringArrayDto
	var err error

	strArrayDto2,
		err = strArrayDto1.CopyOut(
		ePrefix.XCpy(
			"strArrayDto2<-strArrayDto1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if !strArrayDto2.Equal(&strArrayDto1) {
		t.Errorf("%v\n"+
			"Error: strArrayDto1 != strArrayDto2\n"+
			"Expected strArrayDto1 to be equal to strArrayDto2\n"+
			"because strArrayDto2 is a copy of strArrayDto1.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	strArrayDto1.StrArray[2] = "3-Str#"

	if strArrayDto2.Equal(&strArrayDto1) {
		t.Errorf("%v\n"+
			"Error: strArrayDto1 == strArrayDto2\n"+
			"Expected strArrayDto1 to be NOT EQUAL to strArrayDto2\n"+
			"because string array element strArrayDto2[2] was changed.\n"+
			"HOWEVER, THEY ARE REPORTED AS EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStringArrayDto_Equal_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_Equal_000200()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	var err error

	strArrayDto1 := StringArrayDto{}

	strArrayDto1.SetStringArray(
		beginningArray)

	var strArrayDto2 StringArrayDto

	strArrayDto2,
		err = strArrayDto1.CopyOut(
		ePrefix.XCpy(
			"strArrayDto2<-strArrayDto1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if !strArrayDto2.Equal(&strArrayDto1) {
		t.Errorf("%v\n"+
			"Error: strArrayDto1 != strArrayDto2\n"+
			"Expected strArrayDto1 to be equal to strArrayDto2\n"+
			"because strArrayDto2 is a copy of strArrayDto1.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	err = strArrayDto1.DeleteAtIndex(
		2,
		ePrefix.XCpy(
			"strArrayDto1[2]"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	if strArrayDto2.Equal(&strArrayDto1) {
		t.Errorf("%v\n"+
			"Error: strArrayDto1 == strArrayDto2\n"+
			"Expected strArrayDto1 to be NOT EQUAL to strArrayDto2\n"+
			"because string array element strArrayDto2[2] was deleted.\n"+
			"HOWEVER, THEY ARE REPORTED AS EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStringArrayDto_InsertAtIndex_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_InsertAtIndex_000100()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	insertStr := "Inserted String"

	expectedArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"Inserted String",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	expectedStrArrayDto := StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	strArrayDto := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	err := strArrayDto.InsertAtIndex(
		insertStr,
		3,
		ePrefix.XCpy(
			"strArrayDto<-insertStr"))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto.GetStringArray())

		return
	}

}

func TestStringArrayDto_InsertAtIndex_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_InsertAtIndex_000200()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	insertStr := "Inserted String"

	expectedArray := []string{
		"Inserted String",
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	expectedStrArrayDto := StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	strArrayDto := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	err := strArrayDto.InsertAtIndex(
		insertStr,
		0,
		ePrefix.XCpy(
			"strArrayDto[0]<-insertStr"))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto.GetStringArray())

		return
	}

	strArrayDto = StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	err = strArrayDto.InsertAtIndex(
		insertStr,
		-1,
		ePrefix.XCpy(
			"strArrayDto[-1]<-insertStr"))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto.GetStringArray())

		return
	}

}

func TestStringArrayDto_InsertAtIndex_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_InsertAtIndex_000300()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	insertStr := "Inserted String"

	expectedArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"Inserted String",
		"6-Str",
	}

	expectedStrArrayDto := StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	strArrayDto := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	lastIdx := len(beginningArray) - 1

	err := strArrayDto.InsertAtIndex(
		insertStr,
		lastIdx,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]<-insertStr",
				lastIdx)))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto.GetStringArray())

		return
	}

	return
}

func TestStringArrayDto_InsertAtIndex_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_InsertAtIndex_000400()",
		"")

	beginningArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
	}

	insertStr := "Inserted String"

	expectedArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"4-Str",
		"5-Str",
		"6-Str",
		"Inserted String",
	}

	expectedStrArrayDto := StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	strArrayDto := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	lastIdx := len(beginningArray)

	err := strArrayDto.InsertAtIndex(
		insertStr,
		lastIdx,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]<-insertStr",
				lastIdx)))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto.GetStringArray())

		return
	}

	return
}
