package main

import "fmt"

// Code Here
func DRAWXYZ(n int) string {
	var result string 
	for i := 1; i <= n*n ; i++ {
		if i%3 == 0 {
			result += ("X")
		} else if i%2 == 0 {
			result += ("Z")
		} else {
			result += ("Y")
		}
		if i%n == 0 {
			result += ("\n")
		}
	}
	return result
}

func main() {
	fmt.Println(DRAWXYZ(4)) // input program
}
