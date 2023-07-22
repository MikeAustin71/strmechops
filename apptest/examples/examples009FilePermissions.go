package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
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
