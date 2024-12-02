// 功能：使用通道实现单对单通知
// 实现：非缓冲区通道，读取时阻塞，直到有写入
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	values := make([]byte, 32*1024)
	if _, err := rand.Read(values); err != nil {
		panic(err)
	}

	// fmt.Println(values)

	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		fmt.Println(1)
		done <- struct{}{} // 通知排序已完成
		fmt.Println(2)
	}()

	fmt.Println(3)
	d := <-done // 等待排序完成
	fmt.Println(d)
	fmt.Println(4)
	fmt.Println(values[0], values[len(values)-1])
}
