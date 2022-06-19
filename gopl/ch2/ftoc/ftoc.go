package main

import (
	"fmt"
	"os"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	a := "100"
	fmt.Printf("%g F = %g C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))
	fmt.Printf("%g, %g \n", a, b)
	openTest()
	var x int
	print(&x)
	print("\n=================\n")
	print(x)
	var p = &x
	//print(p)
	fmt.Printf("\n====%g=====\n", p)
	print(*p)
	fmt.Printf("\n%g", *f())
	//print(f())
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
func openTest() {
	f, err := os.Open("/Users/jiangkuan/project/self/go/gopl/ch2/ftoc/temp")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
	f.Close()
}

const b string = "1000"

func f() *int {
	x := 1
	return &x
}
