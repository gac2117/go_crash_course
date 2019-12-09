package main

import "fmt"

func main() {
	// Arrays must be fixed length and name the types
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	dogArr := [2]string{"German Shepard", "Poodle"}

	// Slices
	veggieSlice := []string{"Tomato", "Carrot"}

	fmt.Println(fruitArr)
	fmt.Println(dogArr)
	fmt.Println(len(veggieSlice))
}
