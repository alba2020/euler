package main

import "math"

func setTrue(xs []bool) {
	for i := range xs {
		xs[i] = true
	}
}

func setPrimes(xs []bool) {
	// setTrue(xs)
	// xs[0] = false
	// xs[1] = false
	var n int = len(xs) - 1
	for k := 2; k < n; k += 2 {
		// xs[k] = false -- default for evens
		xs[k+1] = true // initial value for odds
	}
	limit := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 3; i <= limit; i += 2 {
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
