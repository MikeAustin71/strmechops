package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

// textLineSpecAverageTimeElectron
//
// Provides helper methods for type TextLineSpecAverageTime
type textLineSpecAverageTimeElectron struct {
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
func (txtLineAvgTimeElectron *textLineSpecAverageTimeElectron) calcAverageDuration(
	txtLineAvgTimer *TextLineSpecAverageTime,
	errPrefDto *ePref.ErrPrefixDto) (
	avgDuration int64,
	maximumTimeDuration int64,
	minimumTimeDuration int64,
	err error) {

	if txtLineAvgTimeElectron.lock == nil {
		txtLineAvgTimeElectron.lock = new(sync.Mutex)
	}

	txtLineAvgTimeElectron.lock.Lock()

	defer txtLineAvgTimeElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeElectron."+
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

	bigZero := big.NewInt(0)

	cmpValue := txtLineAvgTimer.numberOfDurationEvents.Cmp(
		bigZero)

	if cmpValue == -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"'txtLineAvgTimer.numberOfDurationEvents' is less than zero!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

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

	cmpValue = txtLineAvgTimer.maximumTimeDuration.Cmp(
		bigZero)

	if cmpValue == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"'txtLineAvgTimer.maximumTimeDuration' is less than zero!\n"+
			"txtLineAvgTimer.maximumTimeDuration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.maximumTimeDuration.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	cmpValue = txtLineAvgTimer.minimumTimeDuration.Cmp(
		bigZero)

	if cmpValue == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"'txtLineAvgTimer.minimumTimeDuration' is less than zero!\n"+
			"txtLineAvgTimer.minimumTimeDuration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.minimumTimeDuration.Text(10))

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

	cmpValue = txtLineAvgTimer.minimumTimeDuration.Cmp(maxInt64)

	if cmpValue == 1 {

		err = fmt.Errorf("%v\n"+
			"Calculation Error: Minimum Time Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"Minimum Time Duration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.minimumTimeDuration.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err

	}

	cmpValue = txtLineAvgTimer.maximumTimeDuration.Cmp(maxInt64)

	if cmpValue == 1 {

		err = fmt.Errorf("%v\n"+
			"Calculation Error: Maximum Time Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"Maximum Time Duration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.maximumTimeDuration.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	minimumTimeDuration =
		txtLineAvgTimer.maximumTimeDuration.Int64()

	minimumTimeDuration =
		txtLineAvgTimer.maximumTimeDuration.Int64()

	return avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err
}

//	empty
//
//	Receives a pointer to an instance of
//	TextLineSpecAverageTime and proceeds to set all the
//	internal member variables to their uninitialized or
//	zero states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The data values of all member variables contained in
//	input parameter 'txtLineAvgTimer' will be deleted and
//	reset to their zero values.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineAvgTimeElectron *textLineSpecAverageTimeElectron) empty(
	txtLineAvgTimer *TextLineSpecAverageTime) {

	if txtLineAvgTimeElectron.lock == nil {
		txtLineAvgTimeElectron.lock = new(sync.Mutex)
	}

	txtLineAvgTimeElectron.lock.Lock()

	defer txtLineAvgTimeElectron.lock.Unlock()

	if txtLineAvgTimer == nil {

		return
	}

	txtLineAvgTimer.numberOfDurationEvents.SetInt64(0)

	txtLineAvgTimer.totalDurationNanoSecs.SetInt64(0)

	txtLineAvgTimer.maximumTimeDuration.SetInt64(0)

	txtLineAvgTimer.minimumTimeDuration.SetInt64(0)

	txtLineAvgTimer.applyAbbreviatedReportFormat = false

	return
}

//	equal
//
//	Receives pointers to two TextLineSpecAverageTime
//	instances and proceeds to compare the member data
//	elements to determine whether they are equal.
//
//	If the data elements of both input parameters
//	'txtLineAvgTimerOne' and 'txtLineAvgTimerTwo' are
//	verified as equal in all respects, this method
//	returns a boolean value of 'true'. Otherwise, this
//	method returns 'false'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineAvgTimerOne			*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal member data variables contained in this
//		instance will be compared to those contained in a
//		second TextLineSpecAverageTime instance passed as
//		input parameter, 'txtLineAvgTimerOne'. If all
//		corresponding member data variables are equivalent,
//		this method returns 'true'.
//
//	txtLineAvgTimerTwo			*TextLineSpecAverageTime
//
//		A pointer to a second instance of
//		TextLineSpecAverageTime. The internal member data
//		variables contained in this instance will be
//		compared to those contained in input parameter,
//		'txtLineAvgTimerOne'. If all corresponding member
//		data variables are equivalent, this method returns
//		'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all member variable data values in
//		TextLineSpecAverageTime instances 'txtLineAvgTimerOne'
//		and 'txtLineAvgTimerTwo' are determined to be equal,
//		this method returns 'true'.
//
//		Otherwise, a value of 'false' is returned.
func (txtLineAvgTimeElectron *textLineSpecAverageTimeElectron) equal(
	txtLineAvgTimerOne *TextLineSpecAverageTime,
	txtLineAvgTimerTwo *TextLineSpecAverageTime) bool {

	if txtLineAvgTimeElectron.lock == nil {
		txtLineAvgTimeElectron.lock = new(sync.Mutex)
	}

	txtLineAvgTimeElectron.lock.Lock()

	defer txtLineAvgTimeElectron.lock.Unlock()

	if txtLineAvgTimerOne == nil {

		return false
	}

	if txtLineAvgTimerTwo == nil {

		return false
	}

	if txtLineAvgTimerOne.numberOfDurationEvents.Cmp(
		&txtLineAvgTimerTwo.numberOfDurationEvents) != 0 {

		return false
	}

	if txtLineAvgTimerOne.totalDurationNanoSecs.Cmp(
		&txtLineAvgTimerTwo.totalDurationNanoSecs) != 0 {

		return false
	}

	if txtLineAvgTimerOne.maximumTimeDuration.Cmp(
		&txtLineAvgTimerTwo.maximumTimeDuration) != 0 {

		return false
	}

	if txtLineAvgTimerOne.minimumTimeDuration.Cmp(
		&txtLineAvgTimerTwo.minimumTimeDuration) != 0 {

		return false
	}

	if txtLineAvgTimerOne.applyAbbreviatedReportFormat !=
		txtLineAvgTimerTwo.applyAbbreviatedReportFormat {

		return false
	}

	return true
}

//	testValidityOfTxtLineAvgTimer
//
//	Receives a pointer to an instance of
//	TextLineSpecAverageTime and performs a diagnostic
//	analysis to determine if that instance is valid in
//	all respects.
//
//	If the input parameter 'txtLineAvgTimer' is
//	determined to be invalid, this method will return
//	a boolean flag ('isValid') of 'false'. In addition,
//	an instance of type error ('err') will be returned
//	configured with an appropriate error message.
//
//	If the input parameter 'txtLineAvgTimer' is valid,
//	this method will return a boolean flag ('isValid') of
//	'true' and the returned error type ('err') will be
//	set to 'nil'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal member variable data values for this
//		instance will be analyzed to determine if they
//		are valid in all respects.
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
//	isValid						bool
//
//		If input parameter 'txtLineAvgTimer' is judged to
//		be valid in all respects, this return parameter
//		will be set to 'true'.
//
//		If input parameter 'txtLineAvgTimer' is found to
//		be invalid, this return parameter will be set to
//		'false'.
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
func (txtLineAvgTimeElectron *textLineSpecAverageTimeElectron) testValidityOfTxtLineAvgTimer(
	txtLineAvgTimer *TextLineSpecAverageTime,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLineAvgTimeElectron.lock == nil {
		txtLineAvgTimeElectron.lock = new(sync.Mutex)
	}

	txtLineAvgTimeElectron.lock.Lock()

	defer txtLineAvgTimeElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeElectron."+
			"testValidityOfTxtLineAvgTimer()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	bigIntZero := big.NewInt(0)

	if txtLineAvgTimer.totalDurationNanoSecs.Cmp(bigIntZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: TextLineSpecAverageTime instance is invalid!\n"+
			"'txtLineAvgTimer.totalDurationNanoSecs' is a negative value\n"+
			"txtLineAvgTimer.totalDurationNanoSecs = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.totalDurationNanoSecs.Text(10))

		return isValid, err
	}

	if txtLineAvgTimer.numberOfDurationEvents.Cmp(bigIntZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: TextLineSpecAverageTime instance is invalid!\n"+
			"'txtLineAvgTimer.numberOfDurationEvents' is a negative value\n"+
			"txtLineAvgTimer.numberOfDurationEvents = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.numberOfDurationEvents.Text(10))

		return isValid, err
	}

	if txtLineAvgTimer.maximumTimeDuration.Cmp(bigIntZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: TextLineSpecAverageTime instance is invalid!\n"+
			"'txtLineAvgTimer.maximumTimeDuration' is a negative value\n"+
			"txtLineAvgTimer.maximumTimeDuration = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.maximumTimeDuration.Text(10))

		return isValid, err
	}

	if txtLineAvgTimer.minimumTimeDuration.Cmp(bigIntZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: TextLineSpecAverageTime instance is invalid!\n"+
			"'txtLineAvgTimer.minimumTimeDuration' is a negative value\n"+
			"txtLineAvgTimer.minimumTimeDuration = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.minimumTimeDuration.Text(10))

		return isValid, err
	}

	maxInt64 := big.NewInt(0).SetInt64(math.MaxInt64)

	cmpValue := txtLineAvgTimer.minimumTimeDuration.Cmp(maxInt64)

	if cmpValue == 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Minimum Time Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"txtLineAvgTimer.minimumTimeDuration is invalid.\n"+
			"txtLineAvgTimer.minimumTimeDuration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.minimumTimeDuration.Text(10))

		return isValid, err
	}

	cmpValue = txtLineAvgTimer.maximumTimeDuration.Cmp(maxInt64)

	if cmpValue == 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Maximum Time Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"txtLineAvgTimer.maximumTimeDuration is invalid.\n"+
			"txtLineAvgTimer.maximumTimeDuration = %v\n",
			ePrefix.String(),
			txtLineAvgTimer.maximumTimeDuration.Text(10))

		return isValid, err
	}

	if txtLineAvgTimer.numberOfDurationEvents.Cmp(bigIntZero) == 0 &&
		txtLineAvgTimer.totalDurationNanoSecs.Cmp(bigIntZero) != 0 {

		err = fmt.Errorf("%v\n"+
			"Error: 'numberOfDurationEvents' is zero, but\n"+
			"'totalDurationNanoSecs' is NOT zero!\n"+
			"numberOfDurationEvents = '%v'\n"+
			" totalDurationNanoSecs = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.numberOfDurationEvents.Text(10),
			txtLineAvgTimer.totalDurationNanoSecs.Text(10))

		return isValid, err
	}

	if txtLineAvgTimer.totalDurationNanoSecs.Cmp(bigIntZero) == 0 &&
		txtLineAvgTimer.numberOfDurationEvents.Cmp(bigIntZero) != 0 {

		err = fmt.Errorf("%v\n"+
			"Error: 'totalDurationNanoSecs' is zero, but\n"+
			"'numberOfDurationEvents' is NOT zero!\n"+
			" totalDurationNanoSecs = '%v'\n"+
			"numberOfDurationEvents = '%v'\n",
			ePrefix.String(),
			txtLineAvgTimer.totalDurationNanoSecs.Text(10),
			txtLineAvgTimer.numberOfDurationEvents.Text(10))

		return isValid, err
	}

	isValid = true

	return isValid, err
}
