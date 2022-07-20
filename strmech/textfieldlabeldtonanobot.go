package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textFieldLabelDtoNanobot - Provides helper methods for
// TextFieldLabelDto
type textFieldLabelDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldLabelDto to a destination instance of
// TextFieldLabelDto.
func (txtLabelDtoNanobot *textFieldLabelDtoNanobot) copy(
	destinationTxtLabelDto *TextFieldLabelDto,
	sourceTxtLabelDto *TextFieldLabelDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLabelDtoNanobot.lock == nil {
		txtLabelDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelDtoNanobot.lock.Lock()

	defer txtLabelDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldLabelDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtLabelDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtLabelDto' is a nil pointer!\n",
			ePrefix.String())

	}

	if destinationTxtLabelDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtLabelDto' is a nil pointer!\n",
			ePrefix.String())

	}

	destinationTxtLabelDto.FormatType =
		sourceTxtLabelDto.FormatType

	destinationTxtLabelDto.LeftMarginStr =
		sourceTxtLabelDto.LeftMarginStr

	destinationTxtLabelDto.FieldText =
		sourceTxtLabelDto.FieldText

	destinationTxtLabelDto.FieldLength =
		sourceTxtLabelDto.FieldLength

	destinationTxtLabelDto.FieldJustify =
		sourceTxtLabelDto.FieldJustify

	destinationTxtLabelDto.RightMarginStr =
		sourceTxtLabelDto.RightMarginStr

	destinationTxtLabelDto.LineTerminator =
		sourceTxtLabelDto.LineTerminator

	destinationTxtLabelDto.MaxLineLength =
		sourceTxtLabelDto.MaxLineLength

	destinationTxtLabelDto.TurnAutoLineLengthBreaksOn =
		sourceTxtLabelDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldLabelDtoNanobot.
//
func (txtLabelDtoNanobot textFieldLabelDtoNanobot) ptr() *textFieldLabelDtoNanobot {

	if txtLabelDtoNanobot.lock == nil {
		txtLabelDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelDtoNanobot.lock.Lock()

	defer txtLabelDtoNanobot.lock.Unlock()

	return &textFieldLabelDtoNanobot{
		lock: new(sync.Mutex),
	}
}
