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
