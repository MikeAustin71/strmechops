package strmech

import (
	"sync"
)

type textLineSpecLinesCollectionNanobot struct {
	lock *sync.Mutex
}

// emptyCollection - Deletes all the text lines from an instance
// TextLineSpecLinesCollection.
//
func (txtLinesColNanobot *textLineSpecLinesCollectionNanobot) emptyCollection(
	textLinesCol *TextLineSpecLinesCollection) {

	if txtLinesColNanobot.lock == nil {
		txtLinesColNanobot.lock = new(sync.Mutex)
	}

	txtLinesColNanobot.lock.Lock()

	defer txtLinesColNanobot.lock.Unlock()

	if textLinesCol == nil {
		return
	}

	for i := 0; i < len(textLinesCol.textLines); i++ {

		if textLinesCol.textLines[i] == nil {
			continue
		}

		textLinesCol.textLines[i].Empty()

		textLinesCol.textLines[i] = nil
	}

	textLinesCol.textLines = nil

	return
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
