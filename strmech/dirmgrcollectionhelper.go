package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type dirMgrCollectionHelper struct {
	lock *sync.Mutex
}

// copyCollection
//
// This method copies all the Director Manager objects
// contained in a 'source' Directory Manager Collection
// to a 'destination' Directory Manager Collection.
//
// All Directory Manager objects will be copied from
// input parameter 'sourceDMgrCollection' to input
// parameter 'destinationDMgrCollection'. Both instances
// are of type DirMgrCollection. The type copy operation
// employed is a 'deep' copy operation.
//
// If input parameter 'sourceDMgrCollection' contains
// an empty Directory Manager collection, an error will
// be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data fields contained in input parameter
//	'destinationDMgrCollection' will be deleted an
//	overwritten with new data values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationDMgrCollection	*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection. All
//		data fields contained in input parameter
//		'sourceDMgrCollection' will be copied to
//		corresponding data fields in
//		'destinationDMgrCollection'.
//
//		All original pre-existing data values contained
//		in 'destinationDMgrCollection' will be deleted
//		and overwritten.
//
//	sourceDMgrCollection *DirMgrCollection,
//
//		A pointer to an instance of DirMgrCollection. All
//		data fields contained in input parameter
//		'sourceDMgrCollection' will be copied to
//		corresponding data fields in
//		'destinationDMgrCollection'.
//
//		If the Directory Manager Collection contained
//		in 'sourceDMgrCollection' is empty.
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
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrColHelper *dirMgrCollectionHelper) copyCollection(
	destinationDMgrCollection *DirMgrCollection,
	sourceDMgrCollection *DirMgrCollection,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrCollectionHelper.copyCollection()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if sourceDMgrCollection == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceDMgrCollection' is a nil pointer!\n",
			ePrefix.String())
	}

	if destinationDMgrCollection == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationDMgrCollection' is a nil pointer!\n",
			ePrefix.String())
	}

	lenSourceDMgrs := len(sourceDMgrCollection.dirMgrs)

	if lenSourceDMgrs == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceDMgrCollection' is invalid!\n"+
			"The 'sourceDMgrCollection' collection of Directory Managers is Empty.\n",
			ePrefix.String())
	}

	destinationDMgrCollection.dirMgrs = make([]DirMgr, lenSourceDMgrs)

	for i := 0; i < lenSourceDMgrs; i++ {

		err = destinationDMgrCollection.dirMgrs[i].
			CopyIn(
				&sourceDMgrCollection.dirMgrs[i],
				ePrefix.XCpy(
					"sourceDMgrCollection.dirMgrs[i]"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error: destinationDMgrCollection.dirMgrs[%v].CopyIn()\n"+
				"sourceDMgrCollection.dirMgrs index= '%v'\n"+
				"sourceDMgrCollection DirMgr= '%v'\n"+
				"Error= \n%v\n",
				funcName,
				i,
				i,
				sourceDMgrCollection.dirMgrs[i].absolutePath,
				err.Error())
		}

	}

	return err
}

// peekOrPopAtIndex
//
// Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the
// Directory Manager Collection passed as input parameter
// 'dirMgrs'.
//
// If input parameter 'deleteIndex' is set to 'false',
// this method will function as a 'Peek' method and
// therefore, the original Directory Manager ('DirMgr')
// object will NOT be deleted from the Directory Manager
// Collection ('DirMgrCollection') array.
//
// If input parameter 'deleteIndex' is set to 'true',
// this method will function as a 'Pop' method and
// therefore, the original Directory Manager ('DirMgr')
// object WILL BE DELETED from the Directory Manager
// Collection ('DirMgrCollection') array. The deletion
// operation will be performed on the Directory Manager
// object residing at the Directory Manager Collection
// array index identified by input parameter 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryMgrs				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		A 'peek' or 'pop' operation will be performed
//		on the DirMgr object specified by the array index
//		'idx'.
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be deleted from the File
//		Manager Collection encapsulated by the instance
//		of FileMgrCollection passed by input parameter
//		'fMgrs'.
//
//		If this value is less than zero or greater than
//		the last index in the array, an error will be
//		returned.
//
//	deleteIndex					bool
//
//		If this boolean value is set to 'true', the File
//		Manager object residing at the File Manager
//		Collection index identified by input parameter
//		'idx', will be deleted from File Manager
//		Collection 'fMgrs'.
//
//		If 'deleteIndex' is set to 'false', no deletion
//		occur and the File Manager object residing at
//		File Manager Collection index 'idx' will remain.
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at array index 'idx' will be returned
//		through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrColHelper *dirMgrCollectionHelper) peekOrPopAtIndex(
	directoryMgrs *DirMgrCollection,
	idx int,
	deleteIndex bool,
	errPrefDto *ePref.ErrPrefixDto) (
	DirMgr,
	error) {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrCollectionHelper.peekOrPopAtIndex()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return DirMgr{}, err
	}

	if directoryMgrs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'directoryMgrs' is a nil pointer!\n",
			ePrefix.String())

		return DirMgr{}, err
	}

	if directoryMgrs.dirMgrs == nil {
		directoryMgrs.dirMgrs = make([]DirMgr, 0)
	}

	arrayLen := len(directoryMgrs.dirMgrs)

	if arrayLen == 0 {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: The Directory Manager Collection array encapsulated by 'directoryMgrs' is EMPTY!\n",
				ePrefix.String())
	}

	if idx < 0 {
		return DirMgr{},
			fmt.Errorf("%v"+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'",
				ePrefix.String(),
				idx)
	}

	if idx >= arrayLen {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				" last array index of the Directory Collection array.\n"+
				"Index Out-Of-Range!\n"+
				"idx= '%v' "+
				"Last Array Index= '%v' ",
				ePrefix.String(),
				idx,
				arrayLen-1)
	}

	var deepCopyDirMgr DirMgr

	deepCopyDirMgr,
		err = directoryMgrs.dirMgrs[idx].CopyOut(
		ePrefix.XCpy(fmt.Sprintf(
			"directoryMgrs.directoryMgrs[%v]",
			idx)))

	if err != nil {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Error: directoryMgrs.dirMgrs[%v].CopyOut()\n"+
				"directoryMgrs.dirMgrs index = '%v'\n"+
				"directoryMgrs.dirMgrs[%v] = '%v'\n"+
				"Error= \n%v\n",
				funcName,
				idx,
				idx,
				idx,
				directoryMgrs.dirMgrs[idx].absolutePath,
				err.Error())
	}

	if deleteIndex == false {
		return deepCopyDirMgr, err
	}

	// deleteIndex == true

	if arrayLen == 1 {

		directoryMgrs.dirMgrs = make([]DirMgr, 0, 1)

	} else if idx == 0 {
		// arrayLen > 1 and requested idx = 0
		directoryMgrs.dirMgrs = directoryMgrs.dirMgrs[1:]

	} else if idx == arrayLen-1 {
		// arrayLen > 1 and requested idx = last element index
		directoryMgrs.dirMgrs = directoryMgrs.dirMgrs[0 : arrayLen-1]

	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		directoryMgrs.dirMgrs =
			append(directoryMgrs.dirMgrs[0:idx], directoryMgrs.dirMgrs[idx+1:]...)
	}

	return deepCopyDirMgr, err
}
