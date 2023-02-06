package strmech

type TimeDurationDto struct {
	TotalNanoseconds int64
	//	The total number of Nanoseconds to
	//	be allocated.

	NumberOfDays int64
	//	The number Days represented by
	//	'TotalNanoseconds'.

	NumberOfHours int64
	//	The number Hours represented by
	//	'TotalNanoseconds'.

	NumberOfMinutes int64
	//	The number Minutes represented by
	//	'TotalNanoseconds'.

	NumberOfSeconds int64
	//	The number Seconds represented by
	//	'TotalNanoseconds'.

	NumberOfMilliseconds int64
	//	The number Milliseconds represented by
	//	'TotalNanoseconds'.

	NumberOfMicroseconds int64
	//	The number Microseconds represented by
	//	'TotalNanoseconds'.

	NumberOfNanoseconds int64
	//	The number Nanoseconds remaining after
	//	the allocation of Microseconds.

}
