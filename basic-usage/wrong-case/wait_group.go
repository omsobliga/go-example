/*
Add方法的调用位置是不合适的。 此例子程序的打印结果并不总是100，而可能是0到100间的任何一个值。 原因是没有任何一个Add方法调用可以确保发生在唯一的Wait方法调用之前，结果导致没有任何一个Done方法调用可以确保发生在唯一的Wait方法调用返回之前。
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	fmt.Println("等待片刻...")
	wg.Wait()
	fmt.Println(atomic.LoadInt32(&x))
}
