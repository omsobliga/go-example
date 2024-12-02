/*
实现原理类似 ErrorPool
*/
package main

import (
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

type T struct {
	value int
}

func main() {
	p := pool.NewWithResults[*T]().WithMaxGoroutines(10)
	for i := 0; i < 10; i++ {
		p.Go(func() *T {
			return handleWithResult(i)
		})
	}
	res := p.Wait()
	for _, t := range res {
		fmt.Println(t.value)
	}
}

func handleWithResult(i int) *T {
	return &T{value: i * i}
}
