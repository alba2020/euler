package main

import "fmt"

func diff(n int) int {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
		sum += i
	}
	return sum*sum - sumOfSquares
}

func main() {
	fmt.Println(diff(10))
	fmt.Println(diff(100))
}
