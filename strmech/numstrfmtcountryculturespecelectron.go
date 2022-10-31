package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtCountryCultureSpecElectron
//
// Provides helper methods for type
// NumStrFmtCountryCultureSpec
type numStrFmtCountryCultureSpecElectron struct {
	lock *sync.Mutex
}

//	copyNumberFieldSpec
//
//	Copies a source Number String Number Field
//	Specification (NumStrNumberFieldSpec) to a destination
//	Number String Number Field Specification.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'sourceNumFieldSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationNumFieldSpec			*NumStrNumberFieldSpec
//
//		A pointer to a NumStrNumberFieldSpec
//		instance. The Number String Number Field
//		Specification from input parameter
//		'sourceNumFieldSpec' will be copied to this
//		parameter 'destinationNumFieldSpec'.
//
//		All data values contained in
//		'destinationNumFieldSpec' will be deleted and
//		reset to new values extracted from
//		'sourceNumFieldSpec'.
//
//	sourceNumFieldSpec				*NumStrNumberFieldSpec
//
//		A pointer to a NumStrNumberFieldSpec
//		instance. The Number String Number Field
//		Specification from this parameter
//		'sourceNumFieldSpec' will be copied to the
//		destination input parameter,
//		'destinationNumFieldSpec'.
//
//		The data values contained in 'sourceNumFieldSpec'
//		WILL NOT be modified by this method.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrFmtCountryCultureElectron *numStrFmtCountryCultureSpecElectron) copyNumberFieldSpec(
	destinationNumFieldSpec *NumStrNumberFieldSpec,
	sourceNumFieldSpec *NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrFmtCountryCultureElectron.lock == nil {
		nStrFmtCountryCultureElectron.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureElectron.lock.Lock()

	defer nStrFmtCountryCultureElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtCountryCultureSpecElectron."+
			"copyNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNumFieldSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNumFieldSpec' is invalid!\n"+
			"'countryCultureSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNumFieldSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNumFieldSpec' is invalid!\n"+
			"'sourceNumFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	destinationNumFieldSpec.Empty()

	return destinationNumFieldSpec.CopyIn(
		sourceNumFieldSpec,
		ePrefix.XCpy(
			"<-signedNumFieldSpec"))
}
