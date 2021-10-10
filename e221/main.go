package main

import (
	"fmt"
	"math"
	"sort"
)

type Number struct {
	a int
	p int
	q int
	r int
}

func newNumber(p, q, r int) Number {
	return Number{
		a: p * q * r,
		p: p,
		q: q,
		r: r,
	}
}

func (n Number) isValid() bool {
	left := 1 / float64(n.a)
	right := 1/float64(n.p) + 1/float64(n.q) + 1/float64(n.r)
	if math.Abs(left-right) < epsilon {
		return true
	}
	return false
}

const limit = 900
const epsilon = 0.000000001

func main() {
	found := make([]Number, 1)
	foundx := make([]int, 1)

	for p := 1; p < limit; p++ {
		for q := -1; q > -limit; q-- {
			for r := q; r > -limit; r-- {
				n := newNumber(p, q, r)
				if n.isValid() {
					found = append(found, n)
				}
				a := p * q * r
				left := 1 / float64(a)
				right := 1/float64(p) + 1/float64(q) + 1/float64(r)
				if math.Abs(left-right) < epsilon {
					foundx = append(foundx, a)
					// found = append(found, newNumber(p, q, r))
				}
			}
		}
	}
	sort.Slice(found, func(i, j int) bool {
		return found[i].a < found[j].a
	})
	sort.Ints(foundx)

	for i, f := range foundx {
		if i == 150 {
			fmt.Printf("%d - %d\n", i, f)
			fmt.Println(found[i])
		}
	}
	fmt.Println("total", len(found), len(foundx))
}
