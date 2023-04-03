package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func GoRou() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			time.Sleep(1 * time.Second)
			fmt.Println("current", i)
		}(i)

	}

}

func FakeGoRou() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		func(i int) {
			defer wg.Done()

			time.Sleep(1 * time.Second)
			fmt.Println("current", i)
		}(i)

	}

}
func main() {
	startTime := time.Now()
	GoRou()
	// FakeGoRou()

	wg.Wait()
	spendTime := time.Since(startTime)
	fmt.Println("Main Done use", spendTime)
}
