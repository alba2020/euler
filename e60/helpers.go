package main

import (
	"fmt"
	"math"
	"strconv"
)

func conc1(a, b int) int {
	s := fmt.Sprintf("%d%d", a, b)
	n, _ := strconv.ParseInt(s, 10, 32)
	return int(n)
}

// concatenate two integers
func conc(a, b int) int {
	// size1 := int(math.Log10(float64(a))) + 1
	size2 := int(math.Log10(float64(b))) + 1
	return a*int(math.Pow10(size2)) + b
}

func check(xs ...int) bool {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			num1 := conc(xs[i], xs[j])
			if !primes[num1] {
				return false
			}
			num2 := conc(xs[j], xs[i])
			if !primes[num2] {
				return false
			}
		}
	}
	return true
}

func check2(i, j int) bool {
	if !primes[conc(i, j)] {
		return false
	}
	if !primes[conc(j, i)] {
		return false
	}
	return true
}
