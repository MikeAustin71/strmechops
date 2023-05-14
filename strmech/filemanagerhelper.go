package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"strings"
	"sync"
)

type fileMgrHelper struct {
	fMgr FileMgr

	lock *sync.Mutex
}

// closeFile
//
// Helper method for File Manager (FileMgr) Type. It is
// designed to close a file. The version of the
// 'file close' operation first checks to verify whether
// the file exists on disk.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. This method
//		will close the open file pointer associated with
//		this instance of FileMgr.
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
func (fMgrHlpr *fileMgrHelper) closeFile(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"closeFile()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	errs := make([]error, 0)

	var err2 error
	fileDoesExist := false

	fileDoesExist,
		err2 = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	if err2 != nil {

		errs = append(errs, err2)

		err2 = nil

		if fMgr.filePtr != nil {
			err2 = new(fileMgrHelperElectron).
				lowLevelCloseFile(fMgr, "fMgr", ePrefix)
		}

		if err2 != nil {
			errs = append(errs, err2)
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

	if !fileDoesExist {

		err2 = nil

		if fMgr.filePtr != nil {
			err2 = fMgr.filePtr.Close()
		}

		if err2 != nil {
			errs = append(errs, err2)
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

	// fileDoesExist == true
	return new(fileMgrHelperElectron).
		lowLevelCloseFile(fMgr, "fMgr", ePrefix)
}

// copyFileSetup
//
// Helper method designed to perform setup and
// validation checks before initiating a copy operation.
//
// This method is designed to be called immediately
// before method fileMgrHelper.lowLevelCopyByIO() and
// fileMgrHelper.lowLevelCopyByLink().
func (fMgrHlpr *fileMgrHelper) copyFileSetup(
	srcFMgr *FileMgr,
	destFMgr *FileMgr,
	createDestinationDir bool,
	deleteExistingDestFile bool,
	errPrefDto *ePref.ErrPrefixDto,
	srcFMgrLabel string,
	destFMgrLabel string) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper.copyFileSetup()",
		"")

	if err != nil {
		return err
	}

	if srcFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			srcFMgrLabel)
	}

	if destFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			destFMgrLabel)
	}

	var fMgrFileDoesExist, targetFMgrFileDoesExist bool

	fMgrFileDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			srcFMgr,
			PreProcPathCode.None(),
			ePrefix,
			srcFMgrLabel+".absolutePathFileName")

	if err != nil {
		return err
	}

	if !fMgrFileDoesExist {
		return fmt.Errorf("%v\n"+
			"Error: Source File ('%v') DOES NOT EXIST!\n"+
			"Source File ('%v')='%v'\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName)
	}

	err = new(fileMgrHelperElectron).
		lowLevelCloseFile(srcFMgr, srcFMgrLabel, ePrefix)

	if err != nil {
		return err
	}

	if !srcFMgr.actualFileInfo.Mode().IsRegular() {

		return fmt.Errorf("%v\n"+
			"Error: %v file is NOT a 'Regular' File!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName)
	}

	if srcFMgr.actualFileInfo.Mode().IsDir() {

		return fmt.Errorf("%v\n"+
			"Error: %v is NOT a file! It is a DIRECTORY!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName)
	}

	targetFMgrFileDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		destFMgr,
		PreProcPathCode.None(),
		ePrefix,
		destFMgrLabel+".absolutePathFileName")

	if err != nil {
		return err
	}

	if fMgrFileDoesExist && targetFMgrFileDoesExist {

		if os.SameFile(srcFMgr.actualFileInfo.origFileInfo,
			destFMgr.actualFileInfo.origFileInfo) ||
			strings.ToLower(srcFMgr.absolutePathFileName) ==
				strings.ToLower(destFMgr.absolutePathFileName) {

			return fmt.Errorf("%v\n"+
				"Error: %v FileInfo is same as %v FileInfo.\n"+
				"These two files are one in the same file!\n"+
				"%v='%v'\n"+
				"%v='%v'\n",
				ePrefix.String(),
				srcFMgrLabel,
				destFMgrLabel,
				srcFMgrLabel,
				srcFMgr.absolutePathFileName,
				destFMgrLabel,
				destFMgr.absolutePathFileName)
		}

	} else if strings.ToLower(srcFMgr.absolutePathFileName) ==
		strings.ToLower(destFMgr.absolutePathFileName) {

		return fmt.Errorf("%v\n"+
			"Error: %v path file name is same as %v path file name.\n"+
			"These two files are one in the same file!\n"+
			"%v='%v'\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcFMgrLabel,
			destFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName)
	}

	err = new(fileMgrHelperElectron).
		lowLevelCloseFile(destFMgr, destFMgrLabel, ePrefix)

	if err != nil {
		return err
	}

	if targetFMgrFileDoesExist {

		if destFMgr.actualFileInfo.Mode().IsDir() {
			return fmt.Errorf("%v\n"+
				"Error: %v is NOT A FILE! IT IS A DIRECTORY!!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				destFMgrLabel,
				destFMgrLabel,
				destFMgr.absolutePathFileName)
		}

		if !destFMgr.actualFileInfo.Mode().IsRegular() {

			return fmt.Errorf("%v\n"+
				"Error: %v exists and is NOT a 'regular' file.\n"+
				"%v='%v'\n",
				ePrefix.String(),
				destFMgrLabel,
				destFMgrLabel,
				destFMgr.absolutePathFileName)

		}

		if deleteExistingDestFile {

			err = new(fileMgrHelperElectron).
				lowLevelDeleteFile(destFMgr, destFMgrLabel, ePrefix)

			if err != nil {
				return err
			}
		}

	} else {

		targetFMgrFileDoesExist, err = destFMgr.dMgr.
			DoesThisDirectoryExist(ePrefix.XCpy(
				"destFMgr.dMgr"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error returned by %v.dMgr.DoesThisDirectoryExist()\n"+
				"%v='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				destFMgrLabel,
				destFMgrLabel,
				destFMgr.absolutePathFileName,
				err.Error())
		}

		if !targetFMgrFileDoesExist && createDestinationDir {

			err = destFMgr.dMgr.MakeDir(
				ePrefix.XCpy(
					"destFMgr.dMgr"))

			if err != nil {

				return fmt.Errorf("%v\n"+
					"Error returned by %v.dMgr.MakeDir()\n"+
					"%v.dMgr='%v'\n"+
					"Error='%v'\n",
					ePrefix.String(),
					destFMgrLabel,
					destFMgrLabel,
					destFMgr.absolutePathFileName,
					err.Error())
			}

		} else if !targetFMgrFileDoesExist && !createDestinationDir {

			return fmt.Errorf("%v\n"+
				"Error: %v directory DOES NOT EXIST!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				destFMgrLabel,
				destFMgrLabel,
				destFMgr.absolutePathFileName)
		}
	}

	return nil
}

// createDirectory
//
// Creates a Directory
func (fMgrHlpr *fileMgrHelper) createDirectory(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"createDirectory()",
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
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	err = fMgr.dMgr.MakeDir(
		ePrefix.XCpy(
			"fMgr.dMgr"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	return nil
}

// createFile
//
// FileMgr helper method. Creates a file identified by
// input parameter 'fMgr'.
//
// The file open operation uses create/truncate open
// codes.
//
// If input parameter 'createTheDirectory' is set to
// 'true', this method will create the directory if it
// does not previously exist.
func (fMgrHlpr *fileMgrHelper) createFile(
	fMgr *FileMgr,
	createTheDirectory bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"createFile()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	var createTruncateAccessCtrl FileAccessControl

	createTruncateAccessCtrl,
		err =
		new(FileAccessControl).NewReadWriteCreateTruncateAccess(
			ePrefix.XCpy("createTruncateAccessCtrl<-"))

	if err != nil {
		return err
	}

	err = createTruncateAccessCtrl.
		IsValidInstanceError(ePrefix.XCpy(
			"createTruncateAccessCtrl"))

	if err != nil {
		return err
	}

	_,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	var directoryPathDoesExist bool

	directoryPathDoesExist,
		err =
		fMgr.dMgr.DoesThisDirectoryExist(
			ePrefix.XCpy(
				"fMgr.dMgr"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Non-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
			"fMgr.dMgr='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath,
			err.Error())
	}

	if !directoryPathDoesExist && createTheDirectory {

		err = fMgr.dMgr.MakeDir(
			ePrefix.XCpy(
				"fMgr.dMgr"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error returned by fMgr.dMgr.MakeDir().\n"+
				"fMgr.dMgr='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				fMgr.dMgr.absolutePath,
				err.Error())
		}

	} else if !directoryPathDoesExist && !createTheDirectory {

		return fmt.Errorf("%v\n"+
			"Error: File Manager Directory Path DOES NOT EXIST!\n"+
			"fMgr.dMgr='%v'\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath)
	}

	err = new(fileMgrHelperElectron).
		lowLevelCloseFile(fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	return new(fileMgrHelperElectron).
		lowLevelOpenFile(
			fMgr,
			createTruncateAccessCtrl,
			ePrefix)
}

// deleteFile - Helper method used to delete
// the file identified by 'fMgr'.
func (fMgrHlpr *fileMgrHelper) deleteFile(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileMgrHelper." +
		"deleteFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"deleteFile()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	var pathFileNameDoesExist bool

	fMgrHelperAtom := new(fileMgrHelperAtom)

	pathFileNameDoesExist,
		err = fMgrHelperAtom.
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	fMgrHelperElectron := new(fileMgrHelperElectron)

	if err != nil {

		_ = fMgrHelperElectron.
			lowLevelCloseFile(fMgr, "fMgr", ePrefix)

		return err
	}

	if !pathFileNameDoesExist {
		_ = fMgrHelperElectron.
			lowLevelCloseFile(fMgr, "fMgr", ePrefix)
		return nil
	}

	err = fMgrHelperElectron.
		lowLevelCloseFile(fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	err = fMgrHelperElectron.
		lowLevelDeleteFile(fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	pathFileNameDoesExist,
		err = fMgrHelperAtom.
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Non-Path error returned after file deletion.\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	if pathFileNameDoesExist {

		return fmt.Errorf("%v\n"+
			"Error: Attempted file deletion FAILED!.\n"+
			"File still exists.\n"+
			"fMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	return nil
}

// lowLevelCopyByIOBuffer
//
// Helper method which implements the "Copy IO"
// operation. Specifically, this method calls
// 'io.CopyBuffer'.
//
// This is a low level method which performs no set up or
// validation services before initiating the "Copy By IO"
// operation.
//
// If the input parameter 'localCopyBuffer' is an array
// of bytes ([]byte) supplied by the calling function.
// The copy operation stages through the provided buffer
// (if one is required) rather than allocating a
// temporary one.
//
// If the length of 'localCopyBuffer' is zero or if
// 'localCopyBuffer' is nil, this method will return an
// error.
//
// The best practice is to call fileMgrHelper.copyFileSetup()
// or similar method first, before calling this method.
//
//	Reference:
//	  https://golang.org/pkg/io/#CopyBuffer
//	  https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func (fMgrHlpr *fileMgrHelper) lowLevelCopyByIOBuffer(
	srcFMgr *FileMgr,
	destFMgr *FileMgr,
	localCopyBuffer []byte,
	errPrefDto *ePref.ErrPrefixDto,
	srcFMgrLabel string,
	destFMgrLabel string) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"lowLevelCopyByIOBuffer()",
		"")

	if err != nil {
		return err
	}

	if srcFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			srcFMgrLabel)
	}

	if destFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			destFMgrLabel)
	}

	if localCopyBuffer == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'localCopyBuffer' is nil!\n",
			ePrefix.String())
	}

	if len(localCopyBuffer) == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'localCopyBuffer' is a ZERO Length Array!\n",
			ePrefix.String())
	}

	if len(srcFMgrLabel) == 0 {

		srcFMgrLabel = "srcFMgr"
	}

	if len(destFMgrLabel) == 0 {

		destFMgrLabel = "destFMgr"
	}

	var targetFMgrFileDoesExist bool

	// First, open the source file

	var srcPtr *os.File

	srcPtr,
		err = os.Open(srcFMgr.absolutePathFileName)

	if err != nil {

		if srcPtr != nil {
			_ = srcPtr.Close()
		}

		return fmt.Errorf("%v\n"+
			"Error returned by os.Open(%v.absolutePathFileName)\n"+
			"%v.absolutePathFileName='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			err.Error())
	}

	// Second, open the destination file
	var destPtr *os.File

	destPtr,
		err = os.Create(destFMgr.absolutePathFileName)

	if err != nil {

		_ = srcPtr.Close()

		if destPtr != nil {
			_ = destPtr.Close()
		}

		return fmt.Errorf("%v\n"+
			"Error returned by os.Create(%v.absolutePathFileName)\n"+
			"%v.absolutePathFileName='%v'\n"+
			"Error= \n%v'\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error())
	}

	var bytesCopied int64

	bytesCopied,
		err = io.CopyBuffer(
		destPtr,
		srcPtr,
		localCopyBuffer)

	if err != nil {
		_ = srcPtr.Close()
		_ = destPtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned by io.CopyBuffer(%v, %v)\n"+
			"%v='%v'\n%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error())
	}

	errs := make([]error, 0, 5)

	err = destPtr.Sync()

	if err != nil {

		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by %v.Sync()\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error()))
	}

	err = srcPtr.Close()

	if err != nil {

		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by srcPtr.Close()\n"+
			"srcPtr=%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			err.Error()))
	}

	srcPtr = nil

	err = destPtr.Close()

	if err != nil {

		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by destPtr.Close()\n"+
			"destPtr=%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error()))
	}

	destPtr = nil

	if len(errs) > 0 {
		return new(StrMech).ConsolidateErrors(errs)
	}

	targetFMgrFileDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		destFMgr,
		PreProcPathCode.None(),
		ePrefix.XCpy("destFMgr"),
		destFMgrLabel+".absolutePathFileName")

	if err != nil {
		return fmt.Errorf("%v\n"+
			"After Copy IO operation, %v "+
			"generated non-path error!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			err.Error())
	}

	if !targetFMgrFileDoesExist {

		return fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = %v='%v'\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgr.absolutePathFileName)
	}

	srcFileSize := srcFMgr.actualFileInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file!\n"+
			"Source File Bytes='%v'   Bytes Coped='%v'\n"+
			"Source File=%v='%v'\nDestination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName)

		return err
	}

	if destFMgr.actualFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File Bytes='%v'   Destination File Bytes='%v'\n"+
			"Source File=%v='%v'\nDestination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			destFMgr.actualFileInfo.Size(),
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName)

		return err
	}

	return nil
}

// lowLevelCopyByIO
//
// Helper method which implements the "Copy IO" operation.
//
// This is a low level method which performs no set up
// validation before initiating the "Copy By IO" operation.
//
// If the input parameter 'localBufferSize' is greater
// than zero, the copy operation will be implemented
// using a local buffer of type []byte. The size of
// this []byte array is therefore specified by
// 'localBufferSize'.
//
// The best practice is to call
// fileMgrHelper.copyFileSetup() or similar method first,
// before calling this method.
//
//	Reference:
//	  https://golang.org/pkg/io/#CopyBuffer
//	  https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func (fMgrHlpr *fileMgrHelper) lowLevelCopyByIO(
	srcFMgr *FileMgr,
	destFMgr *FileMgr,
	localBufferSize int,
	errPrefDto *ePref.ErrPrefixDto,
	srcFMgrLabel string,
	destFMgrLabel string) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"fileMgrHelper()",
		"")

	if err != nil {
		return err
	}

	if srcFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			srcFMgrLabel)
	}

	if destFMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			destFMgrLabel)
	}

	if len(srcFMgrLabel) == 0 {

		srcFMgrLabel = "srcFMgr"
	}

	if len(destFMgrLabel) == 0 {

		destFMgrLabel = "destFMgr"
	}

	var targetFMgrFileDoesExist bool

	// First, open the source file

	var srcPtr *os.File

	srcPtr,
		err = os.Open(srcFMgr.absolutePathFileName)

	if err != nil {

		if srcPtr != nil {
			_ = srcPtr.Close()
		}

		return fmt.Errorf("%v\n"+
			"Error returned by os.Open(%v.absolutePathFileName)\n"+
			"%v.absolutePathFileName='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			err.Error())
	}

	var destPtr *os.File

	destPtr,
		err = os.Create(destFMgr.absolutePathFileName)

	if err != nil {

		_ = srcPtr.Close()

		if destPtr != nil {
			_ = destPtr.Close()
		}

		return fmt.Errorf("%v\n"+
			"Error returned by os.Create(%v.absolutePathFileName)\n"+
			"%v.absolutePathFileName='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error())
	}

	var byteBuff []byte

	if localBufferSize > 0 {
		byteBuff = make([]byte, localBufferSize)
	}

	var bytesCopied int64

	bytesCopied,
		err = io.CopyBuffer(destPtr, srcPtr, byteBuff)

	if err != nil {

		byteBuff = nil

		_ = srcPtr.Close()

		_ = destPtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned by io.CopyBuffer(%v, %v)\n"+
			"%v='%v'\n%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			srcFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error())
	}

	errs := make([]error, 0, 5)

	err = destPtr.Sync()

	if err != nil {

		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by %v.Sync()\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgrLabel, destFMgr.absolutePathFileName,
			err.Error()))
	}

	byteBuff = nil

	err = srcPtr.Close()

	if err != nil {
		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by srcPtr.Close()\n"+
			"srcPtr=%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			err.Error()))
	}

	srcPtr = nil

	err = destPtr.Close()

	if err != nil {

		errs = append(errs, fmt.Errorf("%v\n"+
			"Error returned by destPtr.Close()\n"+
			"destPtr=%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error()))
	}

	destPtr = nil

	if len(errs) > 0 {
		return new(StrMech).ConsolidateErrors(errs)
	}

	targetFMgrFileDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		destFMgr,
		PreProcPathCode.None(),
		ePrefix,
		destFMgrLabel+".absolutePathFileName")

	if err != nil {
		return fmt.Errorf("%v\n"+
			"After Copy IO operation, %v "+
			"generated non-path error!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			err.Error())
	}

	if !targetFMgrFileDoesExist {

		return fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = %v='%v'\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgr.absolutePathFileName)
	}

	srcFileSize := srcFMgr.actualFileInfo.Size()

	if bytesCopied != srcFileSize {
		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file!\n"+
			"Source File Bytes='%v'   Bytes Coped='%v'\n"+
			"Source File=%v='%v'\nDestination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName)

		return err
	}

	if destFMgr.actualFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File Bytes='%v'   Destination File Bytes='%v'\n"+
			"Source File=%v='%v'\nDestination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			destFMgr.actualFileInfo.Size(),
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName)

		return err
	}

	return nil
}

// lowLevelCopyByLink
//
// Helper method designed to copy a file by 'link' using
// method os.Link(src, dst).
//
// "os.Link(src, dst)" is the only method employed to
// copy a designated file. If "os.Link(src, dst)" fails,
// an error is returned.
//
// By creating a 'linked' file, changing the contents of
// one file will be reflected in the second. The two
// linked files are 'mirrors' of each other.
//
// If the 'mirror' feature causes concerns, use the
// "Copy By IO' technique employed by
// fileMgrHelper.lowLevelCopyByIO().
//
// Be advised that this method does not perform parameter
// validation or setup services prior to initiating the
// 'Copy By Link' operation. It would be prudent to call
// 'fileMgrHelper.copyFileSetup()' first, before calling
// this method.
//
//	Reference
//	 https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	 https://golang.org/pkg/os/#Link
func (fMgrHlpr *fileMgrHelper) lowLevelCopyByLink(
	srcFMgr *FileMgr,
	destFMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto,
	srcFMgrLabel string,
	destFMgrLabel string) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper.lowLevelCloseFile()",
		"")

	if err != nil {
		return err
	}

	if srcFMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			srcFMgrLabel)
	}

	if destFMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			destFMgr)
	}

	if len(srcFMgrLabel) == 0 {

		srcFMgrLabel = "srcFMgr"
	}

	if len(destFMgrLabel) == 0 {

		destFMgrLabel = "destFMgr"
	}

	err = os.Link(
		srcFMgr.absolutePathFileName,
		destFMgr.absolutePathFileName)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by os.Link("+
			"%v.absolutePathFileName, %v.absolutePathFileName)\n"+
			"%v.absolutePathFileName=%v\n"+
			"%v.absolutePathFileName=%v\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFMgrLabel,
			destFMgrLabel,
			srcFMgrLabel,
			srcFMgr.absolutePathFileName,
			destFMgrLabel,
			destFMgr.absolutePathFileName,
			err.Error())
	}

	var destinationFileDoesExist bool

	destinationFileDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		destFMgr,
		PreProcPathCode.None(),
		ePrefix,
		destFMgrLabel+".absolutePathFileName")

	if err != nil {
		return fmt.Errorf("%v\n"+
			"After Copy By Link operation, %v "+
			"generated non-path error!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			destFMgrLabel,
			err.Error())
	}

	if !destinationFileDoesExist {
		return fmt.Errorf("%v\n"+
			"ERROR: After Copy By Link operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = %v='%v'\n",
			ePrefix.String(),
			destFMgrLabel,
			destFMgr.absolutePathFileName)
	}

	return nil
}

// moveFile
//
// Helper method designed to move a file to a new
// destination. The operation is performed in two steps.
// First, the source file is copied to a destination
// specified by parameter, 'targetFMgr'. Second, if the
// copy operation is successful, the source file ('fMgr')
// will be deleted.
func (fMgrHlpr *fileMgrHelper) moveFile(
	fMgr *FileMgr,
	targetFMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper.moveFile()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	if targetFMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetFMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	var sourceFilePathDoesExist, targetFilePathDoesExist bool

	fHelperAtom := new(fileMgrHelperAtom)

	sourceFilePathDoesExist,
		err = fHelperAtom.doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	fHelperElectron := new(fileMgrHelperElectron)

	err = fHelperElectron.lowLevelCloseFile(
		fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	if !sourceFilePathDoesExist {
		return fmt.Errorf("%v\n"+
			"Error: Source File Manager 'fMgr' DOES NOT EXIST!\n"+
			"fMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	if fMgr.actualFileInfo.Mode().IsDir() {
		return fmt.Errorf("%v\n"+
			"Error: Source File 'fMgr' is NOT A FILE!\n"+
			"IT IS A DIRECTORY!!\n"+
			"fMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	if !fMgr.actualFileInfo.Mode().IsRegular() {
		return fmt.Errorf("%v\n"+
			"Error: Source File Invalid. "+
			"The Source File 'fMgr' IS NOT A REGULAR FILE!\n"+
			"fMgr='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)
	}

	targetFilePathDoesExist,
		err = fHelperAtom.doesFileMgrPathFileExist(
		targetFMgr,
		PreProcPathCode.None(),
		ePrefix,
		"targetFMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	err = fHelperElectron.lowLevelCloseFile(
		targetFMgr, "targetFMgr", ePrefix)

	if err != nil {
		return err
	}

	if !targetFilePathDoesExist {

		err = targetFMgr.dMgr.MakeDir(
			ePrefix.XCpy(
				"targetFMgr.dMgr"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error returned by targetFMgr."+
				"dMgr.MakeDir()\n"+
				"targetFMgr.dMgr='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				targetFMgr.dMgr.absolutePath,
				err.Error())
		}

	} else {
		// targetFilePathDoesExist == true
		err = fHelperElectron.lowLevelDeleteFile(
			targetFMgr, "targetFMgr", ePrefix)

		if err != nil {
			return err
		}

	}

	// Close the source file
	err = fHelperElectron.lowLevelCloseFile(fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	err = fHelperElectron.lowLevelCloseFile(
		targetFMgr, "targetFMgr", ePrefix)

	if err != nil {
		return err
	}

	// Now, open the source file
	srcFilePtr, err := os.Open(fMgr.absolutePathFileName)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned from os.Open(fMgr.absolutePathFileName).\n"+
			"fMgr='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			fMgr.absolutePathFileName,
			err.Error())
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outFilePtr, err := os.Create(targetFMgr.absolutePathFileName)

	if err != nil {

		_ = srcFilePtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned from os.Create("+
			"targetFMgr.absolutePathFileName)\n"+
			"targetFMgr='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			targetFMgr.absolutePathFileName,
			err.Error())
	}

	bytesToRead := fMgr.actualFileInfo.Size()

	bytesCopied, err := io.Copy(outFilePtr, srcFilePtr)

	if err != nil {

		_ = srcFilePtr.Close()
		_ = outFilePtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned from io.Copy("+
			"outFilePtr, srcFilePtr).\n"+
			"outFile='%v'\n"+
			"srcFile='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			targetFMgr.absolutePathFileName,
			fMgr.absolutePathFileName,
			err.Error())
	}

	errs := make([]error, 0, 10)

	err = outFilePtr.Sync()

	if err != nil {
		errs = append(errs, fmt.Errorf("%v\n"+
			"Error flushing buffers!\n"+
			"Error returned from outFilePtr.Sync()\n"+
			"outFilePtr=targetFMgr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			targetFMgr.absolutePathFileName,
			err.Error()))
	}

	err = srcFilePtr.Close()

	if err != nil {
		errs = append(errs, err)
	}

	err = outFilePtr.Close()

	if err != nil {
		errs = append(errs, err)
	}

	if bytesToRead != bytesCopied {
		err = fmt.Errorf("%v\n"+
			"Error: Bytes copied not equal bytes "+
			"in source file!\n"+
			"Bytes To Read='%v'   Bytes Copied='%v'\n"+
			"Source File 'fMgr'='%v'\n"+
			"Target File targetFMgr='%v'\n",
			ePrefix.String(),
			bytesToRead,
			bytesCopied,
			fMgr.absolutePathFileName,
			targetFMgr.absolutePathFileName)

		errs = append(errs, err)
	}

	strMech := new(StrMech)

	if len(errs) > 0 {

		return strMech.ConsolidateErrors(errs)
	}

	err = fHelperElectron.lowLevelDeleteFile(fMgr, "fMgr", ePrefix)

	if err != nil {
		return err
	}

	sourceFilePathDoesExist,
		err = fHelperAtom.doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix,
		"#2 fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	errs = make([]error, 0, 10)

	if sourceFilePathDoesExist {
		err = fmt.Errorf("%v\n"+
			"Error: Move operation failed. Source File was copied, \n"+
			"but was NOT deleted after copy operation.\n"+
			"Source File 'fMgr'='%v'",
			ePrefix.String(),
			fMgr.absolutePathFileName)

		errs = append(errs, err)
	}

	targetFilePathDoesExist,
		err = fHelperAtom.doesFileMgrPathFileExist(
		targetFMgr,
		PreProcPathCode.None(),
		ePrefix,
		"#2 targetFMgr.absolutePathFileName")

	if err != nil {
		errs = append(errs, err)
		return strMech.ConsolidateErrors(errs)
	}

	if !targetFilePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: The move operation failed!\n"+
			"The target file DOES NOT EXIST!\n"+
			"targetFMgr='%v'\n",
			ePrefix.String(),
			targetFMgr.absolutePathFileName)

		errs = append(errs, err)
	}

	return strMech.ConsolidateErrors(errs)
}

// openFile - Helper method used to open files specified
// by FileMgr.
//
// If the directory does not previously exist, it will be
// created provided that input parameter, 'createTheDirectory',
// is set equal to 'true'.
//
// If the file does not previously exist, it will be created
// provided that input parameter, 'createTheFile', is set
// equal to 'true'.
//
// If the file does previously exist, it will be closed before
// being re-opened using the input parameter, 'openFileAccessCtrl',
// as a file open access specification.
func (fMgrHlpr *fileMgrHelper) openFile(
	fMgr *FileMgr,
	openFileAccessCtrl FileAccessControl,
	createTheDirectory bool,
	createTheFile bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper.openFile()",
		"")

	if err != nil {
		return err
	}

	var filePathDoesExist, directoryPathDoesExist bool

	if fMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	err = openFileAccessCtrl.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Input Parameter 'openFileAccessCtrl' is INVALID!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).doesFileMgrPathFileExist(
		fMgr,
		PreProcPathCode.None(),
		ePrefix.XCpy("filePathDoesExist<-fMgr"),
		"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	directoryPathDoesExist, err =
		fMgr.dMgr.DoesThisDirectoryExist(
			ePrefix.XCpy(
				"fMgr.dMgr"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Non-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
			"fMgr.dMgr='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath,
			err.Error())
	}

	fMgrHelperElectron := new(fileMgrHelperElectron)

	if !directoryPathDoesExist && createTheDirectory {

		err = fMgr.dMgr.MakeDir(
			ePrefix.XCpy(
				"fMgr.dMgr"))

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Error returned by fMgr.dMgr.MakeDir().\n"+
				"fMgr.dMgr='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				fMgr.dMgr.absolutePath,
				err.Error())
		}

	} else if !directoryPathDoesExist && !createTheDirectory {

		return fmt.Errorf("%v\n"+
			"Error: File Manager Directory Path DOES NOT EXIST!\n"+
			"createTheDirectory == false\n"+
			"fMgr.dMgr='%v'\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath)
	}

	if !filePathDoesExist && !createTheFile {

		return fmt.Errorf("%v\n"+
			"Error: The File Manager File (fMgr) DOES NOT EXIST!\n"+
			"createTheFile == false\n"+
			"File (fMgr)='%v'\n",
			ePrefix.String(),
			fMgr.absolutePathFileName)

	} else if !filePathDoesExist && createTheFile {

		var createTruncateAccessCtrl FileAccessControl

		createTruncateAccessCtrl,
			err =
			new(FileAccessControl).NewReadWriteCreateTruncateAccess(
				ePrefix.XCpy("createTruncateAccessCtrl<-"))

		if err != nil {
			return err
		}

		err = fMgrHelperElectron.
			lowLevelOpenFile(fMgr,
				createTruncateAccessCtrl,
				ePrefix.XCpy("fMgr<-createTruncateAccessCtrl"))

		if err != nil {
			return err
		}

	} // if !filePathDoesExist && !createTheFile

	// At this point, the file exists on disk. Close it!

	err = fMgrHelperElectron.lowLevelCloseFile(
		fMgr,
		"fMgr",
		ePrefix.XCpy("fMgr"))

	if err != nil {
		return err
	}

	// Now, open the file with the specified access configuration
	return fMgrHelperElectron.lowLevelOpenFile(
		fMgr,
		openFileAccessCtrl,
		ePrefix)
}

// readFileSetup
//
// Helper method designed to provide standard setup for
// methods reading data from a file identified by
// FileMgr.
func (fMgrHlpr *fileMgrHelper) readFileSetup(
	fMgr *FileMgr,
	readAccessCtrl FileAccessControl,
	createTheDirectory bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var err, err2 error
	var filePathDoesExist, dirPathDoesExist bool

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper.readFileSetup()",
		"")

	if err != nil {
		return err
	}

	if fMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())
	}

	err = readAccessCtrl.IsValidInstanceError(
		ePrefix.XCpy("readAccessCtrl"))

	if err != nil {
		return err
	}

	var fNewOpenType FileOpenType

	fNewOpenType,
		err = readAccessCtrl.GetFileOpenType(
		ePrefix.XCpy(
			"readAccessCtrl"))

	if err != nil {
		return err
	}

	if fNewOpenType != FOpenType.TypeReadWrite() &&
		fNewOpenType != FOpenType.TypeReadOnly() {

		return fmt.Errorf("%v\n"+
			"Error: Input Parameter readAccessCtrl is NOT an "+
			"open file 'READ' Type.\n"+
			"readAccessCtrl File Open Type='%v'\n",
			ePrefix.String(),
			fNewOpenType.String())
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix.XCpy("filePathDoesExist<-"),
			"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	dirPathDoesExist, err =
		fMgr.dMgr.DoesThisDirectoryExist(
			ePrefix.XCpy(
				"fMgr.dMgr"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Non-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
			"fMgr.dMgr='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath,
			err.Error())
	}

	if !dirPathDoesExist && createTheDirectory {

		err = fMgr.dMgr.MakeDir(
			ePrefix.XCpy(
				"fMgr.dMgr"))

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Error returned by fMgr.dMgr.MakeDir().\n"+
				"fMgr.dMgr='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				fMgr.dMgr.absolutePath,
				err.Error())
		}

	} else if !dirPathDoesExist && !createTheDirectory {

		return fmt.Errorf("%v\n"+
			"Error: File Manager (fMgr) Directory Path DOES NOT EXIST!\n"+
			"Directory Path (fMgr.dMgr)='%v'\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath)
	}

	fMgrHelperElectron := new(fileMgrHelperElectron)

	if !filePathDoesExist {

		err = fMgrHelperElectron.
			lowLevelCloseFile(fMgr, "fMgr", ePrefix)

		if err != nil {
			return err
		}

		var createTruncateAccessCtrl FileAccessControl

		createTruncateAccessCtrl,
			err = new(FileAccessControl).
			NewReadWriteCreateTruncateAccess(ePrefix.XCpy(
				"createTruncateAccessCtrl<-"))

		if err != nil {
			return err
		}

		err = fMgrHelperElectron.lowLevelOpenFile(fMgr, createTruncateAccessCtrl, ePrefix)

		if err != nil {
			return err
		}

		err = fMgrHelperElectron.lowLevelCloseFile(fMgr, "fMgr", ePrefix)

		if err != nil {
			return err
		}
	}

	invalidAccessType := true

	var fOpenType FileOpenType

	fOpenType,
		err2 = fMgr.fileAccessStatus.GetFileOpenType(ePrefix.XCpy(
		"fOpenType<-fMgr.fileAccessStatus"))

	if err2 == nil {
		if fOpenType == FOpenType.TypeReadWrite() ||
			fOpenType == FOpenType.TypeReadOnly() {

			invalidAccessType = false
		}
	}

	if !fMgr.isFilePtrOpen ||
		fMgr.filePtr == nil ||
		invalidAccessType ||
		err2 != nil {

		err = fMgrHelperElectron.lowLevelCloseFile(fMgr, "fMgr", ePrefix)

		if err != nil {
			return err
		}

		// The file exists, just open it for read/write
		// access.
		err = fMgrHelperElectron.lowLevelOpenFile(
			fMgr, readAccessCtrl, ePrefix)

		if err != nil {
			return err
		}
	}

	if fMgr.fileBufRdr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: fMgr.fileBufRdr == nil\n",
			ePrefix.String())

		return err
	}

	return nil
}

// setFileMgrPathFileName
//
// This method will configure an instance of FileMgr
// using the path, file name and file extension
// supplied by input parameter 'pathFileNameExt'.
//
// If 'pathFileNameExt' is submitted as an empty
// string or a string consisting of all white spaces,
// the returned boolean parameter 'isEmpty' will be set
// to 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. This method
//		will configure this FileMgr instance using the
//		path, file name and file extension supplied by
//		input parameter 'pathFileNameExt'.
//
//	pathFileNameExt				string
//
//		This string holds the path, file name and file
//		extension used to configure the instance of
//		FileMgr passed as input paramter 'fMgr'.
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
//	isEmpty						bool
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
func (fMgrHlpr *fileMgrHelper) setFileMgrPathFileName(
	fMgr *FileMgr,
	pathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelper."+
			"setFileMgrPathFileName()",
		"")

	if err != nil {
		return isEmpty, err
	}

	if fMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return isEmpty, err
	}

	errCode := 0

	errCode,
		_,
		pathFileNameExt =
		new(fileHelperElectron).
			isStringEmptyOrBlank(pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' is a zero length or empty string!\n",
			ePrefix.String())

		return isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' consists entirely of blank spaces!\n",
			ePrefix.String())

		return isEmpty, err
	}

	fh := new(FileHelper)

	adjustedPathFileNameExt := fh.AdjustPathSlash(pathFileNameExt)

	adjustedFileNameExt,
		isEmptyFileName,
		err2 :=
		fh.CleanFileNameExtStr(
			adjustedPathFileNameExt,
			ePrefix.XCpy(
				"adjustedPathFileNameExt"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt).\n"+
			"adjustedPathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedPathFileNameExt,
			err2.Error())

		return isEmpty, err
	}

	if isEmptyFileName {

		err = fmt.Errorf("%v\n"+
			"Error: File Name returned from fh.CleanFileNameExtStr(adjustedPathFileNameExt)\n"+
			"is a Zero Length String!.\n"+
			"pathFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedPathFileNameExt)

		return isEmpty, err
	}

	remainingPathStr :=
		strings.TrimSuffix(
			adjustedPathFileNameExt,
			adjustedFileNameExt)

	var dMgr DirMgr

	errCode, _, remainingPathStr =
		new(fileHelperElectron).
			isStringEmptyOrBlank(
				remainingPathStr)

	if errCode < 0 {

		dMgr = DirMgr{}

	} else {

		dMgr, err2 = new(DirMgr).New(
			remainingPathStr,
			ePrefix.XCpy(
				"remainingPathStr"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned from DirMgr{}.NewFromPathFileNameExtStr("+
				"remainingPathStr).\n"+
				"remainingPathStr='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				remainingPathStr,
				err2.Error())

			return isEmpty, err
		}
	}

	isEmpty,
		err =
		new(fileMgrHelperAtom).
			setFileMgrDirMgrFileName(
				fMgr,
				&dMgr,
				adjustedFileNameExt,
				ePrefix)

	return isEmpty, err
}

// writeFileSetup - Helper method designed to provide
// standard setup for methods writing data to the file
// identified by FileMgr.
func (fMgrHlpr *fileMgrHelper) writeFileSetup(
	fMgr *FileMgr,
	writeAccessCtrl FileAccessControl,
	createTheDirectory bool,
	errorPrefix interface{}) error {

	if fMgrHlpr.lock == nil {
		fMgrHlpr.lock = new(sync.Mutex)
	}

	fMgrHlpr.lock.Lock()

	defer fMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {
		return err
	}

	var filePathDoesExist, dirPathDoesExist bool

	if fMgr == nil {
		return errors.New("%v\n" +
			"Error: Input parameter 'fMgr' is a nil pointer!\n")
	}

	err = writeAccessCtrl.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by writeAccessCtrl.IsValidInstanceError()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	var fNewOpenType FileOpenType

	fNewOpenType,
		err = writeAccessCtrl.GetFileOpenType(
		ePrefix.XCpy("fNewOpenType<-writeAccessCtrl"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by writeAccessCtrl.GetFileOpenType()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	if fNewOpenType != FOpenType.TypeReadWrite() &&
		fNewOpenType != FOpenType.TypeWriteOnly() {

		return fmt.Errorf("%v\n"+
			"Error: Input Parameter writeAccessCtrl is NOT an "+
			"open file 'WRITE' Type.\n"+
			"readAccessCtrl File Open Type='%v'\n",
			ePrefix.String(),
			fNewOpenType.String())
	}

	filePathDoesExist,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	if err != nil {
		return err
	}

	dirPathDoesExist, err =
		fMgr.dMgr.DoesThisDirectoryExist(
			ePrefix.XCpy(
				"fMgr.dMgr"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Non-Path error returned by fMgr.dMgr.DoesThisDirectoryExist()\n"+
			"fMgr.dMgr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath,
			err.Error())
	}

	if !dirPathDoesExist && createTheDirectory {

		err = fMgr.dMgr.MakeDir(
			ePrefix.XCpy(
				"fMgr.dMgr"))

		if err != nil {
			return fmt.Errorf("%v\n"+
				"Error returned by fMgr.dMgr.MakeDir().\n"+
				"fMgr.dMgr='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				fMgr.dMgr.absolutePath,
				err.Error())
		}

	} else if !dirPathDoesExist && !createTheDirectory {

		return fmt.Errorf("%v\n"+
			"Error: File Manager (fMgr) Directory Path DOES NOT EXIST!\n"+
			"Directory Path (fMgr.dMgr)='%v'\n",
			ePrefix.String(),
			fMgr.dMgr.absolutePath)
	}

	fMgrHelperElectron := new(fileMgrHelperElectron)

	if !filePathDoesExist {

		err = fMgrHelperElectron.lowLevelCloseFile(
			fMgr,
			"fMgr",
			ePrefix.XCpy("fMgr"))

		if err != nil {
			return err
		}

		var createTruncateAccessCtrl FileAccessControl

		createTruncateAccessCtrl,
			err =
			new(FileAccessControl).
				NewReadWriteCreateTruncateAccess(ePrefix.XCpy(
					"createTruncateAccessCtrl<-"))

		if err != nil {
			return err
		}

		err = fMgrHelperElectron.lowLevelOpenFile(
			fMgr, createTruncateAccessCtrl, ePrefix)

		if err != nil {
			return err
		}

		err = fMgrHelperElectron.
			lowLevelCloseFile(
				fMgr,
				"fMgr",
				ePrefix)

		if err != nil {
			return err
		}
	}

	invalidAccessType := true

	var err2 error

	var fOpenType FileOpenType

	fOpenType,
		err2 = fMgr.fileAccessStatus.GetFileOpenType(
		ePrefix.XCpy("fOpenType<-fMgr.fileAccessStatus"))

	if err2 == nil {
		if fOpenType == fOpenType.TypeReadWrite() ||
			fOpenType == fOpenType.TypeWriteOnly() {

			invalidAccessType = false
		}
	}

	if !fMgr.isFilePtrOpen ||
		fMgr.filePtr == nil ||
		invalidAccessType ||
		err2 != nil {

		err = fMgrHelperElectron.
			lowLevelCloseFile(
				fMgr,
				"fMgr",
				ePrefix)

		if err != nil {
			return err
		}

		// The file exists, just open it for read/write
		// access.
		err = fMgrHelperElectron.lowLevelOpenFile(
			fMgr,
			writeAccessCtrl,
			ePrefix)

		if err != nil {
			return err
		}
	}

	if fMgr.fileBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: fMgr.fileBufWriter == nil\n",
			ePrefix.String())

		return err
	}

	return nil
}
