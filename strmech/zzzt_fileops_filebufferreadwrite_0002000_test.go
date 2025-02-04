package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"runtime"
	"strings"
	"testing"
)

func TestFileBufferReadWrite_ReadAllText_010100(t *testing.T) {

	funcName := "TestFileBufferReadWrite_ReadAllText_010100()"

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
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var dirHelper = new(DirHelper)

	err = dirHelper.DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var baseReadFile string

	baseReadFile,
		err = fOpsTestUtil.
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var expectedNumOfLines int
	var baseReadFileBytes, adjustedBaseReadFileBytes int64
	var fileExistsOnDisk bool
	var fHelper FileHelper

	expectedNumOfLines = 22

	fileExistsOnDisk,
		baseReadFileBytes,
		err = fHelper.GetBytesInFile(
		baseReadFile,
		ePrefix.XCpy("<-baseReadFile"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if !fileExistsOnDisk {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"Base Read File does NOT Exist on Disk!\n"+
			"Base Read Files= %v\n",
			funcName,
			baseReadFile)

		return
	}

	adjustedBaseReadFileBytes =
		baseReadFileBytes - int64(expectedNumOfLines*2)

	var targetReadFile string

	targetReadFile,
		err = fOpsTestUtil.GetCompositeDir(
		"\\fileOpsTest\\trashDirectory\\FileBuffReadWriteInput01.txt",
		ePrefix)

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var targetReadFileBytes int64

	targetReadFileBytes,
		err = new(FileHelper).CopyFileByIoBuffer(
		baseReadFile,
		targetReadFile,
		nil,
		false,
		ePrefix.XCpy("targetReadFile<-baseReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if targetReadFileBytes != baseReadFileBytes {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\nCopy Operation baseReadFile->targetReadFile FAILED!\n"+
			"\nExpected targetReadFile Size= %v bytes"+
			"\nActual targetReadFile Size  = %v\n",
			funcName,
			baseReadFileBytes,
			targetReadFileBytes)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = fOpsTestUtil.GetCompositeDir(
		"\\fileOpsTest\\trashDirectory\\FileBuffReadWriteOutput01.txt",
		ePrefix)

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var fBufReadWrite *FileBufferReadWrite
	var readerFileInfoPlus FileInfoPlus

	readerFileInfoPlus,
		_,
		fBufReadWrite,
		err = new(FileBufferReadWrite).
		NewPathFileNamesReadWrite(
			targetReadFile,
			false, // openReadFileReadWrite,
			1024,  // readerBuffSize
			targetWriteFile,
			false, //openWriteFileReadWrite
			1024,  // writerBuffSize
			true,  // truncateExistingWriteFile
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	var closeFunc = func() {
		err = fBufReadWrite.Close()
		if err != nil {
			t.Errorf("\n"+
				"Error Returned by %v"+
				"\nError closing 'fBufReadWrite' object."+
				"\ntargetReadFile: %v"+
				"\ntargetWriteFile: %v"+
				"\n%v\n",
				funcName,
				targetReadFile,
				targetWriteFile,
				err.Error())

			return

		}

	}

	defer closeFunc()

	var verifiedReadFile, verifiedWriteFile, sTemp string

	verifiedReadFile = strings.ToLower(fBufReadWrite.GetReadFile())

	sTemp = strings.ToLower(targetReadFile)

	if verifiedReadFile != sTemp {

		t.Errorf("Error: %v\n"+
			"Target Read File Path and Name is NOT equal to\n"+
			"File Path and File Name configured in FileBufferReadWrite.\n"+
			"Initial input Read File Path and File Name: \n+"+
			"%v\n"+
			"Actual Path and File Name configured in FileBufferReadWrite:\n"+
			"%v\n",
			funcName,
			sTemp,
			verifiedReadFile)

		return
	}

	verifiedWriteFile = strings.ToLower(fBufReadWrite.GetWriteFile())

	sTemp = strings.ToLower(targetWriteFile)

	if verifiedWriteFile != sTemp {

		t.Errorf("Error: %v\n"+
			"Target Write File Path and Name is NOT equal to\n"+
			"File Path and File Name configured in FileBufferWriteWrite.\n"+
			"Initial input Write File Path and File Name: \n+"+
			"%v\n"+
			"Actual Path and File Name configured in FileBufferWriteWrite:\n"+
			"%v\n",
			funcName,
			sTemp,
			verifiedWriteFile)

		return
	}

	if readerFileInfoPlus.Size() != baseReadFileBytes {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\nInitial Reader File Size is incorrect.\n"+
			"\nExpected File Size = %v bytes"+
			"\nActual File Size   = %v bytes\n",
			funcName,
			baseReadFileBytes,
			readerFileInfoPlus.Size())

		return
	}

	var readEndOfLineDelimiters StringArrayDto

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
		"[EOL]")

	var outputLinesArray StringArrayDto
	var numOfLinesRead int
	var numOfBytesRead int64

	numOfLinesRead,
		numOfBytesRead,
		err = fBufReadWrite.ReadAllTextLines(
		50000,
		&readEndOfLineDelimiters,
		&outputLinesArray,
		ePrefix.XCpy(funcName))

	if err != nil {
		t.Errorf("\n"+
			"Error returned by %v\n"+
			"%v\n",
			funcName,
			err.Error())
		return
	}

	if numOfLinesRead != expectedNumOfLines {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\nThe number of lines 'read' is incorrect.\n"+
			"\nExpected Number of Lines read  = %v"+
			"\nActual number of lines read    = %v\n",
			funcName,
			expectedNumOfLines,
			numOfLinesRead)

		return

	}

	if numOfBytesRead != adjustedBaseReadFileBytes {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\nThe number of bytes 'read' by ReadAllTextLines() is incorrect.\n"+
			"\nExpected Number of bytes read  = %v bytes"+
			"\nActual number of bytes read    = %v bytes\n",
			funcName,
			adjustedBaseReadFileBytes,
			numOfBytesRead)

		return

	}

	var numOfBytesWritten int64
	// Windows Output Format
	var writeEndOfLineChars string

	if runtime.GOOS == "windows" {
		writeEndOfLineChars = "\r\n"
	} else {
		writeEndOfLineChars = "\n"
	}

	numOfBytesWritten,
		err = fBufReadWrite.WriteTextOrNumbers(
		outputLinesArray,
		writeEndOfLineChars,
		writeEndOfLineChars,
		ePrefix.XCpy(funcName))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = fBufReadWrite.Close()

	if err != nil {
		t.Errorf("Error returned from fBufReadWrite.Close()\n"+
			"Target Read File: %v\n"+
			"Target Write File: %v\n"+
			"\n%v\n",
			targetReadFile,
			targetWriteFile,
			err.Error())
		return
	}

	adjustedBaseReadFileBytes = baseReadFileBytes + 2

	if numOfBytesWritten != adjustedBaseReadFileBytes {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\nThe number of bytes 'written' is incorrect.\n"+
			"\nExpected Number of bytes written = %v bytes"+
			"\nActual number of bytes written   = %v\n",
			funcName,
			adjustedBaseReadFileBytes,
			numOfBytesWritten)

		return

	}

	err = dirHelper.DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n"+
			"Error Returned by %v"+
			"\n%v\n",
			funcName,
			err.Error())

		return
	}

	return
}
