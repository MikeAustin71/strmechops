package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

// DateTimeHelper
//
// Provides helper methods for Type DateTimeHelper.
type dateTimeHelperAtom struct {
	lock *sync.Mutex
}

//	allocateInt64TimeDuration
//
//	Receives a time duration in the form of an int64
//	value and breaks down that total duration by days,
//	hours, minutes, seconds, milliseconds, microseconds
//	and nanoseconds.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	totalTimeDuration			int64
//
//		The total time duration to be allocated as days,
//		hours, minutes, seconds, milliseconds,
//		microseconds and nanoseconds.
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
//	allocatedTimeDuration	TimeDurationDto
//
//
//		If this method completes successfully, an
//		instance of TimeDurationDto will be returned
//		containing the allocated time broken down by
//		days, hours, minutes, seconds, milliseconds,
//		microseconds and nanoseconds.
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
func (dateTimeHelpAtom *dateTimeHelperAtom) allocateInt64TimeDuration(
	totalTimeDuration int64,
	errPrefDto *ePref.ErrPrefixDto) (
	allocatedTimeDuration TimeDurationDto,
	err error) {

	if dateTimeHelpAtom.lock == nil {
		dateTimeHelpAtom.lock = new(sync.Mutex)
	}

	dateTimeHelpAtom.lock.Lock()

	defer dateTimeHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dateTimeHelperAtom."+
			"allocateInt64TimeDuration()",
		"")

	if err != nil {
		return allocatedTimeDuration, err
	}

	if totalTimeDuration < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'totalTimeDuration' is invalid!\n"+
			"The value of 'totalTimeDuration' is less than zero.\n"+
			"totalTimeDuration = %v\n",
			ePrefix.String(),
			totalTimeDuration)

		return allocatedTimeDuration, err
	}

	if totalTimeDuration == 0 {

		return allocatedTimeDuration, err
	}

	// microsecondNanoseconds - Number of Nanoseconds in a Microsecond.
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	microsecondNanoseconds := int64(time.Microsecond)

	// millisecondNanoseconds - Number of Nanoseconds in a MilliSecond.
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	millisecondNanoseconds := int64(time.Millisecond)

	// secondNanoseconds - Number of Nanoseconds in a Second
	secondNanoseconds := int64(time.Second)

	// minuteNanoseconds - Number of Nanoseconds in a minute
	minuteNanoseconds := int64(time.Minute)

	// hourNanoseconds - Number of Nanoseconds in an hour
	hourNanoseconds := int64(time.Hour)

	// dayNanoseconds - Number of Nanoseconds in a standard
	// 24-hour day.
	dayNanoseconds := int64(time.Hour) * int64(24)

	allocatedTimeDuration.TotalNanoseconds = totalTimeDuration

	if totalTimeDuration >= dayNanoseconds {

		allocatedTimeDuration.NumberOfDays =
			totalTimeDuration / dayNanoseconds

		totalTimeDuration =
			totalTimeDuration -
				(allocatedTimeDuration.NumberOfDays * dayNanoseconds)
	}

	if totalTimeDuration >= hourNanoseconds {

		allocatedTimeDuration.NumberOfHours =
			totalTimeDuration / hourNanoseconds

		totalTimeDuration =
			totalTimeDuration -
				(allocatedTimeDuration.NumberOfHours * hourNanoseconds)
	}

	if totalTimeDuration >= minuteNanoseconds {

		allocatedTimeDuration.NumberOfMinutes =
			totalTimeDuration / minuteNanoseconds

		totalTimeDuration =
			totalTimeDuration -
				(allocatedTimeDuration.NumberOfMinutes * minuteNanoseconds)
	}

	if totalTimeDuration >= secondNanoseconds {

		allocatedTimeDuration.NumberOfSeconds =
			totalTimeDuration / secondNanoseconds

		totalTimeDuration = totalTimeDuration -
			(allocatedTimeDuration.NumberOfSeconds * secondNanoseconds)
	}

	if totalTimeDuration >= millisecondNanoseconds {

		allocatedTimeDuration.NumberOfMilliseconds =
			totalTimeDuration / millisecondNanoseconds

		totalTimeDuration =
			totalTimeDuration -
				(allocatedTimeDuration.NumberOfMilliseconds * millisecondNanoseconds)
	}

	if totalTimeDuration >= microsecondNanoseconds {
		allocatedTimeDuration.NumberOfMicroseconds =
			totalTimeDuration / microsecondNanoseconds

		totalTimeDuration =
			totalTimeDuration -
				(allocatedTimeDuration.NumberOfMicroseconds * microsecondNanoseconds)
	}

	allocatedTimeDuration.NumberOfNanoseconds = totalTimeDuration

	return allocatedTimeDuration, err
}
