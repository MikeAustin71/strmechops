package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type negNumSignSearchNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingNegNumSearchSpec' to input parameter
// 'targetNegNumSearchSpec'. Both instances are of type
// NegativeNumberSearchSpec.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'targetNegNumSearchSpec' will be
// overwritten.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetNegNumSearchSpec        *NegativeNumberSearchSpec
//	   - A pointer to a NegativeNumberSearchSpec instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'incomingNegNumSearchSpec'.
//
//	     'targetNegNumSearchSpec' is the target of this copy
//	     operation.
//
//
//	incomingNegNumSearchSpec      *NegativeNumberSearchSpec
//	   - A pointer to another NegativeNumberSearchSpec instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'targetNegNumSearchSpec'.
//
//	     'incomingNegNumSearchSpec' is the source for this copy
//	     operation.
//
//	     If 'incomingNegNumSearchSpec' is determined to be invalid,
//	     an error will be returned.
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
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (negNumSearchNanobot *negNumSignSearchNanobot) copyIn(
	targetNegNumSearchSpec *NegativeNumberSearchSpec,
	incomingNegNumSearchSpec *NegativeNumberSearchSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetNegNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNegNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	negNumSearchAtom := negNumSearchSpecAtom{}

	_,
		err2 =
		negNumSearchAtom.
			testValidityOfNegNumSearchSpec(
				incomingNegNumSearchSpec,
				nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingNegNumSearchSpec' failed!\n"+
			"This instance of NegativeNumberSearchSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	// Reset all targetNegNumSearchSpec member
	//  variables to their zero values
	negNumSearchAtom.
		empty(targetNegNumSearchSpec)

	targetNegNumSearchSpec.negNumSignPosition =
		incomingNegNumSearchSpec.negNumSignPosition

	var lenLeadingNegNumSignSymbols,
		lenTrailingNegNumSignSymbols int

	if targetNegNumSearchSpec.negNumSignPosition ==
		NumSignSymPos.Before() {

		lenLeadingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray)

		targetNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i] =
				incomingNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i]
		}

		targetNegNumSearchSpec.foundLeadingNegNumSign =
			incomingNegNumSearchSpec.foundLeadingNegNumSign

		targetNegNumSearchSpec.foundLeadingNegNumSignIndex =
			incomingNegNumSearchSpec.foundLeadingNegNumSignIndex

	} else if targetNegNumSearchSpec.negNumSignPosition ==
		NumSignSymPos.BeforeAndAfter() {

		lenTrailingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray)

		targetNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i] =
				incomingNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i]
		}

		targetNegNumSearchSpec.foundTrailingNegNumSign =
			incomingNegNumSearchSpec.foundTrailingNegNumSign

		targetNegNumSearchSpec.foundTrailingNegNumSignIndex =
			incomingNegNumSearchSpec.foundTrailingNegNumSignIndex

	} else {
		// Must be targetNegNumSearchSpec.negNumSignPosition ==
		//            NumSignSymPos.After()

		// Leading data elements
		lenLeadingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray)

		targetNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {

			targetNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i] =
				incomingNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i]
		}

		targetNegNumSearchSpec.foundLeadingNegNumSign =
			incomingNegNumSearchSpec.foundLeadingNegNumSign

		targetNegNumSearchSpec.foundLeadingNegNumSignIndex =
			incomingNegNumSearchSpec.foundLeadingNegNumSignIndex

		// Trailing Data Elements
		lenTrailingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray)

		targetNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i] =
				incomingNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i]
		}

		targetNegNumSearchSpec.foundTrailingNegNumSign =
			incomingNegNumSearchSpec.foundTrailingNegNumSign

		targetNegNumSearchSpec.foundTrailingNegNumSignIndex =
			incomingNegNumSearchSpec.foundTrailingNegNumSignIndex

	}

	targetNegNumSearchSpec.foundFirstNumericDigitInNumStr =
		incomingNegNumSearchSpec.foundFirstNumericDigitInNumStr

	targetNegNumSearchSpec.foundNegNumSignSymbols =
		incomingNegNumSearchSpec.foundNegNumSignSymbols

	err =
		targetNegNumSearchSpec.leadingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"targetNegNumSearchSpec.leadingNegNumSignSymbols"))

	if err != nil {

		return err

	}

	err =
		targetNegNumSearchSpec.trailingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"targetNegNumSearchSpec.trailingNegNumSignSymbols"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'negNumSearchSpec'. a pointer to an instance of
// NegativeNumberSearchSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The input parameter 'negNumSearchSpec' is determined to be
// invalid, this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	negNumSearchSpec           *NegativeNumberSearchSpec
//	   - A pointer to an instance of NegativeNumberSearchSpec. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of NegativeNumberSearchSpec.
//
//	     If the member variable data values encapsulated by
//	     'negNumSearchSpec' are found to be invalid, this method will
//	     return an error
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
//	copyOfNegNumSearchSpec     NegativeNumberSearchSpec
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'negNumSearchSpec' will be created and returned
//	     in a new instance of NegativeNumberSearchSpec.
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
func (negNumSearchNanobot *negNumSignSearchNanobot) copyOut(
	negNumSearchSpec *NegativeNumberSearchSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	copyOfNegNumSearchSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return copyOfNegNumSearchSpec, err

	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return copyOfNegNumSearchSpec, err
	}

	var err2 error

	_,
		err2 =
		negNumSearchSpecAtom{}.ptr().
			testValidityOfNegNumSearchSpec(
				negNumSearchSpec,
				nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'negNumSearchSpec' failed!\n"+
			"This instance of NegativeNumberSearchSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return copyOfNegNumSearchSpec, err
	}

	copyOfNegNumSearchSpec.negNumSignPosition =
		negNumSearchSpec.negNumSignPosition

	var lenLeadingNegNumSignSymbols,
		lenTrailingNegNumSignSymbols int

	if copyOfNegNumSearchSpec.negNumSignPosition ==
		NumSignSymPos.Before() {

		lenLeadingNegNumSignSymbols =
			len(negNumSearchSpec.leadingNegNumSignSymbols.CharsArray)

		copyOfNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i] =
				negNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i]
		}

		copyOfNegNumSearchSpec.foundLeadingNegNumSign =
			negNumSearchSpec.foundLeadingNegNumSign

		copyOfNegNumSearchSpec.foundLeadingNegNumSignIndex =
			negNumSearchSpec.foundLeadingNegNumSignIndex

	} else if copyOfNegNumSearchSpec.negNumSignPosition ==
		NumSignSymPos.BeforeAndAfter() {

		lenTrailingNegNumSignSymbols =
			len(negNumSearchSpec.trailingNegNumSignSymbols.CharsArray)

		copyOfNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i] =
				negNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i]
		}

		copyOfNegNumSearchSpec.foundTrailingNegNumSign =
			negNumSearchSpec.foundTrailingNegNumSign

		copyOfNegNumSearchSpec.foundTrailingNegNumSignIndex =
			negNumSearchSpec.foundTrailingNegNumSignIndex

	} else {
		// Must be copyOfNegNumSearchSpec.negNumSignPosition ==
		//            NumSignSymPos.After()

		// Leading data elements
		lenLeadingNegNumSignSymbols =
			len(negNumSearchSpec.leadingNegNumSignSymbols.CharsArray)

		copyOfNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i] =
				negNumSearchSpec.leadingNegNumSignSymbols.CharsArray[i]
		}

		copyOfNegNumSearchSpec.foundLeadingNegNumSign =
			negNumSearchSpec.foundLeadingNegNumSign

		copyOfNegNumSearchSpec.foundLeadingNegNumSignIndex =
			negNumSearchSpec.foundLeadingNegNumSignIndex

		// Trailing Data Elements
		lenTrailingNegNumSignSymbols =
			len(negNumSearchSpec.trailingNegNumSignSymbols.CharsArray)

		copyOfNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i] =
				negNumSearchSpec.trailingNegNumSignSymbols.CharsArray[i]
		}

		copyOfNegNumSearchSpec.foundTrailingNegNumSign =
			negNumSearchSpec.foundTrailingNegNumSign

		copyOfNegNumSearchSpec.foundTrailingNegNumSignIndex =
			negNumSearchSpec.foundTrailingNegNumSignIndex

	}

	copyOfNegNumSearchSpec.foundFirstNumericDigitInNumStr =
		negNumSearchSpec.foundFirstNumericDigitInNumStr

	copyOfNegNumSearchSpec.foundNegNumSignSymbols =
		negNumSearchSpec.foundNegNumSignSymbols

	err =
		copyOfNegNumSearchSpec.leadingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"copyOfNegNumSearchSpec.leadingNegNumSignSymbols"))

	if err != nil {

		return copyOfNegNumSearchSpec, err

	}

	err =
		copyOfNegNumSearchSpec.trailingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"copyOfNegNumSearchSpec.trailingNegNumSignSymbols"))

	return copyOfNegNumSearchSpec, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable names and corresponding values
// contained in the 'negNumSearchSpec' instance of
// NegativeNumberSearchSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	negNumSearchSpec       *NegativeNumberSearchSpec
//	   - A pointer to an instance of NegativeNumberSearchSpec.
//	     Formatted text output will be generated listing the member
//	     variable names and their corresponding values. The
//	     formatted text can then be used for screen displays, file
//	     output or printing.
//
//	     No data validation is performed on this instance of
//	     NegativeNumberSearchSpec.
//
//
//	tagDescription         string
//	   - An optional string containing a tag or text description
//	     which will be included in the formatted text output
//	     returned by this method.
//
//	     If this parameter is submitted as an empty string, no
//	     text description will be applied to the formatted text
//	     output and no error will be generated.
//
//
//	errPrefDto             *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	strings.Builder
//	   - If this method completes successfully, an instance of
//	     strings.Builder will be returned. This instance contains
//	     the formatted text output listing the member variable
//	     names and their corresponding values for input parameter
//	     'negNumSearchSpec' . This formatted text can them be used
//	     for text displays, file output or printing.
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
func (negNumSearchNanobot *negNumSignSearchNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	negNumSearchSpec *NegativeNumberSearchSpec,
	tagDescription string,
	errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return err

	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		256 - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 16)
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Total available Length of Output Line
	const maxLineLen = 70

	// Max Label Field Length = 30
	const maxLabelFieldLen = 31

	txtFormatCol := TextFormatterCollection{}

	solidLineDto := TextLineSolidDto{
		FormatType:                 TxtFieldType.SolidLine(),
		LeftMarginStr:              " ",
		SolidLineChars:             "=",
		SolidLineCharRepeatCount:   maxLineLen - 2,
		RightMarginStr:             " ",
		TurnLineTerminationOff:     false,
		LineTerminator:             "",
		MaxLineLength:              -1,
		TurnAutoLineLengthBreaksOn: false,
		lock:                       nil,
	}

	err = txtFormatCol.SetStdFormatParamsLine1Col(
		"",
		maxLineLen,
		TxtJustify.Center(),
		"",
		false,
		"",
		-1,
		false,
		ePrefix.XCpy(
			"1-Column Setup"))

	if err != nil {
		return err
	}

	// Leading Title Marquee

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"NegativeNumberSearchSpec",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	if len(tagDescription) > 0 {

		// Title Line 1
		err = txtFormatCol.AddLine1Col(
			"tagDescription",
			ePrefix.XCpy(
				"Top-Title Line 2"))

		if err != nil {
			return err
		}

	}

	// Title Line 3
	err = txtFormatCol.AddLine1Col(
		"Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return err
	}

	// Title Line  4 Date/Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 4"))

	if err != nil {
		return err
	}

	// Filler Line '========='
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// End Of Marquee

	// Begin Label Parameter Pairs

	colonSpace := ": "

	// Set up 2-Column Parameters
	err = txtFormatCol.SetStdFormatParamsLine2Col(
		" ",
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		-1,
		TxtJustify.Left(),
		"",
		false,
		"",
		maxLineLen,
		true,
		ePrefix.XCpy(
			"Set 2-Column Params"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.negNumSignPosition
	//  - negNumSignPosition NumSignSymbolPosition

	txtStrLabel := "Negative Number Sign Position"

	if !negNumSearchSpec.negNumSignPosition.XIsValid() {
		negNumSearchSpec.negNumSignPosition =
			NumSignSymPos.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.negNumSignPosition,
		ePrefix.XCpy(
			"negNumSignPosition"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.leadingNegNumSignSymbols
	//  RuneArrayDto

	txtStrLabel = "Leading Negative Number Sign"

	txtStrParam :=
		negNumSearchSpec.leadingNegNumSignSymbols.
			GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "Leading Negative Number Sign is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.trailingNegNumSignSymbols
	//  RuneArrayDto

	txtStrLabel = "Trailing Negative Number Sign"

	txtStrParam =
		negNumSearchSpec.trailingNegNumSignSymbols.
			GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "Trailing Negative Number Sign is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundFirstNumericDigitInNumStr
	//  RuneArrayDto

	txtStrLabel = "foundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"foundFirstNumericDigitInNumStr"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundNegNumSignSymbols
	//  RuneArrayDto

	txtStrLabel = "foundNegNumSignSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundNegNumSignSymbols,
		ePrefix.XCpy(
			"foundNegNumSignSymbols"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundLeadingNegNumSign
	//  RuneArrayDto

	txtStrLabel = "foundLeadingNegNumSign"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundLeadingNegNumSign,
		ePrefix.XCpy(
			"foundLeadingNegNumSign"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundLeadingNegNumSignIndex
	//  RuneArrayDto

	txtStrLabel = "foundLeadingNegNumSignIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundLeadingNegNumSignIndex,
		ePrefix.XCpy(
			"foundLeadingNegNumSignIndex"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundTrailingNegNumSign
	//  RuneArrayDto

	txtStrLabel = "foundTrailingNegNumSign"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundTrailingNegNumSign,
		ePrefix.XCpy(
			"foundTrailingNegNumSign"))

	if err != nil {
		return err
	}

	// Build negNumSearchSpec.foundTrailingNegNumSignIndex
	//  RuneArrayDto

	txtStrLabel = "foundTrailingNegNumSignIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		negNumSearchSpec.foundTrailingNegNumSignIndex,
		ePrefix.XCpy(
			"foundTrailingNegNumSignIndex"))

	if err != nil {
		return err
	}

	// Trailing Title Marquee
	// Top Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"NegativeNumberSearchSpec",
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return err
	}

	if len(tagDescription) > 0 {

		// Title Line 1
		err = txtFormatCol.AddLine1Col(
			"tagDescription",
			ePrefix.XCpy(
				"Bottom-Title Line 2"))

		if err != nil {
			return err
		}

	}

	// Title Line 3
	err = txtFormatCol.AddLine1Col(
		"End Of Parameter Listing",
		ePrefix.XCpy(
			"Bottom-Title Line 3"))

	if err != nil {
		return err
	}

	// Filler =======
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Blank Line
	txtFormatCol.AddLineBlank(
		2,
		"")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// ptr - Returns a pointer to a new instance of
// negNumSignSearchNanobot.
func (negNumSearchNanobot negNumSignSearchNanobot) ptr() *negNumSignSearchNanobot {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	return &negNumSignSearchNanobot{
		lock: new(sync.Mutex),
	}
}

// setLeadingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	negNumSearchSpec           *NegativeNumberSearchSpec
//	   - A pointer to an instance of NegativeNumberSearchSpec. This
//	     instance will be configured as a Leading Negative Number
//	     Sign Specification. All previous configuration data will be
//	     deleted and replaced with a new Leading Negative Number
//	     Sign configuration.
//
//
//	leadingNegNumSignSymbols   []rune
//	   - An array of runes identifying the character or characters
//	     which comprise the Leading Negative Number Symbol used in
//	     configuring the NegativeNumberSearchSpec instance,
//	     'negNumSearchSpec'.
//
//	     If this array is empty (zero length) or includes array
//	     elements containing a zero value, an error will be
//	     returned.
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
// -----------------------------------------------------------------
//
// Return Values
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSignAtom := negNumSearchSpecAtom{}

	negNumSignAtom.empty(
		negNumSearchSpec)

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.leadingNegNumSignSymbols.CharsArray,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	err = negNumSearchSpec.leadingNegNumSignSymbols.SetCharacterSearchType(
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"negNumSearchSpec.leadingNegNumSignSymbols"+
				"<-TextCharSearchType.LinearTargetStartingIndex()"))

	negNumSearchSpec.negNumSignPosition = NumSignSymPos.Before()

	return err
}

// setLeadingAndTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading and Trailing Negative Number Sign Specification.
// All internal member variables are then configured using the
// input parameter 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted and replaced with the new
// configuration specifications.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, in the US and Canada
// parentheses "()" are used to indicate negative numeric
// values. Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	negNumSearchSpec           *NegativeNumberSearchSpec
//	   - A pointer to an instance of NegativeNumberSearchSpec. This
//	     instance will be configured as a Leading and Trailing
//	     Negative Number Sign Specification. All previous
//	     configuration data will be deleted and replaced with a new
//	     Leading and Trailing Negative Number Sign configuration.
//
//
//	leadingNegNumSignSymbols   []rune
//	   - An array of runes identifying the character or characters
//	     which comprise the Leading Negative Number Symbol used in
//	     configuring the NegativeNumberSearchSpec instance,
//	     'negNumSearchSpec'.
//
//	     If this array is empty (zero length) or includes array
//	     elements containing a zero value, an error will be
//	     returned.
//
//
//	trailingNegNumSignSymbols  []rune
//	   - An array of runes identifying the character or characters
//	     which comprise the Trailing Negative Number Symbol used in
//	     configuring the NegativeNumberSearchSpec instance,
//	     'negNumSearchSpec'.
//
//	     If this array is empty (zero length) or includes array
//	     elements containing a zero value, an error will be
//	     returned.
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
// -----------------------------------------------------------------
//
// Return Values
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingAndTrailingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingAndTrailingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.leadingNegNumSignSymbols.CharsArray,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.trailingNegNumSignSymbols.CharsArray,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSearchSpec.negNumSignPosition = NumSignSymPos.BeforeAndAfter()

	err = negNumSearchSpec.leadingNegNumSignSymbols.
		SetCharacterSearchType(
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				"negNumSearchSpec.leadingNegNumSignSymbols"+
					"<-TextCharSearchType.LinearTargetStartingIndex()"))

	if err != nil {
		return err
	}

	err = negNumSearchSpec.trailingNegNumSignSymbols.
		SetCharacterSearchType(
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				"negNumSearchSpec.trailingNegNumSignSymbols"+
					"<-TextCharSearchType.LinearTargetStartingIndex()"))

	return err
}

// setTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Trailing Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	negNumSearchSpec           *NegativeNumberSearchSpec
//	   - A pointer to an instance of NegativeNumberSearchSpec. This
//	     instance will be configured as a Trailing Negative Number
//	     Sign Specification. All previous configuration data will
//	     be deleted and replaced with a new Trailing Negative Number
//	     Sign configuration.
//
//
//	trailingNegNumSignSymbols  []rune
//	   - An array of runes identifying the character or characters
//	     which comprise the Trailing Negative Number Symbol used in
//	     configuring the NegativeNumberSearchSpec instance,
//	     'negNumSearchSpec'.
//
//	     If this array is empty (zero length) or includes array
//	     elements containing a zero value, an error will be
//	     returned.
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
// -----------------------------------------------------------------
//
// Return Values
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setTrailingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setTrailingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSearchSpecAtom{}.ptr().empty(
		negNumSearchSpec)

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.trailingNegNumSignSymbols.CharsArray,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSearchSpec.negNumSignPosition = NumSignSymPos.After()

	err = negNumSearchSpec.trailingNegNumSignSymbols.
		SetCharacterSearchType(
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				"negNumSearchSpec.trailingNegNumSignSymbols"+
					"<-TextCharSearchType.LinearTargetStartingIndex()"))

	return err
}
