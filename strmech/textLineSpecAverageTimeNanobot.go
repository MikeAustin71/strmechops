package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textLineSpecAverageTimeNanobot
//
// Provides helper methods for type TextLineSpecAverageTime
type textLineSpecAverageTimeNanobot struct {
	lock *sync.Mutex
}

func (txtLineAvgTimeNanobot *textLineSpecAverageTimeNanobot) copy(
	destinationAvgTimer *TextLineSpecAverageTime,
	sourceAvgTimer *TextLineSpecAverageTime,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineAvgTimeNanobot.lock == nil {
		txtLineAvgTimeNanobot.lock = new(sync.Mutex)
	}

	txtLineAvgTimeNanobot.lock.Lock()

	defer txtLineAvgTimeNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeNanobot."+
			"copy()",
		"")

	if err != nil {
		return err
	}

	if destinationAvgTimer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceAvgTimer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textLineSpecAverageTimeElectron).empty(
		destinationAvgTimer)

	destinationAvgTimer.numberOfDurationEvents.Set(
		&sourceAvgTimer.numberOfDurationEvents)

	destinationAvgTimer.totalDurationNanoSecs.Set(
		&sourceAvgTimer.totalDurationNanoSecs)

	return err
}
