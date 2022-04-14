package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecLinesCollectionAtom struct {
	lock *sync.Mutex
}

// emptyCollection - Deletes all the text line objects from a text
// line collection maintained by an instance of
// TextLineSpecLinesCollection.
//
func (txtLinesColAtom *textLineSpecLinesCollectionAtom) emptyCollection(
	textLinesCol *TextLineSpecLinesCollection) {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	if textLinesCol == nil {
		return
	}

	lenTxtLineCol := len(textLinesCol.textLines)

	if lenTxtLineCol == 0 {
		textLinesCol.textLines = nil

		return
	}

	for i := 0; i < lenTxtLineCol; i++ {

		if textLinesCol.textLines[i] == nil {
			continue
		}

		textLinesCol.textLines[i].Empty()

		textLinesCol.textLines[i] = nil
	}

	textLinesCol.textLines = nil

	return
}

// ptr - Returns a pointer to a new instance of
// textLineSpecLinesCollectionAtom.
//
func (txtLinesColAtom textLineSpecLinesCollectionAtom) ptr() *textLineSpecLinesCollectionAtom {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	return &textLineSpecLinesCollectionAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTextLinesCollection - Receives a pointer to an
// instance of TextLineSpecLinesCollection ('textLineCol') and
// performs a diagnostic analysis to determine if that instance is
// valid in all respects.
//
// If the input parameter 'textLineCol' is determined to be invalid,
// this method will return a boolean flag ('isValid') of 'false'.
// In addition, an instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'textLineCol' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLineCol                *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection.
//       This object will be subjected to diagnostic analysis in
//       order to determine if all the member variables contain
//       valid data values.
//
//       Be advised, if the text lines collection of this
//       TextLineSpecLinesCollection instance is empty and contains
//       zero array elements, the TextLineSpecLinesCollection
//       instance will be declared, 'invalid'.
//
//       If any of the ITextLineSpecification objects contained in
//       the text lines collection encapsulated within
//       'textLineCol' are found to be invalid, the entire
//       TextLineSpecLinesCollection instance ('textLineCol') will
//       be declared invalid.
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
//  isValid                    bool
//     - If input parameter 'textLineCol' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'textLineCol' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'textLineCol' is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If input parameter, 'textLineCol' is found to be invalid,
//       this return parameter will be configured with an appropriate
//       error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinesColAtom *textLineSpecLinesCollectionAtom) testValidityOfTextLinesCollection(
	textLineCol *TextLineSpecLinesCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionAtom."+
			"testValidityOfTextLinesCollection()",
		"")

	if err != nil {
		return isValid, err
	}

	if textLineCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLineCol' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenTxtLines := len(textLineCol.textLines)

	if lenTxtLines == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Text Lines Collection ('textLineCol') is empty!\n"+
			"There are zero member elements in the Text Lines Collection array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenTxtLines; i++ {

		if textLineCol.textLines[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Line member element is invalid!\n"+
				"textLineCol.textLines[%v] has a 'nil' value.\n",
				ePrefix.XCpy(fmt.Sprintf(
					"textLineCol.textLines[%v]",
					i)), i)

			return isValid, err
		}

		err = textLineCol.textLines[i].IsValidInstanceError(
			ePrefix.XCpy(fmt.Sprintf(
				"textLineCol.textLines[%v]",
				i)))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
