package main

import "fmt"

func ping(pingChan chan <- string, msg string) {
	pingChan <- msg
}

func pong(pingChan <- chan string, pongChan chan <- string) {
	msg := <- pingChan
	pongChan <- msg
}

func main() {
	msg := "msg"
	pingChan := make(chan string, 1)
	pongChan := make(chan string, 1)
	ping(pingChan, msg)
	pong(pingChan, pongChan)
	fmt.Println(<-pongChan)
}