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
	txtSpecAtom := textSpecificationAtom{}

	for i := 0; i < lenTextFieldFmtDtos; i++ {

		fieldContentsText,
			err = txtSpecAtom.
			convertParamEmptyInterfaceToString(
				txtFieldFmtDtos[i].FieldContents,
				fmt.Sprintf("txtFieldFmtDtos[%v].FieldContents",
					i),
				ePrefix.XCpy(
					fmt.Sprintf("txtFieldFmtDtos[%v].FieldContents",
						i)))

		if err != nil {
			return err
		}

		txtFieldSpecLabel,
			err = TextFieldSpecLabel{}.NewTextLabel(
			fieldContentsText,
			txtFieldFmtDtos[i].FieldLength,
			txtFieldFmtDtos[i].FieldJustify,
			ePrefix.XCpy(
				"txtFieldSpecLabel<-fieldContentsText"))

		if err != nil {
			return err
		}

		txtStdLine.textFields =
			append(txtStdLine.textFields,
				&txtFieldSpecLabel)
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

// setTxtSpecStandardLine - Reconfigures all the data values for
// input parameter 'txtStdLine', a pointer to an instance of
// TextLineSpecStandardLine.
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
