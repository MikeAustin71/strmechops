package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// charSearchTargetInputParametersDtoNanobot - Provides helper methods for type
// CharSearchTargetInputParametersDto.
//
type charSearchTargetInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTargetInputParms' to input parameter
// 'destinationTargetInputParms'. Both instances are of type
// CharSearchTargetInputParametersDto.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTargetInputParms'
// will be overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTargetInputParms     *CharSearchTargetInputParametersDto
//     - A pointer to a CharSearchTargetInputParametersDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTargetInputParms'.
//
//       'destinationTargetInputParms' is the destination for this
//       copy operation.
//
//
//  sourceTargetInputParms          *CharSearchTargetInputParametersDto
//     - A pointer to another CharSearchTargetInputParametersDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTargetInputParms'.
//
//       'sourceTargetInputParms' is the source for this copy
//       operation.
//
//       If 'sourceTargetInputParms' is determined to be invalid,
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
func (searchTargetInputParmsNanobot *charSearchTargetInputParametersDtoNanobot) copyIn(
	destinationTargetInputParms *CharSearchTargetInputParametersDto,
	sourceTargetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTargetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTargetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTargetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTargetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	targetInputParmsAtom := charSearchTargetInputParametersDtoAtom{}

	targetInputParmsAtom.empty(
		destinationTargetInputParms)

	if sourceTargetInputParms.TargetString != nil {

		err = destinationTargetInputParms.TargetString.CopyIn(
			sourceTargetInputParms.TargetString,
			ePrefix.XCpy("destinationTargetInputParms<-"+
				"sourceTargetInputParms"))

		if err != nil {
			return err
		}
	}

	destinationTargetInputParms.TargetInputParametersName =
		sourceTargetInputParms.TargetInputParametersName

	destinationTargetInputParms.TargetStringName =
		sourceTargetInputParms.TargetStringName

	destinationTargetInputParms.TargetStringLength =
		sourceTargetInputParms.TargetStringLength

	destinationTargetInputParms.TargetStringLengthName =
		sourceTargetInputParms.TargetStringLengthName

	destinationTargetInputParms.TargetStringStartingSearchIndex =
		sourceTargetInputParms.TargetStringStartingSearchIndex

	destinationTargetInputParms.TargetStringStartingSearchIndexName =
		sourceTargetInputParms.TargetStringStartingSearchIndexName

	destinationTargetInputParms.TargetStringSearchLength =
		sourceTargetInputParms.TargetStringSearchLength

	destinationTargetInputParms.TargetStringSearchLengthName =
		sourceTargetInputParms.TargetStringSearchLengthName

	destinationTargetInputParms.TargetStringAdjustedSearchLength =
		sourceTargetInputParms.TargetStringAdjustedSearchLength

	destinationTargetInputParms.TargetStringDescription1 =
		sourceTargetInputParms.TargetStringDescription1

	destinationTargetInputParms.TargetStringDescription2 =
		sourceTargetInputParms.TargetStringDescription2

	destinationTargetInputParms.FoundFirstNumericDigitInNumStr =
		sourceTargetInputParms.FoundFirstNumericDigitInNumStr

	destinationTargetInputParms.TextCharSearchType =
		sourceTargetInputParms.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'targetInputParms'. a pointer to an instance of
// CharSearchTargetInputParametersDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The input parameter 'targetInputParms' is determined to be invalid,
// this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms                *CharSearchTargetInputParametersDto
//     - A pointer to an instance of CharSearchTargetInputParametersDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTargetInputParametersDto.
//
//       If the member variable data values encapsulated by
//       'targetInputParms' are found to be invalid, this method will
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyTargetInputParms        CharSearchTargetInputParametersDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'targetInputParms' will be created and
//       returned in a new instance of CharSearchTargetInputParametersDto.
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
func (searchTargetInputParmsNanobot *charSearchTargetInputParametersDtoNanobot) copyOut(
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyTargetInputParms CharSearchTargetInputParametersDto,
	err error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	deepCopyTargetInputParms.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyTargetInputParms, err
	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyTargetInputParms, err
	}

	if targetInputParms.TargetString != nil {

		err = deepCopyTargetInputParms.TargetString.CopyIn(
			targetInputParms.TargetString,
			ePrefix.XCpy("deepCopyTargetInputParms<-"+
				"targetInputParms"))

		if err != nil {
			return deepCopyTargetInputParms, err
		}

	}

	deepCopyTargetInputParms.TargetInputParametersName =
		targetInputParms.TargetInputParametersName

	deepCopyTargetInputParms.TargetStringName =
		targetInputParms.TargetStringName

	deepCopyTargetInputParms.TargetStringLength =
		targetInputParms.TargetStringLength

	deepCopyTargetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringLengthName

	deepCopyTargetInputParms.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	deepCopyTargetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringStartingSearchIndexName

	deepCopyTargetInputParms.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	deepCopyTargetInputParms.TargetStringSearchLengthName =
		targetInputParms.TargetStringSearchLengthName

	deepCopyTargetInputParms.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	deepCopyTargetInputParms.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	deepCopyTargetInputParms.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

	deepCopyTargetInputParms.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	deepCopyTargetInputParms.TextCharSearchType =
		targetInputParms.TextCharSearchType

	return deepCopyTargetInputParms, err
}

// getFormattedText - Returns formatted text output detailing the
// member variable values contained in the 'targetInputParms'
// instance of CharSearchTargetInputParametersDto.
//
func (searchTargetInputParmsNanobot charSearchTargetInputParametersDtoNanobot) getFormattedText(
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(512)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"getFormattedText()",
		"")

	if err != nil {

		return strBuilder, err
	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 36
	const maxLabelFieldLen = 36

	targetParmsElectron :=
		charSearchTargetInputParametersDtoElectron{}

	txtStrParam :=
		targetInputParms.TargetInputParametersName

	// Title Marquee

	err =
		targetParmsElectron.buildFormattedMarqueeTitle(
			"=",
			"CharSearchTargetInputParametersDto",
			txtStrParam,
			maxFieldLen,
			&strBuilder,
			ePrefix.XCpy(
				""))

	var colonSpacer *TextFieldSpecLabel

	colonSpacer,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		": ",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"colonSpacer"))

	if err != nil {

		return strBuilder, err
	}

	err =
		targetParmsElectron.buildFormattedTargetString(
			maxLabelFieldLen,
			colonSpacer,
			targetInputParms,
			-1,
			&strBuilder,
			ePrefix.XCpy("Target String"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringName

	txtStrParam = targetInputParms.TargetStringName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringName is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringName",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringName"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringLength

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringLength)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringLength is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringLength",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringLength"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringLengthName

	txtStrParam = targetInputParms.TargetStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringLengthName is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringLengthName",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringLengthName"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringStartingSearchIndex

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringStartingSearchIndex)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringStartingSearchIndex is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringStartingSearchIndex",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringStartingSearchIndex"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringStartingSearchIndexName

	txtStrParam =
		targetInputParms.TargetStringStartingSearchIndexName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringStartingSearchIndexName is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringStartingSearchIndexName",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringStartingSearchIndexName"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringSearchLength

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringSearchLength)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringSearchLength is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringSearchLength",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringSearchLength"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringSearchLengthName

	txtStrParam =
		targetInputParms.TargetStringSearchLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringSearchLengthName is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringSearchLengthName",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringSearchLengthName"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringAdjustedSearchLength

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringAdjustedSearchLength)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringAdjustedSearchLength is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringAdjustedSearchLength",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringAdjustedSearchLength"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringDescription1
	txtStrParam = targetInputParms.TargetStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription1 is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringDescription1",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringDescription1"))

	if err != nil {

		return strBuilder, err
	}

	// TargetStringDescription2
	txtStrParam = targetInputParms.TargetStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription2 is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TargetStringDescription2",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TargetStringDescription2"))

	if err != nil {

		return strBuilder, err
	}

	// FoundFirstNumericDigitInNumStr

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.FoundFirstNumericDigitInNumStr)

	if len(txtStrParam) == 0 {
		txtStrParam = "FoundFirstNumericDigitInNumStr is EMPTY!"
	}

	err = targetParmsElectron.
		buildFormattedParameterText(
			"FoundFirstNumericDigitInNumStr",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"FoundFirstNumericDigitInNumStr"))

	if err != nil {

		return strBuilder, err
	}

	// TextCharSearchType
	searchType := targetInputParms.TextCharSearchType

	if !searchType.XIsValid() {
		searchType = CharSearchType.None()
	}

	txtStrParam =
		fmt.Sprintf("%v",
			searchType.String())

	err = targetParmsElectron.
		buildFormattedParameterText(
			"TextCharSearchType",
			maxLabelFieldLen,
			colonSpacer,
			txtStrParam,
			-1,
			&strBuilder,
			ePrefix.XCpy(
				"TextCharSearchType"))

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelNanobot.
//
func (searchTargetInputParmsNanobot charSearchTargetInputParametersDtoNanobot) ptr() *charSearchTargetInputParametersDtoNanobot {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	return &charSearchTargetInputParametersDtoNanobot{
		lock: new(sync.Mutex),
	}
}
