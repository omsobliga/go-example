// 对象池创建设计模式用于根据预期需求准备和保留多个实例。
package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	ID int
}

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- &Object{ID: i}
	}
	return &p
}

func main() {
	p := New(5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case obj := <-*p:
				fmt.Println("get poll", i, obj)
				*p <- obj
			default:
				fmt.Println("no poll", i)
				return
			}
		}(i)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case obj := <-*p:
				fmt.Println("get poll", i, obj)
				*p <- obj
			default:
				fmt.Println("no poll", i)
				return
			}
		}(i)
	}
	wg.Wait()
}
