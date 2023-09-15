package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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

	var fIoReader FileIoReader
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
