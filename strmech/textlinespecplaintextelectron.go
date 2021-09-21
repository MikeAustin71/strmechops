package strmech

import "sync"

type textLineSpecPlainTextElectron struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextLineSpecPlainText and proceeds to reset the data values for
// member values to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'plainTextLine' will be deleted and reset to their zero values.
//
func (txtLinePlainTextElectron *textLineSpecPlainTextElectron) empty(
	plainTextLine *TextLineSpecPlainText) {

	if txtLinePlainTextElectron.lock == nil {
		txtLinePlainTextElectron.lock = new(sync.Mutex)
	}

	txtLinePlainTextElectron.lock.Lock()

	defer txtLinePlainTextElectron.lock.Unlock()

	if plainTextLine == nil {
		return
	}

	plainTextLine.leftMarginChars = nil

	plainTextLine.rightMarginChars = nil

	plainTextLine.textString = ""

	plainTextLine.turnLineTerminatorOff = false

	plainTextLine.newLineChars = nil

	return
}

// equal - Receives pointers to two TextLineSpecPlainText
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'plainTxtLineOne'
// and 'plainTxtLineTwo' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise, this method
// returns 'false'.
//
func (txtLinePlainTextElectron *textLineSpecPlainTextElectron) equal(
	plainTxtLineOne *TextLineSpecPlainText,
	plainTxtLineTwo *TextLineSpecPlainText) bool {

	if txtLinePlainTextElectron.lock == nil {
		txtLinePlainTextElectron.lock = new(sync.Mutex)
	}

	txtLinePlainTextElectron.lock.Lock()

	defer txtLinePlainTextElectron.lock.Unlock()

	if plainTxtLineOne == nil ||
		plainTxtLineTwo == nil {

		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		plainTxtLineOne.leftMarginChars,
		plainTxtLineTwo.leftMarginChars) {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		plainTxtLineOne.rightMarginChars,
		plainTxtLineTwo.rightMarginChars) {
		return false
	}

	if plainTxtLineOne.textString !=
		plainTxtLineTwo.textString {
		return false
	}

	if plainTxtLineOne.turnLineTerminatorOff !=
		plainTxtLineTwo.turnLineTerminatorOff {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		plainTxtLineOne.newLineChars,
		plainTxtLineTwo.newLineChars) {
		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textLineSpecPlainTextElectron.
//
func (txtLinePlainTextElectron textLineSpecPlainTextElectron) ptr() *textLineSpecPlainTextElectron {

	if txtLinePlainTextElectron.lock == nil {
		txtLinePlainTextElectron.lock = new(sync.Mutex)
	}

	txtLinePlainTextElectron.lock.Lock()

	defer txtLinePlainTextElectron.lock.Unlock()

	return &textLineSpecPlainTextElectron{
		lock: new(sync.Mutex),
	}
}
