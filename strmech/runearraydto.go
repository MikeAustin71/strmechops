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
// Each instance is capable of performing three types of text
// character searches:
//   (1) Linear Target Character Search
//        Each Character in the Target String must match each
//        corresponding character in the RuneArrayDto instance
//        beginning at the specified Target String Starting
//        Index.
//
//        Linear Target Character Search Type is the default
//        setting for RuneArrayDto.
//
//   (2) Single Target Character Search
//       A single character in the Target String must mach any
//       character is the RuneArrayDto rune array. A single
//       character in the Target String is therefore compared
//       against all individual characters in the RuneArrayDto
//       instance searching for the first case of a match. Only
//       one designated character in the Target String is included
//       in the search operation.
//
//   (3) LinearEndOfString
//       Designates the search type as a Linear End Of String Search.
//       With this type of search operation, the entire Target Search
//       String will be searched from left to right for the
//       first occurrence of the Test String.
//
//       The search will begin the Target String Starting Index and
//       proceed left to right until (1) an instance of the entire
//       Test String is located or (2) the end of the Target Search
//       String is encountered.
//
//       This is a linear search, so a 'Match' requires that each
//       character in Target Search String must correspond to a
//       matching character in the Test String.
//
//            Example
//                                       1         2         3
//                      Index  0123456789012345678901234567890
//             Target String: "Hey, Xray-4 is the call sign."
//             Target String Starting Index: 0
//               Test String: "Xray-4"
//
//       In this example of a Linear End of String Search, the search
//       operation will begin comparing corresponding characters in
//       the Target Search String and the Test String beginning at
//       index zero. The comparison will fail at index zero, but the
//       search algorithm will continue attempting to find the Test
//       String at indexes 1,2, 3 & 4. The Test String will be found
//       beginning at index number 5 and the search algorithm will
//       terminate at that point with a successful outcome or 'Match'
//       result.
//
// Reference method: RuneArrayDto.SetCharacterSearchType()
//
type RuneArrayDto struct {
	CharsArray []rune
	// An array of runes used to store text characters.

	Description1 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of RuneArrayDto.

	Description2 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of RuneArrayDto.

	charSearchType CharacterSearchType
	// Defines the type of text character search which will be
	// performed on the text characters contained in member
	// variable 'CharsArray'. This enumeration value has four
	// possible settings:
	//     CharSearchType.None() - Invalid
	//     CharSearchType.LinearTargetStartingIndex() - Default, Valid
	//     CharSearchType.SingleTargetChar()  - Valid
	//     CharSearchType.LinearEndOfString() - Valid
	// This variable may only be set by calling method:
	//   RuneArrayDto.SetCharacterSearchType

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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
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

	charsArrayDto.Description1 =
		incomingChars.Description1

	charsArrayDto.Description2 =
		incomingChars.Description2

	charsArrayDto.charSearchType =
		incomingChars.charSearchType

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

// Empty - Resets all internal member variables for the current
// instance of RuneArrayDto to their initial or zero values.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of RuneArrayDto.
///
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

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	runeArrayDtoAtom{}.ptr().empty(
		charsArrayDto)

	charsArrayDto.lock.Unlock()

	charsArrayDto.lock = nil
}

// EmptyCharsArray - This method will empty or delete the contents
// of the internal rune array maintained by the current instance of
// RuneArrayDto.
//
// Only the rune array will be deleted and overwritten. All other
// internal member variables will remain unchanged.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all pre-existing characters contained in
// the internal member variable 'RuneArrayDto.CharsArray'.
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
func (charsArrayDto *RuneArrayDto) EmptyCharsArray() {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	runeArrayDtoElectron{}.ptr().emptyCharsArray(
		charsArrayDto)
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

	return runeArrayDtoAtom{}.ptr().equal(
		charsArrayDto,
		incomingRuneArrayDto)
}

// EqualCharArrays - Compares the current instance of RuneArrayDto
// to another instance ('incomingRuneArrayDto') in order to
// determine whether their internal character rune arrays are
// equivalent. Basically, this is a comparison between the text
// characters encapsulated by 'incomingRuneArrayDto' and the
// current instance of RuneArrayDto, 'charsArrayDto'.
//
// This method will compare the member variables 'CharArrays'
// contained within the two instances of RuneArrayDto. If both
// instances of RuneArrayDto have equivalent member variable
// 'CharArrays', this method will return a boolean value of
// 'true'.
//
// If the rune arrays differ in any respect, this method will
// return a value of 'false'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  incomingRuneArrayDto       *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. The text
//       characters contained in internal member variable
//       rune array, 'CharsArray', will be compared to those
//       contained in the current instance of RuneArrayDto,
//       'charsArrayDto'. If they are equal in all respects,
//       this method will return a boolean value of 'true'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - This boolean flag will transmit the results of a
//       comparison between the character rune arrays contained in
//       input parameter 'incomingRuneArrayDto' and the current
//       instance of RunArrayDto. If the comparison shows that the
//       member variables 'CharsArray' are equivalent in all
//       respects, this return value will be set to 'true'.
//
//       If the member variable 'CharsArray' differs in any way
//       between instances 'incomingRuneArrayDto' and the current
//       instance of RuneArrayDto, this will return 'false'.
//
func (charsArrayDto *RuneArrayDto) EqualCharArrays(
	incomingRuneArrayDto *RuneArrayDto) bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return runeArrayDtoElectron{}.ptr().
		equalCharArrays(
			charsArrayDto,
			incomingRuneArrayDto)
}

// GetCharSearchType - Returns the value of internal member
// variable, 'charSearchType', or Character Search Type.
//
// The Character Search Type is an enumeration value used to
// specify the type target string search algorithm applied by the
// current instance of RuneArrayDto.
//
//       The Character Search Type must be set to one of the
//       following enumeration values:
//
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
// Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
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
//   TextCharSearchType.SingleTargetChar()
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
//   TextCharSearchType.LinearEndOfString()
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
func (charsArrayDto *RuneArrayDto) GetCharSearchType() CharacterSearchType {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return charsArrayDto.charSearchType
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

// GetCharacterString - Returns a string populated with the text
// characters taken from the rune array internal member variable
// contained in the current instance of RuneArrayDto.
//
// The internal member variable is identified as
//   RuneArrayDto.CharsArray
//
// This method is identical in functionality to method:
//  RuneArrayDto.String()
//
// RuneArrayDto.String() implements the Stringer interface.
//
//
// ----------------------------------------------------------------
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
//  string
//     - This method returns a string of characters extracted from
//       the internal member variable rune array
//       'RuneArrayDto.CharsArray'.
//
func (charsArrayDto *RuneArrayDto) GetCharacterString() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return string(charsArrayDto.CharsArray)
}

// IsValidCharacterSearchType - Returns a boolean value signaling
// whether Character Search Type member variable ('charSearchType')
// for this RuneArrayDto instance is valid.
//
// The Character Search Type variable is styled as:
//   runeArrayDto.charSearchType
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If member variable 'charSearchType' (Character Search
//       Type) for the current RuneArrayDto instnce is judged to be
//       valid in all respects, this return parameter will be set
//       to 'true'.
//
//       If member variable 'charSearchType' (Character Search
//       Type) is found to be invalid, this return parameter will
//       be set to 'false'.
//
func (charsArrayDto *RuneArrayDto) IsValidCharacterSearchType() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	isValid,
		_ := runeArrayDtoQuark{}.ptr().isValidCharacterSearchType(
		charsArrayDto,
		nil)

	return isValid
}

// NewNumericCharacters - Returns a new instance of RuneArrayDto
// configured with numeric characters (0 through 9 inclusive).
//
//
// An array of numeric digits in useful in search operations looking
// for numeric characters classified as delimiters.
//
// ----------------------------------------------------------------
//
// Default Settings
//
// The character search type for the returned instance of
// RuneArrayDto is automatically set to:
//
//    CharSearchType.SingleTargetChar()
//
//  - A Single Target Character Search Type means that a single
//    character in the Target Search String will be compared to
//    all characters in the Test String.
//
//    If a single Target String character equals any character in
//    the Test String, a 'Match' or successful search outcome will
//    be declared.
//
//    The search operation is limited to a single designated Target
//    Search String character. Each and every one of the Test
//    String Characters will be compared to this single designated
//    Target String Search Character. The search operation will
//    terminate when a matching character is first identified in
//    the Test String or when the end of the Test String is
//    encountered.
//
//      Example #1
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//       Target String Starting Index: 5
//         Test String: "ZFXyURJK"
//
//    In this example of a Single Target Character Search, the
//    search will begin and end at Target Search String index
//    number 5. Since one of the Test String Characters ('X')
//    matches the 'X' character at index number 5 in the Target
//    Search String, the search operation is classified as a
//    success. A matching character was found.
//
//      Example #2
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//       Target String Starting Index: 0
//         Test String: "ZFXyURJK"
//
//    In this second example of a Single Target Character Search,
//    the search will begin and end at Target Search String index
//    number 0. Since NONE of the Test String Characters matches
//    the 'H' character at index number 0 in the Target Search
//    String, the search operation is classified as a failure. No
//    matching character was found.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  RuneArrayDto
//     - This method returns an instance of RuneArrayDto configured
//       with upper and lower case letters of the Latin Alphabet -
//       English Version (a-z, A-Z).
//
func (charsArrayDto RuneArrayDto) NewNumericCharacters() RuneArrayDto {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	newRuneArrayDto := RuneArrayDto{}

	newRuneArrayDto.CharsArray = []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9'}

	newRuneArrayDto.charSearchType = CharSearchType.SingleTargetChar()

	return newRuneArrayDto
}

func (charsArrayDto RuneArrayDto) NewLatinAlphabet() RuneArrayDto {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	newRuneArrayDto := RuneArrayDto{}

	newRuneArrayDto.CharsArray = []rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z'}

	newRuneArrayDto.charSearchType = CharSearchType.SingleTargetChar()

	return newRuneArrayDto
}

// NewRunes - Receives an array of runes and returns a new
// instance of RuneArrayDto.
//
// The calling function is also required to provide a valid
// Character Search Type as an input parameter.
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
//  charSearchType             CharacterSearchType
//     - An enumeration value used to specify the type target
//       string search algorithm applied by the returned instance
//       of RuneArrayDto.
//
//
//       The Character Search Type must be set to one of the
//       following enumeration values:
//
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewRunes(
	charArray []rune,
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
			"NewRunes()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(charArray) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'charArray' is invalid!\n" +
			"'charArray' is an empty runes array and has an\n" +
			"array length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

	}

	if !charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charSearchType' is invalid!\n"+
			"'charSearchType' must be set to one of three enumerations:\n"+
			"(1) CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"(2) CharacterSearchType(0).SingleTargetChar()\n"+
			"(3) CharacterSearchType(0).LinearEndOfString()\n"+
			"'charSearchType' string  value = '%v'\n"+
			"'charSearchType' integer value = '%v'\n",
			ePrefix.String(),
			charSearchType.String(),
			charSearchType.XValueInt())

		return newRuneArrayDto, err
	}

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	if err != nil {
		return newRuneArrayDto, err
	}

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewRunesAllParams - Receives an array of runes and returns a new
// instance of RuneArrayDto. This method also requires input
// parameters for all internal member variables in the Rune Array
// type.
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
//  description1               string
//
//     - Users have the option of configuring a text string to
//       describe the function or purpose of the text characters
//       configured for the returned instance of RuneArrayDto. This
//       parameter configures the first of two description strings.
//
//
//  description2               string
//
//     - Users have the option of configuring a text string to
//       describe the function or purpose of the text characters
//       configured for the returned instance of RuneArrayDto. This
//       parameter configures the second of two description strings.
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
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewRunesAllParams(
	charArray []rune,
	description1 string,
	description2 string,
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
			"NewRunesAllParams()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(charArray) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'charArray' is invalid!\n" +
			"'charArray' is an empty runes array and has an\n" +
			"array length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

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

		return newRuneArrayDto, err
	}

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	if err != nil {
		return newRuneArrayDto, err
	}

	newRuneArrayDto.Description1 = description1

	newRuneArrayDto.Description2 = description2

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewRunesPtr - Receives an array of runes and returns a new
// instance of RuneArrayDto.
//
// The calling function is also required to provide a valid
// Character Search Type as an input parameter.
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
//  charSearchType             CharacterSearchType
//     - An enumeration value used to specify the type target
//       string search algorithm applied by the returned instance
//       of RuneArrayDto.
//
//
//       The Character Search Type must be set to one of the
//       following enumeration values:
//
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            *RuneArrayDto
//     - If this method completes successfully a pointer to a new,
//       fully populated instance of RuneArrayDto will be returned
//       to the calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewRunesPtr(
	charArray []rune,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	newRuneArrayDto *RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newRuneArrayDto = &RuneArrayDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"NewRunes()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(charArray) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'charArray' is invalid!\n" +
			"'charArray' is an empty runes array and has an\n" +
			"array length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

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

		return newRuneArrayDto, err
	}

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	if err != nil {
		return newRuneArrayDto, err
	}

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewString - Creates and returns a new instance of RuneArrayDto
// based on an input parameter of type string.
//
// The calling function is also required to provide a valid
// Character Search Type as an input parameter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  stringChars                string
//     - A character string used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
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
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
//
func (charsArrayDto RuneArrayDto) NewString(
	stringChars string,
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
			"NewStringAllParams()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(stringChars) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'stringChars' is invalid!\n" +
			"'stringChars' is an empty string and has a string\n" +
			"length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

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

		return newRuneArrayDto, err
	}

	charArray := []rune(stringChars)

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewStringAllParams - Creates and returns a new instance of
// RuneArrayDto using a string input parameter. This method also
// requires input parameters for all internal member variables
// contained in the RuneArrayDto type.
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
//  description1               string
//
//     - Users have the option of configuring a text string to
//       describe the function or purpose of the text characters
//       configured for the returned instance of RuneArrayDto. This
//       parameter configures the first of two description strings.
//
//
//  description2               string
//
//     - Users have the option of configuring a text string to
//       describe the function or purpose of the text characters
//       configured for the returned instance of RuneArrayDto. This
//       parameter configures the second of two description
//       strings.
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
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewStringAllParams(
	stringChars string,
	description1 string,
	description2 string,
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
			"NewStringAllParams()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(stringChars) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'stringChars' is invalid!\n" +
			"'stringChars' is an empty string and has a string\n" +
			"length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

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

		return newRuneArrayDto, err
	}

	charArray := []rune(stringChars)

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	newRuneArrayDto.Description1 = description1

	newRuneArrayDto.Description2 = description2

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewStringPtr - Creates and returns a pointer to a new instance
// of RuneArrayDto based on an input parameter of type string.
//
// The calling function is also required to provide a valid
// Character Search Type as an input parameter.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  stringChars                string
//     - A character string used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
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
//        TextCharSearchType.LinearTargetStartingIndex()
//        TextCharSearchType.SingleTargetChar()
//        TextCharSearchType.LinearEndOfString()
//
//    Character Search Type Options
//
//    TextCharSearchType.LinearTargetStartingIndex()
//    - Designates the search type as a Linear Target Starting
//      Index Search Type. This means that each character in the
//      Target Search String will be compared to each corresponding
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
//      declared, if and only if, the search begins at Target
//      String index number 5. If the search begins at an any index
//      other than 5, no match will be declared and the search will
//      be classified as unsuccessful.
//
//      NOTE: Linear Target Starting Index is the default search
//            type.
//
//
//   TextCharSearchType.SingleTargetChar()
//    - Designates the search type as a Single Target Character
//      Search Type. This means that a single character in the
//      Target Search String will be compared to all characters in
//      the Test String.
//
//      If a single Target String character equals any character in
//      the Test String, a 'Match' or successful search outcome
//      will be declared.
//
//      The search will proceed from left to right in the Target
//      String. Each Target String Character will be compared to
//      all characters in the Test String looking for the first
//      matching Test String Character. The search will terminate
//      when a matching character is first identified or when the
//      end of the Target String is encountered.
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
//   TextCharSearchType.LinearEndOfString()
//    - Designates the search type as a Linear End Of String
//      Search. With this type of search operation, the entire
//      Target Search String will be searched from left to right
//      for the first occurrence of the Test String.
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
//      In this example of a Linear End of String Search, the
//      search operation will begin comparing corresponding
//      characters in the Target Search String and the Test String
//      beginning at index zero. The comparison will fail at index
//      zero, but the search algorithm will continue attempting to
//      find the Test String at indexes 1,2, 3 & 4. The Test String
//      will be found beginning at index number 5 and the search
//      algorithm will terminate at that point with a successful
//      outcome or 'Match' result.
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
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
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
// ----------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            *RuneArrayDto
//     - If this method completes successfully a pointer to a new,
//       fully populated instance of RuneArrayDto will be returned
//       to the calling function.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
//
func (charsArrayDto RuneArrayDto) NewStringPtr(
	stringChars string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	newRuneArrayDto *RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newRuneArrayDto = &RuneArrayDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"NewStringAllParams()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(stringChars) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'stringChars' is invalid!\n" +
			"'stringChars' is an empty string and has a string\n" +
			"length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

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

		return newRuneArrayDto, err
	}

	charArray := []rune(stringChars)

	err = runeArrayDtoElectron{}.ptr().
		setRuneArray(
			newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// SearchForTextCharacterString - This method performs a search
// operation on a string of text characters passed by the calling
// function ('targetSearchString').
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usually involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//    Target String        - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
// A comparison of text characters contained in the Target Search
// String and the Test String serves as the basis for determining
// a 'Match' condition or successful outcome from a text character
// search algorithm. The specific criterion for determining a
// 'Match' condition vary between the different Character Search
// Types.
//
// Character Search Type is a series of enumeration values
// specifying the type of text character search algorithm applied
// by the current instance of RuneArrayDto.
//
// For this method, the Character Search Type must be set to one of
// the following enumeration values:
//
//  TextCharSearchType.LinearTargetStartingIndex()
//  TextCharSearchType.SingleTargetChar()
//  TextCharSearchType.LinearEndOfString()
//
//
// ----------------------------------------------------------------
//
// In this search operation the Target String is searched beginning
// at the specified Starting Search Index ('startingSearchIndex')
// in the host Target Search String.
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
//  targetInputParms           CharSearchTargetInputParametersDto
//     - This Dto, or Data Transfer Object, contains the input
//       parameters and specifications associated with the Target
//       Search String which will be used in completing the Linear
//       Target Starting Index Search operation.
//
//       The CharSearchTargetInputParametersDto structure consists
//       of the following member variable data elements.
//
//       type CharSearchTargetInputParametersDto struct {
//
//        TargetInputParametersName           string
//         - The Name, Label or descriptive Tag associated with this
//           instance of CharSearchTargetInputParametersDto. If empty,
//           this string will be defaulted to "TargetInputParameters"
//
//        TargetString                        *RuneArrayDto
//         - A pointer to the RuneArrayDto containing the Target
//           Search String text characters used in the search
//           algorithm. Target Characters are compared against
//           Test Characters to determine if a 'Match' condition
//           exists.
//
//        TargetStringName                    string
//         - The label or name of the 'TargetString' parameter.
//           Used in error and informational messages.
//
//        TargetStringLength                  int
//         - Actual number of text characters in the entire
//           Target Search String ('TargetString').
//
//        TargetStringLengthName              string
//         - The label or name of the 'TargetStringLength' parameter.
//           Used in error and informational messages.
//
//        TargetStringStartingSearchIndex     int
//         - The index in 'TargetString' at which the search
//           operation begins.
//
//        TargetStringStartingSearchIndexName string
//         - The label or name of the
//           TargetStringStartingSearchIndex parameter.
//           Used in error and informational messages.
//
//        TargetStringSearchLength            int
//         - The actual number of characters within the Target
//           Search String that are included in the search
//           operation. This value may be less than the actual
//           length of the Target Search String.
//
//        TargetStringSearchLengthName        string
//         - The label or name of the TargetStringSearchLength
//           parameter. Used in error and informational
//           messages.
//
//        TargetStringAdjustedSearchLength    int
//         - The adjusted or corrected Target String Search
//           Length. This value is guaranteed to be equal to or
//           less than the actual Target String Length.
//
//        TargetStringDescription1            string
//         - First of two optional description strings
//           describing the Target Search String in the context
//           of the current search operation.
//
//        TargetStringDescription2            string
//         - Second of two optional description strings
//           describing the Target Search String in the context
//           of the current search operation.
//
//        FoundFirstNumericDigitInNumStr      bool
//         - When set to 'true' this signals that the first
//           numeric digit has been identified in the text
//           characters specified by 'TargetString'
//
//        CollectionTestObjIndex int
//         - If the Test String object resides in a collection
//           of Test String objects, this parameter will record
//           the array index of the current Test String object
//           in the collection.
//
//        TextCharSearchType                  CharacterSearchType
//         - Optional. An enumeration value signaling the type
//           of text character search algorithm used to conduct
//           this search operation. When set to a valid value,
//           this specification will override the search
//           specification contained in the Test Input
//           Parameters Data Transfer Object.
//
//          Valid CharSearch Type values are listed as follows:
//          TextCharSearchType.None() - Invalid Value
//          TextCharSearchType.LinearTargetStartingIndex() - Default
//          TextCharSearchType.SingleTargetChar()
//          TextCharSearchType.LinearEndOfString()
//       }
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  CharSearchResultsDto
//     - If this method completes successfully, a fully populated
//       instance of CharSearchResultsDto will be returned to the
//       calling function.
//
//       This structure contains information and data elements
//       detailing the results of the text character search
//       operation performed by this method.
//
//       The CharSearchTestInputParametersDto structure consists
//       of the following member variable data elements.
//
//       type CharSearchResultsDto struct {
//
//        SearchResultsName string
//         - Optional. The Name, Label or descriptive Tag associated with
//           the current instance of CharSearchResultsDto.
//
//        SearchResultsFunctionChain string
//         - Optional. A listing of the functions which led to the
//           creation of this instance of CharSearchResultsDto.
//
//        FoundSearchTarget bool
//         - Signals a successful search outcome. If set to 'true' the
//           Test String character or characters were found in the Target
//           Search String.
//
//        FoundSearchTargetOnPreviousSearch bool
//         - Signals that the Search Target was located in a previous
//           search operation.
//
//        FoundFirstNumericDigitInNumStr bool
//         - When set to 'true' this signals that the search operation
//           has identified the first numeric digit in a string of text
//           characters.
//
//        TargetInputParametersName string
//         - The Name, Label or descriptive Tag associated with an
//           instance of CharSearchTargetInputParametersDto.
//
//        TargetStringLength int
//         - Actual number of text characters in the entire
//           Target Search String ('TargetString').
//
//        TargetStringSearchLength int
//         - The actual number of characters within the Target
//           Search String that are included in the search
//           operation. This value may be less than the actual
//           length of the Target Search String.
//
//        TargetStringAdjustedSearchLength int
//         - The adjusted or corrected Target String Search
//           Length. This value is guaranteed to be equal to or
//           less than the actual Target String Length.
//
//        TargetStringStartingSearchIndex int
//         - The index in 'TargetString' at which the search
//           operation begins.
//
//        TargetStringFirstFoundIndex int
//         - The index of the first character position in the
//           Target Search String occupied by the first
//           character in the Test String.
//
//        TargetStringLastFoundIndex int
//         - The last character position in the Target Search
//           String occupied by the last character in the Test
//           String.
//
//        TargetStringLastSearchIndex int
//         - The index in Target Search String occupied by the
//           last Target character searched. If the Search
//           Target was found, this value is equal to the
//           'TargetStringLastFoundIndex'. If the Search Target
//           was NOT found this value is equal to the
//           'TargetStringStartingSearchIndex'. This value is
//           useful in computing the next index to be searched
//           in the Target String.
//
//        TargetStringNextSearchIndex int
//         - The starting point for the next search operation.
//           If the entire Target String was included in the
//           last search, this value is set to -1.
//
//            Example-1:
//            String = "Hello"
//            String Length = 5
//            Last Search Index = 4
//            TargetStringNextSearchIndex = -1
//
//            Example-2:
//            String = "Hello"
//            String Length = 5
//            Last Search Index = 2
//            TargetStringNextSearchIndex = 3
//
//        TargetStringDescription1 string
//         - First of two optional description strings
//           describing the Target Search String in the context
//           of the current search operation.
//
//        TargetStringDescription2 string
//         - Second of two optional description strings
//           describing the Target Search String in the context
//           of the current search operation.
//
//        TestInputParametersName string
//         - The Name, Label or descriptive Tag associated with an
//           instance of CharSearchTestInputParametersDto.
//
//        TestStringLength int
//         - Actual number of text characters in the entire Test
//           String ('TestString').
//
//        TestStringStartingIndex int
//         - The starting index in the Test String where the
//           search operation will begin.
//
//        TestStringFirstFoundIndex int
//         - The index number in Test String of the first test
//           character to be located in the Target Search String.
//
//        TestStringLastFoundIndex int
//         - The index number in the Test String occupied by the
//           last Test Character to be located in the Target
//           String.
//
//        TestStringDescription1 string
//         - First of two optional description strings
//           describing the Test String in the context of the
//           current search operation.
//
//        TestStringDescription2 string
//         - Second of two optional description strings
//           describing the Test String in the context of the
//           current search operation.
//
//        CollectionTestObjIndex int
//         - If the Test String object resides in a collection
//           of Test String objects, this parameter will record
//           the array index of the current Test String object
//           in the collection.
//
//        ReplacementString *RuneArrayDto
//         - A pointer to the Rune Array Data Transfer Object
//           containing the Replacement Characters to be
//           substituted for existing characters in a Target
//           String.
//
//        RemainderString *RuneArrayDto
//         - A pointer to the Rune Array Data Transfer Object
//           containing the remaining characters in a Target
//           String which were NOT included in the search
//           operation and which remain to be searched in future
//           search operations. This string is also used in 'cut'
//           operations where Target String is divided based on
//           string delimiters.
//
//        NumValueType NumericValueType
//         - Optional. This enumeration value specifies the type
//           of numeric value for this Test Parameter as either
//           an integer or floating point value.
//
//           Possible values are listed as follows:
//            NumValType.None()
//            NumValType.FloatingPoint()
//            NumValType.Integer()
//
//        NumStrFormatType NumStrFormatTypeCode
//         - Optional. This enumeration value specifies the
//           Output Format Type for a number.
//
//           Possible values are listed as follows:
//            NumStrFmtType.None()
//            NumStrFmtType.AbsoluteValue()
//            NumStrFmtType.Binary()
//            NumStrFmtType.CountryCulture()
//            NumStrFmtType.Currency()
//            NumStrFmtType.Binary()
//            NumStrFmtType.Hexadecimal()
//            NumStrFmtType.Octal()
//            NumStrFmtType.ScientificNotation()
//
//        NumSymbolLocation NumericSymbolLocation
//         - Optional. This enumeration value specifies the
//           relative location of a numeric symbol.
//
//           Possible values are listed as follows:
//            NumSymLocation.None()
//            NumSymLocation.Before()
//            NumSymLocation.Interior()
//            NumSymLocation.After()
//
//        NumSymbolClass NumericSymbolClass
//         - Optional. This enumeration value specifies the
//           Number Symbol Classification for a text character.
//
//           Possible values are listed as follows:
//            NumSymClass.None()
//            NumSymClass.NumberSign()
//            NumSymClass.CurrencySign()
//            NumSymClass.IntegerSeparator()
//            NumSymClass.DecimalSeparator()
//
//        NumSignValue NumericSignValueType
//         - Optional. This enumeration value specifies the
//           number sign value.
//
//           Possible values are listed as follows:
//            NumSignVal.None()
//            NumSignVal.Negative()
//            NumSignVal.Zero()
//            NumSignVal.Positive()
//
//        PrimaryNumSignPosition NumSignSymbolPosition
//         - Optional This enumeration value specifies the
//           relative position of positive and negative number
//           sign symbols in a number string. This is the
//           Primary Type Code for Number Signs. Cases involving
//           both 'Leading' and 'Trailing' symbols also make use
//           of the 'SecondaryNumSignPosition'.
//
//           Possible values are listed as follows:
//            NumSignSymPos.None()
//            NumSignSymPos.Before()
//            NumSignSymPos.After()
//            NumSignSymPos.BeforeAndAfter()
//
//        SecondaryNumSignPosition NumSignSymbolPosition
//         - Optional. This enumeration value specifies the
//           relative position of positive and negative number
//           sign symbols in a number string. This value is used
//           in searches involving number signs which occur both
//           before and after the numeric value.
//
//           Possible values are listed as follows:
//            NumSignSymPos.None()
//            NumSignSymPos.Before()
//            NumSignSymPos.After()
//            NumSignSymPos.BeforeAndAfter()
//
//        TextCharSearchType CharacterSearchType
//         - Required. An enumeration value signaling the type
//           of text character search algorithm used to conduct
//           this search operation.
//
//           Possible values are listed as follows:
//            CharSearchType.None() - Invalid value
//            CharSearchType.LinearTargetStartingIndex() - Default
//            CharSearchType.SingleTargetChar()
//            CharSearchType.LinearEndOfString()
//       }
//
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
func (charsArrayDto *RuneArrayDto) SearchForTextCharacterString(
	targetInputParms CharSearchTargetInputParametersDto,
	testConfigDto CharSearchTestConfigDto,
	errorPrefix interface{}) (
	CharSearchRuneArrayResultsDto,
	error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	errorResults := CharSearchRuneArrayResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SearchForTextCharacterString()",
		"")

	if err != nil {

		return errorResults, err

	}

	testInputParms := CharSearchTestInputParametersDto{}

	testInputParms.Empty()

	testInputParms.TestInputParametersName =
		"RuneArrayDto"

	testInputParms.LoadTestConfigDto(
		testConfigDto)

	testInputParms.TextCharSearchType =
		CharSearchType.None()

	err = targetInputParms.ValidateCharSearchType(
		ePrefix)

	if err == nil {
		// This is an override condition. The
		// Target Search String is specifying the
		// Character Search Type.
		testInputParms.TextCharSearchType =
			targetInputParms.TextCharSearchType
	} else {
		// err != nil

		if !charsArrayDto.charSearchType.XIsValid() {
			err = fmt.Errorf("%v\n"+
				"Error: Both Input Parameters Search Type and\n"+
				"RuneArrayDto Search Type are invalid.\n"+
				"The Search Operation is Terminating!\n"+
				"Input Parameters Search Type String  = '%v'\n"+
				"Input Parameters Search Type Integer = '%v'\n"+
				"RuneArrayDto Search Type String  = '%v'\n"+
				"RuneArrayDto Search Type Integer = '%v'\n",
				ePrefix.String(),
				testInputParms.TextCharSearchType.String(),
				testInputParms.TextCharSearchType.XValueInt(),
				charsArrayDto.charSearchType.String(),
				charsArrayDto.charSearchType.XValueInt())
			return errorResults, err

		}

		testInputParms.TextCharSearchType =
			charsArrayDto.charSearchType
	}

	testInputParms.TestString = charsArrayDto

	testInputParms.TestStringName = "RuneArrayDto.CharsArray"

	testInputParms.TestStringLengthName =
		"RuneArrayDto Length"

	testInputParms.TestStringStartingIndex = 0

	err = testInputParms.ValidateTestParameters(
		ePrefix)

	if err != nil {

		return errorResults, err

	}

	runeArrayNanobot := runeArrayDtoNanobot{}

	return runeArrayNanobot.characterSearchExecutor(
		targetInputParms,
		testInputParms,
		ePrefix)

}

// SetCharacterSearchType - Sets the internal member variable used
// to track the Text Character Search Type.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usually involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//    Target String        - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
// A comparison of text characters contained in the Target Search
// String and the Test String serves as the basis for determining
// a 'Match' condition or successful outcome from a text character
// search algorithm. The specific criterion for determining a
// 'Match' condition vary between the different Character Search
// Types.
//
// Character Search Type is a series of enumeration values
// specifying the type of text character search algorithm applied
// by the current instance of RuneArrayDto.
//
// For this method, the Character Search Type must be set to one of
// the following enumeration values:
//
//  TextCharSearchType.LinearTargetStartingIndex()
//  TextCharSearchType.SingleTargetChar()
//  TextCharSearchType.LinearEndOfString()
//
//
// ----------------------------------------------------------------
//
// Character Search Type Options
//
//  TextCharSearchType.LinearTargetStartingIndex()
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
// TextCharSearchType.SingleTargetChar()
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
// TextCharSearchType.LinearEndOfString()
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
// For the current instance of RuneArrayDto, all pre-existing data
// in the internal rune array member variable, 'CharsArray' will be
// deleted and overwritten with new data.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes used to populate the internal member
//       variable rune array for the current instance of
//       RuneArrayDto.
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
	errorPrefix interface{}) error {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SetRuneArray()",
		"")

	if err != nil {
		return err
	}

	return runeArrayDtoElectron{}.ptr().
		setRuneArray(
			charsArrayDto,
			charArray,
			ePrefix.XCpy(
				"charsArrayDto"))
}

// SetRuneArrayFromString - Receives a string as an input parameter
// and proceeds to populate the internal rune array with new
// character data for the current instance of RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// For the current instance of RuneArrayDto, all pre-existing data
// in the internal rune array member variable, 'CharsArray' will be
// deleted and overwritten with new data.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charString                 string
//     - An string of characters which will be used to populate
//       the internal member variable rune array for the current
//       instance of RuneArrayDto.
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
func (charsArrayDto *RuneArrayDto) SetRuneArrayFromString(
	charString string,
	errorPrefix interface{}) error {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SetRuneArrayFromString()",
		"")

	if err != nil {
		return err
	}

	charArray := []rune(charString)

	return runeArrayDtoElectron{}.ptr().
		setRuneArray(
			charsArrayDto,
			charArray,
			ePrefix.XCpy(
				"charsArrayDto"))
}

// String - Returns the internal rune array maintained by this
// instance of RuneArrayDto as a string.
//
// This method implements the Stringer interface.
//
//
// ----------------------------------------------------------------
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
//  string
//     - This method returns a string of characters extracted from
//       the internal member variable rune array
//       'RuneArrayDto.CharsArray'.
//
func (charsArrayDto *RuneArrayDto) String() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return string(charsArrayDto.CharsArray)
}
