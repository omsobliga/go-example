package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	const N = 10
	var wg sync.WaitGroup
	wg.Add(N)
	var values [N]int32

	for i := 0; i < N; i ++ {
		i := i
		go func() {
			values[i] = rand.Int31n(50)
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(values)
}