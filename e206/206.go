package main

import (
	"fmt"
	"os"
)

// Find the unique positive integer whose square has the form 1_2_3_4_5_6_7_8_9_0,
// where each “_” is a single digit.

func chop(p *uint64) uint64 {
	rem := *p % 10
	*p -= rem
	*p /= 10
	return rem
}

func main() {
	var i uint64 = 1
	outer: for i = 123_456_789; i < 10_000_000_000; i++ {
		res := i * i
		if chop(&res) == 0 {
			var j uint64
			for j = 9; j > 0; j-- {
				_ = chop(&res)
				if chop(&res) != j {
					continue outer
				}
			}
			fmt.Printf("%d * %d = %d\n", i, i, i*i)
			os.Exit(0)
		}
	}
	fmt.Println("not found")
}
