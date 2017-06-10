package main

import "fmt"

func euclid(m, n int) int {
  var r int
  if m < n {
    r = m
    m = n
    n = r
  }
  r = m % n
  if r == 0 {
    return n
  } else {
    m = n
    n = r
    return euclid(m, n)
  }
}

func main() {
  fmt.Println(euclid(2166, 6099))
}
