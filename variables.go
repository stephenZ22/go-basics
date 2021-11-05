package main

import "fmt"

var c, ruby, python bool

func main(){
  var i int
  fmt.Println(i, c, ruby, python)

  c = true
  ruby = true

  fmt.Println(c, ruby, python)

  c := false

  fmt.Println(c)
}
