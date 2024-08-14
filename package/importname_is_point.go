// 一个完整引入声明语句形式的引入名importname可以是一个句点(.)。
// 这样的引入称为句点引入。使用被句点引入的包中的导出代码要素时，限定标识符的前缀必须省略。
// 在这个例子中，Println和Now函数调用不需要带任何前缀。
// 一般来说，句点引入不推荐使用，因为它们会导致较低的代码可读性。

package main

import (
	. "fmt"
	. "time"
)

func main() {
	Println("Current time:", Now())
}
