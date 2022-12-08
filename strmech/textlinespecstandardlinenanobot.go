package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineNanobot struct {
	lock *sync.Mutex
}

// addTextFields - STOP@COL68
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	txtStdLine                 *TextLineSpecStandardLine
//	   - A pointer to an instance of TextLineSpecStandardLine which
//	     encapsulates the Text Fields Collection. Deep copies of
//	     text fields contained in input parameter 'textFields' will
//	     be added to the Text Fields Collection for this object.
//
//
//	textFields                 []ITextFieldSpecification
//	   - An array of text field objects which implement the
//	     ITextFieldSpecification interface. A deep copy of each
//	     Text Field object will be added to the text field
//	     collection maintained by input parameter,
//	     'txtStdLine'.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'textField' as a pointer to 'textField' (&textField).
//
//	     If member variable data values contained in this
//	     'textField' parameter are found to be invalid, an error
//	     will be returned.
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
//	lastIndexId                int
//	   - If this method completes successfully, the internal array
//	     index of the last text field object in the Text Field
//	     array maintained by parameter 'txtStdLine' will be
//	     returned as an integer value.
//
//	     In the event of an error, 'lastIndexId' will be set to a
//	     value of minus one (-1).
//
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) addTextFields(
	txtStdLine *TextLineSpecStandardLine,
	textFields *[]ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (
	lastIndexId int,
	err error) {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot.addTextFields()",
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

	if textFields == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is a nil pointer!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	copyCnt := 0

	for idx, val := range *textFields {

		err = val.IsValidInstanceError(
			ePrefix.XCpy(
				fmt.Sprintf("textFields[%v]",
					idx)))

		if err != nil {
			return lastIndexId, err
		}

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCpy(fmt.Sprintf(
				"textFields[%v]->newTextField",
				idx)))

		if err != nil {
			return lastIndexId, err
		}

		txtStdLine.textFields =
			append(txtStdLine.textFields,
				newTextField)

		copyCnt++
	}

	if copyCnt == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'textFields' is empty!\n",
			ePrefix.String())

		return lastIndexId, err
	}

	lastIndexId = len(txtStdLine.textFields) - 1

	return lastIndexId, err
}

// copyIn - Copies all data from input parameter 'incomingStdLine'
// to input parameter 'targetStdLine'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing data fields for input parameter 'targetStdLine'
// will be overwritten and deleted.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	targetStdLine              *TextLineSpecStandardLine
//	   - A pointer to an instance of TextLineSpecStandardLine. Data
//	     extracted from input parameter 'incomingStdLine' will
//	     be copied to this input parameter, 'targetStdLine'. If
//	     this method completes successfully, all member data
//	     variables encapsulated in 'targetStdLine' will be
//	     identical to those contained in input parameter,
//	     'incomingStdLine'.
//
//	     Be advised that the pre-existing data fields in input
//	     parameter 'targetStdLine' will be overwritten and
//	     deleted.
//
//
//	incomingStdLine            *TextLineSpecStandardLine
//	   - A pointer to an instance of TextLineSpecStandardLine.
//
//	     All data values in this TextLineSpecStandardLine instance
//	     will be copied to input parameter 'targetStdLine'.
//
//	     The original member variable data values encapsulated in
//	     'incomingStdLine' will remain unchanged and will NOT be
//	     overwritten or deleted.
//
//	     If 'incomingStdLine' contains invalid member data
//	     variables, this method will return an error.
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) copyIn(
	targetStdLine *TextLineSpecStandardLine,
	incomingStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStdLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingStdLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(incomingStdLine.newLineChars) == 0 {

		err =
			textSpecificationMolecule{}.ptr().
				setDefaultNewLineChars(
					&incomingStdLine.newLineChars,
					ePrefix)

		if err != nil {
			return err
		}

	}

	txtStdLineAtom :=
		textLineSpecStandardLineAtom{}

	_,
		err = txtStdLineAtom.
		testValidityOfTextLineSpecStdLine(
			incomingStdLine,
			false, // allowZeroLengthTextFieldsArray
			ePrefix.XCpy(
				"incomingStdLine"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetStdLine.newLineChars,
		&incomingStdLine.newLineChars,
		true,
		ePrefix.XCpy(
			"incomingStdLine.newLineChars->"+
				"targetStdLine.newLineChars"))

	if err != nil {
		return err
	}

	targetStdLine.turnLineTerminatorOff =
		incomingStdLine.turnLineTerminatorOff

	targetStdLine.numOfStdLines =
		incomingStdLine.numOfStdLines

	targetStdLine.textLineReader = nil

	_,
		err = txtStdLineAtom.
		copyTextFields(
			&targetStdLine.textFields,
			&incomingStdLine.textFields,
			ePrefix.XCpy(
				"incomingStdLine.textFields->"+
					"targetStdLine.textFields"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'txtStdLine'.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	txtStdLine          *TextLineSpecStandardLine
//	   - A pointer to an instance of TextLineSpecStandardLine. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of TextLineSpecStandardLine.
//
//	     If the member variable data values encapsulated by this
//	     'txtStdLine' are found to be invalid, this method will
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
//	TextLineSpecStandardLine
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'txtStdLine' will be created and returned
//	     in a new instance of TextLineSpecStandardLine.
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) copyOut(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecStandardLine, error) {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot."+
			"copyOut()",
		"")

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecStandardLine{}, err
	}

	if len(txtStdLine.newLineChars) == 0 {

		err =
			textSpecificationMolecule{}.ptr().
				setDefaultNewLineChars(
					&txtStdLine.newLineChars,
					ePrefix)

		if err != nil {
			return TextLineSpecStandardLine{}, err
		}

	}

	txtStdLineAtom :=
		textLineSpecStandardLineAtom{}

	_,
		err = txtStdLineAtom.
		testValidityOfTextLineSpecStdLine(
			txtStdLine,
			false, // allowZeroLengthTextFieldsArray
			ePrefix.XCpy(
				"txtStdLine"))

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	newStdLine.textLineReader = nil

	err = strMechPreon{}.ptr().
		copyRuneArrays(
			&newStdLine.newLineChars,
			&txtStdLine.newLineChars,
			true,
			ePrefix.XCpy(
				"txtStdLine.newLineChars->"+
					"newStdLine.newLineChars"))

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	newStdLine.turnLineTerminatorOff =
		txtStdLine.turnLineTerminatorOff

	newStdLine.numOfStdLines = txtStdLine.numOfStdLines

	_,
		err = txtStdLineAtom.
		copyTextFields(
			&newStdLine.textFields,
			&txtStdLine.textFields,
			ePrefix.XCpy(
				"txtStdLine.textFields->"+
					"newStdLine.textFields"))

	return newStdLine, err
}

// insertTextFieldAtIndex - Receives an instance of Text Field in
// the form of an ITextFieldSpecification. This Test Field object
// is then inserted into the Text Fields Collection maintained by
// the TextLineSpecStandardLine instance passed as input parameter,
// 'txtStdLine'.
//
// The Text Field input parameter, 'iTextField', is inserted into
// the internal Text Field collection array at the array element
// index position indicated by input parameter, 'indexId'.
//
// After this method completes, the number of elements in the Text
// Fields Collection will be increased by one.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	txtStdLine                 *TextLineSpecStandardLine
//	   - A pointer to an instance of type TextLineSpecStandardLine.
//	     The Text Field provided by input parameter 'iTextField'
//	     will be inserted at the array element index position
//	     indicated by input parameter 'indexId'.
//
//
//	iTextField                 ITextFieldSpecification
//	   - A text field object which implements the
//	     ITextFieldSpecification interface. A deep copy of this
//	     object will be inserted into the Text Field collection
//	     maintained by the instance of TextLineSpecStandardLine
//	     pointed to by input parameter 'txtStdLine'.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'iTextField' as a pointer to the Text Field (&textField).
//
//	     If the 'iTextField' parameter are found to be invalid, an
//	     error will be returned.
//
//
//	indexId                    int
//	   - This index number designates the array element index in
//	     the Text Fields Collection of the 'txtStdLine' instance at
//	     which the Text Fields parameter, 'iTextField' will be
//	     inserted. This means that 'iTextField' will be inserted
//	     immediately BEFORE the array element specified by
//	     'indexId' in the final Text Fields Array.
//
//	     If the value of 'indexId' is less than zero, it will be
//	     reset to zero. This means that the 'iTextField' object
//	     will be inserted in the first array element position of
//	     the Text Fields Collection maintained by parameter,
//	     'txtStdLine'.
//
//	     If the value of 'indexId' is greater the last array
//	     element index in the 'txtStdLine' Text Fields Collection,
//	     the 'iTextField' object will be appended to the end of
//	     that Text Fields Collection.
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
//	lengthTextFields           int
//	   - If this method completes successfully, 'lengthTextFields'
//	     will be set to the length of the Text Fields collections
//	     contained in input parameter, 'txtStdLine'. This means
//	     that the value of 'lengthTextFields' will be one greater
//	     than the original length of the 'txtStdLine' Text Fields
//	     Collection.
//
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) insertTextFieldAtIndex(
	txtStdLine *TextLineSpecStandardLine,
	iTextField ITextFieldSpecification,
	indexId int,
	errPrefDto *ePref.ErrPrefixDto) (
	lengthTextFields int,
	err error) {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot."+
			"insertTextFieldAtIndex()",
		"")

	if err != nil {
		return lengthTextFields, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return lengthTextFields, err
	}

	_,
		err = textLineSpecStandardLineAtom{}.ptr().
		testValidityOfTextLineSpecStdLine(
			txtStdLine,
			true, // allowZeroLengthTextFieldsArray
			ePrefix.XCpy(
				"txtStdLine"))

	if err != nil {
		return lengthTextFields, err
	}

	if iTextField == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'iTextField' is 'nil'!\n",
			ePrefix.String())

		return lengthTextFields, err
	}

	err = iTextField.IsValidInstanceError(
		ePrefix.XCpy(
			"iTextField is invalid!"))

	if err != nil {
		return lengthTextFields, err
	}

	var newTextField ITextFieldSpecification

	newTextField,
		err = iTextField.CopyOutITextField(
		ePrefix.XCpy("iTextField->newTextField"))

	if err != nil {
		return lengthTextFields, err
	}

	lengthTextFields = len(txtStdLine.textFields)

	if lengthTextFields == 0 ||
		indexId >= lengthTextFields {

		txtStdLine.textFields = append(
			txtStdLine.textFields,
			newTextField)

		lengthTextFields++

		return lengthTextFields, err
	}

	if indexId < 0 {
		indexId = 0
	}

	var oldTextField ITextFieldSpecification

	oldTextField,
		err = txtStdLine.textFields[indexId].CopyOutITextField(
		ePrefix.XCpy(fmt.Sprintf(
			"oldTextField<-txtStdLine.textFields[%v]",
			indexId)))

	if err != nil {
		return lengthTextFields, err
	}

	// arr := []int{1,2,3,4,5}
	// arr[:2]         [1,2]
	// arr[2:])        [3,4,5]
	// 	orig = append(orig[:index+1], orig[index:]...)

	if txtStdLine.textFields[indexId] != nil {

		txtStdLine.textFields[indexId].Empty()

		txtStdLine.textFields[indexId] = nil

	}

	txtStdLine.textFields = append(
		txtStdLine.textFields[:indexId+1],
		txtStdLine.textFields[indexId:]...)

	txtStdLine.textFields[indexId+1] =
		oldTextField

	txtStdLine.textFields[indexId] =
		newTextField

	lengthTextFields++

	return lengthTextFields, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineNanobot.
func (txtStdLineNanobot textLineSpecStandardLineNanobot) ptr() *textLineSpecStandardLineNanobot {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	return &textLineSpecStandardLineNanobot{
		lock: new(sync.Mutex),
	}
}

// setTextFieldFmtStdLine
//
// Designed to configure a Standard Line containing one
// or more text field columns.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values in the
//	TextLineSpecStandardLine instance passed as input
//	parameter 'txtStdLine'.
//
// ----------------------------------------------------------------
//
// # Type TextFieldFormatDto
//
//	This method employs type TextFieldFormatDto to define
//	text fields for the two column layout configured for
//	the returned instance of TextLineSpecStandardLine.
//
//	The type TextFieldFormatDto structure is defined as
//	follows:
//
//		type TextFieldFormatDto struct {
//
//			LeftMarginStr string
//				One or more characters used to create a left
//				margin for this Text Field.
//
//				If this parameter is set to an empty string, no
//				left margin will be configured for this Text
//				Field.
//
//			FieldContents interface{}
//				This parameter may contain one of several
//				specific data types. This empty interface type
//				will be converted to a string and configured as
//				the text column content within a text line.
//
//				Supported types which may be submitted through
//				this empty interface parameter are listed as
//				follows:
//
//				   time.Time (Converted using default format)
//				   string
//				   bool
//				   uint, uint8, uint16, uint32, uint64,
//				   int, int8, int16, int32, int64
//				   float32, float64
//				   *big.Int *big.Float
//				   fmt.Stringer (types that support this interface)
//				   TextInputParamFieldDateTimeDto
//				         (Converts date time to string. The best way
//				          to transmit and configure date time values.)
//
//			 FieldLength int
//				The length of the text field in which the
//				'FieldContents' will be displayed. If
//				'FieldLength' is less than the length of the
//				'FieldContents' string, it will be automatically
//				set equal to the 'FieldContents' string length.
//
//				To automatically set the value of 'FieldLength'
//				to the length of 'FieldContents', set this
//				parameter to a value of minus one (-1).
//
//				If this parameter is submitted with a value less
//				than minus one (-1) or greater than 1-million
//				(1,000,000), an error will be returned.
//
//			 FieldJustify TextJustify
//				An enumeration which specifies the justification
//				of the 'FieldContents' string within the text
//				field length specified by 'FieldLength'.
//
//				Text justification can only be evaluated in the
//				context of a text label ('FieldContents'), field
//				length ('FieldLength') and a Text Justification
//				object of type TextJustify. This is because text
//				labels with a field length equal to or less than
//				the length of the text label string will never
//				use text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length is greater than the length of
//				the text label string, text justification must be
//				equal to one of these three valid values:
//
//				    TextJustify(0).Left()
//				    TextJustify(0).Right()
//				    TextJustify(0).Center()
//
//				Users can also specify the abbreviated text
//				justification enumeration syntax as follows:
//
//				    TxtJustify.Left()
//				    TxtJustify.Right()
//				    TxtJustify.Center()
//
//			RightMarginStr string
//				One or more characters used to create a right
//				margin for this Text Field.
//
//				If this parameter is set to an empty string, no
//				right margin will be configured for this Text
//				Field.
//		}
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
//	txtFieldFmtDtos				[]TextFieldFormatDto
//
//		An array of TextFieldFormatDto objects containing
//		all the text field content and formatting
//		specifications necessary to format multiple text
//		field columns in the returned instance of
//		TextLineSpecStandardLine.
//
//		For more information on Type TextFieldFormatDto,
//		see above.
//
//	numOfStdLines				int
//
//		An integer value specifying the number of
//		repetitions for a standard line text formatted for
//		screen display, file output or printing.
//
//		A 'numOfStdLines' value of 1 means the line will
//		be output once, a value of 2 signals the line
//		will be repeated or output twice, a value of '3'
//		signals the line will be output 3-times and so
//		on.
//
//		If the 'numOfStdLines' value is set to zero, no
//		text line will be formatted for text display, file
//		output or printing.
//
//		If this value is set to a value less than zero,
//		it will be automatically reset to a value of one
//		('1').
//
//		The following examples illustrate the use of
//		'numOfStdLines':
//
//		  Example #1:
//		   Standard Line Text = "Hello World"
//		   numOfStdLines = 1
//		   Text Output:
//		     "Hello World"
//
//		  Example #2:
//		   Standard Line Text = "Hello World"
//		   numOfStdLines = 3
//		   Text Output:
//		     "Hello World"
//		     "Hello World"
//		     "Hello World"
//
//	newLineChars []rune
//
//		An array of runes which contains the text
//		characters which will be applied as line
//		termination characters for each line of text
//		produced by the returned instance of
//		TextLineSpecStandardLine.
//
//		By default, each line of text generated by
//		TextLineSpecStandardLine will be terminated with
//		a new line character ('\n'). However, this
//		parameter allows the user to specify the
//		character or characters to be used as a line
//		termination sequence for each line of text
//		produced by the returned instance of
//		TextLineSpecStandardLine.
//
//		If this parameter is submitted as a 'nil' value
//		or, if 'newLineChars' is a zero length array,
//		this method will set 'newLineChars' to the
//		default new line termination character ('\n').
//
//	turnLineTerminatorOff		bool
//
//		The 'turnLineTerminatorOff' flag controls whether
//		a line termination character or characters will
//		be automatically appended to each line of text
//		produced by TextLineSpecStandardLine.
//
//		When the boolean flag 'turnLineTerminatorOff' is
//		set to 'false', line terminators as defined by
//		parameter 'lineTerminator' will be applied as a
//		line termination sequence for each line of text
//		produced by TextLineSpecStandardLine.
//
//		When this boolean value is set to 'true', it
//		turns off or cancels the automatic generation of
//		line terminators for each line of text produced
//		by TextLineSpecStandardLine.
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) setTextFieldFmtStdLine(
	txtStdLine *TextLineSpecStandardLine,
	txtFieldFmtDtos []TextFieldFormatDto,
	numOfStdLines int,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot."+
			"setTextFieldFmtStdLine()",
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

	lenTextFieldFmtDtos := len(txtFieldFmtDtos)

	if lenTextFieldFmtDtos == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFmtDtos' is invalid!\n"+
			"'txtFieldFmtDtos' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	if numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value less than one (1).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if numOfStdLines > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value greater than one-million (1,000,000).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if len(newLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newLineChars' is invalid!\n"+
			"'newLineChars' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}
	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineChars,
		ePrefix.XCpy(
			"newLineChars"))

	if err != nil {
		return err
	}

	_ =
		textLineSpecStandardLineMolecule{}.ptr().
			emptyStandardLine(
				txtStdLine,
				ePrefix.XCpy(
					"empty->txtStdLine"))

	var fieldContentsText string
	var txtFieldSpecLabel TextFieldSpecLabel

	for i := 0; i < lenTextFieldFmtDtos; i++ {

		if txtFieldFmtDtos[i].GetLeftMarginLength() > 0 {

			fieldContentsText =
				txtFieldFmtDtos[i].GetLeftMarginStr()

			txtFieldSpecLabel,
				err = TextFieldSpecLabel{}.NewTextLabel(
				fieldContentsText,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"Left Margin String"))

			if err != nil {
				return err
			}

			txtStdLine.textFields =
				append(txtStdLine.textFields,
					&txtFieldSpecLabel)

		}

		txtFieldSpecLabel,
			err = txtFieldFmtDtos[i].GetFieldContentTextLabel(
			ePrefix.XCpy(
				fmt.Sprintf("txtFieldFmtDtos[%v]."+
					"FieldContents",
					i)))

		if err != nil {
			return err
		}

		txtStdLine.textFields =
			append(txtStdLine.textFields,
				&txtFieldSpecLabel)

		if txtFieldFmtDtos[i].GetRightMarginLength() > 0 {

			fieldContentsText =
				txtFieldFmtDtos[i].GetRightMarginStr()

			txtFieldSpecLabel,
				err = TextFieldSpecLabel{}.NewTextLabel(
				fieldContentsText,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"Right Margin String"))

			if err != nil {
				return err
			}

			txtStdLine.textFields =
				append(txtStdLine.textFields,
					&txtFieldSpecLabel)

		}

	}

	txtStdLine.numOfStdLines =
		numOfStdLines

	txtStdLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&txtStdLine.newLineChars,
		&newLineChars,
		true,
		ePrefix.XCpy(
			"txtStdLine.newLineChars<-newLineChars"))

	return err
}

// setTxtSpecStandardLine
//
// Reconfigures all the data values for input parameter
// 'txtStdLine', a pointer to an instance of
// TextLineSpecStandardLine.
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
//	numOfStdLines				int
//
//		An integer value specifying the number of
//		repetitions for a standard line text formatted for
//		screen display, file output or printing.
//
//		A 'numOfStdLines' value of 1 means the line will
//		be output once, a value of 2 signals the line
//		will be repeated or output twice, a value of '3'
//		signals the line will be output 3-times and so
//		on.
//
//		If the 'numOfStdLines' value is set to zero, no
//		text line will be formatted for text display, file
//		output or printing.
//
//		If this value is set to a value less than zero,
//		it will be automatically reset to a value of one
//		('1').
//
//		The following examples illustrate the use of
//		'numOfStdLines':
//
//		  Example #1:
//		   Standard Line Text = "Hello World"
//		   numOfStdLines = 1
//		   Text Output:
//		     "Hello World"
//
//		  Example #2:
//		   Standard Line Text = "Hello World"
//		   numOfStdLines = 3
//		   Text Output:
//		     "Hello World"
//		     "Hello World"
//		     "Hello World"
//
//	textFields					[]ITextFieldSpecification
//
//		'textFields' is a collection of objects
//		implementing the ITextLineSpecification
//		interface. These text fields are assembled by the
//		TextLineSpecStandardLine type and formatted as a
//		single line of text. This single line of text is
//		output one or more times as specified by input
//		parameter, 'numOfStdLines'.
//
//		Text fields are the building blocks used to
//		assemble a standard line of text.
//
//		If this parameter is submitted as a 'nil' value
//		or a zero length array, an error will be
//		returned.
//
//		If any of the objects contained in this
//		collection are invalid, an error will be
//		returned.
//
//	newLineChars []rune
//
//		An array of runes which contains the text
//		characters which will be applied as line
//		termination characters for each line of text
//		produced by the returned instance of
//		TextLineSpecStandardLine.
//
//		By default, each line of text generated by
//		TextLineSpecStandardLine will be terminated with
//		a new line character ('\n'). However, this
//		parameter allows the user to specify the
//		character or characters to be used as a line
//		termination sequence for each line of text
//		produced by the returned instance of
//		TextLineSpecStandardLine.
//
//		If this parameter is submitted as a 'nil' value
//		or, if 'newLineChars' is a zero length array,
//		this method will set 'newLineChars' to the
//		default new line termination character ('\n').
//
//	turnLineTerminatorOff		bool
//
//		The 'turnLineTerminatorOff' flag controls whether
//		a line termination character or characters will
//		be automatically appended to each line of text
//		produced by TextLineSpecStandardLine.
//
//		When the boolean flag 'turnLineTerminatorOff' is
//		set to 'false', line terminators as defined by
//		parameter 'lineTerminator' will be applied as a
//		line termination sequence for each line of text
//		produced by TextLineSpecStandardLine.
//
//		When this boolean value is set to 'true', it
//		turns off or cancels the automatic generation of
//		line terminators for each line of text produced
//		by TextLineSpecStandardLine.
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
func (txtStdLineNanobot *textLineSpecStandardLineNanobot) setTxtSpecStandardLine(
	txtStdLine *TextLineSpecStandardLine,
	numOfStdLines int,
	textFields []ITextFieldSpecification,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineNanobot.lock == nil {
		txtStdLineNanobot.lock = new(sync.Mutex)
	}

	txtStdLineNanobot.lock.Lock()

	defer txtStdLineNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineNanobot."+
			"setTxtSpecStandardLine()",
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

	if len(textFields) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"'textFields' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	if numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value less than one (1).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if numOfStdLines > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value greater than one-million (1,000,000).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if len(newLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newLineChars' is invalid!\n"+
			"'newLineChars' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	_,
		err =
		textLineSpecStandardLineElectron{}.ptr().
			testValidityOfTextFields(
				&textFields,
				false, // allowZeroLengthTextFieldsArray
				ePrefix.XCpy("Input parameter textFields is invalid!"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineChars,
		ePrefix.XCpy(
			"newLineChars"))

	if err != nil {
		return err
	}

	_ =
		textLineSpecStandardLineMolecule{}.ptr().
			emptyStandardLine(
				txtStdLine,
				ePrefix.XCpy(
					"empty->txtStdLine"))

	_,
		err = textLineSpecStandardLineAtom{}.ptr().
		copyTextFields(
			&txtStdLine.textFields,
			&textFields,
			ePrefix.XCpy(
				"textFields->"+
					"txtStdLine.textFields"))

	if err != nil {
		return err
	}

	txtStdLine.numOfStdLines =
		numOfStdLines

	txtStdLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&txtStdLine.newLineChars,
		&newLineChars,
		true,
		ePrefix.XCpy(
			"newLineChars->txtStdLine.newLineChars"))

	return err
}
