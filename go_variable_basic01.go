package main

import "fmt"

// 在函数外部声明的变量为全局变量或者包变量
// 在函数内部声明的变量为局部变量

// 标准形式 全局变量
var name, addr string = "张学友", "中国香港"
var age int = 62

// 标准形式利用类型推断
var school = "北京大学"
var since = 1898

// 标准形式的省略初始值, 利用各自类型的默认值声明
var language string
var compiled bool
var number int

func main() {
	fmt.Println(name, addr, age, school, language, compiled, number)

	// 短声明形式只可在函数内部使用，用来声明局部变量
	// 注意！！！
	// golang 当一个局部变量被声明之后至少要被使用一次
	// 否则编译器将会报错， 全局变量则不受此限制

	bookName := "局外人"
	fmt.Println(bookName)
}
