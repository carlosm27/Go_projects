package main

import (
  "fmt"
)

func main () {
  a := make([]int, 1, 2)
  a[0] = 1

  notCopy1 := a[:]
  fmt.Println(notCopy1)
  notCopy1 = append(a, 1)
  fmt.Println(notCopy1)

  //notCopy2 := a[:]
  //notCopy2 = append(a, 2)
  //fmt.Println(notCopy1) // same as notCopy2
  //fmt.Println(notCopy2)


}
