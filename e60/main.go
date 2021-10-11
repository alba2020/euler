package main

import (
	"fmt"
	"time"
)

var primes []bool

const limit = 9999

func setup() {
	start := time.Now()
	primes = createPrimes(99_999_999)
	fmt.Println("primes created", time.Since(start).Seconds())
}
func main() {
	setup()

	for i := 1; i < limit; i++ {
		if !primes[i] {
			continue
		}
		for j := i + 2; j < limit; j += 2 {
			if !primes[j] {
				continue
			}
			if !check2(j, i) {
				continue
			}
			for k := j + 2; k < limit; k += 2 {
				if !primes[k] {
					continue
				}
				if !check2(k, j) {
					continue
				}
				if !check2(k, i) {
					continue
				}
				for m := k + 2; m < limit; m += 2 {
					if !primes[m] {
						continue
					}
					if !check2(m, k) {
						continue
					}
					if !check2(m, j) {
						continue
					}
					if !check2(m, i) {
						continue
					}
					for n := m + 2; n < limit; n += 2 {
						if !primes[n] {
							continue
						}
						if !check2(n, m) {
							continue
						}
						if !check2(n, k) {
							continue
						}
						if !check2(n, j) {
							continue
						}
						if !check2(n, i) {
							continue
						}
						fmt.Println(i, j, k, m, n, i+j+k+m+n)
					}
				}
			}
		}
	}
}
