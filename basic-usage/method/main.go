// 调用方法时，Go 会自动处理值和指针之间的转换。
// 想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值， 你都可以使用指针来调用方法。

package main

import (
	"fmt"
	"unsafe"
)

type rect struct {
	height int
	weight int
}

func (r rect) area() int {
	fmt.Println(unsafe.Pointer(&r))
	return r.height * r.weight
}

func (r *rect) perim() int {
	fmt.Println(unsafe.Pointer(r))
	return 2 * (r.height + r.weight)
}

func main() {
	r := rect{height: 100, weight: 200}
	fmt.Println(r.area())  // 会产生一个拷贝
	fmt.Println(r.area())
	fmt.Println(r.perim())
	fmt.Println(r.perim())

	rp := &rect{height: 200, weight: 300}
	fmt.Println(rp.area())  // 会产生一个拷贝
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
	fmt.Println(rp.perim())
}
