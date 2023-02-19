package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type fileHelperDirector struct {
	lock *sync.Mutex
}

// CopyFileByIoByLink
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a new destination file and
// using "io.Copy(out, in)" to copy the contents. This is
// accomplished by calling 'CopyFileByIo()'.
// If  the call to 'CopyFileByIo()' fails, this method
// will attempt a second copy method.
//
// The second attempt to copy the designated file will be
// accomplished by creating a 'hard link' to the source
// file. The second, 'hard link', attempt will call
// method, 'CopyFileByLink()'.
//
// If both attempted file copy operations fail, an error
// will be returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelpDirector *fileHelperDirector) copyFileByIoByLink(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"copyFileByIoByLink()",
		"")

	if err != nil {
		return err
	}

	fHelperMech := fileHelperMechanics{}

	err = fHelperMech.copyFileByIo(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	// fh.CopyFileByIo() failed. Try
	// fh.CopyFileByLink()

	var err2 error

	err2 = fHelperMech.
		copyFileByLink(
			src,
			dst,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By IO failed, an error was returned\n"+
			"by CopyFileByLink(src, dst)\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			src,
			dst,
			err2.Error())

	}

	return err
}

// copyFileByLinkByIo
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a 'hard link' to the source file.
// The 'hard link' attempt will call 'FileHelper.CopyFileByLink()'.
//
// If that 'hard link' operation fails, this method will call
// 'CopyFileByIo()'.
//
// CopyFileByIo() will create a new destination file and attempt
// to write the contents of the source file to the new destination
// file using "io.Copy(out, in)".
//
// If both attempted file copy operations fail, an error will be
// returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelpDirector *fileHelperDirector) copyFileByLinkByIo(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"copyFileByLinkByIo()",
		"")

	if err != nil {
		return err
	}

	fHelperMech := new(fileHelperMechanics)

	err = fHelperMech.copyFileByLink(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	var err2 error

	// Copy by Link Failed. Try CopyFileByIo()
	err2 = fHelperMech.copyFileByIo(
		src,
		dst,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link failed, an error was returned by fh.CopyFileByIo(src, dst).\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix,
			src,
			dst,
			err2.Error())

		return err
	}

	return err
}

// openDirectory
//
// Opens a directory and returns the associated 'os.File'
// pointer. This method will open a directory designated
// by input parameter, 'directoryPath'.
//
// The input parameter 'createDir' determines the action
// taken if 'directoryPath' does not exist. If
// 'createDir' is set to 'true' and 'directoryPath' does
// not currently exist, this method will attempt to
// create 'directoryPath'. Directories created in this
// manner are configured with Open Type of 'Read-Write'
// and a Permission code of 'drwxrwxrwx'.
//
// Alternatively, if 'createDir' is set to 'false' and
// 'directoryPath' does NOT exist, an error will be
// returned.
//
// Regardless of whether the target directory path
// already exists or is created by this method, the
// returned os.File pointer is opened with the
// 'Read-Only' attribute (O_RDONLY) and a
// permission code of zero ("----------").
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The caller is responsible for calling "Close()" on the
// returned os.File pointer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		A string containing the path name of the
//		directory which will be opened.
//
//	createDir					bool
//
//		Determines what action will be taken if
//		'directoryPath' does NOT exist. If 'createDir' is
//		set to 'true' and 'directoryPath' does NOT exist,
//		this method will attempt to create
//		'directoryPath'. Alternatively, if 'createDir' is
//		set to false and 'directoryPath' does NOT exist,
//		this method will terminate and an error will be
//		returned.
//
//		Directories created in this manner will have an
//		Open Type of 'Read-Write' and a Permission code
//		of 'drwxrwxrwx'. This differs from the Open Type
//		and permission mode represented by the returned
//		os.File pointer.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the directory designated by input
//		parameter 'directoryPath'.
//
//		If successful, the returned os.File pointer is
//		opened with the 'Read-Only' attribute (O_RDONLY)
//		and a permission code of zero ("----------").
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//			The caller is responsible for calling
//			"Close()" on this os.File pointer.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelpDirector *fileHelperDirector) openDirectory(
	directoryPath string,
	createDir bool,
	errPrefDto *ePref.ErrPrefixDto) (
	*os.File,
	error) {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"openDirectory()",
		"")

	if err != nil {
		return nil, err
	}

	var directoryPathDoesExist bool
	var dirPathFInfo FileInfoPlus

	fHelpMolecule := fileHelperMolecule{}

	directoryPath,
		directoryPathDoesExist,
		dirPathFInfo,
		err = fHelpMolecule.doesPathFileExist(
		directoryPath,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"directoryPath")

	if err != nil {
		return nil, err
	}

	if directoryPathDoesExist &&
		!dirPathFInfo.IsDir() {

		return nil,
			fmt.Errorf("%v\n"+
				"ERROR: 'directoryPath' does exist, but\n"+
				"IT IS NOT A DIRECTORY!\n"+
				"directoryPath='%v'\n",
				ePrefix.String(),
				directoryPath)
	}

	if directoryPathDoesExist && dirPathFInfo.Mode().IsRegular() {
		return nil,
			fmt.Errorf("%v\n"+
				"ERROR: 'directoryPath' does exist, but\n"+
				"it is classifed as a REGULAR File!\n"+
				"directoryPath='%v'\n",
				ePrefix.String(),
				directoryPath)
	}

	if !directoryPathDoesExist {

		if !createDir {
			return nil,
				fmt.Errorf("%v\n"+
					"Error 'directoryPath' DOES NOT EXIST!\n"+
					"directoryPath='%v'\n",
					ePrefix.String(),
					directoryPath)
		}

		// Parameter 'createDir' must be 'true'.
		// The error signaled that the path does not exist. So, create the directory path
		err = new(fileHelperMechanics).makeDirAll(
			directoryPath,
			ePrefix)

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Attmpted creation of 'directoryPath' FAILED!\n"+
					"directoryPath='%v'\n"+
					"Error=\n%v\n",
					ePrefix.String(),
					directoryPath,
					err.Error())
		}

		// Verify that the directory exists and get
		// the associated file info object.
		_,
			directoryPathDoesExist,
			dirPathFInfo,
			err = fHelpMolecule.doesPathFileExist(
			directoryPath,
			PreProcPathCode.None(), // Take No Pre-Processing Action
			ePrefix,
			"directoryPath")

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"Error occurred verifying existance of "+
					"newly created 'directoryPath'!\n"+
					"Non-Path error returned by os.Stat(directoryPath)\n"+
					"directoryPath='%v'\n"+
					"Error=\n%v\n",
					ePrefix.String(),
					directoryPath,
					err.Error())
		}

		if !directoryPathDoesExist {
			return nil, fmt.Errorf("%v\n"+
				"Error: Verification of newly created "+
				"directoryPath FAILED!\n"+
				"'directoryPath' DOES NOT EXIST!\n"+
				"directoryPath='%v'\n",
				ePrefix.String(),
				directoryPath)
		}

		if !dirPathFInfo.IsDir() {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Input Paramter 'directoryPath' is NOT a directory!\n"+
					"directoryPath='%v'\n",
					ePrefix.String(),
					directoryPath)
		}

		if dirPathFInfo.Mode().IsRegular() {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: 'directoryPath' does exist, but\n"+
					"it is classifed as a REGULAR File!\n"+
					"directoryPath='%v'\n",
					ePrefix.String(),
					directoryPath)
		}
	}

	filePtr, err := os.Open(directoryPath)

	if err != nil {
		return nil,
			fmt.Errorf("%v\n"+
				"directoryPath='%v'\n"+
				"File Open Error: %v\n",
				ePrefix.String(),
				directoryPath,
				err.Error())
	}

	if filePtr == nil {
		return nil, fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n",
			ePrefix.String())
	}

	return filePtr, nil
}
