package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// range返回一对值: 索引和索引的值，如果不需要索引，只用索引的值可以用 "_" 标识符声明为空标识符
	// 不可以使用temp临时变量，因为go不允许存在没有使用的变量
	// 空标识符可用于任何语法需要变量名但程序逻辑不需要的时候
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

/**
s := "" 只能用在函数内部 不能用于包变量
var s string 没有赋值变量默认初始化零值机制 ""
var s = ""
var s string = ""
*/
