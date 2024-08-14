// 事实上，一个引入声明语句的完整形式为：
// import importname "path/to/package"

// 其中引入名importname是可选的，它的默认值为被引入的包的包名（不是目录名）。
/*
import fmt "fmt"        // <=> import "fmt"
import rand "math/rand" // <=> import "math/rand"
import time "time"      // <=> import "time"
*/

// 如果一个包引入声明中的importname没有省略，则限定标识符使用的前缀必须为importname，而不是被引入的包的名称。

package main

import (
	format "fmt"
	random "math/rand"
	"time"
)

func main() {
	random.Seed(time.Now().UnixNano())
	format.Print("一个随机数:", random.Uint32(), "\n")

	// 下面这行编译不通过，因为rand不可识别。
	/*
	fmt.Print("一个随机数:", rand.Uint32(), "\n")
	*/
}

