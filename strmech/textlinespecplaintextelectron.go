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
