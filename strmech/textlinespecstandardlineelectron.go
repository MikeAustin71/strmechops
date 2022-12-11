package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
}

// addTextFieldDtosToArray
//
// This method will append multiple text fields to the
// end of the internal array of text field objects
// maintained by the instance of TextLineSpecStandardLine
// passed as input parameter 'txtStdLine'.
//
// The text fields will be constructed using text field
// specifications from the 'textFieldDtos' data transfer
// object. Thereafter, the new constructed text fields
// will be appended to the end of the internal text field
// array contained within 'txtStdLine'.
//
// If the method completes successfully, the internal
// array index of the last Text Field object will be
// returned to the calling function.
//
// Remember that the last array index is equal the array
// length minus one (array length - 1).
//
// This method performs low level operations which
// actually constructs the text fields and appends them
// to the end of the internal text field array contained
// in 'txtStdLine'.
//
// ----------------------------------------------------------------
//
// # ITextFieldFormatDto Interface
//
//		This method processes objects implementing the
//		ITextFieldFormatDto interface to define text field
//		specifications used to generate multi-column lines of
//		text.
//
//		These text fields are then bundled to configure a
//		line of text returned as an instance of
//		TextLineSpecStandardLine.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//				TextFieldFormatDtoBigFloat
//				TextFieldFormatDtoDate
//				TextFieldFormatDtoLabel
//				TextFillerFieldFormatDto
//
//		The most frequently used type is the
//		TextFieldFormatDtoLabel structure which is defined
//		as follows:
//
//			type TextFieldFormatDtoLabel struct {
//
//				LeftMarginStr string
//					One or more characters used to create a left
//					margin for this Text Field.
//
//					If this parameter is set to an empty string, no
//					left margin will be configured for this Text
//					Field.
//
//				FieldContents interface{}
//					This parameter may contain one of several
//					specific data types. This empty interface type
//					will be converted to a string and configured as
//					the text column content within a text line.
//
//					Supported types which may be submitted through
//					this empty interface parameter are listed as
//					follows:
//
//					   time.Time (Converted using default format)
//					   string
//					   bool
//					   uint, uint8, uint16, uint32, uint64,
//					   int, int8, int16, int32, int64
//					   float32, float64
//					   *big.Int *big.Float
//					   fmt.Stringer (types that support this interface)
//					   TextInputParamFieldDateTimeDto
//					         (Converts date time to string. The best way
//					          to transmit and configure date time values.)
//
//				 FieldLength int
//					The length of the text field in which the
//					'FieldContents' will be displayed. If
//					'FieldLength' is less than the length of the
//					'FieldContents' string, it will be automatically
//					set equal to the 'FieldContents' string length.
//
//					To automatically set the value of 'FieldLength'
//					to the length of 'FieldContents', set this
//					parameter to a value of minus one (-1).
//
//					If this parameter is submitted with a value less
//					than minus one (-1) or greater than 1-million
//					(1,000,000), an error will be returned.
//
//					Field Length Examples
//
//						Example-1
//	 			        FieldContents String = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Center()
//							Text Field String =
//								"   Hello World!   "
//
//						Example-2
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Left()
//							Text Field String =
//								"Hello World!      "
//
//						Example-3
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = -1
//							FieldJustify = TxtJustify.Center() // Ignored
//							Text Field String =
//								"Hello World!"
//
//						Example-4
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 2
//							FieldJustify = TxtJustify.Center()
//								Ignored, because FieldLength Less
//								Than FieldContents String Length.
//							Text Field String =
//								"Hello World!"
//
//				 FieldJustify TextJustify
//					An enumeration which specifies the justification
//					of the 'FieldContents' string within the text
//					field length specified by 'FieldLength'.
//
//					Text justification can only be evaluated in the
//					context of a text label ('FieldContents'), field
//					length ('FieldLength') and a Text Justification
//					object of type TextJustify. This is because text
//					labels with a field length equal to or less than
//					the length of the text label string will never
//					use text justification. In these cases, text
//					justification is completely ignored.
//
//					If the field length is greater than the length of
//					the text label string, text justification must be
//					equal to one of these three valid values:
//
//					    TextJustify(0).Left()
//					    TextJustify(0).Right()
//					    TextJustify(0).Center()
//
//					Users can also specify the abbreviated text
//					justification enumeration syntax as follows:
//
//					    TxtJustify.Left()
//					    TxtJustify.Right()
//					    TxtJustify.Center()
//
//					Text Justification Examples
//
//						Example-1
//	 			        FieldContents String = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Center()
//							Text Field String =
//								"   Hello World!   "
//
//						Example-2
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Left()
//							Text Field String =
//								"Hello World!      "
//
//						Example-3
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = -1
//							FieldJustify = TxtJustify.Center() // Ignored
//							Text Field String =
//								"Hello World!"
//
//						Example-4
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 2
//							FieldJustify = TxtJustify.Center()
//								Ignored, because FieldLength Less
//								Than FieldContents String Length.
//							Text Field String =
//								"Hello World!"
//
//				RightMarginStr string
//					One or more characters used to create a right
//					margin for this Text Field.
//
//					If this parameter is set to an empty string, no
//					right margin will be configured for this Text
//					Field.
//			}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtStdLine					*TextLineSpecStandardLine
//
//		A pointer to an instance of
//		TextLineSpecStandardLine. Data extracted from the
//		following input parameters will be used to
//		completely reconfigure this instance of
//		TextLineSpecStandardLine.
//
//		All pre-existing data values contained in this
//		TextLineSpecStandardLine instance will be deleted
//		and overwritten with new values.
//
//	txtFieldFmtDtos				[]ITextFieldFormatDto
//
//		An array of objects implementing the
//		ITextFieldFormatDto interface. These Text Field
//		formatting objects contain all the text field
//		content and formatting specifications necessary
//		to format multiple text field columns in the
//		instance of TextLineSpecStandardLine passed as
//		input parameter 'txtStdLine'.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//			TextFieldFormatDtoBigFloat
//			TextFieldFormatDtoDate
//			TextFieldFormatDtoLabel
//			TextFillerFieldFormatDto
//
//		For additional information on the
//		ITextFieldFormatDto interface, see above
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
//	lastIndexId					int
//
//		If this method completes successfully, the
//		internal array index of the last text field
//		object for the current TextLineSpecStandardLine
//		instance will be returned as an integer value.
//
//		Remember that the last array index is equal the
//		array length minus one (array length - 1).
//
//		In the event of an error, 'lastIndexId' will be
//		set to a value of minus one (-1).
//
//	err							error
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
func (txtStdLineElectron *textLineSpecStandardLineElectron) addTextFieldDtosToArray(
	txtStdLine *TextLineSpecStandardLine,
	textFieldDtos []ITextFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	lastIndexId int,
	err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"addTextFieldDtosToArray()",
		"")

	if err != nil {
		return lastIndexId, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	if textFieldDtos == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFieldDtos' is a nil pointer!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	lenTextFieldFmtDtos := len(textFieldDtos)

	if lenTextFieldFmtDtos == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFieldDtos' is invalid!\n"+
			"'textFieldDtos' is an empty or zero length array.\n",
			ePrefix.String())

		return lastIndexId, err
	}

	var fieldContentsText string
	var txtFieldSpecLabel TextFieldSpecLabel

	for i := 0; i < lenTextFieldFmtDtos; i++ {

		if textFieldDtos[i].GetLeftMarginLength() > 0 {

			fieldContentsText =
				textFieldDtos[i].GetLeftMarginStr()

			txtFieldSpecLabel,
				err = TextFieldSpecLabel{}.NewTextLabel(
				fieldContentsText,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"Left Margin String"))

			if err != nil {
				return lastIndexId, err
			}

			txtStdLine.textFields =
				append(txtStdLine.textFields,
					&txtFieldSpecLabel)

		}

		txtFieldSpecLabel,
			err = textFieldDtos[i].GetFieldContentTextLabel(
			ePrefix.XCpy(
				fmt.Sprintf("textFieldDtos[%v]."+
					"FieldContents",
					i)))

		if err != nil {
			return lastIndexId, err
		}

		txtStdLine.textFields =
			append(txtStdLine.textFields,
				&txtFieldSpecLabel)

		if textFieldDtos[i].GetRightMarginLength() > 0 {

			fieldContentsText =
				textFieldDtos[i].GetRightMarginStr()

			txtFieldSpecLabel,
				err = TextFieldSpecLabel{}.NewTextLabel(
				fieldContentsText,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"Right Margin String"))

			if err != nil {
				return lastIndexId, err
			}

			txtStdLine.textFields =
				append(txtStdLine.textFields,
					&txtFieldSpecLabel)

		}
	}

	lastIndexId = len(txtStdLine.textFields) - 1

	return lastIndexId, err
}

// deleteTextField - Deletes a member of the Text Fields
// Collection. The array element to be deleted is designated by
// input parameter 'indexId'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	txtStdLine                 *TextLineSpecStandardLine
//	   - A pointer to an instance of TextLineSpecStandardLine which
//	     encapsulates the Text Fields Collection. The member of
//	     this collection designated by parameter, 'indexId' will be
//	     deleted.
//
//
//	indexId                    int
//	   - The index number of the array element in the Text Fields
//	     Collection which will be deleted.
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
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtStdLineElectron *textLineSpecStandardLineElectron) deleteTextField(
	txtStdLine *TextLineSpecStandardLine,
	indexId int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"testValidityOfTextFields()",
		"")

	if err != nil {
		return err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if indexId < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'indexId' is invalid!\n"+
			"'indexId' is less than zero.\n"+
			"indexId = '%v'\n",
			ePrefix.String(),
			indexId)

		return err
	}

	lenTextFieldCollection := len(txtStdLine.textFields)

	if lenTextFieldCollection == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Field Collection, 'txtStdLine.textFields' is EMPTY!\n",
			ePrefix.String())

		return err
	}

	lastIdx := lenTextFieldCollection - 1

	if indexId > lastIdx {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'indexId' is invalid!\n"+
			"'indexId' is greater than the last index\n"+
			"in the Text Fields Collection.\n"+
			"Last index in collection = '%v'\n"+
			"indexId = '%v'\n",
			ePrefix.String(),
			lastIdx,
			indexId)

		return err

	}

	if indexId == 0 {

		if txtStdLine.textFields[0] != nil {

			txtStdLine.textFields[0].Empty()

			txtStdLine.textFields[0] = nil

		}

		txtStdLine.textFields = txtStdLine.textFields[1:]

	} else if indexId == lastIdx {

		if txtStdLine.textFields[lastIdx] != nil {

			txtStdLine.textFields[lastIdx].Empty()

			txtStdLine.textFields[lastIdx] = nil

		}

		txtStdLine.textFields = txtStdLine.textFields[0:lastIdx]

	} else {

		if txtStdLine.textFields[indexId] != nil {

			txtStdLine.textFields[indexId].Empty()

			txtStdLine.textFields[indexId] = nil

		}

		txtStdLine.textFields = append(
			txtStdLine.textFields[0:indexId],
			txtStdLine.textFields[indexId+1:]...)

	}

	return err
}

// emptyTextFields - Receives a pointer to an array of Text Fields
// and proceeds to empty or delete the elements contained in that
// array.
//
// Text Fields is an array of the ITextFieldSpecification objects.
//
// The ITextFieldSpecification interface defines a text field used
// in conjunction with the type, TextLineSpecStandardLine. This
// type contains an array of text field or ITextFieldSpecification
// objects. Text fields are the building blocks of lines of text
// which are formatted by TextLineSpecStandardLine for text
// displays, file output or printing.
//
// This method will call method 'Empty()' on each of the elements
// contained in the Text Field array passed as input parameter,
// 'textFields'. Upon completion the concrete 'textFields' array
// will be set to 'nil'.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// This method deletes and overwrites all the member elements, and
// their data values, contained in input parameter,
// 'textFields'. Upon completion, the concrete instance of the
// 'textFields' array is also set to 'nil'.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	textFields                 *[]ITextFieldSpecification
//	   - A pointer to an array of text fields.
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
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtStdLineElectron *textLineSpecStandardLineElectron) emptyTextFields(
	textFields *[]ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"emptyTextFields()",
		"")

	if err != nil {
		return err
	}

	if textFields == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter textFields is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	concreteTxtFields := *textFields

	lenTextFields := len(concreteTxtFields)

	if lenTextFields == 0 {
		return err
	}

	for i := 0; i < lenTextFields; i++ {

		if concreteTxtFields[i] == nil {
			continue
		}

		concreteTxtFields[i].Empty()

		concreteTxtFields[i] = nil
	}

	*textFields = nil

	endingLenTextFields := len(*textFields)

	if endingLenTextFields != 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Empty Text Fields Array Operation Failed!\n"+
			"'textFields' has an array length greater than zero.\n"+
			"Begining Length 'textFields' = '%v'\n"+
			"  Ending Length 'textFields' = '%v'\n",
			ePrefix.String(),
			lenTextFields,
			endingLenTextFields)

	}

	return err
}

// equalTextFieldArrays - Compares two text fields arrays and
// returns a boolean value of 'true' if the two arrays are equal
// in all respects.
//
// To qualify as equivalent arrays of Text Fields, both arrays must
// be equal in length and contain Text Fields of equal values.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	textFields01               *[]ITextFieldSpecification
//	   - If parameter 'textFields01' is equivalent, in all
//	     respects, to parameter 'textFields02', this method will
//	     return a boolean value of 'true'.
//
//	     If this pointer to an array of ITextFieldSpecification's
//	     is nil, a value of 'false' will be returned.
//
//
//	textFields02               *[]ITextFieldSpecification
//	   - If parameter 'textFields02' is equivalent, in all
//	     respects, to parameter 'textFields01', this method will
//	     return a boolean value of 'true'.
//
//	     If this pointer to an array of ITextFieldSpecification's
//	     is nil, a value of 'false' will be returned.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	areEqual                   bool
//	   - If Text Field Arrays 'textFields01' and 'textFields02' are
//	     equal in all respects, a boolean value of 'true' is
//	     returned. Otherwise, the return value is 'false'.
func (txtStdLineElectron *textLineSpecStandardLineElectron) equalTextFieldArrays(
	textFields01 *[]ITextFieldSpecification,
	textFields02 *[]ITextFieldSpecification) (
	areEqual bool) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	if textFields01 == nil ||
		textFields02 == nil {
		return false
	}

	txtF01 := *textFields01

	txtF02 := *textFields02

	if txtF01 == nil &&
		txtF02 == nil {

		return true
	}

	lenTxtFields01 := len(txtF01)

	if lenTxtFields01 != len(txtF02) {
		return false
	}

	if lenTxtFields01 == 0 {
		return true
	}

	for i := 0; i < lenTxtFields01; i++ {

		if txtF01[i] == nil &&
			txtF02[i] != nil {

			return false

		} else if txtF02[i] == nil &&
			txtF01[i] != nil {

			return false

		} else if txtF02[i] == nil &&
			txtF01[i] == nil {

			continue

		} else {
			// txtF01 and txtF02 are both
			// != nil
			if !txtF01[i].EqualITextField(
				txtF02[i]) {

				return false

			}

		}
	}

	return true
}

// testValidityOfTextFields - This method receives an array of
// ITextFieldSpecification objects and subjects that array to
// an analysis which determines whether the array, and all of its
// member elements, is valid in all respects.
//
// Type TextLineSpecStandardLine is a text specification for a
// standard line of text. A standard line of text comprises a
// single line of text which may be repeated one or more times.
//
// TextLineSpecStandardLine encapsulates an array of
// ITextFieldSpecification objects which are used to format text
// fields within a single line of text. Essentially, a standard
// text line is a collection of text fields which implement the
// ITextFieldSpecification interface. Text fields can be thought of
// as the building blocks for a standard line of text.
//
// ------------------------------------------------------------------
//
// Input Parameters
//
//	textFields                 *[]ITextFieldSpecification
//	   - 'textFields' is a pointer to a collection of objects
//	     implementing the ITextLineSpecification interface. These
//	     text fields are assembled by the TextLineSpecStandardLine
//	     type and formatted as a single line of text. Text fields
//	     are the building blocks for a standard line of text
//	     characters produced by type TextLineSpecStandardLine.
//
//	     If this pointer is nil, an error will be returned.
//
//	     If this array is 'nil' or has a zero length, an error will
//	     be returned.
//
//	     If any of the ITextFieldSpecification objects within this
//	     array are found to be invalid, an error will be returned.
//
//	     If any of the individual elements within this array as a
//	     'nil' value, an error will be returned.
//
//
//	allowZeroLengthTextFieldsArray  bool
//	   - When set to 'true', no error will be generated if the
//	     input parameter 'textFields' contains a zero length array.
//
//	     Conversely, if 'allowZeroLengthTextFieldsArray' is set to
//	     'false', an error WILL be returned if 'textFields' is a
//	     zero length array.
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
// -----------------------------------------------------------------
//
// Return Values
//
//	lengthTextFields           int
//	   - If input parameter 'textFields' is judged to be valid in
//	     all respects, this return parameter will be set to the
//	     array length of 'textFields'.
//
//	   - If input parameter 'textFields' is found to be invalid,
//	     this return parameter will be set to 'false'.
//
//
//	err                        error
//	   - If input parameter 'textFields' is judged to be valid in
//	     all respects, this return parameter will be set to 'nil'.
//
//	     If input parameter, 'textFields' is found to be invalid,
//	     this return parameter will be configured with an appropriate
//	     error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtStdLineElectron *textLineSpecStandardLineElectron) testValidityOfTextFields(
	textFields *[]ITextFieldSpecification,
	allowZeroLengthTextFieldsArray bool,
	errPrefDto *ePref.ErrPrefixDto) (
	lengthTextFields int,
	err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"testValidityOfTextFields()",
		"")

	if err != nil {
		return lengthTextFields, err
	}

	if textFields == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'textFields' is a 'nil' pointer!\n",
			ePrefix.String())

		return lengthTextFields, err

	}

	lengthTextFields = 0

	var err2 error

	for idx, val := range *textFields {

		if val == nil {
			err = fmt.Errorf("%v - ERROR\n"+
				"textFields[%v] is 'nil'!\n",
				idx,
				ePrefix.String())

			return lengthTextFields, err
		}

		err2 = val.IsValidInstanceError(
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: 'textFields' element is invalid!\n"+
				"textFields[%v] failed validity test.\n"+
				"Validity Error:\n%v\n",
				ePrefix.String(),
				idx,
				err2.Error())

			return lengthTextFields, err
		}

		lengthTextFields++
	}

	if lengthTextFields == 0 &&
		!allowZeroLengthTextFieldsArray {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is empty!\n"+
			"'textFields' is a zero length array.\n",
			ePrefix.String())

		return lengthTextFields, err
	}

	return lengthTextFields, err
}
