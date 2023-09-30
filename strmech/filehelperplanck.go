package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type fileHelperPlanck struct {
	lock *sync.Mutex
}

// makeDirAllPerm
//
// Creates a directory path along with any necessary
// parent paths and configures permissions for the
// new directory path.
//
// If the target directory path already exists, this
// method does nothing and returns.
//
// The input parameter 'permission' is of type
// 'FilePermissionConfig'. See method the documentation
// for method 'FilePermissionConfig.New()' for an
// explanation of permission codes.
//
// If you wish to grant total access to a directory,
// consider setting permission code as follows:
//
//	FilePermissionConfig{}.New("drwxrwxrwx")
//
// If the parent directories in parameter 'dirPath' do
// not yet exist, this method will create them.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		dirPath						string
//
//			This string contains the directory path which
//			will be created by this method.
//
//		permission					FilePermissionConfig
//
//			An instance of FilePermissionConfig containing
//			the permission specifications for the new
//			directory to be created from input paramter,
//			'dirPath'.
//
//			The easiest way to configure permissions is
//			to call FilePermissionConfig.New() with
//			a mode string ('modeStr').
//
//			The first character of the 'modeStr' designates the
//			'Entry Type'. Currently, only two 'Entry Type'
//			characters are supported. Therefore, the first
//			character in the 10-character input parameter
//			'modeStr' MUST be either a "-" indicating a file, or
//			a "d" indicating a directory.
//
//			The remaining nine characters in the 'modeStr'
//			represent unix permission bits and consist of three
//			group fields each containing 3-characters. Each
//			character in the three group fields may consist of
//			'r' (Read-Permission), 'w' (Write-Permission), 'x'
//			(Execute-Permission) or '-' signaling no permission or
//			no access allowed. A typical 'modeStr' authorizing
//			permission for full access to a file would be styled
//			as:
//
//			Directory Example: "drwxrwxrwx"
//
//			Groups: - Owner/User, Group, Other
//			From left to right
//			First Characters is Entry Type index 0 ("-")
//
//			First Char index 0 =     "-"   Designates a file
//
//			First Char index 0 =     "d"   Designates a directory
//
//			Char indexes 1-3 = Owner "rwx" Authorizing 'Read',
//		                                  Write' & Execute Permissions for 'Owner'
//
//			Char indexes 4-6 = Group "rwx" Authorizing 'Read', 'Write' & Execute
//		                                  Permissions for 'Group'
//
//			Char indexes 7-9 = Other "rwx" Authorizing 'Read', 'Write' & Execute
//		                                  Permissions for 'Other'
//
//	        -----------------------------------------------------
//	               Directory Mode String Permission Codes
//	        -----------------------------------------------------
//	          Directory
//				10-Character
//				 'modeStr'
//				 Symbolic		  Directory Access
//				  Format	   Permission Descriptions
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
//				                others have no permissions
//
//				Note: drwxrwxrwx - identifies permissions for directory
//
//
//	dirPathLabel				string
//
//		The name or label associated with input parameter
//		'dirPath' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dirPath" will be
//		automatically applied.
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
func (fHelperPlanck *fileHelperPlanck) makeDirAllPerm(
	dirPath string,
	permission FilePermissionConfig,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelperPlanck.lock == nil {
		fHelperPlanck.lock = new(sync.Mutex)
	}

	fHelperPlanck.lock.Lock()

	defer fHelperPlanck.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "fileHelperPlanck." +
		"makeDirAllPerm()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "dirPath"
	}

	errCode := 0

	errCode,
		_,
		dirPath = new(fileHelperElectron).
		isStringEmptyOrBlank(dirPath)

	if errCode == -1 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is an empty string!\n",
			ePrefix.String(),
			dirPathLabel)
	}

	if errCode == -2 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' consists of blank spaces!\n",
			ePrefix.String(),
			dirPathLabel)
	}

	var err2 error

	err2 = permission.IsValidInstanceError(
		ePrefix.XCpy("permission"))

	if err2 != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'permission' is INVALID!\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	var dirPermCode os.FileMode

	dirPermCode,
		err2 = permission.
		GetCompositePermissionMode(
			ePrefix.XCpy(
				"permission"))

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"ERROR: INVALID Permission Code!\n"+
			"Error returned by permission.GetCompositePermissionMode()\n"+
			"Error=\n%v\n",
			funcName,
			err2.Error())
	}

	dirPath,
		err2 = new(fileHelperProton).
		makeAbsolutePath(
			dirPath,
			ePrefix.XCpy("dirPath"))

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath(dirPath).\n"+
			"dirPath='%v'\n"+
			"Error=\n%v\n",
			funcName,
			dirPath,
			err2.Error())
	}

	err2 = os.MkdirAll(dirPath, dirPermCode)

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"Error return from os.MkdirAll(%v, permission).\n"+
			"%v= '%v'\n"+
			"Permission Code ('dirPermCode') = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dirPathLabel,
			dirPathLabel,
			dirPath,
			dirPermCode,
			err2.Error())
	}

	var pathDoesExist bool

	_,
		pathDoesExist,
		_,
		err2 = new(fileHelperMolecule).doesPathFileExist(
		dirPath,
		PreProcPathCode.None(), // Take no Pre-Processing Action
		ePrefix,
		dirPathLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"After creating the %v directory, an error\n"+
			"was returned by fileHelperMolecule.doesPathFileExist()\n"+
			"Error=\n%v\n",
			funcName,
			dirPathLabel,
			err2.Error())

		return err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dirPathLabel,
			dirPath)
	}

	return err
}
