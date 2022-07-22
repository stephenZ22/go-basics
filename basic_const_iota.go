package main

import "fmt"

const (
	k = 3 // 在此处，iota == 0

	m float32 = iota + .5 // m float32 = 1 + .5
	n                     // n float32 = 2 + .5

	p    = 9          // 在此处，iota == 3
	q    = iota * 2   // q = 4 * 2
	_                 // _ = 5 * 2
	r                 // r = 6 * 2
	s, t = iota, iota // s, t = 7, 7
	u, v              // u, v = 8, 8
	_, w              // _, w = 9, 9
)

const x = iota // x = 0 （iota == 0）
const (
	y = iota // y = 0 （iota == 0）
	z        // z = 1
)

// 实际业务中常用定义
const (
	Failed    = iota - 1 // == -1
	Unknown              // == 0
	Succeeded            // == 1
)

func main() {
	fmt.Println("m =>", m)
	fmt.Println("n =>", n)
	fmt.Println("q =>", q)
	fmt.Println("r =>", r)
	fmt.Println("s =>", t)
	fmt.Println("u =>", u)
	fmt.Println("v =>", v)
	fmt.Println("w =>", w)
	fmt.Println("Failed =>", Failed)
	fmt.Println("Unknown =>", Unknown)
	fmt.Println("Succeeded =>", Succeeded)
}
