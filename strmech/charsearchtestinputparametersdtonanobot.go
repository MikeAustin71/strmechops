package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchTestInputParametersDtoNanobot - Provides helper methods for type
// CharSearchTestInputParametersDto.
//
type charSearchTestInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTestInputParms' to input parameter
// 'destinationTestInputParms'. Both instances are of type
// CharSearchTestInputParametersDto.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTestInputParms'
// will be overwritten.
//
// Also, NO validation is performed on 'sourceTestInputParms'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTestInputParms     *CharSearchTestInputParametersDto
//     - A pointer to a CharSearchTestInputParametersDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTestInputParms'.
//
//       'destinationTestInputParms' is the destination for this
//       copy operation.
//
//
//  sourceTestInputParms          *CharSearchTestInputParametersDto
//     - A pointer to another CharSearchTestInputParametersDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTestInputParms'.
//
//       'sourceTestInputParms' is the source for this copy
//       operation.
//
//       If 'sourceTestInputParms' is determined to be invalid,
//       an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) copyIn(
	destinationTestInputParms *CharSearchTestInputParametersDto,
	sourceTestInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTestInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTestInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTestInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTestInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchTestInputParametersDtoAtom{}.ptr().empty(
		destinationTestInputParms)

	if sourceTestInputParms.TestString != nil {

		err = destinationTestInputParms.TestString.CopyIn(
			sourceTestInputParms.TestString,
			ePrefix.XCpy("destinationTestInputParms<-"+
				"sourceTestInputParms"))

		if err != nil {
			return err
		}
	}

	destinationTestInputParms.TestInputParametersName =
		sourceTestInputParms.TestInputParametersName

	destinationTestInputParms.TestStringName =
		sourceTestInputParms.TestStringName

	destinationTestInputParms.TestStringLength =
		sourceTestInputParms.TestStringLength

	destinationTestInputParms.TestStringLengthName =
		sourceTestInputParms.TestStringLengthName

	destinationTestInputParms.TestStringStartingIndex =
		sourceTestInputParms.TestStringStartingIndex

	destinationTestInputParms.TestStringStartingIndexName =
		sourceTestInputParms.TestStringStartingIndexName

	destinationTestInputParms.TestStringDescription1 =
		sourceTestInputParms.TestStringDescription1

	destinationTestInputParms.TestStringDescription2 =
		sourceTestInputParms.TestStringDescription2

	destinationTestInputParms.CollectionTestObjIndex =
		sourceTestInputParms.CollectionTestObjIndex

	destinationTestInputParms.NumValueType =
		sourceTestInputParms.NumValueType

	destinationTestInputParms.NumStrFormatType =
		sourceTestInputParms.NumStrFormatType

	destinationTestInputParms.NumSymLocation =
		sourceTestInputParms.NumSymLocation

	destinationTestInputParms.NumSymbolClass =
		sourceTestInputParms.NumSymbolClass

	destinationTestInputParms.NumSignValue =
		sourceTestInputParms.NumSignValue

	destinationTestInputParms.PrimaryNumSignPosition =
		sourceTestInputParms.PrimaryNumSignPosition

	destinationTestInputParms.SecondaryNumSignPosition =
		sourceTestInputParms.SecondaryNumSignPosition

	destinationTestInputParms.TextCharSearchType =
		sourceTestInputParms.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'testInputParms', a pointer to an instance of
// CharSearchTestInputParametersDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
//
// NO validation is performed on 'testInputParms'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms                *CharSearchTestInputParametersDto
//     - A pointer to an instance of CharSearchTestInputParametersDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTestInputParametersDto.
//
//       If the member variable data values encapsulated by
//       'testInputParms' are found to be invalid, this method will
//       return an error
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  deepCopyTestInputParms        CharSearchTestInputParametersDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'testInputParms' will be created and
//       returned in a new instance of
//       CharSearchTestInputParametersDto.
//
//
//  err                             error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) copyOut(
	testInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyTestInputParms CharSearchTestInputParametersDto,
	err error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	charSearchTestInputParametersDtoAtom{}.ptr().empty(
		&deepCopyTestInputParms)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return deepCopyTestInputParms, err

	}

	if testInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'testInputParms' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyTestInputParms, err
	}

	if testInputParms.TestString != nil {

		err = deepCopyTestInputParms.TestString.CopyIn(
			testInputParms.TestString,
			ePrefix.XCpy("deepCopyTestInputParms<-"+
				"testInputParms"))

		if err != nil {
			return deepCopyTestInputParms, err
		}
	}

	deepCopyTestInputParms.TestInputParametersName =
		testInputParms.TestInputParametersName

	deepCopyTestInputParms.TestStringName =
		testInputParms.TestStringName

	deepCopyTestInputParms.TestStringLength =
		testInputParms.TestStringLength

	deepCopyTestInputParms.TestStringLengthName =
		testInputParms.TestStringLengthName

	deepCopyTestInputParms.TestStringStartingIndex =
		testInputParms.TestStringStartingIndex

	deepCopyTestInputParms.TestStringStartingIndexName =
		testInputParms.TestStringStartingIndexName

	deepCopyTestInputParms.TestStringDescription1 =
		testInputParms.TestStringDescription1

	deepCopyTestInputParms.TestStringDescription2 =
		testInputParms.TestStringDescription2

	deepCopyTestInputParms.CollectionTestObjIndex =
		testInputParms.CollectionTestObjIndex

	deepCopyTestInputParms.NumValueType =
		testInputParms.NumValueType

	deepCopyTestInputParms.NumStrFormatType =
		testInputParms.NumStrFormatType

	deepCopyTestInputParms.NumSymLocation =
		testInputParms.NumSymLocation

	deepCopyTestInputParms.NumSymbolClass =
		testInputParms.NumSymbolClass

	deepCopyTestInputParms.NumSignValue =
		testInputParms.NumSignValue

	deepCopyTestInputParms.PrimaryNumSignPosition =
		testInputParms.PrimaryNumSignPosition

	deepCopyTestInputParms.SecondaryNumSignPosition =
		testInputParms.SecondaryNumSignPosition

	deepCopyTestInputParms.TextCharSearchType =
		testInputParms.TextCharSearchType

	return deepCopyTestInputParms, err
}

// getFormattedText - Returns formatted text output detailing the
// member variable values contained in the 'testInputParms'
// instance of charSearchTestInputParametersDtoNanobot.
//
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) getFormattedText(
	testInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(512)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"getFormattedText()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if testInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'testInputParms' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 36
	const maxLabelFieldLen = 24

	txtStrParam :=
		testInputParms.TestInputParametersName

	textSpecNanoBot := textSpecificationNanobot{}
	// Title Marquee

	err = textSpecNanoBot.buildFormattedMarqueeTitle(
		"=",
		"CharSearchTestInputParametersDto",
		TxtJustify.Center(),
		txtStrParam,
		TxtJustify.Center(),
		time.Now(),
		"Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		TxtJustify.Center(),
		maxFieldLen,
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	colonSpace := ": "

	// TestInputParametersName
	txtStrLabel := "TestInputParametersName"

	txtStrParam = testInputParms.TestInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestInputParametersName is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// Build Formatted Test String

	txtStrLabel = "Test String"

	lenTxtStrParam := 0

	if testInputParms.TestString != nil {

		if testInputParms.TestString.GetRuneArrayLength() > 0 {
			txtStrParam =
				testInputParms.TestString.GetCharacterString()
		} else {

			txtStrParam = "Test String is EMPTY!"

		}
	} else {
		// testInputParms.TestString == nil
		txtStrParam = "Test String has a nil pointer (Not Set)!"
	}

	lenTxtStrParam = len(txtStrParam)

	if lenTxtStrParam >= (maxFieldLen - maxLabelFieldLen - 2) {
		// We need two Lines of Text

		err = textSpecNanoBot.buildFormattedSingleTextField(
			"",
			txtStrLabel,
			maxLabelFieldLen,
			TxtJustify.Right(),
			colonSpace,
			"\n",
			&strBuilder,
			ePrefix.XCpy(
				"Test String Label"))

		if err != nil {

			return strBuilder, err
		}

		// Now write test string on second line

		err = textSpecNanoBot.buildFormattedSingleTextField(
			" ",
			txtStrParam,
			-1,
			TxtJustify.Left(),
			"",
			"\n",
			&strBuilder,
			ePrefix.XCpy(
				"Test String Parameter"))

		if err != nil {

			return strBuilder, err
		}

	} else {
		// We need only one line of text
		err = textSpecNanoBot.buildFormattedSingleParameterText(
			"",
			txtStrLabel,
			maxLabelFieldLen,
			TxtJustify.Right(),
			colonSpace,
			txtStrParam,
			-1,
			TxtJustify.Left(),
			"",
			"\n",
			&strBuilder,
			ePrefix)

		if err != nil {

			return strBuilder, err
		}
	}

	// TestStringName

	txtStrLabel = "TestStringName"

	txtStrParam = testInputParms.TestStringName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringName is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringLength

	txtStrLabel = "TestStringLength"

	txtStrParam =
		fmt.Sprintf("%v",
			testInputParms.TestStringLength)

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringLengthName
	txtStrLabel = "TestStringLengthName"

	txtStrParam = testInputParms.TestStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringLengthName is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringStartingIndex

	txtStrLabel = "TestStringStartingIndex"

	txtStrParam =
		fmt.Sprintf("%v",
			testInputParms.TestStringStartingIndex)

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringStartingIndexName
	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam = testInputParms.TestStringStartingIndexName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringStartingIndexName is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringDescription1
	txtStrLabel = "TestStringDescription1"

	txtStrParam = testInputParms.TestStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription1 is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TestStringDescription2
	txtStrLabel = "TestStringDescription2"

	txtStrParam = testInputParms.TestStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription2 is EMPTY!"
	}

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-2,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// CollectionTestObjIndex

	txtStrLabel = "CollectionTestObjIndex"

	txtStrParam =
		fmt.Sprintf("%v",
			testInputParms.CollectionTestObjIndex)

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// NumValueType

	txtStrLabel = "NumValueType"

	if !testInputParms.NumValueType.XIsValid() {
		testInputParms.NumValueType = NumValType.None()
	}

	txtStrParam = testInputParms.NumValueType.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// NumStrFormatType

	txtStrLabel = "NumStrFormatType"

	if !testInputParms.NumStrFormatType.XIsValid() {
		testInputParms.NumStrFormatType = NumStrFmtType.None()
	}

	txtStrParam = testInputParms.NumStrFormatType.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// NumSymLocation

	txtStrLabel = "NumSymLocation"

	if !testInputParms.NumSymLocation.XIsValid() {
		testInputParms.NumSymLocation = NumSymLocation.None()
	}

	txtStrParam = testInputParms.NumSymLocation.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// NumSymbolClass

	txtStrLabel = "NumSymbolClass"

	if !testInputParms.NumSymbolClass.XIsValid() {
		testInputParms.NumSymbolClass = NumSymClass.None()
	}

	txtStrParam = testInputParms.NumSymbolClass.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// NumSignValue

	txtStrLabel = "NumSignValue"

	if !testInputParms.NumSignValue.XIsValid() {
		testInputParms.NumSignValue = NumSignVal.None()
	}

	txtStrParam = testInputParms.NumSignValue.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// SecondaryNumSignPosition

	txtStrLabel = "SecondaryNumSignPosition"

	if !testInputParms.SecondaryNumSignPosition.XIsValid() {
		testInputParms.SecondaryNumSignPosition =
			NumSignSymPos.None()
	}

	txtStrParam = testInputParms.SecondaryNumSignPosition.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	// TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !testInputParms.TextCharSearchType.XIsValid() {
		testInputParms.TextCharSearchType =
			CharSearchType.None()
	}

	txtStrParam = testInputParms.TextCharSearchType.String()

	err = textSpecNanoBot.buildFormattedSingleParameterText(
		"",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		&strBuilder,
		ePrefix)

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// charSearchTestInputParametersDtoNanobot.
//
func (searchTestInputParmsNanobot charSearchTestInputParametersDtoNanobot) ptr() *charSearchTestInputParametersDtoNanobot {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	return &charSearchTestInputParametersDtoNanobot{
		lock: new(sync.Mutex),
	}
}
