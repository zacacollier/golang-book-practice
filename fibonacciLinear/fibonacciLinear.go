package main

import "fmt"

func fib(n int) int {
  var n2, n1, next int
  n2 = 1
  n1 = 0
  for n > 1 {
    next = n2
    n2 = n2 + n1
    n1 = next
    fmt.Println(n2)
    n--
  }
  return n2
}

func main() {
  // JavaScript tops out at 9007199254740991 (just returns "Infinity")
  fmt.Println(fib(10000))
}
