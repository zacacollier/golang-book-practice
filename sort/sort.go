package main

import "fmt"

func sweep(numbers []int) {
  N := len(numbers)
  i1 := 0
  i2 := 1

  for i2 < N {
    n1 := numbers[i1]
    n2 := numbers[i2]
    fmt.Println("comparing:", n1, n2)
    if (n1 > n2) {
      numbers[i1] = n2
      numbers[i2] = n1
    }
    i1++
    i2++
  }
}


func main() {
  var numbers = []int{5, 4, 3, 2, 1, 0}
  fmt.Println(numbers)
  for i := range numbers {
    sweep(numbers)
    i++
  }
  fmt.Println(numbers)
}
