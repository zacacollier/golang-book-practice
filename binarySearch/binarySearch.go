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
  low := 0
  high := len(seq) - 1
  var mid int
 //  if (i > seq[high]) {
 //    return -1
 //  }
  for low <= high {
    mid = low + Round(float64(high - low)) / 2
    if (seq[mid] < i) {
      low = mid + 1
      fmt.Println("low:", low, "mid:", mid, "high:", high)
      fmt.Println("seq:", seq)
    } else {
      high = mid - 1
    }
  }
  if (low == len(seq) || seq[low] != i) {
    return -1
  } else {
    return low
  }
}

func main() {
  fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7))
  fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
