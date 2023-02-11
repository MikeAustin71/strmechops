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

//	copy
//
//	Copies all data from input parameter 'sourceAvgTimer'
//	to input parameter 'destinationAvgTimer'.
//
//	Be advised that the pre-existing data fields in input
//	parameter 'destinationAvgTimer' will be overwritten and
//	deleted.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationAvgTimer			*TextLineSpecAverageTime
//
//		A pointer to an instance of
//		TextLineSpecAverageTime. Data extracted from input
//		parameter 'incomingTimerLines' will be copied to
//		this input parameter, 'destinationAvgTimer'. If
//		this method completes successfully, all member
//		data variables encapsulated in 'destinationAvgTimer'
//		will be identical to those contained in input
//		parameter 'sourceAvgTimer'.
//
//		Be advised that the pre-existing data fields in
//		input parameter 'destinationAvgTimer' will be
//		overwritten and deleted.
//
//	sourceAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecTimerLines.
//
//		All data values in this TextLineSpecAverageTime
//		instance will be copied to input parameter
//		'destinationAvgTimer'.
//
//		The original member variable data values
//		encapsulated in 'sourceAvgTimer' will remain
//		unchanged and will NOT be overwritten or deleted.
//
//		If 'sourceAvgTimer' contains invalid member data
//		variables, this method will return an error.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
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

	_,
		err = new(textLineSpecAverageTimeElectron).
		testValidityOfTxtLineAvgTimer(
			sourceAvgTimer,
			ePrefix.XCpy("sourceAvgTimer"))

	if err != nil {
		return err
	}

	new(textLineSpecAverageTimeElectron).empty(
		destinationAvgTimer)

	destinationAvgTimer.numberOfDurationEvents.Set(
		&sourceAvgTimer.numberOfDurationEvents)

	destinationAvgTimer.totalDurationNanoSecs.Set(
		&sourceAvgTimer.totalDurationNanoSecs)

	destinationAvgTimer.maximumTimeDuration.Set(
		&sourceAvgTimer.maximumTimeDuration)

	destinationAvgTimer.minimumTimeDuration.Set(
		&sourceAvgTimer.minimumTimeDuration)

	destinationAvgTimer.applyAbbreviatedReportFormat =
		sourceAvgTimer.applyAbbreviatedReportFormat

	return err
}
