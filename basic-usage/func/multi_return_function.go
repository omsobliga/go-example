package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	fmt.Println(vals())
}