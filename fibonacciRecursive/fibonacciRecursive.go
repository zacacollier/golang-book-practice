package main

import "fmt"

func fibRecursive(n int) int {
  if n <= 2 {
    return 1
  }
  return fibRecursive(n - 2) + fibRecursive(n - 1)
}

func main() {
  fmt.Println(fibRecursive(8))
}
