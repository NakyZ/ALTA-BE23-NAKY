package main

import "fmt"

func swap(a *int, b *int) { // code
temp := *a
	*a = *b 
 	*b = temp
}

func main() { // input number
 var a = 10
 var b = 20

 fmt.Println("--------------------")
 fmt.Println("Input	: a=", a, "b=", b) 
 swap(&a, &b) 

 fmt.Println("Output	:  a=", a, "b=", b) 
 fmt.Print("--------------------")
}
