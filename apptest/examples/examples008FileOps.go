package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"strings"
)

type MainFileOpsTest008 struct {
	input string
}

func (dirOpsTest007 MainFileOpsTest008) GetFiles01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.MainDirOpsTest007()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	targetDir := "D:\\t00"

	var numOFilesLocated int
	var filesLocated strmech.FileMgrCollection
	var err error

	_, // numOfDirectoriesLocated
		_,                //isParentDirectoryIncluded
		_,                //directoriesLocated
		numOFilesLocated, // numOfFilesLocated,
		filesLocated,     // filesLocated
		err = new(strmech.DirHelper).
		GetSubDirsFilesInParentDir(
			targetDir,                       // directoryPath
			false,                           // getSubdirectories
			false,                           // includeParentDirectory
			false,                           // includeSubDirCurrentDirOneDot
			false,                           // includeSubDirParentDirTwoDots
			true,                            // getRegularFiles
			true,                            // getSymLinksFiles
			true,                            // getOtherNonRegularFiles
			strmech.FileSelectionCriteria{}, // subDirSelectCharacteristics
			strmech.FileSelectionCriteria{}, // subDirSelectCharacteristics
			"targetDir",                     // directoryPathLabel
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" Number Of Files Located: %v\n",
		numOFilesLocated)

	strBuilder := strings.Builder{}

	err = filesLocated.
		GetTextListing(
			" ",
			" ",
			80,
			'-',
			"Directory "+targetDir,
			true,
			&strBuilder,
			ePrefix.XCpy("<-filesLocated"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Println(strBuilder.String())

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}
