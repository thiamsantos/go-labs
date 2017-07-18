package main

import (
  "fmt"
  "math"
)

func abs(num float64) float64 {
  if (num < 0) {
    return num * -1
  }
  return num
}

func Sqrt(x float64) float64 {
  var start float64 = 1.0
  var count int = 0

  for count < 1000 {
    count += 1
    var approximation float64 = abs(start - ((start * start - x)/ (2*start)))
    if abs(start - approximation) < 1e-15 {
      break
    }
    start = approximation
  }

  fmt.Println("loops:", count)
  return start
}

func main() {
  fmt.Println("native:", math.Sqrt(2))
  fmt.Println("custom:", Sqrt(2))
}
