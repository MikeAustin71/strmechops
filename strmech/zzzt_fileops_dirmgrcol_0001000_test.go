package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDirMgrCollection_GetTextListingAbsPath_000100(t *testing.T) {

	funcName := "TestDirMgrCollection_GetTextListingAbsPath_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	osPathSepStr := string(os.PathSeparator)

	targetDirString := "\\fileOpsTest\\filesForTest\\levelbfilesfortest"

	targetDirString = strings.Replace(
		targetDirString,
		"\\",
		osPathSepStr,
		-1)

	var targetDir string
	var err error

	targetDir,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			targetDirString,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !strings.Contains(targetDir, "filesForTest") {

		t.Errorf("\n%v\n"+
			"Error: Returned 'targetDir' string does NOT\n"+
			"contain '%v'\n"+
			"targetDir= '%v'\n"+
			"targetDirString= '%v'\n",
			funcName,
			targetDirString,
			targetDir,
			targetDirString)

		return

	}

	var numOfDirectoriesLocated int
	var isParentDirectoryIncluded bool
	var directoriesLocated DirMgrCollection

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

	leftMargin := " "
	rightMargin := ""
	maxLineLength := 90
	solidLineChar := "-"

	netFieldLength := maxLineLength -
		len(leftMargin) -
		len(rightMargin) - 1

	var totalBytesNumStr string

	totalBytesNumStr,
		err = directoriesLocated.GetTotalBytesCommaSeparated(
		ePrefix.XCpy("directoriesLocated"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	topTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        0,
	}

	err = topTitle.AddTitleLineStrings(
		ePrefix,
		"Selected Directory",
		"    ",
		targetDir,
		"    ")

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	dateFmtStr := new(DateTimeHelper).
		GetDateTimeFormat(
			2)

	err = topTitle.AddTitleLineDateTimeStr(
		time.Now(),
		dateFmtStr,
		ePrefix.XCpy("<-time.Now()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	bottomTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        1,
	}

	err =
		bottomTitle.AddTitleLineStrings(
			ePrefix.XCpy("bottomTitle"),
			fmt.Sprintf("Number of Directories: %v",
				directoriesLocated.GetNumOfDirs()),
			fmt.Sprintf("Total Bytes in all Directories: %v",
				totalBytesNumStr))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	strBuilder := strings.Builder{}

	err = directoriesLocated.
		GetTextListingAbsPath(
			leftMargin,
			rightMargin,
			maxLineLength,
			topTitle,
			bottomTitle,
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
		"Number of Directories: 5") {

		t.Errorf("\n%v\n"+
			"Error: The Text Listing did NOT contain:\n"+
			"'Number of Directories: 4'\n"+
			" directoriesLocated.GetTextListingAbsPath().\n"+
			"targetDir='%v'\n"+
			"outputStr=\n%v\n",
			ePrefix.String(),
			targetDir,
			outputStr)

		return

	}

	if numOfDirectoriesLocated != 5 {

		t.Errorf("\n%v\n"+
			"Error: Expected numOfDirectoriesLocated == 10\n"+
			"Instead, numOfDirectoriesLocated = '%v'\n"+
			"targetDir='%v'\n",
			ePrefix.String(),
			numOfDirectoriesLocated,
			targetDir)

		return

	}

	if isParentDirectoryIncluded == true {

		t.Errorf("\n%v\n"+
			"Error: Expected isParentDirectoryIncluded == 'false'\n"+
			"Instead, isParentDirectoryIncluded = 'true'\n"+
			"targetDir='%v'\n",
			ePrefix.String(),
			targetDir)

		return

	}

	expectedTotalBytes := "12,249"

	if expectedTotalBytes != totalBytesNumStr {

		t.Errorf("\n%v\n"+
			"Error: Expected Total Bytes != Actual Total Bytes\n"+
			"Expected Total Bytes = '%v'\n"+
			"  Actual Total Bytes = '%v'\n",
			ePrefix.String(),
			expectedTotalBytes,
			totalBytesNumStr)

		return

	}

	return
}

func TestDirMgrCollection_GetDirProfile_000100(t *testing.T) {

	funcName := "TestDirMgrCollection_GetTextListingAbsPath_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetDir, trashDirectory, outputFile string
	var err error

	targetDir,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			FILEOpsBaseLevelFilesForTest,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	trashDirectory,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			FILEOpsBaseTrashDirectory,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	outputFile = trashDirectory +
		string(os.PathSeparator) +
		"testGetDirProfile_000100.txt"

	err = new(DirHelper).DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fHelper := new(FileHelper)

	err = fHelper.DeleteDirOrFile(
		outputFile,
		ePrefix.XCpy("outputFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	// We know outputFile does NOT exist on
	// disk.

	var dirProfile DirectoryProfile
	var directoryPathDoesExist bool

	directoryPathDoesExist,
		dirProfile,
		err = new(DirHelper).
		GetDirectoryProfile(
			targetDir,
			false,
			false,
			FileSelectionCriteria{},
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !directoryPathDoesExist {

		t.Errorf("\n%v\n"+
			"Error: Test Directory Path Does NOT Exist on\n"+
			"attached storage volume!\n"+
			"Base Test Directory = '%v'\n"+
			"targetDir='%v'\n",
			funcName,
			FILEOpsBaseLevelFilesForTest,
			targetDir)

		return

	}

	strBuilder1 := strings.Builder{}

	leftMargin := " "
	rightMargin := ""
	maxLineLength := 90

	netFieldLength := maxLineLength -
		len(leftMargin) -
		len(rightMargin)

	solidLineChar := "-"

	topTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        0,
	}

	err = topTitle.AddTitleLineStrings(
		ePrefix.XCpy("title lines"),
		"Directory Metrics")

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	dateFmtStr := new(DateTimeHelper).
		GetDateTimeFormat(
			2)

	err = topTitle.AddTitleLineDateTimeStr(
		time.Now(),
		dateFmtStr,
		ePrefix.XCpy("<-time.Now()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	bottomTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        0,
		NumTrailingBlankLines:        0,
	}

	err = dirProfile.GetTextListing(
		leftMargin,
		rightMargin,
		maxLineLength,
		topTitle,
		bottomTitle,
		&strBuilder1,
		ePrefix.XCpy("<-dirProfile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesWritten int64

	numOfBytesWritten,
		err = fHelper.
		WriteStrOpenClose(
			outputFile,
			true,
			true,
			strBuilder1.String(),
			ePrefix.XCpy("outputFile<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesRead int64

	strBuilder2 := strings.Builder{}

	numOfBytesRead,
		err = fHelper.ReadFileStrBuilderOpenClose(
		outputFile,
		&strBuilder2,
		ePrefix.XCpy("outputFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numOfBytesRead != numOfBytesWritten {

		t.Errorf("\n%v\n"+
			"Error: Bytes Written to output file do NOT\n"+
			"equal the Bytes Read from that output file!\n"+
			"Output File= '%v'\n"+
			"Bytes Written= '%v'\n"+
			"Bytes Read= '%v'\n",
			funcName,
			outputFile,
			numOfBytesWritten,
			numOfBytesRead)

		return

	}

	err = fHelper.DeleteDirOrFile(
		outputFile,
		ePrefix.XCpy("Final Delete outputFile"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

}

func TestDirMgrCollection_GetDirProfile_000200(t *testing.T) {

	funcName := "TestDirMgrCollection_GetTextListingAbsPath_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var targetDir, trashDirectory, outputFile string
	var err error

	targetDir,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			FILEOpsBaseFilesForTest,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	trashDirectory,
		err = new(fileOpsTestUtility).
		GetCompositeDir(
			FILEOpsBaseTrashDirectory,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = new(DirHelper).DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	outputFile = trashDirectory +
		string(os.PathSeparator) +
		"testGetDirProfile_000200.txt"

	var dirProfile DirectoryProfile
	var directoryPathDoesExist bool

	directoryPathDoesExist,
		dirProfile,
		err = new(DirHelper).
		GetDirectoryTreeProfile(
			targetDir,
			false,
			false,
			false,
			FileSelectionCriteria{},
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	if !directoryPathDoesExist {

		t.Errorf("\n%v\n"+
			"Error: Test Directory Does NOT Exist!\n"+
			"Test Directory = '%v'\n",
			ePrefix.String(),
			targetDir)

		return
	}

	strBuilder := strings.Builder{}

	leftMargin := " "
	rightMargin := ""
	maxLineLength := 90
	solidLineChar := "-"

	netFieldLength := maxLineLength -
		len(leftMargin) -
		len(rightMargin) - 1

	topTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        0,
	}

	err = topTitle.AddTitleLineStrings(
		ePrefix,
		"Directory Metrics")

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	dateFmtStr := new(DateTimeHelper).
		GetDateTimeFormat(
			2)

	err = topTitle.AddTitleLineDateTimeStr(
		time.Now(),
		dateFmtStr,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	bottomTitle := TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        0,
		NumTrailingBlankLines:        0,
	}

	err = dirProfile.GetTextListing(
		leftMargin,
		rightMargin,
		maxLineLength,
		topTitle,
		bottomTitle,
		&strBuilder,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesWritten int64

	numOfBytesWritten,
		err = new(FileHelper).
		WriteStrOpenClose(
			outputFile,
			true,
			true,
			strBuilder.String(),
			ePrefix.XCpy("outputFile<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var outputFileMgr FileMgr

	outputFileMgr,
		err = new(FileMgr).New(
		outputFile,
		ePrefix.XCpy("outputFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var bytesRead []byte

	bytesRead,
		err = outputFileMgr.ReadAllFileBytes(
		ePrefix.XCpy("outputFileMgr"))

	if err != nil {

		_ = outputFileMgr.CloseThisFile(
			nil)

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	_ = outputFileMgr.CloseThisFile(
		nil)

	err = new(DirHelper).DeleteAllInParentDirectory(
		trashDirectory,
		ePrefix.XCpy("trashDirectory"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var numOfBytesRead int64

	numOfBytesRead = int64(len(bytesRead))

	if numOfBytesRead != numOfBytesWritten {

		t.Errorf("\n%v\n"+
			"Error: Bytes Written to output file do NOT\n"+
			"equal the Bytes Read from that output file!\n"+
			"Output File= '%v'\n"+
			"Bytes Written= '%v'\n"+
			"Bytes Read= '%v'\n",
			funcName,
			outputFile,
			numOfBytesWritten,
			numOfBytesRead)

		return

	}

	return
}
