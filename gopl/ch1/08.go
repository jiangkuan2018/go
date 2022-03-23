package main

import "fmt"

func main() {
	heads := 0
	tails := 0
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
	res := Signum(10)
	fmt.Println(res)

}
func coinflip() string {
	return "heads"
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
	}
}
