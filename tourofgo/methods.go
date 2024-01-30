package main

import "fmt"

type Vertex struct {
	x, y float64
}

func (v Vertex) mulFunc() float64 {
	return v.x * v.y
}
func main() {

	m := Vertex{3.0, 8.0}
	fmt.Println(m.mulFunc())

}
