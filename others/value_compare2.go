// 动态类型均为同一个不可比较类型的两个接口值的比较将产生一个恐慌
package main

func main() {
	type T struct {
		a interface{}
		b int
	}
	var x interface{} = []int{}
	var y = T{a: x}
	var z = [3]T{{a: y}}

	// 这三个比较均会产生一个恐慌。
	_ = x == x
	_ = y == y
	_ = z == z
}
