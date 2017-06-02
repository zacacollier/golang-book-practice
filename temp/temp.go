package main

import (
  "fmt"
)

func main() {
  fmt.Print("What's the temperature in Fahrenheit? ")
  var input float64
  fmt.Scanf("%f", &input)

  output := int((input - 32) * 5 / 9)

  fmt.Println("That is", output, "degrees Celsius u American dumbass")
}
