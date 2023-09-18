package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
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
			"Target Read File was not found on an attached storage drive.\n"+
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

	err = targetIoReader.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoReader.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoReader!\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetReadFile,
			err.Error())

		return
	}

	err = targetIoWriter.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoWriter!\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile,
			err.Error())

		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

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

func TestFileIoWriter_Seek_000300(t *testing.T) {

	funcName := "TestFileIoWriter_Seek_000300()"

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
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
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

	var compareFile string

	compareFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFileWriteSeek12.txt",
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	doesFileExist,
		_,
		err = fHelper.
		DoesFileInfoPlusExist(
			compareFile,
			ePrefix.XCpy("compareFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if doesFileExist == false {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The Comparison File Does NOT Exist!\n"+
			"The Comparison File was not found on an attached storage drive.\n"+
			"Comparison File: %v\n",
			ePrefix.String(),
			dashLineStr,
			compareFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileIoWriter_Seek_000300.txt",
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
			2048,
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
			2048,  // Default Buffer Size
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoReader.
		WriteTo(
			targetIoWriter)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = targetIoReader.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoReader.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoReader!\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetReadFile,
			err.Error())

		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoReader.WriteTo()\n"+
			"The Number of Bytes Processed is NOT EQUAL\n"+
			"to the size of the Target Read File.\n"+
			"Number of Bytes Processed= '%v'\n"+
			"    Target Readfile Size = '%v'\n"+
			" Target Read File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			numOfBytesProcessed,
			readFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	var i64RequestedWriteFileOffset, i64ActualWriteFileOffset int64

	i64RequestedWriteFileOffset = 12

	i64ActualWriteFileOffset,
		err = targetIoWriter.Seek(
		i64RequestedWriteFileOffset,
		io.SeekStart)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if i64RequestedWriteFileOffset !=
		i64ActualWriteFileOffset {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Seek()\n"+
			"The Requested File Offset is NOT EQUAL\n"+
			"to the Actual File Offset!\n"+
			"Requested File Offset= '%v'\n"+
			"   Actual File Offset= '%v'\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			i64RequestedWriteFileOffset,
			i64ActualWriteFileOffset,
			targetWriteFile)

		return
	}

	var testStr = "Hello World - How are you?"
	lenTestStr := len(testStr)

	var bytesToWrite = []byte(testStr)
	var localNumOfBytesWritten int

	localNumOfBytesWritten,
		err = targetIoWriter.Write(bytesToWrite)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: targetIoWriter.Write(bytesToWrite)\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	err = targetIoWriter.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoWriter!\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile,
			err.Error())

		return
	}

	if localNumOfBytesWritten != lenTestStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Write(bytesToWrite)\n"+
			"Expected Bytes Written DOES NOT MATCH\n"+
			"Actual Bytes Written!\n"+
			"Expected Bytes Written= '%v'\n"+
			"  Actual Bytes Written= '%v'\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			lenTestStr,
			localNumOfBytesWritten,
			targetWriteFile)

		return
	}

	var reasonFilesNotEqual string
	var filesAreEqual bool

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.CompareFiles(
		compareFile,
		targetWriteFile,
		ePrefix.XCpy(
			"Target Files Comparison"))

	if err != nil {

		t.Errorf(" %v\n"+
			"Error Return from fHelper.CompareFiles()\n"+
			"  targetReadFile= %v\n"+
			" targetWriteFile= %v\n"+
			"     compareFile= %v\n"+
			"Reason: %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile,
			compareFile,
			reasonFilesNotEqual)

		return
	}

	if !filesAreEqual {

		t.Errorf("%v\n"+
			"Error: Comparison and Write Files are NOT equal!\n"+
			"Reason: %v\n"+
			"  Comparison File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			compareFile,
			targetWriteFile)

		return

	}

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile-Final Delete"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}
}

func TestFileIoWriter_Append_000400(t *testing.T) {

	funcName := "TestFileIoWriter_Append_000400()"

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+10)

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile string
	var err error
	var fOpsTestUtil = new(fileOpsTestUtility)

	targetReadFile,
		err = fOpsTestUtil.
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
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

	var compareFile string

	compareFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFileAppendedText.txt",
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	doesFileExist,
		_,
		err = fHelper.
		DoesFileInfoPlusExist(
			compareFile,
			ePrefix.XCpy("compareFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if doesFileExist == false {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The Comparison File Does NOT Exist!\n"+
			"The Comparison File was not found on an attached storage drive.\n"+
			"Comparison File: %v\n",
			ePrefix.String(),
			dashLineStr,
			compareFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = fOpsTestUtil.
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileIoWriter_Append_000400.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var fMgrReadFile FileMgr

	fMgrReadFile,
		err = new(FileMgr).
		New(
			targetReadFile,
			ePrefix.XCpy("fMgrReadFile<-targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoReader FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(FileIoReader).
		NewFileMgr(
			&fMgrReadFile,
			false, // openFileReadWrite
			512,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var fMgrWriteFile FileMgr

	fMgrWriteFile,
		err = new(FileMgr).
		New(
			targetWriteFile,
			ePrefix.XCpy("fMgrWriteFile<-targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoWriter FileIoWriter

	_,
		targetIoWriter,
		err = new(FileIoWriter).
		NewFileMgr(
			&fMgrWriteFile,
			false, // openFileReadWrite
			512,
			true, // Truncate Existing File
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoWriter.
		ReadFrom(
			targetIoReader)

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	err = targetIoReader.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoReader.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoReader!\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetReadFile,
			err.Error())

		return
	}

	err = targetIoWriter.Close()

	if err != nil {
		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoWriter!\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile,
			err.Error())
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.ReadFrom()\n"+
			"The Number of Bytes Processed is NOT EQUAL\n"+
			"to the size of the Target Read File.\n"+
			"Number of Bytes Processed= '%v'\n"+
			"    Target Readfile Size = '%v'\n"+
			" Target Read File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			numOfBytesProcessed,
			readFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	var targetIoWriterTwo FileIoWriter

	_,
		targetIoWriterTwo,
		err = new(FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			256,
			false, // Truncate Existing File
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var testStr = "End-Of-File Text. Hello!"
	lenTestStr := len(testStr)

	var bytesToWrite = []byte(testStr)
	var localNumOfBytesWritten int

	localNumOfBytesWritten,
		err = targetIoWriterTwo.Write(bytesToWrite)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: targetIoWriterTwo.Write(bytesToWrite)\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		_ = targetIoWriterTwo.Close()

		return
	}

	err = targetIoWriterTwo.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriterTwo.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetIoWriter #2!\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile,
			err.Error())

		return
	}

	if localNumOfBytesWritten != lenTestStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoWriter.Write(bytesToWrite)\n"+
			"Expected Bytes Written DOES NOT MATCH\n"+
			"Actual Bytes Written!\n"+
			"Expected Bytes Written= '%v'\n"+
			"  Actual Bytes Written= '%v'\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			lenTestStr,
			localNumOfBytesWritten,
			targetWriteFile)

		return
	}

	var reasonFilesNotEqual string
	var filesAreEqual bool

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.CompareFiles(
		compareFile,
		targetWriteFile,
		ePrefix.XCpy(
			"Target Files Comparison"))

	if err != nil {

		t.Errorf(" %v\n"+
			"Error Return from fHelper.CompareFiles()\n"+
			"  targetReadFile= %v\n"+
			" targetWriteFile= %v\n"+
			"     compareFile= %v\n"+
			"Reason: %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile,
			compareFile,
			reasonFilesNotEqual)

		return
	}

	if !filesAreEqual {

		t.Errorf("%v\n"+
			"Error: Comparison and Write Files are NOT equal!\n"+
			"Reason: %v\n"+
			"  Comparison File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			compareFile,
			targetWriteFile)

		return

	}

	err = new(FileHelper).
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("Final Delete-targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	return
}
