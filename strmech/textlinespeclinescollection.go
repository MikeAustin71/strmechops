package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextLineSpecLinesCollection struct {
	textLines []ITextLineSpecification
	lock      *sync.Mutex
}

// AddTextLine - Adds a ITextLineSpecification object to the
// collection of text lines maintained by this instance of
// TextLineSpecLinesCollection.
//
// A deep copy of this ITextLineSpecification object ('textLine')
// is added to the end of the array of ITextLineSpecification
// objects encapsulated in this instance of
// TextLineSpecLinesCollection.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLine                   ITextLineSpecification
//     - A text line object which implements the
//       ITextLineSpecification interface. A deep copy of this
//       object will be added to the text lines collection
//       maintained by this instance of
//       TextLineSpecLinesCollection.
//
//       If member variable data values contained in this
//       'textLine' parameter are found to be invalid, an error
//       will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) AddTextLine(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.AddTextLine()",
		"")

	if err != nil {
		return err
	}

	err = textLine.IsValidInstanceError(
		ePrefix.XCtx("textLine"))

	if err != nil {
		return err
	}

	var newTextLine ITextLineSpecification

	newTextLine,
		err = textLine.CopyOutITextLine(
		ePrefix.XCtx(
			"textLine->newTextLine"))

	if err != nil {
		return err
	}

	txtLinesCol.textLines =
		append(
			txtLinesCol.textLines,
			newTextLine)

	return err
}

// Empty - Empties the text line collection and resets all member
// variables to their initial or zero values.
//
// Call this method when you intend to delete the
// TextLineSpecLinesCollection permanently as it will not be
// available for immediate reuse.
//
// If you wish to delete the text line collection and immediately
// reuse this TextLineSpecLinesCollection instance, use method
// TextLineSpecLinesCollection.EmptyTextLines() instead.
//
func (txtLinesCol *TextLineSpecLinesCollection) Empty() {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	textLineSpecLinesCollectionNanobot{}.ptr().
		emptyCollection(txtLinesCol)

	txtLinesCol.lock.Unlock()

	txtLinesCol.lock = nil

	return
}

// EmptyTextLines - Empties the text line collection and resets all member
// variables to their initial or zero values.
//
// This method only deletes the current text line collection. This
// means that the TextLineSpecLinesCollection instance is
// immediately for reuse and new text lines may be added to the now
// empty collection.
//
func (txtLinesCol *TextLineSpecLinesCollection) EmptyTextLines() {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	textLineSpecLinesCollectionNanobot{}.ptr().
		emptyCollection(txtLinesCol)

	return
}

// GetNumberOfTextLines - Returns the number of text lines
// encapsulated by the current TextLineSpecLinesCollection
// instance.
//
// Analyzing the number of text lines in the collection provides
// verification that text lines exist and are ready for formatting.
// Once properly formatted text lines may be presented for text
// display, file output or printing.
//
func (txtLinesCol *TextLineSpecLinesCollection) GetNumberOfTextLines() int {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	return len(txtLinesCol.textLines)
}

// GetTextLines - Returns a deep copy of the text fields contained
// in the current TextLineSpecStandardLine instance.
//
// These text fields are returned in an array of
// ITextFieldSpecification objects.
//
// If the text field collection maintained by the current
// TextLineSpecStandardLine instance is empty (contains zero
// elements), an error will be returned.
//
// If any of the text fields within the collection maintained by
// the current TextLineSpecStandardLine instance are invalid,
// an error will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  []ITextFieldSpecification
//     - If this method completes successfully, a deep copy of the
//       text field collection maintained by the current
//       TextLineSpecStandardLine instance will be returned. These
//       text fields are returned as an array of objects
//       implementing the ITextFieldSpecification interface.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) GetTextLines(
	errorPrefix interface{}) (
	[]ITextLineSpecification,
	error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.GetTextLines()",
		"")

	if err != nil {
		return nil, err
	}

	lenTxtLines := len(txtLinesCol.textLines)

	if lenTxtLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The text lines collection is empty!\n"+
			"TextLineSpecLinesCollection.textLines contains zero text line objects!\n",
			ePrefix.String())

		return nil, err
	}

	newTextLines := make([]ITextLineSpecification, lenTxtLines)

	for i := 0; i < lenTxtLines; i++ {

		if txtLinesCol.textLines[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Line element txtLinesCol.textLines[%v]\n"+
				"has a 'nil' value!\n",
				ePrefix.String(),
				i)

			return nil, err
		}

		err = txtLinesCol.textLines[i].IsValidInstanceError(
			ePrefix.XCtx(
				fmt.Sprintf(
					"txtLinesCol.textLines[%v] invalid",
					i)))

		if err != nil {
			return nil, err
		}

		newTextLine,
			err2 := txtLinesCol.textLines[i].CopyOutITextLine(
			ePrefix.XCtx(
				fmt.Sprintf(
					"txtLinesCol.textLines[%v] copy error",
					i)))

		if err2 != nil {
			return nil, err2
		}

		newTextLines[i] = newTextLine
	}

	return newTextLines, err
}

// ReplaceTextLine - Receives an object which implements the
// ITextLineSpecification interface. This object will replace an
// existing text line object within the text line collection
// maintained by this TextLineSpecLinesCollection instance.
//
// The text line object to be replaced must exist at the index
// specified by input parameter, 'replaceAtIndex'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLine                   ITextLineSpecification
//     - A text line object which implements the
//       ITextLineSpecification interface. A deep copy of this
//       object will replace an existing element within the
//       text lines collection maintained by this instance of
//       TextLineSpecLinesCollection. The text line object to
//       be replaced is identified by the collection element index
//       supplied by input parameter 'replaceAtIndex'.
//
//       If member variable data values contained in this
//       'textLine' parameter are found to be invalid, an error
//       will be returned.
//
//
//  replaceAtIndex             int
//     - The index of an element within the text lines collection
//       maintained by the current TextLineSpecLinesCollection
//       instance which will be replaced by input parameter
//       'textLine'.
//
//       Remember that the text fields collection maintained by
//       the current TextLineSpecLinesCollection instance is a zero
//       based array. Therefore, the first index in the collection
//       is zero (0).
//
//       If 'replaceAtIndex' proves to be an invalid index, an error
//       will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) ReplaceTextLine(
	textLine ITextLineSpecification,
	replaceAtIndex int,
	errorPrefix interface{}) (
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.ReplaceTextLine()",
		"")

	if err != nil {
		return err
	}

	lenOfTextLinesCol := len(txtLinesCol.textLines)

	if lenOfTextLinesCol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The text lines collection is empty and contains no text lines!\n"+
			"First add some text lines before trying to replace a text line.\n",
			ePrefix.String())

		return err
	}

	err = textLine.IsValidInstanceError(
		ePrefix.XCtx("Input Parameter: textLine"))

	if err != nil {
		return err
	}

	if replaceAtIndex < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is invalid!\n"+
			"'replaceAtIndex' is less than zero (0).\n"+
			"replaceAtIndex = '%v'\n",
			ePrefix.String(),
			replaceAtIndex)
	}

	lenOfTextLinesCol--

	if replaceAtIndex > lenOfTextLinesCol {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is out of range and invalid!\n"+
			"'replaceAtIndex' is greater than the maximum collection index.\n"+
			"The last element in the text lines collection is index '%v'.\n"+
			"Input parameter 'replaceAtIndex' = '%v'\n",
			ePrefix.String(),
			lenOfTextLinesCol,
			replaceAtIndex)

		return err
	}

	var newTextLine ITextLineSpecification

	newTextLine,
		err = textLine.CopyOutITextLine(
		ePrefix.XCtx(
			"newTextLine"))

	if err != nil {
		return err
	}

	if txtLinesCol.textLines[replaceAtIndex] != nil {

		txtLinesCol.textLines[replaceAtIndex].Empty()

		txtLinesCol.textLines[replaceAtIndex] = nil

	}

	txtLinesCol.textLines[replaceAtIndex] = newTextLine

	return err
}
