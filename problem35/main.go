package main

import (
	"fmt"
	"strconv"
)

func makeData(n int) *[]int {
	data := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		data[i] = i + 2
	}
	return &data
}

func filterPrimes(xs *[]int) {
	for i := 0; i < len(*xs); i++ {
		number := (*xs)[i] // take number
		if number == 0 {
			continue
		}
		for j := 1; j*number < len(*xs)-i; j++ {
			(*xs)[i+j*number] = 0
		}
	}
}

func deflate(xs *[]int) *[]int {
	primes := make([]int, 0)
	for _, val := range *xs {
		if val != 0 {
			primes = append(primes, val)
		}
	}
	return &primes
}

func primesTo(n int) *PrimesList {
	data := makeData(n)
	filterPrimes(data)
	ps := deflate(data)
	return (*PrimesList)(ps)
}
// ---------------------------
type PrimesList []int

func (ps *PrimesList) contains(x int) bool {
	for _, val := range *ps {
		if val == x {
			return true
		}
	}
	return false
}

func (ps *PrimesList) println() {
	fmt.Println(*ps)
}

func (ps *PrimesList) isCircular(n int) bool {
	for _, val := range rotations(n) {
		if !ps.contains(val) {
			return false
		}
	}
	return true
}

func (ps *PrimesList) circular() []int {
	var cs []int
	for _, val := range *ps {
		if ps.isCircular(val) {
			cs = append(cs, val)
		}
	}
	return cs
}
// ---------------------------
func rotations(n int) []int {
	s := strconv.Itoa(n)
	//fmt.Println("s = ", s)
	var res []int
	for _ = range s {
		s1 := fmt.Sprintf("%v%c", s[1:], s[0])
		i1, _ := strconv.Atoi(s1)
		res = append(res, i1)
		s = s1
	}
	return res
}

func main() {
	//ps := primesTo(7)
	//ps.println()
	//fmt.Println(ps.contains(1))
	//fmt.Println(ps.contains(2))
	//fmt.Println(rotations(197))
	ps := primesTo(1000000)
	//ps.println()
	cs := ps.circular()
	//fmt.Println(cs)
	fmt.Println("len", len(cs))
}
