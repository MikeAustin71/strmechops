package examples

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"os"
	"strings"
)

type MainFileReadWriteTest010 struct {
	input string
}

func (fileReadWriteTest010 MainFileReadWriteTest010) FileBufferReader03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileReadWriteTest010.FileBufferReader03()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fBufReader strmech.FileBufferReader

	_,
		fBufReader,
		err = new(strmech.FileBufferReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			512,
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	bytesReadBuff := make([]byte, 425)

	var totalBytesRead, localBytesRead int
	var err2 error

	for {

		localBytesRead,
			err2 = fBufReader.Read(
			bytesReadBuff)

		totalBytesRead += localBytesRead

		if err2 == io.EOF {

			break
		}

		if err2 != nil {

			fmt.Printf("\n%v\n"+
				"Processing error returned by\n"+
				"fBufReader.Read(bytesReadBuff)"+
				"while reading the file.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return
		}

	}

	if totalBytesRead != 1228 {

		fmt.Printf("\n%v\n"+
			"Error Reading File!\n"+
			"Expected to read 1,228 bytes.\n"+
			"Instead, total bytes read = '%v'\n"+
			"Target File = '%v'\n",
			ePrefix.String(),
			totalBytesRead,
			targetReadFile)

		return

	} else {

		fmt.Printf("\nTotal Bytes Read = '1228'\n" +
			"The file read is CORRECT!\n")
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}

func (fileReadWriteTest010 MainFileReadWriteTest010) FileBufferReader04() {

	funcName := "MainFileReadWriteTest010.FileBuffReadWrite03()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\"+
			"Main010FileReader04.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var readEndOfLineDelimiters,
		outputLinesArray strmech.StringArrayDto
	var expectedNumOfBytesWritten, numOfLinesRead int
	var originalFileSize, numOfBytesRead int64
	var shouldReadAndWriteFilesBeEqual,
		useWindowsLineTerminationChars,
		shouldFinalDeleteWriteFile bool
	var writeEndOfLineChars string

	shouldReadAndWriteFilesBeEqual = true

	useWindowsLineTerminationChars = true

	shouldFinalDeleteWriteFile = true

	if useWindowsLineTerminationChars {

		// Windows Output Format
		writeEndOfLineChars = "\r\n"

	} else {

		// Unix Output Format
		writeEndOfLineChars = "\n"
	}

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
		-1, // maxNumOfTextLines
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	lenTextLineArray := len(outputLinesArray.StrArray)

	fmt.Printf("\n %v\n"+
		"%v\n"+
		" fHelper.ReadTextLines() Stats\n"+
		"                originalFileSize = %v\n"+
		"                  numOfBytesRead = %v\n"+
		" Number of Bytes Text Line Array = %v\n"+
		"                  numOfLinesRead = %v\n"+
		"   Actual length Text Line Array = %v\n\n",
		ePrefix.String(),
		dashLineStr,
		originalFileSize,
		numOfBytesRead,
		outputLinesArray.GetTotalBytesInStrings(),
		numOfLinesRead,
		lenTextLineArray)

	outputLinesArray.AppendSuffix(writeEndOfLineChars)

	expectedNumOfBytesWritten =
		int(outputLinesArray.GetTotalBytesInStrings())

	var fBufWriter strmech.FileBufferWriter

	_,
		fBufWriter,
		err = new(strmech.FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			512,
			true,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var totalNumOfBytesWritten, localNumOfBytesWritten int

	var bytesToWrite []byte
	var err2, err3 error
	lenStrArray := len(outputLinesArray.StrArray)

	for i := 0; i < lenStrArray; i++ {

		bytesToWrite = make([]byte, 0)

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err3 =
			fBufWriter.Write(bytesToWrite)

		if err3 != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err3.Error())

			err = errors.Join(err, err2)

			err3 = fBufWriter.Flush(ePrefix)

			err = errors.Join(err, err3)

			err3 = fBufWriter.FlushAndClose(ePrefix)

			err = errors.Join(err, err3)

			fmt.Printf("%v",
				err.Error())

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten

	}

	err = fBufWriter.FlushAndClose(
		ePrefix.XCpy("fBufWriter"))

	if err != nil {

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}

	}

	fmt.Printf("\n %v\n"+
		"%v"+
		" Write Stats\n"+
		"Expected Num Of Bytes Written = %v\n"+
		"  Actual Num Of Bytes Written = %v\n\n",
		ePrefix.String(),
		dashLineStr,
		expectedNumOfBytesWritten,
		totalNumOfBytesWritten)

	if shouldReadAndWriteFilesBeEqual == true {

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

			fmt.Printf(" %v\n"+
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

			fmt.Printf("%v\n"+
				"Error: Read and Write Files are NOT equal!\n"+
				"Reason: %v\n",
				ePrefix.String(),
				reasonFilesNotEqual)

			return

		} else {

			fmt.Printf(" %v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile)

		}
	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}
	}

	// --------------------------------------------
	// Successful Completion
	// --------------------------------------------
	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// FileBuffReadWrite01
//
// Example of FileBufferReadWrite
func (fileReadWriteTest010 MainFileReadWriteTest010) FileBuffReadWrite01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileReadWriteTest010.FileBuffReadWrite01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\FileBuffReadWrite01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var newFBuffReadWrite strmech.FileBufferReadWrite
	var readerFileInfoPlus strmech.FileInfoPlus
	var writerFileInfoPlus strmech.FileInfoPlus

	readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err = new(strmech.FileBufferReadWrite).
		NewPathFileNames(
			targetReadFile,
			false, // openReadFileReadWrite,
			512,   // readerBuffSize
			targetWriteFile,
			false, //openWriteFileReadWrite
			512,   // writerBuffSize
			true,  // truncateExistingWriteFile
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nNew File Stats\n"+
		"FileBufferReadWrite.NewPathFileNames()\n"+
		"Reader File Size= %v\n"+
		"Writer File Size= %v\n"+
		" Target Read File: %v\n"+
		"Target Write File: %v\n\n",
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size(),
		targetReadFile,
		targetWriteFile)

	var totalBytesRead, totalBytesWritten int

	totalBytesRead,
		totalBytesWritten,
		err = newFBuffReadWrite.ReadWriteAll(
		true,
		ePrefix.XCpy(
			"newFBuffReadWrite"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if totalBytesRead != totalBytesWritten {

		fmt.Printf("%v\n"+
			"Error: totalBytesRead != totalBytesWritten\n"+
			" Read File= %v\n"+
			"Write File= %v\n"+
			"   Total Bytes Read= %v\n"+
			"Total Bytes Written= %v\n",
			ePrefix.String(),
			targetReadFile,
			targetWriteFile,
			totalBytesRead,
			totalBytesWritten)

		return
	}

	var fHelper = new(strmech.FileHelper)
	var doesFileExist bool

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.DoesFileInfoPlusExist(
		targetWriteFile,
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if !doesFileExist {

		fmt.Printf("%v\n"+
			"Error: Target Write File DOES NOT EXIST!\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			targetWriteFile)

		return
	}

	if writerFileInfoPlus.Size() != readerFileInfoPlus.Size() {

		fmt.Printf("%v\n"+
			"Error: The size of the Read and Write Files\n"+
			"DO NOT MATCH!\n"+
			"Read File Size= %v\n"+
			"Write File Size= %v\n",
			ePrefix.String(),
			readerFileInfoPlus.Size(),
			writerFileInfoPlus.Size())

		return
	}

	fmt.Printf("Final Read & Write File Sizes\n"+
		"Read File Size= %v\n"+
		"Write File Size= %v\n\n",
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size())

	var filesAreEqual bool
	var reasonFilesNotEqual string

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.CompareFiles(
		targetReadFile,
		targetWriteFile,
		ePrefix.XCpy("Target Files"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if !filesAreEqual {

		fmt.Printf("%v\n"+
			"Error: Read and Write Files are NOT equal!\n"+
			"Reason: %v\n",
			ePrefix.String(),
			reasonFilesNotEqual)

		return
	}

	err = fHelper.DeleteDirOrFile(
		targetWriteFile,
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	// --------------------------------------------
	// Successful Completion
	// --------------------------------------------
	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// FileBuffReadWrite02
//
// Reads text lines
func (fileReadWriteTest010 MainFileReadWriteTest010) FileBuffReadWrite02() {

	funcName := "MainFileReadWriteTest010.FileBuffReadWrite02()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc2.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\FileBuffReadWrite02.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var newFBuffReadWrite strmech.FileBufferReadWrite
	var readerFileInfoPlus, writerFileInfoPlus strmech.FileInfoPlus

	readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err = new(strmech.FileBufferReadWrite).
		NewPathFileNames(
			targetReadFile,
			false, // openReadFileReadWrite,
			512,   // readerBuffSize
			targetWriteFile,
			false, //openWriteFileReadWrite
			512,   // writerBuffSize
			true,  // truncateExistingWriteFile
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nStats From NewPathFileNames\n"+
		"Reader File Size: %v\n"+
		"Writer File Size: %v\n\n",
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size())

	var numOfBytesWritten, numOfBytesRead int64
	var numOfLinesProcessed, numOfBatchesProcessed int
	var readEndOfLineDelimiters strmech.StringArrayDto
	var writeEndOfLineChars = "\r\n"

	readEndOfLineDelimiters.AddManyStrings(
		"\r",
		"\r\r",
		"[EOL]")

	maxNumOfTextLines := -1

	numOfLinesProcessed,
		numOfBytesRead,
		numOfBytesWritten,
		err = newFBuffReadWrite.
		ReadWriteTextLines(
			&readEndOfLineDelimiters,
			writeEndOfLineChars,
			maxNumOfTextLines, // numTextLinesPerBatch
			true,              // autoFlushAndCloseOnExit
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Stats From newFBuffReadWrite.ReadWriteTextLines()\n"+
		"   Number Of Lines Processed: %v\n"+
		" Number Of Batches Processed: %v\n"+
		"Maximum Number of Text Lines: %v\n"+
		"        Number Of Bytes Read: %v\n"+
		"     Number Of Bytes Written: %v\n\n",
		numOfLinesProcessed,
		numOfBatchesProcessed,
		maxNumOfTextLines,
		numOfBytesRead,
		numOfBytesWritten)

	fHelper := new(strmech.FileHelper)

	var doesFileExist bool

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Write File Stats\n"+
		"   Write File Does Exist: %v\n"+
		"Write File Size in Bytes: %v\n"+
		"         Write File Name: %v\n\n",
		doesFileExist,
		writerFileInfoPlus.Size(),
		writerFileInfoPlus.Name())

	err = newFBuffReadWrite.IsValidInstanceError(
		ePrefix)

	if err == nil {

		fmt.Printf("%v\n"+
			"newFBuffReadWrite Status Error\n"+
			"It was expected that 'newFBuffReadWrite' would\n"+
			"be closed and invalid. However, after being\n"+
			"closed, 'newFBuffReadWrite' is showing as valid.\n"+
			"This is an error condition!\n",
			ePrefix)

		return
	}

	// ------ Trailing Marquee

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// FileBuffReadWrite02B
//
// Reads text lines using [EOL]
func (fileReadWriteTest010 MainFileReadWriteTest010) FileBuffReadWrite02B() {

	funcName := "MainFileReadWriteTest010.FileBuffReadWrite02B()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile2.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\FileBuffReadWrite02B.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var newFBuffReadWrite strmech.FileBufferReadWrite
	var readerFileInfoPlus, writerFileInfoPlus strmech.FileInfoPlus

	readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err = new(strmech.FileBufferReadWrite).
		NewPathFileNames(
			targetReadFile,
			false, // openReadFileReadWrite,
			512,   // readerBuffSize
			targetWriteFile,
			false, //openWriteFileReadWrite
			512,   // writerBuffSize
			true,  // truncateExistingWriteFile
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nNew File Stats\n"+
		"FileBufferReadWrite.NewPathFileNames()\n"+
		"Reader File Size= %v\n"+
		"Writer File Size= %v\n"+
		" Target Read File: %v\n"+
		"Target Write File: %v\n\n",
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size(),
		targetReadFile,
		targetWriteFile)

	var numOfBytesWritten, numOfBytesRead int64
	var numOfLinesProcessed, numOfBatchesProcessed int
	var readEndOfLineDelimiters strmech.StringArrayDto
	var writeEndOfLineChars = "\r\n"

	readEndOfLineDelimiters.AddManyStrings(
		"\r",
		"\r\r",
		"[EOL]")

	maxNumOfTextLines := -1

	numOfLinesProcessed,
		numOfBytesRead,
		numOfBytesWritten,
		err = newFBuffReadWrite.
		ReadWriteTextLines(
			&readEndOfLineDelimiters,
			writeEndOfLineChars,
			maxNumOfTextLines, // numTextLinesPerBatch
			true,              // autoFlushAndCloseOnExit
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nStats From newFBuffReadWrite.ReadWriteTextLines()\n"+
		"   Number Of Lines Processed: %v\n"+
		" Number Of Batches Processed: %v\n"+
		"Maximum Number Of Text Lines: %v\n"+
		"        Number Of Bytes Read: %v\n"+
		"     Number Of Bytes Written: %v\n"+
		" Target Read File: %v\n"+
		"Target Wriet File: %v\n",
		numOfLinesProcessed,
		numOfBatchesProcessed,
		maxNumOfTextLines,
		numOfBytesRead,
		numOfBytesWritten,
		targetReadFile,
		targetWriteFile)

	fHelper := new(strmech.FileHelper)

	var doesFileExist bool

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nWrite File Stats from FileHelper.DoesFileInfoPlusExist()\n"+
		"   Write File Does Exist: %v\n"+
		"Write File Size in Bytes: %v\n"+
		"         Write File Name: %v\n\n",
		doesFileExist,
		writerFileInfoPlus.Size(),
		writerFileInfoPlus.Name())

	err = newFBuffReadWrite.IsValidInstanceError(
		ePrefix)

	if err == nil {

		fmt.Printf("%v\n"+
			"newFBuffReadWrite Status Error\n"+
			"It was expected that 'newFBuffReadWrite' would\n"+
			"be closed and invalid. However, after being\n"+
			"closed, 'newFBuffReadWrite' is showing as valid.\n"+
			"This is an error condition!\n",
			ePrefix)

		return
	}

	var finalWriteFileInfo strmech.FileInfoPlus
	var newFileIoReader strmech.FileIoReader

	finalWriteFileInfo,
		newFileIoReader,
		err = new(strmech.FileIoReader).
		NewPathFileName(
			targetWriteFile,
			false,
			1024,
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nWrite File Stats from FileIoReader\n"+
		"Final Write File Size: %v\n"+
		"         Write File Name: %v\n\n",
		finalWriteFileInfo.Size(),
		targetWriteFile)

	var numOfLinesRead int
	var outputLinesArray = &strmech.StringArrayDto{}
	readEndOfLineDelimiters.Empty()

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
		"[EOL]")

	numOfLinesRead,
		numOfBytesRead,
		err = newFileIoReader.ReadAllTextLines(
		&readEndOfLineDelimiters,
		outputLinesArray,
		-1,   // maxNumOfTextLines
		true, // autoCloseOnExit
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nWrite File Stats from FileIoReader.ReadAllTextLines()\n"+
		" numOfLinesRead: %v\n"+
		" numOfBytesRead: %v\n"+
		"Write File Name: %v\n\n",
		numOfLinesRead,
		numOfBytesRead,
		targetWriteFile)

	// ------ Trailing Marquee

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// FileBuffReadWrite03
//
// Reads text lines
func (fileReadWriteTest010 MainFileReadWriteTest010) FileBuffReadWrite03() {

	funcName := "MainFileReadWriteTest010.FileBuffReadWrite03()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function:\n"+
		"    %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var shouldReadAndWriteFilesBeEqual,
		useWindowsLineTerminationChars,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	useWindowsLineTerminationChars = true

	shouldFinalDeleteWriteFile = true

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\FileBuffReadWrite03.txt",
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fHelper := new(strmech.FileHelper)

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var newFBuffReadWrite strmech.FileBufferReadWrite
	var readerFileInfoPlus, writerFileInfoPlus strmech.FileInfoPlus

	readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err = new(strmech.FileBufferReadWrite).
		NewPathFileNames(
			targetReadFile,
			false, // openReadFileReadWrite,
			512,   // readerBuffSize
			targetWriteFile,
			false, //openWriteFileReadWrite
			512,   // writerBuffSize
			true,  // truncateExistingWriteFile
			ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\n %v\n"+
		"%v\n"+
		" New File Stats\n"+
		" FileBufferReadWrite.NewPathFileNames()\n"+
		" Reader File Size= %v\n"+
		" Writer File Size= %v\n"+
		"  Target Read File: %v\n"+
		" Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size(),
		targetReadFile,
		targetWriteFile)

	var numOfBytesWritten, numOfTextLineBytes int64
	var numOfLinesProcessed int
	var readEndOfLineDelimiters strmech.StringArrayDto
	var writeEndOfLineChars string

	if useWindowsLineTerminationChars {

		// Windows Output Format
		writeEndOfLineChars = "\r\n"

	} else {

		// Unix Output Format
		writeEndOfLineChars = "\n"
	}

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
		"[EOL]")

	maxNumOfTextLines := -1

	numOfLinesProcessed,
		numOfTextLineBytes,
		numOfBytesWritten,
		err = newFBuffReadWrite.
		ReadWriteTextLines(
			&readEndOfLineDelimiters,
			writeEndOfLineChars,
			maxNumOfTextLines, // numTextLinesPerBatch
			true,              // autoFlushAndCloseOnExit
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" %v\n"+
		"%v"+
		" Stats From newFBuffReadWrite.ReadWriteTextLines()\n"+
		" Number Of Text Lines Processed: %v\n"+
		"   Maximum Number of Text Lines: %v\n"+
		"      Number Of Text Line Bytes: %v\n"+
		"        Number Of Bytes Written: %v\n"+
		"  Target Read File: %v\n"+
		" Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		numOfLinesProcessed,
		maxNumOfTextLines,
		numOfTextLineBytes,
		numOfBytesWritten,
		targetReadFile,
		targetWriteFile)

	var doesFileExist bool

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" %v\n"+
		"%v\n"+
		" Write File Stats from DoesFileInfoPlusExist()\n"+
		"    Write File Does Exist: %v\n"+
		" Write File Size in Bytes: %v\n"+
		"          Write File Name: %v\n"+
		"	 	Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		doesFileExist,
		writerFileInfoPlus.Size(),
		writerFileInfoPlus.Name(),
		targetWriteFile)

	err = newFBuffReadWrite.IsValidInstanceError(
		ePrefix)

	if err == nil {

		fmt.Printf("%v\n"+
			"newFBuffReadWrite Status Error\n"+
			"It was expected that 'newFBuffReadWrite' would\n"+
			"be closed and invalid. However, after being\n"+
			"closed, 'newFBuffReadWrite' is showing as valid.\n"+
			"This is an error condition!\n",
			ePrefix)

		return
	}

	if shouldReadAndWriteFilesBeEqual == true {

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

			fmt.Printf(" %v\n"+
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

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}
	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}
	}

	// ------ Trailing Marquee

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// FileBuffReadWrite04
//
// Testing the seek method
func (fileReadWriteTest010 MainFileReadWriteTest010) FileBuffReadWrite04() {

	funcName := "MainFileReadWriteTest010.FileBuffReadWrite04()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function:\n"+
		"    %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\Main10FileBuffReadWrite04.txt",
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fHelper := new(strmech.FileHelper)

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fBufReader strmech.FileBufferReader
	var fBufWriter strmech.FileBufferWriter

	_,
		fBufReader,
		err = new(strmech.FileBufferReader).
		NewPathFileName(
			targetReadFile,
			false, // Open File Read/Write
			512,   // bufSize
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	_,
		fBufWriter,
		err = new(strmech.FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false,
			512,
			false,
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetOffset, offsetFromFileStart int64

	targetOffset = 122

	offsetFromFileStart,
		err = fBufReader.Seek(
		targetOffset,
		io.SeekStart)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if offsetFromFileStart != targetOffset {

		fmt.Printf("%v\n"+
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
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	numOfBytesWritten,
		err = fBufWriter.WriteTextOrNumbers(
		&strBuilder,
		"",   // writeEndOfLineChars
		"",   // writeEndOfTextChars
		true, // autoFlushAndCloseOnExit
		ePrefix.XCpy("fBufWriter"))

	if numOfBytesRead != numOfBytesWritten {
		fmt.Printf(" %v\n"+
			"%v\n"+
			" Error: numOfBytesRead != numOfBytesWritten\n"+
			"    numOfBytesRead = %v\n"+
			" numOfBytesWritten = %v\n",
			ePrefix.String(),
			dashLineStr,
			numOfBytesRead,
			numOfBytesWritten)

		return
	}

	fmt.Printf("%v\n"+
		"%v\n"+
		"Number of Bytes Read matches\n"+
		"Number of bytes written.\n"+
		"   Number of Bytes Read = %v\n"+
		"Number of Bytes Written = %v\n\n",
		ePrefix.String(),
		dashLineStr,
		numOfBytesRead,
		numOfBytesWritten)

	// ------ Trailing Marquee

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (fileReadWriteTest010 MainFileReadWriteTest010) FileBufWriter01() {

	funcName := "Main010.FileBufWriter01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var shouldReadAndWriteFilesBeEqual,
		useWindowsLineTerminationChars,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	useWindowsLineTerminationChars = true

	shouldFinalDeleteWriteFile = true

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010FileBufWriter01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)
	var outputLinesArray strmech.StringArrayDto
	var numOfLinesRead, expectedNumOfBytesWritten int
	var i64numOfBytesRead int64
	var readEndOfLineDelimiters strmech.StringArrayDto
	var writeEndOfLineChars string

	if useWindowsLineTerminationChars {

		// Windows Output Format
		writeEndOfLineChars = "\r\n"

	} else {

		// Unix Output Format
		writeEndOfLineChars = "\n"

	}

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
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
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	outputLinesArray.AppendSuffix(writeEndOfLineChars)

	expectedNumOfBytesWritten =
		int(i64numOfBytesRead) + numOfLinesRead

	var fBufWriter strmech.FileBufferWriter

	_,
		fBufWriter,
		err = new(strmech.FileBufferWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			512,
			true,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var totalNumOfBytesWritten, localNumOfBytesWritten int
	var err2 error
	var bytesToWrite []byte

	for i := 0; i < numOfLinesRead; i++ {

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err2 =
			fBufWriter.Write(bytesToWrite)

		if err2 != nil {

			err = fmt.Errorf(" %v\n"+
				" Error returned by fBufWriter.Write(bytesToWrite)\n"+
				" Bytes To Write = '%v'\n"+
				" Index = '%v'\n"+
				" Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err2.Error())

			err2 = fBufWriter.FlushAndClose(nil)

			err = errors.Join(err, err2)

			fmt.Printf("%v\n",
				err.Error())

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten
	}

	var err3 error

	err3 = fBufWriter.Flush(ePrefix)

	if err3 != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Flush(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err3.Error())

		err = errors.Join(err, err2)
	}

	err3 = fBufWriter.FlushAndClose(ePrefix)

	if err3 != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Close(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err3.Error())

		err = errors.Join(err, err2)
	}

	if err != nil {

		fmt.Printf("%v\n"+
			"Errors returned from Flush() and Close()\n"+
			"Errors= \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if expectedNumOfBytesWritten != totalNumOfBytesWritten {

		fmt.Printf(" %v\n"+
			"%v\n"+
			" Error: Expected Bytes Written != Actual Bytes Written\n"+
			" Expected Bytes Written = '%v'\n"+
			"   Actual Bytes Written = '%v'\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			expectedNumOfBytesWritten,
			totalNumOfBytesWritten,
			targetReadFile,
			targetWriteFile)

		return
	}

	fmt.Printf(" %v\n"+
		"%v\n"+
		" After fBufWriter.Write() Sequence\n"+
		" Expected Number Of Bytes Written: %v\n"+
		"   Actual Number of Bytes Written: %v\n"+
		"  Target Read File: %v\n"+
		" Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		expectedNumOfBytesWritten,
		totalNumOfBytesWritten,
		targetReadFile,
		targetWriteFile)

	var filesAreEqual bool

	if shouldReadAndWriteFilesBeEqual == true {

		var reasonFilesNotEqual string

		filesAreEqual,
			reasonFilesNotEqual,
			err = fHelper.CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Target Files Comparison"))

		if err != nil {

			fmt.Printf(" %v\n"+
				" Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				" Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}

	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (fileReadWriteTest010 MainFileReadWriteTest010) FileBufWriter02() {

	funcName := "Main010.FileBufWriter02()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var shouldReadAndWriteFilesBeEqual,
		useWindowsLineTerminationChars,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	useWindowsLineTerminationChars = true

	shouldFinalDeleteWriteFile = true

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010FileBufWriter02.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)
	var outputLinesArray strmech.StringArrayDto
	var numOfLinesRead, expectedNumOfBytesWritten int
	var i64numOfBytesRead int64
	var readEndOfLineDelimiters strmech.StringArrayDto
	var writeEndOfLineChars string

	if useWindowsLineTerminationChars {

		// Windows Output Format
		writeEndOfLineChars = "\r\n"

	} else {

		// Unix Output Format
		writeEndOfLineChars = "\n"

	}

	readEndOfLineDelimiters.AddManyStrings(
		"\n",
		"\r\n",
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
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	outputLinesArray.AppendSuffix(writeEndOfLineChars)

	expectedNumOfBytesWritten =
		int(i64numOfBytesRead) + numOfLinesRead

	var filePtr *os.File
	var err2 error

	filePtr,
		err2 = fHelper.OpenFileWriteOnly(
		targetWriteFile,
		true,
		ePrefix)

	if err2 != nil {

		fmt.Printf("%v\n"+
			"Error: fHelper.OpenFileWriteOnly()\n"+
			"Atempty to open targetWriteFile Failed!\n"+
			"targetWriteFile= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err2.Error())

		return
	}

	var fBufWriter strmech.FileBufferWriter

	fBufWriter,
		err = new(strmech.FileBufferWriter).
		NewIoWriter(
			filePtr,
			512,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var totalNumOfBytesWritten, localNumOfBytesWritten int
	var bytesToWrite []byte

	expectedNumOfBytesWritten = 0

	for i := 0; i < numOfLinesRead; i++ {

		bytesToWrite = []byte(outputLinesArray.StrArray[i])

		expectedNumOfBytesWritten += len(bytesToWrite)

		localNumOfBytesWritten,
			err2 =
			fBufWriter.Write(bytesToWrite)

		if err2 != nil {

			err = fmt.Errorf(" %v\n"+
				" Error returned by fBufWriter.Write(bytesToWrite)\n"+
				" Bytes To Write = '%v'\n"+
				" Index = '%v'\n"+
				" Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err2.Error())

			err2 = fBufWriter.FlushAndClose(nil)

			err = errors.Join(err, err2)

			fmt.Printf("%v\n",
				err.Error())

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten
	}

	var err3 error

	err3 = fBufWriter.Close()

	if err3 != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Close()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err3.Error())

		err = errors.Join(err, err2)
	}

	filePtr = nil

	if err != nil {

		fmt.Printf("%v\n"+
			"Errors returned from Close()\n"+
			"which were executed on the\n"+
			"targetWriteFile.\n"+
			"targetWriteFile= '%v'\n"+
			"Errors= \n%v\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	if expectedNumOfBytesWritten != totalNumOfBytesWritten {

		fmt.Printf(" %v\n"+
			"%v\n"+
			" Error: Expected Bytes Written != Actual Bytes Written\n"+
			" Expected Bytes Written = '%v'\n"+
			"   Actual Bytes Written = '%v'\n"+
			"  Target Read File: %v\n"+
			" Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			expectedNumOfBytesWritten,
			totalNumOfBytesWritten,
			targetReadFile,
			targetWriteFile)

		return
	}

	fmt.Printf(" %v\n"+
		"%v\n"+
		" After fBufWriter.Write() Sequence\n"+
		" Expected Number Of Bytes Written: %v\n"+
		"   Actual Number of Bytes Written: %v\n"+
		"  Target Read File: %v\n"+
		" Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		expectedNumOfBytesWritten,
		totalNumOfBytesWritten,
		targetReadFile,
		targetWriteFile)

	var filesAreEqual bool

	if shouldReadAndWriteFilesBeEqual == true {

		var reasonFilesNotEqual string

		filesAreEqual,
			reasonFilesNotEqual,
			err = fHelper.CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Target Files Comparison"))

		if err != nil {

			fmt.Printf(" %v\n"+
				" Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				" Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}

	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

// WriteFileBytes01
//
// Runs test on FileHelper.WriteFileBytes()
func (fileReadWriteTest010 MainFileReadWriteTest010) WriteFileBytes01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileReadWriteTest010.WriteFileBytes01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var strMechOpsBaseDir string

	strMechOpsBaseDir,
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	targetOutputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput_2.txt"

	targetOutputFileName = strings.Replace(
		targetOutputFileName,
		"\\",
		osPathSepStr,
		-1)

	fhHelper := new(strmech.FileHelper)

	var numBytesRead int64
	var bytesRead []byte

	bytesRead,
		numBytesRead,
		err = fhHelper.ReadFileBytes(
		targetInputFileName,
		ePrefix.XCpy("targetInputFileName->"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Number of Bytes Read: %v\n\n",
		numBytesRead)

	fmt.Printf("Length of 'bytesRead' Array: %v\n\n",
		len(bytesRead))

	fmt.Printf("%v",
		string(bytesRead))

	var bytesWritten int64

	bytesWritten,
		err = fhHelper.
		WriteFileBytes(
			targetOutputFileName,
			false,
			bytesRead,
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Wrote Bytes to %v\n",
		targetInputFileName)

	fmt.Printf("Bytes Written: %v\n",
		bytesWritten)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// WriteFileBytes02
//
// This method tests FileMgr.WriteBytesToFile()
func (fileReadWriteTest010 MainFileReadWriteTest010) WriteFileBytes02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileReadWriteTest010.WriteFileBytes02()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var err error
	var strMechOpsBaseDir string

	strMechOpsBaseDir,
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	var inputFileMgr strmech.FileMgr

	inputFileMgr,
		err = new(strmech.FileMgr).New(
		targetInputFileName,
		ePrefix.XCpy("targetInputFileName"))

	if err != nil {

		_ = inputFileMgr.CloseThisFile(
			ePrefix.XCpy("inputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fileBytes []byte

	fileBytes,
		err = inputFileMgr.ReadAllFileBytes(
		ePrefix.XCpy("Read targetInputFileName"))

	if err != nil {

		_ = inputFileMgr.CloseThisFile(
			ePrefix.XCpy("inputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = inputFileMgr.CloseThisFile(
		ePrefix.XCpy("targetInputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	targetOutputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput_2.txt"

	targetOutputFileName = strings.Replace(
		targetOutputFileName,
		"\\",
		osPathSepStr,
		-1)

	var outputFileMgr strmech.FileMgr

	outputFileMgr,
		err = new(strmech.FileMgr).New(
		targetOutputFileName,
		ePrefix.XCpy("targetOutputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.DeleteThisFile(
		ePrefix.XCpy("targetOutputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesWritten int

	numOfBytesWritten,
		err = outputFileMgr.WriteBytesToFile(
		fileBytes,
		true,
		ePrefix.XCpy("<-outputFileMgr"))

	if err != nil {

		_ = outputFileMgr.CloseThisFile(
			ePrefix.XCpy("outputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.CloseThisFile(
		ePrefix.XCpy("outputFileMgr"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nOutput File:\n"+
		"%v\n"+
		"Cycle 1 Number Of Bytes Written: %v\n\n",
		outputFileMgr.GetAbsolutePathFileName(),
		numOfBytesWritten)

	fileBytes = make([]byte, 0)

	fileBytes = []byte("\nHello World!!\n")

	numOfBytesWritten,
		err = outputFileMgr.WriteBytesToFile(
		fileBytes,
		false,
		ePrefix.XCpy("<-outputFileMgr"))

	if err != nil {

		_ = outputFileMgr.CloseThisFile(
			ePrefix.XCpy("outputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.CloseThisFile(
		ePrefix.XCpy("outputFileMgr"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nOutput File:\n"+
		"%v\n"+
		"Cycle 2 Number Of Bytes Written: %v\n\n",
		outputFileMgr.GetAbsolutePathFileName(),
		numOfBytesWritten)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (fileReadWriteTest010 MainFileReadWriteTest010) IoReadWrite01() {

	funcName := "Main010.IoReadWrite01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n\n",
		ePrefix.String())

	var shouldReadAndWriteFilesBeEqual,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	shouldFinalDeleteWriteFile = true

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010IoReadwrite01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var i64NumOfBytesRead, i64NumOfBytesWritten int64

	var readFileInfoPlus,
		writeFileInfoPlus strmech.FileInfoPlus

	var targetIoReader strmech.FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(strmech.FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			4096,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var readStrBuilder = new(strings.Builder)

	i64NumOfBytesRead,
		err = targetIoReader.
		ReadAllToStrBuilder(
			readStrBuilder,
			true,
			ePrefix.XCpy("readStrBuilder<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoWriter strmech.FileIoWriter

	writeFileInfoPlus,
		targetIoWriter,
		err = new(strmech.FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	i64NumOfBytesWritten,
		err = targetIoWriter.
		WriteTextOrNumbers(
			readStrBuilder,
			"",
			"",
			true, // autoCloseOnExit
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if i64NumOfBytesRead != i64NumOfBytesWritten {

		fmt.Printf(" %v\n"+
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

	fmt.Printf(" %v\n"+
		"%v\n"+
		" After targetIoWriter.WriteTextOrNumbers() Sequence\n"+
		" Expected Number Of Bytes Written: %v\n"+
		"   Actual Number of Bytes Written: %v\n"+
		"  Target Read File: %v\n"+
		" Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		i64NumOfBytesRead,
		i64NumOfBytesWritten,
		targetReadFile,
		targetWriteFile)

	var fHelper = new(strmech.FileHelper)

	writeFileInfoPlus,
		err = fHelper.
		GetFileInfoPlus(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if readFileInfoPlus.Size() != writeFileInfoPlus.Size() {

		fmt.Printf("%v\n"+
			"Error: Target Read File Size != Target Write File Size!\n"+
			" Target Read File Size= %v\n"+
			"Target Write File Size= %v\n"+
			" Target Read File= %v\n"+
			"Target Write File= %v\n\n",
			ePrefix.String(),
			readFileInfoPlus.Size(),
			writeFileInfoPlus.Size(),
			targetReadFile,
			targetWriteFile)

		return
	}

	fmt.Printf("%v\n"+
		"%v\n"+
		"!!! Success !!!\n"+
		"Target Read File Size == Target Write File Size!\n"+
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

	var filesAreEqual bool

	if shouldReadAndWriteFilesBeEqual == true {

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

			fmt.Printf(" %v\n"+
				" Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				" Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}

	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}

		fmt.Printf("%v\n"+
			"%v\n"+
			"Successfully Deleted Target Write File.\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile)

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (fileReadWriteTest010 MainFileReadWriteTest010) IoWriteTo01() {

	funcName := "Main010.IoWriteTo01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n\n",
		ePrefix.String())

	var shouldReadAndWriteFilesBeEqual,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	shouldFinalDeleteWriteFile = true

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010IoWriteTo01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var readFileInfoPlus strmech.FileInfoPlus

	var targetIoReader strmech.FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(strmech.FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			1024,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoWriter strmech.FileIoWriter

	_,
		targetIoWriter,
		err = new(strmech.FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoReader.
		WriteTo(
			targetIoWriter)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		fmt.Printf("%v\n"+
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

	fmt.Printf("%v\n"+
		"%v\n"+
		"!!! SUCCESS !!!\n"+
		"targetIoReader.WriteTo()\n"+
		"The Number of Bytes Processed is EQUAL\n"+
		"to the size of the Target Read File.\n"+
		"Number of Bytes Processed= '%v'\n"+
		"    Target Readfile Size = '%v'\n"+
		" Target Read File: %v\n"+
		"Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		numOfBytesProcessed,
		readFileInfoPlus.Size(),
		targetReadFile,
		targetWriteFile)

	err = targetIoReader.Close()

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error closing 'targetIoReader'\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n\n",
			ePrefix.String(),
			targetReadFile,
			err.Error())

		return
	}

	err = targetIoWriter.Close()

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error closing 'targetIoWriter'\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	var filesAreEqual bool

	var fHelper = new(strmech.FileHelper)

	if shouldReadAndWriteFilesBeEqual == true {

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

			fmt.Printf(" %v\n"+
				" Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				" Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}

	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}

		fmt.Printf("%v\n"+
			"%v\n"+
			"Successfully Deleted Target Write File.\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile)

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (fileReadWriteTest010 MainFileReadWriteTest010) IoReadFrom01() {

	funcName := "Main010.IoReadFrom01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n\n",
		ePrefix.String())

	var shouldReadAndWriteFilesBeEqual,
		shouldFinalDeleteWriteFile bool

	shouldReadAndWriteFilesBeEqual = true

	shouldFinalDeleteWriteFile = true

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010IoWriteTo01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var readFileInfoPlus strmech.FileInfoPlus

	var targetIoReader strmech.FileIoReader

	readFileInfoPlus,
		targetIoReader,
		err = new(strmech.FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			4096,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoWriter strmech.FileIoWriter

	_,
		targetIoWriter,
		err = new(strmech.FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoWriter.
		ReadFrom(
			targetIoReader)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		fmt.Printf("%v\n"+
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

	fmt.Printf("%v\n"+
		"%v\n"+
		"!!! SUCCESS !!!\n"+
		"targetIoWriter.ReadFrom()\n"+
		"The Number of Bytes Processed is EQUAL\n"+
		"to the size of the Target Read File.\n"+
		"Number of Bytes Processed= '%v'\n"+
		"    Target Readfile Size = '%v'\n"+
		" Target Read File: %v\n"+
		"Target Write File: %v\n\n",
		ePrefix.String(),
		dashLineStr,
		numOfBytesProcessed,
		readFileInfoPlus.Size(),
		targetReadFile,
		targetWriteFile)

	err = targetIoReader.Close()

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error closing 'targetIoReader'\n"+
			"Target Read File: %v\n"+
			"Error=\n%v\n\n",
			ePrefix.String(),
			targetReadFile,
			err.Error())

		return
	}

	err = targetIoWriter.Close()

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error closing 'targetIoWriter'\n"+
			"Target Write File: %v\n"+
			"Error=\n%v\n\n",
			ePrefix.String(),
			targetWriteFile,
			err.Error())

		return
	}

	var filesAreEqual bool

	var fHelper = new(strmech.FileHelper)

	if shouldReadAndWriteFilesBeEqual == true {

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

			fmt.Printf(" %v\n"+
				" Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				" Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf(" %v\n"+
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

		} else {

			fmt.Printf(" %v\n"+
				"%v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				dashLineStr,
				targetReadFile,
				targetWriteFile)

		}

	}

	if shouldFinalDeleteWriteFile == true {

		err = fHelper.
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}

		fmt.Printf("%v\n"+
			"%v\n"+
			"Successfully Deleted Target Write File.\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile)

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

// IoWriteSeek01
//
// Tests FileIoWriter.Seek()
func (fileReadWriteTest010 MainFileReadWriteTest010) IoWriteSeek01() {

	funcName := "Main010.IoWriteSeek01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	dashLineStr := " " + strings.Repeat("-",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n\n",
		ePrefix.String())

	var shouldFinalDeleteWriteFile bool

	shouldFinalDeleteWriteFile = false

	var targetReadFile string
	var err error

	var exUtil = ExampleUtility{}

	targetReadFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFile.txt",
		ePrefix.XCpy("targetInputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exUtil.GetCompositeDirectory(
		"\\fileOpsTest\\trashDirectory\\Main010IoWriteSeek01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoReader strmech.FileIoReader
	var readFileInfoPlus strmech.FileInfoPlus

	readFileInfoPlus,
		targetIoReader,
		err = new(strmech.FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, // openFileReadWrite
			1024,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetIoWriter strmech.FileIoWriter

	_,
		targetIoWriter,
		err = new(strmech.FileIoWriter).
		NewPathFileName(
			targetWriteFile,
			false, // openFileReadWrite
			4096,
			true,
			ePrefix.XCpy("targetIoWriter<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesProcessed int64

	numOfBytesProcessed,
		err = targetIoReader.
		WriteTo(
			targetIoWriter)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if numOfBytesProcessed != readFileInfoPlus.Size() {

		fmt.Printf("%v\n"+
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
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if i64RequestedWriteFileOffset !=
		i64ActualWriteFileOffset {

		fmt.Printf("%v\n"+
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
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if localNumOfBytesWritten != lenTestStr {

		fmt.Printf("%v\n"+
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

	if shouldFinalDeleteWriteFile == true {

		err = new(strmech.FileHelper).
			DeleteDirOrFile(
				targetWriteFile,
				ePrefix.XCpy("Final Delete-targetWriteFile"))

		if err != nil {
			fmt.Printf("\n%v\n\n",
				err.Error())
			return
		}

		fmt.Printf("%v\n"+
			"%v\n"+
			"Successfully Deleted Target Write File.\n"+
			"Target Write File= '%v'\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile)

	}

	var err2 error

	err2 = targetIoReader.Close()

	err = errors.Join(err, err2)

	err2 = targetIoWriter.Close()

	err = errors.Join(err, err2)

	if err != nil {

		fmt.Printf("%v\n"+
			"%v\n"+
			"Error returned by targetIoReader.Close()\n"+
			"and/or targetIoWriter.Close().\n"+
			"All Errors= \n%v\n",
			ePrefix.String(),
			dashLineStr,
			err.Error())

		fmt.Printf("\n\n" + breakStr + "\n")

		fmt.Printf("\nERRORS - Function Execution FAILED!\n"+
			" Function: %v\n",
			ePrefix.String())

		return
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf(" !!! Successful Completion !!!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}
