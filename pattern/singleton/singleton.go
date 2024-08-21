// 单例创建设计模式将类型的实例化限制为单个对象。
package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var (
	once sync.Once
	ins  singleton
)

func New() singleton {
	once.Do(func() {
		ins = singleton{}
	})
	return ins
}

func main() {
	ins := New()
	ins["key"] = "value"
	fmt.Println(ins["key"])
	ins2 := New()
	fmt.Println(ins2["key"])
}
