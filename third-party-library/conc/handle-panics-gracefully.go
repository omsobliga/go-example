/*
处理逻辑
1. 捕获子协程的 panic 并存储起来
2. 捕获到 panic 后，不影响其他 goroutine 执行
3. wg.Wait 前把 panic 抛出
*/
package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"time"
)

func main() {
	var wg conc.WaitGroup
	wg.Go(doSomethingThatMightPanic1)
	wg.Go(doSomethingThatMightPanic2)
	wg.Go(doSomethingNoPanic)
	wg.Go(doSomethingLongTime)
	fmt.Println("Waiting")
	wg.Wait()
	fmt.Println("End")
}

func doSomethingThatMightPanic1() {
	panic("test panic1")
}

func doSomethingThatMightPanic2() {
	panic("test panic2")
}

func doSomethingNoPanic() {
	fmt.Println("doSomethingNoPanic")
}

func doSomethingLongTime() {
	time.Sleep(3 * time.Second)
	fmt.Println("doSomethingLongTime")
}
