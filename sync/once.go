package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup
	value := 0
	doSomething := func() {
		value++
		fmt.Println("hello")
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(doSomething)
			fmt.Println("world")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(value)
}
