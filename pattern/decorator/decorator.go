// 装饰器结构模式允许动态扩展现有对象的功能，而无需改变其内部结构。
package main

import "log"

type Object func(int) int

func LogDecorate(f Object) Object {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)
		result := f(n)
		log.Println("Execution is completed with the result", result)
		return result
	}
}

func AddOne(n int) int {
	return n + 1
}

func main() {
	f := LogDecorate(AddOne)
	n := f(1)
	f(n)
}
