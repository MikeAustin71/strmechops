package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"os"
	"strings"
)

type MainFilePermissionsTest009 struct {
	input string
}

func (filePermissionsTest MainFilePermissionsTest009) PermissionStr01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFilePermissionsTest009.PermissionStr01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var targetReadFileName string
	var err error
	var exampleUtil = ExampleUtility{}

	targetReadFileName,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\badassPlane.txt",
		ePrefix.XCpy("readFileAddOn"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)

	var targetReadFilePtr *os.File

	targetReadFilePtr,
		err = fHelper.OpenFileComponents(
		ePrefix.XCpy("targetReadFilePtr"),
		targetReadFileName,
		false,
		"-r--r--r--",
		strmech.FOpenType.TypeReadOnly())

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fInfo os.FileInfo

	fInfo,
		err = targetReadFilePtr.Stat()

	if err != nil {

		_ = targetReadFilePtr.Close()

		fmt.Printf("%v\n"+
			"Error returned by targetReadFilePtr.Stat()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"File Mode: 0%o\n"+
		"File Permissions: %s\n"+
		"File: %v\n\n",
		ePrefix.String(),
		fInfo.Mode(),
		fInfo.Mode().Perm(),
		targetReadFileName)

	err = targetReadFilePtr.Close()

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error returned by targetReadFilePtr.Close()\n"+
			"Target Read File = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetReadFileName,
			err.Error())

		return
	}

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (filePermissionsTest MainFilePermissionsTest009) PermissionStr02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFilePermissionsTest009.PermissionStr02()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var targetReadFileName string
	var err error
	var exampleUtil = ExampleUtility{}

	targetReadFileName,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\badassPlane.txt",
		ePrefix.XCpy("readFileAddOn"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)

	var targetReadFilePtr *os.File
	var readOpenFilePermissions = "-r--r--r--"

	targetReadFilePtr,
		err = fHelper.OpenFileComponents(
		ePrefix.XCpy("targetReadFilePtr"),
		targetReadFileName,
		false,
		readOpenFilePermissions,
		strmech.FOpenType.TypeReadWrite())

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fInfo os.FileInfo

	fInfo,
		err = targetReadFilePtr.Stat()

	if err != nil {

		_ = targetReadFilePtr.Close()

		fmt.Printf("%v\n"+
			"Error returned by targetReadFilePtr.Stat()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"Target Read File: %v\n"+
		"Target Read File Mode: 0%o\n"+
		"Requested Target Read File Permissions: %s\n"+
		"   Actual Target Read File Permissions: %s\n"+
		"Target Read File is Open for reading.\n\n",
		ePrefix.String(),
		targetReadFileName,
		fInfo.Mode(),
		readOpenFilePermissions,
		fInfo.Mode().Perm())

	var targetWriteFileName string

	targetWriteFileName,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\testOutput.txt",
		ePrefix.XCpy("targetWriteFileName<-"))

	if err != nil {

		_ = targetReadFilePtr.Close()

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFilePtr *os.File
	var writeOpenFilePermissions = "--w--w--w-"

	targetWriteFilePtr,
		err = fHelper.OpenFileComponents(
		ePrefix.XCpy("targetWriteFilePtr<-targetWriteFileName"),
		targetWriteFileName,
		true,
		writeOpenFilePermissions,
		strmech.FOpenType.TypeWriteOnly(),
		strmech.FOpenMode.ModeCreate(),
		strmech.FOpenMode.ModeTruncate())

	if err != nil {

		_ = targetReadFilePtr.Close()

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var writeFInfo os.FileInfo

	writeFInfo,
		err = targetWriteFilePtr.Stat()

	if err != nil {

		_ = targetReadFilePtr.Close()

		_ = targetWriteFilePtr.Close()

		fmt.Printf("%v\n"+
			"Error returned by targetWriteFilePtr.Stat()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"Before Read/Write Operation\n"+
		"Target Write File: %v\n"+
		"Target Write File Mode: 0%o\n"+
		"Requested Target Write File Permissions: %s\n"+
		"   Actual Target Write File Permissions: %s\n"+
		"Target Write File is Open for writing.\n\n",
		ePrefix.String(),
		targetWriteFileName,
		writeFInfo.Mode(),
		writeOpenFilePermissions,
		fInfo.Mode().Perm())

	var readWriteBuffer = make([]byte, 256)

	var numOfBytesRead, numOfBytesWritten,
		totalNumOfBytesRead, totalNumOfBytesWritten, cycleCount int
	var readErr, writeErr error

	for {

		cycleCount++

		numOfBytesRead,
			readErr = targetReadFilePtr.Read(
			readWriteBuffer)

		if readErr != nil &&
			readErr != io.EOF {

			fmt.Printf("%v\n"+
				"Error Reading Target Read File!\n"+
				"Cycle Count= %v\n"+
				"Target Read File= %v\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				cycleCount,
				targetReadFileName,
				readErr.Error())

			_ = targetReadFilePtr.Close()

			_ = targetWriteFilePtr.Close()

			return
		}

		if numOfBytesRead > 0 {
			totalNumOfBytesRead += numOfBytesRead

			numOfBytesWritten,
				writeErr = targetWriteFilePtr.
				Write(
					readWriteBuffer[0:numOfBytesRead])

			if writeErr != nil {

				fmt.Printf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"Target Write File = '%v'\n"+
					"Write Error=\n%v\n",
					ePrefix.String(),
					targetWriteFileName,
					writeErr.Error())

				_ = targetReadFilePtr.Close()

				_ = targetWriteFilePtr.Close()

				return
			}

			if numOfBytesRead != numOfBytesWritten {

				fmt.Printf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"numOfBytesRead != numOfBytesWritten\n"+
					"Target Write File = '%v'\n"+
					"numOfBytesRead = %v\n"+
					"numOfBytesWritten = %v\n"+
					"Write Error=\n%v\n",
					ePrefix.String(),
					numOfBytesRead,
					numOfBytesWritten,
					targetWriteFileName,
					writeErr.Error())

				_ = targetReadFilePtr.Close()

				_ = targetWriteFilePtr.Close()

				return
			}

			totalNumOfBytesWritten += numOfBytesWritten

		}

		if readErr == io.EOF {

			break
		}

	}

	fmt.Printf("%v\n"+
		"Successful Read/Write Operation!\n"+
		"Target Read File: %v\n"+
		"Target Write File: %v\n"+
		"   Number of Bytes Read: %v\n"+
		"Number of Bytes Written: %v\n\n",
		ePrefix.String(),
		targetReadFileName,
		targetWriteFileName,
		totalNumOfBytesRead,
		totalNumOfBytesWritten)

	var errs []error
	var err2 error

	err2 = targetReadFilePtr.Close()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error Closing Target Read File!\n"+
			"targetReadFilePtr.Close()\n"+
			"Target Read File = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetReadFileName,
			err2.Error())

		errs = append(errs, err)
	}

	err2 = targetWriteFilePtr.Close()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error Closing Target Write File!\n"+
			"targetWriteFilePtr.Close()\n"+
			"Target Write File = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetWriteFileName,
			err2.Error())

		errs = append(errs, err)
	}

	if len(errs) > 0 {

		err = new(strmech.StrMech).ConsolidateErrors(errs)

		fmt.Printf("%v\n"+
			"Errors Returned while closing Read and Write Files!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	writeFInfo,
		err = fHelper.GetFileInfo(
		targetWriteFileName,
		ePrefix.XCpy("After Close-targetWriteFileName"))

	if err != nil {

		fmt.Printf("\n%v\n\n",
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"After Read/Write Operation And File is Close.\n"+
		"Target Write File: %v\n"+
		"Target Write File Mode: 0%o\n"+
		"Requested Target Write File Permissions: %s\n"+
		"   Actual Target Write File Permissions: %s\n\n",
		ePrefix.String(),
		targetWriteFileName,
		writeFInfo.Mode(),
		writeOpenFilePermissions,
		writeFInfo.Mode().Perm())

	var fileModeChangeStr = "-r--r--r--"

	err = fHelper.
		ChmodPermissionStr(
			targetWriteFileName,
			fileModeChangeStr,
			ePrefix.XCpy("targetWriteFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	writeFInfo,
		err = fHelper.GetFileInfo(
		targetWriteFileName,
		ePrefix.XCpy("After Close-targetWriteFileName"))

	if err != nil {

		fmt.Printf("\n%v\n\n",
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"After Read/Write Operation, Target Write\n"+
		"File is Closed and the call to Change File\n"+
		"Mode, this is the current permissions status.\n"+
		"Medhod Call: fHelper.ChmodPermissionStr()\n"+
		"Target Write File: %v\n"+
		"Target Write File Mode: 0%o\n"+
		"          Requested File Mode Change: %s\n"+
		"Actual Target Write File Permissions: %s\n\n",
		ePrefix.String(),
		targetWriteFileName,
		writeFInfo.Mode(),
		fileModeChangeStr,
		writeFInfo.Mode().Perm())

	errs = make([]error, 0)

	err,
		writeErr = fHelper.DeleteDirFile(
		targetWriteFileName,
		ePrefix.XCpy("targetWriteFileName"))

	if err != nil {

		errs = append(
			errs, err)

	}

	if writeErr != nil {

		errs = append(errs,
			fmt.Errorf("%v\n"+
				"Error From fHelper.DeleteDirFile(targetWriteFileName)\n"+
				"Target Write File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				targetWriteFileName,
				writeErr))

	}

	if len(errs) > 0 {

		err = new(strmech.StrMech).ConsolidateErrors(errs)

		fmt.Printf("%v\n"+
			"Errors Returned while Deleting Write File!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}
