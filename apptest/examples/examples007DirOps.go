package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"os"
	"strings"
)

type MainDirOpsTest007 struct {
	input string
}

func (dirOpsTest007 MainDirOpsTest007) GetDirs01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.GetDirs01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	targetDir := "..\\fileOpsTest\\filesForTest"

	osPathSepStr := string(os.PathSeparator)

	targetDir = strings.Replace(
		targetDir,
		"\\",
		osPathSepStr,
		-1)

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
			"Directory "+targetDir,
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

// GetDirs02
//
// Testing Get Directory Tree
func (dirOpsTest007 MainDirOpsTest007) GetDirs02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.GetDirs02()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	targetDir := "..\\fileOpsTest\\filesForTest"
	osPathSepStr := string(os.PathSeparator)

	targetDir = strings.Replace(
		targetDir,
		"\\",
		osPathSepStr,
		-1)

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
		GetSubDirsFilesInDirTree(
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
			"Directory "+targetDir,
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

// GetDirProfile01
//
// Testing Get Directory Profile for parent directory.
func (dirOpsTest007 MainDirOpsTest007) GetDirProfile01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.GetDirProfile01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	targetDir := "..\\fileOpsTest\\filesForTest\\levelfilesfortest"

	osPathSepStr := string(os.PathSeparator)

	targetDir = strings.Replace(
		targetDir,
		"\\",
		osPathSepStr,
		-1)

	targetOutputFile := "..\\apptest\\examples\\mike.txt"

	targetOutputFile = strings.Replace(
		targetOutputFile,
		"\\",
		osPathSepStr,
		-1)

	var err error
	var dirProfile strmech.DirectoryProfile
	var directoryPathDoesExist bool

	directoryPathDoesExist,
		dirProfile,
		err = new(strmech.DirHelper).
		GetDirectoryProfile(
			targetDir,
			false,
			false,
			strmech.FileSelectionCriteria{},
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Directory Does Exist: %v\n",
		directoryPathDoesExist)

	strBuilder := strings.Builder{}

	err = dirProfile.GetTextListing(
		" ",
		"",
		80,
		'-',
		"Directory Metrics",
		true,
		&strBuilder,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Println(strBuilder.String())

	var numOfBytesWritten int

	numOfBytesWritten,
		err = new(strmech.FileHelper).
		WriteStrOpenClose(
			targetOutputFile,
			true,
			true,
			strBuilder.String(),
			ePrefix.XCpy("targetOutputFile<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\n\nNumber of Bytes Written: %v\n",
		numOfBytesWritten)

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}

// GetDirTreeProfile01
//
// Returns directory stats on an entire directory tree.
func (dirOpsTest007 MainDirOpsTest007) GetDirTreeProfile01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.GetDirProfile01()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var strMechOpsBaseDir string

	strMechOpsBaseDir,
		err = GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetDir := strMechOpsBaseDir +
		"\\fileOpsTest\\filesForTest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir"

	//	"\\fileOpsTest\\filesForTest\\levelfilesfortest"

	osPathSepStr := string(os.PathSeparator)

	targetDir = strings.Replace(
		targetDir,
		"\\",
		osPathSepStr,
		-1)

	targetOutputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\mike.txt"

	targetOutputFileName = strings.Replace(
		targetOutputFileName,
		"\\",
		osPathSepStr,
		-1)

	var dirProfile strmech.DirectoryProfile
	var directoryPathDoesExist bool

	directoryPathDoesExist,
		dirProfile,
		err = new(strmech.DirHelper).
		GetDirectoryTreeProfile(
			targetDir,
			false,
			false,
			false,
			strmech.FileSelectionCriteria{},
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Directory Does Exist: %v\n",
		directoryPathDoesExist)

	strBuilder := strings.Builder{}

	err = dirProfile.GetTextListing(
		" ",
		"",
		80,
		'-',
		"Directory Metrics",
		true,
		&strBuilder,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Println(strBuilder.String())

	var numOfBytesWritten int

	numOfBytesWritten,
		err = new(strmech.FileHelper).
		WriteStrOpenClose(
			targetOutputFileName,
			true,
			true,
			strBuilder.String(),
			ePrefix.XCpy("targetOutputFileName<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\n\nNumber of Bytes Written: %v\n",
		numOfBytesWritten)

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}
