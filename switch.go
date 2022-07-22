package main

import (
  "fmt"
  "runtime"
)

func main(){
  fmt.Print("Go runs on ")
  os := runtime.GOOS
  switch os{
  case "darwinxx":
    fmt.Println("OS x.")
  case "linux":
    fmt.Println("Linux")
  default:
    fmt.Printf("%s.\n", os)
  }

}
