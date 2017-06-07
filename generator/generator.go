package main

import "fmt"

func makeGenerator() func() int {
  i := int(0)
  return func() (ret int) {
    ret = i
    i += 2
    return
  }
}
func main() {
  var list [10]int
  newEven := makeGenerator()
  for value := range list[:] {
    fmt.Println(value, newEven())
  }
}
