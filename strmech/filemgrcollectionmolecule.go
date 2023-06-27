package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type FileMgrCollectionMolecule struct {
	lock *sync.Mutex
}

// fmtTextListingAllFiles
//
// Formats and returns a text listing of all file paths
// and file names contained in the File Manager
// Collection passed as input parameter
// 'fileMgrCollection'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgrCollection			*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection.
//		This method will format and return a text listing
//		comprised of all file paths and file names in the
//		File Manager (FileMgr) array encapsulated by
//		this File Manager Collection instance.
//
//	leftMargin					string
//
//		This string serves as the left margin for all
//		text lines.
//
//	rightMargin					string
//
//		This string serves as the right margin for all
//		text lines.
//
//
//	maxLineLength				int
//
//		This integer value defines the maximum line
//		length for all text lines. If this value is
//		less than 10, an error will be returned.
//
//	solidLineChar				rune
//
//		This single character will be used to construct
//		'line-breaks' after the title line. Examples:
//			'-'	"----------------------------"
//			'='	"============================"
//			'*'	"****************************"
//
//
//	titleLine					string
//
//		The text in this string will be formatted as the
//		title for the text listing display.
//
//	addDateTimeLine				bool
//
//		When set to 'true' a text line will be added for
//		current date and time expressed as a local time
//		value.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder.
//		The text listing for directory absolute paths
//		will be added to this instance of
//		strings.Builder.
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
func (fMgrColMolecule *FileMgrCollectionMolecule) fmtTextListingAllFiles(
	fileMgrCol *FileMgrCollection,
	leftMargin string,
	rightMargin string,
	maxLineLength int,
	solidLineChar rune,
	titleLine string,
	addDateTimeLine bool,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fMgrColMolecule.lock == nil {
		fMgrColMolecule.lock = new(sync.Mutex)
	}

	fMgrColMolecule.lock.Lock()

	defer fMgrColMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileMgrCollectionMolecule.fmtTextListingAllFiles()",
		"")

	if err != nil {
		return err
	}

	if fileMgrCol == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileMgrCol' is invalid!\n"+
			"'fileMgrCol' is a nil pointer.\n",
			ePrefix.String())

	}

	if strBuilder == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

	}

	if maxLineLength < 10 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is less than '10'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	if maxLineLength > 1999 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' is greater than '1999'.\n"+
			"'maxLineLength' = %v\n",
			ePrefix.String(),
			maxLineLength)

	}

	if solidLineChar == 0 {

		solidLineChar = '-'

	}

	if len(titleLine) == 0 {

		titleLine = "File Listing"

	}

	lenLeftMar := len(leftMargin)

	lenRightMar := len(rightMargin)

	netTextLineLen := maxLineLength - lenLeftMar - lenRightMar

	if netTextLineLen < 5 {

		return fmt.Errorf("%v\n"+
			"Error: The Net Text Line length is less than 5.\n"+
			"The Net Text Line length is calculated by subtracting\n"+
			"the left and right margin lengths from the maximum\n"+
			"line length.\n"+
			"Net Text Line Length = Maximum Line Length minus\n"+
			"Left Margin Length minus Right Margin Length\n"+
			"Net Text Line Length = '%v'\n"+
			"Maximum Line Length  = '%v'\n"+
			"Left Margin Length   = '%v'\n"+
			"Right Margin Length  = '%v'\n",
			ePrefix.String(),
			maxLineLength,
			netTextLineLen,
			lenLeftMar,
			lenRightMar)

	}

	lenFMgrCol := len(fileMgrCol.fileMgrs)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	thisReqCapacity := lenFMgrCol * 110

	netRequiredCapacity :=
		thisReqCapacity - netCapacityStrBuilder

	if netRequiredCapacity > 0 {

		strBuilder.Grow(netRequiredCapacity + 256)
	}

	txtFormatCol := TextFormatterCollection{}

	var titleMarquee TextLineTitleMarqueeDto

	titleMarquee,
		err = new(TextLineTitleMarqueeDto).
		NewBasicTitleMarqueeDto(
			leftMargin,
			rightMargin,
			leftMargin,
			rightMargin,
			maxLineLength,
			string(solidLineChar),
			ePrefix.XCpy("titleMarquee<-"),
			titleLine+"\n",
			fmt.Sprintf(
				"Number of Files: %v\n",
				lenFMgrCol))

	if err != nil {
		return err
	}

	if addDateTimeLine == true {

		err = titleMarquee.AddTitleLineDateTimeStr(
			time.Now(),
			"Monday 2006-01-02 15:04:05.000000000 -0700 MST",
			ePrefix)

		if err != nil {
			return err
		}
	}

	titleMarquee.NumTrailingBlankLines = 2

	err = txtFormatCol.AddTextTitleMarqueeDto(
		titleMarquee,
		ePrefix.XCpy("<-titleMarquee"))

	if err != nil {
		return err
	}

	err = txtFormatCol.
		SetStdFormatParamsLine1Col(
			leftMargin,
			netTextLineLen,
			TxtJustify.Left(),
			rightMargin,
			false,
			"",
			maxLineLength,
			true,
			ePrefix)

	if err != nil {
		return err
	}

	if lenFMgrCol == 0 {

		extraMargin := "  "
		err = txtFormatCol.AddLine1Col(
			extraMargin+
				"The File Manager Collection is Empty!",
			ePrefix)

		if err != nil {
			return err
		}

		err = txtFormatCol.AddLine1Col(
			extraMargin+
				"There are zero File Manager Objects in the Collection!",
			ePrefix)

		if err != nil {
			return err
		}

	} else {

		for i := 0; i < lenFMgrCol; i++ {

			err = txtFormatCol.AddLine1Col(
				fileMgrCol.fileMgrs[i].
					GetAbsolutePathFileName(),
				ePrefix.XCpy(
					fmt.Sprintf("fileMgrCol.fileMgrs[%v]",
						i)))

			if err != nil {
				return err
			}

		}
	}

	txtFormatCol.AddLineBlank(1, "")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// peekOrPopAtIndex
//
// Returns a deep copy of the File Manager ('FileMgr')
// object located at array index 'idx' in the File
// Manager Collection passed as input parameter 'fMgrs'.
//
// If input parameter 'deleteIndex' is set to 'false',
// this method will function as a 'Peek' method and
// therefore, the original File Manager ('FileMgr')
// object will NOT be deleted from the File Manager
// Collection ('FileMgrCollection') array.
//
// If input parameter 'deleteIndex' is set to 'true',
// this method will function as a 'Pop' method and
// therefore, the original File Manager ('FileMgr')
// object WILL BE DELETED from the File Manager
// Collection ('FileMgrCollection') array. The
// deletion operation will be performed on the File
// Manager object residing at the File Manager
// Collection array index identified by input parameter
// 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgrs						*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection.
//		The File Manager object 'fMgr' specified by
//		the array index 'idx' will be deleted will be
//		deleted from the File Manager Collection.
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be deleted from the File
//		Manager Collection encapsulated by the instance
//		of FileMgrCollection passed by input parameter
//		'fMgrs'.
//
//		If this value is less than zero or greater than
//		the last index in the array, an error will be
//		returned.
//
//	deleteIndex					bool
//
//		If this boolean value is set to 'true', the File
//		Manager object residing at the File Manager
//		Collection index identified by input parameter
//		'idx', will be deleted from File Manager
//		Collection 'fMgrs'.
//
//		If 'deleteIndex' is set to 'false', no deletion
//		occur and the File Manager object residing at
//		File Manager Collection index 'idx' will remain.
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
//	FileMgr
//
//		If this method completes successfully without
//		error, a deep copy of the File Manager object
//		residing at array index 'idx' will be returned
//		through this parameter.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fMgrColMolecule *FileMgrCollectionMolecule) peekOrPopAtIndex(
	fMgrs *FileMgrCollection,
	idx int,
	deleteIndex bool,
	errPrefDto *ePref.ErrPrefixDto) (
	FileMgr,
	error) {

	if fMgrColMolecule.lock == nil {
		fMgrColMolecule.lock = new(sync.Mutex)
	}

	fMgrColMolecule.lock.Lock()

	defer fMgrColMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileMgrCollectionElectron." +
		"peekOrPopAtIndex()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return FileMgr{}, err
	}

	if fMgrs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgrs' is a nil pointer!\n",
			ePrefix.String())

		return FileMgr{}, err
	}

	if fMgrs.fileMgrs == nil {
		fMgrs.fileMgrs = make([]FileMgr, 0, 50)
	}

	arrayLen := len(fMgrs.fileMgrs)

	if arrayLen == 0 {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: The File Manager Collection array encapsulated by 'fMgrs' is EMPTY!\n",
				ePrefix.String())
	}

	if idx < 0 {
		return FileMgr{},
			fmt.Errorf("%v"+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'",
				ePrefix.String(),
				idx)
	}

	if idx >= arrayLen {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				" last array index of the File Manager Collection array.\n"+
				"Index Out-Of-Range!\n"+
				"idx= '%v' "+
				"Last Array Index= '%v' ",
				ePrefix.String(),
				idx,
				arrayLen-1)
	}

	var deepCopyFileMgr FileMgr

	deepCopyFileMgr,
		err = fMgrs.fileMgrs[idx].CopyOut(
		ePrefix.XCpy(fmt.Sprintf(
			"fMgrs.fileMgrs[%v]",
			idx)))

	if err != nil {
		return FileMgr{},
			fmt.Errorf("%v\n"+
				"Error: fMgrs.fileMgrs[%v].CopyOut()\n"+
				"fMgrs.fileMgrs index = '%v'\n"+
				"fMgrs.fileMgrs[%v] = '%v'\n"+
				"Error= \n%v\n",
				funcName,
				idx,
				idx,
				idx,
				fMgrs.fileMgrs[idx].absolutePathFileName,
				err.Error())
	}

	if deleteIndex == false {
		return deepCopyFileMgr, err
	}

	// deleteIndex == true

	if arrayLen == 1 {

		fMgrs.fileMgrs = make([]FileMgr, 0, 1)

	} else if idx == 0 {
		// arrayLen > 1 and requested idx = 0
		fMgrs.fileMgrs = fMgrs.fileMgrs[1:]

	} else if idx == arrayLen-1 {
		// arrayLen > 1 and requested idx = last element index
		fMgrs.fileMgrs = fMgrs.fileMgrs[0 : arrayLen-1]

	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		fMgrs.fileMgrs =
			append(fMgrs.fileMgrs[0:idx], fMgrs.fileMgrs[idx+1:]...)
	}

	return deepCopyFileMgr, err
}
