package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileMgrCollectionMechanics struct {
	lock *sync.Mutex
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fMgrColMech *FileMgrCollectionMechanics) addFileMgr(
	fMgrCollection *FileMgrCollection,
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrColMech.lock == nil {
		fMgrColMech.lock = new(sync.Mutex)
	}

	fMgrColMech.lock.Lock()

	defer fMgrColMech.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileMgrCollectionMechanics."+
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
