package basics;


func ExampleRange() {
	numbers := []int{1, 2, 3, 4, 5}

	// Simple range loop
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}


	// Range loop with ignored index
	for _, value := range numbers {
		println("Value:", value)
	}

	// Range with Map
	personAge := map[string]int{"Alice": 30, "Bob": 25}
	for name, age := range personAge {
		println("Name:", name, "Age:", age)
	}
}