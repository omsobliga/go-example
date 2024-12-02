/*
协程池
通过 pool.Go 启动协程时，先判断是否达到限制，若没有达到，则起一个子协程，并把任务添加到 tasks 队列。
若达到限制，不再创建新的协程，只是把任务添加到 tasks 队列，若 tasks 已有值则阻塞。

备注
for range在遍历channel时，如果channel已经关闭，则会自动退出循环。
如果channel未关闭，for range会一直等待，直到接收到新的元素或者channel被关闭为止。
*/
package main

import (
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	p := pool.New().WithMaxGoroutines(10)
	for i := 0; i < 10; i++ {
		p.Go(func() {
			handle(i)
		})
	}
	p.Wait()
}

func handle(i int) {
	fmt.Println(i)
}
