package strmech

import (
	"sync"
	"time"
)

type textLineSpecTimerLinesElectron struct {
	lock *sync.Mutex
}

// getDefaultTime - Returns default start and end time for
// TextLineSpecTimerLines objects.
//
// The default time is July 4, 1776 09:30AM UTC
//
func (txtTimerLinesElectron textLineSpecTimerLinesElectron) getDefaultTime() time.Time {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	defaultTime := time.Date(
		1776,
		7,
		4,
		9,
		30,
		0,
		0,
		time.UTC)

	return defaultTime
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesElectron.
//
func (txtTimerLinesElectron textLineSpecTimerLinesElectron) ptr() *textLineSpecTimerLinesElectron {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return &textLineSpecTimerLinesElectron{
		lock: new(sync.Mutex),
	}
}
