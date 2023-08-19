package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestFileBufferWriter_Write_000100(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_000100()"

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

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_000100.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fHelper = new(FileHelper)
	var textLinesArray StringArrayDto
	var expectedNumOfBytesWritten int

	_,
		_,
		_,
		err = fHelper.ReadLines(
		targetReadFile,
		-1,
		&textLinesArray,
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textLinesArray.AppendSuffix("\n")

	expectedNumOfBytesWritten =
		int(textLinesArray.GetTotalBytesInStrings())

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

	lenStrArray := len(textLinesArray.StrArray)

	for i := 0; i < lenStrArray; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(textLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err =
			fBufWriter.Write(bytesToWrite)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err.Error())

			_ = fBufWriter.Flush(nil)

			_ = fBufWriter.FlushAndClose(nil)

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten

	}

	var err2 error
	var errs []error

	err2 = fBufWriter.Flush(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Flush(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		errs = append(errs, err)
	}

	err2 = fBufWriter.FlushAndClose(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Close(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		errs = append(errs, err)
	}

	if len(errs) > 0 {

		err2 = new(StrMech).ConsolidateErrors(errs)

		t.Errorf("%v\n"+
			"Errors returned from Flush() and Close()\n"+
			"Errors= \n%v\n",
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

	err2 = fHelper.DeleteDirFile(
		targetWriteFile,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirFile(targetWriteFile)\n"+
			"targetWriteFile = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err2.Error())

	}

	if len(errs) > 0 {

		err2 = new(StrMech).ConsolidateErrors(errs)

		t.Errorf("%v\n"+
			"Errors returned from fHelper.DeleteDirFile(targetWriteFile)\n"+
			"Errors= \n%v\n",
			ePrefix.String(),
			err2.Error())

		return
	}

	return
}

func TestFileBufferWriter_Write_000200(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_000200()"

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

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_000200.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fHelper = new(FileHelper)
	var textLinesArray StringArrayDto
	var numOfLinesRead, expectedNumOfBytesWritten int
	var i64numOfBytesRead int64

	_,
		numOfLinesRead,
		i64numOfBytesRead,
		err = fHelper.ReadLines(
		targetReadFile,
		-1,
		&textLinesArray,
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textLinesArray.AppendSuffix("\n")

	expectedNumOfBytesWritten =
		int(i64numOfBytesRead) + numOfLinesRead

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

	for i := 0; i < numOfLinesRead; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(textLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err =
			fBufWriter.Write(bytesToWrite)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err.Error())

			_ = fBufWriter.FlushAndClose(nil)

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten

	}

	var err2 error
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

	_,
		fBufWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			512,
			false,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	for i := 0; i < numOfLinesRead; i++ {

		bytesToWrite = []byte(textLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err =
			fBufWriter.Write(bytesToWrite)

		if err != nil {

			t.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err.Error())

			_ = fBufWriter.Flush(nil)

			_ = fBufWriter.FlushAndClose(nil)

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten

	}

	err2 = fBufWriter.FlushAndClose(ePrefix)

	if err2 != nil {

		t.Errorf("%v\n"+
			"Error returned by fBufWriter.Close(ePrefix) #2\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

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

	if int64(totalNumOfBytesWritten) != actualFileSize {

		t.Errorf("\n%v\n"+
			"Error: totalNumOfBytesWritten != actualFileSize\n"+
			"totalNumOfBytesWritten= '%v'\n"+
			"        actualFileSize= '%v'\n",
			ePrefix.String(),
			totalNumOfBytesWritten,
			actualFileSize)

	}

	err2 = fHelper.DeleteDirFile(
		targetWriteFile,
		ePrefix)

	if err2 != nil {

		t.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirFile(targetWriteFile)\n"+
			"targetWriteFile = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err2.Error())
	}

	return
}

func TestFileBufferWriter_Write_000300(t *testing.T) {

	funcName := "TestFileBufferWriter_Write_000300()"

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

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileBufferWriter_Write_000300.txt",
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

	var endOfLineDelimiters = &StringArrayDto{}
	endOfLineDelimiters.PushStr("\r\n")
	endOfLineDelimiters.PushStr("\n")

	var numOfLinesProcessed, numOfBatchesProcessed int
	var numBytesRead, numBytesWritten int64
	var expectedBytes = 1184
	var expectedLines = 22

	numOfLinesProcessed,
		numOfBatchesProcessed,
		numBytesRead,
		numBytesWritten,
		err = newFBuffReadWrite.
		ReadWriteTextLines(
			endOfLineDelimiters,
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

	if int64(expectedBytes) != numBytesRead {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedBytes NOT EQUAL TO readerFileInfoPlus.Size\n"+
			"expectedBytes= '%v'\n"+
			"readerFileInfoPlus.Size= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedBytes,
			readerFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	if numOfBatchesProcessed < 1 {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"Number Of Batches Processed is INVALID!\n"+
			"numOfBatchesProcessed= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			numOfBatchesProcessed,
			targetReadFile,
			targetWriteFile)

		return

	}

	if int64(expectedBytes) != numBytesWritten {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedBytes NOT EQUAL TO numBytesWritten\n"+
			"  expectedBytes= '%v'\n"+
			"numBytesWritten= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedBytes,
			numBytesWritten,
			targetReadFile,
			targetWriteFile)

		return
	}

	var filesAreEqual = false
	var reasonFilesNotEqual string
	var fHelper = new(FileHelper)

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

	err = fHelper.DeleteDirFile(
		targetWriteFile,
		ePrefix)

	if err != nil {

		t.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirFile(targetWriteFile)\n"+
			"Attempted Target File Deletion FAILED!\n"+
			"targetWriteFile = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	return
}
