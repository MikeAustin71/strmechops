package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecTimerLinesMolecule_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesMolecule_getFormattedText_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesMolecule := textLineSpecTimerLinesMolecule{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		timerLines01,
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		nil,
		ePrefix.XCtx(
			"timerLines01-Test#2"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		&timerLines02,
		ePrefix.XCtx(
			"timerLines02-Test#1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}
