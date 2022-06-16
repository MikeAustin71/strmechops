package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// RuneArrayDto - A Rune Array Data Transfer Object. This type is
// designed to transfer a rune array between various functions and
// types.
//
// This type serves as a wrapper for a rune array and allows methods
// to pass pointers to RuneArrayDto and then perform operations on
// the internal rune array, 'CharsArray'.
//
// For example, this rune array can be used to carry out text
// character searches by passing pointers to the RuneArrayDto.
//
// Each instance is capable of performing two types of text
// character searches:
//   (1) Linear Target Character Search
//        Each Character in the Target String must match each
//        corresponding character in the RuneArrayDto instance.
//        Linear Target Character Search Type is the default
//        setting for RuneArrayDto.
//
//   (2) Single Target Character Search
//       A single character in the Target String must mach any
//       character is the RuneArrayDto rune array. A single
//       character in the Target String is therefore compared
//       against all individual characters in the RuneArrayDto
//       instance searching for the first case of a match.
//
// Reference method: RuneArrayDto.SetCharacterSearchType()
//
type RuneArrayDto struct {
	CharsArray     []rune
	Description1   string
	Description2   string
	charSearchType CharacterSearchType // Two Possible Settings:
	//                                     CharSearchType.LinearTargetStartingIndex()
	//                                     CharSearchType.SingleTargetChar()
	//                                     Default = CharSearchType.LinearTargetStartingIndex()
	lock *sync.Mutex
}

// AddChar - Adds a single text character of type rune to the end
// of the rune array (RuneArrayDto.CharsArray) encapsulated by the
// current instance of RuneArrayDto
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charToAdd                  rune
//     - The text character which will be appended to the end of
//       the rune array encapsulated by the current instance of
//       RuneArrayDto.
//
func (charsArrayDto *RuneArrayDto) AddChar(
	charToAdd rune) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	charsArrayDto.CharsArray =
		append(charsArrayDto.CharsArray, charToAdd)
}

// CopyIn - Copies the internal rune array from an incoming
// instance of RuneArrayDto ('incomingChars') to the internal
// rune array of the current RuneArrayDto instance
// ('charsArrayDto').
//
// IMPORTANT
// The internal rune array for the current RuneArrayDto instance
// ('charsArrayDto') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingChars              *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. This method will
//       NOT change the values of internal rune array contained in
//       this instance.
//
//       All data values in this RuneArrayDto instance will be
//       copied to current RuneArrayDto instance ('charsArrayDto').
//
//       If parameter 'incomingChars' is 'nil' pointer or if its
//       internal rune array has a length of zero, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) CopyIn(
	incomingChars *RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	if incomingChars == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingChars' is a 'nil' pointer!\n",
			ePrefix.String())

		return err

	}

	lenIncomingCharsArray := len(incomingChars.CharsArray)

	if lenIncomingCharsArray == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingChars' is invalid!\n"+
			"The length of 'incomingChars.CharsArray' is zero.\n"+
			"The 'incomingChars' internal Rune Array is empty.",
			ePrefix.String())

		return err

	}

	charsArrayDto.CharsArray =
		make([]rune, lenIncomingCharsArray)

	for i := 0; i < lenIncomingCharsArray; i++ {
		charsArrayDto.CharsArray[i] =
			incomingChars.CharsArray[i]
	}

	return err
}

// CopyOut - Returns a deep copy of the current RuneArrayDto
// instance.
//
// If the current RuneArrayDto instance contains an empty rune
// array, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  copyOfRuneArrayDto         copyOfRuneArrayDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current RuneArrayDto instance.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) CopyOut(
	errorPrefix interface{}) (
	copyOfRuneArrayDto RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"CopyOut()",
		"")

	if err != nil {
		return copyOfRuneArrayDto, err
	}

	lenOfLocalCharsArray := len(charsArrayDto.CharsArray)

	if lenOfLocalCharsArray == 0 {

		err = fmt.Errorf("%v\n" +
			"ERROR: The current instance of RuneArrayDto has an empty rune array!\n" +
			"The length of 'charsArrayDto.CharsArray' is zero.\n" +
			ePrefix.String())

		return copyOfRuneArrayDto, err
	}

	copyOfRuneArrayDto.CharsArray =
		make([]rune, lenOfLocalCharsArray)

	for i := 0; i < lenOfLocalCharsArray; i++ {
		copyOfRuneArrayDto.CharsArray[i] =
			charsArrayDto.CharsArray[i]
	}

	return copyOfRuneArrayDto, err
}

// Empty - Resets the internal member variable rune array to 'nil'.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete the pre-existing characters in the
// internal rune array for the current instance of RuneArrayDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (charsArrayDto *RuneArrayDto) Empty() {

	charsArrayDto.lock.Lock()

	charsArrayDto.CharsArray = nil

	charsArrayDto.Description1 = ""

	charsArrayDto.Description2 = ""

	charsArrayDto.charSearchType = CharSearchType.LinearTargetStartingIndex()

	charsArrayDto.lock.Unlock()

	charsArrayDto.lock = nil
}

// Equal - Receives a pointer to another instance of RuneArrayDto
// and proceeds to compare its internal member variables to those
// of the current RuneArrayDto instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingRuneArrayDto       *RuneArrayDto
//     - A pointer to an external instance of RuneArrayDto. The
//       internal member variable data values in this instance will
//       be compared to those in the current instance of
//       RuneArrayDto. The results of this comparison will
//       be returned to the calling function as a boolean value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingRuneArrayDto' are equivalent in
//       all respects to those contained in the current instance of
//       RuneArrayDto, this return value will be set to 'true'.
//
//       Otherwise, this method will return 'false'.
//
func (charsArrayDto *RuneArrayDto) Equal(
	incomingRuneArrayDto *RuneArrayDto) bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	if incomingRuneArrayDto == nil {
		return false
	}

	lenOfCharsArray := len(charsArrayDto.CharsArray)

	if lenOfCharsArray !=
		len(incomingRuneArrayDto.CharsArray) {

		return false
	}

	if lenOfCharsArray == 0 {
		return true
	}

	for i := 0; i < lenOfCharsArray; i++ {

		if charsArrayDto.CharsArray[i] !=
			incomingRuneArrayDto.CharsArray[i] {
			return false
		}

	}

	return true
}

// GetRuneArrayDescription1 - Returns 'Description1', the optional
// description string associated with the current instance of RuneArrayDto.
//
// Users have the option of configuring a text string to describe
// the function or purpose of the text characters configured for
// the current instance of RuneArrayDto.
//
// An instance of RuneArrayDto has two description strings which
// may be used to describe operational scenarios.
//
// This method returns the current value of internal member
// variable:
//    RuneArrayDto.Description1
//
func (charsArrayDto *RuneArrayDto) GetRuneArrayDescription1() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return charsArrayDto.Description1
}

// GetRuneArrayDescription2 - Returns 'Description2', the optional
// description string associated with the current instance of RuneArrayDto.
//
// Users have the option of configuring a text string to describe
// the function or purpose of the text characters configured for
// the current instance of RuneArrayDto.
//
// An instance of RuneArrayDto has two description strings which
// may be used to describe operational scenarios.
//
// This method returns the current value of internal member
// variable:
//    RuneArrayDto.Description2
//
func (charsArrayDto *RuneArrayDto) GetRuneArrayDescription2() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return charsArrayDto.Description2
}

// GetRuneArrayLength - Returns the length of the internal
// rune array, 'CharsArray' as an integer value.
//
func (charsArrayDto *RuneArrayDto) GetRuneArrayLength() int {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return len(charsArrayDto.CharsArray)
}

// GetRuneArray - Returns a deep copy of the internal rune array,
// 'CharsArray' configured for the current instance of
// RuneArrayDto.
//
// If the length of the internal rune array 'CharsArray' is zero,
// a value of 'nil' will be returned.
//
func (charsArrayDto *RuneArrayDto) GetRuneArray() []rune {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	lenCharsArray := len(charsArrayDto.CharsArray)

	if lenCharsArray == 0 {
		return nil
	}

	newCharsArray := make([]rune, lenCharsArray)

	for i := 0; i < lenCharsArray; i++ {
		newCharsArray[i] =
			charsArrayDto.CharsArray[i]
	}

	return newCharsArray
}

// NewRuneArray - Receives an array of runes and returns a new
// instance of RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
//
//
//  description                string
//
//     - Users have the option of configuring a text string to
//       describe the function or purpose of the text characters
//       configured for the returned instance of RuneArrayDto.
//
//
//
//
//  charSearchType             CharacterSearchType
//     - An enumeration value used to specify the type target
//       string search algorithm applied by the returned instance
//       of RuneArrayDto.
//
//
//       The Character Search Type must be set to one of the
//       following enumeration values:
//
//        CharSearchType.LinearTargetStartingIndex()
//        CharSearchType.SingleTargetChar()
//        CharSearchType.LinearEndOfString()
//
// Character Search Type Options
//
//    CharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting Index
//      Search Type. This means that each character in the Target
//      Search String will be compared to each corresponding
//      character in the Test String beginning at a specified
//      starting index in the Target Search String.
//
//      The search will proceed for from left to right in Test
//      Character Sequence.
//
//      If the Test Characters are NOT found in the Target Search
//      String beginning at the designated Target String Starting
//      Index, the search outcome will be unsuccessful and NO match
//      will be declared.
//
//      A 'Match', or successful search outcome, is defined as the
//      case where each character in the Target String matches each
//      corresponding character in the Test String beginning at the
//      designated Target String Starting Index.
//
//
//        Example
//                                  1         2         3
//                 Index  0123456789012345678901234567890
//        Target String: "Hey, Xray-4 is the call sign."
//        Target String Starting Index: 5
//          Test String: "Xray"
//
//      In this example of a Linear Target Starting Index Search, a
//      match between the Target String and Test String will be
//      declared, if and only if, the search begins at Target String
//      index number 5. If the search begins at an any index other
//      than 5, no match will be declared and the search will be
//      classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   CharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the Target
//      Search String will be compared to all characters in the Test
//      String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome will
//      be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to all
//      characters in the Test String looking for the first matching
//      Test String Character. The search will terminate when a
//      matching character is first identified or when the end of the
//      Target String is encountered.
//
//
//        Example
//                                   1         2         3
//                  Index  0123456789012345678901234567890
//         Target String: "Hey, Xray-4 is the call sign."
//         Target String Starting Index: 0
//           Test String: "ZFXyURJK"
//
//      In this example of a Single Target Character Search, the
//      search will terminate at Target String index numbers 5
//      because it is the first Target String index to match one
//      of the Test String Characters ('X').
//
//
//   CharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String Search.
//      With this type of search operation, the entire Target Search
//      String will be searched from left to right for the
//      first occurrence of the Test String.
//
//      The search will begin the Target String Starting Index and
//      proceed left to right until (1) an instance of the entire
//      Test String is located or (2) the end of the Target Search
//      String is encountered.
//
//      This is a linear search, so a 'Match' requires that each
//      character in Target Search String must correspond to a
//      matching character in the Test String.
//
//           Example
//                                      1         2         3
//                     Index  0123456789012345678901234567890
//            Target String: "Hey, Xray-4 is the call sign."
//            Target String Starting Index: 0
//              Test String: "Xray-4"
//
//      In this example of a Linear End of String Search, the search
//      operation will begin comparing corresponding characters in
//      the Target Search String and the Test String beginning at
//      index zero. The comparison will fail at index zero, but the
//      search algorithm will continue attempting to find the Test
//      String at indexes 1,2, 3 & 4. The Test String will be found
//      beginning at index number 5 and the search algorithm will
//      terminate at that point with a successful outcome or 'Match'
//      result.
//
//
//   For more information see the source code comments for type,
//   CharacterSearchType.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewRuneArray(
	charArray []rune,
	description string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	newRuneArrayDto RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"NewRuneArray()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	lenOfCharArray := len(charArray)

	if lenOfCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid!\n"+
			"'charArray' is empty and has a length of zero.\n",
			ePrefix.String())

		return newRuneArrayDto, err
	}

	if !charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charSearchType' is invalid!\n"+
			"'charSearchType' must be set to one of two enumerations:\n"+
			"CharacterSearchType(0).LinearTargetStartingIndex()\n OR"+
			"CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"'charSearchType' string  value = '%v'\n"+
			"'charSearchType' integer value = '%v'\n",
			ePrefix.String(),
			charSearchType.String(),
			charSearchType.XValueInt())

		return newRuneArrayDto, err
	}

	newRuneArrayDto.CharsArray =
		make([]rune, lenOfCharArray)

	for i := 0; i < lenOfCharArray; i++ {
		newRuneArrayDto.CharsArray[i] =
			charArray[i]
	}

	newRuneArrayDto.Description1 = description

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// SearchForTextCharacterString - This method performs a search
// operation on a string of text characters passed by the calling
// function ('targetSearchString'). This Target String is searched
// beginning at the specified Starting Search Index
// ('startingSearchIndex') in the host Target Search String.
//
// If the Target Search String characters beginning at the Starting
// Search Index match all the text characters configured in the
// current instance of RuneArrayDto, the search operation is
// classified as successful and the return values are configured
// accordingly.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetSearchString              *RuneArrayDto
//     - A pointer to an external instance of RuneArrayDto. The
//       text characters configured for 'targetSearchString' will
//       be compared against those configured in the current
//       RuneArrayDto instance.
//
//
//  startingSearchIndex             int
//     - The zero based index in 'targetSearchString' where the
//       search operation will begin
//
//
//  errorPrefix                     interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  foundRuneArrayDtoChars          bool
//     - This method performs a search operation on a Target Search
//       String containing text characters ('targetSearchString').
//
//       Beginning at the  specified Starting Search Index in the
//       Target Search String, the text characters are compared to
//       those configured in the current RuneArrayDto instance. If
//       both sets of text characters match, the search operation
//       is classified as successful and return parameter
//       'foundRuneArrayDtoChars' is set to 'true'.
//
//       If the characters contained in the current RuneArrayDto
//       instance are not found in the Target Search String
//       beginning at the specified Starting Search Index, this
//       parameter will be set to 'false'.
//
//        Example-1:
//          Target Search String: "Xray-4 is the call sign"
//          Starting Search Index: 0
//          Current RuneArrayDto Instance: "Xray"
//          foundRuneArrayDtoChars = 'true'
//
//        Example-2:
//          Target Search String: "Coyote-4 is the call sign"
//          Starting Search Index: 0
//          Current RuneArrayDto Instance: "Xray"
//          foundRuneArrayDtoChars = 'false'
//
//
//  lastSearchIndex                 int
//     - If the search operation performed by this method is
//       successful, it means that all the characters configured in
//       the current RuneArrayDto instance were found in the host
//       Target Search String. In this case, 'lastSearchIndex' will
//       be set to the index in Target Search String
//       ('targetSearchString') occupied by the last character in
//       text string configured for the current RuneArrayDto
//       instance.
//
//       If the search operation performed by this method is
//       unsuccessful, 'lastSearchIndex' is set to the value of
//       input parameter 'startingSearchIndex'.
//
//        Example-1:
//          Target Search String: "Xray-4 is the call sign"
//          Starting Search Index: 0
//          Current RuneArrayDto Instance: "Xray"
//          foundRuneArrayDtoChars = 'true'
//          lastSearchIndex = 3 (The 'y' in "Xray")
//
//        Example-2:
//          Target Search String: "Coyote-4 is the call sign"
//          Starting Search Index: 0
//          Current RuneArrayDto Instance: "Xray"
//          foundRuneArrayDtoChars = 'false'
//          lastSearchIndex = 0 (Starting Search Index)
//
//
func (charsArrayDto *RuneArrayDto) SearchForTextCharacterString(
	targetSearchString *RuneArrayDto,
	targetStartingSearchIndex int,
	targetSearchLength int,
	errorPrefix interface{}) (
	foundRuneArrayDtoChars bool,
	lastTargetSearchIndex int,
	lastTestStingIndex int,
	searchType CharacterSearchType,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundRuneArrayDtoChars = false

	lastTargetSearchIndex = targetStartingSearchIndex

	lastTestStingIndex = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SearchForTextCharacterString()",
		"")

	if err != nil {

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if !charsArrayDto.charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of RuneArrayDto is invalid!\n"+
			"The Character Search Type is invalid. Character Search Type\n"+
			"must be set to one of these enumeration values:\n"+
			"  CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"  CharacterSearchType(0).SingleTargetChar()\n"+
			"  CharacterSearchType(0).LinearEndOfString()\n"+
			"The Character Search Type for this instance of RuneArrayDto is"+
			" Character Search Type   String Name: %v\n"+
			" Character Search Type Integer Value: %v\n",
			ePrefix.String(),
			charsArrayDto.charSearchType.String(),
			charsArrayDto.charSearchType.XValueInt())

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err
	}

	runeArrayNanobot := runeArrayDtoNanobot{}

	if charsArrayDto.charSearchType == CharSearchType.LinearTargetStartingIndex() {

		return runeArrayNanobot.linearTargetStartingIndexSearch(
			charsArrayDto,
			"RuneArrayDto",
			targetSearchString,
			"targetSearchString",
			targetStartingSearchIndex,
			"targetStartingSearchIndex",
			targetSearchLength,
			"targetSearchLength",
			ePrefix.XCpy(
				"charsArrayDto"))

	}

	// Must Be
	// charsArrayDto.charSearchType == CharSearchType.SingleTargetChar()

	return foundRuneArrayDtoChars,
		lastTargetSearchIndex,
		lastTestStingIndex,
		searchType,
		err
}

// SetCharacterSearchType - Sets the internal member variable used
// to track the Character Search Type.
//
// Character Search Type is a series of enumeration values
// specifying the type of text character search algorithm applied
// by the current instance of RuneArrayDto.
//
// The Character Search Type must be set to one of the following
// enumeration values:
//
//  CharSearchType.LinearTargetStartingIndex()
//  CharSearchType.SingleTargetChar()
//  CharSearchType.LinearEndOfString()
//
//
// ----------------------------------------------------------------
//
// Character Search Type Options
//
//  CharSearchType.LinearTargetStartingIndex()
//  - Designates the search type as a Linear Target Starting Index
//    Search Type. This means that each character in the Target
//    Search String will be compared to each corresponding
//    character in the Test String beginning at a specified
//    starting index in the Target Search String.
//
//    The search will proceed for from left to right in Test
//    Character Sequence.
//
//    If the Test Characters are NOT found in the Target Search
//    String beginning at the designated Target String Starting
//    Index, the search outcome will be unsuccessful and NO match
//    will be declared.
//
//    A 'Match', or successful search outcome, is defined as the
//    case where each character in the Target String matches each
//    corresponding character in the Test String beginning at the
//    designated Target String Starting Index.
//
//
//      Example
//                                1         2         3
//               Index  0123456789012345678901234567890
//      Target String: "Hey, Xray-4 is the call sign."
//      Target String Starting Index: 5
//        Test String: "Xray"
//
//    In this example of a Linear Target Starting Index Search, a
//    match between the Target String and Test String will be
//    declared, if and only if, the search begins at Target String
//    index number 5. If the search begins at an any index other
//    than 5, no match will be declared and the search will be
//    classified as unsuccessful.
//
//    NOTE: Linear Target Starting Index is the default search
//          type.
//
//
// CharSearchType.SingleTargetChar()
//  - Designates the search type as a Single Target Character
//    Search Type. This means that a single character in the Target
//    Search String will be compared to all characters in the Test
//    String.
//
//    If a single Target String character equals any character in
//    the Test String, a 'Match' or successful search outcome will
//    be declared.
//
//    The search will proceed from left to right in the Target
//    String. Each Target String Character will be compared to all
//    characters in the Test String looking for the first matching
//    Test String Character. The search will terminate when a
//    matching character is first identified or when the end of the
//    Target String is encountered.
//
//
//      Example
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//       Target String Starting Index: 0
//         Test String: "ZFXyURJK"
//
//    In this example of a Single Target Character Search, the
//    search will terminate at Target String index numbers 5
//    because it is the first Target String index to match one
//    of the Test String Characters ('X').
//
//
// CharSearchType.LinearEndOfString()
//  - Designates the search type as a Linear End Of String Search.
//    With this type of search operation, the entire Target Search
//    String will be searched from left to right for the
//    first occurrence of the Test String.
//
//    The search will begin the Target String Starting Index and
//    proceed left to right until (1) an instance of the entire
//    Test String is located or (2) the end of the Target Search
//    String is encountered.
//
//    This is a linear search, so a 'Match' requires that each
//    character in Target Search String must correspond to a
//    matching character in the Test String.
//
//         Example
//                                    1         2         3
//                   Index  0123456789012345678901234567890
//          Target String: "Hey, Xray-4 is the call sign."
//          Target String Starting Index: 0
//            Test String: "Xray-4"
//
//    In this example of a Linear End of String Search, the search
//    operation will begin comparing corresponding characters in
//    the Target Search String and the Test String beginning at
//    index zero. The comparison will fail at index zero, but the
//    search algorithm will continue attempting to find the Test
//    String at indexes 1,2, 3 & 4. The Test String will be found
//    beginning at index number 5 and the search algorithm will
//    terminate at that point with a successful outcome or 'Match'
//    result.
//
//
// For more information see the source code comments for type,
// CharacterSearchType.
//
func (charsArrayDto *RuneArrayDto) SetCharacterSearchType(
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SetCharacterSearchType()",
		"")

	if err != nil {
		return err
	}

	if !charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charSearchType' is invalid!\n"+
			"'charSearchType' must be set to one of three enumerations:\n"+
			"CharacterSearchType(0).LinearTargetStartingIndex()\n OR"+
			"CharacterSearchType(0).SingleTargetChar()\n"+
			"CharacterSearchType(0).LinearEndOfString()\n"+
			"'charSearchType' string  value = '%v'\n"+
			"'charSearchType' integer value = '%v'\n",
			ePrefix.String(),
			charSearchType.String(),
			charSearchType.XValueInt())

		return err
	}

	charsArrayDto.charSearchType = charSearchType

	return err
}

// SetDescription1 - Set the first text description associated with
// the current instance of RuneArrayDto.
//
// The text description is optional. No error checking is performed
// on input parameter 'runeArrayDtoDesc'
//
// This method will set the internal member variable:
//    RuneArrayDto.Description1
//
func (charsArrayDto *RuneArrayDto) SetDescription1(
	runeArrayDtoDesc string) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	charsArrayDto.Description1 = runeArrayDtoDesc

	return
}

// SetDescription2 - Set the second text description associated
// with the current instance of RuneArrayDto.
//
// The text description is optional. No error checking is performed
// on input parameter 'runeArrayDtoDesc'
//
// This method will set the internal member variable:
//    RuneArrayDto.Description2
//
func (charsArrayDto *RuneArrayDto) SetDescription2(
	runeArrayDtoDesc string) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	charsArrayDto.Description2 = runeArrayDtoDesc

	return
}

// SetRuneArray - Deletes the internal rune array for the current
// instance of RuneArrayDto and repopulates it with a deep copy
// extracted from input parameter, 'charArray'.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All pre-existing data in the current instance of RuneArrayDto
// will be deleted by this method.
//
// The pre-existing data in the internal rune array member
// variable, 'CharsArray' will be deleted and overwritten with new
// data.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) SetRuneArray(
	charArray []rune,
	errorPrefix interface{}) (
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SetRuneArray()",
		"")

	if err != nil {
		return err
	}

	lenOfCharArray := len(charArray)

	if lenOfCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid!\n"+
			"'charArray' is empty and has a length of zero.\n",
			ePrefix.String())

		return err
	}

	charsArrayDto.CharsArray =
		make([]rune, lenOfCharArray)

	for i := 0; i < lenOfCharArray; i++ {
		charsArrayDto.CharsArray[i] =
			charArray[i]
	}

	return err
}
