package main

import "fmt"


func student(text string, substitution map[rune]rune) string { // code
	result := []rune{}
	for _, char := range text {
		newChar, ok := substitution[char]
		if !ok {
			newChar = char
		}
		result = append(result, newChar)
	}
	return string(result)
}

func decrypt( text string, substitution map[rune]rune) string {
	invertedSub := make(map[rune]rune)
	for key, value := range substitution {
		invertedSub[value] = key
	}
	return student(text, invertedSub)
}

func main() {
	Nama  := "rizky"
	substitution := make(map[rune]rune)
	for i := 'a'; i <= 'z'; i++ {
		newChar := i + 3
		if newChar > 'z' {
			newChar = newChar - 'z' + 'a' - 1
		}
		substitution[i] = newChar
	}
	for i := 'A'; i <= 'Z'; i++ {
		substitution[i] = rune(i + 3)
	}

	ciphertext := student(Nama, substitution)
	fmt.Println("Nama				= ", Nama)
	fmt.Println("Hasil Pergeseran	= ", ciphertext)

}
