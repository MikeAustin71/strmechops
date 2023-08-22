package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"testing"
)

func TestFileBufferReader_Read_000100(t *testing.T) {

	funcName := "TestFileBufferReader_Read_000100()"

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
