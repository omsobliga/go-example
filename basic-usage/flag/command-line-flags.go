package main

import (
	"flag"
	"fmt"
)

func main() {
	word := flag.String("word", "default_str", "a string")
	numb := flag.Int("numb", -1, "a int")
	fork := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "default_str", "a string")
	flag.Parse()

	fmt.Println("word:", *word)
	fmt.Println("numb:", *numb)
	fmt.Println("fork:", *fork)
	fmt.Println("svar:", svar)
	fmt.Println("flag.Args():", flag.Args())
}
