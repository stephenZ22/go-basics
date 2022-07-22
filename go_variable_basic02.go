package main

import (
	"fmt"
)

const b = 50

var a = 100

func main() {
	fmt.Println()

	// 这里的a 函数外部的全局变量（同名）进行了遮挡
	a := true

	{
		// 重新声明 左边的a b
		//右边的a 为当前代码块外层的bool类型的a
		// 右边的b 为 全局变量 const b

		a, b := a, b-10

		//
		//
		a, c := !a, b/2
		fmt.Println(a, b, c)
	}

	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c) 编译报错 c 由上文代码块定义外部不可见

}
