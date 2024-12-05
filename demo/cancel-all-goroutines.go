// 当某个协程异常后，取消其他协程
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func handle(wg *sync.WaitGroup, ctx context.Context, quit chan bool, i int) {
	defer wg.Done()
	if i%2 == 0 {
		fmt.Println("send quit")
		quit <- true
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	default:
	}
	fmt.Println("handle", i)
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan bool)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go handle(&wg, ctx, quit, i)
	}

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("received quit")
				cancel()
			case <-done:
				fmt.Println("received done")
				return
			default:
			}
		}
	}()

	wg.Wait()
	done <- true
}
