package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
	for idx, args := range os.Args {
		res := strconv.Itoa(idx) + args
		fmt.Println(res)
	}
}
