package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strconv"
	"sync"
)

type filePermissionConfigNanobot struct {
	lock *sync.Mutex
}

// setFileModeByComponents
//
// Sets the value of the current FilePermissionConfig
// instance by initializing the internal FileMode data
// field (FilePermissionConfig.fileMode). The final
// FileMode value is computed by integrating the
// 'entryType' FileMode with the unix permission symbolic
// values provided by the input parameter,
// 'unixPermissionStr'. This approach allows the user
// to created custom File Permissions.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// Incorrect or invalid File Permissions can cause
// extensive damage. If you don't know what you are
// doing, you would be well advised to use one of the
// other methods in this type which provide additional
// safeguards.
//
// If you decide to proceed, be guided by the wisdom of
// Davy Crockett:
//
//	"Be always sure you are right - then go ahead."
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		entryType					OsFilePermissionCode
//
//			The code which makes up the first character in a
//			10-digit unix permission character string.
//
//			This a wrapper for os.FileMode constants.
//			Reference:
//				https://golang.org/pkg/os/#FileMode
//
//			Select this value with caution.
//			See the warning above.
//
//		unixPermissionStr			string
//
//			A 9-character string containing the unix
//			permission bits expressed as three groups of
//		 	3-characters each.
//
//			The 9-characters are constituents of the three
//			Symbolic Groups or User Classes:
//				(1) Owners/Users
//				(2) Groups
//				(3) Others.
//			Each group has three characters which may be 'r',
//			'w', 'x'. If a permission is not set, that
//			character position contains a '-'.
//
//		 	Unix Permission String Options:
//
//		      9-Character          File Access
//		      Notation             Permission Descriptions
//		-----------------------------------------------------------
//		      ---------            File - no permissions
//		      rwx------            File - read, write, & execute only for owner
//		      rwxrwx---            File - read, write, & execute for owner and group
//		      rwxrwxrwx            File - read, write, & execute for owner, group and others
//		      --x--x--x            File - execute
//		      -w--w--w-            File - write
//		      -wx-wx-wx            File - write & execute
//		      r--r--r--            File - read
//		      r-xr-xr-x            File - read & execute
//		      rw-rw-rw-            File - read & write
//		      rwxr-----            File - Owner can read, write, & execute. Group can only read;
//
//
//			Note: drwxrwxrwx - identifies permissions for
//	                        directory
//
//
//		errPrefDto					*ePref.ErrPrefixDto
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
func (fPermConfigNanobot *filePermissionConfigNanobot) setFileModeByComponents(
	fPerm *FilePermissionConfig,
	entryType OsFilePermissionCode,
	unixPermissionTextStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fPermConfigNanobot.lock == nil {
		fPermConfigNanobot.lock = new(sync.Mutex)
	}

	fPermConfigNanobot.lock.Lock()

	defer fPermConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigNanobot."+
			"setFileModeByComponents()",
		"")

	if err != nil {
		return err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(unixPermissionTextStr) == 10 {
		unixPermissionTextStr = unixPermissionTextStr[1:]
	}

	if len(unixPermissionTextStr) != 9 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'unixPermissionTextStr' must contain 9-Characters.\n"+
			"This unixPermissionTextStr contains %v-characters.\n"+
			"unixPermissionTextStr '%v'\n",
			ePrefix.String(),
			len(unixPermissionTextStr),
			unixPermissionTextStr)

		return err
	}

	fModeEntryType := os.FileMode(entryType)

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	_, ok := mOsPermissionCodeToString[fModeEntryType]

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Input parameter 'entryType' is an INVALID os.FileMode!\n"+
			"entryType decimal value='%s' \n"+
			"octal value='%s' \n",
			ePrefix.String(),
			strconv.FormatInt(int64(entryType), 10),
			strconv.FormatInt(int64(entryType), 8))

		return err
	}

	fPermCfgElectron := filePermissionConfigElectron{}

	var ownerInt, groupInt, otherInt int

	ownerInt,
		err = fPermCfgElectron.
		convertGroupToDecimal(
			unixPermissionTextStr[0:3],
			"owner",
			ePrefix.XCpy(
				"ownerInt<-unixPermissionTextStr[0:3]"))

	if err != nil {
		return err
	}

	groupInt,
		err = fPermCfgElectron.
		convertGroupToDecimal(
			unixPermissionTextStr[3:6],
			"group",
			ePrefix.XCpy("groupInt<-"+
				"unixPermissionTextStr[3:6]"))

	if err != nil {
		return err
	}

	otherInt,
		err = fPermCfgElectron.
		convertGroupToDecimal(
			unixPermissionTextStr[6:],
			"other",
			ePrefix.XCpy("otherInt<-"+
				"unixPermissionTextStr[6:]"))

	if err != nil {
		return err
	}

	ownerInt *= 100
	groupInt *= 10
	permission := ownerInt + groupInt + otherInt

	fMode := os.FileMode(
		new(NumberConversions).
			ConvertOctalToDecimal(permission))

	fPerm.fileMode = fModeEntryType | fMode
	fPerm.isInitialized = true

	return err
}

// setByFileMode
//
// Sets the permission codes for the FilePermissionConfig
// instance passed as input parameter 'fPerm'.
//
// Using input parameter 'fMode' of type 'os.FileMode'.
// If the value does not include a valid os mode
// constant, an error will be returned.
//
// If successful, this method will assign the os.FileMode
// input value to the internal data field,
// 'fPerm.fileMode'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#FileMode
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		The internal FileMode data field for this
//		instance will be reset using the permission codes
//		contained in input parameter 'modeStr'.
//
//	fMode						os.FileMode
//
//		An instance of os.FileMode containing file or
//		directory permission codes. These permission
//		codes will be used to reset the internal
//		FileMode data field in the 'fPerm' instance of
//		FilePermissionConfig.
//
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
func (fPermConfigNanobot *filePermissionConfigNanobot) setByFileMode(
	fPerm *FilePermissionConfig,
	fMode os.FileMode,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fPermConfigNanobot.lock == nil {
		fPermConfigNanobot.lock = new(sync.Mutex)
	}

	fPermConfigNanobot.lock.Lock()

	defer fPermConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigNanobot."+
			"setByFileMode()",
		"")

	if err != nil {
		return err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	tFMode := fMode

	mask := os.FileMode(0777)

	// The &^ operator is bit clear (AND NOT):
	// in the expression z = x &^ y, each bit of z is 0
	// if the corresponding bit of y is 1; otherwise it
	// equals the corresponding bit of x
	entryType := tFMode &^ mask

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	_, ok := mOsPermissionCodeToString[entryType]

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMode' contains an invalid\n"+
			"'EntryType' otherwise known as an os mode constant.\n"+
			"'fMode' value  = %v\n"+
			"'fMode' string = %v\n",
			ePrefix.String(),
			fMode,
			fMode.String())

		return err
	}

	fPerm.fileMode = fMode
	fPerm.isInitialized = true

	return nil
}

// setFileModeByOctalDigits
//
// Sets the permissions value of the FilePermissionConfig
// instance passed as input parameter 'fPerm'. The
// internal FileMode data field (fPerm.fileMode) is reset
// to the value represented by input parameter,
// 'octalFileModeCode'. Any previous internal FileMode
// value is overwritten.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and
// converted to a decimal value. Therefore, x:= int(0777)
// will mean that 'x' is set equal to 511. If you set
// x:= int(777), x will be set equal to '777'. For purposes
// of this method enter the octal code as x:= int(777).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		The internal FileMode data field for this
//		instance will be reset using the permission codes
//		contained in input parameter 'modeStr'.
//
//	octalFileModeCode			int
//
//		This parameter contains the integer value of the
//		of the permission code which will be used to
//		initialize the 'fPerm' FilePermissionConfig
//		instance (fPerm.fileMode). The integer digits in
//		'octalFileModeCode' represent the octal value for
//		the file permissions.
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
//
//		A partial list of valid file permission value
//	 	examples are shown as follows:
//
//	 ____________________________________________________________________________
//
//	          Input Parameter
//	              integer            Equivalent
//	 Octal    'octalFileModeCode'    Symbolic      File Access
//	 Digits        value             Notation      Permission Descriptions
//	 0000 	         0               ----------    File - no permissions
//	 0700 	       700               -rwx------    File - read, write, & execute only for owner
//	 0770 	       770               -rwxrwx---    File - read, write, & execute for owner and group
//	 0777 	       777               -rwxrwxrwx    File - read, write, & execute for owner, group and others
//	 0111 	       111               ---x--x--x    File - execute
//	 0222 	       222               --w--w--w-    File - write
//	 0333 	       333               --wx-wx-wx    File - write & execute
//	 0444 	       444               -r--r--r--    File - read
//	 0555 	       555               -r-xr-xr-x    File - read & execute
//	 0666 	       666               -rw-rw-rw-    File - read & write
//	 0740 	       740               -rwxr-----    File - Owner can read, write, & execute. Group can only read;
//	                                                      others have no permissions
//
//	 drwxrwxrwx    Directory - read, write, & execute for owner, group and others
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
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
func (fPermConfigNanobot *filePermissionConfigNanobot) setFileModeByOctalDigits(
	fPerm *FilePermissionConfig,
	octalFileModeCode int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fPermConfigNanobot.lock == nil {
		fPermConfigNanobot.lock = new(sync.Mutex)
	}

	fPermConfigNanobot.lock.Lock()

	defer fPermConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigNanobot."+
			"setFileModeByOctalDigits()",
		"")

	if err != nil {
		return err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	decimalVal := new(NumberConversions).ConvertOctalToDecimal(octalFileModeCode)

	tFMode := os.FileMode(decimalVal)

	mask := os.FileMode(0777)

	entryType := tFMode &^ mask

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	_, ok := mOsPermissionCodeToString[entryType]

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'octalFileModeCode' contains an invalid\n"+
			"'EntryType' otherwise known as an os mode constant.\n"+
			"entryType decimal value='%s' \n"+
			"entryType octal value='%s' \n"+
			"'octalFileModeCode' decimal value  = %s\n"+
			"'octalFileModeCode' decimal value  = %s\n",
			ePrefix.String(),
			strconv.FormatInt(int64(entryType), 10),
			strconv.FormatInt(int64(entryType), 8),
			strconv.FormatInt(int64(octalFileModeCode), 10),
			strconv.FormatInt(int64(octalFileModeCode), 8))

		return err
	}

	fPerm.fileMode = tFMode
	fPerm.isInitialized = true

	return nil
}

// setFileModeByTextCode
//
// Sets the internal FileMode data field using input
// parameter 'modeStr'. 'modeStr' is a 10-character
// string containing the read, write and execute
// permissions for the three user classes, 'Owner',
// 'Group' and 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix permission codes.
//
//	Reference:
//
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//	https://en.wikipedia.org/wiki/File_system_permissions
//	https://www.redhat.com/sysadmin/linux-file-permissions-explained
//
// The first character of the 'modeStr' designates the
// 'Entry Type'. Currently, only two 'Entry Type'
// characters are supported. Therefore, the first
// character in 'modeStr' must consist of a hyphen ("-")
// designating a file, or a "d" designating a directory.
//
// The remaining nine characters in the 'modeStr' are
// styled as unix permission bits. These nine characters
// are divided into three group fields each containing
// 3-permission characters. Each character field may be
// populated with a 'r' (Read-Permission), 'w'
// (Write-Permission), 'x' (Execute-Permission) or '-'
// signaling no permission or no access allowed. A
// typical 'modeStr' authorizing permission for full
// access to a file would be styled as:
//
//					"-rwxrwxrwx"
//
//		User Classes: Owner, Group, Other
//	 Note: Owner is also referred to as User/Owner meaning
//	       the User who owns the file.
//
//		From left to right
//
//		Char index 0     = Entry Type. Must be either a "-" or a "d"
//		Char indexes 1-3 = Owner  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Owner'
//		Char indexes 4-6 = Group  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Group'
//		Char indexes 7-9 = Other  "rwx"  Authorizing 'Read', 'Write' & Execute Permissions for 'Other'
//
// The Symbolic notation provided by input parameter
// 'modeStr' MUST conform to the options presented below.
// The first character or 'Entry Type' is listed as "-".
// However, in practice, the caller may set the first
// character as either a "-", specifying a file, or a
// "d", specifying a directory. No other first character
// types are currently supported.
//
// After the first character, the remaining 9-characters
// are constituents of the three Symbolic Groups or user
// classes: Owners, Users, Groups & Others. Each user
// class has three characters which may be 'r', 'w', 'x'.
// If a permission is not set, the character position
// contains a '-'.
//
//		'modeStr'
//		Symbolic    Octal           File Access
//		Format     Notation        Permission Descriptions
//		------------------------------------------------------------
//
//		----------   0000           File - no permissions
//		-rwx------   0700           File - read, write, & execute only for owner
//		-rwxrwx---   0770           File - read, write, & execute for owner and group
//		-rwxrwxrwx   0777           File - read, write, & execute for owner, group and others
//		---x--x--x   0111           File - execute
//		--w--w--w-   0222           File - write only
//		--wx-wx-wx   0333           File - write & execute
//		-r--r--r--   0444           File - read only
//		-r-xr-xr-x   0555           File - read & execute
//		-rw-rw-rw-   0666           File - read & write
//		-rwxr-----   0740           File - Owner can read, write, & execute. Group can only read;
//	                                     others have no permissions
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		The internal FileMode data field for this
//		instance will be reset using the permission codes
//		contained in input parameter 'modeStr'.
//
//	modeStr						string
//
//		'modeStr' is a 10-character string containing the
//		read, write and execute permissions for the three
//		groups or user classes:
//
//			(1)	'Owner/User'
//
//			(2)	'Group'
//
//			(3)	'Other'
//
//		This 10-character string will be used to
//		configure the internal FileMode data field for
//		the instance of FilePermissionConfig passed as
//		input parameter 'fPerm'.
//
//		'modeStr' must conform to the symbolic notation
//		options shown above. Failure to comply with this
//		requirement will generate an error. As indicated,
//		'modeStr' must consist of 10-characters.
//
//		The first character in 'modeStr' may be '-'
//		specifying a fle or 'd' specifying a directory.
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
func (fPermConfigNanobot *filePermissionConfigNanobot) setFileModeByTextCode(
	fPerm *FilePermissionConfig,
	modeStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fPermConfigNanobot.lock == nil {
		fPermConfigNanobot.lock = new(sync.Mutex)
	}

	fPermConfigNanobot.lock.Lock()

	defer fPermConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigNanobot."+
			"setFileModeByTextCode()",
		"")

	if err != nil {
		return err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(modeStr) != 10 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'modeStr' MUST contain 10-characters.\n"+
			"This 'modeStr' contains %v-characters.\n"+
			"modeStr='%v' ",
			ePrefix.String(),
			len(modeStr),
			modeStr)

		return err
	}

	firstChar := string(modeStr[0])

	if firstChar != "-" &&
		firstChar != "d" {

		err = fmt.Errorf("%v\n"+
			"Error: First character of input parameter, 'modeStr' MUST BE 'd' or '-'.\n"+
			"This first character in 'modeStr' is = '%v'",
			ePrefix.String(),
			firstChar)

		return err
	}

	fPermElectron := filePermissionConfigElectron{}

	var ownerInt, groupInt, otherInt int

	var err2 error

	ownerInt,
		err2 = fPermElectron.convertGroupToDecimal(
		modeStr[1:4],
		"owner",
		ePrefix.XCpy(
			"ownerInt<-modeStr[1:4]"))

	if err2 != nil {

		err = fmt.Errorf(
			"%v\n"+
				"'owner' integer code errror for modeStr[1:4].\n"+
				"Error: %v",
			ePrefix.String(),
			err2.Error())

		return err
	}

	groupInt,
		err2 = fPermElectron.convertGroupToDecimal(
		modeStr[4:7],
		"group",
		ePrefix.XCpy(
			"groupInt<-modeStr[4:7]"))

	if err2 != nil {

		err = fmt.Errorf(
			"%v\n"+
				"'group' integer code errror for modeStr[4:7].\n"+
				"Error: %v",
			ePrefix.String(),
			err2.Error())

		return err
	}

	otherInt,
		err2 = fPermElectron.convertGroupToDecimal(
		modeStr[7:],
		"other",
		ePrefix.XCpy(
			"otherInt<-modeStr[7:]"))

	if err2 != nil {

		err = fmt.Errorf(
			"%v\n"+
				"'other' integer code errror for modeStr[7:].\n"+
				"Error: %v",
			ePrefix.String(),
			err2.Error())

		return err
	}

	ownerInt *= 100
	groupInt *= 10
	permission := ownerInt + groupInt + otherInt

	entryType := 0

	fMode := permission

	if firstChar == "d" {

		entryType = new(NumberConversions).
			ConvertDecimalToOctal(int(os.ModeDir))

		fMode = entryType | permission
	}

	fPerm.fileMode = os.FileMode(
		new(NumberConversions).ConvertOctalToDecimal(fMode))

	fPerm.isInitialized = true

	return nil
}
