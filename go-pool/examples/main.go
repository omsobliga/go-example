package main

import (
	"fmt"
	"go-pool"
	"sync"
)

func doSomething(wg *sync.WaitGroup, i int) {
	fmt.Println("do something", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	pool, err := go_pool.NewPool(32)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		err = pool.Submit(func() {
			doSomething(&wg, i)
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	wg.Wait()
}
