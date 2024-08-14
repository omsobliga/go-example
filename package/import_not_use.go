// 除了匿名引入，其它引入必须在代码中被使用一次
package main

import (
	"net/http" // error: 引入未被使用
	. "time"   // error: 引入未被使用
)

import (
	format "fmt"  // okay: 下面被使用了一次
	_ "math/rand" // okay: 匿名引入
)

func main() {
	format.Println() // 使用"fmt"包
}
