// 在任何比较中，第一个比较值必须能被赋值给第二个比较值的类型，或者反之。
// 切片/映射/函数类型为不可比较类型，但是它们的值可以和类型不确定的预声明nil标识符比较。
// 此外：
// 两个类型不确定的布尔值可以相互比较。
// 两个类型不确定的数字值可以相互比较。
// 两个类型不确定的字符串值可以相互比较。
// 两个类型不确定的nil值不能相互比较。

package main

// 一些类型为不可比较类型的变量。
var s []int
var m map[int]int
var f func()()
var t struct {x []int}
var a [5]map[int]int

func main() {
	// 这些比较编译不通过。
	/*
	_ = s == s
	_ = m == m
	_ = f == f
	_ = t == t
	_ = a == a
	_ = nil == nil
	_ = s == interface{}(nil)
	_ = m == interface{}(nil)
	_ = f == interface{}(nil)
	*/

	// 这些比较编译都没问题。
	_ = s == nil
	_ = m == nil
	_ = f == nil
	_ = 123 == interface{}(nil)
	_ = true == interface{}(nil)
	_ = "abc" == interface{}(nil)
}
