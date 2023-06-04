package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type dirMgrHelperPreon struct {
	lock *sync.Mutex
}

// validateDirMgr
//
// This method performs a comprehensive analysis to
// determine if an instance of DirMgr is valid.
//
// Users have the option to configure the validity test
// to require that the Directory Manager directory path
// actually exists on disk.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		This instance of DirMgr will be analyzed to
//		determine if all data values are valid.
//
//	pathMustExist				bool
//
//		If this parameter is set to 'true', the directory
//		path contained in 'dMgr' must exist on disk as
//		requirement for validation.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
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
//	finalDirMgrLabel			string
//
//		The final formatted version of input parameter
//		'dMgrLabel'.
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	pathDoesExist               bool
//
//		If this return parameter is set to 'true' it
//		signals that the directory path contained in the
//		Directory Manager instance 'dMgr' actually exists
//		on disk.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrHlprPreon *dirMgrHelperPreon) validateDirMgr(
	dMgr *DirMgr,
	pathMustExist bool,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	finalDirMgrLabel string,
	pathDoesExist bool,
	err error) {

	if dMgrHlprPreon.lock == nil {
		dMgrHlprPreon.lock = new(sync.Mutex)
	}

	dMgrHlprPreon.lock.Lock()

	defer dMgrHlprPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathDoesExist = false

	funcName := "dirMgrHelperPreon." +
		"validateDirMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if len(dMgrLabel) == 0 {

		finalDirMgrLabel = "dMgr"

	} else {

		finalDirMgrLabel = dMgrLabel
	}

	if dMgr == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter '%v' is a nil pointer!\n",
			ePrefix.String(),
			finalDirMgrLabel)

		return finalDirMgrLabel, pathDoesExist, err
	}

	var err2 error

	err2 = new(dirMgrHelperBoson).
		isDirMgrValid(
			dMgr,
			ePrefix.XCpy(finalDirMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter '%v' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			finalDirMgrLabel,
			err2.Error())

		return finalDirMgrLabel, pathDoesExist, err
	}

	pathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.AbsolutePath(),
				finalDirMgrLabel,
				ePrefix.XCpy(finalDirMgrLabel))

	if err != nil {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if !pathMustExist {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"The current DirMgr path DOES NOT EXIST!\n"+
			"%v.absolutePath='%v'\n",
			ePrefix.String(),
			finalDirMgrLabel,
			dMgr.absolutePath)

	}

	return finalDirMgrLabel, pathDoesExist, err
}
