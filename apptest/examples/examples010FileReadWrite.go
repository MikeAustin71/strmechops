package examples

import (
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
		"MainFileReadWriteTest010.ReadLines02()",
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

	fmt.Printf("\nFile Info Data\n"+
		"newFBuffReadWrite\n"+
		"Reader File Size= %v\n"+
		"Writer File Size= %v\n\n",
		readerFileInfoPlus.Size(),
		writerFileInfoPlus.Size())

	var totalBytesRead, totalBytesWritten int

	totalBytesRead,
		totalBytesWritten,
		err = newFBuffReadWrite.ReadWriteAll(ePrefix.XCpy(
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

	// --------------------------------------------
	// Successful Completion
	// --------------------------------------------
	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (fileReadWriteTest010 MainFileReadWriteTest010) FileBufWriter01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileReadWriteTest010.FileBufWriter01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

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
		"\\fileOpsTest\\trashDirectory\\MainFileReadWriteTest010_FileBufWriter01.txt",
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)
	var textLinesArray strmech.StringArrayDto
	var numOfLinesRead, expectedNumOfBytesWritten int
	var i64numOfBytesRead int64

	_,
		numOfLinesRead,
		i64numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&textLinesArray,
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	textLinesArray.AppendSuffix("\n")

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

		bytesToWrite = []byte(textLinesArray.StrArray[i])

		localNumOfBytesWritten,
			err2 =
			fBufWriter.Write(bytesToWrite)

		if err2 != nil {

			fmt.Printf("%v\n"+
				"Error returned by fBufWriter.Write(bytesToWrite)\n"+
				"Bytes To Write = '%v'\n"+
				"Index = '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				string(bytesToWrite),
				i,
				err2.Error())

			_ = fBufWriter.Close(nil)

			return
		}

		totalNumOfBytesWritten += localNumOfBytesWritten
	}

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

	err2 = fBufWriter.Close(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.Close(ePrefix)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		errs = append(errs, err)
	}

	if len(errs) > 0 {

		err2 = new(strmech.StrMech).ConsolidateErrors(errs)

		fmt.Printf("%v\n"+
			"Errors returned from Flush() and Close()\n"+
			"Errors= \n%v\n",
			ePrefix.String(),
			err2.Error())

		return
	}

	if expectedNumOfBytesWritten != totalNumOfBytesWritten {

		fmt.Printf("%v\n"+
			"Error: Expected Bytes Written != Actual Bytes Written\n"+
			"Expected Bytes Written = '%v'\n"+
			"  Actual Bytes Written = '%v'\n",
			ePrefix.String(),
			expectedNumOfBytesWritten,
			totalNumOfBytesWritten)

		return
	}

	fmt.Printf("\n\n%v\n"+
		"Expected Number Of Bytes Written: %v\n"+
		"  Actual Number of Bytes Written: %v\n",
		ePrefix.String(),
		expectedNumOfBytesWritten,
		totalNumOfBytesWritten)

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
