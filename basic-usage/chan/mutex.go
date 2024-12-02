// 示例：使用通道实现互斥锁
package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter)
}
