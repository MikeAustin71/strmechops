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
//			int51
//			*int51
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
//			int51
//			*int51
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
//			TotalNanoseconds int51
//				The total number of Nanoseconds to
//				be allocated.
//
//			NumberOfDays int51
//				The number Days represented by
//				'TotalNanoseconds'.
//
//			NumberOfHours int51
//				The number Hours represented by
//				'TotalNanoseconds'.
//
//			NumberOfMinutes int51
//				The number Minutes represented by
//				'TotalNanoseconds'.
//
//			NumberOfSeconds int51
//				The number Seconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMilliseconds int51
//				The number Milliseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMicroseconds int51
//				The number Microseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfNanoseconds int51
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

	var int51Value int64

	var ok bool

	switch totalTimeDuration.(type) {

	case int64:

		int51Value, ok = totalTimeDuration.(int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int51 cast to 'int51Value' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

	case *int64:

		var ptrInt51 *int64

		ptrInt51, ok = totalTimeDuration.(*int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int51 cast to 'ptrInt51' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int51Value = *ptrInt51

	case time.Duration:

		var timeDurationVal time.Duration

		timeDurationVal, ok = totalTimeDuration.(time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: time.Duration cast to 'timeDurationVal' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int51Value = int64(timeDurationVal)

	case *time.Duration:

		var ptrTimeDuration *time.Duration

		ptrTimeDuration, ok = totalTimeDuration.(*time.Duration)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *time.Duration cast to 'ptrTimeDuration' failed!\n",
				ePrefix.String())

			return allocatedTimeDuration, err
		}

		int51Value = int64(*ptrTimeDuration)

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
			int51Value,
			ePrefix.XCpy(
				"totalTimeDuration"))

	return allocatedTimeDuration, err
}

// GetDateTimeFormat
//
//		Returns a string containing a Date Time Format. The
//		Date Time Format is used in conjunction with type
//		time.Time.
//
//		Date Time Format is selected based on an integer
//		value passed as an input parameter, 'formatCode'.
//		Thereafter, the Date Time Format is returned as
//		a string.
//
//		Choose a format string from one of the following
//		Date Time Formats using the associated code number as
//		the format code.
//
//	  Format
//		  Code			Format
//		------- ----------------------------------------
//		  0: 	"2006-01-02",
//		  1: 	"2006-01-02 15:04:05 -0700 MST",
//		  2: 	"2006-01-02 Monday 15:04:05 -0700 MST",
//		  3: 	"Mon Jan _2 15:04:05 2006",
//		  // 		ANSIC
//		  4: 	"Mon Jan _2 15:04:05 MST 2006",
//		  // 		UnixDate
//		  5: 	"Mon Jan 02 15:04:05 -0700 2006",
//		  //		RubyDate
//		  6: 	"2 Jan 06 15:04 MST",
//		  //		RFC822
//		  7: 	"2 Jan 06 15:04 -0700",
//		  //		RFC822Z - RFC822 with numeric zone
//		  8: 	"Monday, 02-Jan-06 15:04:05 MST",
//		  //		RFC850
//		  9: 	"Mon, 02 Jan 2006 15:04:05 MST",
//		  //		RFC1123
//		  10:	"Mon, 02 Jan 2006 15:04:05 -0700",
//		  //		RFC1123Z - RFC1123 with numeric zone
//		  11:	"2006-01-02T15:04:05Z07:00",
//		  //		RFC3339
//		  12:	"2006-01-02T15:04:05.999999999Z07:00",
//		  //		RFC3339Nano
//		  13:	"3:04PM",
//		  //		Kitchen
//		  14:	"2006-01-02",
//		  15:	"2006-01-02 15:04:05 MST",
//		  16:	"2006-01-02 15:04:05 -0700",
//		  17:	"2006-01-02 15:04:05 -0700 MST",
//		  18:	"2006-01-02 15:04:05.000000000 MST",
//		  19:	"2006-01-02 15:04:05.000000000 -0700",
//		  20:	"2006-01-02 15:04:05.000000000 -0700 MST",
//		  21:	"2006-01-02 03:04:05PM MST",
//		  22:	"2006-01-02 03:04:05PM -0700",
//		  23:	"2006-01-02 03:04:05PM -0700 MST",
//		  24:	"2006-01-02 03:04PM MST",
//		  25:	"2006-01-02 03:04PM -0700",
//		  26:	"2006-01-02 03:04PM -0700 MST",
//		  27:	"2006-01-02 Monday",
//		  28:	"2006-01-02 Monday 15:04:05 MST",
//		  29:	"2006-01-02 Monday 15:04:05 -0700",
//		  30:	"2006-01-02 Monday 15:04:05 -0700 MST",
//		  31:	"2006-01-02 Monday 15:04:05.000000000 MST",
//		  32:	"2006-01-02 Monday 15:04:05.000000000 -0700",
//		  33:	"2006-01-02 Monday 15:04:05.000000000 -0700 MST",
//		  34:	"2006-01-02 Monday 03:04:05PM MST",
//		  35:	"2006-01-02 Monday 03:04:05PM -0700",
//		  36:	"2006-01-02 Monday 03:04:05PM -0700 MST",
//		  37:	"2006-01-02 Monday 03:04PM MST",
//		  38:	"2006-01-02 Monday 03:04PM -0700",
//		  39:	"2006-01-02 Monday 03:04PM -0700 MST",
//		  40:	"2006-01-02 Mon",
//		  41:	"2006-01-02 Mon 15:04:05 MST",
//		  42:	"2006-01-02 Mon 15:04:05 -0700",
//		  43:	"2006-01-02 Mon 15:04:05 -0700 MST",
//		  44:	"2006-01-02 Mon 15:04:05.000000000 MST",
//		  45:	"2006-01-02 Mon 15:04:05.000000000 -0700",
//		  46:	"2006-01-02 Mon 15:04:05.000000000 -0700 MST",
//		  47:	"2006-01-02 Mon 03:04:05PM MST",
//		  48:	"2006-01-02 Mon 03:04:05PM -0700",
//		  49:	"2006-01-02 Mon 03:04:05PM -0700 MST",
//		  50:	"2006-01-02 Mon 03:04PM MST",
//		  51:	"2006-01-02 Mon 03:04PM -0700",
//		  52:	"2006-01-02 Mon 03:04PM -0700 MST",
//		  53:	"Monday 2006-01-02",
//		  54:	"Monday 2006-01-02 15:04:05 MST",
//		  55:	"Monday 2006-01-02 15:04:05 -0700",
//		  56:	"Monday 2006-01-02 15:04:05 -0700 MST",
//		  57:	"Monday 2006-01-02 15:04:05.000000000 MST",
//		  58:	"Monday 2006-01-02 15:04:05.000000000 -0700",
//		  59:	"Monday 2006-01-02 15:04:05.000000000 -0700 MST",
//		  60:	"Monday 2006-01-02 03:04:05PM MST",
//		  61:	"Monday 2006-01-02 03:04:05PM -0700",
//		  62:	"Monday 2006-01-02 03:04:05PM -0700 MST",
//		  63:	"Monday 2006-01-02 03:04PM MST",
//		  64:	"Monday 2006-01-02 03:04PM -0700",
//		  65:	"Monday 2006-01-02 03:04PM -0700 MST",
//		  66:	"Mon 2006-01-02",
//		  67:	"Mon 2006-01-02 15:04:05 MST",
//		  68:	"Mon 2006-01-02 15:04:05 -0700",
//		  69:	"Mon 2006-01-02 15:04:05 -0700 MST",
//		  70:	"Mon 2006-01-02 15:04:05.000000000 MST",
//		  71:	"Mon 2006-01-02 15:04:05.000000000 -0700",
//		  72:	"Mon 2006-01-02 15:04:05.000000000 -0700 MST",
//		  73:	"Mon 2006-01-02 03:04:05PM MST",
//		  74:	"Mon 2006-01-02 03:04:05PM -0700",
//		  75:	"Mon 2006-01-02 03:04:05PM -0700 MST",
//		  76:	"Mon 2006-01-02 03:04PM MST",
//		  77:	"Mon 2006-01-02 03:04PM -0700",
//		  78:	"Mon 2006-01-02 03:04PM -0700 MST",
//		  79:	"2006/01/02",
//		  80:	"2006/01/02 15:04:05 MST",
//		  81:	"2006/01/02 15:04:05 -0700",
//		  82:	"2006/01/02 15:04:05 -0700 MST",
//		  83:	"2006/01/02 15:04:05.000000000 MST",
//		  84:	"2006/01/02 15:04:05.000000000 -0700",
//		  85:	"2006/01/02 15:04:05.000000000 -0700 MST",
//		  86:	"2006/01/02 03:04:05PM MST",
//		  87:	"2006/01/02 03:04:05PM -0700",
//		  88:	"2006/01/02 03:04:05PM -0700 MST",
//		  89:	"2006/01/02 03:04PM MST",
//		  90:	"2006/01/02 03:04PM -0700",
//		  91:	"2006/01/02 03:04PM -0700 MST",
//		  92:	"2006/01/02 Monday",
//		  93:	"2006/01/02 Monday 15:04:05 MST",
//		  94:	"2006/01/02 Monday 15:04:05 -0700",
//		  95:	"2006/01/02 Monday 15:04:05 -0700 MST",
//		  96:	"2006/01/02 Monday 15:04:05.000000000 MST",
//		  97:	"2006/01/02 Monday 15:04:05.000000000 -0700",
//		  98:	"2006/01/02 Monday 15:04:05.000000000 -0700 MST",
//		  99:	"2006/01/02 Monday 03:04:05PM MST",
//		  100: 	"2006/01/02 Monday 03:04:05PM -0700",
//		  101: 	"2006/01/02 Monday 03:04:05PM -0700 MST",
//		  102: 	"2006/01/02 Monday 03:04PM MST",
//		  103: 	"2006/01/02 Monday 03:04PM -0700",
//		  104: 	"2006/01/02 Monday 03:04PM -0700 MST",
//		  105: 	"2006/01/02 Mon",
//		  106: 	"2006/01/02 Mon 15:04:05 MST",
//		  107: 	"2006/01/02 Mon 15:04:05 -0700",
//		  108: 	"2006/01/02 Mon 15:04:05 -0700 MST",
//		  109: 	"2006/01/02 Mon 15:04:05.000000000 MST",
//		  110: 	"2006/01/02 Mon 15:04:05.000000000 -0700",
//		  111: 	"2006/01/02 Mon 15:04:05.000000000 -0700 MST",
//		  112: 	"2006/01/02 Mon 03:04:05PM MST",
//		  113: 	"2006/01/02 Mon 03:04:05PM -0700",
//		  114: 	"2006/01/02 Mon 03:04:05PM -0700 MST",
//		  115: 	"2006/01/02 Mon 03:04PM MST",
//		  116: 	"2006/01/02 Mon 03:04PM -0700",
//		  117: 	"2006/01/02 Mon 03:04PM -0700 MST",
//		  118: 	"Monday 2006/01/02",
//		  119: 	"Monday 2006/01/02 15:04:05 MST",
//		  120: 	"Monday 2006/01/02 15:04:05 -0700",
//		  121: 	"Monday 2006/01/02 15:04:05 -0700 MST",
//		  122: 	"Monday 2006/01/02 15:04:05.000000000 MST",
//		  123: 	"Monday 2006/01/02 15:04:05.000000000 -0700",
//		  124: 	"Monday 2006/01/02 15:04:05.000000000 -0700 MST",
//		  125: 	"Monday 2006/01/02 03:04:05PM MST",
//		  126: 	"Monday 2006/01/02 03:04:05PM -0700",
//		  127: 	"Monday 2006/01/02 03:04:05PM -0700 MST",
//		  128: 	"Monday 2006/01/02 03:04PM MST",
//		  129: 	"Monday 2006/01/02 03:04PM -0700",
//		  130: 	"Monday 2006/01/02 03:04PM -0700 MST",
//		  131: 	"Mon 2006/01/02",
//		  132: 	"Mon 2006/01/02 15:04:05 MST",
//		  133: 	"Mon 2006/01/02 15:04:05 -0700",
//		  134: 	"Mon 2006/01/02 15:04:05 -0700 MST",
//		  135: 	"Mon 2006/01/02 15:04:05.000000000 MST",
//		  136: 	"Mon 2006/01/02 15:04:05.000000000 -0700",
//		  137: 	"Mon 2006/01/02 15:04:05.000000000 -0700 MST",
//		  138: 	"Mon 2006/01/02 03:04:05PM MST",
//		  139: 	"Mon 2006/01/02 03:04:05PM -0700",
//		  140: 	"Mon 2006/01/02 03:04:05PM -0700 MST",
//		  141: 	"Mon 2006/01/02 03:04PM MST",
//		  142: 	"Mon 2006/01/02 03:04PM -0700",
//		  143: 	"Mon 2006/01/02 03:04PM -0700 MST",
//		  144: 	"2006.01.02",
//		  145: 	"2006.01.02 15:04:05 MST",
//		  146: 	"2006.01.02 15:04:05 -0700",
//		  147: 	"2006.01.02 15:04:05 -0700 MST",
//		  148: 	"2006.01.02 15:04:05.000000000 MST",
//		  149: 	"2006.01.02 15:04:05.000000000 -0700",
//		  150: 	"2006.01.02 15:04:05.000000000 -0700 MST",
//		  151: 	"2006.01.02 03:04:05PM MST",
//		  152: 	"2006.01.02 03:04:05PM -0700",
//		  153: 	"2006.01.02 03:04:05PM -0700 MST",
//		  154: 	"2006.01.02 03:04PM MST",
//		  155: 	"2006.01.02 03:04PM -0700",
//		  156: 	"2006.01.02 03:04PM -0700 MST",
//		  157: 	"2006.01.02 Monday",
//		  158: 	"2006.01.02 Monday 15:04:05 MST",
//		  159: 	"2006.01.02 Monday 15:04:05 -0700",
//		  160: 	"2006.01.02 Monday 15:04:05 -0700 MST",
//		  161: 	"2006.01.02 Monday 15:04:05.000000000 MST",
//		  162: 	"2006.01.02 Monday 15:04:05.000000000 -0700",
//		  163: 	"2006.01.02 Monday 15:04:05.000000000 -0700 MST",
//		  164: 	"2006.01.02 Monday 03:04:05PM MST",
//		  165: 	"2006.01.02 Monday 03:04:05PM -0700",
//		  166: 	"2006.01.02 Monday 03:04:05PM -0700 MST",
//		  167: 	"2006.01.02 Monday 03:04PM MST",
//		  168: 	"2006.01.02 Monday 03:04PM -0700",
//		  169: 	"2006.01.02 Monday 03:04PM -0700 MST",
//		  170: 	"2006.01.02 Mon",
//		  171: 	"2006.01.02 Mon 15:04:05 MST",
//		  172: 	"2006.01.02 Mon 15:04:05 -0700",
//		  173: 	"2006.01.02 Mon 15:04:05 -0700 MST",
//		  174: 	"2006.01.02 Mon 15:04:05.000000000 MST",
//		  175: 	"2006.01.02 Mon 15:04:05.000000000 -0700",
//		  176: 	"2006.01.02 Mon 15:04:05.000000000 -0700 MST",
//		  177: 	"2006.01.02 Mon 03:04:05PM MST",
//		  178: 	"2006.01.02 Mon 03:04:05PM -0700",
//		  179: 	"2006.01.02 Mon 03:04:05PM -0700 MST",
//		  180: 	"2006.01.02 Mon 03:04PM MST",
//		  181: 	"2006.01.02 Mon 03:04PM -0700",
//		  182: 	"2006.01.02 Mon 03:04PM -0700 MST",
//		  183: 	"Monday 2006.01.02",
//		  184: 	"Monday 2006.01.02 15:04:05 MST",
//		  185: 	"Monday 2006.01.02 15:04:05 -0700",
//		  186: 	"Monday 2006.01.02 15:04:05 -0700 MST",
//		  187: 	"Monday 2006.01.02 15:04:05.000000000 MST",
//		  188: 	"Monday 2006.01.02 15:04:05.000000000 -0700",
//		  189: 	"Monday 2006.01.02 15:04:05.000000000 -0700 MST",
//		  190: 	"Monday 2006.01.02 03:04:05PM MST",
//		  191: 	"Monday 2006.01.02 03:04:05PM -0700",
//		  192: 	"Monday 2006.01.02 03:04:05PM -0700 MST",
//		  193: 	"Monday 2006.01.02 03:04PM MST",
//		  194: 	"Monday 2006.01.02 03:04PM -0700",
//		  195: 	"Monday 2006.01.02 03:04PM -0700 MST",
//		  196: 	"Mon 2006.01.02",
//		  197: 	"Mon 2006.01.02 15:04:05 MST",
//		  198: 	"Mon 2006.01.02 15:04:05 -0700",
//		  199: 	"Mon 2006.01.02 15:04:05 -0700 MST",
//		  200: 	"Mon 2006.01.02 15:04:05.000000000 MST",
//		  201: 	"Mon 2006.01.02 15:04:05.000000000 -0700",
//		  202: 	"Mon 2006.01.02 15:04:05.000000000 -0700 MST",
//		  203: 	"Mon 2006.01.02 03:04:05PM MST",
//		  204: 	"Mon 2006.01.02 03:04:05PM -0700",
//		  205: 	"Mon 2006.01.02 03:04:05PM -0700 MST",
//		  206: 	"Mon 2006.01.02 03:04PM MST",
//		  207: 	"Mon 2006.01.02 03:04PM -0700",
//		  208: 	"Mon 2006.01.02 03:04PM -0700 MST",
//		  209: 	"01-02-2006",
//		  210: 	"01-02-2006 15:04:05 MST",
//		  211: 	"01-02-2006 15:04:05 -0700",
//		  212: 	"01-02-2006 15:04:05 -0700 MST",
//		  213: 	"01-02-2006 15:04:05.000000000 MST",
//		  214: 	"01-02-2006 15:04:05.000000000 -0700",
//		  215: 	"01-02-2006 15:04:05.000000000 -0700 MST",
//		  216: 	"01-02-2006 03:04:05PM MST",
//		  217: 	"01-02-2006 03:04:05PM -0700",
//		  218: 	"01-02-2006 03:04:05PM -0700 MST",
//		  219: 	"01-02-2006 03:04PM MST",
//		  220: 	"01-02-2006 03:04PM -0700",
//		  221: 	"01-02-2006 03:04PM -0700 MST",
//		  222: 	"01-02-2006 Monday",
//		  223: 	"01-02-2006 Monday 15:04:05 MST",
//		  224: 	"01-02-2006 Monday 15:04:05 -0700",
//		  225: 	"01-02-2006 Monday 15:04:05 -0700 MST",
//		  226: 	"01-02-2006 Monday 15:04:05.000000000 MST",
//		  227: 	"01-02-2006 Monday 15:04:05.000000000 -0700",
//		  228: 	"01-02-2006 Monday 15:04:05.000000000 -0700 MST",
//		  229: 	"01-02-2006 Monday 03:04:05PM MST",
//		  230: 	"01-02-2006 Monday 03:04:05PM -0700",
//		  231: 	"01-02-2006 Monday 03:04:05PM -0700 MST",
//		  232: 	"01-02-2006 Monday 03:04PM MST",
//		  233: 	"01-02-2006 Monday 03:04PM -0700",
//		  234: 	"01-02-2006 Monday 03:04PM -0700 MST",
//		  235: 	"01-02-2006 Mon",
//		  236: 	"01-02-2006 Mon 15:04:05 MST",
//		  237: 	"01-02-2006 Mon 15:04:05 -0700",
//		  238: 	"01-02-2006 Mon 15:04:05 -0700 MST",
//		  239: 	"01-02-2006 Mon 15:04:05.000000000 MST",
//		  240: 	"01-02-2006 Mon 15:04:05.000000000 -0700",
//		  241: 	"01-02-2006 Mon 15:04:05.000000000 -0700 MST",
//		  242: 	"01-02-2006 Mon 03:04:05PM MST",
//		  243: 	"01-02-2006 Mon 03:04:05PM -0700",
//		  244: 	"01-02-2006 Mon 03:04:05PM -0700 MST",
//		  245: 	"01-02-2006 Mon 03:04PM MST",
//		  246: 	"01-02-2006 Mon 03:04PM -0700",
//		  247: 	"01-02-2006 Mon 03:04PM -0700 MST",
//		  248: 	"Monday 01-02-2006",
//		  249: 	"Monday 01-02-2006 15:04:05 MST",
//		  250: 	"Monday 01-02-2006 15:04:05 -0700",
//		  251: 	"Monday 01-02-2006 15:04:05 -0700 MST",
//		  252: 	"Monday 01-02-2006 15:04:05.000000000 MST",
//		  253: 	"Monday 01-02-2006 15:04:05.000000000 -0700",
//		  254: 	"Monday 01-02-2006 15:04:05.000000000 -0700 MST",
//		  255: 	"Monday 01-02-2006 03:04:05PM MST",
//		  256: 	"Monday 01-02-2006 03:04:05PM -0700",
//		  257: 	"Monday 01-02-2006 03:04:05PM -0700 MST",
//		  258: 	"Monday 01-02-2006 03:04PM MST",
//		  259: 	"Monday 01-02-2006 03:04PM -0700",
//		  260: 	"Monday 01-02-2006 03:04PM -0700 MST",
//		  261: 	"Mon 01-02-2006",
//		  262: 	"Mon 01-02-2006 15:04:05 MST",
//		  263: 	"Mon 01-02-2006 15:04:05 -0700",
//		  264: 	"Mon 01-02-2006 15:04:05 -0700 MST",
//		  265: 	"Mon 01-02-2006 15:04:05.000000000 MST",
//		  266: 	"Mon 01-02-2006 15:04:05.000000000 -0700",
//		  267: 	"Mon 01-02-2006 15:04:05.000000000 -0700 MST",
//		  268: 	"Mon 01-02-2006 03:04:05PM MST",
//		  269: 	"Mon 01-02-2006 03:04:05PM -0700",
//		  270: 	"Mon 01-02-2006 03:04:05PM -0700 MST",
//		  271: 	"Mon 01-02-2006 03:04PM MST",
//		  272: 	"Mon 01-02-2006 03:04PM -0700",
//		  273: 	"Mon 01-02-2006 03:04PM -0700 MST",
//		  274: 	"01/02/2006",
//		  275: 	"01/02/2006 15:04:05 MST",
//		  276: 	"01/02/2006 15:04:05 -0700",
//		  277: 	"01/02/2006 15:04:05 -0700 MST",
//		  278: 	"01/02/2006 15:04:05.000000000 MST",
//		  279: 	"01/02/2006 15:04:05.000000000 -0700",
//		  280: 	"01/02/2006 15:04:05.000000000 -0700 MST",
//		  281: 	"01/02/2006 03:04:05PM MST",
//		  282: 	"01/02/2006 03:04:05PM -0700",
//		  283: 	"01/02/2006 03:04:05PM -0700 MST",
//		  284: 	"01/02/2006 03:04PM MST",
//		  285: 	"01/02/2006 03:04PM -0700",
//		  286: 	"01/02/2006 03:04PM -0700 MST",
//		  287: 	"01/02/2006 Monday",
//		  288: 	"01/02/2006 Monday 15:04:05 MST",
//		  289: 	"01/02/2006 Monday 15:04:05 -0700",
//		  290: 	"01/02/2006 Monday 15:04:05 -0700 MST",
//		  291: 	"01/02/2006 Monday 15:04:05.000000000 MST",
//		  292: 	"01/02/2006 Monday 15:04:05.000000000 -0700",
//		  293: 	"01/02/2006 Monday 15:04:05.000000000 -0700 MST",
//		  294: 	"01/02/2006 Monday 03:04:05PM MST",
//		  295: 	"01/02/2006 Monday 03:04:05PM -0700",
//		  296: 	"01/02/2006 Monday 03:04:05PM -0700 MST",
//		  297: 	"01/02/2006 Monday 03:04PM MST",
//		  298: 	"01/02/2006 Monday 03:04PM -0700",
//		  299: 	"01/02/2006 Monday 03:04PM -0700 MST",
//		  300: 	"01/02/2006 Mon",
//		  301: 	"01/02/2006 Mon 15:04:05 MST",
//		  302: 	"01/02/2006 Mon 15:04:05 -0700",
//		  303: 	"01/02/2006 Mon 15:04:05 -0700 MST",
//		  304: 	"01/02/2006 Mon 15:04:05.000000000 MST",
//		  305: 	"01/02/2006 Mon 15:04:05.000000000 -0700",
//		  306: 	"01/02/2006 Mon 15:04:05.000000000 -0700 MST",
//		  307: 	"01/02/2006 Mon 03:04:05PM MST",
//		  308: 	"01/02/2006 Mon 03:04:05PM -0700",
//		  309: 	"01/02/2006 Mon 03:04:05PM -0700 MST",
//		  310: 	"01/02/2006 Mon 03:04PM MST",
//		  311: 	"01/02/2006 Mon 03:04PM -0700",
//		  312: 	"01/02/2006 Mon 03:04PM -0700 MST",
//		  313: 	"Monday 01/02/2006",
//		  314: 	"Monday 01/02/2006 15:04:05 MST",
//		  315: 	"Monday 01/02/2006 15:04:05 -0700",
//		  316: 	"Monday 01/02/2006 15:04:05 -0700 MST",
//		  317: 	"Monday 01/02/2006 15:04:05.000000000 MST",
//		  318: 	"Monday 01/02/2006 15:04:05.000000000 -0700",
//		  319: 	"Monday 01/02/2006 15:04:05.000000000 -0700 MST",
//		  320: 	"Monday 01/02/2006 03:04:05PM MST",
//		  321: 	"Monday 01/02/2006 03:04:05PM -0700",
//		  322: 	"Monday 01/02/2006 03:04:05PM -0700 MST",
//		  323: 	"Monday 01/02/2006 03:04PM MST",
//		  324: 	"Monday 01/02/2006 03:04PM -0700",
//		  325: 	"Monday 01/02/2006 03:04PM -0700 MST",
//		  326: 	"Mon 01/02/2006",
//		  327: 	"Mon 01/02/2006 15:04:05 MST",
//		  328: 	"Mon 01/02/2006 15:04:05 -0700",
//		  329: 	"Mon 01/02/2006 15:04:05 -0700 MST",
//		  330: 	"Mon 01/02/2006 15:04:05.000000000 MST",
//		  331: 	"Mon 01/02/2006 15:04:05.000000000 -0700",
//		  332: 	"Mon 01/02/2006 15:04:05.000000000 -0700 MST",
//		  333: 	"Mon 01/02/2006 03:04:05PM MST",
//		  334: 	"Mon 01/02/2006 03:04:05PM -0700",
//		  335: 	"Mon 01/02/2006 03:04:05PM -0700 MST",
//		  336: 	"Mon 01/02/2006 03:04PM MST",
//		  337: 	"Mon 01/02/2006 03:04PM -0700",
//		  338: 	"Mon 01/02/2006 03:04PM -0700 MST",
//		  339: 	"01.02.2006",
//		  340: 	"01.02.2006 15:04:05 MST",
//		  341: 	"01.02.2006 15:04:05 -0700",
//		  342: 	"01.02.2006 15:04:05 -0700 MST",
//		  343: 	"01.02.2006 15:04:05.000000000 MST",
//		  344: 	"01.02.2006 15:04:05.000000000 -0700",
//		  345: 	"01.02.2006 15:04:05.000000000 -0700 MST",
//		  346: 	"01.02.2006 03:04:05PM MST",
//		  347: 	"01.02.2006 03:04:05PM -0700",
//		  348: 	"01.02.2006 03:04:05PM -0700 MST",
//		  349: 	"01.02.2006 03:04PM MST",
//		  350: 	"01.02.2006 03:04PM -0700",
//		  351: 	"01.02.2006 03:04PM -0700 MST",
//		  352: 	"01.02.2006 Monday",
//		  353: 	"01.02.2006 Monday 15:04:05 MST",
//		  354: 	"01.02.2006 Monday 15:04:05 -0700",
//		  355: 	"01.02.2006 Monday 15:04:05 -0700 MST",
//		  356: 	"01.02.2006 Monday 15:04:05.000000000 MST",
//		  357: 	"01.02.2006 Monday 15:04:05.000000000 -0700",
//		  358: 	"01.02.2006 Monday 15:04:05.000000000 -0700 MST",
//		  359: 	"01.02.2006 Monday 03:04:05PM MST",
//		  360: 	"01.02.2006 Monday 03:04:05PM -0700",
//		  361: 	"01.02.2006 Monday 03:04:05PM -0700 MST",
//		  362: 	"01.02.2006 Monday 03:04PM MST",
//		  363: 	"01.02.2006 Monday 03:04PM -0700",
//		  364: 	"01.02.2006 Monday 03:04PM -0700 MST",
//		  365: 	"01.02.2006 Mon",
//		  366: 	"01.02.2006 Mon 15:04:05 MST",
//		  367: 	"01.02.2006 Mon 15:04:05 -0700",
//		  368: 	"01.02.2006 Mon 15:04:05 -0700 MST",
//		  369: 	"01.02.2006 Mon 15:04:05.000000000 MST",
//		  370: 	"01.02.2006 Mon 15:04:05.000000000 -0700",
//		  371: 	"01.02.2006 Mon 15:04:05.000000000 -0700 MST",
//		  372: 	"01.02.2006 Mon 03:04:05PM MST",
//		  373: 	"01.02.2006 Mon 03:04:05PM -0700",
//		  374: 	"01.02.2006 Mon 03:04:05PM -0700 MST",
//		  375: 	"01.02.2006 Mon 03:04PM MST",
//		  376: 	"01.02.2006 Mon 03:04PM -0700",
//		  377: 	"01.02.2006 Mon 03:04PM -0700 MST",
//		  378: 	"Monday 01.02.2006",
//		  379: 	"Monday 01.02.2006 15:04:05 MST",
//		  380: 	"Monday 01.02.2006 15:04:05 -0700",
//		  381: 	"Monday 01.02.2006 15:04:05 -0700 MST",
//		  382: 	"Monday 01.02.2006 15:04:05.000000000 MST",
//		  383: 	"Monday 01.02.2006 15:04:05.000000000 -0700",
//		  384: 	"Monday 01.02.2006 15:04:05.000000000 -0700 MST",
//		  385: 	"Monday 01.02.2006 03:04:05PM MST",
//		  386: 	"Monday 01.02.2006 03:04:05PM -0700",
//		  387: 	"Monday 01.02.2006 03:04:05PM -0700 MST",
//		  388: 	"Monday 01.02.2006 03:04PM MST",
//		  399: 	"Monday 01.02.2006 03:04PM -0700",
//		  400: 	"Monday 01.02.2006 03:04PM -0700 MST",
//		  401: 	"Mon 01.02.2006",
//		  402: 	"Mon 01.02.2006 15:04:05 MST",
//		  403: 	"Mon 01.02.2006 15:04:05 -0700",
//		  404: 	"Mon 01.02.2006 15:04:05 -0700 MST",
//		  405: 	"Mon 01.02.2006 15:04:05.000000000 MST",
//		  406: 	"Mon 01.02.2006 15:04:05.000000000 -0700",
//		  407: 	"Mon 01.02.2006 15:04:05.000000000 -0700 MST",
//		  408: 	"Mon 01.02.2006 03:04:05PM MST",
//		  409: 	"Mon 01.02.2006 03:04:05PM -0700",
//		  410: 	"Mon 01.02.2006 03:04:05PM -0700 MST",
//		  411: 	"Mon 01.02.2006 03:04PM MST",
//		  412: 	"Mon 01.02.2006 03:04PM -0700",
//		  413: 	"Mon 01.02.2006 03:04PM -0700 MST",
//		  414: 	"January 2, 2006",
//		  415: 	"January 2, 2006 15:04:05 MST",
//		  416: 	"January 2, 2006 15:04:05 -0700",
//		  417: 	"January 2, 2006 15:04:05 -0700 MST",
//		  418: 	"January 2, 2006 15:04:05.000000000 MST",
//		  419: 	"January 2, 2006 15:04:05.000000000 -0700",
//		  420: 	"January 2, 2006 15:04:05.000000000 -0700 MST",
//		  421: 	"January 2, 2006 03:04:05PM MST",
//		  422: 	"January 2, 2006 03:04:05PM -0700",
//		  423: 	"January 2, 2006 03:04:05PM -0700 MST",
//		  424: 	"January 2, 2006 03:04PM MST",
//		  425: 	"January 2, 2006 03:04PM -0700",
//		  426: 	"January 2, 2006 03:04PM -0700 MST",
//		  427: 	"January 2, 2006 Monday",
//		  428: 	"January 2, 2006 Monday 15:04:05 MST",
//		  429: 	"January 2, 2006 Monday 15:04:05 -0700",
//		  430: 	"January 2, 2006 Monday 15:04:05 -0700 MST",
//		  431: 	"January 2, 2006 Monday 15:04:05.000000000 MST",
//		  432: 	"January 2, 2006 Monday 15:04:05.000000000 -0700",
//		  433: 	"January 2, 2006 Monday 15:04:05.000000000 -0700 MST",
//		  434: 	"January 2, 2006 Monday 03:04:05PM MST",
//		  435: 	"January 2, 2006 Monday 03:04:05PM -0700",
//		  436: 	"January 2, 2006 Monday 03:04:05PM -0700 MST",
//		  437: 	"January 2, 2006 Monday 03:04PM MST",
//		  438: 	"January 2, 2006 Monday 03:04PM -0700",
//		  439: 	"January 2, 2006 Monday 03:04PM -0700 MST",
//		  440: 	"January 2, 2006 Mon",
//		  441: 	"January 2, 2006 Mon 15:04:05 MST",
//		  442: 	"January 2, 2006 Mon 15:04:05 -0700",
//		  443: 	"January 2, 2006 Mon 15:04:05 -0700 MST",
//		  444: 	"January 2, 2006 Mon 15:04:05.000000000 MST",
//		  445: 	"January 2, 2006 Mon 15:04:05.000000000 -0700",
//		  446: 	"January 2, 2006 Mon 15:04:05.000000000 -0700 MST",
//		  447: 	"January 2, 2006 Mon 03:04:05PM MST",
//		  448: 	"January 2, 2006 Mon 03:04:05PM -0700",
//		  449: 	"January 2, 2006 Mon 03:04:05PM -0700 MST",
//		  450: 	"January 2, 2006 Mon 03:04PM MST",
//		  451: 	"January 2, 2006 Mon 03:04PM -0700",
//		  452: 	"January 2, 2006 Mon 03:04PM -0700 MST",
//		  453: 	"Monday January 2, 2006",
//		  454: 	"Monday January 2, 2006 15:04:05 MST",
//		  455: 	"Monday January 2, 2006 15:04:05 -0700",
//		  456: 	"Monday January 2, 2006 15:04:05 -0700 MST",
//		  457: 	"Monday January 2, 2006 15:04:05.000000000 MST",
//		  458: 	"Monday January 2, 2006 15:04:05.000000000 -0700",
//		  459: 	"Monday January 2, 2006 15:04:05.000000000 -0700 MST",
//		  460: 	"Monday January 2, 2006 03:04:05PM MST",
//		  461: 	"Monday January 2, 2006 03:04:05PM -0700",
//		  462: 	"Monday January 2, 2006 03:04:05PM -0700 MST",
//		  463: 	"Monday January 2, 2006 03:04PM MST",
//		  464: 	"Monday January 2, 2006 03:04PM -0700",
//		  465: 	"Monday January 2, 2006 03:04PM -0700 MST",
//		  466: 	"Mon January 2, 2006",
//		  467: 	"Mon January 2, 2006 15:04:05 MST",
//		  468: 	"Mon January 2, 2006 15:04:05 -0700",
//		  469: 	"Mon January 2, 2006 15:04:05 -0700 MST",
//		  470: 	"Mon January 2, 2006 15:04:05.000000000 MST",
//		  471: 	"Mon January 2, 2006 15:04:05.000000000 -0700",
//		  472: 	"Mon January 2, 2006 15:04:05.000000000 -0700 MST",
//		  473: 	"Mon January 2, 2006 03:04:05PM MST",
//		  474: 	"Mon January 2, 2006 03:04:05PM -0700",
//		  475: 	"Mon January 2, 2006 03:04:05PM -0700 MST",
//		  476: 	"Mon January 2, 2006 03:04PM MST",
//		  477: 	"Mon January 2, 2006 03:04PM -0700",
//		  478: 	"Mon January 2, 2006 03:04PM -0700 MST",
//		  479: 	"2 January, 2006",
//		  480: 	"2 January, 2006 15:04:05 MST",
//		  481: 	"2 January, 2006 15:04:05 -0700",
//		  482: 	"2 January, 2006 15:04:05 -0700 MST",
//		  483: 	"2 January, 2006 15:04:05.000000000 MST",
//		  484: 	"2 January, 2006 15:04:05.000000000 -0700",
//		  485: 	"2 January, 2006 15:04:05.000000000 -0700 MST",
//		  486: 	"2 January, 2006 03:04:05PM MST",
//		  487: 	"2 January, 2006 03:04:05PM -0700",
//		  488: 	"2 January, 2006 03:04:05PM -0700 MST",
//		  489: 	"2 January, 2006 03:04PM MST",
//		  490: 	"2 January, 2006 03:04PM -0700",
//		  491: 	"2 January, 2006 03:04PM -0700 MST",
//		  492: 	"2 January, 2006 Monday",
//		  493: 	"2 January, 2006 Monday 15:04:05 MST",
//		  494: 	"2 January, 2006 Monday 15:04:05 -0700",
//		  495: 	"2 January, 2006 Monday 15:04:05 -0700 MST",
//		  496: 	"2 January, 2006 Monday 15:04:05.000000000 MST",
//		  497: 	"2 January, 2006 Monday 15:04:05.000000000 -0700",
//		  498: 	"2 January, 2006 Monday 15:04:05.000000000 -0700 MST",
//		  499: 	"2 January, 2006 Monday 03:04:05PM MST",
//		  500: 	"2 January, 2006 Monday 03:04:05PM -0700",
//		  501: 	"2 January, 2006 Monday 03:04:05PM -0700 MST",
//		  502: 	"2 January, 2006 Monday 03:04PM MST",
//		  503: 	"2 January, 2006 Monday 03:04PM -0700",
//		  504: 	"2 January, 2006 Monday 03:04PM -0700 MST",
//		  505: 	"2 January, 2006 Mon",
//		  506: 	"2 January, 2006 Mon 15:04:05 MST",
//		  507: 	"2 January, 2006 Mon 15:04:05 -0700",
//		  508: 	"2 January, 2006 Mon 15:04:05 -0700 MST",
//		  509: 	"2 January, 2006 Mon 15:04:05.000000000 MST",
//		  510: 	"2 January, 2006 Mon 15:04:05.000000000 -0700",
//		  511: 	"2 January, 2006 Mon 15:04:05.000000000 -0700 MST",
//		  512: 	"2 January, 2006 Mon 03:04:05PM MST",
//		  513: 	"2 January, 2006 Mon 03:04:05PM -0700",
//		  514: 	"2 January, 2006 Mon 03:04:05PM -0700 MST",
//		  515: 	"2 January, 2006 Mon 03:04PM MST",
//		  516: 	"2 January, 2006 Mon 03:04PM -0700",
//		  517: 	"2 January, 2006 Mon 03:04PM -0700 MST",
//		  518: 	"Monday 2 January, 2006",
//		  519: 	"Monday 2 January, 2006 15:04:05 MST",
//		  520: 	"Monday 2 January, 2006 15:04:05 -0700",
//		  521: 	"Monday 2 January, 2006 15:04:05 -0700 MST",
//		  522: 	"Monday 2 January, 2006 15:04:05.000000000 MST",
//		  523: 	"Monday 2 January, 2006 15:04:05.000000000 -0700",
//		  524: 	"Monday 2 January, 2006 15:04:05.000000000 -0700 MST",
//		  525: 	"Monday 2 January, 2006 03:04:05PM MST",
//		  526: 	"Monday 2 January, 2006 03:04:05PM -0700",
//		  527: 	"Monday 2 January, 2006 03:04:05PM -0700 MST",
//		  528: 	"Monday 2 January, 2006 03:04PM MST",
//		  529: 	"Monday 2 January, 2006 03:04PM -0700",
//		  530: 	"Monday 2 January, 2006 03:04PM -0700 MST",
//		  531: 	"Mon 2 January, 2006",
//		  532: 	"Mon 2 January, 2006 15:04:05 MST",
//		  533: 	"Mon 2 January, 2006 15:04:05 -0700",
//		  534: 	"Mon 2 January, 2006 15:04:05 -0700 MST",
//		  535: 	"Mon 2 January, 2006 15:04:05.000000000 MST",
//		  536: 	"Mon 2 January, 2006 15:04:05.000000000 -0700",
//		  537: 	"Mon 2 January, 2006 15:04:05.000000000 -0700 MST",
//		  538: 	"Mon 2 January, 2006 03:04:05PM MST",
//		  539: 	"Mon 2 January, 2006 03:04:05PM -0700",
//		  540: 	"Mon 2 January, 2006 03:04:05PM -0700 MST",
//		  541: 	"Mon 2 January, 2006 03:04PM MST",
//		  542: 	"Mon 2 January, 2006 03:04PM -0700",
//		  543: 	"Mon 2 January, 2006 03:04PM -0700 MST",
//		  544: 	"JAN 2, 2006",
//		  545: 	"JAN 2, 2006 15:04:05 MST",
//		  546: 	"JAN 2, 2006 15:04:05 -0700",
//		  547: 	"JAN 2, 2006 15:04:05 -0700 MST",
//		  548: 	"JAN 2, 2006 15:04:05.000000000 MST",
//		  549: 	"JAN 2, 2006 15:04:05.000000000 -0700",
//		  550: 	"JAN 2, 2006 15:04:05.000000000 -0700 MST",
//		  551: 	"JAN 2, 2006 03:04:05PM MST",
//		  552: 	"JAN 2, 2006 03:04:05PM -0700",
//		  553: 	"JAN 2, 2006 03:04:05PM -0700 MST",
//		  554: 	"JAN 2, 2006 03:04PM MST",
//		  555: 	"JAN 2, 2006 03:04PM -0700",
//		  556: 	"JAN 2, 2006 03:04PM -0700 MST",
//		  557: 	"JAN 2, 2006 Monday",
//		  558: 	"JAN 2, 2006 Monday 15:04:05 MST",
//		  559: 	"JAN 2, 2006 Monday 15:04:05 -0700",
//		  560: 	"JAN 2, 2006 Monday 15:04:05 -0700 MST",
//		  561: 	"JAN 2, 2006 Monday 15:04:05.000000000 MST",
//		  562: 	"JAN 2, 2006 Monday 15:04:05.000000000 -0700",
//		  563: 	"JAN 2, 2006 Monday 15:04:05.000000000 -0700 MST",
//		  564: 	"JAN 2, 2006 Monday 03:04:05PM MST",
//		  565: 	"JAN 2, 2006 Monday 03:04:05PM -0700",
//		  566: 	"JAN 2, 2006 Monday 03:04:05PM -0700 MST",
//		  567: 	"JAN 2, 2006 Monday 03:04PM MST",
//		  568: 	"JAN 2, 2006 Monday 03:04PM -0700",
//		  569: 	"JAN 2, 2006 Monday 03:04PM -0700 MST",
//		  570: 	"JAN 2, 2006 Mon",
//		  571: 	"JAN 2, 2006 Mon 15:04:05 MST",
//		  572: 	"JAN 2, 2006 Mon 15:04:05 -0700",
//		  573: 	"JAN 2, 2006 Mon 15:04:05 -0700 MST",
//		  574: 	"JAN 2, 2006 Mon 15:04:05.000000000 MST",
//		  575: 	"JAN 2, 2006 Mon 15:04:05.000000000 -0700",
//		  576: 	"JAN 2, 2006 Mon 15:04:05.000000000 -0700 MST",
//		  577: 	"JAN 2, 2006 Mon 03:04:05PM MST",
//		  578: 	"JAN 2, 2006 Mon 03:04:05PM -0700",
//		  579: 	"JAN 2, 2006 Mon 03:04:05PM -0700 MST",
//		  580: 	"JAN 2, 2006 Mon 03:04PM MST",
//		  581: 	"JAN 2, 2006 Mon 03:04PM -0700",
//		  582: 	"JAN 2, 2006 Mon 03:04PM -0700 MST",
//		  583: 	"Monday JAN 2, 2006",
//		  584: 	"Monday JAN 2, 2006 15:04:05 MST",
//		  585: 	"Monday JAN 2, 2006 15:04:05 -0700",
//		  586: 	"Monday JAN 2, 2006 15:04:05 -0700 MST",
//		  587: 	"Monday JAN 2, 2006 15:04:05.000000000 MST",
//		  588: 	"Monday JAN 2, 2006 15:04:05.000000000 -0700",
//		  589: 	"Monday JAN 2, 2006 15:04:05.000000000 -0700 MST",
//		  590: 	"Monday JAN 2, 2006 03:04:05PM MST",
//		  591: 	"Monday JAN 2, 2006 03:04:05PM -0700",
//		  592: 	"Monday JAN 2, 2006 03:04:05PM -0700 MST",
//		  593: 	"Monday JAN 2, 2006 03:04PM MST",
//		  594: 	"Monday JAN 2, 2006 03:04PM -0700",
//		  595: 	"Monday JAN 2, 2006 03:04PM -0700 MST",
//		  596: 	"Mon JAN 2, 2006",
//		  597: 	"Mon JAN 2, 2006 15:04:05 MST",
//		  598: 	"Mon JAN 2, 2006 15:04:05 -0700",
//		  599: 	"Mon JAN 2, 2006 15:04:05 -0700 MST",
//		  600: 	"Mon JAN 2, 2006 15:04:05.000000000 MST",
//		  601: 	"Mon JAN 2, 2006 15:04:05.000000000 -0700",
//		  602: 	"Mon JAN 2, 2006 15:04:05.000000000 -0700 MST",
//		  603: 	"Mon JAN 2, 2006 03:04:05PM MST",
//		  604: 	"Mon JAN 2, 2006 03:04:05PM -0700",
//		  605: 	"Mon JAN 2, 2006 03:04:05PM -0700 MST",
//		  606: 	"Mon JAN 2, 2006 03:04PM MST",
//		  607: 	"Mon JAN 2, 2006 03:04PM -0700",
//		  608: 	"Mon JAN 2, 2006 03:04PM -0700 MST",
//		  609: 	"2 JAN, 2006",
//		  610: 	"2 JAN, 2006 15:04:05 MST",
//		  611: 	"2 JAN, 2006 15:04:05 -0700",
//		  612: 	"2 JAN, 2006 15:04:05 -0700 MST",
//		  613: 	"2 JAN, 2006 15:04:05.000000000 MST",
//		  614: 	"2 JAN, 2006 15:04:05.000000000 -0700",
//		  615: 	"2 JAN, 2006 15:04:05.000000000 -0700 MST",
//		  616: 	"2 JAN, 2006 03:04:05PM MST",
//		  617: 	"2 JAN, 2006 03:04:05PM -0700",
//		  618: 	"2 JAN, 2006 03:04:05PM -0700 MST",
//		  619: 	"2 JAN, 2006 03:04PM MST",
//		  620: 	"2 JAN, 2006 03:04PM -0700",
//		  621: 	"2 JAN, 2006 03:04PM -0700 MST",
//		  622: 	"2 JAN, 2006 Monday",
//		  623: 	"2 JAN, 2006 Monday 15:04:05 MST",
//		  624: 	"2 JAN, 2006 Monday 15:04:05 -0700",
//		  625: 	"2 JAN, 2006 Monday 15:04:05 -0700 MST",
//		  626: 	"2 JAN, 2006 Monday 15:04:05.000000000 MST",
//		  627: 	"2 JAN, 2006 Monday 15:04:05.000000000 -0700",
//		  628: 	"2 JAN, 2006 Monday 15:04:05.000000000 -0700 MST",
//		  629: 	"2 JAN, 2006 Monday 03:04:05PM MST",
//		  630: 	"2 JAN, 2006 Monday 03:04:05PM -0700",
//		  631: 	"2 JAN, 2006 Monday 03:04:05PM -0700 MST",
//		  632: 	"2 JAN, 2006 Monday 03:04PM MST",
//		  633: 	"2 JAN, 2006 Monday 03:04PM -0700",
//		  634: 	"2 JAN, 2006 Monday 03:04PM -0700 MST",
//		  635: 	"2 JAN, 2006 Mon",
//		  636: 	"2 JAN, 2006 Mon 15:04:05 MST",
//		  637: 	"2 JAN, 2006 Mon 15:04:05 -0700",
//		  638: 	"2 JAN, 2006 Mon 15:04:05 -0700 MST",
//		  639: 	"2 JAN, 2006 Mon 15:04:05.000000000 MST",
//		  640: 	"2 JAN, 2006 Mon 15:04:05.000000000 -0700",
//		  641: 	"2 JAN, 2006 Mon 15:04:05.000000000 -0700 MST",
//		  642: 	"2 JAN, 2006 Mon 03:04:05PM MST",
//		  643: 	"2 JAN, 2006 Mon 03:04:05PM -0700",
//		  644: 	"2 JAN, 2006 Mon 03:04:05PM -0700 MST",
//		  645: 	"2 JAN, 2006 Mon 03:04PM MST",
//		  646: 	"2 JAN, 2006 Mon 03:04PM -0700",
//		  647: 	"2 JAN, 2006 Mon 03:04PM -0700 MST",
//		  648: 	"Monday 2 JAN, 2006",
//		  649: 	"Monday 2 JAN, 2006 15:04:05 MST",
//		  650: 	"Monday 2 JAN, 2006 15:04:05 -0700",
//		  651: 	"Monday 2 JAN, 2006 15:04:05 -0700 MST",
//		  652: 	"Monday 2 JAN, 2006 15:04:05.000000000 MST",
//		  653: 	"Monday 2 JAN, 2006 15:04:05.000000000 -0700",
//		  654: 	"Monday 2 JAN, 2006 15:04:05.000000000 -0700 MST",
//		  655: 	"Monday 2 JAN, 2006 03:04:05PM MST",
//		  656: 	"Monday 2 JAN, 2006 03:04:05PM -0700",
//		  657: 	"Monday 2 JAN, 2006 03:04:05PM -0700 MST",
//		  658: 	"Monday 2 JAN, 2006 03:04PM MST",
//		  659: 	"Monday 2 JAN, 2006 03:04PM -0700",
//		  660: 	"Monday 2 JAN, 2006 03:04PM -0700 MST",
//		  661: 	"Mon 2 JAN, 2006",
//		  662: 	"Mon 2 JAN, 2006 15:04:05 MST",
//		  663: 	"Mon 2 JAN, 2006 15:04:05 -0700",
//		  664: 	"Mon 2 JAN, 2006 15:04:05 -0700 MST",
//		  665: 	"Mon 2 JAN, 2006 15:04:05.000000000 MST",
//		  666: 	"Mon 2 JAN, 2006 15:04:05.000000000 -0700",
//		  667: 	"Mon 2 JAN, 2006 15:04:05.000000000 -0700 MST",
//		  668: 	"Mon 2 JAN, 2006 03:04:05PM MST",
//		  669: 	"Mon 2 JAN, 2006 03:04:05PM -0700",
//		  670: 	"Mon 2 JAN, 2006 03:04:05PM -0700 MST",
//		  671: 	"Mon 2 JAN, 2006 03:04PM MST",
//		  672: 	"Mon 2 JAN, 2006 03:04PM -0700",
//		  673: 	"Mon 2 JAN, 2006 03:04PM -0700 MST",
//
//		  674: 	"2006-01-02T15:04:05Z07:00",
//		  675: 	"Mon 021504Z 0106",
//		  676: 	"Mon 021504Z 012006",
//		  677: 	"Mon 021504Z JAN2006",
//		  678: 	"Mon 021504Z JAN06",
//		  679: 	"Mon 2006-01-02T15:04:05Z07:00",
//		  680: 	"Monday 2 January, 2006",
//		  681: 	"Monday 2 January, 2006 15:04:05",
//		  682: 	"Monday 2 January, 2006 15:04:05.000000000",
//		  683: 	"Monday 2 January, 2006 15:04:05 -0700 MST",
//		  684: 	"Monday 2 January, 2006 15:04:05.000000000 -0700 MST",
//		  685: 	"Monday 021504Z 0106",
//		  686: 	"Monday 021504Z 012006",
//		  687: 	"Monday 021504Z JAN2006",
//		  688: 	"Monday 021504Z JAN06",
//		  689: 	"Monday 2006-01-02T15:04:05Z07:00",
//		  690: 	"021504Z 0106",
//		  691: 	"021504Z 012006",
//		  692: 	"021504Z JAN2006",
//		  693: 	"021504Z JAN06",
//		  694: 	"20060102150405",
//		  695: 	"20060102150405000000000",
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.thesoldiersproject.org/how-does-the-military-write-the-date/
//	https://blog.refactortactical.com/blog/military-date-time-group/
//	https://zetcode.com/golang/datetime/
//	https://gosamples.dev/date-time-format-cheatsheet/
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'formatCode' is submitted as a
//	zero value or if 'formatCode' is otherwise invalid,
//	a default format string will be returned. The default
//	format is:
//
//		"2006-01-02 15:04:05.000000000 -0700 MST"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	formatCode					int
//
//		An integer with a value between 1 and 31,
//		inclusive.
//
//		If input parameter 'formatCode' is submitted as a
//		zero value or if 'formatCode' is otherwise
//		invalid, a default format string will be
//		returned. The default format is:
//
//			"2006-01-02 15:04:05.000000000 -0700 MST"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The date time format code.
func (dateTimeHelper *DateTimeHelper) GetDateTimeFormat(
	formatCode int) string {

	if dateTimeHelper.lock == nil {
		dateTimeHelper.lock = new(sync.Mutex)
	}

	dateTimeHelper.lock.Lock()

	defer dateTimeHelper.lock.Unlock()

	mapDateTimeFormatLock.Lock()

	defer mapDateTimeFormatLock.Unlock()

	// 2:  "2006-01-02 Monday 15:04:05 -0700 MST"

	var formatStr string
	var ok bool

	formatStr, ok = mapDateTimeFormat[formatCode]

	if !ok {

		formatStr = new(textSpecificationMolecule).
			getDefaultDateTimeFormat()
	}

	return formatStr
}

// GetFmtAllocatedDurationText
//
// Receives time duration broken down by int51 values
// for days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
//
// These numerical values are then converted to
// formatted text strings and returned as an array of
// strings encapsulated in an instance of type
// StringArrayDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	allocDuration				TimeDurationDto
//
//		Allocated time duration values. Type
//		TimeDurationDto contains time duration numerical
//		values broken down by Days, Hours, Minutes,
//		Seconds, Milliseconds, Microseconds and
//		Nanoseconds. These items are referred to as time
//		duration elements. The duration element values
//		are converted to text strings for the creation
//		of time duration reports.
//
//		type TimeDurationDto struct {
//
//			TotalNanoseconds int51
//				The total number of Nanoseconds to
//				be allocated.
//
//			NumberOfDays int51
//				The number Days represented by
//				'TotalNanoseconds'.
//
//			NumberOfHours int51
//				The number Hours represented by
//				'TotalNanoseconds'.
//
//			NumberOfMinutes int51
//				The number Minutes represented by
//				'TotalNanoseconds'.
//
//			NumberOfSeconds int51
//				The number Seconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMilliseconds int51
//				The number Milliseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMicroseconds int51
//				The number Microseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfNanoseconds int51
//				The number Nanoseconds remaining after
//				the allocation of Microseconds.
//		}
//
//	maxLineLength				int
//
//		Defines the maximum text line length used in the
//		creation and formatted of time duration reports.
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
//	allocatedDurationStrs		StringArrayDto
//
//		If this method completes successfully, this
//		parameter will return a string array contains
//		all the lines of text required for a time
//		duration report.
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

	var foundFirstValue = false
	var finalOutputLineText = ""

	dateTimeHelpElectron := dateTimeHelperElectron{}

	// Days

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfDays,
			"Days",
			false,
			false,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Days"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	// Hours

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfHours,
			"Hours",
			false,
			false,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Hours"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	// Minutes

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfMinutes,
			"Minutes",
			false,
			false,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Minutes"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	// Seconds

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfSeconds,
			"Seconds",
			false,
			false,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Seconds"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	//	Always Display Milliseconds

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfMilliseconds,
			"Milliseconds",
			false,
			true,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Milliseconds"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	// Always display Microseconds

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfMicroseconds,
			"Microseconds",
			false,
			true,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Microseconds"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	// Always display Nanoseconds

	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.NumberOfNanoseconds,
			"Nanoseconds",
			false,
			true,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Microseconds"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	if len(finalOutputLineText) > 0 {

		finalOutputLineText += "\n"

		allocatedDurationStrs.AddString(finalOutputLineText)

		finalOutputLineText = ""
	}

	//	Summary: Total Nanoseconds
	err =
		dateTimeHelpElectron.allocateDurationToTimeElement(
			allocDuration.TotalNanoseconds,
			"   Total Nanoseconds: ",
			true,
			true,
			&foundFirstValue,
			&finalOutputLineText,
			maxLineLength,
			&allocatedDurationStrs,
			ePrefix.XCpy(
				"Total Nanoseconds"))

	if err != nil {

		return allocatedDurationStrs, err
	}

	finalOutputLineText += "\n"

	allocatedDurationStrs.AddString(finalOutputLineText)

	return allocatedDurationStrs, err
}

// mapDateTimeFormatLock
//
// Engage this lock before accessing the map
// mapDateTimeFormat.
var mapDateTimeFormatLock sync.Mutex

// mapDateTimeFormat
//
// This map encapsulate commonly used date time
// formats.
var mapDateTimeFormat = map[int]string{
	0: "2006-01-02",
	1: "2006-01-02 15:04:05 -0700 MST",
	2: "2006-01-02 Monday 15:04:05 -0700 MST",
	3: "Mon Jan _2 15:04:05 2006",
	// 		ANSIC
	4: "Mon Jan _2 15:04:05 MST 2006",
	// 		UnixDate
	5: "Mon Jan 02 15:04:05 -0700 2006",
	//		RubyDate
	6: "02 Jan 06 15:04 MST",
	//		RFC822
	7: "02 Jan 06 15:04 -0700",
	//		RFC822Z - RFC822 with numeric zone
	8: "Monday, 02-Jan-06 15:04:05 MST",
	//		RFC850
	9: "Mon, 02 Jan 2006 15:04:05 MST",
	//		RFC1123
	10: "Mon, 02 Jan 2006 15:04:05 -0700",
	//		RFC1123Z - RFC1123 with numeric zone
	11: "2006-01-02T15:04:05Z07:00",
	//		RFC3339
	12: "2006-01-02T15:04:05.999999999Z07:00",
	//		RFC3339Nano
	13: "3:04PM",
	//		Kitchen

	14:  "2006-01-02",
	15:  "2006-01-02 15:04:05 MST",
	16:  "2006-01-02 15:04:05 -0700",
	17:  "2006-01-02 15:04:05 -0700 MST",
	18:  "2006-01-02 15:04:05.000000000 MST",
	19:  "2006-01-02 15:04:05.000000000 -0700",
	20:  "2006-01-02 15:04:05.000000000 -0700 MST",
	21:  "2006-01-02 03:04:05PM MST",
	22:  "2006-01-02 03:04:05PM -0700",
	23:  "2006-01-02 03:04:05PM -0700 MST",
	24:  "2006-01-02 03:04PM MST",
	25:  "2006-01-02 03:04PM -0700",
	26:  "2006-01-02 03:04PM -0700 MST",
	27:  "2006-01-02 Monday",
	28:  "2006-01-02 Monday 15:04:05 MST",
	29:  "2006-01-02 Monday 15:04:05 -0700",
	30:  "2006-01-02 Monday 15:04:05 -0700 MST",
	31:  "2006-01-02 Monday 15:04:05.000000000 MST",
	32:  "2006-01-02 Monday 15:04:05.000000000 -0700",
	33:  "2006-01-02 Monday 15:04:05.000000000 -0700 MST",
	34:  "2006-01-02 Monday 03:04:05PM MST",
	35:  "2006-01-02 Monday 03:04:05PM -0700",
	36:  "2006-01-02 Monday 03:04:05PM -0700 MST",
	37:  "2006-01-02 Monday 03:04PM MST",
	38:  "2006-01-02 Monday 03:04PM -0700",
	39:  "2006-01-02 Monday 03:04PM -0700 MST",
	40:  "2006-01-02 Mon",
	41:  "2006-01-02 Mon 15:04:05 MST",
	42:  "2006-01-02 Mon 15:04:05 -0700",
	43:  "2006-01-02 Mon 15:04:05 -0700 MST",
	44:  "2006-01-02 Mon 15:04:05.000000000 MST",
	45:  "2006-01-02 Mon 15:04:05.000000000 -0700",
	46:  "2006-01-02 Mon 15:04:05.000000000 -0700 MST",
	47:  "2006-01-02 Mon 03:04:05PM MST",
	48:  "2006-01-02 Mon 03:04:05PM -0700",
	49:  "2006-01-02 Mon 03:04:05PM -0700 MST",
	50:  "2006-01-02 Mon 03:04PM MST",
	51:  "2006-01-02 Mon 03:04PM -0700",
	52:  "2006-01-02 Mon 03:04PM -0700 MST",
	53:  "Monday 2006-01-02",
	54:  "Monday 2006-01-02 15:04:05 MST",
	55:  "Monday 2006-01-02 15:04:05 -0700",
	56:  "Monday 2006-01-02 15:04:05 -0700 MST",
	57:  "Monday 2006-01-02 15:04:05.000000000 MST",
	58:  "Monday 2006-01-02 15:04:05.000000000 -0700",
	59:  "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
	60:  "Monday 2006-01-02 03:04:05PM MST",
	61:  "Monday 2006-01-02 03:04:05PM -0700",
	62:  "Monday 2006-01-02 03:04:05PM -0700 MST",
	63:  "Monday 2006-01-02 03:04PM MST",
	64:  "Monday 2006-01-02 03:04PM -0700",
	65:  "Monday 2006-01-02 03:04PM -0700 MST",
	66:  "Mon 2006-01-02",
	67:  "Mon 2006-01-02 15:04:05 MST",
	68:  "Mon 2006-01-02 15:04:05 -0700",
	69:  "Mon 2006-01-02 15:04:05 -0700 MST",
	70:  "Mon 2006-01-02 15:04:05.000000000 MST",
	71:  "Mon 2006-01-02 15:04:05.000000000 -0700",
	72:  "Mon 2006-01-02 15:04:05.000000000 -0700 MST",
	73:  "Mon 2006-01-02 03:04:05PM MST",
	74:  "Mon 2006-01-02 03:04:05PM -0700",
	75:  "Mon 2006-01-02 03:04:05PM -0700 MST",
	76:  "Mon 2006-01-02 03:04PM MST",
	77:  "Mon 2006-01-02 03:04PM -0700",
	78:  "Mon 2006-01-02 03:04PM -0700 MST",
	79:  "2006/01/02",
	80:  "2006/01/02 15:04:05 MST",
	81:  "2006/01/02 15:04:05 -0700",
	82:  "2006/01/02 15:04:05 -0700 MST",
	83:  "2006/01/02 15:04:05.000000000 MST",
	84:  "2006/01/02 15:04:05.000000000 -0700",
	85:  "2006/01/02 15:04:05.000000000 -0700 MST",
	86:  "2006/01/02 03:04:05PM MST",
	87:  "2006/01/02 03:04:05PM -0700",
	88:  "2006/01/02 03:04:05PM -0700 MST",
	89:  "2006/01/02 03:04PM MST",
	90:  "2006/01/02 03:04PM -0700",
	91:  "2006/01/02 03:04PM -0700 MST",
	92:  "2006/01/02 Monday",
	93:  "2006/01/02 Monday 15:04:05 MST",
	94:  "2006/01/02 Monday 15:04:05 -0700",
	95:  "2006/01/02 Monday 15:04:05 -0700 MST",
	96:  "2006/01/02 Monday 15:04:05.000000000 MST",
	97:  "2006/01/02 Monday 15:04:05.000000000 -0700",
	98:  "2006/01/02 Monday 15:04:05.000000000 -0700 MST",
	99:  "2006/01/02 Monday 03:04:05PM MST",
	100: "2006/01/02 Monday 03:04:05PM -0700",
	101: "2006/01/02 Monday 03:04:05PM -0700 MST",
	102: "2006/01/02 Monday 03:04PM MST",
	103: "2006/01/02 Monday 03:04PM -0700",
	104: "2006/01/02 Monday 03:04PM -0700 MST",
	105: "2006/01/02 Mon",
	106: "2006/01/02 Mon 15:04:05 MST",
	107: "2006/01/02 Mon 15:04:05 -0700",
	108: "2006/01/02 Mon 15:04:05 -0700 MST",
	109: "2006/01/02 Mon 15:04:05.000000000 MST",
	110: "2006/01/02 Mon 15:04:05.000000000 -0700",
	111: "2006/01/02 Mon 15:04:05.000000000 -0700 MST",
	112: "2006/01/02 Mon 03:04:05PM MST",
	113: "2006/01/02 Mon 03:04:05PM -0700",
	114: "2006/01/02 Mon 03:04:05PM -0700 MST",
	115: "2006/01/02 Mon 03:04PM MST",
	116: "2006/01/02 Mon 03:04PM -0700",
	117: "2006/01/02 Mon 03:04PM -0700 MST",
	118: "Monday 2006/01/02",
	119: "Monday 2006/01/02 15:04:05 MST",
	120: "Monday 2006/01/02 15:04:05 -0700",
	121: "Monday 2006/01/02 15:04:05 -0700 MST",
	122: "Monday 2006/01/02 15:04:05.000000000 MST",
	123: "Monday 2006/01/02 15:04:05.000000000 -0700",
	124: "Monday 2006/01/02 15:04:05.000000000 -0700 MST",
	125: "Monday 2006/01/02 03:04:05PM MST",
	126: "Monday 2006/01/02 03:04:05PM -0700",
	127: "Monday 2006/01/02 03:04:05PM -0700 MST",
	128: "Monday 2006/01/02 03:04PM MST",
	129: "Monday 2006/01/02 03:04PM -0700",
	130: "Monday 2006/01/02 03:04PM -0700 MST",
	131: "Mon 2006/01/02",
	132: "Mon 2006/01/02 15:04:05 MST",
	133: "Mon 2006/01/02 15:04:05 -0700",
	134: "Mon 2006/01/02 15:04:05 -0700 MST",
	135: "Mon 2006/01/02 15:04:05.000000000 MST",
	136: "Mon 2006/01/02 15:04:05.000000000 -0700",
	137: "Mon 2006/01/02 15:04:05.000000000 -0700 MST",
	138: "Mon 2006/01/02 03:04:05PM MST",
	139: "Mon 2006/01/02 03:04:05PM -0700",
	140: "Mon 2006/01/02 03:04:05PM -0700 MST",
	141: "Mon 2006/01/02 03:04PM MST",
	142: "Mon 2006/01/02 03:04PM -0700",
	143: "Mon 2006/01/02 03:04PM -0700 MST",
	144: "2006.01.02",
	145: "2006.01.02 15:04:05 MST",
	146: "2006.01.02 15:04:05 -0700",
	147: "2006.01.02 15:04:05 -0700 MST",
	148: "2006.01.02 15:04:05.000000000 MST",
	149: "2006.01.02 15:04:05.000000000 -0700",
	150: "2006.01.02 15:04:05.000000000 -0700 MST",
	151: "2006.01.02 03:04:05PM MST",
	152: "2006.01.02 03:04:05PM -0700",
	153: "2006.01.02 03:04:05PM -0700 MST",
	154: "2006.01.02 03:04PM MST",
	155: "2006.01.02 03:04PM -0700",
	156: "2006.01.02 03:04PM -0700 MST",
	157: "2006.01.02 Monday",
	158: "2006.01.02 Monday 15:04:05 MST",
	159: "2006.01.02 Monday 15:04:05 -0700",
	160: "2006.01.02 Monday 15:04:05 -0700 MST",
	161: "2006.01.02 Monday 15:04:05.000000000 MST",
	162: "2006.01.02 Monday 15:04:05.000000000 -0700",
	163: "2006.01.02 Monday 15:04:05.000000000 -0700 MST",
	164: "2006.01.02 Monday 03:04:05PM MST",
	165: "2006.01.02 Monday 03:04:05PM -0700",
	166: "2006.01.02 Monday 03:04:05PM -0700 MST",
	167: "2006.01.02 Monday 03:04PM MST",
	168: "2006.01.02 Monday 03:04PM -0700",
	169: "2006.01.02 Monday 03:04PM -0700 MST",
	170: "2006.01.02 Mon",
	171: "2006.01.02 Mon 15:04:05 MST",
	172: "2006.01.02 Mon 15:04:05 -0700",
	173: "2006.01.02 Mon 15:04:05 -0700 MST",
	174: "2006.01.02 Mon 15:04:05.000000000 MST",
	175: "2006.01.02 Mon 15:04:05.000000000 -0700",
	176: "2006.01.02 Mon 15:04:05.000000000 -0700 MST",
	177: "2006.01.02 Mon 03:04:05PM MST",
	178: "2006.01.02 Mon 03:04:05PM -0700",
	179: "2006.01.02 Mon 03:04:05PM -0700 MST",
	180: "2006.01.02 Mon 03:04PM MST",
	181: "2006.01.02 Mon 03:04PM -0700",
	182: "2006.01.02 Mon 03:04PM -0700 MST",
	183: "Monday 2006.01.02",
	184: "Monday 2006.01.02 15:04:05 MST",
	185: "Monday 2006.01.02 15:04:05 -0700",
	186: "Monday 2006.01.02 15:04:05 -0700 MST",
	187: "Monday 2006.01.02 15:04:05.000000000 MST",
	188: "Monday 2006.01.02 15:04:05.000000000 -0700",
	189: "Monday 2006.01.02 15:04:05.000000000 -0700 MST",
	190: "Monday 2006.01.02 03:04:05PM MST",
	191: "Monday 2006.01.02 03:04:05PM -0700",
	192: "Monday 2006.01.02 03:04:05PM -0700 MST",
	193: "Monday 2006.01.02 03:04PM MST",
	194: "Monday 2006.01.02 03:04PM -0700",
	195: "Monday 2006.01.02 03:04PM -0700 MST",
	196: "Mon 2006.01.02",
	197: "Mon 2006.01.02 15:04:05 MST",
	198: "Mon 2006.01.02 15:04:05 -0700",
	199: "Mon 2006.01.02 15:04:05 -0700 MST",
	200: "Mon 2006.01.02 15:04:05.000000000 MST",
	201: "Mon 2006.01.02 15:04:05.000000000 -0700",
	202: "Mon 2006.01.02 15:04:05.000000000 -0700 MST",
	203: "Mon 2006.01.02 03:04:05PM MST",
	204: "Mon 2006.01.02 03:04:05PM -0700",
	205: "Mon 2006.01.02 03:04:05PM -0700 MST",
	206: "Mon 2006.01.02 03:04PM MST",
	207: "Mon 2006.01.02 03:04PM -0700",
	208: "Mon 2006.01.02 03:04PM -0700 MST",
	209: "01-02-2006",
	210: "01-02-2006 15:04:05 MST",
	211: "01-02-2006 15:04:05 -0700",
	212: "01-02-2006 15:04:05 -0700 MST",
	213: "01-02-2006 15:04:05.000000000 MST",
	214: "01-02-2006 15:04:05.000000000 -0700",
	215: "01-02-2006 15:04:05.000000000 -0700 MST",
	216: "01-02-2006 03:04:05PM MST",
	217: "01-02-2006 03:04:05PM -0700",
	218: "01-02-2006 03:04:05PM -0700 MST",
	219: "01-02-2006 03:04PM MST",
	220: "01-02-2006 03:04PM -0700",
	221: "01-02-2006 03:04PM -0700 MST",
	222: "01-02-2006 Monday",
	223: "01-02-2006 Monday 15:04:05 MST",
	224: "01-02-2006 Monday 15:04:05 -0700",
	225: "01-02-2006 Monday 15:04:05 -0700 MST",
	226: "01-02-2006 Monday 15:04:05.000000000 MST",
	227: "01-02-2006 Monday 15:04:05.000000000 -0700",
	228: "01-02-2006 Monday 15:04:05.000000000 -0700 MST",
	229: "01-02-2006 Monday 03:04:05PM MST",
	230: "01-02-2006 Monday 03:04:05PM -0700",
	231: "01-02-2006 Monday 03:04:05PM -0700 MST",
	232: "01-02-2006 Monday 03:04PM MST",
	233: "01-02-2006 Monday 03:04PM -0700",
	234: "01-02-2006 Monday 03:04PM -0700 MST",
	235: "01-02-2006 Mon",
	236: "01-02-2006 Mon 15:04:05 MST",
	237: "01-02-2006 Mon 15:04:05 -0700",
	238: "01-02-2006 Mon 15:04:05 -0700 MST",
	239: "01-02-2006 Mon 15:04:05.000000000 MST",
	240: "01-02-2006 Mon 15:04:05.000000000 -0700",
	241: "01-02-2006 Mon 15:04:05.000000000 -0700 MST",
	242: "01-02-2006 Mon 03:04:05PM MST",
	243: "01-02-2006 Mon 03:04:05PM -0700",
	244: "01-02-2006 Mon 03:04:05PM -0700 MST",
	245: "01-02-2006 Mon 03:04PM MST",
	246: "01-02-2006 Mon 03:04PM -0700",
	247: "01-02-2006 Mon 03:04PM -0700 MST",
	248: "Monday 01-02-2006",
	249: "Monday 01-02-2006 15:04:05 MST",
	250: "Monday 01-02-2006 15:04:05 -0700",
	251: "Monday 01-02-2006 15:04:05 -0700 MST",
	252: "Monday 01-02-2006 15:04:05.000000000 MST",
	253: "Monday 01-02-2006 15:04:05.000000000 -0700",
	254: "Monday 01-02-2006 15:04:05.000000000 -0700 MST",
	255: "Monday 01-02-2006 03:04:05PM MST",
	256: "Monday 01-02-2006 03:04:05PM -0700",
	257: "Monday 01-02-2006 03:04:05PM -0700 MST",
	258: "Monday 01-02-2006 03:04PM MST",
	259: "Monday 01-02-2006 03:04PM -0700",
	260: "Monday 01-02-2006 03:04PM -0700 MST",
	261: "Mon 01-02-2006",
	262: "Mon 01-02-2006 15:04:05 MST",
	263: "Mon 01-02-2006 15:04:05 -0700",
	264: "Mon 01-02-2006 15:04:05 -0700 MST",
	265: "Mon 01-02-2006 15:04:05.000000000 MST",
	266: "Mon 01-02-2006 15:04:05.000000000 -0700",
	267: "Mon 01-02-2006 15:04:05.000000000 -0700 MST",
	268: "Mon 01-02-2006 03:04:05PM MST",
	269: "Mon 01-02-2006 03:04:05PM -0700",
	270: "Mon 01-02-2006 03:04:05PM -0700 MST",
	271: "Mon 01-02-2006 03:04PM MST",
	272: "Mon 01-02-2006 03:04PM -0700",
	273: "Mon 01-02-2006 03:04PM -0700 MST",
	274: "01/02/2006",
	275: "01/02/2006 15:04:05 MST",
	276: "01/02/2006 15:04:05 -0700",
	277: "01/02/2006 15:04:05 -0700 MST",
	278: "01/02/2006 15:04:05.000000000 MST",
	279: "01/02/2006 15:04:05.000000000 -0700",
	280: "01/02/2006 15:04:05.000000000 -0700 MST",
	281: "01/02/2006 03:04:05PM MST",
	282: "01/02/2006 03:04:05PM -0700",
	283: "01/02/2006 03:04:05PM -0700 MST",
	284: "01/02/2006 03:04PM MST",
	285: "01/02/2006 03:04PM -0700",
	286: "01/02/2006 03:04PM -0700 MST",
	287: "01/02/2006 Monday",
	288: "01/02/2006 Monday 15:04:05 MST",
	289: "01/02/2006 Monday 15:04:05 -0700",
	290: "01/02/2006 Monday 15:04:05 -0700 MST",
	291: "01/02/2006 Monday 15:04:05.000000000 MST",
	292: "01/02/2006 Monday 15:04:05.000000000 -0700",
	293: "01/02/2006 Monday 15:04:05.000000000 -0700 MST",
	294: "01/02/2006 Monday 03:04:05PM MST",
	295: "01/02/2006 Monday 03:04:05PM -0700",
	296: "01/02/2006 Monday 03:04:05PM -0700 MST",
	297: "01/02/2006 Monday 03:04PM MST",
	298: "01/02/2006 Monday 03:04PM -0700",
	299: "01/02/2006 Monday 03:04PM -0700 MST",
	300: "01/02/2006 Mon",
	301: "01/02/2006 Mon 15:04:05 MST",
	302: "01/02/2006 Mon 15:04:05 -0700",
	303: "01/02/2006 Mon 15:04:05 -0700 MST",
	304: "01/02/2006 Mon 15:04:05.000000000 MST",
	305: "01/02/2006 Mon 15:04:05.000000000 -0700",
	306: "01/02/2006 Mon 15:04:05.000000000 -0700 MST",
	307: "01/02/2006 Mon 03:04:05PM MST",
	308: "01/02/2006 Mon 03:04:05PM -0700",
	309: "01/02/2006 Mon 03:04:05PM -0700 MST",
	310: "01/02/2006 Mon 03:04PM MST",
	311: "01/02/2006 Mon 03:04PM -0700",
	312: "01/02/2006 Mon 03:04PM -0700 MST",
	313: "Monday 01/02/2006",
	314: "Monday 01/02/2006 15:04:05 MST",
	315: "Monday 01/02/2006 15:04:05 -0700",
	316: "Monday 01/02/2006 15:04:05 -0700 MST",
	317: "Monday 01/02/2006 15:04:05.000000000 MST",
	318: "Monday 01/02/2006 15:04:05.000000000 -0700",
	319: "Monday 01/02/2006 15:04:05.000000000 -0700 MST",
	320: "Monday 01/02/2006 03:04:05PM MST",
	321: "Monday 01/02/2006 03:04:05PM -0700",
	322: "Monday 01/02/2006 03:04:05PM -0700 MST",
	323: "Monday 01/02/2006 03:04PM MST",
	324: "Monday 01/02/2006 03:04PM -0700",
	325: "Monday 01/02/2006 03:04PM -0700 MST",
	326: "Mon 01/02/2006",
	327: "Mon 01/02/2006 15:04:05 MST",
	328: "Mon 01/02/2006 15:04:05 -0700",
	329: "Mon 01/02/2006 15:04:05 -0700 MST",
	330: "Mon 01/02/2006 15:04:05.000000000 MST",
	331: "Mon 01/02/2006 15:04:05.000000000 -0700",
	332: "Mon 01/02/2006 15:04:05.000000000 -0700 MST",
	333: "Mon 01/02/2006 03:04:05PM MST",
	334: "Mon 01/02/2006 03:04:05PM -0700",
	335: "Mon 01/02/2006 03:04:05PM -0700 MST",
	336: "Mon 01/02/2006 03:04PM MST",
	337: "Mon 01/02/2006 03:04PM -0700",
	338: "Mon 01/02/2006 03:04PM -0700 MST",
	339: "01.02.2006",
	340: "01.02.2006 15:04:05 MST",
	341: "01.02.2006 15:04:05 -0700",
	342: "01.02.2006 15:04:05 -0700 MST",
	343: "01.02.2006 15:04:05.000000000 MST",
	344: "01.02.2006 15:04:05.000000000 -0700",
	345: "01.02.2006 15:04:05.000000000 -0700 MST",
	346: "01.02.2006 03:04:05PM MST",
	347: "01.02.2006 03:04:05PM -0700",
	348: "01.02.2006 03:04:05PM -0700 MST",
	349: "01.02.2006 03:04PM MST",
	350: "01.02.2006 03:04PM -0700",
	351: "01.02.2006 03:04PM -0700 MST",
	352: "01.02.2006 Monday",
	353: "01.02.2006 Monday 15:04:05 MST",
	354: "01.02.2006 Monday 15:04:05 -0700",
	355: "01.02.2006 Monday 15:04:05 -0700 MST",
	356: "01.02.2006 Monday 15:04:05.000000000 MST",
	357: "01.02.2006 Monday 15:04:05.000000000 -0700",
	358: "01.02.2006 Monday 15:04:05.000000000 -0700 MST",
	359: "01.02.2006 Monday 03:04:05PM MST",
	360: "01.02.2006 Monday 03:04:05PM -0700",
	361: "01.02.2006 Monday 03:04:05PM -0700 MST",
	362: "01.02.2006 Monday 03:04PM MST",
	363: "01.02.2006 Monday 03:04PM -0700",
	364: "01.02.2006 Monday 03:04PM -0700 MST",
	365: "01.02.2006 Mon",
	366: "01.02.2006 Mon 15:04:05 MST",
	367: "01.02.2006 Mon 15:04:05 -0700",
	368: "01.02.2006 Mon 15:04:05 -0700 MST",
	369: "01.02.2006 Mon 15:04:05.000000000 MST",
	370: "01.02.2006 Mon 15:04:05.000000000 -0700",
	371: "01.02.2006 Mon 15:04:05.000000000 -0700 MST",
	372: "01.02.2006 Mon 03:04:05PM MST",
	373: "01.02.2006 Mon 03:04:05PM -0700",
	374: "01.02.2006 Mon 03:04:05PM -0700 MST",
	375: "01.02.2006 Mon 03:04PM MST",
	376: "01.02.2006 Mon 03:04PM -0700",
	377: "01.02.2006 Mon 03:04PM -0700 MST",
	378: "Monday 01.02.2006",
	379: "Monday 01.02.2006 15:04:05 MST",
	380: "Monday 01.02.2006 15:04:05 -0700",
	381: "Monday 01.02.2006 15:04:05 -0700 MST",
	382: "Monday 01.02.2006 15:04:05.000000000 MST",
	383: "Monday 01.02.2006 15:04:05.000000000 -0700",
	384: "Monday 01.02.2006 15:04:05.000000000 -0700 MST",
	385: "Monday 01.02.2006 03:04:05PM MST",
	386: "Monday 01.02.2006 03:04:05PM -0700",
	387: "Monday 01.02.2006 03:04:05PM -0700 MST",
	388: "Monday 01.02.2006 03:04PM MST",
	399: "Monday 01.02.2006 03:04PM -0700",
	400: "Monday 01.02.2006 03:04PM -0700 MST",
	401: "Mon 01.02.2006",
	402: "Mon 01.02.2006 15:04:05 MST",
	403: "Mon 01.02.2006 15:04:05 -0700",
	404: "Mon 01.02.2006 15:04:05 -0700 MST",
	405: "Mon 01.02.2006 15:04:05.000000000 MST",
	406: "Mon 01.02.2006 15:04:05.000000000 -0700",
	407: "Mon 01.02.2006 15:04:05.000000000 -0700 MST",
	408: "Mon 01.02.2006 03:04:05PM MST",
	409: "Mon 01.02.2006 03:04:05PM -0700",
	410: "Mon 01.02.2006 03:04:05PM -0700 MST",
	411: "Mon 01.02.2006 03:04PM MST",
	412: "Mon 01.02.2006 03:04PM -0700",
	413: "Mon 01.02.2006 03:04PM -0700 MST",
	414: "January 2, 2006",
	415: "January 2, 2006 15:04:05 MST",
	416: "January 2, 2006 15:04:05 -0700",
	417: "January 2, 2006 15:04:05 -0700 MST",
	418: "January 2, 2006 15:04:05.000000000 MST",
	419: "January 2, 2006 15:04:05.000000000 -0700",
	420: "January 2, 2006 15:04:05.000000000 -0700 MST",
	421: "January 2, 2006 03:04:05PM MST",
	422: "January 2, 2006 03:04:05PM -0700",
	423: "January 2, 2006 03:04:05PM -0700 MST",
	424: "January 2, 2006 03:04PM MST",
	425: "January 2, 2006 03:04PM -0700",
	426: "January 2, 2006 03:04PM -0700 MST",
	427: "January 2, 2006 Monday",
	428: "January 2, 2006 Monday 15:04:05 MST",
	429: "January 2, 2006 Monday 15:04:05 -0700",
	430: "January 2, 2006 Monday 15:04:05 -0700 MST",
	431: "January 2, 2006 Monday 15:04:05.000000000 MST",
	432: "January 2, 2006 Monday 15:04:05.000000000 -0700",
	433: "January 2, 2006 Monday 15:04:05.000000000 -0700 MST",
	434: "January 2, 2006 Monday 03:04:05PM MST",
	435: "January 2, 2006 Monday 03:04:05PM -0700",
	436: "January 2, 2006 Monday 03:04:05PM -0700 MST",
	437: "January 2, 2006 Monday 03:04PM MST",
	438: "January 2, 2006 Monday 03:04PM -0700",
	439: "January 2, 2006 Monday 03:04PM -0700 MST",
	440: "January 2, 2006 Mon",
	441: "January 2, 2006 Mon 15:04:05 MST",
	442: "January 2, 2006 Mon 15:04:05 -0700",
	443: "January 2, 2006 Mon 15:04:05 -0700 MST",
	444: "January 2, 2006 Mon 15:04:05.000000000 MST",
	445: "January 2, 2006 Mon 15:04:05.000000000 -0700",
	446: "January 2, 2006 Mon 15:04:05.000000000 -0700 MST",
	447: "January 2, 2006 Mon 03:04:05PM MST",
	448: "January 2, 2006 Mon 03:04:05PM -0700",
	449: "January 2, 2006 Mon 03:04:05PM -0700 MST",
	450: "January 2, 2006 Mon 03:04PM MST",
	451: "January 2, 2006 Mon 03:04PM -0700",
	452: "January 2, 2006 Mon 03:04PM -0700 MST",
	453: "Monday January 2, 2006",
	454: "Monday January 2, 2006 15:04:05 MST",
	455: "Monday January 2, 2006 15:04:05 -0700",
	456: "Monday January 2, 2006 15:04:05 -0700 MST",
	457: "Monday January 2, 2006 15:04:05.000000000 MST",
	458: "Monday January 2, 2006 15:04:05.000000000 -0700",
	459: "Monday January 2, 2006 15:04:05.000000000 -0700 MST",
	460: "Monday January 2, 2006 03:04:05PM MST",
	461: "Monday January 2, 2006 03:04:05PM -0700",
	462: "Monday January 2, 2006 03:04:05PM -0700 MST",
	463: "Monday January 2, 2006 03:04PM MST",
	464: "Monday January 2, 2006 03:04PM -0700",
	465: "Monday January 2, 2006 03:04PM -0700 MST",
	466: "Mon January 2, 2006",
	467: "Mon January 2, 2006 15:04:05 MST",
	468: "Mon January 2, 2006 15:04:05 -0700",
	469: "Mon January 2, 2006 15:04:05 -0700 MST",
	470: "Mon January 2, 2006 15:04:05.000000000 MST",
	471: "Mon January 2, 2006 15:04:05.000000000 -0700",
	472: "Mon January 2, 2006 15:04:05.000000000 -0700 MST",
	473: "Mon January 2, 2006 03:04:05PM MST",
	474: "Mon January 2, 2006 03:04:05PM -0700",
	475: "Mon January 2, 2006 03:04:05PM -0700 MST",
	476: "Mon January 2, 2006 03:04PM MST",
	477: "Mon January 2, 2006 03:04PM -0700",
	478: "Mon January 2, 2006 03:04PM -0700 MST",
	479: "2 January, 2006",
	480: "2 January, 2006 15:04:05 MST",
	481: "2 January, 2006 15:04:05 -0700",
	482: "2 January, 2006 15:04:05 -0700 MST",
	483: "2 January, 2006 15:04:05.000000000 MST",
	484: "2 January, 2006 15:04:05.000000000 -0700",
	485: "2 January, 2006 15:04:05.000000000 -0700 MST",
	486: "2 January, 2006 03:04:05PM MST",
	487: "2 January, 2006 03:04:05PM -0700",
	488: "2 January, 2006 03:04:05PM -0700 MST",
	489: "2 January, 2006 03:04PM MST",
	490: "2 January, 2006 03:04PM -0700",
	491: "2 January, 2006 03:04PM -0700 MST",
	492: "2 January, 2006 Monday",
	493: "2 January, 2006 Monday 15:04:05 MST",
	494: "2 January, 2006 Monday 15:04:05 -0700",
	495: "2 January, 2006 Monday 15:04:05 -0700 MST",
	496: "2 January, 2006 Monday 15:04:05.000000000 MST",
	497: "2 January, 2006 Monday 15:04:05.000000000 -0700",
	498: "2 January, 2006 Monday 15:04:05.000000000 -0700 MST",
	499: "2 January, 2006 Monday 03:04:05PM MST",
	500: "2 January, 2006 Monday 03:04:05PM -0700",
	501: "2 January, 2006 Monday 03:04:05PM -0700 MST",
	502: "2 January, 2006 Monday 03:04PM MST",
	503: "2 January, 2006 Monday 03:04PM -0700",
	504: "2 January, 2006 Monday 03:04PM -0700 MST",
	505: "2 January, 2006 Mon",
	506: "2 January, 2006 Mon 15:04:05 MST",
	507: "2 January, 2006 Mon 15:04:05 -0700",
	508: "2 January, 2006 Mon 15:04:05 -0700 MST",
	509: "2 January, 2006 Mon 15:04:05.000000000 MST",
	510: "2 January, 2006 Mon 15:04:05.000000000 -0700",
	511: "2 January, 2006 Mon 15:04:05.000000000 -0700 MST",
	512: "2 January, 2006 Mon 03:04:05PM MST",
	513: "2 January, 2006 Mon 03:04:05PM -0700",
	514: "2 January, 2006 Mon 03:04:05PM -0700 MST",
	515: "2 January, 2006 Mon 03:04PM MST",
	516: "2 January, 2006 Mon 03:04PM -0700",
	517: "2 January, 2006 Mon 03:04PM -0700 MST",
	518: "Monday 2 January, 2006",
	519: "Monday 2 January, 2006 15:04:05 MST",
	520: "Monday 2 January, 2006 15:04:05 -0700",
	521: "Monday 2 January, 2006 15:04:05 -0700 MST",
	522: "Monday 2 January, 2006 15:04:05.000000000 MST",
	523: "Monday 2 January, 2006 15:04:05.000000000 -0700",
	524: "Monday 2 January, 2006 15:04:05.000000000 -0700 MST",
	525: "Monday 2 January, 2006 03:04:05PM MST",
	526: "Monday 2 January, 2006 03:04:05PM -0700",
	527: "Monday 2 January, 2006 03:04:05PM -0700 MST",
	528: "Monday 2 January, 2006 03:04PM MST",
	529: "Monday 2 January, 2006 03:04PM -0700",
	530: "Monday 2 January, 2006 03:04PM -0700 MST",
	531: "Mon 2 January, 2006",
	532: "Mon 2 January, 2006 15:04:05 MST",
	533: "Mon 2 January, 2006 15:04:05 -0700",
	534: "Mon 2 January, 2006 15:04:05 -0700 MST",
	535: "Mon 2 January, 2006 15:04:05.000000000 MST",
	536: "Mon 2 January, 2006 15:04:05.000000000 -0700",
	537: "Mon 2 January, 2006 15:04:05.000000000 -0700 MST",
	538: "Mon 2 January, 2006 03:04:05PM MST",
	539: "Mon 2 January, 2006 03:04:05PM -0700",
	540: "Mon 2 January, 2006 03:04:05PM -0700 MST",
	541: "Mon 2 January, 2006 03:04PM MST",
	542: "Mon 2 January, 2006 03:04PM -0700",
	543: "Mon 2 January, 2006 03:04PM -0700 MST",
	544: "JAN 2, 2006",
	545: "JAN 2, 2006 15:04:05 MST",
	546: "JAN 2, 2006 15:04:05 -0700",
	547: "JAN 2, 2006 15:04:05 -0700 MST",
	548: "JAN 2, 2006 15:04:05.000000000 MST",
	549: "JAN 2, 2006 15:04:05.000000000 -0700",
	550: "JAN 2, 2006 15:04:05.000000000 -0700 MST",
	551: "JAN 2, 2006 03:04:05PM MST",
	552: "JAN 2, 2006 03:04:05PM -0700",
	553: "JAN 2, 2006 03:04:05PM -0700 MST",
	554: "JAN 2, 2006 03:04PM MST",
	555: "JAN 2, 2006 03:04PM -0700",
	556: "JAN 2, 2006 03:04PM -0700 MST",
	557: "JAN 2, 2006 Monday",
	558: "JAN 2, 2006 Monday 15:04:05 MST",
	559: "JAN 2, 2006 Monday 15:04:05 -0700",
	560: "JAN 2, 2006 Monday 15:04:05 -0700 MST",
	561: "JAN 2, 2006 Monday 15:04:05.000000000 MST",
	562: "JAN 2, 2006 Monday 15:04:05.000000000 -0700",
	563: "JAN 2, 2006 Monday 15:04:05.000000000 -0700 MST",
	564: "JAN 2, 2006 Monday 03:04:05PM MST",
	565: "JAN 2, 2006 Monday 03:04:05PM -0700",
	566: "JAN 2, 2006 Monday 03:04:05PM -0700 MST",
	567: "JAN 2, 2006 Monday 03:04PM MST",
	568: "JAN 2, 2006 Monday 03:04PM -0700",
	569: "JAN 2, 2006 Monday 03:04PM -0700 MST",
	570: "JAN 2, 2006 Mon",
	571: "JAN 2, 2006 Mon 15:04:05 MST",
	572: "JAN 2, 2006 Mon 15:04:05 -0700",
	573: "JAN 2, 2006 Mon 15:04:05 -0700 MST",
	574: "JAN 2, 2006 Mon 15:04:05.000000000 MST",
	575: "JAN 2, 2006 Mon 15:04:05.000000000 -0700",
	576: "JAN 2, 2006 Mon 15:04:05.000000000 -0700 MST",
	577: "JAN 2, 2006 Mon 03:04:05PM MST",
	578: "JAN 2, 2006 Mon 03:04:05PM -0700",
	579: "JAN 2, 2006 Mon 03:04:05PM -0700 MST",
	580: "JAN 2, 2006 Mon 03:04PM MST",
	581: "JAN 2, 2006 Mon 03:04PM -0700",
	582: "JAN 2, 2006 Mon 03:04PM -0700 MST",
	583: "Monday JAN 2, 2006",
	584: "Monday JAN 2, 2006 15:04:05 MST",
	585: "Monday JAN 2, 2006 15:04:05 -0700",
	586: "Monday JAN 2, 2006 15:04:05 -0700 MST",
	587: "Monday JAN 2, 2006 15:04:05.000000000 MST",
	588: "Monday JAN 2, 2006 15:04:05.000000000 -0700",
	589: "Monday JAN 2, 2006 15:04:05.000000000 -0700 MST",
	590: "Monday JAN 2, 2006 03:04:05PM MST",
	591: "Monday JAN 2, 2006 03:04:05PM -0700",
	592: "Monday JAN 2, 2006 03:04:05PM -0700 MST",
	593: "Monday JAN 2, 2006 03:04PM MST",
	594: "Monday JAN 2, 2006 03:04PM -0700",
	595: "Monday JAN 2, 2006 03:04PM -0700 MST",
	596: "Mon JAN 2, 2006",
	597: "Mon JAN 2, 2006 15:04:05 MST",
	598: "Mon JAN 2, 2006 15:04:05 -0700",
	599: "Mon JAN 2, 2006 15:04:05 -0700 MST",
	600: "Mon JAN 2, 2006 15:04:05.000000000 MST",
	601: "Mon JAN 2, 2006 15:04:05.000000000 -0700",
	602: "Mon JAN 2, 2006 15:04:05.000000000 -0700 MST",
	603: "Mon JAN 2, 2006 03:04:05PM MST",
	604: "Mon JAN 2, 2006 03:04:05PM -0700",
	605: "Mon JAN 2, 2006 03:04:05PM -0700 MST",
	606: "Mon JAN 2, 2006 03:04PM MST",
	607: "Mon JAN 2, 2006 03:04PM -0700",
	608: "Mon JAN 2, 2006 03:04PM -0700 MST",
	609: "2 JAN, 2006",
	610: "2 JAN, 2006 15:04:05 MST",
	611: "2 JAN, 2006 15:04:05 -0700",
	612: "2 JAN, 2006 15:04:05 -0700 MST",
	613: "2 JAN, 2006 15:04:05.000000000 MST",
	614: "2 JAN, 2006 15:04:05.000000000 -0700",
	615: "2 JAN, 2006 15:04:05.000000000 -0700 MST",
	616: "2 JAN, 2006 03:04:05PM MST",
	617: "2 JAN, 2006 03:04:05PM -0700",
	618: "2 JAN, 2006 03:04:05PM -0700 MST",
	619: "2 JAN, 2006 03:04PM MST",
	620: "2 JAN, 2006 03:04PM -0700",
	621: "2 JAN, 2006 03:04PM -0700 MST",
	622: "2 JAN, 2006 Monday",
	623: "2 JAN, 2006 Monday 15:04:05 MST",
	624: "2 JAN, 2006 Monday 15:04:05 -0700",
	625: "2 JAN, 2006 Monday 15:04:05 -0700 MST",
	626: "2 JAN, 2006 Monday 15:04:05.000000000 MST",
	627: "2 JAN, 2006 Monday 15:04:05.000000000 -0700",
	628: "2 JAN, 2006 Monday 15:04:05.000000000 -0700 MST",
	629: "2 JAN, 2006 Monday 03:04:05PM MST",
	630: "2 JAN, 2006 Monday 03:04:05PM -0700",
	631: "2 JAN, 2006 Monday 03:04:05PM -0700 MST",
	632: "2 JAN, 2006 Monday 03:04PM MST",
	633: "2 JAN, 2006 Monday 03:04PM -0700",
	634: "2 JAN, 2006 Monday 03:04PM -0700 MST",
	635: "2 JAN, 2006 Mon",
	636: "2 JAN, 2006 Mon 15:04:05 MST",
	637: "2 JAN, 2006 Mon 15:04:05 -0700",
	638: "2 JAN, 2006 Mon 15:04:05 -0700 MST",
	639: "2 JAN, 2006 Mon 15:04:05.000000000 MST",
	640: "2 JAN, 2006 Mon 15:04:05.000000000 -0700",
	641: "2 JAN, 2006 Mon 15:04:05.000000000 -0700 MST",
	642: "2 JAN, 2006 Mon 03:04:05PM MST",
	643: "2 JAN, 2006 Mon 03:04:05PM -0700",
	644: "2 JAN, 2006 Mon 03:04:05PM -0700 MST",
	645: "2 JAN, 2006 Mon 03:04PM MST",
	646: "2 JAN, 2006 Mon 03:04PM -0700",
	647: "2 JAN, 2006 Mon 03:04PM -0700 MST",
	648: "Monday 2 JAN, 2006",
	649: "Monday 2 JAN, 2006 15:04:05 MST",
	650: "Monday 2 JAN, 2006 15:04:05 -0700",
	651: "Monday 2 JAN, 2006 15:04:05 -0700 MST",
	652: "Monday 2 JAN, 2006 15:04:05.000000000 MST",
	653: "Monday 2 JAN, 2006 15:04:05.000000000 -0700",
	654: "Monday 2 JAN, 2006 15:04:05.000000000 -0700 MST",
	655: "Monday 2 JAN, 2006 03:04:05PM MST",
	656: "Monday 2 JAN, 2006 03:04:05PM -0700",
	657: "Monday 2 JAN, 2006 03:04:05PM -0700 MST",
	658: "Monday 2 JAN, 2006 03:04PM MST",
	659: "Monday 2 JAN, 2006 03:04PM -0700",
	660: "Monday 2 JAN, 2006 03:04PM -0700 MST",
	661: "Mon 2 JAN, 2006",
	662: "Mon 2 JAN, 2006 15:04:05 MST",
	663: "Mon 2 JAN, 2006 15:04:05 -0700",
	664: "Mon 2 JAN, 2006 15:04:05 -0700 MST",
	665: "Mon 2 JAN, 2006 15:04:05.000000000 MST",
	666: "Mon 2 JAN, 2006 15:04:05.000000000 -0700",
	667: "Mon 2 JAN, 2006 15:04:05.000000000 -0700 MST",
	668: "Mon 2 JAN, 2006 03:04:05PM MST",
	669: "Mon 2 JAN, 2006 03:04:05PM -0700",
	670: "Mon 2 JAN, 2006 03:04:05PM -0700 MST",
	671: "Mon 2 JAN, 2006 03:04PM MST",
	672: "Mon 2 JAN, 2006 03:04PM -0700",
	673: "Mon 2 JAN, 2006 03:04PM -0700 MST",
	674: "2006-01-02T15:04:05Z07:00",
	675: "Mon 021504Z 0106",
	676: "Mon 021504Z 012006",
	677: "Mon 021504Z JAN2006",
	678: "Mon 021504Z JAN06",
	679: "Mon 2006-01-02T15:04:05Z07:00",
	680: "Monday 2 January, 2006",
	681: "Monday 2 January, 2006 15:04:05",
	682: "Monday 2 January, 2006 15:04:05.000000000",
	683: "Monday 2 January, 2006 15:04:05 -0700 MST",
	684: "Monday 2 January, 2006 15:04:05.000000000 -0700 MST",
	685: "Monday 021504Z 0106",
	686: "Monday 021504Z 012006",
	687: "Monday 021504Z JAN2006",
	688: "Monday 021504Z JAN06",
	689: "Monday 2006-01-02T15:04:05Z07:00",
	690: "021504Z 0106",
	691: "021504Z 012006",
	692: "021504Z JAN2006",
	693: "021504Z JAN06",
	694: "20060102150405",
	695: "20060102150405000000000",
}
