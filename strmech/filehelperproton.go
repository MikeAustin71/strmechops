package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"sync"
)

type fileHelperProton struct {
	lock *sync.Mutex
}

// addPathSeparatorToEndOfPathStr
//
// Receives a path string as an input parameter. If the
// last character of the path string is not a path
// separator, this method will add a path separator to
// the end of that path string and return it to the
// calling method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		The path string which will be analyzed to
//		determine if the last character is a path
//		separator.
//
//		If the last character is NOT a path separator,
//		this method will add a path separator to the end
//		of that path string and return it to the calling
//		method.
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
//	string
//
//		The path string passed as input parameter,
//		'pathStr' will be analyzed to determine if the
//		last character is a path separator.
//
//		If the last character is NOT a path separator,
//		a path separator will be added to 'pathStr' and
//		returned through this parameter.
//
//		If the last character IS a path separator, no
//		action will be taken and an exact copy of
//		'pathStr' will be returned through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fHelpProton *fileHelperProton) addPathSeparatorToEndOfPathStr(
	pathStr string,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if fHelpProton.lock == nil {
		fHelpProton.lock = new(sync.Mutex)
	}

	fHelpProton.lock.Lock()

	defer fHelpProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"fileHelperProton."+
			"addPathSeparatorToEndOfPathStr()",
		"")

	if err != nil {

		return "", err
	}

	var errCode, lStr int

	errCode,
		lStr,
		pathStr = new(fileHelperElectron).
		isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' is an empty string!\n",
			ePrefix.String())

		return "", err

	}

	if errCode == -2 {

		err = fmt.Errorf(
			"%v\n"+
				"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return "", err
	}

	if pathStr[lStr-1] == os.PathSeparator {
		return pathStr, nil
	}

	var newPathStr string

	if pathStr[lStr-1] == '/' && '/' != os.PathSeparator {

		newPathStr = pathStr[0 : lStr-1]

		newPathStr += string(os.PathSeparator)

		return newPathStr, nil
	}

	if pathStr[lStr-1] == '\\' && '\\' != os.PathSeparator {

		newPathStr = pathStr[0 : lStr-1]

		newPathStr += string(os.PathSeparator)

		return newPathStr, nil
	}

	newPathStr = pathStr + string(os.PathSeparator)

	return newPathStr, nil
}

// makeAbsolutePath
//
// Supply a relative path or any path string and resolve
// that path to an Absolute path. This method calls
// filepath.Abs() to generate the absolute path.
//
// "An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
// By contrast, a relative path starts from some given
// working directory, avoiding the need to provide the
// full absolute path. A filename can be considered as a
// relative path based at the current working directory.
// If the working directory is not the file's parent
// directory, a file not found error will result if the
// file is addressed by its name."
//
//	Wikipedia
//
// Note: Clean() is called on result by filepath.Abs().
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//	https://pkg.go.dev/path/filepath@go1.20.1#Abs
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	relPath						string
//
//		This string holds a relative path. This method
//		will convert this relative path to an absolute
//		path.
//
//		A relative path is defined as follows:
//
//		"A relative path starts from some given working
//		directory, avoiding the need to provide the full
//		absolute path. A filename can be considered as a
//		relative path based at the current working
//		directory. If the working directory is not the
//		file's parent directory, a file not found error
//		will result if the file is addressed by its name."
//			Wikipedia
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
//	string
//
//		If this method completes successfully, this
//		method will convert the relative path received
//		from input parameter '', to an absolute path.
//
//		"An absolute or full path points to the same
//		location in a file system, regardless of the
//		current working directory. To do that, it must
//		include the root directory."	Wikipedia
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

	fHelperElectron := fileHelperElectron{}

	errCode,
		_,
		relPath =
		fHelperElectron.
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
		fHelperElectron.isStringEmptyOrBlank(testRelPath)

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
			"testRelPath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			testRelPath,
			err2.Error())

		return "", err
	}

	return p, err
}

// openFileWithAccessCtrl
//
// This is a wrapper method for fileHelperBoson.openFile().
//
// This method receives an instance of FileAccessControl
// which is used to specify the File Permission codes,
// File Type and File Modes for the target file to be
// opened. It then proceeds to call method
// fileHelperBoson.openFile() which will 'open' the data
// file identified by input parameter 'pathFileName'.
//
// Once 'pathFileName' is opened, this method will return
// a file pointer to the opened file (filePtr *os.File).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The calling routine is responsible for calling
//	"Close()" on this os.File pointer.
//
//	Calling "Close()" multiple times on an os.File
//	pointer will generate a 'panic' error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		pathFileName					string
//
//			A string containing the path and file name of the
//			file which will be opened. If a parent path
//			component does NOT exist, this method will
//			trigger an error.
//
//		createDirectoryPathIfNotExist	bool
//
//			If the directory path element of parameter
//			'pathFileName' does not exist on an attached
//			storage drive, and this parameter is set to
//			'true', this method will attempt to create
//			the directory path.
//
//			If 'createDirectoryPathIfNotExist' is set to
//			'false', and the directory path element of
//			parameter 'pathFileName' does not exist on an
//			attached storage drive, an error will be returned.
//
//		fileAccessCtrl					FileAccessControl
//
//			A concrete instance of FileAccessControl. This
//			instance contains the three components used to
//			configure a file opening operation: File
//			Permission Codes, File Open Type and File Open
//			Modes. Two of these three components are required
//			while the third, File Open Modes, is optional.
//
//		1.	File Permission Codes
//
//			The File Permission Codes is a 10-character
//			string containing the read, write and execute
//			permissions for the three groups or user
//			classes:
//
//			(1)	'Owner/User'
//
//			(2)	'Group'
//
//			(3)	'Other'
//
//			This 10-character string will be used to
//			configure the internal File Permission data
//			field for a configured instance of
//			FileAccessControl.
//
//			The File Permission string must conform to the
//			symbolic notation options shown below. Failure
//			to comply with this requirement will generate
//			an error. As indicated, the File Permission
//			string must consist of 10-characters.
//
//			The first character in the File Permission
//			string may be dash ('-') specifying a file or a
//			'd' specifying a directory.
//
//			The remaining nine characters in the File
//			Permission string represent unix permission
//			bits and consist of three group fields each
//			containing 3-characters. Each character in
//			the three group fields may consist of 'r'
//			(Read-Permission), 'w' (Write-Permission),
//			'x' (Execute-Permission) or dash ('-')
//			signaling no permission or no access allowed.
//			A typical File Permission string authorizing
//			permission for full access to a file would be
//			styled as:
//
//				Example: "-rwxrwxrwx"
//
//			Groups:	-	Owner/User, Group, Other
//
//			From left to right
//			First Characters is Entry Type
//			-----------------------------------------------------
//			First Char index 0	=	"-"   Designates a file
//
//			First Char index 0	=	"d"   Designates a directory
//			-----------------------------------------------------
//
//			Char indexes 1-3	=	Owner "rwx" Authorizing 'Read',
//									Write' & Execute Permissions
//									for 'Owner'
//
//			Char indexes 4-6	= 	Group "rwx" Authorizing 'Read', 'Write' & Execute
//									Permissions for 'Group'
//
//			Char indexes 7-9	=	Other "rwx" Authorizing 'Read', 'Write' & Execute
//									Permissions for 'Other'
//
//			The Symbolic notation provided by input parameter
//			'filePermissionStr' MUST conform to the options
//			presented below. The first character or 'Entry Type'
//			is listed as "-". However, in practice, the caller
//			may set the first character as either a "-",
//			specifying a file, or a "d", specifying a directory.
//			No other first character types are currently
//			supported.
//
//			Three SymbolicGroups:
//
//			The three group types are: User/Owners, Groups & Others.
//
//			Directory Permissions:
//
//				-----------------------------------------------------
//				Directory Mode String Permission Codes
//				-----------------------------------------------------
//				Directory
//				10-Character
//				File Permission
//				String
//				Symbolic		  	Directory Access
//				Format	   		Permission Descriptions
//				----------------------------------------------------
//
//				d---------		no permissions
//				drwx------		read, write, & execute only for owner
//				drwxrwx---		read, write, & execute for owner and group
//				drwxrwxrwx		read, write, & execute for owner, group and others
//				d--x--x--x		execute
//				d-w--w--w-		write
//				d-wx-wx-wx		write & execute
//				dr--r--r--		read
//				dr-xr-xr-x		read & execute
//				drw-rw-rw-		read & write
//				drwxr-----		Owner can read, write, & execute. Group can only read;
//				others have no permissions
//
//				Note: drwxrwxrwx - identifies permissions for directory
//
//			File Permissions:
//
//			-----------------------------------------------------
//			File Mode String Permission Codes
//			-----------------------------------------------------
//
//				10-Character
//				File
//				Permission
//				String
//				Symbolic	Octal		Permission
//				Format		Notation	Descriptions
//				------------------------------------------------------------
//
//				----------	  0000		no permissions
//
//				-rwx------	  0700		read, write, & execute only for owner
//
//				-rwxrwx---	  0770		read, write, & execute for
//							  			owner and group
//
//				-rwxrwxrwx	  0777		read, write, & execute for owner,
//							  			group and others
//
//				---x--x--x	  0111		execute
//
//				--w--w--w-	  0222		write
//
//				--wx-wx-wx	  0333		write & execute
//
//				-r--r--r--	  0444		read
//
//				-r-xr-xr-x	  0555		read & execute
//
//				-rw-rw-rw-	  0666		read & write
//
//				-rwxr-----	  0740		Owner can read, write, &
//										execute. Group can only
//										read; others have no
//										permissions
//
//	 2.	File Open Type
//
//			A file open type is an enumeration specifying
//			the manner in which  the file will be opened.
//			In order to open a file, exactly one of the
//			following File Open Codes MUST be specified:
//
//			FileOpenType(0).TypeReadOnly()
//			FileOpenType(0).TypeWriteOnly()
//			FileOpenType(0).TypeReadWrite()
//
//			-- AND --
//
//	 3.	File Open Mode
//
//	    In addition to a File Open Type, a File Open
//	    Mode may be specified. Zero or more of the
//	    following File Open Modes may optionally be
//	    specified to achieve granular control over
//	    file open behavior.
//
//	    FileOpenMode(0).ModeAppend()
//	    FileOpenMode(0).ModeCreate()
//	    FileOpenMode(0).ModeExclusive()
//	    FileOpenMode(0).ModeSync()
//	    FileOpenMode(0).ModeTruncate()
//
//		pathFileNameLabel				string
//
//			The name or label associated with input parameter
//			'pathFileName' which will be used in error
//			messages returned by this method.
//
//			If this parameter is submitted as an empty
//			string, a default value of "pathFileName" will be
//			automatically applied.
//
//		errPrefDto						*ePref.ErrPrefixDto
//
//			This object encapsulates an error prefix string
//			which is included in all returned error
//			messages. Usually, it contains the name of the
//			calling method or methods listed as a function
//			chain.
//
//			If no error prefix information is needed, set
//			this parameter to 'nil'.
//
//			Type ErrPrefixDto is included in the 'errpref'
//			software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the file designated by input parameter
//		'pathFileName'. This file pointer can
//		subsequently be used for reading content from the
//		subject file. It may NOT be used for writing
//		content to the subject file.
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelpProton *fileHelperProton) openFileWithAccessCtrl(
	pathFileName string,
	createDirectoryPathIfNotExist bool,
	fileAccessCtrl FileAccessControl,
	pathFileNameLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	filePtr *os.File,
	err error) {

	if fHelpProton.lock == nil {
		fHelpProton.lock = new(sync.Mutex)
	}

	fHelpProton.lock.Lock()

	defer fHelpProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"fileHelperProton."+
			"openFileWithAccessCtrl()",
		"")

	if err != nil {

		return filePtr, err
	}

	if len(pathFileNameLabel) == 0 {

		pathFileNameLabel = "pathFileName"
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			&fileAccessCtrl,
			ePrefix.XCpy("fileAccessCtrl"))

	if err != nil {

		return filePtr, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = fileAccessCtrl.GetFileOpenConfig(
		ePrefix.XCpy("fileOpenCfg<-"))

	if err != nil {

		return filePtr, err
	}

	var filePermissionCfg FilePermissionConfig

	filePermissionCfg,
		err = fileAccessCtrl.GetFilePermissionConfig(
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return filePtr, err
	}

	filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			createDirectoryPathIfNotExist,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	return filePtr, err
}
