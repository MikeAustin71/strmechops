package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type dirMgrHelperBoson struct {
	lock *sync.Mutex
}

// isDirMgrValid
//
// This method examines the current DirMgr object to
// determine whether it has been properly configured. If
// the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error
// is returned.
//
// In order to qualify as a valid DirMgr instance the
// DirMgr.path and DirMgr.absolutePath member string
// variables must be populated with alphanumeric
// characters.
func (dMgrHlprBoson *dirMgrHelperBoson) isDirMgrValid(
	dMgr *DirMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprBoson.lock == nil {
		dMgrHlprBoson.lock = new(sync.Mutex)
	}

	dMgrHlprBoson.lock.Lock()

	defer dMgrHlprBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"DirMgr.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	if dMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dirMgr' a 'nil' pointer!\n",
			ePrefix.String())
	}

	fh := new(FileHelper)

	errCode := 0

	errCode,
		_,
		dMgr.path =
		fh.IsStringEmptyOrBlank(dMgr.path)

	if errCode == -1 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'DirMgr'.path is an empty string!\n",
			ePrefix.String())
	}

	if errCode == -2 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'DirMgr'.path consists of blank spaces!\n",
			ePrefix.String())
	}

	errCode,
		_,
		dMgr.absolutePath =
		fh.IsStringEmptyOrBlank(dMgr.absolutePath)

	if errCode == -1 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.absolutePath = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'DirMgr'.absolutePath is an empty string!\n",
			ePrefix.String())
	}

	if errCode == -2 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.absolutePath = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'DirMgr'.absolutePath consists of blank spaces!\n",
			ePrefix.String())
	}

	return err
}
