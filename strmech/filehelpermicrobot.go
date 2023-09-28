package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fileHelperMicrobot struct {
	lock *sync.Mutex
}

// validateSourceFile
//
// This method will validate source file path and file
// name string.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFile					string
//
//		This string holds the path and/or file name of
//	 	the source file. The source file will be
//	 	validated to verify the following:
//
//		(1)	The file must exist on disk.
//
//		(2)	The file must be a 'regular' file.
//
//		(3) The file must NOT be a 'directory'.
//
//	sourceFileLabel				string
//
//		The name or label associated with input parameter
//		'sourceFile' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "sourceFile" will be
//		automatically applied.
//
//	errorOnEmptyFile			bool
//
//		If this parameter is set to 'true' and the source
//		file is empty (contains zero bytes), an error will
//		be returned.
//
//		Conversely, if this parameter is set to 'false'
//		and the source file is empty (contains zero
//		bytes), no error will be returned.
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
//	sourceFileAbsPath			string
//
//		The absolute path and file name generated from
//		input parameter 'sourceFile'.
//
//	srcFInfoPlus				FileInfoPlus
//
//		If the path and file name specified by input
//		parameter 'sourceFile' actually exists on disk,
//		this returned instance of FileInfoPlus will be
//		populated with detailed information on that
//		file.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fHelperMicrobot *fileHelperMicrobot) validateSourceFile(
	sourceFile string,
	sourceFileLabel string,
	errorOnEmptyFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	sourceFileAbsPath string,
	srcFInfoPlus FileInfoPlus,
	err error) {

	if fHelperMicrobot.lock == nil {
		fHelperMicrobot.lock = new(sync.Mutex)
	}

	fHelperMicrobot.lock.Lock()

	defer fHelperMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileHelperMacroBot." +
		"validateSourceFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return sourceFileAbsPath, srcFInfoPlus, err
	}

	if len(sourceFileLabel) == 0 {

		sourceFileLabel = "sourceFile"
	}

	var srcFileDoesExist bool

	fhMolecule := new(fileHelperMolecule)

	sourceFile,
		srcFileDoesExist,
		srcFInfoPlus,
		err = fhMolecule.
		doesPathFileExist(
			sourceFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			sourceFileLabel)

	if err != nil {

		return sourceFileAbsPath, srcFInfoPlus, err
	}

	sourceFileAbsPath = sourceFile

	if !srcFileDoesExist {

		err = fmt.Errorf(
			"%v\n"+
				"Error: Source File DOES NOT EXIST!\n"+
				"%v= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFile)

		return sourceFileAbsPath, srcFInfoPlus, err
	}

	if srcFInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Source File, '%v', is a 'Directory'\n"+
			"and NOT a file!\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFile)

		return sourceFileAbsPath, srcFInfoPlus, err
	}

	if !srcFInfoPlus.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)

		err = fmt.Errorf("%v\n"+
			"Error non-regular source file ='%v'\n"+
			"source file Mode='%v'\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			srcFInfoPlus.Name(),
			srcFInfoPlus.Mode().String(),
			sourceFileLabel,
			sourceFile)

		return sourceFileAbsPath, srcFInfoPlus, err
	}

	if errorOnEmptyFile &&
		srcFInfoPlus.Size() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Source File, '%v', is an EMPTY File!\n"+
			"%v contains zero (0) bytes.\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFileLabel,
			sourceFile)

	}

	return sourceFileAbsPath, srcFInfoPlus, err
}
