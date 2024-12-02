/*
sync/atomic标准库包中提供的Value类型可以用来读取和修改任何类型的值。

类型*Value有几个方法：Load、Store、Swap和CompareAndSwap（其中后两个方法实在Go 1.17中引入的）。 
这些方法均以interface{}做为参数类型，所以传递给它们的实参可以是任何类型的值。 
但是对于一个可寻址的Value类型的值v，一旦v.Store方法（(&v).Store的简写形式）被曾经调用一次，
则传递给值v的后续方法调用的实参的具体类型必须和传递给它的第一次调用的实参的具体类型一致； 
否则，将产生一个恐慌。nil接口类型实参也将导致v.Store()方法调用产生恐慌。
*/
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct {a, b, c int}
	var ta = T{1, 2, 3}
	var v atomic.Value
	v.Store(ta)
	var tb = v.Load().(T)
	fmt.Println(tb)       // {1 2 3}
	fmt.Println(ta == tb) // true

	v.Store("hello") // 将导致一个恐慌
}
