package main

import "fmt"

func sum(nums ...int) int {
	v := 0
	for n := range(nums) {
		v += n
	}
	return v
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}