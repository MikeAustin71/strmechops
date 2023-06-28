package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"strings"
)

type MainDirOpsTest007 struct {
	input string
}

func (dirOpsTest007 MainDirOpsTest007) GetDirs01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.MainDirOpsTest007()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	targetDir := "D:\\T02\\WebSite_15"

	var numOfDirectoriesLocated int
	var isParentDirectoryIncluded bool
	var directoriesLocated strmech.DirMgrCollection
	var err error

	numOfDirectoriesLocated,
		isParentDirectoryIncluded,
		directoriesLocated,
		_, // numOfFilesLocated,
		_, // filesLocated
		err = new(strmech.DirHelper).
		GetSubDirsFilesInParentDir(
			targetDir,                       // directoryPath
			true,                            // getSubdirectories
			false,                           // includeParentDirectory
			false,                           // includeSubDirCurrentDirOneDot
			false,                           // includeSubDirParentDirTwoDots
			false,                           // getRegularFiles
			false,                           // getSymLinksFiles
			false,                           // getOtherNonRegularFiles
			strmech.FileSelectionCriteria{}, // subDirSelectCharacteristics
			strmech.FileSelectionCriteria{}, // subDirSelectCharacteristics
			"targetDir",                     // directoryPathLabel
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" Number Of Subdirectories: %v\n",
		numOfDirectoriesLocated)

	fmt.Printf(" Is Parent Directory Included: %v\n\n",
		isParentDirectoryIncluded)

	strBuilder := strings.Builder{}

	err = directoriesLocated.
		GetTextListingAbsPath(
			" ",
			"",
			80,
			'-',
			"Directory D:\\T02\\WebSite_15",
			true,
			&strBuilder,
			ePrefix.XCpy("<-directoriesLocated"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Println(strBuilder.String())

	/*
		output := strBuilder.String()

		printableChars := new(strmech.StrMech).
			ConvertNonPrintableString(output, true)

		fmt.Println(printableChars)

	*/

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}
