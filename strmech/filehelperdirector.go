package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math"
	"os"
	"strings"
	"sync"
)

type fileHelperDirector struct {
	lock *sync.Mutex
}

// copyFileByIoByLink
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a new destination file and
// using "io.Copy(out, in)" to copy the contents. This is
// accomplished by calling 'CopyFileByIo()'.
// If  the call to 'CopyFileByIo()' fails, this method
// will attempt a second copy method.
//
// The second attempt to copy the designated file will be
// accomplished by creating a 'hard link' to the source
// file. The second, 'hard link', attempt will call
// method, 'CopyFileByLink()'.
//
// If both attempted file copy operations fail, an error
// will be returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
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
func (fHelpDirector *fileHelperDirector) copyFileByIoByLink(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"copyFileByIoByLink()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(fileHelperUtility).copyFileByIo(
		src,
		"src",
		dst,
		"dst",
		true, // createDestDirPathIfNotExist
		ePrefix)

	if err == nil {
		return err
	}

	// fh.CopyFileByIo() failed. Try
	// fh.CopyFileByLink()

	var err2 error

	err2 = new(fileHelperMechanics).
		copyFileByLink(
			src,
			dst,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By IO failed, an error was returned\n"+
			"by CopyFileByLink(src, dst)\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			src,
			dst,
			err2.Error())

	}

	return err
}

// copyFileByLinkByIo
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a 'hard link' to the source file.
// The 'hard link' attempt will call 'FileHelper.CopyFileByLink()'.
//
// If that 'hard link' operation fails, this method will call
// 'CopyFileByIo()'.
//
// CopyFileByIo() will create a new destination file and attempt
// to write the contents of the source file to the new destination
// file using "io.Copy(out, in)".
//
// If both attempted file copy operations fail, an error will be
// returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
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
func (fHelpDirector *fileHelperDirector) copyFileByLinkByIo(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"copyFileByLinkByIo()",
		"")

	if err != nil {
		return err
	}

	err = new(fileHelperMechanics).copyFileByLink(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	var err2 error

	// Copy by Link Failed. Try CopyFileByIo()
	_,
		err2 = new(fileHelperUtility).copyFileByIo(
		src,
		"src",
		dst,
		"dst",
		true, // createDestDirPathIfNotExist
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link failed, an error was returned by fh.CopyFileByIo(src, dst).\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix,
			src,
			dst,
			err2.Error())

		return err
	}

	return err
}

// compareFiles
//
// Receives two strings containing the path and file
// names of two files.
//
// The two files will be compared to determine if their
// contents are identical.
//
// If the compared files are equal with respect to
// content, this method will return a boolean value of
// 'true'.
//
// If the two files differ in file size or file content,
// this method will return 'false'.
//
// If no errors are encountered and the contents of the
// two files are found to be 'NOT EQUAL', this method
// will return a text description of the reason for this
// inequality.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameOne				string
//
//		This string holds the path and file name of
//		File-1. File-1 will be compared with File-2 to
//		determine if the two files are equal in terms
//		of content.
//
//	pathFileNameOneLabel		string
//
//		The name or label associated with input parameter
//		'pathFileNameOne' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileNameOne" will
//		be automatically applied.
//
//	pathFileNameTwo				string
//
//		This string holds the path and file name of
//		File-2. File-2 will be compared with File-1 to
//		determine if the two files are equal in terms
//		of content.
//
//	pathFileNameTwoLabel		string
//
//		The name or label associated with input parameter
//		'pathFileNameTwo' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileNameTwo" will
//		be automatically applied.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	filesAreEqual				bool
//
//		If this return parameter is set to 'true', it
//		signals that the contents of File-1 ('pathFileNameOne')
//		and File-2 ('pathFileNameTwo') are equal.
//
//		If this return parameter is set to 'false', it
//		signals that the contents of File-1 ('pathFileNameOne')
//		and File-2 ('pathFileNameTwo') are NOT equal.
//
//	reasonFilesNotEqual			string
//
//		If the contents of File-1 ('pathFileNameOne') and
//		File-2 ('pathFileNameTwo') are determined to be
//		NOT EQUAL, this returned string will contain text
//		describing the reason for this inequality.
//
//		If File-1 and File-2 are equal in terms of content,
//		this string will be empty.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelpDirector *fileHelperDirector) compareFiles(
	pathFileNameOne string,
	pathFileNameOneLabel string,
	pathFileNameTwo string,
	pathFileNameTwoLabel string,
	errorPrefix interface{}) (
	filesAreEqual bool,
	reasonFilesNotEqual string,
	err error) {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileHelperDirector." +
		"compareFiles()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return filesAreEqual, reasonFilesNotEqual, err
	}

	if len(pathFileNameOneLabel) == 0 {
		pathFileNameOneLabel = "pathFileNameOne"
	}

	if len(pathFileNameTwoLabel) == 0 {
		pathFileNameTwoLabel = "pathFileNameTwo"
	}

	if len(pathFileNameOne) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' is an empty string with a string length of zero.\n",
			ePrefix.String(),
			pathFileNameOneLabel,
			pathFileNameOneLabel)

		return filesAreEqual, reasonFilesNotEqual, err
	}

	if len(pathFileNameTwo) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' is an empty string with a string length of zero.\n",
			ePrefix.String(),
			pathFileNameTwoLabel,
			pathFileNameTwoLabel)

		return filesAreEqual, reasonFilesNotEqual, err
	}

	var fileInfoPlusOne FileInfoPlus
	var err2 error

	var fileBufReaderOne FileBufferReader

	fileInfoPlusOne,
		err2 = new(fileBufferReaderNanobot).
		setPathFileName(
			&fileBufReaderOne,
			"fileBufReaderOne",
			pathFileNameOne,
			pathFileNameOneLabel,
			false, // openFileReadWrite
			4096,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while converting input\n"+
			"parameter '%v' to a FileBufferReader.\n"+
			"'%v' = %v\n"+
			"Error=\n%v\n",
			funcName,
			pathFileNameOneLabel,
			pathFileNameOneLabel,
			pathFileNameOne,
			err2.Error())

		return filesAreEqual, reasonFilesNotEqual, err
	}

	var fileBufReaderTwo FileBufferReader
	var fileInfoPlusTwo FileInfoPlus
	var fileOneBytesRead = make([]byte, 4096)
	var fileTwoBytesRead = make([]byte, 4096)
	var fileOneNumBytesRead, fileTwoNumBytesRead int
	var readCycle, maxCycle int64
	maxCycle = int64(math.MaxInt)
	var errReadOne, errReadTwo error

	var errs []error

	fileInfoPlusTwo,
		err2 = new(fileBufferReaderNanobot).
		setPathFileName(
			&fileBufReaderTwo,
			"fileBufReaderTwo",
			pathFileNameTwo,
			pathFileNameTwoLabel,
			false, // openFileReadWrite
			4096,
			ePrefix)

	if err2 != nil {

		errs = append(
			errs,
			fmt.Errorf("%v\n"+
				"An error occurred while converting input\n"+
				"parameter '%v' to a FileBufferReader.\n"+
				"'%v' = %v\n"+
				"Error=\n%v\n",
				funcName,
				pathFileNameTwoLabel,
				pathFileNameTwoLabel,
				pathFileNameTwo,
				err2.Error()))

		goto finalExit
	}

	if fileInfoPlusOne.Size() != fileInfoPlusTwo.Size() {

		reasonFilesNotEqual = "The files sizes (in bytes) are NOT equal."

		goto finalExit
	}

	for {

		readCycle++

		if readCycle > maxCycle {

			errs = append(
				errs,
				fmt.Errorf("%v\n"+
					"Error: The number of 'read' cycles exceeded the maximum!\n"+
					"The maximum number of 'read' cycles is %v\n",
					ePrefix.String(),
					maxCycle))

			goto finalExit
		}

		fileOneNumBytesRead,
			errReadOne = fileBufReaderOne.
			Read(
				fileOneBytesRead)

		if errReadOne != nil &&
			errReadOne != io.EOF {

			errs =
				append(
					errs,
					fmt.Errorf("%v\n"+
						"Error reading %v\n"+
						"'%v'= %v\n"+
						"Read Error=\n%v\n",
						ePrefix.String(),
						pathFileNameOneLabel,
						pathFileNameOneLabel,
						pathFileNameOne,
						errReadOne.Error()))

			goto finalExit
		}

		fileTwoNumBytesRead,
			errReadTwo = fileBufReaderTwo.
			Read(
				fileTwoBytesRead)

		if errReadTwo != nil &&
			errReadTwo != io.EOF {

			errs = append(
				errs,
				fmt.Errorf("%v\n"+
					"Error reading %v\n"+
					"'%v'= %v\n"+
					"Read Error=\n%v\n",
					ePrefix.String(),
					pathFileNameTwoLabel,
					pathFileNameTwoLabel,
					pathFileNameTwo,
					err2.Error()))

			goto finalExit
		}

		if fileOneNumBytesRead != fileTwoNumBytesRead {

			reasonFilesNotEqual = "Number of bytes read from Files 1 & 2 are not equal"

			goto finalExit
		}

		for i := 0; i < fileTwoNumBytesRead; i++ {

			if fileOneBytesRead[i] != fileTwoBytesRead[i] {

				reasonFilesNotEqual = "Files 1 & 2 content are not equal"

				goto finalExit
			}
		}

		if errReadOne == io.EOF &&
			errReadTwo != io.EOF {

			reasonFilesNotEqual = "Files 1 & 2 Read exits (io.EOF) are not equal"

			goto finalExit
		}

		if errReadOne != io.EOF &&
			errReadTwo == io.EOF {

			reasonFilesNotEqual = "Files 1 & 2 Read exits (io.EOF) are not equal"

			goto finalExit
		}

		if errReadOne == io.EOF &&
			errReadTwo == io.EOF {

			break
		}

	}

	filesAreEqual = true

finalExit:

	err2 = fileBufReaderOne.Close()

	if err2 != nil {

		errs = append(
			errs,
			fmt.Errorf("%v\n"+
				"--------------------------------\n"+
				"Error Closing 'fileBufReaderOne'\n"+
				"Error=\n%v\n",
				ePrefix,
				err2.Error()))

	}

	err2 = fileBufReaderTwo.Close()

	if err2 != nil {

		errs = append(
			errs,
			fmt.Errorf("%v\n"+
				"Error Closing 'fileBufReaderTwo'\n"+
				"Error=\n%v\n",
				ePrefix,
				err2.Error()))

	}

	if len(errs) > 0 {

		err = new(StrMech).ConsolidateErrors(errs)

	}

	return filesAreEqual, reasonFilesNotEqual, err
}

// getPathAndFileNameExt
//
// Breaks out path and fileName+Ext elements from a path
// string. If both path and fileName are empty strings,
// this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string holds the file path, file name and
//		file extension. The file path will be returned
//		as in the first parameter. The file name and file
//		extension will be returned in a second paramter.
//
//	pathFileNameExtLabel		string
//
//		The name or label associated with input parameter
//		'pathFileNameExt' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileNameExt" will
//		be automatically applied.
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
//	pathDir						string
//
//		This returned string will contain the directory
//		path extracted from input parameter 'pathFileNameExt'.
//
//	fileNameExt					string
//
//		This returned string will contain the file name
//		and file extension extracted from input parameter
//		'pathFileNameExt'.
//
//	bothAreEmpty				bool
//
//		If both 'pathDir' and 'fileNameExt' are returned
//		as empty strings, this return parameter will be
//		set to 'true'.
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
func (fHelpDirector *fileHelperDirector) getPathAndFileNameExt(
	pathFileNameExt string,
	pathFileNameExtLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	pathDir string,
	fileNameExt string,
	bothAreEmpty bool,
	err error) {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathDir = ""

	fileNameExt = ""

	bothAreEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"getPathAndFileNameExt()",
		"")

	if err != nil {
		return pathDir, fileNameExt, bothAreEmpty, err
	}

	if len(pathFileNameExtLabel) == 0 {

		pathFileNameExtLabel = "pathFileNameExt"
	}

	trimmedFileNameExt := ""

	errCode := 0

	errCode,
		_,
		trimmedFileNameExt = new(fileHelperElectron).
		isStringEmptyOrBlank(pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is an empty string!\n",
			ePrefix.String(),
			pathFileNameExtLabel)

		return pathDir, fileNameExt, bothAreEmpty, err
	}

	if errCode == -2 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' consists of blank spaces!\n",
			ePrefix.String(),
			pathFileNameExtLabel)

		return pathDir, fileNameExt, bothAreEmpty, err
	}

	xFameExt,
		isEmpty,
		err2 := new(fileHelperNanobot).
		getFileNameWithExt(
			trimmedFileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from getFileNameWithExt(%v).\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			pathFileNameExtLabel,
			pathFileNameExtLabel,
			pathFileNameExt,
			err2.Error())

		return pathDir, fileNameExt, bothAreEmpty, err
	}

	if isEmpty {
		fileNameExt = ""
	} else {
		fileNameExt = xFameExt
	}

	remainingPathStr := strings.TrimSuffix(
		trimmedFileNameExt, fileNameExt)

	if len(remainingPathStr) == 0 {
		pathDir = ""

		if pathDir == "" && fileNameExt == "" {
			bothAreEmpty = true
		} else {
			bothAreEmpty = false
		}

		return pathDir, fileNameExt, bothAreEmpty, err
	}

	xPath,
		isEmpty,
		err2 :=
		new(fileHelperMechanics).
			getPathFromPathFileName(
				remainingPathStr,
				ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from getPathFromPathFileName(remainingPathStr).\n"+
			"remainingPathStr= '%v'\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			pathFileNameExtLabel,
			pathFileNameExt,
			remainingPathStr,
			err2.Error())

		return pathDir, fileNameExt, bothAreEmpty, err
	}

	if isEmpty {
		pathDir = ""
	} else {
		pathDir = xPath
	}

	if pathDir == "" && fileNameExt == "" {
		bothAreEmpty = true
	} else {
		bothAreEmpty = false
	}

	return pathDir, fileNameExt, bothAreEmpty, err
}

// moveFile
//
// Copies a file from source to destination and, if
// successful, then deletes the original source file.
//
// The copy procedure will be carried out using the
// 'Copy By Io' technique. See FileHelper.CopyFileByIo().
//
// The 'move' operation will create the destination file,
// but it will NOT create the destination directory. If
// the destination directory does NOT exist, an error will
// be returned.
//
// If this copy operation fails, the method will return an
// error, and the source file will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the source file after the
//	copy from source to destination file is completed.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the path and file name of the
//		source file. This source file will be copied to
//		the destination path and file name specified by
//		input parameter 'dst'.
//
//		After the source file ('src') is copied to the
//		destination file ('dst'), the source file ('src')
//		will be deleted.
//
//	dst							string
//
//		This string holds the path and file name of the
//		destination file. The source file ('src') will be
//		copied to this destination file ('dst') and the
//		source file will be deleted.
//
//		If the directory path for the destination file
//		('dst') does not previously exist, an error will
//		be returned, the source file ('src') will not be
//		copied to the destination file ('dst') and the
//		source file will NOT be deleted.
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
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelpDirector *fileHelperDirector) moveFile(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperDirector."+
			"moveFile()",
		"")

	if err != nil {
		return err
	}

	// ============================
	// Perform the copy operation!
	// Use Copy By IO Procedure
	// ============================

	var err2 error

	_,
		err2 = new(fileHelperUtility).copyFileByIo(
		src,
		"src",
		dst,
		"dst",
		true, // createDestDirPathIfNotExist
		ePrefix)

	if err2 != nil {
		// Copy Operation Failed. Return an error
		// and DO NOT delete the source file!
		err = fmt.Errorf("%v\n"+
			"Error: Copy operation FAILED!\n"+
			"Source File='%v'\n"+
			"Destination File='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			src,
			dst,
			err2.Error())

		return err
	}

	// CopyFileByIo operation was apparently successful.
	// Now, verify that destination file exists.

	var dstFileDoesExist bool
	var fHelpMolecule = new(fileHelperMolecule)

	_,
		dstFileDoesExist,
		_,
		err = fHelpMolecule.
		doesPathFileExist(
			dst,
			PreProcPathCode.None(), // Take no Pre-Processing action
			ePrefix,
			"dst")

	if err != nil {
		return err
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy Operation, destination file "+
			"DOES NOT EXIST!\n"+
			"Therefore, the copy operation FAILED! Source file was NOT deleted.\n"+
			"destination file='%v'\n",
			ePrefix.String(),
			dst)

		return err
	}

	// Successful copy operation has been verified.
	// Time to delete the source file.
	err2 = os.Remove(src)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Copy operation succeeded, but attempted deletion of source file FAILED!\n"+
			"Source File='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			src,
			err2.Error())

		return err
	}

	var srcFileDoesExist bool

	_,
		srcFileDoesExist,
		_,
		err = fHelpMolecule.doesPathFileExist(
		src,
		PreProcPathCode.None(), // Take No Pre-Processing Action
		ePrefix,
		"src")

	if err != nil {
		return err
	}

	if srcFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Verification Error: File 'src' still exists!\n"+
			"src='%v'\n",
			ePrefix.String(),
			src)

		return err
	}

	// Success, source was copied to destination
	// AND the source file was deleted.

	// Done and we are out of here!
	return err
}

// openDirectory
//
// Opens a directory and returns the associated 'os.File'
// pointer. This method will open a directory designated
// by input parameter, 'directoryPath'.
//
// The input parameter 'createDir' determines the action
// taken if 'directoryPath' does not exist. If
// 'createDir' is set to 'true' and 'directoryPath' does
// not currently exist, this method will attempt to
// create 'directoryPath'. Directories created in this
// manner are configured with Open Type of 'Read-Write'
// and a Permission code of 'drwxrwxrwx'.
//
// Alternatively, if 'createDir' is set to 'false' and
// 'directoryPath' does NOT exist, an error will be
// returned.
//
// Regardless of whether the target directory path
// already exists or is created by this method, the
// returned os.File pointer is opened with the
// 'Read-Only' attribute (O_RDONLY) and a
// permission code of zero ("----------").
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The caller is responsible for calling "Close()" on the
// returned os.File pointer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		A string containing the path name of the
//		directory which will be opened.
//
//	createDir					bool
//
//		Determines what action will be taken if
//		'directoryPath' does NOT exist. If 'createDir' is
//		set to 'true' and 'directoryPath' does NOT exist,
//		this method will attempt to create
//		'directoryPath'. Alternatively, if 'createDir' is
//		set to false and 'directoryPath' does NOT exist,
//		this method will terminate and an error will be
//		returned.
//
//		Directories created in this manner will have an
//		Open Type of 'Read-Write' and a Permission code
//		of 'drwxrwxrwx'. This differs from the Open Type
//		and permission mode represented by the returned
//		os.File pointer.
//
//	directoryPathLabel			string
//
//		The name or label associated with input parameter
//		'directoryPath' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "directoryPath" will
//		be automatically applied.
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
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the directory designated by input
//		parameter 'directoryPath'.
//
//		If successful, the returned os.File pointer is
//		opened with the 'Read-Only' attribute (O_RDONLY)
//		and a permission code of zero ("----------").
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//			The caller is responsible for calling
//			"Close()" on this os.File pointer.
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
func (fHelpDirector *fileHelperDirector) openDirectory(
	directoryPath string,
	createDir bool,
	directoryPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	*os.File,
	error) {

	if fHelpDirector.lock == nil {
		fHelpDirector.lock = new(sync.Mutex)
	}

	fHelpDirector.lock.Lock()

	defer fHelpDirector.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "fileHelperDirector." +
		"openDirectory()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return nil, err
	}

	if len(directoryPathLabel) == 0 {
		directoryPathLabel = "directoryPath"
	}

	var directoryPathDoesExist bool
	var dirPathFInfo FileInfoPlus

	fHelpMolecule := fileHelperMolecule{}

	directoryPath,
		directoryPathDoesExist,
		dirPathFInfo,
		err = fHelpMolecule.doesPathFileExist(
		directoryPath,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		directoryPathLabel)

	if err != nil {
		return nil, err
	}

	if directoryPathDoesExist &&
		!dirPathFInfo.IsDir() {

		return nil,
			fmt.Errorf("%v\n"+
				"ERROR: 'directoryPath' does exist, but\n"+
				"IT IS NOT A DIRECTORY!\n"+
				"%v= '%v'\n",
				ePrefix.String(),
				directoryPathLabel,
				directoryPath)
	}

	if directoryPathDoesExist && dirPathFInfo.Mode().IsRegular() {
		return nil,
			fmt.Errorf("%v\n"+
				"ERROR: 'directoryPath' does exist, but\n"+
				"it is classifed as a REGULAR File!\n"+
				"directoryPath='%v'\n",
				ePrefix.String(),
				directoryPath)
	}

	if !directoryPathDoesExist {

		if !createDir {
			return nil,
				fmt.Errorf("%v\n"+
					"Error '%v' DOES NOT EXIST!\n"+
					"%v= '%v'\n",
					ePrefix.String(),
					directoryPathLabel,
					directoryPathLabel,
					directoryPath)
		}

		// Parameter 'createDir' must be 'true'.
		// The error signaled that the path does not exist. So, create the directory path
		err = new(fileHelperPreon).makeDirAll(
			directoryPath,
			directoryPathLabel,
			ePrefix)

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Attmpted creation of '%v' FAILED!\n"+
					"%v= '%v'\n"+
					"Error=\n%v\n",
					funcName,
					directoryPathLabel,
					directoryPathLabel,
					directoryPath,
					err.Error())
		}

		// Verify that the directory exists and get
		// the associated file info object.
		_,
			directoryPathDoesExist,
			dirPathFInfo,
			err = fHelpMolecule.doesPathFileExist(
			directoryPath,
			PreProcPathCode.None(), // Take No Pre-Processing Action
			ePrefix,
			directoryPathLabel)

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"Error occurred verifying existance of "+
					"newly created '%v'!\n"+
					"Non-Path error returned by os.Stat(%v)\n"+
					"%v= '%v'\n"+
					"Error=\n%v\n",
					funcName,
					directoryPathLabel,
					directoryPathLabel,
					directoryPathLabel,
					directoryPath,
					err.Error())
		}

		if !directoryPathDoesExist {
			return nil, fmt.Errorf("%v\n"+
				"Error: Verification of newly created "+
				"%v FAILED!\n"+
				"'%v' DOES NOT EXIST!\n"+
				"%v= '%v'\n",
				ePrefix.String(),
				directoryPathLabel,
				directoryPathLabel,
				directoryPathLabel,
				directoryPath)
		}

		if !dirPathFInfo.IsDir() {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Input Paramter '%v' is NOT a directory!\n"+
					"%v= '%v'\n",
					ePrefix.String(),
					directoryPathLabel,
					directoryPathLabel,
					directoryPath)
		}

		if dirPathFInfo.Mode().IsRegular() {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: '%v' does exist, but\n"+
					"it is classifed as a REGULAR File!\n"+
					"%v= '%v'\n",
					ePrefix.String(),
					directoryPathLabel,
					directoryPathLabel,
					directoryPath)
		}
	}

	filePtr, err := os.Open(directoryPath)

	if err != nil {
		return nil,
			fmt.Errorf("%v\n"+
				"Error returned by os.Open(%v)\n"+
				"%v= '%v'\n"+
				"File Open Error: \n%v\n",
				ePrefix.String(),
				directoryPathLabel,
				directoryPathLabel,
				directoryPath,
				err.Error())
	}

	if filePtr == nil {
		err = fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile(%v) returned a\n"+
			"'nil' file pointer!\n",
			ePrefix.String(),
			directoryPathLabel)
	}

	return filePtr, err
}
