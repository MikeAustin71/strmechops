package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// DirectoryStatsDto
//
// This type is used to accumulate and disseminate
// information and statistics on a directory tree.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	To properly initialize an instance of
//	DirectoryStatsDto, call method:
//		DirectoryStatsDto.New()
type DirectoryStatsDto struct {
	dMgr DirMgr
	// Identifies the parent directory associated with
	// this directory information.

	numOfFiles uint64
	// The number of files (all types) residing
	// within this directory ('dMgr').

	numOfSubDirs uint64
	// The number of subdirectories residing
	// within this directory

	numOfBytes uint64
	// The total number of bytes for all files
	// contained in this directory.

	isInitialized bool
	// Signals whether this instance of
	// has been properly initialized.

	lock *sync.Mutex
}

// GetDirectory
//
// Returns a deep copy of the Directory Manager object
// ('DirMgr') which identifies the directory path
// associated with these directory statistics.
func (dirStats *DirectoryStatsDto) GetDirectory(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirectoryStatsDto.GetDirectory()",
		"")

	if err != nil {

		return DirMgr{}, err
	}

	return dirStats.dMgr.CopyOut(
		ePrefix)
}

// IsInitialized
//
// Returns the initialized status as a boolean value.
//
// If this return value is set to 'false', the internal
// data structures have NOT been properly initialized
// and use of the internal data values may generate
// errors.
func (dirStats *DirectoryStatsDto) IsInitialized() bool {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	return dirStats.isInitialized
}

// New
//
// Returns a new instance of DirectoryStatsDto properly
// initialized with the Directory Manager associated
// with the target directory statistics.
func (dirStats *DirectoryStatsDto) New(
	dMgr DirMgr,
	errorPrefix interface{}) (
	DirectoryStatsDto,
	error) {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "DirectoryStatsDto.New()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return DirectoryStatsDto{}, err
	}

	var err2 error

	_,
		_,
		err2 = new(dirMgrHelperPreon).validateDirMgr(
		&dMgr,
		false, // Path is NOT required to exit on disk
		"dMgr",
		ePrefix.XCpy("dMgr"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return DirectoryStatsDto{}, err
	}

	newDirStats := DirectoryStatsDto{}

	err = newDirStats.dMgr.CopyIn(
		&dMgr,
		ePrefix.XCpy("dMgr"))

	if err == nil {
		newDirStats.isInitialized = true
	}

	return newDirStats, err
}

// NumOfFiles
//
// Returns the number of files found in the target
// directory.
func (dirStats *DirectoryStatsDto) NumOfFiles() uint64 {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	return dirStats.numOfFiles
}

// NumOfSubDirs
//
// Returns the number of subdirectories residing in the
// target directory.
func (dirStats *DirectoryStatsDto) NumOfSubDirs() uint64 {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	return dirStats.numOfSubDirs
}

// NumOfBytes
//
// Returns the total number of bytes contained
// in the target directory.
func (dirStats *DirectoryStatsDto) NumOfBytes() uint64 {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	return dirStats.numOfBytes
}

// SetDirectory
//
// Configures the internal Directory Manager object
// ('DirMgr') which identifies the directory path
// associated with these directory statistics.
func (dirStats *DirectoryStatsDto) SetDirectory(
	dMgr DirMgr,
	errorPrefix interface{}) error {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirectoryStatsDto.SetDirectory()",
		"")

	if err != nil {

		return err
	}

	err = dirStats.dMgr.CopyIn(
		&dMgr,
		ePrefix.XCpy("dMgr"))

	return err
}

// SetIsInitialized
//
// Sets the internal 'isInitialized' flag. Call this
// method to set the status of the DirectoryStatsDto
// instance.
//
// If input parameter 'isInitialized' is set to 'true',
// it signals that all the internal data values are
// valid.
func (dirStats *DirectoryStatsDto) SetIsInitialized(
	isInitialized bool) {

	if dirStats.lock == nil {
		dirStats.lock = new(sync.Mutex)
	}

	dirStats.lock.Lock()

	defer dirStats.lock.Unlock()

	dirStats.isInitialized = isInitialized

	return
}

// DirTreeCopyStats
//
// The data elements in this structure are used
// to accumulate statistics and information
// related to a file copy operation performed on
// source and destination directory trees.
type DirTreeCopyStats struct {
	TotalDirsScanned uint64
	// The total number of directories scanned
	// during the current directory tree copy
	// operation.

	DirsCopied uint64
	// The number of directories copied.

	DirsCreated uint64
	// The number of target directories created.

	TotalSubDirs uint64
	// The total number of subdirectories identified
	// during the directory tree copy operation. This
	// does NOT include the parent directory.

	TotalFilesProcessed uint64
	// The total number of files processed during
	// the directory tree copy operation.

	FilesCopied uint64
	// The total number of files copied to the
	// target directory tree during the directory
	// tree copy operation.

	FileBytesCopied uint64
	// The total number of file bytes copied to the
	// target directory tree during the directory
	// tree copy operation.

	FilesNotCopied uint64
	// The total number of files scanned and
	// processed, but NOT copied to the target
	// directory tree during the directory tree
	// copy operation.

	FileBytesNotCopied uint64
	// The total number of bytes associated with
	// files scanned and processed, but NOT copied
	// to the target directory tree during the
	// directory tree copy operation.

	SubDirsDocumented uint64
	// The number of subdirectories identified
	// and returned in a Directory Manager
	// Collection. Does NOT include the parent
	// directory. Subdirectories are only
	// documented if requested. This computation
	// value is therefore optional.

	CopiedFilesDocumented uint64
	// The number of copied files documented
	// by adding a File Manager object to a
	// returned File Manager Collection.

	Errors []error
	// An array of errors associated with the
	// calculation of these statistics.
}

func (dTreeCopyStats *DirTreeCopyStats) AddDirCopyStats(
	dCopyStats DirectoryCopyStats) {

	dTreeCopyStats.TotalDirsScanned++

	dTreeCopyStats.DirsCopied++

	dTreeCopyStats.DirsCreated += dCopyStats.DirsCreated

	dTreeCopyStats.TotalFilesProcessed += dCopyStats.TotalFilesProcessed

	dTreeCopyStats.FilesCopied += dCopyStats.FilesCopied

	dTreeCopyStats.FileBytesCopied += dCopyStats.FileBytesCopied

	dTreeCopyStats.FilesNotCopied += dCopyStats.FilesNotCopied

	dTreeCopyStats.FileBytesNotCopied += dCopyStats.FileBytesNotCopied

	dTreeCopyStats.SubDirsDocumented +=
		dCopyStats.SubDirsDocumented

	dTreeCopyStats.CopiedFilesDocumented +=
		dCopyStats.CopiedFilesDocumented

	if len(dCopyStats.Errors) > 0 {

		dTreeCopyStats.Errors = append(
			dTreeCopyStats.Errors,
			dCopyStats.Errors...)
	}

}

// DirectoryProfile
//
// This structure contains status and statistical
// information on a single directory.
type DirectoryProfile struct {
	ParentDirAbsolutePath string
	// The absolute directory path for the
	// directory described by this profile
	// information.

	ParentDirManager DirMgr
	// An instance of DirMgr encapsulating the
	// Directory Path and associated parameters
	// for the directory described by this profile
	// information.

	ParentDirIsIncludedInStats bool
	// If this parameter is set to 'true', it
	// signals that the directory statistics and
	// information provided by this instance of
	// DirectoryProfile includes metrics from
	// the parent directory.

	IsDirectoryTreeStats bool
	// If this parameter is set to 'true', it
	// signals that the metrics included in this
	// instance of DirectoryProfile are compiled
	// for a directory tree, and not an individual
	// directory.
	//
	// Conversely, if this parameter is set to
	// 'false', it signals that these directory
	// metrics describe a single directory and
	// not a directory tree.

	ParentDirExistsOnStorageDrive bool
	// If 'true', this paramter signals
	// that the parent directory actually
	// exists on a storage drive.

	DirTotalFiles uint64
	// The number of total files residing in
	// the subject directory. This includes
	// Regular Files, SymLink Files and
	// Non-Regular Files. It does NOT include
	// directory entry files.

	DirTotalFileBytes uint64
	// The size of all files residing in the
	// subject directory expressed in bytes.
	// This includes Regular Files, SymLink
	// Files and Non-Regular Files. It does
	// NOT include directory entry files.

	DirSubDirectories uint64
	// The number of subdirectories residing
	// within the subject directory. This

	SubDirsIncludeCurrentDirOneDot bool
	// All directories include an os.FileInfo entry for
	// the current directory. The current directory name
	// is always denoted as single dot ('.').
	//
	// When data element, 'SubDirsIncludeCurrentDirOneDot',
	// is set to 'true', the one dot current directory ('.')
	// will be included in the directory profile information
	// and counted as a separate subdirectory.

	SubDirsIncludeParentDirTwoDot bool
	// All directories include an os.FileInfo entry for
	// the parent directory. The parent directory name
	// is always denoted as two dots ('..').
	//
	// When data element, 'SubDirsIncludeParentDirTwoDot',
	// is set to 'true', the two dot ('..') parent directory,
	// will be included in the directory profile information
	// and counted as a separate subdirectory.

	DirRegularFiles uint64
	// The number of 'Regular' Files residing
	// within the subject Directory. Regular
	// files include text files, image files
	// and executable files. Reference:
	// https://www.computerhope.com/jargon/r/regular-file.htm

	DirRegularFileBytes uint64
	// The total size of all 'Regular' files
	// residing in the subject directory expressed
	// in bytes.

	DirSymLinkFiles uint64
	// The number of SymLink files residing in the
	// subject directory.

	DirSymLinkFileBytes uint64
	// The total size of all SymLink files
	// residing in the subject directory
	// expressed in bytes.

	DirNonRegularFiles uint64
	// The total number of Non-Regular files residing
	// in the subject directory.
	//
	// Non-Regular files include directories, device
	// files, named pipes, sockets, and symbolic links.

	DirNonRegularFileBytes uint64
	// The total size of all Non-Regular files residing
	// in the subject directory expressed in bytes.

	Errors []error
	// An array of errors associated with the
	// calculation of these statistics.

	lock *sync.Mutex
}

// AddDirProfileStats
//
// Receives an incoming instance of DirectoryProfile and
// proceeds to add the directory metrics to the directory
// metric values contained in the current instance of
// DirectoryProfile.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingDirProfile			DirectoryProfile
//
//		A concrete instance of DirectoryProfile. The
//		Directory statistics contained in this instance
//		will be added to those contained in the current
//		instance of AddDirProfileStats.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (dirProfile *DirectoryProfile) AddDirProfileStats(
	incomingDirProfile DirectoryProfile) {

	if dirProfile.lock == nil {
		dirProfile.lock = new(sync.Mutex)
	}

	dirProfile.lock.Lock()

	defer dirProfile.lock.Unlock()

	dirProfile.DirTotalFiles +=
		incomingDirProfile.DirTotalFiles

	dirProfile.DirTotalFileBytes +=
		incomingDirProfile.DirTotalFileBytes

	dirProfile.DirSubDirectories +=
		incomingDirProfile.DirSubDirectories

	dirProfile.DirRegularFiles +=
		incomingDirProfile.DirRegularFiles

	dirProfile.DirRegularFileBytes +=
		incomingDirProfile.DirRegularFileBytes

	dirProfile.DirSymLinkFiles +=
		incomingDirProfile.DirSymLinkFiles

	dirProfile.DirSymLinkFileBytes +=
		incomingDirProfile.DirSymLinkFileBytes

	dirProfile.DirNonRegularFiles +=
		incomingDirProfile.DirNonRegularFiles

	dirProfile.DirNonRegularFileBytes +=
		incomingDirProfile.DirNonRegularFileBytes

	dirProfile.Errors = append(
		dirProfile.Errors,
		incomingDirProfile.Errors...)

	return
}

// GetTextListing
//
// Returns a string containing a text listing displaying
// the directory metrics contained in the current
// instance of DirectoryProfile.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leftMargin					string
//
//		This string serves as the left margin for all
//		text lines.
//
//	rightMargin					string
//
//		This string serves as the right margin for all
//		text lines.
//
//
//	maxLineLength				int
//
//		This integer value defines the maximum line
//		length for all text lines. If this value is
//		less than 10, an error will be returned.
//
//	topTitleDisplay				TextLineTitleMarqueeDto
//
//		Contains specifications for the top tile display
//		including title lines and solid line breaks.
//
//		If no title is required, set this parameter to an
//		empty instance of TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no top title text lines are required, and the
//		solid line breaks are still necessary, simply
//		leave the title lines collection empty.
//
//	bottomTitleDisplay			TextLineTitleMarqueeDto
//
//		Contains specifications for the bottom tile
//		display including title lines and solid line
//		breaks.
//
//		If no bottom title is required, set this
//		parameter to an empty instance of
//		TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no bottom title text lines are required, and
//		the solid line breaks are still necessary, simply
//		leave the title lines collection empty.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder.
//		The text listing for directory absolute paths
//		will be added to this instance of
//		strings.Builder.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dirProfile *DirectoryProfile) GetTextListing(
	leftMargin string,
	rightMargin string,
	maxLineLength int,
	topTitleDisplay TextLineTitleMarqueeDto,
	bottomTitleDisplay TextLineTitleMarqueeDto,
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if dirProfile.lock == nil {
		dirProfile.lock = new(sync.Mutex)
	}

	dirProfile.lock.Lock()

	defer dirProfile.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirectoryProfile." +
		"GetTextListing()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

	}

	if maxLineLength < 10 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is less than '10'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	if maxLineLength > 1999 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is greater than '1999'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	lenCol1Field := len("SubDirsIncludeCurrentDirTwoDot")
	col1RightMarginStr := ": "
	lenCol1RightMargin := len(col1RightMarginStr)

	multiLineLeftMargin :=
		strings.Repeat(" ",
			lenCol1Field+
				lenCol1RightMargin+
				2)

	lenLeftMar := len(leftMargin)
	lenRightMar := len(rightMargin)
	lenCol1 := lenCol1Field +
		lenCol1RightMargin

	lenCol2 := 27

	lenLineTerminator := 1

	actualTextLineLen :=
		lenLeftMar +
			lenRightMar +
			lenCol1 +
			lenCol2 +
			lenLineTerminator

	if actualTextLineLen > maxLineLength {

		return fmt.Errorf("%v\n"+
			"Error: The Actual Text Line length is greater than."+
			"the Maximum Line Length!\n"+
			"The Actual Text Line length is calculated by adding\n"+
			"the left and right margin lengths, the Column1\n"+
			"and Column2 lengths and the Line Terminator length.\n"+
			"Actual Text Line Length =  Left Margin Length plus\n"+
			"Right Margin Length plus Column-1 Length plus Column-2\n"+
			"Length plus Line Terminator Length.\n"+
			"    Maximum Line Length = '%v'\n"+
			"Actual Text Line Length = '%v'\n"+
			"     Left Margin Length = '%v'\n"+
			"    Right Margin Length = '%v'\n"+
			"        Column-1 Length = '%v'\n"+
			"        Column-2 Length = '%v'\n"+
			" Line Terminator Length = '%v'\n",
			ePrefix.String(),
			maxLineLength,
			actualTextLineLen,
			lenLeftMar,
			lenRightMar,
			lenCol1,
			lenCol2,
			lenLineTerminator)

	}

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	numOfTitleLines := 12

	thisReqCapacity := (actualTextLineLen * 17) +
		len(dirProfile.Errors) +
		(numOfTitleLines * maxLineLength)

	netRequiredCapacity :=
		thisReqCapacity - netCapacityStrBuilder

	if netRequiredCapacity > 0 {

		strBuilder.Grow(netRequiredCapacity + 64)
	}

	solidLineChar := "-"

	txtFormatCol := TextFormatterCollection{}

	if topTitleDisplay.IsValidInstance() {

		err = txtFormatCol.AddTextTitleMarqueeDto(
			topTitleDisplay,
			ePrefix.XCpy("<-topTitleDisplay"))

		if err != nil {
			return err
		}

		solidLineChar = topTitleDisplay.LeadingSolidLineChar

		if len(topTitleDisplay.TrailingSolidLineChar) > 0 {
			solidLineChar = topTitleDisplay.TrailingSolidLineChar
		}

		if len(solidLineChar) == 0 {
			solidLineChar = "-"
		}
	}

	effectiveLineLen := maxLineLength - lenLeftMar - lenRightMar - 1

	txtFormatCol.AddFieldLabel(
		leftMargin,
		"Parent Directory",
		effectiveLineLen,
		TxtJustify.Center(),
		rightMargin,
		"\n\n",
		maxLineLength,
		false,
		"")

	lenTxtParam := len(dirProfile.ParentDirAbsolutePath)

	if lenTxtParam > effectiveLineLen {

		txtFormatCol.AddFieldLabel(
			leftMargin,
			dirProfile.ParentDirAbsolutePath,
			-1,
			TxtJustify.Left(),
			rightMargin,
			"\n",
			maxLineLength,
			true,
			leftMargin)

	} else {

		txtFormatCol.AddFieldLabel(
			leftMargin,
			dirProfile.ParentDirAbsolutePath,
			-1,
			TxtJustify.Center(),
			rightMargin,
			"\n",
			maxLineLength,
			false,
			"")

	}

	if len(topTitleDisplay.TrailingSolidLineChar) > 0 {

		solidLineChar = topTitleDisplay.TrailingSolidLineChar

	} else if len(topTitleDisplay.LeadingSolidLineChar) > 0 {

		solidLineChar = topTitleDisplay.LeadingSolidLineChar

	}

	txtFormatCol.AddLineSolid(
		leftMargin,
		solidLineChar,
		effectiveLineLen,
		rightMargin,
		true,
		"\n",
		maxLineLength,
		false,
		"")

	err = txtFormatCol.
		SetStdFormatParamsLine2Col(
			leftMargin,
			lenCol1Field,
			TxtJustify.Right(),
			col1RightMarginStr,
			lenCol2,
			TxtJustify.Left(),
			"",
			false,
			"",
			maxLineLength,
			true,
			multiLineLeftMargin,
			ePrefix)

	if err != nil {
		return err
	}

	var txtStrLabel, txtStrParam string

	txtStrLabel = "ParentDirIsIncludedInStats"
	txtStrParam = fmt.Sprintf("%v",
		dirProfile.ParentDirIsIncludedInStats)

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrLabel = "IsDirectoryTreeStats"
	txtStrParam = fmt.Sprintf("%v",
		dirProfile.IsDirectoryTreeStats)

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrLabel = "ParentDirExistsOnStorageDrive"
	txtStrParam = fmt.Sprintf("%v",
		dirProfile.ParentDirExistsOnStorageDrive)

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	var intSep IntegerSeparatorSpec

	intSep,
		err = new(IntegerSeparatorSpec).
		NewUnitedStatesDefaults(
			ePrefix.XCpy(
				"intSep<-"))

	if err != nil {
		return err
	}

	var delimitedNumStr string

	// DirTotalFiles

	txtStrLabel = "DirTotalFiles"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirTotalFiles),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirTotalFileBytes

	txtStrLabel = "DirTotalFileBytes"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirTotalFileBytes),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirSubDirectories

	txtStrLabel = "DirSubDirectories"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirSubDirectories),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// SubDirsIncludeCurrentDirOneDot

	txtStrLabel = "SubDirsIncludeCurrentDirOneDot"
	txtStrParam = fmt.Sprintf("%v",
		dirProfile.SubDirsIncludeCurrentDirOneDot)

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// SubDirsIncludeParentDirTwoDot

	txtStrLabel = "SubDirsIncludeParentDirTwoDot"
	txtStrParam = fmt.Sprintf("%v",
		dirProfile.SubDirsIncludeParentDirTwoDot)

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirRegularFiles

	txtStrLabel = "DirRegularFiles"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirRegularFiles),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirRegularFileBytes

	txtStrLabel = "DirRegularFileBytes"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirRegularFileBytes),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirSymLinkFiles

	txtStrLabel = "DirSymLinkFiles"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirSymLinkFiles),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirSymLinkFiles

	txtStrLabel = "DirSymLinkFileBytes"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirSymLinkFileBytes),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	// DirNonRegularFiles

	txtStrLabel = "DirNonRegularFiles"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirNonRegularFiles),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrLabel = "DirNonRegularFileBytes"

	delimitedNumStr,
		err = intSep.
		GetFmtIntSeparatedNumStr(
			fmt.Sprintf("%v",
				dirProfile.DirNonRegularFileBytes),
			ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	txtStrParam = delimitedNumStr

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(txtStrLabel))

	if err != nil {
		return err
	}

	lenErrors := len(dirProfile.Errors)

	if lenErrors == 0 {

		txtStrLabel = "Errors"

		txtStrParam = "-- NONE --"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(txtStrLabel))

		if err != nil {
			return err
		}

	} else {

		txtFormatCol.AddFieldLabel(
			leftMargin,
			"Errors:",
			lenCol1Field,
			TxtJustify.Right(),
			" ",
			"",
			maxLineLength,
			false,
			"")

		txtFormatCol.AddLineBlank(
			1,
			"\n")

		var errStr1, errStr2 string

		for i := 0; i < lenErrors; i++ {

			errStr1 = fmt.Sprintf(leftMargin+"%v",
				dirProfile.Errors[i].Error())

			errStr2 = strings.Replace(
				errStr1,
				"\n",
				"\n"+leftMargin,
				-1)

			txtFormatCol.AddAdHocText(
				"",
				errStr2,
				"",
				false,
				"\n",
				maxLineLength,
				true,
				leftMargin)

		}

	}

	if bottomTitleDisplay.IsValidInstance() {

		err = txtFormatCol.AddTextTitleMarqueeDto(
			bottomTitleDisplay,
			ePrefix.XCpy("<-bottomTitleDisplay"))

		if err != nil {
			return err
		}

	} else {

		// Final Text Line Build
		txtFormatCol.AddLineBlank(1, "")

	}

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err

}

// SetDirMgr
//
// Receives an instance of Directory Manager (DirMgr) and
// proceeds to the set the DirectoryProfile profile
// member variables:
//
//	DirectoryProfile.ParentDirManager
//	DirectoryProfile.ParentDirAbsolutePath
//
// Receives a string containing a directory path and proceeds
// to convert that path to an absolute directory path before
// configuring the DirectoryProfile profile
// member variables:
//
//	DirectoryProfile.ParentDirManager
//	DirectoryProfile.ParentDirAbsolutePath
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of Directory Manager
//		(DirMgr). The directory path information
//		contained in this instance will be used to the
//		set the following DirectoryProfile profile member
//		variables:
//
//			DirectoryProfile.ParentDirManager
//			DirectoryProfile.ParentDirAbsolutePath
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dirProfile *DirectoryProfile) SetDirMgr(
	dMgr *DirMgr,
	dMgrLabel string,
	errorPrefix interface{}) error {

	if dirProfile.lock == nil {
		dirProfile.lock = new(sync.Mutex)
	}

	dirProfile.lock.Lock()

	defer dirProfile.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirectoryProfile.SetDirMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a nil pointer.\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel)

		return err
	}

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&dirProfile.ParentDirManager,
		dMgr,
		ePrefix.XCpy(
			"dirProfile.ParentDirManager<-dMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error occurred while copying %v Directory Manager"+
			"to Directory Profile.\n"+
			"new(dirMgrHelperBoson).copyDirMgrs()\n"+
			"%v Absolute Path = '%v'\n"+
			"Error = \n%v\n",
			funcName,
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err.Error())
	}

	dirProfile.ParentDirAbsolutePath =
		dirProfile.ParentDirManager.absolutePath

	dirProfile.ParentDirIsIncludedInStats = true

	return err
}

// SetDirPath
//
// Receives a string containing a directory path and proceeds
// to convert that path to an absolute directory path before
// configuring the DirectoryProfile profile
// member variables:
//
//	DirectoryProfile.ParentDirManager
//	DirectoryProfile.ParentDirAbsolutePath
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains a directory path which will
//		be converted to an absolute directory path before
//		configuring the following DirectoryProfile
//		profile member variables:
//
//			DirectoryProfile.ParentDirManager
//			DirectoryProfile.ParentDirAbsolutePath
//
//	dirPathLabel				string
//
//		The name or label associated with input parameter
//		'dirPath', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dirPath" will be
//		automatically applied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dirProfile *DirectoryProfile) SetDirPath(
	dirPath string,
	dirPathLabel string,
	errorPrefix interface{}) error {

	if dirProfile.lock == nil {
		dirProfile.lock = new(sync.Mutex)
	}

	dirProfile.lock.Lock()

	defer dirProfile.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "DirectoryProfile.SetDirMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "dirPath"
	}

	if len(dirPath) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is empty with a zero (0) string length.\n",
			ePrefix.String(),
			dirPathLabel,
			dirPathLabel)

		return err
	}

	isEmpty,
		err := new(dirMgrHelperNanobot).
		setDirMgr(
			&dirProfile.ParentDirManager,
			dirPath,
			"dirProfile.ParentDirManager",
			dirPathLabel,
			ePrefix)

	if err != nil {
		return err
	}

	if isEmpty {

		err = fmt.Errorf("%v"+
			"ERROR: %v returned an EMPTY DirMgr\n"+
			"pathStr='%v'\n",
			ePrefix.String(),
			dirPathLabel,
			dirPath)

		return err
	}

	dirProfile.ParentDirAbsolutePath =
		dirProfile.ParentDirManager.absolutePath

	dirProfile.ParentDirIsIncludedInStats = true

	return err
}

// DirectoryCopyStats
//
// The data elements in this structure are used
// to accumulate statistics and information
// related to files copied from a single source
// directory to a single destination or target
// directory.
type DirectoryCopyStats struct {
	DirsCreated uint64
	// The number of new directories created.

	TotalFilesProcessed uint64
	// The total number of files processed.
	// Does NOT include directory entries.

	FilesCopied uint64
	// The number of files copied. Does
	// NOT include directory entries.

	FileBytesCopied uint64
	// The number of file bytes copied.
	// Does NOT include directory entries.

	FilesNotCopied uint64
	// The number of files processed, but
	// NOT copied. Does NOT include directory
	// entries.

	FileBytesNotCopied uint64
	// The number of bytes associated with
	// files processed but NOT copied. Does
	// NOT include directory entries.

	SubDirs uint64
	// The total number of subdirectories identified
	// during the directory tree copy operation. This
	// does NOT include the parent directory.

	SubDirsDocumented uint64
	// The number of subdirectories identified
	// and returned in a Directory Manager
	// Collection. Does NOT include the parent
	// directory. Subdirectories are only
	// documented if requested. This computation
	// value is therefore optional.

	CopiedFilesDocumented uint64
	// The number of copied files documented
	// by adding a File Manager object to a
	// returned File Manager Collection.

	Errors []error
	// An array of errors associated with the
	// calculation of these statistics.
}

// DirectoryDeleteFileInfo - structure used
// to delete files in a directory specified
// by 'StartPath'. Deleted files will be selected
// based on 'DeleteFileSelectCriteria' value.
//
// 'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
// type which contains FileNamePatterns strings and the
// FilesOlderThan or FilesNewerThan date time parameters
// which can be used as file selection criteria.
type DirectoryDeleteFileInfo struct {
	StartPath string
	// The starting directory path for the deletion
	// operation.

	Directories DirMgrCollection
	// The Directories actively scanned and included in
	// this file deletion operation.

	ErrReturns []error
	// This array of errors includes all errors, both fatal
	// and non-fatal, encountered in the deletion operation.
	// If a fatal error occurred, it will be the last error
	// in the array.

	DeleteFileSelectCriteria FileSelectionCriteria
	// The File Selection criteria used to select the files
	// deleted in the file deletion operation.

	DeletedFiles FileMgrCollection
	// A collection of File Manager (FileMgr) objects
	// identifying the files that were actually deleted in
	// the file deletion operation.
}

type DirectoryMoveStats struct {
	TotalSrcFilesProcessed   uint64
	SourceFilesMoved         uint64
	SourceFileBytesMoved     uint64
	SourceFilesRemaining     uint64
	SourceFileBytesRemaining uint64
	SourceSubDirsMoved       uint64
	SourceSubDirsRemaining   uint64
	TotalDirsProcessed       uint64
	TargetDirsCreated        uint64
	SourceOriginalSubDirs    uint64
	SourceDirWasDeleted      bool
	ComputeError             error
}

// DeleteDirFilesStats
//
// The data elements in this structure are used
// to accumulate statistics and information
// related to the deletion of files from a single
// target directory.
type DeleteDirFilesStats struct {
	TotalFilesProcessed uint64
	// The total number of files processed.
	// Does NOT include directory entries.

	FilesDeleted uint64
	// The number of files deleted. Does
	// NOT include directory entries.

	FilesDeletedBytes uint64
	// The number of file bytes deleted.
	// Does NOT include directory entries.

	FilesRemaining uint64
	// The number of files processed, but
	// NOT deleted. Does NOT include directory
	// entries.

	FilesRemainingBytes uint64
	// The number of bytes associated with
	// files processed but NOT copied. Does
	// NOT include directory entries.

	TotalSubDirectories uint64
	// Total SubDirectories processed

	TotalDirsScanned uint64
	// Total Directories Scanned.

	NumOfDirsWhereFilesDeleted uint64
	// The number of parent directories and
	// subdirectories where files were deleted.

	DirectoriesDeleted uint64
	// The number of directories deleted.

	SubDirsDocumented uint64
	// The number of subdirectories identified
	// and returned in a Directory Manager
	// Collection. Does NOT include the parent
	// directory. Subdirectories are only
	// documented if requested. This computation
	// value is therefore optional.

	DeletedFilesDocumented uint64
	// The number of deleted files documented
	// by adding a File Manager object to a
	// returned File Manager Collection.

	Errors []error
	// An array of errors associated with the
	// calculation of these statistics.
}

// AddStats
//
// Receives another instance of DeleteDirFilesStats and
// adds those deletion statistics to those contained in
// the current instance of DeleteDirFilesStats.
func (delDirFileStats *DeleteDirFilesStats) AddStats(
	delDirFStats2 DeleteDirFilesStats) {

	delDirFileStats.TotalFilesProcessed +=
		delDirFStats2.TotalFilesProcessed

	delDirFileStats.FilesDeleted +=
		delDirFStats2.FilesDeleted

	delDirFileStats.FilesDeletedBytes +=
		delDirFStats2.FilesDeletedBytes

	delDirFileStats.FilesRemaining +=
		delDirFStats2.FilesRemaining

	delDirFileStats.FilesRemainingBytes +=
		delDirFStats2.FilesRemainingBytes

	delDirFileStats.TotalSubDirectories +=
		delDirFStats2.TotalSubDirectories

	delDirFileStats.TotalDirsScanned +=
		delDirFStats2.TotalDirsScanned

	delDirFileStats.NumOfDirsWhereFilesDeleted +=
		delDirFStats2.NumOfDirsWhereFilesDeleted

	delDirFileStats.DirectoriesDeleted +=
		delDirFStats2.DirectoriesDeleted

	delDirFileStats.SubDirsDocumented +=
		delDirFStats2.SubDirsDocumented

	delDirFileStats.DeletedFilesDocumented +=
		delDirFStats2.DeletedFilesDocumented

	if len(delDirFStats2.Errors) > 0 {

		delDirFileStats.Errors =
			append(delDirFileStats.Errors,
				delDirFStats2.Errors...)

	}

}
