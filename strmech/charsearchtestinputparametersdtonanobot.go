package strmech

import "sync"

// charSearchTestInputParametersDtoNanobot - Provides helper methods for type
// CharSearchTestInputParametersDto.
//
type charSearchTestInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// charSearchTestInputParametersDtoNanobot.
//
func (searchTestInputParmsNanobot charSearchTestInputParametersDtoNanobot) ptr() *charSearchTestInputParametersDtoNanobot {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	return &charSearchTestInputParametersDtoNanobot{
		lock: new(sync.Mutex),
	}
}
