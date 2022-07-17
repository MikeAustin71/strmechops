package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textFormatterCollectionElectron - Provides helper methods for type
// TextFormatterCollection
type textFormatterCollectionElectron struct {
	lock *sync.Mutex
}

func (txtSolidLineElectron *textFormatterCollectionElectron) cfgNewStdTxtLineParameters(
	txtFmtCollection *TextFormatterCollection,
	newStdFmtParams TextFmtParamsLineColumnsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtSolidLineElectron.lock == nil {
		txtSolidLineElectron.lock = new(sync.Mutex)
	}

	txtSolidLineElectron.lock.Lock()

	defer txtSolidLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFormatterCollectionElectron."+
			"findStdTxtLineParameters()",
		"")

	if err != nil {

		return err

	}

	if txtFmtCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtCollection' is invalid!\n"+
			"'txtFmtCollection' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenStdTxtLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	foundStdTxtLineColFmt := false

	if lenStdTxtLineCol > 0 {

		for i := 0; i < lenStdTxtLineCol; i++ {

			if txtFmtCollection.stdTextLineParamCollection[i].
				FormatType == TxtFieldType.Line1Column() {

				txtFmtCollection.stdTextLineParamCollection[i].
					CopyIn(newStdFmtParams)

				foundStdTxtLineColFmt = true

			}

		}
	}

	if !foundStdTxtLineColFmt {

		txtFmtCollection.stdTextLineParamCollection =
			append(
				txtFmtCollection.stdTextLineParamCollection,
				newStdFmtParams)

	}

	return err
}

func (txtSolidLineElectron *textFormatterCollectionElectron) findStdTxtLineParameters(
	txtFmtCollection *TextFormatterCollection,
	searchForTextFieldType TextFieldType,
	errPrefDto *ePref.ErrPrefixDto) (
	foundTxtFormatter bool,
	lineColsFormatter TextFmtParamsLineColumnsDto,
	err error) {

	if txtSolidLineElectron.lock == nil {
		txtSolidLineElectron.lock = new(sync.Mutex)
	}

	txtSolidLineElectron.lock.Lock()

	defer txtSolidLineElectron.lock.Unlock()

	foundTxtFormatter = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFormatterCollectionElectron."+
			"findStdTxtLineParameters()",
		"")

	if err != nil {

		return foundTxtFormatter,
			lineColsFormatter,
			err

	}

	if txtFmtCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtCollection' is invalid!\n"+
			"'txtFmtCollection' is a 'nil' pointer.\n",
			ePrefix.String())

		return foundTxtFormatter,
			lineColsFormatter,
			err
	}

	if !searchForTextFieldType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'searchForTextFieldType' is invalid!\n"+
			"searchForTextFieldType String Value  = '%v'\n"+
			"searchForTextFieldType Integer Value = '%v'\n",
			ePrefix.String(),
			searchForTextFieldType.String(),
			searchForTextFieldType.XValueInt())

		return foundTxtFormatter,
			lineColsFormatter,
			err
	}

	lenOfStdLineFmtParams := len(txtFmtCollection.stdTextLineParamCollection)

	if lenOfStdLineFmtParams == 0 {

		return foundTxtFormatter, lineColsFormatter, err
	}

	for i := 0; i < lenOfStdLineFmtParams; i++ {

		if txtFmtCollection.stdTextLineParamCollection[i].FormatType ==
			searchForTextFieldType {

			lineColsFormatter = txtFmtCollection.
				stdTextLineParamCollection[i].CopyOut()

			foundTxtFormatter = true

			return foundTxtFormatter, lineColsFormatter, err

		}
	}

	return foundTxtFormatter, lineColsFormatter, err
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
