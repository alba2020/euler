package main

import "fmt"

func main() {
	var a uint8 = 12
	var b uint8 = 2

	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", b)	
	fmt.Printf("%b \n", a & b)
}
