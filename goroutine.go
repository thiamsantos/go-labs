package main

import (
	"fmt"
	"time"
)

func numbers(start time.Time) {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nnumber took %s\n", time.Since(start))
}

func alphabet(start time.Time) {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	fmt.Printf("\nalphabet took %s\n", time.Since(start))
}

func main() {
	start := time.Now()
	go numbers(start)
	go alphabet(start)
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	fmt.Printf("main took %s", time.Since(start))
}
