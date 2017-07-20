package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func abs(num float64) float64 {
  if num < 0 {
    return num * -1
  }
  return num
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }

  var start float64 = 1.0
  var count int = 0

  for count < 1000 {
    count += 1
    var approximation float64 = abs(start - ((start*start - x) / (2 * start)))
    if abs(start-approximation) < 1e-15 {
      break
    }
    start = approximation
  }

  return start, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
