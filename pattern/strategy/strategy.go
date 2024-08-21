// 策略行为设计模式可以在运行时选择算法的行为。
package main

import "fmt"

type Operate interface{
	Apply(int, int) int
}

type Operation struct {
	operate Operate
}

func (o *Operation) Operate(left, right int) int {
	return o.operate.Apply(left, right)
}

type Add struct{}

func (a *Add) Apply(left, right int) int {
	return left + right
}

type Multiply struct{}

func (m *Multiply) Apply(left, right int) int {
	return left * right
}

func main() {
	operation := Operation{operate: &Add{}}
	fmt.Println(operation.Operate(3, 4))

	operation = Operation{operate: &Multiply{}}
	fmt.Println(operation.Operate(3, 4))
}