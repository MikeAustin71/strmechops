package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestFileBufferReader_Read_000100(t *testing.T) {

	funcName := "TestFileBufferReader_Read_000100()"

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

	var fBufReader FileBufferReader

	_,
		fBufReader,
		err = new(FileBufferReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			256,
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	bytesReadBuff := make([]byte, 125)

	var totalBytesRead, localBytesRead int
	var err2 error
	var cycleCount int

	for {

		localBytesRead,
			err2 = fBufReader.Read(
			bytesReadBuff)

		totalBytesRead += localBytesRead

		if err2 == io.EOF {

			break
		}

		if err2 != nil {

			t.Errorf("\n%v\n"+
				"Processing error returned by\n"+
				"fBufReader.Read(bytesReadBuff)"+
				"while reading the file.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return
		}

		cycleCount++

		if cycleCount > 30 {

			t.Errorf("\n%v\n"+
				"Error: Read Cycle Count Exceeded Maximum!\n"+
				"Expected less than 30-read cycles.\n"+
				"This cycle count has been execeed.\n"+
				"The read cycle has entered an enless loop!\n"+
				"Target File = '%v'\n",
				ePrefix.String(),
				targetReadFile)

			return

		}
	}

	if totalBytesRead != 1228 {

		t.Errorf("\n%v\n"+
			"Error Reading File!\n"+
			"Expected to read 1,228 bytes.\n"+
			"Instead, total bytes read = '%v'\n"+
			"Target File = '%v'\n",
			ePrefix.String(),
			totalBytesRead,
			targetReadFile)

		return
	}

	return
}

func TestFileBufferReader_Seek_000100(t *testing.T) {

	funcName := "TestFileBufferReader_Read_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fHelper := new(FileHelper)

	var doesFileExist bool
	var readerFileInfoPlus FileInfoPlus

	doesFileExist,
		readerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetReadFile,
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if doesFileExist == false {

		t.Errorf("%v\n"+
			"Error: The test 'read' file does NOT exist!\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			targetReadFile)

		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\"+funcName+".txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fBufReader FileBufferReader
	_,
		fBufReader,
		err = new(FileBufferReader).
		NewPathFileName(
			targetReadFile,
			false, // Open File Read/Write
			512,   // bufSize
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	var fBufWriter FileBufferWriter

	_,
		fBufWriter,
		err = new(FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false,
			512,
			false,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	var targetOffset, offsetFromFileStart int64

	targetOffset = 122

	offsetFromFileStart,
		err = fBufReader.Seek(
		targetOffset,
		io.SeekStart)

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if offsetFromFileStart != targetOffset {

		t.Errorf("%v\n"+
			"Error: fBufReader.Seek()\n"+
			"The target offset is NOT equal to\n"+
			"the actual file offset.\n"+
			"Target File Offset: %v\n"+
			"Actual File Offset: %v\n",
			ePrefix.String(),
			targetOffset,
			offsetFromFileStart)

		return
	}

	var strBuilder strings.Builder
	var numOfBytesRead, numOfBytesWritten int64

	numOfBytesRead,
		err = fBufReader.ReadAllToStrBuilder(
		&strBuilder,
		true,
		ePrefix.XCpy("fBufReader"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	numOfBytesWritten,
		err = fBufWriter.WriteTextOrNumbers(
		&strBuilder,
		"", // writeEndOfLineChars
		"", // writeEndOfTextChars
		true,
		ePrefix.XCpy("fBufWriter"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if numOfBytesRead != numOfBytesWritten {
		t.Errorf(" %v\n"+
			" Error: numOfBytesRead != numOfBytesWritten\n"+
			"    numOfBytesRead = %v\n"+
			" numOfBytesWritten = %v\n",
			ePrefix.String(),
			numOfBytesRead,
			numOfBytesWritten)

		return
	}

	var writerFileInfoPlus FileInfoPlus

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedFileSize := readerFileInfoPlus.Size() -
		targetOffset

	actualFileSize := writerFileInfoPlus.Size()

	if expectedFileSize != actualFileSize {

		t.Errorf("%v\n"+
			"First Test\n"+
			"Error: Incorrect Target Write File size!\n"+
			"Expected Target Write File Size: '%v'\n"+
			"  Actual Target Write File Size: '%v'\n",
			ePrefix.String(),
			expectedFileSize,
			actualFileSize)

		return
	}

	// "This is paragraph-2 and line number-1."
	var numOfLinesRead int
	var readEndOfLineDelimiters,
		outputLinesArray StringArrayDto

	readEndOfLineDelimiters.PushStr(
		"\r\n")

	actualFileSize,
		numOfLinesRead,
		_,
		err = fHelper.ReadTextLines(
		targetWriteFile,
		&readEndOfLineDelimiters,
		&outputLinesArray,
		-1,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedFileSize != actualFileSize {

		t.Errorf("%v\n"+
			"Second Test\n"+
			"Error: Incorrect Target Write File size!\n"+
			"Expected Target Write File Size: '%v'\n"+
			"  Actual Target Write File Size: '%v'\n",
			ePrefix.String(),
			expectedFileSize,
			actualFileSize)

		return
	}

	if numOfLinesRead != 3 {

		t.Errorf("%v\n"+
			"Error: Incorrect number of lines read from Target Write File!\n"+
			"Expected Lines from Target Write File: '%v'\n"+
			"  Actual Lines from Target Write File: '%v'\n",
			ePrefix.String(),
			3,
			numOfLinesRead)

		return
	}

	expectedFirstLineOfText :=
		"This is paragraph-2 and line number-1."

	actualFirstLineOfText :=
		outputLinesArray.StrArray[0]

	if expectedFirstLineOfText != actualFirstLineOfText {

		t.Errorf("%v\n"+
			"Error: Incorrect first text line read from Target Write File!\n"+
			"Expected First Text Line from Target Write File: '%v'\n"+
			"  Actual First Text Line from Target Write File: '%v'\n",
			ePrefix.String(),
			expectedFirstLineOfText,
			actualFirstLineOfText)

		return

	}

	return
}
