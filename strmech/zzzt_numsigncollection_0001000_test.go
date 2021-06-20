package strmech

import "testing"

func TestNumberSignSymbolCollection_AddSymbol_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_AddSymbol_000100()"

	var err error
	var nSignSymOne, nSignSymTwo,
		nSignSymThree, nSignSymFour NumberSignSymbolDto

	nSignSymOne,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo,
		err = NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymThree,
		err = NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymFour,
		err = NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignCollection := NumberSignSymbolCollection{}

	err =
		nSignCollection.AddSymbol(
			nSignSymOne,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymTwo,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymThree,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymFour,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='false'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='true'\n",
			ePrefix)
		return
	}

	colLen := nSignCollection.GetCollectionLength()

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '4'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	var collection []NumberSignSymbolDto

	collection,
		err = nSignCollection.GetCollection(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	colLen = len(collection)

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection array colLen == '4'\n"+
			"Instead, actual collection array colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	if !collection[0].Equal(&nSignSymOne) {
		t.Errorf("%v - Error\n"+
			"Expected collection[0].Equal(&nSignSymOne)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[1].Equal(&nSignSymTwo) {
		t.Errorf("%v - Error\n"+
			"Expected collection[1].Equal(&nSignSymTwo)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[2].Equal(&nSignSymThree) {
		t.Errorf("%v - Error\n"+
			"Expected collection[2].Equal(&nSignSymThree)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[3].Equal(&nSignSymFour) {
		t.Errorf("%v - Error\n"+
			"Expected collection[3].Equal(&nSignSymFour)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

}

func TestNumberSignSymbolCollection_AddSymbol_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_AddSymbol_000200()"

	var err error
	var nSignSymOne, nSignSymTwo,
		nSignSymThree, nSignSymFour,
		nSignSymThreeDup NumberSignSymbolDto

	nSignSymOne,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo,
		err = NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymThree,
		err = NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymThreeDup,
		err = NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymFour,
		err = NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignCollection := NumberSignSymbolCollection{}

	err =
		nSignCollection.AddSymbol(
			nSignSymOne,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymTwo,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymThree,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymFour,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddSymbol(
			nSignSymThreeDup,
			ePrefix)

	if err != nil {
		t.Errorf("Error- nSignSymThreeDup\n"+
			"%v", err.Error())
		return
	}

	if nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='false'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='true'\n",
			ePrefix)
		return
	}

	colLen := nSignCollection.GetCollectionLength()

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '4'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	var collection []NumberSignSymbolDto

	collection,
		err = nSignCollection.GetCollection(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	colLen = len(collection)

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection array colLen == '4'\n"+
			"Instead, actual collection array colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	if !collection[0].Equal(&nSignSymOne) {
		t.Errorf("%v - Error\n"+
			"Expected collection[0].Equal(&nSignSymOne)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[1].Equal(&nSignSymTwo) {
		t.Errorf("%v - Error\n"+
			"Expected collection[1].Equal(&nSignSymTwo)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[2].Equal(&nSignSymThree) {
		t.Errorf("%v - Error\n"+
			"Expected collection[2].Equal(&nSignSymThree)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[3].Equal(&nSignSymFour) {
		t.Errorf("%v - Error\n"+
			"Expected collection[3].Equal(&nSignSymFour)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}
}

func TestNumberSignSymbolCollection_AddNewSymbol_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_AddNewSymbol_000100()"

	var err error
	var nSignSymOne, nSignSymTwo,
		nSignSymThree, nSignSymFour NumberSignSymbolDto

	nSignSymOne,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo,
		err = NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymThree,
		err = NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymFour,
		err = NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignCollection := NumberSignSymbolCollection{}

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='false'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='true'\n",
			ePrefix)
		return
	}

	colLen := nSignCollection.GetCollectionLength()

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '4'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	var collection []NumberSignSymbolDto

	collection,
		err = nSignCollection.GetCollection(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	colLen = len(collection)

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection array colLen == '4'\n"+
			"Instead, actual collection array colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	if !collection[0].Equal(&nSignSymOne) {
		t.Errorf("%v - Error\n"+
			"Expected collection[0].Equal(&nSignSymOne)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[1].Equal(&nSignSymTwo) {
		t.Errorf("%v - Error\n"+
			"Expected collection[1].Equal(&nSignSymTwo)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[2].Equal(&nSignSymThree) {
		t.Errorf("%v - Error\n"+
			"Expected collection[2].Equal(&nSignSymThree)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[3].Equal(&nSignSymFour) {
		t.Errorf("%v - Error\n"+
			"Expected collection[3].Equal(&nSignSymFour)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

}

func TestNumberSignSymbolCollection_AddNewSymbol_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_AddNewSymbol_000200()"

	var err error
	var nSignSymOne, nSignSymTwo,
		nSignSymThree, nSignSymFour NumberSignSymbolDto

	nSignSymOne,
		err = NumberSignSymbolDto{}.New(
		"+",
		"",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymTwo,
		err = NumberSignSymbolDto{}.New(
		"-",
		"",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymThree,
		err = NumberSignSymbolDto{}.New(
		"",
		"+",
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignSymFour,
		err = NumberSignSymbolDto{}.New(
		"",
		"-",
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	nSignCollection := NumberSignSymbolCollection{}

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Duplicate of #3
	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='false'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='true'\n",
			ePrefix)
		return
	}

	colLen := nSignCollection.GetCollectionLength()

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '4'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	var collection []NumberSignSymbolDto

	collection,
		err = nSignCollection.GetCollection(
		ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	colLen = len(collection)

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection array colLen == '4'\n"+
			"Instead, actual collection array colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	if !collection[0].Equal(&nSignSymOne) {
		t.Errorf("%v - Error\n"+
			"Expected collection[0].Equal(&nSignSymOne)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[1].Equal(&nSignSymTwo) {
		t.Errorf("%v - Error\n"+
			"Expected collection[1].Equal(&nSignSymTwo)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[2].Equal(&nSignSymThree) {
		t.Errorf("%v - Error\n"+
			"Expected collection[2].Equal(&nSignSymThree)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

	if !collection[3].Equal(&nSignSymFour) {
		t.Errorf("%v - Error\n"+
			"Expected collection[3].Equal(&nSignSymFour)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix)
	}

}

func TestNumberSignSymbolCollection_EmptyCollection_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_AddNewSymbol_000100()"

	var err error

	nSignCollection := NumberSignSymbolCollection{}

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	if nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='false'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='true'\n",
			ePrefix)
		return
	}

	colLen := nSignCollection.GetCollectionLength()

	if colLen != 4 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '4'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

	nSignCollection.EmptyCollection()

	if !nSignCollection.IsCollectionEmpty() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.IsCollectionEmpty()=='true'\n"+
			"Instead, nSignCollection.IsCollectionEmpty()=='false'\n",
			ePrefix)
		return
	}

	colLen = nSignCollection.GetCollectionLength()

	if colLen != 0 {
		t.Errorf("%v - Error\n"+
			"Expected collection length colLen == '0'\n"+
			"Instead, actual collection length colLen == '%v'\n",
			ePrefix,
			colLen)
		return
	}

}

func TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000100"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                   012345678
	hostRunes := []rune(" +1234.56")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex != 1 {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == 1\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber != true {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'true'\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'false'\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000200"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                   012345678
	hostRunes := []rune(" +1234.56")
	startIndex := 2

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if foundLeadingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundLeadingNumSign == 'false'\n"+
			"Instead foundLeadingNumSign == 'true'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == 0\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber != false {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'false'\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'true'\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000300"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                   0123456789
	hostRunes := []rune(" -+1234.56")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Round # 1 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	//                     0123456789
	//hostRunes := []rune(" -+1234.56")
	startIndex = 2

	foundLeadingNumSign =
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Round #2 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex != 2 {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == 2\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber != true {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'true'\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'false'\n",
			ePrefix)

		return
	}

}

func TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000400"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                   0123456789
	hostRunes := []rune(" -+1234.56")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Round # 1 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	//                     0123456789
	//hostRunes := []rune(" -+1234.56")
	startIndex = 2

	foundLeadingNumSign =
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Round #2 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex != 2 {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == 2\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex)

		return
	}

	if nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber != true {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'true'\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'false'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[0].numSignPosition != NumSymPos.Before() {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].numSignPosition == NumSymPos.Before()\n"+
			"Instead, nSignCollection.numSignSymbols[0].numSignPosition == '%v'\n",
			ePrefix,
			nSignCollection.numSignSymbols[0].numSignPosition.String())

		return
	}

	foundNumberSign,
		nSignSymDto,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign != true {
		t.Errorf("%v - Error\n"+
			"Result: nSignCollection.GetFoundNumberSignSymbol()\n"+
			"Expected foundNumberSign == 'true'\n"+
			"Instead, foundNumberSign == 'false'\n",
			ePrefix)
		return
	}

	if collectionIndex != 0 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '0'\n"+
			"Instead, collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

	actualNumSignPos := nSignSymDto.GetNumSignSymPosition()

	if actualNumSignPos != NumSymPos.Before() {
		t.Errorf("%v - Error\n"+
			"Expected actualNumSignPos == NumSymPos.Before()\n"+
			"Instead, actualNumSignPos == '%v'\n",
			ePrefix,
			actualNumSignPos.String())
		return
	}

	if nSignSymDto.GetLeadingNumSignFoundInNumber() != true {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymDto.GetLeadingNumSignFoundInNumber() == 'true'\n"+
			"Instead, nSignSymDto.GetLeadingNumSignFoundInNumber() == 'false'\n",
			ePrefix)
		return
	}

	if nSignSymDto.GetLeadingNumSignFoundIndex() != 2 {
		t.Errorf("%v - Error\n"+
			"Expected nSignSymDto.GetLeadingNumSignFoundIndex() == '2'\n"+
			"Instead, nSignSymDto.GetLeadingNumSignFoundIndex() == '%v'\n",
			ePrefix,
			nSignSymDto.GetLeadingNumSignFoundIndex())
		return
	}

}

func TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsLeadingNumSignAtHostIndex_000500"

	nSignCollection := NumberSignSymbolCollection{}

	//                   0123456789
	hostRunes := []rune(" -+1234.56")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if foundLeadingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundLeadingNumSign == 'false'\n"+
			"Instead foundLeadingNumSign == 'true'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == true {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'false'\n"+
			"Instead foundNumberSign == 'true'\n",
			ePrefix)

		return
	}

	if numSignSymbol.IsEmpty() == false {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.IsEmpty() == 'true'\n"+
			"Instead numSignSymbol.IsEmpty() == 'false'\n",
			ePrefix)

		return
	}

	if collectionIndex != -1 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '-1'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000100(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000100"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                             1
	//                   01234567890
	hostRunes := []rune(" 1234.56+  ")
	startIndex := 8

	foundTrailingNumSign :=
		nSignCollection.IsTrailingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundTrailingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundTrailingNumSign == 'true'\n"+
			"Instead foundTrailingNumSign == 'false'\n",
			ePrefix)

		return
	}

	if nSignCollection.numSignSymbols[2].trailingNumSignFoundIndex != startIndex {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[2].trailingNumSignFoundIndex == '%v'\n"+
			"Instead, nSignCollection.numSignSymbols[2].trailingNumSignFoundIndex == '%v'\n",
			ePrefix,
			startIndex,
			nSignCollection.numSignSymbols[2].trailingNumSignFoundIndex)

		return
	}

	if nSignCollection.numSignSymbols[2].trailingNumSignFoundInNumber != true {
		t.Errorf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[2].trailingNumSignFoundInNumber == 'true'\n"+
			"Instead, nSignCollection.numSignSymbols[2].trailingNumSignFoundInNumber == 'false'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == false {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'true'\n"+
			"Instead foundNumberSign == 'false'\n",
			ePrefix)

		return
	}

	numSignCharStr := string(numSignSymbol.GetTrailingNumSignChars())

	if numSignCharStr != "+" {
		t.Errorf("%v - Error\n"+
			"Expected numSignCharStr == '+'\n"+
			"Instead numSignCharStr == '%v'\n",
			ePrefix,
			numSignCharStr)

		return
	}

	if collectionIndex != 2 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '2'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000200(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000200"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                             1
	//                   01234567890
	hostRunes := []rune(" 1234.56+  ")
	startIndex := 7

	foundTrailingNumSign :=
		nSignCollection.IsTrailingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if foundTrailingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundTrailingNumSign == 'false'\n"+
			"Instead foundTrailingNumSign == 'true'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == true {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'false'\n"+
			"Instead foundNumberSign == 'true'\n",
			ePrefix)

		return
	}

	if numSignSymbol.IsEmpty() == false {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.IsEmpty() == 'true'\n"+
			"Instead numSignSymbol.IsEmpty() == 'false'\n",
			ePrefix)

		return
	}

	if collectionIndex != -1 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '-1'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000300(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000300"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                             1
	//                   01234567890
	hostRunes := []rune(" 1234.56+  ")
	startIndex := 9

	foundTrailingNumSign :=
		nSignCollection.IsTrailingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if foundTrailingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundTrailingNumSign == 'false'\n"+
			"Instead foundTrailingNumSign == 'true'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == true {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'false'\n"+
			"Instead foundNumberSign == 'true'\n",
			ePrefix)

		return
	}

	if numSignSymbol.IsEmpty() == false {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.IsEmpty() == 'true'\n"+
			"Instead numSignSymbol.IsEmpty() == 'false'\n",
			ePrefix)

		return
	}

	if collectionIndex != -1 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '-1'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000400(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000400"

	nSignCollection := NumberSignSymbolCollection{}

	//                             1
	//                   01234567890
	hostRunes := []rune(" 1234.56+  ")
	startIndex := 8

	foundTrailingNumSign :=
		nSignCollection.IsTrailingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if foundTrailingNumSign {
		t.Errorf("%v - Error\n"+
			"Expected foundTrailingNumSign == 'false'\n"+
			"Instead foundTrailingNumSign == 'true'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == true {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'false'\n"+
			"Instead foundNumberSign == 'true'\n",
			ePrefix)

		return
	}

	if numSignSymbol.IsEmpty() == false {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.IsEmpty() == 'true'\n"+
			"Instead numSignSymbol.IsEmpty() == 'false'\n",
			ePrefix)

		return
	}

	if collectionIndex != -1 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '-1'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}

}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000500(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000500"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	// Index = 0
	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 1
	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 2
	err =
		nSignCollection.AddNewSymbol(
			"(",
			")",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 3
	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 5
	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                             11
	//                   012345678901
	hostRunes := []rune(" (1234.56)  ")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		t.Errorf("%v - Error\n"+
			"nSignCollection.IsLeadingNumSignAtHostIndex()\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		return
	}

	startIndex = 9

	foundTrailingNumSign :=
		nSignCollection.IsTrailingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundTrailingNumSign {
		t.Errorf("%v - Error\n"+
			"nSignCollection.IsTrailingNumSignAtHostIndex()\n"+
			"Expected foundTrailingNumSign == 'true'\n"+
			"Instead foundTrailingNumSign == 'false'\n",
			ePrefix)

		return
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == false {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'true'\n"+
			"Instead foundNumberSign == 'false'\n",
			ePrefix)

		return
	}

	numSignCharStr := string(numSignSymbol.GetLeadingNumSignChars())

	if numSignCharStr != "(" {
		t.Errorf("%v - Error\n"+
			"Expected Leading numSignCharStr == '('\n"+
			"Instead Leading numSignCharStr == '%v'\n",
			ePrefix,
			numSignCharStr)

		return
	}

	numSignCharStr = string(numSignSymbol.GetTrailingNumSignChars())

	if numSignCharStr != ")" {
		t.Errorf("%v - Error\n"+
			"Expected Trailing numSignCharStr == ')'\n"+
			"Instead Trailing numSignCharStr == '%v'\n",
			ePrefix,
			numSignCharStr)

		return
	}

	if collectionIndex != 2 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '2'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}
}

func TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000600(t *testing.T) {

	ePrefix := "TestNumberSignSymbolCollection_IsTrailingNumSignAtHostIndex_000600"

	nSignCollection := NumberSignSymbolCollection{}
	var err error

	// Index = 0
	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 1
	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 2
	err =
		nSignCollection.AddNewSymbol(
			"(",
			")",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 3
	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	// Index = 5
	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		t.Errorf("%v", err.Error())
		return
	}

	//                             11
	//                   012345678901
	hostRunes := []rune(" (1234.56 ) ")

	for i := 0; i < 2; i++ {
		_ =
			nSignCollection.IsLeadingNumSignAtHostIndex(
				hostRunes,
				i)
	}

	for j := 9; j < 12; j++ {

		_ =
			nSignCollection.IsTrailingNumSignAtHostIndex(
				hostRunes,
				j)
	}

	foundNumberSign,
		numSignSymbol,
		collectionIndex :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign == false {
		t.Errorf("%v - Error\n"+
			"Expected foundNumberSign == 'true'\n"+
			"Instead foundNumberSign == 'false'\n",
			ePrefix)

		return
	}

	numSignCharStr := string(numSignSymbol.GetLeadingNumSignChars())

	if numSignCharStr != "(" {
		t.Errorf("%v - Error\n"+
			"Expected Leading numSignCharStr == '('\n"+
			"Instead Leading numSignCharStr == '%v'\n",
			ePrefix,
			numSignCharStr)

		return
	}

	if numSignSymbol.GetLeadingNumSignFoundIndex() != 1 {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.GetLeadingNumSignFoundIndex() == '1'\n"+
			"Instead numSignSymbol.GetLeadingNumSignFoundIndex() == '%v'\n",
			ePrefix,
			numSignSymbol.GetLeadingNumSignFoundIndex())

		return
	}

	numSignCharStr = string(numSignSymbol.GetTrailingNumSignChars())

	if numSignCharStr != ")" {
		t.Errorf("%v - Error\n"+
			"Expected Trailing numSignCharStr == ')'\n"+
			"Instead Trailing numSignCharStr == '%v'\n",
			ePrefix,
			numSignCharStr)

		return
	}

	if numSignSymbol.GetTrailingNumSignFoundIndex() != 10 {
		t.Errorf("%v - Error\n"+
			"Expected numSignSymbol.GetTrailingNumSignFoundIndex() == '10'\n"+
			"Instead numSignSymbol.GetTrailingNumSignFoundIndex() == '%v'\n",
			ePrefix,
			numSignSymbol.GetTrailingNumSignFoundIndex())

		return
	}

	if collectionIndex != 2 {
		t.Errorf("%v - Error\n"+
			"Expected collectionIndex == '2'\n"+
			"Instead collectionIndex == '%v'\n",
			ePrefix,
			collectionIndex)

		return
	}
}
