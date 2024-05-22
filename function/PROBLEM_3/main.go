package main

import "fmt"

// Code Here
func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%2d ", i*j)
		}
		fmt.Println("")
	}
}

func main() {
	cetakTablePerkalian(7) // Input Program
}
