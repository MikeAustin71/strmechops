package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type fileHelperMechanics struct {
	lock *sync.Mutex
}

// copyFileByIo
//
// Copies file from source path and file name to
// destination path and file name.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this
// method does NOT rely on the creation of symbolic
// links. Instead, a new destination file is created and
// the contents of the source file are written to the new
// destination file using "io.Copy()".
//
// "io.Copy()" is the only method used to copy the
// designated source file. If this method fails, an error
// is returned.
//
// If source file is equivalent to the destination file,
// no action will be taken and no error will be returned.
//
// If the destination file does not exist, this method
// will create. However, it will NOT create the
// destination directory. If the destination directory
// does NOT exist, this method will abort the copy
// operation and return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFile					string
//
//		This string holds the path and/or file name of
//	 	the source file. This source file will be copied
//		to the destination file.
//
//	destinationFile				string
//
//		This string holds the path and/or the file name
//		of the destination file. The source file taken
//		from input parameter 'sourceFile' will be copied
//		to this destination file.
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
func (fileHelpMech *fileHelperMechanics) copyFileByIo(
	sourceFile string,
	destinationFile string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"copyFileByIo()",
		"")

	if err != nil {
		return err
	}

	var err2, err3 error
	var srcFileDoesExist, dstFileDoesExist bool
	var srcFInfo, dstFileInfo FileInfoPlus

	fhMolecule := new(fileHelperMolecule)

	sourceFile,
		srcFileDoesExist,
		srcFInfo,
		err = fhMolecule.
		doesPathFileExist(
			sourceFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"sourceFile")

	if err != nil {
		return err
	}

	if !srcFileDoesExist {

		err = fmt.Errorf(
			"%v\n"+
				"Error: Source File DOES NOT EXIST!\n"+
				"sourceFile='%v'\n",
			ePrefix.String(),
			sourceFile)

		return err
	}

	destinationFile,
		dstFileDoesExist,
		dstFileInfo,
		err = fhMolecule.
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"destinationFile")

	if err != nil {
		return err
	}

	var areSameFile bool

	areSameFile,
		err2 = new(fileHelperNanobot).areSameFile(
		sourceFile,
		destinationFile,
		ePrefix.XCpy(
			"areSameFile<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error occurred during path file name comparison.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			sourceFile, destinationFile,
			err2.Error())

		return err
	}

	if areSameFile {

		err = fmt.Errorf("%v\n"+
			"Error: The source and destination file\n"+
			"are the same. They are equivalent.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n",
			ePrefix.String(),
			sourceFile,
			destinationFile)

		return err
	}

	if srcFInfo.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Source File is 'Directory' and NOT a file!\n"+
			"Source File='%v'\n",
			ePrefix.String(),
			sourceFile)

		return err
	}

	if !srcFInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)

		err = fmt.Errorf("%v\n"+
			"Error non-regular source file ='%v'\n"+
			"source file Mode='%v'\n",
			ePrefix.String(),
			srcFInfo.Name(),
			srcFInfo.Mode().String())

		return err
	}

	if dstFileDoesExist && dstFileInfo.Mode().IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: 'destinationFile' is a Directory and NOT a File!\n"+
			"destinationFile='%v'",
			ePrefix.String(),
			destinationFile)

		return err
	}

	if dstFileDoesExist && !dstFileInfo.Mode().IsRegular() {
		err = fmt.Errorf("%v\n"+
			"Error: 'destinationFile' is NOT a 'Regular' File!\n"+
			"destinationFile='%v'\n",
			ePrefix.String(),
			destinationFile)

		return err
	}

	// If the destination file does NOT exist, this is not a problem
	// since it will be created later. If the destination 'Path' does
	// not exist, an error return will be triggered.

	// Create a new destination file and copy source
	// file contents to the destination file.

	// First, open the source file
	inSrcPtr, err2 := os.Open(sourceFile)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from os.Open(sourceFile) sourceFile='%v'\n"+
			"Error='%v'",
			ePrefix.String(),
			sourceFile,
			err2.Error())
		return err
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outDestPtr, err2 := os.Create(destinationFile)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from os.Create(destinationFile)\n"+
			"destinationFile='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		_ = inSrcPtr.Close()

		return err
	}

	bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

	if err2 != nil {

		_ = inSrcPtr.Close()

		_ = outDestPtr.Close()

		err = fmt.Errorf("%v\n"+
			"Error returned from io.Copy(destination, source)\n"+
			"destination='%v'\n"+
			"source='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			sourceFile,
			err2.Error())

		return err
	}

	errs := make([]error, 0)

	// flush file buffers inSrcPtr memory
	err2 = outDestPtr.Sync()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Sync()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile, err2.Error())

		errs = append(errs, err3)
	}

	err2 = inSrcPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from inSrcPtr.Close()\n"+
			"inSrcPtr=source='%v'\nError='%v'\n",
			ePrefix.String(),
			sourceFile,
			err2.Error())

		errs = append(errs, err3)
	}

	inSrcPtr = nil

	err2 = outDestPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Close()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		errs = append(errs, err3)
	}

	outDestPtr = nil

	if len(errs) > 0 {

		return new(StrMech).ConsolidateErrors(errs)
	}

	_,
		dstFileDoesExist,
		dstFileInfo,
		err2 = fhMolecule.
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.None(), // Do NOT alter path
			ePrefix,
			"destinationFile")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy IO operation, destinationFile "+
			"generated non-path error!\n"+
			"destinationFile='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		return err
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = 'destinationFile' = '%v'\n",
			ePrefix.String(),
			destinationFile)

		return err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file!\n"+
			"Source File Bytes='%v'\n"+
			"Bytes Coped='%v'\n"+
			"Source File=sourceFile='%v'\n"+
			"Destination File=destinationFile='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			sourceFile,
			destinationFile)

		return err
	}

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File Bytes='%v'\n"+
			"Destination File Bytes='%v'\n"+
			"Source File=sourceFile='%v'\n"+
			"Destination File=destinationFile='%v'\n",
			ePrefix.String(),
			srcFileSize,
			dstFileInfo.Size(),
			sourceFile,
			destinationFile)

	}

	return err
}

// copyFileByLink
//
// Copies a file from source to destination by means of
// creating a 'hard link' to the source file,
// "os.Link(src, dst)".
//
// Note: This method of copying files does NOT create a
// new destination file and write the contents of the
// source file to destination file. (See CopyFileByIo
// Below).  Instead, this method performs the copy
// operation by creating a hard symbolic link to the
// source file.
//
// By creating a 'linked' file, changing the contents
// of one file will be reflected in the second. The
// two linked files are 'mirrors' of each other.
//
// Consider using CopyFileByIo() if this 'mirror' feature
// causes problems.
//
// "os.Link(src, dst)" is the only method employed to
// copy a designated file. If "os.Link(src, dst)" fails,
// an err is returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// REQUIREMENT: The destination Path must previously
// exist. The destination file need NOT exist as it will
// be created. If the destination file currently exists,
// it will first be deleted and a new linked file will be
// created.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file. which will be copied to the
//	 	destination file passed through input parameter
//	 	'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file passed through
//	 	input parameter 'src' which will be copied to the
//	 	destination file identified by this parameter,
//	 	'dst'.
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
func (fileHelpMech *fileHelperMechanics) copyFileByLink(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"copyFileByLink()",
		"")

	if err != nil {
		return err
	}

	var err2 error
	var srcFileDoesExist, dstFileDoesExist bool
	var srcFInfo, dstFInfo FileInfoPlus

	fHelpMolecule := fileHelperMolecule{}
	src,
		srcFileDoesExist,
		srcFInfo,
		err = fHelpMolecule.doesPathFileExist(
		src,
		PreProcPathCode.AbsolutePath(), // Covert to Absolute Path
		ePrefix,
		"src")

	if err != nil {
		return err
	}

	dst,
		dstFileDoesExist,
		dstFInfo,
		err = fHelpMolecule.doesPathFileExist(
		dst,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"dst")

	if err != nil {
		return err
	}

	areSameFile, err2 :=
		new(fileHelperNanobot).areSameFile(
			src,
			dst,
			ePrefix.XCpy("areSameFile<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error occurred during path file name comparison.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n"+
			"Error='%v'\n",
			ePrefix,
			src,
			dst,
			err2.Error())

		return err
	}

	if areSameFile {

		err = fmt.Errorf("%v\n"+
			"Error: The source and destination file"+
			" are the same - equivalent.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n",
			ePrefix.String(),
			src,
			dst)

		return err
	}

	if !srcFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'src' file DOES NOT EXIST!\n"+
			"src='%v'\n",
			ePrefix.String(),
			src)

		return err
	}

	if srcFInfo.IsDir() {

		err = fmt.Errorf("%v\n"+
			"ERROR: Source File (src) is a 'Directory' NOT A FILE!\n"+
			"Source File (src)='%v'\n",
			ePrefix.String(),
			src)

		return err
	}

	if !srcFInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)

		err = fmt.Errorf("%v\n"+
			"Error: Non-regular source file.\n"+
			"Source File Name='%v'\n"+
			"Source File Mode='%v'\n",
			ePrefix.String(),
			srcFInfo.Name(),
			srcFInfo.Mode().String())

		return err
	}

	// If the destination file does NOT exist - this is not a problem
	// because the destination file will be created later.

	if dstFileDoesExist {
		// The destination file exists. This IS a problem. Link will
		// fail when attempting to create a link to an existing file.

		if dstFInfo.IsDir() {

			err = fmt.Errorf("%v\n"+
				"Error: The destination file ('dst') is NOT A FILE.\n"+
				"It is a DIRECTORY!\n"+
				"Destination File ('dst') = '%v'\n",
				ePrefix.String(),
				dst)

			return err
		}

		if !dstFInfo.Mode().IsRegular() {

			err = fmt.Errorf("%v\n"+
				"Error: The destination file ('dst') is NOT A REGULAR FILE.\n"+
				"Destination File ('dst') = '%v'\n",
				ePrefix.String(),
				dst)

			return err
		}

		err2 = os.Remove(dst)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: The target destination file exists and could NOT be deleted!\n"+
				"destination file='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dst,
				err2.Error())

			return err
		}

		dst,
			dstFileDoesExist,
			_,
			err = fHelpMolecule.
			doesPathFileExist(
				dst,
				PreProcPathCode.None(), // Apply no pre-processing conversion to 'dst'
				ePrefix,
				"dst")

		if err != nil {
			return err
		}

		if dstFileDoesExist {

			err = fmt.Errorf("%v\n"+
				"Error: Deletion of preexisting "+
				"destination file failed!\n"+
				"The copy link operation cannot proceed!\n"+
				"destination file='%v' ",
				ePrefix.String(),
				dst)

			return err
		}
	}

	err2 = os.Link(src, dst)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"os.Link(src, dst) FAILED!\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			src,
			dst,
			err2.Error())

		return err
	}

	dst,
		dstFileDoesExist,
		_,
		err2 = fHelpMolecule.
		doesPathFileExist(
			dst,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"dst")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link Operation, a non-path error was returned on 'dst'.\n"+
			"Error='%v'",
			ePrefix,
			err2.Error())

		return err
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link Operation, the destination file DOES NOT EXIST!\n"+
			"Destination File= dst = %v",
			ePrefix.String(),
			dst)

	}

	return err
}

// deleteFilesWalkDirectory
//
// This method 'walks' the directory tree searching for
// files which match the file selection criteria
// specified by input parameter 'fileSelectCriteria'.
// When a file matching said 'fileSelectCriteria' is
// found, that file is deleted.
//
// This method returns file information on files deleted.
//
// If a file matches the File Selection Criteria
// ('fileSelectCriteria') it is deleted and its file
// information is recorded in the returned instance of
// DirectoryDeleteFileInfo,
// 'DirectoryDeleteFileInfo.DeletedFiles'.
//
// By the way, if ALL the file selection criterion are
// set to zero values or 'Inactive', then ALL FILES in
// the directory are selected, deleted and returned in
// the field, 'DirectoryDeleteFileInfo.DeletedFiles'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If all the file selection criterion in the
//	FileSelectionCriteria object are 'Inactive' or
//	'Not Set' (set to their zero or default values), then
//	all the files processed in the directory tree WILL BE
//	DELETED!
//
//	Information on the deleted file will be returned in the
//	file manager collection,
//	'DirectoryDeleteFileInfo.DeletedFiles'.
//
//		Example:
//			FileNamePatterns  = ZERO Length Array
//			filesOlderThan    = time.Time{}
//			filesNewerThan    = time.Time{}
//
//	In this example, all the selection criterion are
//	'Inactive' and therefore all the files encountered
//	in the target directory will be selected and returned
//	as 'Found Files'.
//
//	This same effect can be achieved by simply creating an
//	empty file selection instance:
//
//		FileSelectionCriteria{}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	startPath					string
//
//		A string consisting of the starting path or
//		directory from which the file search operation
//		will commence.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories the
//		'startPath' will be searched, and all files
//		within those directories WILL BE DELETED.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string    // An array of strings containing File Name Patterns
//		  FilesOlderThan       time.Time   // Match files with older modification date times
//		  FilesNewerThan       time.Time   // Match files with newer modification date times
//		  SelectByFileMode     FilePermissionConfig // Match file mode (os.FileMode).
//		  SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//		}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria Type are described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or more search
//			patterns. If a file name matches any one of the search
//			pattern strings, it is deemed to be a 'match' for the
//			search pattern criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file modification
//			date times in order to determine whether the file is
//			older than the 'FilesOlderThan' file selection
//			criterion. If the file is older than the
//			'FilesOlderThan' date time, that file is considered
//			a 'match'	for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file modification
//			date time in order to determine whether the file is newer
//			than the 'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than the
//			'FilesNewerThan' date time, that file is considered a
//			'match' for this file selection criterion.
//
//			If the value of 'FilesNewerThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an os.FileMode. The
//			file selection criterion allows for the selection of files
//			by File Mode.
//
//			File modes are compared to the value of 'SelectByFileMode'.
//			If the File Mode for a given file is equal to the value of
//	 		'SelectByFileMode', that file is considered to be a 'match'
//	 		for this file selection criterion. Examples for setting
//	 		SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the file selection
//		criteria above are applied in determining a 'match' for file
//		selection purposes. 'SelectCriterionMode' may be set to one of
//		two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will not be judged as 'selected' unless all
//			the active selection criterion are satisfied. In other words, if
//			three active search criterion are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//			selected unless it has satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection criterion is satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will be selected if any one of the active file
//			selection criterion is satisfied. In other words, if three active
//			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//			and 'FilesNewerThan', then a file will be selected if it satisfies any
//			one of the three criterion in this example.
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
//	DirectoryDeleteFileInfo
//
//		If successful, files matching the file selection criteria input
//		parameter shown above will be deleted and returned in a
//		'DirectoryDeleteFileInfo' object. The file manager
//		'DirectoryDeleteFileInfo.DeletedFiles' contains information on all files
//		deleted during this operation.
//
//		Note:
//		It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//		to determine if any internal system errors were encountered while processing
//		the directory tree.
//
//		__________________________________________________________________________________________________
//
//		type DirectoryDeleteFileInfo struct {
//
//			StartPath             string
//
//		  		The starting path or directory for the file
//		  		search.
//
//			Directories           DirMgrCollection
//
//				Directory Manager instances found during the
//				directory tree search.
//
//			DeletedFiles          FileMgrCollection
//
//				Contains File Managers for Deleted Files matching
//				file selection criteria.
//
//			ErrReturns            []error
//
//				Internal System errors encountered during the search
//				and file deletion operations. This includes type
//				*PathError objects created by low level system
//				function calls.
//
//			FileSelectCriteria    FileSelectionCriteria
//
//				The File Selection Criteria submitted as an
//				input parameter to this method.
//		}
//
//		__________________________________________________________________________________________________
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
func (fileHelpMech *fileHelperMechanics) deleteFilesWalkDirectory(
	startPath string,
	fileSelectCriteria FileSelectionCriteria,
	errPrefDto *ePref.ErrPrefixDto) (
	DirectoryDeleteFileInfo,
	error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deleteFilesInfo := DirectoryDeleteFileInfo{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"deleteFilesWalkDirectory()",
		"")

	if err != nil {
		return deleteFilesInfo, err
	}

	errCode := 0

	fHelperElectron := new(fileHelperElectron)

	errCode, _, startPath =
		fHelperElectron.isStringEmptyOrBlank(startPath)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startPath' is an empty string!",
			ePrefix.String())

		return deleteFilesInfo, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startPath' consists of blank spaces!\n",
			ePrefix.String())

		return deleteFilesInfo, err
	}

	startPath = new(fileHelperAtom).adjustPathSlash(startPath)

	strLen := len(startPath)

	if startPath[strLen-1] == os.PathSeparator {
		startPath = startPath[0 : strLen-1]
	}

	startPath,
		err = new(fileHelperProton).makeAbsolutePath(
		startPath,
		ePrefix.XCpy("startPath<-"))

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath(startPath).\n"+
			"startPath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			startPath,
			err.Error())

		return deleteFilesInfo, err
	}

	if !new(fileHelperNanobot).doesFileExist(startPath) {

		err = fmt.Errorf("%v\n"+
			"Error - startPath DOES NOT EXIST!\n"+
			"startPath='%v'",
			ePrefix.String(),
			startPath)

		return deleteFilesInfo, err
	}

	deleteFilesInfo.StartPath = startPath

	deleteFilesInfo.DeleteFileSelectCriteria =
		fileSelectCriteria

	var err2 error

	err2 = fp.Walk(
		deleteFilesInfo.StartPath,
		new(fileHelperMolecule).makeFileHelperWalkDirDeleteFilesFunc(
			&deleteFilesInfo))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fp.Walk(deleteFilesInfo.StartPath - \n"+
			"MakeFileHelperWalkDirDeleteFilesFunc"+
			"(&deleteFilesInfo)).\n"+
			"startPath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			startPath,
			err.Error())

		return deleteFilesInfo, err
	}

	return deleteFilesInfo, nil
}

// FindFilesWalkDirectory
//
// This method returns file information on files residing
// in a specified directory tree identified by the input
// parameter, 'startPath'.
//
// This method 'walks the directory tree' locating all
// files in the directory tree which match the file
// selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is
// included in the returned field:
//
//	'DirectoryTreeInfo.FoundFiles'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If all the file selection criterion in the
//	FileSelectionCriteria object are 'Inactive' or
//	'Not Set' (set to their zero or default values), then
//	all the files processed in the directory tree will be
//	selected and returned as 'Found Files'.
//
//	  Example:
//	     FileNamePatterns  = ZERO Length Array
//	     filesOlderThan    = time.Time{}
//	     filesNewerThan    = time.Time{}
//
//	  In this example, all the selection criterion are
//	  'Inactive' and therefore all the files encountered
//	  in the target directory will be selected and returned
//	  as 'Found Files' stored in the following member
//	  variable:
//
//		'DirectoryTreeInfo.FoundFiles'.
//
//	  This same effect can be achieved by simply creating
//	  an empty file selection instance:
//
//	          FileSelectionCriteria{}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	startPath					string
//
//		A string consisting of the starting path or
//		directory from which the find files search
//		operation will commence.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories and
//		files will be returned from the 'startPath'
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or more search
//			patterns. If a file name matches any one of the search
//			pattern strings, it is deemed to be a 'match' for the
//			search pattern criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file modification
//			date times in order to determine whether the file is
//			older than the 'FilesOlderThan' file selection
//			criterion. If the file is older than the
//			'FilesOlderThan' date time, that file is considered
//			a 'match'	for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file modification
//			date time in order to determine whether the file is newer
//			than the 'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than the
//			'FilesNewerThan' date time, that file is considered a
//			'match' for this file selection criterion.
//
//			If the value of 'FilesNewerThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an os.FileMode. The
//			file selection criterion allows for the selection of files
//			by File Mode.
//
//			File modes are compared to the value of 'SelectByFileMode'.
//			If the File Mode for a given file is equal to the value of
//	 		'SelectByFileMode', that file is considered to be a 'match'
//	 		for this file selection criterion. Examples for setting
//	 		SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the file selection
//		criteria above are applied in determining a 'match' for file
//		selection purposes. 'SelectCriterionMode' may be set to one of
//		two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will not be judged as 'selected' unless all
//			the active selection criterion are satisfied. In other words, if
//			three active search criterion are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//			selected unless it has satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection criterion is satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will be selected if any one of the active file
//			selection criterion is satisfied. In other words, if three active
//			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//			and 'FilesNewerThan', then a file will be selected if it satisfies any
//			one of the three criterion in this example.
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
// ------------------------------------------------------------------------
//
// Return Values:
//
//	DirectoryTreeInfo
//
//		If successful, files matching the file selection
//		criteria input parameter shown above will be
//		returned in a 'DirectoryTreeInfo' object. The
//		field 'DirectoryTreeInfo.FoundFiles' contains
//		information on all the files in the specified
//		path or directory tree which match the file
//		selection criteria.
//
//		Note:
//
//		It's a good idea to check the returned field
//		'DirectoryTreeInfo.ErrReturns' to determine if
//		any internal system errors were encountered while
//		processing the directory tree.
//
//	        ________________________________________________
//
//	        type DirectoryTreeInfo struct {
//
//	          StartPath             string
//	          	The starting path or directory for the
//	          	file search.
//
//	          dirMgrs               []DirMgr
//	          	Directories found during directory tree
//	          	search are stored here.
//
//	          FoundFiles            []FileWalkInfo
//				Files matching the file search selection
//				criteria are stored here.
//
//	          ErrReturns            []string
//				Internal System errors encountered during
//				the file search operation are stored here
//				as text messages.
//
//	          FileSelectCriteria    FileSelectionCriteria
//	          	The File Selection Criteria submitted as an
//				input parameter to this method.
//	        }
//
//	        ________________________________________________
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
//		BE ADVISED
//
//		It's a good idea to check the returned field
//		'DirectoryTreeInfo.ErrReturns' to determine if
//		any internal system errors were encountered while
//		processing the directory tree.
func (fileHelpMech *fileHelperMechanics) findFilesWalkDirectory(
	startPath string,
	fileSelectCriteria FileSelectionCriteria,
	errPrefDto *ePref.ErrPrefixDto) (
	DirectoryTreeInfo,
	error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	findFilesInfo := DirectoryTreeInfo{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"findFilesWalkDirectory()",
		"")

	if err != nil {
		return findFilesInfo, err
	}

	errCode := 0

	errCode, _, startPath =
		new(fileHelperElectron).
			isStringEmptyOrBlank(startPath)

	if errCode == -1 {

		return findFilesInfo,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'startPath' is an empty string!\n",
				ePrefix.String())
	}

	if errCode == -2 {

		return findFilesInfo,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'startPath' consists of blank spaces!\n",
				ePrefix.String())
	}

	startPath = new(fileHelperAtom).
		removePathSeparatorFromEndOfPathString(
			startPath)

	startPath, err = new(fileHelperProton).
		makeAbsolutePath(
			startPath,
			ePrefix)

	if err != nil {

		return findFilesInfo,
			fmt.Errorf("%v\n"+
				"Error returned by fh.MakeAbsolutePath(startPath).\n"+
				"startPath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				startPath,
				err.Error())
	}

	if !new(fileHelperNanobot).doesFileExist(startPath) {

		return findFilesInfo,
			fmt.Errorf("%v\n"+
				"Error - startPath DOES NOT EXIST!\n"+
				"startPath='%v'\n",
				ePrefix.String(),
				startPath)
	}

	findFilesInfo.StartPath = startPath

	findFilesInfo.FileSelectCriteria = fileSelectCriteria

	err = fp.Walk(findFilesInfo.StartPath,
		new(fileHelperMolecule).
			makeFileHelperWalkDirFindFilesFunc(
				&findFilesInfo))

	if err != nil {

		return findFilesInfo,
			fmt.Errorf("%v\n"+
				"Error returned from fp.Walk(findFilesInfo.StartPath,\n"+
				"fh.MakeFileHelperWalkDirFindFilesFunc"+
				"(&findFilesInfo)).\n"+
				"startPath='%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				startPath,
				err.Error())
	}

	return findFilesInfo, nil
}

// getFileNameWithoutExt
//
// Returns the file name without the path or extension.
//
// If the returned File Name is an empty string, return
// parameter 'isEmpty' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # Usage Examples
//
//	     Actual Path Plus File Name: = "./pathfilego/003_filehelper/common/xt_dirmgr_01_test.go"
//	             Returned File Name: = "dirmgr_01_test"
//
//	Actual File Name Plus Extension: "newerFileForTest_01.txt"
//	             Returned File Name: "newerFileForTest_01"
//
//	Actual File Name Plus Extension: "newerFileForTest_01"
//	             Returned File Name: "newerFileForTest_01"
//
//	Actual File Name Plus Extension: ".gitignore"
//	             Returned File Name: ".gitignore"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string holds the path, file name and file
//		extension. This method will extract the file name
//		from this string and return it to the calling
//		function.
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
//	fName						string
//
//			This return parameter contains the file name
//			extracted from input parameter
//			'pathFileNameExt'.
//
//	isEmpty						bool
//
//		If the returned File Name is an empty string,
//		'isEmpty' is set to 'true'.
//
//	err							error
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
func (fileHelpMech *fileHelperMechanics) getFileNameWithoutExt(
	pathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	fName string,
	isEmpty bool,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	fName = ""
	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"getFileNameWithoutExt()",
		"")

	if err != nil {
		return fName, isEmpty, err
	}

	errCode := 0

	fHelperElectron := new(fileHelperElectron)

	errCode,
		_,
		pathFileNameExt = fHelperElectron.
		isStringEmptyOrBlank(pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' is an empty string!\n",
			ePrefix.String())

		return fName, isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' consists of blank spaces!\n",
			ePrefix.String())

		return fName, isEmpty, err
	}

	testPathFileNameExt := new(fileHelperAtom).
		adjustPathSlash(pathFileNameExt)

	errCode,
		_,
		testPathFileNameExt = fHelperElectron.
		isStringEmptyOrBlank(testPathFileNameExt)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Adjusted path version of 'pathFileNameExt', 'testPathFileNameExt'\n"+
			"is an empty string!\n",
			ePrefix.String())

		return fName, isEmpty, err
	}

	fileNameExt,
		isFileNameExtEmpty,
		err2 :=
		new(fileHelperNanobot).
			getFileNameWithExt(
				testPathFileNameExt,
				ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from getFileNameWithExt(testPathFileNameExt)\n"+
			"testPathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return fName, isEmpty, err
	}

	if isFileNameExtEmpty {

		isEmpty = true

		fName = ""

		return fName, isEmpty, err
	}

	var dotIdxs []int

	dotIdxs,
		err2 = new(fileHelperAtom).
		getDotSeparatorIndexesInPathStr(
			fileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from getDotSeparatorIndexesInPathStr(fileNameExt).\n"+
			"fileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fileNameExt,
			err2.Error())

		return fName, isEmpty, err
	}

	lDotIdxs := len(dotIdxs)

	if lDotIdxs == 1 &&
		dotIdxs[lDotIdxs-1] == 0 {
		// Outlier Case: .gitignore
		fName = fileNameExt[0:]

		if fName == "" {
			isEmpty = true
		} else {
			isEmpty = false
		}

		return fName, isEmpty, err
	}

	// Primary Case: filename.ext
	if lDotIdxs > 0 {
		fName = fileNameExt[0:dotIdxs[lDotIdxs-1]]

		if fName == "" {
			isEmpty = true
		} else {
			isEmpty = false
		}

		return fName, isEmpty, err
	}

	// Secondary Case: filename
	fName = fileNameExt

	if fName == "" {
		isEmpty = true
	} else {
		isEmpty = false
	}

	return fName, isEmpty, err
}

// getPathFromPathFileName
//
// Returns the path from a path and file name string. If
// the returned path is an empty string, return parameter
// 'isEmpty' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # Usage Examples
//
//	pathFileNameExt = ""                  returns isEmpty==true  err==nil
//	pathFileNameExt = "D:\"               returns "D:\"
//	pathFileNameExt = "."                 returns ".\"
//	pathFileNameExt = "..\"               returns "..\"
//	pathFileNameExt = "...\"              returns ERROR
//
//	pathFileNameExt = ".\pathfile\003_filehelper\wt_HowToRunTests.md"
//	                                      returns ".\pathfile\003_filehelper"
//
//	pathFileNameExt = "someFile.go"       returns ""
//	pathFileNameExt = "..\dir1\dir2\.git" returns "..\dir1\dir2"
//	                                       '.git' is assumed to be a file.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt  string
//
//		This is an input parameter. The method expects to
//		receive a single, properly formatted path and
//		file name string delimited by dots ('.') and path
//		Separators ('/' or '\').
//
//		On Windows, if the 'pathFileNameExt' string
//		contains valid volume designations
//		(Example: "D:"), these are returned as part of
//		the path.
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
//	path						string
//
//		This is the directory path extracted from the
//		input parameter 'pathFileNameExt'. If successful,
//		the 'path' string that is returned by this method
//		WILL NOT include a trailing path separator
//		('/' or '\' depending on the os).
//
//		Example Return 'path': "./pathfile/003_filehelper"
//
//	isEmpty						bool
//
//		If the method determines that it cannot extract a
//		valid directory path from input parameter
//		'pathFileNameExt', this boolean value will be set
//		to 'true'. Failure to extract a valid directory
//		path will occur if the input parameter
//		'pathFileNameExt' is not properly formatted as a
//		valid path and file name.
//
//	err							error
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
func (fileHelpMech *fileHelperMechanics) getPathFromPathFileName(
	pathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirPath string,
	isEmpty bool,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	dirPath = ""
	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"getPathFromPathFileName()",
		"")

	if err != nil {
		return dirPath, isEmpty, err
	}

	errCode := 0

	errCode,
		_,
		pathFileNameExt =
		new(fileHelperElectron).isStringEmptyOrBlank(
			pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' is an empty string!\n",
			ePrefix.String())

		return dirPath, isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' consists of blank spaces!\n",
			ePrefix.String())

		return dirPath, isEmpty, err
	}

	var testPathStr string
	var isDirEmpty bool
	var err2 error

	testPathStr,
		isDirEmpty,
		err2 = new(fileHelperNanobot).cleanDirStr(
		pathFileNameExt,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.CleanDirStr(pathFileNameExt).\n"+
			"pathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			pathFileNameExt,
			err2.Error())

		return dirPath, isEmpty, err
	}

	if isDirEmpty {
		dirPath = ""
		isEmpty = true
		err = nil
		return dirPath, isEmpty, err
	}

	lTestPathStr := len(testPathStr)

	if lTestPathStr == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: AdjustPathSlash was applied to 'pathStr'.\n"+
			"The 'testPathStr' string is a Zero Length string!\n",
			ePrefix.String())

		return dirPath, isEmpty, err
	}

	fHelperAtom := new(fileHelperAtom)

	var slashIdxs []int

	slashIdxs,
		err2 = fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			testPathStr,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetPathSeparatorIndexesInPathStr(testPathStr).\n"+
			"testPathStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathStr,
			err2.Error())

		return dirPath, isEmpty, err
	}

	lSlashIdxs := len(slashIdxs)

	firstGoodChar, lastGoodChar, err2 :=
		new(fileHelperMolecule).
			getFirstLastNonSeparatorCharIndexInPathStr(
				testPathStr,
				ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr("+
			"testPathStr).\n"+
			"testPathStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathStr,
			err2.Error())

		return dirPath, isEmpty, err
	}

	var dotIdxs []int

	dotIdxs, err2 = fHelperAtom.
		getDotSeparatorIndexesInPathStr(
			testPathStr,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetDotSeparatorIndexesInPathStr(testPathStr).\n"+
			"testPathStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathStr,
			err2.Error())

		return dirPath, isEmpty, err
	}

	lDotIdxs := len(dotIdxs)

	var finalPathStr string

	volName := fp.VolumeName(testPathStr)

	if testPathStr == volName {

		finalPathStr = testPathStr

	} else if strings.Contains(testPathStr, "...") {

		err = fmt.Errorf("%v\n"+
			"Error: PATH CONTAINS INVALID Dot Characters!\n"+
			"testPathStr='%v'\n",
			ePrefix.String(),
			testPathStr)

		return dirPath, isEmpty, err

	} else if firstGoodChar == -1 || lastGoodChar == -1 {

		absPath, err2 := new(fileHelperProton).
			makeAbsolutePath(
				testPathStr,
				ePrefix)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned from fh.MakeAbsolutePath(testPathStr).\n"+
				"testPathStr='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				testPathStr,
				err2.Error())

			return dirPath, isEmpty, err
		}

		if absPath == "" {

			err = fmt.Errorf("%v\n"+
				"Error: Could not convert 'testPathStr' to Absolute path!\n"+
				"testPathStr='%v'\n",
				ePrefix.String(),
				testPathStr)

			return dirPath, isEmpty, err
		}

		finalPathStr = testPathStr

	} else if lSlashIdxs == 0 {

		// No path separators but alphanumeric chars are present
		dirPath = ""
		isEmpty = true
		err = nil
		return dirPath, isEmpty, err

	} else if lDotIdxs == 0 {

		//path separators are present but there are no dots in the string

		if slashIdxs[lSlashIdxs-1] == lTestPathStr-1 {
			// Trailing path separator
			finalPathStr = testPathStr[0:slashIdxs[lSlashIdxs-2]]
		} else {
			finalPathStr = testPathStr
		}

	} else if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] {
		// format: ./dir1/dir2/fileName.ext
		finalPathStr = testPathStr[0:slashIdxs[lSlashIdxs-1]]

	} else if dotIdxs[lDotIdxs-1] < slashIdxs[lSlashIdxs-1] {

		finalPathStr = testPathStr

	} else {

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH STRING.\n"+
			"testPathStr='%v'\n",
			ePrefix.String(),
			testPathStr)

		return dirPath, isEmpty, err
	}

	if len(finalPathStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Processed path is a Zero Length String!\n",
			ePrefix.String())

		return dirPath, isEmpty, err
	}

	//Successfully isolated and returned a valid
	// directory path from 'pathFileNameExt'
	dirPath = finalPathStr

	if len(dirPath) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil

	return dirPath, isEmpty, err
}

// isPathString
//
// Attempts to determine whether a string is a path
// string designating a directory (and not a path file
// name file extension string).
//
// If the path exists on disk, this method will examine
// the associated file information and determine whether
// the path string represents a directory.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//	pathStr						string
//
//		The path string to be analyzed. This will method
//		will determine whether 'pathStr' is a directory
//		path.
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
// ------------------------------------------------------------------------
//
// Return Values:
//
//	isPathStr       			bool
//
//		If the input parameter, 'pathStr' is determined
//		to be a directory path, this return value is set
//		to 'true'. Here, a 'directory path' is defined as
//		a true directory and the path does NOT contain a
//		file name.
//
//	cannotDetermine bool
//
//		If the method cannot determine whether the input
//		parameter 'pathStr' is a valid directory path,
//		this return value will be set to 'true'.
//
//		The 'cannotDetermine=true' condition occurs with
//		path names like 'D:\DirA\common'. The method
//		cannot determine whether 'common' is a file name
//		or a directory name.
//
//
//	testPathStr					string
//
//		Input parameter 'pathStr' is subjected to
//		cleaning routines designed to exclude extraneous
//		characters from the analysis. testPathFileStr'
//		is the actual string on which the analysis was
//		performed.
//
//	err							error
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
func (fileHelpMech *fileHelperMechanics) isPathString(
	pathStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	isPathStr bool,
	cannotDetermine bool,
	testPathStr string,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"isPathString()",
		"")

	if err != nil {
		return isPathStr, cannotDetermine, testPathStr, err
	}

	testPathStr = ""
	isPathStr = false
	cannotDetermine = false

	errCode := 0

	errCode,
		_,
		pathStr =
		new(fileHelperElectron).
			isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' is an empty string!\n",
			ePrefix.String())

		return isPathStr, cannotDetermine, testPathStr, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return isPathStr, cannotDetermine, testPathStr, err
	}

	if strings.Contains(pathStr, "...") {

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH STRING!\n"+
			"pathStr='%v'\n",
			ePrefix.String(),
			pathStr)

		return isPathStr, cannotDetermine, testPathStr, err
	}

	testPathStr = new(fileHelperAtom).
		adjustPathSlash(pathStr)

	var pathFileType PathFileTypeCode

	pathFileType,
		_,
		err = new(fileHelperNanobot).
		isPathFileString(
			testPathStr,
			ePrefix)

	if err != nil {

		return isPathStr, cannotDetermine, testPathStr, err
	}

	if pathFileType == PathFileType.Path() {

		isPathStr = true

		cannotDetermine = false

		err = nil

		return isPathStr, cannotDetermine, testPathStr, err
	}

	if pathFileType == PathFileType.Indeterminate() {

		isPathStr = false

		cannotDetermine = true

		err = nil

		return isPathStr, cannotDetermine, testPathStr, err
	}

	isPathStr = false

	cannotDetermine = false

	err = nil

	return isPathStr, cannotDetermine, testPathStr, err
}

// makeDirAll
//
// Creates a directory named path, along with any
// necessary parent directories. In other words, all
// directories in the path are created.
//
// The permission bits 'drwxrwxrwx' are used for all
// directories that the method creates.
//
// If path is a directory which already exists, this
// method does nothing and returns and error value of
// 'nil'.
//
// Note:
//
// This method calls MakeDirAllPerm()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains the name of the directory
//		path which will be created by this method.
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
func (fileHelpMech *fileHelperMechanics) makeDirAll(
	dirPath string,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"makeDirAll()",
		"")

	if err != nil {
		return err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "dirPath"
	}

	var permission FilePermissionConfig

	permission,
		err = new(FilePermissionConfig).
		New("drwxrwxrwx",
			ePrefix)

	if err != nil {
		return fmt.Errorf(
			"Error FilePermissionConfig.New()\n"+
				"Permission Code: \"drwxrwxrwx\"\n"+
				"Error=\n%v\n",
			err.Error())
	}

	err = new(fileHelperNanobot).makeDirAllPerm(
		dirPath,
		permission,
		dirPathLabel,
		ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// readLines
//
// Reads a file and returns each line in a target file as
// an element of a string array.
//
// Multiple custom end of line delimiters may be utilized
// to determine the end of each line of text read from
// the target file. End of line delimiters are specified
// by input parameter 'endOfLineDelimiters', an instance
// of StringArrayDto. 'endOfLineDelimiters' contains an
// array of strings any one of which may be used to
// identify, delimit and separate individual lines of
// text read from the target file.
//
// This method is designed to open a target file, read
// the entire contents of that file, separate the file
// contents into individual lines of text and return
// those text lines in a string array encapsulated by
// an instance of StringArrayDto instance passed as
// input parameter 'outputLinesArray'.
//
// The returned individual lines of text will NOT
// include the end of line delimiters. End of line
// delimiters will therefore be stripped and deleted
// from the end of each configured text line.
//
// It naturally follows that this method will read the
// entire contents of the target file into memory when
// writing said contents to the StringArrayDto instance
// 'outputLinesArray'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the target file ('pathFileName') into
//		memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and writing their contents
//		to the output  instance of StringArrayDto,
//		'outputLinesArray'.
//
//	(2)	This method will open the target file, read the
//		entire contents of that file and automatically
//		close the target file.
//
//		The user is NOT required to manually close the
//		target file.
//
//	(3)	If the target file to be read does not exist on
//		an attached storage drive, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		A string containing the path and file name of the
//		target input file. The contents of this file will
//		be read, line by line, with each text line added
//		as an individual array element in the string
//		array returned by parameter 'strArray'.
//
//		After reading the file contents, the target input
//		file will be automatically closed and rendered
//		ready in all respects for future read/write
//		operations.
//
//	pathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'pathFileName' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileName" will be
//		automatically applied.
//
//	endOfLineDelimiters				*StringArrayDto
//
//		A pointer to an instance of StringArrayDto.
//		'endOfLineDelimiters' encapsulates a string
//		array which contains the end-of-line delimiters
//		that will be used to identify and separate
//		individual lines of text.
//
//		Users have the flexibility to specify multiple
//		end-of-line delimiters for used in parsing text
//		lines extracted from file identified by
//		'pathFileName'.
//
//	writeEndOfLineChars			string
//
//		This string contains the end-of-line characters
//		which will be configured for each line of text
//		written to the output destination specified by
//		the internal io.Writer object.
//
//		On Windows, line-endings are terminated with a
//		combination of a carriage return (ASCII 0x0d or
//		\r) and a newline(\n), also referred to as CR/LF
//		(\r\n).
//
//		On UNIX, text file line-endings are terminated
//		with a newline character (ASCII 0x0a, represented
//		by the \n escape sequence in most languages),
//		also referred to as a linefeed (LF).
//
//		On the Mac Classic (Mac systems using any system
//		prior to Mac OS X), line-endings are terminated
//		with a single carriage return (\r or CR). (Mac OS
//		X uses the UNIX convention.)
//
//		If 'writeEndOfLineChars' is submitted as an empty
//		or zero length string, no end-of-line characters
//		will be written to the io.Writer output
//		destination and no error will be returned.
//
//	outputLinesArray *StringArrayDto,
//
//		A pointer to an instance of StringArrayDto.
//		Lines of text read from the file specified
//		by 'pathFileName' will be stored as
//		individual strings in the string array
//		encapsulated by 'outputLinesArray'.
//
//	maxNumOfLines				int
//
//		Specifies the maximum number of text lines which
//		will be read from the file identified by
//		'pathFileName'.
//
//		If 'maxNumOfLines' is set to a value less than
//		zero (0) (Example: minus-one (-1) ),
//		'maxNumOfLines' will be automatically reset to
//		math.MaxInt(). This means all text lines existing
//		in the file identified by 'pathFileName' will be
//		read and processed. Reading all the text lines in
//		a file 'may' have memory implications depending
//		on the size of the file and the memory resources
//		available to your computer.
//
//		If 'maxNumOfLines' is set to a value of zero
//		('0'), no text lines will be read from the file
//		identified by 'pathFileName', and no error will be
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
//	originalFileSize			int64
//
//		The original file size in bytes of the file
//		specified by input parameter 'pathFileName'.
//
//	numOfLinesRead				int
//
//		This integer value contains the number of text
//		lines read from the file specified by input
//		parameter 'pathFileName'. This value also
//		specifies the number of array elements added to
//		the string array encapsulated by
//		'outputLinesArray'.
//
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the target input file 'pathFileName'
//		and added to the string array encapsulated by
//		'outputLinesArray'.
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
func (fileHelpMech *fileHelperMechanics) readLines(
	pathFileName string,
	pathFileNameLabel string,
	endOfLineDelimiters *StringArrayDto,
	writeEndOfLineChars string,
	outputLinesArray *StringArrayDto,
	maxNumOfLines int,
	errPrefDto *ePref.ErrPrefixDto) (
	originalFileSize int64,
	numOfLinesRead int,
	numOfBytesRead int64,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileHelperMechanics.readLines()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	if len(pathFileNameLabel) == 0 {

		pathFileNameLabel = "pathFileName"
	}

	if endOfLineDelimiters == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endOfLineDelimiters' is invalid!\n"+
			"endOfLineDelimiters' is a nil pointer.\n",
			ePrefix.String())

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	if len(endOfLineDelimiters.StrArray) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endOfLineDelimiters' is invalid!\n"+
			"endOfLineDelimiters' contains a zero length string array.\n"+
			"There are no End-Of-Line delimiters available for text\n"+
			"line separation and identification.\n",
			ePrefix.String())

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err

	}

	if outputLinesArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'outputLinesArray' is invalid!\n"+
			"outputLinesArray' is a nil pointer.\n",
			ePrefix.String())

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	var fInfoPlus FileInfoPlus
	var pathFileDoesExist bool
	var err2 error

	pathFileName,
		pathFileDoesExist,
		fInfoPlus,
		err2 =
		new(fileHelperMolecule).
			doesPathFileExist(
				pathFileName,
				PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
				ePrefix,
				pathFileNameLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while testing for the existance\n"+
			"of '%v' on an attached storage drive.\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName,
			err2.Error())

	}

	if !pathFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"The path and file name do NOT exist on an attached\n"+
			"storage drive. Therefore the contents cannot be read.\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName)

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is directory and NOT a file name.\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName)

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	originalFileSize = fInfoPlus.Size()

	var filePermissionCfg FilePermissionConfig

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		"-r--r--r--",
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fileOpenCfg<-"),
		FOpenType.TypeReadOnly())

	if err != nil {

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	var filePtr *os.File

	defer func() {

		if filePtr != nil {
			_ = filePtr.Close()
		}

	}()

	filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	if err != nil {

		return originalFileSize,
			numOfLinesRead,
			numOfBytesRead,
			err
	}

	numOfLinesRead,
		numOfBytesRead,
		err = new(fileHelperMolecule).
		readerScanLines(
			filePtr,
			pathFileNameLabel,
			endOfLineDelimiters,
			writeEndOfLineChars,
			outputLinesArray,
			maxNumOfLines,
			ePrefix.XCpy("filePtr->"))

	return originalFileSize,
		numOfLinesRead,
		numOfBytesRead,
		err
}
