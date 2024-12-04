// 字符串是由字节构成的，因此索引它们得到的是字节，而不是字符。
// 字符串甚至可能不包含字符。事实上，“字符”的定义是模棱两可的，试图通过定义字符串由字符组成来解决歧义是错误的。
// 在 Go 中，字符串实际上是一个只读的字节切片。字符串可以包含任意字节，不一定非得是 Unicode 文本或 UTF-8 文本。
// 在 Go 中，rune 是 int32 类型的别名，用于表示 Unicode 码点。

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%s, %+q, %x, ", runeValue, runeValue, runeValue)
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%s, %+q, %x, ", runeValue, runeValue, runeValue)
		fmt.Printf("%#U starts at %d, width %d\n", runeValue, i, width)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
