package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fileHelperDirector struct {
	lock *sync.Mutex
}

// CopyFileByIoByLink
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

	fHelperMech := fileHelperMechanics{}

	err = fHelperMech.copyFileByIo(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	// fh.CopyFileByIo() failed. Try
	// fh.CopyFileByLink()

	var err2 error

	err2 = fHelperMech.
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

	fHelperMech := new(fileHelperMechanics)

	err = fHelperMech.copyFileByLink(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	var err2 error

	// Copy by Link Failed. Try CopyFileByIo()
	err2 = fHelperMech.copyFileByIo(
		src,
		dst,
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
