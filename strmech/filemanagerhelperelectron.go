package strmech

import (
	"bufio"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
	"time"
)

// fileMgrHelperElectron
//
// Provides helper methods for Type
// fileMgrHelper.
type fileMgrHelperElectron struct {
	lock *sync.Mutex
}

// lowLevelCloseFile - Helper method for the File
// Manager Type.
//
// This method Closes the file identified by 'fMgr'.
//
// This method does check to verify that the file
// does exist.
//
// This method will call "fMgrHlpr.flushBytesToDisk()"
func (fMgrHlprElectron *fileMgrHelperElectron) lowLevelCloseFile(
	fMgr *FileMgr,
	fMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprElectron.lock == nil {
		fMgrHlprElectron.lock = new(sync.Mutex)
	}

	fMgrHlprElectron.lock.Lock()

	defer fMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileMgrHelperElectron." +
		"lowLevelCloseFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			fMgrLabel)
	}

	errs := make([]error, 0)
	var err2, err3 error

	if fMgr.filePtr == nil {

		fMgr.isFilePtrOpen = false
		fMgr.fileAccessStatus.Empty()

		fMgr.fileBufRdr = nil
		fMgr.fileBufWriter = nil
		fMgr.fileBytesWritten = 0
		fMgr.buffBytesWritten = 0

		return nil
	}

	// fMgr.filePtr != nil
	err2 = new(fileMgrHelperBoson).
		flushBytesToDisk(fMgr, ePrefix)

	if err2 != nil {
		errs = append(errs, err2)
	}

	if fMgr.filePtr != nil {

		err3 = fMgr.filePtr.Close()

		if err3 != nil {
			err2 = fmt.Errorf("%v\n"+
				"Error returned by %v.filePtr.Close()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				fMgrLabel,
				err3.Error())

			errs = append(errs, err2)
		}
	}

	fMgr.filePtr = nil
	fMgr.isFilePtrOpen = false
	fMgr.fileAccessStatus.Empty()
	fMgr.fileBufRdr = nil
	fMgr.fileBufWriter = nil
	fMgr.fileBytesWritten = 0
	fMgr.buffBytesWritten = 0

	return new(StrMech).ConsolidateErrors(errs)
}

// lowLevelDeleteFile - Helper method which is designed
// to delete the file identified by input parameter, 'fMgr'.
//
// This method performs no validation on 'fMgr' and does
// not verify existence of the 'fMgr' file.
func (fMgrHlprElectron *fileMgrHelperElectron) lowLevelDeleteFile(
	targetFMgr *FileMgr,
	targetFileLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprElectron.lock == nil {
		fMgrHlprElectron.lock = new(sync.Mutex)
	}

	fMgrHlprElectron.lock.Lock()

	defer fMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperElectron."+
			"lowLevelDeleteFile()",
		"")

	if err != nil {
		return err
	}

	if targetFMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	for i := 0; i < 3; i++ {

		err2 := os.Remove(targetFMgr.absolutePathFileName)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"\nError: os.Remove(%v.absolutePathFileName) "+
				"returned an error.\n"+
				"%v='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				targetFileLabel,
				targetFileLabel,
				targetFMgr.absolutePathFileName,
				err2.Error())

		} else {

			return nil
		}

		time.Sleep(50 * time.Millisecond)
	}

	return err
}

// lowLevelDoesFileExist - Tests input parameter
// 'pathFileName' to determine if the file exists
// on disk.
//
// The method makes three calls to os.Stat(pathFileName)
// to ensure that non-path errors are real and not due
// to operating system timing issues.
func (fMgrHlprElectron *fileMgrHelperElectron) lowLevelDoesFileExist(
	pathFileName string,
	directoryPath string,
	errPrefDto *ePref.ErrPrefixDto,
	pathFileNameLabel string,
	directoryPathLabel string) (
	pathFileDoesExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if fMgrHlprElectron.lock == nil {
		fMgrHlprElectron.lock = new(sync.Mutex)
	}

	fMgrHlprElectron.lock.Lock()

	defer fMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	pathFileDoesExist = false

	funcName := "fileMgrHelperElectron." +
		"lowLevelDoesFileExist()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return pathFileDoesExist, fInfoPlus, err
	}

	if len(pathFileNameLabel) == 0 {
		pathFileNameLabel = "pathFileName"
	}

	if len(directoryPathLabel) == 0 {
		directoryPathLabel = "dirPath"
	}

	errCode := 0

	fHelpElectron := new(fileHelperElectron)

	errCode,
		_,
		pathFileName = fHelpElectron.
		isStringEmptyOrBlank(pathFileName)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter %v is an EMPTY STRING!\n",
			ePrefix.String(),
			pathFileNameLabel)

		return pathFileDoesExist, fInfoPlus, err
	}

	errCode,
		_,
		directoryPath = fHelpElectron.
		isStringEmptyOrBlank(directoryPath)

	var err2 error
	var info os.FileInfo

	for i := 0; i < 3; i++ {

		pathFileDoesExist = false
		fInfoPlus = FileInfoPlus{}
		err = nil

		info, err2 = os.Stat(pathFileName)

		if err2 != nil {

			if os.IsNotExist(err2) {

				pathFileDoesExist = false
				fInfoPlus = FileInfoPlus{}
				err = nil
				return pathFileDoesExist, fInfoPlus, err
			}

			// err == nil and err != os.IsNotExist(err)
			// This is a non-path error. The non-path error will be
			// tested up to 3-times before it is returned.
			err = fmt.Errorf("%v\n"+
				"Non-Path error returned by os.Stat(pathFileName)\n"+
				"%v= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				pathFileNameLabel,
				pathFileName,
				err2.Error())

			fInfoPlus = FileInfoPlus{}

			pathFileDoesExist = false

		} else {
			// err == nil
			// The path really does exist!
			pathFileDoesExist = true
			err = nil

			if errCode < 0 {
				// If the directoryPath is an empty string,
				// only record the os.FileInfo data, not the path.
				//
				fInfoPlus = new(FileInfoPlus).
					NewFromFileInfo(info)

			} else {

				// directoryPath is a string with length greater than zero,
				// record the os.FileInfo data AND the path.
				fInfoPlus,
					err2 = new(FileInfoPlus).
					NewFromPathFileInfo(directoryPath, info)

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Error returned by FileInfoPlus{}.NewFromPathFileInfo(directoryPath, info)\n"+
						"directoryPath= '%v'\n"+
						"Error= \n%v\n",
						ePrefix,
						directoryPath,
						err2.Error())

					fInfoPlus = FileInfoPlus{}
				}
			}

			return pathFileDoesExist, fInfoPlus, err
		}

		time.Sleep(30 * time.Millisecond)
	}

	return pathFileDoesExist, fInfoPlus, err
}

// lowLevelOpenFile - Helper method which is designed
// to open a file. Unlike similar methods, this method
// does not check for the existence of the file, does
// not try to close an existing file and will not
// create the directory path.
func (fMgrHlprElectron *fileMgrHelperElectron) lowLevelOpenFile(
	fMgr *FileMgr,
	fileAccessCtrl FileAccessControl,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprElectron.lock == nil {
		fMgrHlprElectron.lock = new(sync.Mutex)
	}

	fMgrHlprElectron.lock.Lock()

	defer fMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileMgrHelperElectron." +
		"lowLevelOpenFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			&fileAccessCtrl,
			ePrefix.XCpy(
				"fileAccessCtrl"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Parameter 'fileAccessCtrl' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	var fOpenParam int

	var fPermParam os.FileMode

	fOpenParam,
		fPermParam,
		err =
		fileAccessCtrl.GetFileOpenAndPermissionCodes(
			ePrefix.XCpy("fileAccessCtrl"))

	if err != nil {

		fMgr.fileAccessStatus.Empty()

		return err
	}

	fMgr.fileAccessStatus = fileAccessCtrl.CopyOut()

	fOpenType, err2 := fMgr.fileAccessStatus.GetFileOpenType(
		ePrefix.XCpy("fMgr.fileAccessStatus"))

	if err2 != nil {
		return fmt.Errorf(funcName+
			"\n%v\n", err2.Error())
	}

	fMgr.filePtr, err =
		os.OpenFile(
			fMgr.absolutePathFileName,
			fOpenParam,
			fPermParam)

	if err != nil {

		fMgr.filePtr = nil
		fMgr.isFilePtrOpen = false
		fMgr.fileAccessStatus.Empty()
		fMgr.fileBufRdr = nil
		fMgr.fileBufWriter = nil
		fMgr.fileBytesWritten = 0
		fMgr.buffBytesWritten = 0

		return fmt.Errorf("%v\n"+
			"Error opening file from os.OpenFile():\n"+
			"fMgr.absolutePathFileName= '%v'\n"+
			"Error= '%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName, err.Error())
	}

	if fMgr.filePtr == nil {
		return fmt.Errorf("%v\n"+
			"Error: os.OpenFile() returned a 'nil' file pointer!\n"+
			"FileMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	fMgr.isFilePtrOpen = true

	if fOpenType == FOpenType.TypeReadWrite() {

		if fMgr.fileBufWriter == nil {
			if fMgr.fileWriterBufSize > 0 {
				fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
			} else {
				fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
			}
		}

		if fMgr.fileBufRdr == nil {
			if fMgr.fileRdrBufSize > 0 {
				fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
			} else {
				fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
			}
		}
	}

	if fOpenType == FOpenType.TypeWriteOnly() {

		if fMgr.fileBufWriter == nil {
			if fMgr.fileWriterBufSize > 0 {
				fMgr.fileBufWriter = bufio.NewWriterSize(fMgr.filePtr, fMgr.fileWriterBufSize)
			} else {
				fMgr.fileBufWriter = bufio.NewWriter(fMgr.filePtr)
			}
		}
	}

	if fOpenType == FOpenType.TypeReadOnly() {

		if fMgr.fileBufRdr == nil {
			if fMgr.fileRdrBufSize > 0 {
				fMgr.fileBufRdr = bufio.NewReaderSize(fMgr.filePtr, fMgr.fileRdrBufSize)
			} else {
				fMgr.fileBufRdr = bufio.NewReader(fMgr.filePtr)
			}
		}
	}

	return nil
}
