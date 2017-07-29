package main

import "fmt"

func digits(num int, digitsChannel chan int) {
	for num != 0 {
		digit := num % 10
		digitsChannel <- digit
		num /= 10
	}
	close(digitsChannel)
}

func calculateByDigit(num int, fn func(int) int) int {
	digitsChannel := make(chan int)
	go digits(num, digitsChannel)
	sum := 0

	for digit := range digitsChannel {
		sum += fn(digit)
	}

	return sum
}

func square(num int) int {
	return num * num
}

func cube(num int) int {
	return num * num * num
}

func calculateSquare(num int, squareResult chan<- int) {
	squareResult <- calculateByDigit(num, square)
}

func calculateCube(num int, cubeResult chan<- int) {
	cubeResult <- calculateByDigit(num, cube)
}

func main() {
	cubeResult := make(chan int)
	squareResult := make(chan int)

	num := 589

	go calculateSquare(num, squareResult)
	go calculateCube(num, cubeResult)

	result := <-cubeResult + <-squareResult

	fmt.Println("result:", result)
}
