package main

import "fmt"

func runTaskWithResult(f func() interface{}) interface{} {
	return f()
}

func task1(i int) int {
	return i * i
}

func main() {
	r1 := runTaskWithResult(func() interface{} { return task1(1) })
	r2 := runTaskWithResult(func() interface{} { return task1(2) })
	r3 := runTaskWithResult(func() interface{} { return task1(3) })
	fmt.Println(r1, r2, r3)
}
