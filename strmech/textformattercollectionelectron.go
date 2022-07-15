package strmech

import (
	"sync"
)

// textFormatterCollectionElectron - Provides helper methods for type
// TextFormatterCollection
type textFormatterCollectionElectron struct {
	lock *sync.Mutex
}

func (txtSolidLineElectron *textFormatterCollectionElectron) findStdTxtLineParameters(
	txtFmtCollection *TextFormatterCollection,
	searchForTextFieldType TextFieldType) (
	foundTxtFormatter bool,
	lineColsFormatter TextFmtParamsLineColumns) {

	if txtSolidLineElectron.lock == nil {
		txtSolidLineElectron.lock = new(sync.Mutex)
	}

	txtSolidLineElectron.lock.Lock()

	defer txtSolidLineElectron.lock.Unlock()

	foundTxtFormatter = false

	if txtFmtCollection == nil {
		return foundTxtFormatter, lineColsFormatter
	}

	if !searchForTextFieldType.XIsValid() {
		return foundTxtFormatter, lineColsFormatter
	}

	lenOfStdLineFmtParams := len(txtFmtCollection.stdTextLineParamCollection)

	if lenOfStdLineFmtParams == 0 {

		return foundTxtFormatter, lineColsFormatter
	}

	for i := 0; i < lenOfStdLineFmtParams; i++ {

		if txtFmtCollection.stdTextLineParamCollection[i].formatType ==
			searchForTextFieldType {

			lineColsFormatter = txtFmtCollection.
				stdTextLineParamCollection[i].CopyOut()

			foundTxtFormatter = true

			return foundTxtFormatter, lineColsFormatter

		}
	}

	return foundTxtFormatter, lineColsFormatter
}

// ptr - Returns a pointer to a new instance of
// textFormatterCollectionElectron.
//
func (txtSolidLineElectron textFormatterCollectionElectron) ptr() *textFormatterCollectionElectron {

	if txtSolidLineElectron.lock == nil {
		txtSolidLineElectron.lock = new(sync.Mutex)
	}

	txtSolidLineElectron.lock.Lock()

	defer txtSolidLineElectron.lock.Unlock()

	return &textFormatterCollectionElectron{
		lock: new(sync.Mutex),
	}
}
