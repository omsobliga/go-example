// 在任何一个给定时刻，一个函数调用最多只能和一个未恢复的恐慌相关联。 如果一个调用正和一个未恢复的恐慌相关联，则
// - 在此恐慌被恢复之后，此调用将不再和任何恐慌相关联。
// - 当在此函数调用中产生了一个新的恐慌，此新恐慌将替换原来的未被恢复的恐慌做为和此函数调用相关联的恐慌。

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println(recover()) // 3
	}()
	defer panic(3) // 将替换恐慌2
	defer panic(2) // 将替换恐慌1
	defer panic(1) // 将替换恐慌0
	panic(0)
}
