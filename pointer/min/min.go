package main

import "fmt"

func getminmax(numbers *int) (int, int) { // code 
	 min := *numbers
	 max := *numbers

for i := 0; i < 6; i++ {
		if (*numbers+i) < min {
		min = (*numbers+i)
	}
if (*numbers+i) > max {
max = (*numbers+i) 
}
	}
	return min, max 
}

func main() { // input number
var numbers [6]int

fmt.Println("Masukkan 6 angka:")
for i := 0; i < 6; i++ {
 fmt.Scan(&numbers[i]) 
}

min, max := getminmax(&numbers[0]) 

fmt.Println(max, "Nomer maksimal") 
 fmt.Println(min, "Nomer Minimal") 
}