package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"testing"
)

func TestFileIoWriter_ReadWrite_000100(t *testing.T) {

	funcName := "TestFileIoWriter_ReadWrite_000100()"

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+10)

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var trashDirectory string

	trashDirectory,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			FILEOpsBaseTrashDirectory,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = new(DirHelper).DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile = trashDirectory +
		string(os.PathSeparator) +
		"TestFileIoWriter_ReadWrite_000100.txt"

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var i64NumOfBytesRead, i64NumOfBytesWritten int64

	var readFileInfoPlus,
		writeFileInfoPlus FileInfoPlus

	var targetIoReader FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			4096,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var readStrBuilder = new(strings.Builder)

	i64NumOfBytesRead,
		err = targetIoReader.
		ReadAllToStrBuilder(
			readStrBuilder,
			true, // autoCloseOnExit
			ePrefix.XCpy("readStrBuilder<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetIoWriter FileIoWriter

	writeFileInfoPlus,
		targetIoWriter,
		err = new(FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,  // Default Buffer Size
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fHelper = new(FileHelper)

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	i64NumOfBytesWritten,
		err = targetIoWriter.
		WriteTextOrNumbers(
			readStrBuilder,
			"",
			"",
			true, // autoCloseOnExit
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if i64NumOfBytesRead != i64NumOfBytesWritten {

		t.Errorf(" %v\n"+
			"%v\n"+
			" Error: Expected Bytes Written != Actual Bytes Written\n"+
			" Expected Bytes Written = '%v'\n"+
			"   Actual Bytes Written = '%v'\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			i64NumOfBytesRead,
			i64NumOfBytesWritten,
			targetReadFile,
			targetWriteFile)

		return

	}

	writeFileInfoPlus,
		err = fHelper.
		GetFileInfoPlus(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if readFileInfoPlus.Size() != writeFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Target Read File Size != Target Write File Size!\n"+
			" Target Read File Size= %v\n"+
			"Target Write File Size= %v\n"+
			" Target Read File= %v\n"+
			"Target Write File= %v\n\n",
			ePrefix.String(),
			dashLineStr,
			readFileInfoPlus.Size(),
			writeFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	var filesAreEqual bool
	var reasonFilesNotEqual string

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.
		CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Target Files Comparison"))

	if err != nil {

		t.Errorf(" %v\n"+
			" Error Return from fHelper.CompareFiles()\n"+
			"  targetReadFile= %v\n"+
			" targetWriteFile= %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile)

		return
	}

	if !filesAreEqual {

		t.Errorf(" %v\n"+
			"%v\n"+
			" Error: Read and Write Files are NOT equal!\n"+
			" Reason: %v\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n\n",
			ePrefix.String(),
			dashLineStr,
			reasonFilesNotEqual,
			targetReadFile,
			targetWriteFile)

		return

	}

	return
}

func TestFileIoWriter_ReadFrom_000200(t *testing.T) {

	funcName := "TestFileIoWriter_ReadFrom_000200()"

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+10)

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}
	var doesFileExist bool
	var fHelper = new(FileHelper)
	var readFileInfoPlus FileInfoPlus

	doesFileExist,
		readFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetReadFile,
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if doesFileExist == false {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The Target Read File Does NOT Exist!\n"+
			"Target Read File was not found on attached storage drive.\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			dashLineStr,
			targetReadFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileIoWriter_ReadFrom_000200.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetIoReader FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			4096,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetIoWriter FileIoWriter

	_,
		targetIoWriter,
		err = new(FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,  // Default Buffer Size
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoWriter.
		ReadFrom(
			targetIoReader)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.ReadFrom()\n"+
			"numOfBytesProcessed != readFileInfoPlus Size!\n"+
			"numOfBytesProcessed= %v\n"+
			"targetReadFile Size= %v\n",
			ePrefix.String(),
			dashLineStr,
			numOfBytesProcessed,
			readFileInfoPlus.Size())

		return
	}

	var writeFileInfoPlus FileInfoPlus

	writeFileInfoPlus,
		err = fHelper.
		GetFileInfoPlus(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if readFileInfoPlus.Size() != writeFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Target Read File Size != Target Write File Size!\n"+
			" Target Read File Size= %v\n"+
			"Target Write File Size= %v\n"+
			" Target Read File= %v\n"+
			"Target Write File= %v\n\n",
			ePrefix.String(),
			dashLineStr,
			readFileInfoPlus.Size(),
			writeFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	var filesAreEqual bool
	var reasonFilesNotEqual string

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.
		CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Target Files Comparison"))

	if err != nil {

		t.Errorf(" %v\n"+
			" Error Return from fHelper.CompareFiles()\n"+
			"  targetReadFile= %v\n"+
			" targetWriteFile= %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile)

		return
	}

	if !filesAreEqual {

		t.Errorf(" %v\n"+
			"%v\n"+
			" Error: Read and Write Files are NOT equal!\n"+
			" Reason: %v\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n\n",
			ePrefix.String(),
			dashLineStr,
			reasonFilesNotEqual,
			targetReadFile,
			targetWriteFile)

		return

	}

	return
}
