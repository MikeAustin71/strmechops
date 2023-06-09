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
	TotalDirsScanned    uint64
	DirsCopied          uint64
	DirsCreated         uint64
	TotalFilesProcessed uint64
	FilesCopied         uint64
	FileBytesCopied     uint64
	FilesNotCopied      uint64
	FileBytesNotCopied  uint64
	ComputeError        error
}

func (dTreeCopyStats *DirTreeCopyStats) AddDirCopyStats(
	dCopyStats DirectoryCopyStats) {

	dTreeCopyStats.DirsCreated += dCopyStats.DirsCreated

	dTreeCopyStats.TotalFilesProcessed += dCopyStats.TotalFilesProcessed

	dTreeCopyStats.FilesCopied += dCopyStats.FilesCopied

	dTreeCopyStats.FileBytesCopied += dCopyStats.FileBytesCopied

	dTreeCopyStats.FilesNotCopied += dCopyStats.FilesNotCopied

	dTreeCopyStats.FileBytesNotCopied += dCopyStats.FileBytesNotCopied

}

type DirectoryCopyStats struct {
	DirsCreated         uint64
	TotalFilesProcessed uint64
	FilesCopied         uint64
	FileBytesCopied     uint64
	FilesNotCopied      uint64
	FileBytesNotCopied  uint64
	ComputeError        error
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
	StartPath                string
	Directories              DirMgrCollection
	ErrReturns               []error
	DeleteFileSelectCriteria FileSelectionCriteria
	DeletedFiles             FileMgrCollection
}

type DirectoryMoveStats struct {
	TotalSrcFilesProcessed   uint64
	SourceFilesMoved         uint64
	SourceFileBytesMoved     uint64
	SourceFilesRemaining     uint64
	SourceFileBytesRemaining uint64
	TotalDirsProcessed       uint64
	DirsCreated              uint64
	NumOfSubDirectories      uint64
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
