package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

// DateTimeHelper
//
// Provides helper methods for date time calculations.
type DateTimeHelper struct {
	lock *sync.Mutex
}

//	AllocateTimeDuration
//
//	Receives a time duration value and proceeds to break
//	down that time duration value by days, hours,
//	minutes, seconds, milliseconds, microseconds and
//	nanoseconds.
//
//	The time duration value is submitted as a type of
//	empty interface (interface{}). This means that the
//	time duration value may be submitted as one of four
//	supported types:
//
//			int64
//			*int64
//			time.Duration
//			*time.Duration
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	totalTimeDuration			interface{}
//
//		The total time duration to be allocated as days,
//		hours, minutes, seconds, milliseconds,
//		microseconds and nanoseconds.
//
//		This parameter will accept one of four types:
//
//			int64
//			*int64
//			time.Duration
//			*time.Duration
//
//		If 'totalTimeDuration' is not submitted as one
//		of the four supported types, an error will be
//		returned.
//
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
//	numOfDays					int64
//
//		The number Days represented by input parameter
//		'totalTimeDuration'.
//
//	numOfHours					int64
//
//		The number Hours represented by input parameter
//		'totalTimeDuration'.
//
//	numOfMinutes				int64
//
//		The number of Minutes represented by input
//		parameter 'totalTimeDuration'.
//
//	numOfSeconds				int64
//
//		The number of Seconds represented by input
//		parameter 'totalTimeDuration'.
//
//	numOfMilliseconds			int64
//
//		The number of Milliseconds represented by input
//		parameter 'totalTimeDuration'.
//
//	numOfMicroseconds			int64
//
//		The number Microseconds represented by input
//		parameter 'totalTimeDuration'.
//
//	numOfNanoseconds			int64
//
//		The number of remaining Nanoseconds represented
//		by input parameter 'totalTimeDuration'.
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
func (dateTimeHelper *DateTimeHelper) AllocateTimeDuration(
	totalTimeDuration interface{},
	errorPrefix interface{}) (
	numOfDays int64,
	numOfHours int64,
	numOfMinutes int64,
	numOfSeconds int64,
	numOfMilliseconds int64,
	numOfMicroseconds int64,
	numOfNanoseconds int64,
	err error) {

	if dateTimeHelper.lock == nil {
		dateTimeHelper.lock = new(sync.Mutex)
	}

	dateTimeHelper.lock.Lock()

	defer dateTimeHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DateTimeHelper."+
			"AllocateTimeDuration()",
		"")

	if err != nil {

		return numOfDays,
			numOfHours,
			numOfMinutes,
			numOfSeconds,
			numOfMilliseconds,
			numOfMicroseconds,
			numOfNanoseconds,
			err
	}

	var int64Value int64

	var ok bool

	switch totalTimeDuration.(type) {

	case int64:

		int64Value, ok = totalTimeDuration.(int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int64 cast to 'int64Value' failed!\n",
				ePrefix.String())

			return numOfDays,
				numOfHours,
				numOfMinutes,
				numOfSeconds,
				numOfMilliseconds,
				numOfMicroseconds,
				numOfNanoseconds,
				err
		}

	case *int64:

		var ptrInt64 *int64

		ptrInt64, ok = totalTimeDuration.(*int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int64 cast to 'ptrInt64' failed!\n",
				ePrefix.String())

			return numOfDays,
				numOfHours,
				numOfMinutes,
				numOfSeconds,
				numOfMilliseconds,
				numOfMicroseconds,
				numOfNanoseconds,
				err
		}

		int64Value = *ptrInt64

	case time.Duration:

		var timeDurationVal time.Duration

		timeDurationVal, ok = totalTimeDuration.(time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: time.Duration cast to 'timeDurationVal' failed!\n",
				ePrefix.String())

			return numOfDays,
				numOfHours,
				numOfMinutes,
				numOfSeconds,
				numOfMilliseconds,
				numOfMicroseconds,
				numOfNanoseconds,
				err
		}

		int64Value = int64(timeDurationVal)

	case *time.Duration:

		var ptrTimeDuration *time.Duration

		ptrTimeDuration, ok = totalTimeDuration.(*time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *time.Duration cast to 'ptrTimeDuration' failed!\n",
				ePrefix.String())

			return numOfDays,
				numOfHours,
				numOfMinutes,
				numOfSeconds,
				numOfMilliseconds,
				numOfMicroseconds,
				numOfNanoseconds,
				err
		}

		int64Value = int64(*ptrTimeDuration)

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'totalTimeDuration' is an invalid type!\n"+
			"'totalTimeDuration' is unsupported type '%T'\n",
			ePrefix.String(),
			totalTimeDuration)

		return numOfDays,
			numOfHours,
			numOfMinutes,
			numOfSeconds,
			numOfMilliseconds,
			numOfMicroseconds,
			numOfNanoseconds,
			err
	}

	numOfDays,
		numOfHours,
		numOfMinutes,
		numOfSeconds,
		numOfMilliseconds,
		numOfMicroseconds,
		numOfNanoseconds,
		err = new(dateTimeHelperAtom).
		allocateInt64TimeDuration(
			int64Value,
			ePrefix.XCpy(
				"totalTimeDuration"))

	return numOfDays,
		numOfHours,
		numOfMinutes,
		numOfSeconds,
		numOfMilliseconds,
		numOfMicroseconds,
		numOfNanoseconds,
		err
}
