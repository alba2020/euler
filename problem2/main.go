package main

import "fmt"

func main() {

	x := 1
	y := 2
	sum := 0

	for x < 4000000 {
		fmt.Printf("%d ", x)
		if x % 2 == 0 {
			sum += x
		}
		y, x = x + y, y
	}
	fmt.Println()
	fmt.Println("sum", sum)
}
