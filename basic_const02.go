package main

import "fmt"

const (
	X float32 = 3.14
	Y         // 这里必须只有一个标识符
	Z         // 这里必须只有一个标识符

	A, B = "Go", "Ruby"
	C, _
)

func main() {
	fmt.Println("X =>", X)
	fmt.Println("Y =>", Y)
	fmt.Println("Z =>", Z)
	fmt.Println("A =>", A)
	fmt.Println("B =>", B)
	fmt.Println("C =>", C)
}
