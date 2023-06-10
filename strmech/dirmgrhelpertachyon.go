package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperTachyon struct {
	lock *sync.Mutex
}

// getFileInfosFromDirectory
//
// Receives an instance of DirMgr and proceeds to extract
// os.FileInfo data DirMgr's absolute path.
//
// Upon completion an array of FileInfoPlus objects is
// returned containing os.FileInfo information on all files
// and directories in the DirMgr's absolute path.
//
// Type FileInfoPlus implements the os.FileInfo interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will be
//		searched to extract and return os.FileInfo
//		information on all files and directories contained
//		therein.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
//	fileInfos					[]FileInfoPlus
//
//		If this method completes successfully, this
//		method will return an array of FileInfoPlus
//		objects containing os.FileInfo data on all files
//		and directories contained in the directory path
//		specified by input parameter 'dMgr'.
//
//		Type FileInfoPlus implements the os.FileInfo
//		interface, but provides additional file information
//		over and above that provided by the standard
//		os.FileInfo interface.
//
//		The os.FileInfo interface is defined as follows:
//
//	 	type FileInfo interface {
//			 Name() string       // base name of the file
//			 Size() int64        // length in bytes for regular files; system-dependent for others
//			 Mode() FileMode     // file mode bits
//			 ModTime() time.Time // modification time
//			 IsDir() bool        // abbreviation for Mode().IsDir()
//			 Sys() interface{}   // underlying data source (can return nil)
//	 	}
//
//	nonfatalErrs				[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If non-fatal errors are encountered during
//		processing, the returned error Type will
//		encapsulate appropriate error messages.
//
//		Non-fatal errors usually involve failures
//		associated with reading individual files.
//
//		The returned error messages will incorporate
//		the method chain and text passed by input
//		parameter, 'errPrefDto'. The 'errPrefDto' text
//		will be prefixed or attached to the beginning of
//		the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
//
//	fatalErr					error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'.
//
//		If a fatal error is encountered during
//		processing, this returned error Type will
//		encapsulate an appropriate error message. This
//		returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errPrefDto'. The 'errPrefDto' text will be
//		prefixed or attached to the	beginning of the error
//		message.
func (dMgrHlprTachyon *dirMgrHelperTachyon) getFileInfosFromDirectory(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	fileInfos []FileInfoPlus,
	lenFileInfos int,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprTachyon.lock == nil {
		dMgrHlprTachyon.lock = new(sync.Mutex)
	}

	dMgrHlprTachyon.lock.Lock()

	defer dMgrHlprTachyon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperElectron.getFileInfosFromDirectory()",
		"")

	if err != nil {

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "sourceDMgr"
	}

	var dMgrHlprPreon = new(dirMgrHelperPreon)

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if fatalErr != nil {

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(%v.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	var nameFileInfo os.FileInfo
	osPathSepStr := string(os.PathSeparator)
	fip := FileInfoPlus{}

	lenFileInfos = len(nameDirEntries)

	fileInfos = make([]FileInfoPlus, lenFileInfos)

	var osFileInfo os.FileInfo

	for i := 0; i < lenFileInfos; i++ {

		osFileInfo,
			err2 = nameDirEntries[i].Info()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Error Returned by nameDirEntry.Info().\n"+
				"The conversion of DirEntry to os.FileInfo Failed."+
				"%v= '%v'\n"+
				"FileName= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath,
				dMgr.absolutePath+osPathSepStr+nameFileInfo.Name(),
				err2.Error())

			nonfatalErrs = append(nonfatalErrs, err)

			continue
		}

		fileInfos = append(fileInfos, fip.NewFromFileInfo(osFileInfo))
	}

	lenFileInfos = len(fileInfos)

	return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
}

// getSubdirectories
//
// This method receives an instance of DirMgr ('dMgr')
// and proceeds to identify all the subdirectories located
// within the directory path specified by this DirMgr
// instance.
//
// Any subdirectories located in the 'dMgr' path will be
// returned as a Directory Manager Collection.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will be
//		searched to extract and return os.FileInfo
//		information on all files and directories contained
//		therein.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
//	subDirectories				DirMgrCollection
//
//		If this method completes successfully, this
//		method will return an instance of DirMgrCollection
//		populated with an array of 'DirMgr' objects
//		identifying the subdirectories contained in the
//		directory path specified by input parameter
//		'dMgr'.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//	dTreeStats					DirectoryStatsDto
//
//		The DirectoryStatsDto structure is used to
//		accumulate and disseminate statistical
//		information relating to a specific directory
//		tree.
//
//		type DirectoryStatsDto struct {
//
//			dMgr DirMgr
//				Identifies the parent directory associated with
//				this directory information.
//
//			numOfFiles uint64
//				The number of files (all types) residing
//				within this directory ('dMgr').
//
//			numOfSubDirs uint64
//				The number of subdirectories residing
//				within this directory
//
//			numOfBytes uint64
//				The total number of bytes for all files
//				contained in this directory.
//				isInitialized bool
//		}
//
//		Type DirectoryStatsDto contains public methods
//		for retrieving the specified directory statistics
//		and information.
//
//		If this method completes successfully, this
//		returned instance of DirectoryStatsDto will
//		contain information on files and directories
//		contained in the directory tree specified by
//		input parameter 'targetBaseDir'.
//
//	fatalErr					error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'.
//
//		If a fatal error is encountered during
//		processing, this returned error Type will
//		encapsulate an appropriate error message. This
//		returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errPrefDto'. The 'errPrefDto' text will be
//		prefixed or attached to the	beginning of the error
//		message.
func (dMgrHlprTachyon *dirMgrHelperTachyon) getSubdirectories(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	subdirectories DirMgrCollection,
	dTreeStats DirectoryStatsDto,
	fatalErr error) {

	if dMgrHlprTachyon.lock == nil {
		dMgrHlprTachyon.lock = new(sync.Mutex)
	}

	dMgrHlprTachyon.lock.Lock()

	defer dMgrHlprTachyon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperAtom." +
		"getSubdirectories()"

	subdirectories.dirMgrs = make([]DirMgr, 0)

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return subdirectories, dTreeStats, fatalErr
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		fatalErr = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if fatalErr != nil {

		return subdirectories, dTreeStats, fatalErr
	}

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(%v.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return subdirectories, dTreeStats, fatalErr
	}

	dTreeStats,
		err2 = new(DirectoryStatsDto).
		New(
			*dMgr,
			ePrefix.XCpy("dMgr"))

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Failed to create a new instance of DirectoryStatsDto!\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return subdirectories, dTreeStats, fatalErr
	}

	if len(nameDirEntries) == 0 {

		return subdirectories, dTreeStats, fatalErr
	}

	var fInfo os.FileInfo

	for _, dirEntry := range nameDirEntries {

		fInfo,
			err2 = dirEntry.Info()

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Conversion of Direct Entry to os.FileInfo Failed!\n"+
				"Error returned by dirEntry.Info().\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath,
				err2.Error())

			return subdirectories, dTreeStats, fatalErr
		}

		if fInfo.IsDir() {

			dTreeStats.numOfSubDirs++
			dTreeStats.numOfBytes += uint64(fInfo.Size())

			err2 = subdirectories.
				AddDirMgrByKnownPathDirName(
					dMgr.absolutePath,
					fInfo.Name(),
					ePrefix.XCpy("dMgr"))

		} else {
			// This must be a file

			dTreeStats.numOfFiles++
			dTreeStats.numOfBytes += uint64(fInfo.Size())

		}

	}

	return subdirectories, dTreeStats, fatalErr
}
