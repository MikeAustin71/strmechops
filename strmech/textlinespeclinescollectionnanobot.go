package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecLinesCollectionNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTextLineCol' to input parameter 'targetTextLineCol'.
//
// Be advised that the pre-existing data fields in input parameter
// 'targetTextLineCol' will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetTextLineCol          *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection.
//       Data extracted from input parameter 'incomingTextLineCol'
//       will be copied to this input parameter,
//       'targetTextLineCol'. If this method completes
//       successfully, all member data variables encapsulated in
//       'targetTextLineCol' will be identical to those contained
//       in input parameter, 'incomingTextLineCol'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetTextLineCol' will be overwritten and
//       deleted.
//
//
//  incomingTextLineCol        *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection.
//
//       All data values in this TextLineSpecLinesCollection
//       instance will be copied to input parameter
//       'targetTextLineCol'.
//
//       The original member variable data values encapsulated in
//       'incomingTextLineCol' will remain unchanged and will NOT
//       be overwritten or deleted.
//
//       If 'incomingTextLineCol' contains invalid member data
//       variables, this method will return an error.
//
//       If 'incomingTextLineCol' contains an empty, or zero
//       length, Text Lines Collection, an error will be returned.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) copyIn(
	targetTextLineCol *TextLineSpecLinesCollection,
	incomingTextLineCol *TextLineSpecLinesCollection,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionNanobot."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTextLineCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTextLineCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTextLineCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTextLineCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	_,
		err = txtLinesColAtom.
		testValidityOfTextLinesCollection(
			incomingTextLineCol,
			ePrefix.XCtx("incomingTextLineCol"))

	if err != nil {
		return err
	}

	txtLinesColAtom.
		emptyCollection(targetTextLineCol)

	lenTxtLineCol := len(incomingTextLineCol.textLines)

	if lenTxtLineCol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTextLineCol' is invalid!\n"+
			"The 'incomingTextLineCol' text lines collection is empty.\n"+
			"There is nothing to copy.\n",
			ePrefix.XCtxEmpty().String())

		return err
	}

	targetTextLineCol.textLines =
		make([]ITextLineSpecification, lenTxtLineCol)

	var newTextLine ITextLineSpecification

	for i := 0; i < lenTxtLineCol; i++ {

		if incomingTextLineCol.textLines[i] == nil {

			err = fmt.Errorf("%v\n"+
				"Error: Text Line element incomingTextLineCol.textLines[%v]\n"+
				"has a 'nil' value!\n",
				ePrefix.String(),
				i)

			return err
		}

		newTextLine,
			err = incomingTextLineCol.textLines[i].CopyOutITextLine(
			ePrefix.XCtx(
				fmt.Sprintf(
					"incomingTextLineCol.textLines[%v] copy error",
					i)))

		if err != nil {
			return err
		}

		targetTextLineCol.textLines[i] = newTextLine
	}

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecLinesCollectionNanobot.
//
func (txtLinesColNanobot textLineSpecLinesCollectionNanobot) ptr() *textLineSpecLinesCollectionNanobot {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	return &textLineSpecLinesCollectionNanobot{
		lock: new(sync.Mutex),
	}
}
