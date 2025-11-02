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
}