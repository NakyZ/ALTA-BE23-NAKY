package main

import "fmt"

func binarysearch(array []int, x int) int { // code here
  nilaibawah := 0
  nilaiatas := len(array) - 1

  for nilaibawah <= nilaiatas {
    mid := (nilaibawah + nilaiatas) / 2

    if array[mid] > x {
		nilaiatas = mid - 1
    } else if array[mid] < x {
		nilaibawah = mid + 1
    } else {
      return mid
    }
  }

  return -1
}

func main() { // input data
  arr := []int{1, 1, 3, 5, 5, 6, 7}
  x := 3

  index := binarysearch(arr, x)

  if index == -1 {
    fmt.Println("Hasil Tidak ditemukan.")
  } else {
    fmt.Println("Hasilnya adalah	=  ",index)
 }
}
