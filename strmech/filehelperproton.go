package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	fp "path/filepath"
	"sync"
)

type fileHelperProton struct {
	lock *sync.Mutex
}

// makeAbsolutePath - Supply a relative path or any path
// string and resolve that path to an Absolute path.
// Note: Clean() is called on result by fp.Abs().
func (fHelpProton *fileHelperProton) makeAbsolutePath(
	relPath string,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if fHelpProton.lock == nil {
		fHelpProton.lock = new(sync.Mutex)
	}

	fHelpProton.lock.Lock()

	defer fHelpProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperProton."+
			"makeAbsolutePath()",
		"")

	if err != nil {
		return "", err
	}

	errCode := 0

	errCode, _, relPath =
		new(fileHelperElectron).
			isStringEmptyOrBlank(relPath)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'relPath' is an empty string!\n",
			ePrefix.String())

		return "", err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'relPath' consists of blank spaces!\n",
			ePrefix.String())

		return "", err
	}

	testRelPath :=
		new(fileHelperAtom).adjustPathSlash(relPath)

	errCode, _, testRelPath =
		new(fileHelperElectron).isStringEmptyOrBlank(testRelPath)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'relPath' adjusted for path Separators is an EMPTY string!\n",
			ePrefix.String())

		return "", err
	}

	var err2 error
	var p string

	p, err2 = fp.Abs(testRelPath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fp.Abs(testRelPath).\n"+
			"testRelPath='%v'\nError='%v'\n",
			ePrefix.String(),
			testRelPath,
			err2.Error())

		return "", err
	}

	return p, err
}
