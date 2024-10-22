package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	height, weight float64
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c *circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

func (r *rect) area() float64 {
	return r.height * r.weight
}

func (r *rect) perim() float64 {
	return 2 * (r.height + r.weight)
}

func main() {
	r := &rect{height: 1, weight: 2}
	c := &circle{radius: 1}
	measure(r)
	measure(c)
}