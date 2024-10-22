// Go支持对于结构体(struct)和接口(interfaces)的嵌入(embedding) 以表达一种更加无缝的组合(composition)类型
// 类似于其他语言的继承

package main

import "fmt"

type base struct {
	num int
}

func (b *base) describe() {
	fmt.Println("base.describe", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe()
}

func main() {
	// 当创建含有嵌入的结构体，必须对嵌入进行显式的初始化； 在这里使用嵌入的类型当作字段的名字
	c := container{
		base: base{
			num: 10,
		},
		str: "abc",
	}
	fmt.Println("container", c)
	fmt.Println("container.base.num", c.base.num)
	// 可以直接在 c 上访问 base 中定义的字段
	fmt.Println("container.num", c.num)
	fmt.Println("container.str", c.str)

	// 直接在 c 上 调用了一个从 base 嵌入的方法
	c.base.describe()
	c.describe()

	var d describer = &c
	d.describe()
}