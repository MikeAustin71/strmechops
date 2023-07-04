package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock this mutex before accessing any of these maps.
var fileOperationCodeMapsLock sync.Mutex

// mFileOperationCodeToString - This map is used to map enumeration values
// to enumeration names stored as strings for Type FileOperationCode.
var mFileOperationCodeToString = map[FileOperationCode]string{
	FileOperationCode(0):  "None",
	FileOperationCode(1):  "MoveSourceFileToDestinationFile",
	FileOperationCode(2):  "MoveSourceFileToDestinationDir",
	FileOperationCode(3):  "DeleteDestinationFile",
	FileOperationCode(4):  "DeleteSourceFile",
	FileOperationCode(5):  "DeleteSourceAndDestinationFiles",
	FileOperationCode(6):  "CopySourceToDestinationByHardLinkByIo",
	FileOperationCode(7):  "CopySourceToDestinationByIoByHardLink",
	FileOperationCode(8):  "CopySourceToDestinationByHardLink",
	FileOperationCode(9):  "CopySourceToDestinationByIo",
	FileOperationCode(10): "CreateSourceDir",
	FileOperationCode(11): "CreateSourceDirAndFile",
	FileOperationCode(12): "CreateSourceFile",
	FileOperationCode(13): "CreateDestinationDir",
	FileOperationCode(14): "CreateDestinationDirAndFile",
	FileOperationCode(15): "CreateDestinationFile",
}

// mValidFileOperationCodeToString
// This map is used to map valid enumeration values to
// valid enumeration names stored as strings for Type
// FileOperationCode.
var mValidFileOperationCodeToString = map[FileOperationCode]string{
	FileOperationCode(0):  "None",
	FileOperationCode(1):  "MoveSourceFileToDestinationFile",
	FileOperationCode(2):  "MoveSourceFileToDestinationDir",
	FileOperationCode(3):  "DeleteDestinationFile",
	FileOperationCode(4):  "DeleteSourceFile",
	FileOperationCode(5):  "DeleteSourceAndDestinationFiles",
	FileOperationCode(6):  "CopySourceToDestinationByHardLinkByIo",
	FileOperationCode(7):  "CopySourceToDestinationByIoByHardLink",
	FileOperationCode(8):  "CopySourceToDestinationByHardLink",
	FileOperationCode(9):  "CopySourceToDestinationByIo",
	FileOperationCode(10): "CreateSourceDir",
	FileOperationCode(11): "CreateSourceDirAndFile",
	FileOperationCode(12): "CreateSourceFile",
	FileOperationCode(13): "CreateDestinationDir",
	FileOperationCode(14): "CreateDestinationDirAndFile",
	FileOperationCode(15): "CreateDestinationFile",
}

// mFileOperationCodeStringToCode - This map is used to map enumeration names
// stored as strings to enumeration values for Type FileOperationCode.
var mFileOperationCodeStringToCode = map[string]FileOperationCode{
	"None":                                  FileOperationCode(0),
	"MoveSourceFileToDestinationFile":       FileOperationCode(1),
	"MoveSourceFileToDestinationDir":        FileOperationCode(2),
	"DeleteDestinationFile":                 FileOperationCode(3),
	"DeleteSourceFile":                      FileOperationCode(4),
	"DeleteSourceAndDestinationFiles":       FileOperationCode(5),
	"CopySourceToDestinationByHardLinkByIo": FileOperationCode(6),
	"CopySourceToDestinationByIoByHardLink": FileOperationCode(7),
	"CopySourceToDestinationByHardLink":     FileOperationCode(8),
	"CopySourceToDestinationByIo":           FileOperationCode(9),
	"CreateSourceDir":                       FileOperationCode(10),
	"CreateSourceDirAndFile":                FileOperationCode(11),
	"CreateSourceFile":                      FileOperationCode(12),
	"CreateDestinationDir":                  FileOperationCode(13),
	"CreateDestinationDirAndFile":           FileOperationCode(14),
	"CreateDestinationFile":                 FileOperationCode(15),
}

// mFileOperationCodeLwrCaseStringToCode - This map is used to map enumeration names
// stored as lower case strings to enumeration values for Type FileOperationCode.
// This map is used for case-insensitive look-ups.
var mFileOperationCodeLwrCaseStringToCode = map[string]FileOperationCode{
	"none":                                  FileOperationCode(0),
	"movesourcefiletodestinationfile":       FileOperationCode(1),
	"movesourcefiletodestinationdir":        FileOperationCode(2),
	"deletedestinationfile":                 FileOperationCode(3),
	"deletesourcefile":                      FileOperationCode(4),
	"deletesourceanddestinationfiles":       FileOperationCode(5),
	"copysourcetodestinationbyhardlinkbyio": FileOperationCode(6),
	"copysourcetodestinationbyiobyhardlink": FileOperationCode(7),
	"copysourcetodestinationbyhardlink":     FileOperationCode(8),
	"copysourcetodestinationbyio":           FileOperationCode(9),
	"createsourcedir":                       FileOperationCode(10),
	"createsourcedirandfile":                FileOperationCode(11),
	"createsourcefile":                      FileOperationCode(12),
	"createdestinationdir":                  FileOperationCode(13),
	"createdestinationdirandfile":           FileOperationCode(14),
	"createdestinationfile":                 FileOperationCode(15),
}

// FileOperationCode - Integer enumeration. Signals
// the type of operation to be performed on a file.
//
// Usage:
//
// ----------------------------------------------------
//
// To designate a File Operation Code, use the form:
//
//	FileOperationCode(0).MoveSourceFileToDestinationFile()
//
// To access the File Operation Code using strictly dot
// notation, use either the private variable:
//
//	FileOpCode.MoveSourceFileToDestinationFile()
//
// OR the public global variable:
//
//	FileOpCode.MoveSourceFileToDestinationFile()
//
// Listing Of File Operation Codes:
//
// ----------------------------------------------------
//
//	FileOperationCode(0).None()
//	FileOperationCode(0).MoveSourceFileToDestinationFile()
//	FileOperationCode(0).MoveSourceFileToDestinationDir()
//	FileOperationCode(0).DeleteDestinationFile()
//	FileOperationCode(0).DeleteSourceFile()
//	FileOperationCode(0).DeleteSourceAndDestinationFiles()
//	FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//	FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//	FileOperationCode(0).CopySourceToDestinationByHardLink()
//	FileOperationCode(0).CopySourceToDestinationByIo()
//	FileOperationCode(0).CreateSourceDir()
//	FileOperationCode(0).CreateSourceDirAndFile()
//	FileOperationCode(0).CreateSourceFile()
//	FileOperationCode(0).CreateDestinationDir()
//	FileOperationCode(0).CreateDestinationDirAndFile()
//	FileOperationCode(0).CreateDestinationFile()
//
// FileOperationCode has been adapted to function as an enumeration of valid
// File Operation Code values. Since Go does not directly support enumerations,
// the 'FileOperationCode' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//	Jeffrey Richter Using Reflection to implement enumerated types
//	       https://www.youtube.com/watch?v=DyXJy_0v0_U
type FileOperationCode int

var enumFileOperationCodeLock sync.Mutex

// None - No operation (NOOP) No File Operation is performed.
//
// Usage:
//
//	FileOperationCode(0).None()
//			OR
//	FileOpCode.None()
func (fop FileOperationCode) None() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(0)
}

// MoveSourceFileToDestinationFile - Moves the source file to the destination
// file and then deletes the original source file.
//
// Usage:
//
//	FileOperationCode(0).MoveSourceFileToDestinationFile()
//			OR
//	FileOpCode.MoveSourceFileToDestinationFile()
func (fop FileOperationCode) MoveSourceFileToDestinationFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(1)
}

// MoveSourceFileToDestinationDir - Moves the source file to the destination
// directory and then deletes the original source file. The destination file
// name is the same as the source file name.
//
// Usage:
//
//	FileOperationCode(0).MoveSourceFileToDestinationDir()
//			OR
//	FileOpCode.MoveSourceFileToDestinationDir()
func (fop FileOperationCode) MoveSourceFileToDestinationDir() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(2)
}

// DeleteDestinationFile
//
// # Deletes the Destination file if it exists
//
// Usage:
//
//	FileOperationCode(0).DeleteDestinationFile()
//			OR
//	FileOpCode.DeleteDestinationFile()
func (fop FileOperationCode) DeleteDestinationFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(3)
}

// DeleteSourceFile
//
// # Deletes the Source file if it exists
//
// Usage:
//
//	FileOperationCode(0).DeleteSourceFile()
//			OR
//	FileOpCode.DeleteSourceFile()
func (fop FileOperationCode) DeleteSourceFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(4)
}

// DeleteSourceAndDestinationFiles - Deletes both the Source and Destination files
// if they exist.
//
// Usage:
//
//	FileOperationCode(0).DeleteSourceAndDestinationFiles()
//			OR
//	FileOpCode.DeleteSourceAndDestinationFiles()
func (fop FileOperationCode) DeleteSourceAndDestinationFiles() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(5)
}

// CopySourceToDestinationByHardLinkByIo - Copies the Source File to the
// Destination using two copy attempts. The first copy is by Hard Link.
// If the first copy attempt fails, a second copy attempt is initiated
// by creating a new file and copying the contents by 'io.Copy'.
//
// An error is returned only if both copy attempts fail. The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()()
//			OR
//	FileOpCode.CopySourceToDestinationByHardLinkByIo()
func (fop FileOperationCode) CopySourceToDestinationByHardLinkByIo() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(6)
}

// CopySourceToDestinationByIoByHardLink - Copies the Source File to the Destination
// using two copy attempts. The first copy is by 'io.Copy' which creates a new file
// and copies the contents to the new file. If the first attempt fails, a second
// copy attempt is initiated using 'copy by hard link'.
//
// An error is returned only if both copy attempts fail. The source file is
// unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//			OR
//	FileOpCode.CopySourceToDestinationByIoByHardLink()
func (fop FileOperationCode) CopySourceToDestinationByIoByHardLink() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(7)
}

// CopySourceToDestinationByHardLink - Copies the Source File to the Destination
// using one copy mode. The only copy attempt utilizes 'Copy by Hard Link'. If
// this attempted copy fails, an error is returned.  The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByHardLink()
//			OR
//	FileOpCode.CopySourceToDestinationByHardLink()
func (fop FileOperationCode) CopySourceToDestinationByHardLink() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(8)
}

// CopySourceToDestinationByIo - Copies the Source File to the Destination
// using only one copy mode. The only copy attempt is initiated using 'Copy by IO'
// or 'io.Copy'.  If this fails an error is returned. The source file is unaffected.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Usage:
//
//	FileOperationCode(0).CopySourceToDestinationByIo()
//			OR
//	FileOpCode.CopySourceToDestinationByIo()
func (fop FileOperationCode) CopySourceToDestinationByIo() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(9)
}

// CreateSourceDir - Creates the Source Directory
//
// Usage:
//
//	FileOperationCode(0).CreateSourceDir()
//			OR
//	FileOpCode.CreateSourceDir()
func (fop FileOperationCode) CreateSourceDir() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(10)
}

// CreateSourceDirAndFile - Creates the Source Directory and File
//
// Usage:
//
//	FileOperationCode(0).CreateSourceDirAndFile()
//			OR
//	FileOpCode.CreateSourceDirAndFile()
func (fop FileOperationCode) CreateSourceDirAndFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(11)
}

// CreateSourceFile - Creates the Source File
//
// Usage:
//
//	FileOperationCode(0).CreateSourceFile()
//			OR
//	FileOpCode.CreateSourceFile()
func (fop FileOperationCode) CreateSourceFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(12)
}

// CreateDestinationDir - Creates the Destination Directory
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationDir()
//			OR
//	FileOpCode.CreateDestinationDir()
func (fop FileOperationCode) CreateDestinationDir() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(13)
}

// CreateDestinationDirAndFile - Creates the Destination Directory and File
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationDirAndFile()
//			OR
//	FileOpCode.CreateDestinationDirAndFile()
func (fop FileOperationCode) CreateDestinationDirAndFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(14)
}

// CreateDestinationFile - Creates the Destination File
//
// Usage:
//
//	FileOperationCode(0).CreateDestinationFile()
//			OR
//	FileOpCode.CreateDestinationFile()
func (fop FileOperationCode) CreateDestinationFile() FileOperationCode {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return FileOperationCode(15)
}

// IsValid - If the value of the current FileOperationCode is 'invalid',
// this method will return an error. If the FileOperationCode is 'valid',
// this method will return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
func (fop FileOperationCode) IsValid() error {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	fileOperationCodeMapsLock.Lock()

	defer fileOperationCodeMapsLock.Unlock()

	_, ok := mValidFileOperationCodeToString[fop]

	if !ok {
		ePrefix := "FileOperationCode.IsValidInstanceError() "

		return fmt.Errorf("%v\n"+
			"Error: File Operation Code INVALID!.\n"+
			"Unknown Value= %v\n",
			ePrefix,
			int(fop))
	}

	return nil
}

// ParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of FileOperationCode is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
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
//	                       exact match. Therefore, 'movesourcefiletodestination' will NOT
//	                       match the enumeration name, 'MoveSourceFileToDestinationFile'.
//
//	                       If 'false' a case-insensitive search is conducted
//	                       for the enumeration name. In this case, 'movesourcefiletodestination'
//	                       will match enumeration name 'MoveSourceFileToDestinationFile'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	FileOperationCode - Upon successful completion, this method will return a new
//	                    instance of FileOperationCode set to the value of the enumeration
//	                    matched by the string search performed on input parameter,
//	                    'valueString'.
//
//	error             - If this method completes successfully, the returned error
//	                    Type is set equal to 'nil'. If an error condition is encountered,
//	                    this method will return an error Type which encapsulates an
//	                    appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//		t, err := FileOperationCode(0).ParseString("MoveSourceFileToDestinationFile", true)
//	                                 OR
//		t, err := FileOperationCode(0).ParseString("movesourcefiletodestination", false)
//
//		For all the cases shown above,
//		t is now equal to FileOperationCode(0).MoveSourceFileToDestinationFile()
func (fop FileOperationCode) ParseString(
	valueString string,
	caseSensitive bool) (
	FileOperationCode,
	error) {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	fileOperationCodeMapsLock.Lock()

	defer fileOperationCodeMapsLock.Unlock()

	ePrefix := "FileOperationCode.ParseString() "

	if len(valueString) < 3 {
		return FileOperationCode(0),
			fmt.Errorf("%v\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"valueString= '%v'\n",
				ePrefix,
				valueString)
	}

	var ok bool
	var result FileOperationCode

	if caseSensitive {

		result, ok = mFileOperationCodeStringToCode[valueString]

		if !ok {

			return FileOperationCode(0),
				fmt.Errorf("%v\n"+
					"'valueString' did NOT MATCH a FileOperationCode.\n"+
					"valueString= '%v'\n",
					ePrefix,
					valueString)
		}

	} else {

		result, ok = mFileOperationCodeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {

			return FileOperationCode(0),
				fmt.Errorf("%v\n"+
					"'valueString' did NOT MATCH a FileOperationCode.\n"+
					"valueString='%v'\n",
					ePrefix,
					valueString)
		}
	}

	return result, nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FileOperationCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//	string - The string label or description for the current enumeration
//	         value. If, the FileOperationCode value is invalid, this method
//	         will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FileOperationCode(0).MoveSourceFileToDestinationFile()
//	str := t.String()
//	    str is now equal to "MoveSourceFileToDestinationFile"
func (fop FileOperationCode) String() string {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	fileOperationCodeMapsLock.Lock()

	defer fileOperationCodeMapsLock.Unlock()

	str, ok := mFileOperationCodeToString[fop]

	if !ok {
		return ""
	}

	return str
}

// Value - This is a utility method which is not part of the
// enumerations supported by this type. It returns the numeric
// value of the enumeration associated with the current FileOperationCode
// instance.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
func (fop FileOperationCode) Value() int {

	enumFileOperationCodeLock.Lock()

	defer enumFileOperationCodeLock.Unlock()

	return int(fop)
}

// FileOpCode - Internal or private global variable of type FileOperationCode.
//
// Usage:
//
//	FileOpCode.None()
//	FileOpCode.MoveSourceFileToDestinationFile()
//	FileOpCode.MoveSourceFileToDestinationDir()
//	FileOpCode.DeleteDestinationFile()
//	FileOpCode.DeleteSourceFile()
//	FileOpCode.DeleteSourceAndDestinationFiles()
//	FileOpCode.CopySourceToDestinationByHardLinkByIo()
//	FileOpCode.CopySourceToDestinationByIoByHardLink()
//	FileOpCode.CopySourceToDestinationByHardLink()
//	FileOpCode.CopySourceToDestinationByIo()
//	FileOpCode.CreateSourceDir()
//	FileOpCode.CreateSourceDirAndFile()
//	FileOpCode.CreateSourceFile()
//	FileOpCode.CreateDestinationDir()
//	FileOpCode.CreateDestinationDirAndFile()
//	FileOpCode.CreateDestinationFile()
const FileOpCode = FileOperationCode(0)
