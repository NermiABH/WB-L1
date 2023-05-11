// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x, y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Point структура точки
type Point struct {
	X float64
	Y float64
}

// NewPoint это конструктор в который мы подставляем параметры
func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

// Distance находит расстояние между текущей точкой и подставляемой
func (p *Point) Distance(d *Point) float64 {
	return math.Sqrt(math.Pow(p.X-d.X, 2) + math.Pow(p.Y-d.Y, 2))
}

func main() {
	A := NewPoint(2, 4)
	B := NewPoint(5, 1)
	fmt.Println(A.Distance(B))
}
