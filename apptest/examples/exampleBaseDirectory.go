package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"strings"
)

// GetBaseDirectory
//
// Used to search for and return the base directory
// path to the 'strmechops' project.
func GetBaseDirectory(
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
