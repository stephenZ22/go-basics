package main

import "fmt"

// 常量 声明关键字 const

// 声明单独的包级常量
const Pi = 3.14
const Pai = Pi

// 声明一组包级常量
const (
	Age        = 25
	Name       = "张学友"
	Profession = "歌手"
	No         = !Yes
	Yes        = true
)

// 同时声明多个包级常量
const Addr, Number, Gender = "中国香港", "13100000000", "男"

func main() {
	fmt.Println(Pi)
	fmt.Println(Pai)
	fmt.Println(Age)
	fmt.Println(Name)
	fmt.Println(Profession)
	fmt.Println(Addr)
	fmt.Println(Number)
	fmt.Println(Gender)
	fmt.Println("No is", No)
	// fmt.Println("Yes is", Yes)

	// 声明局部常量
	const DoublePi = 2 * Pi
	fmt.Println("DoublePi is", DoublePi)
}
