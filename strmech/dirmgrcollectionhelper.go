package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type dirMgrCollectionHelper struct {
	lock *sync.Mutex
}

// copyCollection
//
// This method copies all the Director Manager objects
// contained in a 'source' Directory Manager Collection
// to a 'destination' Directory Manager Collection.
//
// All Directory Manager objects will be copied from
// input parameter 'sourceDMgrCollection' to input
// parameter 'destinationDMgrCollection'. Both instances
// are of type DirMgrCollection. The type copy operation
// employed is a 'deep' copy operation.
//
// If input parameter 'sourceDMgrCollection' contains
// an empty Directory Manager collection, an error will
// be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data fields contained in input parameter
//	'destinationDMgrCollection' will be deleted an
//	overwritten with new data values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationDMgrCollection	*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection. All
//		data fields contained in input parameter
//		'sourceDMgrCollection' will be copied to
//		corresponding data fields in
//		'destinationDMgrCollection'.
//
//		All original pre-existing data values contained
//		in 'destinationDMgrCollection' will be deleted
//		and overwritten.
//
//	sourceDMgrCollection *DirMgrCollection,
//
//		A pointer to an instance of DirMgrCollection. All
//		data fields contained in input parameter
//		'sourceDMgrCollection' will be copied to
//		corresponding data fields in
//		'destinationDMgrCollection'.
//
//		If the Directory Manager Collection contained
//		in 'sourceDMgrCollection' is empty.
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
func (dMgrColHelper *dirMgrCollectionHelper) copyCollection(
	destinationDMgrCollection *DirMgrCollection,
	sourceDMgrCollection *DirMgrCollection,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrCollectionHelper.copyCollection()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if sourceDMgrCollection == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceDMgrCollection' is a nil pointer!\n",
			ePrefix.String())
	}

	if destinationDMgrCollection == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationDMgrCollection' is a nil pointer!\n",
			ePrefix.String())
	}

	lenSourceDMgrs := len(sourceDMgrCollection.dirMgrs)

	if lenSourceDMgrs == 0 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceDMgrCollection' is invalid!\n"+
			"The 'sourceDMgrCollection' collection of Directory Managers is Empty.\n",
			ePrefix.String())
	}

	destinationDMgrCollection.dirMgrs = make([]DirMgr, lenSourceDMgrs)

	for i := 0; i < lenSourceDMgrs; i++ {

		err = destinationDMgrCollection.dirMgrs[i].
			CopyIn(
				&sourceDMgrCollection.dirMgrs[i],
				ePrefix.XCpy(
					"sourceDMgrCollection.dirMgrs[i]"))

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error: destinationDMgrCollection.dirMgrs[%v].CopyIn()\n"+
				"sourceDMgrCollection.dirMgrs index= '%v'\n"+
				"sourceDMgrCollection DirMgr= '%v'\n"+
				"Error= \n%v\n",
				funcName,
				i,
				i,
				sourceDMgrCollection.dirMgrs[i].absolutePath,
				err.Error())
		}

	}

	return err
}

// equalDMgrCollections
//
// This method receives pointers to two instances of
// DirMgrCollection and proceeds to analyze all members
// of each Directory Manager Collection to determine if
// the collections are equal in all respects.
//
// If any of the Directory Manager (DirMgr) objects in
// the two collections are not equal, this method returns
// a boolean value of 'false'.
//
// A value of 'true' is only returned if all Directory
// Manager objects in both collections are equal in all
// respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgrCollectionOne			*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection. All
//		the Directory Manager objects in this Directory
//		Manager Collection will be compared to the
//		Directory Manager Collection contained in input
//		parameter 'dMgrCollectionTwo' to determine if
//		all the Directory Manager objects are equivalent.
//
//	dMgrCollectionTwo			*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection. All
//		the Directory Manager objects in this Directory
//		Manager Collection will be compared to the
//		Directory Manager Collection contained in input
//		parameter 'dMgrCollectionOne' to determine if
//		all the Directory Manager objects are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		Two Directory Manager Collections from input
//		parameters 'dMgrCollectionOne' and
//		'dMgrCollectionTwo' are compared to determine
//		if they are equal in all respects.
//
//		If any of the Directory Manager (DirMgr) objects
//		in the two collections are not equal, this method
//		returns a boolean value of 'false'.
//
//		A value of 'true' is only returned if all
//	 	Directory Manager objects in both collections are
//	 	equal in all respects.
func (dMgrColHelper *dirMgrCollectionHelper) equalDMgrCollections(
	dMgrCollectionOne *DirMgrCollection,
	dMgrCollectionTwo *DirMgrCollection) bool {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	if dMgrCollectionOne == nil ||
		dMgrCollectionTwo == nil {

		return false
	}

	collectionLen := len(dMgrCollectionOne.dirMgrs)

	if collectionLen != len(dMgrCollectionTwo.dirMgrs) {

		return false
	}

	// Collection array lengths are equal

	for i := 0; i < collectionLen; i++ {

		if !dMgrCollectionOne.dirMgrs[i].Equal(
			&dMgrCollectionTwo.dirMgrs[i]) {

			return false
		}

	}

	return true
}

// fmtTextListingAbsPath
//
// Formats and returns a text listing of all directory
// absolute paths contained in the Directory Manager
// Collection passed as input parameter 'dMgrCollection'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgrCollection				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		This method will format and return a text listing
//		comprised of all absolute directory paths in the
//		Directory Manager (DirMgr) array encapsulated by
//		this Directory Manager Collection instance.
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
//	topTitleDisplay				TextLineTitleMarqueeDto
//
//		Contains specifications for the top tile display
//		including title lines and solid line breaks.
//
//		If no title is required, set this parameter to an
//		empty instance of TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no top title text lines are required, and the
//		solid line breaks are still necessary, simply
//		leave the title lines collection empty.
//
//	bottomTitleDisplay			TextLineTitleMarqueeDto
//
//		Contains specifications for the bottom tile
//		display including title lines and solid line
//		breaks.
//
//		If no bottom title is required, set this
//		parameter to an empty instance of
//		TextLineTitleMarqueeDto.
//
//		Example:
//			titleMarquee = 	TextLineTitleMarqueeDto{}
//
//		All TextLineTitleMarqueeDto member data values
//		are public. Just set the data values as
//		necessary during creation of the
//		TextLineTitleMarqueeDto instance. Afterward, use
//		the 'Add' methods to add title lines to the
//		TextLineTitleMarqueeDto collection.
//
//		If no bottom title text lines are required, and
//		the solid line breaks are still necessary, simply
//		leave the title lines collection empty.
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
func (dMgrColHelper *dirMgrCollectionHelper) fmtTextListingAbsPath(
	dMgrCollection *DirMgrCollection,
	leftMargin string,
	rightMargin string,
	maxLineLength int,
	topTitleDisplay TextLineTitleMarqueeDto,
	bottomTitleDisplay TextLineTitleMarqueeDto,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrCollectionHelper.fmtTextListingAllFiles()",
		"")

	if err != nil {
		return err
	}

	if dMgrCollection == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgrCollection' is invalid!\n"+
			"'dMgrCollection' is a nil pointer.\n",
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

	lenLeftMar := len(leftMargin)

	lenRightMar := len(rightMargin)

	lenLineTerminator := 1

	netTextLineLen :=
		maxLineLength -
			lenLeftMar -
			lenRightMar -
			lenLineTerminator

	if netTextLineLen < 10 {

		return fmt.Errorf("%v\n"+
			"Error: The Net Text Line length is less than 10.\n"+
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

	lenDMgrCol := len(dMgrCollection.dirMgrs)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	thisReqCapacity := lenDMgrCol * 110

	netRequiredCapacity :=
		thisReqCapacity - netCapacityStrBuilder

	if netRequiredCapacity > 0 {

		strBuilder.Grow(netRequiredCapacity + 256)
	}

	txtFormatCol := TextFormatterCollection{}

	solidLineChar := "-"

	if topTitleDisplay.IsValidInstance() {

		err = txtFormatCol.AddTextTitleMarqueeDto(
			topTitleDisplay,
			ePrefix.XCpy("<-topTitleDisplay"))

		if err != nil {
			return err
		}

		solidLineChar = topTitleDisplay.LeadingSolidLineChar

		if len(topTitleDisplay.TrailingSolidLineChar) > 0 {
			solidLineChar = topTitleDisplay.TrailingSolidLineChar
		}

		if len(solidLineChar) == 0 {
			solidLineChar = "-"
		}

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
			leftMargin+"  ",
			ePrefix)

	if err != nil {
		return err
	}

	if lenDMgrCol == 0 {

		extraMargin := "  "
		err = txtFormatCol.AddLine1Col(
			extraMargin+
				"The Directory Manager Collection is Empty!",
			ePrefix)

		if err != nil {
			return err
		}

		err = txtFormatCol.AddLine1Col(
			extraMargin+
				"There are zero Directory Manager Objects in the Collection!",
			ePrefix)

		if err != nil {
			return err
		}

	} else {

		for i := 0; i < lenDMgrCol; i++ {

			err = txtFormatCol.AddLine1Col(
				dMgrCollection.dirMgrs[i].GetPathAbsolute(),
				ePrefix.XCpy(
					fmt.Sprintf("dMgrCollection.dirMgrs[%v]",
						i)))
			if err != nil {
				return err
			}

		}
	}

	if bottomTitleDisplay.IsValidInstance() {

		err = txtFormatCol.AddTextTitleMarqueeDto(
			bottomTitleDisplay,
			ePrefix.XCpy("<-bottomTitleDisplay"))

		if err != nil {
			return err
		}

	} else {

		// Final Text Line Build
		txtFormatCol.AddLineBlank(1, "")

	}

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// newEmptyDMgrCollection
//
// Creates and returns a new, empty and properly
// initialized Directory Manager Collection
// ('DirMgrCollection') instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	DirMgrCollection
//
//		This method returns a concrete instance of
//		DirMgrCollection. The returned instance consists
//		of an empty and properly initialized instance of
//		DirMgrCollection.
func (dMgrColHelper *dirMgrCollectionHelper) newEmptyDMgrCollection() DirMgrCollection {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	newDirMgrCol := DirMgrCollection{}

	newDirMgrCol.dirMgrs = make([]DirMgr, 0)

	return newDirMgrCol
}

// peekOrPopAtIndex
//
// Returns a deep copy of the Directory Manager
// ('DirMgr') object located at array index 'idx' in the
// Directory Manager Collection passed as input parameter
// 'dirMgrs'.
//
// If input parameter 'deleteIndex' is set to 'false',
// this method will function as a 'Peek' method and
// therefore, the original Directory Manager ('DirMgr')
// object will NOT be deleted from the Directory Manager
// Collection ('DirMgrCollection') array.
//
// If input parameter 'deleteIndex' is set to 'true',
// this method will function as a 'Pop' method and
// therefore, the original Directory Manager ('DirMgr')
// object WILL BE DELETED from the Directory Manager
// Collection ('DirMgrCollection') array. The deletion
// operation will be performed on the Directory Manager
// object residing at the Directory Manager Collection
// array index identified by input parameter 'idx'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryMgrs				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		A 'peek' or 'pop' operation will be performed
//		on the DirMgr object specified by the array index
//		'idx'.
//
//	idx							int
//
//		This integer value specifies the index of the
//		array element which will be extracted from the
//		Directory Manager Collection encapsulated by the
//		instance of DirMgrCollection passed by input
//		parameter 'directoryMgrs'.
//
//		If this value is less than zero an error will be
//		returned.
//
//		If 'idx' exceeds the last index in the collection,
//		an io.EOF (End-Of-File) error will be returned.
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
//	DirMgr
//
//		If this method completes successfully without
//		error, a deep copy of the Directory Manager
//		object residing at array index 'idx' will be
//		returned through this parameter.
//
//	ArrayColErrorStatus
//
//		This structure provides detailed error
//		information related to the completion of this
//		method. This structure is designed to convey
//		error status for operations involving arrays
//		or collections of objects
//
//		type ArrayColErrorStatus struct {
//
//			IsProcessingError bool
//				When set to 'true', this parameter signals
//				that an error was encountered during a
//				routine array or object collection
//				processing operation. In this case an
//				appropriate error message describing the
//				error will be recorded in data element
//				'ProcessingError'.
//
//			IsIndexOutOfBounds bool
//				When set to 'true', this parameter signals
//				that the index value used to access the array
//				or object collection was less than zero or
//				greater than the last index in the
//				array/collection.
//
//			IsArrayCollectionEmpty bool
//				When set to 'true', this parameter signals
//				that array or objects collections is empty.
//
//			IsErrorFree bool
//				When set to 'true', this parameter signals that
//				no errors were encountered in the most recent
//				array or collection operation. This also means
//				that data element 'ProcessingError' is set to
//				'nil'.
//
//			ProcessingError	error
//				If no errors were encountered in the most recent
//				array or object collection processing operation,
//				this error parameter will be set to nil.
//
//				If errors are encountered during an array or
//				object collection processing operation, this
//				error Type will encapsulate an appropriate error
//				message.
//		}
//
//		If this method completes successfully,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'true'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		be set to 'nil'.
//
//		If errors are encountered during processing,
//		ArrayColErrorStatus.IsErrorFree will
//		be set to 'false'. In addition,
//		ArrayColErrorStatus.ProcessingError will
//		encapsulate an appropriate error message.
//		This returned error message will incorporate
//		the method chain and text passed by input
//		parameter, 'errorPrefix'.
//
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgrColHelper *dirMgrCollectionHelper) peekOrPopAtIndex(
	directoryMgrs *DirMgrCollection,
	idx int,
	deleteIndex bool,
	errPrefDto *ePref.ErrPrefixDto) (
	DirMgr,
	ArrayColErrorStatus) {

	if dMgrColHelper.lock == nil {
		dMgrColHelper.lock = new(sync.Mutex)
	}

	dMgrColHelper.lock.Lock()

	defer dMgrColHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var errStatus ArrayColErrorStatus

	var err error

	funcName := "dirMgrCollectionHelper.peekOrPopAtIndex()"

	ePrefix,
		errStatus.ProcessingError =
		ePref.ErrPrefixDto{}.NewFromErrPrefDto(
			errPrefDto,
			funcName,
			"")

	if errStatus.ProcessingError != nil {

		errStatus.IsProcessingError = true

		return DirMgr{}, errStatus
	}

	if directoryMgrs == nil {

		errStatus.IsProcessingError = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'directoryMgrs' is a nil pointer!\n",
				ePrefix.String())

		return DirMgr{}, errStatus
	}

	if directoryMgrs.dirMgrs == nil {
		directoryMgrs.dirMgrs = make([]DirMgr, 0)
	}

	arrayLen := len(directoryMgrs.dirMgrs)

	if arrayLen == 0 {

		errStatus.IsArrayCollectionEmpty = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: The Directory Manager Collection array encapsulated by 'directoryMgrs' is EMPTY!\n",
				ePrefix.String())

		return DirMgr{}, errStatus
	}

	if idx < 0 {

		errStatus.IsIndexOutOfBounds = true

		errStatus.ProcessingError =
			fmt.Errorf("%v"+
				"Error: Input Parameter 'idx' is less than zero. "+
				"Index Out-Of-Range! idx='%v'",
				ePrefix.String(),
				idx)

		return DirMgr{}, errStatus
	}

	if idx >= arrayLen {

		errStatus.IsIndexOutOfBounds = true

		errStatus.ProcessingError =
			fmt.Errorf("%v"+
				"Error: Input Parameter 'idx' is greater than the\n"+
				"last index in the Directory Manager Collection array.\n"+
				"Index Out-Of-Range! idx='%v'",
				ePrefix.String(),
				idx)

		return DirMgr{}, errStatus
	}

	var deepCopyDirMgr DirMgr

	deepCopyDirMgr,
		err =
		directoryMgrs.dirMgrs[idx].CopyOut(
			ePrefix.XCpy(fmt.Sprintf(
				"directoryMgrs.directoryMgrs[%v]",
				idx)))

	if err != nil {

		errStatus.IsProcessingError = true

		errStatus.ProcessingError =
			fmt.Errorf("%v\n"+
				"Error: directoryMgrs.dirMgrs[%v].CopyOut()\n"+
				"directoryMgrs.dirMgrs index = '%v'\n"+
				"directoryMgrs.dirMgrs[%v] = '%v'\n"+
				"Error= \n%v\n",
				funcName,
				idx,
				idx,
				idx,
				directoryMgrs.dirMgrs[idx].absolutePath,
				err.Error())

		return DirMgr{}, errStatus
	}

	if deleteIndex == false {

		errStatus.IsErrorFree = true

		return deepCopyDirMgr, errStatus
	}

	// deleteIndex == true

	if arrayLen == 1 {

		directoryMgrs.dirMgrs = make([]DirMgr, 0, 1)

	} else if idx == 0 {
		// arrayLen > 1 and requested idx = 0
		directoryMgrs.dirMgrs = directoryMgrs.dirMgrs[1:]

	} else if idx == arrayLen-1 {
		// arrayLen > 1 and requested idx = last element index
		directoryMgrs.dirMgrs = directoryMgrs.dirMgrs[0 : arrayLen-1]

	} else {
		// arrayLen > 1 and idx is in between
		// first and last elements
		directoryMgrs.dirMgrs =
			append(directoryMgrs.dirMgrs[0:idx], directoryMgrs.dirMgrs[idx+1:]...)
	}

	errStatus.IsErrorFree = true

	return deepCopyDirMgr, errStatus
}
