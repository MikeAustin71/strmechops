package strmech

import "sync"

type textLineSpecStandardLineMolecule struct {
	lock *sync.Mutex
}

// copyOut - Returns a deep copy of the input parameter
// 'txtFieldLabel'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtFieldLabel       *TextFieldSpecLabel
//     - A pointer to an instance of TextFieldSpecLabel. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextFieldSpecLabel.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
//  TextFieldSpecLabel
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtFieldLabel' will be created and
//       returned in a new instance of TextFieldSpecLabel.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) copyOut(
	txtStdLine *TextLineSpecStandardLine) (
	TextLineSpecStandardLine, error) {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	return *txtStdLine, nil
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineMolecule.
//
func (txtStdLineMolecule textLineSpecStandardLineMolecule) ptr() *textLineSpecStandardLineMolecule {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	return &textLineSpecStandardLineMolecule{
		lock: new(sync.Mutex),
	}
}
