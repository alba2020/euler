package main

import (
	"fmt"
	"sort"
	"time"
)

type Number struct {
	a int64
	p int64
	q int64
	r int64
}

func newNumber(p, q, r int64) Number {
	return Number{
		a: p * q * r,
		p: p,
		q: q,
		r: r,
	}
}

func (n Number) isValid() bool {
	return n.q*n.r+n.p*n.r+n.p*n.q == 1
}

const limit = 30

func main() {
	start := time.Now()

	found := make([]Number, 1)
	var p, q, r int64
	for p = 1; p < limit; p++ {
		for q = -1; q > -limit; q-- {
			for r = q; r > -limit; r-- {
				n := newNumber(p, q, r)
				if n.isValid() {
					found = append(found, n)
				}
			}
		}
	}
	sort.Slice(found, func(i, j int) bool {
		return found[i].a < found[j].a
	})

	for i, f := range found {
		// fmt.Println("limit", limit)
		// if i == 101 {
		fmt.Printf("%d - %d\n", i, f)
		// }
	}
	fmt.Println("total", len(found))
	fmt.Println("elapsed", time.Since(start).Seconds())
}
