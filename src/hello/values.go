package main

import "fmt"

func main() {
	fmt.Println("go is fun")
	fmt.Println("1 + 2 = ", 1+2) // Good Old C Syntax
	fmt.Println(true || false)

	var var1 = "Some Random Text' " // Like Js and Dart
	fmt.Println(var1)

	var v1, v2 int = 1, 2
	fmt.Println("Declare Multiple variables at once ", v1, ", ", v2)

	var d = true
	fmt.Println("Go Infers type of Variables : ", d)

	var e int
	fmt.Println("int type is Initialized with zero", e)

	f := "Shortand :="
	fmt.Println(f, "for declaring and initializing a variable")

	const test = 255
	fmt.Println("const decare a constant value ", test)

	fmt.Println("Numeric constants have no type until its given one, e.g by explicit converion ", 3e10/5, int64(3e10/5))
}
