package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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

type DirTreeCopyStats struct {
	TotalDirsScanned uint64
	// The total number of directories scanned
	// during the current directory tree copy
	// operation.

	DirsCopied uint64
	// The number of directories copied.

	DirsCreated uint64
	// The number of target directories created.

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

	SubDirs uint64
	// The total number of subdirectories identified
	// during the directory tree copy operation. This
	// does NOT include the parent directory.

	ComputeError error
	// Errors related to computations or
	// conflicted category counts.
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

	dTreeCopyStats.SubDirs +=
		dCopyStats.SubDirsDocumented

}

// DirectoryProfile
//
// This structure contains status and statistical
// information on a single directory.
type DirectoryProfile struct {
	DirAbsolutePath string
	// The absolute directory path for the
	// directory described by this profile
	// information.

	DirExistsOnStorageDrive bool
	// If 'true', this paramter signals
	// that the directory actually exists on
	// a storage drive.

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

	DirSubDirectoriesBytes uint64
	// The total size of all Subdirectory entries
	// residing in the subject directory expressed
	// in bytes.

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

	ComputeError error
	// Computational or processing errors will be
	// recorded through this parameter.

	lock *sync.Mutex
}

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

	dirProfile.DirSubDirectoriesBytes +=
		incomingDirProfile.DirSubDirectoriesBytes

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

	return
}

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

	FilesNotCopied uint64
	// The number of files NOT copied.
	// Does NOT include directory entries.
	//
	FileBytesNotCopied uint64
	// The number of bytes associated with
	// files processed but NOT copied. Does
	// NOT include directory entries.

	SubDirsDocumented uint64
	// The number of subdirectories identified
	// and returned in a Directory Manager
	// Collection. Does NOT include the parent
	// directory. Subdirectories are only
	// documented if requested. This computation
	// is therefore optional.

	ComputeError error
	// Errors related to computations or
	// conflicted category counts.
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

type DeleteDirFilesStats struct {
	TotalFilesProcessed        uint64
	FilesDeleted               uint64
	FilesDeletedBytes          uint64
	FilesRemaining             uint64
	FilesRemainingBytes        uint64
	TotalSubDirectories        uint64
	TotalDirsScanned           uint64
	NumOfDirsWhereFilesDeleted uint64
	DirectoriesDeleted         uint64
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
}
