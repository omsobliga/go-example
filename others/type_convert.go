package main

func main() {
	// 类型[]int、IntSlice和MySlice共享底层类型：[]int。
	type IntSlice []int
	type MySlice  []int
	type Foo = struct{n int `foo`}
	type Bar = struct{n int `bar`}

	var s  = []int{}
	var is = IntSlice{}
	var ms = MySlice{}
	var x map[Bar]Foo
	var y map[Foo]Bar

	// 这两行隐式转换编译不通过。
	/*
	is = ms
	ms = is
	*/

	// 必须使用显式转换。
	is = IntSlice(ms)
	ms = MySlice(is)
	x = map[Bar]Foo(y)
	y = map[Foo]Bar(x)

	// 这些隐式转换是没问题的。
	s = is
	is = s
	s = ms
	ms = s
}
