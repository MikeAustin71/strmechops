package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecLinesCollectionAtom_emptyCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollectionAtom_emptyCollection_000100()",
		"")

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	txtLinesColAtom.emptyCollection(
		nil)

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol01.textLines[1] = nil

	txtLinesColAtom.emptyCollection(
		&txtLinesCol01)

}

func TestTextLineSpecLinesCollectionAtom_equalCollections_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollectionAtom_equalCollections_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLinesCol02 TextLineSpecLinesCollection
	_,
		txtLinesCol02,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	areEqual := txtLinesColAtom.equalCollections(
		&txtLinesCol01,
		&txtLinesCol02)

	if !areEqual {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.equalCollections()\n"+
			"Expected txtLinesCol01 == txtLinesCol02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	areEqual = txtLinesColAtom.equalCollections(
		&txtLinesCol01,
		nil)

	if areEqual {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.equalCollections()\n"+
			"Expected areEqual == 'false' because"+
			"'textLinesCol02' is nil.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.String())

		return

	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	txtLinesCol04 := TextLineSpecLinesCollection{}

	areEqual = txtLinesColAtom.equalCollections(
		&txtLinesCol03,
		&txtLinesCol04)

	if !areEqual {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.equalCollections()\n"+
			"Expected txtLinesCol03 == txtLinesCol04\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	var txtLinesCol05, txtLinesCol06 TextLineSpecLinesCollection

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		txtLinesCol06,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol06.textLines[3].Empty()

	areEqual = txtLinesColAtom.equalCollections(
		&txtLinesCol05,
		&txtLinesCol06)

	if areEqual {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.equalCollections()\n"+
			"Expected areEqual == 'false' because"+
			"txtLinesCol06.textLines[3].Empty() was called.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.String())

		return

	}

}

func TestTextLineSpecLinesCollectionAtom_insertTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollectionAtom_insertTextLine_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	_,
		err = txtLinesColAtom.insertTextLine(
		&txtLinesCol01,
		&stdLine01,
		3,
		ePrefix.XCpy(
			"txtLinesCol01[3] = stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesColAtom.insertTextLine(
		nil,
		&stdLine01,
		3,
		ePrefix.XCpy(
			"nil = stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"insertTextLine()\n"+
			"because input parameter 'txtLinesCol' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01.Empty()

	_,
		err = txtLinesColAtom.insertTextLine(
		&txtLinesCol01,
		&stdLine01,
		3,
		ePrefix.XCpy(
			"nil = stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"insertTextLine()\n"+
			"because input parameter 'stdLine01' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLinesCol02 TextLineSpecLinesCollection
	_,
		txtLinesCol02,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol02.textLines[3].Empty()

	_,
		err = txtLinesColAtom.insertTextLine(
		&txtLinesCol02,
		&stdLine02,
		3,
		ePrefix.XCpy(
			"txtLinesCol02.textLines[3].Empty()"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"insertTextLine()\n"+
			"because input parameter txtLinesCol02.textLines[3].Empty() is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollectionAtom_peekPopTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollectionAtom_peekPopTextLine_000100()",
		"")

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	_,
		err := txtLinesColAtom.peekPopTextLine(
		nil,
		3,
		false,
		ePrefix.XCpy(
			"textLinesCol==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"peekPopTextLine()\n"+
			"because input parameter 'txtLinesCol' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecLinesCollectionAtomS_testValidityOfTextLinesCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollectionAtomS_testValidityOfTextLinesCollection_000100()",
		"")

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	_,
		err := txtLinesColAtom.testValidityOfTextLinesCollection(
		nil,
		ePrefix.XCpy(
			"textLineCol==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"testValidityOfTextLinesCollection()\n"+
			"because input parameter 'txtLinesCol' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol01 TextLineSpecLinesCollection

	_,
		txtLinesCol01,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol01.textLines[3] = nil

	_,
		err = txtLinesColAtom.testValidityOfTextLinesCollection(
		&txtLinesCol01,
		ePrefix.XCpy(
			"txtLinesCol01.textLines[3]==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesColAtom."+
			"testValidityOfTextLinesCollection()\n"+
			"because input parameter txtLinesCol01.textLines[3] = nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}
