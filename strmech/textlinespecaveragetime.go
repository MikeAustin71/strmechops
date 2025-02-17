package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"
)

// TextLineSpecAverageTime
//
// This type is primarily used for timing code
// execution. It is designed to compute and record
// a large number of code executions and produce
// a text report detailing the average duration
// of all executions.
//
// While the primary usage is envisioned as computing
// average duration for code executions, this type
// may be used to compute average time for any series
// of events.
//
// Two types or text reports are provided. One type
// presents the average time duration along with the
// maximum and minimum time durations in the timing
// event series.
type TextLineSpecAverageTime struct {
	numberOfDurationEvents       big.Int
	totalDurationNanoSecs        big.Int
	maximumTimeDuration          big.Int
	minimumTimeDuration          big.Int
	applyAbbreviatedReportFormat bool
	textLineReader               *strings.Reader
	lock                         *sync.Mutex
}

//	AddDurationEvent
//
//	Adds a time duration value to the internal average
//	time counters maintained by the current instance of
//	TextLineSpecAverageTime.
//
//	This time duration data is recorded for purposes of
//	later generating an average time duration report for
//	a series of timing events.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	eventDuration				time.Duration
//
//		The time duration associated with an event. This
//		duration value will be added to the internal
//		average time counters maintained by the
//		current instance of TextLineSpecAverageTime.
//
//		This time duration data is recorded for purposes
//		of later generating an average time duration
//		report for a series of timing events.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (txtLineAvgTime *TextLineSpecAverageTime) AddDurationEvent(
	eventDuration time.Duration,
	errorPrefix interface{}) error {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime."+
			"AddDurationEvent()",
		"")

	if err != nil {

		return err
	}

	return new(textLineSpecAverageTimeMechanics).
		addDurationEvent(
			txtLineAvgTime,
			eventDuration,
			ePrefix.XCpy(
				"txtLineAvgTime"))
}

//	AddStartStopEvent
//
//	Receives an event start time and event end time. This
//	method will then proceed to compute a time duration
//	from these input time values and record the time
//	duration in the internal average time counters
//	maintained by the current instance of
//	TextLineSpecAverageTime.
//
//	This time duration data is recorded for purposes of
//	later generating an average time duration report for
//	a series of timing events.
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
//
//		The time at which a specific event started.
//
//	endTime						time.Time
//
//		The time at which a specific event ended.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (txtLineAvgTime *TextLineSpecAverageTime) AddStartStopEvent(
	startTime time.Time,
	endTime time.Time,
	errorPrefix interface{}) error {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime."+
			"AddStartStopEvent()",
		"")

	if err != nil {

		return err
	}

	return new(textLineSpecAverageTimeMechanics).
		addStartStopEvent(
			txtLineAvgTime,
			startTime,
			endTime,
			ePrefix.XCpy(
				"txtLineAvgTime"))
}

//	CalcAvgTimeDuration
//
//	Calculates the average duration of a timer event
//	series recorded in the current instance of
//	TextLineSpecAverageTime.
//
//	Average time duration is calculated by dividing the
//	total time duration by the number of separate timing
//	events.
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
//		for all recorded timing events in the current
//		instance of TextLineSpecAverageTime. Average
//		duration is calculated by dividing the total
//		time duration by the number of separate timing
//		events.
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
//	numberOfTimingEvents		int64
//
//		The number of timing events recorded in the
//		current instance of TextLineSpecAverageTime.
//
//		The number of timing events is used to
//		calculate average time duration for the time
//		series.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CalcAvgTimeDuration(
	errorPrefix interface{}) (
	avgDuration int64,
	maximumTimeDuration int64,
	minimumTimeDuration int64,
	numberOfTimingEvents int64,
	err error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime."+
			"CalcAvgTimeDuration()",
		"")

	if err != nil {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			numberOfTimingEvents,
			err
	}

	avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err = new(textLineSpecAverageTimeAtom).
		calcAverageDuration(
			txtLineAvgTime,
			ePrefix.XCpy(
				"txtLineAvgTime"))

	if err != nil {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			numberOfTimingEvents,
			err
	}

	maxInt64Value := big.NewInt(
		math.MaxInt64)

	comparison := txtLineAvgTime.
		numberOfDurationEvents.Cmp(maxInt64Value)

	if comparison == 1 {
		err = fmt.Errorf("%v\n"+
			"Error: The number of timing events exceeds\n"+
			"the maximum value for a 64-bit integer (int64).\n"+
			"Number Of Timing Events = %v\n",
			ePrefix.String(),
			txtLineAvgTime.numberOfDurationEvents.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			numberOfTimingEvents,
			err
	}

	numberOfTimingEvents =
		txtLineAvgTime.numberOfDurationEvents.Int64()

	return avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		numberOfTimingEvents,
		err
}

//	CalcAvgTimeDurationDetail
//
//	Calculates and returns a detailed breakdown of the
//	average duration, maximum duration and minimum
//	duration for a timer event series recorded in the
//	current instance of TextLineSpecAverageTime.
//
//	Average time duration is calculated by dividing the
//	total time duration by the number of separate timing
//	events.
//
//	The data returned by this method breaks down
//	average time duration, maximum time duration and
//	minimum time duration by days, hours, minutes,
//	seconds, milliseconds, microseconds and nanoseconds.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The detail data for average time duration, maximum
//	time duration and minimum time duration is returned
//	as a type TimeDurationDto. This structure provides
//	a breakdown of time duration data by days, hours,
//	minutes, seconds, milliseconds, microseconds and
//	nanoseconds.
//
//		type TimeDurationDto struct {
//
//			TotalNanoseconds int64
//				The total number of Nanoseconds to
//				be allocated.
//
//			NumberOfDays int64
//				The number Days represented by
//				'TotalNanoseconds'.
//
//			NumberOfHours int64
//				The number Hours represented by
//				'TotalNanoseconds'.
//
//			NumberOfMinutes int64
//				The number Minutes represented by
//				'TotalNanoseconds'.
//
//			NumberOfSeconds int64
//				The number Seconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMilliseconds int64
//				The number Milliseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMicroseconds int64
//				The number Microseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfNanoseconds int64
//				The number Nanoseconds remaining after
//				the allocation of Microseconds.
//		}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	allocatedAvgDuration		TimeDurationDto
//
//		If this method completes successfully, this
//		parameter will return the average time duration
//		for all recorded timing events in the current
//		instance of TextLineSpecAverageTime. Average
//		duration is calculated by dividing the total
//		time duration by the number of separate timing
//		events.
//
//		Type TimeDurationDto will present average time
//		duration by days, hours, minutes, seconds,
//		milliseconds, microseconds and nanoseconds.
//
//	allocatedMaxDuration		TimeDurationDto
//
//		If this method completes successfully, this
//		parameter will return the maximum time duration
//		for all recorded timing events in the current
//		instance of TextLineSpecAverageTime.
//
//		Type TimeDurationDto will present maximum time
//		duration by days, hours, minutes, seconds,
//		milliseconds, microseconds and nanoseconds.
//
//	allocatedMinDuration		TimeDurationDto
//
//		If this method completes successfully, this
//		parameter will return the minimum time duration
//		for all recorded timing events in the current
//		instance of TextLineSpecAverageTime.
//
//		Type TimeDurationDto will present minimum time
//		duration by days, hours, minutes, seconds,
//		milliseconds, microseconds and nanoseconds.
//
//	numberOfTimingEvents		int64
//
//		The number of timing events recorded in the
//		current instance of TextLineSpecAverageTime.
//
//		The number of timing events is used to
//		calculate average time duration for the time
//		series.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CalcAvgTimeDurationDetail(
	errorPrefix interface{}) (
	allocatedAvgDuration TimeDurationDto,
	allocatedMaxDuration TimeDurationDto,
	allocatedMinDuration TimeDurationDto,
	numberOfTimingEvents int64,
	err error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime."+
			"CalcAvgTimeDurationDetail()",
		"")

	if err != nil {

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	var avgDuration, maximumTimeDuration,
		minimumTimeDuration int64

	avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err = new(textLineSpecAverageTimeAtom).
		calcAverageDuration(
			txtLineAvgTime,
			ePrefix.XCpy(
				"txtLineAvgTime"))

	if err != nil {

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	dateTimeHelper := new(DateTimeHelper)

	allocatedAvgDuration,
		err = dateTimeHelper.AllocateTimeDuration(
		avgDuration,
		ePrefix.XCpy(
			"allocatedAvgDuration<-avgDuration"))

	if err != nil {

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	allocatedMaxDuration,
		err = dateTimeHelper.AllocateTimeDuration(
		maximumTimeDuration,
		ePrefix.XCpy(
			"allocatedMaxDuration<-maximumTimeDuration"))

	if err != nil {

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	allocatedMinDuration,
		err = dateTimeHelper.AllocateTimeDuration(
		minimumTimeDuration,
		ePrefix.XCpy(
			"allocatedMinDuration<-minimumTimeDuration"))

	if err != nil {

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	maxInt64Value := big.NewInt(
		math.MaxInt64)

	comparison := txtLineAvgTime.
		numberOfDurationEvents.Cmp(maxInt64Value)

	if comparison == 1 {
		err = fmt.Errorf("%v\n"+
			"Error: The number of timing events exceeds\n"+
			"the maximum value for a 64-bit integer (int64).\n"+
			"Number Of Timing Events = %v\n",
			ePrefix.String(),
			txtLineAvgTime.numberOfDurationEvents.Text(10))

		return allocatedAvgDuration,
			allocatedMaxDuration,
			allocatedMinDuration,
			numberOfTimingEvents,
			err
	}

	numberOfTimingEvents =
		txtLineAvgTime.numberOfDurationEvents.Int64()

	return allocatedAvgDuration,
		allocatedMaxDuration,
		allocatedMinDuration,
		numberOfTimingEvents,
		err
}

//	CopyIn
//
//	Copies all the data fields from an incoming instance
//	of TextLineSpecAverageTime ('incomingAvgTimer') to
//	the data fields of the current TextLineSpecAverageTime
//	instance ('txtLineAvgTime').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the current instance of
//	TextLineSpecAverageTime.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	incomingAvgTimer			*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		All data value contained in this instance will be
//		copied to current instance of
//		TextLineSpecAverageTime, 'txtLineAvgTime'.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CopyIn(
	incomingAvgTimer *TextLineSpecAverageTime,
	errorPrefix interface{}) error {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLineSpecAverageTimeNanobot).copy(
		txtLineAvgTime,
		incomingAvgTimer,
		ePrefix.XCpy("txtLineAvgTime<-"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	TextLineSpecAverageTime instance.
//
//	If the current TextLineSpecAverageTime instance
//	contains invalid member variable data values, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextLineSpecAverageTime
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the current
//		TextLineSpecAverageTime instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecAverageTime,
	error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deepCopyTxtLineAvgTimer := TextLineSpecAverageTime{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.CopyOut()",
		"")

	if err != nil {
		return deepCopyTxtLineAvgTimer, err
	}

	err = new(textLineSpecAverageTimeNanobot).copy(
		&deepCopyTxtLineAvgTimer,
		txtLineAvgTime,
		ePrefix.XCpy("deepCopyTxtLineAvgTimer<-"+
			"txtLineAvgTime"))

	return deepCopyTxtLineAvgTimer, err
}

//	CopyOutITextLine
//
//	Returns a deep copy of the current
//	TextLineSpecAverageTime instance cast as a type
//	ITextLineSpecification.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ITextLineSpecification
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextLineSpecAverageTime instance cast as an
//		ITextLineSpecification object.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deepCopyTxtLineAvgTimer := TextLineSpecAverageTime{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.CopyOutITextLine()",
		"")

	if err != nil {
		return ITextLineSpecification(&deepCopyTxtLineAvgTimer),
			err
	}

	err = new(textLineSpecAverageTimeNanobot).copy(
		&deepCopyTxtLineAvgTimer,
		txtLineAvgTime,
		ePrefix.XCpy("deepCopyTxtLineAvgTimer<-"+
			"txtLineAvgTime"))

	return ITextLineSpecification(&deepCopyTxtLineAvgTimer),
		err
}

//	CopyOutPtr
//
//	Returns a pointer to a deep copy of the current
//	TextLineSpecAverageTime instance.
//
//	If the current TextLineSpecAverageTime instance
//	contains invalid member variable data values, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*TextLineSpecAverageTime
//
//		If this method completes successfully, this
//		parameter will return a pointer to a deep copy of
//		the current TextLineSpecAverageTime instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecAverageTime,
	error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deepCopyTxtLineAvgTimer := TextLineSpecAverageTime{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.CopyOutPtr()",
		"")

	if err != nil {
		return &deepCopyTxtLineAvgTimer, err
	}

	err = new(textLineSpecAverageTimeNanobot).copy(
		&deepCopyTxtLineAvgTimer,
		txtLineAvgTime,
		ePrefix.XCpy("deepCopyTxtLineAvgTimer<-"+
			"txtLineAvgTime"))

	return &deepCopyTxtLineAvgTimer, err
}

//	Empty
//
//	Resets all internal member variables to their initial
//	or zero states.
//
//	Effectively, this method resets all the internal
//	event timer counters to zero.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineAvgTime *TextLineSpecAverageTime) Empty() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	new(textLineSpecAverageTimeElectron).empty(
		txtLineAvgTime)

	txtLineAvgTime.lock.Unlock()

	txtLineAvgTime.lock = nil

	return
}

//	Equal
//
//	Receives a pointer to another instance of
//	TextLineSpecAverageTime and proceeds to compare the
//	member variables to those of the current
//	TextLineSpecAverageTime instance in order to
//	determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables of both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingAvgTimer			*TextLineSpecAverageTime
//
//		A pointer to an incoming instance of
//		TextLineSpecAverageTime. This method will compare
//		all member variable data values in this instance
//		against those contained in the current instance
//		of TextLineSpecAverageTime. If the data values in
//		both instances are found to be equal in all
//		respects, this method will return a boolean value
//		of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingAvgTimer' are equal
//		in all respects to those contained in the current
//		instance of TextLineSpecAverageTime, this method
//		will return a boolean value of 'true'. Otherwise,
//		a value of 'false' will be returned.
func (txtLineAvgTime *TextLineSpecAverageTime) Equal(
	incomingAvgTimer *TextLineSpecAverageTime) bool {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	return new(textLineSpecAverageTimeElectron).equal(
		txtLineAvgTime,
		incomingAvgTimer)
}

// EqualITextLine
//
// Receives an object implementing the
// ITextLineSpecification interface and proceeds to
// compare the member variables to those of the current
// TextLineSpecAverageTime instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison
// is returned. If the member variables from both
// instances are equal in all respects, this flag is set
// to 'true'. Otherwise, this method returns 'false'.
//
// This method is required by interface
// ITextLineSpecification.
func (txtLineAvgTime *TextLineSpecAverageTime) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	avgTimer, ok := iTextLine.(*TextLineSpecAverageTime)

	if !ok {
		return false
	}

	return new(textLineSpecAverageTimeElectron).equal(
		txtLineAvgTime,
		avgTimer)
}

// GetDefaultReportFormat
//
// Returns a boolean value specifying the default report
// format type.
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// If the returned boolean value is set to 'true', it
// signals that all reporting methods will return timing
// reports using the Full Report Format.
//
// If the returned boolean value is set to 'false', it
// signals that all reporting methods will return timing
// reports using the Abbreviated Report Format.
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format".
//
// Users have the option to control the report format
// applied by the current instance of
// TextLineSpecAverageTime. The default report format
// is controlled by calling the following methods:
//
//	TextLineSpecAverageTime.SetFullReportFormat()
//	TextLineSpecAverageTime.SetAbbreviatedReportFormat()
//
// When the instances of TextLineSpecAverageTime are first
// created, the default report format is set to the Full
// Report Format.
//
// To generate a text report of timing data, call one of
// the following methods:
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isFullReportFormat			bool
//
//		If the returned boolean value is set to 'true',
//		it signals that all reporting methods will return
//		timing reports using the Full Report Format.
//
//		If the returned boolean value is set to 'false',
//		it signals that all reporting methods will return
//		timing reports using the Abbreviated Report Format.
//
//		TextLineSpecAverageTime produces two types reports.
//
//		A "Full" report presents timing data on average time
//		duration, minimum time duration and maximum time
//		duration.
//
//		An "Abbreviated" report presents timing data only on
//		average duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
func (txtLineAvgTime *TextLineSpecAverageTime) GetDefaultReportFormat() (isFullReportFormat bool) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	return !txtLineAvgTime.applyAbbreviatedReportFormat
}

// GetFormattedText
//
// Configures the lines of text produced by this instance
// of TextLineSpecAverageTime and returns them as a
// single string.
//
// This method will compile the calculated timing data
// as formatted text for screen displays, file output or
// printing.
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// The type of report prepared by this method will depend
// on the current default report format setting.
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format".
//
// Users have the option to control the report format
// applied by the current instance of
// TextLineSpecAverageTime. The default report format
// is controlled by calling the following methods:
//
//	TextLineSpecAverageTime.SetFullReportFormat()
//	TextLineSpecAverageTime.SetAbbreviatedReportFormat()
//
// If the text report format produced by this method is
// unsuitable, the user may acquire the raw numerical
// data and create a custom report format. This
// statistical timing data can be acquired by calling
// the following method:
//
//	TextLineSpecAverageTime.CalculateAvgDuration()
//
// This method fulfills requirements of interface
// ITextLineSpecification.
//
// Methods which return formatted text are listed as
// follows:
//
//	TextLineSpecAverageTime.String()
//	TextLineSpecAverageTime.GetFormattedText()
//	TextLineSpecAverageTime.TextBuilder()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The report format generated by this method is
//	controlled by the current default report format
//	setting.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The formatted text line output generated by the
//		current instance of TextLineSpecAverageTime.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	strBuilder := strings.Builder{}

	err = new(textLineSpecAverageTimeMechanics).
		getFormattedText(
			txtLineAvgTime,
			&strBuilder,
			txtLineAvgTime.applyAbbreviatedReportFormat,
			ePrefix.XCpy(
				"strBuilder<-txtLineAvgTime"))

	if err != nil {
		return "", err
	}

	return strBuilder.String(), err
}

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecAverageTime
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this method
//	will return a boolean value of 'false'.
//
//	If all data elements are determined to be valid, this
//	method returns a boolean value of 'true'.
//
//	This method is functionally equivalent to method:
//
//		TextLineSpecAverageTime.IsValidInstanceError()
//
//	The difference between that method and this is that
//	this method takes no input parameters and returns a
//	boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal member data variables
//		contained in the current instance of
//		TextLineSpecAverageTime are found to be invalid,
//		this method will return a boolean value of
//		'false'.
//
//		Conversely, if all internal member data variables
//		contained in the current instance of
//		TextLineSpecAverageTime are found to be valid,
//		this method returns a boolean value of 'true'.
func (txtLineAvgTime *TextLineSpecAverageTime) IsValidInstance() bool {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	isValid,
		_ := new(textLineSpecAverageTimeElectron).
		testValidityOfTxtLineAvgTimer(
			txtLineAvgTime,
			nil)

	return isValid
}

//	IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current instance of
//	TextLineSpecAverageTime to determine if they are
//	valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error containing an appropriate
//	error message.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If any of the internal member data variables
//		contained in the current instance of
//		TextLineSpecAverageTime are found to be invalid,
//		this method will return an error containing an
//		appropriate error message.
//
//		If an error message is returned, the text value
//		of input parameter 'errorPrefix' (error prefix)
//		will be inserted or prefixed at the beginning of
//		the error message.
func (txtLineAvgTime *TextLineSpecAverageTime) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineSpecAverageTimeElectron).
		testValidityOfTxtLineAvgTimer(
			txtLineAvgTime,
			ePrefix.XCpy(
				"txtLineAvgTime"))

	return err
}

//	New
//
//	Returns an initialized instance of
//	TextLineSpecAverageTime ready to receive and process
//	a series event durations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	To properly utilize an instance of
//	TextLineSpecAverageTime, it should be created with one
//	of the 'New' methods.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextLineSpecAverageTime
//
//		If this method completes successfully, an
//		initialized instance of TextLineSpecAverageTime
//		will be returned.
//
//		This new instance will be ready in all respects
//		to receive and process event durations.
func (txtLineAvgTime *TextLineSpecAverageTime) New() TextLineSpecAverageTime {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	newAvgTimer := TextLineSpecAverageTime{}

	newAvgTimer.numberOfDurationEvents.SetInt64(0)
	newAvgTimer.totalDurationNanoSecs.SetInt64(0)

	return newAvgTimer
}

//	Read
//
//	Implements the io.Reader interface for type
//	TextLineSpecAverageTime.
//
//	The formatted text line string generated by the
//	current instance of TextLineSpecAverageTime will be
//	written to the byte buffer 'p'. If the length of 'p'
//	is less than the length of the formatted text line
//	string, multiple calls to this method will write the
//	remaining unread characters to the byte buffer 'p'.
//
//	Read() supports buffered 'read' operations.
//
//	This method reads up to len(p) bytes into p. It
//	returns the number of bytes read (0 <= n <= len(p))
//	and any error encountered. Even if read returns
//	n < len(p), it may use all of p as scratch space
//	during the call.
//
//	If some data is available but not len(p) bytes,
//	readBytes() conventionally returns what is available
//	instead of waiting for more.
//
//	When this method encounters an error or end-of-file
//	condition after successfully reading n > 0 bytes, it
//	returns the number of bytes read. It may return the
//	(non-nil) error from the same call or return the
//	error (and n == 0) from a subsequent call.
//
//	An instance of this general case is that a Reader
//	returning a non-zero number of bytes at the end of
//	the input stream may return either 'err == EOF' or
//	'err == nil'. The next read operation should return
//	'0, EOF'.
//
//	Callers should always process the n > 0 bytes
//	returned before considering the error 'err'. Doing so
//	correctly handles I/O errors that happen after reading
//	some bytes and also both of the allowed EOF behaviors.
//
//	The last read operation performed on the formatted
//	text string will always return 'n==0' and
//	'err==io.EOF'.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p							[]byte
//
//		The byte buffer into which the formatted text line string
//		generated by the current TextLineSpecAverageTime instance
//		will be written.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	n							int
//
//		The number of bytes written to byte buffer 'p'.
//
//		Read() reads up to len(p) bytes into p. It returns
//		the number of bytes read (0 <= n <= len(p)) and any error
//		encountered. Even if Read() returns n < len(p), it may use
//		all of 'p' as scratch space during the call. If some
//		data is available but not len(p) bytes, Read()
//		conventionally returns what is available instead of
//		waiting for more.
//
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will
//		encapsulate an error message.
//
//		When Read() encounters an error or end-of-file condition
//		after successfully reading n > 0 bytes, it returns the
//		number of bytes read. It may return the (non-nil) error
//		from the same call or return the error (and n == 0) from
//		a subsequent call. An instance of this general case is
//		that a Reader returning a non-zero number of bytes at the
//		end of the input stream may return either err == EOF or
//		err == nil. The next read operation should return 0, EOF.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example # 1
//
//		p := make([]byte, 50)
//
//		var n, readBytesCnt int
//		sb := strings.Builder{}
//
//		for {
//
//		  n,
//		  err = avgTimer01.Read(p)
//
//		  if n == 0 {
//		    break
//		  }
//
//		  sb.Write(p[:n])
//		  readBytesCnt += n
//		}
//
//		if err != nil &&
//		  err != io.EOF {
//		   return fmt.Errorf(
//		    "Error Returned From avgTimer01.Read(p)\n"+
//		    "Error = \n%v\n",
//		     err.Error())
//		}
//
//		fmt.Printf("Text Line String: %s\n",
//		              sb.String())
//
//		fmt.Printf("Number of bytes Read: %v\n",
//		              readBytesCnt)
//
//	Example # 2
//
//		p := make([]byte, 50)
//
//		var n, readBytesCnt int
//		var actualStr string
//
//		for {
//
//		  n,
//		  err = avgTimer01.Read(p)
//
//		  if n == 0 {
//		    break
//		  }
//
//		  actualStr += string(p[:n])
//		  readBytesCnt += n
//		}
//
//		if err != nil &&
//		  err != io.EOF {
//		   return fmt.Errorf(
//		    "Error Returned From avgTimer01.Read(p)\n"+
//		    "Error = \n%v\n",
//		     err.Error())
//		}
//
//		fmt.Printf("Text Line String: %v\n",
//		              actualStr)
//
//		fmt.Printf("Number of bytes Read: %v\n",
//		              readBytesCnt)
func (txtLineAvgTime *TextLineSpecAverageTime) Read(
	p []byte) (
	n int,
	err error) {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecBlankLines.Read()",
		"")

	if txtLineAvgTime.textLineReader == nil {

		strBuilder := strings.Builder{}

		err = new(textLineSpecAverageTimeMechanics).
			getFormattedText(
				txtLineAvgTime,
				&strBuilder,
				false,
				ePrefix.XCpy(
					"strBuilder<-txtLineAvgTime"))

		if err != nil {
			return n, err
		}

		txtLineAvgTime.textLineReader =
			strings.NewReader(strBuilder.String())

		if txtLineAvgTime.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(strBuilder.String())\n"+
				"returned a nil pointer.\n"+
				"txtLineAvgTime.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = new(textSpecificationAtom).
		readBytes(
			txtLineAvgTime.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtLineAvgTime.textLineReader"))

	if err == io.EOF {

		txtLineAvgTime.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize
//
// This method will reset the internal member variable
// 'TextLineSpecAverageTime.textLineReader' to its
// initial zero state of 'nil'.
//
// This method is rarely used. It provides a means of
// reinitializing the internal strings.Reader in case an
// error occurs during a read operation initiated by
// method TextLineSpecAverageTime.Read().
//
// Calling this method cleans up the residue from an
// aborted read operation and allows the calling
// function to start a new read operation.
//
// If any errors are returned by method
// TextLineSpecAverageTime.Read() which are NOT equal to
// io.EOF, call this method,
// TextLineSpecAverageTime.ReaderInitialize(), to reset
// the internal reader for future read operations.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (txtLineAvgTime *TextLineSpecAverageTime) ReaderInitialize() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	txtLineAvgTime.textLineReader = nil

	return
}

// SetAbbreviatedReportFormat
//
// Calling this method will set the default report format
// to the "Abbreviated Report Format".
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// The type of report prepared by this method will depend
// on the current default report format setting.
//
// This method will set the default report format to the
// "Abbreviated Report Format". This means that future
// calls to the following methods will yield a timing
// report presented with the Abbreviated Report Format:
//
//	TextLineSpecAverageTime.String()
//	TextLineSpecAverageTime.GetFormattedText()
//	TextLineSpecAverageTime.TextBuilder()
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format". However, users have the option to designate
// the "Abbreviated Report Format" as the default report
// format by calling this method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (txtLineAvgTime *TextLineSpecAverageTime) SetAbbreviatedReportFormat() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	txtLineAvgTime.applyAbbreviatedReportFormat = true
}

// SetFullReportFormat
//
// Calling this method will set the default report format
// to the "Full Report Format".
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// The type of report prepared by this method will depend
// on the current default report format setting.
//
// This method will set the default report format to the
// "Full Report Format". This means that future calls to
// the following methods will yield a timing report
// presented with the Full Report Format:
//
//	TextLineSpecAverageTime.String()
//	TextLineSpecAverageTime.GetFormattedText()
//	TextLineSpecAverageTime.TextBuilder()
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format". However, users may find it necessary to
// reinstate the "Full Report Format" as the default
// report format by calling this method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (txtLineAvgTime *TextLineSpecAverageTime) SetFullReportFormat() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	txtLineAvgTime.applyAbbreviatedReportFormat = false
}

//	SetInitializeTimerToZero
//
//	Reinitializes the internal timers to zero for the
//	current instance of TextLineSpecAverageTime.
//
//	After this method completes, the current instance
//	of TextLineSpecAverageTime is ready in all respects
//	to begin processing a new series of time duration
//	events.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All previously collected time duration and event data
//	will be deleted by this method.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	None
func (txtLineAvgTime *TextLineSpecAverageTime) SetInitializeTimerToZero() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	new(textLineSpecAverageTimeElectron).empty(
		txtLineAvgTime)

	return
}

// String
//
// Configures the lines of text produced by this instance
// of TextLineSpecAverageTime and returns them as a
// single string.
//
// This method will compile the calculated timing data
// as formatted text for screen displays, file output or
// printing.
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// The type of report prepared by this method will depend
// on the current default report format setting.
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format".
//
// Users have the option to control the report format
// applied by the current instance of
// TextLineSpecAverageTime. The default report format
// is controlled by calling the following methods:
//
//	TextLineSpecAverageTime.SetFullReportFormat()
//	TextLineSpecAverageTime.SetAbbreviatedReportFormat()
//
// If the text report format produced by this method is
// unsuitable, the user may acquire the raw numerical
// data and create a custom report format. This
// statistical timing data can be acquired by calling
// the following method:
//
//	TextLineSpecAverageTime.CalculateAvgDuration()
//
// This method fulfills requirements of interface
// ITextLineSpecification.
//
// Methods which return formatted text are listed as
// follows:
//
//	TextLineSpecAverageTime.GetFormattedText()
//	TextLineSpecAverageTime.TextBuilder()
//	TextLineSpecAverageTime.String()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The report format generated by this method is
//	controlled by the current default report format
//	setting.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The formatted text line output generated by the
//		current instance of TextLineSpecAverageTime.
//
//		These text lines comprise a report of the
//		recorded average duration timing data collected
//		by the current instance of
//		TextLineSpecAverageTime.
//
//		The data format used in the preparation of this
//		report is controlled by the current default
//		report format.
//
//		If an error condition is encountered, this string
//		will contain an appropriate error message. This
//		error message will contain the word 'Error'.
func (txtLineAvgTime TextLineSpecAverageTime) String() string {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var formattedText string

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecAverageTime.String()",
		"")

	strBuilder := strings.Builder{}

	err := new(textLineSpecAverageTimeMechanics).
		getFormattedText(
			&txtLineAvgTime,
			&strBuilder,
			txtLineAvgTime.applyAbbreviatedReportFormat,
			ePrefix.XCpy(
				"formattedText<-txtLineAvgTime Formatted Text"))

	if err != nil {

		formattedText = fmt.Sprintf("%v\n",
			err.Error())

		return formattedText
	}

	return formattedText
}

// TextBuilder
//
// Configures the lines of text produced by this
// instance of TextLineSpecAverageTime, and writes it to
// an instance of strings.Builder.
//
// This method will compile the calculated timing data
// as formatted text for screen displays, file output or
// printing.
//
// TextLineSpecAverageTime produces two types reports.
//
//	(1)	A "Full" report presents timing data on average
//		time duration, minimum time duration and maximum
//		time duration.
//
//	(2)	The second type of report is an "Abbreviated"
//		report which only presents timing data on average
//		duration. It does NOT include timing data on
//		minimum time duration and maximum time duration.
//
// The type of report prepared by this method will depend
// on the current default report format setting.
//
// By default, all instances of TextLineSpecAverageTime
// are initially created with a default of "Full Report
// Format".
//
// Users have the option to control the report format
// applied by the current instance of
// TextLineSpecAverageTime. The default report format
// is controlled by calling the following methods:
//
//	TextLineSpecAverageTime.SetFullReportFormat()
//	TextLineSpecAverageTime.SetAbbreviatedReportFormat()
//
// If the text report format produced by this method is
// unsuitable, the user may acquire the raw numerical
// data and create a custom report format. This
// statistical timing data can be acquired by calling
// the following method:
//
//	TextLineSpecAverageTime.CalculateAvgDuration()
//
// This method fulfills requirements of interface
// ITextLineSpecification.
//
// Methods which return formatted text are listed as
// follows:
//
//	TextLineSpecAverageTime.String()
//	TextLineSpecAverageTime.GetFormattedText()
//	TextLineSpecAverageTime.TextBuilder()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The report format generated by this method is
//	controlled by the current default report format
//	setting.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of *strings.Builder. The
//		formatted text characters produced by this method
//		will be written to this instance of
//		strings.Builder.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineAvgTime *TextLineSpecAverageTime) TextBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecAverageTime.TextBuilder()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	err = new(textLineSpecAverageTimeMechanics).
		getFormattedText(
			txtLineAvgTime,
			strBuilder,
			txtLineAvgTime.applyAbbreviatedReportFormat,
			ePrefix.XCpy(
				"strBuilder<-txtLineAvgTime Formatted Text"))

	if err != nil {
		return err
	}

	return err
}

// TextLineSpecName
//
// Returns Text Line Specification Name.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (txtLineAvgTime *TextLineSpecAverageTime) TextLineSpecName() string {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	return "AverageTime"
}

// TextTypeName
//
// Returns a string specifying the type of Text Line
// specification.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (txtLineAvgTime *TextLineSpecAverageTime) TextTypeName() string {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	return "TextLineSpecAverageTime"

}
