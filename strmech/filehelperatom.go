package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type fileHelperAtom struct {
	lock *sync.Mutex
}

// setFileModeByTextCode
//
// Sets the internal FileMode data field using input
// parameter 'modeStr'. 'modeStr' is a 10-character
// string containing the read, write and execute
// permissions for the three groups, 'Owner/User',
// 'Group' and 'Other'.
//
// The text codes used in the 'modeStr' mimic the
// Unix permission codes.
//
//	Reference:
//
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//	https://en.wikipedia.org/wiki/File_system_permissions
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
//	"-rwxrwxrwx"
//
//	Groups: - Owner, Group, Other
//	From left to right
//
//	Char index 0     = Entry Type. Must be either a "-" or
//						a "d"
//
//	Char indexes 4-6 = Group  "rwx"  Authorizing 'Read',
//						'Write' & Execute Permissions for
//						'Group'
//
//	Char indexes 1-3 = Owner  "rwx"  Authorizing 'Read',
//						'Write' And Execute Permissions
//						for 'Owner'
//
//	Char indexes 7-9 = Other  "rwx"  Authorizing 'Read',
//						'Write' & Execute Permissions for
//						'Other'
//
// The Symbolic notation provided by input parameter
// 'modeStr' MUST conform to the options presented below.
// The first character or 'Entry Type' is listed as "-".
// However, in practice, the caller may set the first
// character as either a "-", specifying a file, or a "d",
// specifying a directory. No other first character types
// are currently supported.
//
// After the first character, the remaining 9-characters
// are constituents of the three Symbolic Groups:
//
//	Owners/Users
//	Groups
//	Others
//
// Each group has three characters which may be 'r', 'w',
// 'x'. If a permission is not set, the character position
// contains a '-'.
//
//		 'modeStr'
//		 Symbolic	 Octal			File Access
//		 Format	 	Notation	Permission Descriptions
//		------------------------------------------------------
//
//	  ----------   0000		File - no permissions
//	  -rwx------   0700		File - read, write, & execute only for
//	 								owner
//	  -rwxrwx---   0770		File - read, write, & execute for owner
//	 								and group
//	  -rwxrwxrwx   0777		File - read, write, & execute for owner,
//	 								 group and others
//	  ---x--x--x   0111		File - execute
//	  --w--w--w-   0222		File - write only
//	  --wx-wx-wx   0333		File - write & execute
//	  -r--r--r--   0444		File - read only
//	  -r-xr-xr-x   0555		File - read & execute
//	  -rw-rw-rw-   0666		File - read & write
//	  -rwxr-----   0740		File - Owner can read, write, & execute.
//	 								 Group can only read;
//	                   				 Others have no permissions
//
//	  drwxrwxrwx   20000000777	File - Directory - read, write, & execute for
//	 									 Owner, Group and Others
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//	modeStr  string - 'modeStr' must conform to the symbolic notation options shown
//	                  above. Failure to comply with this requirement will generate an
//	                  error. As indicated, 'modeStr' must consist of 10-characters.
//	                  The first character in 'modeStr' may be '-' specifying a fle or
//	                  'd' specifying a directory.
//
//	Reference:
//	How to use special permissions: the setuid, setgid and sticky bits
//	https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
func (fileHelpAtom *fileHelperAtom) setFileModeByTextCode(
	modeStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpAtom.lock == nil {
		fileHelpAtom.lock = new(sync.Mutex)
	}

	fileHelpAtom.lock.Lock()

	defer fileHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"setFileModeByTextCode()",
		"")

	if err != nil {
		err = fmt.Errorf("%v\n", err.Error())

		return formattedText, err
	}

	if len(modeStr) != 10 {
		return fmt.Errorf(ePrefix+
			"Error: Input parameter 'modeStr' MUST contain 10-characters. This 'modeStr' "+
			"contains %v-characters. modeStr='%v' ", len(modeStr), modeStr)
	}

	firstChar := string(modeStr[0])

	if firstChar != "-" &&
		firstChar != "d" {
		return fmt.Errorf(ePrefix+
			"Error: First character of input parameter, 'modeStr' MUST BE 'd' or '-'. "+
			"This first character = '%v'", firstChar)
	}

	ownerInt, err := fPerm.convertGroupToDecimal(modeStr[1:4], "owner")

	if err != nil {
		return fmt.Errorf(ePrefix+"'ownerInt' Error: %v", err.Error())
	}

	groupInt, err := fPerm.convertGroupToDecimal(modeStr[4:7], "group")

	if err != nil {
		return fmt.Errorf(ePrefix+"groupInt Error: %v", err.Error())
	}

	otherInt, err := fPerm.convertGroupToDecimal(modeStr[7:], "other")

	if err != nil {
		return fmt.Errorf(ePrefix+"otherInt Error: %v", err.Error())
	}

	ownerInt *= 100
	groupInt *= 10
	permission := ownerInt + groupInt + otherInt

	entryType := 0

	fMode := permission

	fh := FileHelper{}

	if firstChar == "d" {
		entryType = fh.ConvertDecimalToOctal(int(os.ModeDir))
		fMode = entryType | permission
	}

	fPerm.fileMode = os.FileMode(fh.ConvertOctalToDecimal(fMode))
	fPerm.isInitialized = true

	return nil
}
