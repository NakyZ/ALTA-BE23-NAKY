package main

import "fmt"


func main() {
	//input
	var bilangan int
	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&bilangan)

	//kode disini
	fmt.Println("Faktor bilangan", bilangan, "adalah = ")
	for a := 1; a <= bilangan; a++ {
		if bilangan%a == 0 {
			fmt.Println(a)
		}
	}
}