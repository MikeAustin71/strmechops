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
//  targetRuneArray            []rune
//     - All of the data in the input parameter rune array,
//       'sourceRuneArray', will be copied to this parameter,
//       'targetRuneArray'. All of the pre-existing data in
//       'targetRuneArray' will be deleted and replaced.
//
//
//  sourceRuneArray            []rune
//     - The contents of this rune array will be copied to input
//       parameter, 'targetRuneArray'.
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

		return
	}

	if targetRuneArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetRuneArray' is a nil pointer!\n",
			ePrefix.String())

		return
	}

	if *sourceRuneArray == nil {
		*targetRuneArray = nil
		return
	}

	// At this point, sourceRuneArray
	// IS NOT 'nil'!
	lenSrcRuneAry := len(*sourceRuneArray)

	if lenSrcRuneAry == 0 &&
		setZeroLenArrayToNil == true {

		*targetRuneArray = nil
		return

	} else if lenSrcRuneAry == 0 &&
		setZeroLenArrayToNil == false {

		*targetRuneArray = make([]rune, 0)
		return
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
//  targetIntArray             []rune
//     - All of the data in the input parameter rune array,
//       'sourceIntArray', will be copied to this parameter,
//       'targetIntArray'. All of the pre-existing data in
//       'targetIntArray' will be deleted and replaced.
//
//
//  sourceIntArray            []rune
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
//     - All of the data in the input parameter rune array,
//       'sourceUintArray', will be copied to this parameter,
//       'targetUintArray'. All of the pre-existing data in
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
//     - An array of runes. This rune array will searched to
//       identify the the beginning index of input parameter
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

// isTargetRunesIndex - Receives a host rune array and a starting
// index to that array. Beginning with the starting index this
// method determines whether the target rune array exists in the
// the host rune array beginning at the starting index.
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
//     - An array of runes. This rune array will searched to
//       determine if the target runes array is present at the
//       'hostStartIndex.
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
