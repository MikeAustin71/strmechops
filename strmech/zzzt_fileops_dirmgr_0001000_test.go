package strmech

import (
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
