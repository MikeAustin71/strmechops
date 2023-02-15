package strmech

import (
	"sync"
)

type fileHelperAtom struct {
	lock *sync.Mutex
}
