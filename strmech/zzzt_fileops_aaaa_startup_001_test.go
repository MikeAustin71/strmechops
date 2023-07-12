package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
)

const (
	//FILEOpsBaseTestLogDir   = "/fileOpsTest/logTest"
	FILEOpsBaseFilesForTest = "/fileOpsTest/filesForTest"
	//FILEOpsBaseTrashDirectory = "/fileOpsTest/trashDirectory"
	FILEOpsRelTestLogDir     = "../fileOpsTest/logTest"
	FILEOpsRelFilesForTest   = "../fileOpsTest/filesForTest"
	FILEOpsRelTrashDirectory = "../fileOpsTest/trashDirectory"
	//commonDir = "../../pathFileOps"
)

type fileOpsTestUtility struct {
	Input string
}

// GetCompositeDir
//
// Locates and returns the package base directory path to
// the 'strmechops' project.
func (fOpsTestUtil *fileOpsTestUtility) GetCompositeDir(
	localDir string,
	errorPrefix interface{}) (
	string,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileOpsTestUtility."+
			"GetBaseDirectory()",
		"")

	if err != nil {
		return "", err
	}

	var baseDir string

	baseDir,
		err = new(DirHelper).GetCurrentDir(
		ePrefix.XCpy("baseDir<-"))

	if err != nil {
		return baseDir, err
	}

	if len(baseDir) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Returned 'baseDir' is invalid!\n"+
			"'baseDir' is an empty string with zero string length.\n",
			ePrefix.String())

		return baseDir, err
	}

	baseDir = strings.ToLower(baseDir)

	targetDir := "strmechops"

	strmechopsIdx := strings.Index(baseDir, targetDir)

	if strmechopsIdx == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Returned 'baseDir' is invalid!\n"+
			"'baseDir' is an empty string with zero string length.\n",
			ePrefix.String())

		return baseDir, err

	}

	osPathSepStr := string(os.PathSeparator)

	endStrmechopsIdx := strmechopsIdx + len(targetDir)

	strMechOpsBaseDir := baseDir[0:endStrmechopsIdx]

	compositeDir :=
		strMechOpsBaseDir + localDir

	compositeDir = strings.Replace(
		compositeDir,
		"\\",
		osPathSepStr,
		-1)

	return compositeDir, err
}
