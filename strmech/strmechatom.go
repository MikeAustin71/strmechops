package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type strMechAtom struct {
	lock *sync.Mutex
}

// breakTextAtLineLength - Breaks string text into lines. Takes a
// string and inserts a line delimiter character (a.k.a 'rune') at
// the specified line length ('lineLength').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be parsed into text lines. If
//       'targetStr' is an empty string an error will be returned.
//        If 'targetStr' consists entirely of white space, this
//        method will return a string consisting of a new-line
//        character and an error value of 'nil'.
//
//
//  lineLength          int
//     - The maximum length of each line.
//
//
//  lineDelimiter       rune
//     - The line delimiter character which will be inserted at the
//       end of a line break.
//
//
//  Note: If the caller specifies a line length of 50, the line delimiter
//  character may be placed in the 51st character position depending upon
//  the word breaks.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string
//       parameter will contain the text with line breaks delimited
//       by the input parameter 'lineDelimiter'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
func (sMechAtom *strMechAtom) breakTextAtLineLength(
	targetStr string,
	lineLength int,
	lineDelimiter rune,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechAtom." +
			"breakTextAtLineLength()")

	targetLen := len(targetStr)

	if targetLen == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetStr' is a ZERO "+
				"LENGTH STRING!\n",
				ePrefix.String())
	}

	if lineLength < 5 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'lineLength' is LESS THAN "+
				"5-CHARACTERS!\n"+
				"lineLength='%v'\n",
				ePrefix.String(),
				lineLength)
	}

	if lineDelimiter == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'lineDelimiter' is "+
				"ZERO VALUE!\n",
				ePrefix.String())
	}

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(targetStr) {
		return "\n", nil
	}

	var err error

	var b strings.Builder
	b.Grow(((targetLen / lineLength) * targetLen) + 50)

	var begIdx, endWrdIdx, actualLastIdx, beginWrdIdx int
	var isAllOneWord, isAllSpaces bool
	sMechElectron := strMechElectron{}

	for begIdx < targetLen && begIdx > -1 {

		// skip spaces at the beginning of the line
		begIdx, err = sMechElectron.findFirstNonSpaceChar(
			targetStr,
			begIdx,
			targetLen-1,
			ePrefix.XCtx(
				fmt.Sprintf(
					"\ntargetStr='%v'\n"+
						"begIdx='%v'\n"+
						"actualLastIdx='%v'\n",
					targetStr,
					begIdx,
					actualLastIdx)))

		if err != nil {
			return "", err
		}

		ePrefix.SetCtxEmpty()

		if begIdx == -1 {

			if b.Len() == 0 {
				b.WriteRune(lineDelimiter)
			}

			break // Exit loop
		}

		if begIdx == targetLen-1 {
			b.WriteByte(targetStr[begIdx])
			b.WriteRune(lineDelimiter)
			break
		}

		actualLastIdx = begIdx + lineLength - 1

		if actualLastIdx >= targetLen {
			actualLastIdx = targetLen - 1
		}

		// Find the last complete word in this string segment
		beginWrdIdx, endWrdIdx, isAllOneWord, isAllSpaces, err =
			sOpsQuark.findLastWord(
				targetStr,
				begIdx,
				actualLastIdx,
				ePrefix.XCtx(
					fmt.Sprintf(
						"\ntargetStr='%v'\n"+
							"begIdx='%v'\n"+
							"actualLastIdx='%v'\n",
						targetStr,
						begIdx,
						actualLastIdx)))

		if err != nil {
			return "", err
		}

		ePrefix.SetCtxEmpty()

		if isAllSpaces {
			// This string segment is all spaces
			// write a line delimiter and continue
			begIdx = actualLastIdx + 1

		} else if isAllOneWord {
			// This string segment is all one word
			// and contains NO spaces.

			if actualLastIdx+1 >= targetLen {
				// If this is end of the main string,
				// just write the remaining segment and
				// exit.
				//
				b.WriteString(targetStr[begIdx:])
				b.WriteRune(lineDelimiter)
				break

			} else if actualLastIdx-begIdx+1 <= lineLength {
				// If this string segment is less than the specified
				// line length, just write the entire line segment.
				// Be careful, we may be at the end of the main
				// string.

				if actualLastIdx+1 >= targetLen {
					// This is the end of the main string,
					// just exit.
					b.WriteString(targetStr[begIdx:])
					b.WriteRune(lineDelimiter)
					break

				} else {

					b.WriteString(targetStr[begIdx : actualLastIdx+1])
					begIdx = actualLastIdx + 1
				}

			} else {
				// Out of options. Nothing left to do but hyphenate
				// the word.
				b.WriteString(targetStr[begIdx : actualLastIdx-1])
				b.WriteRune('-')
				begIdx = actualLastIdx

			}

		} else {
			// The segment is NOT All spaces nor is it all one word.

			if endWrdIdx+1 >= targetLen {
				// Are we at the end of targetStr
				b.WriteString(targetStr[begIdx:])
				b.WriteRune(lineDelimiter)
				break

			} else if targetStr[endWrdIdx+1] != ' ' {
				// This word crosses a line break boundary. Try not to split the word.

				// Find  the end of the last word.
				idx, err :=
					sOpsQuark.findLastNonSpaceChar(
						targetStr,
						begIdx,
						beginWrdIdx-1,
						ePrefix.XCtx(
							fmt.Sprintf(
								"\ntargetStr='%v'\n"+
									"begIdx='%v'\n"+
									"actualLastIdx='%v'\n",
								targetStr,
								begIdx,
								actualLastIdx)))

				if err != nil {
					return "", err
				}

				ePrefix.SetCtxEmpty()

				if idx == -1 {
					begIdx = beginWrdIdx
					// Do not write end of line delimiter
					// Set bigIdx to beginning of word and
					// loop again
					continue

				} else {
					// Success we found the end of the last word.
					b.WriteString(targetStr[begIdx : idx+1])
					begIdx = idx + 1
				}

			} else {
				// The word does not cross a line break boundary.
				// The next character after the last word is a
				// space.

				b.WriteString(targetStr[begIdx : endWrdIdx+1])
				begIdx = endWrdIdx + 1
			}
		}

		b.WriteRune(lineDelimiter)

	}

	return b.String(), nil
}

// CopyIn - Copies string information from input parameter
// 'incomingStrOps' to input parameter 'targetStrOps'.
//
// Be advised that the data fields in 'targetStrOps' will be
// overwritten.
//
func (sMechAtom *strMechAtom) copyIn(
	targetStrMech *StrMech,
	incomingStrMech *StrMech,
	ePrefix *ePref.ErrPrefixDto) (
	err error) {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechAtom." +
			"copyIn()")

	if targetStrMech == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStrMech' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if incomingStrMech == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingStrMech' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetStrMech.stringDataMutex == nil {
		targetStrMech.stringDataMutex = new(sync.Mutex)
	}

	if incomingStrMech.stringDataMutex == nil {
		incomingStrMech.stringDataMutex = new(sync.Mutex)
	}

	targetStrMech.StrIn = incomingStrMech.StrIn
	targetStrMech.StrOut = incomingStrMech.StrOut

	targetStrMech.cntBytesWritten = 0
	targetStrMech.cntBytesRead = 0
	targetStrMech.stringData = incomingStrMech.stringData

	return err
}

// CopyOut - Creates a 'deep' copy of input parameter
// 'strOps', an instance of StrMech.
//
func (sMechAtom *strMechAtom) copyOut(
	strMech *StrMech,
	ePrefix *ePref.ErrPrefixDto) (
	*StrMech,
	error) {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechAtom." +
			"copyOut()")

	var err error

	newStrOps := StrMech{}

	if strMech == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strMech' is a 'nil' pointer!\n",
			ePrefix.String())

		return &newStrOps, err
	}

	newStrOps.StrIn = strMech.StrIn
	newStrOps.StrOut = strMech.StrOut
	newStrOps.stringData = strMech.stringData

	newStrOps.stringDataMutex = new(sync.Mutex)

	return &newStrOps, err
}

// extractDataField - Extracts a data field string from a larger target string ('targetStr').
// The target string is searched for a data field. If the 'leadingKeyWordDelimiter' parameter
// is populated, the data field MUST contain this leading key word, otherwise an empty data field
// is returned.
//
// If 'leadingKeyWordDelimiter' is an empty string, the search for the data field will begin at
// 'targetStr' index, 'startIdx'.
//
// The returned data field must occur in 'targetStr' prior to a comment or End-Of-Line character.
//
// The extracted data field MUST be preceded by one of the characters specified in input
// parameter, 'leadingFieldSeparators'. In addition, the data field must be immediately
// followed by one of the characters in input parameter 'trailingFieldSeparators' or a comment
// or an End-Of-Line character.
//
// ------------------------------------------------------------------------
//
// Input Values
//
//  targetStr               string   - The target string from which the data field will be extracted.
//
//
//  leadingKeyWordDelimiters []string- Data fields are often preceded by field names or field designators.
//                                       The 'leadingKeyWordDelimiters' parameter is a string array
//                                       containing 'Key Word Delimiters'. A Key Word Delimiter may be
//                                       a Key Word string or a character which identifies and immediately
//                                       precedes the data field. If multiple Key Word Delimiters exist
//                                       in 'targetStr' the first instance of a key word in targetStr'
//                                       will be designated as the Key Word Delimiter.
//
//                                       If this parameter is populated, the search for a data field
//                                       will begin immediately after the first located Key Word
//                                       Delimiter string. If none of Key Words in this string array
//                                       are located in 'targetStr', an empty string will be returned
//                                       for data field. If this parameter is populated, at least one
//                                       of the Key Words MUST exist in 'targetStr' before a data field
//                                       will be extracted and returned.
//
//                                       If this parameter is an empty string array, the search for a
//                                       data field will begin at the string index designated by
//                                       parameter, 'startIdx'.
//
//
//  startIdx                int      - The string index in parameter 'targetStr' from which the search for
//                                       a data field will begin. Note that the starting index will be adjusted
//                                       according to the existence of a Key Word Delimiter as explained
//                                       above.
//
//
//  leadingFieldSeparators  []string - An array of characters or groups of characters which delimit the
//                                       leading edge of the data field.
//
//
//  trailingFieldSeparators []string - An array of characters or groups of characters which delimit the
//                                       end of a data field.
//
//
//  commentDelimiters       []string - Comments effectively terminate the search for a data field. This
//                                       array stores comment characters or phrases which signal the beginning
//                                       of a comment.
//
//
//  endOfLineDelimiters     []string - Those characters or groups of characters which mark the end of a line.
//                                       Generally this includes characters like 'new line' or 'carriage return'.
//                                       End of line characters will terminate the search for a data field.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods listed
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
//  DataFieldProfileDto - If successful, this method returns a structure containing
//                        characteristics describing the extracted data field.
//
//    type DataFieldProfileDto struct {
//       TargetStr                      string //  The string from which the data field is extracted.
//       TargetStrLength                int    //  Length of 'TargetStr'
//       TargetStrStartIndex            int    //  The index with in 'TargetStr' from which the search for a data field was initiated.
//       TargetStrLastGoodIndex         int    //  Last valid index in target string which is less than the target string length and is NOT an 'End Of Field' or 'End Of Line' Delimiter.
//       LeadingKeyWordDelimiter        string //  The Leading Key Word Delimiter which is used to identify the beginning of the field search.
//       LeadingKeyWordDelimiterIndex   int    //  Index of the found Leading Key Word Delimiter.
//       DataFieldStr                   string //  The extracted data field string.
//       DataFieldIndex                 int    //  The index in 'TargetStr' where the data field begins.
//       DataFieldLength                int    //  The length of the extracted data field string.
//       DataFieldTrailingDelimiter     string //  The trailing character which marked the end of the data field. A zero value indicates end of string encountered.
//       DataFieldTrailingDelimiterType DataFieldTrailingDelimiterType // A constant or enumeration type used to describe the type of delimiter used to mark the end of a data field.
//       NextTargetStrIndex             int    //  The index in 'TargetStr' immediately following the extracted data field.
//       CommentDelimiter               string //  If a Comment Delimiter is detected it is stored here.
//       CommentDelimiterIndex          int    //  If a Comment Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
//       EndOfLineDelimiter             string //  If an End-Of-Line Delimiter is detected it is captured and stored here.
//       EndOfLineDelimiterIndex        int    //  If an End-Of-Line Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
//     }
//
//
//   error
//     - If the method completes successfully and no errors are encountered
//       this return value is set to 'nil'. Otherwise, if errors are encountered
//       this return value will contain an appropriate error message.
//
//       The most likely source of errors are invalid input parameters.
//       Input parameters 'targetStr', 'startIdx', 'leadingFieldSeparators',
//       'trailingFieldSeparators' and 'endOfStringDelimiters' are required input
//       parameters and must be populated with valid data.
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
//  ePrefix := "TestStrOps_ExtractDataField_01() "
//  endOfLineDelimiters := []string{"\n"}
//  commentDelimiters := []string{"#"}
//  leadingFieldDelimiters := []string{
//  "\t",
//  "\r",
//  "\f",
//  "\v",
//  " "}
//
//  trailingFieldDelimiters := []string{
//  "\t",
//  "\r",
//  "\f",
//  "\v",
//  " "}
//
//  targetStr := " Zone:\t America/Chicago\t Link:\t US/Central\t\n"
//  startIdx := 0
//  leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
//
//  datDto,
//  err :=
//    StrMech{}.Ptr().
//        ExtractDataField(
//           targetStr,
//           leadingKeyWordDelimiters,
//           startIdx,
//           leadingFieldDelimiters,
//           trailingFieldDelimiters,
//           commentDelimiters,
//           endOfLineDelimiters,
//           ePrefix)
//
//  -----------------------------------------------
//  datDto.DataFieldStr is now equal to:
//          "America/Chicago"
//
//
func (sMechAtom *strMechAtom) extractDataField(
	targetStr string,
	leadingKeyWordDelimiters []string,
	startIdx int,
	leadingFieldSeparators []string,
	trailingFieldSeparators []string,
	commentDelimiters []string,
	endOfLineDelimiters []string,
	ePrefix *ePref.ErrPrefixDto) (
	DataFieldProfileDto,
	error) {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechAtom." +
			"extractDataField()")

	newDataDto := DataFieldProfileDto{}.New()
	newDataDto.TargetStr = targetStr
	newDataDto.TargetStrLength = len(targetStr)
	newDataDto.TargetStrStartIndex = startIdx
	newDataDto.LeadingKeyWordDelimiter = ""

	lenTargetStr := len(targetStr)

	if lenTargetStr == 0 {
		return newDataDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'targetStr' is an EMPTY string!\n",
				ePrefix.String())
	}

	if startIdx < 0 {
		return newDataDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'startIdx' is less than zero!\n"+
				"startIdx='%v'\n",
				ePrefix.String(),
				startIdx)
	}

	if startIdx >= lenTargetStr {

		return newDataDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'startIdx' is out-of-bounds!\n"+
				"startIdx='%v'\t\tLast TargetStr Index='%v'\n"+
				"Length Of TargetStr='%v'\n",
				ePrefix.String(),
				startIdx, lenTargetStr-1, lenTargetStr)
	}

	lenLeadingFieldSeparators := len(leadingFieldSeparators)

	if lenLeadingFieldSeparators == 0 {

		return newDataDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'leadingFieldSeparators' is a zero length array!\n"+
				"'leadingFieldSeparators' are required!\n",
				ePrefix.String())
	}

	validTestDelimiterExists := false

	for i := 0; i < lenLeadingFieldSeparators; i++ {

		if len(leadingFieldSeparators[i]) == 0 {
			continue
		}

		validTestDelimiterExists = true

	}

	if !validTestDelimiterExists {
		newDataDto.ConvertToErrorState()
		return newDataDto,
			fmt.Errorf("%v\n"+
				"Error: Parameter 'leadingFieldSeparators' Delimiters Array "+
				"consists entirely of empty strings!\n",
				ePrefix.String())
	}

	lenTrailingFieldSeparators := len(trailingFieldSeparators)

	if lenTrailingFieldSeparators == 0 {

		return newDataDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input Parameter 'trailingFieldSeparators' is a zero length array!\n"+
				"'trailingFieldSeparators' are required!\n",
				ePrefix.String())
	}

	validTestDelimiterExists = false

	for i := 0; i < lenTrailingFieldSeparators; i++ {

		if len(trailingFieldSeparators[i]) == 0 {
			continue
		}

		validTestDelimiterExists = true
	}

	if !validTestDelimiterExists {
		newDataDto.ConvertToErrorState()
		return newDataDto,
			fmt.Errorf("%v\n"+
				"Error: Parameter 'trailingFieldSeparators' Delimiters Array consists entirely of empty strings!\n",
				ePrefix.String())
	}

	targetStrRunes := []rune(targetStr)
	lenTargetStr = len(targetStrRunes)
	lastGoodTargetStrIdx := lenTargetStr - 1

	lenOfEndOfLineDelimiters := len(endOfLineDelimiters)
	delimiterIdx := -1
	delimiterValue := ""
	validTestDelimiterExists = false

	// Check End-Of-Line Delimiters
	if lenOfEndOfLineDelimiters > 0 {

		for b := 0; b < lenOfEndOfLineDelimiters; b++ {

			if len(endOfLineDelimiters[b]) == 0 {
				continue
			}

			validTestDelimiterExists = true

			eolDelimiterIdx := strings.Index(targetStr[startIdx:], endOfLineDelimiters[b])

			if eolDelimiterIdx == -1 {
				continue
			}

			if delimiterIdx == -1 ||
				eolDelimiterIdx < delimiterIdx {
				delimiterIdx = eolDelimiterIdx
				delimiterValue = endOfLineDelimiters[b]
			}
		}

		if !validTestDelimiterExists {
			newDataDto.ConvertToErrorState()
			return newDataDto,
				fmt.Errorf("%v\n"+
					"Error: End-Of-Line Delimiters Array consists "+
					"entirely of empty strings!\n",
					ePrefix.String())
		}

		if delimiterIdx > -1 {
			// Valid End-Of-Line Delimiter does exist
			delimiterIdx += startIdx
			newDataDto.EndOfLineDelimiter = delimiterValue
			newDataDto.EndOfLineDelimiterIndex = delimiterIdx

			delimiterIdx-- // Compute last good Target String Index

			if delimiterIdx < lastGoodTargetStrIdx {
				// End-Of-Line Index is less than or equal to 'lastGoodTargetStrIds'
				newDataDto.DataFieldTrailingDelimiter = delimiterValue
				newDataDto.DataFieldTrailingDelimiterType = DfTrailDelimiter.EndOfLine()
				lastGoodTargetStrIdx = delimiterIdx
			}
		} // End of if delimiterIdx > -1 {
	} // End of if lenOfEndOfLineDelimiters > 0 {

	if startIdx > lastGoodTargetStrIdx ||
		lastGoodTargetStrIdx < 0 {

		newDataDto.TargetStrLastGoodIndex = lastGoodTargetStrIdx

		return newDataDto, nil
	}

	lenCommentDelimiters := len(commentDelimiters)

	// Check Comment Delimiters
	if lenCommentDelimiters > 0 {

		delimiterIdx = -1
		delimiterValue = ""
		validTestDelimiterExists = false

		for b := 0; b < lenCommentDelimiters; b++ {

			if len(commentDelimiters[b]) == 0 {
				continue
			}

			validTestDelimiterExists = true

			commentIdx := strings.Index(targetStr[startIdx:], commentDelimiters[b])

			if commentIdx == -1 {
				continue
			}

			if delimiterIdx == -1 ||
				commentIdx < delimiterIdx {
				delimiterIdx = commentIdx
				delimiterValue = commentDelimiters[b]
			}
		}

		if !validTestDelimiterExists {
			newDataDto.ConvertToErrorState()
			return newDataDto,
				fmt.Errorf("%v\n"+
					"Error: Comment Delimiters Array consists "+
					"entirely of empty strings!\n",
					ePrefix.String())
		}

		if delimiterIdx > -1 {

			delimiterIdx += startIdx
			newDataDto.CommentDelimiter = delimiterValue
			newDataDto.CommentDelimiterIndex = delimiterIdx
			delimiterIdx--

			if delimiterIdx < lastGoodTargetStrIdx {

				// Comment Index is less than or equal to 'lastGoodTargetStrIds'
				newDataDto.DataFieldTrailingDelimiter = delimiterValue
				newDataDto.DataFieldTrailingDelimiterType = DfTrailDelimiter.Comment()
				lastGoodTargetStrIdx = delimiterIdx
			}
		}
	}

	newDataDto.TargetStrLastGoodIndex = lastGoodTargetStrIdx

	if startIdx > lastGoodTargetStrIdx ||
		lastGoodTargetStrIdx < 0 {

		newDataDto.ConvertToErrorState()

		return newDataDto, nil
	}

	lenLeadingKeyWordDelimiters := len(leadingKeyWordDelimiters)

	// Check Leading Key Word Delimiters
	if lenLeadingKeyWordDelimiters > 0 {
		delimiterIdx = -1
		delimiterValue = ""
		validTestDelimiterExists = false

		for k := 0; k < lenLeadingKeyWordDelimiters; k++ {

			if len(leadingKeyWordDelimiters[k]) == 0 {
				// Zero length strings are not processed
				continue
			}

			validTestDelimiterExists = true

			tempKeyWordIdx := strings.Index(targetStr[startIdx:], leadingKeyWordDelimiters[k])

			if tempKeyWordIdx == -1 {
				continue
			}

			if delimiterIdx == -1 ||
				tempKeyWordIdx < delimiterIdx {

				delimiterIdx = tempKeyWordIdx
				delimiterValue = leadingKeyWordDelimiters[k]
			}
		}

		if !validTestDelimiterExists {
			newDataDto.ConvertToErrorState()
			return newDataDto,
				fmt.Errorf("%v\n"+
					"Error: Leading Key Word Delimiters Array "+
					"consists entirely of empty strings!\n",
					ePrefix.String())
		}

		if delimiterIdx == -1 {
			// Key Word Delimiters were requested,
			// but none were found. Exit!
			return newDataDto, nil
		}

		if delimiterIdx > -1 {
			// All of the key word delimiters were zero
			// length strings. Therefore, ignore
			// key word delimiters.
			delimiterIdx += startIdx

			if delimiterIdx >= lastGoodTargetStrIdx {
				// Key Word Delimiter was found but it is
				// located beyond the last good character index.
				// Probably located inside a comment or after a new-line.
				return newDataDto, nil
			}

			newDataDto.LeadingKeyWordDelimiter = delimiterValue
			newDataDto.LeadingKeyWordDelimiterIndex = delimiterIdx

			startIdx = len(delimiterValue) + delimiterIdx
		} // End of if delimiterIdx > -1
	} // End of if lenLeadingKeyWordDelimiters > 0

	//////////////////////////////
	// Main Target String Loop
	//////////////////////////////
	fieldDataRunes := make([]rune, 0, 20)
	firstDataFieldIdx := -1

	i := startIdx

	for i <= lastGoodTargetStrIdx {

		if firstDataFieldIdx == -1 {

			for j := 0; j < lenLeadingFieldSeparators; j++ {

				idxLeadingFieldSep := strings.Index(targetStr[i:], leadingFieldSeparators[j])

				if idxLeadingFieldSep != 0 {
					continue
				}

				// Found a leading Field Separator - skip it
				i += len(leadingFieldSeparators[j])

				goto cycleMainTargetLoop
			}

		} else {

			for k := 0; k < lenTrailingFieldSeparators; k++ {

				idxTrailingFieldSep := strings.Index(targetStr[i:], trailingFieldSeparators[k])

				if idxTrailingFieldSep != 0 {
					continue
				}

				newDataDto.DataFieldTrailingDelimiter = trailingFieldSeparators[k]

				newDataDto.DataFieldTrailingDelimiterType = DfTrailDelimiter.EndOfField()

				goto exitMainTargetLoop
			}

		}

		if firstDataFieldIdx == -1 {
			firstDataFieldIdx = i
		}

		fieldDataRunes = append(fieldDataRunes, targetStrRunes[i])

		i++

	cycleMainTargetLoop:
	}

exitMainTargetLoop:

	if len(fieldDataRunes) == 0 {
		return newDataDto, nil
	}

	if newDataDto.DataFieldTrailingDelimiterType == DfTrailDelimiter.Unknown() {
		newDataDto.DataFieldTrailingDelimiterType = DfTrailDelimiter.EndOfString()
	}

	newDataDto.DataFieldStr = string(fieldDataRunes)
	newDataDto.DataFieldLength = len(newDataDto.DataFieldStr)
	newDataDto.DataFieldIndex = firstDataFieldIdx
	newDataDto.TargetStrLastGoodIndex = lastGoodTargetStrIdx
	nextIdx := newDataDto.DataFieldIndex + newDataDto.DataFieldLength

	if nextIdx > lastGoodTargetStrIdx {
		newDataDto.NextTargetStrIndex = -1
	} else {
		newDataDto.NextTargetStrIndex = nextIdx
	}

	return newDataDto, nil
}

// ExtractNumericDigits - Examines an input parameter 'targetStr' to identify and extract the
// first instance of a number string. The number string will be comprised of one or more
// consecutive numeric digits (0-9) and may include leading, trailing or interior non-numeric
// characters as specified by input parameters.
//
// The search for this number string will be started at the index specified by input parameter
// 'startIdx'. Beginning at 'startIdx' the 'targetStr' will be searched to identify and extract
// the first instance of a number string.
//
// A number string is usually defined a string of consecutive numeric digits. However, this
// method allows the caller to include additional non-numeric characters as identified by
// input parameters	'keepLeadingChars', 'keepInteriorChars' and 'keepTrailingChars'.
//
// 'keepLeadingChars' is a string of characters which will be prefixed to the number string
// if those characters exist in 'targetStr' and immediately precede the number string.
//
// 'keepInteriorChars' is a string of characters which, if they exist within the number string,
// will be retained and presented in the final extracted number string.
//
// 'keepTrailingChars' is a string of characters which will be suffixed to the end of the
// final extracted number string.  To qualify, the designated 'keepTrailingChars' must immediately
// follow the number string contained in 'targetStr'.
//
// If successfully located within 'targetStr' the first instance of a number string along with
// characteristics describing that number string are returned in a Type 'NumStrProfileDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string    - The target string to be searched for the first instance of
//                                  a number string. A number string is usually defined as a
//                                  string comprised of one or more consecutive numeric digits.
//                                  Additional parameters provided by this method will allow
//                                  the caller to insert specified non-numeric characters at
//                                  the beginning, end or interior of a number string.
//
//  startIdx               int    - The starting index in input parameter 'targetStr'
//                                  from which the search for a number string will be
//                                  initiated. This useful in extracting multiple number
//                                  strings form a single 'targetStr'.
//
//  keepLeadingChars    string    - This string contains non-numeric characters which will be
//                                  retained as a prefix to the final number string extracted
//                                  from the 'targetStr' parameter. To be included, these characters
//                                  must exist in 'targetStr' and must immediately precede the
//                                  first instance of a number string.
//
//                                  For example, if the target string is "Hello $123789 world" and
//                                  parameter 'keepLeadingChars' includes the USA currency character,
//                                  '$', the returned number string would be '$123789'.  If no currency
//                                  character was included in 'keepLeadingChars', the returned number
//                                  string would be '123789'. It is worth noting that if the target
//                                  string was '$ 123789' and a currency symbol, '$', was included
//                                  in 'keepLeadingChars', the returned number string would still be
//                                  '123789' because 'keepLeadingChars' characters must immediately
//                                  precede the string of numeric digits in 'targetStr'.
//
//                                  Specifically, if the plus ('+') and minus ('-') sign are NOT
//                                  included in 'keepLeadingChars' those leading number signs will
//                                  never be included in the final number string.
//
//                                  Leading characters will not be repeated. If for some reason you
//                                  wanted to retain two leading currency symbols ("$$") it would be
//                                  necessary to include two currency characters in 'keepLeadingChars'.
//
//  keepInteriorChars   string    - This string contains non-numeric characters which will be retained
//                                  as valid characters within the final extracted number string. The
//                                  characters must exist withing the first instance of a number string
//                                  located in 'targetStr'. Such interior characters might include
//                                  thousands separators (commas) or decimal points (periods).
//
//                                  For example, if a comma and a period are included in 'keepInteriorChars'
//                                  and the target string is "Hello word 123,456,789.25 !", the returned
//                                  number string would be "123,456,789.25".  If the comma character was
//                                  NOT included in the 'keepInteriorChars' string, the returned number
//                                  string would be '123', since the number string extraction parser
//                                  would break on the comma, a non-numeric digit.
//
//                                  'keepInteriorChars' will NOT allow multiple non-numeric characters
//                                  to exist within the interior of the final extracted number string.
//                                  Only single non-numeric characters are allowed within a number string.
//
//  keepTrailingChars   string    - This string contains non-numeric characters which should be retained
//                                  at the end of the final number string. By default, a non-numeric
//                                  character will mark the end of a number string. However, if the caller
//                                  elects to use parameter 'keepTrailingChars' to retain non-numeric
//                                  characters such as a trailing right-parenthesis, then those non-numeric
//                                  characters will be retained in the final extracted number string.
//
//                                  Trailing characters will not be repeated. If for some reason you
//                                  wanted to retain two closing parentheses symbols ("))") it would be
//                                  necessary to include closing parentheses characters in 'keepTrailingChars'.
//
//                                  It should be emphasized that 'keepTrailingChars' must immediately
//                                  follow the first instance of a number string in parameter, 'targetStr'.
//
//                                  Example #1:
//                                    Target String = "Hello world, (1234). Today is new day."
//                                    keepLeadingChars = "("
//                                    keepInteriorChars = ""
//                                    keepTrailingChars= ")"
//                                    Extracted Number String = "(1234)"
//
//                                  Example #2:
//                                    Target String = "Hello world, USA GDP growth is projected at 1.8%."
//                                    keepLeadingChars = ""
//                                    keepInteriorChars = "."
//                                    keepTrailingChars= "%"
//                                    Extracted Number String = "1.8%"
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods listed
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
//  NumStrProfileDto    - If successful, this method will return a type 'NumStrProfileDto'
//                        populated with the extracted number string and additional profile
//                        information related to the extracted number string.
//
//     type NumStrProfileDto struct {
//
//           TargetStr            string   //  The original target string which is scanned for a
//                                         //    number string
//
//           TargetStrStartIndex  int      //  The starting index in 'TargetStr' from which the
//                                         //    number string search was initiated.
//
//           LeadingSignIndex     int      //  The string index of a leading sign in 'NumStr' below. If a
//                                         //    leading sign character is NOT present in 'NumStr' this
//                                         //    value is set to -1
//
//           LeadingSignChar      string   //  If a leading sign character (plus '+' or minus '-')
//                                         //    exists in data field 'NumStr' (below), it is stored
//                                         //    in this string.
//
//           FirstNumCharIndex    int      //  The index in 'TargetStr' (above) where the first character
//                                         //    of the extracted number string is located.
//
//           NextTargetStrIndex   int      //  The index of the next character in 'TargetStr' immediately
//                                         //    following the extracted number string.
//
//           NumStrLen            int      //  The length of the extracted number string.
//
//           NumStr               string   //  The number string extracted from 'TargetStr'.
//     }
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If 'startIndex' is less than zero or if 'startIndex'
//       exceeds the last character index in 'targetStr', an error
//       will be returned.
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
//  ePrefix := "TestStrOps_ExtractNumericDigits_06() "
//  targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
//  startIndex := 0
//  keepLeadingChars := "$("
//  keepInteriorChars := ",."
//  keepTrailingChars := ")"
//
//
//  nStrDto,
//  err :=
//    StrMech{}.Ptr().
//         ExtractNumericDigits(
//             targetStr,
//             startIndex,
//             keepLeadingChars,
//             keepInteriorChars,
//             keepTrailingChars,
//             ePrefix)
//
//  ----------------------------------------
//
//  nStrDto.NumStr is now equal to:
//     "$(1,250,364.33)"
//
//
func (sMechAtom *strMechAtom) extractNumericDigits(
	targetStr string,
	startIndex int,
	keepLeadingChars string,
	keepInteriorChars string,
	keepTrailingChars string,
	ePrefix *ePref.ErrPrefixDto) (
	NumStrProfileDto,
	error) {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechAtom." +
			"extractNumericDigits()")

	nStrDto := NumStrProfileDto{}.New()
	nStrDto.TargetStr = targetStr
	nStrDto.StartIndex = startIndex

	lenTargetStr := len(targetStr)

	if lenTargetStr == 0 {
		return nStrDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'targetStr' is "+
				"an empty string!\n",
				ePrefix.String())
	}

	if startIndex < 0 {
		return nStrDto, fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'startIndex' is less than zero!\n"+
			"startIndex='%v'",
			ePrefix.String(),
			startIndex)
	}

	if startIndex >= lenTargetStr {
		return nStrDto, fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'startIndex' is INVALID!\n"+
			"'startIndex' exceeds the last character index in "+
			"'targetStr'\n"+
			"startIndex='%v'\tlast character index='%v'\n"+
			"targetStr='%v'\n",
			ePrefix.String(),
			startIndex, lenTargetStr-1, targetStr)
	}

	targetRunes := []rune(targetStr)
	lenTargetStr = len(targetRunes)

	keepLeadingRunes := make([]rune, 0)
	lenKeepLeadingRunes := 0

	keepInteriorRunes := make([]rune, 0)
	lenKeepInteriorRunes := 0

	keepTrailingRunes := make([]rune, 0)
	lenKeepTrailingRunes := 0

	if len(keepLeadingChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepLeadingChars); a++ {

			if keepLeadingChars[a] >= '0' &&
				keepLeadingChars[a] <= '9' {
				continue
			}

			keepLeadingRunes = append(keepLeadingRunes, rune(keepLeadingChars[a]))

		}

		lenKeepLeadingRunes = len(keepLeadingRunes)
	}

	if len(keepInteriorChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepInteriorChars); a++ {

			if keepInteriorChars[a] >= '0' &&
				keepInteriorChars[a] <= '9' {
				continue
			}

			keepInteriorRunes = append(keepInteriorRunes, rune(keepInteriorChars[a]))

		}

		lenKeepInteriorRunes = len(keepInteriorRunes)
	}

	if len(keepTrailingChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepTrailingChars); a++ {

			if keepTrailingChars[a] >= '0' &&
				keepTrailingChars[a] <= '9' {
				continue
			}

			keepTrailingRunes = append(keepTrailingRunes, rune(keepTrailingChars[a]))

		}

		lenKeepTrailingRunes = len(keepTrailingRunes)
	}

	numberRunesCaptured := make([]rune, 0, 20)
	lenNumberRunesCaptured := 0

	leadingCharRunesCaptured := make([]rune, 0, 20)
	lenLeadingCharRunesCaptured := 0

	firstNumericDigitIdx := -1

	for e := startIndex; e < lenTargetStr; e++ {

		if targetRunes[e] >= '0' &&
			targetRunes[e] <= '9' &&
			firstNumericDigitIdx == -1 {
			// Target has at least one numeric
			// digit - and we found it.
			firstNumericDigitIdx = e
			break
		}
	}

	if firstNumericDigitIdx == -1 {
		// There are no numeric digits
		// in this target string.
		// EXIT HERE!!!
		return nStrDto, nil
	}

	firstNumStrCharIdx := -1
	leadingSignChar := ""

	// Check for leading non-numeric characters that
	// need to be retained at the front of the number
	// string.
	if lenKeepLeadingRunes > 0 &&
		startIndex < firstNumericDigitIdx {

		for f := firstNumericDigitIdx - 1; f >= startIndex; f-- {

			for g := 0; g < lenKeepLeadingRunes; g++ {

				if keepLeadingRunes[g] == targetRunes[f] {

					if keepLeadingRunes[g] == '+' ||
						keepLeadingRunes[g] == '-' {

						// This is a leading sign char
						leadingSignChar = string(targetRunes[f])

						leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
						// Delete Leading Sign character. It will not be repeated in
						// future searches. Only one leading sign char per number string.

						keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
						lenKeepLeadingRunes--

						firstNumStrCharIdx = f

						// Now delete the alternative leading sign character.
						// There are only two - plus or minus
						nextSignChar := '-'

						if leadingSignChar == "-" {
							nextSignChar = '+'
						}

						// Leading sign char has been found. Now delete the
						// alternative lead sign char to avoid duplications
						for m := 0; m < lenKeepLeadingRunes; m++ {
							if keepLeadingRunes[m] == nextSignChar {
								keepLeadingRunes = append(keepLeadingRunes[0:m], keepLeadingRunes[m+1:]...)
								lenKeepLeadingRunes--
							}
						}

						break

					} else {

						// Standard Keep Leading Rune character found
						leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
						// Delete Leading Rune character. It will not be repeated in
						// future searches

						firstNumStrCharIdx = f

						keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
						lenKeepLeadingRunes--
						break
					}
				}
			}

			t := len(leadingCharRunesCaptured)

			if t > lenLeadingCharRunesCaptured {
				lenLeadingCharRunesCaptured = t
				continue
			}

			break
		}
	}

	leadingSignIndex := -1

	if lenLeadingCharRunesCaptured > 0 {

		for h := lenLeadingCharRunesCaptured - 1; h >= 0; h-- {

			if leadingCharRunesCaptured[h] == '+' ||
				leadingCharRunesCaptured[h] == '-' {

				numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
				leadingSignIndex = lenNumberRunesCaptured
				lenNumberRunesCaptured++

			} else {
				numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
				lenNumberRunesCaptured++
			}
		}
	}

	// Main Number String Extraction Loop
	isEndOfNumStr := false

	for i := firstNumericDigitIdx; i < lenTargetStr; i++ {

		if !isEndOfNumStr {

			if targetRunes[i] >= '0' && targetRunes[i] <= '9' {

				numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
				continue
			}

			for j := 0; j < lenKeepInteriorRunes; j++ {

				if targetRunes[i] == keepInteriorRunes[j] {

					if i+1 >= lenTargetStr ||
						(targetRunes[i+1] < '0' || targetRunes[i+1] > '9') {
						// We are either at the end of string or the next char
						// is NOT a numeric character.
						goto trailChar
					}

					numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])

					goto numDigitLoop
				}
			}

		}

	trailChar:
		isEndOfNumStr = true

		for k := 0; k < lenKeepTrailingRunes; k++ {

			if targetRunes[i] == keepTrailingRunes[k] {
				numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
				// Only one instance of a keep trailing rune character is captured.
				// Delete the keep trailing rune character to prevent repeat captures.
				keepTrailingRunes = append(keepLeadingRunes[0:k], keepTrailingRunes[k+1:]...)
				lenKeepTrailingRunes--
				goto numDigitLoop
			}

		}

		// Non-numeric character and Non-Trailing Character: Exit the Loop
		break

	numDigitLoop:
	}

	if len(numberRunesCaptured) > 0 {
		nStrDto.NumStr = string(numberRunesCaptured)

		if firstNumStrCharIdx > -1 {
			nStrDto.FirstNumCharIndex = firstNumStrCharIdx
		} else {
			nStrDto.FirstNumCharIndex = firstNumericDigitIdx
		}

		nStrDto.NumStrLen = len(nStrDto.NumStr)
		nStrDto.LeadingSignChar = leadingSignChar
		nStrDto.LeadingSignIndex = leadingSignIndex
		nStrDto.NextTargetStrIndex =
			nStrDto.FirstNumCharIndex + nStrDto.NumStrLen

		if nStrDto.NextTargetStrIndex >= len(targetStr) {
			nStrDto.NextTargetStrIndex = -1
		}
	}

	return nStrDto, nil
}

// ptr - Returns a pointer to a new instance of
// strMechAtom.
//
func (sMechAtom strMechAtom) ptr() *strMechAtom {

	if sMechAtom.lock == nil {
		sMechAtom.lock = new(sync.Mutex)
	}

	sMechAtom.lock.Lock()

	defer sMechAtom.lock.Unlock()

	return &strMechAtom{
		lock: new(sync.Mutex),
	}
}
