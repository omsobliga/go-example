// 并发发起请求，再读取结果
// invokeService 负责执行任务

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

func invokeService(f func() interface{}) chan interface{} {
	ch := make(chan interface{})
	go func() {
		ch <- f()
	}()
	return ch

}

func f(i int) Result {
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
	return Result{Value: strconv.Itoa(i)}
}

func main() {
	c1 := invokeService(func() interface{} { return f(1) })
	c2 := invokeService(func() interface{} { return f(2) })
	c3 := invokeService(func() interface{} { return f(3) })
	for i := 0; i < 3; i++ {
		select {
		case r := <-c1:
			fmt.Println(r.(Result).Value)
		case r := <-c2:
			fmt.Println(r.(Result).Value)
		case r := <-c3:
			fmt.Println(r.(Result).Value)
		}
	}
}
