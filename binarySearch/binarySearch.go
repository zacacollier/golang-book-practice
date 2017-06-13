package main

import "fmt"

func Round(val float64) int {
  return int(val)
}

func binarySearch(seq []int, i int) int {
  low := 0
  high := len(seq) - 1
  for low <= high {
    mid := low + Round(float64(high - low)) / 2
    // If `i` is above the current midpoint,
    // shift the starting point `low` forward by 1.
    if (seq[mid] < i) {
      low = mid + 1
      fmt.Println("low:", low, "mid:", mid, "high:", high)
      fmt.Println("seq:", seq)
    } else {
      // If `i` is below the current midpoint,
      // shift the endpoint `high` backward by 1
      high = mid - 1
    }
  }
  // If `low` is accumulated to len(seq),
  // the `if` condition on line 11 will
  // be out of range (runtime error)
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
