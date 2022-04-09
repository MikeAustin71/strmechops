package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecTimerLinesNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesNanobot_copyIn_000100()",
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

	var timerLines02 *TextLineSpecTimerLines

	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCtx(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesNanobot := textLineSpecTimerLinesNanobot{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			nil,
			ePrefix.XCtx(
				"timerLines02<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtTimerLinesNanobot.copyIn(
			nil,
			timerLines01,
			ePrefix.XCtx(
				"nil<-timerLines01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'targetTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			&timerLines03,
			ePrefix.XCtx(
				"timerLines02<-timerLines03 (invalid)"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

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
