package main

import "fmt"

// func swap(a, b *int) {
// 	tmp := *a
// 	*a = *b
// 	*b = tmp
// }

// var errParams = errors.New("参数错误")

// func area(l, w int) (int, error) {
// 	if l <= 0 || w <= 0 {
// 		return 0, errParams
// 	}

// 	return l * w, nil
// }

func addNumber(a int, b int) int {
	fmt.Println("---Function addNumber()----")
	return a + b
}

func addNumber2(a int, b int) (c int) {
	fmt.Println("---Function addNumber2()----")
	c = a + b
	return
}

func addOne(a, b, c int) (int, int, int) {
	fmt.Println("---Function addOne()----")
	return a + 1, b + 1, c + 1
}

func addTwo(a *int) {
	fmt.Println("---Function addTwo()----")
	t := *a
	*a = t + 2
}

func main() {
	a := 1
	b := 2
	c := 3
	d := 4

	x := addNumber(a, b)
	fmt.Println(x)

	y := addNumber2(a, b)
	fmt.Println(y)

	z, v, n := addOne(a, b, c)
	fmt.Println(z, v, n)

	addTwo(&d)
	fmt.Println("Current d:", d)
}

// func main() {
// 	a, b := 1, 2
// 	fmt.Println(a, b)

// 	swap(&a, &b)
// 	fmt.Println(a, b)

// 	x, y := 3, 6
// 	area, err := area(x, y)

// 	if err != nil {
// 		fmt.Println("异常")
// 	}

// 	fmt.Println("面积为:", area)
// }
