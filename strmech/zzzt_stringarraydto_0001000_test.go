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

func TestStringArrayDto_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_CopyIn_000100()",
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

	var strArrayDto01, strArrayDto02 StringArrayDto

	strArrayDto01 = StringArrayDto{}.New()

	strArrayDto01.SetStringArray(
		beginningArray)

	desc1 := "strArrayDto01 Desc-1"
	desc2 := "strArrayDto01 Desc-2"

	strArrayDto01.SetDescription1(desc1)

	strArrayDto01.SetDescription1(desc2)

	err := strArrayDto02.CopyIn(
		&strArrayDto01,
		ePrefix.XCpy(
			"strArrayDto02<-strArrayDto01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())

		return
	}

	if !strArrayDto02.Equal(&strArrayDto01) {

		t.Errorf("\n%v\n"+
			"Error: strArrayDto01!=strArrayDto02\n"+
			"Expected that strArrayDto01 would equal strArrayDto02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	err = strArrayDto02.CopyIn(
		&strArrayDto01,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from strArrayDto02.CopyIn()\n"+
			"because Error Prefix is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStringArrayDto_GetRuneArrayCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_GetRuneArrayCollection_000100()",
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

	expectedRuneArrays := RuneArrayCollection{}

	lenBeginningArray := len(beginningArray)
	var err error

	for i := 0; i < lenBeginningArray; i++ {

		err = expectedRuneArrays.AddRuneArrayString(
			beginningArray[i],
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"expectedRuneArrays<-beginningArray[%v]",
					i)))

		if err != nil {
			t.Errorf("%v",
				err.Error())

			return
		}
	}

	var strArrayDto1 StringArrayDto

	strArrayDto1 = StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	actualRuneArrays := strArrayDto1.GetRuneArrayCollection()

	if !expectedRuneArrays.Equal(&actualRuneArrays) {

		t.Errorf("\n%v\n"+
			"Error: Expected expectedRuneArrays===actualRuneArrays\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	var expectedArrayDto StringArrayDto

	expectedArrayDto = expectedRuneArrays.GetStringArrayDto()

	if !expectedArrayDto.Equal(&strArrayDto1) {

		t.Errorf("\n%v\n"+
			"Error: Expected expectedArrayDto===strArrayDto1\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

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

	lenBeginningArray := len(beginningArray)

	expectedArray := make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] = beginningArray[i]
	}

	expectedArray = append(
		expectedArray,
		insertStr)

	zeroBasedIndex := lenBeginningArray + 5

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
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]<-insertStr",
				lenBeginningArray)))

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

	expectedArray = make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] = beginningArray[i]
	}

	expectedArray = append(
		[]string{insertStr},
		expectedArray...)

	zeroBasedIndex = -5

	expectedStrArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	strArrayDto2 := StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	err = strArrayDto2.InsertAtIndex(
		insertStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf("strArrayDto[%v]<-insertStr",
				zeroBasedIndex)))

	if err != nil {
		t.Errorf("%v",
			err.Error())
		return
	}

	if !strArrayDto2.Equal(
		&expectedStrArrayDto) {

		t.Errorf("%v\n"+
			"Error: strArrayDto2!=expectedStrArrayDto\n"+
			"expectedStrArrayDto = \n%v\n"+
			"Actual strArrayDto2 = \n%v\n",
			ePrefix.String(),
			expectedStrArrayDto.GetStringArray(),
			strArrayDto2.GetStringArray())

		return
	}

	err = strArrayDto2.InsertAtIndex(
		insertStr,
		zeroBasedIndex,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from strArrayDto05.ReplaceAtIndex()\n"+
			"because Error Prefix is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestStringArrayDto_InsertAtIndex_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStringArrayDto_InsertAtIndex_000500()",
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

func TestStringArrayDto_PopAtIndex_000100(t *testing.T) {

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
		"7-Str",
		"8-Str",
		"9-Str",
		"10-Str",
	}

	expectedArray := []string{
		"1-Str",
		"2-Str",
		"3-Str",
		"5-Str",
		"6-Str",
		"7-Str",
		"8-Str",
		"9-Str",
		"10-Str",
	}

	lenExpectedArray := len(expectedArray)

	var strArrayDto1 StringArrayDto

	desc1 := "Hello1"
	desc2 := "Goodbye2"

	strArrayDto1 = StringArrayDto{}.NewStringArray(
		beginningArray,
		"",
		"")

	strArrayDto1.SetDescription1(desc1)

	strArrayDto1.SetDescription2(desc2)

	actualDesc1 := strArrayDto1.GetDescription1()

	if actualDesc1 != desc1 {
		t.Errorf("\n%v\n"+
			"Error: Expected Description1 is NOT EQUAL\n"+
			"to Actual Description1\n"+
			"Expected Description1 = '%v'\n"+
			"  Actual Description1 = '%v'\n",
			ePrefix.String(),
			desc1,
			actualDesc1)

		return
	}

	actualDesc2 := strArrayDto1.GetDescription2()

	if actualDesc2 != desc2 {
		t.Errorf("\n%v\n"+
			"Error: Expected Description2 is NOT EQUAL\n"+
			"to Actual Description2\n"+
			"Expected Description2 = '%v'\n"+
			"  Actual Description2 = '%v'\n",
			ePrefix.String(),
			desc2,
			actualDesc2)

		return
	}

	var poppedStr4 string
	var err error

	zeroBasedIndex := 3

	poppedStr4,
		_,
		err = strArrayDto1.PopAtIndex(
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto1[%v]",
				zeroBasedIndex)))

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	if poppedStr4 != beginningArray[zeroBasedIndex] {
		t.Errorf("\n%v\n"+
			"Popped String is NOT EQUAL to target string!\n"+
			"Target String = '%v'\n"+
			"Popped String = '%v'\n",
			ePrefix.String(),
			beginningArray[zeroBasedIndex],
			poppedStr4)

		return
	}

	finalArrayLen := strArrayDto1.GetStringArrayLength()

	if finalArrayLen != lenExpectedArray {
		t.Errorf("\n%v\n"+
			"Error: After popping index %v, the string\n"+
			"array should have a length of '%v'.\n"+
			"Instead, the actual final array length is '%v'\n",
			ePrefix.String(),
			zeroBasedIndex,
			lenExpectedArray,
			finalArrayLen)

		return
	}

	return
}
func TestStringArrayDto_ReplaceAtIndex_000100(t *testing.T) {

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
		"7-Str",
		"8-Str",
		"9-Str",
		"10-Str",
	}

	replacementStr := "newReplaceStr"

	lenBeginningArray := len(beginningArray)

	expectedArray := make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	var expectedArrayDto, strArrayDto01, strArrayDto02,
		strArrayDto03, strArrayDto04, strArrayDto05 StringArrayDto

	strArrayDto01 = StringArrayDto{}.New()

	strArrayDto01.SetStringArray(
		beginningArray)

	zeroBasedIndex := 6

	expectedArray[zeroBasedIndex] = replacementStr

	err := strArrayDto01.ReplaceAtIndex(
		replacementStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto01[%v]=replacementStr",
				zeroBasedIndex)))

	if err != nil {

		t.Errorf("%v",
			err.Error())

		return
	}

	expectedArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	if !expectedArrayDto.Equal(&strArrayDto01) {
		t.Errorf("\n%v\n"+
			"Error: expectedArrayDto!=strArrayDto01\n"+
			"Replace strArrayDto01[%v] FAILED!\n",
			ePrefix.String(),
			zeroBasedIndex)

		return
	}

	expectedArray = make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	zeroBasedIndex = 0

	expectedArray[zeroBasedIndex] = replacementStr

	strArrayDto02 = StringArrayDto{}.New()

	strArrayDto02.SetStringArray(
		beginningArray)

	err = strArrayDto02.ReplaceAtIndex(
		replacementStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto02[%v]=replacementStr",
				zeroBasedIndex)))

	if err != nil {

		t.Errorf("%v",
			err.Error())

		return
	}

	expectedArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	if !expectedArrayDto.Equal(&strArrayDto02) {
		t.Errorf("\n%v\n"+
			"Error: expectedArrayDto!=strArrayDto02\n"+
			"Replace strArrayDto02[%v] FAILED!\n",
			ePrefix.String(),
			zeroBasedIndex)

		return
	}

	expectedArray = make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	zeroBasedIndex = lenBeginningArray - 1

	expectedArray[zeroBasedIndex] = replacementStr

	strArrayDto03 = StringArrayDto{}.New()

	strArrayDto03.SetStringArray(
		beginningArray)

	err = strArrayDto03.ReplaceAtIndex(
		replacementStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto03[%v]=replacementStr",
				zeroBasedIndex)))

	if err != nil {

		t.Errorf("%v",
			err.Error())

		return
	}

	expectedArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	if !expectedArrayDto.Equal(&strArrayDto03) {
		t.Errorf("\n%v\n"+
			"Error: expectedArrayDto!=strArrayDto03\n"+
			"Replace strArrayDto03[%v] FAILED!\n",
			ePrefix.String(),
			zeroBasedIndex)

		return
	}

	expectedArray = make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	zeroBasedIndex = 99

	expectedArray = append(expectedArray,
		replacementStr)

	strArrayDto04 = StringArrayDto{}.New()

	strArrayDto04.SetStringArray(
		beginningArray)

	err = strArrayDto04.ReplaceAtIndex(
		replacementStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto04[%v]=replacementStr",
				zeroBasedIndex)))

	if err != nil {

		t.Errorf("%v",
			err.Error())

		return
	}

	expectedArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	if !expectedArrayDto.Equal(&strArrayDto04) {
		t.Errorf("\n%v\n"+
			"Error: expectedArrayDto!=strArrayDto04\n"+
			"Replace strArrayDto04[%v] FAILED!\n",
			ePrefix.String(),
			zeroBasedIndex)

		return
	}

	if strArrayDto04.StrArray[lenBeginningArray] !=
		replacementStr {
		t.Errorf("\n%v\n"+
			"Error: Expected strArrayDto04.StrArray[%v]\n"+
			"would be equal to '%v'\n"+
			"Instead, strArrayDto04.StrArray[%v]= '%v'\n",
			ePrefix.String(),
			zeroBasedIndex,
			replacementStr,
			zeroBasedIndex,
			strArrayDto04.StrArray[lenBeginningArray])

		return

	}

	expectedArray = make([]string, lenBeginningArray)

	for i := 0; i < lenBeginningArray; i++ {
		expectedArray[i] =
			beginningArray[i]
	}

	zeroBasedIndex = -1

	expectedArray = append(
		[]string{replacementStr},
		expectedArray...)

	strArrayDto05 = StringArrayDto{}.New()

	strArrayDto05.SetStringArray(
		beginningArray)

	err = strArrayDto05.ReplaceAtIndex(
		replacementStr,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"strArrayDto05[%v]=replacementStr",
				zeroBasedIndex)))

	if err != nil {

		t.Errorf("%v",
			err.Error())

		return
	}

	expectedArrayDto = StringArrayDto{}.NewStringArray(
		expectedArray,
		"",
		"")

	if !expectedArrayDto.Equal(&strArrayDto05) {
		t.Errorf("\n%v\n"+
			"Error: expectedArrayDto!=strArrayDto05\n"+
			"Replace strArrayDto05[%v] FAILED!\n",
			ePrefix.String(),
			zeroBasedIndex)

		return
	}

	if strArrayDto05.StrArray[0] !=
		replacementStr {

		t.Errorf("\n%v\n"+
			"Error: Expected strArrayDto05.StrArray[%v]\n"+
			"would be equal to '%v'\n"+
			"Instead, strArrayDto05.StrArray[%v]= '%v'\n",
			ePrefix.String(),
			0,
			replacementStr,
			0,
			strArrayDto05.StrArray[0])

		return

	}

	err = strArrayDto05.ReplaceAtIndex(
		replacementStr,
		3,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from strArrayDto05.ReplaceAtIndex()\n"+
			"because Error Prefix is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}
