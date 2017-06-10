package main

import "fmt"

func fibRecursive(n int) int {
  // (starts from 1)
  if n <= 1 {
    return 1
  }
  return fibRecursive(n - 2) + fibRecursive(n - 1)
}

func main() {
  fmt.Println(fibRecursive(4))
  // starts bottlenecking right about here
  // fmt.Println(fibRecursive(40))
}
