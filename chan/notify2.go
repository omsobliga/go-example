// 功能：使用通道单对单通知
// 实现：非缓冲区通道，写入时阻塞，直到有读取
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// 此信号通道也可以缓冲为1。如果这样，则在下面
	// 这个协程创建之前，我们必须向其中写入一个值。

	go func() {
		// 模拟一个工作负载。
		time.Sleep(time.Second * 2)

		// 使用一个接收操作来通知主协程。
		fmt.Println(1)
		d := <-done
		fmt.Println(d)
		fmt.Println(2)
	}()

	fmt.Println(3)
	done <- struct{}{} // 阻塞在此，等待通知
	time.Sleep(time.Second * 2)
	fmt.Println(4)
}
