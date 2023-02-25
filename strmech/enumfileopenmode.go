package strmech

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// Lock this mutex before accessing any of these maps.
var fileOpenModeMapsLock sync.Mutex

// mFileOpenModeToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOpenMode.
var mFileOpenModeToString = map[FileOpenMode]string{
	FileOpenMode(-1):          "ModeNone",
	FileOpenMode(os.O_APPEND): "ModeAppend",
	FileOpenMode(os.O_CREATE): "ModeCreate",
	FileOpenMode(os.O_EXCL):   "ModeExclusive",
	FileOpenMode(os.O_SYNC):   "ModeSync",
	FileOpenMode(os.O_TRUNC):  "ModeTruncate",
}

// mFileOpenModeToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOpenMode.
var mValidFileOpenModeToString = map[FileOpenMode]string{
	FileOpenMode(os.O_APPEND): "ModeAppend",
	FileOpenMode(os.O_CREATE): "ModeCreate",
	FileOpenMode(os.O_EXCL):   "ModeExclusive",
	FileOpenMode(os.O_SYNC):   "ModeSync",
	FileOpenMode(os.O_TRUNC):  "ModeTruncate",
}

// mFileOpenModeStringToInt - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOpenMode.
var mFileOpenModeStringToInt = map[string]int{
	"ModeNone":      -1,
	"ModeAppend":    os.O_APPEND,
	"ModeCreate":    os.O_CREATE,
	"ModeExclusive": os.O_EXCL,
	"ModeSync":      os.O_SYNC,
	"ModeTruncate":  os.O_TRUNC,
}

// mFileOpenModeLwrCaseStringToInt - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOpenMode.
// This map is used for case-insensitive look-ups.
var mFileOpenModeLwrCaseStringToInt = map[string]int{
	"modenone":      -1,
	"modeappend":    os.O_APPEND,
	"modecreate":    os.O_CREATE,
	"modeexclusive": os.O_EXCL,
	"modesync":      os.O_SYNC,
	"modetruncate":  os.O_TRUNC,
}

// FileOpenMode - To further control the file open operation, one
// or more FileOpenMode values may be or'd with a FileOpenType
// code in order to control behavior.
//
// In addition, one of the three  codes may be or'd with
// zero or more of the following File Open Modes (Type: 'FileOpenMode')
// to better control file open behavior.
//
// FileOpenMode has been adapted to function as an enumeration of valid
// File Open Mode values. Since Go does not directly support enumerations,
// the 'FileOpenMode' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//	Jeffrey Richter Using Reflection to implement enumerated types
//	       https://www.youtube.com/watch?v=DyXJy_0v0_U
//
// These FileOpenMode methods used as enumerators for os mode constants:
//
//	FileOpenMode(0).ModeNone()
//	FileOpenMode(0).ModeAppend()
//	FileOpenMode(0).ModeTypeCreate()
//	FileOpenMode(0).ModeExclusive()
//	FileOpenMode(0).ModeSync()
//	FileOpenMode(0).ModeTruncate()
//
//	Reference CONSTANTS: https://golang.org/pkg/os/
//
// The FileOpenType type is used in conjunction with FileOpenMode to specify
// file permissions. Reference the 'FileOpenType' source code documentation.
// The methods used to specify File Open Types are listed as follows:
//
//	FileOpenType(0).TypeReadOnly()
//	FileOpenType(0).TypeWriteOnly()
//	FileOpenType(0).TypeReadWrite()
//
//	Reference CONSTANTS: https://golang.org/pkg/os/
type FileOpenMode int

var enumFileOpenModeLock sync.Mutex

// ModeNone - No File Open Mode is active
func (fOpenMode FileOpenMode) ModeNone() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(-1)
}

// ModeAppend - append data to the file when writing.
func (fOpenMode FileOpenMode) ModeAppend() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(os.O_APPEND)
}

// ModeCreate - create a new file if none exists.
func (fOpenMode FileOpenMode) ModeCreate() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(os.O_CREATE)
}

// ModeExclusive - used with FileOpenControlMode(0).Create(), file must not exist.
func (fOpenMode FileOpenMode) ModeExclusive() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(os.O_EXCL)
}

// ModeSync - open for synchronous I/O.
func (fOpenMode FileOpenMode) ModeSync() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(os.O_SYNC)
}

// ModeTruncate - if possible, truncate file when opened.
func (fOpenMode FileOpenMode) ModeTruncate() FileOpenMode {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return FileOpenMode(os.O_TRUNC)
}

// IsValid - If the value of the current FileOpenMode is 'invalid',
// this method will return an error. If the FileOpenMode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
func (fOpenMode FileOpenMode) IsValid() error {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	_, ok := mValidFileOpenModeToString[fOpenMode]

	if !ok {
		ePrefix := "FileOpenMode.IsValidInstanceError() "
		return fmt.Errorf(ePrefix+
			"Error: Ivalid FileOpenMode! Current FileOpenMode='%v'",
			fOpenMode)
	}

	return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOpenMode is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case-sensitive and will require an
//	                       exact match. Therefore, 'append' will NOT
//	                       match the enumeration name, 'Append'.
//
//	                       If 'false' a case-insensitive search is conducted
//	                       for the enumeration name. In this case, 'append'
//	                       will match enumeration name 'Append'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	FileOpenMode - Upon successful completion, this method will return a new
//	               instance of FileOpenMode set to the value of the enumeration
//	               matched by the string search performed on input parameter,
//	               'valueString'.
//
//	error        - If this method completes successfully, the returned error
//	               Type is set equal to 'nil'. If an error condition is encountered,
//	               this method will return an error Type which encapsulates an
//	               appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t, err := FileOpenMode(0).ParseString("Append", true)
//	                     OR
//	t, err := FileOpenMode(0).ParseString("ModeAppend", true)
//	                     OR
//	t, err := FileOpenMode(0).ParseString("ModeAppend()", true)
//	                     OR
//	t, err := FileOpenMode(0).ParseString("Append()", true)
//	                     OR
//	t, err := FileOpenMode(0).ParseString("append", false)
//
//	In any case shown above, t is now equal to FileOpenMode(0).Append()
func (fOpenMode FileOpenMode) ParseString(
	valueString string,
	caseSensitive bool) (FileOpenMode, error) {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	fileOpenModeMapsLock.Lock()

	defer fileOpenModeMapsLock.Unlock()

	ePrefix := "FileOpenMode.ParseString() "

	var result FileOpenMode

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! valueString='%v' ", valueString)
	}

	var ok bool
	var idx int

	if caseSensitive {

		if !strings.HasPrefix(valueString, "Mode") {
			valueString = "Mode" + valueString
		}

		idx, ok = mFileOpenModeStringToInt[valueString]

		if !ok {
			return FileOpenMode(0),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenMode. valueString='%v' ", valueString)
		}

		result = FileOpenMode(idx)

	} else {

		valueString = strings.ToLower(valueString)

		if !strings.HasPrefix(valueString, "mode") {
			valueString = "mode" + valueString
		}

		idx, ok = mFileOpenModeLwrCaseStringToInt[valueString]

		if !ok {
			return FileOpenMode(-1),
				fmt.Errorf(ePrefix+
					"'valueString' did NOT MATCH a FileOpenMode. valueString='%v' ", valueString)
		}

		result =
			FileOpenMode(idx)
	}

	return result, nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOpenMode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//	string - The string label or description for the current enumeration
//	         value. If, the FileOpenMode value is invalid, this method will
//	         return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOpenMode(0).ModeAppend()
//	str := t.String()
//	    str is now equal to 'ModeAppend'
func (fOpenMode FileOpenMode) String() string {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	fileOpenModeMapsLock.Lock()

	defer fileOpenModeMapsLock.Unlock()

	str, ok := mFileOpenModeToString[fOpenMode]

	if !ok {
		return ""
	}

	return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOpenMode
// instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (fOpenMode FileOpenMode) Value() int {

	enumFileOpenModeLock.Lock()

	defer enumFileOpenModeLock.Unlock()

	return int(fOpenMode)
}

// FOpenMode - This public global variable allows
// easy access to the enumerations of the FileOpenMode
// type using the dot operator.
//
//	Example:
//
//		FOpenMode.ModeNone()
//		FOpenMode.ModeAppend()
//		FOpenMode.ModeCreate()
//		FOpenMode.ModeExclusive()
//		FOpenMode.ModeSync()
//		FOpenMode.ModeTruncate()
const FOpenMode = FileOpenMode(0)
