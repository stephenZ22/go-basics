package main

import "fmt"

func main(){
  sum := 1
  for ; sum <= 100; {
    fmt.Println("初始sum:", sum)
    sum += sum
    fmt.Println("完成sum:", sum)
  }
  fmt.Println(sum)
}
