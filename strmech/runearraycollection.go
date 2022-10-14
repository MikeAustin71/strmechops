package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// RuneArrayCollection - A collection of Rune Array Dto objects.
//
// Essentially, this is a collection, or an array, of rune arrays.
type RuneArrayCollection struct {
	runeArrayDtoCol []RuneArrayDto

	lock *sync.Mutex
}

// AddCollection - Receives an array of RuneArrayDto objects and
// adds them to the collection maintained by the current instance
// of RuneArrayCollection.
//
// Input parameter 'runeArrayDtoCol' consists of an array of
// RuneArrayDto objects. Deep copies of these objects will be used
// to populate the new Rune Array Collection.
//
// If any member of the RuneArrayDto array, 'runeArrayDtoCol', is
// judged to be invalid, an error will be returned.
//
// Valid RuneArrayDto array members must satisfy two criteria in
// order to be classified as 'valid'.
//
//	(1) The RuneArrayDto internal Character Array must have a
//	    length greater than zero.
//
//	(2) The Character Search Type associated with each
//	    RuneArrayDto must be valid.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDtoCol            []RuneArrayDto
//	   - An array of RuneArrayDto objects which be used to replace
//	     the existing Rune Array Dto Collection maintained by the
//	     current instance of RuneArrayCollection.
//
//	     Deep copies of each element in parameter 'runeArrayDtoCol'
//	     will be used to populate the new Rune Array Dto Collection
//	     contained in the current instance of RuneArrayCollection.
//
//	     If this parameter is submitted as a 'nil' or zero length
//	     array, an error will be returned.
//
//	     If any member element of this array is classified as
//	     invalid, an error will be returned.
//
//	     Valid RuneArrayDto array members must satisfy two criteria
//	     in order to be classified as 'valid'.
//
//	     (1) The RuneArrayDto internal Character Array must have a
//	         length greater than zero.
//
//	     (2) The Character Search Type associated with each
//	         RuneArrayDto must be valid.
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
// ------------------------------------------------------------------------
//
// Return Values
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
func (runeArrayCol *RuneArrayCollection) AddCollection(
	runeArrayDtoCol []RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddCollection()",
		"")

	if err != nil {
		return err
	}

	lenIncomingRuneArray := len(runeArrayDtoCol)

	if lenIncomingRuneArray == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'runeArrayDtoCol' is EMPTY!\n",
			ePrefix.String())

		return err
	}

	for i := 0; i < lenIncomingRuneArray; i++ {

		err = runeArrayDtoCol[i].IsValidCharacterArrayError(
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayDtoCol[%v] Character Array Error!",
					i)))

		if err != nil {
			return err
		}

		err = runeArrayDtoCol[i].IsValidCharacterSearchTypeError(
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayDtoCol[%v] Character Serach Type Error!",
					i)))

		if err != nil {
			return err
		}
	}

	runeArrayCol.runeArrayDtoCol = append(
		runeArrayCol.runeArrayDtoCol,
		runeArrayDtoCol...)

	return err
}

// AddLatinAlphabetEnglishDto - Adds a RuneArrayDto to the Rune
// Array Collection. This RuneArrayDto is populated with the Latin
// Alphabet (English Version). The total number of characters
// is 52 comprised of 26 lower case letters and 26 upper case
// letters.
//
// An array of alphabetic characters in useful in search
// operations looking for alphabetic characters classified as
// delimiters.
//
// ----------------------------------------------------------------
//
// # Default Settings
//
// The character search type for the RuneArrayDto instance is
// automatically set to:
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
// -----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (runeArrayCol *RuneArrayCollection) AddLatinAlphabetEnglishDto() {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	runeArrayCol.runeArrayDtoCol =
		append(
			runeArrayCol.runeArrayDtoCol,
			RuneArrayDto{}.NewLatinAlphabet())

}

// AddNumericDigitsDto - Adds a RuneArrayDto to the Rune Array
// Collection. This RuneArrayDto is populated with numeric digits
// zero (0) through nine (9).
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
// -----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (runeArrayCol *RuneArrayCollection) AddNumericDigitsDto() {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	runeArrayCol.runeArrayDtoCol =
		append(
			runeArrayCol.runeArrayDtoCol,
			RuneArrayDto{}.NewNumericCharacters())
}

// AddRuneArrayDto - Receives an instance of RuneArrayDto and
// appends that instance to the RuneArrayDto Collection.
//
// This differs from method:
//
//	RuneArrayCollection.AddRuneArrayDtoDeepCopy()
//
// This method appends the passed RuneArrayDto instance to the
// collection. The Deep Copy method appends a copy of the
// RuneArrayDto to the collection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// RuneArrayDto instances added to this collection MUST HAVE a
// valid Character Search Type. Reference RuneArrayDto member
// variable 'RuneArrayDto.charSearchType'.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               RuneArrayDto
//	   - An instance of RuneArrayDto. This instance will be
//	     appended to the RuneArrayDto collection maintained by this
//	     instance of RuneArrayCollection.
//
//	     If 'runeArrayDto' has an invalid character search type
//	     (runeArrayDto.charSearchType), an error will be returned.
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
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayDto(
	runeArrayDto RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayDto()",
		"")

	if err != nil {
		return err
	}

	lenRuneDtoChars := len(runeArrayDto.CharsArray)

	if lenRuneDtoChars == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"runeArrayDto.CharsArray has a length of zero",
			ePrefix.String())

		return err

	}

	if !runeArrayDto.charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"Member vaiable Character Search Type is invalid.\n"+
			"runeArrayDto.charSearchType String Value = '%v'\n"+
			"runeArrayDto.charSearchType Integer Value = '%v'\n",
			ePrefix.String(),
			runeArrayDto.charSearchType.String(),
			runeArrayDto.charSearchType.XValueInt())

		return err

	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, runeArrayDto)

	return err
}

// AddRuneArrayDtoDeepCopy - Receives an instance of RuneArrayDto and
// appends a deep copy of that instance to the RuneArrayDto
// Collection.
//
// This differs from method:
//
//	RuneArrayCollection.AddRuneArrayDto()
//
// This method appends a deep copy of the passed RuneArrayDto
// instance to the collection. The 'AddRuneArrayDto()' method
// appends the actual RuneArrayDto instance to the collection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// RuneArrayDto instances added to this collection MUST HAVE a
// valid Character Search Type. Reference RuneArrayDto member
// variable 'RuneArrayDto.charSearchType'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               RuneArrayDto
//	   - An instance of RuneArrayDto. A deep copy of this instance
//	     will be appended to the RuneArrayDto collection maintained
//	     by this instance of RuneArrayCollection.
//
//	     If 'runeArrayDto' has an invalid character search type
//	     (runeArrayDto.charSearchType), an error will be returned.
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
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayDtoDeepCopy(
	runeArrayDto RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayDto()",
		"")

	if err != nil {
		return err
	}

	lenRuneDtoChars := len(runeArrayDto.CharsArray)

	if lenRuneDtoChars == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"runeArrayDto.CharsArray has a length of zero",
			ePrefix.String())

		return err

	}

	if !runeArrayDto.charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"Member vaiable Character Search Type is invalid.\n"+
			"runeArrayDto.charSearchType String Value = '%v'\n"+
			"runeArrayDto.charSearchType Integer Value = '%v'\n",
			ePrefix.String(),
			runeArrayDto.charSearchType.String(),
			runeArrayDto.charSearchType.XValueInt())

		return err

	}

	var deepCopyRuneArrayDto RuneArrayDto

	deepCopyRuneArrayDto,
		err = runeArrayDto.CopyOut(
		ePrefix.XCpy(
			"deepCopyRuneArrayDto<-runeArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, deepCopyRuneArrayDto)

	return err
}

// AddRuneArrayRunes - Adds a new RuneArrayDto instance to the
// Rune Array Collection. The new RuneArrayDto is created from
// a rune array input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - A rune array used to create the new instance of
//	     RuneArrayDto which is then added to the Rune Array
//	     collection.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayRunes(
	charArray []rune,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayRunes()",
		"")

	if err != nil {
		return err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewRunes(
		charArray,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return err
}

// AddRuneArrayRunesDesc - Adds a new RuneArrayDto instance to the
// Rune Array Collection. The new RuneArrayDto is created from
// a rune array input parameter. Users also have the option of
// attaching text descriptions to the new RuneArrayDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - A rune array used to create the new instance of
//	     RuneArrayDto which is then added to the Rune Array
//	     collection.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	description1               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the RuneArrayDto
//	     instance added to the Rune Array Collection. This
//	     parameter configures the first of two description
//	     strings.
//
//	     This parameter is optional and not required. Zero length
//	     strings will NOT return an error.
//
//
//	description2               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the RuneArrayDto
//	     instance added to the Rune Array Collection. This
//	     parameter configures the second of two description
//	     strings.
//
//	     This parameter is optional and not required. Zero length
//	     strings will NOT return an error.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayRunesDesc(
	charArray []rune,
	description1 string,
	description2 string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayRunesDesc()",
		"")

	if err != nil {
		return err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewRunesAllParams(
		charArray,
		description1,
		description2,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return err
}

// AddRuneArrayString - Adds a new RuneArrayDto instance to the
// Rune Array Collection. The new RuneArrayDto is created from
// an input parameter of type string.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to create the new instance of
//	     RuneArrayDto which is then added to the Rune Array
//	     collection.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayString(
	stringChars string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayString()",
		"")

	if err != nil {
		return err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewString(
		stringChars,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return err
}

// AddRuneArrayStringDesc - Adds a new RuneArrayDto instance to the
// Rune Array Collection. The new RuneArrayDto is created from an
// input parameter of type string. Users also have the option of
// attaching text descriptions to the new RuneArrayDto instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to create the new instance of
//	     RuneArrayDto which is then added to the Rune Array
//	     collection.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	description1               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the RuneArrayDto
//	     instance added to the Rune Array Collection. This
//	     parameter configures the first of two description
//	     strings.
//
//	     If this descriptive text is not required, set this
//	     parameter to an empty string.
//
//
//	description2               string
//
//	   - Users have the option of configuring a text string to
//	     describe the function or purpose of the RuneArrayDto
//	     instance added to the Rune Array Collection. This
//	     parameter configures the second of two description
//	     strings.
//
//	     If this descriptive text is not required, set this
//	     parameter to an empty string.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) AddRuneArrayStringDesc(
	stringChars string,
	description1 string,
	description2 string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArrayStringDesc()",
		"")

	if err != nil {
		return err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewStringAllParams(
		stringChars,
		description1,
		description2,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return err
}

//	AddRunesDefault
//
//	Receives a rune array and proceeds to create a new
//	instance of RuneArrayDto which is then added to the
//	collection maintained by the current instance of
//	RuneArrayCollection.
//
//	When creating the RuneArrayDto object to be added
//	to the collection, the Character Search Type is
//	automatically defaulted to:
//
//		CharSearchType.LinearTargetStartingIndex()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	charArray					[]rune
//
//		This rune array is used to construct the new
//		RuneArrayDto object which will be added to the
//		collection maintained by the current instance
//		of RuneArrayCollection.
//
//		If this rune array is 'nil' or has a zero
//		length, an error will be returned.
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
// -----------------------------------------------------------------
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
func (runeArrayCol *RuneArrayCollection) AddRunesDefault(
	charArray []rune,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRunesDefault()",
		"")

	if err != nil {
		return err
	}

	if len(charArray) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is zero length\n"+
			"rune array!\n",
			ePrefix.String())

		return err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto = RuneArrayDto{}.NewRunesDefault(
		charArray)

	runeArrayCol.runeArrayDtoCol =
		append(runeArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return err
}

// CopyIn - Copies the data fields from an incoming instance of
// RuneArrayCollection ('incomingRuneArrayCol') to the data fields
// of the current RuneArrayCollection instance ('runeArrayCol').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current RuneArrayCollection instance
// ('runeArrayCol') will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingRuneArrayCol       *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. This
//	     method will NOT change the values of internal member
//	     variables contained in this instance.
//
//	     All data values in this RuneArrayCollection instance
//	     will be copied to current RuneArrayCollection
//	     instance ('runeArrayCol').
//
//	     If parameter 'incomingRuneArrayCol' is determined to be
//	     invalid, an error will be returned.
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
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) CopyIn(
	incomingRuneArrayCol *RuneArrayCollection,
	errorPrefix interface{}) error {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(runeArrayCollectionNanobot).
		copyIn(
			runeArrayCol,
			incomingRuneArrayCol,
			ePrefix.XCpy(
				"runeArrayCol<-incomingRuneArrayCol"))
}

// CopyOut - Returns a deep copy of the current RuneArrayCollection
// instance.
//
// If the current RuneArrayCollection instance contains invalid
// member variables, this method will return an error.
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
//	RuneArrayCollection
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current RuneArrayCollection instance.
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
func (runeArrayCol *RuneArrayCollection) CopyOut(
	errorPrefix interface{}) (
	RuneArrayCollection,
	error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"CopyOut()",
		"")

	if err != nil {
		return RuneArrayCollection{}, err
	}

	return new(runeArrayCollectionNanobot).
		copyOut(
			runeArrayCol,
			ePrefix.XCpy(
				"<-runeArrayCol"))
}

// DeleteCollectionElement - Deletes a member element of the Rune
// Arrays Collection maintained by the current instance of
// RuneArrayCollection. The element to be deleted is specified by
// input parameter 'zeroBasedIndex' and represents the array index
// for the element to be deleted.
//
// If the current Rune Array Collection is empty and has zero
// elements, an error will be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised, this method WILL DELETE DATA!
//
// The Rune Arrays Collection member element specified by input
// parameter 'zeroBasedIndex' WILL BE DELETED.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Rune
//	     Arrays Collection which will be deleted.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol *RuneArrayCollection) DeleteCollectionElement(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PopIndexElement()",
		"")

	if err != nil {
		return err
	}

	return runeArrayCollectionQuark{}.ptr().
		deleteCollectionElement(
			runeArrayCol,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayCol.runeArrayDtoCol[%v]",
					zeroBasedIndex)))
}

// Empty - Empties the internal RuneArrayDto collection and sets
// its value to 'nil'. This method will leave the current
// instance of RuneArrayCollection in an invalid state and
// unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will members of the internal RuneArrayDto
// collection. After completion, the internal RuneArrayDto
// collection will have a value of 'nil'.
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
//	NONE
func (runeArrayCol *RuneArrayCollection) Empty() {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	runeArrayCollectionAtom{}.ptr().
		empty(runeArrayCol)

	runeArrayCol.lock.Unlock()

	runeArrayCol.lock = nil

	return
}

// Equal - Receives a pointer to another instance of
// RuneArrayCollection and proceeds to compare its internal
// member variables to those of the current
// RuneArrayCollection instance in order to determine if they
// are equivalent.
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
//	incomingRuneArrayCol       *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. The
//	     internal member variable data values in this instance will
//	     be compared to those in the current instance of
//	     RuneArrayCollection. The results of this comparison
//	     will be returned to the calling functions as a boolean
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the internal member variable data values contained in
//	     input parameter 'incomingRuneArrayCol' are equivalent
//	     in all respects to those contained in the current instance
//	     of RuneArrayCollection, this return value will be set
//	     to 'true'.
//
//	     Otherwise, this method will return 'false'.
func (runeArrayCol *RuneArrayCollection) Equal(
	incomingRuneArrayCol *RuneArrayCollection) bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	return runeArrayCollectionAtom{}.ptr().
		equal(
			runeArrayCol,
			incomingRuneArrayCol)
}

// GetCollection - Returns a deep copy of the Rune Array Collection
// maintained by the current instance of RuneArrayCollection.
//
// If the Rune Array Collection is empty or has zero elements, an
// error will be returned.
//
// This operation is nondestructive meaning that the Rune Array
// Collection maintained by the current instance of
// RuneArrayCollection is unchanged by this method. No data will be
// deleted.
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
//	deepCopyRuneArrayCol       []RuneArray
//	   - This method will return an array of RuneArrayDto objects
//	     representing a deep copy of the Rune Array Collection
//	     maintained by the current instance of RuneArrayCollection.
//
//	     If the current Rune Array Collection is empty, this
//	     return parameter will be set to 'nil' and an error will
//	     be returned.
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
func (runeArrayCol *RuneArrayCollection) GetCollection(
	errorPrefix interface{}) (
	deepCopyRuneArrayCol []RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	deepCopyRuneArrayCol = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"GetCollection()",
		"")

	if err != nil {
		return deepCopyRuneArrayCol, err
	}

	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneArrayCol == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Rune Array Dto Collection, 'runeArrayCol.runeArrayDtoCol' is EMPTY!\n",
			ePrefix.String())

		return deepCopyRuneArrayCol, err
	}

	deepCopyRuneArrayCol =
		make([]RuneArrayDto, lenRuneArrayCol)

	for i := 0; i < lenRuneArrayCol; i++ {

		deepCopyRuneArrayCol[i],
			err = runeArrayCol.runeArrayDtoCol[i].CopyOut(
			ePrefix.XCpy(
				fmt.Sprintf("deepCopyRuneArrayCol"+
					"<-runeArrayCol.runeArrayDtoCol[%v]",
					i)))

		if err != nil {
			deepCopyRuneArrayCol = nil
			return deepCopyRuneArrayCol, err
		}

	}

	return deepCopyRuneArrayCol, err
}

// GetNumberOfRuneArrayDtos - Returns the number of elements in the
// RuneArrayDto collection. The returned integer value is therefore
// equal to the length of the internal array of RuneArrayDto
// objects.
func (runeArrayCol *RuneArrayCollection) GetNumberOfRuneArrayDtos() int {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	return len(runeArrayCol.runeArrayDtoCol)
}

// GetStringArrayDto - Returns an instance of StringArrayDto
// created from the current instance of RuneArrayCollection.
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
//	NONE
func (runeArrayCol *RuneArrayCollection) GetStringArrayDto() StringArrayDto {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	newStrArray := StringArrayDto{}

	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneArrayCol == 0 {
		return newStrArray
	}

	newStrArray.StrArray =
		make([]string, lenRuneArrayCol)

	for i := 0; i < lenRuneArrayCol; i++ {
		newStrArray.StrArray[i] =
			runeArrayCol.runeArrayDtoCol[i].GetCharacterString()
	}

	return newStrArray
}

// IsNOP - Stands for 'Is No Operation'.
//
// If the Rune Array Dto Collection for the current instance of
// RuneArrayCollection is empty (has a length of zero), it signals
// that this instance is an empty placeholder that performs no
// operations.
//
// If the current instance of RuneArrayCollection is classified as
// 'No Operation', this method returns 'true'.
//
// Otherwise, this method returns 'false' signaling that the
// current instance of RuneArrayCollection if fully populated,
// functional and ready to perform assigned tasks.
func (runeArrayCol *RuneArrayCollection) IsNOP() bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	if len(runeArrayCol.runeArrayDtoCol) == 0 {

		// This instance is a NOP
		return true
	}

	// This instance is NOT a NOP.
	// Open for business and ready for action.
	return false
}

// IsValidInstance - Returns a boolean value signaling whether the
// current instance of RuneArrayCollection is valid.
//
// There is only one criterion for classifying an instance of
// RuneArrayCollection as valid. It must contain a Rune Array Dto
// Collection where the number of elements is greater than zero.
//
// If the length of internal member variable
// 'RuneArrayCollection.runeArrayDtoCol' is greater than zero, this
// method will return 'true'.
//
// If the length of internal member variable
// RuneArrayCollection.runeArrayDtoCol is equal to zero, this
// method will return 'false'.
//
// This method is identical in function to method:
//
//	RuneArrayCollection.IsValidInstanceError
//
// The only difference is that this method returns a boolean value,
// while 'IsValidInstanceError()' returns an error.
func (runeArrayCol *RuneArrayCollection) IsValidInstance() bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	isValid,
		_ :=
		runeArrayCollectionElectron{}.ptr().
			testValidityRuneArrayCollection(
				runeArrayCol.runeArrayDtoCol,
				nil)

	return isValid
}

// IsValidInstanceError - Returns an error if the current instance
// of RuneArrayCollection is invalid.
//
// There is only one criterion for classifying an instance of
// RuneArrayCollection as valid. It must contain a Rune Array Dto
// Collection where the number of elements is greater than zero.
//
// If the length of internal member variable
// 'RuneArrayCollection.runeArrayDtoCol' is greater than zero, this
// method will return 'nil' signaling "No Error".
//
// If the length of internal member variable
// RuneArrayCollection.runeArrayDtoCol is equal to zero, this
// method will return an error containing an appropriate error
// message.
//
// This method is identical in function to method:
//
//	RuneArrayCollection.IsValidInstance
//
// The only difference is that this method returns an error, while
// 'IsValidInstance()' returns a boolean value.
func (runeArrayCol *RuneArrayCollection) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err =
		runeArrayCollectionElectron{}.ptr().
			testValidityRuneArrayCollection(
				runeArrayCol.runeArrayDtoCol,
				ePrefix.XCpy(
					"runeArrayCol"))

	return err
}

// New - Returns a new uninitialized instance of
// RuneArrayCollection. The internal RuneArrayDto collection will
// be empty.
//
// In this state, the returned instance of RuneArrayCollection is
// invalid an unusable. It will then be necessary to add
// RuneArrayDto objects to the collection using 'Add' or 'Set'
// methods.
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
//	RuneArrayCollection
//	   - This method returns an empty an uninitialized instance of
//	     RuneArrayCollection. The internal RuneArrayDto Collection
//	     is empty.
func (runeArrayCol RuneArrayCollection) New() RuneArrayCollection {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	newRuneArrayCol := RuneArrayCollection{}

	runeArrayCollectionAtom{}.ptr().
		empty(&newRuneArrayCol)

	return newRuneArrayCol
}

// NewColMemberString - Creates a new instance of
// RuneArrayCollection which includes a single RuneArrayDto object
// in the Rune Array Dto Collection.
//
// This single RuneArrayDto collection member is generated from the
// string and character search type input parameters. Users also
// have the option of creating tags or descriptive text for the new
// RuneArrayDto object.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	stringChars                string
//	   - A character string used to create the new instance of
//	     RuneArrayDto which is then added to the Rune Array Dto
//	     Collection.
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
//	     string search algorithm applied by the RuneArrayDto
//	     instance added to the Rune Array Collection.
//
//	     If 'charSearchType' is invalid, an error will be returned.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol RuneArrayCollection) NewColMemberString(
	stringChars string,
	description1 string,
	description2 string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	newRuneArrayCol RuneArrayCollection,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"NewColMemberString()",
		"")

	if err != nil {
		return newRuneArrayCol, err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewStringAllParams(
		stringChars,
		description1,
		description2,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return newRuneArrayCol, err
	}

	newRuneArrayCol.runeArrayDtoCol =
		append(newRuneArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return newRuneArrayCol, err
}

// NewColMemberRunes - Creates a new instance of
// RuneArrayCollection which includes a single RuneArrayDto object
// in the Rune Array Dto Collection.
//
// This single RuneArrayDto collection member is generated from the
// rune array and character search type passed as input parameters.
// Users also have the option of creating tags or descriptive text
// for the new RuneArrayDto object.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	charArray                  []rune
//	   - An array of runes used to populate a new instance of
//	     RuneArrayDto  which is then added to the Rune Array Dto
//	     Collection.
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
//	     string search algorithm applied by the RuneArrayDto
//	     instance added to the Rune Array Collection.
//
//	     If 'charSearchType' is invalid, an error will be returned.
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
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (runeArrayCol RuneArrayCollection) NewColMemberRunes(
	charArray []rune,
	description1 string,
	description2 string,
	charSearchType CharacterSearchType,
	errorPrefix interface{}) (
	newRuneArrayCol RuneArrayCollection,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"NewColMemberRunes()",
		"")

	if err != nil {
		return newRuneArrayCol, err
	}

	var newRuneArrayDto RuneArrayDto

	newRuneArrayDto,
		err = RuneArrayDto{}.NewRunesAllParams(
		charArray,
		description1,
		description2,
		charSearchType,
		ePrefix.XCpy(
			"newRuneArrayDto"))

	if err != nil {
		return newRuneArrayCol, err
	}

	newRuneArrayCol.runeArrayDtoCol =
		append(newRuneArrayCol.runeArrayDtoCol, newRuneArrayDto)

	return newRuneArrayCol, err
}

// PeekAtFirstElement - Returns a deep copy of the first element in
// the Rune Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// This 'peek' operation is nondestructive meaning that the Rune
// Array Collection maintained by the current instance of
// RuneArrayCollection is unchanged by this method. No data will be
// deleted.
//
// If the current Rune Array Collection is empty and has zero
// elements, an error will be returned.
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
//	firstRuneArrayDto          RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     first Rune Array Data Transfer Object 'RuneArrayDto' in
//	     the Rune Array Collection maintained by the current
//	     RuneArrayCollection instance will be returned to the
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
func (runeArrayCol *RuneArrayCollection) PeekAtFirstElement(
	errorPrefix interface{}) (
	firstRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PeekAtFirstElement()",
		"")

	if err != nil {
		return firstRuneArrayDto, err
	}

	firstRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			0,
			false,
			ePrefix.XCpy(
				"firstRuneArrayDto<-"+
					"runeArrayCol.runeArrayDtoCol[0]"))

	return firstRuneArrayDto, err
}

// PeekAtIndexElement - Returns a deep copy of the Rune Array
// Collection member element specified by input parameter,
// 'zeroBasedIndex'.
//
// This 'peek' operation is nondestructive meaning that the Rune
// Array Collection maintained by the current instance of
// RuneArrayCollection is unchanged by this method.
//
// If the Rune Array collection maintained by the current
// RuneArrayCollection instance is empty (contains zero elements),
// an error will be returned.
//
// Remember that indexes in the Rune Array collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the index of the last member element in the collection, an
// error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Rune
//	     Arrays Collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Rune Arrays Collection is empty, an
//	     error will be returned.
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
//	targetRuneArrayDto         RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     Rune Array Data Transfer Object 'RuneArrayDto' specified
//	     by index parameter 'zeroBasedIndex' in the Rune Array
//	     Collection maintained by the current RuneArrayCollection
//	     instance will be returned to the calling function.
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
func (runeArrayCol *RuneArrayCollection) PeekAtIndexElement(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	targetRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PeekAtIndexElement()",
		"")

	if err != nil {
		return targetRuneArrayDto, err
	}

	targetRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			zeroBasedIndex,
			false,
			ePrefix.XCpy(
				fmt.Sprintf("targetRuneArrayDto<-"+
					"runeArrayCol.runeArrayDtoCol[%v]",
					zeroBasedIndex)))

	return targetRuneArrayDto, err
}

// PeekAtLastElement - Returns a deep copy of the last element in
// the Rune Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// This 'peek' operation is nondestructive meaning that the Rune
// Array Collection maintained by the current instance of
// RuneArrayCollection is unchanged by this method. No data will be
// deleted.
//
// If the current Rune Array Collection is empty and has zero
// elements, an error will be returned.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	lastRuneArrayDto           RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     last Rune Array Data Transfer Object 'RuneArrayDto' in
//	     the Rune Array Collection maintained by the current
//	     RuneArrayCollection instance will be returned to the
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
func (runeArrayCol *RuneArrayCollection) PeekAtLastElement(
	errorPrefix interface{}) (
	lastRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PeekAtLastElement()",
		"")

	if err != nil {
		return lastRuneArrayDto, err
	}

	lastIdx := len(runeArrayCol.runeArrayDtoCol) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Rune Array Dto Collection is empty!\n",
			ePrefix.String())

		return lastRuneArrayDto, err
	}

	lastRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			lastIdx,
			false,
			ePrefix.XCpy(
				fmt.Sprintf("lastRuneArrayDto<-"+
					"runeArrayCol.runeArrayDtoCol[%v]",
					lastIdx)))

	return lastRuneArrayDto, err
}

// PopFirstElement - Returns a deep copy of the first element in
// the Rune Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// If the current Rune Array Collection is empty and has zero
// elements, an error will be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This 'pop' operation is destructive meaning that after returning
// a deep copy of the first element in the Rune Array Collection,
// this method will proceed to DELETE the first element in the Rune
// Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// Be advised, this method WILL DELETE DATA!
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	firstRuneArrayDto          RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     first Rune Array Data Transfer Object 'RuneArrayDto' in
//	     the Rune Array Collection maintained by the current
//	     RuneArrayCollection instance will be returned to the
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
func (runeArrayCol *RuneArrayCollection) PopFirstElement(
	errorPrefix interface{}) (
	firstRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PopFirstElement()",
		"")

	if err != nil {
		return firstRuneArrayDto, err
	}

	firstRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			0,
			true,
			ePrefix.XCpy(
				"firstRuneArrayDt<-"+
					"runeArrayCol.runeArrayDtoCol[0]"))

	return firstRuneArrayDto, err
}

// PopIndexElement - Returns a deep copy of the Rune Array
// Collection member element specified by input parameter,
// 'zeroBasedIndex'.
//
// If the Rune Array collection maintained by the current
// RuneArrayCollection instance is empty (contains zero elements),
// an error will be returned.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the index of the last member element in the collection, an
// error will be returned.
//
// Remember that indexes in the Rune Array collection are zero
// based. This means the first element in the collection is index
// zero.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This 'pop' operation is destructive meaning that after returning
// a deep copy of the first element in the Rune Array Collection,
// this method will proceed to DELETE the array element in the Rune
// Array Collection specified by input parameter 'zeroBasedIndex'.
//
// Be advised, this method WILL DELETE DATA!
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Rune
//	     Arrays Collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Rune Arrays Collection is empty, an
//	     error will be returned.
//
//	     Afterwards, this method will proceed to DELETE the array
//	     element in the Rune Arrays Collection specified by this
//	     input parameter.
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
//	targetRuneArrayDto         RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     Rune Array Data Transfer Object 'RuneArrayDto' specified
//	     by index parameter 'zeroBasedIndex' in the Rune Array
//	     Collection maintained by the current RuneArrayCollection
//	     instance will be returned to the calling function.
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
func (runeArrayCol *RuneArrayCollection) PopIndexElement(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	targetRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PopIndexElement()",
		"")

	if err != nil {
		return targetRuneArrayDto, err
	}

	targetRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			zeroBasedIndex,
			true,
			ePrefix.XCpy(
				fmt.Sprintf("targetRuneArrayDto<-"+
					"runeArrayCol.runeArrayDtoCol[%v]",
					zeroBasedIndex)))

	return targetRuneArrayDto, err
}

// PopLastElement - Returns a deep copy of the last element in
// the Rune Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// If the current Rune Array Collection is empty and has zero
// elements, an error will be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This 'pop' operation is destructive meaning that after returning
// a deep copy of the last element in the Rune Array Collection,
// this method will proceed to DELETE the last element in the Rune
// Array Collection maintained by the current instance of
// RuneArrayCollection.
//
// Be advised, this method WILL DELETE DATA!
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	lastRuneArrayDto           RuneArrayDto
//	   - If this method completes successfully, a deep copy of the
//	     last Rune Array Data Transfer Object 'RuneArrayDto' in
//	     the Rune Array Collection maintained by the current
//	     RuneArrayCollection instance will be returned to the
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
func (runeArrayCol *RuneArrayCollection) PopLastElement(
	errorPrefix interface{}) (
	lastRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"PopLastElement()",
		"")

	if err != nil {
		return lastRuneArrayDto, err
	}

	lastIdx := len(runeArrayCol.runeArrayDtoCol) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Rune Array Dto Collection is empty!\n",
			ePrefix.String())

		return lastRuneArrayDto, err
	}

	lastRuneArrayDto,
		err = runeArrayCollectionAtom{}.ptr().
		peekPopRuneArrayCol(
			runeArrayCol,
			lastIdx,
			true,
			ePrefix.XCpy(
				fmt.Sprintf("lastRuneArrayDto<-"+
					"runeArrayCol.runeArrayDtoCol[%v]",
					lastIdx)))

	return lastRuneArrayDto, err
}

// ReplaceAtIndex - Replaces a member of the Rune Array Collection
// with a user supplied RuneArrayDto.
//
// A deep copy of input parameter runeArrayDto will be used to make
// the replacement.
//
// The Rune Array Collection member element to be replaced is
// determined by input parameter 'replaceAtIndex'.
//
// When the operation is completed the old member element at index
// 'replaceAtIndex' will be deleted
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               RuneArrayDto
//	   - A deep copy of this RuneArrayDto instance will be created
//	     and added to the Rune Array Collection for the current
//	     instance of RuneArrayCollection.
//
//	     If 'runeArrayDto' contains a zero length character array,
//	     an error will be returned. Likewise, if 'runeArrayDto'
//	     contains an invalid Character Search Type
//	     'charSearchType', an error will be returned.
//
//
//	replaceAtIndex             int
//	   - The index of an element within the  Rune Array Collection
//	     maintained by the current RuneArrayDto instance. The
//	     array element at this index will be deleted and replaced
//	     by a deep copy of input parameter, 'runeArrayDto'.
//
//	     If 'replaceAtIndex' proves to be an invalid index, an error
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
// ------------------------------------------------------------------------
//
// Return Values
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
func (runeArrayCol *RuneArrayCollection) ReplaceAtIndex(
	runeArrayDto RuneArrayDto,
	replaceAtIndex int,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"ReplaceAtIndex()",
		"")

	if err != nil {
		return err
	}

	if len(runeArrayCol.runeArrayDtoCol) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Rune Array Collection is Empty!\n"+
			"'runeArrayCol.runeArrayDtoCol' has a length of zero.\n",
			ePrefix.String())

		return err
	}

	err = runeArrayDto.IsValidCharacterArrayError(
		ePrefix.XCpy(
			"runeArrayDto"))

	if err != nil {
		return err
	}

	err = runeArrayDto.IsValidCharacterSearchTypeError(
		ePrefix.XCpy(
			"runeArrayDto"))

	if err != nil {
		return err
	}

	// Length is already validated. It MUST BE
	// greater than zero.
	lenOfChars := runeArrayDto.GetRuneArrayLength()

	if replaceAtIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is invalid!\n"+
			"'replaceAtIndex' is less than zero (0).\n"+
			"replaceAtIndex = '%v'\n",
			ePrefix.String(),
			replaceAtIndex)

		return err
	}

	if replaceAtIndex >= lenOfChars {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is out of range and invalid!\n"+
			"'replaceAtIndex' is greater than the maximum Rune Array Collection index.\n"+
			"The last element in the Rune Array Collection is index '%v'.\n"+
			"Input parameter 'replaceAtIndex' = '%v'\n",
			ePrefix.String(),
			lenOfChars-1,
			replaceAtIndex)

		return err
	}

	runeArrayCol.runeArrayDtoCol[replaceAtIndex].Empty()

	runeArrayCol.runeArrayDtoCol[replaceAtIndex],
		err = runeArrayDto.CopyOut(
		ePrefix.XCpy(
			fmt.Sprintf("newRuneArray<-"+
				"runeArrayCol.runeArrayDtoCol[%v]",
				replaceAtIndex)))

	return err
}

func (runeArrayCol *RuneArrayCollection) SearchForTextCharacters(
	targetInputParms CharSearchTargetInputParametersDto,
	errorPrefix interface{}) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	errorSearchResults :=
		CharSearchRuneArrayResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"SearchForTextCharacters()",
		"")

	if err != nil {

		return errorSearchResults, err
	}

	lenRuneDtoCollection := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneDtoCollection == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Rune Array Collection is Empty!\n"+
			"runeArrayCol.runeArrayDtoCol has a length of zero.\n",
			ePrefix.String())

		return errorSearchResults, err
	}

	err = targetInputParms.IsValidInstanceError(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return errorSearchResults, err
	}

	var dtoSearchResults CharSearchRuneArrayResultsDto

	testConfigDto := CharSearchTestConfigDto{}.New()
	testConfigDto.RequestFoundTestCharacters = true
	for i := 0; i < lenRuneDtoCollection; i++ {

		dtoSearchResults,
			err = runeArrayCol.runeArrayDtoCol[i].
			SearchForTextCharacterString(
				targetInputParms,
				testConfigDto,
				ePrefix.XCpy(
					fmt.Sprintf("runeArrayCol.runeArrayDtoCol[%v]",
						i)))

		if err != nil {
			return errorSearchResults, err
		}

		if dtoSearchResults.FoundSearchTarget {

			dtoSearchResults.CollectionTestObjIndex = i

			return dtoSearchResults, err
		}
	}

	return errorSearchResults, err
}

// SetCollection - Deletes all members of the existing Rune Array
// Dto Collection maintained by the current instance of
// RuneArrayCollection and proceeds to instantiate a new collection
// based on the array of RuneArrayDto objects passed to this
// method.
//
// Input parameter 'runeArrayDtoCol' consists of an array of
// RuneArrayDto objects. Deep copies of these objects will be used
// to populate the new Rune Array Collection.
//
// If any member of the RuneArrayDto array, 'runeArrayDtoCol', is
// judged to be invalid, an error will be returned.
//
// Valid RuneArrayDto array members must satisfy two criteria in
// order to be classified as 'valid'.
//
//	(1) The RuneArrayDto internal Character Array must have a
//	    length greater than zero.
//
//	(2) The Character Search Type associated with each
//	    RuneArrayDto must be valid.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will DELETE all member elements in the existing
// collection of RuneArrayDto objects contained in the current
// instance of RuneArrayCollection.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDtoCol            []RuneArrayDto
//	   - An array of RuneArrayDto objects which be used to replace
//	     the existing Rune Array Dto Collection maintained by the
//	     current instance of RuneArrayCollection.
//
//	     Deep copies of each element in parameter 'runeArrayDtoCol'
//	     will be used to populate the new Rune Array Dto Collection
//	     contained in the current instance of RuneArrayCollection.
//
//	     If this parameter is submitted as a 'nil' or zero length
//	     array, an error will be returned.
//
//	     If any member element of this array is classified as
//	     invalid, an error will be returned.
//
//	     Valid RuneArrayDto array members must satisfy two criteria
//	     in order to be classified as 'valid'.
//
//	     (1) The RuneArrayDto internal Character Array must have a
//	         length greater than zero.
//
//	     (2) The Character Search Type associated with each
//	         RuneArrayDto must be valid.
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
// ------------------------------------------------------------------------
//
// Return Values
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
func (runeArrayCol *RuneArrayCollection) SetCollection(
	runeArrayDtoCol []RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"SetCollection()",
		"")

	if err != nil {
		return err
	}

	lenIncomingRuneArray := len(runeArrayDtoCol)

	if lenIncomingRuneArray == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'runeArrayDtoCol' is EMPTY!\n",
			ePrefix.String())

		return err
	}

	for i := 0; i < lenIncomingRuneArray; i++ {

		err = runeArrayDtoCol[i].IsValidCharacterArrayError(
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayDtoCol[%v] Character Array Error!",
					i)))

		if err != nil {
			return err
		}

		err = runeArrayDtoCol[i].IsValidCharacterSearchTypeError(
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayDtoCol[%v] Character Serach Type Error!",
					i)))

		if err != nil {
			return err
		}
	}

	runeArrayCollectionAtom{}.ptr().
		empty(runeArrayCol)

	runeArrayCol.runeArrayDtoCol =
		make([]RuneArrayDto, lenIncomingRuneArray)

	for j := 0; j < lenIncomingRuneArray; j++ {

		runeArrayCol.runeArrayDtoCol[j],
			err = runeArrayDtoCol[j].CopyOut(
			ePrefix.XCpy(
				fmt.Sprintf(
					"runeArrayCol.runeArrayDtoCol[%v]"+
						"<-runeArrayDtoCol[%v]",
					j,
					j)))

		if err != nil {
			return err
		}
	}

	return err
}
