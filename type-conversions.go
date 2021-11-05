package main

import (
  "fmt"
  "math"
)


func main() {
  var x, y int = 3, 4

  var f float64 = math.Sqrt(float64(x*x + y*y))
  z := uint(f)
  var c = uint(f)

  fmt.Println(x, y, f)
  fmt.Printf("%T %T %T\n", f, z, c)
}
