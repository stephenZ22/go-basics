package main

import "fmt"

// 匿名函数作为参数，实现回调函数
func sumNumbers(numbers []int, callback func(int)) {
	i := 0
	for _, n := range numbers {
		i += n
	}

	callback(i)
}

func main() {
	// 小括号代表匿名函数立即执行
	func() {
		fmt.Println("我是一个无需参数无返回值的匿名函数")
	}()

	// 可以将匿名函数赋值给函数变量
	doubleNumber := func(n int) int {
		fmt.Println("doubleNumber function")
		return n * 2
	}

	n := doubleNumber(2)
	fmt.Println(n)

	sumNumbers([]int{1, 2, 3, 4, 5}, func(i int) {
		if i > 5 {
			fmt.Println("总和大于五")
		} else if i == 5 {
			fmt.Println("总和等于五")
		} else {
			fmt.Println("总和小于五")
		}
	})
}
