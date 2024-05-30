package main

import "fmt"

// code here
func FindMaxSumSubArray(k int, arr []int) int {
	sizesum, sum := k, 0
	maxSum, start := 0, 0
	for i, num := range arr {
		sum += num
	if i >= sizesum {
			if sum > maxSum {
				maxSum = sum
			}
			sum = arr[start]
			start ++
		}
	}
	return maxSum
}
// input program
func main() {
fmt.Println("==============================")
	arr := []int{2,1,5,1,3,2} // input angka
	k := 3 // input nilai K
	fmt.Println("Find Max Sum Sub Array	= ", k)
	fmt.Println("Output			= ", FindMaxSumSubArray(k,arr))
fmt.Println("==============================")
}
