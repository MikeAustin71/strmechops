package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// negNumSearchSpecCollectionNanobot - Provides helper methods for type
// NegNumSearchSpecCollection.
//
type negNumSearchSpecCollectionNanobot struct {
	lock *sync.Mutex
}

func (negNumSearchColNanobot negNumSearchSpecCollectionNanobot) copyIn(
	targetNegNumSearchCol *NegNumSearchSpecCollection,
	incomingNegNumSearchCol *NegNumSearchSpecCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchColNanobot.lock == nil {
		negNumSearchColNanobot.lock = new(sync.Mutex)
	}

	negNumSearchColNanobot.lock.Lock()

	defer negNumSearchColNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecCollectionNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetNegNumSearchCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNegNumSearchCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchCol' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	_,
		err2 = negNumSearchSpecCollectionAtom{}.ptr().
		testValidityOfNegNumSearchCol(
			incomingNegNumSearchCol,
			nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingNegNumSearchCol' failed!\n"+
			"This instance of NegNumSearchSpecCollection is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	targetNegNumSearchCol.Empty()

	lenIncomingNegNumSearchCol :=
		len(incomingNegNumSearchCol.negNumSearchSpecsCol)

	var deepCpyNegNumSearchSpec NegativeNumberSearchSpec

	for i := 0; i < lenIncomingNegNumSearchCol; i++ {

		deepCpyNegNumSearchSpec,
			err2 = incomingNegNumSearchCol.
			negNumSearchSpecsCol[i].CopyOut(
			nil)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Copy Operation from incomingNegNumSearchCol failed!\n"+
				"The input parameter targetNegNumSearchCol is now invalid.\n"+
				"incomingNegNumSearchCol.negNumSearchSpecsCol[%v].CopyOut()\n"+
				"returned the following error:\n"+
				"%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return err
		}

		targetNegNumSearchCol.negNumSearchSpecsCol =
			append(targetNegNumSearchCol.negNumSearchSpecsCol,
				deepCpyNegNumSearchSpec)
	}

	return err
}

// ptr - Returns a pointer to a new instance of
// negNumSearchColNanobot.
//
func (negNumSearchColNanobot negNumSearchSpecCollectionNanobot) ptr() *negNumSearchSpecCollectionNanobot {

	if negNumSearchColNanobot.lock == nil {
		negNumSearchColNanobot.lock = new(sync.Mutex)
	}

	negNumSearchColNanobot.lock.Lock()

	defer negNumSearchColNanobot.lock.Unlock()

	return &negNumSearchSpecCollectionNanobot{
		lock: new(sync.Mutex),
	}
}
