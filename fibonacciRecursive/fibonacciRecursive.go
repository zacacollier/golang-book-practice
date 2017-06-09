package main

import "fmt"

func fibRecursive(n int) int {
  if n <= 2 {
    return 1
  }
  return fibRecursive(n - 2) + fibRecursive(n - 1)
}

func main() {
  fmt.Println(fibRecursive(5))
  fmt.Println(fibRecursive(15))
  fmt.Println(fibRecursive(25))
  fmt.Println(fibRecursive(35))
  // starts bottlenecking right about here
  fmt.Println(fibRecursive(40))
}
