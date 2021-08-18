package strmech

import "sync"

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
}

// emptyTextFields - Receives an array of objects implementing the
// ITextFieldSpecification interface. The ITextFieldSpecification
// describes text fields which serve as the building blocks for a
// single line of text.
//
// Text Fields are encapsulated in the TextLineSpecStandardLine and
// used to produce single lines of text for text display, file output
// or printing.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the text fields stored in the text field collection
// maintained by input parameter 'txtStdLine' will be deleted.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) emptyTextFields(
	txtFields []ITextFieldSpecification) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	if txtFields == nil {
		return
	}

	lenTextFields := len(txtFields)

	if lenTextFields == 0 {

		txtFields = nil

		return
	}

	for i := 0; i < lenTextFields; i++ {

		if txtFields[i] == nil {
			continue
		}

		txtFields[i].Empty()

		txtFields[i] = nil
	}

	txtFields = nil

	return
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineElectron.
//
func (txtStdLineElectron textLineSpecStandardLineElectron) ptr() *textLineSpecStandardLineElectron {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	return &textLineSpecStandardLineElectron{
		lock: new(sync.Mutex),
	}
}
