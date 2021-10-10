package main

import (
	"fmt"
	"math"
)

func setTrue(xs []bool) {
	for i := range xs {
		xs[i] = true
	}
}

func setPrimes(xs []bool) {
	setTrue(xs)
	xs[0] = false
	xs[1] = false
	var n int = len(xs) - 1
	limit := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= limit; i++ {
		if xs[i] == true {
			for j := i * i; j <= n; j += i {
				xs[j] = false
			}
		}
	}
}

func createPrimes(max int) []bool {
	xs := make([]bool, max+1)
	setPrimes(xs)
	return xs
}

func main() {
	xs := createPrimes(100)
	for i, isPrime := range xs {
		if isPrime {
			fmt.Println(i)
		}
	}
}
