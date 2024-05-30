package main

import (
	"fmt"
	"strings"
)

//input
func Palindrome(s string) bool {
	s = strings.ToLower(s) 
	for a := 0; a < len(s)/2; a++ {
		if s[a] != s[len(s)-a-1] {
			return false
		}
	}
	return true
}

//kode disini
func main() {
	var input string
	fmt.Print("Ketik Kalimat yang mau di cek = ")
	fmt.Scanln(&input)

	if Palindrome(input) {
		fmt.Println(input, "Termasuk Palindrome")
	} else {
		fmt.Println(input, "Tidak Termasuk Palindrome")
	}
}
