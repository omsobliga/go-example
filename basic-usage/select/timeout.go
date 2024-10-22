package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "result1"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout1")
	}

	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "result2"
	}()

	select {
	case msg := <-ch2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout2")
	}
}
