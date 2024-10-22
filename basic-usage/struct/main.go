package main

import "fmt"

type Person struct {
	name string
	age int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(Person{"john", 20})
	fmt.Println(Person{name: "blob"})
	fmt.Println(Person{name: "alis", age: 10})

	fmt.Println(newPerson("hni"))

	p := Person{name: "ele", age: 10}
	fmt.Println(p)

	p.age = 20
	fmt.Println(p)

	_p := &p
	p.age = 30
	fmt.Println(*_p)
}