package main

import "fmt"

func main() {
	type Point struct {
		x, y int
	}
	var p Point
	p.x = 1
	p.y = 2
	fmt.Println(p)
}

func Exp() int {
	return 1
}
