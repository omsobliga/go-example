// 示例：使用通道实现定时器
package main

import (
	"fmt"
	"time"
)

func afterDuration(t time.Duration) <- chan struct{} {
	done := make(chan struct{})
	go func() {
		time.Sleep(t)
		done <- struct{}{}
	}()
	return done
}

func main() {
	fmt.Println("1")
	<- afterDuration(time.Second)
	fmt.Println("2")
	<- afterDuration(time.Second)
	fmt.Println("3")
}