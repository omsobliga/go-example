package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atomic_add() {
	var n int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&n, 1)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}

func atomic_add2() {
	var n atomic.Int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			n.Add(1)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n.Load())
}

func add() {
	var n int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			n += 1
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(n)
}

func main() {
	atomic_add()
	atomic_add2()
	add()
}
