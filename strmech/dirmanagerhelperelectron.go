package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
	"time"
)

type dirMgrHelperElectron struct {
	lock *sync.Mutex
}

// lowLevelDoesDirectoryExist
//
// This method tests for the existence of directory path.
func (dMgrHlprElectron *dirMgrHelperElectron) lowLevelDoesDirectoryExist(
	dirPath,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirPathDoesExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	dirPathDoesExist = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperElectron."+
			"lowLevelDoesDirectoryExist()",
		"")

	if err != nil {
		return dirPathDoesExist, fInfoPlus, err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "DirMgr"
	}

	errCode := 0

	errCode,
		_,
		dirPath = new(FileHelper).
		IsStringEmptyOrBlank(dirPath)

	if errCode < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input paramter %v is an empty string!\n",
			ePrefix.String(),
			dirPathLabel)

		return dirPathDoesExist, fInfoPlus, err
	}

	var err2 error
	var info os.FileInfo

	for i := 0; i < 3; i++ {

		dirPathDoesExist = false
		fInfoPlus = FileInfoPlus{}
		err = nil

		info,
			err2 = os.Stat(dirPath)

		if err2 != nil {

			if os.IsNotExist(err2) {

				dirPathDoesExist = false
				fInfoPlus = FileInfoPlus{}
				err = nil

				return dirPathDoesExist, fInfoPlus, err
			}

			// err == nil and err != os.IsNotExist(err)
			// This is a non-path error. The non-path error will be
			// tested up to 3-times before it is returned.
			err = fmt.Errorf("%v\n"+
				"Non-Path error returned by os.Stat(%v)\n"+
				"%v= %v\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirPathLabel,
				dirPathLabel,
				dirPath,
				err2.Error())

			fInfoPlus = FileInfoPlus{}

			dirPathDoesExist = false

		} else {
			// err == nil
			// The path really does exist!
			dirPathDoesExist = true
			err = nil

			fInfoPlus,
				err2 = new(FileInfoPlus).
				NewFromPathFileInfo(dirPath, info)

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					err2.Error())

				fInfoPlus = FileInfoPlus{}
			}

			return dirPathDoesExist, fInfoPlus, err
		}

		time.Sleep(30 * time.Millisecond)
	}

	return dirPathDoesExist, fInfoPlus, err
}
