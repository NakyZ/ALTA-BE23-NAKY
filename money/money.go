package main

import "fmt"

func moneychange(money int)[] int { // code here
	a := []int{100, 20, 1}
    change := []int{}
    for _, denom := range a {
        count := money / denom
        money -= count * denom
        for i := 0; i < count; i++ {
            change = append(change, denom)
        }
    }
   return change 
}

func main()  { // input data
	input := 123
	change := moneychange(input)
	fmt.Println(change)
}  