package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"testing"
)

func TestFileMgr_Write_000100(t *testing.T) {

	funcName := "TestFileMgr_Write_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fMgrReadFile FileMgr

	fMgrReadFile,
		err = new(FileMgr).New(
		targetReadFile,
		ePrefix.XCpy("fMgrReadFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var readFileSize int64

	readFileSize = fMgrReadFile.GetFileSize()

	err =
		fMgrReadFile.OpenThisFileComponents(
			ePrefix.XCpy("fMgrReadFile"),
			false,
			"-r--r--r--",
			FOpenType.TypeReadOnly())

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var readBuf = make([]byte, 128)
	var numOfBytesRead, totalNumOfBytesRead int
	var strArray StringArrayDto

	for i := 0; i < 256; i++ {

		if i == 255 {

			t.Errorf("%v\n"+
				"Error: Read cycles have exceeded Maximum Limit!\n"+
				"Cycle Index = '%v'\n",
				ePrefix.String(),
				i)

			return
		}

		numOfBytesRead,
			err = fMgrReadFile.Read(readBuf)

		if err != nil &&
			err != io.EOF {

			_ = fMgrReadFile.CloseThisFile(
				ePrefix.XCpy("fMgrReadFile"))

			t.Errorf("%v\n"+
				"Error returned by fMgrReadFile.Read(readBuf)\n"+
				"Read Cycle Index = '%v'\n"+
				"Read File = '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				i,
				fMgrReadFile.GetAbsolutePathFileName(),
				err.Error())

			return
		}

		if numOfBytesRead > 0 {

			strArray.PushBytes(readBuf[0:numOfBytesRead])

		}

		totalNumOfBytesRead += numOfBytesRead

		if err == io.EOF {

			break
		}

	}

	err = fMgrReadFile.CloseThisFile(
		ePrefix.XCpy("fMgrReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileMgr_Write_000100.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fMgrWriteFile FileMgr

	fMgrWriteFile,
		err = new(FileMgr).New(
		targetWriteFile,
		ePrefix.XCpy("fMgrWriteFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err =
		fMgrWriteFile.OpenThisFileComponents(
			ePrefix.XCpy("targetWriteFile"),
			true,
			"--w--w--w-",
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeTruncate())

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenStrArray := strArray.GetStringArrayLength()

	var numOfBytesWritten, totalNumOfBytesWritten int

	for j := 0; j < lenStrArray; j++ {

		numOfBytesWritten,
			err = fMgrWriteFile.Write(
			[]byte(strArray.StrArray[j]))

		if err != nil {

			_ = fMgrWriteFile.CloseThisFile(
				nil)

			t.Errorf("%v\n"+
				"Error returned by fMgrWriteFile.Write(readBuf)\n"+
				"Read Cycle Index = '%v'\n"+
				"Write File = '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				j,
				fMgrWriteFile.GetAbsolutePathFileName(),
				err.Error())

			return
		}

		totalNumOfBytesWritten += numOfBytesWritten

	}

	err = fMgrWriteFile.CloseThisFile(
		ePrefix.XCpy("fMgrWriteFile Final Close"))

	if err != nil {

		t.Errorf("%v\n"+
			"Errors occurred while flushing the buffer\n"+
			"and closing the Write File!\n"+
			"Write File = '%v'\n"+
			"Errors = \n%v\n",
			ePrefix.String(),
			fMgrWriteFile.GetAbsolutePathFileName(),
			err.Error())

		return
	}

	if totalNumOfBytesRead != totalNumOfBytesWritten {

		t.Errorf("%v\n"+
			"Error:  totalNumOfBytesRead != totalNumOfBytesWritten\n"+
			"        totalNumOfBytesRead = '%v'\n"+
			"     totalNumOfBytesWritten = '%v'\n"+
			"Total Bytes in String Array = '%v'\n",
			ePrefix.String(),
			totalNumOfBytesRead,
			totalNumOfBytesWritten,
			strArray.GetTotalBytesInStrings())

		return
	}

	var writeFileSize int64

	writeFileSize = fMgrWriteFile.GetFileSize()

	if readFileSize != writeFileSize {

		t.Errorf("%v\n"+
			"Error: readFileSize != writeFileSize\n"+
			" readFileSize = '%v'\n"+
			"writeFileSize = '%v'\n",
			ePrefix.String(),
			readFileSize,
			writeFileSize)

		return
	}

	err = fMgrWriteFile.
		DeleteThisFile(
			ePrefix.XCpy("fMgrWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if fMgrWriteFile.DoesFileExist() == true {

		t.Errorf("%v\n"+
			"Error: After fMgrWriteFile was deleted\n"+
			"it was found to still exist on disk.\n"+
			"fMgrWriteFile = '%v'\n",
			ePrefix.String(),
			fMgrWriteFile.GetAbsolutePathFileNameLc())

		return
	}

	return
}
