package strmech

import "sync"

// textLineSpecTitleMarqueeElectron
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeElectron struct {
	lock *sync.Mutex
}

//	empty
//
//	Receives a pointer to an instance of
//	TextLineSpecTitleMarquee and proceeds to reset the
//	data values for member variables to their initial
//	or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values contained in
//	input parameter 'txtLineTitleMarquee' will be
//	deleted and reset to their zero values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineTitleMarquee			*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. All the internal member
//		variables contained in this instance will be
//		deleted and reset to their zero values.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
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

	txtLineTitleMarquee.leadingBlankLines.Empty()

	txtLineTitleMarquee.leadingSolidLines.Empty()

	txtLineTitleMarquee.topTitleBlankLines.Empty()

	txtLineTitleMarquee.titleLines.Empty()

	txtLineTitleMarquee.bottomTitleBlankLines.Empty()

	txtLineTitleMarquee.trailingSolidLines.Empty()

	txtLineTitleMarquee.trailingBlankLines.Empty()

	return
}

// equal
//
// Receives a pointer to two instances of
// TextLineSpecTitleMarquee and proceeds to compare
// their member variables in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison
// is returned. If the member variables for both
// instances are equal in all respects, this flag is set
// to 'true'. Otherwise, this method returns 'false'.
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

	if !txtLineTitleOne.titleLines.Equal(
		&txtLineTitleTwo.titleLines) {

		return false
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
