package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

type letterRange struct {
  Start byte
  End   byte
}

func (lr letterRange) InRange(letter byte) bool {
  return lr.Start <= letter && letter <= lr.End
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
  firstLowerCaseRange := letterRange{'a', 'm'}
  firstUpperCaseRange := letterRange{'A', 'M'}

  secondLowerCaseRange := letterRange{'n', 'z'}
  secondUpperCaseRange := letterRange{'N', 'Z'}

  length, err := rot.r.Read(b)

  for i := 0; i < length; i++ {
    if firstLowerCaseRange.InRange(b[i]) || firstUpperCaseRange.InRange(b[i]) {
      b[i] += 13
    } else if secondLowerCaseRange.InRange(b[i]) || secondUpperCaseRange.InRange(b[i]) {
      b[i] -= 13
    }
  }

  return length, err
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
