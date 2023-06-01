package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// fileMgrHelperBoson
//
// Provides helper methods for Type
// fileMgrHelper.
type fileMgrHelperBoson struct {
	lock *sync.Mutex
}

// emptyFileMgr
//
// Helper method designed to "empty" or set all data
// fields in an instance of FileMgr to their zero or
// initial values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values contained in
//	input parameter 'fMgr' will be deleted and reset to
//	their zero values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. All the
//		internal member variables contained in this
//		instance will be deleted and reset to their zero
//		values.
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
func (fMgrHlprBoson *fileMgrHelperBoson) emptyFileMgr(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprBoson.lock == nil {
		fMgrHlprBoson.lock = new(sync.Mutex)
	}

	fMgrHlprBoson.lock.Lock()

	defer fMgrHlprBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperBoson."+
			"emptyFileMgr()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	fMgr.isInitialized = false
	fMgr.dMgr = DirMgr{}
	fMgr.originalPathFileName = ""
	fMgr.absolutePathFileName = ""
	fMgr.isAbsolutePathFileNamePopulated = false
	fMgr.doesAbsolutePathFileNameExist = false
	fMgr.fileName = ""
	fMgr.isFileNamePopulated = false
	fMgr.fileExt = ""
	fMgr.isFileExtPopulated = false
	fMgr.fileNameExt = ""
	fMgr.isFileNameExtPopulated = false
	fMgr.filePtr = nil
	fMgr.isFilePtrOpen = false
	fMgr.fileAccessStatus.Empty()
	fMgr.actualFileInfo = FileInfoPlus{}
	fMgr.fileBufRdr = nil
	fMgr.fileBufWriter = nil
	fMgr.fileBytesWritten = 0
	fMgr.buffBytesWritten = 0
	fMgr.fileRdrBufSize = 0
	fMgr.fileWriterBufSize = 0

	return nil
}

// equalFileMgrs
//
// This method receives two instances of FileMgr and
// proceeds to compare the internal data values to
// determine if they are equal in all respects.
//
// If all internal data values in the two instances are
// equivalent, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrOne						*FileMgr
//
//		A pointer to an incoming instance of FileMgr.
//		The data fields contained in this instance will
//		be compared to corresponding data fields in the
//		FileMgr instance passed as input parameter
//		'fMgrTwo'. If all data fields are equivalent in
//		all respects, this method returns a boolean value
//		of 'true'.
//
//	fMgrTwo						*FileMgr
//
//		A pointer to an incoming instance of FileMgr.
//		The data fields contained in this instance will
//		be compared to corresponding data fields in the
//		FileMgr instance passed as input parameter
//		'fMgrOne'. If all data fields are equivalent in
//		all respects, this method returns a boolean value
//		of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the corresponding data fields contained in
//		input parameters 'fMgrOne' and 'fMgrTwo' are
//		equivalent,	this return value is set to 'true'.
//
//		If the data fields in these two FileMgr instances
//		are not equal in all respects, this return value
//		is set to 'false'.
func (fMgrHlprBoson *fileMgrHelperBoson) equalFileMgrs(
	fMgrOne *FileMgr,
	fMgrTwo *FileMgr) bool {

	if fMgrHlprBoson.lock == nil {
		fMgrHlprBoson.lock = new(sync.Mutex)
	}

	fMgrHlprBoson.lock.Lock()

	defer fMgrHlprBoson.lock.Unlock()

	if fMgrOne == nil ||
		fMgrTwo == nil {

		return false
	}

	if fMgrOne.isInitialized != fMgrTwo.isInitialized ||
		fMgrOne.originalPathFileName != fMgrTwo.originalPathFileName ||
		fMgrOne.isAbsolutePathFileNamePopulated != fMgrTwo.isAbsolutePathFileNamePopulated ||
		fMgrOne.doesAbsolutePathFileNameExist != fMgrTwo.doesAbsolutePathFileNameExist ||
		fMgrOne.absolutePathFileName != fMgrTwo.absolutePathFileName ||
		fMgrOne.fileName != fMgrTwo.fileName ||
		fMgrOne.isFileNamePopulated != fMgrTwo.isFileNamePopulated ||
		fMgrOne.fileExt != fMgrTwo.fileExt ||
		fMgrOne.isFileExtPopulated != fMgrTwo.isFileExtPopulated ||
		fMgrOne.fileNameExt != fMgrTwo.fileNameExt ||
		fMgrOne.isFileNameExtPopulated != fMgrTwo.isFileNameExtPopulated ||
		fMgrOne.filePtr != fMgrTwo.filePtr ||
		fMgrOne.isFilePtrOpen != fMgrTwo.isFilePtrOpen ||
		fMgrOne.fileRdrBufSize != fMgrTwo.fileRdrBufSize ||
		fMgrOne.fileWriterBufSize != fMgrTwo.fileWriterBufSize {

		return false
	}

	if !fMgrOne.fileAccessStatus.Equal(&fMgrTwo.fileAccessStatus) {

		return false
	}

	if !fMgrOne.dMgr.Equal(&fMgrTwo.dMgr) {

		return false
	}

	if !new(dirMgrHelper).equal(
		&fMgrOne.dMgr,
		&fMgrTwo.dMgr) {

		return false
	}

	if !fMgrOne.actualFileInfo.Equal(&fMgrTwo.actualFileInfo) {

		return false
	}

	return true
}

// flushBytesToDisk
//
//	Helper method which is designed to flush all buffers
//	and write all data in memory to the file identified
//	by the instance of FileMgr passed as input paramter
//	'fMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr.
//
//		This method will flush all buffers and write all
//	 	data in memory to the file identified by 'fMgr'.
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
func (fMgrHlprBoson *fileMgrHelperBoson) flushBytesToDisk(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprBoson.lock == nil {
		fMgrHlprBoson.lock = new(sync.Mutex)
	}

	fMgrHlprBoson.lock.Lock()

	defer fMgrHlprBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperBoson."+
			"flushBytesToDisk()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: FileMgr instance is invalid!\n"+
			"Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var errs = make([]error, 0)

	var err2, err3 error

	if fMgr.filePtr != nil &&
		fMgr.fileBufWriter != nil {

		err3 = fMgr.fileBufWriter.Flush()

		if err3 != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned from fMgr.fileBufWriter."+
				"Flush().\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err3.Error())

			errs = append(errs, err2)
		}
	}

	if fMgr.filePtr != nil &&
		(fMgr.fileBytesWritten > 0 ||
			fMgr.buffBytesWritten > 0) {

		err3 = fMgr.filePtr.Sync()

		if err3 != nil {
			err2 = fmt.Errorf("%v\n"+
				"Error returned from fMgr.filePtr.Sync()\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err3.Error())

			errs = append(errs, err2)
		}
	}

	return new(StrMech).ConsolidateErrors(errs)
}
