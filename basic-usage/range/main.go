package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for v := range s {
		fmt.Println(v)
	}

	for i, v := range s {
		fmt.Println(i, v)
	}

	kv := map[string]int{"a": 1, "b": 2}
	for k, v := range kv {
		fmt.Println(k, v)
	}

	for k := range kv {
		fmt.Println(k)
	}

	for i, c := range "os" {
		fmt.Println(i, c)
	}

}