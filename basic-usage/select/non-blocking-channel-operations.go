// 常规的通过通道发送和接收数据是阻塞的。
// 然而，我们可以使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收，甚至是非阻塞的多路 select。

package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case ch1 <- "msg":
		fmt.Println("send message to ch1")
	default:
		fmt.Println("send failed")
	}

	select {
	case <- ch1:
		fmt.Println("received message from ch1")
	default:
		fmt.Println("receive failed")
	}

	select {
	case <- ch1:
		fmt.Println("received message from ch1")
	case <- ch2:
		fmt.Println("received message from ch2")
	default:
		fmt.Println("both receive failed")
	}
}