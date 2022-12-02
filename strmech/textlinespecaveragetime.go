package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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
	numberOfDurationEvents big.Int
	totalDurationNanoSecs  big.Int
	lock                   *sync.Mutex
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
//	All previously collected time event duration data
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

	txtLineAvgTime.numberOfDurationEvents.SetInt64(0)
	txtLineAvgTime.totalDurationNanoSecs.SetInt64(0)

	return
}
