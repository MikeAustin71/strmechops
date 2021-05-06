package strmech

import "sync"

// NumberSignChars - Used to identify numbers as either positive or
// negative values. These are most often used to classify numeric
// values contained in number strings. Numbers signs may be defined
// by a leading number sign, a trailing number sign or a
// combination of leading and trailing number signs.
//
type NumberSignChars struct {
	leadingNumSignChars  []rune
	trailingNumSignChars []rune
	lock                 *sync.Mutex
}

// GetLeadingNumSignChars - Returns an array of runes which
// represent the leading number sign characters contained in this
// instance of NumberSignChars.
//
func (numSignChars *NumberSignChars) GetLeadingNumSignChars() []rune {

	if numSignChars.lock == nil {
		numSignChars.lock = new(sync.Mutex)
	}

	numSignChars.lock.Lock()

	defer numSignChars.lock.Unlock()

	var leadingNumSignChars []rune

	lenNumSignChars := len(numSignChars.leadingNumSignChars)

	if numSignChars.leadingNumSignChars == nil ||
		lenNumSignChars == 0 {
		return leadingNumSignChars
	}

	leadingNumSignChars = make([]rune, lenNumSignChars)

	copy(
		leadingNumSignChars,
		numSignChars.leadingNumSignChars)

	return leadingNumSignChars
}

// GetTrailingNumSignChars - Returns an array of runes which
// represent the trailing number sign characters contained in this
// instance of NumberSignChars.
//
func (numSignChars *NumberSignChars) GetTrailingNumSignChars() []rune {

	if numSignChars.lock == nil {
		numSignChars.lock = new(sync.Mutex)
	}

	numSignChars.lock.Lock()

	defer numSignChars.lock.Unlock()

	var trailingNumSignChars []rune

	lenNumSignChars := len(numSignChars.trailingNumSignChars)

	if numSignChars.trailingNumSignChars == nil ||
		lenNumSignChars == 0 {
		return trailingNumSignChars
	}

	trailingNumSignChars = make([]rune, lenNumSignChars)

	copy(
		trailingNumSignChars,
		numSignChars.trailingNumSignChars)

	return trailingNumSignChars
}

// GetNumSignChars - Returns two arrays of runes representing the
// leading and trailing number sign characters contained in this
// instance of NumberSignChars.
//
func (numSignChars *NumberSignChars) GetNumSignChars() (
	leadingNumSignChars []rune,
	trailingNumSignChars []rune) {

	if numSignChars.lock == nil {
		numSignChars.lock = new(sync.Mutex)
	}

	numSignChars.lock.Lock()

	defer numSignChars.lock.Unlock()

	lenNumSignChars := len(numSignChars.trailingNumSignChars)

	if numSignChars.trailingNumSignChars != nil &&
		lenNumSignChars > 0 {
		trailingNumSignChars = make([]rune, lenNumSignChars)

		copy(
			trailingNumSignChars,
			numSignChars.trailingNumSignChars)
	}

	lenNumSignChars = len(numSignChars.leadingNumSignChars)

	if numSignChars.leadingNumSignChars != nil &&
		lenNumSignChars > 0 {
		leadingNumSignChars = make([]rune, lenNumSignChars)

		copy(
			leadingNumSignChars,
			numSignChars.leadingNumSignChars)
	}

	return leadingNumSignChars, trailingNumSignChars
}

// New - Returns a pointer to a new instance of NegativeNumberSign.
// This instance is populated based on the values of input
// parameters, 'leadingNumSignChars' and
// 'trailingNumSignChars'.
//
func (numSignChars NumberSignChars) New(
	leadingNumSignChars []rune,
	trailingNumSignChars []rune) NumberSignChars {

	if numSignChars.lock == nil {
		numSignChars.lock = new(sync.Mutex)
	}

	numSignChars.lock.Lock()

	defer numSignChars.lock.Unlock()

	newNegNumSign := NumberSignChars{}

	lenNumSign := len(leadingNumSignChars)

	if lenNumSign > 0 {

		newNegNumSign.leadingNumSignChars =
			make([]rune, lenNumSign)

		copy(
			newNegNumSign.leadingNumSignChars,
			leadingNumSignChars)
	}

	lenNumSign = len(trailingNumSignChars)

	if lenNumSign > 0 {

		newNegNumSign.trailingNumSignChars =
			make([]rune, len(trailingNumSignChars))

		copy(
			newNegNumSign.trailingNumSignChars,
			trailingNumSignChars)
	}

	return newNegNumSign
}

// SetNumSign - Overwrites and deletes the current values
// for internal member variables:
//   NegativeNumberSign.leadingNumSignChars
//   NegativeNumberSign.trailingNumSignChars
//
// New values are supplied to these internal member variables
// by input parameters, 'leadingNumSignChars' and
// 'trailingNumSignChars'.
//
// Effectively, this method resets the values for leading and
// trailing negative number signs maintained by the current
// instance of NegativeNumberSign.
//
func (numSignChars *NumberSignChars) SetNumSign(
	leadingNegativeNumSignChars []rune,
	trailingNegativeNumSignChars []rune) {

	if numSignChars.lock == nil {
		numSignChars.lock = new(sync.Mutex)
	}

	numSignChars.lock.Lock()

	defer numSignChars.lock.Unlock()

	lenNumSign := len(leadingNegativeNumSignChars)

	numSignChars.leadingNumSignChars = nil

	if lenNumSign > 0 {

		numSignChars.leadingNumSignChars =
			make([]rune, lenNumSign)

		copy(
			numSignChars.leadingNumSignChars,
			leadingNegativeNumSignChars)
	}

	lenNumSign = len(trailingNegativeNumSignChars)

	numSignChars.trailingNumSignChars = nil

	if lenNumSign > 0 {

		numSignChars.trailingNumSignChars =
			make([]rune, len(trailingNegativeNumSignChars))

		copy(
			numSignChars.trailingNumSignChars,
			trailingNegativeNumSignChars)
	}

}
