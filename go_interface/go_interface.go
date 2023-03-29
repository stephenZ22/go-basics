// go_interface.go

package main

import (
	"fmt"
	"unsafe"
)

// Golang 中接口可以看作是一组方法的集合
// 当一个类型 T 实现了 I 接口的全部方法 我们可以认为
// T 类型实现了这个接口
// 这里暂时可以认为 T类型是 I 接口的子类（但要清楚 Golang 没有class 继承的概念）

// 定义一个 Person 接口
type Person interface {
	walk()
	take()
	speak()
}

// 定义一个Human 结构体
type Human struct {
	Name string
	Age  int
}

// 定义一个Adult 结构体 引入Human
type Adult struct {
	Human
	WorkAddrs string
}

// Adult 以下三个方法 实现了 Person Interface 中定义的函数
func (a Adult) walk() {
	fmt.Println(a.Name, "Adult walk ")
}

func (a Adult) take() {
	fmt.Println(a.Name, "Adult take ")
}

func (a Adult) speak() {
	fmt.Println(a.Name, "Adult speak")
}

func (a Adult) work() {
	fmt.Println(a.Name, "Adult Work hard")
}

// 定义一个Child 结构体 引入Human
type Child struct {
	Human
	SchoolAddrs string
}

// Child 以下三个方法 实现了 Person Interface 中定义的函数
//
func (c Child) walk() {
	fmt.Println(c.Name, "Child walk ")
}

func (c Child) take() {
	fmt.Println(c.Name, "Child take ")
}

func (c Child) speak() {
	fmt.Println(c.Name, "Child speak")
}

func (c Child) study() {
	fmt.Println(c.Name, "Child study")
}

// 接受类型为 Person interface的参数
func personDo(p Person) {
	p.walk()
	p.take()
	p.speak()

	// 类型断言
	// 根据类型执行对应类型的不同方法
	if c, ok := p.(Child); ok {
		c.study()
	} else if c, ok := p.(Adult); ok {
		c.work()
	}

}

func main() {

	a := Adult{
		Human: Human{
			Name: "李建军",
			Age:  45,
		},
		WorkAddrs: "西二旗",
	}
	fmt.Println(unsafe.Pointer(&a.Human.Age))
	fmt.Println(unsafe.Pointer(&a.Age))

	personDo(a)

	c := Child{
		Human: Human{
			Name: "李子涵",
			Age:  9,
		},
		SchoolAddrs: "人大附小",
	}

	personDo(c)
}
