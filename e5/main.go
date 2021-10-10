package main

import "fmt"

var x10 = []int{7, 8, 9, 10}
var x20 = []int{7, 11, 13, 16, 17, 18, 19, 20}

func check(i int, xs []int) bool {
	for _, x := range xs {
		if i%x != 0 {
			return false
		}
	}
	return true
}

const limit = 1_000_000_000

func main() {
	for i := 2; i < limit; i++ {
		if check(i, x20) {
			fmt.Println(i)
			break
		}
	}
}
