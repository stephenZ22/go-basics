package main

import "fmt"

func sum(arr [5]int) (sum int) {
	sum = 0
	for i, n := range arr {
		fmt.Println(i, n)
		sum += n
	}

	return
}

func main() {
	var names = [5]string{"张学友", "刘德华", "黎明", "郭富城", "成龙"}
	names1 := [5]string{"张学友", "刘德华", "黎明", "郭富城", "成龙"}
	// 由编译器推断数组长度
	names2 := [...]string{"张学友", "刘德华", "黎明", "郭富城", "成龙"}

	names3 := [...]string{}
	// 下面这行赋值编译报错
	// names3[0] = "李修贤"

	fmt.Printf("names: %T, %v\n", names, names)
	fmt.Printf("names1: %T, %v\n", names1, names1)
	fmt.Printf("names2: %T, %v\n", names2, names2)
	fmt.Printf("names2: %T, %v\n", names3, names3)

}
