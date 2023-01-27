package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtSpecElectron - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecElectron struct {
	lock *sync.Mutex
}

//	testValidityNumStrFormatSpec
//
//	Performs a diagnostic review of the data values
//	encapsulated in an instance of NumStrFormatSpec
//	to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberStrFmtSpec			*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		All member variable data values contained in
//		this instance will be reviewed to determine if
//		they are valid.
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
//		If any of the internal member data variables
//		contained in 'numberStrFmtSpec' are found to be
//		invalid, this method will return a boolean value
//		of 'false'.
//
//		If all internal member data variables contained
//		in 'numberStrFmtSpec' are found to be valid, this
//		method returns a boolean value of 'true'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, or
//		if input parameter 'numberStrFmtSpec' contains
//		invalid data elements, an error will be returned.
//		The returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrFmtSpecElectron *numStrFmtSpecElectron) testValidityNumStrFormatSpec(
	numberStrFmtSpec *NumStrFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrFmtSpecElectron.lock == nil {
		nStrFmtSpecElectron.lock = new(sync.Mutex)
	}

	nStrFmtSpecElectron.lock.Lock()

	defer nStrFmtSpecElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecElectron."+
			"testValidityNumStrFormatSpec()",
		"")

	if err != nil {
		return isValid, err
	}

	if numberStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberStrFmtSpec' is invalid!\n"+
			"'numberStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return isValid, err
	}

	err = numberStrFmtSpec.decSeparator.
		IsValidInstanceError(
			ePrefix.XCpy(
				"numberStrFmtSpec.decSeparator"))

	if err != nil {
		return isValid, err
	}

	err = numberStrFmtSpec.intSeparatorSpec.
		IsValidInstanceError(
			ePrefix.XCpy(
				"numberStrFmtSpec.intSeparatorSpec"))

	if err != nil {
		return isValid, err
	}

	if numberStrFmtSpec.numberSymbolsGroup.IsNOP() {

		err = fmt.Errorf("%v\n"+
			"'numberStrFmtSpec.numberSymbolsGroup' is not configured.\n"+
			"'numberStrFmtSpec.numberSymbolsGroup' is NOP, Not Operational.\n",
			ePrefix.String())

		return isValid, err

	}

	isValid = true

	return isValid, err
}
