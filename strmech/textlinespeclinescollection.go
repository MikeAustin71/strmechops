package strmech

import (
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
