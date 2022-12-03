package strmech

import "sync"

// textLineSpecTitleMarqueeElectron
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeElectron struct {
	lock *sync.Mutex
}

func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) empty(
	txtLineTitleMarquee *TextLineSpecTitleMarquee) {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	if txtLineTitleMarquee == nil {
		return
	}

	txtLineTitleMarquee.standardTitleLeftMargin = ""

	txtLineTitleMarquee.standardTitleRightMargin = ""

	txtLineTitleMarquee.standardMaxLineLen = 0

	txtLineTitleMarquee.standardTextFieldLen = 0

	txtLineTitleMarquee.leadingBlankLines =
		TextLineSpecBlankLines{}

	txtLineTitleMarquee.leadingSolidLines =
		TextLineSpecSolidLine{}

	txtLineTitleMarquee.topTitleBlankLines =
		TextLineSpecBlankLines{}

	txtLineTitleMarquee.titleLines = nil

	txtLineTitleMarquee.bottomTitleBlankLines =
		TextLineSpecBlankLines{}

	txtLineTitleMarquee.trailingSolidLines =
		TextLineSpecSolidLine{}

	txtLineTitleMarquee.trailingBlankLines =
		TextLineSpecBlankLines{}

	return
}

func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) equal(
	txtLineTitleOne *TextLineSpecTitleMarquee,
	txtLineTitleTwo *TextLineSpecTitleMarquee) bool {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	if txtLineTitleOne.standardTitleLeftMargin !=
		txtLineTitleTwo.standardTitleLeftMargin {

		return false
	}

	if txtLineTitleOne.standardTitleRightMargin !=
		txtLineTitleTwo.standardTitleRightMargin {

		return false
	}

	if txtLineTitleOne.standardMaxLineLen !=
		txtLineTitleTwo.standardMaxLineLen {

		return false
	}

	if txtLineTitleOne.standardTextFieldLen !=
		txtLineTitleTwo.standardTextFieldLen {

		return false
	}

	if !txtLineTitleOne.leadingBlankLines.Equal(
		&txtLineTitleTwo.leadingBlankLines) {

		return false
	}

	if !txtLineTitleOne.leadingSolidLines.Equal(
		&txtLineTitleTwo.leadingSolidLines) {

		return false
	}

	if !txtLineTitleOne.topTitleBlankLines.Equal(
		&txtLineTitleTwo.topTitleBlankLines) {

		return false
	}

	lenTitleLines := len(txtLineTitleOne.titleLines)

	if lenTitleLines != len(txtLineTitleTwo.titleLines) {

		return false
	}

	for i := 0; i < lenTitleLines; i++ {

		if !txtLineTitleOne.titleLines[i].Equal(
			&txtLineTitleTwo.titleLines[i]) {

			return false
		}

	}

	if !txtLineTitleOne.bottomTitleBlankLines.Equal(
		&txtLineTitleTwo.bottomTitleBlankLines) {

		return false
	}

	if !txtLineTitleOne.trailingSolidLines.Equal(
		&txtLineTitleTwo.trailingSolidLines) {

		return false
	}

	if !txtLineTitleOne.trailingBlankLines.Equal(
		&txtLineTitleTwo.trailingBlankLines) {

		return false
	}

	return true
}
