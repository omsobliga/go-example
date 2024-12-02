// 示例：使用 select 实现定时器
package main

import (
	"fmt"
	"time"
)

type T = struct{}

func tick(d time.Duration) <-chan T {
	c := make(chan T)
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- T{}:
			default:
			}
		}
	}()
	return c
}

func main() {
	t := time.Now()
	for range tick(time.Second) {
		fmt.Println("tick", time.Since(t))
	}
}
