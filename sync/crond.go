// crond 用法
//  一个c.Wait()调用将首先将当前协程推入到c所维护的等待协程队列；
// 一个c.Signal()调用将唤醒并移除c所维护的等待协程队列中的第一个协程（如果此队列不为空的话）
// 一个c.Broadcast()调用将唤醒并移除c所维护的等待协程队列中的所有协程（如果此队列不为空的话）
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要

	const N = 10
	var values [N]string

	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < N; i++ {
		d := time.Second * time.Duration(rand.Intn(10)) / 10
		go func(i int) {
			time.Sleep(d) // 模拟一个工作负载
			cond.L.Lock()
			// 下面的修改必须在cond.L被锁定的时候执行
			values[i] = string('a' + i)
			cond.Broadcast() // 可以在cond.L被解锁后发出通知
			cond.L.Unlock()
			// 上面的通知也可以在cond.L未锁定的时候发出。
			//cond.Broadcast() // 上面的调用也可以放在这里
		}(i)
	}

	// 此函数必须在cond.L被锁定的时候调用。
	checkCondition := func() bool {
		fmt.Println(values)
		for i := 0; i < N; i++ {
			if values[i] == "" {
				return false
			}
		}
		return true
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for !checkCondition() {
		cond.Wait() // 必须在cond.L被锁定的时候调用
	}
}
