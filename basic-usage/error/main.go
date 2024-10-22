// 符合 Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。
// 这与 Java、Ruby 使用的异常（exception） 以及在 C 语言中有时用到的重载 (overloaded) 的单返回/错误值有着明显的不同。
// Go 语言的处理方式能清楚的知道哪个函数返回了错误，并使用跟其他（无异常处理的）语言类似的方式来处理错误。

// 错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口。errors.New 使用给定的错误信息构造一个基本的 error 值。
// 也可以通过实现 Error() 方法来自定义 error 类型。
// 返回错误值为 nil 代表没有错误。

package main

import (
	"errors"
	"fmt"
)

func f1(age int) (int, error) {
	if age == 42 {
		return -1, errors.New("cant work with 42")
	} else {
		return age + 3, nil
	}
}

type argError struct {
	arg  int
	desc string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.desc)
}

func f2(age int) (int, *argError) {
	if age == 42 {
		return -1, &argError{arg: age, desc: "cant work with 42"}
	} else {
		return age + 3, nil
	}
}

func main() {
	for _, i := range []int{1, 42} {
		if r, e := f1(i); e == nil {
			fmt.Println("f1 worked, age:", r)
		} else {
			fmt.Println("f1 failed, error:", e)
		}
	}
	for _, i := range []int{1, 42} {
		if r, e := f2(i); e == nil {
			fmt.Println("f2 worked, age:", r)
		} else {
			fmt.Println("f2 failed, error:", e)
		}
	}
}
