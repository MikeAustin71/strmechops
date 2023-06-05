package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

// fileMgrHelperAtom
//
// Provides helper methods for Type
// fileMgrHelper.
type fileMgrHelperAtom struct {
	lock *sync.Mutex
}

// copyInFileMgr
//
// This method receives two instances of FileMgr,
// 'destinationFMgr' and 'sourceFMgr'. With the sole
// exception of the os.File pointer, all internal
// member data values will be copied from 'sourceFMgr' to
// 'destinationFMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The internal File Pointer (filePtr *os.File) for the
// 'sourceFMgr' instance of FileMgr will not be copied to
// 'destinationFMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationFMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. With the
//		sole exception of the internal File Pointer
//		(filePtr *os.File), all internal member data
//		elements in 'sourceFMgr' will be copied to the
//		corresponding member data elements in
//		'destinationFMgr'.
//
//	sourceFMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. With the
//		sole exception of the internal File Pointer
//		(filePtr *os.File), all internal member data
//		elements in 'sourceFMgr' will be copied to the
//		corresponding member data elements in
//		'destinationFMgr'.
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
func (fMgrHlprAtom *fileMgrHelperAtom) copyInFileMgr(
	destinationFMgr *FileMgr,
	sourceFMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperAtom.copyInFileMgr()",
		"")

	if err != nil {
		return err
	}

	if destinationFMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationFMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceFMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationFMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	destinationFMgr.isInitialized = sourceFMgr.isInitialized

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&destinationFMgr.dMgr,
		&sourceFMgr.dMgr,
		ePrefix.XCpy(
			"destinationFMgr.dMg<-"+
				"sourceFMgr.dMgr"))

	if err != nil {
		return err
	}

	destinationFMgr.originalPathFileName =
		sourceFMgr.originalPathFileName

	destinationFMgr.absolutePathFileName =
		sourceFMgr.absolutePathFileName

	destinationFMgr.isAbsolutePathFileNamePopulated =
		sourceFMgr.isAbsolutePathFileNamePopulated

	destinationFMgr.doesAbsolutePathFileNameExist =
		sourceFMgr.doesAbsolutePathFileNameExist

	destinationFMgr.fileName = sourceFMgr.fileName

	destinationFMgr.isFileNamePopulated =
		sourceFMgr.isFileNamePopulated

	destinationFMgr.fileExt = sourceFMgr.fileExt

	destinationFMgr.isFileExtPopulated =
		sourceFMgr.isFileExtPopulated

	destinationFMgr.fileNameExt = sourceFMgr.fileNameExt

	destinationFMgr.isFileNameExtPopulated =
		sourceFMgr.isFileNameExtPopulated

	destinationFMgr.filePtr = nil

	destinationFMgr.isFilePtrOpen = false

	destinationFMgr.fileAccessStatus =
		sourceFMgr.fileAccessStatus.CopyOut()

	destinationFMgr.actualFileInfo =
		sourceFMgr.actualFileInfo.CopyOut()

	destinationFMgr.fileBytesWritten = 0

	destinationFMgr.buffBytesWritten = 0

	destinationFMgr.fileRdrBufSize = sourceFMgr.fileRdrBufSize

	destinationFMgr.fileWriterBufSize = sourceFMgr.fileWriterBufSize

	return err
}

// copyOutFileMgr
//
// This method receives an instance of FileMgr and returns
// a deep copy of that instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The internal File Pointer (filePtr *os.File) for the
// fileMgr instance of FileMgr will not be copied.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. This method
//		will construct and return a deep copy of this
//		instance.
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
//	deepCopyOfFileMgr			FileMgr
//
//		If this method completes successfully, a deep
//		copy of FileMgr input parameter 'fileMgr' will
//		be returned through this parameter.
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
func (fMgrHlprAtom *fileMgrHelperAtom) copyOutFileMgr(
	fileMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyOfFileMgr FileMgr,
	err error) {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperAtom.copyOutFileMgr()",
		"")

	if err != nil {
		return deepCopyOfFileMgr, err
	}

	if fileMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyOfFileMgr, err
	}

	deepCopyOfFileMgr.isInitialized = fileMgr.isInitialized

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&deepCopyOfFileMgr.dMgr,
		&fileMgr.dMgr,
		ePrefix.XCpy(
			"deepCopyOfFileMgr.dMgr<-"+
				"fileMgr.dMgr"))

	if err != nil {
		return deepCopyOfFileMgr, err
	}

	deepCopyOfFileMgr.originalPathFileName =
		fileMgr.originalPathFileName

	deepCopyOfFileMgr.absolutePathFileName =
		fileMgr.absolutePathFileName

	deepCopyOfFileMgr.isAbsolutePathFileNamePopulated =
		fileMgr.isAbsolutePathFileNamePopulated

	deepCopyOfFileMgr.doesAbsolutePathFileNameExist =
		fileMgr.doesAbsolutePathFileNameExist

	deepCopyOfFileMgr.fileName = fileMgr.fileName

	deepCopyOfFileMgr.isFileNamePopulated =
		fileMgr.isFileNamePopulated

	deepCopyOfFileMgr.fileExt = fileMgr.fileExt

	deepCopyOfFileMgr.isFileExtPopulated =
		fileMgr.isFileExtPopulated

	deepCopyOfFileMgr.fileNameExt = fileMgr.fileNameExt

	deepCopyOfFileMgr.isFileNameExtPopulated =
		fileMgr.isFileNameExtPopulated

	deepCopyOfFileMgr.filePtr = nil

	deepCopyOfFileMgr.isFilePtrOpen = false

	deepCopyOfFileMgr.fileAccessStatus =
		fileMgr.fileAccessStatus.CopyOut()

	deepCopyOfFileMgr.actualFileInfo =
		fileMgr.actualFileInfo.CopyOut()

	deepCopyOfFileMgr.fileBytesWritten = 0
	deepCopyOfFileMgr.buffBytesWritten = 0
	deepCopyOfFileMgr.fileRdrBufSize = fileMgr.fileRdrBufSize
	deepCopyOfFileMgr.fileWriterBufSize = fileMgr.fileWriterBufSize

	return deepCopyOfFileMgr, err
}

// isFileMgrValid
//
// This method receives a pointer to an instance of
// FileMgr (fMgr) which will be analyzed to determine if
// all the member variables contain valid values.
//
// If input parameter 'fMgr' is determined to be invalid,
// this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr which will be
//		analyzed to determine if all internal member
//		variables contain valid values. If this instance
//		of FileMgr evaluates as invalid, this method will
//		return an error.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If any of the internal member data variables
//		contained in 'fMgr' are found to be invalid, this
//		method will return an error configured with an
//		appropriate message identifying the invalid
//		member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any internal member data values are found to be
//		invalid, the returned error Type will encapsulate
//		an appropriate error message. This returned error
//		message will incorporate the method chain and text
//		passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the
//		beginning of the error message.
func (fMgrHlprAtom *fileMgrHelperAtom) isFileMgrValid(
	fMgr *FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileMgrHelperAtom.isFileMgrValid()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
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

	if fMgr.absolutePathFileName == "" {

		fMgr.isAbsolutePathFileNamePopulated = false

		err = fmt.Errorf("%v\n"+
			" Error: FileMgr absolutePathFileName is EMPTY!\n",
			ePrefix.String())

		return err

	}

	var err2 error

	err2 = new(dirMgrHelperBoson).isDirMgrValid(
		&fMgr.dMgr,
		ePrefix.XCpy("fMgr.dMgr"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"FileMgr Directory Manager INVALID\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return err
	}

	_,
		err = new(fileMgrHelperAtom).
		doesFileMgrPathFileExist(
			fMgr,
			PreProcPathCode.None(),
			ePrefix,
			"fMgr.absolutePathFileName")

	_ = fMgr.dMgr.DoesPathExist()

	_ = fMgr.dMgr.DoesAbsolutePathExist()

	return err
}

// doesFileMgrPathFileExist
//
// Used by FileMgr type to test for the existence of a
// path and file name. In addition, this method performs
// validation on the 'FileMgr' instance.
func (fMgrHlprAtom *fileMgrHelperAtom) doesFileMgrPathFileExist(
	fileMgr *FileMgr,
	preProcessCode PreProcessPathCode,
	errPrefDto *ePref.ErrPrefixDto,
	filePathTitle string) (
	filePathDoesExist bool,
	nonPathError error) {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	filePathDoesExist = false

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileMgrHelperAtom.doesFileMgrPathFileExist()"

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if nonPathError != nil {
		return filePathDoesExist, nonPathError
	}

	if len(filePathTitle) == 0 {
		filePathTitle = "filePath"
	}

	if fileMgr == nil {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return filePathDoesExist, nonPathError
	}

	errCode := 0

	errCode,
		_, fileMgr.absolutePathFileName =
		new(fileHelperElectron).isStringEmptyOrBlank(fileMgr.absolutePathFileName)

	if errCode == -1 {

		fileMgr.isAbsolutePathFileNamePopulated = false

		nonPathError = fmt.Errorf("%v\n"+
			"Error: '%v' is an empty string!\n",
			ePrefix.String(),
			filePathTitle)

		return filePathDoesExist, nonPathError
	}

	if errCode == -2 {

		fileMgr.isAbsolutePathFileNamePopulated = false

		nonPathError = fmt.Errorf("%v\n"+
			"\nError: '%v' consists of blank spaces!\n",
			ePrefix.String(),
			filePathTitle)

		return filePathDoesExist, nonPathError
	}

	if !fileMgr.isInitialized {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: This data structure is NOT initialized.\n"+
			"fileMgr.isInitialized='false'\n",
			ePrefix.String())

		return filePathDoesExist, nonPathError
	}

	var err2, err3 error

	err2 = fileMgr.dMgr.IsValidInstanceError(ePrefix.String())

	if err2 != nil {
		nonPathError = fmt.Errorf("\nFileMgr Directory Manager INVALID!\n"+
			"\nError='%v'", err2.Error())
		return filePathDoesExist, nonPathError
	}

	if preProcessCode == PreProcPathCode.PathSeparator() {

		fileMgr.absolutePathFileName =
			new(FileHelper).
				AdjustPathSlash(fileMgr.absolutePathFileName)

	} else if preProcessCode == PreProcPathCode.AbsolutePath() {

		fileMgr.absolutePathFileName, err2 =
			new(FileHelper).MakeAbsolutePath(
				fileMgr.absolutePathFileName,
				ePrefix)

		if err2 != nil {

			nonPathError = fmt.Errorf("%v\n"+
				"FileHelper{}.MakeAbsolutePath() FAILED!\n"+
				"%v='%v'\n"+
				"%v\n",
				ePrefix.String(),
				filePathTitle,
				fileMgr.absolutePathFileName,
				err2.Error())

			return filePathDoesExist, nonPathError
		}
	}

	var info FileInfoPlus

	filePathDoesExist,
		info,
		nonPathError =
		new(fileMgrHelperElectron).lowLevelDoesFileExist(
			fileMgr.absolutePathFileName,
			fileMgr.dMgr.absolutePath,
			ePrefix,
			"fileMgr.absolutePathFileName",
			"fileMgr.dMgr.absolutePath")

	if nonPathError != nil {
		fileMgr.doesAbsolutePathFileNameExist = false
		fileMgr.actualFileInfo = FileInfoPlus{}
		filePathDoesExist = false
		fileMgr.fileAccessStatus.Empty()
		return filePathDoesExist, nonPathError
	}

	if !filePathDoesExist {
		fileMgr.doesAbsolutePathFileNameExist = false
		fileMgr.actualFileInfo = FileInfoPlus{}
		fileMgr.fileAccessStatus.Empty()
		filePathDoesExist = false
		nonPathError = nil
		_ = fileMgr.dMgr.DoesPathExist()
		_ = fileMgr.dMgr.DoesAbsolutePathExist()
		return filePathDoesExist, nonPathError
	}

	// The path really does exist!
	errs := make([]error, 0, 10)
	filePathDoesExist = true
	nonPathError = nil
	fileMgr.doesAbsolutePathFileNameExist = true

	fileMgr.actualFileInfo,
		err2 =
		new(FileInfoPlus).
			NewFromPathFileInfo(
				fileMgr.dMgr.absolutePath,
				&info,
				ePrefix.XCpy(
					"fileMgr.dMgr.absolutePath"))

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned by FileInfoPlus{}.NewFromPathFileInfo(fileMgr.dMgr.absolutePath, info)\n"+
			"fileMgr.dMgr.absolutePath='%v'\n"+
			"info.Name()='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fileMgr.dMgr.absolutePath,
			info.Name(),
			err2.Error())

		errs = append(errs, err3)

	} else {

		permCode, err2 := new(FilePermissionConfig).
			NewByFileMode(
				fileMgr.actualFileInfo.Mode(),
				ePrefix.XCpy(
					"permCode<-fileMgr"))

		if err2 != nil {
			err3 = fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.NewByFileMode(fileMgr.actualFileInfo.Mode())\n"+
				"Error= \n%v\n",
				funcName,
				err2.Error())

			errs = append(errs, err3)

		} else {

			err2 = fileMgr.
				fileAccessStatus.
				SetFilePermissionCodes(
					permCode,
					ePrefix.XCpy(
						"fileMgr.fileAccessStatus->"))

			if err2 != nil {
				err3 = fmt.Errorf("%v\n"+
					"Error returned by fileMgr.fileAccessStatus.SetFilePermissionCodes(permCode)\n"+
					"Error='%v'\n",
					funcName,
					err2.Error())

				errs = append(errs, err3)
			}

		}
	}

	_ = fileMgr.dMgr.DoesPathExist()

	_ = fileMgr.dMgr.DoesAbsolutePathExist()

	if len(errs) > 0 {
		nonPathError = new(StrMech).ConsolidateErrors(errs)
	}

	return filePathDoesExist, nonPathError
}

// setFileMgrDirMgrFileName
//
// This configures a FileMgr instance based on a
// Directory Manager instance plus a string containing
// the file name and file extension.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. This method
//		will configure this FileMgr instance using two
//		components. First, the directory component will
//		be taken from the DirMgr instance passed as input
//		parameter 'dMgr'. Second, the file name and file
//		extension will be taken from input parameter
//	 	'fileNameExt'.
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of Directory Manager
//		('DirMgr'). The directory name contained in this
//		DirMgr instance will be extracted and used to
//		construct the directory name configured in
//		'fMgr'.
//
//	fileNameExt					string
//
//		This string holds the file name and file
//		extension used to configure the FileMgr instance
//		passed as input parameter 'fMgr'.
//
//		If 'fileNameExt' is submitted as an empty string
//		or a string containing all white space, an error
//		will be returned.
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
//		If input parameter 'fileNameExt' is submitted as
//		an empty string or a string containing all white
//		space, this boolean parameter will return a value
//		of 'true'.
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
func (fMgrHlprAtom *fileMgrHelperAtom) setFileMgrDirMgrFileName(
	fMgr *FileMgr,
	dMgr *DirMgr,
	fileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperAtom."+
			"setFileMgrDirMgrFileName()",
		"")

	if err != nil {
		return false, err
	}

	isEmpty = true
	err = nil

	if fMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return isEmpty, err
	}

	var err2 error
	_,
		_,
		err2 = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			false, // Not required to exist
			"dMgr",
			ePrefix.XCpy(
				"dMgr"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is INVALID!\n"+
			"dMgr.absolutePath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err2.Error())

		return isEmpty, err
	}

	errCode,
		_,
		fileNameExt :=
		new(fileHelperElectron).
			isStringEmptyOrBlank(fileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileNameExt' is a Zero length string!\n",
			ePrefix.String())

		return isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileNameExt' consists entirely of blank spaces!\n",
			ePrefix.String())

		return isEmpty, err
	}

	fh := new(FileHelper)

	adjustedFileNameExt,
		isFileNameEmpty,
		err2 :=
		fh.CleanFileNameExtStr(
			fileNameExt,
			ePrefix.XCpy(
				"adjustedFileNameExt<-fileNameExt"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fh.CleanFileNameExtStr(fileNameExt).\n"+
			"fileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fileNameExt,
			err2.Error())

		return isEmpty, err
	}

	if isFileNameEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: fileName returned from fh.CleanFileNameExtStr(fileNameExt)\n"+
			"is a ZERO length string!\n"+
			"fileNameExt='%v'\n",
			ePrefix.String(),
			fileNameExt)

		return isEmpty, err
	}

	fMgrHelperBoson := new(fileMgrHelperBoson)

	err = fMgrHelperBoson.
		emptyFileMgr(fMgr, ePrefix)

	if err != nil {
		return isEmpty, err
	}

	err = new(dirMgrHelperBoson).
		copyDirMgrs(
			&fMgr.dMgr,
			dMgr,
			ePrefix.XCpy("fMgr.dMgr<-dMgr"))

	if err != nil {
		return isEmpty, err
	}

	s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(
		adjustedFileNameExt,
		ePrefix.XCpy(
			"s<-adjustedFileNameExt"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from fh.GetFileNameWithoutExt(adjustedFileNameExt).\n"+
			"adjustedFileNameExt='%v'\n"+
			"Error='%v'\n ",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		isEmpty = true

		_ = fMgrHelperBoson.
			emptyFileMgr(fMgr, ePrefix)

		return isEmpty, err
	}

	if fNameIsEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: fileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt)\n"+
			"is Zero length string!\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt)

		_ = fMgrHelperBoson.
			emptyFileMgr(fMgr, ePrefix)

		isEmpty = true

		return isEmpty, err
	}

	fMgr.isFileNamePopulated = true
	fMgr.fileName = s

	s,
		extIsEmpty,
		err2 := fh.GetFileExtension(
		adjustedFileNameExt,
		ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from fh.GetFileExt(fileNameAndExt).\n"+
			"fileNameAndExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		isEmpty = true

		_ = fMgrHelperBoson.emptyFileMgr(
			fMgr,
			ePrefix)

		return isEmpty, err
	}

	if !extIsEmpty {
		fMgr.isFileExtPopulated = true
		fMgr.fileExt = s
	}

	if fMgr.isFileNamePopulated {
		fMgr.isFileNameExtPopulated = true
		fMgr.fileNameExt = fMgr.fileName + fMgr.fileExt
	}

	lPath := len(fMgr.dMgr.absolutePath)
	if lPath == 0 {
		fMgr.absolutePathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.absolutePath[lPath-1] == os.PathSeparator {
		fMgr.absolutePathFileName = fMgr.dMgr.absolutePath + fMgr.fileNameExt

	} else {
		fMgr.absolutePathFileName =
			fMgr.dMgr.absolutePath + string(os.PathSeparator) + fMgr.fileNameExt

	}

	lPath = len(fMgr.dMgr.path)

	if lPath == 0 {
		fMgr.originalPathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.path[lPath-1] == os.PathSeparator {
		fMgr.originalPathFileName = fMgr.dMgr.path + fMgr.fileNameExt

	} else {
		fMgr.originalPathFileName = fMgr.dMgr.path + string(os.PathSeparator) + fMgr.fileNameExt
	}

	fMgr.isAbsolutePathFileNamePopulated = true

	_,
		filePathDoesExist,
		fInfoPlus,
		nonPathError :=
		new(fileHelperMolecule).doesPathFileExist(
			fMgr.absolutePathFileName,
			PreProcPathCode.None(), // Do NOT perform pre-processing on path
			ePrefix,
			"fMgr.absolutePathFileName")

	if filePathDoesExist && nonPathError == nil {
		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.actualFileInfo = fInfoPlus.CopyOut()

		err2 = fMgr.actualFileInfo.SetDirectoryPath(
			dMgr.absolutePath,
			ePrefix.XCpy(
				"dMgr.absolutePath"))

		if err2 != nil {
			isEmpty = true
			err = fmt.Errorf("%v\n"+
				"Error returned by fMgr.actualFileInfo.SetDirectoryPath(dMgr.absolutePath)\n"+
				"dMgr.absolutePath='%v'\n"+
				"%v",
				ePrefix.String(),
				dMgr.absolutePath,
				err2.Error())

			_ = fMgrHelperBoson.emptyFileMgr(
				fMgr,
				ePrefix)

			return isEmpty, err
		}

	} else {
		fMgr.doesAbsolutePathFileNameExist = false
		fMgr.actualFileInfo = FileInfoPlus{}
	}

	fMgr.isInitialized = true

	err = nil
	isEmpty = false

	return isEmpty, err
}
