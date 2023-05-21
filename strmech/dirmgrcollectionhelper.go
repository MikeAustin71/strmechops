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

	destinationDMgrCollection.dirMgrs = make([]DirMgr, 0, lenSourceDMgrs)

	var dMgrCopy DirMgr

	for i := 0; i < lenSourceDMgrs; i++ {

		dMgrCopy,
			err = sourceDMgrCollection.dirMgrs[i].CopyOut(ePrefix.XCpy(
			"sourceDMgrCollection.dirMgrs[i]"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error: sourceDMgrCollection.dirMgrs[%v].CopyOut()\n"+
				"sourceDMgrCollection.dirMgrs index= '%v'\n"+
				"sourceDMgrCollection DirMgr= '%v'\n"+
				"Error= \n%v\n",
				funcName,
				i,
				i,
				sourceDMgrCollection.dirMgrs[i].absolutePath,
				err.Error())
		}

		sourceDMgrCollection.dirMgrs =
			append(sourceDMgrCollection.dirMgrs,
				dMgrCopy)
	}

	return err
}
