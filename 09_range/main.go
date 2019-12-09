package main

import "fmt"

func main() {
	ids := []int{33, 76, 39, 45, 2, 44}

	// Loop
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// if not using index i
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	greetings := map[string]string{"Hello": "World", "Goodbye": "Land"}

	for k, v := range greetings {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range greetings {
		fmt.Println("Key: " + k)
	}
}
