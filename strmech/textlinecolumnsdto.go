package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextLineColumnsDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. This should value should
	// always be set to:
	//   TxtFieldType.LineColumns()

	TextFieldsContent []TextFieldsContentDto

	FmtParameters TextFmtParamsLineColumnsDto

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineColumnsDto ('incomingTxtLineCols') to the data fields
// of the current TextLineColumnsDto instance
// ('fmtLineCols').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextLineColumnsDto instance
// ('fmtLineCols') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtLineCols'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtLineCols        TextLineColumnsDto
//     - An instance of TextLineColumnsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextLineColumnsDto instance
//       ('incomingTxtLineCols') will be copied to the current
//       TextLineColumnsDto instance ('fmtLineCols').
//
//       No data validation is performed on input parameter,
//       'incomingTxtLineCols'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (fmtLineCols *TextLineColumnsDto) CopyIn(
	incomingTxtLineCols TextLineColumnsDto) {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	_ = new(textLineColumnsDtoNanobot).
		copy(
			fmtLineCols,
			&incomingTxtLineCols,
			nil)

	return
}

// CopyOut - Returns a deep copy of the current TextLineColumnsDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextLineColumnsDto.
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
//  deepCopyTxtLineColsDto     TextLineColumnsDto
//     - This parameter will return a deep copy of the current
//       TextLineColumnsDto instance.
//
func (fmtLineCols *TextLineColumnsDto) CopyOut() (
	deepCopyTxtLineColsDto TextLineColumnsDto) {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	_ = new(textLineColumnsDtoNanobot).
		copy(
			&deepCopyTxtLineColsDto,
			fmtLineCols,
			nil)

	return deepCopyTxtLineColsDto
}

// Empty - Resets all internal member variables for the current
// instance of TextLineColumnsDto to their zero or
// uninitialized states. This method will leave the current
// instance of TextLineColumnsDto in an invalid state and
// unavailable for immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextLineColumnsDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
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
func (fmtLineCols *TextLineColumnsDto) Empty() {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	new(textLineColumnsDtoMolecule).
		empty(
			fmtLineCols)

	fmtLineCols.lock.Unlock()

	fmtLineCols.lock = nil
}

// Equal - Receives another instance of TextLineColumnsDto and
// proceeds to compare the member variables to those of the current
// TextLineColumnsDto instance in order to determine if they are
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
//  incomingTxtLineCols        TextLineColumnsDto
//     - An incoming instance of TextLineColumnsDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextLineColumnsDto. If the data values in both
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
//       parameter 'incomingTxtLineCols' are equal in all respects
//       to those contained in the current instance of
//       TextLineColumnsDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (fmtLineCols *TextLineColumnsDto) Equal(
	incomingTxtLineCols TextLineColumnsDto) bool {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	return new(textLineColumnsDtoMolecule).
		equal(
			fmtLineCols,
			&incomingTxtLineCols)
}

// GetNumberOfFieldFormatParams - Returns the number of Field
// Format Parameters configured for the current instance of
// TextLineColumnsDto.
//
// The number of Field Format Parameters should always match the
// number of Text Fields configured for the current instance of
// TextLineColumnsDto.
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
//  int
//     - This method returns an integer value specifying the number
//       of Field Format Parameters array maintained by this
//       instance of TextLineColumnsDto:
//
//       TextFmtParamsLineColumnsDto.FmtParameters.
//           FieldFormatParams
//
func (fmtLineCols *TextLineColumnsDto) GetNumberOfFieldFormatParams() int {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	return fmtLineCols.FmtParameters.GetNumOfFieldFmtParams()
}

// GetNumberOfTextFields - Returns the number of Text Field Content
// Data Transfer objects residing in the Text Field Content array
// maintained by this instance of TextLineColumnsDto.
//
// Effectively, this is the number of Text Fields currently
// configured, in this instance of TextLineColumnsDto.
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
//  int
//     - This returned integer value specifies the number of Text
//       Fields currently configured in this instance of
//       TextLineColumnsDto:
//          TextLineColumnsDto.TextFieldsContent
//
func (fmtLineCols *TextLineColumnsDto) GetNumberOfTextFields() int {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	return len(fmtLineCols.TextFieldsContent)
}

// GetTextFieldType - Returns the internal member variable which
// stores the value of Text Format Type. For type
// TextLineColumnsDto, this value should be set to
// TxtFieldType.Line1Column().
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
//  TextFieldType
//     - This method returns an enumeration value specifying the
//       Text Field Type associated with this instance of
//       TextLineColumnsDto:
//          TextLineColumnsDto.FormatType
//
func (fmtLineCols *TextLineColumnsDto) GetTextFieldType() TextFieldType {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	return fmtLineCols.FormatType
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineColumnsDto instance
// to determine if they are valid.
//
// If all data elements evaluate as classified as valid, this
// method returns a boolean value of 'true'. If any data element
// is invalid, this method returns 'false'.
//
// Comparatively little validation is performed. Primarily, this
// method checks to ensure that the number Text Fields matches
// the number Field Format Paramter objects configured for this
// instance of TextLineColumnsDto.
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
//       of TextLineSpecStandardLine are valid, this returned
//       boolean value is set to 'true'. If any data values are
//       invalid, this return parameter is set to 'false'.
//
func (fmtLineCols *TextLineColumnsDto) IsValidInstance() (
	isValid bool) {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	isValid = false

	lenTxtFields := len(fmtLineCols.TextFieldsContent)

	lenFmtParams := len(fmtLineCols.FmtParameters.FieldFormatParams)

	if lenTxtFields !=
		lenFmtParams {

		return isValid
	}

	isValid = true

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineColumnsDto instance
// to determine if they are valid.
//
// If any data elements evaluates as invalid, this method will
// return an error.
//
// Comparatively little validation is performed. Primarily, this
// method checks to ensure that the number Text Fields matches
// the number Field Format Paramter objects configured for this
// instance of TextLineColumnsDto.
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
//     - If any of the internal member data variables contained in
//       the current instance of TextLineColumnsDto are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (fmtLineCols *TextLineColumnsDto) IsValidInstanceError(
	errorPrefix interface{}) error {

	if fmtLineCols.lock == nil {
		fmtLineCols.lock = new(sync.Mutex)
	}

	fmtLineCols.lock.Lock()

	defer fmtLineCols.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	lenTxtFields := len(fmtLineCols.TextFieldsContent)

	lenFmtParams := len(fmtLineCols.FmtParameters.FieldFormatParams)

	if lenTxtFields !=
		lenFmtParams {

		err = fmt.Errorf("%v\n"+
			"Error: The number of Text Fields DOES NOT MATCH\n"+
			"the number of Field Format Parameters configured\n"+
			"for this instance of TextLineColumnsDto.\n"+
			"Number of Text Fields = '%v'\n"+
			"Number of Field Format Parameters = '%v'\n",
			ePrefix.String(),
			lenTxtFields,
			lenFmtParams)
	}

	return err
}

// textLineColumnsDtoNanobot - Provides helper methods for
// TextLineColumnsDto.
type textLineColumnsDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextLineColumnsDto to a destination instance of
// TextLineColumnsDto.
func (lineColsDtoNanobot *textLineColumnsDtoNanobot) copy(
	destinationTxtLineColsDto *TextLineColumnsDto,
	sourceTxtLineColsDto *TextLineColumnsDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if lineColsDtoNanobot.lock == nil {
		lineColsDtoNanobot.lock = new(sync.Mutex)
	}

	lineColsDtoNanobot.lock.Lock()

	defer lineColsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineColumnsDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtLineColsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtLineColsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtLineColsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtLineColsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textLineColumnsDtoMolecule).
		empty(destinationTxtLineColsDto)

	destinationTxtLineColsDto.FormatType =
		sourceTxtLineColsDto.FormatType

	lenItems := len(sourceTxtLineColsDto.TextFieldsContent)

	if lenItems > 0 {

		destinationTxtLineColsDto.TextFieldsContent =
			make([]TextFieldsContentDto, lenItems)

		for i := 0; i < lenItems; i++ {

			destinationTxtLineColsDto.TextFieldsContent[i].CopyIn(
				sourceTxtLineColsDto.TextFieldsContent[i])

		}
	}

	destinationTxtLineColsDto.FmtParameters.CopyIn(
		sourceTxtLineColsDto.FmtParameters)

	return err
}

// textLineColumnsDtoMolecule - Provides helper methods for
// TextLineColumnsDto.
type textLineColumnsDtoMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of TextLineColumnsDto
// and proceeds to set all the internal member variables to their
// zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextLineColumnsDto.
//
func (lineColsDtoMolecule *textLineColumnsDtoMolecule) empty(
	txtLineColsDto *TextLineColumnsDto) {

	if lineColsDtoMolecule.lock == nil {
		lineColsDtoMolecule.lock = new(sync.Mutex)
	}

	lineColsDtoMolecule.lock.Lock()

	defer lineColsDtoMolecule.lock.Unlock()

	if txtLineColsDto == nil {
		return
	}

	txtLineColsDto.FormatType = TxtFieldType.None()

	lenItems := len(txtLineColsDto.TextFieldsContent)

	for i := 0; i < lenItems; i++ {

		txtLineColsDto.TextFieldsContent[i].Empty()
	}

	txtLineColsDto.TextFieldsContent = nil

	txtLineColsDto.FmtParameters.Empty()

}

// equal - Receives pointers to two instances of TextLineColumnsDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextLineColumnsDto are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
//
func (lineColsDtoMolecule *textLineColumnsDtoMolecule) equal(
	txtLineColsDto1 *TextLineColumnsDto,
	txtLineColsDto2 *TextLineColumnsDto) bool {

	if lineColsDtoMolecule.lock == nil {
		lineColsDtoMolecule.lock = new(sync.Mutex)
	}

	lineColsDtoMolecule.lock.Lock()

	defer lineColsDtoMolecule.lock.Unlock()

	if txtLineColsDto1 == nil ||
		txtLineColsDto2 == nil {
		return false
	}

	if txtLineColsDto1.FormatType !=
		txtLineColsDto2.FormatType {

		return false
	}

	lenItems := len(txtLineColsDto1.TextFieldsContent)

	if lenItems != len(txtLineColsDto2.TextFieldsContent) {

		return false
	}

	for i := 0; i < lenItems; i++ {

		if !txtLineColsDto1.TextFieldsContent[i].Equal(
			txtLineColsDto2.TextFieldsContent[i]) {

			return false
		}
	}

	if !txtLineColsDto1.FmtParameters.Equal(
		txtLineColsDto1.FmtParameters) {

		return false
	}

	return true
}
