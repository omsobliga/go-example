// 同时等待多个通道操作
package main

import (
	"fmt"
	"time"
)

func worker1(c chan string) {
	time.Sleep(time.Second)
	c <- "worker1"
}

func worker2(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "worker2"
}

func main() {
	chan1 := make(chan string, 1)
	chan2 := make(chan string, 2)
	go worker1(chan1)
	go worker2(chan2)
	for i := 0; i < 2; i++ {
		select {
		case msg := <-chan1:
			fmt.Println(msg)
		case msg := <-chan2:
			fmt.Println(msg)
		}
	}
}
