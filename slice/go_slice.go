package main

import "fmt"

func sum(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return
}

func printSlice(s []int) {
	fmt.Printf("Type: %T, Value: %v, len: %d, cap: %d\n",

		s, s, len(s), cap(s)*2,
	)
}

func main() {

	// 初始化一个切片
	as := []int{1, 2, 3, 4}
	fmt.Printf("%p\n", &as)
	// 使用append函数对切片进行追加
	as = append(as, 5)
	fmt.Printf("%p\n", &as)

	// 使用make函数初始切片
	// 每个方法接受三个参数（类型，长度，容量）
	// 其中容量参数可以省略，但长度及容量参数同时存在，长度值需小于等于容量值
	// 容量为切片底层数组的长度
	bs := make([]int, 4, 5)
	cs := []int{6, 7, 8, 9, 10}

	bs[0] = 1
	bs[1] = 2
	bs[2] = 3
	bs[3] = 4
	bs = append(bs, 5)
	bs = append(bs, cs...)

	printSlice(as)
	printSlice(bs)
	printSlice(cs)

	ds := make([]int, len(bs), cap(bs)*2)
	// 使用copy函数对切片进行复制
	copy(ds, bs)
	printSlice(ds)

	fmt.Printf("Sum as: %d\n", as)
	fmt.Printf("Sum bs: %d\n", bs)
	fmt.Printf("Sum cs: %d\n", cs)
	fmt.Printf("Sum ds: %d\n", ds)

	// 切片截取
	ds2 := ds[:3]
	printSlice(ds2)

	ds3 := ds[3:]
	printSlice(ds3)

	ds4 := ds[3:7]
	printSlice(ds4)

	ds5 := ds[:]
	printSlice(ds5)
}
