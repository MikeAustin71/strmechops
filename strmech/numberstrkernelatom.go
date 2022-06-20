package strmech

import "sync"

type numberStrKernelAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelAtom.
//
func (numStrKernelAtom numberStrKernelAtom) ptr() *numberStrKernelAtom {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	return &numberStrKernelAtom{
		lock: new(sync.Mutex),
	}
}
