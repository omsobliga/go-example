// 并发发起请求，再读取结果
// invokeService 负责执行任务，并支持超时结束任务

package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

type Result struct {
	Value string
}

func invokeService(ctx context.Context, f func() interface{}) chan interface{} {
	ch := make(chan interface{})
	go func() {
		time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			ch <- Result{}
			return
		default:
		}
		ch <- f()
	}()
	return ch
}

func f(i int) Result {
	return Result{Value: strconv.Itoa(i)}
}

func main() {
	ctx, canc := context.WithCancel(context.Background())
	c1 := invokeService(ctx, func() interface{} { return f(1) })
	c2 := invokeService(ctx, func() interface{} { return f(2) })
	c3 := invokeService(ctx, func() interface{} { return f(3) })

	timeout := time.After(200 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case r := <-c1:
			fmt.Println(r.(Result).Value)
		case r := <-c2:
			fmt.Println(r.(Result).Value)
		case r := <-c3:
			fmt.Println(r.(Result).Value)
		case <-timeout:
			canc()
			break
		}
	}
}
