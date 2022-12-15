package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecFillerMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTxtFiller' to input parameter 'incomingTxtFiller'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in 'targetTxtFiller' will be
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	targetTxtFiller            *TextFieldSpecFiller
//	   - A pointer to an instance of TextFieldSpecFiller. Data
//	     extracted from input parameter 'incomingTxtFiller' will
//	     be copied to this input parameter, 'targetTxtFiller'. If
//	     this method completes successfully, all member data
//	     variables encapsulated in 'targetTxtFiller' will be
//	     identical to those contained in input parameter,
//	     'incomingTxtFiller'.
//
//	     Be advised that the pre-existing data fields in input
//	     parameter 'targetTxtFiller' will be overwritten and
//	     deleted.
//
//
//	incomingTxtFiller          *TextFieldSpecFiller
//	   - A pointer to an instance of TextFieldSpecFiller.
//
//	     All data values in this TextFieldSpecFiller instance
//	     will be copied to input parameter 'targetTxtFiller'.
//
//	     The original member variable data values encapsulated in
//	     'incomingTxtFiller' will remain unchanged and will NOT be
//	     overwritten or deleted.
//
//	     If 'incomingTxtFiller' contains invalid member data
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
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) copyIn(
	targetTxtFiller *TextFieldSpecFiller,
	incomingTxtFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTxtFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTxtFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTxtFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err =
		textFieldSpecFillerAtom{}.ptr().
			isValidTextFieldSpecFiller(
				incomingTxtFiller,
				ePrefix.XCpy("incomingTxtFiller validation - "))

	if err != nil {
		return err
	}

	// Setting zero length array to nil.
	err = strMechPreon{}.ptr().copyRuneArrays(
		&targetTxtFiller.fillerCharacters,
		&incomingTxtFiller.fillerCharacters,
		true,
		ePrefix.XCpy("targetTxtFiller.fillerCharacters=Target "+
			"<-incomingTxtFiller.fillerCharacters=Source"))

	if err != nil {
		return err
	}

	targetTxtFiller.fillerCharsRepeatCount =
		incomingTxtFiller.fillerCharsRepeatCount

	targetTxtFiller.textLineReader = nil

	return nil
}

// copyOut - Returns a deep copy of the input parameter
// 'txtFieldFiller'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	txtFieldFiller      *TextFieldSpecFiller
//	   - A pointer to an instance of TextFieldSpecFiller. A deep
//	     copy of the internal member variables will be created
//	     and returned in a new instance of TextFieldSpecFiller.
//
//	     If the member variable data values encapsulated by this
//	     'txtFieldFiller' are found to be invalid, this method will
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
//	TextFieldSpecFiller
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'txtFieldFiller' will be created and
//	     returned in a new instance of TextFieldSpecFiller.
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
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) copyOut(
	txtFieldFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecFiller, error) {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule.copyOut()",
		"")

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	if txtFieldFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return TextFieldSpecFiller{}, err
	}

	_,
		err =
		textFieldSpecFillerAtom{}.ptr().
			isValidTextFieldSpecFiller(
				txtFieldFiller,
				ePrefix.XCpy("txtFieldFiller validation - "))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	newTxtFieldFiller := TextFieldSpecFiller{}

	// Set zero length array to nil = true
	err = strMechPreon{}.ptr().copyRuneArrays(
		&newTxtFieldFiller.fillerCharacters,
		&txtFieldFiller.fillerCharacters,
		true,
		ePrefix.XCpy(
			"newTxtFieldFiller.fillerCharacters=Target "+
				"<-txtFieldFiller.fillerCharacters=Source"))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	newTxtFieldFiller.fillerCharsRepeatCount =
		txtFieldFiller.fillerCharsRepeatCount

	newTxtFieldFiller.textLineReader = nil

	newTxtFieldFiller.lock = new(sync.Mutex)

	return newTxtFieldFiller, nil
}

// empty - Receives a pointer to an instance of TextFieldSpecLabel
// and proceeds to set all internal member variables to
// their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'txtFieldLabel' will be overwritten and deleted.
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) empty(
	txtFieldFiller *TextFieldSpecFiller) {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	if txtFieldFiller == nil {
		return
	}

	txtFieldFiller.fillerCharacters = nil

	txtFieldFiller.fillerCharsRepeatCount = 0

	txtFieldFiller.textLineReader = nil

	return
}

// equal - Receives a pointer to two instances of
// TextFieldSpecFiller and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) equal(
	txtFieldFiller *TextFieldSpecFiller,
	incomingTxtFieldFiller *TextFieldSpecFiller) bool {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	if txtFieldFiller == nil ||
		incomingTxtFieldFiller == nil {
		return false
	}

	if txtFieldFiller.fillerCharsRepeatCount !=
		incomingTxtFieldFiller.fillerCharsRepeatCount {
		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtFieldFiller.fillerCharacters,
		incomingTxtFieldFiller.fillerCharacters) {
		return false
	}

	return true
}

// getFormattedText - Returns the formatted text generated by the
// instance of TextFieldSpecFiller passed as an input parameter.
//
// If the length of the Filler Characters array (fillerCharacters
// []rune) is zero, this method returns an empty string.
//
// If the Filler Characters Repeat Count is less than one (+1),
// this method returns an empty string.
//
// If the Filler Characters Repeat Count is greater than
// one-million (+1,000,000), this method returns an empty string.
//
// The length of the final formatted text string is the product of:
//
//	Filler Characters Array Length  X
//	             Filler Characters Repeat Count
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) getFormattedText(
	txtFieldFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	formattedText = ""

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule.getFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	if txtFieldFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return formattedText, err
	}

	_,
		err =
		textFieldSpecFillerAtom{}.ptr().
			isValidTextFieldSpecFiller(
				txtFieldFiller,
				ePrefix.XCpy("txtFieldFiller validation - "))

	if err != nil {
		return formattedText, err
	}

	for i := 0; i < txtFieldFiller.fillerCharsRepeatCount; i++ {
		formattedText += string(txtFieldFiller.fillerCharacters)
	}

	return formattedText, err
}

// newEmpty - Returns a new unpopulated instance of
// TextFieldSpecFiller. All the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the TextFieldSpecFiller is invalid.
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) newEmpty() TextFieldSpecFiller {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	newFillerField := TextFieldSpecFiller{}

	newFillerField.fillerCharacters = nil

	newFillerField.fillerCharsRepeatCount = 0

	newFillerField.textLineReader = nil

	newFillerField.lock = new(sync.Mutex)

	return newFillerField
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecFillerMolecule.
func (txtFieldFillerMolecule textFieldSpecFillerMolecule) ptr() *textFieldSpecFillerMolecule {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	return &textFieldSpecFillerMolecule{
		lock: new(sync.Mutex),
	}
}
