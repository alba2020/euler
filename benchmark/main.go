package main

import (
	"fmt"
	"time"
)

func calc(x int) int {
	return x*x + 2*x + 1 + 3*x + 123 - x*12 + x
}

func f1(xs *[]int) {
	limit := len(*xs)
	for i := 0; i < limit; i++ {
		(*xs)[i] = calc(i)
		//fmt.Println(*xs)
	}
}

func sum(xs *[]int) (s int) {
	for _, val := range *xs {
		s += val
	}
	return
}


func f2(xs *[]int) {
	input := make(chan int)
	output := make(chan int)
	control := make(chan int)

	go func(input, output, control chan int) {
		for {
			select {
			case x := <- input:
				output <- calc(x)
			case z := <- control:
				if z == 0 {
					return
				}
			}
		}
	}(input, output, control)

	limit := len(*xs)
	for i := 0; i < limit; i++ {
		input <- i
		(*xs)[i] = <- output
	}
	control <- 0
}

func main() {
	xs := make([]int, 1000000)

	start := time.Now()
	f1(&xs)
	elapsed := time.Since(start)
	fmt.Printf("time: %s\n", elapsed)
	//fmt.Println(xs)
	fmt.Println(sum(&xs))

	xs1 := make([]int, 1000000)

	start1 := time.Now()
	f2(&xs1)
	elapsed1 := time.Since(start1)
	fmt.Printf("time: %s\n", elapsed1)

	fmt.Println(sum(&xs1))
}
