/*
有序输出结果

实现原理：
- 每次 stream.Go 时，生成一个 resultChannel
- 把 resultChannel 加入到 stream.queue 中
- 任务的执行完成后，把结果写入到 resultChannel
*/
package main

import (
	"fmt"
	"github.com/sourcegraph/conc/stream"
)

func main() {
	s := stream.New().WithMaxGoroutines(10)
	var res []int
	for i := 0; i < 10; i++ {
		s.Go(func() stream.Callback {
			return func() {
				fmt.Println(i * i)
				res = append(res, i*i)
			}
		})
	}
	s.Wait()
	fmt.Println(res)
}
