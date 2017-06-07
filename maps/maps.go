package main

import "fmt"

func main() {
  // make(map[key type]value type)
  elements := map[string]string {
  "H": "Hydrogen",
  "He": "Helium",
  "Li": "Lithium",
  "Be": "Beryllium",
  "B": "Boron",
  "C": "Carbon",
  "N": "Nitrogen",
  "O": "Oxygen",
  "F": "Flourine",
  "Ne": "Neon",
}

  if name, ok := elements["Un"]; ok {
    fmt.Println(name,ok)
  }
}
