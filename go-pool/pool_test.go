package go_pool

import (
	"math"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	n = 10000000
)

var curMem uint64

func demoFunc() {
	time.Sleep(time.Duration(10) * time.Millisecond)
}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			demoFunc()
			wg.Done()
		}()
	}
	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem := mem.TotalAlloc/1024/1024 - curMem
	t.Logf("memory usage: %d MB", curMem)
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	pool, err := NewPool(math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < n; i++ {
		wg.Add(1)
		err = pool.Submit(func() {
			demoFunc()
			wg.Done()
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	wg.Wait()
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem := mem.TotalAlloc/1024/1024 - curMem
	t.Logf("memory usage: %d MB", curMem)
}
