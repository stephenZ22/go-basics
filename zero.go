package main

import "fmt"

func main(){
  var (
    i int
    f float64
    v bool
    s string
  )

  fmt.Printf("%v %v %v %q\n", i, f, v, s)
  fmt.Printf("%T %T %T %T\n", i, f, v, s)

}
