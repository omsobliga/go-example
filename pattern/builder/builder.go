// 构建器模式将复杂对象的构建与表示分离开来，因此同一构建过程可以创建不同的表示。
package main

import "fmt"

type Speed int

type Color string

type Builder interface {
	Build() Interface
	Color(Color) Builder
	Speed(Speed) Builder
}

type Interface interface {
	Drive()
}

type CarBuilder struct {
	speed Speed
	color Color
}

func (c *CarBuilder) Build() Interface {
	return &Car{
		speed: c.speed,
		color: c.color,
	}
}

func (c *CarBuilder) Color(color Color) Builder {
	c.color = color
	return c
}

func (c *CarBuilder) Speed(speed Speed) Builder {
	c.speed = speed
	return c
}

type Car struct {
	speed Speed
	color Color
}

func (c *Car) Drive() {
	fmt.Printf("Car is driving, color: %s, speed: %d\n", c.color, c.speed)
}

func NewCarBuilder() Builder {
	return &CarBuilder{}
}

func main() {
	carBuilder := NewCarBuilder()
	carBuilder.Color("red").Speed(100).Build().Drive()
	carBuilder.Color("white").Speed(200).Build().Drive()
}
