package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key and values
	emails["Bob"] = "bob@email.com"
	emails["Sharon"] = "sharon@email.com"
	emails["Grace"] = "grace@email.com"

	// Declare map and add kv
	greetings := map[string]string{"Hello": "World", "Goodbye": "Land"}

	fmt.Println(emails, greetings)

	delete(emails, "Bob")
	fmt.Println(emails)
}
