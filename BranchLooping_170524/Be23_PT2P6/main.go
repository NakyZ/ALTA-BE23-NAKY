package main

import (
	"fmt"
	"strconv"
	"strings"
)

//input
func fullprime(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
	return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	i := 5
	for i*i <= number {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func isFullPrime(n int) bool {
	str := strconv.Itoa(n)
	digits := strings.Split(str, "")
	for _, digit := range digits {
		digitInt, err := strconv.Atoi(digit)
		if err != nil {
			return false
		}
		if !fullprime(digitInt) {
			return false
		}
	}
	return fullprime(n)
}

//kode disini
func main() {
	var input int
	fmt.Print("Masukkan angka =  ")
	fmt.Scan(&input)
	if isFullPrime(input) {
		fmt.Println(input, "True")
	} else {
		fmt.Println(input, "False")
	}
}
