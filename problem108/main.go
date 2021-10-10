package main

import (
	"fmt"
	"time"
)

const epsilon float64 = 0.0000000000001

func solve(n int) (nsolutions int) {
	first_x := n + 1
	last_x := first_x
	//last_x := int(1/(1/float64(n)-1/float64(first_x))) * 2
	for 1/float64(n)-1/float64(first_x) < 1/float64(last_x) {
		last_x *= 2
	}

	//fmt.Printf("from %d to %d\n", first_x, last_x)
	//fmt.Println("n = ", n)
	//fmt.Printf("old: %d last: %d\n", last_x_old, last_x)

	var y int

	for x := first_x; x < last_x; x++ {

		//k := float64(1)/float64(n) - float64(1)/float64(x)
		//fmt.Println(n, x, k)
		n64 := float64(n)
		x64 := float64(x)

		y_float := 1 / (1/n64 - 1/x64)
		//y_float := x64 * n64  / (x64 - n64)
		diff := y_float - float64(int(y_float))
		//diff := y_float - math.Trunc(y_float)
		//fmt.Printf("x = %v, diff = %v\n", x, diff)
		if diff < epsilon {
			y = int(y_float)
			//fmt.Println("solution", x, y)
			if x == y {
				nsolutions += 2
				//fmt.Println("+2")
			} else {
				nsolutions += 1
				//fmt.Println("+1")
			}
		} else if (1 - diff) < epsilon {
			y = int(y_float) + 1
			//fmt.Println("solution", x, y)
			if x == y {
				nsolutions += 2
				//fmt.Println("+2")
			} else {
				nsolutions += 1
				//fmt.Println("+1")
			}
		}
	}
	//fmt.Println("total ns", nsolutions)
	return nsolutions / 2
}

func solve2(n int) (nsolutions int) {
	//first_x := n + 1
	//last_x := first_x
	//for 1/float64(n)-1/float64(first_x) < 1/float64(last_x) {
	//	last_x *= 2
	//}

	var y int
	x := n + 1

	//for x := first_x; x < last_x; x++ {
	for {
		if x*n%(x-n) == 0 {
			y = x * n / (x - n)
			if x > y {
				break
			}
			nsolutions++
			//fmt.Println(x, y)
			//if x == y {
			//	nsolutions += 2
			//} else {
			//	nsolutions += 1
			//}
		}
		x++
	}

	return
}

func solve3(n int) (nsolutions int) {
	var y int
	x := n + 1

	for {
		if x*n%(x-n) == 0 {
			y = x * n / (x - n)
			if x > y {
				break
			}
			fmt.Println("x", x, "y", y, "gcd x y", gcd(x, y), "x-n", x-n)
			nsolutions++
		}
		x++
	}
	return
}

func gcd(a, b int) int {
	for b > 0 {
		a %= b
		a, b = b, a
	}
	return a
}

func main() {
	start := time.Now()

	for i := 4; i <= 15; i++ {
		n := solve3(i)
		if n > 1 {
			fmt.Printf("i = %d solutions = %d\n-----\n", i, n)
		}
	}

	//fmt.Println("gcd 15 5", gcd(21, 14))

	elapsed1 := time.Since(start)
	fmt.Printf("time: %s\n", elapsed1)
}
