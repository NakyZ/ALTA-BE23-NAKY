package main

import "fmt"

func compare(a, b string) string {
// Code Here
	nilaiA := len(a)
	nilaiB := len(b)

	start := 0
	end := 0
	
	for i := 0; i < nilaiA; i++ {
		for j := 0; j < nilaiB; j++ {
			length := 0
			for i2 := i; i2 < nilaiA && j < nilaiB && a[i2] == b[j]; i2, j = i2+1, j+1 {
				length++ 
			}
			if length > start {
				start = length
				end = i
			}
		}
	}
	return a[end : end+start]
}

// Input Program
func main() {
fmt.Println("====================")
	fmt.Println(compare("AKA", "AKASHI"), "\t") // Jika Input "AKASHI" maka Outputnya adalah "AKA"
	fmt.Println(compare("KANG", "KANGOORO"), "\t") // Jika Input "KANGOORO" maka Outputnya adalah "KANG"
fmt.Println("====================")
}
