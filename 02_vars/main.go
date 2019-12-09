package main

import "fmt"

func main() {
	// using var
	var age = 36
	const isCool = true

	//shorthand method - only in function
	name := "Grace"
	size := 1.3

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", size)
}
