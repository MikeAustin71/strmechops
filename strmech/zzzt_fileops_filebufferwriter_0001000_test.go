package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestFileBufferWriter_Append_001000(t *testing.T) {

	funcName := "TestFileBufferWriter_Append_001000()"

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
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Append_001000.txt",
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

	var targetBufioReader FileBufferReader

	readFileInfoPlus,
		targetBufioReader,
		err = new(FileBufferReader).
		NewFileMgr(
			&fMgrReadFile,
			false, // openFileReadWrite
			512,
			ePrefix.XCpy("targetBufioReader<-"))

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

	var targetBufioWriter FileBufferWriter

	_,
		targetBufioWriter,
		err = new(FileBufferWriter).
		NewFileMgr(
			&fMgrWriteFile,
			false, // openFileReadWrite
			512,
			true, // Truncate Existing File
			ePrefix.XCpy("targetBufioWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetBufioWriter.
		ReadFrom(
			targetBufioReader)

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	err = targetBufioReader.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioReader.Close()\n"+
			"Error returned while attempting\n"+
			"to close TargetBufioReader!\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			targetReadFile,
			err.Error())

		return
	}

	err = targetBufioWriter.Close()

	if err != nil {
		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioWriter.Close()\n"+
			"Error returned while attempting\n"+
			"to close targetBufioWriter!\n"+
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
			"Error: targetBufioWriter.ReadFrom()\n"+
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

	var targetBufioWriterTwo FileBufferWriter

	_,
		targetBufioWriterTwo,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			256,
			false, // Truncate Existing File
			ePrefix.XCpy("targetBufioWriterTwo<-"))

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
		err = targetBufioWriterTwo.Write(bytesToWrite)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: targetBufioWriterTwo.Write(bytesToWrite)\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		_ = targetBufioWriterTwo.Close()

		return
	}

	err = targetBufioWriterTwo.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioWriterTwo.Close()\n"+
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
			"Error: targetBufioWriter.Write(bytesToWrite)\n"+
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

func TestFileBufferWriter_Write_001100(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_001100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var trashDirectory string
	var err error

	var fOpsTestUtil = new(fileOpsTestUtility)

	trashDirectory,
		err = fOpsTestUtil.
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

	var targetReadFile string

	targetReadFile,
		err = fOpsTestUtil.
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fHelper = new(FileHelper)

	var doesFileExist bool
	var readerFileInfoPlus FileInfoPlus

	doesFileExist,
		readerFileInfoPlus,
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
			"Error: The Target Read File Does NOT Exist!\n"+
			"Target Read File was not found on attached storage drive.\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			targetReadFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = fOpsTestUtil.
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_001100.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var outputLinesArray, readEndOfLineDelimiters StringArrayDto
	var expectedNumOfBytesWritten int
	var writeFileTextLineTerminator = "\r\n"

	readEndOfLineDelimiters.AddManyStrings(
		"\r\n",
		"\r\r",
		"[EOL]")

	_,
		_,
		_,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&readEndOfLineDelimiters,
		&outputLinesArray,
		-1, // maxNumOfLines
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	outputLinesArray.AppendSuffix(
		writeFileTextLineTerminator)

	expectedNumOfBytesWritten =
		int(readerFileInfoPlus.Size())

	var fBufWriter FileBufferWriter

	_,
		fBufWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			512,
			true,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var totalNumOfBytesWritten, localNumOfBytesWritten int

	var bytesToWrite []byte

	var err2 error

	lenStrArray := len(outputLinesArray.StrArray)

	for i := 0; i < lenStrArray; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err2 =
			fBufWriter.Write(bytesToWrite)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err2.Error())

			err2 = fBufWriter.Flush("After Error: fBufWriter")

			if err2 != nil {
				err = errors.Join(err, err2)
			}

			err2 = fBufWriter.FlushCloseRelease("After Error: fBufWriter")

			if err2 != nil {
				err = errors.Join(err, err2)
			}

			t.Errorf("%v\n", err.Error())

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten

	}

	err2 = fBufWriter.Flush(ePrefix.XCpy(
		"After Loop fBufWriter"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Flush()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		err = errors.Join(err, err2)
	}

	err2 = fBufWriter.Close()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Close()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		err = errors.Join(err, err2)
	}

	if err != nil {

		t.Errorf("%v\n"+
			"Errors returned from Flush() and Close()\n"+
			"Errors= \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if expectedNumOfBytesWritten != totalNumOfBytesWritten {

		t.Errorf("%v\n"+
			"Error: Expected Bytes Written != Actual Bytes Written\n"+
			"Expected Bytes Written = '%v'\n"+
			"  Actual Bytes Written = '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			totalNumOfBytesWritten)

		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	var writerFileInfoPlus FileInfoPlus

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if !doesFileExist {

		t.Errorf("%v\n"+
			"Error: Target Write File Does NOT Exist!\n"+
			"The Target Write File Was NOT Found on attached storage drives\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			targetWriteFile)

		return
	}

	if int64(expectedNumOfBytesWritten) != writerFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"Error: The Expected Number Of Bytes Written\n"+
			"Does NOT match the size of the 'write' file.\n"+
			"  Expected Number Of Bytes Written= '%v'\n"+
			"Size Of Target Write File in Bytes= '%v'\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			writerFileInfoPlus.Size(),
			targetWriteFile)

		return
	}

	var reasonFilesNotEqual string
	var filesAreEqual bool

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.CompareFiles(
		targetReadFile,
		targetWriteFile,
		ePrefix.XCpy(
			"Target Files Comparison"))

	if err != nil {

		t.Errorf(" %v\n"+
			"Error Return from fHelper.CompareFiles()\n"+
			"  targetReadFile= %v\n"+
			" targetWriteFile= %v\n"+
			"Reason: %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile,
			reasonFilesNotEqual)

		return
	}

	if !filesAreEqual {

		t.Errorf(" %v\n"+
			" Error: Read and Write Files are NOT equal!\n"+
			" Reason: %v\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			targetReadFile,
			targetWriteFile)

		return

	}

	return
}

func TestFileBufferWriter_Write_002200(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_002200()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile, compareFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc2.txt",
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var doesFileExist bool
	var fHelper = new(FileHelper)
	var targetReadFileInfo FileInfoPlus

	doesFileExist,
		targetReadFileInfo,
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
			"Error: The Target Read File Does NOT Exist!\n"+
			"Target Read File was not found on attached storage drive.\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			targetReadFile)

		return
	}

	compareFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFuncBlankLastLine.txt",
			ePrefix.XCpy("compareFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var compareFileInfo FileInfoPlus

	doesFileExist,
		compareFileInfo,
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
			"Error: The Comparison Read File Does NOT Exist!\n"+
			"The Comparison File was not found on attached storage drive.\n"+
			"Comparison File: %v\n",
			ePrefix.String(),
			compareFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_002200.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var outputLinesArray, readEndOfLineDelimiters StringArrayDto
	var numOfLinesRead int
	var i64numOfBytesRead int64
	var writeFileTextLineTerminator = "\r\n"
	var expectedNumOfLinesRead = int64(23)
	var expectedNumOfBytesWritten = compareFileInfo.Size()
	var expectedNumOfTextLineBytesRead = expectedNumOfBytesWritten -
		(expectedNumOfLinesRead * int64(len(writeFileTextLineTerminator)))

	readEndOfLineDelimiters.AddManyStrings(
		"\r\n",
		"\r\r",
		"[EOL]")

	_,
		numOfLinesRead,
		i64numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&readEndOfLineDelimiters,
		&outputLinesArray,
		-1,
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if i64numOfBytesRead != expectedNumOfTextLineBytesRead {

		t.Errorf("\n%v\n"+
			"Error: Number of Bytes Read from Target Read File\n"+
			"Does NOT match the number of bytes in Target Read File.\n"+
			"            Target Read File Size in Bytes= %v\n"+
			"Number of Bytes Read from Target Read File= %v\n",
			ePrefix.String(),
			targetReadFileInfo.Size(),
			i64numOfBytesRead)

		return
	}

	outputLinesArray.AppendSuffix(
		writeFileTextLineTerminator)

	var fBufWriter FileBufferWriter

	_,
		fBufWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			512,
			true,
			ePrefix.XCpy("targetWriteFile"))

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

	var localNumOfBytesWritten int
	var totalNumOfBytesWritten int64
	var err2 error

	var bytesToWrite []byte

	for i := 0; i < numOfLinesRead; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err2 =
			fBufWriter.Write(bytesToWrite)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err2.Error())

			err2 = fBufWriter.FlushCloseRelease(ePrefix.XCpy(
				"After Error-fBufWriter"))

			err = errors.Join(err, err2)

			t.Errorf("%v\n",
				err.Error())

			return
		}

		totalNumOfBytesWritten += int64(localNumOfBytesWritten)

	}

	err2 = fBufWriter.Close()

	if err2 != nil {

		t.Errorf("%v\n"+
			"Error returned by fBufWriter.Close(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		return
	}

	if expectedNumOfBytesWritten != totalNumOfBytesWritten {

		t.Errorf("%v\n"+
			"Error: Expected Bytes Written != Actual Bytes Written\n"+
			"Expected Bytes Written = '%v'\n"+
			"  Actual Bytes Written = '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			totalNumOfBytesWritten)

		return
	}

	var fileInfoPlus FileInfoPlus

	fileInfoPlus,
		err = fHelper.GetFileInfoPlus(
		targetWriteFile,
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualFileSize int64

	actualFileSize = fileInfoPlus.Size()

	if totalNumOfBytesWritten != actualFileSize {

		t.Errorf("\n%v\n"+
			"Error: totalNumOfBytesWritten != actualFileSize\n"+
			"totalNumOfBytesWritten= '%v'\n"+
			"        actualFileSize= '%v'\n",
			ePrefix.String(),
			totalNumOfBytesWritten,
			actualFileSize)

	}

	var reasonFilesNotEqual string
	var filesAreEqual bool

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.CompareFiles(
		compareFile,
		targetWriteFile,
		ePrefix.XCpy(
			"compareFile vs targetWriteFile"))

	if err != nil {

		t.Errorf(" %v\n"+
			"Error Return from fHelper.CompareFiles()\n"+
			"     compareFile= %v\n"+
			" targetWriteFile= %v\n"+
			"Reason: %v\n",
			ePrefix.String(),
			compareFile,
			targetWriteFile,
			reasonFilesNotEqual)

		return
	}

	if !filesAreEqual {

		t.Errorf(" %v\n"+
			" Error: Comparison File and Write File are NOT equal!\n"+
			" Reason: %v\n"+
			"   Comparison File: %v\n"+
			" Target Write File: %v\n\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			compareFile,
			targetWriteFile)

		return

	}

	return
}

func TestFileBufferWriter_Write_003300(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_003300()"

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

	var fHelper = new(FileHelper)
	var doesFileExist bool
	var originalReaderFileInfoPlus FileInfoPlus

	doesFileExist,
		originalReaderFileInfoPlus,
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
			"Error: The Target Read File Does NOT Exist!\n"+
			"Target Read File was not found on attached storage drive.\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			targetReadFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_003300.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var newFBuffReadWrite FileBufferReadWrite
	var readerFileInfoPlus FileInfoPlus

	readerFileInfoPlus,
		_,
		newFBuffReadWrite,
		err = new(FileBufferReadWrite).
		NewPathFileNames(
			targetReadFile,
			false,
			512,
			targetWriteFile,
			false,
			1024,
			true,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if originalReaderFileInfoPlus.Size() !=
		readerFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"Error: The original 'reader' file size in bytes\n"+
			"does NOT match the file size returned by method\n"+
			"FileBufferReadWrite).NewPathFileNames().\n"+
			"Target Read File: %v\n"+
			"Original Size in Bytes: %v\n"+
			"NewPathFileNames() File Size in Bytes: %v\n",
			ePrefix.String(),
			targetReadFile,
			originalReaderFileInfoPlus.Size(),
			readerFileInfoPlus.Size())

		return
	}

	var readEndOfLineDelimiters = &StringArrayDto{}
	var writeEndOfLineChars = "\r\n"
	readEndOfLineDelimiters.PushStr("\r\n")
	readEndOfLineDelimiters.PushStr("\n")

	var numOfLinesProcessed int
	var numBytesRead, numBytesWritten int64
	var expectedLines = 22

	var expectedReadFileSize = int(originalReaderFileInfoPlus.Size())

	var expectedBytesRead = expectedReadFileSize - (expectedLines * len(writeEndOfLineChars))

	var expectedBytesWritten = expectedReadFileSize

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	numOfLinesProcessed,
		numBytesRead,
		numBytesWritten,
		err = newFBuffReadWrite.
		ReadWriteTextLines(
			readEndOfLineDelimiters,
			writeEndOfLineChars,
			-1,
			true,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedLines != numOfLinesProcessed {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedLines NOT EQUAL TO numOfLinesProcessed\n"+
			"      expectedLines= '%v'\n"+
			"numOfLinesProcessed= '%v'\n",
			ePrefix.String(),
			expectedLines,
			numOfLinesProcessed)

		return
	}

	if int64(expectedBytesRead) != numBytesRead {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedBytesRead NOT EQUAL TO numBytesRead\n"+
			"expectedBytesRead= '%v'\n"+
			"     numBytesRead= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedBytesRead,
			numBytesRead,
			targetReadFile,
			targetWriteFile)

		return
	}

	if int64(expectedReadFileSize) != readerFileInfoPlus.Size() {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedReadFileSize NOT EQUAL TO readerFileInfoPlus.Size()\n"+
			"          expectedBytes= '%v'\n"+
			"readerFileInfoPlus.Size= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedReadFileSize,
			readerFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	if int64(expectedBytesWritten) != numBytesWritten {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedBytes NOT EQUAL TO numBytesWritten\n"+
			"expectedBytesWritten= '%v'\n"+
			"     numBytesWritten= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedBytesWritten,
			numBytesWritten,
			targetReadFile,
			targetWriteFile)

		return
	}

	var filesAreEqual = false
	var reasonFilesNotEqual string

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.
		CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !filesAreEqual {

		t.Errorf("\n%v\n"+
			"Error: FileHelper.CompareFiles()\n"+
			"Target Read and Write Files ARE NOT EQUAL!\n"+
			"Reason Files Are NOT Equal= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			targetReadFile,
			targetWriteFile)

		return

	}

	return
}

func TestFileBufferWriter_ReadFrom_004400(t *testing.T) {

	funcName := "TestFileBufferWriter_ReadFrom_004400()"

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
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_ReadFrom_004400.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetIoReader *FileIoReader

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

	var targetBufioWriter FileBufferWriter

	_,
		targetBufioWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			256,   // Default Buffer Size
			true,  // Truncate Existing File
			ePrefix.XCpy("targetBufioWriter<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetBufioWriter.
		ReadFrom(
			targetIoReader)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())

		_ = targetIoReader.Close()

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

		_ = targetBufioWriter.Close()

		return
	}

	err = targetBufioWriter.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioWriter.Close()\n"+
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
			"Error: targetBufioWriter.ReadFrom()\n"+
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

func TestFileBufferWriter_Seek_005500(t *testing.T) {

	funcName := "TestFileBufferWriter_Seek_005500()"

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
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Seek_005500.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetBufioReader FileBufferReader

	readFileInfoPlus,
		targetBufioReader,
		err = new(FileBufferReader).
		NewPathFileName(
			targetReadFile,
			true, // openFileReadWrite
			1024, // Buffer Size
			ePrefix.XCpy("targetBufioReader<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetBufioWriter FileBufferWriter
	var err2 error

	_,
		targetBufioWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			2048,  // Default Buffer Size
			true,  // Truncate Existing File
			ePrefix.XCpy("targetBufioWriter<-"))

	if err != nil {

		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetBufioReader.
		WriteTo(
			targetBufioWriter)

	if err != nil {

		err2 = targetBufioReader.Close()

		if err2 != nil {

			err = errors.Join(err, err2)

		}

		t.Errorf("\n%v\n",
			err.Error())

		return
	}

	err = targetBufioReader.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioReader.Close()\n"+
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

	err = targetBufioWriter.Flush(
		ePrefix.XCpy("targetBufioWriter"))

	if err != nil {

		err2 = targetBufioWriter.Close()

		err = errors.Join(err, err2)

		t.Errorf("\n%v\n",
			err.Error())

		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioReader.WriteTo()\n"+
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
		err = targetBufioWriter.Seek(
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
			"Error: targetBufioWriter.Seek()\n"+
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
		err = targetBufioWriter.Write(bytesToWrite)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: targetBufioWriter.Write(bytesToWrite)\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	err = targetBufioWriter.Close()

	if err != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetBufioWriter.Close()\n"+
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
			"Error: targetBufioWriter.Write(bytesToWrite)\n"+
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

	return
}
