package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// textLineSpecAverageTimeElectron
//
// Provides helper methods for type TextLineSpecAverageTime
type textLineSpecAverageTimeElectron struct {
	lock *sync.Mutex
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
