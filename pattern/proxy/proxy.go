// 代理模式提供了一个控制另一个对象访问的对象，拦截所有调用。
package main

import "fmt"

type Object struct {}

func (o *Object) Do(action string) {
	fmt.Printf("I can, %s\n", action)
}

type ObjectProxy struct {
	object *Object
}

func (o *ObjectProxy) Do(action string) {
	if o.object == nil {
		o.object = &Object{}
	}
	if action == "Run" {
		o.object.Do(action)
	}
}

func main() {
	proxy := &ObjectProxy{}
	proxy.Do("Run")
}