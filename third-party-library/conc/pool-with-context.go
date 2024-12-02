/*
通过 context 管理所有子协程，当错误发生时，终止其他子协程
*/
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	p := pool.New().WithMaxGoroutines(10).WithContext(context.Background()).WithCancelOnError()
	for i := 0; i < 10; i++ {
		p.Go(func(ctx context.Context) error {
			return handleWithContext(ctx, i)
		})
	}
	errs := p.Wait()
	fmt.Println(errs)
}

func handleWithContext(ctx context.Context, i int) error {
	select {
	case <-ctx.Done():
		return errors.New(fmt.Sprintf("%s: %d", ctx.Err(), i))
	default:
	}

	if i == 5 {
		return errors.New("err")
	}

	fmt.Println(i)
	return nil
}
