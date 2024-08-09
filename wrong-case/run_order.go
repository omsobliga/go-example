// 一个协程的创建发生在此协程中的任何代码执行之前
package main

import (
	"fmt"
	"time"
)

var x, y int

func main() {
	go func() {
		fmt.Println(x) // 可能打印出0、123，或其它值
	}()
	go func() {
		fmt.Println(y) // 可能打印出0、789，或其它值
	}()
	x, y = 123, 789
	time.Sleep(time.Second)
}

