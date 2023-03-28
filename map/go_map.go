package main

import "fmt"

// 遍历指定类型的键值对
func PrintfMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}

func main() {
	// 使用make内置函数声明map
	var liming = make(map[string]string)
	liming["name"] = "李明"
	liming["age"] = "18"
	liming["addrs"] = "北京"

	// 使用短变量声明并初始化map
	lilei := map[string]string{
		"name":  "李雷",
		"age":   "19",
		"addrs": "广州",
	}
	// 下面这行赋值编译错误，键值类型需和声明是保持一致
	// lilei["weight"] = 19
	PrintfMap(lilei)

	// 获取map键值长度
	fmt.Printf("李明 len:%d\n", len(liming))

	PrintfMap(liming)
	fmt.Printf("李明 len:%d\n", len(liming))
	fmt.Println("========================================")

	// 修改map键值元素
	liming["addrs"] = "上海"
	PrintfMap(liming)
	fmt.Printf("李明 len:%d\n", len(liming))
	fmt.Println("========================================")

	// 删除键
	delete(liming, "addrs")
	PrintfMap(liming)
	fmt.Printf("李明 len:%d\n", len(liming))

}
