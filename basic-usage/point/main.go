package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPoint(val *int) {
	*val = 0
}

func main() {
	val := 1
	zeroVal(val)
	fmt.Println(val)

	zeroPoint(&val)
	fmt.Println(val)
}
