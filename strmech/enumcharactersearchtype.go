package strmech

import "sync"

type CharacterSearchType int

var lockCharacterSearchType sync.Mutex

// None - Signals that the CharacterSearchType specification is
// empty or uninitialized. This is an error condition
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) None() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(-1)
}

// LinearTargetChars - Signals that each character in the Target
// Search String will be compared to each character in the Test
// String. The search will proceed for from left to right in
// Test Character Sequence.
//
// A 'Match' or successful search outcome is defined as the case
// where each character in the Target String matches each
// corresponding character in the Test String.
//
//  Example
//                             1         2         3
//            Index  0123456789012345678901234567890
//   Target String: "Hey, Xray-4 is the call sign."
//     Test String: "Xray"
//
//   In a Linear Target Characters Search, a match between
//   the Target String and Test String will be declared, if,
//   and only if, the search begins at Target String index
//   5.
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) LinearTargetChars() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(0)
}

// SingleTargetChar - Signals that a single character in the Target
// Search String will be compared to all characters in the Test
// String.
//
// If a single Target String character equals any character in the
// Test String, a 'Match' or successful search outcome will
// be declared.
//
// The search will proceed from left to right in the Target String.
// Each Target String Character will be compared to all characters
// in the Test String looking for the first matching Test String
// Character.
//
//
//  Example
//                             1         2         3
//            Index  0123456789012345678901234567890
//   Target String: "Hey, Xray-4 is the call sign."
//     Test String: "XZyURJK"
//
//   In a Single Target Character Search, a match between the
//   Target String and Test String will be declared at Target
//   String indexes 5 and 8.
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) SingleTargetChar() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(1)
}
