package main

import "fmt"

func simpleequations(a, b, c int)int {
	if a == 0 {
		if b == 0 {
			fmt.Println("Persamaan tidak memiliki solusi.")
			return -1
		} else {
			x := -c / b
			fmt.Println("Solusi persamaan adalah x =", x)
			return x
		}
	} else {
		x := (-b * b + 4 * a * c) / (2 * a)
		fmt.Println("Solusi persamaan adalah x =", x)
		return x
	}
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scanln(&c)
}
