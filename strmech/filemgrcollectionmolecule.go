package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileMgrCollectionMolecule struct {
	lock *sync.Mutex
}

// peekOrPopAtIndex
//
// Returns a deep copy of the File Manager ('FileMgr')
// object located at array index 'idx' in the File
// Manager Collection passed as input parameter 'fMgrs'.
//
// If input parameter 'deleteIndex' is set to 'false',
// this method will function as a 'Peek' method and
// therefore, the original File Manager ('FileMgr')
// object will NOT be deleted from the File Manager
// Collection ('FileMgrCollection') array.
//
// If input parameter 'deleteIndex' is set to 'true',
// this method will function as a 'Pop' method and
// therefore, the original File Manager ('FileMgr')
// object WILL BE DELETED from the File Manager
// Collection ('FileMgrCollection') array. The
// deletion operation will be performed on the File
// Manager object residing at the File Manager
// Collection array index identified by input parameter
// 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrs						*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection.
//		The File Manager object 'fMgr' specified by
//		the array index 'idx' will be deleted will be
//		deleted from the File Manager Collection.
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
func (fMgrColMolecule *FileMgrCollectionMolecule) peekOrPopAtIndex(
	fMgrs *FileMgrCollection,
	idx int,
	deleteIndex bool,
	errPrefDto *ePref.ErrPrefixDto) (
	FileMgr,
	error) {

	if fMgrColMolecule.lock == nil {
		fMgrColMolecule.lock = new(sync.Mutex)
	}

	fMgrColMolecule.lock.Lock()

	defer fMgrColMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileMgrCollectionElectron."+
			"peekOrPopAtIndex()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	if fMgrs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgrs' is a nil pointer!\n",
			ePrefix.String())

		return FileMgr{}, err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: The File Manager Collection array encapsulated by 'fMgrs' is EMPTY!\n",
				ePrefix.String())
	}

	if idx < 0 {
		return FileMgr{},
			fmt.Errorf("%v"+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'",
				ePrefix.String(),
				idx)
	}

	if idx >= arrayLen {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				" last array index of the collection array.\n"+
				"Index Out-Of-Range!\n"+
				"idx= '%v' "+
				"Last Array Index= '%v' ",
				ePrefix.String(),
				idx,
				arrayLen-1)
	}

	var deepCopyFileMgr FileMgr

	deepCopyFileMgr,
		err = fMgrs.fileMgrs[idx].CopyOut(
		ePrefix.XCpy(fmt.Sprintf(
			"fMgrs.fileMgrs[%v]",
			idx)))

	if err != nil ||
		deleteIndex == false {
		return deepCopyFileMgr, err
	}

	// deleteIndex == true

	if arrayLen == 1 {

		fMgrs.fileMgrs = make([]FileMgr, 0, 5)

	} else if idx == 0 {
		// arrayLen > 1 and requested idx = 0
		fMgrs.fileMgrs = fMgrs.fileMgrs[1:]

	} else if idx == arrayLen-1 {
		// arrayLen > 1 and requested idx = last element index
		fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]

	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		fMgrs.fileMgrs =
			append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)
	}

	return deepCopyFileMgr, err
}
