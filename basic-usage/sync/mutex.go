package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Counter struct {
	lock sync.Mutex
	n    uint64
}

func (c *Counter) Increase(delta uint64) {
	defer c.lock.Unlock()
	c.lock.Lock()
	c.n += delta
}

func (c *Counter) Value() uint64 {
	defer c.lock.Unlock()
	c.lock.Lock()
	return c.n
}

func main() {
	var c Counter
	for i := 0; i < 1000; i++ {
		go func() {
			c.Increase(1)
		}()
	}
	for c.Value() < 1000 {
		runtime.Gosched()
	}
	fmt.Println(c.Value())
}
