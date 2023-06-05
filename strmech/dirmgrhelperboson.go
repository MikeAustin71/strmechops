package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type dirMgrHelperBoson struct {
	lock *sync.Mutex
}

// copyDirMgrs
//
// Receives a pointer to an incoming DirMgr object
// ('sourceDMgrIn') as an input parameter and copies the
// data values from the incoming object to the input
// parameter, 'destinationDMgr'.
//
// When the copy operation is completed, the
// 'destinationDMgr' object is configured as a duplicate
// of the incoming DirMgr object ('sourceDMgrIn').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'sourceDMgrIn' is invalid, this
//	method will return an error.
//
//	If input parameter 'sourceDMgrIn' does NOT exist on
//	disk, no error will be returned and the copy operation
//	will proceed to completion.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationDMgr				*DirMgr
//
//		A pointer to an instance of DirMgr. All the data
//		values contained in input parameter
//		'sourceDMgrIn' will be copied to the
//		corresponding member data values contained in
//		'destinationDMgr'.
//
//		When the copy operation is completed, all data
//		values in 'destinationDMgr' will duplicate
//		corresponding data values contained in
//		'sourceDMgrIn'.
//
//	sourceDMgrIn				*DirMgr
//
//		A pointer to an instance of DirMgr. Data values
//		contained in this instance will be copied to
//		corresponding member data values encapsulated
//		by input parameter 'destinationDMgr'.
//
//		When the copy operation is completed, all data
//		values in 'destinationDMgr' will duplicate
//		corresponding data values contained in
//		'sourceDMgrIn'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (dMgrHlprBoson *dirMgrHelperBoson) copyDirMgrs(
	destinationDMgr *DirMgr,
	sourceDMgrIn *DirMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprBoson.lock == nil {
		dMgrHlprBoson.lock = new(sync.Mutex)
	}

	dMgrHlprBoson.lock.Lock()

	defer dMgrHlprBoson.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperBoson." +
		"copyDirMgrs()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if destinationDMgr == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'destinationDMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceDMgrIn == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'sourceDMgrIn' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	err = new(dirMgrHelperPlanck).isDirMgrValid(
		sourceDMgrIn,
		"sourceDMgrIn",
		ePrefix.XCpy("sourceDMgrIn"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input paramter 'sourceDMgrIn' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	destinationDMgr.isInitialized = sourceDMgrIn.isInitialized
	destinationDMgr.originalPath = sourceDMgrIn.originalPath
	destinationDMgr.path = sourceDMgrIn.path
	destinationDMgr.isPathPopulated = sourceDMgrIn.isPathPopulated
	destinationDMgr.doesPathExist = sourceDMgrIn.doesPathExist
	destinationDMgr.parentPath = sourceDMgrIn.parentPath
	destinationDMgr.isParentPathPopulated = sourceDMgrIn.isParentPathPopulated
	destinationDMgr.absolutePath = sourceDMgrIn.absolutePath
	destinationDMgr.isAbsolutePathPopulated = sourceDMgrIn.isAbsolutePathPopulated
	destinationDMgr.doesAbsolutePathExist = sourceDMgrIn.doesAbsolutePathExist
	destinationDMgr.isAbsolutePathDifferentFromPath = sourceDMgrIn.isAbsolutePathDifferentFromPath
	destinationDMgr.directoryName = sourceDMgrIn.directoryName
	destinationDMgr.volumeName = sourceDMgrIn.volumeName
	destinationDMgr.isVolumePopulated = sourceDMgrIn.isVolumePopulated
	destinationDMgr.actualDirFileInfo =
		sourceDMgrIn.actualDirFileInfo.CopyOut()

	return err
}
