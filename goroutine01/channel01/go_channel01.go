package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// c := make(chan int)
	// c2 := make(chan int)
	// wg.Add(3)

	// go bbb(c)
	// go aaa(c, c2)

	// for y := range c2 {
	// 	fmt.Println("Y:", y)
	// }
	c3 := make(chan int)
	wg.Add(1)

	go func(c3 chan int) {
		defer wg.Done()

		x := <-c3
		fmt.Println(x)
	}(c3)
	c3 <- 1

	wg.Wait()
	fmt.Println("Main done")
}

func aaa(c <-chan int, c2 chan<- int) {
	defer wg.Done()
	defer close(c2)

	for x := range c {
		fmt.Println("计算c平方")
		c2 <- x * x
	}
}

func bbb(c chan<- int) {
	defer wg.Done()
	defer close(c)

	for x := 0; x < 10; x++ {
		fmt.Println("add chan c", x)

		c <- x
	}
}
