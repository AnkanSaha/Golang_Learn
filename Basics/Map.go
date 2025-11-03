package basics

import "fmt"

func Maps () {
	// simple map 
	m1 := map[string]int{"name": 1}

	fmt.Println(m1["name"], len(m1))


	// map with make
	m2 := make(map[string]int)

	
	// Append Data
	m2["Test"] = 2
	// fmt.Println(m2["Test"], m2)

	value, ok := m2["Tests"]
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}