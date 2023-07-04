package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// fileAccessControlElectron
//
// Provides helper methods for Type
// FileOpenConfig.
type fileOpenConfigElectron struct {
	lock *sync.Mutex
}

// testValidityFileOpenConfig
//
// Receives a pointer to an instance of FileOpenConfig
// and performs a diagnostic analysis to determine if
// that instance is valid in all respects.
//
// If the input parameter 'fOpenCfg' is determined to
// be invalid, this method will return a boolean flag
// ('isValid') of 'false'. In addition, an instance of
// type error ('err') will be returned configured with an
// appropriate error message.
//
// If the input parameter 'fOpenCfg' is valid, this
// method will return a boolean flag ('isValid') of
// 'true' and the returned error type ('err') will be set
// to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOpenCfg					*FileOpenConfig
//
//		A pointer to an instance of FileOpenConfig.
//		This object will be subjected to diagnostic
//		analysis in order to determine if all the
//		member data variables contain valid values.
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
//		contained in the input parameter 'fOpenCfg'
//		are found to be invalid, this return parameter
//		will be set to 'false'.
//
//		If all the internal member data variables
//		contained in the input parameter 'fOpenCfg'
//		are found to be valid, this return parameter
//		will be set to 'true'.
//
//	err							error
//
//		If any of the internal member data variables
//		contained in the input parameter 'fOpenCfg'
//		are found to be invalid, this method will return
//		an error configured with an appropriate message
//		identifying the invalid member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any internal member data values are found to be
//		invalid, the returned error Type will encapsulate
//		an appropriate error message. This returned error
//		message will incorporate the method chain and text
//		passed by input parameter, 'errPrefDto'. The
//		'errPrefDto' text will be attached to the
//		beginning of the error message.
func (fOpenCfgElectron *fileOpenConfigElectron) testValidityFileOpenConfig(
	fOpenCfg *FileOpenConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if fOpenCfgElectron.lock == nil {
		fOpenCfgElectron.lock = new(sync.Mutex)
	}

	fOpenCfgElectron.lock.Lock()

	defer fOpenCfgElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileOpenConfigElectron." +
		"testValidityFileOpenConfig()"

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return isValid, err
	}

	if fOpenCfg == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOpenConfig object is invalid!\n"+
			"Input parameter 'fOpenCfg' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if fOpenCfg.fileOpenModes == nil {
		fOpenCfg.fileOpenModes = make([]FileOpenMode, 0)
	}

	if !fOpenCfg.isInitialized {
		return isValid,
			fmt.Errorf("%v\n"+
				"Error: The current FileOpenConfig instance has\n"+
				"NOT been properly initialized.\n",
				ePrefix)
	}

	err = fOpenCfg.fileOpenType.IsValid()

	if err != nil {
		return isValid,
			fmt.Errorf("%v\n"+
				"Error: The File Open Type is INVALID!.\n"+
				"Error=\n%v\n",
				ePrefix,
				err.Error())
	}

	lenFileOpenModes := len(fOpenCfg.fileOpenModes)

	if fOpenCfg.fileOpenType == FOpenType.TypeNone() &&
		lenFileOpenModes > 1 {

		return isValid,
			fmt.Errorf("%v\n"+
				"Error: Current FileOpenConfig has Type='None'\n"+
				"and multiple File Open Modes!\n"+
				"Number Of File Open Modes = %v\n",
				ePrefix,
				lenFileOpenModes)
	}

	if fOpenCfg.fileOpenType == FOpenType.TypeNone() &&
		lenFileOpenModes == 1 &&
		fOpenCfg.fileOpenModes[0] != FileOpenMode(0).ModeNone() {

		return isValid,
			fmt.Errorf("%v\n"+
				"Error: Current FileOpenConfig has Type='None' and "+
				"a valid File Open Mode\n",
				ePrefix)
	}

	if fOpenCfg.fileOpenType != FOpenType.TypeNone() &&
		lenFileOpenModes > 1 {

		for i := 0; i < lenFileOpenModes; i++ {
			if fOpenCfg.fileOpenModes[i] == FileOpenMode(0).ModeNone() {

				return isValid,
					fmt.Errorf("%v\n"+
						"Error: The File Open Status has multiple File Open Modes\n"+
						"one of which is 'None'.\n"+
						"fOpenCfg.fileOpenModes[%v] == ModeNone\n	"+
						"Please resolve this conflict.\n",
						ePrefix,
						i)
			}
		}

	}

	for i := 0; i < lenFileOpenModes; i++ {

		err = fOpenCfg.fileOpenModes[i].XIsValid()

		if err != nil {

			return isValid,
				fmt.Errorf("%v\n"+
					"Error: A File Open Mode is INVALID!\n"+
					"Index='%v'\n"+
					"Invalid Error=\n%v\n ",
					ePrefix,
					i,
					err.Error())
		}

	}

	isValid = true

	return isValid, nil
}
