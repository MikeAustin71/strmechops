package strmech

import "sync"

type fileHelperPreon struct {
	lock *sync.Mutex
}
