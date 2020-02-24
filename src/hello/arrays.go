package main

import "fmt"

func main() {
	var v [5]int

	for i := 0; i < 5; i++ {
		fmt.Println(v[i])
		v[i] = 5*i + 2
		fmt.Println(v[i])
	}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
