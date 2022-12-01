package strmech

import (
	"math/big"
	"sync"
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
type TextLineSpecAverageTime struct {
	numberOfCycles        big.Int
	totalDurationNanoSecs big.Int
	averageTimeNanoSecs   big.Int
	lock                  *sync.Mutex
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
func (txtLineAvgTime *TextLineSpecAverageTime) New() {

	if txtLineAvgTime.lock == nil {
		txtLineAvgTime.lock = new(sync.Mutex)
	}

	txtLineAvgTime.lock.Lock()

	defer txtLineAvgTime.lock.Unlock()

	newAvgTimer := TextLineSpecAverageTime{}

	newAvgTimer.numberOfCycles.SetInt64(0)
	newAvgTimer.totalDurationNanoSecs.SetInt64(0)
	newAvgTimer.averageTimeNanoSecs.SetInt64(0)

}
