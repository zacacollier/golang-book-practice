package main

import (
  "fmt"
)

func main() {
  for i := 1; i <= 100; i++ {
    switch i % 3 {
    case 0:
        if i % 5 == 0 {
          fmt.Println("FizzBuzz")
        } else {
          fmt.Println("Fizz")
        }
    default:
      fmt.Println(i)
    }
  }
}
