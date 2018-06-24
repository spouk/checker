package check

import "sync"

type Check struct {
	sync.Pool
	sync.RWMutex
}
