package strmech

import (
	"strings"
	"sync"
	"time"
)

// FileSelectionCriteria - Used is selecting file names. These
// data fields specify the criterion used to determine if a
// file should be selected for some type of operation.
// Example: find files or delete files operations
type FileSelectionCriteria struct {
	// FileNamePatterns - a string array containing one or more file matching
	// patterns. Example '*.txt' '*.log' 'common*.*'
	FileNamePatterns []string

	// FilesOlderThan - Used to select files with a modification less than this date time
	FilesOlderThan time.Time

	// FilesNewerThan - // Used to select files with a modification greater than this date time
	FilesNewerThan time.Time

	// SelectByFileMode - Used to select files with equivalent os.FileMode values.
	// To select by File Mode, set the FilePermissionCfg type to the desired value
	//  Examples:
	//    fsc := FileSelectionCriteria{}
	//
	//    err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
	//    err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
	//
	// Note: os.FileMode is an uint32 type.
	SelectByFileMode FilePermissionConfig

	// SelectCriterionMode - Can be one of three values:
	//
	// FileSelectMode.None()      = No Operation - No File Select Criterion
	//                                   mode selected
	//
	// FileSelectMode.ANDSelect() = select a file only if ALL
	//                                   the selection criterion are satisfied.
	//
	// FileSelectMode.ORSelect()  = select a file if only ONE
	//                                   of the selection criterion are satisfied.
	//
	// SEE TYPE 'FileSelectCriterionMode'
	SelectCriterionMode FileSelectCriterionMode

	lock *sync.Mutex
}

// ArePatternsActive - surveys the FileNamePatterns string
// array to determine if there are currently any active search
// file pattern string.
//
// A search file pattern is considered active if the string
// length of the pattern string is greater than zero.
func (fsc *FileSelectionCriteria) ArePatternsActive() bool {

	if fsc.lock == nil {
		fsc.lock = new(sync.Mutex)
	}

	fsc.lock.Lock()

	defer fsc.lock.Unlock()

	lPats := len(fsc.FileNamePatterns)

	if lPats == 0 {
		return false
	}

	isActive := false

	for i := 0; i < lPats; i++ {
		fsc.FileNamePatterns[i] =
			strings.TrimRight(strings.TrimLeft(fsc.FileNamePatterns[i], " "), " ")
		if fsc.FileNamePatterns[i] != "" {
			isActive = true
		}

	}

	return isActive
}

// IsSelectionCriteriaActive
//
// This method returns a boolean value signaling whether
// the current instance of FileSelectionCriteria is active
// and engaged.
//
// If the returned boolean value is set to 'true', means
// that the File Selection Criterion are configured to
// select a subset of available files.
//
// By contrast a return value of 'false' signals that no
// file selection criteria have been configured meaning
// that all files will be selected.
func (fsc *FileSelectionCriteria) IsSelectionCriteriaActive() bool {

	if fsc.lock == nil {
		fsc.lock = new(sync.Mutex)
	}

	fsc.lock.Lock()

	defer fsc.lock.Unlock()

	namePatternIsActive := false

	lPats := len(fsc.FileNamePatterns)

	var tempPattern string

	for i := 0; i < lPats; i++ {

		tempPattern =
			strings.TrimRight(
				strings.TrimLeft(fsc.FileNamePatterns[i], " "), " ")

		if tempPattern != "" {
			namePatternIsActive = true
		}
	}

	if namePatternIsActive == true {

		return true
	}

	if !fsc.FilesNewerThan.IsZero() {

		return true
	}

	if !fsc.FilesOlderThan.IsZero() {

		return true
	}

	return false
}
