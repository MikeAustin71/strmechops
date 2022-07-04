package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// RuneArrayCollection - A collection of Rune Array Dto objects.
//
// Essentially, this is a collection, or an array, of rune arrays.
//
type RuneArrayCollection struct {
	RuneArrayDtoCol []RuneArrayDto

	lock *sync.Mutex
}

// AddLatinAlphabetEnglishDto - Adds a RuneArrayDto to the Rune Array
// Collection. This RuneArrayDto is populated with the Latin
// Alphabet (English Version). The total number of characters
// is 52 comprised of 26 lower case letters and 26 upper case
// letters.
//
// An array of alphabetic characters in useful in search
// operations looking for alphabetic characters classified as
// delimiters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (runeArrayCol *RuneArrayCollection) AddLatinAlphabetEnglishDto() {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	latinAlphabetRuneArray := []rune{
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

	lenLatinAlphabet := len(latinAlphabetRuneArray)

	runeArrayDto := RuneArrayDto{}

	runeArrayDto.CharsArray = make([]rune, lenLatinAlphabet)

	for i := 0; i < lenLatinAlphabet; i++ {
		runeArrayDto.CharsArray[i] = latinAlphabetRuneArray[i]
	}

	runeArrayDto.charSearchType =
		CharSearchType.SingleTargetChar()

	runeArrayCol.RuneArrayDtoCol =
		append(
			runeArrayCol.RuneArrayDtoCol,
			runeArrayDto)

}

// AddNumericDigitsDto - Adds a RuneArrayDto to the Rune Array
// Collection. This RuneArrayDto is populated with numeric digits
// zero (0) through nine (9).
//
// An array of numeric digits in useful in search operations looking
// for numeric characters classified as delimiters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (runeArrayCol *RuneArrayCollection) AddNumericDigitsDto() {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	numericDigitsArray := []rune{
		'0',
		'1',
		'2',
		'3',
		'4',
		'5',
		'6',
		'7',
		'8',
		'9',
	}

	runeArrayDto := RuneArrayDto{}

	runeArrayDto.CharsArray = make([]rune, 10)

	for i := 0; i < 10; i++ {
		runeArrayDto.CharsArray[i] = numericDigitsArray[i]
	}

	runeArrayDto.charSearchType =
		CharSearchType.SingleTargetChar()

	runeArrayCol.RuneArrayDtoCol =
		append(
			runeArrayCol.RuneArrayDtoCol,
			runeArrayDto)
}

// AddRuneArrayDto - Receives an instance of RuneArrayDto and
// appends that instance to the RuneArrayDto Collection.
//
// This differs from method:
//  RuneArrayCollection.AddRuneArrayDtoDeepCopy()
//
// This method appends the passed RuneArrayDto instance to the
// collection. The Deep Copy method appends a copy of the
// RuneArrayDto to the collection.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// RuneArrayDto instances added to this collection MUST HAVE a
// valid Character Search Type. Reference RuneArrayDto member
// variable 'RuneArrayDto.charSearchType'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayDto               RuneArrayDto
//     - An instance of RuneArrayDto. This instance will be
//       appended to the RuneArrayDto collection maintained by this
//       instance of RuneArrayCollection.
//
//       If 'runeArrayDto' has an invalid character search type
//       (runeArrayDto.charSearchType), an error will be returned.
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
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
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

	runeArrayCol.RuneArrayDtoCol =
		append(runeArrayCol.RuneArrayDtoCol, runeArrayDto)

	return err
}

// AddRuneArrayDtoDeepCopy - Receives an instance of RuneArrayDto and
// appends a deep copy of that instance to the RuneArrayDto
// Collection.
//
// This differs from method:
//  RuneArrayCollection.AddRuneArrayDto()
//
// This method appends a deep copy of the passed RuneArrayDto
// instance to the collection. The 'AddRuneArrayDto()' method
// appends the actual RuneArrayDto instance to the collection.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// RuneArrayDto instances added to this collection MUST HAVE a
// valid Character Search Type. Reference RuneArrayDto member
// variable 'RuneArrayDto.charSearchType'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayDto               RuneArrayDto
//     - An instance of RuneArrayDto. A deep copy of this instance
//       will be appended to the RuneArrayDto collection maintained
//       by this instance of RuneArrayCollection.
//
//       If 'runeArrayDto' has an invalid character search type
//       (runeArrayDto.charSearchType), an error will be returned.
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
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
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

	runeArrayCol.RuneArrayDtoCol =
		append(runeArrayCol.RuneArrayDtoCol, deepCopyRuneArrayDto)

	return err
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
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingRuneArrayCol       *RuneArrayCollection
//     - A pointer to an instance of RuneArrayCollection. The
//       internal member variable data values in this instance will
//       be compared to those in the current instance of
//       RuneArrayCollection. The results of this comparison
//       will be returned to the calling functions as a boolean
//       value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingRuneArrayCol' are equivalent
//       in all respects to those contained in the current instance
//       of RuneArrayCollection, this return value will be set
//       to 'true'.
//
//       Otherwise, this method will return 'false'.
//
func (runeArrayCol *RuneArrayCollection) Equal(
	incomingRuneArrayCol *RuneArrayCollection) bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	if incomingRuneArrayCol == nil {
		return false
	}

	lenOfRuneArrayDtoCol := len(runeArrayCol.RuneArrayDtoCol)

	if lenOfRuneArrayDtoCol !=
		len(incomingRuneArrayCol.RuneArrayDtoCol) {

		return false
	}

	// Collection Lengths are Equal!
	if lenOfRuneArrayDtoCol == 0 {
		return true
	}

	for i := 0; i < lenOfRuneArrayDtoCol; i++ {

		if !runeArrayCol.RuneArrayDtoCol[i].Equal(
			&incomingRuneArrayCol.RuneArrayDtoCol[i]) {

			return false
		}
	}

	// All elements of the current instance collection are
	// equal to all corresponding elements of the incoming
	// instance collection.
	return true
}

// GetNumberOfRuneArrayDtos - Returns the number of elements in the
// RuneArrayDto collection. The returned integer value is therefore
// equal to the length of the internal array of RuneArrayDto
// objects.
//
func (runeArrayCol *RuneArrayCollection) GetNumberOfRuneArrayDtos() int {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	return len(runeArrayCol.RuneArrayDtoCol)
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
//
func (runeArrayCol *RuneArrayCollection) IsNOP() bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	if len(runeArrayCol.RuneArrayDtoCol) == 0 {

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
// 'RuneArrayCollection.RuneArrayDtoCol' is greater than zero, this
// method will return 'true'.
//
// If the length of internal member variable
// RuneArrayCollection.RuneArrayDtoCol is equal to zero, this
// method will return 'false'.
//
// This method is identical in function to method:
//   RuneArrayCollection.IsValidInstanceError
//
// The only difference is that this method returns a boolean value,
// while 'IsValidInstanceError()' returns an error.
//
func (runeArrayCol *RuneArrayCollection) IsValidInstance() bool {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	if len(runeArrayCol.RuneArrayDtoCol) > 0 {
		return true
	}

	return false
}

// IsValidInstanceError - Returns an error if the current instance
// of RuneArrayCollection is invalid.
//
// There is only one criterion for classifying an instance of
// RuneArrayCollection as valid. It must contain a Rune Array Dto
// Collection where the number of elements is greater than zero.
//
// If the length of internal member variable
// 'RuneArrayCollection.RuneArrayDtoCol' is greater than zero, this
// method will return 'nil' signaling "No Error".
//
// If the length of internal member variable
// RuneArrayCollection.RuneArrayDtoCol is equal to zero, this
// method will return an error containing an appropriate error
// message.
//
// This method is identical in function to method:
//   RuneArrayCollection.IsValidInstance
//
// The only difference is that this method returns an error, while
// 'IsValidInstance()' returns a boolean value.
//
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

	if len(runeArrayCol.RuneArrayDtoCol) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: This instance of RuneArrayCollection contains\n"+
			"an empty collection of RuneArrayDto objects. The length\n"+
			"of RuneArrayCollection.RuneArrayDtoCol is zero.\n"+
			"This instance of RuneArrayCollection is therefore invalid!\n",
			ePrefix.String())

	}

	return err
}

func (runeArrayCol *RuneArrayCollection) SearchForTextCharacters(
	targetInputParms CharSearchTargetInputParametersDto,
	errorPrefix interface{}) (
	CharSearchResultsDto,
	error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	errorSearchResults := CharSearchResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"SearchForTextCharacters()",
		"")

	if err != nil {

		return errorSearchResults, err
	}

	lenRuneDtoCollection := len(runeArrayCol.RuneArrayDtoCol)

	if lenRuneDtoCollection == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The Rune Array Collection is Empty!\n"+
			"runeArrayCol.RuneArrayDtoCol has a length of zero.\n",
			ePrefix.String())

		return errorSearchResults, err
	}

	err = targetInputParms.IsValidInstanceError(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return errorSearchResults, err
	}

	var dtoSearchResults CharSearchResultsDto

	for i := 0; i < lenRuneDtoCollection; i++ {

		targetInputParms.CollectionTestObjIndex = i

		dtoSearchResults,
			err = runeArrayCol.RuneArrayDtoCol[i].
			SearchForTextCharacterString(
				targetInputParms,
				ePrefix.XCpy(
					fmt.Sprintf("runeArrayCol.RuneArrayDtoCol[%v]",
						i)))

		if err != nil {
			return dtoSearchResults, err
		}

		targetInputParms.CollectionTestObjIndex = -1

		if dtoSearchResults.FoundSearchTarget {

			return dtoSearchResults, err
		}
	}

	return errorSearchResults, err
}
