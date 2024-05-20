package main

import "fmt"

//input
func exponentation(base int, pangkat int)int  {
	result := 1
	for a := 0; a < pangkat; a++ {
		result *= base
	}
	return result
}

//kode disini
func main()  {
	var base, pangkat int
	fmt.Print("Masukan Basis = ")
	fmt.Scanln(&base)
	fmt.Println("Masukan pangkat = ")
	fmt.Scanln(&pangkat)

	result := exponentation (base, pangkat)
	fmt.Printf("hasilnya adalah %d\n", result )
}