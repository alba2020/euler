package main

import (
	"strconv"
	"fmt"
	"strings"
	"os"
	"log"
	"bufio"
	"time"
)

func createNumbers(from, to int) []string{
	numbers := make([]string, to - from + 1)
	for i := 0; i <= to - from; i++ {
		numbers[i] = strconv.Itoa(from + i)
	}
	return numbers
}

func check(n string, numbers []string) (res []string){
	for _, number := range numbers {
		i0 := strings.Index(number, string(n[0]))
		i1 := strings.Index(number, string(n[1]))
		i2 := strings.Index(number, string(n[2]))

		if i0 < 0 || i1 < 0 || i2 < 0 {
			continue
		}

		if i0 < i1 && i1 < i2 {
			res = append(res, number)
		}
	}
	return
}

func readData(filename string) (res []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
		//fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	start := time.Now()

	pins := readData("problem79/pins.txt")
	//fmt.Println(pins)

	//os.Exit(0)

	numbers := createNumbers(64000000, 80000000)
	//numbers := []string{"123487", "53174", "31674"}
	//fmt.Println(numbers)

	for _, pin := range pins {
		fmt.Println(len(numbers))
		numbers = check(pin, numbers)
	}
	fmt.Println(numbers)

	elapsed := time.Since(start)
	fmt.Printf("time: %s", elapsed)
}
