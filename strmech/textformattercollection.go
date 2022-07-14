package strmech

import "sync"

type TextFormatterCollection struct {
	formatterCol []ITTextFormatter

	lock *sync.Mutex
}
