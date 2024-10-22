// 一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。
package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "aa"
	ch <- "bb"
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}