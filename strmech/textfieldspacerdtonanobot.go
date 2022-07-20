package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textFieldSpacerDtoNanobot - Provides helper methods for
// TextFieldSpacerDto.
type textFieldSpacerDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldSpacerDto to a destination instance of
// TextFieldSpacerDto.
func (txtSpacerDtoNanobot *textFieldSpacerDtoNanobot) copy(
	destinationTxtSpacerDto *TextFieldSpacerDto,
	sourceTxtSpacerDto *TextFieldSpacerDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtSpacerDtoNanobot.lock == nil {
		txtSpacerDtoNanobot.lock = new(sync.Mutex)
	}

	txtSpacerDtoNanobot.lock.Lock()

	defer txtSpacerDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpacerDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtSpacerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtSpacerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	if destinationTxtSpacerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtSpacerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	destinationTxtSpacerDto.FormatType =
		sourceTxtSpacerDto.FormatType

	destinationTxtSpacerDto.LeftMarginStr =
		sourceTxtSpacerDto.LeftMarginStr

	destinationTxtSpacerDto.FieldLength =
		sourceTxtSpacerDto.FieldLength

	destinationTxtSpacerDto.RightMarginStr =
		sourceTxtSpacerDto.RightMarginStr

	destinationTxtSpacerDto.LineTerminator =
		sourceTxtSpacerDto.LineTerminator

	destinationTxtSpacerDto.MaxLineLength =
		sourceTxtSpacerDto.MaxLineLength

	destinationTxtSpacerDto.TurnAutoLineLengthBreaksOn =
		sourceTxtSpacerDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpacerDtoNanobot.
//
func (txtSpacerDtoNanobot textFieldSpacerDtoNanobot) ptr() *textFieldSpacerDtoNanobot {

	if txtSpacerDtoNanobot.lock == nil {
		txtSpacerDtoNanobot.lock = new(sync.Mutex)
	}

	txtSpacerDtoNanobot.lock.Lock()

	defer txtSpacerDtoNanobot.lock.Unlock()

	return &textFieldSpacerDtoNanobot{
		lock: new(sync.Mutex),
	}
}
