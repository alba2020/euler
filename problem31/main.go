package main

import "fmt"

func c(sum int, vals []int, history []int) int {
	if sum == 0 {
		//fmt.Println("ok", history)
		return 1
	}
	if sum < 0 {
		return 0
	}

	if len(vals) == 0 {
		return 0
	}

	return c(sum-vals[0], vals, append(history, vals[0])) + c(sum, vals[1:], history)
}

func main() {
	n := c(200, []int{1, 2, 5, 10, 20, 50, 100, 200}, []int{})
	fmt.Println("n", n)
}
