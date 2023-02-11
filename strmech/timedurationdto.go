package strmech

import "sync"

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

	lock *sync.Mutex
}

func (timeDurDto *TimeDurationDto) Equal(
	incomingTimeDuration *TimeDurationDto) bool {

	if timeDurDto.lock == nil {
		timeDurDto.lock = new(sync.Mutex)
	}

	timeDurDto.lock.Lock()

	defer timeDurDto.lock.Unlock()

	if timeDurDto.TotalNanoseconds !=
		incomingTimeDuration.TotalNanoseconds {

		return false
	}

	if timeDurDto.NumberOfDays !=
		incomingTimeDuration.NumberOfDays {

		return false
	}

	if timeDurDto.NumberOfHours !=
		incomingTimeDuration.NumberOfHours {

		return false
	}

	if timeDurDto.NumberOfMinutes !=
		incomingTimeDuration.NumberOfMinutes {

		return false
	}

	if timeDurDto.NumberOfSeconds !=
		incomingTimeDuration.NumberOfSeconds {

		return false
	}

	if timeDurDto.NumberOfMilliseconds !=
		incomingTimeDuration.NumberOfMilliseconds {

		return false
	}

	if timeDurDto.NumberOfMicroseconds !=
		incomingTimeDuration.NumberOfMicroseconds {

		return false
	}

	if timeDurDto.NumberOfNanoseconds !=
		incomingTimeDuration.NumberOfNanoseconds {

		return false
	}

	return true
}
