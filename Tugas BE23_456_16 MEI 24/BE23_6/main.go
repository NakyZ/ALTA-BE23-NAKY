package main

import "fmt"

func main() {
	// input
    var HargaBeli float64
	var Diskon float64

	// kode disini
    fmt.Print(" Masukan Harga Beli = ")
    fmt.Scan(&HargaBeli)

    fmt.Print("Masukan Diskon (10 Menjadi 10%) = ")
    fmt.Scan(&Diskon)

    TotalDiskon := HargaBeli * Diskon / 100
    HargaTertinggi := HargaBeli - TotalDiskon

    fmt.Println("Total harga yang didapat",HargaTertinggi)
}