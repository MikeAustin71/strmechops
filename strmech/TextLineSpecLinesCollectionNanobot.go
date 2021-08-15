package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecLinesCollectionNanobot struct {
	lock *sync.Mutex
}

// emptyCollection - Deletes all the text lines from an instance
// TextLineSpecLinesCollection.
//
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) emptyCollection(
	textLinesCol *TextLineSpecLinesCollection,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionNanobot."+
			"emptyCollection()",
		"")

	if err != nil {
		return err
	}

	if textLinesCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLinesCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	for i := 0; i < len(textLinesCol.textLines); i++ {

		textLinesCol.textLines[i].Empty()

		textLinesCol.textLines[i] = nil
	}

	textLinesCol.textLines = nil

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecLinesCollectionNanobot.
//
func (txtLinesColNanobot textLineSpecLinesCollectionNanobot) ptr() *textLineSpecLinesCollectionNanobot {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	return &textLineSpecLinesCollectionNanobot{
		lock: new(sync.Mutex),
	}
}
