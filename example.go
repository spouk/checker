package check

import (
	"sync"
	"log"
	"os"
)

const (
	checker = "[check]"
)

type Check struct {
	sync.Pool
	sync.RWMutex
	log *log.Logger
}

func NewBrancCheck(count int) *Check {
	return &Check{
		log: log.New(os.Stderr, checker, log.Lshortfile | log.Ltime),
	}
}
func (c *Check) ShowVersion() {
	c.log.Printf("version check 0.0.1")
}
