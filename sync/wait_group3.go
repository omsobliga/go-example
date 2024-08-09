package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const N = 10
	var wgA sync.WaitGroup
	var wgB sync.WaitGroup
	wgA.Add(N)
	wgB.Add(1)

	var values [N]int32

	for i := 0; i < N; i ++ {
		i := i
		go func() {
			wgB.Wait()
			values[i] = rand.Int31n(50)
			fmt.Println(i)
			wgA.Done()
		}()
	}

	wgB.Done()
	wgA.Wait()
	fmt.Println(values)
}