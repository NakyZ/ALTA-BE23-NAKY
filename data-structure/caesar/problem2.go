package main

import "fmt"

// code here
func caesar(offset string, input int) string {
	ciphertext := ""
	for _, char := range offset {
		if char == ' ' {
			ciphertext += " "
		} else {
			newCharCode := int(char) + input 
			if newCharCode > 'z' {
				newCharCode = (newCharCode-'a')%26 + 'a'
			}
			ciphertext += string(rune(newCharCode))
		}
	}
	return ciphertext
}
// input program
func main() {
	TextAwal := "Naky" // input nama
	input := 3 // input data huruf yang akan di offset

	TextAkhir := caesar(TextAwal, input)
fmt.Println("====================")
	fmt.Println("Text Awal	=   ", TextAwal)
	fmt.Println("Text Akhir	=   ", TextAkhir)
fmt.Println("====================")
}
