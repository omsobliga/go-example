/*
实现原理
- WithErrors 后生成一个 ErrorPool 对象
- ErrorPool 中 errs 字段负责记录子协程的错误
- ErrorPool.Go 执行任务时，把返回的错误追加添加到 errs 字段
*/
package main

import (
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	p := pool.New().WithMaxGoroutines(10).WithErrors()
	for i := 0; i < 10; i++ {
		p.Go(func() error {
			return handleWithError(i)
		})
	}
	errs := p.Wait()
	fmt.Println(errs)
}

func handleWithError(i int) error {
	fmt.Println(i)
	if i < 5 {
		return errors.New(fmt.Sprintf("error: %d", i))
	} else {
		return nil
	}
}
