package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type strMechMolecule struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// strMechMolecule.
//
func (sMechMolecule strMechMolecule) ptr() *strMechMolecule {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	return &strMechMolecule{
		lock: new(sync.Mutex),
	}
}

// extractNumRunes - Receives an array of runes and extracts the
// numeric digits as text characters.
func (sMechMolecule *strMechMolecule) extractNumRunes(
	targetSearchString RuneArrayDto,
	targetSearchStringName string,
	startingSearchIndex int,
	characterSearchLength int,
	negativeNumSearchSpecsCol NegNumSearchSpecCollection,
	decimalSeparatorSpec DecimalSeparatorSpec,
	numParsingTerminators RuneArrayCollection,
	requestRemainderRunesString bool,
	ePrefDto *ePref.ErrPrefixDto) (
	searchResults CharSearchNumStrParseResultsDto,
	numStrKernel NumberStrKernel,
	err error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"strMechMolecule.extractNumRunes()",
		"")

	if err != nil {

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms := CharSearchTargetInputParametersDto{}.New()

	targetInputParms.TargetString = &targetSearchString

	targetInputParms.TargetStringLength = targetSearchString.GetRuneArrayLength()

	if targetInputParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSearchString' is invalid.\n"+
			"'targetSearchString' has an array length of zero!\n",
			ePrefix.String())

		return searchResults,
			numStrKernel,
			err

	}

	sMechPreon := strMechPreon{}
	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		targetInputParms.TargetString.CharsArray,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' should contain valid characters.\n"+
			"A validity test on this rune array produced the following error:\n"+
			"%v\n",
			ePrefix.String(),
			targetSearchStringName,
			targetSearchStringName,
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetInputParametersName = "extractNumRunes"

	if len(targetSearchStringName) == 0 {
		targetInputParms.TargetStringName = "targetSearchString"
	} else {
		targetInputParms.TargetStringName = targetSearchStringName
	}

	targetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringName + "Length"

	targetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringName + "StartingSearchIndex"

	targetInputParms.TargetStringSearchLength = -1

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.TargetStringStartingSearchIndex =
		startingSearchIndex

	if targetInputParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value less than zero!\n"+
			"startingSearchIndex = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringStartingSearchIndex)

		return searchResults,
			numStrKernel,
			err
	}

	if targetInputParms.TargetStringStartingSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value greater than the last index\n"+
			"of %v!\n"+
			"startingSearchIndex = '%v'\n"+
			"%v last index = %v\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringStartingSearchIndex,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringLength-1)

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetStringSearchLength = characterSearchLength

	if targetInputParms.TargetStringSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'characterSearchLength' is invalid.\n"+
			"'characterSearchLength' has a value of zero.\n",
			ePrefix.String())

		return searchResults,
			numStrKernel,
			err
	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return searchResults,
			numStrKernel,
			err

	}

	err2 = negativeNumSearchSpecsCol.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeNumSearchSpecsCol' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"negativeNumSearchSpecsCol"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	err2 = decimalSeparatorSpec.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorSpec' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"decimalSeparatorSpec"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.RequestRemainderString =
		requestRemainderRunesString

	searchResults.Empty()

	searchResults.SearchResultsName = "Number Runes Extraction Results"

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	searchResults.FoundNumericDigits = false
	searchResults.FoundNonZeroValue = false
	searchResults.FoundDecimalSeparatorSymbols = false
	searchResults.FoundDecimalDigits = false
	searchResults.NumSignValue = NumSignVal.Zero()
	searchResults.NumValueType = NumValType.None()
	searchResults.RemainderString.Empty()
	searchResults.DecimalSeparatorSearchResults.Empty()
	searchResults.NegativeNumberSymbolSearchResults.Empty()
	searchResults.ParsingTerminatorSearchResults.Empty()

	var negNumSearchResults CharSearchResultsDto
	var decimalSepSearchResults CharSearchDecimalSeparatorResultsDto
	var parsingTerminationResults CharSearchResultsDto

	for i := targetInputParms.TargetStringStartingSearchIndex; i < targetInputParms.TargetStringAdjustedSearchLength; i++ {

		targetInputParms.TargetStringCurrentSearchIndex = i

		if targetSearchString.CharsArray[i] >= '0' &&
			targetSearchString.CharsArray[i] <= '9' {

			targetInputParms.FoundFirstNumericDigitInNumStr = true

			if targetSearchString.CharsArray[i] > '0' {

				targetInputParms.FoundNonZeroValue = true

				if searchResults.NumSignValue == NumSignVal.Zero() {
					searchResults.NumSignValue = NumSignVal.Positive()
				}
			}

			if !searchResults.DecimalSeparatorSearchResults.
				FoundDecimalSeparatorSymbols {

				err = numStrKernel.AddIntegerDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}

			} else {

				err = numStrKernel.AddFractionalDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}
			}

			continue
		}

		// Check for Parsing Terminators
		// All Parsing Operations Cease if Delimiter is Found
		//if !numParsingTerminators.IsNOP() &&
		//	targetInputParms.FoundFirstNumericDigitInNumStr {
		//
		//	parsingTerminationResults,
		//		err = numParsingTerminators.SearchForTextCharacters(
		//		targetInputParms,
		//		ePrefix.XCpy(
		//			"numParsingTerminators"))
		//
		//	if parsingTerminationResults.FoundSearchTarget {
		//
		//		i = parsingTerminationResults.
		//			TargetStringLastSearchIndex
		//
		//		if i+1 < targetInputParms.TargetStringLength {
		//
		//			targetInputParms.
		//				TargetStringNextSearchIndex = i + 1
		//
		//		} else {
		//
		//			targetInputParms.
		//				TargetStringNextSearchIndex = -1
		//
		//		}
		//
		//		goto computeExitStats
		//	}
		//}

		// Check for Negative Number Sign Symbol
		if !searchResults.NegativeNumberSymbolSearchResults.
			FoundNegativeNumberSymbols {

			searchResults.NegativeNumberSymbolSearchResults,
				err = negativeNumSearchSpecsCol.
				SearchForNegNumSignSymbols(
					targetInputParms,
					ePrefix)

			if searchResults.NegativeNumberSymbolSearchResults.
				FoundNegativeNumberSymbols {

				i =
					negNumSearchResults.TargetStringLastFoundIndex

				searchResults.NumSignValue = NumSignVal.Negative()

				continue
			}
		}

		// Check for Decimal Separators
		//if !foundDecimalSeparators {
		//
		//	decimalSepSearchResults,
		//		err = decimalSeparatorSpec.SearchForDecimalSeparator(
		//		targetInputParms,
		//		ePrefix.XCpy(
		//			"targetInputParms"))
		//
		//	if err != nil {
		//
		//		return searchResults,
		//			numStrKernel,
		//			err
		//
		//	}
		//
		//	if decimalSepSearchResults.FoundDecimalSeparatorSymbols {
		//
		//		foundDecimalSeparators = true
		//
		//		i = decimalSepSearchResults.TargetStringLastSearchIndex
		//
		//		searchResults.TargetStringLastSearchIndex = i
		//
		//		continue
		//	}
		//}
	}

	searchResults.TargetStringNextSearchIndex =
		searchResults.TargetStringLastSearchIndex + 1

	if searchResults.TargetStringNextSearchIndex >=
		targetInputParms.TargetStringLength {

		searchResults.TargetStringNextSearchIndex = -1
	}

computeExitStats:

	// Slice Examples
	//arr := []int{1,2,3,4,5}
	//
	//fmt.Println(arr[:2])        // [1,2]
	//
	//fmt.Println(arr[2:])        // [3,4,5]
	//
	//fmt.Println(arr[2:3])        // [3]
	//
	//fmt.Println(arr[:])            // [1,2,3,4,5]

	searchResults.LoadTargetBaseInputParameters(
		targetInputParms)

	searchResults.RemainderString.Empty()

	if searchResults.FoundSearchTarget == true {

		if requestRemainderRunesString {

			if searchResults.TargetStringNextSearchIndex == -1 {
				// All characters have been searched.
				// There is no Remainder String
			} else {
				// All characters HAVE NOT BEEN SEARCHED
				// There is a Remainder String

				searchResults.RemainderString.CharsArray =
					append(
						searchResults.RemainderString.CharsArray,
						targetSearchString.
							CharsArray[searchResults.TargetStringLastSearchIndex:]...)
			}

		}

		numStrKernel.RationalizeFractionalIntegerDigits()

		if searchResults.FoundNonZeroValue == true {
			// Value is Nonzero

			if searchResults.NumSignValue == NumSignVal.Negative() {
				numStrKernel.numberSign = NumSignVal.Negative()
			} else {
				numStrKernel.numberSign = NumSignVal.Positive()
			}

		} else {
			// Value is Zero
			numStrKernel.numberSign = NumSignVal.Zero()
		}

	} else {
		// searchResults.FoundSearchTarget == false
		// Didn't find any numeric digits!
		numStrKernel.Empty()
		searchResults.TargetStringLastSearchIndex = -1
		searchResults.TargetStringLastFoundIndex = -1
		searchResults.TargetStringNextSearchIndex = 0

		if requestRemainderRunesString {

			searchResults.RemainderString.CharsArray =
				append(
					searchResults.RemainderString.CharsArray,
					targetSearchString.
						CharsArray[searchResults.TargetStringLastSearchIndex:]...)

			if startingSearchIndex == 0 {

				searchResults.RemainderString.CharsArray =
					make([]rune, len(targetSearchString.
						CharsArray))

				copy(
					searchResults.RemainderString.CharsArray,
					targetSearchString.CharsArray)

			} else {

				startingSearchIndex--

				searchResults.RemainderString.CharsArray =
					append(
						searchResults.RemainderString.CharsArray,
						targetSearchString.
							CharsArray[startingSearchIndex:]...)
			}
		}

	}

	return searchResults,
		numStrKernel,
		err
}

// strCenterInStr - returns a string which includes a left pad blank string plus
// the original string ('strToCenter'), plus a right pad blank string.
//
// The returned string will effectively center the original string ('strToCenter')
// in a field of specified length ('fieldLen').
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//     - This string will be centered in a text field. The text
//       field length is defined by input parameter, 'fieldLen'.
//
//
//  fieldLen            int
//     - Defines the length of a text field in which 'strToCenter'
//       will be centered.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - This returned string contains 'strToCenter' with the
//       necessary left-pad and right-pad number of spaces
//       required for centering. The total length of this string
//       will be equal to input parameter, 'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_StrCenterInStr_02() "
//  strToCenter := "Hello"
//  fieldLen := 15
//
//  su := StrMech{}
//  centeredStr, err := su.strCenterInStr(
//  strToCenter,
//  fieldLen,
//  ePrefix)
//
//  ---------------------------------------------
//                               123456789012345
//  centeredStr is now equal to "     Hello     "
//  'Hello' is centered in a field of length 15
//  with left and right pad of 5-spaces.
//
func (sMechMolecule *strMechMolecule) strCenterInStr(
	strToCenter string,
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	ePrefix,
		err := ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechMolecule.strCenterInStr()",
		"")

	if err != nil {
		return "", err
	}

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White "+
				"Space or an EMPTY String!\n",
				ePrefix.String())
	}

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: 'fieldLen' = '%v' strToCenter Length= '%v'.\n"+
				"'fieldLen' is shorter than 'strToCenter' Length!\n",
				ePrefix.String(),
				fieldLen,
				sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen - sLen) / 2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}

	return leftPadStr + strToCenter + rightPadStr, nil
}

// strLeftJustify - Creates a new string, left-justified, within
// a wider text field or output string. The text to be left
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the
// output string and spaces are padded to the right in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//     - The text content which will be left-justified in the
//       output string returned by this method.
//
//  fieldLen            int
//     - Defines the length of the output string in which input
//       parameter 'strToJustify' will be left-justified.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - The output string resulting from the 'left-justify'
//       operation. Input parameter, 'strToJustify' will be
//       left-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrLeftJustify_01() "
//  fieldLen = 15
//  strToJustify    = "Hello World"
//  su := StrMech{}
//  justifiedStr, err := su.strLeftJustify(
//                           strToJustify,
//                           fieldLen,
//                           ePrefix)
//
//  ------------------------------------------------
//                                  123456789012345
//  'justifiedStr' is now equal to "Hello World    "
//  The string length of 'justifiedStr' is 15
//
//
func (sMechMolecule *strMechMolecule) strLeftJustify(
	strToJustify string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechMolecule." +
			"strLeftJustify()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Length of string to left justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				ePrefix.String(),
				strLen, fieldLen)
	}

	rightPadLen := fieldLen - strLen

	rightPadStr := strings.Repeat(" ", rightPadLen)

	return strToJustify + rightPadStr, nil
}

// strPadLeftToCenter - Returns a blank string which allows
// centering of the target string in a fixed length field. A blank
// string is defined here as a string consisting entirely of white
// space characters (' ' 0x20). This string can be prefixed or
// added to the beginning of a text string to achieve a centering
// effect when displayed.
//
// Therefore, the blank string returned by this method only
// constitutes the left-padding of white space necessary to center
// the input parameter, 'strToCenter'. It does not include the
// right padding of white space.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//    - The content or text string which will be centered.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be centered.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - The output string resulting from the string centering
//       operation. This string will consist entirely of white
//       space (' ' 0x20 characters). It will represent the left
//       padding necessary to center the text string,
//       'strToCenter'. See 'Example Usage' section below.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// Assume that total field length ('fieldlen') is 70. Further,
// assume that the string to Center ('strToCenter') is
// 10-characters in length. In order to center a 10-character
// string in a 70-character field, 30-space characters would need
// to be positioned on each side of the string to center. This
// method only returns the left margin, or in this example, a
// string consisting of 30-spaces.
//
//
//  ePrefix := "TestStrOps_StrPadLeftToCenter_02() "
//  //              12345
//  strToCenter := "Hello"
//  fieldLen := 15
//  su := StrMech{}
//  padStr, err := su.StrPadLeftToCenter(
//                   strToCenter,
//                   fieldLen,
//                   ePrefix)
//
//  -------------------------------------------------------------
//                          12345
//  'padStr' is now equal to "     "
//  'padStr' consists of 5-spaces.
//  padStr + strToCenter will yield a centered string.
//
//
func (sMechMolecule *strMechMolecule) strPadLeftToCenter(
	strToCenter string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechMolecule." +
			"strPadLeftToCenter()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	sLen := sOpsQuark.getRuneCountInStr(strToCenter)

	if sLen > fieldLen {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input Parameter String To Center "+
				"('strToCenter')\n"+
				"is longer than Field Length.\n",
				ePrefix.String())
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// strRightJustify - Creates a new string, right-justified, within
// a wider text field or output string. The text to be right
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the right side of
// the output string and spaces are padded to the left in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//    - The content or text string which will be right justified.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be right-justified.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - The output string resulting from the 'right-justify'
//       operation. Input parameter, 'strToJustify' will be
//       right-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  If the total field length ('fieldLen') is specified as
//  10-characters and the length of string to justify
//  ('strToJustify') is 5-characters, then this method would return
//  a string consisting of 5-space characters plus the
//  'strToJustify'.
//
//  ePrefix := "TestStrOps_StrRightJustify_05() "
//  strToJustify := "12345"
//  fieldLen := 10
//
//  su := StrMech{}
//  strRightJustified, err :=
//   su.StrRightJustify(
//               strToJustify,
//               fieldLen,
//               ePrefix)
//
//  --------------------------------------------------------
//                                       1234567890
//  'strRightJustified' is now equal to "     12345"
//  The string length of 'strRightJustified' is 10
//
//
func (sMechMolecule *strMechMolecule) strRightJustify(
	strToJustify string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"sMechMolecule." +
			"strRightJustify()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("%s\n"+
				"Error: Length of string to right justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				ePrefix.String(),
				strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)

	return leftPadStr + strToJustify, nil
}
