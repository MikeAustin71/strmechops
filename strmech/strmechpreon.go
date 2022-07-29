package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type strMechPreon struct {
	lock *sync.Mutex
}

// copyRuneArrays - Copies a source rune array to a target rune
// array.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that all the data in 'targetRuneArray' will be
// deleted and replaced.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetRuneArray            *[]rune
//     - A pointer to the target rune array. All the data in the
//       input parameter rune array, 'sourceRuneArray', will be
//       copied to this parameter, 'targetRuneArray'.
//
//       All the pre-existing data in 'targetRuneArray' will be
//       deleted and replaced with data copied from 'sourceRuneArray'.
//
//       If the 'targetRuneArray' pointer is 'nil', an error will
//       be returned.
//
//
//  sourceRuneArray            *[]rune
//     - A pointer to the source rune array. The contents of this
//       rune array will be copied to input parameter,
//       'targetRuneArray'.
//
//       If the 'sourceRuneArray' pointer is 'nil', an error will
//       be returned.
//
//       If the concrete rune array pointed to by this pointer is
//       nil, 'targetRuneArray' concrete rune array is set to 'nil'
//       and NO error is returned.
//
//
//  setZeroLenArrayToNil       bool
//     - If sourceRuneArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'true', 'targetRuneArray'
//       will be set to 'nil'.
//
//       If sourceRuneArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'false',
//       'targetRuneArray' will be set to a zero length array.
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
func (sMechPreon *strMechPreon) copyRuneArrays(
	targetRuneArray *[]rune,
	sourceRuneArray *[]rune,
	setZeroLenArrayToNil bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.copyRuneArrays()",
		"")

	if err != nil {
		return err
	}

	if sourceRuneArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceRuneArray' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if targetRuneArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetRuneArray' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if *sourceRuneArray == nil {
		*targetRuneArray = nil
		return err
	}

	// At this point, concrete sourceRuneArray
	// IS NOT 'nil'!
	lenSrcRuneAry := len(*sourceRuneArray)

	if lenSrcRuneAry == 0 &&
		setZeroLenArrayToNil == true {

		*targetRuneArray = nil

		return err

	} else if lenSrcRuneAry == 0 &&
		setZeroLenArrayToNil == false {

		*targetRuneArray = make([]rune, 0)

		return err
	}

	*targetRuneArray = make([]rune, lenSrcRuneAry)

	itemsCopied := copy(*targetRuneArray, *sourceRuneArray)

	if itemsCopied != lenSrcRuneAry {
		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"Runes copied does not equal length of Source Rune Array\n"+
			"Length Source Rune Array: '%v'\n"+
			"  Number of Runes Copied: '%v'\n",
			ePrefix.String(),
			lenSrcRuneAry,
			itemsCopied)
	}

	return err
}

// copyIntegerArrays - Copies a source integer array to a target
// integer array.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that all the data in 'targetIntArray' will be
// deleted and replaced.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetIntArray             *[]int
//     - All the data in the input parameter rune array,
//       'sourceIntArray', will be copied to this parameter,
//       'targetIntArray'. All the pre-existing data in
//       'targetIntArray' will be deleted and replaced.
//
//
//  sourceIntArray             *[]int
//     - The contents of this rune array will be copied to input
//       parameter, 'targetIntArray'.
//
//
//  setZeroLenArrayToNil       bool
//     - If sourceIntArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'true', 'targetIntArray'
//       will be set to 'nil'.
//
//       If sourceIntArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'false', 'targetIntArray'
//       will be set to a zero length array.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
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
func (sMechPreon *strMechPreon) copyIntegerArrays(
	targetIntArray *[]int,
	sourceIntArray *[]int,
	setZeroLenArrayToNil bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.copyIntegerArrays()",
		"")

	if err != nil {
		return err
	}

	if sourceIntArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceIntArray' is a nil pointer!\n",
			ePrefix.String())

		return
	}

	if targetIntArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetIntArray' is a nil pointer!\n",
			ePrefix.String())

		return
	}

	if *sourceIntArray == nil {
		*targetIntArray = nil
		return
	}

	// At this point, sourceIntArray
	// IS NOT 'nil'!
	lenSrcIntAry := len(*sourceIntArray)

	if lenSrcIntAry == 0 &&
		setZeroLenArrayToNil == true {

		*targetIntArray = nil
		return

	} else if lenSrcIntAry == 0 &&
		setZeroLenArrayToNil == false {

		*targetIntArray = make([]int, 0)
		return
	}

	*targetIntArray = make([]int, lenSrcIntAry)

	itemsCopied := copy(*targetIntArray, *sourceIntArray)

	if itemsCopied != lenSrcIntAry {
		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"Runes copied does not equal length of Source "+
			"Integer Array\n"+
			"Length Source Integer Array: '%v'\n"+
			"  Number of Integers Copied: '%v'\n",
			ePrefix.String(),
			lenSrcIntAry,
			itemsCopied)
	}

	return err
}

// copyUnsignedIntArrays - Copies a source unsigned integer array
// to a target unsigned integer array.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that all the data in 'targetUintArray' will be
// deleted and replaced.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetUintArray            []rune
//     - All the data in the input parameter rune array,
//       'sourceUintArray', will be copied to this parameter,
//       'targetUintArray'. All the pre-existing data in
//       'targetUintArray' will be deleted and replaced.
//
//
//  sourceUintArray            []rune
//     - The contents of this rune array will be copied to input
//       parameter, 'targetUintArray'.
//
//
//  setZeroLenArrayToNil       bool
//     - If sourceUintArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'true', 'targetUintArray'
//       will be set to 'nil'.
//
//       If sourceUintArray is NOT 'nil', has a zero length and
//       'setZeroLenArrayToNil' is set to 'false',
//       'targetUintArray' will be set to a zero length array.
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
func (sMechPreon *strMechPreon) copyUnsignedIntArrays(
	targetUintArray *[]uint,
	sourceUintArray *[]uint,
	setZeroLenArrayToNil bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.copyUnsignedIntArrays()",
		"")

	if err != nil {
		return err
	}

	if sourceUintArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceUintArray' is a nil pointer!\n",
			ePrefix.String())

		return
	}

	if targetUintArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetUintArray' is a nil pointer!\n",
			ePrefix.String())

		return
	}

	if *sourceUintArray == nil {
		*targetUintArray = nil
		return
	}

	// At this point, sourceUintArray
	// IS NOT 'nil'!
	lenSrcUintAry := len(*sourceUintArray)

	if lenSrcUintAry == 0 &&
		setZeroLenArrayToNil == true {

		*targetUintArray = nil
		return

	} else if lenSrcUintAry == 0 &&
		setZeroLenArrayToNil == false {

		*targetUintArray = make([]uint, 0)
		return
	}

	*targetUintArray = make([]uint, lenSrcUintAry)

	itemsCopied := copy(*targetUintArray, *sourceUintArray)

	if itemsCopied != lenSrcUintAry {
		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"Runes copied does not equal length of Source "+
			"Unsigned Integer Array\n"+
			"Length Source Unsigned Integer Array: '%v'\n"+
			"  Number of Unsigned Integers Copied: '%v'\n",
			ePrefix.String(),
			lenSrcUintAry,
			itemsCopied)
	}

	return err
}

// equalRuneArrays - Receives two rune arrays and proceeds to
// determine if they are equal.
//
// If the two rune arrays are equivalent, this method returns
// 'true'. Otherwise, the method returns 'false'.
//
// If one array is 'nil' and the other is a zero length array,
// this method will return 'true'.
//
func (sMechPreon *strMechPreon) equalRuneArrays(
	runeAryOne []rune,
	runeAryTwo []rune) (
	areEqual bool) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	lenAryOne := len(runeAryOne)

	lenAryTwo := len(runeAryTwo)

	if lenAryOne != lenAryTwo {
		return false
	}

	if lenAryOne == 0 {
		// They are equal but both have a zero length.
		return true
	}

	for i := 0; i < lenAryOne; i++ {
		if runeAryOne[i] != runeAryTwo[i] {
			return false
		}
	}

	return true
}

// equalIntArrays - Receives two integer arrays and proceeds to
// determine if they are equal.
//
// If the two integer arrays are equivalent, this method returns
// 'true'. Otherwise, this method returns 'false'.
//
// If one array is 'nil' and the other is a zero length array,
// this method will return 'true'.
//
func (sMechPreon *strMechPreon) equalIntArrays(
	integerAryOne []int,
	integerAryTwo []int) (
	areEqual bool) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	lenAryOne := len(integerAryOne)

	lenAryTwo := len(integerAryTwo)

	if lenAryOne != lenAryTwo {
		return false
	}

	if lenAryOne == 0 {
		// They are equal but both have a zero length.
		return true
	}

	for i := 0; i < lenAryOne; i++ {
		if integerAryOne[i] != integerAryTwo[i] {
			return false
		}
	}

	return true
}

// equalUintArrays - Receives two unsigned integer arrays and
// proceeds to determine if they are equal.
//
// If the two unsigned integer arrays are equivalent, this method
// returns 'true'. Otherwise, this method returns 'false'.
//
// If one array is 'nil' and the other is a zero length array,
// this method will return 'true'.
//
func (sMechPreon *strMechPreon) equalUintArrays(
	uintAryOne []uint,
	uintAryTwo []uint) (
	areEqual bool) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	lenAryOne := len(uintAryOne)

	lenAryTwo := len(uintAryTwo)

	if lenAryOne != lenAryTwo {
		return false
	}

	if lenAryOne == 0 {
		// They are equal but both have a zero length.
		return true
	}

	for i := 0; i < lenAryOne; i++ {
		if uintAryOne[i] != uintAryTwo[i] {
			return false
		}
	}

	return true
}

// findRunesInRunes - Locates an array of target runes within an
// array of host runes.
//
// If the target runes are located, this method returns the index
// within the host runes ('foundIndex') where the target runes were
// located.
//
// If the target runes are NOT located within the hast runes array,
// this method returns a 'foundIndex' value of -1.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hostRunes           []rune
//     - An array of runes. This rune array will be searched to
//       identify the beginning index of input parameter
//       'targetRunes'. If 'hostRunes' is a zero length array, an
//       error will be returned.
//
//
//  hostStartIndex      int
//     - The starting index within the host runes array where
//       the search operation will commence. If 'hostStartIndex' is
//       less than zero, it will be automatically set to zero.
//
//       If the 'hostStartIndex' is greater than or equal to the
//       length of 'hostRunes', the return value of 'foundIndex'
//       will be set to -1 and no error will be returned.
//
//
//  targetRunes         []rune
//     - The object of the search. The 'hostRunes' will be searched
//       beginning at the 'hostRunes' starting index to determine
//       whether these 'targetRunes' exists in the 'hostRunes'
//       array. If 'targetRunes' is a zero length array, an error
//       will be returned.
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
//  foundIndex          int
//     - If the 'targetRunes' array is located within the
//       'hostRunes' array, this parameter will contain the index
//       in the 'hostRunes' array where 'targetRunes' array begins.
//       If 'targetRunes' are located within the 'hostRunes' array,
//       this parameter will always be set to an integer value
//       greater than or equal to zero.
//
//       If the 'targetRunes' array is NOT located within the
//       'hostRunes' array, this parameter will be set to an
//       integer value of negative one (-1).
//
//
//  err                 error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (sMechPreon *strMechPreon) findRunesInRunes(
	hostRunes []rune,
	hostStartIndex int,
	targetRunes []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	foundIndex int,
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundIndex = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.findRunesInRunes()",
		"")

	if err != nil {
		return foundIndex, err
	}

	lenHostRunes := len(hostRunes)

	if lenHostRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'hostRunes' is an empty array!\n",
			ePrefix.String())

		return foundIndex, err
	}

	lenTargetRunes := len(targetRunes)

	if lenTargetRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetRunes' is an empty array!\n",
			ePrefix.String())

		return foundIndex, err
	}

	if hostStartIndex < 0 {
		hostStartIndex = 0
	}

	if hostStartIndex >= lenHostRunes {
		return foundIndex, err
	}

	lastHostIndex := lenHostRunes - 1

	lastTargetIndex := lenTargetRunes - 1

	if lastTargetIndex+hostStartIndex > lastHostIndex {
		return foundIndex, err
	}

	for i := hostStartIndex; i < lenHostRunes; i++ {

		if hostRunes[i] != targetRunes[0] {
			continue
		}

		if lastTargetIndex+i > lastHostIndex {
			return -1, err
		}

		j := i

		hitCount := 0

		for k := 0; k < lenTargetRunes; k++ {
			if targetRunes[k] == hostRunes[j] {
				hitCount++
				j++
			} else {
				break
			}
		}

		if hitCount == lenTargetRunes {
			foundIndex = i
			return foundIndex, err
		}

	}

	return foundIndex, err
}

// getRepeatRuneArray - Returns a slice of runes containing a single
// rune array repeated multiple times.
//
// -----------------------------------------------------------------
//
// Important
//
// The maximum number of times that the input parameter rune array
// ('repeatRuneChars') will be repeated is limited only by
// available memory.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  runeRepeatCount            int
//     - An integer value specifying the number of times a rune
//       array will be repeated in the slice of runes returned by
//       this method.
//
//       If 'runeRepeatCount' is less than zero (0), an error will
//       be returned.
//
//       If 'runeRepeatCount' is set to zero (0), the returned
//       slice runes will be set to a 'nil' value and NO error will
//       be returned.
//
//       BE CAREFUL, the maximum value of 'runeRepeatCount' is
//       limited only by available memory.
//
//
//  repeatRuneChars            []rune
//     - This rune array specifies the text characters which will
//       be repeated multiple times in the slice of runes returned
//       by this method.
//
//       If this value is equal to integer zero (0), an error will
//       be returned.
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
//  runeChars                  []rune
//     - If this method completes successfully, this parameter will
//       return a slice of runes containing a single text character
//       repeated multiple times.
//
//       The single text character is specified by input parameter,
//       'repeatRuneChar'. The number of times this character is
//       repeated will be specified by input parameter,
//       'runeRepeatCount'.
//
//       If this method completes successfully, the length of
//       'runeChars' will be equal to 'repeatRuneChar'.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (sMechPreon *strMechPreon) getRepeatRuneArray(
	runeRepeatCount int,
	repeatRuneChars []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	runeChars []rune,
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	runeChars = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon."+
			"getRepeatRuneArray()",
		"")

	if err != nil {
		return runeChars, err
	}

	if runeRepeatCount < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeRepeatCount' is invalid!\n"+
			"'runeRepeatCount' has a value less than zero (0).\n"+
			"runeRepeatCount = '%v'\n",
			ePrefix.String(),
			runeRepeatCount)

		return runeChars, err
	}

	lenRepeatRuneChars := len(repeatRuneChars)

	if lenRepeatRuneChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'repeatRuneChars' is INVALID!\n"+
			"The lenght of 'repeatRuneChars' is equal to zero (0).\n",
			ePrefix.String())

		return runeChars, err
	}

	if runeRepeatCount == 0 {

		return runeChars, err
	}

	runeChars = make([]rune,
		lenRepeatRuneChars*runeRepeatCount)

	j := 0

	for i := 0; i < runeRepeatCount; i++ {
		for k := 0; k < lenRepeatRuneChars; k++ {
			runeChars[j] = repeatRuneChars[k]
			j++
		}

	}

	return runeChars, err
}

// getRepeatRuneChar - Returns a slice of runes containing a single
// rune character repeated multiple times.
//
// -----------------------------------------------------------------
//
// Important
//
// The maximum number of times that the input parameter rune
// character ('repeatRuneChar') will be repeated is limited only by
// available memory.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  runeRepeatCount            int
//     - An integer value specifying the number of times a rune
//       character will be repeated in the slice of runes returned
//       by this method.
//
//       If 'runeRepeatCount' is less than zero (0), an error will
//       be returned.
//
//       If 'runeRepeatCount' is set to zero (0), the returned
//       slice runes will be set to a 'nil' value and NO error will
//       be returned.
//
//       BE CAREFUL, the maximum value of 'runeRepeatCount' is
//       limited only by available memory.
//
//
//  repeatRuneChar             rune
//     - This rune specifies the text character which will be
//       repeated multiple times in the slice of runes returned by
//       this method.
//
//       If this value is equal to integer zero (0), an error will
//       be returned.
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
//  runeChars                  []rune
//     - If this method completes successfully, this parameter will
//       return a slice of runes containing a single text character
//       repeated multiple times.
//
//       The single text character is specified by input parameter,
//       'repeatRuneChar'. The number of times this character is
//       repeated will be specified by input parameter,
//       'runeRepeatCount'.
//
//       If this method completes successfully, the length of
//       'runeChars' will be equal to 'repeatRuneChar'.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (sMechPreon *strMechPreon) getRepeatRuneChar(
	runeRepeatCount int,
	repeatRuneChar rune,
	errPrefDto *ePref.ErrPrefixDto) (
	runeChars []rune,
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	runeChars = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.getRepeatRuneChar()",
		"")

	if err != nil {
		return runeChars, err
	}

	if runeRepeatCount < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeRepeatCount' is invalid!\n"+
			"'runeRepeatCount' has a value less than zero (0).\n"+
			"runeRepeatCount = '%v'\n",
			ePrefix.String(),
			runeRepeatCount)

		return runeChars, err
	}

	if repeatRuneChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'repeatRuneChar' is INVALID!\n"+
			"'repeatRuneChar' is equal to zero (0).\n",
			ePrefix.String())

		return runeChars, err
	}

	if runeRepeatCount == 0 {

		return runeChars, err
	}

	runeChars = make([]rune, runeRepeatCount)

	for i := 0; i < runeRepeatCount; i++ {
		runeChars[i] = repeatRuneChar
	}

	return runeChars, err
}

// isTargetRunesIndex - Receives a host rune array and a starting
// index to that array. Beginning with the starting index this
// method determines whether the target rune array exists in the
// host rune array beginning at the starting index.
//
// If the target rune array is found in the host array at the host
// array starting index, this method returns true.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  hostRunes           []rune
//     - An array of runes. This rune array will be searched to
//       determine if the target runes array is present at the
//       'hostStartIndex'.
//
//       If 'hostRunes' is a zero length array, this method will
//       return 'false'.
//
//
//  hostStartIndex      int
//     - The starting index within the host runes array where
//       the search operation will commence. If 'hostStartIndex' is
//       less than zero, it will be automatically set to zero.
//
//       If the 'hostStartIndex' is greater than or equal to the
//       length of 'hostRunes', this method will return 'false'.
//
//
//  targetRunes         []rune
//     - The object of the search. The 'hostRunes' will be searched
//       beginning at the 'hostRunes' starting index to determine
//       whether these 'targetRunes' exists beginning that starting
//       index. If the target rune array is NOT found beginning at
//       the staring index in the host runes array, this method will
//       return 'false'.
//
//       If the target runes array IS found in the host runes array
//       beginning at the host runes starting index, this method
//       will return 'true'.
//
//       If 'targetRunes' is an empty array, this method will
//       return 'false'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isTargetRunesIndex  bool
//     - A boolean flag signaling whether the target runes array
//       was found in the host runes array beginning at the host
//       runes staring index.
//
//       If the target runes array is found at the staring index in
//       the host runes array, this method will return 'true'.
//
func (sMechPreon *strMechPreon) isTargetRunesIndex(
	hostRunes []rune,
	hostStartIndex int,
	targetRunes []rune) (
	isTargetRunesIndex bool) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	isTargetRunesIndex = false

	lenHostRunes := len(hostRunes)

	if lenHostRunes == 0 {
		return isTargetRunesIndex
	}

	lenTargetRunes := len(targetRunes)

	if lenTargetRunes == 0 {
		return isTargetRunesIndex
	}

	if hostStartIndex < 0 {
		hostStartIndex = 0
	}

	if hostStartIndex >= lenHostRunes {
		return isTargetRunesIndex
	}

	lastHostIndex := lenHostRunes - 1

	lastTargetIndex := lenTargetRunes - 1

	if lastTargetIndex+hostStartIndex > lastHostIndex {
		return isTargetRunesIndex
	}

	for k := 0; k < lenTargetRunes; k++ {

		if targetRunes[k] != hostRunes[hostStartIndex] {

			return isTargetRunesIndex

		}

		hostStartIndex++
	}

	isTargetRunesIndex = true

	return isTargetRunesIndex
}

// setRepeatRuneChar - Receives a pointer to a rune array and
// proceeds to populate that rune array with a single rune
// character which is repeated multiple times as specified by
// input parameter, 'runeRepeatCount'.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that all the data in 'targetRuneArray' will be
// deleted and replaced.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetRuneArray            *[]rune
//     - A pointer to a rune target rune array. This rune array
//       will be configured with a single rune which is repeated
//       multiple times as specified by the input parameter,
//       'runeRepeatCount'.
//
//       All the pre-existing data in 'targetRuneArray' will be
//       deleted and replaced.
//
//       If the 'targetRuneArray' pointer is 'nil', an error will
//       be returned.
//
//
//  runeRepeatCount            int
//     - An integer value specifying the number of times a rune
//       character will be repeated in the rune array specified by
//       input parameter, 'targetRuneArray'
//
//       If 'runeRepeatCount' is less than zero (0), an error will
//       be returned.
//
//       If 'runeRepeatCount' is set to zero (0), input parameter
//       'targetRuneArray' will be set to 'nil', and NO error will
//       be returned.
//
//       BE CAREFUL, the maximum value of 'runeRepeatCount' is
//       limited only by available memory.
//
//
//  repeatRuneChar             rune
//     - This rune specifies the text character which will be
//       repeated in input parameter 'targetRuneArray'.
//
//       If this value is equal to integer zero (0), an error will
//       be returned.
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
func (sMechPreon *strMechPreon) setRepeatRuneChar(
	targetRuneArray *[]rune,
	runeRepeatCount int,
	repeatRuneChar rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon.setRepeatRuneChar()",
		"")

	if err != nil {
		return err
	}

	if targetRuneArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetRuneArray' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if runeRepeatCount < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeRepeatCount' is invalid!\n"+
			"'runeRepeatCount' has a value less than zero (0).\n"+
			"runeRepeatCount = '%v'\n",
			ePrefix.String(),
			runeRepeatCount)

		return err
	}

	if repeatRuneChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'repeatRuneChar' is INVALID!\n"+
			"'repeatRuneChar' is equal to zero (0).\n",
			ePrefix.String())

		return err
	}

	if runeRepeatCount == 0 {

		*targetRuneArray = nil

		return err
	}

	*targetRuneArray = make([]rune, runeRepeatCount)

	sourceArray := make([]rune, runeRepeatCount)

	for i := 0; i < runeRepeatCount; i++ {
		sourceArray[i] = repeatRuneChar
	}

	itemsCopied := copy(*targetRuneArray, sourceArray)

	if itemsCopied != runeRepeatCount {
		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"Runes copied does not equal length of 'runeRepeatCount'\n"+
			"Length 'runeRepeatCount': '%v'\n"+
			"  Number of Runes Copied: '%v'\n",
			ePrefix.String(),
			runeRepeatCount,
			itemsCopied)
	}

	return err
}

// testValidityOfRuneCharArray - Performs a diagnostic analysis on
// an array of runes to determine if the characters are valid.
//
// If the rune array is equal to 'nil', the array is judged to be
// invalid and an error will be returned.
//
// If the rune array is a zero length array, the array is judged to
// be invalid and an error will be returned.
//
// If any of the array elements are equal to integer zero
// (char==0), that character element invalidates the entire array
// and an error will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes consisting entirely of alphanumeric
//       characters. This method will evaluate this array to
//       determine whether it is valid.
//
//       If the rune array is equal to 'nil', the array is judged
//       to be invalid and an error will be returned.
//
//       If the rune array is a zero length array, the array is
//       judged to be invalid and an error will be returned.
//
//       If any of the array elements are equal to integer zero
//       (char == 0), that character element invalidates the entire
//       array and an error will be returned.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If the input parameter 'charArray' is determined to be
//       valid, this parameter will be set to 'true'. If
//       'charArray' is invalid, this parameter will be set to
//       'false'.
//
//
//  err                        error
//     - If the input parameter 'charArray' is determined to be
//       valid, this parameter will be set to 'nil'.
//
//       If 'charArray' is invalid, the returned error Type will
//       encapsulate an appropriate error message. This returned
//       error message will also incorporate the method chain and
//       text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (sMechPreon *strMechPreon) testValidityOfRuneCharArray(
	charArray []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon."+
			"testValidityOfRuneCharArray()",
		"")

	if err != nil {
		return isValid, err
	}

	if charArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid.\n"+
			"'charArray' is equal to 'nil'!\n",
			ePrefix.String())

		return isValid, err
	}

	lenCharArray := len(charArray)

	if lenCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid.\n"+
			"'charArray' is a zero length array!\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenCharArray; i++ {

		if charArray[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'charArray' is invalid.\n"+
				"'charArray' contains an invalid character element!\n"+
				"charArray[%v] == 0",
				ePrefix.String(),
				i)

			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}

// testValidityOfRuneIntArray - Performs a diagnostic analysis on
// an array of runes to determine if all the character values
// in the array constitute integer digits '0' (0x30) through '9'
// (0x39), inclusive.
//
// If the rune array is equal to 'nil', the array is judged to be
// invalid and an error will be returned.
//
// If the rune array is a zero length array, the array is judged to
// be invalid and an error will be returned.
//
// If any of the array elements are equal to an integer value of
// zero (0), the array is judged to be invalid and an error will be
// returned.
//
// If any of the array elements specify text characters which are
// NOT integer digit characters zero ('0' or 0x30) through nine
// ('9' or 0x39) inclusive, the array is judged to be invalid and
// an error will be returned.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intDigitsArray             []rune
//     - An array of runes consisting entirely of numeric text
//       characters which represent integer digits zero ('0' or
//       0x30) through ('9' or 0x39) inclusive. This method will
//       evaluate this array to determine whether it is valid.
//
//       If the rune array is 'nil' or a zero length array, the
//       array is judged to be invalid and an error will be
//       returned.
//
//       If any of the array elements are equal to an integer
//       value zero (0), the array is judged to be invalid and an
//       error will be returned.
//
//       If any of the array elements specify text characters which
//       are NOT integer digit characters zero ('0' or 0x30)
//       through nine ('9' or 0x39) inclusive, the array is judged
//       to be invalid and an error will be returned.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If the input parameter 'intDigitsArray' is determined to
//       be valid, this parameter will be set to 'true'. If
//       'intDigitsArray' is invalid, this parameter will be set
//       to 'false'.
//
//
//  err                        error
//     - If the input parameter 'intDigitsArray' is determined to
//       be valid, this parameter will be set to 'nil'.
//
//       If 'intDigitsArray' is invalid, the returned error Type
//       will encapsulate an appropriate error message. This
//       returned error message will also incorporate the method
//       chain and text passed by input parameter, 'errPrefDto'.
//       The 'errPrefDto' text will be attached to the beginning of
//       the error message.
//
func (sMechPreon *strMechPreon) testValidityOfRuneIntArray(
	intDigitsArray []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechPreon."+
			"testValidityOfRuneIntArray()",
		"")

	if err != nil {
		return isValid, err
	}

	if intDigitsArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intDigitsArray' is invalid.\n"+
			"'intDigitsArray' is equal to 'nil'!\n",
			ePrefix.String())

		return isValid, err
	}

	lenCharArray := len(intDigitsArray)

	if lenCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intDigitsArray' is invalid.\n"+
			"'charArray' is a zero length array!\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenCharArray; i++ {

		if intDigitsArray[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'intDigitsArray' is invalid.\n"+
				"'intDigitsArray' contains an invalid character element!\n"+
				"intDigitsArray[%v] == 0",
				ePrefix.String(),
				i)

			return isValid, err
		}

		if intDigitsArray[i] < '0' ||
			intDigitsArray[i] > '9' {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'intDigitsArray' is invalid.\n"+
				"'intDigitsArray' contains an invalid character element!"+
				"This text character is NOT an inter digit '0'-'9', inclusive.\n"+
				"intDigitsArray[%v] == '%v''",
				ePrefix.String(),
				i,
				string(intDigitsArray[i]))

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// strMechPreon.
//
func (sMechPreon strMechPreon) ptr() *strMechPreon {

	if sMechPreon.lock == nil {
		sMechPreon.lock = new(sync.Mutex)
	}

	sMechPreon.lock.Lock()

	defer sMechPreon.lock.Unlock()

	return &strMechPreon{
		lock: new(sync.Mutex),
	}
}
