package basics

import "fmt"

func FuncWithGenetics[T int | string](status []T) int {
	count := 0
	for index, value := range status {
		fmt.Println("Genetics Index", index, "Genetics Value", value)
		count += 1
	}

	return count
}



func UseGenetics (){
	var newSlice = []string{"efewfw"}
	fmt.Println("Total value Genetics", FuncWithGenetics(newSlice))
}