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
// This method will validate the source path and
// file name string.
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
//		source file.
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

// validateDestinationFile
//
// This method will validate destination path and
// file name string.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationFile				string
//
//		This string holds the path and/or file name of
//	 	the destination file. The destination file will
//	 	be validated to verify the following:
//
//		(1)	The file must be a 'regular' file.
//
//		(2) The file must NOT be a 'directory'.
//
//	destinationFileLabel		string
//
//		The name or label associated with input parameter
//		'destinationFile' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "destinationFile" will
//		be automatically applied.
//
//	createDirectoryPathIfNotExist	bool
//
//		If the directory path element of parameter
//		'destinationFile' does not exist on an attached
//		storage drive and this parameter is set to
//		'true', this method will attempt to create
//		the directory path.
//
//		If 'createDirectoryPathIfNotExist' is set to
//		'false', and the directory path element of
//		parameter 'pathFileName' does not exist on an
//		attached storage drive, an error will be
//		returned.
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
//	destinationFileAbsPath		string
//
//		The absolute path and file name generated from
//		input parameter 'destinationFile'.
//
//	destFileDoesExist			bool
//
//		If this return parameter is set to 'true', it
//		signals that the file identified by input
//		parameter 'destinationFile' already exists on
//		an attached storage disk.
//
//	destFInfoPlus				FileInfoPlus
//
//		If the path and file name specified by input
//		parameter 'destinationFile' actually exists on an
//		attached storage disk, this returned instance of
//		FileInfoPlus will be populated with detailed
//		information on that destination file.
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
func (fHelperMicrobot *fileHelperMicrobot) validateDestinationFile(
	destinationFile string,
	destinationFileLabel string,
	createDirectoryPathIfNotExist bool,
	errPrefDto *ePref.ErrPrefixDto) (
	destinationFileAbsPath string,
	destFileDoesExist bool,
	destFInfoPlus FileInfoPlus,
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

		return destinationFileAbsPath,
			destFileDoesExist,
			destFInfoPlus,
			err
	}

	if len(destinationFileLabel) == 0 {

		destinationFileLabel = "destinationFile"
	}

	var dstFileInfo FileInfoPlus

	var fhMolecule = new(fileHelperMolecule)

	destinationFile,
		destFileDoesExist,
		dstFileInfo,
		err = fhMolecule.
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			destinationFileLabel)

	if err != nil {

		return destinationFileAbsPath,
			destFileDoesExist,
			destFInfoPlus,
			err
	}

	if destFileDoesExist && dstFileInfo.Mode().IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: '%v' is a Directory and NOT a File!\n"+
			"%v='%v'",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile)

		return destinationFileAbsPath,
			destFileDoesExist,
			destFInfoPlus,
			err
	}

	if destFileDoesExist && !dstFileInfo.Mode().IsRegular() {
		err = fmt.Errorf("%v\n"+
			"Error: Destination File is NOT a 'Regular' File!\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFile)

		return destinationFileAbsPath,
			destFileDoesExist,
			destFInfoPlus,
			err
	}

	var err2 error

	if !destFileDoesExist {
		// The destination path and/or file does NOT
		// exist on disk.
		var directoryPath, fileNameExt string
		var bothAreEmpty, dirPathExists bool

		directoryPath,
			fileNameExt,
			bothAreEmpty,
			err2 = new(fileHelperDirector).
			getPathAndFileNameExt(
				destinationFile,
				destinationFileLabel,
				ePrefix.XCpy("<-"+destinationFileLabel))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Destination File parameter '%v' is invalid!\n"+
				"An error occurred while breaking '%v'\n"+
				"into directory path, file name and file extension\n"+
				"components.\n"+
				"Error = \n%v\n",
				funcName,
				destinationFileLabel,
				destinationFileLabel,
				err2.Error())

			return destinationFileAbsPath,
				destFileDoesExist,
				destFInfoPlus,
				err
		}

		if len(fileNameExt) == 0 || bothAreEmpty {

			err = fmt.Errorf("%v\n"+
				"Error: Input parameter destination file, '%v',\n"+
				"is invalid! No valid file name could be extracted\n"+
				"from '%v'.\n"+
				"%v= '%v'\n"+
				"Directory Path Element= '%v'\n"+
				"File Name Element= '%v'\n",
				ePrefix.String(),
				destinationFileLabel,
				destinationFileLabel,
				destinationFileLabel,
				destinationFile,
				directoryPath,
				fileNameExt)

			return destinationFileAbsPath,
				destFileDoesExist,
				destFInfoPlus,
				err
		}

		dirPathExists,
			dstFileInfo,
			err2 = new(fileHelperAtom).doesDirectoryExist(
			directoryPath,
			destinationFileLabel+" Dir Path",
			ePrefix.XCpy(destinationFileLabel+" Dir Path"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fileHelperAtom.doesDirectoryExist()\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				funcName,
				destinationFileLabel,
				destinationFile,
				err2.Error())

			return destinationFileAbsPath,
				destFileDoesExist,
				destFInfoPlus,
				err
		}

		if !dstFileInfo.IsDir() {

			err = fmt.Errorf("%v\n"+
				"Error: The directory path extracted from destination\n"+
				"file path, '%v', is NOT a valid directory.\n"+
				"Destination File '%v' is therefore invalid!\n"+
				"%v= '%v'\n",
				ePrefix.String(),
				destinationFileLabel,
				destinationFileLabel,
				destinationFileLabel,
				destinationFile)

			return destinationFileAbsPath,
				destFileDoesExist,
				destFInfoPlus,
				err
		}

		if !dirPathExists {
			// The Destination File directory path does NOT
			// exist on an attached storage volume.

			if createDirectoryPathIfNotExist {

				err2 = new(fileHelperPreon).makeDirAll(
					directoryPath,
					destinationFileLabel+" directoryPath",
					ePrefix)

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Attempted creation of directory path for the"+
						"destination file, '%v', failed!\n"+
						"%v= '%v'\n"+
						"%v Directory Path= '%v'\n"+
						"Error returned by fileHelperMechanics.makeDirAll()\n"+
						"Error=\n%v\n",
						funcName,
						destinationFileLabel,
						destinationFileLabel,
						destinationFile,
						destinationFileLabel,
						directoryPath,
						err2.Error())

					return destinationFileAbsPath,
						destFileDoesExist,
						destFInfoPlus,
						err
				}

			} else {
				// The Path File Name Directory DOES NOT EXIST
				// on an attached storage drive and
				// createDirectoryPathIfNotExist = 'false'.

				err = fmt.Errorf("%v\n"+
					"Error: The Destination File, %v, Directory\n"+
					"does NOT exist on an attached storage drive and\n"+
					"Input Parameter 'createDirectoryPathIfNotExist'\n"+
					"was set to 'false'. Therefore the file cannot be\n"+
					"opened.\n"+
					"%v= '%v\n"+
					"%v Directory = '%v'\n",
					ePrefix.String(),
					destinationFileLabel,
					destinationFileLabel,
					destinationFile,
					destinationFileLabel,
					directoryPath)

				return destinationFileAbsPath,
					destFileDoesExist,
					destFInfoPlus,
					err
			}
		}
	}

	return destinationFileAbsPath,
		destFileDoesExist,
		destFInfoPlus,
		err
}
