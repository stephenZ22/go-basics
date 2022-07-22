package main

import (
	"fmt"
	"math/rand"
)

func printOdd(n uint32) {
	if n%2 == 0 {
		fmt.Println(n, "为偶数")
	} else {
		fmt.Println(n, "为奇数")
	}
}

func printFuzzBuzz(n uint32) {
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FuzzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fuzz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("Unknown")
	}
}
func main() {
	if a := rand.Int(); a > 1 {
		fmt.Println("2 > 1")
	} else if a == 1 {
		fmt.Println("a == 1")
	}

	// 下面这行输出 会编译错误
	// 未声明 变量 a
	// fmt.Println(a)
	// 有由此可知 if 语句中 initStatement 说声明的变量只在当前if语句代码块可见
	// 随着 if 语句的结束 说声明的变量也将结束其生命周期

	// 省略 initStatement

	c := rand.Uint32()
	printOdd(c)
	printFuzzBuzz(c)

}
