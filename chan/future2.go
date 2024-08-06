// 功能：使用通道实现 future/promise
// 实现：写通道作为函数参数
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan <- int32) {
	time.Sleep(time.Second * 3)	
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	a, b := make(chan int32), make(chan int32)
	go longTimeRequest(a)
	go longTimeRequest(b)
	fmt.Println(sumSquares(<-a, <-b))
}
