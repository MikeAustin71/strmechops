package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

// TextFieldDto - The Text Field Data Transfer Object specifies the
// parameters necessary to construct a formatted line of text
// consisting of a left margin, a text field, a right margin and a
// line termination sequence.
//
// Example Layout
//
//         " " + "Inflation Rates" + " " + "\n"
//
type TextFieldDto struct {
	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldDateTime time.Time
	// If this Text Field is designated as a 'DateTime' Field, this
	// time value will be used to populate the Text Field.

	FieldDateTimeFormat string
	// If this Text Field is designated as a 'DateTime' Field, this
	// string will be used to format the Date/Time.
	//
	// If 'FieldDateTime' is set to a value greater than zero and
	// this 'ParamValueDateTimeFormat' string is empty (has a zero
	// length), a default Date/Time format string will be applied
	// as follows:
	//         "2006-01-02 15:04:05.000000000 -0700 MST"

	FieldText string
	// The Text Field string or contents. If this string is empty
	// (has a zero (0) length) and is designated as a Label, Filler
	// or Spacer Text Field, an error will be generated.
	//
	// This string represents the contents of the Text Field.

	FieldLength int
	// Used to format Label Text Fields. This is the length of the
	// text field in which the 'FieldText' will be displayed. If
	// 'FieldLength' is less than the length of the 'FieldText'
	// string, it will be automatically set equal to the
	// 'FieldText' string length.
	//
	// To automatically set the value of 'FieldLength' to the
	// length of 'FieldText', set this parameter to a value of
	// minus one (-1).
	//
	// If this parameter is submitted with a value less than
	// minus one (-1) or greater than 1-million (1,000,000), an
	// error will be returned.
	//
	// If TextFieldType is set to 'Filler' or 'Spacer',
	// 'FieldLength' will be used to specify the number of Filler
	// or Spacer characters in the line.
	//
	// If TextFieldType is set to 'BlankLine', 'FieldLength' will
	// be used to specify the number of blank lines created.

	FieldJustify TextJustify
	// An enumeration which specifies the justification of the
	// 'FieldText' string within the text field specified by
	// 'FieldLength'.
	//
	// Text justification can only be evaluated in the context of
	// a text label, field length and a Text Justification object
	// of type TextJustify. This is because text labels with a
	// field length equal to or less than the length of the text
	// label never use text justification. In these cases, text
	// justification is completely ignored.
	//
	// If the field length is greater than the length of the text
	// label, text justification must be equal to one of these
	// three valid values:
	//     TextJustify(0).Left()
	//     TextJustify(0).Right()
	//     TextJustify(0).Center()
	//
	// You can also use the abbreviated text justification
	// enumeration syntax as follows:
	//
	//     TxtJustify.Left()
	//     TxtJustify.Right()
	//     TxtJustify.Center()

	FieldType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Field Specification which will to configure the final text
	// field output. If this parameter is invalid, an error will be
	// generated.
	//
	// Possible values are listed as follows:
	//   TxtFieldType.None()      - Invalid
	//   TxtFieldType.Label()     - Valid
	//   TxtFieldType.DateTime()  - Valid
	//   TxtFieldType.Filler()    - Valid
	//   TxtFieldType.Spacer()    - Valid
	//   TxtFieldType.BlankLine() - Valid

	RightMarginStr string
	// The contents of the string will be used as the right margin
	// for the Text Field.
	//
	// If no right margin is required, set 'RightMarginStr' to a
	// zero length or empty string, and no right margin will be
	// created.

	LineTerminator string
	// This string holds the character or characters which will be
	// used to terminate the formatted line of text output.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If no Line Terminator is required, set 'lineTerminator' to
	// a zero length or empty string and no line termination
	// characters will be created.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldDto ('incomingTxtFieldDto') to the data fields
// of the current TextFieldDto instance ('txtFieldDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFieldDto instance
// ('txtFieldDto') will be deleted and overwritten.
//
// NO data validation is performed on input parameter
// 'incomingTxtFieldDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFieldDto     *TextFieldDto
//     - A pointer to an instance of TextFieldDto. This
//       method will NOT change the data values of member variables
//       contained in this instance.
//
//       All data values in this TextFieldDto instance
//       will be copied to the current TextFieldDto
//       instance ('txtFieldDto').
//
//       No data validation is performed on input parameter
//       'incomingTxtFieldDto'.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
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
func (txtFieldDto *TextFieldDto) CopyIn(
	incomingTxtFieldDto *TextFieldDto,
	errorPrefix interface{}) error {

	if txtFieldDto.lock == nil {
		txtFieldDto.lock = new(sync.Mutex)
	}

	txtFieldDto.lock.Lock()

	defer txtFieldDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLabelParameterValueFieldDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	if incomingTxtFieldDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtFieldDto' is invalid!\n"+
			"'incomingTxtFieldDto' has a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	txtFieldDto.LeftMarginStr =
		incomingTxtFieldDto.LeftMarginStr

	txtFieldDto.FieldDateTime =
		incomingTxtFieldDto.FieldDateTime

	txtFieldDto.FieldDateTimeFormat =
		incomingTxtFieldDto.FieldDateTimeFormat

	txtFieldDto.FieldText =
		incomingTxtFieldDto.FieldText

	txtFieldDto.FieldLength =
		incomingTxtFieldDto.FieldLength

	txtFieldDto.FieldJustify =
		incomingTxtFieldDto.FieldJustify

	txtFieldDto.FieldType =
		incomingTxtFieldDto.FieldType

	txtFieldDto.RightMarginStr =
		incomingTxtFieldDto.RightMarginStr

	txtFieldDto.LineTerminator =
		incomingTxtFieldDto.LineTerminator

	return err
}

// CopyOut - Returns a deep copy of the current
// TextFieldDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method does NOT perform data validation checks on the
// current instance of TextFieldDto before creating and returning
// the deep copy of this instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  TextFieldDto
//     - This method will return a deep copy of the current
//       TextFieldDto instance.
//
func (txtFieldDto *TextFieldDto) CopyOut() TextFieldDto {

	if txtFieldDto.lock == nil {
		txtFieldDto.lock = new(sync.Mutex)
	}

	txtFieldDto.lock.Lock()

	defer txtFieldDto.lock.Unlock()

	newCopyTextFieldDto := TextFieldDto{}

	newCopyTextFieldDto.LeftMarginStr =
		txtFieldDto.LeftMarginStr

	newCopyTextFieldDto.FieldDateTime =
		txtFieldDto.FieldDateTime

	newCopyTextFieldDto.FieldDateTimeFormat =
		txtFieldDto.FieldDateTimeFormat

	newCopyTextFieldDto.FieldText =
		txtFieldDto.FieldText

	newCopyTextFieldDto.FieldLength =
		txtFieldDto.FieldLength

	newCopyTextFieldDto.FieldJustify =
		txtFieldDto.FieldJustify

	newCopyTextFieldDto.FieldType =
		txtFieldDto.FieldType

	newCopyTextFieldDto.RightMarginStr =
		txtFieldDto.RightMarginStr

	newCopyTextFieldDto.LineTerminator =
		txtFieldDto.LineTerminator

	return newCopyTextFieldDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldDto to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldDto. All member variable data
// values will be reset to their zero or uninitialized states.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtFieldDto *TextFieldDto) Empty() {

	if txtFieldDto.lock == nil {
		txtFieldDto.lock = new(sync.Mutex)
	}

	txtFieldDto.lock.Lock()

	txtFieldDto.LeftMarginStr = ""

	txtFieldDto.FieldDateTime = time.Time{}

	txtFieldDto.FieldDateTimeFormat = ""

	txtFieldDto.FieldText = ""

	txtFieldDto.FieldLength = 0

	txtFieldDto.FieldJustify = TxtJustify.None()

	txtFieldDto.FieldType = TxtFieldType.None()

	txtFieldDto.RightMarginStr = ""

	txtFieldDto.LineTerminator = ""

	txtFieldDto.lock.Unlock()

	txtFieldDto.lock = nil
}

// Equal - Receives a pointer to another instance of TextFieldDto
// and proceeds to compare the member variables to those of the
// current TextFieldDto instance in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFieldDto   *TextFieldDto
//     - A pointer to an incoming instance of TextFieldDto. This
//       method will compare all member variable data values in
//       this instance against those contained in the current
//       instance of TextFieldDto. If the data values in both
//       instances are found to be equal in all respects, this
//       method will return a boolean value of 'true'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the member variable data values contained in input
//       parameter 'incomingTxtFieldDto' are equal in all respects
//       to those contained in the current instance of
//       TextFieldDto, this method will return a boolean value of
//       'true'. Otherwise a value of 'false' will be returned to
//       the calling function.
//
func (txtFieldDto *TextFieldDto) Equal(
	incomingTxtFieldDto *TextFieldDto) bool {

	if txtFieldDto.lock == nil {
		txtFieldDto.lock = new(sync.Mutex)
	}

	txtFieldDto.lock.Lock()

	defer txtFieldDto.lock.Unlock()

	if incomingTxtFieldDto == nil {
		return false
	}

	if txtFieldDto.LeftMarginStr !=
		incomingTxtFieldDto.LeftMarginStr {

		return false
	}

	if txtFieldDto.FieldDateTime !=
		incomingTxtFieldDto.FieldDateTime {

		return false
	}

	if txtFieldDto.FieldDateTimeFormat !=
		incomingTxtFieldDto.FieldDateTimeFormat {

		return false
	}

	if txtFieldDto.FieldText !=
		incomingTxtFieldDto.FieldText {

		return false
	}

	if txtFieldDto.FieldLength !=
		incomingTxtFieldDto.FieldLength {

		return false
	}

	if txtFieldDto.FieldJustify !=
		incomingTxtFieldDto.FieldJustify {

		return false
	}

	if txtFieldDto.FieldType !=
		incomingTxtFieldDto.FieldType {

		return false
	}

	if txtFieldDto.RightMarginStr !=
		incomingTxtFieldDto.RightMarginStr {

		return false
	}

	if txtFieldDto.LineTerminator !=
		incomingTxtFieldDto.LineTerminator {

		return false
	}

	return true
}
