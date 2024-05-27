package main

import "fmt"

// code here
func removeduplicate(arr []int) int {
    if len(arr) == 0 {
        return 0
    }
    i := 0
    for j := 1; j < len(arr); j++ {
        if arr[j] != arr[i] {
            i++
            arr[i] = arr[j]
        }
    }
    return i + 1
}
// input program
func main() {
fmt.Println("===================================")
    arr := []int{1, 2, 3, 3, 3, 6, 9} // input angka yang akan di gunakan
    length := removeduplicate(arr)
    fmt.Println("Input		= ", length)
    fmt.Println("Angka		= ", arr[:length])
fmt.Println("===================================")
}
