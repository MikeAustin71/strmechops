package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type strMechPreon struct {
	lock *sync.Mutex
}

// equalRuneArrays - Receives two rune arrays and proceeds to
// determine if they are equal.
//
// If the two rune arrays are equivalent, this method returns
// 'true'. Otherwise, the method returns 'false'.
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
//       'targetRunes'.
//
//  hostStartIndex      int
//     - The starting index within the host runes array where
//       the search operation will commence.
//
//  targetRunes         []rune
//     - The object of the search. The 'hostRunes' will be searched
//       beginning at the 'hostRunes' starting index to determine
//       whether these 'targetRunes' exists in the 'hostRunes'
//       array.
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
