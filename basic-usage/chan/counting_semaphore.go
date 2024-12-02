// 示例：使用通道实现计数信号量，限制最大并发数
// 不足：创建子协程后再获得位置，协程数量会大于 10

package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(customerId int) {
	log.Println("顾客#", customerId, "进入酒吧")
	seat := <-bar
	log.Println("顾客#", customerId, "获得位置#", seat)
	time.Sleep(time.Second * time.Duration(2*rand.Intn(6)))
	bar <- seat
	log.Println("顾客#", customerId, "离开了位置#", seat)
}

func main() {
	bar := make(Bar, 10) // 酒吧 10 个位置
	for i := 0; i < 10; i++ {
		bar <- Seat(i)
	}

	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 200)
		go bar.ServeCustomer(i)
		log.Println("协程数量", runtime.NumGoroutine())
	}

	for {
		time.Sleep(time.Second)
	}
}
