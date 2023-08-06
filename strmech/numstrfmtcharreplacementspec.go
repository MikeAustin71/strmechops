package strmech

import "sync"

// NumStrFmtCharReplacementSpec
//
// The Number String Character Replacement Format
// Specification is used to replace characters in a
// text string with numeric values, usually integer
// values.
//
// As such the format is useful in formatting such
// numbers as telephone numbers, identification numbers
// and inventory numbers.
//
//	Telephone Number Example:
//
//		NumStrFmtCharReplacementSpec.NumberFormat =
//			"(NNN) NNN-NNNN"
//
//		NumStrFmtCharReplacementSpec.NumReplacementChar =
//			'N'
//
//		Integer Digits: 0115550101
//
//		Formatted Number String: (011) 555-0101
//
// This type is used by type NumberStrKernel to generate
// formatted number strings using the 'Character
// Replacement' technique. Reference method:
//
//	NumberStrKernel.FmtCharReplacementStr()
type NumStrFmtCharReplacementSpec struct {
	NumberFormat string
	//	This string should contain the Number Replacement
	//	Character defined in member variable
	//	'NumReplacementChar'. The Number Replacement
	//	Character will be replaced by numeric digits
	//	in the NumberFormat string.
	//
	//	Example:
	//		NumberFormat = "(NNN) NNN-NNNN"
	//		'NumReplacementChar' = 'N'
	//		Formatted Number String: "(NNN) NNN-NNNN"
	//
	//		The letter 'N' will be replaced with numeric
	//		digits. See Type NumberStrKernel, Method:
	//			NumberStrKernel.FmtCharReplacementStr()
	//

	NumReplacementChar rune
	//	This rune character will serve as a placeholder
	//	in the NumberFormat string described above. Every
	//	instance of this character will be replaced by a
	//	numeric digit character.

	lock *sync.Mutex
}
