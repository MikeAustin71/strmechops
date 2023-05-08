package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileMgrCollectionElectron struct {
	lock *sync.Mutex
}

//	deleteAtIndex
//
// This method deletes an array element from the File
// Manager Collection encapsulated by the instance of
// FileMgrCollection passed as input parameter 'fMgrs'.
//
// The index of the array element to be deleted is
// specified by input parameter 'idx'.
//
// If this method completes successfully, the File
// Manager Collection ('fMgrs') array will have a
// length which is one less than the starting array
// length.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgrColElectron *FileMgrCollectionElectron) deleteAtIndex(
	fMgrs *FileMgrCollection,
	idx int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrColElectron.lock == nil {
		fMgrColElectron.lock = new(sync.Mutex)
	}

	fMgrColElectron.lock.Lock()

	defer fMgrColElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileMgrCollectionElectron."+
			"deleteAtIndex()",
		"")

	if err != nil {
		return err
	}

	if fMgrs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgrs' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 5)
	}

	if idx < 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'idx' is less than zero.\n"+
			"Index Out-Of-Range!\n"+
			"idx='%v'\n",
			ePrefix.String(),
			idx)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return fmt.Errorf("%v\n"+
			"Error: The File Manager Collection, 'FileMgrCollection', is EMPTY!\n",
			ePrefix.String())
	}

	if idx >= arrayLen {
		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'idx' is greater than the "+
			"last index in the File Manager Collection.\n"+
			"Index Out-Of-Range!\n"+
			"idx= '%v'\n"+
			"Last Array Index= '%v' ",
			ePrefix.String(),
			idx,
			arrayLen-1)
	}

	if arrayLen == 1 {
		fMgrs.fileMgrs = make([]FileMgr, 0, 100)
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

	return err
}

// addFileMgr
//
// Adds a deep copy of the FileMgr object passed as input
// parameter 'fMgr' to the File Manager Collection
// maintained by a FileMgrCollection instance passed as
// input parameter 'fMgrCollection'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No validation is performed on input parameter 'fMgr'
//	before a deep copy of 'fMgr' is added to the
//	File Manager Collection passed by input parameter
//	'fMgrCollection'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrCollection				*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection.
//		A deep copy of the File Manager object 'fMgr'
//		will be added to the File Manager Collection
//		maintained by fMgrCollection''.
//
//
//	fMgr						FileMgr
//
//		A pointer to an instance of FileMgr. A deep copy
//		of this FileMgr object will be added to the File
//		Manager Collection maintained by the instance of
//		FileMgrCollection passed by input parameter
//		'fMgrCollection'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgrColElectron *FileMgrCollectionElectron) addFileMgr(
	fMgrCollection *FileMgrCollection,
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrColElectron.lock == nil {
		fMgrColElectron.lock = new(sync.Mutex)
	}

	fMgrColElectron.lock.Lock()

	defer fMgrColElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileMgrCollectionElectron."+
			"addFileMgr()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fMgrCollection == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgrCollection' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fMgrCollection.fileMgrs == nil {
		fMgrCollection.fileMgrs = make([]FileMgr, 0)
	}

	var deepCopyFMgr FileMgr

	deepCopyFMgr,
		err = fMgr.CopyOut(
		ePrefix.XCpy("<-fMgr"))

	if err != nil {
		return err
	}

	fMgrCollection.fileMgrs =
		append(fMgrCollection.fileMgrs, deepCopyFMgr)

	return err
}
