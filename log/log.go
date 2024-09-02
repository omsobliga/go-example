package main

import (
	"log"
)

func main() {
	log.Println("hello world")
	v := "bob"
	log.Printf("hello %s\n", v)
	log.Fatalln("hello world, fatal log")
	log.Panicln("hello world, panic log")
}