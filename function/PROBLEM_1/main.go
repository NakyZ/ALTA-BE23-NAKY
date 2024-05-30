package main

import "fmt"

func playwirhasterisk(n int) string {
	//Code Here
	var result string
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			result += (" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			result += ("*")
		}
		result += ("\n")
	}
	return result
}
func main() {
	fmt.Println(playwirhasterisk(5)) //input Program
}
