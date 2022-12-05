package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

type TextLineSpecTitleMarquee struct {
	standardTitleLeftMargin  string
	standardTitleRightMargin string
	standardMaxLineLen       int
	standardTextFieldLen     int
	leadingBlankLines        TextLineSpecBlankLines
	leadingSolidLines        TextLineSpecSolidLine
	topTitleBlankLines       TextLineSpecBlankLines
	titleLines               TextLineSpecLinesCollection
	bottomTitleBlankLines    TextLineSpecBlankLines
	trailingSolidLines       TextLineSpecSolidLine
	trailingBlankLines       TextLineSpecBlankLines
	textLineReader           *strings.Reader

	lock *sync.Mutex
}

// CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumberStrKernel ('incomingNumStrKernel') to the	data
//	fields of the current NumberStrKernel instance
//	('numStrKernel').
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// All the data fields in current NumberStrKernel instance
// ('numStrKernel') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTitleMarquee		*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. This method will NOT
//		change the values of internal member data
//		variables contained in this instance.
//
//		All data values in this TextLineSpecTitleMarquee
//		instance will be copied to the current
//		TextLineSpecTitleMarquee instance
//		('txtLineSpecTitleMarquee').
//
//		If parameter 'incomingTitleMarquee' is determined
//		to be invalid, an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyIn(
	incomingTitleMarquee *TextLineSpecTitleMarquee,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLineSpecTitleMarqueeNanobot).
		copy(txtLineSpecTitleMarquee,
			incomingTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee<-"+
					"incomingTitleMarquee"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	TextLineSpecTitleMarquee instance.
//
//	If the current TextLineSpecTitleMarquee instance
//	contains invalid member variables, this method will
//	return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a deep
//		copy of the current TextLineSpecTitleMarquee
//		instance will be returned through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyTitleMarquee, err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return deepCopyTitleMarquee, err
}

//	CopyOutITextLine
//
//	Returns a deep copy of the current
//	TextLineSpecTitleMarquee instance cast as a type
//	ITextLineSpecification.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ITextLineSpecification
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextLineSpecTitleMarquee instance cast as an
//		ITextLineSpecification object.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOutITextLine()",
		"")

	if err != nil {
		return ITextLineSpecification(&deepCopyTitleMarquee), err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return ITextLineSpecification(&deepCopyTitleMarquee), err
}

//	CopyOutPtr
//
//	Returns a pointer to a deep copy of the current
//	TextLineSpecTitleMarquee instance.
//
//	If the current TextLineSpecTitleMarquee instance
//	contains invalid member variables, this method will
//	return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*TextLineSpecTitleMarquee
//
//		If this method completes successfully, a pointer
//		to a deep copy of the current
//		TextLineSpecTitleMarquee instance will be
//		returned through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOutPtr()",
		"")

	if err != nil {
		return &deepCopyTitleMarquee, err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return &deepCopyTitleMarquee, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of TextLineSpecTitleMarquee to their initial
//	or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of TextLineSpecTitleMarquee.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Empty() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	new(textLineSpecTitleMarqueeElectron).
		empty(txtLineSpecTitleMarquee)

	txtLineSpecTitleMarquee.lock.Unlock()

	txtLineSpecTitleMarquee.lock = nil
}

//	Equal
//
//	Receives a pointer to another instance of
//	TextLineSpecTitleMarquee and proceeds to compare the
//	member variables to those of the current
//	TextLineSpecTitleMarquee instance in order to
//	determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables of both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTitleMarquee	*TextLineSpecTitleMarquee
//
//		A pointer to an incoming instance of
//		TextLineSpecTitleMarquee. This method will
//		compare all member variable data values in this
//		instance against those contained in the current
//		instance of TextLineSpecTitleMarquee. If the data
//		values in both instances are found to be equal in
//		all respects, this method will return a boolean
//		value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingTitleMarquee' are equal
//		in all respects to those contained in the current
//		instance of TextLineSpecTitleMarquee, this method
//		will return a boolean value of 'true'. Otherwise,
//		a value of 'false' will be returned.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Equal(
	incomingTitleMarquee *TextLineSpecTitleMarquee) bool {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	return new(textLineSpecTitleMarqueeElectron).
		equal(
			txtLineSpecTitleMarquee,
			incomingTitleMarquee)
}

//	EqualITextLine
//
//	Receives an object implementing the
//	ITextLineSpecification interface and proceeds to
//	compare the member variables to those of the current
//	TextLineSpecTitleMarquee instance in order to
//	determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables from both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
//	This method is required by interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	iTextLine	ITextLineSpecification
//
//		An object implementing the ITextLineSpecification
//		interface. If this object proves to be equal in
//		all respects to the current instance of
//		TextLineSpecTitleMarquee, this method will return
//		'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If input parameter 'iTextLine' is judged to be
//		equal in all respects to the current instance of
//		TextLineSpecTitleMarquee, this return parameter
//		will be set to 'true'.
//
//		If 'iTextLine' is NOT equal to the current
//		instance of TextLineSpecTitleMarquee, this return
//		parameter will be set to 'false'.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	txtTitleMarqueeTwo, ok := iTextLine.(*TextLineSpecTitleMarquee)

	if !ok {
		return false
	}

	return new(textLineSpecTitleMarqueeElectron).
		equal(
			txtLineSpecTitleMarquee,
			txtTitleMarqueeTwo)
}

//	GetFormattedText
//
//	Returns the formatted text generated by this Text
//	Line Specification Title Marquee
//	(TextLineSpecTitleMarquee) for text display, file
//	output, screen display and printing.
//
//	This method is similar to
//	TextLineSpecTitleMarquee.String() with the sole
//	difference being that this method returns an error.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The formatted text lines generated by the current
//		instance of TextLineSpecTitleMarquee.
//
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	strBuilder := strings.Builder{}

	_,
		_,
		err = new(textLineSpecTitleMarqueeMolecule).
		getFormattedText(
			&strBuilder,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee"))

	return strBuilder.String(), err
}

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecTitleMarquee
//	instance to determine if they are valid.
//
//	If all data element evaluate as valid, this method
//	returns 'true'. If any data element is invalid, this
//	method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If all data elements encapsulated by the current
//		instance of TextLineSpecTitleMarquee are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) IsValidInstance() (
	isValid bool) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	isValid,
		_ = new(textLineSpecTitleMarqueeElectron).
		testValidityTitleMarquee(
			txtLineSpecTitleMarquee,
			nil)

	return isValid
}

//	IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecTitleMarquee
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this method
//	will return an error.
//
//	This method fulfills requirements of
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//
//		If the current instance of
//		TextLineSpecTitleMarquee is judged to be valid in
//		all respects, the returned error Type is set
//		equal to 'nil'.
//
//		If input parameter 'TextLineSpecTitleMarquee' is
//		found to be invalid, the returned error Type will
//		encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errorPrefix' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineSpecTitleMarqueeElectron).
		testValidityTitleMarquee(
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee"))

	return err
}

//	Read
//
//	Implements the io.Reader interface for type
//	TextLineSpecTitleMarquee.
//
//	The formatted text line string generated by the
//	current instance of TextLineSpecTitleMarquee will be
//	written to the byte buffer 'p'. If the length of 'p'
//	is less than the length of the formatted text line
//	string, multiple calls to this method will write the
//	remaining unread characters to the byte buffer 'p'.
//
//	Read() supports buffered 'read' operations.
//
//	This method reads up to len(p) bytes into p. It
//	returns the number of bytes read (0 <= n <= len(p))
//	and any error encountered. Even if read returns
//	n < len(p), it may use all of p as scratch space
//	during the call.
//
//	If some data is available but not len(p) bytes,
//	readBytes() conventionally returns what is available
//	instead of waiting for more.
//
//	When this method encounters an error or end-of-file
//	condition after successfully reading n > 0 bytes, it
//	returns the number of bytes read. It may return the
//	(non-nil) error from the same call or return the
//	error (and n == 0) from a subsequent call. An
//	instance of this general case is that a Reader
//	returning a non-zero number of bytes at the end of
//	the input stream may return either err == EOF or
//	err == nil. The next read operation should return 0,
//	EOF.
//
//	Callers should always process the n > 0 bytes returned
//	before considering the error err. Doing so correctly
//	handles I/O errors that happen after reading some bytes
//	and also both of the allowed EOF behaviors.
//
//	The last read operation performed on the formatted text
//	string will always return n==0 and err==io.EOF.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p							[]byte
//
//		The byte buffer into which the formatted text
//		line string generated by the current
//		TextLineSpecTitleMarquee instance will be
//		written.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	n							int
//
//		The number of bytes written to byte buffer 'p'.
//
//		Read() reads up to len(p) bytes into p. It
//		returns the number of bytes read
//		(0 <= n <= len(p)) and any error encountered.
//		Even if Read() returns n < len(p), it may use all
//		of 'p' as scratch space during the call. If some
//		data is available but not len(p) bytes, Read()
//		conventionally returns what is available instead
//		of waiting for more.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		When Read() encounters an error or end-of-file
//		condition after successfully reading n > 0 bytes,
//		it returns the number of bytes read. It may
//		return the (non-nil) error from the same call or
//		return the error (and n == 0) from a subsequent
//		call. An instance of this general case is that a
//		Reader returning a non-zero number of bytes at
//		the end of the input stream may return either
//		err == EOF or err == nil. The next read operation
//		should return 0, EOF.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example # 1
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	sb := strings.Builder{}
//
//	for {
//
//	  n,
//	  err = txtTitleMarquee.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  sb.Write(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtTitleMarquee.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %s\n",
//	              sb.String())
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
//
//	Example # 2
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	var actualStr string
//
//	for {
//
//	  n,
//	  err = txtTitleMarquee.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  actualStr += string(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtTitleMarquee.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %v\n",
//	              actualStr)
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Read(
	p []byte) (
	n int,
	err error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecTitleMarquee.Read()",
		"")

	if txtLineSpecTitleMarquee.textLineReader == nil {

		strBuilder := strings.Builder{}

		_,
			_,
			err = new(textLineSpecTitleMarqueeMolecule).
			getFormattedText(
				&strBuilder,
				txtLineSpecTitleMarquee,
				ePrefix.XCpy("txtLineSpecTitleMarquee"))

		if err != nil {
			return n, err
		}

		txtLineSpecTitleMarquee.textLineReader =
			strings.NewReader(strBuilder.String())

		if txtLineSpecTitleMarquee.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtLineSpecTitleMarquee.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}

		strBuilder.Reset()
	}

	n,
		err = new(textLineSpecTitleMarqueeMolecule).
		readBytes(
			txtLineSpecTitleMarquee.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> "+
					"txtLineSpecTitleMarquee.textLineReader"))

	if err == io.EOF {

		txtLineSpecTitleMarquee.textLineReader = nil

	}

	return n, err
}

//	ReaderInitialize
//
//	This method will reset the internal member variable
//	'TextLineSpecTitleMarquee.textLineReader' to its
//	initial zero state of 'nil'. Effectively, this resets
//	the internal strings.Reader object for use in future
//	read operations.
//
//	This method is rarely used or needed. It provides a
//	means of reinitializing the internal strings.Reader
//	object in case an error occurs during a read
//	operation initiated by method
//	TextLineSpecTitleMarquee.Read().
//
//	Calling this method cleans up the residue from an
//	aborted read operation and prepares the
//	strings.Reader object for future read operations.
//
//	If any errors are returned by method
//	TextLineSpecTitleMarquee.Read() which are NOT equal
//	to io.EOF, call this method,
//	TextLineSpecTitleMarquee.ReaderInitialize(), to reset
//	and prepare the internal reader for future read
//	operations.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) ReaderInitialize() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	txtLineSpecTitleMarquee.textLineReader = nil

	return
}

//	NewAllParams
//
//	Creates and returns a new instance of
//	TextLineSpecTitleMarquee.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	standardTitleLeftMargin		string
//
//		The standard left margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardTitleRightMargin		string
//
//		The standard right margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardMaxLineLen					int
//
//		The maximum number of characters allowed on
//		a text title line. This maximum limit will be
//		applied to the length of all text lines generated
//		by the returned instance of
//		TextLineSpecTitleMarquee.
//
//	standardTextFieldLen		int
//
//		The standard field length applied to Text
//		Title Lines in the 'TitleLines' array unless
//		overridden by user customizations.
//
//		If the standardTextFieldLen exceeds the value of
//		the Maximum Available Text Field Length, it will
//		be reset and defaulted to the Maximum Available
//		Text Field Length.
//
//		The Maximum Available Text Field Length is
//		calculated as follows:
//
//		Maximum Available Text Field Length =
//			TextLineTitleMarqueeDto.StandardMaxLineLen -
//			1 -
//			len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//			len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//	numLeadingBlankLines		int
//
//		The number of blank lines or 'new lines'
//		inserted above the Leading Solid Line.
//
//	leadingSolidLineChar		string
//
//		The character used to create the Leading
//		Solid Line displayed above the Title
//		Lines.
//
//	numLeadingSolidLines		int
//
//		The Number of Leading Solid Lines to display
//		above the Title Lines.
//
//	numTopTitleBlankLines		int
//
//		The number of blank lines or 'new lines' to
//		insert immediately above the Title Lines
//		Display.
//
//	titleLines					TextLineSpecLinesCollection
//
//		A collection of text line objects containing all
//		specifications necessary to display the Text
//		Title Lines.
//
//		If this parameter is empty with zero Text Line
//		member elements, no error will be returned.
//		However, the user will be responsible for
//		populating the title lines using the 'Add'
//		methods on the returned instance of
//		TextLineSpecTitleMarquee.
//
//	numBottomTitleBlankLines	int
//
//		The number of blank lines or 'new lines' to
//		insert immediately below the Title Lines Display.
//
//	trailingSolidLineChar		string
//
//		The character used to create the Trailing Solid
//		Line displayed below the Title Lines.
//
//	numTrailingSolidLines		int
//
//		The Number of Trailing Solid Lines to display
//		below the Title Lines.
//
//	numTrailingBlankLines		int
//
//		The number of blank lines or 'new lines' inserted
//		after the Trailing Solid Line.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a new
//		instance of TextLineSpecTitleMarquee will be
//		returned.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) NewAllParams(
	standardTitleLeftMargin string,
	standardTitleRightMargin string,
	standardMaxLineLen int,
	standardTextFieldLen int,
	numLeadingBlankLines int,
	leadingSolidLineChar string,
	numLeadingSolidLines int,
	numTopTitleBlankLines int,
	titleLines TextLineSpecLinesCollection,
	numBottomTitleBlankLines int,
	trailingSolidLineChar string,
	numTrailingSolidLines int,
	numTrailingBlankLines int,
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var newTxtLineTitle TextLineSpecTitleMarquee
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"NewAllParams()",
		"")

	if err != nil {
		return newTxtLineTitle, err
	}

	titleMarqueeDto := TextLineTitleMarqueeDto{
		StandardTitleLeftMargin:  standardTitleLeftMargin,
		StandardTitleRightMargin: standardTitleRightMargin,
		StandardMaxLineLen:       standardMaxLineLen,
		StandardTextFieldLen:     standardTextFieldLen,
		NumLeadingBlankLines:     numLeadingBlankLines,
		LeadingSolidLineChar:     leadingSolidLineChar,
		NumLeadingSolidLines:     numLeadingSolidLines,
		NumTopTitleBlankLines:    numTopTitleBlankLines,
		NumBottomTitleBlankLines: numBottomTitleBlankLines,
		TrailingSolidLineChar:    trailingSolidLineChar,
		NumTrailingSolidLines:    numTrailingSolidLines,
		NumTrailingBlankLines:    numTrailingBlankLines,
	}

	if titleLines.GetNumberOfTextLines() == 0 {

		titleMarqueeDto.TitleLines.Empty()

	} else {

		err = titleMarqueeDto.TitleLines.CopyIn(
			&titleLines,
			ePrefix.XCpy("<-titleLines"))

		if err != nil {
			return newTxtLineTitle, err
		}

	}

	err = new(textLineSpecTitleMarqueeMechanics).
		setTxtLineTitleMarquee(
			&newTxtLineTitle,
			titleMarqueeDto,
			ePrefix.XCpy(
				"newTxtLineTitle"))

	return newTxtLineTitle, err
}
