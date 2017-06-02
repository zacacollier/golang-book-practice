package main

import (
  "fmt"
  "strconv"
)

func main() {
  fmt.Print("How many miles? ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 0.3048
  strconv.FormatFloat(output, 'f', 2, 64)
  fmt.Println("That's", output, "meters, you imperialist swine.")
}
