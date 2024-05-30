package main

import "fmt"

// Code Here
func ubahHuruf(i string) string {
	var output string
	for _, char := range i {
		if char >= ('A') && char <= ('Z') {
			if char == ('A') {
				output += ("K")
			} else if char == ('B') {
				output += ("L")
			} else if char == ('C') {
				output += ("M")
			} else if char == ('D') {
				output += ("N")
			} else if char == ('E') {
				output += ("O")
			} else if char == ('F') {
				output += ("P")
			} else if char == ('G') {
				output += ("Q")
			} else if char == ('H') {
				output += ("R")
			} else if char == ('I') {
				output += ("S")
			} else if char == ('J') {
				output += ("T")
			} else if char == ('K') {
				output += ("U")
			} else if char == ('L' ){
				output += ("V")
			} else if char == ('M') {
				output += ("W")
			} else if char == ('N') {
				output += ("X")
			} else if char == ('O' ){
				output += ("Y")
			} else if char == ('P' ){
				output += ("Z")
			} else if char == ('Q') {
				output += ("A")
			} else if char == ('R' ){
				output += ("B")
			} else if char == ('S') {
				output += ("C")
			} else if char == ('T' ){
				output += ("D")
			} else if char == ('U') {
				output += ("E")
			} else if char == ('V') {
				output += ("F")
			} else if char == ('W' ){
				output += ("G")
			} else if char == ('X') {
				output += ("H")
			} else if char == ('Y' ){
				output += ("I")
			} else if char == ('Z' ){
				output += ("J")
			}
		} else if char == (' ') {
			output += (" ")
		} else {
			output += string(char)
		}
	}
	return output
}

func main() {
	i := ("SEPULSA OKE	") // Input Program
	fmt.Println("Huruf Input = ", i)
	fmt.Println("Huruf Output = ", ubahHuruf(i))
}
