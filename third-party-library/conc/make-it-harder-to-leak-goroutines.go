/*
通过 wg 管理 goroutine，避免了 goroutine 被泄露
*/
package main

import (
	"github.com/sourcegraph/conc"
	"time"
)

func main() {
	var wg conc.WaitGroup
	defer wg.Wait()

	startTheThing(&wg)
}

func startTheThing(wg *conc.WaitGroup) {
	wg.Go(func() {
		time.Sleep(1 * time.Second)
	})
}
