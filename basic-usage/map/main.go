package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("empty", m)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("set", m)

	v := m["k1"]
	fmt.Println("get", v)

	v = m["k3"]
	fmt.Println("get not exist key", v)

	v2, exist := m["k1"]
	fmt.Println("get", v2, "exist", exist)

	v2, exist = m["k3"]
	fmt.Println("get not exist key", v2, "exist", exist)

	delete(m, "k1")
	fmt.Println("delete", m)

	n := map[string]int{"a": 1, "b": 2}
	fmt.Println(n)
}
