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
	var outputLinesArray, readEndOfLineDelimiters StringArrayDto
	var expectedNumOfBytesWritten int

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

	outputLinesArray.AppendSuffix("\n")

	expectedNumOfBytesWritten =
		int(outputLinesArray.GetTotalBytesInStrings())

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

	lenStrArray := len(outputLinesArray.StrArray)

	for i := 0; i < lenStrArray; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

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

	err2 = fHelper.DeleteDirOrFile(
		targetWriteFile,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirOrFile(targetWriteFile)\n"+
			"targetWriteFile = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err2.Error())

	}

	if len(errs) > 0 {

		err2 = new(StrMech).ConsolidateErrors(errs)

		t.Errorf("%v\n"+
			"Errors returned from fHelper.DeleteDirOrFile(targetWriteFile)\n"+
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
	var outputLinesArray, readEndOfLineDelimiters StringArrayDto
	var numOfLinesRead, expectedNumOfBytesWritten int
	var i64numOfBytesRead int64

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

	outputLinesArray.AppendSuffix("\n")

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

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

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

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

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

	err2 = fHelper.DeleteDirOrFile(
		targetWriteFile,
		ePrefix)

	if err2 != nil {

		t.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirOrFile(targetWriteFile)\n"+
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

	var readEndOfLineDelimiters = &StringArrayDto{}
	var writeEndOfLineChars = "\r\n"
	readEndOfLineDelimiters.PushStr("\r\n")
	readEndOfLineDelimiters.PushStr("\n")

	var numOfLinesProcessed int
	var numBytesRead, numBytesWritten int64
	var expectedBytesRead = 1184
	var expectedFileSize = 1228
	var expectedBytesWritten = 1228

	var expectedLines = 22

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

	if int64(expectedFileSize) != readerFileInfoPlus.Size() {

		t.Errorf("\n%v\n"+
			"Error: newFBuffReadWrite.ReadWriteTextLines()\n"+
			"expectedFileSize NOT EQUAL TO readerFileInfoPlus.Size()\n"+
			"          expectedBytes= '%v'\n"+
			"readerFileInfoPlus.Size= '%v'\n"+
			" Target Read File= '%v'\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			expectedFileSize,
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

	err = new(FileHelper).DeleteDirOrFile(
		targetWriteFile,
		ePrefix)

	if err != nil {

		t.Errorf("%v\n"+
			"Error return from fHelper.DeleteDirOrFile(targetWriteFile)\n"+
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
