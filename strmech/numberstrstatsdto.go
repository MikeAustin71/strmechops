package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberStrStatsDto
//
// A data transport type designed to store and transmit
// information on a numeric value.
type NumberStrStatsDto struct {
	NumOfIntegerDigits uint64
	//	The total number of integer digits to the left of
	//	the radix point or, decimal point, in the subject
	//	numeric value.

	NumOfSignificantIntegerDigits uint64
	//	The number of nonzero integer digits to the left
	//	of the radix point, or decimal point, in the
	//	subject numeric value.

	NumOfFractionalDigits uint64
	//	The total number of fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumOfSignificantFractionalDigits uint64
	//	The number of nonzero fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumberValueType NumericValueType
	//	This enumeration value specifies whether the
	//	subject numeric value is classified either as an
	//	integer or a floating point value.
	//
	//	Possible enumeration values are listed as
	//	follows:
	//  	NumValType.None()
	//  	NumValType.FloatingPoint()
	//  	NumValType.Integer()

	NumberSign NumericSignValueType
	//	An enumeration specifying the number sign
	//	associated with the numeric value. Possible
	//	values are listed as follows:
	//      NumSignVal.None()		= Invalid Value
	//      NumSignVal.Negative()	= -1
	//      NumSignVal.Zero()		=  0
	//      NumSignVal.Positive()	=  1

	IsZeroValue bool
	//	If 'false', the Numeric Value is greater than or
	//	less than zero ('0').
	//
	//	If 'true', the Numeric Value is equal to zero.

	lock *sync.Mutex
}

// numberStrStatsDtoNanobot
//
// Provides helper methods for type NumberStrStatsDto.
type numberStrStatsDtoNanobot struct {
	lock *sync.Mutex
}

// copyNumStatsDto
//
// Copies NumberStrStatsDto information from a source
// object to a destination object.
func (numStrStatsNanobot *numberStrStatsDtoNanobot) copyNumStatsDto(
	destinationNumStatsDto *NumberStrStatsDto,
	sourceNumStatsDto *NumberStrStatsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrStatsNanobot.lock == nil {
		numStrStatsNanobot.lock = new(sync.Mutex)
	}

	numStrStatsNanobot.lock.Lock()

	defer numStrStatsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrStatsDtoNanobot."+
			"copyNumStatsDto()",
		"")

	if err != nil {

		return err

	}

	if destinationNumStatsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationNumStatsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNumStatsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceNumStatsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	destinationNumStatsDto.NumOfIntegerDigits =
		sourceNumStatsDto.NumOfIntegerDigits

	destinationNumStatsDto.NumOfSignificantIntegerDigits =
		sourceNumStatsDto.NumOfSignificantIntegerDigits

	destinationNumStatsDto.NumOfFractionalDigits =
		sourceNumStatsDto.NumOfFractionalDigits

	destinationNumStatsDto.NumOfSignificantFractionalDigits =
		sourceNumStatsDto.NumOfSignificantFractionalDigits

	destinationNumStatsDto.NumberValueType =
		sourceNumStatsDto.NumberValueType

	destinationNumStatsDto.NumberSign =
		sourceNumStatsDto.NumberSign

	destinationNumStatsDto.IsZeroValue =
		sourceNumStatsDto.IsZeroValue

	return err
}
