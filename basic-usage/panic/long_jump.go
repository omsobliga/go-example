// 用例3：使用panic/recover函数调用模拟长程跳转
// 使用panic/recover函数调用来模拟跨函数跳转，尽管一般这种方式并不推荐使用。
// 这种跳转方式的可读性不高，代码效率也不是很高，唯一的好处是它有时可以使代码看上去不是很啰嗦。

package main

import "fmt"

func main() {
	n := func () (result int)  {
		defer func() {
			if v := recover(); v != nil {
				if n, ok := v.(int); ok {
					result = n
				}
			}
		}()

		func () {
			func () {
				func () {
					// ...
					panic(123) // 用恐慌来表示成功返回
				}()
				// ...
			}()
		}()
		// ...
		return 0
	}()
	fmt.Println(n) // 123
}
