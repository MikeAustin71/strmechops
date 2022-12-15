package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayCollectionNanobot - Provides helper methods for type
// RuneArrayCollection.
type runeArrayCollectionNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceRuneArrayCol' to input parameter
// 'destinationRuneArrayCol'. Both instances are of type
// RuneArrayCollection.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationRuneArrayCol'
// will be overwritten.
//
// Also, DATA VALIDATION IS PERFORMED on input parameter,
// 'sourceRuneArrayCol'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationRuneArrayCol    *RuneArrayCollection
//	   - A pointer to a RuneArrayCollection instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNegNumResults'.
//
//	     'destinationNegNumResults' is the destination for this
//	     copy operation.
//
//
//	sourceRuneArrayCol         *RuneArrayCollection
//	   - A pointer to another RuneArrayCollection instance. All the
//	     member variable data values from this object will be
//	     copied to corresponding member variables in
//	     'destinationNegNumResults'.
//
//	     'sourceRuneArrayCol' is the source for this copy
//	     operation.
//
//	     If 'sourceRuneArrayCol' is empty an error will be
//	     returned.
//
//	     If 'sourceRuneArrayCol' contains RuneArrayDto objects with
//	     zero length character arrays, an error will be returned.
//
//	     If 'sourceRuneArrayCol' contains RuneArrayDto objects with
//	     invalid Character Search Type designations, an error will
//	     be returned.
//
//
//	errPrefDto                   *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (runeArrayColNanobot *runeArrayCollectionNanobot) copyIn(
	destinationRuneArrayCol *RuneArrayCollection,
	sourceRuneArrayCol *RuneArrayCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if runeArrayColNanobot.lock == nil {
		runeArrayColNanobot.lock = new(sync.Mutex)
	}

	runeArrayColNanobot.lock.Lock()

	defer runeArrayColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayCollectionNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationRuneArrayCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationRuneArrayCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceRuneArrayCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceRuneArrayCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err =
		runeArrayCollectionElectron{}.ptr().
			testValidityRuneArrayCollection(
				sourceRuneArrayCol.runeArrayDtoCol,
				ePrefix.XCpy(
					"sourceRuneArrayCol.runeArrayDtoCol"))

	if err != nil {

		return err

	}

	// lenSourceCol was validated. It MUST BE greater
	// than zero.
	lenSourceCol := len(sourceRuneArrayCol.runeArrayDtoCol)

	runeArrayCollectionAtom{}.ptr().
		empty(destinationRuneArrayCol)

	destinationRuneArrayCol.runeArrayDtoCol =
		make([]RuneArrayDto, lenSourceCol)

	for i := 0; i < lenSourceCol; i++ {

		err = destinationRuneArrayCol.runeArrayDtoCol[i].CopyIn(
			&sourceRuneArrayCol.runeArrayDtoCol[i],
			ePrefix.XCpy(
				fmt.Sprintf("destinationRuneArrayCol"+
					"<-sourceRuneArrayCol.runeArrayDtoCol[%v]",
					i)))

		if err != nil {
			return err
		}
	}

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'runeArrayCol', a pointer to an instance of
// RuneArrayCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// DATA VALIDATION IS PERFORMED on 'runeArrayCol'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayCol               *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. A deep
//	     copy of the internal member variables contained in this
//	     instance will be created and returned in a new instance of
//	     RuneArrayCollection.
//
//	     DATA VALIDATION IS performed on 'runeArrayCol'.
//
//	     If 'runeArrayCol' is empty with a length of zero,
//	     an error will be returned.
//
//	     If 'runeArrayCol' contains RuneArrayDto objects with zero
//	     length character arrays, an error will be returned.
//
//	     If 'runeArrayCol' contains RuneArrayDto objects with
//	     invalid Character Search Type designations, an error will
//	     be returned.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyRuneArrayCol       RuneArrayCollection
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'runeArrayCol' will be created and
//	     returned in a new instance of RuneArrayCollection.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (runeArrayColNanobot *runeArrayCollectionNanobot) copyOut(
	runeArrayCol *RuneArrayCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyRuneArrayCol RuneArrayCollection,
	err error) {

	if runeArrayColNanobot.lock == nil {
		runeArrayColNanobot.lock = new(sync.Mutex)
	}

	runeArrayColNanobot.lock.Lock()

	defer runeArrayColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayCollectionNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyRuneArrayCol, err

	}

	if runeArrayCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayCol' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyRuneArrayCol, err
	}

	_,
		err =
		runeArrayCollectionElectron{}.ptr().
			testValidityRuneArrayCollection(
				runeArrayCol.runeArrayDtoCol,
				ePrefix.XCpy(
					"runeArrayCol.runeArrayDtoCol"))

	if err != nil {

		return deepCopyRuneArrayCol, err

	}
	// lenRuneArrayCol was validated. It MUST BE greater
	// than zero.
	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	runeArrayCollectionAtom{}.ptr().
		empty(&deepCopyRuneArrayCol)

	deepCopyRuneArrayCol.runeArrayDtoCol =
		make([]RuneArrayDto, lenRuneArrayCol)

	for i := 0; i < lenRuneArrayCol; i++ {

		err = deepCopyRuneArrayCol.runeArrayDtoCol[i].CopyIn(
			&runeArrayCol.runeArrayDtoCol[i],
			ePrefix.XCpy(
				fmt.Sprintf("deepCopyRuneArrayCol"+
					"<-runeArrayDtoCol.runeArrayDtoCol[%v]",
					i)))

		if err != nil {
			return deepCopyRuneArrayCol, err
		}
	}

	return deepCopyRuneArrayCol, err
}

// ptr - Returns a pointer to a new instance of
// runeArrayCollectionNanobot.
func (runeArrayColNanobot runeArrayCollectionNanobot) ptr() *runeArrayCollectionNanobot {

	if runeArrayColNanobot.lock == nil {
		runeArrayColNanobot.lock = new(sync.Mutex)
	}

	runeArrayColNanobot.lock.Lock()

	defer runeArrayColNanobot.lock.Unlock()

	return &runeArrayCollectionNanobot{
		lock: new(sync.Mutex),
	}
}
