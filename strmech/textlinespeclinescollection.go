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
		ePrefix.XCpy("textLine"))

	if err != nil {
		return err
	}

	var newTextLine ITextLineSpecification

	newTextLine,
		err = textLine.CopyOutITextLine(
		ePrefix.XCpy(
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

// CopyIn - Copies the text line collection from an incoming
// instance of TextLineSpecLinesCollection
// ('incomingTxtLinesCol') to the current
// TextLineSpecLinesCollection instance ('txtLinesCol').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All data values in current TextLineSpecLinesCollection instance
// ('txtLinesCol') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtLinesCol        *TextLineSpecLinesCollection
//     - A pointer to an instance of TextLineSpecLinesCollection.
//       This method will NOT change the data values contained in
//       this instance.
//
//       All text line collection member elements in this
//       TextLineSpecLinesCollection instance will be copied to the
//       current TextLineSpecLinesCollection instance ('txtLinesCol').
//
//       If 'incomingTextLineCol' contains invalid member data
//       variables, this method will return an error.
//
//       If 'incomingTextLineCol' contains an empty, or zero
//       length, Text Lines Collection, an error will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
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
func (txtLinesCol *TextLineSpecLinesCollection) CopyIn(
	incomingTxtLinesCol *TextLineSpecLinesCollection,
	errorPrefix interface{}) error {

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
		"TextLineSpecLinesCollection.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecLinesCollectionNanobot{}.ptr().
		copyIn(
			txtLinesCol,
			incomingTxtLinesCol,
			ePrefix.XCpy(
				"incomingTxtLinesCol->"+
					"txtLinesCol"))
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecLinesCollection instance.
//
// If the current TextLineSpecLinesCollection instance contains
// invalid member variables, this method will return an error.
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
//  TextLineSpecLinesCollection
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecLinesCollection instance.
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
func (txtLinesCol *TextLineSpecLinesCollection) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecLinesCollection,
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
		"TextLineSpecLinesCollection.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecLinesCollection{}, err
	}

	return textLineSpecLinesCollectionNanobot{}.ptr().
		copyOut(
			txtLinesCol,
			ePrefix.XCpy(
				"txtLinesCol->"))
}

// DeleteTextLineMember - Deletes a member of the Text Lines
// collection encapsulated by the current instance of
// TextLineSpecLinesCollection.
//
// Input parameter 'zeroBasedIndex' is used to specify the
// collection member which will be deleted. If the operation
// is completed successfully, the total number of member elements
// in the collection will be reduced by one (1).
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in collection, an error will be
// returned. Also, if this method is called on an empty collection,
// an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// ----------------------------------------------------------------
//
// BE ADVISED
//
// If you delete the last element in the collection, the current
// instance TextLineSpecLinesCollection will be rendered invalid,
// and cannot be used until more Text Line elements are added to
// the collection
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  zeroBasedIndex             int
//     - Specifies the index of the member element in the Text
//       Lines collection which will be deleted. If this input
//       parameter is found to be invalid or if the Text Lines
//       collection is empty, an error will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) DeleteTextLineMember(
	zeroBasedIndex int,
	errorPrefix interface{}) error {

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
		"TextLineSpecLinesCollection."+
			"DeleteTextLineMember()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecLinesCollectionElectron{}.ptr().
		deleteTextLineElement(
			txtLinesCol,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"Delete Element txtLinesCol[%v]",
					zeroBasedIndex)))

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

	textLineSpecLinesCollectionAtom{}.ptr().
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

	textLineSpecLinesCollectionAtom{}.ptr().
		emptyCollection(txtLinesCol)

	return
}

// Equal - Receives a pointer to another instance of
// TextLineSpecLinesCollection and proceeds to compare the member
// variables to those of the current TextLineSpecLinesCollection
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
func (txtLinesCol *TextLineSpecLinesCollection) Equal(
	textLinesCol02 *TextLineSpecLinesCollection) bool {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	return textLineSpecLinesCollectionAtom{}.ptr().
		equalCollections(
			txtLinesCol,
			textLinesCol02)
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

// GetTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// If any of the text lines within the collection maintained by
// the current TextLineSpecLinesCollection instance are invalid,
// an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in the collection, an error will be
// returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  zeroBasedIndex             int
//     - Specifies the index of the member element in the Text
//       Lines collection which will be returned as a deep copy
//       of the original. If this input parameter is found to be
//       invalid or if the Text Lines collection is empty, an
//       error will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ITextLineSpecification
//     - If this method completes successfully, a deep copy of the
//       Text Line member element specified by input parameter
//       'zeroBasedIndex' will be returned.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) GetTextLine(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"DeleteTextLineMember()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			zeroBasedIndex,
			false,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesCol[%v]",
					zeroBasedIndex)))

	return iTextLineSpec, err
}

// GetTextLineCollection - Returns a deep copy of the text lines
// contained in the current TextLineSpecLinesCollection instance.
//
// These text lines are returned in an array of
// ITextLineSpecification objects.
//
// If the text line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// If any of the text lines within the collection maintained by
// the current TextLineSpecLinesCollection instance are invalid,
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
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  []ITextLineSpecification
//     - If this method completes successfully, a deep copy of the
//       text line collection maintained by the current
//       TextLineSpecLinesCollection instance will be returned.
//       These text lines are returned as an array of objects
//       implementing the ITextLineSpecification interface.
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
func (txtLinesCol *TextLineSpecLinesCollection) GetTextLineCollection(
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
		"TextLineSpecLinesCollection."+
			"GetTextLineCollection()",
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

	newTextLines :=
		make([]ITextLineSpecification, lenTxtLines)

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
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesCol.textLines[%v] invalid",
					i)))

		if err != nil {
			return nil, err
		}

		newTextLine,
			err2 := txtLinesCol.textLines[i].CopyOutITextLine(
			ePrefix.XCpy(
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

// InsertTextLine - Receives a Text Line instance in the form of a
// type ITextLineSpecification. This Text Line object is then
// inserted into Text Lines Collection maintained by the current
// instance of TextLineSpecLinesCollection.
//
// The Text Input input parameter, 'textLine', is inserted into
// the internal Text Lines array at the array element index
// position indicated by input parameter, 'indexId'.
//
// After this method completes, the number of elements in the Text
// Lines Collection will be increased by one.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  textLine                   ITextLineSpecification
//     - A Text Line object which implements the
//       ITextLineSpecification interface. A deep copy of this
//       text field will be inserted into the Text Lines Collection
//       maintained by the current instance of
//       TextLineSpecLinesCollection.
//
//       After the insertion operation is completed, the
//       'textLine' object will be located at array element
//       'indexId' immediately BEFORE the original array element
//       located at that array index.
//
//       NOTE: You will need to pass the concrete instance of
//       'textLine' as a pointer to the Text Line (&textLine).
//
//       If the 'iTextField' parameter are found to be invalid, an
//       error will be returned.
//
//
//  indexId                    int
//     - This index number designates the array element index in
//       the Text Lines Collection of the 'txtLinesCol' instance at
//       which the Text Line parameter, 'textLine' will be
//       inserted. This means that 'textLine' will be inserted
//       immediately BEFORE the array element specified by
//       'indexId' in the final Text Lines Array.
//
//       If the value of 'indexId' is less than zero, it will be
//       reset to zero. This means that the 'textLine' object
//       will be inserted in the first array element position of
//       the Text Fields Collection maintained by the current
//       TextLineSpecLinesCollection instance.
//
//       If the value of 'indexId' is greater the last array
//       element index in the Text Fields Collection, the
//       'textLine' object will be appended to the end of the Text
//       Lines Collection.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  lastIndexId                int
//     - If this method completes successfully, the internal array
//       index of the last text line object for the Text Lines
//       Collection maintained by the current
//       TextLineSpecLinesCollection instance will be returned as
//       an integer value.
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
//
func (txtLinesCol *TextLineSpecLinesCollection) InsertTextLine(
	textLine ITextLineSpecification,
	indexId int,
	errorPrefix interface{}) (
	lastIndexId int,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"InsertTextLine()",
		"")

	if err != nil {
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

	lastIndexId = len(txtLinesCol.textLines)

	if lastIndexId == 0 ||
		indexId >= lastIndexId {

		txtLinesCol.textLines = append(
			txtLinesCol.textLines,
			newTextLine)

		return lastIndexId, err
	}

	if indexId < 0 {

		indexId = 0

	}

	var oldTextLine ITextLineSpecification

	oldTextLine,
		err = txtLinesCol.textLines[indexId].CopyOutITextLine(
		ePrefix.XCpy(fmt.Sprintf(
			"oldTextLine<-txtLinesCol.textLines[%v]",
			indexId)))

	if err != nil {
		return lastIndexId, err
	}

	// arr := []int{1,2,3,4,5}
	// arr[:2]         [1,2]
	// arr[2:])        [3,4,5]
	// 	orig = append(orig[:index+1], orig[index:]...)

	txtLinesCol.textLines[indexId].Empty()
	txtLinesCol.textLines[indexId] = nil

	txtLinesCol.textLines = append(
		txtLinesCol.textLines[:indexId+1],
		txtLinesCol.textLines[indexId:]...)

	txtLinesCol.textLines[indexId+1] =
		oldTextLine

	txtLinesCol.textLines[indexId] =
		newTextLine

	return lastIndexId, err
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecLinesCollection
// instance to determine if they are valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ------------------------------------------------------------------------
//
// BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return 'false'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If all data elements encapsulated by the current instance
//       of TextLineSpecLinesCollection are valid, this returned
//       boolean value is set to 'true'. If any data values are
//       invalid, this return parameter is set to 'false'.
//
func (txtLinesCol *TextLineSpecLinesCollection) IsValidInstance() (
	isValid bool) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	isValid,
		_ = textLineSpecLinesCollectionAtom{}.ptr().
		testValidityOfTextLinesCollection(
			txtLinesCol,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecLinesCollection
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// ------------------------------------------------------------------------
//
// BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix         interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextLineSpecLinesCollection are found
//       to be invalid, this method will return an error.
//
//       Also, if the current instance of
//       TextLineSpecLinesCollection contains zero Text Line members
//       in the collection, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtLinesCol *TextLineSpecLinesCollection) IsValidInstanceError(
	errorPrefix interface{}) error {

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
		"TextLineSpecLinesCollection."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		testValidityOfTextLinesCollection(
			txtLinesCol,
			ePrefix.XCpy(
				"txtLinesCol"))

	return err
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
//       NOTE: You will need to pass the concrete instance of
//       'textLine' as a pointer to the Text Line (&textLine).
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
		"TextLineSpecLinesCollection."+
			"ReplaceTextLine()",
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

	if textLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLine' is 'nil' and invalid!\n",
			ePrefix.String())

		return err

	}

	err = textLine.IsValidInstanceError(
		ePrefix.XCpy("Input Parameter: textLine"))

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

	if replaceAtIndex < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is out of range and invalid!\n"+
			"'replaceAtIndex' is less than zero. The first index in the collection\n"+
			"is always zero.\n"+
			"Input parameter 'replaceAtIndex' = '%v'\n",
			ePrefix.String(),
			replaceAtIndex)

		return err
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
		ePrefix.XCpy(
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
