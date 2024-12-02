// 示例：使用通道 + select 实现峰值限制

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
		select {
		case seat := <-bar:
			log.Println("顾客#", i, "获得位置#", seat)
			go bar.ServeCustomer(i, seat)
		default:
			log.Print("顾客#", i, "不愿等待而离去")
		}
		log.Println("协程数量", runtime.NumGoroutine())
	}

	for {
		time.Sleep(time.Second)
	}
}
