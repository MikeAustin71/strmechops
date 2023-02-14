package strmech

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Lock this Mutex before accessing these maps
var osPermissionCodeMapLock sync.Mutex

var mOsPermissionCodeToString = map[os.FileMode]string{
	os.FileMode(0):    "ModeNone",
	os.ModeDir:        "ModeDir",
	os.ModeAppend:     "ModeAppend",
	os.ModeExclusive:  "ModeExclusive",
	os.ModeTemporary:  "ModeTemporary",
	os.ModeSymlink:    "ModeSymlink",
	os.ModeDevice:     "ModeDevice",
	os.ModeNamedPipe:  "ModeNamedPipe",
	os.ModeSocket:     "ModeSocket",
	os.ModeSetuid:     "ModeSetuid",
	os.ModeSetgid:     "ModeSetgid",
	os.ModeCharDevice: "ModeCharDevice",
	os.ModeSticky:     "ModeSticky",
	os.ModeIrregular:  "ModeIrregular",
}

var mOsPermissionCodeToLetter = map[os.FileMode]string{
	os.FileMode(0):    "-",
	os.ModeDir:        "d",
	os.ModeAppend:     "a",
	os.ModeExclusive:  "l",
	os.ModeTemporary:  "T",
	os.ModeSymlink:    "L",
	os.ModeDevice:     "D",
	os.ModeNamedPipe:  "p",
	os.ModeSocket:     "S",
	os.ModeSetuid:     "u",
	os.ModeSetgid:     "g",
	os.ModeCharDevice: "c",
	os.ModeSticky:     "t",
	os.ModeIrregular:  "?",
}

var mOsPermissionStringToCode = map[string]os.FileMode{
	"ModeNone":       os.FileMode(0),
	"ModeDir":        os.ModeDir,
	"ModeAppend":     os.ModeAppend,
	"ModeExclusive":  os.ModeExclusive,
	"ModeTemporary":  os.ModeTemporary,
	"ModeSymlink":    os.ModeSymlink,
	"ModeDevice":     os.ModeDevice,
	"ModeNamedPipe":  os.ModeNamedPipe,
	"ModeSocket":     os.ModeSocket,
	"ModeSetuid":     os.ModeSetuid,
	"ModeSetgid":     os.ModeSetgid,
	"ModeCharDevice": os.ModeCharDevice,
	"ModeSticky":     os.ModeSticky,
	"ModeIrregular":  os.ModeIrregular,
}

var mOsPermissionLwrCaseStringToCode = map[string]os.FileMode{
	"modenone":       os.FileMode(0),
	"modedir":        os.ModeDir,
	"modeappend":     os.ModeAppend,
	"modeexclusive":  os.ModeExclusive,
	"modetemporary":  os.ModeTemporary,
	"modesymlink":    os.ModeSymlink,
	"modedevice":     os.ModeDevice,
	"modenamedpipe":  os.ModeNamedPipe,
	"modesocket":     os.ModeSocket,
	"modesetuid":     os.ModeSetuid,
	"modesetgid":     os.ModeSetgid,
	"modechardevice": os.ModeCharDevice,
	"modesticky":     os.ModeSticky,
	"modeirregular":  os.ModeIrregular,
}

var mOsPermissionLetterToCode = map[string]os.FileMode{
	"-": os.FileMode(0),
	"d": os.ModeDir,
	"a": os.ModeAppend,
	"l": os.ModeExclusive,
	"T": os.ModeTemporary,
	"L": os.ModeSymlink,
	"D": os.ModeDevice,
	"p": os.ModeNamedPipe,
	"S": os.ModeSocket,
	"u": os.ModeSetuid,
	"g": os.ModeSetgid,
	"c": os.ModeCharDevice,
	"t": os.ModeSticky,
	"?": os.ModeIrregular,
}

// OsFilePermissionCode
//
// An enumeration of the os File Mode constant values:
//
//	 Method           os.FileMode           Associated
//	  Name             Constant             Letter Code
//	______________________________________________________________________
//	ModeNone()        os.ModeNone           "-" is a file
//	ModeDir()         os.ModeDir            "d" is a directory
//	ModeAppend()      os.ModeAppend         "a" append-only
//	ModeExclusive()   os.ModeExclusive      "l" exclusive use
//	ModeTemporary()   os.ModeTemporary      "T" temporary file; Plan 9 only
//	ModeSymlink()     os.ModeSymlink        "L" symbolic link
//	ModeDevice()      os.ModeDevice         "D" device file
//	ModeNamedPipe()   os.ModeNamedPipe      "p" named pipe (FIFO)
//	ModeSocket()      os.ModeSocket         "S" Unix domain socket
//	ModeSetuid()      os.ModeSetuid         "u" setuid
//	ModeSetgid()      os.ModeSetgid         "g" setgid
//	ModeCharDevice()  os.ModeCharDevice     "c" Unix character device, when ModeDevice is set
//	ModeSticky()      os.ModeSticky         "t" sticky
//	ModeIrregular()   os.ModeIrregular      "?" non-regular file; nothing else is known about this file
//
// For more information on os Mode Constants Reference:
// https://golang.org/pkg/os/#pkg-constants
//
// OsFilePermissionCode has been adapted to function as an enumeration of valid
// File Permission Code values. Since Go does not directly support enumerations,
// the 'OsFilePermissionCode' type has been configured to function in a manner
// similar to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//	Jeffrey Richter Using Reflection to implement enumerated types
//	       https://www.youtube.com/watch?v=DyXJy_0v0_U
type OsFilePermissionCode os.FileMode

var osPermissionCodeTypeLock sync.Mutex

// ModeNone
//
// "-" No Permission Set
//
// There is no os constant for 'None'. However since the
// zero value is used extensively to identity a 'file'
// within the context of permission descriptions, it is
// added here. 'ModeNone' therefore represents both a
// zero value and the 'file' designation.
func (osPerm OsFilePermissionCode) ModeNone() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.FileMode(0)

}

// ModeDir
//
// Letter Code= "d" is a directory - alias for os.ModeDir
func (osPerm OsFilePermissionCode) ModeDir() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeDir
}

// ModeAppend
//
// Letter Code= "a" append-only - alias for os.ModeAppend
func (osPerm OsFilePermissionCode) ModeAppend() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeAppend
}

// ModeExclusive
//
// Letter Code= "l" (Letter 'l' as in lima)
// exclusive use   - alias for os.ModeExclusive
func (osPerm OsFilePermissionCode) ModeExclusive() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeExclusive
}

// ModeTemporary
//
// Letter Code= "T" temporary file;
// Plan 9 only  - alias for os.ModeTemporary
func (osPerm OsFilePermissionCode) ModeTemporary() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeTemporary
}

// ModeSymlink
//
// Letter Code= "L" symbolic link
//
// Alias for os.ModeSymlink
func (osPerm OsFilePermissionCode) ModeSymlink() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeSymlink
}

// ModeDevice
//
// Letter Code= "D" device file
//
// Alias for os.ModeDevice
func (osPerm OsFilePermissionCode) ModeDevice() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeDevice
}

// ModeNamedPipe
//
// Letter Code= "p" named pipe (FIFO)
//
// Alias for os.ModeNamedPipe
func (osPerm OsFilePermissionCode) ModeNamedPipe() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeNamedPipe
}

// ModeSocket
//
// Letter Code= "S" Unix domain socket
//
// Alias for os.ModeSocket
func (osPerm OsFilePermissionCode) ModeSocket() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeSocket
}

// ModeSetuid
//
// Letter Code= "u" setuid - alias for os.ModeSetuid
//
// When the setuid bit is used, the behavior described
// above it's modified so that when an executable is
// launched, it does not run with the privileges of the
// user who launched it, but with that of the file owner
// instead. So, for example, if an executable has the
// setuid bit set on it, and it's owned by root, when
// launched by a normal user, it will run with root
// privileges. It should be clear why this represents
// a potential security risk, if not used correctly.
func (osPerm OsFilePermissionCode) ModeSetuid() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeSetuid
}

// ModeSetgid
//
// Letter Code= "g" setgid - alias for os.ModeSetgid
//
// Unlike the setuid bit, the setgid bit has effect on
// both files and directories. In the first case, the
// file which has the setgid bit set, when executed,
// instead of running with the privileges of the group
// of the user who started it, runs with those of the
// group which owns the file: in other words, the group
// ID of the process will be the same of that of the
// file.
//
// When used on a directory, instead, the setgid bit
// alters the standard behavior so that the group of the
// files created inside said directory, will not be that
// of the user who created them, but that of the parent
// directory itself. This is often used to ease the
// sharing of files (files will be modifiable by all the
// users that are part of said group).
func (osPerm OsFilePermissionCode) ModeSetgid() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeSetgid
}

// ModeCharDevice
//
// Letter Code= "c" Unix character device, when
// ModeDevice is set.
//
// Alias for os.ModeCharDevice
func (osPerm OsFilePermissionCode) ModeCharDevice() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeCharDevice
}

// ModeSticky
//
// Letter Code= "t" sticky - alias for os.ModeSticky
//
// The sticky bit works in a different way: while it has
// no effect on files, when used on a directory, all the
// files in said directory will be modifiable only by
// their owners. A typical case in which it is used,
// involves the /tmp directory. Typically, this directory
// is writable by all users on the system, so to make
// impossible for one user to delete the files of another
// one.
func (osPerm OsFilePermissionCode) ModeSticky() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeSticky
}

// ModeIrregular
//
// Letter Code= "?" non-regular file; nothing else is
// known about this file.
//
// Alias for os.ModeIrregular
func (osPerm OsFilePermissionCode) ModeIrregular() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.ModeIrregular
}

// Equal
//
// Compares the current OsFilePermissionCode instance to
// another OsFilePermission instance passed as an input
// parameter. If the two are equal in all respects, this
// method returns 'true'.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (osPerm OsFilePermissionCode) Equal(osPerm2 OsFilePermissionCode) bool {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	if osPerm == osPerm2 {
		return true
	}

	return false
}

// GetFileModeLetterCode
//
// Returns the single alphabetic character associated
// with this os.FileMode. All os.FileMode's are
// associated with a single letter used in unix
// permission strings.
//
//	                     Letter
//	File Mode             Code     Description
//	__________________________________________
//	os.ModeNone           "-"      is a file
//	os.ModeDir            "d"      is a directory
//	os.ModeAppend         "a"      append-only
//	os.ModeExclusive      "l"      exclusive use
//	os.ModeTemporary      "T"      temporary file; Plan 9 only
//	os.ModeSymlink        "L"      symbolic link
//	os.ModeDevice         "D"      device file
//	os.ModeNamedPipe      "p"      named pipe (FIFO)
//	os.ModeSocket         "S"      Unix domain socket
//	os.ModeSetuid         "u"      setuid
//	os.ModeSetgid         "g"      setgid
//	os.ModeCharDevice     "c"      Unix character device, when ModeDevice is set
//	os.ModeSticky         "t"      sticky
//	os.ModeIrregular      "?"      non-regular file; nothing else is known about this file
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (osPerm OsFilePermissionCode) GetFileModeLetterCode() (string, error) {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	ePrefix := "OsFilePermissionCode.FileModeLetterCode() "

	letter, ok := mOsPermissionCodeToLetter[os.FileMode(osPerm)]

	if !ok {
		return "", fmt.Errorf(ePrefix+
			"The current OsFilePermissionCode FileMode value is INVALID! "+
			"OsFilePermissionCode decimal value=%s  octal value=%s",
			strconv.FormatInt(int64(osPerm), 10),
			strconv.FormatInt(int64(osPerm), 8))
	}

	return letter, nil
}

// GetNewFromFileMode
//
// Creates and returns a new OsFilePermissionCode
// instance generated from the os.FileMode type input
// parameter ('fMode'). If the input os.FileMode value is
// invalid, an error is returned.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (osPerm OsFilePermissionCode) GetNewFromFileMode(
	fMode os.FileMode) (OsFilePermissionCode, error) {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	newFilePerm := OsFilePermissionCode(fMode)

	isValid,
		_ := new(osFilePermissionCodeNanobot).
		isValidOsFPermCode(newFilePerm)

	if !isValid {
		ePrefix := "OsFilePermissionCode.GetNewFromFileMode() "

		return OsFilePermissionCode(0),
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fMode' is an INVALID File Mode!\n"+
				"OsFilePermissionCode Octal Value='%s'\n",
				ePrefix,
				strconv.FormatInt(int64(newFilePerm), 8))
	}

	return newFilePerm, nil
}

// GetNewFromLetterCode - Creates a new OsFilePermissionCode instance based on an
// associated 'letter code'. The letter code consists of a single character
// representing an os.FileMode. This single character is useful in configuring
// unix permission strings.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (osPerm OsFilePermissionCode) GetNewFromLetterCode(
	letterCode string) (OsFilePermissionCode, error) {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	fModeValue, ok := mOsPermissionLetterToCode[letterCode]

	if !ok {
		ePrefix := "OsFilePermissionCode.GetNewFromLetterCode() "

		return OsFilePermissionCode(0),
			fmt.Errorf("%v\n"+
				"Error: 'letterCode' is INVALID! "+
				"letterCode='%v'",
				ePrefix,
				letterCode)
	}

	return OsFilePermissionCode(fModeValue), nil
}

// IsValid
//
// If the value of the current OsFilePermissionCode
// instance is 'invalid', this method will return
// an error.
//
// If the OsFilePermissionCode instance is 'valid', this method will return
// a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
func (osPerm OsFilePermissionCode) IsValid() error {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	_,
		err := new(osFilePermissionCodeNanobot).
		isValidOsFPermCode(osPerm)

	return err
}

// ParseString
//
// Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful,
// a new instance of OsFilePermissionCode is returned set
// to the value of the associated enumeration.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString					string
//
//		A string which will be matched against the
//		enumeration string values. If 'valueString' is
//		equal to one of the enumeration names, this
//		method will proceed to successful completion.
//
//	caseSensitive				bool
//
//		If 'true' the search for enumeration names will
//		be case-sensitive and will require an exact
//		match. Therefore, 'modedir' will NOT match the
//		enumeration name, 'ModeDir'.
//
//		If 'false' a case-insensitive search is conducted
//		for the enumeration name. In this case, 'modedir'
//		will match enumeration name 'ModeDir'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	OsFilePermissionCode
//
//		Upon successful completion, this method will
//		return a new instance of OsFilePermissionCode
//		set to the value of the enumeration matched by
//		the string search performed on input parameter,
//		'valueString'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//		If an error condition is encountered, this
//		method will return an error Type which
//		encapsulates an appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//		t, err := OsFilePermissionCode(0).ParseString("ModeDir", true)
//	                           OR
//		t, err := OsFilePermissionCode(0).ParseString("ModeDir()", true)
//	                           OR
//		t, err := OsFilePermissionCode(0).ParseString("Dir()", true)
//	                           OR
//		t, err := OsFilePermissionCode(0).ParseString("Dir", true)
//	                           OR
//		t, err := OsFilePermissionCode(0).ParseString("modedir", false)
//
//		For all the cases shown above,
//	 t is now equal to OsFilePermissionCode(0).ModeDir()
func (osPerm OsFilePermissionCode) ParseString(
	valueString string,
	caseSensitive bool) (OsFilePermissionCode, error) {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	ePrefix := "OsFilePermissionCode.ParseString() "

	valueString = strings.TrimLeft(strings.TrimRight(valueString, " "), " ")

	lenValueStr := len(valueString)

	if lenValueStr == 0 {
		return OsFilePermissionCode(0),
			errors.New(ePrefix +
				"Error: Input parameter 'valueString' is an empty string and therefore INVALID!\n")
	}

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	result := OsFilePermissionCode(0)

	if lenValueStr < 3 {
		return result,
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters valueString='%v' ", valueString)
	}

	var permCode os.FileMode
	var ok bool

	if caseSensitive {

		if !strings.HasPrefix(valueString, "Mode") {
			valueString = "Mode" + valueString
		}

		permCode, ok = mOsPermissionStringToCode[valueString]

		if !ok {
			return result,
				fmt.Errorf(ePrefix+
					"Error: Invalid Permission Code!\n"+
					"valueString='%v'\n", valueString)
		}

		result = OsFilePermissionCode(permCode)

	} else {

		valueString = strings.ToLower(valueString)

		if !strings.HasPrefix(valueString, "mode") {
			valueString = "mode" + valueString
		}

		permCode, ok = mOsPermissionLwrCaseStringToCode[valueString]

		if !ok {
			return result,
				fmt.Errorf(ePrefix+
					"Error: Invalid Permission Code!\n"+
					"valueString='%v'\n", valueString)
		}

		result = OsFilePermissionCode(permCode)

	}

	return result, nil
}

// String
//
// Returns a string with the name of the enumeration
// associated with this instance of
// 'OsFilePermissionCode'.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//	string
//
//		The string label or description for the current
//		enumeration value. If, the OsFilePermissionCode
//		value is invalid, this method will return an
//		empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= OsFilePermissionCode(0).ModeDir()
//	str := t.String()
//	    str is now equal to "ModeDir"
func (osPerm OsFilePermissionCode) String() string {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	label, ok := mOsPermissionCodeToString[os.FileMode(osPerm)]

	if !ok {
		return ""
	}

	return label
}

// Value - Returns the value of the OsFilePermissionCode instance
// as type os.FileMode.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (osPerm OsFilePermissionCode) Value() os.FileMode {

	osPermissionCodeTypeLock.Lock()

	defer osPermissionCodeTypeLock.Unlock()

	return os.FileMode(osPerm)
}

// OsFilePermCode - public global variable of type OsFilePermissionCode.
// Provides alternative, easier access to OsFilePermissionCode enumeration
// values.
//
// Usage:
//
//	OsFilePermCode.ModeNone()
//	OsFilePermCode.ModeDir()
//	OsFilePermCode.ModeAppend()
//	OsFilePermCode.ModeExclusive()
//	OsFilePermCode.ModeTemporary()
//	OsFilePermCode.ModeSymlink()
//	OsFilePermCode.ModeDevice()
//	OsFilePermCode.ModeNamedPipe()
//	OsFilePermCode.ModeSocket()
//	OsFilePermCode.ModeSetuid()
//	OsFilePermCode.ModeSetgid()
//	OsFilePermCode.ModeCharDevice()
//	OsFilePermCode.ModeSticky()
//	OsFilePermCode.ModeIrregular()
const OsFilePermCode = OsFilePermissionCode(0)

// osFilePermissionCodeNanobot - Provides helper methods for
// enumeration OsFilePermissionCode.
type osFilePermissionCodeNanobot struct {
	lock *sync.Mutex
}

// isValidOsFPermCode
//
// If the value of the OsFilePermissionCode input
// parameter 'fPermCode' is 'invalid', this method will
// return a boolean value of 'false' and an error.
//
// If the OsFilePermissionCode instance is 'valid', this
// method will return a boolean value of 'true' and an
// error value of 'nil'.
func (ofFPermCodeNanobot *osFilePermissionCodeNanobot) isValidOsFPermCode(
	fPermCode OsFilePermissionCode) (
	isValid bool,
	err error) {

	if ofFPermCodeNanobot.lock == nil {
		ofFPermCodeNanobot.lock = new(sync.Mutex)
	}

	ofFPermCodeNanobot.lock.Lock()

	defer ofFPermCodeNanobot.lock.Unlock()

	osPermissionCodeMapLock.Lock()

	defer osPermissionCodeMapLock.Unlock()

	isValid = false

	_, isValid = mOsPermissionCodeToString[os.FileMode(fPermCode)]

	if !isValid {

		err = fmt.Errorf(
			"%v\n"+
				"The current OsFilePermissionCode is INVALID!\n"+
				"OsFilePermissionCode Octal Value='%s'\n",
			"OsFilePermissionCode.IsValid()",
			strconv.FormatInt(int64(fPermCode), 8))
	}

	return isValid, err
}
