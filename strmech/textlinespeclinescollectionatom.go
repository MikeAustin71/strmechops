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

// equalCollections - Analyzes two Text Line Collections to
// determine if they are equal in all respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  textLinesCol01             *TextLineSpecLinesCollection
//     - The instance of TextLineSpecLinesCollection will be
//       compared to a second TextLineSpecLinesCollection
//       instance (textLinesCol02) in order to determine if
//       both instances are equal in all respects.
//
//
//  textLinesCol02             *TextLineSpecLinesCollection
//     - The instance of TextLineSpecLinesCollection will be
//       compared to a second TextLineSpecLinesCollection
//       instance (textLinesCol01) in order to determine if
//       both instances are equal in all respects.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If input parameters 'textLinesCol01' and 'textLinesCol02'
//       are determined to be equal in all respects, this returned
//       boolean value is set to 'true'.
//
//       If an error occurs or if the two instance of
//       TextLineSpecLinesCollection are NOT equivalent, this
//       returned boolean value is set to 'false'.

func (txtLinesColAtom *textLineSpecLinesCollectionAtom) equalCollections(
	textLinesCol01 *TextLineSpecLinesCollection,
	textLinesCol02 *TextLineSpecLinesCollection) bool {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	if textLinesCol01 == nil ||
		textLinesCol02 == nil {

		return false
	}

	lenTextLinesCol01 := len(textLinesCol01.textLines)

	if lenTextLinesCol01 !=
		len(textLinesCol02.textLines) {

		return false
	}

	if lenTextLinesCol01 == 0 {
		return true
	}

	for i := 0; i < lenTextLinesCol01; i++ {

		if !textLinesCol01.textLines[i].EqualITextLine(
			textLinesCol02.textLines[i]) {

			return false
		}

	}

	return true
}

// insertTextLine - Receives an instance of Text Line in the form
// of a type ITextLineSpecification. This Text Line object is then
// inserted into Text Lines Collection maintained by the
// TextLineSpecLinesCollection instance passed as input parameter,
// 'textLinesCol'.
//
// The Text Line input parameter, 'textLine', is inserted into
// the internal Text Lines collection array at the array element
// index position indicated by input parameter, 'zeroBasedIndex'.
//
// After this method completes, the number of elements in the Text
// Lines Collection will be increased by one.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  textLinesCol               *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection.
//       The Text Line Collection encapsulated by this instance
//       will have an additional Text Line member inserted into the
//       Collection. The Text Line object to be inserted is
//       specified by input parameter, 'textLine'. This Text Line
//       will be inserted into the Text Lines Collection at the
//       array index designated by input parameter,
//       'zeroBasedIndex'.
//
//
//  textLine                   ITextLineSpecification
//     - A Text Line object which implements the
//       ITextLineSpecification interface. A deep copy of this
//       Text Line will be inserted into the Text Lines Collection
//       maintained by the TextLineSpecLinesCollection instance,
//       'textLinesCol'.
//
//       After the insertion operation is completed, the
//       'textLine' object will be located at array element
//       'zeroBasedIndex' immediately BEFORE the original array
//       element previously located at that array index.
//
//       NOTE: You will need to pass the concrete instance of
//       'textLine' as a pointer to the Text Line (&textLine).
//
//       If the 'textLine' parameter is found to be invalid, an
//       error will be returned.
//
//
//  zeroBasedIndex             int
//     - This index number designates the array element index in
//       the Text Lines Collection ('txtLinesCol') at which the
//       Text Line parameter, 'textLine' will be inserted. This
//       means that 'textLine' will be inserted immediately
//       BEFORE the array element specified by 'zeroBasedIndex'
//       in the final Text Lines Array.
//
//       For example, if 'zeroBasedIndex' is set to '4', the
//       original Text Line object at index '4' will be moved to
//       index position '5' after the insertion operation is completed.
//
//       If the value of 'zeroBasedIndex' is less than zero, it
//       will be reset to zero. This means that the 'textLine'
//       object will be inserted in the first array element
//       position of the Text Fields Collection maintained by the
//       current TextLineSpecLinesCollection instance.
//
//       If the value of 'zeroBasedIndex' is greater the last array
//       element index in the Text Fields Collection, the
//       'textLine' object will be appended to the end of the Text
//       Lines Collection.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  lastIndexId                int
//     - If this method completes successfully, the internal array
//       index of the last text line object for the Text Lines
//       Collection maintained by the TextLineSpecLinesCollection
//       instance 'textLinesCol' will be returned as an integer
//       value. Remember, this is a zero based index value which is
//       always one less than the length of the Text Line
//       Collection.
//
//       In the event of an error, 'lastIndexId' will be set to a
//       value of minus one (-1).
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesColAtom *textLineSpecLinesCollectionAtom) insertTextLine(
	textLinesCol *TextLineSpecLinesCollection,
	textLine ITextLineSpecification,
	zeroBasedIndex int,
	errPrefDto *ePref.ErrPrefixDto) (
	lastIndexId int,
	err error) {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionAtom."+
			"insertTextLine()",
		"")

	if err != nil {
		return lastIndexId, err
	}

	if textLinesCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'textLinesCol' is a nil pointer!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	if textLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLine' is 'nil' and invalid!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	err = textLine.IsValidInstanceError(
		ePrefix.XCpy("Input Parameter: textLine"))

	if err != nil {
		return lastIndexId, err
	}

	var newTextLine ITextLineSpecification

	newTextLine,
		err = textLine.CopyOutITextLine(
		ePrefix.XCpy("textLine->newTextLine"))

	if err != nil {
		return lastIndexId, err
	}

	lastIndexId = len(textLinesCol.textLines)

	if lastIndexId == 0 ||
		zeroBasedIndex >= lastIndexId {

		textLinesCol.textLines = append(
			textLinesCol.textLines,
			newTextLine)

		return lastIndexId, err
	}

	if zeroBasedIndex < 0 {

		zeroBasedIndex = 0

	}

	var oldTextLine ITextLineSpecification

	oldTextLine,
		err = textLinesCol.textLines[zeroBasedIndex].
		CopyOutITextLine(
			ePrefix.XCpy(fmt.Sprintf(
				"oldTextLine<-textLinesCol.textLines[%v]",
				zeroBasedIndex)))

	if err != nil {

		lastIndexId = -1

		return lastIndexId, err
	}

	// arr := []int{1,2,3,4,5}
	// arr[:2]         [1,2]
	// arr[2:])        [3,4,5]
	// 	orig = append(orig[:index+1], orig[index:]...)

	if textLinesCol.textLines[zeroBasedIndex] != nil {

		textLinesCol.textLines[zeroBasedIndex].Empty()

		textLinesCol.textLines[zeroBasedIndex] = nil

	}

	textLinesCol.textLines = append(
		textLinesCol.textLines[:zeroBasedIndex+1],
		textLinesCol.textLines[zeroBasedIndex:]...)

	textLinesCol.textLines[zeroBasedIndex+1] =
		oldTextLine

	textLinesCol.textLines[zeroBasedIndex] =
		newTextLine

	return lastIndexId, err
}

// peekPopTextLine - Performs either a 'Peek' or 'Pop' operation
// on an array element in the Text Lines Collection as specified
// by the input parameter, 'popTextLine'.
//
// A 'Pop' operation returns a deep copy of the designated Text
// Line element in the Text Lines Collection and then DELETES that
// designated array element. The designated array element is
// specified by input parameter, 'zeroBasedIndex'.
//
// On the other hand, a 'Peek' operation will return a deep copy of
// the designated Text Line in the Text lines Collection and WILL
// NOT delete that array element. The designated array element
// therefore remains in the collection after the 'Peek' operation
// is completed.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLinesCol               *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection. A
//       deep copy of the designated Text Line in the Text Lines
//       Collection for this instance of TextLineSpecLinesCollection
//       will be returned to the calling function. The returned Text
//       Line is designated by input parameter, 'zeroBasedIndex'.
//
//       Depending on the value of input parameter, 'popTextLine',
//       either a 'Peek' or 'Pop' operation will be performed on
//       the designated Text Line in the Text Lines Collection.
//
//
//  zeroBasedIndex             int
//     - The index number of the array element in the Text Lines
//       Collection on which the 'Pop' or 'Peek' operation will be
//       performed.
//
//
//  popTextLine                bool
//     - If this parameter is set to 'true', it signals that a
//       'Pop' operation will be performed on the designated Text
//       Line in the Text Lines Collection encapsulated in
//       parameter 'textLinesCol'. A 'Pop' operation will DELETE
//       the designated Text Field from the Text Fields Collection.
//
//       If this parameter is set to 'false', it signals that a
//       'Peek' operation will be performed on the designated Text
//       Line in the Text Lines Collection encapsulated in
//       parameter 'textLinesCol'. A 'Peek' operation means that
//       the designated Text Line element in the Text Lines
//       Collection WILL NOT be deleted and will remain in the
//       collection.
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
//  iTextLine                  ITextLineSpecification
//     - If this method completes successfully, a deep copy of
//       if the designated member of the Text Lines Collection
//       will be returned to the calling function. The returned
//       object will implement the ITextLineSpecification
//       interface.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinesColAtom *textLineSpecLinesCollectionAtom) peekPopTextLine(
	textLinesCol *TextLineSpecLinesCollection,
	zeroBasedIndex int,
	popTextLine bool,
	errPrefDto *ePref.ErrPrefixDto) (
	iTextLine ITextLineSpecification,
	err error) {

	if txtLinesColAtom.lock == nil {
		txtLinesColAtom.lock = new(sync.Mutex)
	}

	txtLinesColAtom.lock.Lock()

	defer txtLinesColAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionAtom."+
			"peekPopTextLine()",
		"")

	if err != nil {
		return iTextLine, err
	}

	if textLinesCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'textLinesCol' is a nil pointer!\n",
			ePrefix.String())

		return iTextLine, err
	}

	lenTextLinesCollection := len(textLinesCol.textLines)

	lastIdx := lenTextLinesCollection - 1

	if lenTextLinesCollection == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Lines Collection, 'textLinesCol.textLines' is EMPTY!\n",
			ePrefix.String())

		return iTextLine, err
	}

	if zeroBasedIndex < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is less than zero.\n"+
			"indexId = '%v'\n",
			ePrefix.String(),
			zeroBasedIndex)

		return iTextLine, err
	}

	if zeroBasedIndex > lastIdx {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is greater than the last index\n"+
			"in the Text Lines Collection.\n"+
			"Last index in collection = '%v'\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			lastIdx,
			zeroBasedIndex)

		return iTextLine, err
	}

	if textLinesCol.textLines[zeroBasedIndex] == nil {

		err = fmt.Errorf("%v - ERROR\n"+
			"The requested Text Lines array element is invalid!\n"+
			"textLinesCol.textLines[%v] has a 'nil' value.\n",
			ePrefix.String(),
			zeroBasedIndex)

		return iTextLine, err
	}

	iTextLine,
		err = textLinesCol.textLines[zeroBasedIndex].CopyOutITextLine(
		ePrefix.XCpy(
			fmt.Sprintf(
				"textLinesCol.textLines[%v]",
				zeroBasedIndex)))

	if err != nil {
		return iTextLine, err
	}

	if !popTextLine {
		// popTextLine == false
		// This means a 'peek' operation is
		// being performed and the designated
		// array element will NOT be deleted.
		return iTextLine, err
	}

	err =
		textLineSpecLinesCollectionElectron{}.ptr().
			deleteTextLineElement(
				textLinesCol,
				zeroBasedIndex,
				ePrefix.XCpy(
					fmt.Sprintf(
						"Deleting textLinesCol.textLines[%v]",
						zeroBasedIndex)))

	return iTextLine, err
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

	var err2 error

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

		err2 = textLineCol.textLines[i].IsValidInstanceError(
			ePrefix.XCpy(fmt.Sprintf(
				"textLineCol.textLines[%v]",
				i)))

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"ERROR: textLineCol.textLines[%v] is invalid!\n"+
				"%v\n",
				ePrefix.String(),
				i,
				err2)

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
