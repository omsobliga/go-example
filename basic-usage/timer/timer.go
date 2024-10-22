// 定时器：在未来某一时刻执行
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(time.Second * 2)

	go func() {
		<- timer2.C
		fmt.Println("timer2 fired")
	}()

	timer2.Stop()
	fmt.Println("timer2 stop")

	time.Sleep(time.Second * 2)
}
