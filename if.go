package main

import "fmt"

var a int = 10

func compare(x int) string{
  if x > 10 {
    fmt.Println("x 大于 10")
  } else if x == 10 {
    fmt.Println("x 等于 10")
  } else {
    fmt.Println("x 小于10")
  }
  return string(x)
}

func main (){

  fmt.Println(compare(5))
}
