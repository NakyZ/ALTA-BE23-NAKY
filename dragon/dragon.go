package main

import "fmt"

func dragonloowater(dragonhead []int, knightheighy []int) int { // code here
	sum := -15
	for _, val := range dragonhead {
		sum += val 
	}

	for _, val := range knightheighy {
		sum += val 
	}
	return sum
}

func main() { // input data
	dragonhead := [2]int{7, 2}
	knightheighy := [4]int{2, 1, 8, 5 }
	result := dragonloowater(dragonhead[:], knightheighy[:])
	fmt.Println("Hasil	=   ", result)
}