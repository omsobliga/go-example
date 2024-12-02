// 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 "ping"。

package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		msg := <- messages
		fmt.Println(msg)
	}()
	messages <- "ping"
}
