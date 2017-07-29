package main

import "fmt"

func fibonacci() func() int {
	previous, current := 0, 1

	return func() int {
		result := previous
		previous, current = current, previous+current
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
