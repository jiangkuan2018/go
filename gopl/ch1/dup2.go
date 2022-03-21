package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println(counts)
}

func countLines(f *os.File, count map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		count[str]++
	}
	for line, n := range count {
		if n > 1 {
			fmt.Println(line)
		}
	}
}
