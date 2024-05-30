package main

import "fmt"

//input
func primenumber(number int)bool  {
	if number <= 1{
		return false
	}
	for a := 2; a*a <= number; a++ {
		if number%a ==0 {
			return false
		}
	}
	return true
}	

//kode disini
func main()  {
	var bilangan int 
	fmt.Print("Masukan bilangan = ")
	fmt.Scan(&bilangan)

	if primenumber(bilangan){
		fmt.Println(bilangan, "Berikut adalah bilangan Prima")
	} else {
		fmt.Println(bilangan, "Berikut adalah bilangan bukan Prima")
	}
}
