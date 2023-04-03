package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func threeAZ() {
	defer wg.Done()

	for count := 0; count < 3; count++ {

		for c := 'A'; c < 'A'+26; c++ {
			fmt.Printf("%c ", c)
		}
	}
}

func threeaz() {
	defer wg.Done()

	for count := 0; count < 3; count++ {

		for c := 'a'; c < 'a'+26; c++ {
			fmt.Printf("%c ", c)
		}
	}
}
func main() {
	wg.Add(2)

	go threeAZ()
	go threeaz()

	wg.Wait()
	fmt.Println("\nMain done")
}
