// 使用通道来同步协程之间的执行状态
// 当子协程结束后，主协程再退出

package main

import (
	"fmt"
	"time"
)

func worker(done chan int) {
	fmt.Println("worker running")
	time.Sleep(time.Second)
	done <- 1
}

func main() {
	done := make(chan int, 1)
	go worker(done)
	<- done
}