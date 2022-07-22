package main

import (
	"fmt"
	"math/rand"
)

func switch01() {
	switch n := rand.Intn(100); n % 9 {
	case 0:
		fmt.Println("被9整除")
	case 1, 3:
		fmt.Println("余 1, 3")
	// case 3: 这里编译失败 因为 3 以及在上一个case语句出现
	// 	fmt.Println("余 3")
	case 5:
		fmt.Println("余 5")
	case 7:
		fmt.Println("余 7")
	default:
		fmt.Println("余 2, 4, 6, 8, 9")
	}
}

func switch02() {

	switch n := 54; n % 9 {
	case 0:
		fmt.Println(n)
		fallthrough
	case 1, 3:
		n := 99
		fmt.Println(n)
	case 5:
		fmt.Println(n)
	case 7:
		fmt.Println(n)
	default:
		fmt.Println(n)
	}
}

func main() {
	switch01()
	switch02()
}
