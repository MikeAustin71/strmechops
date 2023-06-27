package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textStrBuilderAtom - Provides helper methods for type
// TextStrBuilder.
type textStrBuilderAtom struct {
	lock *sync.Mutex
}

func (txtBuilderAtom *textStrBuilderAtom) preBuildScreening(
	textParams *textStrBuilderParamsDto) error {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		textParams.errPrefDto,
		"textStrBuilderAtom."+
			"preBuildScreening()",
		"")

	if err != nil {
		return err
	}

	if textParams.strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: %v Text Builder\n"+
			"Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String(),
			textParams.sourceTag)

		if err != nil {
			return err
		}

	}

	if textParams.lenLineTerminatorStr > 0 {
		// Terminator String Length > 0
		// There is a command for line termination
		textParams.turnLineTerminationOff = false
	}

	if textParams.turnLineTerminationOff == true {

		textParams.lenLineTerminatorStr = 0
		textParams.lineTerminatorStr = ""

	} else {
		// MUST BE
		// textParams.turnLineTerminationOff == false

		if textParams.lenLineTerminatorStr == 0 {
			textParams.lenLineTerminatorStr = 1
			textParams.lineTerminatorStr = "\n"
		}
	}

	var adjustedMaximumLineLength = textParams.maxLineLength -
		textParams.lenLeftMarginStr -
		textParams.lenRightMarginStr -
		textParams.lenLineTerminatorStr

	if textParams.turnAutoLineLengthBreaksOn == false {

		textParams.maxLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true

		if textParams.maxLineLength < 1 {
			err = fmt.Errorf("%v\n"+
				"Error: Maximum Line Length is invalid!\n"+
				"'%v.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Maximum Line Length has a value less than one (1)\n"+
				"When Automatic Line Breaks are engaged, Maximum\n"+
				"Line Length must be set to a meaningful number.\n"+
				"Maximum Line Length = '%v'\n",
				ePrefix.String(),
				textParams.sourceDtoTag,
				textParams.maxLineLength)

			return err
		}

		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"This %v Text Line has turned on automatic line breaks. This\n"+
				"means that Line Termination Character(s) will be automatically inserted\n"+
				"when a string of text exceeds the Adjusted Maximum Line Length.\n"+
				"'%v.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				"minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters. With automatic line"+
				"breaks engaged, Adjusted Maximum Line Length MUST BE greater than\n"+
				"or equal to five (5).\n"+
				"SET Maximum Line Length to a meaningful number.\n"+
				"               Maximum Line Length =  %v \n"+
				"               Left Margin Length  = (%v)\n"+
				"               Right Margin Length = (%v)\n"+
				"Line Termination Characters Length = (%v)\n"+
				"                                   ------\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				textParams.sourceTag,
				textParams.sourceDtoTag,
				textParams.maxLineLength,
				textParams.lenLeftMarginStr,
				textParams.lenRightMarginStr,
				textParams.lenLineTerminatorStr,
				adjustedMaximumLineLength)

			return err
		}
	}

	totalStrLen := textParams.lenLineTerminatorStr +
		textParams.lenTextStr +
		textParams.lenRightMarginStr +
		textParams.lenLineTerminatorStr +
		(textParams.lenTextStr / adjustedMaximumLineLength)

	netCapacityStrBuilder :=
		textParams.strBuilder.Cap() -
			textParams.strBuilder.Len()

	requiredCapacity :=
		totalStrLen - netCapacityStrBuilder

	if requiredCapacity > 0 {

		textParams.strBuilder.Grow(requiredCapacity + 16)
	}

	txtBuilderElectron := textStrBuilderElectron{}

	textParams.currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.leftMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				fmt.Sprintf("%v.LeftMarginStr",
					textParams.sourceDtoTag)))

	if err != nil {
		return err
	}

	textParams.currentLineLength,
		textParams.lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeText(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.leftMarginStr,
			textParams.textStr,
			textParams.rightMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			textParams.multiLineLeftMarginStr,
			ePrefix.XCpy(
				fmt.Sprintf("%v",
					textParams.sourceTag)))

	if err != nil {
		return err
	}

	textParams.currentLineLength,
		textParams.lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.rightMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				fmt.Sprintf("%v.RightMarginStr",
					textParams.sourceDtoTag)))

	if err != nil {
		return err
	}

	if !textParams.lastWriteWasLineTerminator &&
		textParams.lenLineTerminatorStr > 0 {
		textParams.strBuilder.WriteString(
			textParams.lineTerminatorStr)
	}

	return err
}
