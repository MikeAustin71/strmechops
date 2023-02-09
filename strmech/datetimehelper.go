package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strconv"
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
	allocatedTimeDuration TimeDurationDto,
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

		return allocatedTimeDuration,
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

			return allocatedTimeDuration, err
		}

	case *int64:

		var ptrInt64 *int64

		ptrInt64, ok = totalTimeDuration.(*int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int64 cast to 'ptrInt64' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int64Value = *ptrInt64

	case time.Duration:

		var timeDurationVal time.Duration

		timeDurationVal, ok = totalTimeDuration.(time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: time.Duration cast to 'timeDurationVal' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int64Value = int64(timeDurationVal)

	case *time.Duration:

		var ptrTimeDuration *time.Duration

		ptrTimeDuration, ok = totalTimeDuration.(*time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *time.Duration cast to 'ptrTimeDuration' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int64Value = int64(*ptrTimeDuration)

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'totalTimeDuration' is an invalid type!\n"+
			"'totalTimeDuration' is unsupported type '%T'\n",
			ePrefix.String(),
			totalTimeDuration)

		return allocatedTimeDuration, err
	}

	allocatedTimeDuration,
		err = new(dateTimeHelperAtom).
		allocateInt64TimeDuration(
			int64Value,
			ePrefix.XCpy(
				"totalTimeDuration"))

	return allocatedTimeDuration, err
}

// GetDefaultDateTimeFormat
//
// # Returns a string containing a default Date Time Format
//
// "2006-01-02 15:04:05.000000000 -0700 MST"
func (dateTimeHelper *DateTimeHelper) GetDefaultDateTimeFormat() string {

	if dateTimeHelper.lock == nil {
		dateTimeHelper.lock = new(sync.Mutex)
	}

	dateTimeHelper.lock.Lock()

	defer dateTimeHelper.lock.Unlock()

	return new(textSpecificationMolecule).
		getDefaultDateTimeFormat()
}

// GetFmtAllocatedDurationText
//
// Receives time duration broken down by int64 values
// for days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// These numerical values are then converted to
// formatted text strings and returned as an array of
// strings encapsulated in an instance of type
// StringArrayDto.
func (dateTimeHelper *DateTimeHelper) GetFmtAllocatedDurationText(
	allocDuration TimeDurationDto,
	maxLineLength int,
	errorPrefix interface{}) (
	allocatedDurationStrs StringArrayDto,
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
			"GetFmtAllocatedDurationText()",
		"")

	if err != nil {

		return allocatedDurationStrs, err
	}

	var tLine, newOutputLine, finalOutputLine string
	var numStrWithIntSeps []rune
	var lenNewOutputLine, lenFinalOutputLine int
	var foundFirstValue = false

	// Reduce Max Line Length by 1
	// to account for new line character.
	maxLineLength--

	var nStrIntSeparator IntegerSeparatorSpec

	nStrIntSepMolecule := integerSeparatorSpecMolecule{}

	if allocDuration.NumberOfDays > 0 {

		tLine =
			strconv.FormatInt(
				allocDuration.NumberOfDays, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(tLine),
				ePrefix.XCpy("NumberOfDays"))

		if err != nil {
			return allocatedDurationStrs, err
		}

		newOutputLine = fmt.Sprintf(
			"%v Days ",
			string(numStrWithIntSeps))

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(finalOutputLine)

		if lenNewOutputLine >= maxLineLength {

			newOutputLine += "\n"

			allocatedDurationStrs.AddString(
				newOutputLine)

			newOutputLine = ""
			lenNewOutputLine = 0

		} else {

			finalOutputLine += newOutputLine
		}

		foundFirstValue = true
	}

	if len(finalOutputLine) >= maxLineLength {

		finalOutputLine += "\n"

		allocatedDurationStrs.AddString(
			finalOutputLine)

		finalOutputLine = ""

		lenFinalOutputLine = 0

	}

	if allocDuration.NumberOfHours > 0 ||
		foundFirstValue == true {

		tLine =
			strconv.FormatInt(
				allocDuration.NumberOfHours, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(tLine),
				ePrefix.XCpy("NumberOfHours"))

		if err != nil {
			return allocatedDurationStrs, err
		}

		newOutputLine = fmt.Sprintf(
			"%v Hours ",
			string(numStrWithIntSeps))

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(finalOutputLine)

		if lenNewOutputLine+lenFinalOutputLine >=
			maxLineLength {

			if lenFinalOutputLine > 0 {

				finalOutputLine += "\n"

				allocatedDurationStrs.AddString(
					finalOutputLine)

				finalOutputLine = ""

				lenFinalOutputLine = 0

			}

			if lenNewOutputLine >= maxLineLength {

				newOutputLine += "\n"

				allocatedDurationStrs.AddString(
					newOutputLine)

				newOutputLine = ""

				lenNewOutputLine = 0

			} else {

				finalOutputLine += newOutputLine

			}

		} else {

			// MUST BE -
			// lenNewOutputLine + lenFinalOutputLine
			//	LESS THAN maxLineLen

			finalOutputLine += newOutputLine

			lenFinalOutputLine = len(finalOutputLine)
		}

		foundFirstValue = true

	}

	if allocDuration.NumberOfMinutes > 0 ||
		foundFirstValue == true {

		tLine =
			strconv.FormatInt(
				allocDuration.NumberOfMinutes, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(tLine),
				ePrefix.XCpy("NumberOfMinutes"))

		if err != nil {
			return allocatedDurationStrs, err
		}

		newOutputLine = fmt.Sprintf(
			"%v Minutes ",
			string(numStrWithIntSeps))

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(finalOutputLine)

		if lenNewOutputLine+lenFinalOutputLine >=
			maxLineLength {

			if lenFinalOutputLine > 0 {

				finalOutputLine += "\n"

				allocatedDurationStrs.AddString(
					finalOutputLine)

				finalOutputLine = ""

				lenFinalOutputLine = 0

			}

			if lenNewOutputLine >= maxLineLength {

				newOutputLine += "\n"

				allocatedDurationStrs.AddString(
					newOutputLine)

				newOutputLine = ""

				lenNewOutputLine = 0

			} else {

				finalOutputLine += newOutputLine

			}

		} else {

			// MUST BE -
			// lenNewOutputLine + lenFinalOutputLine
			//	LESS THAN maxLineLen

			finalOutputLine += newOutputLine

			lenFinalOutputLine = len(finalOutputLine)
		}

		foundFirstValue = true

	}

	// Always display Seconds

	tLine =
		strconv.FormatInt(
			allocDuration.NumberOfSeconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(tLine),
			ePrefix.XCpy("NumberOfSeconds"))

	if err != nil {
		return allocatedDurationStrs, err
	}

	newOutputLine = fmt.Sprintf(
		"%v Seconds ",
		string(numStrWithIntSeps))

	lenNewOutputLine = len(newOutputLine)

	lenFinalOutputLine = len(finalOutputLine)

	if lenNewOutputLine+lenFinalOutputLine >=
		maxLineLength {

		if lenFinalOutputLine > 0 {

			finalOutputLine += "\n"

			allocatedDurationStrs.AddString(
				finalOutputLine)

			finalOutputLine = ""

			lenFinalOutputLine = 0

		}

		if lenNewOutputLine >= maxLineLength {

			newOutputLine += "\n"

			allocatedDurationStrs.AddString(
				newOutputLine)

			newOutputLine = ""

			lenNewOutputLine = 0

		} else {

			finalOutputLine += newOutputLine

		}

	} else {

		// MUST BE -
		// lenNewOutputLine + lenFinalOutputLine
		//	LESS THAN maxLineLen

		finalOutputLine += newOutputLine

		lenFinalOutputLine = len(finalOutputLine)
	}

	// Always display Milliseconds
	tLine =
		strconv.FormatInt(
			allocDuration.NumberOfMilliseconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(tLine),
			ePrefix.XCpy("NumberOfMilliseconds"))

	if err != nil {
		return allocatedDurationStrs, err
	}

	newOutputLine = fmt.Sprintf(
		"%v Milliseconds ",
		string(numStrWithIntSeps))

	lenNewOutputLine = len(newOutputLine)

	lenFinalOutputLine = len(finalOutputLine)

	if lenNewOutputLine+lenFinalOutputLine >=
		maxLineLength {

		if lenFinalOutputLine > 0 {

			finalOutputLine += "\n"

			allocatedDurationStrs.AddString(
				finalOutputLine)

			finalOutputLine = ""

			lenFinalOutputLine = 0

		}

		if lenNewOutputLine >= maxLineLength {

			newOutputLine += "\n"

			allocatedDurationStrs.AddString(
				newOutputLine)

			newOutputLine = ""

			lenNewOutputLine = 0

		} else {

			finalOutputLine += newOutputLine

		}

	} else {

		// MUST BE -
		// lenNewOutputLine + lenFinalOutputLine
		//	LESS THAN maxLineLen

		finalOutputLine += newOutputLine

		lenFinalOutputLine = len(finalOutputLine)
	}

	// Always display Microseconds
	tLine =
		strconv.FormatInt(
			allocDuration.NumberOfMicroseconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(tLine),
			ePrefix.XCpy("NumberOfMicroseconds"))

	if err != nil {
		return allocatedDurationStrs, err
	}

	newOutputLine = fmt.Sprintf(
		"%v Microseconds ",
		string(numStrWithIntSeps))

	lenNewOutputLine = len(newOutputLine)

	lenFinalOutputLine = len(finalOutputLine)

	if lenNewOutputLine+lenFinalOutputLine >=
		maxLineLength {

		if lenFinalOutputLine > 0 {

			finalOutputLine += "\n"

			allocatedDurationStrs.AddString(
				finalOutputLine)

			finalOutputLine = ""

			lenFinalOutputLine = 0

		}

		if lenNewOutputLine >= maxLineLength {

			newOutputLine += "\n"

			allocatedDurationStrs.AddString(
				newOutputLine)

			newOutputLine = ""

			lenNewOutputLine = 0

		} else {

			finalOutputLine += newOutputLine

		}

	} else {

		// MUST BE -
		// lenNewOutputLine + lenFinalOutputLine
		//	LESS THAN maxLineLen

		finalOutputLine += newOutputLine

		lenFinalOutputLine = len(finalOutputLine)
	}

	// Always display Nanoseconds
	tLine =
		strconv.FormatInt(
			allocDuration.NumberOfNanoseconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(tLine),
			ePrefix.XCpy("NumberOfNanoseconds"))

	if err != nil {
		return allocatedDurationStrs, err
	}

	newOutputLine = fmt.Sprintf(
		"%v Nanoseconds ",
		string(numStrWithIntSeps))

	lenNewOutputLine = len(newOutputLine)

	lenFinalOutputLine = len(finalOutputLine)

	if lenNewOutputLine+lenFinalOutputLine >=
		maxLineLength {

		if lenFinalOutputLine > 0 {

			finalOutputLine += "\n"

			allocatedDurationStrs.AddString(
				finalOutputLine)

			finalOutputLine = ""

			lenFinalOutputLine = 0

		}

		if lenNewOutputLine >= maxLineLength {

			newOutputLine += "\n"

			allocatedDurationStrs.AddString(
				newOutputLine)

			newOutputLine = ""

			lenNewOutputLine = 0

		} else {

			finalOutputLine += newOutputLine

		}

	} else {

		// MUST BE -
		// lenNewOutputLine + lenFinalOutputLine
		//	LESS THAN maxLineLen

		finalOutputLine += newOutputLine

		lenFinalOutputLine = len(finalOutputLine)
	}

	lenFinalOutputLine = len(finalOutputLine)

	if lenFinalOutputLine > 0 {

		finalOutputLine += "\n"

		allocatedDurationStrs.AddString(
			finalOutputLine)

		finalOutputLine = ""

		lenFinalOutputLine = 0

	}

	finalOutputLine = fmt.Sprintf(
		"Total Nanoseconds: %v\n",
		string(numStrWithIntSeps))

	allocatedDurationStrs.AddString(
		finalOutputLine)

	return allocatedDurationStrs, err
}
