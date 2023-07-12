package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"os"
	"strings"
	"time"
)

type MainFileOpsTest008 struct {
	input string
}

func (dirOpsTest008 MainFileOpsTest008) GetFiles01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainDirOpsTest007.GetFiles01()",
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
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetDir := strMechOpsBaseDir +
		"\\fileOpsTest\\filesForTest\\levelfilesfortest\\level_01_dir\\level_02_dir\\level_03_dir"

	osPathSepStr := string(os.PathSeparator)

	targetDir = strings.Replace(
		targetDir,
		"\\",
		osPathSepStr,
		-1)

	var numOFilesLocated int
	var filesLocated strmech.FileMgrCollection

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

	leftMargin := " "
	rightMargin := ""
	maxLineLength := 80
	solidLineChar := "-"

	netFieldLength := maxLineLength -
		len(leftMargin) -
		len(rightMargin) - 1

	topTitle := strmech.TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    strmech.TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   strmech.TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        0,
	}

	err = topTitle.AddTitleLineStrings(
		ePrefix,
		"Selected Files\n",
		targetDir+"\n")

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	dateFmtStr := new(strmech.DateTimeHelper).
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

	bottomTitle := strmech.TextLineTitleMarqueeDto{
		StandardSolidLineLeftMargin:  leftMargin,
		StandardSolidLineRightMargin: rightMargin,
		StandardTitleLeftMargin:      leftMargin,
		StandardTitleRightMargin:     rightMargin,
		StandardMaxLineLen:           maxLineLength,
		StandardTextFieldLen:         netFieldLength,
		StandardTextJustification:    strmech.TxtJustify.Center(),
		NumLeadingBlankLines:         1,
		LeadingSolidLineChar:         solidLineChar,
		NumLeadingSolidLines:         1,
		NumTopTitleBlankLines:        0,
		TitleLines:                   strmech.TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:     0,
		TrailingSolidLineChar:        solidLineChar,
		NumTrailingSolidLines:        1,
		NumTrailingBlankLines:        1,
	}

	var delimitedNumStr string

	delimitedNumStr,
		err = filesLocated.
		GetTotalFileBytesCommaSeparated(
			ePrefix.XCpy("<-delimitedNumStr"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = bottomTitle.AddTitleLineStrings(
		ePrefix,
		fmt.Sprintf("Total File Bytes: %v",
			delimitedNumStr))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	strBuilder := strings.Builder{}

	err = filesLocated.
		GetTextListing(
			leftMargin,
			rightMargin,
			maxLineLength,
			topTitle,
			bottomTitle,
			&strBuilder,
			ePrefix.XCpy("<-filesLocated"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	/*
		output := strBuilder.String()

		printableChars := new(strmech.StrMech).
			ConvertNonPrintableString(output, true)

		fmt.Println(printableChars)

	*/

	fmt.Println(strBuilder.String())

	fmt.Printf(breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}

// ReadFiles01
//
// Runs test on FileHelper.ReadFileStrBuilderOpenClose()
func (dirOpsTest008 MainFileOpsTest008) ReadFiles01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileOpsTest008.ReadFiles01()",
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
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	fhHelper := new(strmech.FileHelper)

	strBuilder := new(strings.Builder)

	var numBytesRead int64

	numBytesRead,
		err = fhHelper.ReadFileStrBuilderOpenClose(
		targetInputFileName,
		strBuilder,
		ePrefix.XCpy("targetInputFileName->"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Number of Bytes Read: %v\n\n",
		numBytesRead)

	fmt.Printf("Length of 'strBuilder' string: %v\n\n",
		strBuilder.Len())

	fmt.Printf("%v",
		strBuilder.String())

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// ReadFiles02
//
// Runs test on FileHelper.ReadFileBytes()
func (dirOpsTest008 MainFileOpsTest008) ReadFiles02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileOpsTest008.ReadFiles02()",
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
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	fhHelper := new(strmech.FileHelper)

	var numBytesRead int64
	var bytesRead []byte

	bytesRead,
		numBytesRead,
		err = fhHelper.ReadFileBytes(
		targetInputFileName,
		ePrefix.XCpy("targetInputFileName->"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Number of Bytes Read: %v\n\n",
		numBytesRead)

	fmt.Printf("Length of 'bytesRead' Array: %v\n\n",
		len(bytesRead))

	fmt.Printf("%v",
		string(bytesRead))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// WriteFileBytes01
//
// Runs test on FileHelper.WriteFileBytes()
func (dirOpsTest008 MainFileOpsTest008) WriteFileBytes01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileOpsTest008.WriteFileBytes01()",
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
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	targetOutputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput_2.txt"

	targetOutputFileName = strings.Replace(
		targetOutputFileName,
		"\\",
		osPathSepStr,
		-1)

	fhHelper := new(strmech.FileHelper)

	var numBytesRead int64
	var bytesRead []byte

	bytesRead,
		numBytesRead,
		err = fhHelper.ReadFileBytes(
		targetInputFileName,
		ePrefix.XCpy("targetInputFileName->"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Number of Bytes Read: %v\n\n",
		numBytesRead)

	fmt.Printf("Length of 'bytesRead' Array: %v\n\n",
		len(bytesRead))

	fmt.Printf("%v",
		string(bytesRead))

	var bytesWritten int64

	bytesWritten,
		err = fhHelper.
		WriteFileBytes(
			targetOutputFileName,
			false,
			bytesRead,
			ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("Wrote Bytes to %v\n",
		targetInputFileName)

	fmt.Printf("Bytes Written: %v\n",
		bytesWritten)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// WriteFileBytes02
//
// This method tests FileMgr.WriteBytesToFile()
func (dirOpsTest008 MainFileOpsTest008) WriteFileBytes02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileOpsTest008.WriteFileBytes02()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	var err error
	var strMechOpsBaseDir string

	strMechOpsBaseDir,
		err = ExampleUtility{}.GetBaseDirectory(
		ePrefix.XCpy("strMechOpsBaseDir<-"))

	fmt.Printf("strMechOpsBaseDir: %v\n",
		strMechOpsBaseDir)

	targetInputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput.txt"

	osPathSepStr := string(os.PathSeparator)

	targetInputFileName = strings.Replace(
		targetInputFileName,
		"\\",
		osPathSepStr,
		-1)

	var inputFileMgr strmech.FileMgr

	inputFileMgr,
		err = new(strmech.FileMgr).New(
		targetInputFileName,
		ePrefix.XCpy("targetInputFileName"))

	if err != nil {

		_ = inputFileMgr.CloseThisFile(
			ePrefix.XCpy("inputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fileBytes []byte

	fileBytes,
		err = inputFileMgr.ReadAllFileBytes(
		ePrefix.XCpy("Read targetInputFileName"))

	if err != nil {

		_ = inputFileMgr.CloseThisFile(
			ePrefix.XCpy("inputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = inputFileMgr.CloseThisFile(
		ePrefix.XCpy("targetInputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	targetOutputFileName :=
		strMechOpsBaseDir +
			"\\apptest\\examples\\testoutput_2.txt"

	targetOutputFileName = strings.Replace(
		targetOutputFileName,
		"\\",
		osPathSepStr,
		-1)

	var outputFileMgr strmech.FileMgr

	outputFileMgr,
		err = new(strmech.FileMgr).New(
		targetOutputFileName,
		ePrefix.XCpy("targetOutputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.DeleteThisFile(
		ePrefix.XCpy("targetOutputFileName"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var numOfBytesWritten int

	numOfBytesWritten,
		err = outputFileMgr.WriteBytesToFile(
		fileBytes,
		true,
		ePrefix.XCpy("<-outputFileMgr"))

	if err != nil {

		_ = outputFileMgr.CloseThisFile(
			ePrefix.XCpy("outputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.CloseThisFile(
		ePrefix.XCpy("outputFileMgr"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nOutput File:\n"+
		"%v\n"+
		"Cycle 1 Number Of Bytes Written: %v\n\n",
		outputFileMgr.GetAbsolutePathFileName(),
		numOfBytesWritten)

	fileBytes = make([]byte, 0)

	fileBytes = []byte("\nHello World!!\n")

	numOfBytesWritten,
		err = outputFileMgr.WriteBytesToFile(
		fileBytes,
		false,
		ePrefix.XCpy("<-outputFileMgr"))

	if err != nil {

		_ = outputFileMgr.CloseThisFile(
			ePrefix.XCpy("outputFileMgr"))

		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	err = outputFileMgr.CloseThisFile(
		ePrefix.XCpy("outputFileMgr"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\nOutput File:\n"+
		"%v\n"+
		"Cycle 2 Number Of Bytes Written: %v\n\n",
		outputFileMgr.GetAbsolutePathFileName(),
		numOfBytesWritten)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}
