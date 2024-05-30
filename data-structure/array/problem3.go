package main

import "fmt"

// code here
func arrayunique(arrayA, arrayB []int)[] int {
	result  := make([]int, 0)
	for _, num := range arrayA {
		found := false
		for _, num2 := range arrayB{
			if num == num2 {
				found = true 
				break
			}
		}
		if !found {
			result = append(result, num)
		}
	}
	return result
}

// input program
func main()  {
fmt.Println("====================")
	arrayA := []int {1,2,3,4} // input data A
	arrayB := []int {1,3,5,10,16} // input data B
	fmt.Println(arrayunique(arrayA, arrayB))
fmt.Println("====================")
}