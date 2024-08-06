// 示例：使用通道实现计数信号量，限制最大并发数
// 获得位置后再创建子协程，协程数量不会大于 10
// 不足：仍然会大量创建和销毁协程

package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(customerId int, seat Seat) {
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
		log.Println("顾客#", i, "进入酒吧")
		seat := <-bar
		log.Println("顾客#", i, "获得位置#", seat)
		go bar.ServeCustomer(i, seat)
		log.Println("协程数量", runtime.NumGoroutine())
	}

	for {
		time.Sleep(time.Second)
	}
}
