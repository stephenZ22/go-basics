package main

import "fmt"

const X float32 = 3.14

// 同样可以使用显式类型转换
// const X = float32(3.14)

// 当然这里的现实类型转换必须是合法的
const a uint = 128

// 这里的 256 不能 显式转换为uint8
// error: 256溢出uint8
// const c = uint8(256)

const (
	//
	A, B int64 = -3, 5

	Y float32 = 2.718
)

// 常量的自动补全

func main() {
	fmt.Println("a =>", a)
	fmt.Println("A =>", A)
	fmt.Println("B =>", B)
	fmt.Println("Y =>", Y)
}
