package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

type dirMgrHelperMolecule struct {
	lock *sync.Mutex
}

// lowLevelCopyFile
//
// This low level helper method is designed
// to copy files from a source file to a destination file.
//
// No validation or error checking is performed on the input
// parameters.
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelCopyFile(
	src string,
	srcFInfo os.FileInfo,
	dst string,
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

	if len(src) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			srcLabel)
	}

	if len(dst) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			dstLabel)
	}

	if !srcFInfo.Mode().IsRegular() {

		return fmt.Errorf("%v\n"+
			"Error: %v is a Non-Regular File and cannot be copied!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcLabel,
			srcLabel,
			src)
	}

	// First, open the source file
	inSrcPtr, err := os.Open(src)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned from os.Open(src)\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcLabel,
			src,
			err.Error())
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outDestPtr, err := os.Create(dst)

	if err != nil {

		_ = inSrcPtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned from os.Create(destinationFile)\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dst,
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
			dst,
			srcLabel,
			src,
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
			dst,
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
			src,
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
			dst,
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
			dst,
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
			dst,
			err.Error())
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = '%v' = '%v'\n",
			ePrefix.String(),
			dstLabel,
			dst)

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
			src,
			dstLabel,
			dst)

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
			src,
			dstLabel,
			dst)
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
	errPrefDto *ePref.ErrPrefixDto) (dirCreated bool, err error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"dirMgrHelperMolecule."+
			"lowLevelMakeDir()",
		"")

	if err != nil {
		return false, err
	}

	dirCreated = false
	err = nil

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
