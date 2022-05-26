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
