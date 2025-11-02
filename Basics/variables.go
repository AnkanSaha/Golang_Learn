package basics

import "fmt"

func Variables() {
	// Variable declaration
	var name string = "John"
	var age int = 30
	var height float64 = 5.9

	// Print the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)

	// Short variable declaration
	country := "USA"
	fmt.Println("Country:", country)

	// Array declaration
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Slice declaration
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Fruits:", fruits)

	// Map declaration
	person := map[string]string{
		"firstName": "Jane",
		"lastName":  "Doe",
	}
	fmt.Println("Person:", person["firstName"])
}