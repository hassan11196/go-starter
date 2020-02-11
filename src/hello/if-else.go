package main

import (
	"fmt"
)

func main() {
	num := 5
	if num%2 == 1 {
		fmt.Println("num ", num, " is odd")
	} else {
		fmt.Println("num ", num, " is even")
	}
	fmt.Println("Sadly Go has no Ternary Operator :'( ")

	if t1 := 5; t1 < 10 {
		fmt.Println("statements can precede conditionals and they will be available in all brnaches")
	}

}
