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

// cfgNewStdTxtLineParameters - Configures a new
// TextFmtParamsLineColumnsDto object in the Standard Format
// Parameters collection maintained of the TextFormatterCollection.
//
// If an existing TextFmtParamsLineColumnsDto cannot be located,
// the new Standard Format Parameters object is appended to the
// Standard Format Parameters Collection.
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

	numOfNewStdFmtParams := newStdFmtParams.GetNumOfFieldFmtParams()

	lenStdTxtLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	foundStdTxtLineColFmt := false

	if lenStdTxtLineCol > 0 {

		for i := 0; i < lenStdTxtLineCol; i++ {

			if txtFmtCollection.stdTextLineParamCollection[i].
				GetNumOfFieldFmtParams() == numOfNewStdFmtParams {

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
	targetNumOfTextLineColumns int,
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

	if targetNumOfTextLineColumns < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNumOfTextLineColumns' is invalid!\n"+
			"'targetNumOfTextLineColumns' has a value less than one (1).\n"+
			"targetNumOfTextLineColumns Value  = '%v'\n",
			ePrefix.String(),
			targetNumOfTextLineColumns)

		return foundTxtFormatter,
			lineColsFormatter,
			err
	}

	lenOfStdLineFmtParams := len(txtFmtCollection.stdTextLineParamCollection)

	if lenOfStdLineFmtParams == 0 {

		return foundTxtFormatter, lineColsFormatter, err
	}

	for i := 0; i < lenOfStdLineFmtParams; i++ {

		if txtFmtCollection.stdTextLineParamCollection[i].GetNumOfFieldFmtParams() ==
			targetNumOfTextLineColumns {

			lineColsFormatter = txtFmtCollection.
				stdTextLineParamCollection[i].CopyOut()

			foundTxtFormatter = true

			return foundTxtFormatter, lineColsFormatter, err

		}
	}

	return foundTxtFormatter, lineColsFormatter, err
}
