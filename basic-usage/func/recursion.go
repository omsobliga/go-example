// 递归

package main

import "fmt"

func fib(i int) int {
	if i >= 2 {
		return fib(i - 1) + fib(i - 2)
	} else if i == 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
}