package check

import "sync"

type Check struct {
	sync.Pool
	sync.RWMutex
}
type WorkerCheck struct {

}
func NewCheck(countWorker int) *Check{
	return &Check{}
}