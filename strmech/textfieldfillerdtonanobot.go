package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textFieldFillerDtoNanobot - Provides helper methods for
// TextFieldFillerDto.
type textFieldFillerDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldFillerDto to a destination instance of
// TextFieldFillerDto.
func (txtFillerDtoNanobot *textFieldFillerDtoNanobot) copy(
	destinationTxtFillerDto *TextFieldFillerDto,
	sourceTxtFillerDto *TextFieldFillerDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFillerDtoNanobot.lock == nil {
		txtFillerDtoNanobot.lock = new(sync.Mutex)
	}

	txtFillerDtoNanobot.lock.Lock()

	defer txtFillerDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFillerDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFillerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	if destinationTxtFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFillerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	destinationTxtFillerDto.FormatType =
		sourceTxtFillerDto.FormatType

	destinationTxtFillerDto.LeftMarginStr =
		sourceTxtFillerDto.LeftMarginStr

	destinationTxtFillerDto.FillerCharacters =
		sourceTxtFillerDto.FillerCharacters

	destinationTxtFillerDto.FillerCharsRepeatCount =
		sourceTxtFillerDto.FillerCharsRepeatCount

	destinationTxtFillerDto.RightMarginStr =
		sourceTxtFillerDto.RightMarginStr

	destinationTxtFillerDto.LineTerminator =
		sourceTxtFillerDto.LineTerminator

	destinationTxtFillerDto.MaxLineLength =
		sourceTxtFillerDto.MaxLineLength

	destinationTxtFillerDto.TurnAutoLineLengthBreaksOn =
		sourceTxtFillerDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpacerDtoNanobot.
//
func (txtFillerDtoNanobot textFieldFillerDtoNanobot) ptr() *textFieldFillerDtoNanobot {

	if txtFillerDtoNanobot.lock == nil {
		txtFillerDtoNanobot.lock = new(sync.Mutex)
	}

	txtFillerDtoNanobot.lock.Lock()

	defer txtFillerDtoNanobot.lock.Unlock()

	return &textFieldFillerDtoNanobot{
		lock: new(sync.Mutex),
	}
}
