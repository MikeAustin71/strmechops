package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"testing"
)

func TestFileHelper_ReadLines_000100(t *testing.T) {

	funcName := "TestFileHelper_ReadLines_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile, trashDirectory, outputFile string
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

	outputFile = trashDirectory +
		string(os.PathSeparator) +
		"testReadLines_000100.txt"

	var fHelper = new(FileHelper)

	var outputLinesArray,
		endOfLineDelimiters StringArrayDto

	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int

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

	if originalFileSize != 1299 {

		t.Errorf("%v\n"+
			"Error: Original File Size of the\n"+
			"target read file is incorrect!\n"+
			"The original file size of the\n"+
			"the target read file should be %v-bytes.\n"+
			"Instead, original file size is %v-bytes\n",
			ePrefix.String(),
			1299,
			originalFileSize)

		return

	}

	if numOfLinesRead != 23 {

		t.Errorf("%v\n"+
			"Error: Number of lines read from the\n"+
			"target read file is incorrect!\n"+
			"Number of lines parsed and read from\n"+
			"the target read file should be '%v'.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			23,
			numOfLinesRead)

		return

	}

	if numOfBytesRead != 1184 {

		t.Errorf("%v\n"+
			"Error: Number of bytes read from the\n"+
			"target read file is incorrect!\n"+
			"Number of bytes read from the target\n"+
			"read file should be %v.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			1184,
			numOfBytesRead)

	}

	var numBytesWritten int64

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		outputFile,
		true,
		true,
		outputLinesArray.ConcatenateStrings("\n"),
		ePrefix.XCpy("outputFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numBytesWritten != 1207 {

		t.Errorf("%v\n"+
			"Error: Number of bytes written to output\n"+
			"file is incorrect!\n"+
			"Number of bytes written to output file\n"+
			"should be %v.\n"+
			"Instead, Number of Bytes Written= '%v'\n",
			ePrefix.String(),
			numBytesWritten,
			1207)

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

}

func TestFileHelper_ReadTextLines_000100(t *testing.T) {

	funcName := "TestFileHelper_ReadTextLines_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetReadFile, trashDirectory, outputFile string
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

	outputFile = trashDirectory +
		string(os.PathSeparator) +
		"TestFileHelper_ReadTextLines_000100.txt"

	var fHelper = new(FileHelper)

	var outputLinesArray,
		readEndOfLineDelimiters StringArrayDto

	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int

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

	if originalFileSize != 1228 {

		t.Errorf("%v\n"+
			"Error: Original File Size of the\n"+
			"target read file is incorrect!\n"+
			"The original file size of the\n"+
			"the target read file should be %v-bytes.\n"+
			"Instead, original file size is %v-bytes\n",
			ePrefix.String(),
			1228,
			originalFileSize)

		return

	}

	if numOfLinesRead != 22 {

		t.Errorf("%v\n"+
			"Error: Number of lines read from the\n"+
			"target read file is incorrect!\n"+
			"Number of lines parsed and read from\n"+
			"the target read file should be '%v'.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			22,
			numOfLinesRead)

		return

	}

	if numOfBytesRead != 1184 {

		t.Errorf("%v\n"+
			"Error: Number of bytes read from the\n"+
			"target read file is incorrect!\n"+
			"Number of bytes read from the target\n"+
			"read file should be %v.\n"+
			"Instead, Number of Bytes Read= '%v'\n",
			ePrefix.String(),
			1184,
			numOfBytesRead)

	}

	var numBytesWritten int64

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		outputFile,
		true,
		true,
		outputLinesArray.ConcatenateStrings("\r\n"),
		ePrefix.XCpy("outputFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numBytesWritten != 1228 {

		t.Errorf("%v\n"+
			"Error: Number of bytes written to output\n"+
			"file is incorrect!\n"+
			"Number of bytes written to output file\n"+
			"should be %v.\n"+
			"Instead, Number of Bytes Written= '%v'\n",
			ePrefix.String(),
			numBytesWritten,
			1228)

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

}
