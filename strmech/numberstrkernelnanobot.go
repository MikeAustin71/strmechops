package strmech

import "sync"

type numberStrKernelNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelNanobot.
//
func (numStrKernelNanobot numberStrKernelNanobot) ptr() *numberStrKernelNanobot {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	return &numberStrKernelNanobot{
		lock: new(sync.Mutex),
	}
}
