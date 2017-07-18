package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  m := make(map[string]int)

  for _, value := range strings.Fields(s) {
    elem, ok := m[value]
    if ok == true {
      m[value] = elem + 1
      continue
    }

    m[value] = 1
  }
  return m
}

func main() {
  wc.Test(WordCount)
}
