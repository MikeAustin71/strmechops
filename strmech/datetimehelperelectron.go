package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strconv"
	"sync"
)

// DateTimeHelper
//
// Provides helper methods for Type DateTimeHelper.
type dateTimeHelperElectron struct {
	lock *sync.Mutex
}

//	allocateDurationToTimeElement
//
//	This method receives allocated time duration elements
//	and formats these elements for output as text
//	strings.
//
//	Allocated Time Duration is typically broken down for
//	reporting purposes into days, hours, minutes,
//	seconds, milliseconds, microseconds and nanoseconds.
//	This breakdown represents the individual time
//	duration elements passed to this method for
//	formatting and conversion to text strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	allocDurationElement			int64
//
//		The numerical value of the time duration element
//		being processed. This element value will be
//		converted to text and stored in the string array
//		parameter, 'allocatedDurationStrs'.
//
//	allocDurationElementName		string
//
//		The name or text label associated with the time
//		duration element value passed in parameter
//		'allocDurationElement'. This string will be
//		used to label the 'allocDurationElement' value
//		converted to a string and stored in parameter
//		'allocatedDurationStrs'.
//
//	labelNameFirst					bool
//
//		When the boolean value is set to 'true' the
//		label 'allocDurationElementName' is positioned
//		before the numerical value.
//
//		When this parameter is set to 'false', the
//		label 'allocDurationElementName' is positioned
//		after the numerical value.
//
//	alwaysDisplayDurationElement	bool
//
//		This formatting specification determines
//		whether the time duration element value will
//		be converted to text even the element value
//		is zero. This allows the calling function to
//		select whether zero element values will be
//		included in the final formatted text string.
//
//	foundFirstElementValue			*bool
//
//		A pointer to a boolean value.
//
//		This internal flag signals whether previous time
//		duration element values contain non-zero values.
//		In some cases it is preferable to skip a time
//		duration element value if it has a zero value
//		and no non-zero elements have been previously
//		reported.
//
//	finalOutputLineText				*string
//
//		A pointer to the string containing the current
//		output text. Text strings containing new time
//		duration elements are added to this string until
//		such time that the length of
//		'finalOutputLineText' equals or exceeds the
//		maximum line length.
//
//		If 'finalOutputLineText' equals or exceeds the
//		maximum line length, this text is saved to
//		the string array 'allocatedDurationStrs' and
//		'finalOutputLineText' is emptied in preparation
//		for receiving more text output.
//
//	maxLineLength					int
//
//		Defines the maximum length of the text line
//		in which allocated time duration elements will be
//		displayed.
//
//		If 'finalOutputLineText' equals or exceeds the
//		maximum line length ('maxLineLength'), time
//		duration text is saved to the string array
//		'allocatedDurationStrs' and 'finalOutputLineText'
//		is emptied in preparation for receiving more text
//		output.
//
//	allocatedDurationStrs		*StringArrayDto
//
//		Type StringArrayDto maintains an internal array
//		of strings. This string array is used as storage
//		for all time duration text lines created and
//		formatted by this method.
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
//	error
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
func (dateTimeHelpElectron *dateTimeHelperElectron) allocateDurationToTimeElement(
	allocDurationElement int64,
	allocDurationElementName string,
	labelNameFirst bool,
	alwaysDisplayDurationElement bool,
	foundFirstElementValue *bool,
	finalOutputLineText *string,
	maxLineLength int,
	allocatedDurationStrs *StringArrayDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dateTimeHelpElectron.lock == nil {
		dateTimeHelpElectron.lock = new(sync.Mutex)
	}

	dateTimeHelpElectron.lock.Lock()

	defer dateTimeHelpElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dateTimeHelperAtom."+
			"allocateInt64TimeDuration()",
		"")

	if err != nil {
		return err
	}

	// Reduce Max Line Length by 1
	// to account for new line character.
	maxLineLength--

	var nStrIntSeparator IntegerSeparatorSpec

	nStrIntSeparator,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy(
			"nStrIntSeparator<-"))

	if err != nil {
		return err
	}

	var tLine, newOutputLine string
	var numStrWithIntSeps string
	var lenNewOutputLine, lenFinalOutputLine int

	if allocDurationElement > 0 ||
		alwaysDisplayDurationElement == true ||
		*foundFirstElementValue == true {

		*foundFirstElementValue = true

		tLine =
			strconv.FormatInt(
				allocDurationElement, 10)

		numStrWithIntSeps,
			err =
			nStrIntSeparator.GetFmtIntSeparatedNumStr(
				tLine,
				ePrefix.XCpy(
					fmt.Sprintf("numStrWithIntSeps<-Element Name: %v",
						allocDurationElementName)))

		if err != nil {
			return err
		}

		if labelNameFirst == true {

			newOutputLine = fmt.Sprintf(
				"%v %v ",
				allocDurationElementName,
				numStrWithIntSeps)

		} else {

			newOutputLine = fmt.Sprintf(
				"%v %v ",
				numStrWithIntSeps,
				allocDurationElementName)

		}

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(*finalOutputLineText)

		if lenNewOutputLine+lenFinalOutputLine < maxLineLength {

			*finalOutputLineText += newOutputLine

			return err
		}

		if lenNewOutputLine+lenFinalOutputLine >= maxLineLength {

			if lenFinalOutputLine > 0 {

				*finalOutputLineText += "\n"

				allocatedDurationStrs.AddString(
					*finalOutputLineText)

				*finalOutputLineText = newOutputLine

				newOutputLine = ""

			} else {

				// len(finalOutputLineText) == 0
				*finalOutputLineText = newOutputLine

				newOutputLine = ""

			}

		}

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(*finalOutputLineText)

		if lenFinalOutputLine >= maxLineLength {

			*finalOutputLineText += "\n"

			allocatedDurationStrs.AddString(
				*finalOutputLineText)

			*finalOutputLineText = newOutputLine

			newOutputLine = ""

		}

		lenNewOutputLine = len(newOutputLine)

		lenFinalOutputLine = len(*finalOutputLineText)

		if lenNewOutputLine >= maxLineLength {

			if lenFinalOutputLine > 0 {

				*finalOutputLineText += "\n"

				allocatedDurationStrs.AddString(
					*finalOutputLineText)

				*finalOutputLineText = ""

			}

			*finalOutputLineText = newOutputLine

			newOutputLine = ""

			*finalOutputLineText += "\n"

			allocatedDurationStrs.AddString(
				*finalOutputLineText)

			*finalOutputLineText = ""

		}

	} /*
		END OF 	if allocDurationElement > 0 ||
					alwaysDisplayDurationElement == true ||
					foundFirstElementValue == true
	*/

	return err
}
