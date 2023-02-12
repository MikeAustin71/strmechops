package strmech

import "sync"

type filePermissionConfigMechanics struct {
	lock *sync.Mutex
}
