package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textFormatterCollectionNanobot - Provides helper methods for
// TextFormatterCollection.
type textFormatterCollectionNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFormatterCollection to a destination instance of
// TextFormatterCollection.
func (textFmtCollectionNanobot *textFormatterCollectionNanobot) copy(
	destinationTxtFmtCol *TextFormatterCollection,
	sourceTxtFmtCol *TextFormatterCollection,
	errPrefDto *ePref.ErrPrefixDto) error {

	if textFmtCollectionNanobot.lock == nil {
		textFmtCollectionNanobot.lock = new(sync.Mutex)
	}

	textFmtCollectionNanobot.lock.Lock()

	defer textFmtCollectionNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFormatterCollectionNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtFmtCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFmtCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtFmtCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFmtCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textFormatterCollectionMolecule).
		empty(destinationTxtFmtCol)

	lenItems := len(sourceTxtFmtCol.fmtCollection)

	if lenItems > 0 {

		destinationTxtFmtCol.fmtCollection =
			make([]TextFormatterDto, lenItems)

		for i := 0; i < lenItems; i++ {

			destinationTxtFmtCol.fmtCollection[i].CopyIn(
				sourceTxtFmtCol.fmtCollection[i])

		}
	}

	lenItems = len(sourceTxtFmtCol.stdTextLineParamCollection)

	if lenItems > 0 {

		destinationTxtFmtCol.stdTextLineParamCollection =
			make([]TextFmtParamsLineColumnsDto, lenItems)

		for i := 0; i < lenItems; i++ {

			destinationTxtFmtCol.stdTextLineParamCollection[i].CopyIn(
				sourceTxtFmtCol.stdTextLineParamCollection[i])

		}

	}

	return err
}
