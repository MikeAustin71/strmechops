package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"os"
	"strings"
)

type ExampleUtility struct {
	Input string
}

// GetBaseDirectory
//
// Used to search for and return the base directory
// path to the 'strmechops' project.
func (exUtil ExampleUtility) GetBaseDirectory(
	errorPrefix interface{}) (
	string,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirHelper."+
			"GetCurrentDir()",
		"")

	if err != nil {
		return "", err
	}

	var baseDir string

	baseDir,
		err = new(strmech.DirHelper).GetCurrentDir(
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

	/*
		fmt.Printf("Original Current Directory\n"+
			"%v\n\n",
			baseDir)

	*/

	targetDir := "strmechops"

	strmechopsIdx := strings.Index(baseDir, targetDir)

	if strmechopsIdx == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Returned 'baseDir' is invalid!\n"+
			"'baseDir' is an empty string with zero string length.\n",
			ePrefix.String())

		return baseDir, err

	}

	endStrmechopsIdx := strmechopsIdx + len(targetDir)

	strMechOpsBaseDir := baseDir[0:endStrmechopsIdx]

	return strMechOpsBaseDir, err
}

// GetCompositeDirectory
//
// Identifies the base directory for the 'strmechops' project
// and adds the requested subdirectory and/or file name to
// the base directory.
func (exUtil ExampleUtility) GetCompositeDirectory(
	dirFileAddOn string,
	errorPrefix interface{}) (
	string,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirHelper."+
			"GetCurrentDir()",
		"")

	if err != nil {
		return "", err
	}

	var baseDir string

	baseDir,
		err = new(strmech.DirHelper).GetCurrentDir(
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

	/*
		fmt.Printf("Original Current Directory\n"+
			"%v\n\n",
			baseDir)

	*/

	targetDir := "strmechops"

	strmechopsIdx := strings.Index(baseDir, targetDir)

	if strmechopsIdx == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Returned 'baseDir' is invalid!\n"+
			"'baseDir' is an empty string with zero string length.\n",
			ePrefix.String())

		return baseDir, err

	}

	endStrmechopsIdx := strmechopsIdx + len(targetDir)

	dirFileSpec := baseDir[0:endStrmechopsIdx]

	if len(dirFileAddOn) == 0 {

		return dirFileSpec, err
	}

	osPathSepStr := string(os.PathSeparator)

	if dirFileAddOn[0] == '\\' ||
		dirFileAddOn[0] == '/' {

		if len(dirFileAddOn) > 0 {
			dirFileAddOn = dirFileAddOn[1:]
		} else {

			dirFileAddOn = ""
		}

	}

	if len(dirFileAddOn) > 0 {
		dirFileSpec =
			dirFileSpec +
				osPathSepStr +
				dirFileAddOn
	}

	if osPathSepStr == "\\" {

		dirFileSpec = strings.Replace(
			dirFileSpec,
			"/",
			osPathSepStr,
			-1)

	} else {
		// MUST BE osPathSepStr == "/"

		dirFileSpec = strings.Replace(
			dirFileSpec,
			"\\",
			osPathSepStr,
			-1)

	}

	return dirFileSpec, err
}
