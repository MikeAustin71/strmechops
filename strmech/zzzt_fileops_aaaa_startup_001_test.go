package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"runtime"
	"strings"
)

const (
	FILEOpsBaseFilesForTest      = "/fileOpsTest/filesForTest"
	FILEOpsBaseLevelFilesForTest = "/fileOpsTest/filesForTest/levelfilesfortest"
	FILEOpsBaseTrashDirectory    = "/fileOpsTest/trashDirectory"
	FILEOpsRelTestLogDir         = "../fileOpsTest/logTest"
	FILEOpsRelFilesForTest       = "../fileOpsTest/filesForTest"
	FILEOpsRelTrashDirectory     = "../fileOpsTest/trashDirectory"
)

type fileOpsTestUtility struct {
	Input string
}

func (fOpsTestUtil *fileOpsTestUtility) CopyTestFileToTrashDir(
	baseFileName string,
	targetFileName string,
	errorPrefix interface{}) (
	sourcePathFileName string,
	destinationPathFileName string,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileOpsTestUtility."+
			"CopyTestFileToTrashDir()",
		"")

	if err != nil {

		return sourcePathFileName,
			destinationPathFileName,
			err
	}

	if baseFileName[0] == '\\' ||
		baseFileName[0] == '/' {

		baseFileName = baseFileName[1:]
	}

	baseFile := FILEOpsBaseFilesForTest + "/" + baseFileName

	if targetFileName[0] == '\\' ||
		targetFileName[0] == '/' {

		targetFileName = targetFileName[1:]
	}

	targetFile := FILEOpsBaseTrashDirectory + "/" + targetFileName

	sourcePathFileName,
		err = new(fileOpsTestUtilityMolecule).
		makeCompositeDir(
			baseFile,
			ePrefix)

	if err != nil {

		return sourcePathFileName,
			destinationPathFileName,
			err
	}

	destinationPathFileName,
		err = new(fileOpsTestUtilityMolecule).
		makeCompositeDir(
			targetFile,
			ePrefix)

	if err != nil {

		return sourcePathFileName,
			destinationPathFileName,
			err
	}
	var baseReadFileBytes int64
	var fileExistsOnDisk bool
	var fHelper FileHelper

	fileExistsOnDisk,
		baseReadFileBytes,
		err = fHelper.GetBytesInFile(
		sourcePathFileName,
		ePrefix.XCpy(fmt.Sprintf(
			"sourcePathFileName= %v",
			sourcePathFileName)))

	if err != nil {

		return sourcePathFileName,
			destinationPathFileName,
			err
	}

	if !fileExistsOnDisk {

		err = fmt.Errorf("\n\n%v\n"+
			"Error Returned byfHelper.GetBytesInFile(sourcePathFileName)\n"+
			"Base File Name does NOT Exist on Disk!\n"+
			"Base File name= %v\n"+
			"Formatted Base File Name=\n"+
			"  %v",
			ePrefix.String(),
			baseFileName,
			sourcePathFileName)

		return sourcePathFileName,
			destinationPathFileName,
			err
	}

	var targetFileWriteBytes int64

	targetFileWriteBytes,
		err = new(FileHelper).CopyFileByIoBuffer(
		sourcePathFileName,
		destinationPathFileName,
		nil,
		false,
		ePrefix.XCpy("destinationPathFileName<-sourcePathFileName"))

	if err != nil {

		return sourcePathFileName,
			destinationPathFileName,
			err
	}

	if targetFileWriteBytes != baseReadFileBytes {

		err = fmt.Errorf("\n\n%v\n"+
			"Error Returned by %v"+
			"\nCopy Operation sourcePathFileName->destinationPathFileName FAILED!\n"+
			"\nExpected destinationPathFileName Size= %v bytes"+
			"\nActual destinationPathFileName Size  = %v\n"+
			"sourcePathFileName=%v\n"+
			"destinationPathFileName=%v\n",
			ePrefix.String(),
			"new(FileHelper).CopyFileByIoBuffer()",
			baseReadFileBytes,
			targetFileWriteBytes,
			sourcePathFileName,
			destinationPathFileName)
	}

	return sourcePathFileName,
		destinationPathFileName,
		err
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

	return new(fileOpsTestUtilityMolecule).
		makeCompositeDir(
			localDir,
			ePrefix)
}

// GetOpsSys
// returns "windows" or "linux"
func (fOpsTestUtil *fileOpsTestUtility) GetOpsSys() string {
	return runtime.GOOS
}

type fileOpsTestUtilityMolecule struct {
	InputStr string
}

// makeCompositeDir
// Creates a composite directory string from a
// subdirectory string for a directory located in
// the current source code directory tree.
func (fOpsTestUtilMol *fileOpsTestUtilityMolecule) makeCompositeDir(
	localDir string,
	errorPrefix interface{}) (
	string,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileOpsTestUtilityMolecule."+
			"makeCompositeDir()",
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

	if localDir[0] == '\\' ||
		localDir[0] == '/' {

		localDir = localDir[1:]
	}

	compositeDir :=
		strMechOpsBaseDir +
			osPathSepStr +
			localDir

	if runtime.GOOS == "windows" {

		// windows
		compositeDir = strings.Replace(
			compositeDir,
			"/",
			"\\",
			-1)

	} else {

		// Must be linux
		compositeDir = strings.Replace(
			compositeDir,
			"\\",
			"/",
			-1)
	}

	if osPathSepStr == "/" {
		// linux

	} else {

	}

	return compositeDir, err
}
