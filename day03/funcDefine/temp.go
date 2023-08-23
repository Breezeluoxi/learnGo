package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Point2 struct {
	x, y int
}

func (p *Point2) Scale(factor int) {
	p.x *= factor
	p.y *= factor
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.Distance()) // 调用方法，不会修改原始结构体 p

	//p := &Point2{3, 4}
	//p.Scale(2) // 调用方法，会直接修改原始结构体 p
	//fmt.Println(p) // 输出 &{6 8}
}
