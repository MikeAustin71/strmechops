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
// string and inserts a line delimiter character (a.k.a. 'rune') at
// the specified line length ('lineLength').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	targetStr           string
//	   - The string which will be parsed into text lines. If
//	     'targetStr' is an empty string an error will be returned.
//	      If 'targetStr' consists entirely of white space, this
//	      method will return a string consisting of a new-line
//	      character and an error value of 'nil'.
//
//
//	lineLength          int
//	   - The maximum length of each line.
//
//
//	lineDelimiter       rune
//	   - The line delimiter character which will be inserted at the
//	     end of a line break.
//
//
//	Note: If the caller specifies a line length of 50, the line delimiter
//	character may be placed in the 51st character position depending upon
//	the word breaks.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string
//	   - If this method completes successfully, this string
//	     parameter will contain the text with line breaks delimited
//	     by the input parameter 'lineDelimiter'.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
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
			ePrefix.XCpy(
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
				ePrefix.XCpy(
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
						ePrefix.XCpy(
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
	targetStrMech.stringData = incomingStrMech.stringData

	return err
}

// CopyOut - Creates a 'deep' copy of input parameter
// 'strOps', an instance of StrMech.
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
//	targetStr               string   - The target string from which the data field will be extracted.
//
//
//	leadingKeyWordDelimiters []string - Data fields are often preceded by field names or field designators.
//	                                     The 'leadingKeyWordDelimiters' parameter is a string array
//	                                     containing 'Key Word Delimiters'. A Key Word Delimiter may be
//	                                     a Key Word string or a character which identifies and immediately
//	                                     precedes the data field. If multiple Key Word Delimiters exist
//	                                     in 'targetStr' the first instance of a key word in 'targetStr'
//	                                     will be designated as the Key Word Delimiter.
//
//	                                     If this parameter is populated, the search for a data field
//	                                     will begin immediately after the first located Key Word
//	                                     Delimiter string. If none of the Keywords in this string array
//	                                     are located in 'targetStr', an empty string will be returned
//	                                     for data field. If this parameter is populated, at least one
//	                                     of the Keywords MUST exist in 'targetStr' before a data field
//	                                     will be extracted and returned.
//
//	                                     If this parameter is an empty string array, the search for a
//	                                     data field will begin at the string index designated by
//	                                     parameter, 'startIdx'.
//
//
//	startIdx                int      - The string index in parameter 'targetStr' from which the search for
//	                                     a data field will begin. Note that the starting index will be adjusted
//	                                     according to the existence of a Key Word Delimiter as explained
//	                                     above.
//
//
//	leadingFieldSeparators  []string - An array of characters or groups of characters which delimit the
//	                                     leading edge of the data field.
//
//
//	trailingFieldSeparators []string - An array of characters or groups of characters which delimit the
//	                                     end of a data field.
//
//
//	commentDelimiters       []string - Comments effectively terminate the search for a data field. This
//	                                     array stores comment characters or phrases which signal the beginning
//	                                     of a comment.
//
//
//	endOfLineDelimiters     []string - Those characters or groups of characters which mark the end of a line.
//	                                     Generally this includes characters like 'new line' or 'carriage return'.
//	                                     End of line characters will terminate the search for a data field.
//
//
//	ePrefix             *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods listed
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
//	DataFieldProfileDto - If successful, this method returns a structure containing
//	                      characteristics describing the extracted data field.
//
//	  type DataFieldProfileDto struct {
//	     TargetStr                      string //  The string from which the data field is extracted.
//	     TargetStrLength                int    //  Length of 'TargetStr'
//	     TargetStrStartIndex            int    //  The index with in 'TargetStr' from which the search for a data field was initiated.
//	     TargetStrLastGoodIndex         int    //  Last valid index in target string which is less than the target string length and is NOT an 'End Of Field' or 'End Of Line' Delimiter.
//	     LeadingKeyWordDelimiter        string //  The Leading Key Word Delimiter which is used to identify the beginning of the field search.
//	     LeadingKeyWordDelimiterIndex   int    //  Index of the found Leading Key Word Delimiter.
//	     DataFieldStr                   string //  The extracted data field string.
//	     DataFieldIndex                 int    //  The index in 'TargetStr' where the data field begins.
//	     DataFieldLength                int    //  The length of the extracted data field string.
//	     DataFieldTrailingDelimiter     string //  The trailing character which marked the end of the data field. A zero value indicates end of string encountered.
//	     DataFieldTrailingDelimType     DataFieldTrailingDelimiterType // A constant or enumeration type used to describe the type of delimiter used to mark the end of a data field.
//	     NextTargetStrIndex             int    //  The index in 'TargetStr' immediately following the extracted data field.
//	     CommentDelimiter               string //  If a Comment Delimiter is detected it is stored here.
//	     CommentDelimiterIndex          int    //  If a Comment Delimiter is detected, the string index in 'TargetStr' showing its location is stored here.
//	     EndOfLineDelimiter             string //  If an End-Of-Line Delimiter is detected it is captured and stored here.
//	     EndOfLineDelimiterIndex        int    //  If an End-Of-Line Delimiter presents, the string index in 'TargetStr' showing its location, is stored here.
//	  }
//
//
//	 error
//	   - If the method completes successfully and no errors are encountered
//	     this parameter returns 'nil'. Otherwise, if errors are encountered
//	     this return value will contain an appropriate error message.
//
//	     The most likely source of errors are invalid input parameters.
//	     Input parameters 'targetStr', 'startIdx', 'leadingFieldSeparators',
//	     'trailingFieldSeparators' and 'endOfStringDelimiters' are required
//	     parameters and must be populated with valid data.
//
//	     If an error message is returned, the text value for input
//	     parameter 'ePrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	ePrefix := "TestStrOps_ExtractDataField_01() "
//	endOfLineDelimiters := []string{"\n"}
//	commentDelimiters := []string{"#"}
//	leadingFieldDelimiters := []string{
//	"\t",
//	"\r",
//	"\f",
//	"\v",
//	" "}
//
//	trailingFieldDelimiters := []string{
//	"\t",
//	"\r",
//	"\f",
//	"\v",
//	" "}
//
//	targetStr := " Zone:\t America/Chicago\t Link:\t US/Central\t\n"
//	startIdx := 0
//	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
//
//	datDto,
//	err :=
//	  new(StrMech).
//	      ExtractDataField(
//	         targetStr,
//	         leadingKeyWordDelimiters,
//	         startIdx,
//	         leadingFieldDelimiters,
//	         trailingFieldDelimiters,
//	         commentDelimiters,
//	         endOfLineDelimiters,
//	         ePrefix)
//
//	-----------------------------------------------
//	datDto.DataFieldStr is now equal to:
//	        "America/Chicago"
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
				newDataDto.DataFieldTrailingDelimType = DfTrailDelimiter.EndOfLine()
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
				newDataDto.DataFieldTrailingDelimType = DfTrailDelimiter.Comment()
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
			// All the key word delimiters were zero
			// length strings. Therefore, ignore
			// key word delimiters.
			delimiterIdx += startIdx

			if delimiterIdx >= lastGoodTargetStrIdx {
				// Key Word Delimiter was found, but it is
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

				newDataDto.DataFieldTrailingDelimType = DfTrailDelimiter.EndOfField()

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

	if newDataDto.DataFieldTrailingDelimType == DfTrailDelimiter.None() {
		newDataDto.DataFieldTrailingDelimType = DfTrailDelimiter.EndOfString()
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
