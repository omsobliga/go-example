// 一个被延迟调用的函数值是在其调用被推入延迟调用队列之前被估值的。下面这个例子将输出false。
package main

import "fmt"

func main() {
	var f = func () {
		fmt.Println(false)
	}
	defer f()
	f = func () {
		fmt.Println(true)
	}
}

