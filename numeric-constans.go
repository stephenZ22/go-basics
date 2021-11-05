package main

import "fmt"

const (
  Big = 1 << 100
  Small = Big >> 99
)

func needInt(x int) (y int){
  y = x *10 + 1
  return
}

func needFloat(x float64)(y float64){
  y = x * 0.1
  return
}


func main(){
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))

}
