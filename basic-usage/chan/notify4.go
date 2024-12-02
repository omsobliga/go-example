// 功能：使用通道实现多对单通知，单对多通知
// 实现：从一个已关闭的通道可以接收到无穷个值，利用这一特性来实现群发通知
package main

import "log"
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // 阻塞在此，等待通知
	log.Print("Worker#", id, "开始工作")
	// 模拟一个工作负载。
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, "工作完成")
	done <- T{} // 通知主协程（N-to-1）
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(time.Second * 3 / 2)
	// 单对多通知
	close(ready)
	// 等待被多对单通知
	<-done; <-done; <-done
}