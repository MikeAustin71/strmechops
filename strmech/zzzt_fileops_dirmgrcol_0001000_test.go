package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestDirMgrCollection_GetTextListingAbsPath_000100(t *testing.T) {

	funcName := "TestDirMgrCollection_GetTextListingAbsPath_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	targetDir := FILEOPSFilesForTest
	//"../fileOpsTest/filesForTest"
	// Number of Directories: 4

	var numOfDirectoriesLocated int
	var isParentDirectoryIncluded bool
	var directoriesLocated DirMgrCollection
	var err error

	numOfDirectoriesLocated,
		isParentDirectoryIncluded,
		directoriesLocated,
		_, // numOfFilesLocated,
		_, // filesLocated
		err = new(DirHelper).
		GetSubDirsFilesInParentDir(
			targetDir,               // directoryPath
			true,                    // getSubdirectories
			false,                   // includeParentDirectory
			false,                   // includeSubDirCurrentDirOneDot
			false,                   // includeSubDirParentDirTwoDots
			false,                   // getRegularFiles
			false,                   // getSymLinksFiles
			false,                   // getOtherNonRegularFiles
			FileSelectionCriteria{}, // subDirSelectCharacteristics
			FileSelectionCriteria{}, // subDirSelectCharacteristics
			"targetDir",             // directoryPathLabel
			ePrefix)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error returned by:\n"+
			" DirHelper.GetSubDirsFilesInParentDir(targetDir).\n"+
			"targetDir='%v'\n"+
			"Error=\n%v\n",
			funcName,
			targetDir,
			err.Error())

		return
	}

	strBuilder := strings.Builder{}

	err = directoriesLocated.
		GetTextListingAbsPath(
			" ",
			" ",
			80,
			'-',
			targetDir,
			true,
			&strBuilder,
			ePrefix.XCpy("<-directoriesLocated"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error returned by:\n"+
			" directoriesLocated.GetTextListingAbsPath().\n"+
			"targetDir='%v'\n"+
			"Error=\n%v\n",
			funcName,
			targetDir,
			err.Error())

		return
	}

	outputStr := strBuilder.String()

	if !strings.Contains(
		outputStr,
		"Number of Directories: 9") {

		t.Errorf("\n%v\n"+
			"Error: The Text Listing did NOT contain:\n"+
			"'Number of Directories: 9'\n"+
			" directoriesLocated.GetTextListingAbsPath().\n"+
			"targetDir='%v'\n"+
			"outputStr=\n%v\n",
			ePrefix.String(),
			targetDir,
			outputStr)

		return

	}

	if numOfDirectoriesLocated != 9 {

		t.Errorf("\n%v\n"+
			"Error: Expected numOfDirectoriesLocated == 4\n"+
			"Instead, numOfDirectoriesLocated = '%v'\n"+
			"targetDir='%v'\n",
			ePrefix.String(),
			numOfDirectoriesLocated,
			targetDir)

		return

	}

	if isParentDirectoryIncluded == true {

		t.Errorf("\n%v\n"+
			"Error: Expected isParentDirectoryIncluded == 'true'\n"+
			"Instead, isParentDirectoryIncluded = 'false'\n"+
			"targetDir='%v'\n",
			ePrefix.String(),
			targetDir)

		return

	}

	return
}
