package main

import "fmt"

func factors(n int) []int {
	i := 2
	result := make([]int, 0, 16)
	for {
		if n % i == 0 {
			result = append(result, i)
			n /= i
			i = 2
		} else {
			if n == 1 {
				break
			}
			i++
		}
	}
	return result
}

func main() {
	xs := factors(600851475143)
	fmt.Println(xs)
}
