// 在一个代码包中，甚至一个源文件中，可以声明若干名为init的函数。这些init函数必须不带任何输入参数和返回结果。
// 在程序运行时刻，在进入main入口函数之前，每个init函数在此包加载的时候将被（串行）执行并且只执行一遍。
package main

import "fmt"

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")
