package main

import "fmt"

// 定义结构体
// 这里需要注意！！！
// 下面Person 结构体及结构体字段首字母均为大写
// 在Golang中不同于Ruby，命名规范上推荐使用用驼峰命名
// 若一个包中的 函数，结构体，等定义是时 首字母为小写.
// 则该函数、结构体只能在当前包内使用，若为大写则可在其他包Import时在其他包内使用
// 结构体中的字段可以用成员变量来理解，类似于Ruby中类的属性

type Person struct {
	Name    string
	Age     int
	Addrs   string
	StepNum int
}

// 方法
// Golang中将函数和方法做了区分(也肯能是我自己一直没有区分，不过问题不大)
// Golang 的方法 需要一个接收者（receiver）这点和Ruby类似
// 只能为当前包内定义的类型 定义方法
// 同时方法的接收者 可以是一个值 也可以是一个指针
// 同时不支持方法重载

// 实现 Person 的实例方法

// 接收者类型为指针
func (p *Person) addStep(i int) {
	p.StepNum += i
}

// 接收者类型为Value
func (p Person) sayHello() {
	fmt.Printf("%v say: Hello !!!\n", p.Name)
}

func (p Person) afterTenYearAge() int {
	return p.Age + 10
}

func printPerson(p Person) {

	fmt.Printf("Name: %v, Age: %d, Addrs: %v, StepNum: %d\n", p.Name, p.Age, p.Addrs, p.StepNum)
}

func main() {
	// 实例化一个结构体
	liming := Person{Name: "李明", Age: 18, Addrs: "北京", StepNum: 180}

	// 若不对属性进行赋值则会使用对应类型的0值
	lilei := Person{Name: "李磊", Age: 17}

	printPerson(liming)
	printPerson(lilei)

	// 更新结构体实例
	lilei.StepNum = 170
	liming.Age = 20
	fmt.Println("\n===========================\n")
	printPerson(liming)
	printPerson(lilei)

	fmt.Println("\n===========================\n")
	// 调用指针类型
	// 这里没用使用& 进行寻址，但是方法调用用成功
	// 因为编译器做了自动转换
	liming.addStep(20)

	// 调用值类型
	liming.sayHello()
	printPerson(liming)

	lm := &liming
	lm.addStep(25)
	printPerson(liming)

	afterAge := liming.afterTenYearAge()
	fmt.Println(afterAge)

}
