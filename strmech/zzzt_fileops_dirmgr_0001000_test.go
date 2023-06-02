package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestDirMgr_ChangeWorkingDir_000100(t *testing.T) {

	funcName := "TestDirMgr_ChangeWorkingDir_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var err error
	var startDir, checkDir, targetDir string
	fh := FileHelper{}

	startDir,
		err = fh.GetAbsCurrDir(
		ePrefix.XCpy("startDir<-"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: GetAnsCurrDir() Failed\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	targetDir, err = fh.MakeAbsolutePath(
		FILEOpsTestLogDir,
		ePrefix.XCpy("FILEOpsTestLogDir"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: MakeAbsolutePath() Failed\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	err = fh.ChangeWorkingDir(
		targetDir,
		ePrefix.XCpy("targetDir"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: MakeAbsolutePath() Failed\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	checkDir, err = fh.GetAbsCurrDir(
		ePrefix.XCpy("checkDir<-"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: GetAbsCurrDir() #2 Failed\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	if checkDir != targetDir {

		t.Errorf("\n%v\n"+
			"checkDir != targetDir\n"+
			"Error Target Dir is NOT EQUAL to CheckDir (#1)!\n",
			ePrefix.String())
	}

	err = fh.ChangeWorkingDir(
		startDir,
		ePrefix.XCpy("startDir"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: Change To Start Dir Failed!\n"+
			"fh.ChangeWorkingDir(startDir)\n"+
			"%v\n",
			funcName,
			err.Error())
	}

	checkDir,
		err = fh.GetAbsCurrDir(ePrefix.XCpy(
		"checkDir<-"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: GetAbsCurrDir() 3 Failed!\n"+
			"%v\n",
			funcName,
			err.Error())

	}

	if checkDir != startDir {

		t.Errorf("\n%v\n"+
			"Start Dir != CheckDir\n"+
			"Error Target Dir is NOT EQUAL to CheckDir (#2)!\n",
			ePrefix.String())
	}
}
func TestDirMgr_CopyDirectory_000100(t *testing.T) {

	funcName := "TestDirMgr_CopyDirectory_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	targetDir := FILEOPSTrashDirectory + "/TestDirMgr_CopyFilesToDirectory_000100"
	//"../../checkFiles/TestDirMgr_CopyFilesToDirectory_01"

	fh := new(FileHelper)

	msgError,
		lowLevelErr := fh.DeleteDirPathAll(
		targetDir,
		ePrefix.XCpy("targetDir"))

	if msgError != nil {

		t.Errorf("\n%v\n"+
			"Test Setup Message Error returned by\n"+
			"fh.DeleteDirPathAll(targetDir).\n"+
			"targetDir='%v'\n"+
			"Message Error=\n%v\n",
			funcName,
			targetDir,
			msgError.Error())

		return
	}

	if lowLevelErr != nil {

		t.Errorf("\n%v\n"+
			"Test Setup Low Level Error returned by\n"+
			"fh.DeleteDirPathAll(targetDir).\n"+
			"targetDir='%v'\n"+
			"Low Level Error=\n%v\n",
			funcName,
			targetDir,
			lowLevelErr.Error())

		return
	}

	targetDMgr, err := new(DirMgr).New(
		targetDir,
		ePrefix.XCpy("targetDir"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Test Setup Error returned from DirMgr.New(targetDMgr).\n"+
			"targetDMgr='%v'\n"+
			"Error= \n%v\n",
			funcName,
			targetDMgr,
			err.Error())

		return
	}

	srcDir1 := FILEOPSFilesForTest + "/levelfilesfortest"

	//"../../filesfortest/levelfilesfortest"

	srcDMgr, err := new(DirMgr).New(
		srcDir1,
		ePrefix.XCpy("srcDir1"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Test Setup Error returned from new(DirMgr).New(srcDir1).\n"+
			"srcDir1='%v'\n"+
			"Error=\n%v\n",
			funcName,
			srcDir1,
			err.Error())

		return
	}

	fsc := FileSelectionCriteria{}

	var dirCopyStats DirectoryCopyStats
	var errs []error

	dirCopyStats,
		errs = srcDMgr.CopyDirectory(
		targetDMgr,
		fsc, false,
		ePrefix.XCpy("targetDMgr"))

	if len(errs) > 0 {

		t.Errorf("Error returned from srcDMgr.CopyDirectory(targetDMgr, fsc)\n"+
			"targetDir='%v'\nErrors Follow:\n\n%v",
			targetDMgr.GetAbsolutePath(),
			new(StrMech).ConsolidateErrors(errs))

		_,
			_ = fh.DeleteDirPathAll(
			targetDir,
			ePrefix)

		return

	}
	// 5 txt src Files
	/*
	   "../../filesfortest/levelfilesfortest/level_0_0_test.txt"
	   "../../filesfortest/levelfilesfortest/level_0_1_test.txt"
	   "../../filesfortest/levelfilesfortest/level_0_2_test.txt"
	   "../../filesfortest/levelfilesfortest/level_0_3_test.txt"
	   "../../filesfortest/levelfilesfortest/level_0_4_test.txt"
	*/

	fileNames := []string{"level_0_0_test.txt",
		"level_0_1_test.txt",
		"level_0_2_test.txt",
		"level_0_3_test.txt",
		"level_0_4_test.txt"}

	fsc = FileSelectionCriteria{}

	var fMgrCollection FileMgrCollection

	fMgrCollection,
		err = targetDMgr.FindFilesBySelectCriteria(
		fsc,
		ePrefix.XCpy(
			"targetDMgr"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Test Setup Error returned by targetDMgr.FindFilesBySelectCriteria(fsc).\n"+
			"targetDMgr='%v'\n"+
			"Error= \n%v\n",
			funcName,
			targetDMgr.GetAbsolutePath(),
			err.Error())

		_,
			_ = fh.DeleteDirPathAll(
			targetDir,
			nil)

		return
	}

	if fMgrCollection.GetNumOfFileMgrs() != 5 {

		t.Errorf("\n%v\n"+
			"Test Setup Error: Expected to find 5-files in 'targetDir'.\n"+
			"Instead, %v-files were found.\n",
			ePrefix.String(),
			fMgrCollection.GetNumOfFileMgrs())

		_,
			_ = fh.DeleteDirPathAll(targetDir, nil)

		return

	}

	if 5 != dirCopyStats.FilesCopied {

		t.Errorf("\n%v\n"+
			"Test Setup Error: Expected that dirCopyStats.FilesCopied='5'.\n"+
			"Instead, dirCopyStats.FilesCopied='%v'.\n",
			ePrefix.String(),
			dirCopyStats.FilesCopied)

		_,
			_ = fh.DeleteDirPathAll(targetDir, nil)

		return
	}

	var fMgr FileMgr

	for i := 0; i < fMgrCollection.GetNumOfFileMgrs(); i++ {

		fMgr, err = fMgrCollection.PeekFileMgrAtIndex(
			i,
			ePrefix.XCpy(
				fmt.Sprintf("fMgrCollection[%v]",
					i)))

		if err != nil {

			t.Errorf("\n%v\n"+
				"Error returned by fMgrCollection.GetFileMgrAtIndex(%v)\n"+
				"Error=\n%v\n",
				funcName,
				i,
				err.Error())

			_,
				_ = fh.DeleteDirPathAll(targetDir, nil)

			return
		}

		fileName := fMgr.GetFileNameExt()

		foundFile := false

		for k := 0; k < len(fileNames); k++ {

			if fileNames[k] == fileName {

				foundFile = true

			}
		}

		if foundFile == false {

			t.Errorf("\n%v\n"+
				"Error: File NOT Found. Expected to find specific file Name.\n"+
				"However, it WAS NOT FOUND!\n"+
				"FileName='%v'",
				ePrefix.String(),
				fileName)

			_,
				_ = fh.DeleteDirPathAll(targetDir, nil)

			return
		}

	}

	msgError,
		lowLevelErr = fh.DeleteDirPathAll(
		targetDir,
		ePrefix.XCpy("targetDir Cleanup"))

	if msgError != nil {

		t.Errorf("\n%v\n"+
			"Cleanup Message Error returned by\n"+
			"fh.DeleteDirPathAll(targetDir).\n"+
			"targetDir='%v'\n"+
			"Message Error=\n%v\n",
			funcName,
			targetDir,
			msgError.Error())

		return
	}

	if lowLevelErr != nil {

		t.Errorf("\n%v\n"+
			"Cleanup Low Level Error returned by\n"+
			"fh.DeleteDirPathAll(targetDir).\n"+
			"targetDir='%v'\n"+
			"Low Level Error=\n%v\n",
			funcName,
			targetDir,
			lowLevelErr.Error())

		return
	}

	return
}
