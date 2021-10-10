package main

import (
	"strconv"
	"fmt"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Palindrome(n int) bool {
	s1 := strconv.Itoa(n)
	return s1 == Reverse(s1)
}

func main() {
	//s1 := "hello"
	//s2 := "olleh"
	//fmt.Println(s1, s2, s1 == Reverse(s2))

	max := 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			n := i * j
			if Palindrome(n) && n > max {
				max = n
			}
		}
	}

	fmt.Println("max", max)

	//fmt.Println(Palindrome(121))
}

