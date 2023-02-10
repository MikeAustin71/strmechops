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

// allocateDurationToTimeElement
func (dateTimeHelpElectron *dateTimeHelperElectron) allocateDurationToTimeElement(
	allocDurationElement int64,
	allocDurationElementName string,
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

	nStrIntSepMolecule := integerSeparatorSpecMolecule{}

	var tLine, newOutputLine string
	var numStrWithIntSeps []rune
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
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(tLine),
				ePrefix.XCpy(
					fmt.Sprintf("numStrWithIntSeps<-Element Name: %v",
						allocDurationElementName)))

		if err != nil {
			return err
		}

		newOutputLine = fmt.Sprintf(
			"%v %v ",
			string(numStrWithIntSeps),
			allocDurationElementName)

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
