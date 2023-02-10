package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
	"time"
)

// textLineSpecAverageTimeMechanics
//
// Provides helper methods for type TextLineSpecAverageTime
type textLineSpecAverageTimeMechanics struct {
	lock *sync.Mutex
}

//	addDurationEvent
//
//	Adds a time duration value to the internal average
//	time counters maintained by an instance of
//	TextLineSpecAverageTime ('txtLineAvgTimer').
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal average time counters maintained by
//		this instance will be updated with the new data
//		passed as input parameter, 'eventDuration'.
//
//	eventDuration				time.Duration
//
//		The time duration associated with an event. This
//		duration value will be added to the internal
//		average time counters maintained by the
//		TextLineSpecAverageTime instance passed as input
//		parameter, 'txtLineAvgTimer'.
//
//		This time duration data is recorded for purposes
//		of later generating an average time duration
//		report for a series of timing events.
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
func (txtLineAvgTimeMech *textLineSpecAverageTimeMechanics) addDurationEvent(
	txtLineAvgTimer *TextLineSpecAverageTime,
	eventDuration time.Duration,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineAvgTimeMech.lock == nil {
		txtLineAvgTimeMech.lock = new(sync.Mutex)
	}

	txtLineAvgTimeMech.lock.Lock()

	defer txtLineAvgTimeMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeMechanics."+
			"addDurationEvent()",
		"")

	if err != nil {
		return err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_ = txtLineAvgTimer.numberOfDurationEvents.Add(
		&txtLineAvgTimer.numberOfDurationEvents,
		big.NewInt(1))

	bigDuration := big.NewInt(int64(eventDuration))

	_ = txtLineAvgTimer.totalDurationNanoSecs.Add(
		&txtLineAvgTimer.totalDurationNanoSecs,
		bigDuration)

	var comparison int

	comparison = txtLineAvgTimer.maximumTimeDuration.Cmp(
		bigDuration)

	if comparison == -1 {
		txtLineAvgTimer.maximumTimeDuration.Set(bigDuration)
	}

	comparison = txtLineAvgTimer.minimumTimeDuration.Cmp(
		big.NewInt(0))

	if comparison < 1 {

		txtLineAvgTimer.minimumTimeDuration.Set(bigDuration)

	} else {

		comparison = txtLineAvgTimer.minimumTimeDuration.Cmp(bigDuration)

		if comparison == 1 {

			txtLineAvgTimer.minimumTimeDuration.Set(bigDuration)
		}

	}

	return err
}

//	addStartStopEvent
//
//	Receives an event start time and event end time. This
//	method will then proceed to compute a time duration
//	from these input time values and record the time
//	duration in the internal average time counters
//	maintained by an instance of TextLineSpecAverageTime
//	passed as an input parameter ('txtLineAvgTimer').
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal average time counters maintained by
//		this instance will be updated with the new data
//		passed as input parameter, 'eventDuration'.
//
//	startTime					time.Time
//		The time at which a specific event started.
//
//	endTime						time.Time
//		The time at which a specific event ended.
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
func (txtLineAvgTimeMech *textLineSpecAverageTimeMechanics) addStartStopEvent(
	txtLineAvgTimer *TextLineSpecAverageTime,
	startTime time.Time,
	endTime time.Time,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineAvgTimeMech.lock == nil {
		txtLineAvgTimeMech.lock = new(sync.Mutex)
	}

	txtLineAvgTimeMech.lock.Lock()

	defer txtLineAvgTimeMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeMechanics."+
			"addStartStopEvent()",
		"")

	if err != nil {
		return err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if startTime.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTime' is INVALID!\n"+
			"'startTime' has a zero value.\n",
			ePrefix.String())

		return err
	}

	if endTime.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTime' is INVALID!\n"+
			"'endTime' has a zero value.\n",
			ePrefix.String())

		return err
	}

	if startTime.Equal(endTime) {
		return err
	}

	timeFormat :=
		new(textSpecificationMolecule).
			getDefaultDateTimeFormat()

	if !startTime.Before(endTime) {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTime' is INVALID!\n"+
			"'startTime' occurs after 'endTime'.\n"+
			"startTime = %v\n"+
			"  endTime = %v\n",
			ePrefix.String(),
			startTime.Format(timeFormat),
			endTime.Format(timeFormat))

		return err

	}

	eventDuration := endTime.Sub(startTime)

	_ = txtLineAvgTimer.numberOfDurationEvents.Add(
		&txtLineAvgTimer.numberOfDurationEvents,
		big.NewInt(1))

	bigDuration := big.NewInt(int64(eventDuration))

	_ = txtLineAvgTimer.totalDurationNanoSecs.Add(
		&txtLineAvgTimer.totalDurationNanoSecs,
		bigDuration)

	var comparison int

	comparison = txtLineAvgTimer.maximumTimeDuration.Cmp(
		bigDuration)

	if comparison == -1 {
		txtLineAvgTimer.maximumTimeDuration.Set(bigDuration)
	}

	comparison = txtLineAvgTimer.minimumTimeDuration.Cmp(
		big.NewInt(0))

	if comparison < 1 {

		txtLineAvgTimer.minimumTimeDuration.Set(big.NewInt(0))

	} else {
		// minimum vs zero == 1

		comparison = txtLineAvgTimer.minimumTimeDuration.Cmp(bigDuration)

		if comparison == 1 {

			txtLineAvgTimer.minimumTimeDuration.Set(bigDuration)
		}

	}

	return err
}

//	getFormattedText
//
//	Receives an instance of TextLineSpecAverageTime and,
//	using the internal timing counters, calculates the
//	average time duration for all recorded events.
//
//	Average time duration data including the minimum and
//	maximum time duration for this series of timing
//	events is compiled as text and reported to the
//	calling function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal average time counters maintained by
//		this instance will be used to produce an average
//		time duration for all recorded timing events.
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
//	string
//
//		If this method completes successfully, this
//		string will contain a report of average time
//		duration, minimum time duration and maximum
//		time duration for the series of timing events
//		recorded in the instance of
//		TextLineSpecAverageTime passed as input parameter
//		'txtLineAvgTimer'.
//
//	err							error
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
func (txtLineAvgTimeMech *textLineSpecAverageTimeMechanics) getFormattedText(
	txtLineAvgTimer *TextLineSpecAverageTime,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineAvgTimeMech.lock == nil {
		txtLineAvgTimeMech.lock = new(sync.Mutex)
	}

	txtLineAvgTimeMech.lock.Lock()

	defer txtLineAvgTimeMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeMechanics."+
			"getFormattedText()",
		"")

	if err != nil {
		return err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var avgDuration, maximumTimeDuration,
		minimumTimeDuration int64

	avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err = new(textLineSpecAverageTimeAtom).
		calcAverageDuration(
			txtLineAvgTimer,
			ePrefix.XCpy(
				"txtLineAvgTimer"))

	if err != nil {
		return err
	}

	var allocatedAvgDuration, allocatedMaxDuration,
		allocatedMinDuration TimeDurationDto

	allocatedAvgDuration,
		err = new(DateTimeHelper).AllocateTimeDuration(
		avgDuration,
		ePrefix.XCpy(
			"avgDuration"))

	if err != nil {
		return err
	}

	txtLineAvgTimeAtom := textLineSpecAverageTimeAtom{}

	maxLineLength := 60

	err = txtLineAvgTimeAtom.getFullDurationReport(
		txtLineAvgTimer,
		strBuilder,
		allocatedAvgDuration,
		true,
		"Average Duration",
		maxLineLength,
		ePrefix.XCpy(
			"<-allocatedAvgDuration"))

	if err != nil {
		return err
	}

	allocatedMaxDuration,
		err = new(DateTimeHelper).AllocateTimeDuration(
		maximumTimeDuration,
		ePrefix.XCpy(
			"maximumTimeDuration"))

	if err != nil {
		return err
	}

	err = txtLineAvgTimeAtom.getFullDurationReport(
		txtLineAvgTimer,
		strBuilder,
		allocatedMaxDuration,
		false,
		"Maximum Duration",
		maxLineLength,
		ePrefix.XCpy(
			"<-allocatedMaxDuration"))

	if err != nil {
		return err
	}

	allocatedMinDuration,
		err = new(DateTimeHelper).AllocateTimeDuration(
		minimumTimeDuration,
		ePrefix.XCpy(
			"minimumTimeDuration"))

	if err != nil {
		return err
	}

	err = txtLineAvgTimeAtom.getFullDurationReport(
		txtLineAvgTimer,
		strBuilder,
		allocatedMinDuration,
		false,
		"Minimum Duration",
		maxLineLength,
		ePrefix.XCpy(
			"<-allocatedMinDuration"))

	if err != nil {
		return err
	}

	strBuilder.WriteString("\n")

	return err
}
