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
//
//	(1) Linear Target Character Search
//	     Each Character in the Target String must match each
//	     corresponding character in the RuneArrayDto instance
//	     beginning at the specified Target String Starting
//	     Index.
//
//	     Linear Target Character Search Type is the default
//	     setting for RuneArrayDto.
//
//	(2) Single Target Character Search
//	    A single character in the Target String must mach any
//	    character is the RuneArrayDto rune array. A single
//	    character in the Target String is therefore compared
//	    against all individual characters in the RuneArrayDto
//	    instance searching for the first case of a match. Only
//	    one designated character in the Target String is included
//	    in the search operation.
//
//	(3) LinearEndOfString
//	    Designates the search type as a Linear End Of String Search.
//	    With this type of search operation, the entire Target Search
//	    String will be searched from left to right for the
//	    first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the search
//	    operation will begin comparing corresponding characters in
//	    the Target Search String and the Test String beginning at
//	    index zero. The comparison will fail at index zero, but the
//	    search algorithm will continue attempting to find the Test
//	    String at indexes 1,2, 3 & 4. The Test String will be found
//	    beginning at index number 5 and the search algorithm will
//	    terminate at that point with a successful outcome or 'Match'
//	    result.
//
// Reference method: RuneArrayDto.SetCharacterSearchType()
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

//	AddChar
//
//	Adds a single text character of type rune to the end
//	of the rune array (RuneArrayDto.CharsArray)
//	encapsulated by the current instance of RuneArrayDto.
//
//	Another method which is similar in function is:
//
//			RuneArrayDto.ExtendRuneArray()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	charToAdd					rune
//
//		The text character which will be appended to the
//		end of the rune array encapsulated by the current
//		instance of RuneArrayDto.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
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

//	AddString
//
//	Receives a string of text characters and proceeds to
//	add them to the existing rune array contained in the
//	current instance of RuneArrayDto.
//
//	The rune array to which the new characters will be
//	added is specified by member variable:
//
//		RuneArrayDto.CharsArray
//
//	The setting for input parameter 'addTrailingChars'
//	will determine whether the new characters are added
//	to as trailing characters to the end of the existing
//	rune array, or as leading characters as the beginning
//	of the existing rune array.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	charsToAdd					string
//
//		A string containing text characters which be
//		added to the existing rune contained within the
//		current instance of RuneArrayDto:
//
//			RuneArrayDto.CharsArray
//
//	addTrailingChars			bool
//
//		This parameter determines whether 'charsToAdd'
//		will be added to 'RuneArrayDto.CharsArray' as
//		trailing characters or as leading characters.
//
//		If 'addTrailingChars' is set to 'true',
//		'charsToAdd' will be added to the end of the
//		existing rune array as trailing characters.
//
//		If 'addTrailingChars' is set to 'false',
//		'charsToAdd' will be added to the beginning of
//		the existing rune array as leading characters.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (charsArrayDto *RuneArrayDto) AddString(
	charsToAdd string,
	addTrailingChars bool,
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
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(runeArrayDtoAtom).addRunes(
		charsArrayDto,
		[]rune(charsToAdd),
		addTrailingChars,
		ePrefix.XCpy(
			"charsArrayDto<-"))
}

//	CopyIn
//
//	Copies the internal rune array from an incoming
//	instance of RuneArrayDto ('incomingChars') to the
//	internal rune array of the current RuneArrayDto
//	instance ('charsArrayDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The internal rune array for the current RuneArrayDto
//	instance ('charsArrayDto') will be deleted and
//	overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingChars				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		method will NOT change the values of internal
//		rune array contained in this instance.
//
//		All data values in this RuneArrayDto instance
//		will be copied to current RuneArrayDto instance
//		('charsArrayDto').
//
//		If parameter 'incomingChars' is 'nil' pointer or
//		if its internal rune array has a length of zero,
//		an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
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

	return new(runeArrayDtoNanobot).copyRuneArrayDto(
		charsArrayDto,
		incomingChars,
		ePrefix.XCpy(
			"charsArrayDto<-incomingChars"))
}

// CopyOut - Returns a deep copy of the current RuneArrayDto
// instance.
//
// If the current RuneArrayDto instance contains an empty rune
// array, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	copyOfRuneArrayDto         copyOfRuneArrayDto
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current RuneArrayDto instance.
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(runeArrayDtoNanobot).copyRuneArrayDto(
		&copyOfRuneArrayDto,
		charsArrayDto,
		ePrefix.XCpy(
			"copyOfRuneArrayDto<-charsArrayDto"))

	return copyOfRuneArrayDto, err
}

// Empty - Resets all internal member variables for the current
// instance of RuneArrayDto to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of RuneArrayDto.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (charsArrayDto *RuneArrayDto) Empty() {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	new(runeArrayDtoAtom).empty(
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
// # IMPORTANT
//
// This method will delete all pre-existing characters contained in
// the internal member variable 'RuneArrayDto.CharsArray'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (charsArrayDto *RuneArrayDto) EmptyCharsArray() {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	new(runeArrayDtoElectron).emptyCharsArray(
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingRuneArrayDto       *RuneArrayDto
//	   - A pointer to an external instance of RuneArrayDto. The
//	     internal member variable data values in this instance will
//	     be compared to those in the current instance of
//	     RuneArrayDto. The results of this comparison will
//	     be returned to the calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the internal member variable data values contained in
//	     input parameter 'incomingRuneArrayDto' are equivalent in
//	     all respects to those contained in the current instance of
//	     RuneArrayDto, this return value will be set to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (charsArrayDto *RuneArrayDto) Equal(
	incomingRuneArrayDto *RuneArrayDto) bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return new(runeArrayDtoAtom).equal(
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	incomingRuneArrayDto       *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. The text
//	     characters contained in internal member variable
//	     rune array, 'CharsArray', will be compared to those
//	     contained in the current instance of RuneArrayDto,
//	     'charsArrayDto'. If they are equal in all respects,
//	     this method will return a boolean value of 'true'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - This boolean flag will transmit the results of a
//	     comparison between the character rune arrays contained in
//	     input parameter 'incomingRuneArrayDto' and the current
//	     instance of RunArrayDto. If the comparison shows that the
//	     member variables 'CharsArray' are equivalent in all
//	     respects, this return value will be set to 'true'.
//
//	     If the member variable 'CharsArray' differs in any way
//	     between instances 'incomingRuneArrayDto' and the current
//	     instance of RuneArrayDto, this will return 'false'.
func (charsArrayDto *RuneArrayDto) EqualCharArrays(
	incomingRuneArrayDto *RuneArrayDto) bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return new(runeArrayDtoElectron).
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
//	The Character Search Type must be set to one of the
//	following enumeration values:
//
//	 CharSearchType.LinearTargetStartingIndex()
//	 CharSearchType.SingleTargetChar()
//	 CharSearchType.LinearEndOfString()
//
// Character Search Type Options
//
//	 CharSearchType.LinearTargetStartingIndex()
//	 - Designates the search type as a Linear Target Starting Index
//	   Search Type. This means that each character in the Target
//	   Search String will be compared to each corresponding
//	   character in the Test String beginning at a specified
//	   starting index in the Target Search String.
//
//	   The search will proceed for from left to right in Test
//	   Character Sequence.
//
//	   If the Test Characters are NOT found in the Target Search
//	   String beginning at the designated Target String Starting
//	   Index, the search outcome will be unsuccessful and NO match
//	   will be declared.
//
//	   A 'Match', or successful search outcome, is defined as the
//	   case where each character in the Target String matches each
//	   corresponding character in the Test String beginning at the
//	   designated Target String Starting Index.
//
//
//	     Example
//	                               1         2         3
//	              Index  0123456789012345678901234567890
//	     Target String: "Hey, Xray-4 is the call sign."
//	     Target String Starting Index: 5
//	       Test String: "Xray"
//
//	   In this example of a Linear Target Starting Index Search, a
//	   match between the Target String and Test String will be
//	   declared, if and only if, the search begins at Target String
//	   index number 5. If the search begins at an any index other
//	   than 5, no match will be declared and the search will be
//	   classified as unsuccessful.
//
//	   NOTE: Linear Target Starting Index is the default search
//	         type.
//
//
//	CharSearchType.SingleTargetChar()
//	 - Designates the search type as a Single Target Character
//	   Search Type. This means that a single character in the Target
//	   Search String will be compared to all characters in the Test
//	   String.
//
//	   If a single Target String character equals any character in
//	   the Test String, a 'Match' or successful search outcome will
//	   be declared.
//
//	   The search will proceed from left to right in the Target
//	   String. Each Target String Character will be compared to all
//	   characters in the Test String looking for the first matching
//	   Test String Character. The search will terminate when a
//	   matching character is first identified or when the end of the
//	   Target String is encountered.
//
//
//	     Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 0
//	        Test String: "ZFXyURJK"
//
//	   In this example of a Single Target Character Search, the
//	   search will terminate at Target String index numbers 5
//	   because it is the first Target String index to match one
//	   of the Test String Characters ('X').
//
//
//	CharSearchType.LinearEndOfString()
//	 - Designates the search type as a Linear End Of String Search.
//	   With this type of search operation, the entire Target Search
//	   String will be searched from left to right for the
//	   first occurrence of the Test String.
//
//	   The search will begin the Target String Starting Index and
//	   proceed left to right until (1) an instance of the entire
//	   Test String is located or (2) the end of the Target Search
//	   String is encountered.
//
//	   This is a linear search, so a 'Match' requires that each
//	   character in Target Search String must correspond to a
//	   matching character in the Test String.
//
//	        Example
//	                                   1         2         3
//	                  Index  0123456789012345678901234567890
//	         Target String: "Hey, Xray-4 is the call sign."
//	         Target String Starting Index: 0
//	           Test String: "Xray-4"
//
//	   In this example of a Linear End of String Search, the search
//	   operation will begin comparing corresponding characters in
//	   the Target Search String and the Test String beginning at
//	   index zero. The comparison will fail at index zero, but the
//	   search algorithm will continue attempting to find the Test
//	   String at indexes 1,2, 3 & 4. The Test String will be found
//	   beginning at index number 5 and the search algorithm will
//	   terminate at that point with a successful outcome or 'Match'
//	   result.
//
//
//	For more information see the source code comments for type,
//	CharacterSearchType.
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
//
//	RuneArrayDto.Description1
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
//
//	RuneArrayDto.Description2
func (charsArrayDto *RuneArrayDto) GetRuneArrayDescription2() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return charsArrayDto.Description2
}

// GetRuneArrayLength
//
// Returns the length of the internal rune array,
// 'CharsArray' as an integer value.
func (charsArrayDto *RuneArrayDto) GetRuneArrayLength() int {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return len(charsArrayDto.CharsArray)
}

// GetRuneArray
//
// Returns a deep copy of the internal rune array,
// 'CharsArray' configured for the current instance of
// RuneArrayDto.
//
// If the length of the internal rune array 'CharsArray'
// is zero, a value of 'nil' will be returned.
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
//
//	RuneArrayDto.CharsArray
//
// This method is identical in functionality to method:
//
//	RuneArrayDto.String()
//
// RuneArrayDto.String() implements the Stringer interface.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string
//	   - This method returns a string of characters extracted from
//	     the internal member variable rune array
//	     'RuneArrayDto.CharsArray'.
func (charsArrayDto *RuneArrayDto) GetCharacterString() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return string(charsArrayDto.CharsArray)
}

//	ExtendRuneArray
//
//	Adds one or more characters to the beginning or end
//	of the rune array contained in the current instance
//	of RuneArrayDto.
//
//	The name of the rune array member variable is:
//
//			RuneArrayDto.CharsArray
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	fillChar					rune
//
//		This Character will be added to the beginning or
//		the end of the rune array contained in the
//		current instance of RuneArrayDto.
//
//	numOfCharsToAdd				int
//
//		'fillChar' will be added to the rune array
//		'numOfCharsToAdd' of times.
//
//		'numOfCharsToAdd' is equal to the total number of
//		characters which will be added to the rune array.
//
//		If the value of 'numOfCharsToAdd' is less than
//		one (1), an error will be returned.
//
//	addTrailingChars			bool
//
//		If this parameter is set to 'true', the additional
//		characters will be appended to the rune arrays as
//		trailing characters. This means the characters
//		will be added to the end of the rune array.
//
//		If this parameter is set to 'false', the
//		additional characters will be added to the rune
//		array as leading characters. In other words, the
//		characters will be added to the beginning of the
//		rune array.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (charsArrayDto *RuneArrayDto) ExtendRuneArray(
	fillChar rune,
	numOfCharsToAdd int,
	addTrailingChars bool,
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
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	if numOfCharsToAdd < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfCharsToAdd' is invalid!\n"+
			"'numOfCharsToAdd' has a value less than one (1).\n"+
			"numOfCharsToAdd = '%v'\n",
			ePrefix.String(),
			numOfCharsToAdd)

		return err
	}

	newCharArrayLength := len(charsArrayDto.CharsArray)

	newCharArrayLength += numOfCharsToAdd

	return new(numStrMathQuark).extendRunes(
		charsArrayDto,
		charsArrayDto,
		fillChar,
		newCharArrayLength,
		addTrailingChars,
		ePrefix.XCpy(
			"charsArrayDto<-"))
}

// IsAllNumericDigits
//
// If the rune array contained in this instance of
// RuneArrayDto consists entirely of numeric character
// digits, this method return a boolean value of 'true'.
//
// If all rune array member elements do NOT consist
// of numeric character digits in the range '0' through
// '9' inclusive, this method returns 'false'.
func (charsArrayDto *RuneArrayDto) IsAllNumericDigits() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return new(runeArrayDtoQuark).
		isRuneArrayAllNumericDigits(
			charsArrayDto)
}

// IsAllNumericZeros
//
// If the rune array contained in this instance of
// RuneArrayDto consists entirely of zero numeric
// character digits ('0'), this method return a boolean
// value of 'true'.
//
// If all rune array member elements do NOT consist
// exclusively of zero numeric character digits ('0'),
// this method returns 'false'.
func (charsArrayDto *RuneArrayDto) IsAllNumericZeros() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return new(runeArrayDtoQuark).
		isRuneArrayAllNumericZeros(
			charsArrayDto)
}

// IsEmpty
//
// Returns a boolean value of 'true' if the Character
// Array for the current RuneArrayDto is empty or has a
// zero length.
//
// A Character Array with a zero length contains zero
// characters.
func (charsArrayDto *RuneArrayDto) IsEmpty() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	if len(charsArrayDto.CharsArray) == 0 {
		return true
	}

	return false
}

//	IsValidCharacterArray
//
//	Returns a boolean value of 'true' if the Character
//	Array for the current RuneArrayDto instance.
//
//	The Character Array of the current RuneArrayDto
//	instance is stored in member variable:
//
//			RuneArrayDto.CharsArray
//
//	The Character array is judged to be valid if it has
//	an array length greater than zero.
//
//	Conversely, character arrays with a length of zero
//	are classified as invalid.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid                    bool
//
//		If the current RuneArrayDto instance member
//		variable 'CharsArray' (Character Array) is judged
//		to be valid in all respects, this return
//		parameter will be set to 'true'.
//
//	   - If the current RuneArrayDto instance member variable
//	     'CharsArray' (Character Array) is judged to be invalid,
//	     this return parameter will be set to 'false'.
func (charsArrayDto *RuneArrayDto) IsValidCharacterArray() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	isValid,
		_ := new(runeArrayDtoQuark).isValidCharacterArray(
		charsArrayDto,
		nil)

	return isValid
}

// IsValidCharacterArrayError - Returns an error value signaling
// whether Character Array for the current RuneArrayDto instance is
// valid.
//
// If the Character Array is judged to be valid, the returned error
// parameter is set to 'nil'.
//
// If the Character Array length is greater than zero, the
// Character Array is classified as valid.
//
// If the Character Array length is equal to zero, an error will be
// returned and this instance of RuneArrayDto will be considered
// invalid.
//
// The Character Array variable is styled as:
//
//	RuneArrayDto.charSearchType
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If the current RuneArrayDto member variable
//	     'charSearchType' (Character Search Type) is judged to be
//	     valid in all respects, this return parameter will be set
//	     to 'nil'.
//
//	     If the current RuneArrayDto member variable
//	     'charSearchType' (Character Search Type) is found to be
//	     invalid, this return parameter will be configured with an
//	     appropriate error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (charsArrayDto *RuneArrayDto) IsValidCharacterArrayError(
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
			"IsValidCharacterArrayError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(runeArrayDtoQuark).isValidCharacterArray(
		charsArrayDto,
		ePrefix.XCpy(
			"charsArrayDto"))

	return err
}

// IsValidCharacterSearchType - Returns a boolean value signaling
// whether Character Search Type member variable ('charSearchType')
// for this RuneArrayDto instance is valid.
//
// The Character Search Type variable is styled as:
//
//	runeArrayDto.charSearchType
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If member variable 'charSearchType' (Character Search
//	     Type) for the current RuneArrayDto instance is judged to be
//	     valid in all respects, this return parameter will be set
//	     to 'true'.
//
//	     If member variable 'charSearchType' (Character Search
//	     Type) is found to be invalid, this return parameter will
//	     be set to 'false'.
func (charsArrayDto *RuneArrayDto) IsValidCharacterSearchType() bool {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	isValid,
		_ := new(runeArrayDtoQuark).isValidCharacterSearchType(
		charsArrayDto,
		nil)

	return isValid
}

// IsValidCharacterSearchTypeError - Returns an error value
// signaling whether Character Search Type member variable
// ('charSearchType') for this RuneArrayDto instance is valid.
//
// If the Character Search Type is judged to be valid, the returned
// error parameter is set to 'nil'.
//
// The Character Search Type variable is styled as:
//
//	runeArrayDto.charSearchType
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If the current RuneArrayDto member variable
//	     'charSearchType' (Character Search Type) is judged to be
//	     valid in all respects, this return parameter will be set
//	     to 'nil'.
//
//	     If the current RuneArrayDto member variable
//	     'charSearchType' (Character Search Type) is found to be
//	     invalid, this return parameter will be configured with an
//	     appropriate error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (charsArrayDto *RuneArrayDto) IsValidCharacterSearchTypeError(
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
			"IsValidCharacterSearchTypeError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(runeArrayDtoQuark).isValidCharacterSearchType(
		charsArrayDto,
		ePrefix.XCpy(
			"charsArrayDto"))

	return err
}

//	NewRunesDefault
//
//	Receives a rune array and returns a new instance
//	of RuneArrayDto.
//
//	The Character Search type for the new
//	RuneArrayDto is defaulted to:
//
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	runeChars	[]rune
//
//		An array of runes which will comprise the
//		rune array in the returned instance of
//		RuneArrayDto.
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	RuneArrayDto
//
//	This method returns a new instance of RuneArrayDto.
//	The internal rune array is set according to text
//	characters supplied by input parameter, 'runeChars'.
//
//	The Character Search type for the new
//	RuneArrayDto is defaulted to:
//		CharSearchType.LinearTargetStartingIndex()
func (charsArrayDto RuneArrayDto) NewRunesDefault(
	runeChars []rune) RuneArrayDto {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	newRuneArray := RuneArrayDto{}

	_ = new(runeArrayDtoElectron).
		setRuneArray(
			&newRuneArray,
			runeChars,
			nil)

	newRuneArray.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return newRuneArray
}

//	NewStringDefault
//
//	Receives a string and returns a new instance
//	of RuneArrayDto.
//
//	The Character Search type for the new
//	RuneArrayDto is automatically defaulted to:
//
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	stringChars					string
//
//	A string of characters which will comprise the rune
//	array in the returned instance of RuneArrayDto.
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	RuneArrayDto
//
//	This method returns a new instance of RuneArrayDto.
//	The internal rune array is set according to text
//	characters supplied by input parameter, 'stringChars'.
//
//	The Character Search type for the new
//	RuneArrayDto is automatically defaulted to:
//
//		CharSearchType.LinearTargetStartingIndex()
func (charsArrayDto RuneArrayDto) NewStringDefault(
	stringChars string) RuneArrayDto {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	newRuneArray := RuneArrayDto{}

	if len(stringChars) == 0 {

		newRuneArray.CharsArray = nil

	} else {

		newRuneArray.CharsArray = []rune(stringChars)

	}

	newRuneArray.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return newRuneArray
}

// NewNumericCharacters - Returns a new instance of RuneArrayDto
// configured with numeric characters (0 through 9 inclusive).
//
// An array of numeric digits in useful in search operations looking
// for numeric characters classified as delimiters.
//
// ----------------------------------------------------------------
//
// # Default Settings
//
// The character search type for the returned instance of
// RuneArrayDto is automatically set to:
//
//	  CharSearchType.SingleTargetChar()
//
//	- A Single Target Character Search Type means that a single
//	  character in the Target Search String will be compared to
//	  all characters in the Test String.
//
//	  If a single Target String character equals any character in
//	  the Test String, a 'Match' or successful search outcome will
//	  be declared.
//
//	  The search operation is limited to a single designated Target
//	  Search String character. Each and every one of the Test
//	  String Characters will be compared to this single designated
//	  Target String Search Character. The search operation will
//	  terminate when a matching character is first identified in
//	  the Test String or when the end of the Test String is
//	  encountered.
//
//	    Example #1
//	                               1         2         3
//	              Index  0123456789012345678901234567890
//	     Target String: "Hey, Xray-4 is the call sign."
//	     Target String Starting Index: 5
//	       Test String: "ZFXyURJK"
//
//	  In this example of a Single Target Character Search, the
//	  search will begin and end at Target Search String index
//	  number 5. Since one of the Test String Characters ('X')
//	  matches the 'X' character at index number 5 in the Target
//	  Search String, the search operation is classified as a
//	  success. A matching character was found.
//
//	    Example #2
//	                               1         2         3
//	              Index  0123456789012345678901234567890
//	     Target String: "Hey, Xray-4 is the call sign."
//	     Target String Starting Index: 0
//	       Test String: "ZFXyURJK"
//
//	  In this second example of a Single Target Character Search,
//	  the search will begin and end at Target Search String index
//	  number 0. Since NONE of the Test String Characters matches
//	  the 'H' character at index number 0 in the Target Search
//	  String, the search operation is classified as a failure. No
//	  matching character was found.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	RuneArrayDto
//	   - This method returns an instance of RuneArrayDto configured
//	     with upper and lower case letters of the Latin Alphabet -
//	     English Version (a-z, A-Z).
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

// NewNumStr - This method is designed to create a new instance of
// RuneArrayDto configured on a number string passed as an input
// parameter.
//
// The Character Search Type is defaulted to
//
//	CharSearchType.LinearTargetStartingIndex()
//
// This is the character search type typically used when converting
// number strings to numeric values.
//
// For more information on character search types see the source
// code comments for type:
//
//	CharacterSearchType.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numStr                     string
//	   - A string of text characters containing numeric digits
//
//	     If this string is empty or has a zero length, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            RuneArrayDto
//	   - If this method completes successfully a new, fully
//	     populated instance of RuneArrayDto will be returned to the
//	     calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (charsArrayDto RuneArrayDto) NewNumStr(
	numStr string,
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
			"NewNumStr()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	if len(numStr) == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameter 'numStr' is invalid!\n" +
			"'numStr' is an empty string and has a string\n" +
			"length of zero (0).\n" +
			ePrefix.String())

		return newRuneArrayDto, err

	}

	charArray := []rune(numStr)

	err = new(runeArrayDtoElectron).
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	newRuneArrayDto.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return newRuneArrayDto, err
}

// NewRunes - Receives an array of runes and returns a new
// instance of RuneArrayDto.
//
// The calling function is also required to provide a valid
// Character Search Type as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - An array of runes used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto. If 'charSearchType' is invalid, an error
//	     will be returned.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex() (Default)
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            RuneArrayDto
//	   - If this method completes successfully a new, fully
//	     populated instance of RuneArrayDto will be returned to the
//	     calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(runeArrayDtoElectron).
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
// instance of RuneArrayDto. This method requires input parameters
// for all internal member variables in the Rune Array type.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - An array of runes used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	description1               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the text characters
//	     configured for the returned instance of RuneArrayDto. This
//	     parameter configures the first of two description strings.
//
//	     If this descriptive text is not required, set this
//	     parameter to an empty string.
//
//
//	description2               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the text characters
//	     configured for the returned instance of RuneArrayDto. This
//	     parameter configures the second of two description strings.
//
//	     If this descriptive text is not required, set this
//	     parameter to an empty string.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex()
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            RuneArrayDto
//	   - If this method completes successfully a new, fully
//	     populated instance of RuneArrayDto will be returned to the
//	     calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(runeArrayDtoElectron).
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - An array of runes used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex()
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            *RuneArrayDto
//	   - If this method completes successfully a pointer to a new,
//	     fully populated instance of RuneArrayDto will be returned
//	     to the calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			"NewRunesPtr()",
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

	err = new(runeArrayDtoElectron).
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this string is empty or has a zero length, an error
//	     will be returned.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto. If 'charSearchType' is invalid, an error
//	     will be returned.
//
//	     CharSearchType.LinearTargetStartingIndex() is the default
//	     search type.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex()
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            RuneArrayDto
//	   - If this method completes successfully a new, fully
//	     populated instance of RuneArrayDto will be returned to the
//	     calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			"NewString()",
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

	err = new(runeArrayDtoElectron).
		setRuneArray(
			&newRuneArrayDto,
			charArray,
			ePrefix.XCpy(
				"newRuneArrayDto"))

	newRuneArrayDto.charSearchType = charSearchType

	return newRuneArrayDto, err
}

// NewStringAllParams - Creates and returns a new instance of
// RuneArrayDto using a string input parameter. This method
// requires input parameters for all internal member variables
// contained in the RuneArrayDto type.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this string is empty or has a zero length, an error
//	     will be returned.
//
//
//	description1               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the text characters
//	     configured for the returned instance of RuneArrayDto. This
//	     parameter configures the first of two description strings.
//
//
//	description2               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the text characters
//	     configured for the returned instance of RuneArrayDto. This
//	     parameter configures the second of two description
//	     strings.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto. If 'charSearchType' is invalid, an error
//	     will be returned.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex()
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            RuneArrayDto
//	   - If this method completes successfully a new, fully
//	     populated instance of RuneArrayDto will be returned to the
//	     calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(runeArrayDtoElectron).
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to populate a new instance of
//	     RuneArrayDto which is returned to the calling function.
//
//	     If this string is empty or has a zero length, an error
//	     will be returned.
//
//
//	charSearchType             CharacterSearchType
//	   - An enumeration value used to specify the type target
//	     string search algorithm applied by the returned instance
//	     of RuneArrayDto. If 'charSearchType' is invalid, an error
//	     will be returned.
//
//
//	     The Character Search Type must be set to one of the
//	     following enumeration values:
//
//	      CharSearchType.LinearTargetStartingIndex()
//	      CharSearchType.SingleTargetChar()
//	      CharSearchType.LinearEndOfString()
//
//	  Character Search Type Options
//
//	  CharSearchType.LinearTargetStartingIndex()
//	  - Designates the search type as a Linear Target Starting
//	    Index Search Type. This means that each character in the
//	    Target Search String will be compared to each corresponding
//	    character in the Test String beginning at a specified
//	    starting index in the Target Search String.
//
//	    The search will proceed for from left to right in Test
//	    Character Sequence.
//
//	    If the Test Characters are NOT found in the Target Search
//	    String beginning at the designated Target String Starting
//	    Index, the search outcome will be unsuccessful and NO match
//	    will be declared.
//
//	    A 'Match', or successful search outcome, is defined as the
//	    case where each character in the Target String matches each
//	    corresponding character in the Test String beginning at the
//	    designated Target String Starting Index.
//
//
//	      Example
//	                                1         2         3
//	               Index  0123456789012345678901234567890
//	      Target String: "Hey, Xray-4 is the call sign."
//	      Target String Starting Index: 5
//	        Test String: "Xray"
//
//	    In this example of a Linear Target Starting Index Search, a
//	    match between the Target String and Test String will be
//	    declared, if and only if, the search begins at Target
//	    String index number 5. If the search begins at an any index
//	    other than 5, no match will be declared and the search will
//	    be classified as unsuccessful.
//
//	    NOTE: Linear Target Starting Index is the default search
//	          type.
//
//
//	 CharSearchType.SingleTargetChar()
//	  - Designates the search type as a Single Target Character
//	    Search Type. This means that a single character in the
//	    Target Search String will be compared to all characters in
//	    the Test String.
//
//	    If a single Target String character equals any character in
//	    the Test String, a 'Match' or successful search outcome
//	    will be declared.
//
//	    The search will proceed from left to right in the Target
//	    String. Each Target String Character will be compared to
//	    all characters in the Test String looking for the first
//	    matching Test String Character. The search will terminate
//	    when a matching character is first identified or when the
//	    end of the Target String is encountered.
//
//
//	      Example
//	                                 1         2         3
//	                Index  0123456789012345678901234567890
//	       Target String: "Hey, Xray-4 is the call sign."
//	       Target String Starting Index: 0
//	         Test String: "ZFXyURJK"
//
//	    In this example of a Single Target Character Search, the
//	    search will terminate at Target String index numbers 5
//	    because it is the first Target String index to match one
//	    of the Test String Characters ('X').
//
//
//	 CharSearchType.LinearEndOfString()
//	  - Designates the search type as a Linear End Of String
//	    Search. With this type of search operation, the entire
//	    Target Search String will be searched from left to right
//	    for the first occurrence of the Test String.
//
//	    The search will begin the Target String Starting Index and
//	    proceed left to right until (1) an instance of the entire
//	    Test String is located or (2) the end of the Target Search
//	    String is encountered.
//
//	    This is a linear search, so a 'Match' requires that each
//	    character in Target Search String must correspond to a
//	    matching character in the Test String.
//
//	         Example
//	                                    1         2         3
//	                   Index  0123456789012345678901234567890
//	          Target String: "Hey, Xray-4 is the call sign."
//	          Target String Starting Index: 0
//	            Test String: "Xray-4"
//
//	    In this example of a Linear End of String Search, the
//	    search operation will begin comparing corresponding
//	    characters in the Target Search String and the Test String
//	    beginning at index zero. The comparison will fail at index
//	    zero, but the search algorithm will continue attempting to
//	    find the Test String at indexes 1,2, 3 & 4. The Test String
//	    will be found beginning at index number 5 and the search
//	    algorithm will terminate at that point with a successful
//	    outcome or 'Match' result.
//
//
//	 For more information see the source code comments for type,
//	 CharacterSearchType.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newRuneArrayDto            *RuneArrayDto
//	   - If this method completes successfully a pointer to a new,
//	     fully populated instance of RuneArrayDto will be returned
//	     to the calling function.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			"NewStringPtr()",
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

	err = new(runeArrayDtoElectron).
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
// # TERMINOLOGY
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
//	Target String        - A string character or characters which
//	                       will be searched for the occurrence of
//	                       another predefined character or
//	                       characters referred to as a Test
//	                       String.
//
//
//	Test String          - A string character or characters which
//	                       will be used to search for matching
//	                       characters in a Target Search String.
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
//	CharSearchType.LinearTargetStartingIndex()
//	CharSearchType.SingleTargetChar()
//	CharSearchType.LinearEndOfString()
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
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms           CharSearchTargetInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Target
//	     Search String which will be used in completing the Linear
//	     Target Starting Index Search operation.
//
//	     The CharSearchTargetInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTargetInputParametersDto struct {
//
//	      TargetInputParametersName           string
//	       - The Name, Label or descriptive Tag associated with this
//	         instance of CharSearchTargetInputParametersDto. If empty,
//	         this string will be defaulted to "TargetInputParameters"
//
//	      TargetString                        *RuneArrayDto
//	       - A pointer to the RuneArrayDto containing the Target
//	         Search String text characters used in the search
//	         algorithm. Target Characters are compared against
//	         Test Characters to determine if a 'Match' condition
//	         exists.
//
//	      TargetStringName                    string
//	       - The label or name of the 'TargetString' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringLength                  int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringLengthName              string
//	       - The label or name of the 'TargetStringLength' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringStartingSearchIndex     int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringStartingSearchIndexName string
//	       - The label or name of the
//	         TargetStringStartingSearchIndex parameter.
//	         Used in error and informational messages.
//
//	      TargetStringSearchLength            int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringSearchLengthName        string
//	       - The label or name of the TargetStringSearchLength
//	         parameter. Used in error and informational
//	         messages.
//
//	      TargetStringAdjustedSearchLength    int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringDescription1            string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2            string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      FoundFirstNumericDigitInNumStr      bool
//	       - When set to 'true' this signals that the first
//	         numeric digit has been identified in the text
//	         characters specified by 'TargetString'
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      TextCharSearchType                  CharacterSearchType
//	       - Optional. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation. When set to a valid value,
//	         this specification will override the search
//	         specification contained in the Test Input
//	         Parameters Data Transfer Object.
//
//	        Valid CharSearch Type values are listed as follows:
//	        CharSearchType.None() - Invalid Value
//	        CharSearchType.LinearTargetStartingIndex() - Default
//	        CharSearchType.SingleTargetChar()
//	        CharSearchType.LinearEndOfString()
//	     }
//
//
//	errorPrefix                     interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	CharSearchResultsDto
//	   - If this method completes successfully, a fully populated
//	     instance of CharSearchResultsDto will be returned to the
//	     calling function.
//
//	     This structure contains information and data elements
//	     detailing the results of the text character search
//	     operation performed by this method.
//
//	     The CharSearchTestInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchResultsDto struct {
//
//	      SearchResultsName string
//	       - Optional. The Name, Label or descriptive Tag associated with
//	         the current instance of CharSearchResultsDto.
//
//	      SearchResultsFunctionChain string
//	       - Optional. A listing of the functions which led to the
//	         creation of this instance of CharSearchResultsDto.
//
//	      FoundSearchTarget bool
//	       - Signals a successful search outcome. If set to 'true' the
//	         Test String character or characters were found in the Target
//	         Search String.
//
//	      FoundSearchTargetOnPreviousSearch bool
//	       - Signals that the Search Target was located in a previous
//	         search operation.
//
//	      FoundFirstNumericDigitInNumStr bool
//	       - When set to 'true' this signals that the search operation
//	         has identified the first numeric digit in a string of text
//	         characters.
//
//	      TargetInputParametersName string
//	       - The Name, Label or descriptive Tag associated with an
//	         instance of CharSearchTargetInputParametersDto.
//
//	      TargetStringLength int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringSearchLength int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringAdjustedSearchLength int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringStartingSearchIndex int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringFirstFoundIndex int
//	       - The index of the first character position in the
//	         Target Search String occupied by the first
//	         character in the Test String.
//
//	      TargetStringLastFoundIndex int
//	       - The last character position in the Target Search
//	         String occupied by the last character in the Test
//	         String.
//
//	      TargetStringLastSearchIndex int
//	       - The index in Target Search String occupied by the
//	         last Target character searched. If the Search
//	         Target was found, this value is equal to the
//	         'TargetStringLastFoundIndex'. If the Search Target
//	         was NOT found this value is equal to the
//	         'TargetStringStartingSearchIndex'. This value is
//	         useful in computing the next index to be searched
//	         in the Target String.
//
//	      TargetStringNextSearchIndex int
//	       - The starting point for the next search operation.
//	         If the entire Target String was included in the
//	         last search, this value is set to -1.
//
//	          Example-1:
//	          String = "Hello"
//	          String Length = 5
//	          Last Search Index = 4
//	          TargetStringNextSearchIndex = -1
//
//	          Example-2:
//	          String = "Hello"
//	          String Length = 5
//	          Last Search Index = 2
//	          TargetStringNextSearchIndex = 3
//
//	      TargetStringDescription1 string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2 string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TestInputParametersName string
//	       - The Name, Label or descriptive Tag associated with an
//	         instance of CharSearchTestInputParametersDto.
//
//	      TestStringLength int
//	       - Actual number of text characters in the entire Test
//	         String ('TestString').
//
//	      TestStringStartingIndex int
//	       - The starting index in the Test String where the
//	         search operation will begin.
//
//	      TestStringFirstFoundIndex int
//	       - The index number in Test String of the first test
//	         character to be located in the Target Search String.
//
//	      TestStringLastFoundIndex int
//	       - The index number in the Test String occupied by the
//	         last Test Character to be located in the Target
//	         String.
//
//	      TestStringDescription1 string
//	       - First of two optional description strings
//	         describing the Test String in the context of the
//	         current search operation.
//
//	      TestStringDescription2 string
//	       - Second of two optional description strings
//	         describing the Test String in the context of the
//	         current search operation.
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      ReplacementString *RuneArrayDto
//	       - A pointer to the Rune Array Data Transfer Object
//	         containing the Replacement Characters to be
//	         substituted for existing characters in a Target
//	         String.
//
//	      RemainderString *RuneArrayDto
//	       - A pointer to the Rune Array Data Transfer Object
//	         containing the remaining characters in a Target
//	         String which were NOT included in the search
//	         operation and which remain to be searched in future
//	         search operations. This string is also used in 'cut'
//	         operations where Target String is divided based on
//	         string delimiters.
//
//	      NumValueType NumericValueType
//	       - Optional. This enumeration value specifies the type
//	         of numeric value for this Test Parameter as either
//	         an integer or floating point value.
//
//	         Possible values are listed as follows:
//	          NumValType.None()
//	          NumValType.FloatingPoint()
//	          NumValType.Integer()
//
//	      NumStrFormatType NumStrFormatTypeCode
//	       - Optional. This enumeration value specifies the
//	         Output Format Type for a number.
//
//	         Possible values are listed as follows:
//	          NumStrFmtType.None()
//	          NumStrFmtType.AbsoluteValue()
//	          NumStrFmtType.Binary()
//	          NumStrFmtType.CountryCulture()
//	          NumStrFmtType.Currency()
//	          NumStrFmtType.Binary()
//	          NumStrFmtType.Hexadecimal()
//	          NumStrFmtType.Octal()
//	          NumStrFmtType.ScientificNotation()
//
//	      NumSymbolLocation NumericSymbolLocation
//	       - Optional. This enumeration value specifies the
//	         relative location of a numeric symbol.
//
//	         Possible values are listed as follows:
//	          NumSymLocation.None()
//	          NumSymLocation.Before()
//	          NumSymLocation.Interior()
//	          NumSymLocation.After()
//
//	      NumSymbolClass NumericSymbolClass
//	       - Optional. This enumeration value specifies the
//	         Number Symbol Classification for a text character.
//
//	         Possible values are listed as follows:
//	          NumSymClass.None()
//	          NumSymClass.NumberSign()
//	          NumSymClass.CurrencySign()
//	          NumSymClass.IntegerSeparator()
//	          NumSymClass.DecimalSeparator()
//
//	      NumSignValue NumericSignValueType
//	       - Optional. This enumeration value specifies the
//	         number sign value.
//
//	         Possible values are listed as follows:
//	          NumSignVal.None()
//	          NumSignVal.Negative()
//	          NumSignVal.Zero()
//	          NumSignVal.Positive()
//
//	      PrimaryNumSignPosition NumSignSymbolPosition
//	       - Optional This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This is the
//	         Primary Type Code for Number Signs. Cases involving
//	         both 'Leading' and 'Trailing' symbols also make use
//	         of the 'SecondaryNumSignPosition'.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	      SecondaryNumSignPosition NumSignSymbolPosition
//	       - Optional. This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This value is used
//	         in searches involving number signs which occur both
//	         before and after the numeric value.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	      TextCharSearchType CharacterSearchType
//	       - Required. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation.
//
//	         Possible values are listed as follows:
//	          CharSearchType.None() - Invalid value
//	          CharSearchType.LinearTargetStartingIndex() - Default
//	          CharSearchType.SingleTargetChar()
//	          CharSearchType.LinearEndOfString()
//	     }
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
		// targetInputParms.TextCharSearchType is invalid

		err = testInputParms.ValidateCharSearchType(
			ePrefix)

		if err != nil {
			// testInputParms.TextCharSearchType is invalid
			// Use charsArrayDto.charSearchType

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
	}

	testInputParms.TestString = charsArrayDto

	if len(testInputParms.TestStringName) == 0 {

		testInputParms.TestStringName = "RuneArrayDto.CharsArray"

	}

	if len(testInputParms.TestStringLengthName) == 0 {

		testInputParms.TestStringLengthName =
			"RuneArrayDto Length"

	}

	testInputParms.TestStringStartingIndex = 0

	err = testInputParms.ValidateTestParameters(
		ePrefix)

	if err != nil {

		return errorResults, err

	}

	runeArrayNanobot := runeArrayDtoNanobot{}

	var searchResults CharSearchRuneArrayResultsDto

	searchResults,
		err = runeArrayNanobot.characterSearchExecutor(
		targetInputParms,
		testInputParms,
		ePrefix)

	if err != nil {

		return searchResults, err

	}

	if searchResults.FoundSearchTarget {

		err = searchResults.FoundRuneArrayChars.
			CopyIn(
				charsArrayDto,
				ePrefix.XCpy(
					"searchResults.FoundRuneArrayChars"+
						"<-charsArrayDto"))

		if err != nil {

			return searchResults, err

		}

	}

	return searchResults, err
}

// SetCharacterSearchType - Sets the internal member variable used
// to track the Text Character Search Type.
//
// ----------------------------------------------------------------
//
// # TERMINOLOGY
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
//	Target String        - A string character or characters which
//	                       will be searched for the occurrence of
//	                       another predefined character or
//	                       characters referred to as a Test
//	                       String.
//
//
//	Test String          - A string character or characters which
//	                       will be used to search for matching
//	                       characters in a Target Search String.
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
//	CharSearchType.LinearTargetStartingIndex()
//	CharSearchType.SingleTargetChar()
//	CharSearchType.LinearEndOfString()
//
// ----------------------------------------------------------------
//
// Character Search Type Options
//
//	CharSearchType.LinearTargetStartingIndex()
//
//	- Designates the search type as a Linear Target Starting Index
//	  Search Type. This means that each character in the Target
//	  Search String will be compared to each corresponding
//	  character in the Test String beginning at a specified
//	  starting index in the Target Search String.
//
//	  The search will proceed for from left to right in Test
//	  Character Sequence.
//
//	  If the Test Characters are NOT found in the Target Search
//	  String beginning at the designated Target String Starting
//	  Index, the search outcome will be unsuccessful and NO match
//	  will be declared.
//
//	  A 'Match', or successful search outcome, is defined as the
//	  case where each character in the Target String matches each
//	  corresponding character in the Test String beginning at the
//	  designated Target String Starting Index.
//
//
//	    Example
//	                              1         2         3
//	             Index  0123456789012345678901234567890
//	    Target String: "Hey, Xray-4 is the call sign."
//	    Target String Starting Index: 5
//	      Test String: "Xray"
//
//	  In this example of a Linear Target Starting Index Search, a
//	  match between the Target String and Test String will be
//	  declared, if and only if, the search begins at Target String
//	  index number 5. If the search begins at an any index other
//	  than 5, no match will be declared and the search will be
//	  classified as unsuccessful.
//
//	  NOTE: Linear Target Starting Index is the default search
//	        type.
//
// CharSearchType.SingleTargetChar()
//
//   - Designates the search type as a Single Target Character
//     Search Type. This means that a single character in the Target
//     Search String will be compared to all characters in the Test
//     String.
//
//     If a single Target String character equals any character in
//     the Test String, a 'Match' or successful search outcome will
//     be declared.
//
//     The search will proceed from left to right in the Target
//     String. Each Target String Character will be compared to all
//     characters in the Test String looking for the first matching
//     Test String Character. The search will terminate when a
//     matching character is first identified or when the end of the
//     Target String is encountered.
//
//     Example
//     1         2         3
//     Index  0123456789012345678901234567890
//     Target String: "Hey, Xray-4 is the call sign."
//     Target String Starting Index: 0
//     Test String: "ZFXyURJK"
//
//     In this example of a Single Target Character Search, the
//     search will terminate at Target String index numbers 5
//     because it is the first Target String index to match one
//     of the Test String Characters ('X').
//
// CharSearchType.LinearEndOfString()
//
//   - Designates the search type as a Linear End Of String Search.
//     With this type of search operation, the entire Target Search
//     String will be searched from left to right for the
//     first occurrence of the Test String.
//
//     The search will begin the Target String Starting Index and
//     proceed left to right until (1) an instance of the entire
//     Test String is located or (2) the end of the Target Search
//     String is encountered.
//
//     This is a linear search, so a 'Match' requires that each
//     character in Target Search String must correspond to a
//     matching character in the Test String.
//
//     Example
//     1         2         3
//     Index  0123456789012345678901234567890
//     Target String: "Hey, Xray-4 is the call sign."
//     Target String Starting Index: 0
//     Test String: "Xray-4"
//
//     In this example of a Linear End of String Search, the search
//     operation will begin comparing corresponding characters in
//     the Target Search String and the Test String beginning at
//     index zero. The comparison will fail at index zero, but the
//     search algorithm will continue attempting to find the Test
//     String at indexes 1,2, 3 & 4. The Test String will be found
//     beginning at index number 5 and the search algorithm will
//     terminate at that point with a successful outcome or 'Match'
//     result.
//
// For more information see the source code comments for type,
// CharacterSearchType.
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
//
//	RuneArrayDto.Description1
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
//
//	RuneArrayDto.Description2
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

//	SetRuneArray
//
//	Deletes the internal rune array for the current
//	instance of RuneArrayDto and repopulates it with
//	a deep copy extracted from input parameter,
//	'charArray'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of RuneArrayDto, all
//	pre-existing data in the internal rune array
//	member variable, 'CharsArray' will be deleted
//	and overwritten with new data.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	charArray					[]rune
//		An array of runes used to populate the internal
//		member variable rune array for the current
//		instance of RuneArrayDto.
//
//	     If this array is empty or has a zero length, an
//	     error will be returned.
//
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
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

	return new(runeArrayDtoElectron).
		setRuneArray(
			charsArrayDto,
			charArray,
			ePrefix.XCpy(
				"charsArrayDto"))
}

//	SetRunesDefault
//
//	Receives a rune array and proceeds to copy the
//	text characters to the internal rune array
//	maintained by the current instance of
//	RuneArrayDto:
//		RuneArrayDto.CharsArray
//
//	The Character Search type for the current
//	instance of RuneArrayDto will be defaulted
//	to:
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The data fields for the current instance of
//	RuneArrayDto will be modified.
//
//	For the current instance of RuneArrayDto, all
//	pre-existing data in the internal rune array
//	member variable, 'RuneArrayDto.CharsArray',
//	will be deleted and overwritten with new data.
//
//	In addition, the Character Search type for the
//	current instance of RuneArrayDto will be
//	defaulted to:
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	runeChars					[]rune
//
//	An array of runes which will be copied to the
//	member variable, 'RuneArrayDto.CharsArray', in
//	the current instance of RuneArrayDto.
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (charsArrayDto *RuneArrayDto) SetRunesDefault(
	runeChars []rune) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	_ = new(runeArrayDtoElectron).
		setRuneArray(
			charsArrayDto,
			runeChars,
			nil)

	charsArrayDto.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return
}

//	SetStrDefault
//
//	Receives a string and proceeds to copy the text
//	characters to the internal rune array maintained
//	by the current instance of RuneArrayDto:
//		RuneArrayDto.CharsArray
//
//	The Character Search type for the current
//	instance of RuneArrayDto will be defaulted
//	to:
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The data fields for the current instance of
//	RuneArrayDto will be modified.
//
//	For the current instance of RuneArrayDto, all
//	pre-existing data in the internal rune array
//	member variable, 'RuneArrayDto.CharsArray',
//	will be deleted and overwritten with new data.
//
//	In addition, the Character Search type for the
//	current instance of RuneArrayDto will be
//	defaulted to:
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	stringChars					string
//
//	The text characters contained in this string
//	will be copied to the member variable,
//	'RuneArrayDto.CharsArray', in the current
//	instance of RuneArrayDto.
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (charsArrayDto *RuneArrayDto) SetStrDefault(
	stringChars string) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	_ = new(runeArrayDtoElectron).
		setRuneArray(
			charsArrayDto,
			[]rune(stringChars),
			nil)

	charsArrayDto.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return
}

// SetString - Receives a string as an input parameter
// and proceeds to populate the internal rune array with new
// character data for the current instance of RuneArrayDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// For the current instance of RuneArrayDto, all pre-existing data
// in the internal rune array member variable, 'CharsArray' will be
// deleted and overwritten with new data.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charString                 string
//	   - An string of characters which will be used to populate
//	     the internal member variable rune array for the current
//	     instance of RuneArrayDto.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (charsArrayDto *RuneArrayDto) SetString(
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
			"SetString()",
		"")

	if err != nil {
		return err
	}

	charArray := []rune(charString)

	return new(runeArrayDtoElectron).
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string
//	   - This method returns a string of characters extracted from
//	     the internal member variable rune array
//	     'RuneArrayDto.CharsArray'.
func (charsArrayDto RuneArrayDto) String() string {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return string(charsArrayDto.CharsArray)
}
