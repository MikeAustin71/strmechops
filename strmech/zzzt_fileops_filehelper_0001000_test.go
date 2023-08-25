package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"testing"
)

func TestFileHelper_ReadTextLines_000100(t *testing.T) {

	funcName := "TestFileHelper_ReadTextLines_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile, trashDirectory, targetWriteFile,
		compareFile string
	var err error

	targetReadFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc2.txt",
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var doesFileExist bool
	var readerFileInfoPlus FileInfoPlus
	var fHelper = new(FileHelper)

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

	targetWriteFile = trashDirectory +
		string(os.PathSeparator) +
		"TestFileHelper_ReadTextLines_000100.txt"

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

	var expectedTargetWriteFileInfo FileInfoPlus

	expectedTargetWriteFileInfo,
		err = fHelper.GetFileInfoPlus(
		compareFile,
		ePrefix.XCpy("compareFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var outputLinesArray,
		endOfLineDelimiters StringArrayDto

	var writeFileTextLineTerminator = "\r\n"
	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int
	var expectedNumOfLinesRead = 23

	var expectedNumOfBytesWritten = int(expectedTargetWriteFileInfo.Size())

	var expectedNumOfBytesRead = expectedNumOfBytesWritten -
		(expectedNumOfLinesRead * len(writeFileTextLineTerminator))

	endOfLineDelimiters.AddManyStrings(
		"\r",
		"\r\r",
		"[EOL]")

	originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&endOfLineDelimiters,
		&outputLinesArray,
		-1, // maxNumOfTextLines
		ePrefix.XCpy("outputLinesArray"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())

		return
	}

	if originalFileSize != readerFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"Error: Original File Size of the\n"+
			"target read file is incorrect!\n"+
			"The original file size was read\n"+
			"from fHelper.ReadTextLines()\n"+
			"The original file size of the\n"+
			"the target read file should be %v-bytes.\n"+
			"Instead, original file size is %v-bytes\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			readerFileInfoPlus.Size(),
			originalFileSize,
			targetReadFile)

		return

	}

	if numOfLinesRead != expectedNumOfLinesRead {

		t.Errorf("%v\n"+
			"Error: Number of lines read from the\n"+
			"target read file is incorrect!\n"+
			"Number of lines parsed and read from\n"+
			"the target read file should be '%v'.\n"+
			"Instead, Number of Bytes Read= '%v'\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			expectedNumOfLinesRead,
			numOfLinesRead,
			targetReadFile)

		return

	}

	if numOfBytesRead != int64(expectedNumOfBytesRead) {

		t.Errorf("%v\n"+
			"Error: Number of bytes read from the\n"+
			"target read file is incorrect!\n"+
			"Expected Number Of Bytes Read= '%v'\n"+
			"  Actual Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesRead,
			numOfBytesRead)

	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	var numBytesWritten int64

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		targetWriteFile,
		true,
		true,
		outputLinesArray.ConcatenateStrings(
			writeFileTextLineTerminator),
		ePrefix.XCpy("targetWriteFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numBytesWritten != int64(expectedNumOfBytesWritten) {

		t.Errorf("%v\n"+
			"Error: Number of bytes written to output\n"+
			"file is incorrect!\n"+
			"Number of bytes written to output file\n"+
			"should be %v.\n"+
			"Instead, Number of Bytes Written= '%v'\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			numBytesWritten,
			targetWriteFile)

		return
	}

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
		compareFile,
		targetWriteFile,
		ePrefix.XCpy(
			"Target Files Comparison"))

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
			" Error: Compare File and Write File are NOT equal!\n"+
			" Reason: %v\n"+
			"      Compare File: %v\n"+
			" Target Write File: %v\n\n",
			ePrefix.String(),
			reasonFilesNotEqual,
			compareFile,
			targetWriteFile)

		return

	}

	return
}

func TestFileHelper_ReadTextLines_000200(t *testing.T) {

	funcName := "TestFileHelper_ReadTextLines_000200()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile, trashDirectory, targetWriteFile string
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

	var doesFileExist bool
	var targetReaderFileInfoPlus FileInfoPlus
	var fHelper = new(FileHelper)

	doesFileExist,
		targetReaderFileInfoPlus,
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

	targetWriteFile = trashDirectory +
		string(os.PathSeparator) +
		"TestFileHelper_ReadTextLines_000200.txt"

	var outputLinesArray,
		readEndOfLineDelimiters StringArrayDto

	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int

	var writeFileTextLineTerminator = "\r\n"

	var expectedNumOfLinesRead = 22

	var actualReadFileSize = int(targetReaderFileInfoPlus.Size())

	var expectedNumOfBytesWritten = actualReadFileSize

	var expectedNumOfBytesRead = expectedNumOfBytesWritten -
		(expectedNumOfLinesRead * len(writeFileTextLineTerminator))

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
		"[EOL]")

	originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&readEndOfLineDelimiters,
		&outputLinesArray,
		-1, // maxNumOfLines
		ePrefix.XCpy("outputLinesArray"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())

		return
	}

	if originalFileSize != int64(actualReadFileSize) {

		t.Errorf("%v\n"+
			"Error: Original File Size of the\n"+
			"target read file is incorrect!\n"+
			"Original File Size was taken from\n"+
			"fHelper.ReadTextLines()\n"+
			"The original file size of the\n"+
			"the target read file should be %v-bytes.\n"+
			"Instead, original file size is %v-bytes\n",
			ePrefix.String(),
			actualReadFileSize,
			originalFileSize)

		return

	}

	if numOfLinesRead != expectedNumOfLinesRead {

		t.Errorf("%v\n"+
			"Error: Number of lines read from the\n"+
			"target read file is incorrect!\n"+
			"Number of lines parsed and read from\n"+
			"the target read file should be '%v'.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			expectedNumOfLinesRead,
			numOfLinesRead)

		return

	}

	if numOfBytesRead != int64(expectedNumOfBytesRead) {

		t.Errorf("%v\n"+
			"Error: Number of bytes read from the\n"+
			"target read file is incorrect!\n"+
			"Number of bytes read from the target\n"+
			"read file should be %v.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesRead,
			numOfBytesRead)

	}

	var numBytesWritten int64

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		targetWriteFile,
		true,
		true,
		outputLinesArray.ConcatenateStrings("\r\n"),
		ePrefix.XCpy("targetWriteFile<-"))

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

	if numBytesWritten != int64(expectedNumOfBytesWritten) {

		t.Errorf("%v\n"+
			"Error: Number of bytes written to output\n"+
			"file is incorrect!\n"+
			"Number of bytes written to output file\n"+
			"should be %v.\n"+
			"Instead, Number of Bytes Written= '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			numBytesWritten)

		return
	}

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
