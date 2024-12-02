// 当一个字符串被转换为一个字节切片时，结果切片中的底层字节序列是此字符串中存储的字节序列的一份深复制。
// 即Go运行时将为结果切片开辟一块足够大的内存来容纳被复制过来的所有字节。当此字符串的长度较长时，此转换开销是比较大的。
// 同样，当一个字节切片被转换为一个字符串时，此字节切片中的字节序列也将被深复制到结果字符串中。 当此字节切片的长度较长时，此转换开销同样是比较大的。
// 在这两种转换中，必须使用深复制的原因是字节切片中的字节元素是可修改的，但是字符串中的字节是不可修改的，所以一个字节切片和一个字符串是不能共享底层字节序列的。
package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Runes2Bytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func main1() {
	s := "颜色感染是一个有趣的游戏。"
	fmt.Println(s)
	bs := []byte(s) // string -> []byte
	fmt.Println(bs)
	s = string(bs)  // []byte -> string
	fmt.Println(s)
	rs := []rune(s) // string -> []rune
	fmt.Println(rs)
	s = string(rs)  // []rune -> string
	fmt.Println(s)
	rs = bytes.Runes(bs) // []byte -> []rune
	fmt.Println(rs)
	bs = Runes2Bytes(rs) // []rune -> []byte
	fmt.Println(bs)
}

// 标准编译器做了一些优化，从而在某些情形下避免了深复制。 至少这些优化在当前（Go官方工具链1.22版本）是存在的。 这样的情形包括：
// * 一个for-range循环中跟随range关键字的从字符串到字节切片的转换；
// * 一个在映射元素读取索引语法中被用做键值的从字节切片到字符串的转换（注意：对修改写入索引语法无效）；
// * 一个字符串比较表达式中被用做比较值的从字节切片到字符串的转换；
// * 一个（至少有一个被衔接的字符串值为非空字符串常量的）字符串衔接表达式中的从字节切片到字符串的转换。

func main2() {
	var str = "world"
	// 这里，转换[]byte(str)将不需要一个深复制。
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	// 这个string(key)转换仍然需要深复制。
	m[string(key)] = "value"
	// 这里的转换string(key)将不需要一个深复制。
	// 即使key是一个包级变量，此优化仍然有效。
	fmt.Println(m[string(key)]) // value
}

var s string
var x = []byte{1023: 'x'}
var y = []byte{1023: 'y'}

func fc() {
	// 下面的四个转换都不需要深复制。
	if string(x) != string(y) {
		s = (" " + string(x) + string(y))[1:]
	}
}

func fd() {
	// 两个在比较表达式中的转换不需要深复制，
	// 但两个字符串衔接中的转换仍需要深复制。
	// 请注意此字符串衔接和fc中的衔接的差别。
	if string(x) != string(y) {
		s = string(x) + string(y)
	}
}

func main3() {
	fc()
	fd()
	fmt.Println(s)
}

func main() {
	main1()
	main2()
	main3()
}