package main

import "fmt"

func main() {
	s := make([]int, 3)
	fmt.Println("empty", s)

	s[0] = 1
	fmt.Println("set", s)

	fmt.Println("len(s)", len(s))

	s = append(s, 7)
	s = append(s, 8, 9)
	fmt.Println("append", s)

	s1_3 := s[1:3]
	fmt.Println("s1_3", s1_3)

	s1_ := s[1:]
	fmt.Println("s1_", s1_)

	_s3 := s[:3]
	fmt.Println("_s3", _s3)

	twoD := make([][]int, 2)
	for i := 0; i < 2; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD", twoD)
}
