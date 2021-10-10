package main

import (
	"math"
)

const epsilon = 0.0000001

func validFloat(p, q, r int) bool {
	a := p * q * r
	left := 1 / float64(a)
	right := 1/float64(p) + 1/float64(q) + 1/float64(r)
	if math.Abs(left-right) < epsilon {
		return true
	}
	return false
}

func validInt(p, q, r int) bool {
	return q*r+p*r+p*q == 1
}
