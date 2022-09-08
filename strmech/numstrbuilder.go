package strmech

import "sync"

type NumStrBuilder struct {
	lock *sync.Mutex
}
