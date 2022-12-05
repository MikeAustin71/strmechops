package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textLineSpecLinesCollectionNanobot struct {
	lock *sync.Mutex
}

// addTextLine - Adds a ITextLineSpecification object to the
// end of the Text Line collection maintained by this instance of
// TextLineSpecLinesCollection.
//
// A deep copy of this ITextLineSpecification object ('textLine')
// is added to the end of the array of ITextLineSpecification
// objects encapsulated in this instance of
// TextLineSpecLinesCollection.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLinesCol               *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection.
//	     A deep copy of input parameter 'textLine' will be appended
//	     to the end of the Text Lines Collection maintained by this
//	     instance of TextLineSpecLinesCollection.
//
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be appended to the end of the text lines
//	     collection maintained by the 'textLinesCol' instance
//	     of TextLineSpecLinesCollection.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) addTextLine(
	textLinesCol *TextLineSpecLinesCollection,
	textLine ITextLineSpecification,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionNanobot."+
			"addTextLine()",
		"")

	if err != nil {
		return err
	}

	if textLinesCol == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLinesCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if textLine == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLine' is 'nil'\n",
			ePrefix.String())

		return err
	}

	err2 := textLine.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLine' is failed validation test!\n"+
			"Validation Error:\n%v\n",
			ePrefix.String(),
			err2.Error())

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

	textLinesCol.textLines =
		append(
			textLinesCol.textLines,
			newTextLine)

	return err
}

// copyIn - Copies all data from input parameter
// 'incomingTextLineCol' to input parameter 'targetTextLineCol'.
//
// Be advised that the pre-existing data fields in input parameter
// 'targetTextLineCol' will be overwritten and deleted.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	targetTextLineCol          *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection.
//	     Data extracted from input parameter 'incomingTextLineCol'
//	     will be copied to this input parameter,
//	     'targetTextLineCol'. If this method completes
//	     successfully, all member data variables encapsulated in
//	     'targetTextLineCol' will be identical to those contained
//	     in input parameter, 'incomingTextLineCol'.
//
//	     Be advised that the pre-existing data fields in input
//	     parameter 'targetTextLineCol' will be overwritten and
//	     deleted.
//
//
//	incomingTextLineCol        *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection.
//
//	     All data values in this TextLineSpecLinesCollection
//	     instance will be copied to input parameter
//	     'targetTextLineCol'.
//
//	     The original member variable data values encapsulated in
//	     'incomingTextLineCol' will remain unchanged and will NOT
//	     be overwritten or deleted.
//
//	     If 'incomingTextLineCol' contains invalid member data
//	     variables, this method will return an error.
//
//	     If 'incomingTextLineCol' contains an empty, or zero
//	     length, Text Lines Collection, an error will be returned.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
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
			ePrefix.XCpy("incomingTextLineCol"))

	if err != nil {
		return err
	}

	txtLinesColAtom.
		emptyCollection(targetTextLineCol)

	lenTxtLineCol := len(incomingTextLineCol.textLines)

	// If lenTxtLineCol is zero, Validity Check above
	// would fail.

	targetTextLineCol.textLines =
		make([]ITextLineSpecification, lenTxtLineCol)

	var newTextLine ITextLineSpecification

	for i := 0; i < lenTxtLineCol; i++ {

		// If incomingTextLineCol.textLines[i] == nil
		// Validity check above would fail.

		newTextLine,
			err = incomingTextLineCol.textLines[i].CopyOutITextLine(
			ePrefix.XCpy(
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

// copyOut - Returns a deep copy of the input parameter
// 'textLineCol'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLineCol          *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of
//	     TextLineSpecLinesCollection.
//
//	     If the member variable data values encapsulated by this
//	     'textLineCol' are found to be invalid, this method will
//	     return an error
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecLinesCollection
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'textLineCol' will be created and returned
//	     in a new instance of TextLineSpecLinesCollection.
//
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) copyOut(
	textLineCol *TextLineSpecLinesCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecLinesCollection, error) {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtLinesCol := TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionNanobot."+
			"copyOut()",
		"")

	if err != nil {
		return newTxtLinesCol, err
	}

	if textLineCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTextLineCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return newTxtLinesCol, err
	}

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}
	_,
		err = txtLinesColAtom.
		testValidityOfTextLinesCollection(
			textLineCol,
			ePrefix.XCpy("incomingTextLineCol"))

	if err != nil {
		return newTxtLinesCol, err
	}

	lenTxtLineCol := len(textLineCol.textLines)

	// If len(textLineCol.textLines)==0 it would have
	// been caught and returned as an error by the
	// Validity Check above.

	newTxtLinesCol.textLines =
		make([]ITextLineSpecification, lenTxtLineCol)

	var newTextLine ITextLineSpecification

	for i := 0; i < lenTxtLineCol; i++ {

		// If textLineCol.textLines[i] == nil it would have
		// been caught and returned as an error by the
		// Validity Check above.

		newTextLine,
			err = textLineCol.textLines[i].CopyOutITextLine(
			ePrefix.XCpy(
				fmt.Sprintf(
					"incomingTextLineCol.textLines[%v] copy error",
					i)))

		if err != nil {
			return TextLineSpecLinesCollection{}, err
		}

		newTxtLinesCol.textLines[i] = newTextLine
	}

	return newTxtLinesCol, err
}

//	getFormattedText
//
//	Generates formatted text strings for all member of a
//	Text Line Specification Collection (textLinesCol) and
//	writes them to a String Builder (strBuilder) passed
//	as an input parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//			A pointer to an instance of strings.Builder.
//			The formatted text strings generated from
//			input parameter 'textLinesCol' will be
//			written to this instance of strings.Builder.
//
//	textLinesCol				*TextLineSpecLinesCollection
//
//		The formatted text generated from the member Text
//		Line Specifications contained in this collection
//		will be returned in a string array.
//
//		If this instance of TextLineSpecLinesCollection
//		is judged to be invalid, an error will be
//		returned.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) getFormattedText(
	strBuilder *strings.Builder,
	textLinesCol *TextLineSpecLinesCollection,
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
			"getFormattedText()",
		"")

	if err != nil {
		return err
	}

	if textLinesCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLinesCol' (Text Line Collection)\n"+
			"is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' (String Builder)\n"+
			"is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		testValidityOfTextLinesCollection(
			textLinesCol,
			ePrefix.XCpy(
				"textLinesCol"))

	if err != nil {
		return err
	}

	lenColElements := len(textLinesCol.textLines)

	var str string

	for i := 0; i < lenColElements; i++ {

		str,
			err = textLinesCol.textLines[i].GetFormattedText(
			ePrefix.XCpy(
				fmt.Sprintf(
					"textLinesCol.textLines[%v]",
					i)))

		if err != nil {
			return err
		}

		strBuilder.WriteString(str)

	}

	return err

}

//	getFormattedTextStrArray
//
//	Returns a string array containing the formatted text
//	strings generated from an instance of
//	TextLineSpecLinesCollection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strArrayDto					*StringArrayDto
//
//		The formatted text strings generated from
//		'textLineCol' will be written to a string array
//		contained in this instance of StringArrayDto.
//
//	textLinesCol				*TextLineSpecLinesCollection
//
//		The formatted text generated from the member Text
//		Line Specifications contained in this collection
//		will be returned in a string array (strArrayDto).
//
//		If this instance of TextLineSpecLinesCollection
//		is judged to be invalid, an error will be
//		returned.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) getFormattedTextStrArray(
	strArrayDto *StringArrayDto,
	textLinesCol *TextLineSpecLinesCollection,
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
			"getFormattedTextStrArray()",
		"")

	if err != nil {
		return err
	}

	if textLinesCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLinesCol' (Text Line Collection)\n"+
			"is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if strArrayDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strArrayDto' (String Array Dto)\n"+
			"is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		testValidityOfTextLinesCollection(
			textLinesCol,
			ePrefix.XCpy(
				"textLinesCol"))

	if err != nil {
		return err
	}

	lenColElements := len(textLinesCol.textLines)

	var str string

	for i := 0; i < lenColElements; i++ {

		str,
			err = textLinesCol.textLines[i].GetFormattedText(
			ePrefix.XCpy(
				fmt.Sprintf(
					"textLinesCol.textLines[%v]",
					i)))

		if err != nil {
			return err
		}

		strArrayDto.AddString(str)

	}

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecLinesCollectionNanobot.
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
