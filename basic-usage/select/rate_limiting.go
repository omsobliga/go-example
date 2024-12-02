// 示例：使用 chan + select 实现限流
package main

import (
	"fmt"
	"runtime"
	"time"
)

type Request interface{}

const RateLimitPeriod = time.Minute
const RateLimit = 100 // 每分钟 100 次

func handleRequest(req Request) {
	fmt.Println(req.(int))
}

func handleRequests(requests <-chan Request) {
	parallel := make(chan struct{}, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		for range tick.C {
			select {
			case parallel <- struct{}{}:
			default:
			}
		}
	}()

	for r := range requests {
		<-parallel
		go handleRequest(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	for i := 0; ; i++ {
		requests <- i
		if i%5 == 0 {
			fmt.Println("协程数量", runtime.NumGoroutine())
		}
	}
}
