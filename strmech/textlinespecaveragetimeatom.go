package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

type textLineSpecAverageTimeAtom struct {
	lock *sync.Mutex
}

//	calcAverageDuration
//
//	Calculates the average duration of a timer event
//	encapsulated in an instance of TextLineSpecAverageTime
//	passed as input parameter 'txtLineAvgTimer'.
//
//	Average duration is calculated by dividing the total
//	time duration by the number of separate timing events.
//
//	In addition to Average Duration, this method also
//	returns the minimum and maximum time duration for
//	all timing events included in the average time
//	calculation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	avgDuration					int64
//
//		If this method completes successfully, this
//		parameter will return the average time duration
//		for all recorded timing events. Average duration
//		is calculated by dividing the total duration by
//		the number of separate timing events.
//
//	maximumTimeDuration			int64
//
//		If this method completes successfully, this
//		parameter will return the maximum time duration
//		from all recorded timing events.
//
//	minimumTimeDuration			int64
//
//		If this method completes successfully, this
//		parameter will return the minimum time duration
//		from all recorded timing events.
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
func (txtLineAvgTimeAtom *textLineSpecAverageTimeAtom) calcAverageDuration(
	txtLineAvgTimer *TextLineSpecAverageTime,
	errPrefDto *ePref.ErrPrefixDto) (
	avgDuration int64,
	maximumTimeDuration int64,
	minimumTimeDuration int64,
	err error) {

	if txtLineAvgTimeAtom.lock == nil {
		txtLineAvgTimeAtom.lock = new(sync.Mutex)
	}

	txtLineAvgTimeAtom.lock.Lock()

	defer txtLineAvgTimeAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeAtom."+
			"calcAverageDuration()",
		"")

	if err != nil {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	var err2 error

	_,
		err2 = new(textLineSpecAverageTimeElectron).
		testValidityOfTxtLineAvgTimer(
			txtLineAvgTimer,
			ePrefix.XCpy("txtLineAvgTimer"))

	if err2 != nil {

		err = fmt.Errorf("calcAverageDuration()\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"\n%v\n",
			err2.Error())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	bigZero := big.NewInt(0)

	cmpValue := txtLineAvgTimer.numberOfDurationEvents.Cmp(
		bigZero)

	if cmpValue == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"No timer events have been recorded.\n"+
			"'txtLineAvgTimer.numberOfDurationEvents' is equal to zero!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	cmpValue = txtLineAvgTimer.totalDurationNanoSecs.Cmp(
		bigZero)

	if cmpValue == -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"'txtLineAvgTimer.totalDurationNanoSecs' is less than zero!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if cmpValue == 0 {

		avgDuration = 0

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	cmpValue = txtLineAvgTimer.numberOfDurationEvents.Cmp(
		&txtLineAvgTimer.totalDurationNanoSecs)

	if cmpValue == 0 {

		avgDuration = 1

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if cmpValue == 1 {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	// Number of Duration Events is
	// less than total duration

	bigZero.Quo(
		&txtLineAvgTimer.totalDurationNanoSecs,
		&txtLineAvgTimer.numberOfDurationEvents)

	maxInt64 := big.NewInt(0).SetInt64(
		math.MaxInt64)

	cmpValue = bigZero.Cmp(maxInt64)

	if cmpValue == 1 {
		err = fmt.Errorf("%v\n"+
			"Calculation Error: Average Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"Average Duration = %v\n",
			ePrefix.String(),
			bigZero.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	avgDuration = bigZero.Int64()

	minimumTimeDuration =
		txtLineAvgTimer.maximumTimeDuration.Int64()

	minimumTimeDuration =
		txtLineAvgTimer.maximumTimeDuration.Int64()

	return avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err
}
