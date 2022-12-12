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
