package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// negNumSearchSpecColectionAtom - Provides helper methods for type
// NegNumSearchSpecCollection.
//
type negNumSearchSpecColectionAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// negNumSearchColAtom.
//
func (negNumSearchColAtom negNumSearchSpecColectionAtom) ptr() *negNumSearchSpecColectionAtom {

	if negNumSearchColAtom.lock == nil {
		negNumSearchColAtom.lock = new(sync.Mutex)
	}

	negNumSearchColAtom.lock.Lock()

	defer negNumSearchColAtom.lock.Unlock()

	return &negNumSearchSpecColectionAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfNegNumSearchCol - Receives a pointer to an
// instance of NegNumSearchSpecCollection and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'negNumSearchCol' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'negNumSearchCol' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchCol            *NegNumSearchSpecCollection
//     - A pointer to an instance of NegNumSearchSpecCollection.
//       This object will be subjected to diagnostic analysis in
//       order to determine if all the member variables contain
//       valid values.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If input parameter 'negNumSearchCol' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'negNumSearchCol' is found to be
//       invalid, this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'negNumSearchCol' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'negNumSearchCol' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (negNumSearchColAtom *negNumSearchSpecColectionAtom) testValidityOfNegNumSearchCol(
	negNumSearchCol *NegNumSearchSpecCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecColectionAtom."+
			"testValidityOfNegNumSearchCol()",
		"")

	if err != nil {

		return isValid, err
	}

	if negNumSearchCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSearchCol' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenNegNumSearchCol :=
		len(negNumSearchCol.negNumSearchSpecsCol)

	if lenNegNumSearchCol == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Negative Number Search Specification Collection is empty!\n"+
			"'negNumSearchCol.negNumSearchSpecsCol' has a length of zero.\n",
			ePrefix.String())

		return isValid, err

	}

	var err2 error

	for i := 0; i < lenNegNumSearchCol; i++ {

		err2 =
			negNumSearchCol.negNumSearchSpecsCol[i].IsValidInstanceError(
				nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: An element of the Negative Number Search\n"+
				"Specification Collection is invalid!\n"+
				"negNumSearchSpecsCol[%v] generated the following\n"+
				"validation error:\n"+
				"%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
