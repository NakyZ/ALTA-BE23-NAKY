package main

import "fmt"

func main() {
    // input
    var alas float64
	var tinggi float64

    // kode disini
    fmt.Print("Masukan angka alas segitiga = ")
    fmt.Scan(&alas)

    fmt.Print("Masukan angka tinggi segitiga = ")
    fmt.Scan(&tinggi)

    luas := 0.5*alas*tinggi

    fmt.Println("luas segitiga =",luas)
}