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
//	Telephone Number Example: (555) 555-0101
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
	//	'NumFmtReplacementChar'. The Number Replacement
	//	Character will be replaced by numbers.
	//
	//	Example:
	//		NumberFormat = "(NNN) NNN-NNNN
	//		'NumFmtReplacementChar' = 'N'
	//
	//		The letter 'N' will be replaced with numeric
	//		digits. See Type NumberStrKernel
	NumFmtReplacementChar rune

	lock *sync.Mutex
}
