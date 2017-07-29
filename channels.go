package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello go routine awake")
	done <- true
}
func main() {
	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
}
