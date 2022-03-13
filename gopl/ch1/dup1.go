package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//testmap := make(map[string]int)
	//testmap[string('A')] = 1
	//testmap[string('B')] = 2
	//for key, value := range testmap {
	//	fmt.Printf(key, value)
	//	//fmt.Printf("%d\t%s\n", key, value)
	//}
	// make函数创建 slice, map, or chan (only) 类型
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin) // bufio Scanner读取输入并将其拆成行或单词，通常是处理行形式的输入最简单的方法
	fmt.Println(input)
	for input.Scan() { // Scan 函数读到一行时返回true，无输入时返回false
		counts[input.Text()]++ // 读取的内容使用Text函数获取
		for line, n := range counts {
			if n > 2 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}
