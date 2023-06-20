package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperPreon struct {
	lock *sync.Mutex
}

// getSubdirectories
//
// This method receives an instance of DirMgr ('dMgr')
// and proceeds to identify all the subdirectories
// located within the directory path specified by this
// DirMgr instance.
//
// Each subdirectory located in the 'dMgr' parent
// directory will be recorded as a separate DirMgr object
// and added to the Directory Manager collection passed
// as input parameter 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method will only identify, document and
//		return subdirectories located in the top level or
//		parent directory identified by input parameter
//		'dMgr'.
//
//	(2) The top level or parent directory specified by
//		input parameter 'dMgr' will NOT be included in
//		the Directory Manager collection returned by this
//		method ('subDirectories').
//
//	(3)	Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped.
//		These directory entries will not be added or
//		included in the subdirectories collection
//		('subDirectories'). Likewise, these two directory
//		entries will NOT be included in the subdirectory
//		profile and statistical information returned by
//		this method ('dirProfile').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will be
//		searched to extract and return information on
//		subdirectories residing within that directory path.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		The DirMgrCollection contains an array of DirMgr
//		objects.
//
//		This method will scan the top level or parent
//		directory identified by input parameter 'dMgr'.
//		All subdirectories identified in this parent
//		directory will be documented by a new instance
//		of DirMgr which will be added to the
//		'subDirectories' collection.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped.
//		These directory entries will not be added or
//		included in the subdirectories collection
//		('subDirectories'). Likewise, these two directory
//		entries will NOT be included in the subdirectory
//		profile and statistical information returned by
//		this method ('dirProfile').
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
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
//	numOfSubDirsReturned		int
//
//		This integer value returns the number of
//		subdirectories added by this method to the
//		Directory Manager Collection passed as
//		input parameter.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'.
//
//		If a fatal error is encountered during
//		processing, this returned error Type will
//		encapsulate an appropriate error message. This
//		returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errPrefDto'. The 'errPrefDto' text will be
//		prefixed or attached to the	beginning of the error
//		message.
func (dMgrHlprPreon *dirMgrHelperPreon) getSubdirectories(
	dMgr *DirMgr,
	subDirectories *DirMgrCollection,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfSubDirsReturned int,
	err error) {

	if dMgrHlprPreon.lock == nil {
		dMgrHlprPreon.lock = new(sync.Mutex)
	}

	dMgrHlprPreon.lock.Lock()

	defer dMgrHlprPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperPreon." +
		"getSubdirectories()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfSubDirsReturned, err
	}

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is a 'nil' pointer!\n",
			ePrefix.String())

		return numOfSubDirsReturned, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		return numOfSubDirsReturned, err
	}

	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var nonFatalErrs []error
	var fatalErr, err2 error

	fileInfos,
		lenFileInfos,
		nonFatalErrs,
		err2 = new(dirMgrHelperMolecule).
		lowLevelGetFileInfosFromDir(
			dMgr,
			true,                    // getSubdirectoryFileInfos
			false,                   // includeSubDirCurrenDirOneDot
			false,                   // includeSubDirParentDirTwoDots
			false,                   // getRegularFileInfos
			false,                   // getSymLinksFileInfos
			false,                   // getOtherNonRegularFileInfos
			FileSelectionCriteria{}, // subdirectorySelectCharacteristics
			FileSelectionCriteria{}, // fileSelectCharacteristics
			dMgrLabel,
			ePrefix)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error occurred while selecting subdirectories.\n"+
			"%v Absolute Path= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		nonFatalErrs = append(
			nonFatalErrs, fatalErr)

		err = new(StrMech).ConsolidateErrors(nonFatalErrs)

		return numOfSubDirsReturned, err

	}

	if lenFileInfos == 0 {

		return numOfSubDirsReturned, err
	}

	osPathSepStr := string(os.PathSeparator)

	for i := 0; i < lenFileInfos; i++ {

		if fileInfos[i].IsDir() {

			err2 = subDirectories.
				AddDirMgrByKnownPathDirName(
					dMgr.absolutePath,
					fileInfos[i].Name(),
					ePrefix.XCpy(
						"subDirectories<-dMgr"))

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error occurred while adding subdirectory\n"+
					"to 'subDirectories' collection.\n"+
					"%v Absolute Path= '%v'\n"+
					"Subdirectory Name= '%v'\n"+
					"Subdirectory Path= '%v'\n"+
					"Index= '%v'\n"+
					"Error=\n%v\n",
					funcName,
					dMgrLabel,
					dMgr.absolutePath,
					fileInfos[i].Name(),
					dMgr.absolutePath+
						osPathSepStr+
						fileInfos[i].Name(),
					i,
					err2.Error())

				return numOfSubDirsReturned, err
			}

			numOfSubDirsReturned++
		}
	}

	return numOfSubDirsReturned, err
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

	err2 = new(dirMgrHelperPlanck).
		isDirMgrValid(
			dMgr,
			dMgrLabel,
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
			"Error: The current DirMgr path DOES NOT EXIST!\n"+
			"%v.absolutePath='%v'\n",
			ePrefix.String(),
			finalDirMgrLabel,
			dMgr.absolutePath)

	}

	return finalDirMgrLabel, pathDoesExist, err
}
