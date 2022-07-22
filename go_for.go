package main

import "fmt"

func for01() {
	fmt.Println("========= 省略初始化语句=========")

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func for02() {
	fmt.Println("========= 省略步尾语句=========")

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func for03() {
	// break 关键字用于提前跳出循化，不执行步尾语句
	fmt.Println("========= 全部省略=========")
	i := 0
	for {
		if i >= 10 {
			break
		}

		fmt.Println(i)
		i++
	}
}

func for04() {
	// continue 关键字提前结束当前循环，仍执行步尾语句
	fmt.Println("========= continue关键字=========")
	for n := 0; n < 10; n++ {
		if n%2 != 0 {
			continue
		}

		fmt.Println(n)
	}
}

func main() {

	// 依次输出数字 0-9
	// for n := 0; n < 10; n++ {
	// 	fmt.Println(n)
	// }

	// 上述代码块为for循环语句的基本实现
	// 其中for为关键字,
	// n：=0 为初始化语句;
	// n < 10 为条件表达式;
	// n++为步尾语句;
	// 每个for循环中，初始化语句将率先执行，并且只会执行一次
	// 当每次循环开始时，将对条件表达式进行判断，若为false则结束循环
	// 当每次循环结束时，将执行步尾语句.

	// 以上提到的三个部分均可被省略

	//省略初始化语句s
	for01()

	// 省略步尾语句
	for02()

	// 全部省略
	for03()

	// continue 关键字
	for04()

}
