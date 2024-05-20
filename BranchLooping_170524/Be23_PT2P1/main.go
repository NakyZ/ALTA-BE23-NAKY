package main

import "fmt"

func main() {
	//input
    var Studentscore int
    fmt.Print("Masukkan Score anda ( dengan Angka ) = ")
    fmt.Scan(&Studentscore)

	//kode di sini
    if Studentscore >100 { 
		fmt.Println("Nilai anda tidak ada")
	}else if Studentscore >0{
        fmt.Println("Nilai Anda = D")
	}else if Studentscore >34{
        fmt.Println("Nilai Anda = C")
	}else if Studentscore >49{
        fmt.Println("Nilai Anda = B")
	}else if Studentscore >64{
        fmt.Println("Nilai Anda = B+")
	}else if Studentscore >79{
        fmt.Println("Nilai Anda = A")		
	}else if Studentscore <0{
        fmt.Println("Nilai anda tidak ada")		
	}
}
