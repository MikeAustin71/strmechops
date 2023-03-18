package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
	"time"
)

type dirMgrHelperMolecule struct {
	lock *sync.Mutex
}

// deleteAllSubDirectories
//
// The directory identified by the input parameter 'dMgr'
// is treated as the parent directory.
//
// This method will proceed to delete all directories and
// files which are subsidiary to the parent directory,
// or top level directory, identified by 'dMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All subdirectories and files which are subordinate to
// the parent or top level directory identified by 'dMgr'
// will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. All
//		subdirectories and files subsidiary to the parent
//		or top level directory identified by 'dMgr' will
//		be deleted.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
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
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
func (dMgrHlprMolecule *dirMgrHelperMolecule) deleteAllSubDirectories(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	errs []error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	funcName := "dirMgrHelper.doesDirectoryExist() "

	errs = make([]error, 0)

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"

	}

	dirPathDoesExist,
		_,
		err := new(dirMgrHelperAtom).doesDirectoryExist(
		dMgr,
		PreProcPathCode.None(),
		dMgrLabel,
		ePrefix)

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if !dirPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: %v Directory Path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath)

		errs = append(errs, err)

		return errs
	}

	var err2, err3 error

	dirMgrPtr, err := os.Open(dMgr.absolutePath)

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error return by os.Open(dMgr.absolutePath)\n"+
			"dMgr.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err.Error())

		errs = append(errs, err2)

		return errs
	}

	var nameFileInfos []os.FileInfo

	err3 = nil

	osPathSeparatorStr := string(os.PathSeparator)

	for err3 != io.EOF {

		nameFileInfos, err3 = dirMgrPtr.Readdir(10000)

		if err3 != nil && err3 != io.EOF {

			_ = dirMgrPtr.Close()

			err2 = fmt.Errorf("%v\n"+
				"Error returned by dirMgrPtr.Readdirnames(10000).\n"+
				"dMgr.absolutePath='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dMgr.absolutePath,
				err3.Error())

			errs = append(errs, err2)

			return errs
		}

		for _, nameFInfo := range nameFileInfos {

			if nameFInfo.IsDir() {

				err = os.RemoveAll(dMgr.absolutePath + osPathSeparatorStr + nameFInfo.Name())

				if err != nil {

					err2 = fmt.Errorf("%v\n"+
						"Error returned by os.RemoveAll(subDir)\n"+
						"subDir='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						dMgr.absolutePath+osPathSeparatorStr+nameFInfo.Name(),
						err.Error())

					errs = append(errs, err2)

					continue
				}
			}
		}
	}

	if dirMgrPtr != nil {

		err = dirMgrPtr.Close()

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by %vPtr.Close().\n"+
				"%v='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgrLabel,
				dMgr.absolutePath,
				err.Error())

			errs = append(errs, err2)
		}
	}

	return errs
}

// lowLevelCopyFile
//
// This low level helper method is designed
// to copy files from a source file to a destination file.
//
// No validation or error checking is performed on the input
// parameters.
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelCopyFile(
	srcFile string,
	srcFInfo os.FileInfo,
	dstFile string,
	srcLabel string,
	dstLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"dirMgrHelperMolecule."+
			"lowLevelCopyFile()",
		"")

	if err != nil {
		return err
	}

	if len(srcFile) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			srcLabel)
	}

	if len(srcLabel) == 0 {
		srcLabel = "srcFile"
	}

	if len(dstFile) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			dstLabel)
	}

	if len(dstLabel) == 0 {
		dstLabel = "dstFile"
	}

	if !srcFInfo.Mode().IsRegular() {

		return fmt.Errorf("%v\n"+
			"Error: %v is a Non-Regular File and cannot be copied!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcLabel,
			srcLabel,
			srcFile)
	}

	// First, open the source file
	inSrcPtr, err := os.Open(srcFile)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned from os.Open(srcFile)\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcLabel,
			srcFile,
			err.Error())
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outDestPtr, err := os.Create(dstFile)

	if err != nil {

		_ = inSrcPtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned from os.Create(destinationFile)\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstFile,
			err.Error())
	}

	bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

	if err2 != nil {

		_ = inSrcPtr.Close()
		_ = outDestPtr.Close()

		err = fmt.Errorf("%v\n"+
			"Error returned from io.Copy(%v, %v)\n"+
			"%v='%v'\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			srcLabel,
			dstLabel,
			dstFile,
			srcLabel,
			srcFile,
			err2.Error())

		return err
	}

	errs := make([]error, 0)

	// flush file buffers inSrcPtr memory
	err = outDestPtr.Sync()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Sync()\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstFile,
			err.Error())

		errs = append(errs, err2)
	}

	err = inSrcPtr.Close()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from inSrcPtr.Close()\n"+
			"inSrcPtr=source='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFile,
			err.Error())

		errs = append(errs, err2)
	}

	inSrcPtr = nil

	err = outDestPtr.Close()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Close()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstFile,
			err.Error())

		errs = append(errs, err2)
	}

	outDestPtr = nil

	if len(errs) > 0 {

		return new(StrMech).ConsolidateErrors(errs)
	}

	var dstFileDoesExist bool
	var dstFileInfo FileInfoPlus

	dstFileDoesExist,
		dstFileInfo,
		err = new(dirMgrHelperElectron).
		lowLevelDoesDirectoryExist(
			dstFile,
			dstLabel,
			ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: After Copy IO operation, %v\n"+
			"generated non-path error!\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstLabel,
			dstFile,
			err.Error())
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = '%v' = '%v'\n",
			ePrefix.String(),
			dstLabel,
			dstFile)

		return err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes in source file!\n"+
			"Source File Bytes='%v'   Bytes Coped='%v'\n"+
			"Source File=%v='%v'\n"+
			"Destination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			srcLabel,
			srcFile,
			dstLabel,
			dstFile)

		return err
	}

	err = nil

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes\n"+
			"in the destination file!\n"+
			"Source File Bytes='%v'   Destination File Bytes='%v'\n"+
			"Source File=%v='%v'\n"+
			"Destination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			dstFileInfo.Size(),
			srcLabel,
			srcFile,
			dstLabel,
			dstFile)
	}

	return err
}

// lowLevelDeleteDirectoryAll
//
// Helper method designed for use by DirMgr.
//
// This method will delete the designated directory and
// constituent file specified by input parameter 'dMgr',
// as well as all subsidiary directories and files. This
// means that the entire directory tree designated by
// 'dMgr', along with all the contained files, will be
// deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This low-level method will not perform validation
//		services. It assumes that 'dMgr' specifies a
//		directory path which actually exists on disk.
//
//	(2) This method will delete the directory and
//		constituent files identified by input parameter
//		'dMgr'. In addition, all the child directories
//		and files subordinate to the directory designated
//		by 'dMgr' will likewise be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. The entire
//		directory tree identified by this parameter will
//		be deleted along with all the resident files.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelDeleteDirectoryAll(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	funcName := "dirMgrHelper.lowLevelDeleteDirectoryAll() "

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	var err2 error

	if dMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	for i := 0; i < 3; i++ {

		err2 = os.RemoveAll(dMgr.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.RemoveAll(%v.absolutePath)\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgrLabel,
				dMgr.absolutePath,
				err2.Error())

		} else {
			// err2 == nil
			// Deletion was successful
			dMgr.doesAbsolutePathExist = false
			dMgr.doesPathExist = false
			dMgr.actualDirFileInfo = FileInfoPlus{}
			return nil
		}

		time.Sleep(50 * time.Millisecond)
	}

	return err
}

// lowLevelMakeDir
//
// Helper Method used by 'DirMgr'. This method will
// create the directory path including parent directories
// for the path specified by 'dMgr'.
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelMakeDir(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirCreated bool,
	err error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	dirCreated = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"dirMgrHelperMolecule."+
			"lowLevelMakeDir()",
		"")

	if err != nil {
		return dirCreated, err
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is a 'nil' pointer!\n",
			ePrefix)

		return dirCreated, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	dMgrHlprAtom := dirMgrHelperAtom{}

	dMgrPathDoesExist,
		_,
		err :=
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix.XCpy(
				"dMgr"))

	if err != nil {
		return dirCreated, err
	}

	if dMgrPathDoesExist {
		// The directory exists
		// Nothing to do.
		return dirCreated, err
	}

	var fPermCfg FilePermissionConfig

	fPermCfg, err =
		new(FilePermissionConfig).New(
			"drwxrwxrwx",
			ePrefix)

	if err != nil {

		return dirCreated, err
	}

	var modePerm os.FileMode

	modePerm,
		err = fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy(
			"modePerm<-fPermCfg"))

	if err != nil {

		return dirCreated, err
	}

	var err2 error

	err2 = os.MkdirAll(dMgr.absolutePath, modePerm)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.MkdirAll(dMgr.absolutePath, modePerm).\n"+
			"dMgr.absolutePath='%v'\n"+
			"modePerm=\"drwxrwxrwx\"\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err2.Error())

		return dirCreated, err
	}

	dMgrPathDoesExist,
		_,
		err2 =
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix.XCpy(
				"dMgr"))

	if err2 != nil {
		err = fmt.Errorf("Error: After attempted directory creation, "+
			"a non-path error was generated!\n"+
			"%v.absolutePath='%v'\n"+
			"Error='%v'\n",
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())
		return dirCreated, err
	}

	if !dMgrPathDoesExist {
		err = fmt.Errorf("Error: After attempted directory creation,\n"+
			"the directory DOES NOT EXIST!\n"+
			"%v=%v\n", dMgrLabel, dMgr.absolutePath)

		return dirCreated, err
	}

	dirCreated = true
	err = nil

	return dirCreated, err
}
