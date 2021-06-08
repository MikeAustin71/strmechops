package strmech

import "sync"

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
