// 并发发起请求，再读取结果

package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

type Result struct {
	Value string
}

func f(i int) Result {
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
	return Result{Value: strconv.Itoa(i)}
}

func main() {
	c1 := make(chan Result, 1)
	go func() {
		c1 <- f(1)
	}()

	c2 := make(chan Result, 1)
	go func() {
		c2 <- f(2)
	}()

	c3 := make(chan Result, 1)
	go func() {
		c3 <- f(3)
	}()

	for i := 0; i < 3; i++ {
		select {
		case r := <-c1:
			fmt.Println(r.Value)
		case r := <-c2:
			fmt.Println(r.Value)
		case r := <-c3:
			fmt.Println(r.Value)
		}
	}
}
