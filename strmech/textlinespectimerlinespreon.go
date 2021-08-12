package strmech

import "sync"

type textLineSpecTimerLinesPreon struct {
	lock *sync.Mutex
}

// getMaximumTimerLabelLen - Returns the maximum allowable length
// for a text label string describing a timer event element for
// type TextLineSpecTimerLines.
//
// The maximum time label length is the maximum number of
// text characters which can be allocated to 'labelLeftMarginChars'
// plus 'labelRightMarginChars' plus the text label field length
// ('textLabelFieldLen').
//
// The text label field length or 'textLabelFieldLen' value is
// calculated by taking the greater of the following two values:
//
//   (1) The length of the longest text label ('startTimeLabel',
//       'endTimeLabel' or 'timeDurationLabel').
//                     OR
//   (2) The user entered value for 'textLabelFieldLen'
//
// The maximum time label length is currently 55-characters.
//
func (txtTimerLinesPreon *textLineSpecTimerLinesPreon) getMaximumTimerLabelLen() int {

	if txtTimerLinesPreon.lock == nil {
		txtTimerLinesPreon.lock = new(sync.Mutex)
	}

	txtTimerLinesPreon.lock.Lock()

	defer txtTimerLinesPreon.lock.Unlock()

	return 55
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesElectron.
//
func (txtTimerLinesPreon textLineSpecTimerLinesPreon) ptr() *textLineSpecTimerLinesPreon {

	if txtTimerLinesPreon.lock == nil {
		txtTimerLinesPreon.lock = new(sync.Mutex)
	}

	txtTimerLinesPreon.lock.Lock()

	defer txtTimerLinesPreon.lock.Unlock()

	return &textLineSpecTimerLinesPreon{
		lock: new(sync.Mutex),
	}
}
