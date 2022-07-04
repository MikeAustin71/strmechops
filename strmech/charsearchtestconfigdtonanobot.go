package strmech

import "sync"

// charSearchTestConfigDtoNanobot - Provides helper methods for type
// CharSearchTestConfigDto.
//
type charSearchTestConfigDtoNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoNanobot.
//
func (searchTestConfigNanobot charSearchTestConfigDtoNanobot) ptr() *charSearchTestConfigDtoNanobot {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	return &charSearchTestConfigDtoNanobot{
		lock: new(sync.Mutex),
	}
}
