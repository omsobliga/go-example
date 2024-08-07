// 示例：使用 select 实现超时机制
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func request(c chan<- int) {
	time.Sleep(time.Second * 3)
	c <- rand.Intn(100)
}

func requestWithTimeout(timeout time.Duration) {
	c := make(chan int)
	go request(c)

	select {
	case data := <-c:
		fmt.Println(data, "done")
	case <-time.After(timeout):
		fmt.Println("timeout")
	}
}

func main() {
	requestWithTimeout(time.Second)
	requestWithTimeout(time.Second * 2)
	requestWithTimeout(time.Second * 3)
	requestWithTimeout(time.Second * 4)
}
