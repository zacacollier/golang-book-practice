package main

import (
  // "bytes"
  // "encoding/json"
  "fmt"
  // "reflect"
  "regexp"
)
func compileRegExp(reg string) {
  return
}

func findCSSClass(reg string) {
  ex, e := regexp.Compile(reg)
  fmt.Println(ex, e)
  // match, _ := regexp.MatchString(`\t*\"(color)\"(\:\s)\"(#[\w]*)`, reg)
  return
}


func main() {
  compileRegExp(`\t*\"(color)\"(\:\s)\"(#[\w]*)`)
  findCSSClass(`\t*\"(color)\"(\:\s)\"(#[\w]*)`)
}
