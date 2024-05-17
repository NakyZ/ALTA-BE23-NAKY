package main

import "fmt"

func main() {
   // input
	var r float64
	var t float64

	// Kode disini
    fmt.Print(" Masukan angka Radius Tabung = ")
    fmt.Scan(&r)

    fmt.Print(" Masukan angka Tinggi Tabung = ")
    fmt.Scan(&t)

    Luas := 2*3.14*r*(r+t)

    fmt.Println("Luas Tabung = ",Luas )
}