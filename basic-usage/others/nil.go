// 访问一个nil映射将得到此映射的类型的元素类型的零值。

package main

import "fmt"

func main() {
    fmt.Println( (map[string]int)(nil)["key"] ) // 0
    fmt.Println( (map[int]bool)(nil)[123] )     // false
    fmt.Println( (map[int]*int64)(nil)[123] )   // <nil>
}

