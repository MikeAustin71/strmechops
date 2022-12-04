package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecLinesCollection
//
// A collection of objects implementing the
// ITextLineSpecification interface.
type TextLineSpecLinesCollection struct {
	textLines []ITextLineSpecification
	lock      *sync.Mutex
}

// AddTextLine - Adds a ITextLineSpecification object to the
// end of the Text Line collection maintained by this instance of
// TextLineSpecLinesCollection.
//
// A deep copy of this ITextLineSpecification object ('textLine')
// is appended to the end of the array of ITextLineSpecification
// objects maintained by this instance of
// TextLineSpecLinesCollection.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by this instance of
//	     TextLineSpecLinesCollection.
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

	txtLinesColNanobot := textLineSpecLinesCollectionNanobot{}

	err = txtLinesColNanobot.addTextLine(
		txtLinesCol,
		textLine,
		ePrefix.XCpy(
			"txtLinesCol<-textLine"))

	return err
}

// CopyIn - Copies the text line collection from an incoming
// instance of TextLineSpecLinesCollection
// ('incomingTxtLinesCol') to the current
// TextLineSpecLinesCollection instance ('txtLinesCol').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All data values in current TextLineSpecLinesCollection instance
// ('txtLinesCol') will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtLinesCol        *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection.
//	     This method will NOT change the data values contained in
//	     this instance.
//
//	     All text line collection member elements in this
//	     TextLineSpecLinesCollection instance will be copied to the
//	     current TextLineSpecLinesCollection instance ('txtLinesCol').
//
//	     If 'incomingTextLineCol' contains invalid member data
//	     variables, this method will return an error.
//
//	     If 'incomingTextLineCol' contains an empty, or zero
//	     length, Text Lines Collection, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
// ----------------------------------------------------------------
//
// Input Parameters
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
//	TextLineSpecLinesCollection
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextLineSpecLinesCollection instance.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
// # BE ADVISED
//
// If you delete the last element in the collection, the current
// instance TextLineSpecLinesCollection will be rendered invalid,
// and cannot be used until more Text Line elements are added to
// the collection
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be deleted. If this input
//	     parameter is found to be invalid or if the Text Lines
//	     collection is empty, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in the collection, an error will be
// returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( GetTextField() ) is functionally equivalent to
// method:
//
//	TextLineSpecLinesCollection.PeekAtTextField()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Text Lines collection is empty, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.GetTextLine(
//	           2,  // Return Text Line at index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
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
			"GetTextLine()",
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     text line collection maintained by the current
//	     TextLineSpecLinesCollection instance will be returned.
//	     These text lines are returned as an array of objects
//	     implementing the ITextLineSpecification interface.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

		var newTextLine ITextLineSpecification

		newTextLine,
			err = txtLinesCol.textLines[i].CopyOutITextLine(
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesCol.textLines[%v] copy error",
					i)))

		if err != nil {
			return nil, err
		}

		newTextLines[i] = newTextLine
	}

	return newTextLines, err
}

// InsertTextLine - Receives a Text Line instance in the form of a
// type ITextLineSpecification. This Text Line object is then
// inserted into the Text Lines Collection maintained by the
// current instance of TextLineSpecLinesCollection.
//
// The Text Line input parameter, 'textLine', is inserted into
// the internal Text Lines collection at the array element index
// position indicated by input parameter, 'zeroBasedIndex'.
//
// After this method completes, the number of elements in the Text
// Lines Collection will be increased by one.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A Text Line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     text field will be inserted into the Text Lines Collection
//	     maintained by the current instance of
//	     TextLineSpecLinesCollection.
//
//	     After the insertion operation is completed, the
//	     'textLine' object will be located at array element
//	     'zeroBasedIndex' immediately BEFORE the original array
//	      element previously located at that array index.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'textLine' as a pointer to the Text Line (&textLine).
//
//	     If the 'textLine' parameter is found to be invalid, an
//	     error will be returned.
//
//
//	zeroBasedIndex             int
//	   - This index number designates the array element index in
//	     the Text Lines Collection of the current
//	     TextLineSpecLinesCollection instance where the Text Line
//	     parameter, 'textLine' will be inserted. After insertion,
//	     the 'textLine' object will be positioned immediately
//	     BEFORE the original array element previously located at
//	     that array index.
//
//	     If 'zeroBasedIndex' is set to '4', the original Text Line
//	     object at index '4' will be moved to index position '5'
//	     after the insertion operation is completed.
//
//	     If the value of 'zeroBasedIndex' is less than zero, it
//	     will be reset to zero. This means that the 'textLine'
//	     object will be inserted in the first array element
//	     position of the Text Fields Collection maintained by the
//	     current TextLineSpecLinesCollection instance.
//
//	     If the value of 'zeroBasedIndex' is greater the last array
//	     element index in the Text Fields Collection, the
//	     'textLine' object will be appended to the end of the Text
//	     Lines Collection.
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
// -----------------------------------------------------------------
//
// Return Values
//
//	lastIndexId                int
//	   - If this method completes successfully, the internal array
//	     index of the last text line object for the Text Lines
//	     Collection maintained by the current
//	     TextLineSpecLinesCollection instance will be returned as
//	     an integer value. Remember, this is a zero based index
//	     value which is always one less than the length of the Text
//	     Line Collection.
//
//	     In the event of an error, 'lastIndexId' will be set to a
//	     value of minus one (-1).
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesCol *TextLineSpecLinesCollection) InsertTextLine(
	textLine ITextLineSpecification,
	zeroBasedIndex int,
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

	lastIndexId,
		err = new(textLineSpecLinesCollectionAtom).
		insertTextLine(
			txtLinesCol,
			textLine,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesCol[%v]<-textLine",
					zeroBasedIndex)))

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
// # BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	isValid             bool
//	   - If all data elements encapsulated by the current instance
//	     of TextLineSpecLinesCollection are valid, this returned
//	     boolean value is set to 'true'. If any data values are
//	     invalid, this return parameter is set to 'false'.
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
// # BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix         interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	error
//	   - If any of the internal member data variables contained in
//	     the current instance of TextLineSpecLinesCollection are found
//	     to be invalid, this method will return an error.
//
//	     Also, if the current instance of
//	     TextLineSpecLinesCollection contains zero Text Line members
//	     in the collection, this method will return an error.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
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

// New - Returns a new, empty instance of
// TextLineSpecLinesCollection.
//
// The Text Line Collection for this returned instance is empty and
// contains zero member elements.
//
// To add Text Lines to the collection encapsulated by this
// instance of TextLineSpecLinesCollection, call the method
//
//	TextLineSpecLinesCollection.AddTextLine()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecLinesCollection
//	   - This method will return an empty or uninitialized instance
//	     of TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this instance contains zero member
//	     elements.
func (txtLinesCol TextLineSpecLinesCollection) New() TextLineSpecLinesCollection {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	newTxtLineSpec := TextLineSpecLinesCollection{}

	return newTxtLineSpec
}

// NewTextLine - Returns a new instance of
// TextLineSpecLinesCollection. The Text Lines Collection of this
// new TextLineSpecLinesCollection instance will be populated with
// a single Text Line object passed as input parameter, 'textLine'
//
// To add more Text Lines to the Text Lines Collection maintained
// by the returned instance of TextLineSpecLinesCollection, call
// the method
//
//	TextLineSpecLinesCollection.AddTextLine()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by the returned instance of
//	     TextLineSpecLinesCollection.
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
//	newTextLineCol             TextLineSpecLinesCollection
//	   - This method will return a new instance of
//	     TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this returned instance will contain a
//	     single Text Line object passed as input parameter,
//	     'textLine'.
//
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
func (txtLinesCol TextLineSpecLinesCollection) NewTextLine(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	newTextLineCol TextLineSpecLinesCollection,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newTextLineCol = TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.NewTextLine()",
		"")

	if err != nil {
		return newTextLineCol, err
	}

	err = textLineSpecLinesCollectionNanobot{}.ptr().
		addTextLine(
			&newTextLineCol,
			textLine,
			ePrefix.XCpy(
				"newTextLineCol<-textLine"))

	return newTextLineCol, err
}

// NewPtr - Returns a pointer to a new, empty instance of
// TextLineSpecLinesCollection.
//
// The Text Line Collection for this returned instance is empty and
// contains zero member elements.
//
// To add Text Lines to the collection encapsulated by this
// instance of TextLineSpecLinesCollection, call the method
//
//	TextLineSpecLinesCollection.AddTextLine()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecLinesCollection
//	   - This method will return a pointer to an empty or
//	     uninitialized instance of TextLineSpecLinesCollection. The
//	     Text Line Collection encapsulated by this instance
//	     contains zero member elements.
func (txtLinesCol TextLineSpecLinesCollection) NewPtr() *TextLineSpecLinesCollection {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	newTxtLineSpec := TextLineSpecLinesCollection{}

	return &newTxtLineSpec

}

// NewPtrTextLine - Returns a pointer to a new instance of
// TextLineSpecLinesCollection. The Text Lines Collection of this
// new TextLineSpecLinesCollection instance will be populated with
// a single Text Line object passed as input parameter, 'textLine'
//
// To add more Text Lines to the Text Lines Collection maintained
// by the returned instance of TextLineSpecLinesCollection, call
// the method
//
//	TextLineSpecLinesCollection.AddTextLine()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by the returned instance of
//	     TextLineSpecLinesCollection.
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
//	newTextLineCol             *TextLineSpecLinesCollection
//	   - This method will return a pointer to a new instance of
//	     TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this returned instance will contain a
//	     single Text Line object passed as input parameter,
//	     'textLine'.
//
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
func (txtLinesCol TextLineSpecLinesCollection) NewPtrTextLine(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	newTextLineCol *TextLineSpecLinesCollection,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newTextLineCol = &TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.NewPtrTextLine()",
		"")

	if err != nil {
		return newTextLineCol, err
	}

	err = textLineSpecLinesCollectionNanobot{}.ptr().
		addTextLine(
			newTextLineCol,
			textLine,
			ePrefix.XCpy(
				"newTextLineCol<-textLine"))

	return newTextLineCol, err
}

// PeekAtFirstTextLine - Returns a deep copy of the first Text Line
// ('ITextLineSpecification') object in the Text Lines Collection
// ('txtLinesCol.textLines[0]').
//
// As a 'Peek' method, the original Text Line object
// ('txtLinesCol.textLines[0]') WILL NOT be deleted from the Text
// Lines Collection encapsulated by the current instance of
// TextLineSpecLinesCollection.
//
// After completion of this method, the Text Line Collection array
// will remain unchanged.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	iTxtLineSpec              ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     the designated member of the Text Lines Collection
//	     will be returned to the calling function. The returned
//	     object will implement the ITextLineSpecification
//	     interface.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	---------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtFirstTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PeekAtFirstTextLine(
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
			"PeekAtFirstTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			0,
			false,
			ePrefix.XCpy(
				"txtLinesCol[0]"))

	return iTextLineSpec, err
}

// PeekAtLastTextLine - Returns a deep copy of the last Text Line
// ('ITextLineSpecification') object in the Text Lines Collection
// ('txtLinesCol.textLines[lastIdx]').
//
// As a 'Peek' method, the original Text Line object
// ('txtLinesCol.textLines[lastIdx]') WILL NOT be deleted from the Text
// Line Collection encapsulated by this instance of
// TextLineSpecStandardLine.
//
// After completion of this method, the Text Line Collection array
// will remain unchanged.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//	iTxtLineSpec              ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     the last member element of the Text Lines Collection
//	     will be returned to the calling function. The returned
//	     object will implement the ITextLineSpecification
//	     interface.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
// /
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtLastTextLine(
//	         ePrefix.XCpy(
//	         "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PeekAtLastTextLine(
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
			"PeekAtLastTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	lastIdx := len(txtLinesCol.textLines) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Lines Collection is empty!\n",
			ePrefix.String())

		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			lastIdx,
			false,
			ePrefix.XCpy(
				fmt.Sprintf("txtLinesCol[%v]",
					lastIdx)))

	return iTextLineSpec, err
}

// PeekAtTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the index of the last member element in the collection, an
// error will be returned.
//
// -----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( PeekAtTextField() ) is functionally equivalent to
// method:
//
//	TextLineSpecLinesCollection.GetTextField()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Text Lines collection is empty, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtTextLine(
//	           2,  // Return Text Line at index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PeekAtTextLine(
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
			"PeekAtTextLine()",
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

// PopFirstTextLine - Returns a deep copy of the first Text Line
// ('ITextLineSpecification') object in the Text Line Collection
// ('txtLinesCol.textLines[0]').
//
// As a 'Pop' method, the first Text Line object will be deleted
// from the Text Line Collection encapsulated by this instance of
// TextLineSpecStandardLine. Parameter 'remainingNumOfTextLines'
// will be returned to the calling function containing the number
// of array elements still remaining in the Text Line Collection
// after deletion of the first array element.
//
// The number of array elements remaining in the Text Line
// Collection after this operation will always be one less than
// the original number of array elements.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the first member of
// the Text Line Collection will be DELETED and the Text Line
// Collection array will have a length which is one less than the
// original array length.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//	iTxtLineSpec               ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     first member of the Text Lines Collection will be returned
//	     to the calling function. The returned object will
//	     implement the ITextLineSpecification interface.
//
//	     After completion, the first element of Text Lines
//	     Collection will be deleted.
//
//
//	remainingNumOfTextLines    int
//	   - If this method completes successfully, the first array
//	     element in the Text Lines Collection will be deleted.
//	     After deleting that element, this parameter will return
//	     the number of array elements still remaining in the
//	     Text Lines Collection.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopFirstTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE FIRST
//	   // MEMBER OF THE TEXT LINES COLLECTION!!!
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PopFirstTextLine(
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	remainingNumOfTextLines = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopFirstTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			0,
			true,
			ePrefix.XCpy(
				"txtLinesCol[0]"))

	remainingNumOfTextLines = len(txtLinesCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// PopLastTextLine - Returns a deep copy of the last Text Line
// ('ITextLineSpecification') element in the Text Line Collection
// ('txtLinesCol.textLines[lastIndex]').
//
// As a 'Pop' method, the last Text Line object will be deleted
// from the Text Line Collection encapsulated by this instance of
// TextLineSpecStandardLine. Parameter 'remainingNumOfTextLines'
// will be returned to the calling function containing the number
// of array elements still remaining in the Text Line Collection
// after deletion of the last array element.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the last member of
// the Text Line Collection will be DELETED and the Text Line
// Collection array will have a length which is one less than the
// original array length.
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//	iTxtLineSpec               ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     if the last member of the Text Lines Collection will be
//	     returned to the calling function. The returned object
//	     will implement the ITextLineSpecification interface.
//
//	     After completion, the last element of Text Lines
//	     Collection will be deleted.
//
//
//	remainingNumOfTextLines    int
//	   - If this method completes successfully, the last array
//	     element in the Text Lines Collection will be deleted.
//	     After deleting that element, this parameter will return
//	     the number of array elements still remaining in the
//	     Text Lines Collection.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopLastTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE LAST
//	   // MEMBER OF THE TEXT LINES COLLECTION!!!
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PopLastTextLine(
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	remainingNumOfTextLines = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopLastTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	lastIdx := len(txtLinesCol.textLines) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Lines Collection is empty!\n",
			ePrefix.String())

		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			lastIdx,
			true,
			ePrefix.XCpy(
				"txtLinesCol[0]"))

	remainingNumOfTextLines = len(txtLinesCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// PopTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is always
// index zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in the collection, an error will be
// returned.
//
// -----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( PopTextField() ) is similar to the following
// methods:
//
//	TextLineSpecLinesCollection.GetTextField()
//	TextLineSpecLinesCollection.PeekAtTextField()
//
// The sole difference between this method and the two methods
// cited above is that this method deletes the target collection
// member element.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the target member of
// the Text Line Collection specified by input parameter
// 'zeroBasedIndex' will be DELETED.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. After returning a deep copy of the
//	     collection element specified by 'zeroBasedIndex',
//	     that element will be deleted from the collection.
//
//	     If this input parameter is found to be invalid or if the
//	     Text Lines collection is empty, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned. In addition, the
//	     original collection element specified by 'zeroBasedIndex'
//	     will be deleted.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopTextLine(
//	           2,  // Return a copy of and delete Text Line index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE TARGET
//	   // MEMBER OF THE TEXT LINES COLLECTION DESIGNATED
//	   // BY 'zeroBasedIndex'
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesCol *TextLineSpecLinesCollection) PopTextLine(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesCol.lock == nil {
		txtLinesCol.lock = new(sync.Mutex)
	}

	txtLinesCol.lock.Lock()

	defer txtLinesCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = textLineSpecLinesCollectionAtom{}.ptr().
		peekPopTextLine(
			txtLinesCol,
			zeroBasedIndex,
			true,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesCol[%v]",
					zeroBasedIndex)))

	remainingNumOfTextLines = len(txtLinesCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// ReplaceTextLine - Receives an object which implements the
// ITextLineSpecification interface. This object will replace an
// existing text line object within the text line collection
// maintained by this TextLineSpecLinesCollection instance.
//
// The text line object to be replaced must exist at the index
// specified by input parameter, 'replaceAtIndex'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will replace an existing element within the
//	     text lines collection maintained by this instance of
//	     TextLineSpecLinesCollection. The text line object to
//	     be replaced is identified by the collection element index
//	     supplied by input parameter 'replaceAtIndex'.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'textLine' as a pointer to the Text Line (&textLine).
//
//
//	replaceAtIndex             int
//	   - The index of an element within the text lines collection
//	     maintained by the current TextLineSpecLinesCollection
//	     instance which will be replaced by input parameter
//	     'textLine'.
//
//	     Remember that the text fields collection maintained by
//	     the current TextLineSpecLinesCollection instance is a zero
//	     based array. Therefore, the first index in the collection
//	     is zero (0).
//
//	     If 'replaceAtIndex' proves to be an invalid index, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
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

// SetTextLineCollection - Deletes the Text Line Collection for the
// current instance of TextLineSpecLinesCollection. This collection
// is then replaced with the new Text Line Collection passed as
// input parameter, 'newTextLineCol'.
//
// Input parameter 'newTextLineCol' is an array of
// ITextLineSpecification objects.
//
// Only deep copies of the member elements of 'newTextLineCol' will
// be transferred to the new Text Line Collection for the current
// instance of TextLineSpecLinesCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will DELETE the current Text Line Collection and
// replace it with a new collection passed as an input parameter.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	newTextLineCol             []ITextLineSpecification
//	   - An array of ITextLineSpecification objects which will
//	     replace the Text Line Collection for the current instance
//	     of TextLineSpecLinesCollection.
//
//	     The replacement operation will first create deep copies
//	     of 'newTextLineCol' member elements before adding those
//	     deep copies to the TextLineSpecLinesCollection Text Line
//	     Collection.
//
//	     If any of the member elements of this array are found to
//	     be invalid, or if this parameter is passed as a zero
//	     length array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
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
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesCol *TextLineSpecLinesCollection) SetTextLineCollection(
	newTextLineCol []ITextLineSpecification,
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
			"SetTextLineCollection()",
		"")

	if err != nil {
		return err
	}

	lenOfNewTxtLines := len(newTextLineCol)

	if lenOfNewTxtLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newTextLineCol' is invalid!\n"+
			"'newTextLineCol' contains zero text line objects!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	for i := 0; i < lenOfNewTxtLines; i++ {

		if newTextLineCol[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter element newTextLineCol[%v] is invalid!\n"+
				"newTextLineCol[%v] has a 'nil' value.\n",
				ePrefix.String(),
				i,
				i)

			return err
		}

		err2 = newTextLineCol[i].IsValidInstanceError(
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter element newTextLineCol[%v] is invalid!\n"+
				"newTextLineCol[%v] produced the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				i,
				i,
				err2.Error())

			return err
		}
	}

	textLineColAtom := textLineSpecLinesCollectionAtom{}

	textLineColAtom.emptyCollection(
		txtLinesCol)

	txtLinesCol.textLines = make([]ITextLineSpecification, lenOfNewTxtLines)

	for j := 0; j < lenOfNewTxtLines; j++ {

		txtLinesCol.textLines[j],
			err2 = newTextLineCol[j].CopyOutITextLine(
			nil)

		if err2 != nil {

			textLineColAtom.emptyCollection(
				txtLinesCol)

			err = fmt.Errorf("%v\n"+
				"Error: newTextLineCol[%v] Deep Copy Failed!\n"+
				"The Text Line Collection for the Current Instance\n"+
				"of TextLineSpecLinesCollection is now deleted and empty.\n"+
				"The copy operation for newTextLineCol[%v] produced the\n"+
				"following error:\n"+
				"%v\n",
				ePrefix.String(),
				j,
				j,
				err2.Error())

			return err
		}

	}

	return err
}
