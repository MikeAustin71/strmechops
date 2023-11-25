package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestFileIoReader_Read_000100(t *testing.T) {

	funcName := "TestFileIoReader_Read_000100()"

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

	var fIoReader *FileIoReader
	var fInfoPlus FileInfoPlus

	fInfoPlus,
		fIoReader,
		err = new(FileIoReader).
		NewPathFileName(
			targetReadFile,
			false, //openFileReadWrite
			1024,
			ePrefix.XCpy("fIoReader<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if fInfoPlus.Size() != 1228 {

		t.Errorf("\n%v\n"+
			"Error: Test file should contain 1228-bytes.\n"+
			"Instead test file contains %v-bytes.\n"+
			"Target Read File= %v\n",
			ePrefix.String(),
			fInfoPlus.Size(),
			targetReadFile)

		return
	}

	var strBuilder = new(strings.Builder)
	var numOfBytesRead int64

	numOfBytesRead,
		err = fIoReader.
		ReadAllToStrBuilder(
			strBuilder,
			false,
			ePrefix.XCpy("strBuilder"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = fIoReader.Close()

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numOfBytesRead != 1228 {

		t.Errorf("\n%v\n"+
			"Error: Failed to read correct number\n"+
			"of bytes from the target read file.\n"+
			"The expected number of bytes is 1228-bytes.\n"+
			"The actual number of bytes read is %v-bytes.\n"+
			"Target Read File= %v\n",
			ePrefix.String(),
			numOfBytesRead,
			targetReadFile)

		return
	}

	if strBuilder.Len() != 1228 {

		t.Errorf("\n%v\n"+
			"Error: String Builder contains the wrong number\n"+
			"of bytes.\n"+
			"The expected length of String Builder is 1226.\n"+
			"The actual length of String Builder is %v.\n"+
			"Target Read File= %v\n",
			ePrefix.String(),
			strBuilder.Len(),
			targetReadFile)

		return
	}

	return
}

func TestFileIoReader_Seek_000200(t *testing.T) {

	funcName := "TestFileIoReader_Seek_000200()"

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
			"\\fileOpsTest\\filesForTest\\textFilesForTest\\smallTextFilePara2Only.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			"\\fileOpsTest\\trashDirectory\\TestFileIoReader_Seek_000200.txt",
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fHelper = new(FileHelper)

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
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
			2048,
			ePrefix.XCpy("targetIoReader<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var targetIoWriter *FileIoWriter

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

	var actualReaderOffset, targetOffset int64

	targetOffset = 122

	actualReaderOffset,
		err = targetIoReader.Seek(
		targetOffset,
		io.SeekStart)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())

		_ = targetIoReader.Close()

		_ = targetIoWriter.Close()

		return
	}

	if actualReaderOffset != targetOffset {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: targetIoReader.Seek()\n"+
			"The requested offset is NOT equal\n"+
			"to the actual offset."+
			"Requested Offset: %v\n"+
			"   Actual Offset: %v\n"+
			"Target Read File: %v\n",
			ePrefix.String(),
			dashLineStr,
			targetOffset,
			actualReaderOffset,
			targetReadFile)

		_ = targetIoReader.Close()

		_ = targetIoWriter.Close()

		return
	}

	var readContentStr string

	var numOfReaderBytesRead int64

	numOfReaderBytesRead,
		readContentStr,
		err = targetIoReader.ReadAllToString(
		true, // autoCloseOnExit
		ePrefix.XCpy("readContentStr<-"))

	if err != nil {

		t.Errorf("\n%v\n",
			err.Error())

		_ = targetIoReader.Close()

		_ = targetIoWriter.Close()

		return
	}

	var intNumOfBytesWritten int

	var i64NumOfBytesWritten int64

	intNumOfBytesWritten,
		err = targetIoWriter.Write(
		[]byte(readContentStr))

	if err != nil {

		t.Errorf("\n%v\n",
			err.Error())

		_ = targetIoReader.Close()

		_ = targetIoWriter.Close()

		return
	}

	defer func() {

		_ = fHelper.DeleteDirOrFile(
			targetWriteFile,
			nil)

	}()

	var err2, err3 error

	err3 = targetIoReader.Close()

	err2 = errors.Join(err2, err3)

	err3 = targetIoWriter.Close()

	err2 = errors.Join(err2, err3)

	if err2 != nil {

		t.Errorf("%v\n"+
			"%v\n"+
			"Errors occurred while closing the"+
			"targetIoReader and targetIoWriter!\n"+
			"Errors=\n%v\n",
			ePrefix.String(),
			dashLineStr,
			err2.Error())

		return
	}

	i64NumOfBytesWritten = int64(intNumOfBytesWritten)

	if i64NumOfBytesWritten != numOfReaderBytesRead {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The number of bytes read is NOT\n"+
			"EQUAL to the number of bytes written!\n"+
			"   Number of Bytes Read= '%v'\n"+
			"Number of Bytes Written= '%v'\n"+
			" Target Read File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			numOfReaderBytesRead,
			i64NumOfBytesWritten,
			targetReadFile,
			targetWriteFile)

		return
	}

	var writeFileInfoPlus FileInfoPlus

	doesFileExist,
		writeFileInfoPlus,
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
			"%v\n"+
			"Error: After writing data to the Target Write File and\n"+
			"closing the Target Write File, the Target Write File\n"+
			"was not found and Does NOT exist on disk!\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			targetWriteFile)

		return

	}

	var i64ExpectedWriteFileSize int64

	i64ExpectedWriteFileSize =
		readFileInfoPlus.Size() -
			targetOffset

	if i64ExpectedWriteFileSize != writeFileInfoPlus.Size() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The expected size of the Target Write\n"+
			"File DOES NOT MATCH the actual size of the\n"+
			"Target Write File!\n"+
			"Expected Target Write File Size= %v-bytes\n"+
			" Actual Target Write File Size=  %v-bytes\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			i64ExpectedWriteFileSize,
			writeFileInfoPlus.Size(),
			targetWriteFile)

		return
	}

	var filesAreEqual bool
	var reasonFilesNotEqual string

	filesAreEqual,
		reasonFilesNotEqual,
		err = fHelper.
		CompareFiles(
			compareFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Compare File vs Write File"))

	if err != nil {
		t.Errorf("\n%v\n\n",
			err.Error())
		return
	}

	if !filesAreEqual {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: The Comparison File and the Target\n"+
			"Write File are NOT EQUAL!\n"+
			"Reason Files NOT Equal: %v\n"+
			"  Comparison File: %v\n"+
			"Target Write File: %v\n",
			ePrefix.String(),
			dashLineStr,
			reasonFilesNotEqual,
			compareFile,
			targetWriteFile)

		return
	}

	return
}
