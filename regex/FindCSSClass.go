package main

import (
  // "bytes"
  // "encoding/json"
  "fmt"
  "io/ioutil"
  // "os"
  // "github.com/alexflint/go-restructure"
  // "reflect"
  "regexp"
)
type DecodeJSON struct {
  // Start
  _    struct{}   `^`
  Class string    `\t*\"([\w\s]*)\"(\:)`
  Color string    `\t*\"(color)\"`
  _    struct{}   `$`
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func openJson(path string) {
  f, err := ioutil.ReadFile(path)
  check(err)
  fmt.Println(string(f))
  return
}
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
  // compileRegExp(`\t*\"(color)\"(\:\s)\"(#[\w]*)`)
  // findCSSClass(`\t*\"(color)\"(\:\s)\"(#[\w]*)`)
  openJson("./small.json")
}
