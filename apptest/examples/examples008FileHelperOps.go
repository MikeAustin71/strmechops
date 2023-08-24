package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"os"
	"strings"
	"time"
)

type MainFileHelperOpsTest008 struct {
	input string
}

func (fileHlprOpsTest008 MainFileHelperOpsTest008) GetFiles01() {

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
func (fileHlprOpsTest008 MainFileHelperOpsTest008) ReadFiles01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileHelperOpsTest008.ReadFiles01()",
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
// Runs test on FileHelper.Read()
func (fileHlprOpsTest008 MainFileHelperOpsTest008) ReadFiles02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileHelperOpsTest008.ReadFiles02()",
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

// ReadTextLines01
//
// Reads single lines of text from a file.
func (fileHlprOpsTest008 MainFileHelperOpsTest008) ReadTextLines01() {

	funcName := "MainFileHelperOpsTest008.ReadTextLines01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := " " + strings.Repeat("=",
		len(funcName)+6)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var shouldReadAndWriteFilesBeEqual,
		useWindowsLineTerminationChars bool

	shouldReadAndWriteFilesBeEqual = true

	useWindowsLineTerminationChars = true

	var err error
	var targetReadFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var targetWriteFile string

	targetWriteFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\ReadTextLines01.txt",
		ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fHelper := new(strmech.FileHelper)

	err = fHelper.
		DeleteDirOrFile(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var readEndOfLineDelimiters strmech.StringArrayDto

	if useWindowsLineTerminationChars {

		// Windows Output Format
		readEndOfLineDelimiters.PushStr("\r\n")

	} else {

		// Unix Output Format
		readEndOfLineDelimiters.PushStr("\n")
	}

	var strArrayDto strmech.StringArrayDto
	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int

	originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&readEndOfLineDelimiters,
		&strArrayDto,
		-1, // maxNumOfTextLines
		ePrefix.XCpy("strArrayDto"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\n %v\n"+
		"   Original File Size= '%v'\n"+
		" Number of Lines Read= '%v'\n"+
		" Number of Bytes Read= '%v'\n"+
		" Target Read File: %v\n\n",
		ePrefix.String(),
		originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		targetReadFile)

	var numBytesWritten int64

	var numOfLinesToWrite int
	var numOfBytesToWrite int64

	numOfLinesToWrite = strArrayDto.GetStringArrayLength()
	numOfBytesToWrite = strArrayDto.GetTotalBytesInStrings()

	readableTextLineTerm := new(strmech.StrMech).
		ConvertNonPrintableString(
			readEndOfLineDelimiters.StrArray[0],
			false)

	fmt.Printf(" %v\n"+
		" strArrayDto Stats Before File Write\n"+
		" Number of Lines To Write: %v\n"+
		" Number of Bytes To Write: %v\n"+
		" Text Line Terminator: %v\n",
		ePrefix,
		numOfLinesToWrite,
		numOfBytesToWrite,
		readableTextLineTerm)

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		targetWriteFile,
		true,
		true,
		strArrayDto.ConcatenateStrings(
			readEndOfLineDelimiters.StrArray[0]),
		ePrefix.XCpy("targetWriteFile<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" %v\n"+
		"After fHelper.WriteStrOpenClose()"+
		" Number of Bytes Written= %v\n"+
		" Target Write File= %v\n",
		ePrefix.String(),
		numBytesWritten,
		targetWriteFile)

	var doesFileExist bool
	var writerFileInfoPlus strmech.FileInfoPlus

	doesFileExist,
		writerFileInfoPlus,
		err = fHelper.
		DoesFileInfoPlusExist(
			targetWriteFile,
			ePrefix.XCpy("targetWriteFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf(" %v\n"+
		" Write File Stats from FileHelper.DoesFileInfoPlusExist()\n"+
		"    Write File Does Exist: %v\n"+
		" Write File Size in Bytes: %v\n"+
		"          Write File Name: %v\n"+
		"	 	Target Write File: %v\n\n",
		ePrefix.String(),
		doesFileExist,
		writerFileInfoPlus.Size(),
		writerFileInfoPlus.Name(),
		targetWriteFile)

	var filesAreEqual bool

	if shouldReadAndWriteFilesBeEqual == true {

		var reasonFilesNotEqual string

		filesAreEqual,
			reasonFilesNotEqual,
			err = fHelper.CompareFiles(
			targetReadFile,
			targetWriteFile,
			ePrefix.XCpy(
				"Target Files Comparison"))

		if err != nil {

			fmt.Printf(" %v\n"+
				"Error Return from fHelper.CompareFiles()\n"+
				"  targetReadFile= %v\n"+
				" targetWriteFile= %v\n"+
				"Reason: %v\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile,
				reasonFilesNotEqual)

			return
		}

		if !filesAreEqual {

			fmt.Printf("%v\n"+
				"Error: Read and Write Files are NOT equal!\n"+
				"Reason: %v\n",
				ePrefix.String(),
				reasonFilesNotEqual)

			return

		} else {

			fmt.Printf(" %v\n"+
				" SUCCESS! Files are EQUAL!\n"+
				"  Target Read File: %v\n"+
				" Target Write File: %v\n\n",
				ePrefix.String(),
				targetReadFile,
				targetWriteFile)

		}

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

// ReadLines02
//
// Reads single lines of text from a file.
func (fileHlprOpsTest008 MainFileHelperOpsTest008) ReadLines02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainFileHelperOpsTest008.ReadLines02()",
		"")

	breakStr := " " + strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n\n\n")

	var err error
	var targetReadFile, targetOutputFile string
	var exampleUtil = ExampleUtility{}

	targetReadFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\filesForTest\\textFilesForTest\\splitFunc2.txt",
		ePrefix.XCpy("targetReadFile"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	targetOutputFile,
		err = exampleUtil.GetCompositeDirectory(
		"fileOpsTest\\trashDirectory\\testOutput.txt",
		ePrefix.XCpy("readFileAddOn"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	var fHelper = new(strmech.FileHelper)

	var outputLinesArray,
		endOfLineDelimiters strmech.StringArrayDto

	var originalFileSize, numOfBytesRead int64
	var numOfLinesRead int

	endOfLineDelimiters.AddManyStrings(
		"\r",
		"\r\r",
		"[EOL]")

	originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		err = fHelper.ReadTextLines(
		targetReadFile,
		&endOfLineDelimiters,
		&outputLinesArray,
		-1, // maxNumOfLines
		ePrefix.XCpy("outputLinesArray"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("\n%v\n"+
		"  Original File Size= '%v'\n"+
		"Number of Lines Read= '%v'\n"+
		"Number of Bytes Read= '%v'\n\n",
		ePrefix.String(),
		originalFileSize,
		numOfLinesRead,
		numOfBytesRead)

	//outputLinesArray.ConvertToPrintableChars()

	var numBytesWritten int64

	numBytesWritten,
		err = fHelper.WriteStrOpenClose(
		targetOutputFile,
		true,
		true,
		outputLinesArray.ConcatenateStrings("\n"),
		ePrefix.XCpy("targetOutputFile<-"))

	if err != nil {
		fmt.Printf("\n%v\n\n",
			err.Error())
		return
	}

	fmt.Printf("%v\n"+
		"targetOutputFile= %v\n"+
		"Number of Bytes Written= %v\n",
		ePrefix.String(),
		targetOutputFile,
		numBytesWritten)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		" Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}
