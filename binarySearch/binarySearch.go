package main

import "fmt"

func Round(val float64) int {
  if val < 0 {
    return int(val + 0.5)
  }
  return int(val)
}

func binarySearch(seq []int, i int) int {
  seq = seq[:]
  mid := Round(float64(len(seq))) / 2
  last := len(seq) - 1
  fmt.Println(mid, last, seq[0:mid], seq[mid:])
  switch {
  case seq[0] == i:
    return 0
  case seq[mid] == i:
    return mid
  case seq[last] == i:
    return last
  case seq[mid] < i:
    return binarySearch(seq[mid:], i)
  case seq[mid] > i:
    return binarySearch(seq[0:mid], i)
  default:
    return -1
  }
}

func main() {
  fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7))
}
