package main

import "fmt"

func primes(xs *[]int) {
	for i := 0; i < len(*xs); i++ {
		number := (*xs)[i] // take number
		if number == 0 {
			continue
		}
		for j := 1; j * number < len(*xs) - i; j++ {
			(*xs)[i + j * number] = 0
		}
	}
}

func makeData(n int) *[]int{
	data := make([]int, n - 1)
	for i := 0; i < n - 1; i++ {
		data[i] = i + 2
	}
	return &data
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

func main() {
	data := makeData(200000)

	fmt.Println(data)
	primes(data)
	fmt.Println(data)
	ps := deflate(data)
	fmt.Println(ps)
	fmt.Println("len", len(*ps))
	fmt.Println("res", (*ps)[10001 - 1])
	fmt.Println("end")
}
