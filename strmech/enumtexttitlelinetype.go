package strmech

import "sync"

// Lock lockTextTitleLineType before accessing these
// 'maps'.

var mapTextTitleLineTypeCodeToString = map[TextTileLineType]string{
	TextTileLineType(0): "None",
	TextTileLineType(1): "LeadingMarqueeLine",
	TextTileLineType(2): "TitleLine",
	TextTileLineType(3): "TrailingMarqueeLine",
}

var mapTextTileLineTypeStringToCode = map[string]TextTileLineType{
	"None":                TextTileLineType(0),
	"LeadingMarqueeLine":  TextTileLineType(1),
	"TitleLine":           TextTileLineType(2),
	"TrailingMarqueeLine": TextTileLineType(3),
}

var mapTextTileLineTypeLwrCaseStringToCode = map[string]TextTileLineType{
	"none":                TextTileLineType(0),
	"leadingmarqueeline":  TextTileLineType(1),
	"titleline":           TextTileLineType(2),
	"trailingmarqueeline": TextTileLineType(3),
}

type TextTileLineType int

var lockTextTitleLineType sync.Mutex

// None
//
// Signals that 'TextTileLineType' has not been
// initialized and therefore has no value. This is an
// error condition.
func (txtTitleLineType TextTileLineType) None() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(0)
}

// LeadingMarqueeLine
//
// Classifies the line of text as a leading title marquee
// text line.
func (txtTitleLineType TextTileLineType) LeadingMarqueeLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(1)

}

// TitleLine
//
// Classifies the line of text as a main text title line.
// Text titles lines are typically displayed in the center
// of the Title Marquee.
func (txtTitleLineType TextTileLineType) TitleLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(2)

}

// TrailingMarqueeLine
//
// Classifies the line of text as a trailing title marquee
// text line.
func (txtTitleLineType TextTileLineType) TrailingMarqueeLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(3)

}

// String
//
// Returns a string with the name of the enumeration
// associated with this instance of TextTileLineType.
//
// This is a standard utility method and is not part
// of the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
//	t:= TextTileLineType(0).TrailingMarqueeLine()
//	str := t.String()
//	   str is now equal to 'TrailingMarqueeLine'
func (txtTitleLineType TextTileLineType) String() string {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	result, ok :=
		mapTextTitleLineTypeCodeToString[txtTitleLineType]

	if !ok {
		return "Error: TextTileLineType code UNKNOWN!"
	}

	return result
}

// XIsValid
//
// Returns a boolean value signaling whether the current
// TextTileLineType value is valid.
//
// Be advised, the enumeration value "None" is considered
// NOT VALID.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	titleLineVal := TextTileLineType(0).TrailingMarqueeLine()
//
//	isValid := titleLineVal.XIsValid() // isValid == true
//
//	titleLineVal = TextTileLineType(0).None()
//
//	isValid = titleLineVal.XIsValid() // isValid == false
func (txtTitleLineType TextTileLineType) XIsValid() bool {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return new(textTileLineTypeNanobot).
		isValidTextTitleLineType(
			txtTitleLineType)
}

// TitleLineType
//
// A public global constant of type TextTileLineType.
//
// This variable serves as an easier, shorthand technique
// for accessing TextTileLineType values.
//
// Usage:
// TitleLineType.None(),
// TitleLineType.LeadingMarqueeLine(),
// TitleLineType.TitleLine(),
// TitleLineType.TrailingMarqueeLine(),
const TitleLineType = TextTileLineType(0)

// textTileLineTypeNanobot
//
// Provides helper methods for enumeration
// TextTileLineType.
type textTileLineTypeNanobot struct {
	lock *sync.Mutex
}

// isValidTextTitleLineType
//
// Receives an instance of TextTileLineType and returns a
// boolean value signaling whether that TextTileLineType
// instance is valid.
//
// If the passed instance of TextTileLineType is valid,
// this method returns 'true'.
//
// Be advised, the enumeration value "None" is considered
// NOT VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of
// the valid TextTileLineType enumeration.
func (txtTileLineTypeNanobot *textTileLineTypeNanobot) isValidTextTitleLineType(
	txtTileLineType TextTileLineType) bool {

	if txtTileLineTypeNanobot.lock == nil {
		txtTileLineTypeNanobot.lock = new(sync.Mutex)
	}

	txtTileLineTypeNanobot.lock.Lock()

	defer txtTileLineTypeNanobot.lock.Unlock()

	if txtTileLineType < 1 ||
		txtTileLineType > 3 {

		return false
	}

	return true
}
